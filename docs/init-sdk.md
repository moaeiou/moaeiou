# A general tutorial for compiling OpenWrt SDK.
## 🚀 First, you need to prepare a Linux environment.
> Demonstration system: Arch Linux amd64
```
sudo pacman -Syu --needed base-devel make git gawk ncurses openssl zlib zstd wget curl unzip python3 python-setuptools fillup rsync libxslt libxml2 boost findutils gcc g++ unzip patch
```
> OR any pm you love.
## ⏬ Download SDK
> You have two choices
- ImmortalWrt SDK and OpenWrt SDK
[All the mirrors of OpenWrt](https://openwrt.org/mirrors)
[All the mirrors of ImmortalWrt](https://downloads.immortalwrt.org/acknowledgements.html)
```
curl -LO [SDK URL]
```
## 📦 Unzip SDK
```
tar -xvf $SDK_PKG_NAME
```
## 🛠 Init SDK
```
cd $SDK_FOLDER_NAME
make defconfig
./scripts/feeds update -a && ./scripts/feeds install -a
```
## 🎉 Done, SDK avaliable now.
