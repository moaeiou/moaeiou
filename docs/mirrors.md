# 🚀 Tourist of Mirror by MoAEIOU

## 🤔 Q&A

| Question | Anwser |
|  ----  | ----  |
| Time zone | Default UTC+0 |
| Support new feature? | Yes, we supported QUIC, TLS1.3, X25519MLKEM768 and most new features. |
| Support rsync? | Yes, to dateils , please turn to [Here](#rsync). |

## 🗺 Location/Region?

We have some mirrors, Addr:

| Location | Address | Note |
|  ----  | ----  | ---- |
| Los Angeles, CA, U.S. | `https://la.mirrors.867678.xyz/` | Tier 1 but not support IPV6 |
| Singapore | `https://sg.mirrors.867678.xyz/` | Tier 2 , sync every hour |

## 💁 How to use?

Just need change addr from `Origin` to `https://la.mirrors.867678.xyz/` or any others

For example:

- You need to download `https://downloads.openwrt.org/snapshots/targets/x86/64/openwrt-sdk-x86-64_gcc-14.4.0_musl.Linux-x86_64.tar.zst`
- But it only have `5MB/S`, That is so slowly.
- So you can change address to `https://la.mirrors.867678.xyz/sdk/openwrt-sdk-x86-64_gcc-14.4.0_musl.Linux-x86_64.tar.zst`.

## rsync

### 🪣 GitHub Actions

```yaml
- name: Send to mirror
  uses: moaeiou/moaeiou/actions/rsync@main
  with:
    local: ${{ steps.sdk.outputs.sdk_dir }}/bin/packages/x86_64/base/
    remote: floder/
    # Default start from /var/www/html/
    key: ${{ secrets.SSH_KEY }}
    extra: --delete
    server: la.mirrors.867678.xyz
```

### 🧱 Sync to your mirror

Any mirror are all listening port `873`

And option: `--delete` that 1:1 sync any changes from upstream

```bash
rsync -avz --exclude='contento24/' --exclude='0fi/' --exclude='speedtest' --delete rsync://server/sync/ /var/www/html/
```

## 🙏 Acknowledgements

Server、CDN、Domain: <https://colocrossing.com> <https://cloudflare.com> <https://spaceship.com>

Software: <https://nginx.org> <https://github.com/RsyncProject/rsync> <https://github.com/AnInsomniacy/aria2-next> <https://867678.xyz/projects/fancyindex-theme>

Server Donate: <https://github.com/midqwerty-alt> <https://github.com/muchengrui110-source>

## ⚖️ LICENSE

Rsync licensed under the GPL-v3

And mirrors servers using Debian GNU/Linux lisenced under the GPL-v2

Mirrored files using their open source license, it just a file server.
