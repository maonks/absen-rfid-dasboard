package models

type Kartu struct {
	Id   uint   `gorm:"primaryKey"`
	Uid  string `gorm:"uniqueIndex"`
	Nama string
}
