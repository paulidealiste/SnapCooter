// Package collections provides the js style collection functions for the cooter slices
package collections

import "github.com/paulidealiste/SnapCooter/roles"

func Map(cooters []roles.Cooter, f func(roles.Cooter) roles.Cooter) []roles.Cooter {
	vsm := make([]roles.Cooter, len(cooters))
	for i, v := range cooters {
		vsm[i] = f(v)
	}
	return vsm
}
