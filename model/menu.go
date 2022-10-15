package model

type Menu struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	NamaMenu string `gorm:"type:varchar(255)" json:"nama_menu"`
	Harga    int    `gorm:"type:int" json:"harga"`
	Status   string `gorm:"type:varchar(255)" json:"status"`
}