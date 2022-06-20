package main

import "fmt"

func main(){
	//数据作为参数值传递
	array1 := [3]int{1,2,3}
	fmt.Printf("array1 %v\n",array1)
	array2 := modify(array1)
	fmt.Printf("array2 %v\n",array2)
	fmt.Printf("array1 %v\n",array1)
	//引用类型 切片、字典、通道
	yy1 := [3][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}
	fmt.Printf("yy1 %v\n",yy1)
	yy2 := modifySlice(yy1[1])
	fmt.Printf("yy2 %v\n",yy2)
	fmt.Printf("yy1 %v\n",yy1)
}
func modify(param [3]int) [3]int{
	param[1] = 40
	return param
}
func modifySlice(param []int) []int{
	param[1] = 40
	return param
}