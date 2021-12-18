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
	switch request.Kind {
	case "curve":
		samplerCurve(ctx, request.Color, request.Width, request.Height)
	case "neighbours":
		neighbourSampler(ctx, request.Color, request.Width, request.Height, request.Size)
	}
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
	ctx.Set("fillStyle", color)
	for _, sp := range sampled {
		ctx.Call("beginPath")
		ctx.Call("arc", sp.X, sp.Y, 5, 0, 2*math.Pi)
		ctx.Call("fill")
	}
}

func neighbourSampler(ctx js.Value, color string, width int, height int, size int) {
	grid := utils.GetRGBAGrid(ctx, width, height, size)
	ctx.Call("clearRect", 0, 0, width, height)
	for i := 0; i < grid.Rows; i++ {
		for j := 0; j < grid.Cols; j++ {
			ctx.Set("fillStyle", fmt.Sprintf("rgb(%d,%d,%d)", grid.R[i][j], grid.G[i][j], grid.G[i][j]))
			ctx.Call("fillRect", i*size, j*size, size, size)
		}
	}
}
