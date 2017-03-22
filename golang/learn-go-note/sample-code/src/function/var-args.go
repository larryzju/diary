package main

import "fmt"

func VarArgsTest( xs ...int ) {
	for _, x := range xs {
		fmt.Println(x)
	}
}

func main() {
	VarArgsTest( 1,2,34,5,6,7,8 )
}
