package models

import (
	"fmt"
	"github.com/pratyush934/crowdfunding-app/user-service/database"
)

type Role struct {
	ID          int64  `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(10);default:'user'"`
	Description string `gorm:"type:varchar(100)"`
	User        *User  `gorm:"foreignKey:RoleId"`
}

func CreateRole(role *Role) error {

	err := database.DB.Create(role).Error

	if err != nil {
		fmt.Println("Error exist in the method named CreateRole in role.go")
		return err
	}

	return nil
}

func GetRoles(role *[]Role) error {
	err := database.DB.Find(role).Error

	if err != nil {
		fmt.Println("Error exist in the method named GetRoles in role.go")
		return err
	}

	return nil
}

func GetRole(role *Role, roleId int64) error {
	err := database.DB.Where("id=?", roleId).Find(role).Error

	if err != nil {
		fmt.Println("Error exist in the method name GetRole in role.go")
		return err
	}

	return nil
}

func UpdateRole(role *Role) error {
	err := database.DB.Save(role).Error

	if err != nil {
		fmt.Println("Error exist in the method name UpdateRole in role.go")
		return err
	}
	return nil
}
