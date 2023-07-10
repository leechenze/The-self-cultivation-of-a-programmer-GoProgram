package GoTemplate

import (
	"GoWeb/GoTemplate/supplementTemplate"
	"fmt"
	"net/http"
)

func Main() {
	/** 基础模版 */
	// http.HandleFunc("/", baseTemplate.BaseTemplate)
	/** 嵌套模版 */
	// http.HandleFunc("/", nestedTemplate.NestedTemplate)
	/** 模版继承 */
	// http.HandleFunc("/index", inheritTemplate.Index)
	// http.HandleFunc("/home", inheritTemplate.Home)
	/** 模版补充 */
	http.HandleFunc("/", supplementTemplate.SupplementTemplate)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed: %v\n", err)
		return
	} else {
		fmt.Println("9000 serve is running")
	}

}
