package main

import "fmt"

func swap(x, y string) (string, string) {
	fmt.Println("1 ",x," ",y)
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
