package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := 9.5246

	k1 := new(big.Float).SetPrec(2).SetFloat64(a)
	fmt.Println(k1)

}
