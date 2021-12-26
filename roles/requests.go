package roles

import "encoding/json"

type MotionRequest struct {
	Cooters []string `json:"cooters"`
	Document
}

type SamplerRequest struct {
	Color string `json:"color"`
	Size  int    `json:"size"`
	Kind  string `json:"kind"`
	Document
}

type GridRequest struct {
	State         bool   `json:"state"`
	Stroke        string `json:"stroke"`
	StrokeWidth   int    `json:"strokeWidth"`
	TileDimension int    `json:"tileDimension"`
	Document
}

//TranslateMotionRequest
func TranslateMotionRequest(jsmotion string) (MotionRequest, error) {
	mreq := MotionRequest{}
	if err := json.Unmarshal([]byte(jsmotion), &mreq); err != nil {
		return MotionRequest{}, err
	}
	return mreq, nil
}

//TranslateSamplerRequest
func TranslateSamplerRequest(jssampler string) (SamplerRequest, error) {
	sreq := SamplerRequest{}
	if err := json.Unmarshal([]byte(jssampler), &sreq); err != nil {
		return SamplerRequest{}, err
	}
	return sreq, nil
}

//TranslateGridRequest
func TranslateGridRequest(jsgrid string) (GridRequest, error) {
	greq := GridRequest{}
	if err := json.Unmarshal([]byte(jsgrid), &greq); err != nil {
		return GridRequest{}, err
	}
	return greq, nil
}
