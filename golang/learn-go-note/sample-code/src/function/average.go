package main

import `fmt`

func main() {
	s := []float64{1.2,2.3,3.4,4.5,5.6,6.7,7.8,8.9,9.0}
	sum := float64(0.0)
	for _, v := range s {
		sum += v
	}
	fmt.Println( sum / float64(len(s)) )
}


