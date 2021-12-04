package roles

import "encoding/json"

type Document struct {
	CanvasID string `json:"canvasID"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
}

type Setup struct {
	CooterCount int      `json:"cooterCount"`
	Palette     []string `json:"palette"`
	Document
}

//TranslateSetup returns the operative instance of the setup object
func TranslateSetup(jscooter string) (Setup, error) {
	s := Setup{}
	if err := json.Unmarshal([]byte(jscooter), &s); err != nil {
		return Setup{}, err
	}
	return s, nil
}
