package main

import "log"

func main() {
	log.Print(Hoge{})
}

type Hoge struct {
	Foo string
}
