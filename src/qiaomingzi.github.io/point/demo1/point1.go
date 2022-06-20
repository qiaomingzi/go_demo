package main

type Cat struct {
	name string
}

func New(name string) Cat {
	return Cat{name}
}
func (cat *Cat) SetName(name string) {
	cat.name = name
}
func main() {
	//New("miao").SetName("ji ji") //编译器报错 临时结果不可取址
    cat := New("miao")
    cat.SetName("jiji")
}
