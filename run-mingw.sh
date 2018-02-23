#!/bin/bash

# stop as soon as one of steps will fail
set -e -o pipefail

set -x
WINARCH=$1
echo "$0, $1"
cat /etc/issue

echo "[ownstuff]" >> /etc/pacman.conf
echo "SigLevel = Optional TrustAll" >> /etc/pacman.conf
echo "Server = https://martchus.no-ip.biz/repo/arch/ownstuff/os/\$arch" >> /etc/pacman.conf

pacman -Suyy
pacman -S --noconfirm go git
pacman -S --noconfirm mingw-w64-gcc mingw-w64-libffi

pwd

cd /root/
ls

pwd

export CGO_ENABLED=1
export GOOS=windows

git clone github.com/gonuts/dl
sed -i 's/NoLoad/\/\/NoLoad/' dl/dl.go
sed -i 's/NoDelete/\/\/NoDelete/' dl/dl.go
mkdir $GOPATH/src/github.com/gonuts -p
cp -a dl $GOPATH/src/github.com/gonuts/

if [ x"$WINARCH" = x"x64" ]; then
    ### build x64 version
    export GOARCH=x86_64
    go get -v github.com/emirpasic/gods/lists/arraylist
    go get -v github.com/thoas/go-funk
    go get -v github.com/kitech/goplusplus

    # script:
    pwd
    make qtrt- bases qmls extras tools

else
    ### build x32 version dll
    export GOARCH=386
    go get -v github.com/emirpasic/gods/lists/arraylist
    go get -v github.com/thoas/go-funk
    go get -v github.com/kitech/goplusplus

    # script:
    pwd
    make qtrt- bases qmls extras tools
fi


curl -F 'name=@/etc/issue' http://img.vim-cn.com/

