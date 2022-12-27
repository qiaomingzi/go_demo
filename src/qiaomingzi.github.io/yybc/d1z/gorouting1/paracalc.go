package main

import (
	"fmt"
	"github.com/petermattis/goid"
	"runtime"
	"strconv"
	"strings"
)

/**
need to exec : go get -u github.com/petermattis/goid
*/
func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Println("goroutine:", goID(), goid.Get(), sum)
	resultChan <- sum
}
func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 2)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan

	fmt.Println("result:", sum1, sum2, sum1+sum2)
}

func goID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	/*
			string(buf[:n])运行时堆栈 值：
			goroutine 7 [running]:
			main.goID(0x0)
			E:/workspace/goWorkspace/

			strings.TrimPrefix(string(buf[:n]), "goroutine ") 去除前缀值:
			7 [running]:
			main.goID(0x0)
		    E:/workspace/goWorkspace/
	*/
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
