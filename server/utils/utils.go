package utils

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