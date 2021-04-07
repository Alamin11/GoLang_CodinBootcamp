package main

func valiBraces(str string) bool {
	if str == "" {
		return true
	}
	stack := []string{}

	for _, t := range str {
		s := string(t)
		if s == "(" || s == "{" || s == "[" {
			stack = append(stack, s)
		} else {
			if len(stack) == 0 {
				return false
			}
			lastStr := stack[len(stack)-1]

			if s == ")" {
				if lastStr == "(" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}

			}

			if s == "]" {
				if lastStr == "[" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}

			}

			if s == "]" {
				if lastStr == "[" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}

			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {

}
