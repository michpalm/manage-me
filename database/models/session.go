package models

import "time"

type Session struct {
	ID        uint `gorm:"primary_key"`
	Username  string
	Timestamp time.Time
}
