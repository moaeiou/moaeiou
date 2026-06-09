# This is a general tutorial for compiling ImmortalWrt projects.
- First, make sure you have initialized the SDK.
- If this step has not been completed, see [How to init SDK](https://867678.xyz/doc/init/)
## 📝 Clone source code
> Please replace sdk-root with your actual SDK directory.
```
cd sdk-root/package/
git clone [URI]
```
## 🚀 Launch
> Make sure your at SDK root
>
> pn is for project folder name
```
make package/pn/compile V=s -j$(nproc)
```
## ⚠️ WARN
> If you want to be lazy, in SDK outside and use:
```
make sdk-root/package/pn/compile V=s -j$(nproc)
```
> You will get an error message.
>
> This is because make needs to find the makefile, which it usually finds directly in the current directory.
### ⚠️ Some projects require additional binaries; please refer to the How to build section of the project's README for details.
## ⚖️ License
> This Document is licensed under the [MoPL](https://github.com/mokanove/mokanove/blob/main/docs/license.md)
# 🎉 The end.
