package main

import "fmt"

func main() {
	fmt.Println("hello world")

	counter := 0

	for {
		counter++
		fmt.Println("counter:", counter)

		// delay the loop for 1 second
	}

}
