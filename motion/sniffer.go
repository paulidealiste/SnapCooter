package motion

import (
	"syscall/js"

	"github.com/paulidealiste/SnapCooter/roles"
	"github.com/paulidealiste/SnapCooter/utils"
)

func hardBoundaryPassable(x int, y int, width int, height int, size int) bool {
	if x > width || y > height || x < 0 || y < 0 {
		return false
	}
	return true
}

func sniffDestination(x int, y int, ct roles.Cooter, ctx js.Value) bool {
	jsdata := ctx.Call("getImageData", x, y, ct.Size, ct.Size).Get("data")
	godata := make([]byte, 4*ct.Size*ct.Size)
	js.CopyBytesToGo(godata, jsdata)
	r := int(godata[0])
	g := int(godata[1])
	b := int(godata[2])
	if r == 0 && g == 0 && b == 0 {
		return true
	}
	if r != ct.RGB[0] || g != ct.RGB[1] || b != ct.RGB[2] {
		return utils.SimpleProbable(ct.Determination)
	}
	if r == ct.RGB[0] && g == ct.RGB[1] && b == ct.RGB[2] {
		return true
	}
	return false
}
