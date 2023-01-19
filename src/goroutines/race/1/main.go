package main

import (
	"fmt"
	"math/rand"
	"time"
)

var result int

func main() {
	go runProcess("P1 ", 10)
	go runProcess("P2 ", 10)

	var s string
	fmt.Scanln(&s)
	fmt.Println("Final result: ", result)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		z := result
		z++
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		result = z
		fmt.Println(name, "-> ", i, " Partial result: ", result)
	}
}
