package main

func Multiple3And5(number int) int {
	sum := 0
	for i := 3; i <= number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	return sum
}
func main() {

}
