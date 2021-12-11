package engine

import (
	"syscall/js"

	"github.com/paulidealiste/SnapCooter/roles"
	"github.com/paulidealiste/SnapCooter/utils"
)

func DrawSampler(this js.Value, args []js.Value) interface{} {
	request, err := roles.TranslateSamplerRequest(args[0].String())
	if err != nil {
		return map[string]interface{}{
			"error": "Error reading the sampler request!",
		}
	}
	canvas, err := utils.GetCanvas(request.CanvasID)
	if err != nil {
		return map[string]interface{}{
			"error": "Error reading the sampler request!",
		}
	}
	ctx := canvas.Call("getContext", "2d")
	samplerCurve(ctx, request.Color, request.Width, request.Height)
	return nil
}

func samplerCurve(ctx js.Value, color string, width int, height int) {
	ctx.Set("strokeStyle", color)
	ctx.Call("beginPath")
	ctx.Call("moveTo", 10, 10)
	ctx.Call("bezierCurveTo", 10+width/3, 10+height/4, 10+2*(width/5), 10+2*(height/3), width-10, height-10)
	ctx.Call("stroke")
}
