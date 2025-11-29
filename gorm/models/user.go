package models

import (
	"errors"

	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "Male"
	Femail Gender = "Femail"
)

type User struct {
	ID        uint    `gorm:"primaryKey"`
	FirstName string  `gorm:"size:100;not null"`
	LastName  string  `gorm:"size:100;not null"`
	Email     *string `gorm:"size:30;not null;unique"`
	Mobile    string  `gorm:"size:11;not null;unique"`
	Age       uint    `gorm:"not null"`
	Gender    Gender  `gorm:"type:gender"`
	IsActive  bool    `gorm:"default:true"`
	gorm.Model
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Gender != Male && u.Gender != Femail {
		return errors.New("Invalid Gender Type")
	}

	var existsUser User
	err = tx.Where("mobile = ?", u.Mobile).First(existsUser).Error
	if err != nil && existsUser.ID>0{
		return errors.New("mobile number allready exists")
	}

	return nil

}
