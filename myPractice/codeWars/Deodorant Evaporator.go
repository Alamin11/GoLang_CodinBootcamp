package main

import "fmt"

func Evaporator(content float64, evapPerDay int, threshold int) int {
	// your code
	days := 0
	thresholdPercentage := float64(threshold) / 100
	minThresholdValue := content * thresholdPercentage
	evapPercentage := float64(evapPerDay) / 100

	//i:=0;
	for content >= minThresholdValue {
		content = content - content*evapPercentage
		days++
	}
	return days
}
func main() {
	x := Evaporator(10, 10, 10)
	fmt.Println(x)
}
