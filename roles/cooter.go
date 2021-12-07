// Pacakge roles provides an implementation of the various drawing roles
package roles

import (
	"encoding/json"
	"regexp"
	"strconv"
)

//Bearing represents the facing of the cooter in a cartesian grid
const (
	N  string = "N"
	NW string = "NW"
	NE string = "NE"
	E  string = "E"
	W  string = "W"
	SE string = "SE"
	SW string = "SW"
	S  string = "S"
)

// Cooter is the main character usedd for all the work
type Cooter struct {
	ID            int     `json:"ID"`
	Name          string  `json:"Name"`
	Size          int     `json:"Size"`
	Bearing       string  `json:"Bearing"`
	X             int     `json:"X"`
	Y             int     `json:"Y"`
	Color         string  `json:"Color"`
	Determination float64 `json:"Determination"`
	Friendliness  float64 `json:"Friendliness"`
	RGB           []int
}

//TranslateCooter creates a cooter instance from its js json definition
func TranslateCooter(jscooter string) (Cooter, error) {
	cooter := Cooter{}
	if err := json.Unmarshal([]byte(jscooter), &cooter); err != nil {
		return Cooter{}, err
	}
	cooter.RawRGB()
	return cooter, nil
}

//ObtainMap returns the map representation of the cooter struct
func (c *Cooter) ObtainMap() map[string]interface{} {
	var mapped map[string]interface{}
	rep, _ := json.Marshal(c)
	json.Unmarshal(rep, &mapped)
	return mapped
}

//ObtatinJSON returns the json representation of the cooter struct
func (c *Cooter) ObtainJSON() string {
	rep, _ := json.Marshal(c)
	return string(rep)
}

//RawRGB sets the R,G,B slice from the rgb color string
func (c *Cooter) RawRGB() {
	re := regexp.MustCompile("[0-9]+")
	rgbs := re.FindAllString(c.Color, -1)
	r, _ := strconv.Atoi(rgbs[0])
	g, _ := strconv.Atoi(rgbs[2])
	b, _ := strconv.Atoi(rgbs[4])
	c.RGB = []int{r, g, b}
}
