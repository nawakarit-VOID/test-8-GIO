package ui

import (
	"image/color"

	"gioui.org/widget/material"
)

type Mode int

const (
	Auto Mode = iota
	Light
	Dark
)

type ThemeManager struct {
	Theme  *material.Theme
	Mode   Mode
	isDark bool
}

// 🔥 สร้าง Theme Manager
func NewThemeManager() *ThemeManager {
	th := material.NewTheme()

	tm := &ThemeManager{
		Theme: th,
		Mode:  Auto,
	}

	// default
	tm.applyLight()

	return tm
}

// 🔁 Update จาก OS หรือ user
func (tm *ThemeManager) Update(isDarkFromOS bool) {
	var useDark bool

	switch tm.Mode {
	case Auto:
		useDark = isDarkFromOS
	case Light:
		useDark = false
	case Dark:
		useDark = true
	}

	// 👇 กัน reapply ซ้ำ
	if useDark == tm.isDark {
		return
	}

	tm.isDark = useDark

	if useDark {
		tm.applyDark()
	} else {
		tm.applyLight()
	}
}

// 🌞 Light Theme
func (tm *ThemeManager) applyLight() {
	tm.Theme.Palette = material.Palette{
		Bg:         color.NRGBA{R: 245, G: 245, B: 245, A: 255},
		Fg:         color.NRGBA{R: 20, G: 20, B: 20, A: 255},
		ContrastBg: color.NRGBA{R: 0, G: 120, B: 255, A: 255},
		ContrastFg: color.NRGBA{R: 255, G: 255, B: 255, A: 255},
	}
}

// 🌙 Dark Theme
func (tm *ThemeManager) applyDark() {
	tm.Theme.Palette = material.Palette{
		Bg:         color.NRGBA{R: 20, G: 20, B: 20, A: 255},
		Fg:         color.NRGBA{R: 230, G: 230, B: 230, A: 255},
		ContrastBg: color.NRGBA{R: 0, G: 120, B: 255, A: 255},
		ContrastFg: color.NRGBA{R: 255, G: 255, B: 255, A: 255},
	}
}

// 🎨 เปลี่ยน accent สีปุ่ม
func (tm *ThemeManager) SetAccent(c color.NRGBA) {
	p := tm.Theme.Palette
	p.ContrastBg = c
	tm.Theme.Palette = p
}
