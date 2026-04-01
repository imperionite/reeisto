package dto

type CreateTransactionDTO struct {
	ElementID uint    `json:"element_id" binding:"required"`
	Type      string  `json:"type" binding:"required,oneof=buy sell"`
	Quantity  float64 `json:"quantity" binding:"required,gt=0"`
	Price     float64 `json:"price" binding:"required,gt=0"`
<<<<<<< HEAD
	WarehouseLocation  string `json:"warehouse_location"`
=======
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
}