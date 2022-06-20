package main

import "fmt"

type AnimalCategory struct {
	category_name string
	category_count int
}
type Animal struct {
	name string
	AnimalCategory
}
func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s %d",ac.category_name,ac.category_count)
}
func (animal Animal) String() string {
	return fmt.Sprintf("animal:%s [catetory:%s]",animal.name,animal.AnimalCategory)
}
func (ac AnimalCategory) PrintMe() string {
	return fmt.Sprintf("%s %d",ac.category_name,ac.category_count)
}
func main()  {
	category := AnimalCategory{category_name: "cat-category",category_count: 1}
	cat := Animal{name: "cat",AnimalCategory: category}

	//指定方法打印字符串
	fmt.Printf("PrintMe(): %s \n",category.PrintMe())
	//自动匹配String()方法打印
	fmt.Printf("String(): %s \n",category)
	//如果Animal没有String() 会调用 AnimalCategory的String()
	//如果Animal声明了String() 会屏蔽 AnimalCategory的String()
	fmt.Printf("cat: %s \n",cat)
}
