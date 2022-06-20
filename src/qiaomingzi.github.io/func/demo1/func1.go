package main

import "fmt"

type Printer func(param string) (n int,err error)

func PrinterStr(name string) (count int,err error)  {
	return fmt.Println("here is %s",name)
}
func main(){
	var f Printer
	f = PrinterStr
	f("limz")
}
