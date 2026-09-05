# ✋ Install OpenWrt Package

using ssh to connect your router

```bash
ssh root@IP # option: -p port
```

Then download package

```bash
cd /tmp # It is usually stored here. If you have limited memory, please find another location.
wget [URL of Package]
apk add --allow-untrusted ./pkg.apk # --allow-untrusted because software not included in software repositories generally lacks a digital signature and should be ignored.
```

# 🔧 Compiling OpenWrt SDK.

## 🔩 You need a Linux amd64 environment.

Demonstration system: Arch Linux amd64

```
sudo pacman -Syu --needed base-devel make git gawk ncurses openssl zlib zstd wget curl unzip python3 python-setuptools fillup rsync libxslt libxml2 boost findutils gcc g++ unzip patch
```

OR any package manager you using.

## ⏬ Download SDK

[All the mirrors of OpenWrt](https://openwrt.org/mirrors)
[All the mirrors of ImmortalWrt](https://downloads.immortalwrt.org/acknowledgements.html)
[Mirror by MoAEIOU](https://867678.xyz/docs/mirrors)

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
./scripts/feeds update -a
./scripts/feeds install -a
```

## 📝 Clone source code

Please replace sdk-root with your actual SDK directory.

```bash
cd ⚠️sdk-root/package/
git clone ⚠️[URL of project]
```

## 🚀 Build

Make sure you at SDK `/`

```bash
make package/⚠️project_name/compile V=s -j$(nproc)
```

# ⚖️ License

There documents licensed under the [MoPL](https://867678.xyz/docs/mopl)
