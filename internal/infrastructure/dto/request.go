package dto

import "github.com/anawatj/food-ordering/internal/domain/model"

type CreateOrderRequest struct {
	OrderID    string           `json:"order_id" binding:"required"`
	Cart       []model.CartItem `json:"cart" binding:"required"`
	CardNumber string           `json:"card_number" binding:"required"`
}
