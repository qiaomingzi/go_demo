package main

import (
	"fmt"
	"io"
)
func SayHello(name *string, age *int) {
	fmt.Println("Printf: %d, %s, %f, %s", 1, "two", 3.0, "over.")
	fmt.Println("hello %v %v",*name,*age)
}
func Print(w io.Writer)  {
	fmt.Fprint(w,"main_lib->print")
}
