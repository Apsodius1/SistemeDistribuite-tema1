package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func NumarPatratePerfecteHandler(w http.ResponseWriter, r *http.Request) {
	var input []string

	fmt.Println("Server a primit request de la client")
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Serrver proceseasza datele primite de la client")
	numarTotalPatratePerfecte := utils.NumarPatratePerfecte(input)

	response := []string{fmt.Sprintf("%d", numarTotalPatratePerfecte)}


	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	
	fmt.Println("Server trimite ", response, " catre client")

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func ConversieNumereBinareHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server a primit request de la client")

	var input []string
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Serrver proceseasza datele primite de la client")

	rezultat := utils.ConversieNumereBinare(input)

	response := make([]string, 0)
	if len(rezultat) > 0 {
		response = append(response, fmt.Sprintf("%d", rezultat[0]))
	}

	for _, val := range rezultat[1:] {
		response = append(response, fmt.Sprintf("%d", val))
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

func NumarCuvinteVocalePareHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server a primit request de la client")

	var input []string
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	fmt.Println("Serrver proceseasza datele primite de la client")

	count := utils.NumarCuvinteVocalePare(input)
	
	response := []string{fmt.Sprintf("%d", count)}

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

	// Prepare the response array with the sum
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