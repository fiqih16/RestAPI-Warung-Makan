package model

type Menu struct {
	ID         uint64   `gorm:"primary_key:auto_increment" json:"id"`
	CustomerID uint64   `gorm:"not null" json:"-"`
	Customer   Customer `gorm:"foreignkey:CustomerID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"customer"`
	NamaMenu   string   `gorm:"type:varchar(255)" json:"nama_menu"`
	Harga      int      `gorm:"type:int" json:"harga"`
	Status     string   `gorm:"type:varchar(255)" json:"status"`
}