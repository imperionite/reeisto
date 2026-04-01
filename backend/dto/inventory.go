package dto

type CreateInventoryDTO struct {
	ElementID         uint    `json:"element_id" binding:"required"`
	Quantity          float64 `json:"quantity" binding:"required,gte=0"`
	WarehouseLocation string  `json:"warehouse_location" binding:"required,min=2,max=100"`
}