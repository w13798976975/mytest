package main

import (
	"fmt"
	"io"
	"net/http"
	"log"
)

func main() {
	//fmt.Println("")
	//http2()

	m:=map[string]interface{}{}
	//m=make(map[string]interface{})
	m["data"]="string value"
	m["sex"]=false

	for v:=range m{
		fmt.Println(v)
	}
}

func http1() {
	http.HandleFunc("/", SayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func SayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "say hello,this is version 1")
}

func http2(){
	mux:=http.NewServeMux()
	mux.Handle("/",&myHandler{})
	mux.HandleFunc("/hello1",SayHello)
	err:=http.ListenAndServe(":8080",mux)
	if err!=nil{
		log.Fatal(err)
	}
}

type myHandler struct {}

func (*myHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,r.URL.String())
}
