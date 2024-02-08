package model

import (
	"time"
)

type Account struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	Account       string    `gorm:"type:varchar(350);"`
	UserName      string    `gorm:"type:varchar(30);"`
	Password      string    `gorm:"type:varchar(350);"`
	SafeQuestions string    `gorm:"type:varchar(600);"`
	ShortName     string    `gorm:"type:varchar(20);"`
	LoginWebsite  string    `gorm:"type:varchar(100);"`
	FirstName     string    `gorm:"type:varchar(20);"`
	LastName      string    `gorm:"type:varchar(20);"`
	Bod           string    `gorm:"type:varchar(20);"`
	Describe      string    `gorm:"type:varchar(300);"`
	Risk          string    `gorm:"type:varchar(300);"`
	IsDeprecated  bool      `gorm:"type:boolean;default:false"`
	CreateTime    time.Time `gorm:"type:timestamp;default:current_timestamp;not null"`
	UpdateTime    time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp;not null"`
}
