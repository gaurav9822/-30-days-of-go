package main

import "fmt"

const (
	//Overflows int
	Big = 1 << 100

	Small = Big >> 99
)

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//An untyped constant takes the type needed by its context.
	fmt.Println(needFloat(Big))
	fmt.Println(Small)
}
