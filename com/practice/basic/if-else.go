package main

import (
	"fmt"
	"math"
)

func power(x, n, limit float64) float64 {
	//value is assigned before comparison.
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Println("v was: ", v)
		return limit
	}
}

func main() {
	fmt.Println(power(3, 3, 10))

}
