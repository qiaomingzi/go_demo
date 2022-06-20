package main

import "fmt"

func main(){
	fmt.Println(4   |  5)
	//for遍历数组
	var array = [5]int{1,2,3,4,5}
	arrayLen := len(array)
	for i,_ := range array{
		if(i == 0){
			array[i] += array[i]
		} else if(i == (arrayLen - 1)){
			//0010 | 0101 =>0111
			// 2   |  5   =>7
			array[i] ^= array[0]
		}
	}
	fmt.Printf("array:%v len:%d \n",array,arrayLen)
	//for 遍历数组
	array1 := [...]int{1,2,3,4,5}
	array1Len := len(array1)
    for i,e := range array1{
    	if(i == (array1Len - 1)){
			// 1+5,3,5,7,9 =>  6,3,5,7,9
			array1[0] += e
		} else {
			// 1,2+1,3,4,5
			// 1,3,3+2,4,5
			// 1,3,5,4+3,5
			// 1,3,5,7,5+4
			array1[i+1] += e
		}
	}
	fmt.Printf("array1:%v len:%d \n",array1,array1Len)
    //切片
	slice2 := []int{1,2,3,4,5}
	slice2Len := len(slice2)
	for i,e := range slice2{
		if(i == (slice2Len - 1)){
			// 1+15,3,5,10,15 => 16,3,5,10,15
			slice2[0] += e
		} else {
			// 1,2+1,3,4,5
			// 1,3,3+3,4,5
			// 1,3,6,4+6,5
			// 1,3,5,10,5+10
			slice2[i+1] += e
		}
	}
	fmt.Printf("slice:%v len:%d \n",slice2,slice2Len)
}
