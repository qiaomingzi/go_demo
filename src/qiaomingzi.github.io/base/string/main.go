package main

import (
	"fmt"
	"strings"
)

func main(){
	var s1 = "li"
	var s2 = "mignzi"
	//+号连接
	s3 := s1+s2
	fmt.Printf("plus：%s \n",s3)
	//Sprintf
	s3 = fmt.Sprintf("family:%s name:%s",s1,s2)
	fmt.Printf("Sprintf：%s \n",s3)
	//strings.Builder
	var sb strings.Builder
	sb.WriteString(s1)
	sb.WriteString(s2)
	s3 = sb.String()
	fmt.Printf("Builder：%s \n",s3)
}
