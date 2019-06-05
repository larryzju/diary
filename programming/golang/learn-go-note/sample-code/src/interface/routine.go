package main

import (
   "fmt"
   "time"
)

func ready( c chan string, w string, sec int ) {
    time.Sleep( time.Duration(sec) * time.Second )
    fmt.Println( w, "is ready" )
    c <- w
}

func main() {
    c := make(chan string)
    go ready( c, "2sec", 2 )
    go ready( c, "1sec", 1 )
    fmt.Println( "I'm waiting" )
    fmt.Println( <-c )
    fmt.Println( <-c )
    fmt.Println( "over" )
}

