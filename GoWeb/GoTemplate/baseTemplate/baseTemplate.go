package baseTemplate

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func BaseTemplate(w http.ResponseWriter, r *http.Request) {

	/** 定义模版函数 */
	// 要么只能有一个返回值，要么有两个返回值，第二个返回值必须为error
	kua := func(name string) (string, error) {
		return name + "is great", nil
	}

	/** 定义模版数据 */
	str := "模版内容"
	user := User{Name: "trump", Gender: "man", Age: 20}
	map1 := map[string]interface{}{
		"name":   "clinton",
		"gender": "female",
		"age":    30,
	}
	hobbyList := []string{"抽烟", "喝酒", "烫头"}
	hobbyList1 := []string{}

	data := map[string]interface{}{
		"user":       user,
		"str":        str,
		"map1":       map1,
		"hobbyList":  hobbyList,
		"hobbyList1": hobbyList1,
	}

	/** 解析模版 */
	// 声明模版
	// template1, err := template.ParseFiles("./GoTemplate/baseTemplate/baseTemplate.tmpl")
	// 声明具名模版（baseTemplate），注意New中的name参数 必须是后续解析路径的文件名称。
	template1 := template.New("baseTemplate.tmpl")

	/** 注册模版函数 */
	template1.Funcs(template.FuncMap{
		"kua": kua,
	})
	_, err := template1.ParseFiles("./GoTemplate/baseTemplate/baseTemplate.tmpl")

	if err != nil {
		fmt.Printf("parse template failed: %v\n", err)
		return
	}

	/** 渲染模版 */
	err = template1.Execute(w, data)
	if err != nil {
		fmt.Printf("render template failed: %v\n", err)
		return
	}
}
