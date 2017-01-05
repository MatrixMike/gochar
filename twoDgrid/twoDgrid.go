package main

import "fmt"

func main() {
	
		var d [3]string  
		d[1] = "1234567890"
				d[0] = "abcdefghi"
	//[]rune( d[1:2]) = "m"
		[]rune(d[0])[3] = []rune(d[1])[1]
			fmt.Println(string([]rune(d[1])[1]))
			fmt.Println(d, "c")
	
	names := [2]string{
		"xxxxxxxxxx",
		"yyyyyyyyyy",
//		"George",
//		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
//	b := names[1:3]
	fmt.Println(a)

//	b[0] = "XXX"
	fmt.Println(a)
	fmt.Println(" ")
	fmt.Println(names)
}
