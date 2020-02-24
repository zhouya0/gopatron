package Decorator

import (
	"fmt"
	"log"
	"net/http"
)

//装饰器模式，在go里面，看来一般就是指函数的封装。

func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithServerHeader()")
		w.Header().Set("Server", "HelloServer v0.0.1")
		h(w, r)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter是一个带有Write函数的接口
	// 同时，Fprintln的第一个参数也是一个io.Writer接口，只有Write函数
	log.Printf("Recieved Request %s from %s\n", r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "Hello, World!"+r.URL.Path)
}

func serve() {
	// 这个http.HandleFunc的第二个参数就是
	// type HandlerFunc func(ResponseWriter, *Request)
	// 所以hello函数就是符合这个HandlerFunc的
	http.HandleFunc("/v1/hello", WithServerHeader(hello))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
