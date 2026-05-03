#!/bin/bash
set -e

INSTALL_DIR="$HOME/Applications"
ICON_DIR="$HOME/.local/share/icons/hicolor/256x256/apps"
DESKTOP_DIR="$HOME/.local/share/applications"

echo "ตรวจเช็คไฟล์"
sleep 1
[ -f "{{.Name}}.png" ] || { echo "❌ icon.png missing"; exit 1; }
[ -f "{{.Name}}.desktop" ] || { echo "❌ .desktop missing"; exit 1; }
[ -f "{{.Name}}-x86_64.AppImage" ] || { echo "❌ appimage missing"; exit 1; }

echo "CP..."
sleep 1
cp {{.Name}}.png $ICON_DIR
cp {{.Name}}.desktop $DESKTOP_DIR
cp {{.Name}}-x86_64.AppImage $INSTALL_DIR

echo "✅ เสร็จแล้ว"
