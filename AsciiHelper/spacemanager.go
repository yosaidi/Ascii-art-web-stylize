package ascii

func SpaceManager(s string) []string {
	result := []string{}
	helper := ""
	for _, ch := range s {
		if ch == ' ' {
			helper += " "
			result = append(result, helper)
			helper = ""
		} else {
			helper += string(ch)
		}
	}
	result = append(result, helper)
	return result
}
