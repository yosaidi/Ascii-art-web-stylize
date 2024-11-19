package ascii

func AreStringValid(runes []rune) bool {
	for i := 0; i <len(runes)-1; i++ {
		if (runes[i] < 32 || runes[i] > 126) && runes[i] != '\r' &&  runes[i] != '\n' {
			return false
		}
	}
	return true
}
