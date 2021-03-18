package main

import "fmt"

func DateNbDays(a0 float64, a float64, p float64) int {
	// your code
	count := 1
	for i := 1; a0 <= a; i++ {
		a0 += a0 * p / 100 / 360
		if a > a0 {
			count++
		}
	}
	return count
}
func (t Time) AddDate(years int, months int, days int) Time {
	year, month, day := t.Date()
}
func main() {
	days := DateNbDays(100, 101, 0.98)
	fmt.Println(days)
}
