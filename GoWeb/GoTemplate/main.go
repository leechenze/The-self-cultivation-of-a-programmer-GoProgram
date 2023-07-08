package GoTemplate

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模版
	template, err := template.ParseFiles("./GoTemplate/hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed: %v\n", err)
		return
	}
	// 渲染模版
	data := "模版内容"
	err = template.Execute(w, data)
	if err != nil {
		fmt.Printf("render template failed: %v\n", err)
		return
	}
}

func Main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed: %v\n", err)
		return
	} else {
		fmt.Println("9000 serve is running")
	}
}
