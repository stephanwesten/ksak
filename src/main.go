package main

import (
	"fmt"
	"math/rand"
	"time"
)

// print every second the current time and a random character
func main() {
	for {
		fmt.Print(time.Now())
		fmt.Println(" " + string(rand.Intn(32)+32))
		time.Sleep(time.Second)
	}
}
