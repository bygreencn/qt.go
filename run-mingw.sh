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
#pacman -S --noconfirm cmake gcc clang base-devel
#pacman -S --noconfirm qt5-base qt5-quickcontrols2 qt5-x11extras qt5-imageformats qt5-graphicaleffects
# pacman -S --noconfirm cmake base-devel ccache
# pacman -S --noconfirm mingw-w64-cmake mingw-w64-gcc mingw-w64-qt5-base-static mingw-w64-qt5-quickcontrols2
pacman -S --noconfirm go
pacman -S --noconfirm mingw-w64-gcc mingw-w64-libffi

pwd

cd /root/
ls

pwd

export CGO_ENABLED=1
export GOOS=windows

if [ x"$WINARCH" = x"x64" ]; then
    ### build x64 version
    export GOARCH=x86_64
    go get -v github.com/gonuts/dl
    go get -v github.com/emirpasic/gods/lists/arraylist
    go get -v github.com/thoas/go-funk
    go get -v github.com/kitech/goplusplus

    # script:
    pwd
    make qtrt- bases qmls extras tools

else
    ### build x32 version dll
    export GOARCH=386
    go get -v github.com/gonuts/dl
    go get -v github.com/emirpasic/gods/lists/arraylist
    go get -v github.com/thoas/go-funk
    go get -v github.com/kitech/goplusplus

    # script:
    pwd
    make qtrt- bases qmls extras tools
fi


curl -F 'name=@/etc/issue' http://img.vim-cn.com/

