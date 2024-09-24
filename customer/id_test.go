package customer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCustomerID(t *testing.T) {

	customers := CustomerIDs{}
	_, err := customers.GenerateCustomerId()

	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(customers.IDs))
}
