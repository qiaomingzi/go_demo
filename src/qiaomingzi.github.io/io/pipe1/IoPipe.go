package main

import (
	"fmt"
	"io"
	"time"
)

func main(){
	r,w := io.Pipe()
	//每隔1s写入数据
	go func(){
		for{
			time.Sleep(time.Second)
			w.Write([]byte(time.Now().String()))
		}
	}()
	for{
		buff := make([]byte,256)
		n,_ := r.Read(buff)
		fmt.Println(string(buff[:n]))
	}
}
