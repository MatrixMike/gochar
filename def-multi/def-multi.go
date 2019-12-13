package main
// 13/dec/2019
import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
//	for i := 0; i < 10; i++ {
//	        fmt.Println(i)
//	}
	fmt.Println("done")
}
