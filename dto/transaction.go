package dto

type TransactionCreateDTO struct {
	CustomerID uint64 `json:"customer_id,omitempty" form:"customer_id,omitempty"`
	MenuID     uint64 `json:"menu_id,omitempty" form:"menu_id,omitempty"`
	Tanggal    string `json:"tanggal" form:"tanggal" binding:"required"`
	// JumlahBeli int    `json:"jumlah_beli" form:"jumlah_beli" binding:"required"`
	// TotalHarga int    `json:"total_harga" form:"total_harga" binding:"required"`
}