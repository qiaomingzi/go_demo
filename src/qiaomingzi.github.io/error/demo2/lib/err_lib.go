package lib

import (
	"fmt"
	"os"
	"os/exec"
)

func UnderlyingError(err error) error{
	if err == nil{
		return nil
	}
	switch et:=err.(type) {
	case *os.PathError:
		return et.Err
	case *os.LinkError:
		return et.Err
	case *os.SyscallError:
		return et.Err
	case *exec.Error:
		return et.Err
	}
	return nil
}
func ErrorPrint(err error){
	if(err == nil){
		return
	}
	switch err {
	case os.ErrClosed:
		fmt.Println("os.ErrClosed...",err)
	case os.ErrInvalid:
		fmt.Println("os.ErrInvalid...",err)
	case os.ErrPermission:
		fmt.Println("os.ErrPermission...",err)
	default:
		fmt.Println("default...",err)
	}
}
