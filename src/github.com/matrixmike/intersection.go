package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
	fmt.Println(strings.IndexAny("capacitor", "resistor"))
	fmt.Println(strings.IndexAny("capacitor", "condensor"))
	fmt.Println(strings.IndexAny("apacitor", "ondensor"))
}
