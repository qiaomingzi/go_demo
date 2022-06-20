package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int, 3)
	//发送数据到通道
	c <- 1
	c <- 2
	c <- 3
	//从通道中接收一个数据
	cr := <-c
	fmt.Printf("read chan value %v", cr)
	//带有range子句的for语句
	rangefor()
	//单向发送通道
	sendDirectionChan()
	//select子句
	selectChan()
	selectChanClose()
}

//===================for range子句==================
func sendDirectionChan() {
	ch := make(chan int, 1)
	sendSingleChan(ch)
	fmt.Println("")
	fmt.Printf("chan value %v", <-ch)
}
func rangefor() {
	intChan := getInChan()
	for meta := range intChan {
		fmt.Printf("meta:%d ", meta)
	}
}

//===================单向通道==================
func getInChan() <-chan int {
	len := 5
	c := make(chan int, len)
	for i := 0; i < len; i++ {
		c <- i
	}
	close(c)
	return c
}
func sendSingleChan(ch chan<- int) {
	ch <- rand.Intn(1000)
}

//===================select==================
func selectChan() {
	fmt.Println("")
	chanArray := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	//从0-2随机
	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	chanArray[index] <- index
	select {
	case meta1 := <-chanArray[0]:
		fmt.Printf("slect first chan %v", meta1)
	case meta2 := <-chanArray[1]:
		fmt.Printf("slect second chan %v", meta2)
	case meta3 := <-chanArray[2]:
		fmt.Printf("slect third chan %v", meta3)
	default:
		fmt.Printf("slect default")
	}
}
func selectChanClose() {
	fmt.Println("")
	ch := make(chan int,1)
	time.AfterFunc(time.Second,func(){
		close(ch)
	})
	select{
	case v,ok := <-ch:
	  if(!ok){
		  fmt.Printf("slect case ok:%v",ok)
		  break
	  }
	  fmt.Println("select case %v",v)
	}
}
