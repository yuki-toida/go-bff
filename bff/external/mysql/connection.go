package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func Connect(user, pass, host, port, name string) *gorm.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)

	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err.Error())
	}

	return db
}
