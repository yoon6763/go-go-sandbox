package main

import "fmt"

func main() {
	initState()

	for i := 0; i < 100; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}

func initState() {
	fmt.Println("Init")
}
