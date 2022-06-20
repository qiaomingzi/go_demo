package lib

import (
	"fmt"
	"io"
)

func PrintLib(w io.Writer){
	fmt.Fprint(w,"demo3_lib_1 -> Print")
}
