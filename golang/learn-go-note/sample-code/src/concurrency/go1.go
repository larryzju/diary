package main

import "fmt"

func main() {
   c := make(chan int)
   o := make(chan bool)
   defer close(c)
   defer close(o)
   go shower( c, o )
   for i := 0; i < 100; i++ {
      c <- i
   }
   o <- true
}

func shower( c chan int, over chan bool ) {
   for {
       select {
       case j := <-c:
	   fmt.Println( j )
       case <-over:
            break
       }
   }
}
