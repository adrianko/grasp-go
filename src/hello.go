package main

import (
	"fmt"
	"time"
	"container/list"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

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
	
	list := list.New()
	list.PushBack(4)
	list.PushBack(1)
	list.PushFront(7)
	
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	
	fmt.Println("Random numbers: ")
	
	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		fmt.Println(rand.Intn(100))
	}
	
	fmt.Println(add(3, 4))
}
