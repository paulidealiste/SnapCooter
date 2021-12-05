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
		ct.Bearing, ct.X, ct.Y = next(ct.Bearing, ct.X, ct.Y, ct.Size, request.Width, request.Height)
		ctx.Set("fillStyle", ct.Color)
		ctx.Call("fillRect", ct.X, ct.Y, ct.Size, ct.Size)
		ctrs[i] = ct.ObtainJSON()
	}
	return map[string]interface{}{
		"passed": ctrs,
		"error":  erred,
	}
}

func next(b string, cx int, cy int, size int, width int, height int) (string, int, int) {
	var step string
	var x int
	var y int
	nb := b
	passable := false
	for !passable {
		step, x, y, passable = calculateNext(nb, cx, cy, size, width, height)
		if !passable {
			if x > width {
				cx = width - size
			}
			if x < 0 {
				cx = 0
			}
			if y > height {
				cy = height - size
			}
			if y < 0 {
				cy = 0
			}
		}
		nb = bearingFromShortlist(nb)
	}
	return step, x, y
}

func calculateNext(b string, cx int, cy int, size int, width int, height int) (string, int, int, bool) {
	var x int
	var y int
	switch b {
	case roles.E:
		x = cx + size
		y = cy
	case roles.N:
		x = cx
		y = cy - size
	case roles.NE:
		x = cx + size
		y = cy - size
	case roles.NW:
		x = cx - size
		y = cy - size
	case roles.S:
		x = cx
		y = cy + size
	case roles.SE:
		x = cx + size
		y = cy + size
	case roles.SW:
		x = cx - size
		y = cy + size
	case roles.W:
		x = cx - size
		y = cy
	default:
		x = cx
		y = cy
	}
	passable := hardBoundaryPassable(x, y, width, height, size)
	return b, x, y, passable
}

func bearingFromShortlist(b string) string {
	shortlist := collections.StringFilter([]string{roles.E, roles.N, roles.W, roles.S, roles.NE, roles.NW, roles.SW, roles.SE}, func(v string) bool {
		return !strings.Contains(v, b[0:1]) && v != utils.OppositeBearing(b)
	})
	sampled := utils.SampleSlice(shortlist)
	return sampled
}
