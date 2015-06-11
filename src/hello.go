package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, world")
	
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	
	j := 1
	
	fmt.Println("while")
	
	for j < 5 {
		j += j
		fmt.Println(j)
	}
	
	fmt.Println(j)
	fmt.Println(time.Now())
	
	arr := []string{"abc", "def"}
	fmt.Println(arr)
	arr = append(arr, "ghi", "abc")
	fmt.Println(arr)
}
