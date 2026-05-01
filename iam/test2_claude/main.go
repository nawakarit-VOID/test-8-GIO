// card_widget.go — Beautiful custom card widgets with spring press animation
//
// Setup & run:
//   go mod init cardwidget
//   go get gioui.org@latest golang.org/x/exp
//   go run card_widget.go

package main

import (
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"golang.org/x/exp/shiny/materialdesign/icons"
)

// ─────────────────────────────────────────────────────────────────────────────
// CardWidget — gradient card with spring press animation
// ─────────────────────────────────────────────────────────────────────────────

type CardWidget struct {
	button    *widget.Clickable
	icon      *widget.Icon
	title     string
	subtitle  string
	color1    color.NRGBA
	color2    color.NRGBA
	pressed   bool
	pressTime time.Time
}

func NewCard(title, subtitle string, ic *widget.Icon, c1, c2 color.NRGBA) *CardWidget {
	return &CardWidget{
		button:   new(widget.Clickable),
		icon:     ic,
		title:    title,
		subtitle: subtitle,
		color1:   c1,
		color2:   c2,
	}
}

// easeOutBack — snappy spring rebound curve.
func easeOutBack(t float32) float32 {
	const c1 = float32(1.70158)
	const c3 = c1 + 1
	t -= 1
	return 1 + c3*float32(math.Pow(float64(t), 3)) + c1*float32(math.Pow(float64(t), 2))
}

func (c *CardWidget) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	for c.button.Clicked(gtx) {
		c.pressed = true
		c.pressTime = gtx.Now
	}

	scale := float32(1.0)
	if c.pressed {
		dt := float32(gtx.Now.Sub(c.pressTime).Seconds())
		const dur = float32(0.3)
		if dt < dur {
			t := dt / dur
			if t < 0.3 {
				scale = 1.0 - 0.06*(t/0.3)
			} else {
				scale = 0.94 + 0.06*easeOutBack((t-0.3)/0.7)
			}
			gtx.Execute(op.InvalidateCmd{})
		} else {
			c.pressed = false
		}
	}

	w, h := gtx.Dp(unit.Dp(200)), gtx.Dp(unit.Dp(130))

	macro := op.Record(gtx.Ops)
	dims := c.draw(gtx, th, w, h)
	call := macro.Stop()

	if scale != 1.0 {
		cx, cy := float32(w)/2, float32(h)/2
		tr := f32.Affine2D{}.Scale(f32.Pt(cx, cy), f32.Pt(scale, scale))
		op.Affine(tr).Add(gtx.Ops)
	}
	call.Add(gtx.Ops)
	return dims
}

func (c *CardWidget) draw(gtx layout.Context, th *material.Theme, w, h int) layout.Dimensions {
	sz := image.Pt(w, h)
	radius := gtx.Dp(unit.Dp(20))

	rr := clip.RRect{Rect: image.Rectangle{Max: sz}, NE: radius, NW: radius, SE: radius, SW: radius}
	defer rr.Push(gtx.Ops).Pop()

	// Gradient fill
	paint.LinearGradientOp{
		Stop1:  f32.Pt(0, 0),
		Stop2:  f32.Pt(float32(w), float32(h)),
		Color1: c.color1,
		Color2: c.color2,
	}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	// Top-left sheen
	func() {
		sheenRR := clip.RRect{
			Rect: image.Rectangle{Max: image.Pt(w*2/3, h/2)},
			NE:   gtx.Dp(unit.Dp(60)), NW: radius, SE: gtx.Dp(unit.Dp(60)),
		}
		defer sheenRR.Push(gtx.Ops).Pop()
		paint.ColorOp{Color: color.NRGBA{R: 255, G: 255, B: 255, A: 22}}.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)
	}()

	// Content
	gtx.Constraints = layout.Exact(sz)
	layout.Inset{Top: unit.Dp(16), Left: unit.Dp(18), Right: unit.Dp(18), Bottom: unit.Dp(14)}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceBetween, Alignment: layout.Start}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					iconSz := gtx.Dp(unit.Dp(30))
					gtx.Constraints = layout.Exact(image.Pt(iconSz, iconSz))
					return c.icon.Layout(gtx, color.NRGBA{R: 255, G: 255, B: 255, A: 210})
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Dimensions{Size: gtx.Constraints.Min}
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					l := material.H6(th, c.title)
					l.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
					return l.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					l := material.Caption(th, c.subtitle)
					l.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 175}
					return l.Layout(gtx)
				}),
			)
		},
	)

	// Clickable overlay
	c.button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{Size: sz}
	})

	return layout.Dimensions{Size: sz}
}

