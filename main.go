package main

import (
	"fmt"
	"log"

	"github.com/gdimit07/greetings"
	"github.com/gdimit07/hello/customer"
)

func main() {
	logger := log.Default()
	customer_id, err := customer.GetCustomerId()
	if err != nil {
		logger.Fatal(err)
	}

	custom_message := fmt.Sprintf("Your customer ID is: %d", customer_id)
	message, err := greetings.Hello("Bob", custom_message)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println(message)
}
