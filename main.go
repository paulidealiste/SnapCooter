package main

import (
	"syscall/js"

	"github.com/paulidealiste/SnapCooter/engine"
	"github.com/paulidealiste/SnapCooter/motion"
	"github.com/paulidealiste/SnapCooter/utils"
)

func main() {
	js.Global().Set("CooterSetup", js.FuncOf(engine.CooterSetup))
	js.Global().Set("CooterStep", js.FuncOf(motion.CooterStep))
	js.Global().Set("CooterInterpolatedPalettes", js.FuncOf(utils.CooterInterpolatedPalettes))
	js.Global().Set("DrawSampler", js.FuncOf(engine.DrawSampler))
	js.Global().Set("ManageGrid", js.FuncOf(engine.ManageGrid))
	<-make(chan bool)
}
