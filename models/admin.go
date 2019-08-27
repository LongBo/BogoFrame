package db

func GetAdmin() []map[string]string {
	db := NewDb()
	res := Query("SELECT * FROM admin", db)
	return res
}

func SaveAdmin(adminData map[string]string) int64{//添加、更新共用方法
	tablename := "admin"
	db := NewDb()
	res := Save(db, tablename, adminData)
	return res
}

func DeleteAdmin(idList []int64) int64 {
	tablename := "admin"
	db  := NewDb()
	res := DelByIds(db, tablename, idList)
	return res
}
