# 🚀 Tourist of Mirror by MoAEIOU

## 🗺 Location/Region?

We have some mirrors, Addr:

| Location | Address | Note |
|  ----  | ----  |
| Los Angeles, CA, U.S. | `https://l.867678.xyz/` | Tier 1 Mirror and not support IPV6 |
| Singapore | `https://s.867678.xyz/` | Tier 2 Mirror, Every 2 hours sync 1 time |
| Tokyo, Japan | `https://o.867678.xyz/` | Tier 2 Mirror, Every 2 hours sync 1 time |

## 💁 How to use?

Just need change addr from `Original` to `https://l.867678.xyz/` or any others

Such as you need to download `https://downloads.openwrt.org/snapshots/targets/x86/64/openwrt-sdk-x86-64_gcc-14.4.0_musl.Linux-x86_64.tar.zst`

But it only have `5MB/S`, That is so slowly.

So you can change address to `https://l.867678.xyz/sdk/openwrt-sdk-x86-64_gcc-14.4.0_musl.Linux-x86_64.tar.zst`.

Don't belive it fast? Don't worry!

GitHub Actions:
```bash
07/18 13:51:49 [info] Downloading 1 item(s)
[#2eb2f2 89MiB/282MiB(31%) CN:64 DL:94MiB ETA:2s]
[#2eb2f2 150MiB/282MiB(53%) CN:64 DL:78MiB ETA:1s]
[#2eb2f2 231MiB/282MiB(81%) CN:64 DL:79MiB]

07/18 13:51:53 [info] Download complete: /home/runner/work/moaeiou/moaeiou/openwrt-sdk-x86-64_gcc-14.4.0_musl.Linux-x86_64.tar.zst

Download Results:
gid   |stat|avg speed  |path/URI
======+====+===========+=======================================================
2eb2f2|OK  |    82MiB/s|/home/runner/work/moaeiou/moaeiou/openwrt-sdk-x86-64_gcc-14.4.0_musl.Linux-x86_64.tar.zst
```
But the runner at:
```bash
{
  "ip": "52.159.244.162",
  "city": "San Jose",
  "region": "California",
  "country": "US",
  "loc": "37.3394,-121.8950",
  "org": "AS8075 Microsoft Corporation",
  "postal": "95110",
  "timezone": "America/Los_Angeles",
  "readme": "https://ipinfo.io/missingauth"
}
```

Need more fast?
```bash
ANV=$(curl -sS -I https://github.com/AnInsomniacy/aria2-next/releases/latest | grep -i '^location:' | grep -oE '[^/]+$' | tr -d '\r' | sed 's/^v//')
wget -O aria2-next https://github.com/AnInsomniacy/aria2-next/releases/download/v$ANV/aria2-next-$ANV-linux-x86_64
chmod +x ./aria2-next
./aria2-next --max-connection-per-server=1024 --split=88 --min-split-size=1M --continue=true https://l.867678.xyz/filename
```

Rsync

GitHub Actions
```yaml
- name: Send to mirror
  uses: moaeiou/moaeiou/actions/rsync@main
  with:
    local: ${{ steps.sdk.outputs.sdk_dir }}/bin/packages/x86_64/base/
    remote: /var/www/html/luci-app/pm/
    key: ${{ secrets.CC_LA_KEY }}
    extra: --delete
```

Sync from this mirrors to your mirror

```bash
rsync -avz --exclude 'contento24/' rsync://l.867678.xyz/sync/ /path/to/your/mirror
```

Aslo can try Mirror level 2

## 🤔 Q&A

| Question | Anwser |
|  ----  | ----  |
| Time zone | Default UTC+0 |
| How long to sync | Every 4 hours from 00:00(UTC+0) |
| What be mirror | Dynamic version programs, binary, SDK. Turn to the URL to view all it mirror. |
| Are you support new features? | Yes, we supported QUIC, TLS1.3, X25519MLKEM768 and most new features. |
| Are you support rsync? | Yes, we testing it. Maybe after sometimes it will be release. |

## 🙏 Acknowledgements

Server、CDN、Domain: <https://colocrossing.com> <https://cloudflare.com> <https://spaceship.com>

Software: <https://nginx.org> <https://rsync.samba.org> <https://github.com/AnInsomniacy/aria2-next> <https://867678.xyz/project/fancyindex-theme>

Server Donate: <https://github.com/midqwerty-alt> <https://github.com/muchengrui110-source>
