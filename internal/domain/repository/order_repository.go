package repository

import "github.com/anawatj/food-ordering/internal/domain/model"

type OrderRepository interface {
	Save(order *model.Order) error
	GetByID(id string) (*model.Order, error)
	GetAll() ([]*model.Order, error)
}
