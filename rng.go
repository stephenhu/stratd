package main

import (
	"math/rand"
	"time"
)

const (
	ATTRIBUTE_MAX					= 1000
	ATTRIBUTE_MIN					= 50
	ATTRIBUTE_DEFAULT     = 100
	ATTRIBUTE_SIGMA       = 200.0
	ATTRIBUTE_MEAN        = 375.0
)

var rng *rand.Rand


func initRng() {

	seed := rand.NewSource(time.Now().UnixNano())

	rng = rand.New(seed)

} // initRng


func attrStandard(sigma float64, mean float64) int {

	mu := int(rng.NormFloat64() * sigma + mean)

	if mu < 0 {
		return ATTRIBUTE_MIN
	} else if mu > ATTRIBUTE_MAX {
		return ATTRIBUTE_MAX
	} else {
		return mu
	}

} // attrStandard


func attrRange(min int, max int) int {

	n := rng.Intn(max)

	if n < min {

		nn := min + n

		if nn > max {
			return max
		} else {
			return nn
		}

	} else {
		return n
	}

} // attrRange


func attrSelect(attrs ...int) int {

	i := rng.Intn(len(attrs))

	return attrs[i]

} // attrSelect
