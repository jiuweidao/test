package main

import (
	"io"
	"net/http"
)


/*
	简单的web应用
 */
func hello(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "hello widuu")
}

func main() {
	http.HandleFunc("/", hello)  //设定访问的路径
	http.ListenAndServe(":8080", nil) //设定端口和handler
}