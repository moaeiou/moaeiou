# A general tutorial for compiling OpenWrt SDK.
## 🚀 First, you need to prepare a Linux environment.
> Demonstration system: Arch Linux amd64
```
sudo pacman -Syu --needed base-devel make git gawk ncurses openssl zlib zstd wget curl unzip python3 python-setuptools fillup rsync libxslt libxml2 boost findutils gcc g++ unzip patch
```
> OR any pm you love.
## ⏬ Download SDK
[All the mirrors of ImmortalWrt](https://downloads.immortalwrt.org/acknowledgements.html)
```
curl -LO https://downloads.immortalwrt.org/releases/25.12.0-rc2/targets/x86/64/immortalwrt-sdk-25.12.0-rc2-x86-64_gcc-14.3.0_musl.Linux-x86_64.tar.zst
```
## 📦 Unzip SDK
```
tar -xvf immortalwrt-sdk-25.12.0-rc2-x86-64_gcc-14.3.0_musl.Linux-x86_64.tar.zst
```
## 🛠 Init SDK
```
cd immortalwrt-sdk-25.12.0-rc2-x86-64_gcc-14.3.0_musl.Linux-x86_64
make defconfig
./scripts/feeds update -a && ./scripts/feeds install -a
```
## 📝 Clone source code
```
cd package/
git clone [URI]
```
## 🚀 Launch
> Make sure your at SDK root
>
> pn is for project folder name
```
make package/pn/compile V=s -j$(nproc)
```
## 🎉 Done
