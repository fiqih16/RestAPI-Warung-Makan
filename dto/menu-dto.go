package dto

type MenuCreateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	NamaMenu string `json:"nama_menu" form:"nama_menu" binding:"required"`
	Harga    int    `json:"harga" form:"harga" binding:"required"`
	Status   string `json:"status" form:"status" binding:"required"`
}