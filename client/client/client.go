package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func CallAPI(url string, data []string, name string) ([]string, error) {
	fmt.Printf("Client %s Conectat\n", name)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON: %v", err)
	}

	fmt.Printf("Client %s a facut request cu datele: %v\n", name, data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var output []string
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	fmt.Printf("Client %s a primit raspunsul: %v\n", name, output)

	return output, nil
}

func CallAPIInt(url string, data []int, name string) ([]int, error) {
	fmt.Printf("Client %s Conectat\n", name)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON: %v", err)
	}

	fmt.Printf("Client %s a facut request cu datele: %v\n", name, data)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var output []int
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	fmt.Printf("Client %s a primit raspunsul: %v\n", name, output)

	return output, nil
}