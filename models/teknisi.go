package models

type Teknisi struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`
}