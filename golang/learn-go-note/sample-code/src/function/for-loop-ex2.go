package main

import "fmt"

func main() {
	i := 0
	Repeat:
	fmt.Println( i )
	i++
	if i < 10 {
		goto Repeat
	}

}
