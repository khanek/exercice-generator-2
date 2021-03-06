package math

import (
	"math/rand"
	"time"
)

// Abs returns absolute value
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min returns the lowest numer
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Max returns the greates numer
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// RandomInt returns pseudo random integer from range
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func RandomFloat64(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}
