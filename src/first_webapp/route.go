package main

import ("fmt"
		"net/http"
		"html/template"
	)
/*
	该文件定义了不同的请求处理函数	
*/
func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "this is index page")
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world, %s", request.URL.Path[1:])
}

/*返回html页面
	使用html/template解析html文件
*/
func djtest(writer http.ResponseWriter, request *http.Request){
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Printf("error")
        return
	}
	t.Execute(writer, nil)
}