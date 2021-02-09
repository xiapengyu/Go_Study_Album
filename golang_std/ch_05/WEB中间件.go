package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", logging(foo))
	http.ListenAndServe("127.0.0.1:8756", nil)
}

/**
日志中间件，拦截请求，打印完日志后，传入的http.HandlerFunc对象用户服务器调用
*/
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}
