package main

import "fmt"

type  Cat struct{
	name string
	age int
}
func New(name string,age int) Cat {
	return Cat{
		name,age,
	}
}
func (cat Cat) SetName(name string)  {
	cat.name = name
}
func (cat *Cat) SetNamePoint(name string)  {
	cat.name = name
}
func (cat Cat) String()  {
	fmt.Printf("cat:%s %d",cat.name,cat.age)
}

func main()  {
   cat := New("qiu qiu",10)
   fmt.Println(cat);
   cat.SetName("hua hua")
   fmt.Println(cat);
	cat.SetNamePoint("hua hua")
	fmt.Println(cat);
}



