package admin

import (
	"fmt"
	"ncfwxen/common/router"
	"ncfwxen/common/utils"
	"ncfwxen/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	articles := db.GetArticles()
	utils.TmplAdmin(w, "index", articles)
}

func Article(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, this is Article ID:%s!\n", ps.ByName("id"))
}

func DelArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	idlist := []int64{6, 7}
	res := db.DeleteArticle(idlist)
	output := fmt.Sprintf("result : %d rows affected", res)
	utils.JsonReturn(0, "OK", output, w)
}
