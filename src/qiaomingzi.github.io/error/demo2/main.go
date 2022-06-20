package main

import (
	"fmt"
	"os"
)
import tool "qiaomingzi.gihub.io/error/demo2/lib"

func main(){
	r,w,err := os.Pipe()
	if err != nil{
		fmt.Println("os.Pipe() error")
		return
	}
	r.Close()
	n, err := w.Write([]byte("hello"))
	fmt.Println("w.write:",n)
	whatErr := tool.UnderlyingError(err)
	//The pipe is being closed. [type: syscall.Errno]
	fmt.Printf("underlyingError(): %s [type: %T]\n",whatErr,whatErr)
}
