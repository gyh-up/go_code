package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Test struct {
	Id int
	Appid string
}

func main() {
	var user Test
	//dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err := gorm.Open("mysql", "novel_data:pass@novel_data@(159.75.27.239:3306)/novel_cp_data?charset=utf8mb4&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")

	if err!= nil{
		panic(err)
	}
	defer db.Close()

	//db.Table("test_tb_name").First(&user) // 通过数据的指针来创建

	db.Table("mp_users_info_10").Last(&user)
	fmt.Println(user)
}
