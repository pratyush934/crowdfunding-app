package dbBond

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	DB, _ = connectDb()
	return DB
}

func connectDb() (*gorm.DB, error) {

	url := "root:Pratyush@123@/crowdfunding-app?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Println("So bad , it should not had happen")
		return nil, err
	}
	return db, err
}
