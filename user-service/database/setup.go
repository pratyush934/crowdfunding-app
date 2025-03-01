package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	DB = connectDB()
	return DB
}

func connectDB() *gorm.DB {
	url := "root:Pratyush@123@/crowdfunding-app?charset=utf8&parseTime=True&loc=Local"

	//database, err := sql.Open("mysql", "root:Pratyush@123@tcp(localhost:3306)/crowdfunding-app")

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(`Failed while connecting the DataBase crowdfunding-app`)
	} else {
		fmt.Println("We have successfully connected to the Database")
	}

	return db
}
