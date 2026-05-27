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
### ⚠️ Some projects require additional binaries; please refer to the How to build section of the project's README for details.
# 🎉 The end.
