package main

import (
	"fmt"
	"time"
)

const (
	WIDTH  = 100
	HEIGHT = 100
)

func main() {
	h := time.Now().Hour()

	fmt.Println(digits[h])
}
