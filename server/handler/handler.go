package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"tema1/utils"
)

type StringProcessor func([]string) []string

func GenericHandler(processor StringProcessor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Server a primit request de la client")
		var input []string
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		
		fmt.Println("Serrver proceseasza datele primite de la client")

		output := processor(input)

		response, err := json.Marshal(output)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		fmt.Println("Server trimite ", response, " catre client")

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func processInput(input string, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	numarPatratePerfecte := utils.NumarPatratePerfecte([]string{input})

	resultChan <- numarPatratePerfecte
}

func NumarPatratePerfecteHandler(w http.ResponseWriter, r *http.Request) {
	var input []string

	fmt.Println("Server a primit request de la client")
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Server proceseaza datele primite de la client")

	var wg sync.WaitGroup

	resultChan := make(chan int, len(input))

	for _, data := range input {
		wg.Add(1)
		go processInput(data, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	totalPatratePerfecte := 0
	for result := range resultChan {
		totalPatratePerfecte += result
	}

	response := []string{fmt.Sprintf("%d", totalPatratePerfecte)}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Server trimite ", response, " catre client")

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func processBinaryConversion(input string, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	conversionResult := utils.ConversieNumereBinare([]string{input})

	if len(conversionResult) > 0 {
		resultChan <- conversionResult[0]
	}
}

func ConversieNumereBinareHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server a primit request de la client")

	var input []string
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Server proceseaza datele primite de la client")

	var wg sync.WaitGroup

	resultChan := make(chan int, len(input))

	for _, data := range input {
		wg.Add(1)
		go processBinaryConversion(data, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	conversionResults := make([]int, 0)
	for result := range resultChan {
		conversionResults = append(conversionResults, result)
	}

	response := make([]string, len(conversionResults))
	if len(conversionResults) > 0 {
		response[0] = fmt.Sprintf("%d", conversionResults[0])
	}

	for i, val := range conversionResults[1:] {
		response[i+1] = fmt.Sprintf("%d", val)
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Server trimite ", response, " catre client")

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func processNumarCuvinteVocalePare(input string, wg *sync.WaitGroup, resultChan chan int) {
	defer wg.Done()

	numarCuvinteVocalePare := utils.NumarCuvinteVocalePare([]string{input})

	resultChan <- numarCuvinteVocalePare
}

func NumarCuvinteVocalePareHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server a primit request de la client")

	var input []string
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Server proceseaza datele primite de la client")

	var wg sync.WaitGroup

	resultChan := make(chan int, len(input))

	for _, data := range input {
		wg.Add(1)
		go processNumarCuvinteVocalePare(data, &wg, resultChan)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	totalNumarCuvinteVocalePare := 0
	for result := range resultChan {
		totalNumarCuvinteVocalePare += result
	}

	response := []string{fmt.Sprintf("%d", totalNumarCuvinteVocalePare)}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Server trimite ", response, " catre client")

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func CMMDCNumereHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Server a primit request de la client")

	var input []string
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Serrver proceseasza datele primite de la client")

	cmmdc := utils.CMMDCNumere(input)

	response := []string{fmt.Sprintf("%d", cmmdc)}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Server trimite ", response, " catre client")

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func FiltreazaParoleHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Server a primit request de la client")
	var input []string
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Serrver proceseasza datele primite de la client")

	validPasswords := utils.FiltreazaParole(input)

	responseJSON, err := json.Marshal(validPasswords)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Server trimite ", responseJSON, " catre client")

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func GenereazaParoleHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Server a primit request de la client")
	var input []string
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Serrver proceseasza datele primite de la client")
	passwords := utils.GenereazaParole(input)

	responseJSON, err := json.Marshal(passwords)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Server trimite ", responseJSON, " catre client")

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func SumaInverselorHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Server a primit request de la client")

	var input []int
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Server proceseaza datele primite de la client")

	suma := utils.SumaInverselor(input)

	response := []int{suma}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Server trimite ", response, " catre client")

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func NumarCifreNumerelorPrimeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Server a primit request de la client")

	var input []int
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Server proceseaza datele primite de la client")

	numarCifre := utils.NumarCifreNumerelorPrime(input)

	response := []int{numarCifre}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Server trimite ", response, " catre client")

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}

func SumaDublarePrimaCifraHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Server a primit request de la client")

	var input []int
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Server proceseaza datele primite de la client")

	suma := utils.SumaDublarePrimaCifra(input)

	response := []int{suma}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println("Server trimite ", response, " catre client")

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}