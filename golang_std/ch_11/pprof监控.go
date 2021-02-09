package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"  // 为什么用_ , 在讲解http包时有解释。
)

func myFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

func main() {
	http.HandleFunc("/", myFunc)
	//访问http://localhost:8080/debug/pprof/
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
