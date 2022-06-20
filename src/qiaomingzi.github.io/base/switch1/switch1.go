package main
import "fmt"

func main(){
    array1 := [...]int8{1,2,3,4,5}
	fmt.Println("======[编译错误 存在相同的case]=======")
	//存在相同的case
	//switch array1[2] {
	//case 3:
	//	fmt.Println("case 3 \n",3)
	//case 3:
	//	fmt.Println("case 3 \n",3)
	//}
	fmt.Println("======[编译错误 存在相同的case,(byte <- uint)]=======")
	//value := interface{}(byte(127))
	//switch value.(type) {
	//case int8,uint8:
	//	fmt.Println("case int8")
	//case byte:
	//	fmt.Println("case byte")
	//default:
	//	fmt.Println("case default")
	//}

	fmt.Println("======[编译成功 存在相同的case]=======")
	switch array1[2] {
	case array1[2]:
		fmt.Println("case 3 \n",3)
	case array1[2],array1[3]:
		fmt.Println("case 3,4 \n",3,4)
	}
	fmt.Println("======[编译错误 类型不一致]=======")
	//类型不一致,无法case
	//switch 1 + 2 {
	//case array1[0]:
	//	fmt.Printf("case array1[0]:%d \n",array1[0])
	//case array1[1], array1[2]:
	//	fmt.Printf("case array1[1-2]:%d  %d \n",array1[1],array1[2])
	//case array1[3], array1[4]:
	//	fmt.Printf("case array1[3-4]:%d  %d \n",array1[3], array1[4])
	//}

	fmt.Println("======[编译成功 发生自动类型转换 int -> int8]=======")
	//
	switch array1[3]{
	case 1:
		fmt.Println("case 1:%d \n",1)
	case 2,3:
		fmt.Println("case 2,3:%d  %d \n",2,3)
	case 4,5:
		fmt.Printf("case 4,5:%d  %d \n",4,5)
	}

	fmt.Println("======[类型转换]=======")
	 /*
	 int8 -128~127
	 rune int32
	 uint 32 或 64 位
	 byte uint8
	 int 与 uint 一样大小
     uintptr 无符号整型，用于存放一个指针
	 */
	var i int
    var i8 int8
    i = 129
    //i8 = 129  //constant 129 overflows int8
	fmt.Println(int8(i))
	fmt.Println(int(i8))
}
