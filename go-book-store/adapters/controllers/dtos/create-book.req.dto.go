package dtos

type CreateBookReq struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
