package roles

import "encoding/json"

type MotionRequest struct {
	Cooters []string `json:"cooters"`
	Document
}

type SamplerRequest struct {
	Color string
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
