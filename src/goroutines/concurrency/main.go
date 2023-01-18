package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go runProcess("P1 ", 10)
	go runProcess("P2 ", 10)

	var s string
	fmt.Scanln(&s)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "-> ", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
}
