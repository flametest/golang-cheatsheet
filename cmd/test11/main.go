package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"log"
)

type Meta struct {
	x string `json:"x"`
}

type User struct {
	Meta *Meta `json:"meta"`
}


func (m *Meta) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	srcBytes, ok := src.([]byte)
	if !ok {
		return errors.New("only support []byte type")
	}
	return json.Unmarshal(srcBytes, m)
}

func (m *Meta) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

type UserDO struct {
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

	db.Begin()

	user := &User{Meta: &Meta{x: "y"}}
	db.Create(user)
	db.Commit()
}
