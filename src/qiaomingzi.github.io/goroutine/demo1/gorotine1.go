package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	outBysleep()
	outByChan()
	outAtomic()
}

//使用atomic
func outAtomic() {
	fmt.Println("===========atomic 顺序输出=============")
	var index int32
	trigger := func(i int32, fn func()) {
		for {
			n := atomic.LoadInt32(&index);
		    fmt.Sprintf("n：%d i:%d",n,i)
			if n == i {
				fn()
				atomic.AddInt32(&index, 1)
				break
			}
		}
	}

	for i := int32(0); i < 10; i++ {
		go func(i int32) {
			print := func() {
				fmt.Println(i)
			}
			trigger(i, print)
		}(i)
	}
	//让主goroutine等待其他goroutine
	trigger(10, func() {fmt.Println(10)})
}

//使用通道
func outByChan() {
	fmt.Println("===========chan 无序输出=============")
	ch := make(chan struct{}, 10)
	for i := 0; i < 10; i++ {
		go func(param int) {
			fmt.Println(param)
			ch <- struct{}{}
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}

//使用sleep
func outBysleep() {
	fmt.Println("===========sleep 无序输出=============")
	for i := 0; i < 10; i++ {
		go func(param int) {
			fmt.Println(param)
		}(i)
	}
	time.Sleep(1000)
}
