package db

import (
	"database/sql"
	"fmt"
	_ "ncfwxen/common/mysql"
	"ncfwxen/common/utils"
	"ncfwxen/config"
)

func NewDb() *sql.DB {

	dbConf := *config.DbConfig()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbConf.User,
		dbConf.Pwd,
		dbConf.Host,
		dbConf.Port,
		dbConf.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		utils.Nlog.Fatal(err)
	}
	return db
}

func Query(sql string, db *sql.DB) []map[string]string {
	rows, err := db.Query(sql)
	if err != nil {
		utils.Nlog.Fatal(err)
	}

	var res = make([]map[string]string, 0)
	for rows.Next() {
		cls, _ := rows.Columns()
		scanArgs := make([]interface{}, len(cls))
		values := make([]interface{}, len(cls))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		err = rows.Scan(scanArgs...)
		if err != nil {
			utils.Nlog.Fatal(err)
		}
		record := make(map[string]string)
		for i, v := range values {
			if v != nil {
				record[cls[i]] = string(v.([]byte))
			}
		}
		res = append(res, record)
	}
	return res
}

func Insert(sql string, db *sql.DB) int64 {
	result, err := db.Exec(sql)
	if err != nil {
		utils.Nlog.Fatal(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		utils.Nlog.Fatal(err)
	}
	return lastInsertId
}

func Update(sql string, db *sql.DB) int64 {
	result, err := db.Exec(sql)
	if err != nil {
		utils.Nlog.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.Nlog.Fatal(err)
	}
	return rowsAffected
}

func Remove(sql string, db *sql.DB) int64 {
	result, err := db.Exec(sql)
	if err != nil {
		utils.Nlog.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.Nlog.Fatal(err)
	}
	return rowsAffected
}

func Description(table string, db *sql.DB) map[string]string {
	queryString := fmt.Sprintf("DESCRIBE %s", table)
	rows, err := db.Query(queryString)
	if err != nil {
		utils.Nlog.Fatal(err)
	}

	var res = map[string]string{}
	for rows.Next() {
		cls, _ := rows.Columns()
		scanArgs := make([]interface{}, len(cls))
		values := make([]interface{}, len(cls))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		err = rows.Scan(scanArgs...)
		if err != nil {
			utils.Nlog.Fatal(err)
		}
		filed := string(values[0].([]byte))
		fieldType := string(values[1].([]byte))
		res[filed] = fieldType
	}
	return res
}

func Save(db *sql.DB, tablename string, data map[string]string) int64 {
	id, isset := data["id"]
	delete(data, "id")
	description := Description(tablename, db)
	fieldCheck(data, description)
	if !isset { //insert
		keys := ""
		vals := ""
		for i := range data {
			keys += fmt.Sprintf("`%s`,", i)
			vals += fmt.Sprintf("'%s',", data[i])
		}
		keySlice := []byte(keys)
		keySlice = keySlice[:len(keySlice)-1]
		valSlice := []byte(vals)
		valSlice = valSlice[:len(valSlice)-1]
		sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tablename, string(keySlice), string(valSlice))
		return Insert(sql, db)
	} else { //update
		pair := ""
		for i := range data {
			pair += fmt.Sprintf(" `%s` = '%s',", i, data[i])
		}
		pairSlice := []byte(pair)
		pairSlice = pairSlice[:len(pairSlice)-1] // remove last","
		sql := fmt.Sprintf("UPDATE %s SET %s WHERE (`id` = %s) ", tablename, string(pairSlice), id)
		fmt.Println(sql)
		return Update(sql, db)
	}
}

func fieldCheck(data, descirption map[string]string) { //delete the field which doesn't exist in the table
	for i := range data {
		_, isset := descirption[i]
		if !isset {
			delete(data, i)
		}
	}
}

func DelByIds(db *sql.DB, tablename string, idlist []int64) int64 {
	ids := ""
	for _, id := range idlist {
		ids += fmt.Sprintf("%d,", id)
	}
	idSlice := []byte(ids)
	idSlice = idSlice[:len(idSlice)-1]
	sql := fmt.Sprintf("DELETE FROM %s WHERE `id` in (%s) ", tablename, string(idSlice))
	utils.Nlog.Println(sql)
	return Remove(sql, db)
}
