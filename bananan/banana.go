package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune
}
