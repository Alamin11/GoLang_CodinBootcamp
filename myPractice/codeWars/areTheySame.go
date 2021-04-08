package main

import (
	"fmt"
	"math"
)

// func Comp(array1 []int, array2 []int) bool {
// 	// your code
// 	flag := false
// 	if len(array1) == 0 || len(array2) == 0 {
// 		return flag
// 	}
// 	for _, val1 := range array1 {
// 		for _, val2 := range array2 {
// 			if float64(val1) == math.Sqrt(float64(val2)) {
// 				flag = true
// 			} else {
// 				flag = false
// 			}
// 		}

// 	}
// 	return flag
// }

func main() {
	var a1 = []int{121, 144, 19, 161, 19, 144, 15, 11}
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(Comp(a1, a2))

}

//using for loop

func Comp(array1 []int, array2 []int) bool {
	flag := false
	l1 := len(array1)
	l2 := len(array2)
	if l1 == 0 || l2 == 0 {
		return flag
	} else {
		for i := 0; i < l1; i++ {
			for j := 0; j <= l2; j++ {
				if math.Sqrt(float64(array2[i])) == float64(array1[j]) {
					flag = true
					break
				}
			}
		}
		return flag
	}

}
