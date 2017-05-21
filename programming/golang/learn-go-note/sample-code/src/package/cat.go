package main

import(
  "fmt"
  "flag"
)

func main() {
  ip := flag.Int( "s", 10, "size of input" )
  flag.Parse()
  fmt.Printf( "%T %v\n", ip, *ip )
}


