// Copyright (c) 2026 Nawakarit
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License v3.0.
package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {

		w := new(app.Window)

		w.Option(app.Title("test8"), app.Size(unit.Dp(500), unit.Dp(500)))
		//w := &app.Window{} //เหมือนกัน 100%  -- สร้างหน้าต่างเปล่า
		//w.Option(app.Title("Custom Card Widgets ✨"), app.Size(unit.Dp(720), unit.Dp(440)))
		//w := &app.Window{} //เหมือนกัน 100%
		//app.Window = struct ของ window
		//new(...) = สร้าง pointer ไปยัง struct นั้น
		if err := run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0) // ตัวปิดหน้าต่าง
	}()

	app.Main() //เริ่มระบบของ Gio
}

func run(w *app.Window) error {
	var ops op.Ops //ตัวแปรชื่อ ops ที่เป็นประเภท op.Ops //กระดาษรายการคำสั่ง UI แต่ละชิ้น =  เขียนคำสั่งลงไป
	//op.Ops - รายการคำสั่งวาด //var ops op.Ops - รายการคำสั่งนี้ชื่อว่า ops //
	th := material.NewTheme()

	var btn widget.Clickable //จะใช้คู่กับ &btn ตอน layout เสมอ
	//btn := new(widget.Clickable)
	//ตรงนี้ไม่ปุ่ม แต่เป็นตัวแปร ที่ชื่อว่า btn

	for {
		e := w.Event() // รับ event จาก OS
		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err // ตัวปิดหน้าต่าง

		case app.FrameEvent: //ใช้วาด UI
			gtx := app.NewContext(&ops, e) //เอา
			//app.NewContext(&ops, e) - เริ่มเขียนคำสั่ง

			for btn.Clicked(gtx) {
				println("clicked!")
			}

			layout.Flex{
				Axis: layout.Vertical, // แนวตั้ง
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.H6(th, "Title").Layout(gtx)
					})
				}),

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &btn, "Click").Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Button(th, &btn, "Click me").Layout(gtx)
					//Button สร้างอยู่ตรงนี้ แล้วใช้ theme, ฟังชั้น, ชื่อที่แสดง
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Label(th, unit.Sp(16), "Click me").Layout(gtx)
				}),
			)

			layout.Inset{
				Top:    unit.Dp(300),
				Bottom: unit.Dp(10),
				Left:   unit.Dp(20),
				Right:  unit.Dp(20),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, &btn, "Click").Layout(gtx)
			})

			e.Frame(gtx.Ops) //สั่งวาดจริง

		}
	}
}

/*
1. layout.Flex (แถว / คอลัมน์)

layout.Flex{
    Axis: layout.Vertical, // แนวตั้ง
}.Layout(gtx,
    layout.Rigid(func(gtx layout.Context) layout.Dimensions {
        return material.H6(th, "Title").Layout(gtx)
    }),
    layout.Rigid(func(gtx layout.Context) layout.Dimensions {
        return material.Button(th, &btn, "Click").Layout(gtx)
    }),
)

	Title
[Button]

🔥 แนวนอน (row)
layout.Flex{
    Axis: layout.Horizontal,
}.Layout(gtx,
    layout.Rigid(func(gtx layout.Context) layout.Dimensions {
        return material.Button(th, &btn1, "A").Layout(gtx)
    }),
    layout.Rigid(func(gtx layout.Context) layout.Dimensions {
        return material.Button(th, &btn2, "B").Layout(gtx)
    }),
)
[A] [B]


2. layout.Flexed (ยืดพื้นที่)

layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
    layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
        return material.Button(th, &btn1, "Left").Layout(gtx)
    }),
    layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
        return material.Button(th, &btn2, "Right").Layout(gtx)
    }),
)👉 ขวาจะใหญ่กว่า 2 เท่า

3. layout.Inset (padding)

layout.Inset{
    Top: unit.Dp(10),
    Left: unit.Dp(10),
}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
    return material.Button(th, &btn, "Padding").Layout(gtx)
})

4. layout.Center

layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
    return material.H3(th, "Centered").Layout(gtx)
})

5. layout.Stack (ซ้อนกัน)

layout.Stack{}.Layout(gtx,
    layout.Stacked(func(gtx layout.Context) layout.Dimensions {
        return material.Label(th, unit.Sp(20), "Background").Layout(gtx)
    }),
    layout.Stacked(func(gtx layout.Context) layout.Dimensions {
        return material.Button(th, &btn, "Top").Layout(gtx)
    }),
)


*/
