package main

import "fmt"

func main() {
	var a [3]string
	a[0] = "Hello"
	a[1] = "World"
	a[2] = "third"
	fmt.Println("HELLO"[1])
	fmt.Println(string("HELLO"[1]))
	fmt.Println("c%s", []rune("HELLO")[1])
	fmt.Println("%s", "[]rune(\"HELLO\")[1]")
	fmt.Println(string([]rune(a[0])[4]))
}
