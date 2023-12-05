package main

import (
	"fmt"
	"sync"
	"tema/client"
	"tema/config"
)

func makeAPICall(apiEndpoint string, input []string, clientName string, wg *sync.WaitGroup) ([]string, error) {
	defer wg.Done()
	url := "http://localhost:8080/"
	fullURL := url + apiEndpoint

	output, err := client.CallAPI(fullURL, input, clientName)
	if err != nil {
		return nil, fmt.Errorf("Error making API call to %s: %v", apiEndpoint, err)
	}

	fmt.Printf("Output for %s: %v\n\n", apiEndpoint, output)
	return output, nil
}

func makeAPICallInt(apiEndpoint string, input []int, clientName string, wg *sync.WaitGroup) ([]int, error) {
	defer wg.Done()
	url := "http://localhost:8080/"
	fullURL := url + apiEndpoint

	output, err := client.CallAPIInt(fullURL, input, clientName)
	if err != nil {
		return nil, fmt.Errorf("Error making API call to %s: %v", apiEndpoint, err)
	}

	fmt.Printf("Output for %s: %v\n\n", apiEndpoint, output)
	return output, nil
}


func main() {
	var wg sync.WaitGroup
	if err := config.Init(); err != nil {
		fmt.Printf("Error initializing configuration: %v\n", err)
		return
	 }
  
	 clients := config.Clients

	// Exercitiul 1
	// input := []string{"casa", "masa", "trei", "tanc", "4321"}
	// for _, clientName := range clients {
	// 	wg.Add(1)
	// 	go makeAPICall("concatByPosition", input, clientName, &wg)
	// }

	// Exercitiul 2
	input := []string{"abd4g5", "1sdf6fd", "fd2fdsf5"}
	for _, clientName := range clients {
		wg.Add(1)
		go makeAPICall("numarPatratePerfecte", input, clientName, &wg)
	}

	// Exercitiul 3
	// inputInt := []int{12, 13, 14}
	// for _, clientName := range clients {
	// 	wg.Add(1)
	// 	go makeAPICallInt("sumaInverselor", inputInt, clientName, &wg)
	// }


	// Exercitiul 5
	// input := []string{"2dasdas", "12", "dasdas", "1010", "101"}
	// for _, clientName := range clients {
	// 	wg.Add(1)
	// 	go makeAPICall("conversieNumereBinare", input, clientName, &wg)
	// }

	// Exercitiul 8
	// inputInt = []int{23, 17, 15, 3, 18}
	// for _, clientName := range clients {
	// 	wg.Add(1)
	// 	go makeAPICallInt("numarCifreNumerelorPrime", inputInt, clientName, &wg)
	// }

	// Exercitiul 9
	// input := []string{"mama", "iris", "bunica", "ala"}
	// for _, clientName := range clients {
	// 	wg.Add(1)
	// 	go makeAPICall("numarCuvinteVocalePare", input, clientName, &wg)
	// }

	// Exercitiul 10
	input = []string{"24", "16", "8", "aaa", "bbb"}
	for _, clientName := range clients {
		wg.Add(1)
		go makeAPICall("cmmdcNumere", input, clientName, &wg)
	}

	// Exercitiul 12
	// inputInt = []int{23, 43, 26, 74}
	// for _, clientName := range clients {
	// 	wg.Add(1)
	// 	go makeAPICallInt("sumaDublarePrimaCifra", inputInt, clientName, &wg)
	// }

	// Exercitiul 14
	// input = []string{"Ceva12!@", "asa212", "dasdas"}
	// for _, clientName := range clients {
	// 	wg.Add(1)
	// 	go makeAPICall("filtreazaParole", input, clientName, &wg)
	// }

	// Exercitiul 15
	// input = []string{"a", "b", "e", "3", "!", "A"}
	// for _, clientName := range clients {
	// 	wg.Add(1)
	// 	go makeAPICall("genereazaParole", input, clientName, &wg)
	// }

	
	
	wg.Wait()
}