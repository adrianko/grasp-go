package main

import (
	"fmt"
	"time"
	"math"
	"container/list"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x 
	return
}

func loops() {
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
}

func arrayEx() {
	arr := []string{"abc", "def"}
	fmt.Println(arr)
	arr = append(arr, "ghi", "abc")
	fmt.Println(arr)
}

func listEx() {
	list := list.New()
	list.PushBack(4)
	list.PushBack(1)
	list.PushFront(7)
	
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func randInt(max int) (int) {
	rand.Seed(time.Now().UTC().UnixNano())
	
	return rand.Intn(max)
}

func ifScopedVars() (float64) {
	if v:= math.Pow(5, 2); v < 20 {
		return v
	}
	
	return 20
}

func deferred() {
	defer fmt.Println("world")
	
	fmt.Println("hello")
}

func deferredStack() {
	fmt.Println("\n")
	fmt.Println("counting")
	
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	
	fmt.Println("done")
}

func main() {
	fmt.Println("Hello, world")
	/*
	Go basic types
	bool
	string
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintpr
	byte (alias for uint8)
	rune (alias for int32)
	float32 float64
	complex64 complex128
	*/
	i := 3 //implicit type
	var m int = 2 //declared type
	var n int = int(4.0)
	const Hello = "World"
	loops()
	
	fmt.Println(time.Now())
	
	arrayEx()
	listEx()
	
	fmt.Println("Random numbers: ")
	
	for i := 0; i < 5; i++ {
		fmt.Println(randInt(100))
	}
	
	fmt.Println(add(i, m))
	hello := "hello"
	world := "world"
	hello, world = swap(hello, world)
	fmt.Println("hello: " + hello)
	fmt.Println("world: " + world)
	
	fmt.Println(split(n))
	fmt.Println(ifScopedVars())
	deferred()
	deferredStack()
}
