package customer

import (
	"errors"
	"math/rand"
)

func GetCustomerId() (int, error) {
	new_id := rand.Int()
	if new_id%3 == 0 {
		return -1, errors.New("could not generate a new id")
	}
	return new_id, nil
}
