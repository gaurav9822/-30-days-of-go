package main

import (
	"fmt"
	"math"
)

func pow(x, n, limit float64) float64 {
	//value is assigned before comparison.
	if v := math.Pow(x, n); v < limit {
		return v
	}
	return limit
}

func main() {
	fmt.Println(pow(3, 2, 10))

}
