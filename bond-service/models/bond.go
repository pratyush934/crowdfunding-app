package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pratyush934/crowdfunding-app/bond-service/dbBond"
	"gorm.io/gorm"
	"time"
)

type Bond struct {
	ID           uuid.UUID `gorm:"type:char(36);primary_key"`
	SerialNumber string    `gorm:"type:varchar(100);not null"`
	UserId       uuid.UUID `gorm:"type:char(36)"`
	Price        float64   `gorm:"type:float"`
	Status       string    `gorm:"type:varchar(191);default:'active'"`
	MinedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	RedeemedAt   *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (bond *Bond) BeforeCreate(tx *gorm.DB) (err error) {
	bond.ID = uuid.New()
	return
}

func (bond *Bond) Save() (*Bond, error) {
	err := dbBond.DB.Save(&bond).Error
	if err != nil {
		fmt.Println("There is an error in Save method in bond.go/bond-service")
		return &Bond{}, err
	}
	return bond, nil
}

func UpdateBond(bond *Bond) (*Bond, error) {
	err := dbBond.DB.Save(&bond).Error

	if err != nil {
		fmt.Println("There is an error in Update Method the Bond in bong.go/bond-service")
		return &Bond{}, err
	}
	return bond, nil
}

func GetBonds(bonds *[]Bond) error {
	err := dbBond.DB.Find(&bonds).Error

	if err != nil {
		fmt.Println("There is an error in GetBonds method in bond.go/bond-service")
		return err
	}
	return nil
}

func GetBondById(id string) (Bond, error) {
	var bond Bond
	err := dbBond.DB.Where("id=?", id).First(&bond).Error

	if err != nil {
		fmt.Println("There is an error in GetBondById method in bond.go/bond-service")
		return Bond{}, err
	}
	return bond, err
}

func GetBondByUserId(id string) ([]Bond, error) {
	var bond []Bond
	err := dbBond.DB.Where("user_id=?", id).First(&bond).Error

	if err != nil {
		fmt.Println("There is an error in GetBondByUserId method in bond.go/bond-service")
		return nil, err
	}
	return bond, nil
}
