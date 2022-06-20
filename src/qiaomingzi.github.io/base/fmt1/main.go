package main

import (
	"flag"
	"fmt"
)
var age = flag.Int("age",-1,"请输入姓名")
var name = flag.String("name","name default","请输入名称");
func main() {
	flag.Parse()
	fmt.Printf("hello,%v %v \n",*name,*age);
}