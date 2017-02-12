package entity

import "github.com/jinzhu/gorm"

type Team struct {

	gorm.Model

	Name		string		`gorm:"not null"`
	ShortName	string		`gorm:"not null"`
	NumOfMembers	int		`gorm:"not null"`
	AvatarURL	string		`gorm:"not null"`
	Overview	string
	Sponsors	string		`gorm:"not null"`
}
