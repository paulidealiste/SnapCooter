// Package utils provide cooter assemblies utility functions
package utils

import (
	"math"
	"syscall/js"

	"github.com/paulidealiste/SnapCooter/roles"
)

type RGBAGrid struct {
	R    [][]byte
	G    [][]byte
	B    [][]byte
	A    [][]byte
	Rows int
	Cols int
}

type PlacementCell struct {
	I int
	X int
	Y int
	W int
	H int
}

type PlacementPadding struct {
	PT int
	PB int
	PL int
	PR int
}

type PlacementGrid struct {
	Cells   []PlacementCell
	Rows    int
	Cols    int
	Padding PlacementPadding
}

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

func GetRGBACell(ctx js.Value, x int, y int, s int) []byte {
	jsdata := ctx.Call("getImageData", x, y, s, s).Get("data")
	godata := make([]byte, 4*s*s)
	js.CopyBytesToGo(godata, jsdata)
	return godata
}

func GetRGBAGrid(ctx js.Value, w int, h int, s int) RGBAGrid {
	rows := int(math.Floor(float64(h) / float64(s)))
	cols := int(math.Floor(float64(w) / float64(s)))
	r := make([][]byte, rows)
	g := make([][]byte, rows)
	b := make([][]byte, rows)
	a := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		r[i] = make([]byte, cols)
		g[i] = make([]byte, cols)
		b[i] = make([]byte, cols)
		a[i] = make([]byte, cols)
		for j := 0; j < cols; j++ {
			rgbacell := GetRGBACell(ctx, i*s, j*s, s)
			r[i][j] = rgbacell[0]
			g[i][j] = rgbacell[1]
			b[i][j] = rgbacell[2]
			a[i][j] = rgbacell[3]
		}
	}
	grid := RGBAGrid{R: r, G: g, B: b, A: a, Rows: rows, Cols: cols}
	return grid
}

func GetPlacementGrid(w int, h int, s int) PlacementGrid {
	nx := int(math.Floor(float64(w)/float64(s))) - 2
	ny := int(math.Floor(float64(h)/float64(s))) - 2
	px := w - nx*s
	py := h - ny*s

	pl := int(math.Ceil(float64(px)/2.0) - 0.5)
	pt := int(math.Ceil(float64(py)/2.0) - 0.5)
	pr := w - nx*s - pl
	pb := h - ny*s - pt
	cells := make([]PlacementCell, nx*ny)
	for i := 0; i < nx; i++ {
		for j := 0; j < ny; j++ {
			cell := PlacementCell{X: pl + i*s, Y: pt + j*s, W: s, H: s, I: i*ny + j}
			cells[cell.I] = cell
		}
	}

	grid := PlacementGrid{Cells: cells, Rows: ny, Cols: nx, Padding: PlacementPadding{PT: pt, PB: pb, PL: pl, PR: pr}}
	return grid
}
