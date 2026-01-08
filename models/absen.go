package models

import "time"

type Absen struct {
	ID       uint `gorm:"primaryKey"`
	Uid      string
	DeviceId string
	Waktu    time.Time `gorm:"type:timestamp"`
}
