package model

import (
	"time"
)

type Link struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Table0       string    `gorm:"type:varchar(30);"`
	ID0          uint      `gorm:"type:int;"`
	Table1       string    `gorm:"type:varchar(30);"`
	ID1          uint      `gorm:"type:int;"`
	ShortName    string    `gorm:"type:varchar(20);"`
	Describe     string    `gorm:"type:varchar(300);"`
	Risk         string    `gorm:"type:varchar(300);"`
	IsDeprecated bool      `gorm:"type:boolean;default:false"`
	CreateTime   time.Time `gorm:"type:timestamp;default:current_timestamp;not null"`
	UpdateTime   time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp;not null"`
}
