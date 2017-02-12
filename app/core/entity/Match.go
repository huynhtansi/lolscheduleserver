package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)


type MatchStatus uint

const (
	NotStarted 	MatchStatus = 1
	Playing		MatchStatus = 2
	Ended		MatchStatus = 3
)

type MatchType string

const (
	Bo1	MatchType = "Best of 1"
	Bo2	MatchType = "Best of 2"
	Bo3	MatchType = "Best of 3"
	Bo5	MatchType = "Best of 5"
)

type Match struct {
	gorm.Model

	StartTime	time.Time	`gorm:"not null"`
	Team1		Team
	Team2		Team
	Status		MatchStatus	`gorm:"not null"`
	Type		MatchType	`gorm:"not null"`
}
