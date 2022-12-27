package main

import (
	"fmt"
	"os"
	"os/exec"
	lib "qiaomingzi.gihub.io/error/demo2/lib"
)
/**
 r=4 w=3 x=1
 */
func main() {
	closeFd := func(fd *os.File){
	   if fd != nil{
		   fd.Close()
	   }
	}
	{
		fmt.Println("==============文件不存在==============")
		fd,_ := os.Open("/it/must/not/exist")
		_,err := fd.Stat()
		lib.ErrorPrint(err)
		closeFd(fd)
	}
	{
		fmt.Println("==============关闭文件描述符==============")
		fd,err := os.Open(os.Args[0])
		if err != nil {
			fmt.Printf("unexpected error: %s\n", err)
			return
		}
		fd.Close()
		n, err := fd.Read([]byte{})
		lib.ErrorPrint(err)
		fmt.Println("fd.Read:", n)
		closeFd(fd)
	}
    {
		fmt.Println("==============LookPath==============")
		str,err := exec.LookPath("D:/temp/tmp/")
		lib.ErrorPrint(err)
		fmt.Println("exec.LookPath：",str)
	}
}
