package main
// 29.08.2019 15:49:19
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println("Galapagos")
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune
	fmt.Println(strings.Count("banana", "an"))
}
