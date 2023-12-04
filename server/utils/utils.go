package utils

import (
	"math"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"unicode"
)

func ConcatByPosition(input []string) []string {
	output := make([]string, len(input[0]))

	for i := 0; i < len(input[0]); i++ {
		var builder string
		for _, str := range input {
			builder += string(str[i])
		}
		output[i] = builder
	}

	return output
}

func NumarPatratePerfecte(input []string) int {
	numarTotalPatratePerfecte := 0

	for _, str := range input {
		for _, char := range str {
			if unicode.IsDigit(char) {
				digit := int(char - '0')
				if isPatratPerfect(digit) {
					numarTotalPatratePerfecte++
				}
			}
		}
	}

	return numarTotalPatratePerfecte
}

func isPatratPerfect(numar int) bool {
	radical := int(math.Sqrt(float64(numar)))
	return radical*radical == numar
}

func ConversieNumereBinare(input []string) []int {
	var rezultat []int

	for _, str := range input {
		_, err := strconv.ParseInt(str, 2, 0)
		if err == nil {
			decimalValue, _ := strconv.ParseInt(str, 2, 64)
			rezultat = append(rezultat, int(decimalValue))
		}
	}

	return rezultat
}

func NumarCuvinteVocalePare(input []string) int {
	count := 0

	for _, word := range input {
		if isNumarVocalePare(word) {
			count++
		}
	}

	return count
}

func isNumarVocalePare(word string) bool {
	vowelCount := 0

	for i, char := range word {
		if i%2 == 0 && isVowel(char) {
			vowelCount++
		}
	}

	return (vowelCount%2 == 0 && vowelCount != 0)
}

func isVowel(char rune) bool {
	return unicode.In(char, unicode.Latin, unicode.Scripts["Greek"], unicode.Scripts["Cyrillic"]) && (char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U')
}

func CMMDCNumere(input []string) int {
	var numbers []int

	for _, str := range input {
		if num, err := strconv.Atoi(str); err == nil {
			numbers = append(numbers, num)
		}
	}

	if len(numbers) == 0 {
		return 0
	}

	divisors := listDivisors(numbers[0])

	for _, num := range numbers[1:] {
		divisors = intersectLists(divisors, listDivisors(num))
	}

	return maxInList(divisors)
}

func listDivisors(num int) []int {
	divisors := []int{1, num}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			divisors = append(divisors, i, num/i)
		}
	}

	sort.Ints(divisors)
	return divisors
}

func intersectLists(list1, list2 []int) []int {
	var result []int

	i, j := 0, 0
	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			i++
		} else if list1[i] > list2[j] {
			j++
		} else {
			result = append(result, list1[i])
			i++
			j++
		}
	}

	return result
}

func maxInList(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[len(list)-1]
}

func FiltreazaParole(input []string) []string {
	var validPasswords []string
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	digitRegex := regexp.MustCompile(`[0-9]`)
	symbolRegex := regexp.MustCompile(`[!@#$%^&*()-_+=\[\]{}|;:'",.<>/?]`)

	for _, str := range input {
		if lowercaseRegex.MatchString(str) &&
			uppercaseRegex.MatchString(str) &&
			digitRegex.MatchString(str) &&
			symbolRegex.MatchString(str) {
			validPasswords = append(validPasswords, str)
		}
	}

	return validPasswords
}

func GenereazaParole(input []string) []string {

	var passwords []string

	N := rand.Intn(10) + 1

	for i := 0; i < N; i++ {
		M := rand.Intn(11) + 5

		password := ""
		for j := 0; j < M; j++ {
			char := input[rand.Intn(len(input))]

			password += char
		}

		passwords = append(passwords, password)
	}

	return passwords
}

func SumaInverselor(input []int) int {
	sum := 0

	for _, num := range input {
		reversed := reverseNumber(num)
		sum += reversed
	}

	return sum
}

func reverseNumber(num int) int {
	reversed := 0

	for num != 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}

	return reversed
}

func NumarCifreNumerelorPrime(input []int) int {
	var totalDigits int

	for _, num := range input {
		if isPrime(num) {
			totalDigits += countDigits(num)
		}
	}

	return totalDigits
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func countDigits(num int) int {
	digitCount := 0
	for num != 0 {
		num /= 10
		digitCount++
	}
	return digitCount
}

func SumaDublarePrimaCifra(input []int) int {
	var sum int

	for _, num := range input {
		firstDigit := num / 10
		modifiedNum := firstDigit*100 + num
		sum += modifiedNum
	}

	return sum
}