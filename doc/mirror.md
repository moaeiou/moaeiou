# 🚀 A Tourist of Mirror by MoKanove
## 🤔 Q&A
| Question | Anwser |
|  ----  | ----  |
| Time zone | Default always UTC+0 |
| How long to sync | RSYNC auto push |
| What be mirror | Dynamic version programs program, binary, SDK. |
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
> Or force reffer IP
```
sudo -i
whoami # make sure you as root
echo "23.95.28.150 l.867678.xyz" >> /etc/hosts
```
