// Pacakge roles provides an implementation of the various drawing roles
package roles

import (
	"encoding/json"
)

//Bearing represents the facing of the cooter in a cartesian grid
type Bearing string

const (
	N  Bearing = "N"
	NW Bearing = "NW"
	NE Bearing = "NE"
	E  Bearing = "E"
	W  Bearing = "W"
	SE Bearing = "SE"
	SW Bearing = "Sw"
	S  Bearing = "S"
)

// Cooter is the main character usedd for all the work
type Cooter struct {
	ID            int     `json:"ID"`
	Name          string  `json:"Name"`
	Bearing       Bearing `json:"Bearing"`
	X             int     `json:"X"`
	Y             int     `json:"Y"`
	Color         string  `json:"Color"`
	Determination float32 `json:"Determination"`
	Friendliness  float32 `json:"Friendliness"`
}

//TranslateCooter creates a cooter instance from its js json definition
func TranslateCooter(jscooter string) (Cooter, error) {
	cooter := Cooter{}
	if err := json.Unmarshal([]byte(jscooter), &cooter); err != nil {
		return Cooter{}, err
	}
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
