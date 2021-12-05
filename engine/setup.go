//Package engine exposes all the interop functions used directly in javascript
package engine

import (
	"syscall/js"
	"time"

	"github.com/paulidealiste/SnapCooter/roles"
	"github.com/paulidealiste/SnapCooter/utils"
)

func CooterSetup(this js.Value, args []js.Value) interface{} {
	canvas, setup, err := utils.GetCanvasSetup(args[0].String())
	if err != nil {
		return map[string]interface{}{
			"error": "Error reading the setup request!",
		}
	}
	canvas.Set("style", "border: 1px solid grey")
	ctx := canvas.Call("getContext", "2d")
	ctx.Call("clearRect", 0, 0, setup.Width, setup.Height)
	cooters := make([]interface{}, setup.CooterCount)
	for i := 0; i < setup.CooterCount; i++ {
		x := utils.RandomInt(0, setup.Width)
		y := utils.RandomInt(0, setup.Height)
		fill := utils.ColorForPosition(setup.Palette, x, y, setup.Width, setup.Height)
		cc := roles.Cooter{
			ID:            utils.RandomInt(9000, 90000),
			Bearing:       utils.RandomBearing(),
			Name:          utils.RandomName(),
			X:             x,
			Y:             y,
			Color:         fill,
			Determination: 0.5,
			Friendliness:  0.5,
			Size:          setup.CooterSize,
		}
		cooters[i] = cc.ObtainJSON()
		ctx.Set("fillStyle", fill)
		ctx.Call("fillRect", x, y, cc.Size, cc.Size)
		time.Sleep(1 * time.Millisecond)
	}

	return map[string]interface{}{
		"passed": cooters,
		"error":  make([]interface{}, 0),
	}
}
