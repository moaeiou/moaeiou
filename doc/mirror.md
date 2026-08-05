# 🚀 Tourist of Mirror by MoAEIOU
## 🗺 Location/Region?
We have some mirrors, Addr:

Location: Los Angeles, CA, U.S.
```
https://l.867678.xyz/
```
Location: Cloudflare R2
```
https://r2.867678.xyz/
```
## 💁 How to use?
Just need change addr from `Original` to `https://l.867678.xyz/` or any others

Such as you need to download `https://downloads.openwrt.org/snapshots/targets/x86/64/openwrt-sdk-x86-64_gcc-14.4.0_musl.Linux-x86_64.tar.zst`

But it only have `5MB/S`, That is so slowly.

So you can change address to `https://l.867678.xyz/openwrt-sdk-x86-64_gcc-14.4.0_musl.Linux-x86_64.tar.zst`.

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
But Action at:
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
Yes.

Need more fast?
```bash
ANV=$(curl -sS -I https://github.com/AnInsomniacy/aria2-next/releases/latest | grep -i '^location:' | grep -oE '[^/]+$' | tr -d '\r' | sed 's/^v//')
wget -O aria2-next https://github.com/AnInsomniacy/aria2-next/releases/download/v$ANV/aria2-next-$ANV-linux-x86_64
chmod +x ./aria2-next
aria2-next --max-connection-per-server=64 --split=64 --min-split-size=1M --continue=true https://l.867678.xyz/filename
```

## 🤔 Q&A
| Question | Anwser |
|  ----  | ----  |
| Time zone | Default UTC+0 |
| How long to sync | Every 4 hours from 00:00(UTC+0) |
| What be mirror | Dynamic version programs, binary, SDK. Turn to the URL to view all it mirror. |
## 🙏 Acknowledgements
Server: <https://colocrossing.com> <https://cloudflare.com>

Software: <https://nginx.org> <https://github.com/AnInsomniacy/aria2-next> <https://867678.xyz/project/fancyindex-theme>

Donate: <https://github.com/midqwerty-alt>
