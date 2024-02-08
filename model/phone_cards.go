package model

import (
	"time"
)

type PhoneCard struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Number       string    `gorm:"type:varchar(350)"`
	Country      string    `gorm:"type:varchar(20)"`
	Company      string    `gorm:"type:varchar(10)"`
	CardPaidWay  string    `gorm:"type:varchar(40)"`
	CardTopUpWay string    `gorm:"type:varchar(40)"`
	CanTextIn    bool      `gorm:"default:false"`
	CanTextOut   bool      `gorm:"default:false"`
	CanCallIn    bool      `gorm:"default:false"`
	CanCallOut   bool      `gorm:"default:false"`
	CanInternet  bool      `gorm:"default:false"`
	Charge       string    `gorm:"type:varchar(10)"`
	IsAutoCharge bool      `gorm:"default:false"`
	Describe     string    `gorm:"type:varchar(300);"`
	Risk         string    `gorm:"type:varchar(300);"`
	IsDeprecated bool      `gorm:"type:boolean;default:false"`
	SimInfo      string    `gorm:"type:varchar(350);comment:'{\"puk_code\": \"str\", \"sim_number\": \"str\", \"bar_code\": \"str\"}'"`
	CreateTime   time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdateTime   time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp"`
}
