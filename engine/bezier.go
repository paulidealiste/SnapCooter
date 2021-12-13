package engine

import (
	"github.com/paulidealiste/SnapCooter/utils"
)

type Point struct {
	X float64
	Y float64
}

func SamplePoints(control []Point) []Point {
	ts := utils.RangeGen(0, 1, 0.1)
	sampled := make([]Point, len(ts))
	for i, t := range ts {
		sample := Point{}
		SampleBezier(control, t, &sample)
		sampled[i] = sample
	}
	return sampled
}

func SampleBezier(points []Point, t float64, sample *Point) {
	if len(points) == 1 {
		*sample = points[0]
	} else {
		shorts := make([]Point, len(points)-1)
		for i, np := range shorts {
			np.X = (1.0-t)*points[i].X + t*points[i+1].X
			np.Y = (1.0-t)*points[i].Y + t*points[i+1].Y
			shorts[i] = np
		}
		SampleBezier(shorts, t, sample)
	}
}
