package model

import (
	"time"
)

type Cellphone struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Country      string    `gorm:"type:varchar(20);not null"`
	Brand        string    `gorm:"type:varchar(20);not null"`
	Model        string    `gorm:"type:varchar(20);not null"`
	System       string    `gorm:"type:varchar(20);not null"`
	Describe     string    `gorm:"type:varchar(300);"`
	Risk         string    `gorm:"type:varchar(300);"`
	IsDeprecated bool      `gorm:"type:boolean;default:false"`
	CreateTime   time.Time `gorm:"type:timestamp;default:current_timestamp;not null"`
	UpdateTime   time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp;not null"`
}
