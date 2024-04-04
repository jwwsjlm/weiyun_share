package utils

import (
	"math"
	"math/rand/v2"
)

func GetARandom(l int) int {
	maxnum := int(math.Pow10(l) - 1)
	minnum := int(math.Pow10(l - 1))
	randomNumber := rand.IntN(maxnum-minnum+1) + minnum
	return randomNumber
}
