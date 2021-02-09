package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi("-180")
	fmt.Println(a)
	f, _ := strconv.ParseFloat("12", 1)
	fmt.Println(f)

	fmt.Print("hello\n gopher\n")
	fmt.Println(`hello\n gopher\n`)

	var b bytes.Buffer
	b.WriteString("hello")
	b.WriteString(" world!")
	fmt.Println(b.String())
}
