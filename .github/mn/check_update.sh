#!/usr/bin/env bash
# Check the motrix-next GitHub repo for its newest release (stable OR
# pre-release, whichever was published most recently) and rewrite
# PKGBUILD's version fields to match if it differs from what's pinned.
#
# Writes CHANGED=true/false (and NEW_VERSION/NEW_TAG/DEB_ASSET_NAME when
# changed) to version_check_result.env for the workflow to source.
#
# Local testing without hitting the network / rate limits:
#   MOTRIX_RELEASE_FIXTURE=/path/to/releases.json ./check_update.sh

set -euo pipefail

REPO="AnInsomniacy/motrix-next"
PKGBUILD_PATH="PKGBUILD"
RESULT_FILE="version_check_result.env"

fetch_latest_release() {
    if [[ -n "${MOTRIX_RELEASE_FIXTURE:-}" ]]; then
        jq -c '.[0]' "$MOTRIX_RELEASE_FIXTURE"
        return
    fi

    # NOTE: /releases/latest only returns the newest *stable* release and
    # silently skips pre-releases. /releases (list) is sorted newest-first
    # and includes pre-releases, so we take element 0 from that instead.
    #
    # GITHUB_TOKEN (if set) raises the rate limit from 60/hr (anonymous)
    # to 5000/hr — the workflow always has one via secrets.GITHUB_TOKEN.
    local url="https://api.github.com/repos/${REPO}/releases?per_page=1"
    local -a curl_args=(-sf -H "Accept: application/vnd.github+json")
    if [[ -n "${GITHUB_TOKEN:-}" ]]; then
        curl_args+=(-H "Authorization: Bearer ${GITHUB_TOKEN}")
    fi

    local response
    if ! response=$(curl "${curl_args[@]}" "$url"); then
        echo "GitHub API request failed" >&2
        exit 1
    fi

    # GitHub returns an error object (e.g. rate limit) instead of a list.
    if jq -e 'type == "object" and has("message")' <<<"$response" >/dev/null 2>&1; then
        echo "GitHub API error: $(jq -r '.message' <<<"$response")" >&2
        exit 1
    fi

    jq -c '.[0]' <<<"$response"
}

main() {
    local release tag upstream_version sanitized_version deb_asset_name
    local current_gh_version

    release=$(fetch_latest_release)
    tag=$(jq -r '.tag_name' <<<"$release")               # e.g. "v3.9.8-beta.6"
    upstream_version="${tag#v}"                            # "3.9.8-beta.6"
    deb_asset_name=$(jq -r '.assets[] | select(.name | endswith("_amd64.deb")) | .name' <<<"$release" | head -n1)

    if [[ -z "$deb_asset_name" ]]; then
        echo "No amd64 .deb asset found on latest release" >&2
        exit 1
    fi

    # Arch pkgver cannot contain a hyphen; sanitize for the pkgver field
    # while keeping the raw upstream string for building the download URL.
    sanitized_version="${upstream_version//-/_}"

    current_gh_version=$(grep -oP '^_pkgver_gh=\K\S+' "$PKGBUILD_PATH" || true)

    if [[ "$current_gh_version" == "$upstream_version" ]]; then
        echo "Already up to date: $upstream_version"
        echo "CHANGED=false" > "$RESULT_FILE"
        return
    fi

    echo "New version found: '${current_gh_version}' -> '${upstream_version}'"

    sed -i "s/^_pkgver_gh=.*/_pkgver_gh=${upstream_version}/" "$PKGBUILD_PATH"
    sed -i "s/^pkgver=.*/pkgver=${sanitized_version}/" "$PKGBUILD_PATH"
    sed -i "s/^pkgrel=.*/pkgrel=1/" "$PKGBUILD_PATH"

    {
        echo "CHANGED=true"
        echo "NEW_VERSION=${sanitized_version}"
        echo "NEW_TAG=${tag}"
        echo "DEB_ASSET_NAME=${deb_asset_name}"
    } > "$RESULT_FILE"

    echo "PKGBUILD updated to pkgver=${sanitized_version} (_pkgver_gh=${upstream_version})"
}

main "$@"
