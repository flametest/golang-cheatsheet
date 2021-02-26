package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"log"
)

type User struct {
	Id uint64 `gorm:"primary_key column:id"`
	Name string `gorm:"column:name"`
	Money uint64 `gorm:"column:money"`
	Meta []byte `gorm:"column:meta"`
}


func (User) TableName() string  {
	return "User"
}


func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	//m, err := json.Marshal([]string{"xx", "xfx"})
	m, err := json.Marshal(map[string]string{"xx": "y"})
	if err != nil {
		log.Fatal(err.Error())
	}
	//db.Begin()
	//err =db.Unscoped().Save(&User{Id: 1, Name: "xx", Money: 1, Meta: m}).Error
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//db.Commit()

	db.Begin()
	var users = []User{{Name: "xx", Money: 1, Meta: m},
		{Name: "yy", Money: 1, Meta: m}}
	db.Create(&users)
	for _, user1 := range users {
		fmt.Println(user1.Id)
	}
	db.Commit()
}
