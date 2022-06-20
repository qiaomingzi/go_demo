package main

import "fmt"

func main()  {
	//指定固定长度得切片
	s1 := make([]int,5)
	fmt.Printf("the length of s1:%d\n",len(s1));
	fmt.Printf("the captcity of s1:%d\n",cap(s1));
	fmt.Printf("the value of s1:%d\n",s1);
	//指定长度、容量得切片
	s2 := make([]int,5,8)
	fmt.Printf("the length of s2:%d\n",len(s2));
	fmt.Printf("the captcity of s2:%d\n",cap(s2));
	fmt.Printf("the value of s2:%d\n",s2);
	//
	s3 := []int{1,2,3,4,5,6,7,8};
	s4 := s3[3:6]
	fmt.Printf("the length of s4:%d\n",len(s4));
	fmt.Printf("the captcity of s4:%d\n",cap(s4));
	fmt.Printf("the value of s4:%d\n",s4);

	//扩容
	s5 := make([]int,0)
	for i :=1; i < 5;i++{
		s5 = append(s5,i)
		fmt.Printf("s5(%d): len: %d, cap: %d\n", i, len(s5), cap(s5))
	}
	s5_1 := append(s5,make([]int,102)...)
	fmt.Printf("s5_1: len: %d, cap: %d\n", len(s5_1), cap(s5_1))
	s6 := make([]int, 1024)
	s6_1 := append(s6, make([]int, 200)...)
	fmt.Printf("s6_1: len: %d, cap: %d\n", len(s6_1), cap(s6_1))
}
