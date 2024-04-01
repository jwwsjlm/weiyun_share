package utils

import (
	"math/rand/v2"
)

func GetARandom(l int) int {

	number := 0
	for i := 0; i < l; i++ {
		c := rand.IntN(9) + 1

		number = number*10 + c

	}
	return number
}
