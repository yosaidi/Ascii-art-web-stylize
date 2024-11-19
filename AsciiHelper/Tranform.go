package ascii

import (
	"fmt"
	"strings"
)

func Transform(inp, temp string) (string, error) {
	template := temp + ".txt"
	input := inp
	input = strings.ReplaceAll(input, "\r\n", "\n")

	runes := []rune(input)
	isValid := AreStringValid(runes)
	if !isValid {
		return "", fmt.Errorf("invalid character in your input")
	}
	myString := ReadFile(template)
	if myString == nil {
		return "", fmt.Errorf("internal server error")
	}
	lines := strings.Split(input, "\n")
	result := ""
	for i, line := range lines {

		if line == "" {
			if i == len(lines)-1 {
				continue
			}
			result += "\n"
			continue
		}
		newInput := SpaceManager(line)

		for row := 1; row < 9; row++ {
			for _, word := range newInput {
				wordRunes := []rune(word)
				for j := 0; j < len(wordRunes); j++ {
					char := wordRunes[j]
					starts := (int(char) - 32) * 9
					result += myString[starts+row]
				}

			}
			result += "\n"
		}

	}
	return result, nil
}
