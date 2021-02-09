package main

import (
	"fmt"
	"golang_std/ch_08/sharedRPC"
	"math"
	"net"
	"net/rpc"
)

type MyInterface struct {}

func (t *MyInterface) Multiply(arguments *sharedRPC.MyFloats, reply * float64) error {
	*reply = arguments.A1 * arguments.A2
	return nil
}

func (t *MyInterface) Power(arguments *sharedRPC.MyFloats, reply *float64) error {
	*reply = math.Pow(arguments.A1, arguments.A2)
	return nil
}

func main() {
	myInterface := new(MyInterface)
	rpc.Register(myInterface)
	t, err := net.ResolveTCPAddr("tcp4", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}
}
