# ✋ Install OpenWrt Package
using ssh to connect your router
```
ssh root@IP # option: -p port
```
Then download package
```
cd /tmp # It is usually stored here. If you have limited memory, please find another location.
wget [URL of Package]
apk add --allow-untrusted ./pkg.apk # --allow-untrusted because software not included in software repositories generally lacks a digital signature and should be ignored.
```
# 🔧 Compiling OpenWrt SDK.
## 🚀 You need to prepare a Linux amd64 environment.
Demonstration system: Arch Linux amd64
```
sudo pacman -Syu --needed base-devel make git gawk ncurses openssl zlib zstd wget curl unzip python3 python-setuptools fillup rsync libxslt libxml2 boost findutils gcc g++ unzip patch
```
OR any package manager you using.
## ⏬ Download SDK
You have three choices
[All the mirrors of OpenWrt](https://openwrt.org/mirrors)
[All the mirrors of ImmortalWrt](https://downloads.immortalwrt.org/acknowledgements.html)
[Mirror by MoAEIOU](https://867678.xyz/doc/Mirror)
```bash
wget ⚠️[SDK URL]
# Havn't wget? Let's try curl
curl -LO ⚠️[SDK URL]
```
## 📦 Unzip SDK
```bash
tar -xvf ⚠️SDK_PKG_NAME
```
## 🛠 Init SDK
```bash
cd ⚠️SDK_FOLDER_NAME
make defconfig
./scripts/feeds update -a && ./scripts/feeds install -a
```
## 📝 Clone source code
Please replace sdk-root with your actual SDK directory.
```bash
cd ⚠️sdk-root/package/
git clone ⚠️[URL of project]
```
## 🚀 Launch
Make sure you at SDK `/`
```bash
make package/⚠️project_name/compile V=s -j$(nproc)
```
## ⚖️ License
This document was licensed under the [MoPL](https://867678.xyz/doc/MoPL)
