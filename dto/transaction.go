package dto

type TransactionCreateDTO struct {
	MenuID     uint64 `json:"menu_id" form:"menu_id" binding:"required"`
	JumlahBeli int    `json:"jumlah_beli" form:"jumlah_beli" binding:"required"`
	CustomerID uint64 `json:"customer_id,omitempty" form:"customer_id,omitempty"`
	// Customer   model.Customer
}

type TransactionUpdateDTO struct {
	ID         uint64 `json:"id" form:"id"`
	MenuID     uint64 `json:"menu_id" form:"menu_id" binding:"required"`
	JumlahBeli int    `json:"jumlah_beli" form:"jumlah_beli" binding:"required"`
	CustomerID uint64 `json:"customer_id,omitempty" form:"customer_id,omitempty"`
}