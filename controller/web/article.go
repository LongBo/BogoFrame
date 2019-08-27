package web

import (
	"ncfwxen/common/router"
	"ncfwxen/common/utils"
	"ncfwxen/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	articles := db.GetArticles()
	utils.TmplWeb(w, "index", articles)
}

func Article(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	article := db.GetArticle(ps.ByName("id"))
	/*for _, article := range articles {
		for k, v := range article {
			fmt.Fprintf(w, "%s = %s \n", k, v)
		}
	}*/

	utils.TmplWeb(w, "article", article)
}
