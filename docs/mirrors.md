# 🚀 Tourist of Mirror by MoAEIOU

## 🤔 Q&A

| Question             | Anwser                                                                |
| -------------------- | --------------------------------------------------------------------- |
| Time zone            | Default UTC+0                                                         |
| Support new feature? | Yes, we supported QUIC, TLS1.3, X25519MLKEM768 and most new features. |
| Support rsync?       | Yes, to dateils , please turn to [Here](#rsync).                      |

## 🗺 Location/Region?

We have some mirrors, Addr:

| Location              | Address                  | Note                        |
| --------------------- | ------------------------ | --------------------------- |
| Los Angeles, CA, U.S. | `https://la.867678.xyz/` | Tier 1 but not support IPV6 |
| Sydney, NSWales, AUS. | `https://au.867678.xyz/` | Tier 2 , sync every hour    |
| Osaka, Kansai, Japan  | `https://ki.867678.xyz/` | Tier 2 , sync every hour    |
| Singapore  | `https://sg.867678.xyz/` | Tier 2 , sync every hour    |

The `Los Angeles` node was so slowly, so i gave a CDN.

If you need rsync, please use `23.95.28.150`.

## 💁 How to use?

Just need change addr from `Origin` to `https://x.867678.xyz/` or any others.

## rsync

### 🧱 Sync to your mirror

Any mirror are all listening port `873`

And option: `--delete` that 1:1 sync any changes from upstream

```bash
rsync -avz --exclude='0fi/' --exclude='speedtest' --delete rsync://server/sync/ /var/www/html/
```

## 🙏 Acknowledgements

Server、CDN、Domain: <https://colocrossing.com> <https://cloudflare.com> <https://spaceship.com>

Software: <https://nginx.org> <https://github.com/RsyncProject/rsync> <https://github.com/AnInsomniacy/aria2-next> <https://867678.xyz/projects/fancyindex-theme>

## ⚖️ LICENSE

Rsync licensed under the GPL-v3

And mirrors servers using Debian GNU/Linux lisenced under the GPL-v2

Mirrored files using their open source license, it just a file server.
