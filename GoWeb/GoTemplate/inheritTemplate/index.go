package inheritTemplate

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./GoTemplate/inheritTemplate/rootTemplate/root.tmpl", "./GoTemplate/inheritTemplate/index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
	}

	tmpl.ExecuteTemplate(w, "index.tmpl", "indexPage")

}
