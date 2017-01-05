package main

import (
	"fmt"
	"strings"
)

func swap(x, y string) (string, string) {
	fmt.Println("1 ", x, " ", y)
	return y, x
}
func main() {
	var a [4]string
	a[0] = "Hello"
	a[1] = "World"

	a[2] = "resistor"
	a[3] = "capacitor"
	x := 0
	//x := 1
	//	var i interface{}
	//	i = 42
	//describe(i)
	x = x + 1

	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))

	fmt.Println(strings.IndexAny(a[3], a[2]), "K")
	x = strings.IndexAny(a[3], a[2])

	fmt.Println(a[2:4])
	fmt.Print("Index of intersecting character ", x, " ")
	fmt.Println(string([]rune(a[3])[x]))
	fmt.Println("you will find character ", string([]rune(a[3])[x]), " at ", x, " in ", a[3])
	a[2] = "transistor"
	a[3] = "capacitor"
	fmt.Println(strings.IndexAny(a[3], a[2]), "K")
	x = strings.IndexAny(a[3], a[2])

	fmt.Println(a[2:4])
	fmt.Print("Index of intersecting character ", x, " ")
	fmt.Println(string([]rune(a[3])[x]))
	fmt.Println("you will find character ", string([]rune(a[3])[x]), " at ", x, " in ", a[3])

	fmt.Println(strings.IndexAny("capacitor", "condensor"))
	fmt.Println(strings.IndexAny("apacitor", "ondensor"))
}
