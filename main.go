package main

import (
	"fmt"

	go_say_hello "github.com/hakimfauzi23/go-say-hello"
)

func main() {
	fmt.Println(go_say_hello.SayHello())
	fmt.Println(go_say_hello.SayHelloTo("Hakim"))
}
