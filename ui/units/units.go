// Package units provides unit values used across the app
package units

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

var (
	// Label is the unit for the app labels
	Label = unit.Dp(50)

	// FlexInset is the unit for flex insets
	FlexInset = unit.Dp(50)
)

func Inset(top,right,bottom, left float32) layout.Inset {
	return layout.Inset{
		Top:    unit.Dp(top),
		Right:  unit.Dp(right),
		Bottom: unit.Dp(bottom),
		Left:   unit.Dp(left),
	}
}
