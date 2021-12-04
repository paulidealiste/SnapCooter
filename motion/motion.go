//Package motion usees the collection of cooters to calculate all the movements
package motion

import (
	"fmt"
	"log"
	"syscall/js"

	"github.com/paulidealiste/SnapCooter/roles"
	"github.com/paulidealiste/SnapCooter/utils"
)

func CooterStep(this js.Value, args []js.Value) interface{} {
	erred := make([]interface{}, 0)
	request, err := roles.TranslateMotionRequest(args[0].String())
	if err != nil {
		log.Fatal(err)
	}
	actives := make([]roles.Cooter, len(request.Cooters))
	for i, scs := range request.Cooters {
		ct, err := roles.TranslateCooter(scs)
		if err != nil {
			log.Fatal(err)
		}
		actives[i] = ct
	}
	canvas, err := utils.GetCanvas(request.CanvasID)
	if err != nil {
		return map[string]interface{}{
			"error": "Error reading the setup request!",
		}
	}
	ctx := canvas.Call("getContext", "2d")
	canvas.Set("style", "border: 1px solid #419D78")
	ctrs := make([]interface{}, len(actives))
	for i, ct := range actives {
		ct.X = utils.RandomInt(0, request.Width)
		ct.Y = utils.RandomInt(0, request.Height)
		ctx.Set("fillStyle", ct.Color)
		ctx.Call("fillRect", ct.X, ct.Y, 10, 10)
		ctrs[i] = ct.ObtainJSON()
		fmt.Println("")
	}
	return map[string]interface{}{
		"passed": ctrs,
		"error":  erred,
	}
}
