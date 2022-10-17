package dto

type MenuCreateDTO struct {
	// ID         uint64 `json:"id" form:"id" binding:"required"`
	NamaMenu   string `json:"nama_menu" form:"nama_menu" binding:"required"`
	Harga      int    `json:"harga" form:"harga" binding:"required"`
	Status     string `json:"status" form:"status" binding:"required"`
	CustomerID uint64 `json:"customer_id,omitempty" form:"customer_id,omitempty"`
}

type MenuUpdateDTO struct {
	ID         uint64 `json:"id" form:"id" binding:"required"`
	NamaMenu   string `json:"nama_menu" form:"nama_menu" binding:"required"`
	Harga      int    `json:"harga" form:"harga" binding:"required"`
	Status     string `json:"status" form:"status" binding:"required"`
	CustomerID uint64 `json:"customer_id,omitempty" form:"customer_id,omitempty"`
}