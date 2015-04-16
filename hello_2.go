package main

import "fmt"

type Message struct {
	X string
	Y *string
}

func (v Message) Print () {
	if v.Y != nil {
		fmt.Println(v.X, *v.Y)
	} else {
		fmt.Println(v.X)
	}
}

func (v *Message) Store(x, y string) {
	v.X = x
	v.Y = &y
}

func main() {
	m := &Message{}
	m.Print()
	m.Store("Hello", "world")
	m.Print()
}

