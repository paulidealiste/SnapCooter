// Package utils provide cooter assemblies utility functions
package utils

import (
	"math"
	"syscall/js"

	"github.com/paulidealiste/SnapCooter/roles"
)

func GetCanvasSetup(cas string) (js.Value, roles.Setup, error) {
	document := js.Global().Get("document")
	if !document.Truthy() {
		return js.Value{}, roles.Setup{}, js.Error{}
	}
	setup, err := roles.TranslateSetup(cas)
	if err != nil {
		return js.Value{}, roles.Setup{}, js.Error{}
	}
	canvas := document.Call("getElementById", setup.Document.CanvasID)
	if !canvas.Truthy() {
		return js.Value{}, roles.Setup{}, js.Error{}
	}
	return canvas, setup, nil
}

func GetCanvas(cid string) (js.Value, error) {
	document := js.Global().Get("document")
	if !document.Truthy() {
		return js.Value{}, js.Error{}
	}
	canvas := document.Call("getElementById", cid)
	if !canvas.Truthy() {
		return js.Value{}, js.Error{}
	}
	return canvas, nil
}

func LinearScale(pmin int, pmax int, cmin int, cmax int, x int) int {
	fpmin := float64(pmin)
	fpmax := float64(pmax)
	fcmin := float64(cmin)
	fcmax := float64(cmax)
	fx := float64(x)
	i := int(math.Round(((fx - fpmin) / (fpmax - fpmin)) * ((fcmax - fcmin) + fcmin)))
	return i
}

func OppositeBearing(b string) string {
	switch b {
	case roles.E:
		return roles.W
	case roles.N:
		return roles.S
	case roles.NE:
		return roles.SW
	case roles.NW:
		return roles.SE
	case roles.S:
		return roles.N
	case roles.SE:
		return roles.NW
	case roles.SW:
		return roles.NE
	case roles.W:
		return roles.E
	}
	return roles.E
}

func RangeGen(start float64, end float64, step float64) []float64 {
	rng := make([]float64, 1+int(math.Ceil((end-start)/step)))
	i := 0
	for start < end {
		rng[i] = start
		start = start + step
		i = i + 1
	}
	return rng
}
