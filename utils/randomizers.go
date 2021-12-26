package utils

import (
	"math/rand"
	"time"

	"github.com/paulidealiste/SnapCooter/roles"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var rseed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func RandomName() string {
	q := RandomInt(5, 10)
	b := make([]byte, q)
	for i := range b {
		b[i] = charset[rseed.Intn(len(charset))]
	}
	return string(b)
}

func RandomBearing() string {
	q := RandomInt(0, 7)
	switch {
	case q == 0:
		return roles.E
	case q == 1:
		return roles.N
	case q == 2:
		return roles.NE
	case q == 3:
		return roles.NW
	case q == 4:
		return roles.S
	case q == 5:
		return roles.SE
	case q == 6:
		return roles.SW
	case q == 7:
		return roles.W
	}
	return roles.E
}

func SampleSlice(s []string) string {
	q := RandomInt(0, len(s)-1)
	return s[q]
}

func SimpleProbable(probability float64) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64() < probability
}

func GridSampler(grid PlacementGrid, quantity int) []PlacementCell {
	sample := make([]PlacementCell, quantity)
	copies := make([]PlacementCell, len(grid.Cells))
	copy(copies, grid.Cells)
	for i := 0; i < quantity; i++ {
		ri := RandomInt(0, len(copies)-1)
		sample[i] = copies[ri]
		copies[ri] = copies[len(copies)-1]
		copies = copies[:len(copies)-1]
	}
	return sample
}
