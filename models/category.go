package db

import (
	"fmt"
	"strconv"
	"ncfwxen/common/utils"
	)

var DelErr int64 = 1001

func GetCategoryList() []map[string]string {
	db := NewDb()
	res := Query("SELECT * FROM category", db)
	return res
}

func SaveCategory(data map[string]string) int64{
	tablename := "category"
	db := NewDb()
	res := Save(db, tablename, data)
	return res
}

func DeleteCategory(id int64) int64 {
	tablename := "category"
	db  := NewDb()
	numQuery := fmt.Sprintf("SELECT COUNT(*) AS num FROM article WHERE cate_id = %d", id)
	queryRes := Query(numQuery, db)
	articleNum , err := strconv.Atoi(queryRes[0]["num"])
	if err != nil {
		utils.Nlog.Fatal(err)
	}
	if 0 == articleNum {
		utils.Nlog.Println("Process 1")
		idList := []int64{id}
		return DelByIds(db, tablename, idList)
	} else {
		utils.Nlog.Println("Process 2")
		return DelErr
	}
}