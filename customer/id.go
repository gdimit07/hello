package customer

import "math/rand"

func GetCustomerId() int {
	return rand.Int()
}
