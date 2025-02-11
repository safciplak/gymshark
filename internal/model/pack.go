package model

type PackResponse struct {
	OrderAmount int         `json:"orderAmount"`
	Packs       map[int]int `json:"packs"`
	TotalItems  int         `json:"totalItems"`
}

type OrderRequest struct {
	OrderAmount int `json:"orderAmount"`
}

// Interfaces
type PackCalculatorStrategy interface {
	Calculate(orderAmount int, packSizes []int) map[int]int
}

type PackService interface {
	CalculatePacks(orderAmount int) PackResponse
}
