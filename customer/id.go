package customer

import (
	"errors"
	"math/rand"
	"slices"
)

type CustomerIDs struct {
	IDs []int
}

func (customers *CustomerIDs) GenerateCustomerId() (int, error) {
	new_id := rand.Int()
	if new_id%11 == 0 {
		return -1, errors.New("could not generate a new id")
	}

	if slices.Contains(customers.IDs, new_id) {
		return -1, errors.New("customer id already exists, there was a conflict")
	}

	customers.IDs = append(customers.IDs, new_id)

	return new_id, nil
}
