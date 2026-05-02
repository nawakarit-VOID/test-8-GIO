#!/bin/bash
set -e
export PATH=/usr/local/go/bin:$PATH

INPUT="icon.png"
OUTDIR="icons"

echo " 🖼️ สร้าง..ico..."
sleep 1
convert $INPUT -define icon:auto-resize=256,128,64,48,32,16 icon.ico

echo " 🖼️ สร้าง..png..."
sleep 1
[ -f "icon.png" ] || { echo "❌ icon.png missing"; exit 1; }

mkdir -p $OUTDIR

SIZES=(512 256 128 64 32 16)

for SIZE in "${SIZES[@]}"; do
  convert "$INPUT" \
    -resize ${SIZE}x${SIZE} \
    "$OUTDIR/icon-${SIZE}.png"
done

echo "✅ เสร็จแล้ว!"
