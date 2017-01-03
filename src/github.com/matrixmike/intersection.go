package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [4]string
	a[0] = "Hello"
	a[1] = "World"
	a[2] = "resistor"
	a[3] = "capacitor" 
	
	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
	fmt.Println(strings.IndexAny(a[3], a[2])     , "K"    )
	        fmt.Println(string([]rune(a[0])[4]))
	fmt.Println(strings.IndexAny("capacitor", "condensor"))
	fmt.Println(strings.IndexAny("apacitor", "ondensor"))
}
