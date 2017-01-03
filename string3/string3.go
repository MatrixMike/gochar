package main

import "fmt"

func main() {
		var a [3]string
	a[0] = "Hello"
	a[1] = "World"
	a[2] = "third"
	fmt.Println("HELLO"[1])
	fmt.Println(string("HELLO"[1]))
	fmt.Println("c%c", []rune("HELLO")[1])
	fmt.Println("%c []rune(\"HELLO\")[1]")
	fmt.Println(string([]rune(a[0])[4]))
}
