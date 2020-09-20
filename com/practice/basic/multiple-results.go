package main

import "fmt"

func main() {
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}

func swap(a string, b string) (string, string) {
	return b, a
}
