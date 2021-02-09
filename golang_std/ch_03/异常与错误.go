package main

import (
	"fmt"
	"io"
	"os"
	"syscall"
)

func main() {
	fmt.Println(ParseJSON())
}

func CatchErr() {
	err := syscall.Chmod(":invalid path:", 0666)
	if err != nil {
		fmt.Println(err)
		fmt.Println(err.(syscall.Errno))
	}
}

func CopyFile(dstFile, srcFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func ParseJSON() (err error) {
	defer func() {
		if p := recover(); p != nil { //必须在defer函数中直接调用recover
			err = fmt.Errorf("JSON: internal error: %v", p)
		}
	}()
	panic("啊啊啊")
	return
}
