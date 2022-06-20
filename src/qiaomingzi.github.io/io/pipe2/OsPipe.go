package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	inR,inW,_ := os.Pipe()
	outR,outW,_ := os.Pipe()
	fmt.Println(os.Args)
	dir,_ := filepath.Abs(filepath.Dir(os.Args[0]))
	process,_ := os.StartProcess("E:/cygwin64/Cygwin.bat",nil,&os.ProcAttr{
		Files: []*os.File{inR,outW,os.Stderr},
		Dir: dir,
	})
	go func(){
		writer := bufio.NewWriter(inW)
		writer.WriteString("ls -l")
		writer.Flush()
		inW.Close()
		outW.Close()
	}()
	process.Wait()
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(outR)
	fmt.Println(buffer.String())
	outR.Close()
}
