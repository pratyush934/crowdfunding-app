package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pratyush934/crowdfunding-app/user-service/database"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
)

type User struct {
	ID            uuid.UUID `gorm:"type:char(36);primary_key"`
	Name          string    `gorm:"type:varchar(100);not null"`
	UserName      string    `gorm:"type:varchar(100);not null;unique"`
	Email         string    `gorm:"type:varchar(100);not null;unique"`
	Password      string    `gorm:"type:longtext;not null"`
	Phone         string    `gorm:"type:varchar(10);unique"`
	WalletAddress string    `gorm:"type:varchar(191);not null;unique"`
	PublicKey     string    `gorm:"type:longtext;not null"`
	KYCStatus     string    `gorm:"type:varchar(191);default:'pending'"`
	GovtIDHash    string    `gorm:"type:longtext;not null"`
	RoleID        int64     `gorm:"foreignKey:RoleId;default:1"`
	Role          Role      `gorm:"foreignKey:RoleId"`
	TokenBalance  float64   `gorm:"default:0"`
	BondID        string    `gorm:"type:json"`
	StakingAmount float64   `gorm:"default:0"`
	Status        string    `gorm:"type:varchar(191);default:'active'"`
}

func (user *User) BeforeCreate() (err error) {
	user.ID = uuid.New()
	return
}

func (user *User) Save() (*User, error) {
	err := database.DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave() error {
	//generating password
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error exist in the user.go in the function BeforeSave()")
		return err
	}

	user.Password = string(password)
	user.UserName = html.EscapeString(strings.TrimSpace(user.UserName))

	return err
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func GetUsers(user *[]User) (err error) {
	err = database.DB.Find(&user).Error

	if err != nil {
		fmt.Println("Error exist in the GetUsers method in user.go please check")
		return err
	}
	return nil
}

func GetUserByName(userName string) (User, error) {
	var user User
	err := database.DB.Where("username=?", userName).Find(&user).Error

	if err != nil {
		fmt.Println("Error exist in the GerUserByName method in user.go please check")
		return User{}, err
	}
	return user, nil
}

func GetUserById(userId string) (User, error) {
	var user User
	err := database.DB.Where("id=?", userId).Find(&user).Error

	if err != nil {
		fmt.Println("Error exist in GetUserById method in user.go please look at this")
		return User{}, err
	}
	return user, nil
}

func GetUser(user *User, userId string) error {
	err := database.DB.Where("id=?", userId).Find(user).Error

	if err != nil {

		fmt.Println("Error exist in GetUser method in user.go please look at this")
		return err
	}

	return nil
}

func UpdateUser(User *User) error {
	err := database.DB.Omit("password").Updates(User).Error

	if err != nil {
		fmt.Println("Error exist in UpdateUser method in user.go please look at this")
		return err
	}
	return nil
}
