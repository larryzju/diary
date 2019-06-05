package main

import "fmt"

func fib() <-chan int {
    x := make( chan int, 2 )
    go func() {
       for s,b := 0,1;; s,b=b,s+b {
           x <- s
       }
    }()
    return x
}

func main() {
    x := fib()
    for i := 0; i < 10; i++ {
        fmt.Println( <-x )
    }
}
