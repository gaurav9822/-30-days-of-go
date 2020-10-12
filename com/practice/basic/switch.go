package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("windows")
	}
}
