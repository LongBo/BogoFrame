package admin

import (
	"fmt"
	"ncfwxen/common/router"
	"ncfwxen/models"
	"net/http"
	"ncfwxen/common/utils"
)

func CategoryList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "List for category\n")
}

func DelCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	r.ParseForm()
	var id int64 = 1
	res := db.DeleteCategory(id)
	if res == db.DelErr {
		utils.JsonReturn(db.DelErr, "Delete Error : Can't delete this category", "", w)
	} else {
		str := fmt.Sprintf("%d articles", res)
		utils.JsonReturn(0, "OK", str, w)
	}
}

