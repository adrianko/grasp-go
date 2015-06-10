package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
	
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	
	j := 1
	
	for j < 5 {
		j += j
	}
	
	fmt.Println(j)
}
