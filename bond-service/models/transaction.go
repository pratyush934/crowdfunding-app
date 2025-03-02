package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pratyush934/crowdfunding-app/bond-service/dbBond"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	Id              uuid.UUID `gorm:"type:char(36);primary_key"`
	UserId          uuid.UUID `gorm:"type:char(36);"`
	BondId          uuid.UUID `gorm:"type:char(36);"`
	Bond            Bond      `gorm:"foreignKey:BondId;references:ID"`
	Amount          float64   `gorm:"type:float"`
	Status          string    `gorm:"type:varchar(191);default:'pending'"`
	TransactionType string    `gorm:"type:varchar(100);default:'purchase'"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.Id = uuid.New()
	return
}

func (t *Transaction) Save() (*Transaction, error) {
	err := dbBond.DB.Create(&t).Error

	if err != nil {
		fmt.Println("There is an error in the Save method in transaction.go/bond-service")
		return &Transaction{}, err
	}

	return t, nil
}
