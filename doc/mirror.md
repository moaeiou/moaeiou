# 🚀 A Tourist of Mirror by MoAEIOU
## 🤔 Q&A
| Question | Anwser |
|  ----  | ----  |
| Time zone | Default always UTC+0 |
| How long to sync | Every 4 hours from 00:00(UTC+0) |
| What be mirror | Dynamic version programs, binary, SDK. Turn to the URL to view all it mirror|
| Why using mirror | Sometime github speed maybe slowly, yeah. |
## 💁 How to use?
```
https://l.867678.xyz/filename
```
### Bypass CDN
> If you have no network restrictions, you can force l.867678.xyz to point to the following IP address.
```
23.95.28.150
```
> Example using wget bypass CDN
```
wget --no-check-certificate --header="Host: l.867678.xyz" https://23.95.28.150/filename
```
> Aslo can using motrix-next, it is very fast
>
> <https://github.com/AnInsomniacy/aria2-next> Give it a star!
```
ANV=$(curl -sS -I https://github.com/AnInsomniacy/aria2-next/releases/latest | grep -i '^location:' | grep -oE '[^/]+$' | tr -d '\r' | sed 's/^v//')
wget -O aria2-next https://github.com/AnInsomniacy/aria2-next/releases/download/v$ANV/aria2-next-$ANV-linux-x86_64
chmod +x ./aria2-next
./aria2-next --max-connection-per-server=64 --split=64 --min-split-size=1M --continue=true --check-certificate=false --header="Host: l.867678.xyz" https://23.95.28.150/filename
```
> Or force reffer IP
```
sudo -i
whoami # make sure you as root
echo "23.95.28.150 l.867678.xyz" >> /etc/hosts
```
