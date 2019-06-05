package main

import ( "fmt" 
"reflect" )

type Person struct{
    name string "namestr"
    age  int
}

func main() {
    p := &Person{"zhao wenbin", 28}
    t := reflect.TypeOf(p)
    v := reflect.ValueOf(p)
    fmt.Println( t.Elem().Field(0).Tag )
    fmt.Println( v.Elem().Field(0).String() )
}
