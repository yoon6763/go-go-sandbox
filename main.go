package main

import "fmt"

func main() {
	n, _ := fmt.Println("Hello, World!", 42, true)
	fmt.Println(n)

	var aa = 3
	bb := 4
	fmt.Println(aa, bb)

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
