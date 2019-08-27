package db

import "fmt"

func GetArticles() []map[string]string {
	db := NewDb()
	res := Query("SELECT * FROM article", db)
	return res
}

func SaveArticle(articleData map[string]string) int64 { //添加、更新共用方法
	tablename := "article"
	db := NewDb()
	res := Save(db, tablename, articleData)
	return res
}

func DeleteArticle(idList []int64) int64 {
	tablename := "article"
	db := NewDb()
	res := DelByIds(db, tablename, idList)
	return res
}

func GetArticle(id string) []map[string]string {
	db := NewDb()
	res := Query(fmt.Sprintf("SELECT * FROM article where id=%s", id), db)
	return res
}
