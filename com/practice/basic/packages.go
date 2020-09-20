package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//initialize global source used by rand.Intn()
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(10))
}
