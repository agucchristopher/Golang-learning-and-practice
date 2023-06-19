package main

import "fmt"


func main() {
 // Types
 // string
 // bool
 // int
 // int int8 int16 int32 int64
 // uint uint8 uint16 uint32 uint64
 // byte - alias for int32
 //  float32 float64
//  complex64 complex128

// Shorthand
name := "Chris"
size := 1.3

// Using var
var age int32 = 37
var isCool = true
isCool = false
fmt.Println(name, age, isCool)

// To get var type
fmt.Printf("%T\n", size)


}