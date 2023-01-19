package main

import "fmt"

func main() {

	c := generate(30, 200)

	//fan out
	d1 := divide(c)
	d2 := divide(c)
	//fan in

	fmt.Println("Channel d1", <-d1)
	fmt.Println("Channel d2", <-d2)

}

func generate(numbers ...int) chan int {
	channel := make(chan int)
	go func() {
		for _, n := range numbers {
			channel <- n
		}
		close(channel)
	}()
	return channel
}
func divide(input chan int) chan int {
	channel := make(chan int)

	go func() {
		for number := range input {
			channel <- number / 2
		}
		close(channel)
	}()
	return channel
}