// ─────────────────────────────────────────────────────────────────────────────
// App entry point
// ─────────────────────────────────────────────────────────────────────────────

var allCards []*CardWidget

func main() {
	go func() {
		w := new(app.Window)
		w.Option(app.Title("Custom Card Widgets ✨"), app.Size(unit.Dp(720), unit.Dp(440)))
		if err := run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	icStar, _ := widget.NewIcon(icons.ToggleStar)
	icCloud, _ := widget.NewIcon(icons.FileCloudUpload)
	icHeart, _ := widget.NewIcon(icons.ActionFavorite)
	icCode, _ := widget.NewIcon(icons.ActionCode)
	icFlash, _ := widget.NewIcon(icons.ImageFlashOn)
	icMap, _ := widget.NewIcon(icons.MapsMap)

	allCards = []*CardWidget{
		NewCard("Starred", "Your favourites", icStar,
			color.NRGBA{R: 255, G: 140, B: 0, A: 255},
			color.NRGBA{R: 255, G: 75, B: 55, A: 255}),
		NewCard("Cloud Sync", "Everything backed up", icCloud,
			color.NRGBA{R: 25, G: 130, B: 230, A: 255},
			color.NRGBA{R: 0, G: 188, B: 212, A: 255}),
		NewCard("Liked", "Things you love", icHeart,
			color.NRGBA{R: 236, G: 28, B: 100, A: 255},
			color.NRGBA{R: 156, G: 39, B: 176, A: 255}),
		NewCard("Dev Tools", "Code & build", icCode,
			color.NRGBA{R: 32, G: 168, B: 152, A: 255},
			color.NRGBA{R: 60, G: 160, B: 70, A: 255}),
		NewCard("Quick", "Lightning fast", icFlash,
			color.NRGBA{R: 251, G: 192, B: 45, A: 255},
			color.NRGBA{R: 255, G: 130, B: 0, A: 255}),
		NewCard("Explore", "Discover places", icMap,
			color.NRGBA{R: 63, G: 81, B: 181, A: 255},
			color.NRGBA{R: 33, G: 150, B: 243, A: 255}),
	}

	var ops op.Ops
	for {
		e := w.Event()
		switch e := e.(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			drawScene(gtx, th)
			e.Frame(gtx.Ops)
		}
	}
}

func drawScene(gtx layout.Context, th *material.Theme) layout.Dimensions {
	paint.Fill(gtx.Ops, color.NRGBA{R: 15, G: 15, B: 22, A: 255})

	return layout.Inset{Top: unit.Dp(36), Left: unit.Dp(36), Right: unit.Dp(36), Bottom: unit.Dp(36)}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					l := material.H4(th, "My Cards  ✨")
					l.Color = color.NRGBA{R: 235, G: 235, B: 255, A: 255}
					return layout.Inset{Bottom: unit.Dp(32)}.Layout(gtx, l.Layout)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return cardRow(gtx, th, allCards[0:3])
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{Height: unit.Dp(20)}.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return cardRow(gtx, th, allCards[3:6])
				}),
			)
		},
	)
}

func cardRow(gtx layout.Context, th *material.Theme, row []*CardWidget) layout.Dimensions {
	children := make([]layout.FlexChild, 0, len(row)*2-1)
	for i, c := range row {
		card := c
		if i > 0 {
			children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Spacer{Width: unit.Dp(20)}.Layout(gtx)
			}))
		}
		children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return card.Layout(gtx, th)
		}))
	}
	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx, children...)
}
