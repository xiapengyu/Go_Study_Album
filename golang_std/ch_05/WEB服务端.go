package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	//原始的路由匹配，不方便从url中解析参数
	/*http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":8056", nil)*/

	//gorilla/mux：方便从url中解析参数
	r := mux.NewRouter()
	r.HandleFunc("/test/{title}/page/{page}", testFunc)
	//可以指定HTTP方法
	r.HandleFunc("/books/{title}", testFunc).Methods("POST")
	r.HandleFunc("/books/{title}", testFunc).Methods("GET")
	r.HandleFunc("/books/{title}", testFunc).Methods("PUT")
	r.HandleFunc("/books/{title}", testFunc).Methods("DELETE")

	//对域名做限制
	r.HandleFunc("/books/{title}", testFunc).Host("www.hao123.com")

	//对http和https做限制
	r.HandleFunc("/secure", testFunc).Schemes("https")
	r.HandleFunc("/insecure", testFunc).Schemes("http")

	//将请求限定为特定的路径
	br := r.PathPrefix("/books").Subrouter()
	br.HandleFunc("/", testFunc)
	br.HandleFunc("/{title}", testFunc)

	http.ListenAndServe(":8056", nil)
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]
	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}
