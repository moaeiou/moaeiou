#!/usr/bin/env bash

set -euo pipefail

REPO="AnInsomniacy/motrix-next"
PKGBUILD_PATH="PKGBUILD"
RESULT_FILE="version_check_result.env"

fetch_latest_release() {
    if [[ -n "${MOTRIX_RELEASE_FIXTURE:-}" ]]; then
        jq -c '.[0]' "$MOTRIX_RELEASE_FIXTURE"
        return
    fi
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
