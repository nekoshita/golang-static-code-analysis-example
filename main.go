package main

import "log"

func main() {
	log.Print("hello")

	if _, err := foo(); err != nil {
		panic(err)
	}

	hogeYey := HogeYey{
		Foo: "yey",
	}
	_ = hogeYey
}

type HogeYey struct {
	Foo string
}

func foo() (string, error) {
	return "", nil
}
