package inheritTemplate

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./GoTemplate/inheritTemplate/rootTemplate/root.tmpl", "./GoTemplate/inheritTemplate/home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
	}

	tmpl.ExecuteTemplate(w, "home.tmpl", "homePage")

}
