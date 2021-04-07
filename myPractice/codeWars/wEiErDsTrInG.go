package main

import "strings"

func toWeirdCase(str string) string {
	if str = "" {
		return ""
	}
	result := ""
	for _, value := range str {
		v := string(value)
		if v == " " {
			result += v
			index = 0
			continue
		}
		if index%2 == 0 {
			result += strings.ToUpper(v)
		} else {
			result += strings.ToLower(v)
		}
		index++
	}
	return result
}
func main() {

}
