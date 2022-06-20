package main

import "fmt"

type PetI interface {
	SetName(name string)
	Name() string
	Color() string
}
type Pet struct {
	name  string
	age   int
	color string
}
func New(name string,age int,color string) Pet{
	return Pet{
		name: name,
		age:age,
		color: color,
	}
}
func (pet *Pet) SetName(name string) {
	pet.name = name
}
func (pet Pet) SetNameCopy(name string) {
	pet.name = name
}
func (pet Pet) Name() string {
	return pet.name
}
func (pet Pet) Color() string {
	return pet.color
}
//func (pet Pet) String() string {
//	return fmt.Sprintf("pet stuct value %v\n",pet)
//}
func main() {
	//cat := New("miao", 1, "red")
	  cat := Pet{
		  name: "miao",
		  age:1,
		  color: "red",
	  }
     /**
     cat  包含了2个方法：Name() Color()
     *cat 包含了3个方法：Name() Color() SetName(name string)
      */
	_, ok := interface{}(cat).(PetI)
	fmt.Printf("cat implement PetI %v \n", ok)
	_, ok = interface{}(&cat).(PetI)
	fmt.Printf("cat(point) implement PetI %v \n", ok)
}
