package nestedTemplate

import (
	"fmt"
	"html/template"
	"net/http"
)

func NestedTemplate(w http.ResponseWriter, r *http.Request) {
	// 声明嵌套模版和解析模版
	template2, err2 := template.ParseFiles("./GoTemplate/nestedTemplate/nestedTemplate.tmpl", "./GoTemplate/nestedTemplate/childTemplateUl.tmpl")
	if err2 != nil {
		fmt.Printf("parse template failed, err: %v\n", err2)
	}

	/** 渲染模版 */
	err := template2.Execute(w, "hello world")
	if err != nil {
		fmt.Printf("render template failed: %v\n", err)
		return
	}
}
