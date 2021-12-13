package engine

import (
	"fmt"
	"math"
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
	controls := []Point{
		{X: 10, Y: 10},
		{X: float64(10 + width/3), Y: float64(10 + height/4)},
		{X: float64(10 + 2*(width/5)), Y: float64(10 + 2*(height/3))},
		{X: float64(width) - 10, Y: float64(height) - 10},
	}
	sampled := SamplePoints(controls)
	fmt.Println(sampled)
	ctx.Set("fillStyle", color)
	for _, sp := range sampled {
		ctx.Call("beginPath")
		ctx.Call("arc", sp.X, sp.Y, 5, 0, 2*math.Pi)
		ctx.Call("fill")
	}
}
