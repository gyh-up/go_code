package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type msgWrap struct {
	ID         int
	//MsgType    string
	//MsgContent string
	Keyword 	string
}

func NewMySQLClient(addr, user, pass string) (db *gorm.DB, err error) {
	sqlParam := "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@(%s)/%s%s", user, pass, addr, "", sqlParam)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func main() {
	//db, err := gorm.Open("mysql", "novel_data:pass@novel_data@(159.75.27.239:3306)/novel_cp_data?charset=utf8mb4&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", "root:pass4novel@2020@(159.75.27.239:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")


	addr, user, pass := "159.75.27.239:3306", "novel_data", "pass@novel_data"
	db, err := NewMySQLClient(addr, user, pass)
	if err != nil {
		log.Panicln("[错误] 连接MySQL数据库", addr, "失败: ", err)
	}
	fmt.Println(db)


	keyword := 1

	wrap := msgWrap{}
	//db.First(&wrap)
	err = db.Table("novel_cp_data.mp_keyword_reply").
		Select("id, msg_type, msg_content, keyword").
		Where("status = 1 AND priority = 0 AND FIND_IN_SET(?, keyword)", keyword).
		Order("id DESC").
		Find(&wrap).
		Error
	fmt.Printf("%v, %v", wrap, err)
}

