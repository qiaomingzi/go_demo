package main

import "fmt"

func main(){
  m := map[string]int{
  	"one":1,
  	"two":2,
  	"three":3,
  }
  key := "two"
  print(m,key)
  delete(m,key)
  print(m,key)
  m[key] = 2
  print(m,key)

  var m2 map[string]int
  print(m2,key)
  delete(m2,key)
  //m2[key] = 2 //出现panic 向nil的字典添加值
}

func print(m map[string]int,key string)  {
	v,ok := m[key]
	if ok {
		fmt.Printf("value:%d \n",v)
	} else {
		fmt.Printf("value not found \n")
	}
}
