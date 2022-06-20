package main

import "fmt"

type Pet interface {
	Name() string
}
type Cat struct{
	name string
}
func New(name string) Cat{
	return Cat{name}
}
func (cat Cat) Name() string{
	return cat.name
}
func main(){
	var cat *Cat
	cat1 := cat
	fmt.Println("cat is nil")
	fmt.Println("cat1 is nil")
	//iface的实例包装cat1的值的副本,虽然被包装的动态值是`nil`，但是`pet`的值却不会是`nil`，因为这个动态值只是`pet`值的一部分而已
	var pet Pet = cat1
	if pet == nil {
		fmt.Println("pet is nil")
	} else {
		fmt.Println("pet is not nil")
	}
}


