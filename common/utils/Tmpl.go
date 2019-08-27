package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func TmplWeb(w http.ResponseWriter, tmpl string, data interface{}) *template.Template {
	t, _ := template.ParseFiles("views/web/layout.html", fmt.Sprintf("views/web/%s.html", tmpl))

	err := t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		Nlog.Println("executing template:", err)
	}
	return t
}

func TmplAdmin(w http.ResponseWriter, tmpl string, data interface{}) *template.Template {
	t, _ := template.ParseFiles("views/admin/dashboard.html", "views/admin/left.html", fmt.Sprintf("views/admin/%s.html", tmpl))

	err := t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		Nlog.Println("executing template:", err)
	}
	return t
}
