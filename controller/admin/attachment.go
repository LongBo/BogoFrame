package admin

import (
	"fmt"
	"io"
	"io/ioutil"
	"ncfwxen/common/router"
	"ncfwxen/common/utils"
	"ncfwxen/models"
	"net/http"
	"os"
)

func AddAttachment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		utils.Nlog.Println("FileName= [" + part.FileName() + "], FormName=[" + part.FormName() + "]\n")
		if part.FileName() == "" {
			data, _ := ioutil.ReadAll(part)
			utils.Nlog.Println("FormData=[" + string(data) + "]\n")
		} else { //此处可以单独写个文件?
			dir := "bin/attachment"
			_, err = os.Stat(dir)
			if os.IsNotExist(err) {
				os.Mkdir(dir, 0755)
			}
			dst, _ := os.Create(dir + "/" + part.FileName())
			defer dst.Close()
			_, err = io.Copy(dst, part)
			//if err == nil {
			//	data := make(map[string]string)
			//	filename := strings.Split(part.FileName(),".")
			//	data["type"] := filename[len(filename)-1]
			//
			//}
		}
	}
}

func DelAttachment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	idlist := []int64{6, 7}
	res := db.DeleteArticle(idlist)
	output := fmt.Sprintf("result : %d rows affected", res)
	utils.JsonReturn(0, "OK", output, w)
}
