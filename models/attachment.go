package db

func GetAttachment() []map[string]string {
	db := NewDb()
	res := Query("SELECT * FROM attachment", db)
	return res
}

func SaveAttachment(AttachmentData map[string]string) int64{//添加、更新共用方法
	tablename := "attachment"
	db := NewDb()
	res := Save(db, tablename, AttachmentData)
	return res
}

func DeleteAttachment(idList []int64) int64 {
	tablename := "attachment"
	db  := NewDb()
	res := DelByIds(db, tablename, idList)
	return res
}
