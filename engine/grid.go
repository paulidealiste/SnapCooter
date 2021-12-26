package engine

import (
	"syscall/js"

	"github.com/paulidealiste/SnapCooter/roles"
	"github.com/paulidealiste/SnapCooter/utils"
)

func ManageGrid(this js.Value, args []js.Value) interface{} {
	request, err := roles.TranslateGridRequest(args[0].String())
	if err != nil {
		return map[string]interface{}{
			"error": "Error reading the grid request!",
		}
	}
	canvas, err := utils.GetCanvas(request.CanvasID)
	if err != nil {
		return map[string]interface{}{
			"error": "Error reading the grid request!",
		}
	}
	ctx := canvas.Call("getContext", "2d")
	grid := utils.GetPlacementGrid(request.Width, request.Height, request.TileDimension)
	ctx.Call("clearRect", 0, 0, request.Width, request.Height)
	if request.State {
		drawGrid(ctx, request, grid)
	}
	return nil
}

func drawGrid(ctx js.Value, request roles.GridRequest, grid utils.PlacementGrid) {
	ctx.Set("strokeStyle", request.Stroke)
	ctx.Set("lineWidth", request.StrokeWidth)
	ctx.Call("beginPath")
	for i := grid.Padding.PL; i <= request.Width-grid.Padding.PR; i += request.TileDimension {
		ctx.Call("moveTo", i, grid.Padding.PT)
		ctx.Call("lineTo", i, request.Height-grid.Padding.PB)
	}
	for i := grid.Padding.PT; i <= request.Height-grid.Padding.PB; i += request.TileDimension {
		ctx.Call("moveTo", grid.Padding.PL, i)
		ctx.Call("lineTo", request.Width-grid.Padding.PR, i)
	}
	ctx.Call("stroke")
}
