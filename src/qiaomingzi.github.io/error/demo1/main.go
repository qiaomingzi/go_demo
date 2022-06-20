package main

import (
	"errors"
	"fmt"
)

func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("request is empty")
		return
	}
	response = "pong:" + request
	return
}
func main() {
	//卫述语句
	for _, str := range []string{"", "limz"} {
		response, err := echo(str)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println("response:", response)
	}

	//判断错误信息是否相等
	err1 := fmt.Errorf("error %s","cause")
	err2 := errors.New("error cause")
	if err1.Error() == err2.Error(){
		fmt.Println("err1.Error()==err2.Error()")
	}
	if err1 != err2{
		fmt.Println("err1!=err2")
	}
}
