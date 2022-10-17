package model

import "time"

type Transaction struct {
	ID         uint64   `gorm:"primary_key:auto_increment" json:"id"`
	CustomerID uint64   `gorm:"not null" json:"-"`
	Customer   Customer `gorm:"foreignKey:CustomerID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"customer"`
	MenuID     uint64   `gorm:"not null" json:"-"`
	Menu       Menu     `gorm:"foreignKey:MenuID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"menu"`
	Tanggal    time.Time `json:"tanggal"`
	JumlahBeli int `gorm:"type:int" json:"jumlah_beli"`
	TotalBayar int `gorm:"type:int" json:"total_bayar"`
}