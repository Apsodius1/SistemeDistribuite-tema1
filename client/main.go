package main

import (
	"fmt"
	"tema/client"
)

func main() {
	input := []string{"casa", "masa", "trei", "tanc", "4321"}
	clientName := "client1"

	url := "http://localhost:8080/concatByPosition"
	output, err := client.CallAPI(url, input, clientName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Output:", output)
}