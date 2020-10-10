#!/bin/sh
XREPO="ccr.ccs.tencentyun.com"
XPATH="open-platform-repo/ability"
XTAG=`date +%Y-%m-%d-%H-%M`

# 构建并编译
echo "[+] Compiling ${XREPO}/${XPATH}:build"
docker build --no-cache -t ${XREPO}/${XPATH}:build . -f Dockerfile.build

# 解压提取编译好的可执行文件和配置文件
echo "[+] Extracting docker image..."
docker create --name extract ${XREPO}/${XPATH}:build
mkdir -p ./docker/bin ./docker/config ./docker/tools

echo "[+] Extracting binary and config( ability/app.yml/tools)..."
docker cp extract:/open-platform/ability/ability ./docker/bin/
docker cp extract:/open-platform/ability/config/app-dev.yml ./docker/config/
docker cp extract:/open-platform/ability/testcase ./docker/tools/

# 删除多余的解压文件
echo "[+] Deleting extracted files..."
docker rm -f extract

# 删除编译过程镜像
docker image rm ${XREPO}/${XPATH}:build

# 构建发布用的包
echo "[+] Building docker image for release..."
docker build --no-cache -t ${XREPO}/${XPATH}:${XTAG} . -f Dockerfile.release-dev

# 发布到腾讯云 TKE 仓库
echo "[+] Releasing to TKE registry..."
docker login -u REPO_USERNAME -p REPO_PASSWORD ${XREPO}
docker push ${XREPO}/${XPATH}:${XTAG}

# 清理过程输出
rm -fr ./docker

echo "[+] DONE!"
