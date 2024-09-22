package main

import (
	"fmt"

	"github.com/gdimit07/greetings"
	"github.com/gdimit07/hello/customer"
)

func main() {
	customer_id := customer.GetCustomerId()
	custom_message := fmt.Sprintf("Your customer ID is: %d", customer_id)
	message := greetings.Hello("Bob", custom_message)
	fmt.Println(message)
}
