package roles

import "encoding/json"

type MotionRequest struct {
	Cooters []string `json:"cooters"`
	Document
}

//TranslateMotionRequest
func TranslateMotionRequest(jscooter string) (MotionRequest, error) {
	mreq := MotionRequest{}
	if err := json.Unmarshal([]byte(jscooter), &mreq); err != nil {
		return MotionRequest{}, err
	}
	return mreq, nil
}
