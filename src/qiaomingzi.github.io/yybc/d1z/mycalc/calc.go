package main

import (
	"fmt"
	"os"
	"qiaomingzi.gihub.io/yybc/d1z/mycalc/simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE: calc command [argument] ....")
	fmt.Println("\nThe commands are: \n\t add \t  Addtion of two values. \n\t sqrt \t Square root of a non-negative value")
}

func getArgs() {
	fmt.Println("args len:", len(os.Args))
	for i, v := range os.Args {
		fmt.Println("args[%d]=%v", i, v)
	}
}

/**
cd bin
go build "qiaomingzi.gihub.io/yybc/d1z/mycalc"
./mycalc "add" 2 3
./mycalc "sqar" 2

*/
func main() {
	getArgs()
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	if args[0] == "" || args[0] != "add" || args[0] != "sqrt" {
		Usage()
		return
	}
	switch args[0] {
	case "add":
		{
			addUsage := "USAGE: calc add <integer1><integer2>"
			if len(args) != 3 {
				fmt.Println(addUsage)
				return
			}
			v1, err1 := strconv.Atoi(args[1])
			v2, err2 := strconv.Atoi(args[2])
			if err1 != nil || err2 != nil {
				fmt.Println(addUsage)
				return
			}
			result := simplemath.Add(v1, v2)
			fmt.Println("Result: ", result)
			return
		}
	case "sqrt":
		{
			sqrtUsage := "USAGE: calc sqrt <integer1>"
			if len(args) != 2 {
				fmt.Println(sqrtUsage)
				return
			}
			v, err1 := strconv.Atoi(args[1])
			if err1 != nil {
				fmt.Println(sqrtUsage)
				return
			}
			result := simplemath.Sqrt(v)
			fmt.Println("Result: ", result)
			return
		}
	default:
		Usage()
	}
}
