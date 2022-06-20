package main

import (
	"flag"
	"os"
)
import . "qiaomingzi.gihub.io/base/fmt3/lib"
var name = flag.String("name","name default","请输入名称")
var age = flag.Int("age",-1,"请输入年龄")

/*
与main_lib.go中的方法重名了,编译器报错
func Print(w io.Writer)  {
	fmt.Fprint(w,"main->print")
}
*/
func main() {
	flag.Parse()
	SayHello(name,age)
	PrintLib(os.Stdout)
}