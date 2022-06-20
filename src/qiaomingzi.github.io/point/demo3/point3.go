package main

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}
func (dog Dog) Name() string {
	return dog.name
}
func main() {
	dog := Dog{"little dog"}
	dogPoint := &dog
	uintprtVal := uintptr(unsafe.Pointer(dogPoint))
	fmt.Println("&dog=", dogPoint)
	fmt.Println("uintptr=", uintprtVal)

	dogNamePtr := uintprtVal + unsafe.Offsetof(dogPoint.name)
	dogNameP := (*string)(unsafe.Pointer(dogNamePtr)) //转string指针
	fmt.Println("dogNamePtr=", dogNamePtr)
	fmt.Println("dogNameP=&(dogPoint.name)", dogNameP,&(dogPoint.name),(dogNameP==&(dogPoint.name)))
	fmt.Println("*dogNameP=", *dogNameP)

	//修改name值
	*dogNameP = "little cat"
	fmt.Println("dogPoint.name=",dogPoint.name)
	fmt.Println("*dogNameP=",*dogNameP)

	//下面这种不匹配的转换虽然不会引发panic，但是其结果往往不符合预期
	dogNamePint := (*int)(unsafe.Pointer(dogNamePtr))
	fmt.Println("*dogNamePint=",*dogNamePint)
}
