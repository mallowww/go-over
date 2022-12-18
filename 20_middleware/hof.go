package main

import "fmt"

type Decorator func(s string) error

func Use(next Decorator) Decorator {
	return func(c string) error {
		fmt.Println("do something before")
		r := c + ", howdy"
		return next(r)
	}
}

func hello(s string) error {
	fmt.Println("hello", s)
	return nil
}

func main() {
	wrapped := Use(hello)
	w := wrapped("world")
	fmt.Println("end result:", w)
}
