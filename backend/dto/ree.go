package dto

type CreateREEDTO struct {
	Name        string  `json:"name" binding:"required,min=2,max=100"`
	Symbol      string  `json:"symbol" binding:"required,min=1,max=10"`
	Category    string  `json:"category" binding:"required,oneof=Light REE Heavy REE"`
	MarketPrice float64 `json:"market_price" binding:"required,gt=0"`
<<<<<<< HEAD

	Form      string  `json:"form" binding:"omitempty,oneof=oxide metal alloy"`
	PriceUnit string  `json:"price_unit" binding:"omitempty"`
	Purity    float64 `json:"purity" binding:"omitempty,gt=0,lt=100"`
=======
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
}

type UpdateREEDTO struct {
	Name        string  `json:"name" binding:"omitempty,min=2,max=100"`
	Symbol      string  `json:"symbol" binding:"omitempty,min=1,max=10"`
	Category    string  `json:"category" binding:"omitempty,oneof=Light REE Heavy REE"`
	MarketPrice float64 `json:"market_price" binding:"omitempty,gt=0"`
<<<<<<< HEAD

	Form      string  `json:"form" binding:"omitempty,oneof=oxide metal alloy"`
	PriceUnit string  `json:"price_unit" binding:"omitempty"`
	Purity    float64 `json:"purity" binding:"omitempty,gt=0,lt=100"`
}
=======
}
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
