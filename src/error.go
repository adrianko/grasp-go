package main

import (
    "time"
    "fmt"
)

type Error struct {
	When time.Time
	What string
}

func (e *Error) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func err() error {
	 return &Error{time.Now(), "Broken",}
}

func main() {
    fmt.Println(err())
}