package main

import (
	"fmt"
	"strconv"
	"strings"
)

func NbDig(n int, d int) int {
	// your code
	//if
	flag := 0
	for i := 0; i <= n; i++ {
		flag += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}
	return flag
}

func main() {
	NbDig(25, 1)
	fmt.Println(NbDig(550, 5))
	// 	fmt.Println(strings.Count("cheese", "e"))
	// 	fmt.Println(strings.Count("five", "")) // before & after each rune
	//

}
