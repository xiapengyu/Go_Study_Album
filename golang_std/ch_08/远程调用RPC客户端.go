package main

import (
	"fmt"
	"golang_std/ch_08/sharedRPC"
	"net/rpc"
)

func main() {
	c, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	args := sharedRPC.MyFloats{16, -0.5}
	var reply float64

	err = c.Call("MyInterface.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Multiply): %f\n", reply)

	argus := sharedRPC.MyFloats{2, 2}
	err = c.Call("MyInterface.Power", argus, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Power): %f\n", reply)
}
