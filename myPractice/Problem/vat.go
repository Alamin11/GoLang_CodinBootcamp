package main

import (
	"fmt"
	"math"
)

func inclusive(taxParcent, value float64) (tax float64) {

	tax = math.Round(value - (value * 100.0 / (100 + taxParcent)))
	return
}

func exclusive(taxPercent, value float64) (tax float64) {
	tax = math.Round(taxPercent * value / 100)
	return
}

func main() {
	exT := exclusive(15.0, 100)
	inT := inclusive(15.0, 100)
	fmt.Printf("Exclusive tax %.2f\nInclusive tax %.2f", exT, inT)
}
