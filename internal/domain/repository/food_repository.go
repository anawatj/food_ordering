package repository

import "github.com/anawatj/food-ordering/internal/domain/model"

type FoodRepository interface {
	GetAll() ([]*model.Food, error)
	GetByID(id string) (*model.Food, error)
	UpdateQuantity(id string, qty int) error
}
