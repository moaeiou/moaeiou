#!/bin/bash
SECRET_VAL="神秘小Token"

ORGS=("contento24" "morouter" "orgmio")

MY_REPOS=("moaeiou/Mineradio" "moaeiou/fancyindex-theme")

for org in "${ORGS[@]}"; do
  gh secret set CC_LA_KEY -b"$SECRET_VAL" --org "$org" --visibility all
  echo "✅ 组织 $org 配置完成"
done

for repo in "${MY_REPOS[@]}"; do
  gh secret set CC_LA_KEY -b"$SECRET_VAL" -R "$repo"
  echo "✅ 个人仓库配置完成: $repo"
done
