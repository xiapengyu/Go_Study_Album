/**
io/ioutil：提供了一些使用的io工具
fmt：格式化io
*/

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type Person struct {
	Name string
	Age int
	Sex int
}

func main() {
	//ListFiles("F:\\文件\\jetbrains破解工具")
	//ReadFileContent("F:\\文件\\jetbrains破解工具\\安装参数.txt")
	//ReadFileByLine("F:\\文件\\工作日志\\工作日志.txt")
	//WriteFile("F:\\文件\\工作日志\\工作日志1.txt", "a")
	//TraceFile("F:\\logs\\TMP.txt", "b")

	/*person := Person{"XPY", 1, 1}
	fmt.Println(person)
	fmt.Printf("%#v", person)*/

	DbConnection()
}

/**
一次新读取文本内容
*/
func ReadFileContent(name string) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(file))
}

/**
逐行读取文本内容
*/
func ReadFileByLine(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	for ; ; {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			return
		}
		fmt.Println(string(line))
	}
}

/**
写入文件,覆盖
*/
func WriteFile(name, content string) {
	//文件不存在会新建文件，覆盖更新
	err := ioutil.WriteFile(name, []byte(content), 0644)
	if err == nil {
		fmt.Println("success")
	} else {
		fmt.Println(err)
	}
}

func TraceFile(name, content string) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		file.Write([]byte(content))
	}
}

/**
  遍历目录基子目录，打印所有文件名
*/
func ListFiles(path string) {
	//fmt.Println("搜索目录:", path)
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, file := range dir {
			if file.IsDir() {
				fmt.Println(file.Name(), "\\")
				ListFiles(path + "/" + file.Name())
			} else {
				fmt.Println(file.Name())
			}
		}
	}
}

func DbConnection()  {
	//db是一个句柄而非数据库连接，只有使用的时候才会建立数据库连接
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1)/nacos_config")
	//控制连接池
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	if err != nil {
		fmt.Println(err)
		return
	}else{
		for i := 0; i < 10; i++ {
			go func() {
				db.Ping()
			}()
		}
		time.Sleep(20 * time.Second)

		/*fmt.Println("please exec show processlist")
		time.Sleep(10 * time.Second)
		fmt.Println("please exec show processlist again")
		db.Ping()
		time.Sleep(10 * time.Second)*/

	}
}
