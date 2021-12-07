//Package motion usees the collection of cooters to calculate all the movements and boundaries
package motion

import (
	"log"
	"strings"
	"syscall/js"

	"github.com/paulidealiste/SnapCooter/collections"
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
		ct.Bearing, ct.X, ct.Y = next(ct.Bearing, ct, request.Width, request.Height, ctx)
		ctx.Set("fillStyle", ct.Color)
		ctx.Call("fillRect", ct.X, ct.Y, ct.Size, ct.Size)
		ctrs[i] = ct.ObtainJSON()
	}
	return map[string]interface{}{
		"passed": ctrs,
		"error":  erred,
	}
}

func next(b string, ct roles.Cooter, width int, height int, ctx js.Value) (string, int, int) {
	var step string
	var x int
	var y int
	nb := b
	passable := false
	for !passable {
		step, x, y, passable = calculateNext(nb, ct, width, height, ctx)
		if !passable {
			if x > width {
				ct.X = width - ct.Size
			}
			if x < 0 {
				ct.X = 0
			}
			if y > height {
				ct.Y = height - ct.Size
			}
			if y < 0 {
				ct.Y = 0
			}
		}
		nb = bearingFromShortlist(nb)
	}
	return step, x, y
}

func calculateNext(b string, ct roles.Cooter, width int, height int, ctx js.Value) (string, int, int, bool) {
	var x int
	var y int
	switch b {
	case roles.E:
		x = ct.X + ct.Size
		y = ct.Y
	case roles.N:
		x = ct.X
		y = ct.Y - ct.Size
	case roles.NE:
		x = ct.X + ct.Size
		y = ct.Y - ct.Size
	case roles.NW:
		x = ct.X - ct.Size
		y = ct.Y - ct.Size
	case roles.S:
		x = ct.X
		y = ct.Y + ct.Size
	case roles.SE:
		x = ct.X + ct.Size
		y = ct.Y + ct.Size
	case roles.SW:
		x = ct.X - ct.Size
		y = ct.Y + ct.Size
	case roles.W:
		x = ct.X - ct.Size
		y = ct.Y
	default:
		x = ct.X
		y = ct.Y
	}
	passable := hardBoundaryPassable(x, y, width, height, ct.Size)
	if passable {
		passable = sniffDestination(x, y, ct, ctx)
	}
	return b, x, y, passable
}

func bearingFromShortlist(b string) string {
	shortlist := collections.StringFilter([]string{roles.E, roles.N, roles.W, roles.S, roles.NE, roles.NW, roles.SW, roles.SE}, func(v string) bool {
		return !strings.Contains(v, b[0:1]) && v != utils.OppositeBearing(b)
	})
	sampled := utils.SampleSlice(shortlist)
	return sampled
}
