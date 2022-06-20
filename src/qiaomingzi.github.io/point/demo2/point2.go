package main

import "fmt"
type TypeI interface {
	Name() string
}
type Dog struct {
	name string
}
func (dog Dog) Name() string{
	return dog.name
}

func main(){
	//1.字典 ++、--
	map[string]int{"one":1,"two":2}["one"]++
	m := map[string]int{"one":1,"two":2}
	m["one"]++
	fmt.Println(m["one"])

	//2.常量
	const num = 123
	//_ = &num //常量不可取址
	//_ = &(123) //基本类型字面量不取址

	//3.字符串变量
	str := "abc"
	_ = str
	//_ = &str[0] //对字符串索引结果不可寻址
	//_ = &(str[0:2]) //字符串切片不可寻址
	str2 := str
	_ = &str2

	//4.算术操作符结果
	num1 := 10
	num2 := 20
	//_ = &(num1+num2) //算术操作的结果值不可寻址
    //_ = &(10+20)  //算术操作的结果值不可寻址
    _ = (num1 + num2)

    //5.数组索引
    array := [3]int{1,2,3}
    _ = &(array[1])
    //_ = &([3]int{2,3,4}[0]) // 对数组字面量的索引结果值不可寻址
    //_ = &([3]int{2,3,4}[0:2]) //对数组字面量的切片结果值不可寻址
    //_ = &([]int{1,2,3}[0:2]) //对切片字面量的 切片结果值不可取址
    _ = &([]int{1,2,3}[1])

    //6.字典取址
    m1 := map[int]string{0:"zero",1:"one"}
    _ = m1
    //_ = &(map[int]string{0:"zero"}[0]) //对字典字面量的索引结果值不可寻址
	//_ = &(m1[0]) //对字典变量的索引结果不可寻址

	//7.函数取址
	f := func(a,b int) int{
		return a+b
	}
	_ = &f
	//_ = &(func()int{return 1}) //函数字面量不可寻址
	//_ = &(fmt.Sprint()) //标识符代表的函数不可寻址
	//_ = &(fmt.Println("test")) //函数调用的结果值不可寻址

	//结构体取址
	dog := Dog{"xiao dog"}
	_ = &(dog)
    _ = &(dog.name)
    //_ = &(dog.Name()) //方法调用结果值不可取址
    //_ = &(Dog{"little"}.name) //结构体字面量不可取址

    //类型断言
    dogAssert := interface{}(dog)
    _ = dogAssert
    _ = &(dogAssert)
	dogA := dogAssert.(Dog)
	_ = dogA
    //_ = &(dogAssert.(Dog)) //类型断言表达式 结果值不可取址
	//_ = &(dogA.(TypeI))//类型断言表达式 结果值不可取址

	//通道
	ch := make(chan int,1)
	ch <- 1
	//_ = &(<-ch) //接收表达式结果值 不可寻址
}
