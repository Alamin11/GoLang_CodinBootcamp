package main

func splitString(str string) []string {

	result := []string{}

	if str == "" {
		return result
	}

	counter := 0
	tempString := ""
	lastChar := ' '
	for _, c := range str {
		counter++
		tempString += string(c)
		lastChar = c

		if counter == 2 {
			result = append(result, tempString)
			counter = 0
			tempString = ""
		}
	}
	if counter == 1 {
		result = append(result, string(lastChar)+" ")
	}

	return result
}
