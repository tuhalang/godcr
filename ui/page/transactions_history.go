package page

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/raedahgroup/dcrlibwallet"
	"image"
	"image/color"
	"time"
)

type (
	// TransactionsHistory represents for Transactions page of the app
	// It should be fulfill the following conditions:
	// see a list of transactions from my wallet if I have any or an empty page if I have none.
	// sort them by newest or oldest
	// view all of them at once or filter them by received, sent or transfer
	TransactionsHistory struct {
		inset   layout.Inset
		stack   layout.Stack
		theme 	*material.Theme
		heading material.Label
		transactions []dcrlibwallet.Transaction

		gtk        *layout.Context
	}
)

var (
	list              = &layout.List{
		Axis: layout.Vertical,
	}
)

func (p *TransactionsHistory) Init(theme *material.Theme, gtk *layout.Context)  {

	p.heading = theme.H4("Transactions")
	p.theme = theme
	p.inset = layout.UniformInset(unit.Dp(5))
	p.transactions = make([]dcrlibwallet.Transaction, 3)
	go func() {
		time.Sleep(time.Minute)
		p.transactions = append(p.transactions , dcrlibwallet.Transaction{
			Amount: 125,
		})
	}()
	//p.stack.Alignment = layout.W
	p.gtk = gtk
}

func (p *TransactionsHistory) Draw() {
	t := p.theme
	gtk := p.gtk
	layout.Stack{Alignment: layout.SE}.Layout(gtk,
		layout.Expanded(func() {
			layout.Flex{Axis: layout.Vertical}.Layout(gtk,
				layout.Rigid(func() {
					//gtk.Constraints.Width.Min = gtk.Constraints.Width.Max
					layout.UniformInset(unit.Dp(16)).Layout(gtk, func() {
						sz := gtk.Px(unit.Dp(32))
						cs := gtk.Constraints
						gtk.Constraints = layout.RigidConstraints(cs.Constrain(image.Point{X: sz, Y: sz}))
						t.H4("Transactions").Layout(gtk)
					})
				}),
				layout.Rigid(func() {
					layout.UniformInset(unit.Dp(16)).Layout(gtk, func() {
						layout.Flex{}.Layout(gtk,
							layout.Rigid(func() {
								label := t.Label(unit.Dp(16), "Default")
								label.Color = color.RGBA{
									R: 44,
									G: 114,
									B: 255,
									A: 255,
								}
								layout.UniformInset(unit.Dp(5)).Layout(gtk, func() {
									label.Layout(gtk)
								})
							}),
							layout.Rigid(func() {
								layout.UniformInset(unit.Dp(5)).Layout(gtk, func() {
									t.Label(unit.Dp(16), "Wallet 2").Layout(gtk)
								})
							}),
							layout.Rigid(func() {
								layout.UniformInset(unit.Dp(5)).Layout(gtk, func() {
									t.Label(unit.Dp(16), "Wallet 3").Layout(gtk)
								})
							}),
						)
					})
				}),
				layout.Rigid(func() {
					layout.UniformInset(unit.Dp(16)).Layout(gtk, func() {
						layout.Flex{}.Layout(gtk,
							layout.Rigid(func() {
								layout.UniformInset(unit.Dp(5)).Layout(gtk, func() {
									t.Label(unit.Dp(16), "All").Layout(gtk)
								})
							}),
							layout.Rigid(func() {
								layout.UniformInset(unit.Dp(5)).Layout(gtk, func() {
									t.Label(unit.Dp(16), "Newest").Layout(gtk)
								})
							}),
						)
					})
				}),
				layout.Flexed(1, func() {
					layout.UniformInset(unit.Dp(16)).Layout(gtk, func() {
						gtk.Constraints.Width.Min = gtk.Constraints.Width.Max
						p.renderTransactions()
					})
				}),
			)
		}),
	)
}
func rgb(c uint32) color.RGBA {
	return argb((0xff << 24) | c)
}

func argb(c uint32) color.RGBA {
	return color.RGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}

func (p *TransactionsHistory) renderTransactions() {
	t := p.theme
	gtk := p.gtk
	l := layout.List{
		Axis: layout.Vertical,
	}
	l.Layout(gtk, len(p.transactions), func(i int) {
		transaction := p.transactions[i]
		row := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}
		row.Layout(gtk,
			layout.Rigid(func() {
				baseLine := layout.Flex{Axis: layout.Horizontal, Alignment: layout.Baseline}
				baseLine.Layout(gtk,
					layout.Rigid(func() {
						t.Body1(fmt.Sprintf("%d DCR", transaction.Amount)).Layout(gtk)
					}),
					layout.Flexed(1, func() {
						gtk.Constraints.Width.Min = gtk.Constraints.Width.Max
						layout.Align(layout.E).Layout(gtk, func() {
							layout.Inset{Left: unit.Dp(2)}.Layout(gtk, func() {
								t.Caption("Pending").Layout(gtk)
							})
						})
					}),
				)
			}),
		)
	})
}

type fill struct {
	col color.RGBA
}

func (f fill) Layout(gtx *layout.Context) {
	cs := gtx.Constraints
	d := image.Point{X: cs.Width.Min, Y: cs.Height.Min}
	dr := f32.Rectangle{
		Max: f32.Point{X: float32(d.X), Y: float32(d.Y)},
	}
	paint.ColorOp{Color: f.col}.Add(gtx.Ops)
	paint.PaintOp{Rect: dr}.Add(gtx.Ops)
	gtx.Dimensions = layout.Dimensions{Size: d}
}
