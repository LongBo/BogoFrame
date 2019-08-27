package utils

import (
	"encoding/json"
	"net/http"
	"fmt"
	)

var JsonData map[string]interface{}

func init()  {
	JsonData = make(map[string]interface{})
}

func JsonReturn (code int64, message string, data string, w http.ResponseWriter)  {
	if 0 == code {
		JsonData["code"] = code
		JsonData["message"] = message
		JsonData["data"] = data
	} else {
		JsonData["code"] = code
		JsonData["error"] = message
		JsonData["data"] = data
	}
	result, err := json.Marshal(JsonData)
	if (err != nil) {
		Nlog.Println(err)
	}
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	fmt.Fprintf(w,string(result))
	//return string(result)
}