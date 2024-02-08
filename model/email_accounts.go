package model

import (
	"time"
)

type EmailAccount struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	Address       string    `gorm:"type:varchar(350);not null;comment:'the part of the email address before @'"`
	Password      string    `gorm:"type:varchar(350);"`
	SafeQuestions string    `gorm:"type:varchar(600);"`
	Describe      string    `gorm:"type:varchar(300);"`
	Risk          string    `gorm:"type:varchar(300);"`
	FirstName     string    `gorm:"type:varchar(20);not null"`
	LastName      string    `gorm:"type:varchar(20);not null"`
	IsDeprecated  bool      `gorm:"type:boolean;default:false"`
	CreateTime    time.Time `gorm:"type:timestamp;default:current_timestamp;not null"`
	UpdateTime    time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp;not null"`
}
