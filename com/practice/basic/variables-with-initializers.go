package main

import "fmt"

var i, j = 1, 2

func main() {
	//if initializer is present, then no need to give type
	var c, python, java = true, false, "no!"
	fmt.Println(i, c, python, java)
}
