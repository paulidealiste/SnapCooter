// Package utils provide cooter assemblies utility functions
package utils

import (
	"math"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var rseed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomName() string {
	q := RandomInt(5, 10)
	b := make([]byte, q)
	for i := range b {
		b[i] = charset[rseed.Intn(len(charset))]
	}
	return string(b)
}

func LinearScale(pmin int, pmax int, cmin int, cmax int, x int) int {
	fpmin := float64(pmin)
	fpmax := float64(pmax)
	fcmin := float64(cmin)
	fcmax := float64(cmax)
	fx := float64(x)
	i := int(math.Round(((fx - fpmin) / (fpmax - fpmin)) * ((fcmax - fcmin) + fcmin)))
	return i
}
