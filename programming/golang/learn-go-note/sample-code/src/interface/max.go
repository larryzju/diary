package main

import "fmt"

type Numeric interface{}

func Less( l, r Numeric ) bool {
    switch l.(type) {
    case int:
        if _, ok := r.(int); ok {
           return l.(int) < r.(int)
        }
    case float32:
	if _, ok := r.(float32); ok {
           return l.(float32) < r.(float32)
        }
    }
    return false
}


func main() {
    fmt.Println( Less(10,20) )
    fmt.Println( Less(10.0,20.0) )
}

