package djpackage

import ("fmt"
		"net/http")

func index_(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "this is my package page")
}