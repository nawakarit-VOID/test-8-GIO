#!/bin/bash
set -e
export PATH=/usr/local/go/bin:$PATH

echo " ลบไฟล์..icons.sh... "
sleep 0.2
rm -rf buildicons.sh

echo " ลบไฟล์..image.sh... "
sleep 0.2
rm -rf appimagetool-x86_64.AppImage
rm -rf buildimage.sh

echo " ลบไฟล์..flatpak.sh... "
sleep 0.2
rm -rf buildflatpak.sh
rm -rf buildinstall.sh
rm -rf build-dir
rm -rf flatpak
rm -rf .flatpak-builder
rm -rf repo

echo " ลบไฟล์..exe.sh... "
sleep 0.2
rm -rf app.rc
rm -rf buildexe.sh
rm -rf FyneApp.toml
rm -rf rsrc.syso

echo " ลบไฟล์..clear.sh... "
sleep 1
rm -rf clear.sh

echo "✅ เสร็จแล้ว!"
