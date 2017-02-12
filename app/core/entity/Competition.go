package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Competition struct {
	gorm.Model

	Name		string		`gorm:"not null"`
	ShortName	string		`gorm:"not null"`
	Location	string		`gorm:"not null"`
	StartDate	time.Time	`gorm:"not null"`
	EndDate		time.Time	`gorm:"not null"`
	Desciption	string		`gorm:"not null"`

}
