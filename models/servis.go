package models

type Servis struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	NamaServis  string `json:"nama_servis"`
	Deskripsi   string `json:"deskripsi"`
	Harga       int    `json:"harga"`
}