package supplementTemplate

import (
	"fmt"
	"html/template"
	"net/http"
)

func SupplementTemplate(w http.ResponseWriter, r *http.Request) {

	tmpl := template.New("supplementTemplate.tmpl")

	// 定义一个解析为HTML模版的函数
	tmpl.Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	_, err := tmpl.Delims("{[", "]}").
		ParseFiles("./GoTemplate/supplementTemplate/supplementTemplate.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err: %v\n", err)
	}

	data := map[string]interface{}{
		"htmlCon": "<script>alert('恶意脚本')</script>",
		"textCon": "<a href='http://www.leechenze.com'>李晨泽的博客<a/>",
	}
	tmpl.Execute(w, data)

}
