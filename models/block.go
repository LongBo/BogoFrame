package db

func GetBlocks() []map[string]string {
	db := NewDb()
	res := Query("SELECT * FROM block", db)
	return res
}

func SaveBlock(blockData map[string]string) int64{//添加、更新共用方法
	tablename := "block"
	db := NewDb()
	res := Save(db, tablename, blockData)
	return res
}

func DeleteBlock(idList []int64) int64 {
	tablename := "block"
	db  := NewDb()
	res := DelByIds(db, tablename, idList)
	return res
}
