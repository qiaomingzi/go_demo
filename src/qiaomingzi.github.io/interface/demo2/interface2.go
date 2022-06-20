package main

import "fmt"

type PetI interface {
	Name() string
}
type Cat struct {
	name string
}
func (cat Cat) Name() string{
	return cat.name
}
func (cat *Cat) SetName(name string){
	cat.name = name
}
func main(){
	cat := Cat{"miao"}
	var pet PetI = cat
	var pet2 PetI = &cat
	cat2 := cat
	cat.SetName("ji ji") // =>  (&cat).SetName("ji ji")
	fmt.Println(cat.Name())
	fmt.Println(pet.Name())
	fmt.Println(pet2.Name())
	fmt.Println(cat2.Name())
}
