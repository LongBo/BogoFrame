package admin

import (
	"fmt"
	"log"
	"ncfwxen/common/router"
	"ncfwxen/common/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.TmplAdmin(w, "login", nil)
}

func DoLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Panic("form data is error")
	}
	fmt.Fprintln(w, r.Form.Get("user"))
	fmt.Fprintln(w, r.Form.Get("password"))
}
