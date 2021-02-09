package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/index", IndexHandle)

	err := http.ListenAndServe("127.0.0.1:1111", nil)
	checkErr(err)
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("You've requested the url: %s", r.Host+r.URL.Path)
	resultMap := make(map[string]string)
	resultMap["user"] = "Lily"
	resultMap["sex"] = "Female"

	ret, _ := json.Marshal(resultMap)
	//io.WriteString(w, string(ret))
	//w.Write([]byte(ret))
	fmt.Fprintf(w, "%s", ret)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("服务器启动失败", err)
	}
}
