package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//CopyFile("F:\\logs\\TMP.txt", "F:\\logs\\TMP2.txt")
	//ReadLine("F:\\logs\\TMP.txt")
}

/**
文件内容复制
*/
func CopyFile(src, des string) (w int64, err error) {
	//srcFile, err := os.Open(src)
	srcFile, err := os.OpenFile(src, os.O_RDONLY, 0666)
	handleErr(err)
	defer srcFile.Close()

	//文件不存在则新建，可读写，追加模式
	//desFile, err := os.Create(des)
	desFile, err := os.OpenFile(des, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	handleErr(err)
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

/**
逐行读取
*/
func ReadLine(src string) {
	srcFile, error := os.OpenFile(src, os.O_RDONLY, 0666)
	handleErr(error)
	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	//scanner := bufio.NewScanner(srcFile)
	//for scanner.Scan() {
	//	fmt.Println(scanner.Text())
	//}
}

func handleErr(err error) {
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
}
