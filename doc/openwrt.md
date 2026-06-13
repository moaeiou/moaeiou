# Compiling OpenWrt SDK.
## 🚀 You need to prepare a Linux environment.
> Demonstration system: Arch Linux amd64
```
sudo pacman -Syu --needed base-devel make git gawk ncurses openssl zlib zstd wget curl unzip python3 python-setuptools fillup rsync libxslt libxml2 boost findutils gcc g++ unzip patch
```
> OR any package manager you using.
## ⏬ Download SDK
> You have three choices
[All the mirrors of OpenWrt](https://openwrt.org/mirrors)
[All the mirrors of ImmortalWrt](https://downloads.immortalwrt.org/acknowledgements.html)
[Mirror by MoKanove](https://867678.xyz/doc/mirror)
```
curl -LO ⚠️[SDK URL]
# Not have curl? Let's try wget
wget ⚠️[SDK URL]
```
## 📦 Unzip SDK
```
tar -xvf ⚠️SDK_PKG_NAME
```
## 🛠 Init SDK
```
cd ⚠️SDK_FOLDER_NAME
make defconfig
./scripts/feeds update -a && ./scripts/feeds install -a
```
## 📝 Clone source code
> Please replace sdk-root with your actual SDK directory.
```
cd ⚠️sdk-root/package/
git clone ⚠️[URI]
```
## 🚀 Launch
> Make sure your at SDK root
```
make package/⚠️project_name/compile V=s -j$(nproc)
```
## ⚖️ License
> This Document is licensed under the [MoPL](https://867678.xyz/docs/license)
## 🎉 The End.
