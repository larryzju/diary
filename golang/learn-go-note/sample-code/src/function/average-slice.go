package main

import "fmt"

func average( xs []float64 ) ( avg float64 ) {
	length := float64(len(xs))
	for _, x := range xs {
		avg += x / length
	}
	return
}

func main() {
	fmt.Println( average( []float64{1.1,2.2,3.3,4.4} ) )
}
