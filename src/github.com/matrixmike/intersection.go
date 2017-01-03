package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [4]string
	a[0] = "Hello"
	a[1] = "World"
	a[2] = "third"
	a[3] = "capacitor" 
	
	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
	fmt.Println(strings.IndexAny(a[3], "resistor")         )
	fmt.Println(strings.IndexAny("capacitor", "condensor"))
	fmt.Println(strings.IndexAny("apacitor", "ondensor"))
}
