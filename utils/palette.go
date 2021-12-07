package utils

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"syscall/js"
)

type InterpolatedPaletteConfig struct {
	Start string `json:"Start"`
	End   string `json:"End"`
	Count int    `json:"Count"`
	Type  string `json:"Type"`
}

type RGB struct {
	R, G, B float64
}

type HSV struct {
	H, S, V float64
}

func TranslatePaletteConfig(ipconfig string) (InterpolatedPaletteConfig, error) {
	ipp := InterpolatedPaletteConfig{}
	if err := json.Unmarshal([]byte(ipconfig), &ipp); err != nil {
		return InterpolatedPaletteConfig{}, err
	}
	return ipp, nil
}

func CooterInterpolatedPalettes(this js.Value, args []js.Value) interface{} {
	configurations := make([]InterpolatedPaletteConfig, len(args))
	for i, arg := range args {
		cfg, err := TranslatePaletteConfig(arg.String())
		if err != nil {
			log.Fatal(err)
		}
		configurations[i] = cfg
	}
	palettes := make([]interface{}, len(configurations))
	for i, cfg := range configurations {
		palettes[i] = createInterpolatedList(cfg)
	}

	return palettes
}

func createInterpolatedList(cfg InterpolatedPaletteConfig) []interface{} {
	palette := make([]interface{}, cfg.Count)
	factor := 1.0 / float64(cfg.Count)
	for i := 0; i < cfg.Count; i++ {
		start, err := HexToRgb(cfg.Start)
		if err != nil {
			log.Fatal(err)
		}
		end, err := HexToRgb(cfg.End)
		if err != nil {
			log.Fatal(err)
		}
		switch cfg.Type {
		case "RGB":
			t := float64(i) * factor
			rinc := lerpRGB(start, end, t)
			palette[i] = fmt.Sprintf("rgb(%f,%f,%f)", rinc.R, rinc.G, rinc.B)
		case "HSV":
			t := float64(i) * factor
			hstart := RgbToHsv(start)
			hend := RgbToHsv(end)
			hinc := lerpHSV(hstart, hend, t)
			rfin := HsvToRgb(hinc)
			palette[i] = fmt.Sprintf("rgb(%f,%f,%f)", math.Round(rfin.R), math.Round(rfin.G), math.Round(rfin.B))
		}
	}
	return palette
}

func lerpRGB(start RGB, end RGB, t float64) RGB {
	inc := start
	inc.R = float64(inc.R) + t*float64((end.R-start.R))
	inc.G = float64(inc.G) + t*float64((end.G-start.G))
	inc.B = float64(inc.B) + t*float64((end.R-start.B))
	return inc
}

func lerpHSV(start HSV, end HSV, t float64) HSV {
	h := 1.0
	d := end.H - start.H
	if start.H > end.H {
		h3 := end.H
		end.H = start.H
		start.H = h3
		d = -d
		t = 1 - t
	}
	if d > 0.5 {
		start.H = start.H + 1
		h = float64(int((start.H + t*(end.H-start.H))) % 1)
	}
	if d <= 0.5 {
		h = start.H + t*d
	}
	return HSV{h, start.S + t*(end.S-start.S), start.V + t*(end.V-start.V)}
}

func HexToRgb(shex string) (RGB, error) {
	rgb := RGB{}
	if shex[0] != '#' {
		return rgb, errors.New("hex colors must start with a hash")
	}
	switch len(shex) {
	case 7:
		decoded, err := hex.DecodeString(shex[1:])
		if err != nil {
			log.Fatal(err)
		}
		rgb = RGB{float64(decoded[0]), float64(decoded[1]), float64(decoded[2])}
	default:
		return rgb, errors.New("hex string of an unsupported length")

	}
	return rgb, nil
}

func RgbToHsv(rgb RGB) HSV {
	r := (rgb.R / 255)
	g := (rgb.G / 255)
	b := (rgb.B / 255)

	min := math.Min(r, math.Min(g, b))
	max := math.Max(r, math.Max(g, b))
	del := max - min

	v := max

	var h, s float64

	if del == 0 {
		h = 0
		s = 0
	} else {
		s = del / max
		delR := (((max - r) / 6) + (del / 2)) / del
		delG := (((max - g) / 6) + (del / 2)) / del
		delB := (((max - b) / 6) + (del / 2)) / del

		if r == max {
			h = delB - delG
		} else if g == max {
			h = (1.0 / 3.0) + delR - delB
		} else if b == max {
			h = (2.0 / 3.0) + delG - delR
		}

		if h < 0 {
			h += 1
		}
		if h > 1 {
			h -= 1
		}
	}
	hsv := HSV{h, s, v}
	return hsv
}

func HsvToRgb(hsv HSV) RGB {
	var r, g, b float64
	if hsv.S == 0 { //HSV from 0 to 1
		r = hsv.V * 255
		g = hsv.V * 255
		b = hsv.V * 255
	} else {
		h := hsv.H * 6
		if h == 6 {
			h = 0
		} //H must be < 1
		i := math.Floor(h) //Or ... var_i = floor( var_h )
		v1 := hsv.V * (1 - hsv.S)
		v2 := hsv.V * (1 - hsv.S*(h-i))
		v3 := hsv.V * (1 - hsv.S*(1-(h-i)))

		if i == 0 {
			r = hsv.V
			g = v3
			b = v1
		} else if i == 1 {
			r = v2
			g = hsv.V
			b = v1
		} else if i == 2 {
			r = v1
			g = hsv.V
			b = v3
		} else if i == 3 {
			r = v1
			g = v2
			b = hsv.V
		} else if i == 4 {
			r = v3
			g = v1
			b = hsv.V
		} else {
			r = hsv.V
			g = v1
			b = v2
		}

		r = r * 255 //RGB results from 0 to 255
		g = g * 255
		b = b * 255
	}
	rgb := RGB{r, g, b}
	return rgb
}

func ColorForPosition(palette []string, x int, y int, width int, height int) string {
	i := LinearScale(0, width, 0, len(palette)-1, x)
	return palette[i]
}
