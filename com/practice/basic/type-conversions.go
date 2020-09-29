package main

import (
	"fmt"
)

func main() {
	var i int = 42
	//Requires explicit type conversion.
	var f float64 = float64(i)
	fmt.Println(f)
}
