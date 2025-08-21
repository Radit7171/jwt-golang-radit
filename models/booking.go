package models

type Booking struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	TeknisiID     uint   `json:"teknisi_id"`
	ServisID      uint   `json:"servis_id"`
	TanggalBooking string `json:"tanggal_booking"`
	Status        string `json:"status"` // pending, dikerjakan, selesai
}