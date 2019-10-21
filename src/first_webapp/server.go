package main

import ("fmt"
	"net/http")

func main() {
	fmt.Printf("server is running at localhost:8083.....")
	mux := http.NewServeMux()	// 创建多路复用器，添加HandleFunc来处理不同请求

	mux.HandleFunc("/main", handler)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/dj", djtest)

	server := &http.Server{
		Addr: "0.0.0.0:8083",
		Handler: mux,
	}
	server.ListenAndServe()
}
