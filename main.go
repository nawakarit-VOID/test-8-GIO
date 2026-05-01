// Copyright (c) 2026 Nawakarit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := new(app.Window)
		var ops op.Ops

		th := material.NewTheme()
		var btn widget.Clickable

		for {
			e := w.Event() // ✅ ไม่ใช่ range แล้ว

			switch e := e.(type) {

			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				for btn.Clicked(gtx) {
					println("clicked!")
				}

				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &btn, "Click me").Layout(gtx)
				})

				e.Frame(gtx.Ops)
			}
		}
	}()

	app.Main()
}
