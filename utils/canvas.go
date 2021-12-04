package utils

import (
	"syscall/js"

	"github.com/paulidealiste/SnapCooter/roles"
)

func GetCanvasSetup(cas string) (js.Value, roles.Setup, error) {
	document := js.Global().Get("document")
	if !document.Truthy() {
		return js.Value{}, roles.Setup{}, js.Error{}
	}
	setup, err := roles.TranslateSetup(cas)
	if err != nil {
		return js.Value{}, roles.Setup{}, js.Error{}
	}
	canvas := document.Call("getElementById", setup.Document.CanvasID)
	if !canvas.Truthy() {
		return js.Value{}, roles.Setup{}, js.Error{}
	}
	return canvas, setup, nil
}

func GetCanvas(cid string) (js.Value, error) {
	document := js.Global().Get("document")
	if !document.Truthy() {
		return js.Value{}, js.Error{}
	}
	canvas := document.Call("getElementById", cid)
	if !canvas.Truthy() {
		return js.Value{}, js.Error{}
	}
	return canvas, nil
}
