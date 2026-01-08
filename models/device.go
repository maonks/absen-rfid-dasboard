package models

import (
	"time"
)

type Device struct {
	Id       uint   `gorm:"primaryKey"`
	DeviceId string `gorm:"uniqueIndex"`
	LastSeen time.Time
}
