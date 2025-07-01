package memory

import (
	"errors"
	"sync"

	"github.com/anawatj/food-ordering/internal/domain/model"
)

type OrderRepo struct {
	mu     sync.RWMutex
	orders map[string]*model.Order
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{
		orders: make(map[string]*model.Order),
	}
}

func (r *OrderRepo) Save(order *model.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders[order.ID] = order
	return nil
}

func (r *OrderRepo) GetByID(id string) (*model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	order, ok := r.orders[id]
	if !ok {
		return nil, errors.New("order not found")
	}
	return order, nil
}

func (r *OrderRepo) GetAll() ([]*model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*model.Order, 0, len(r.orders))
	for _, o := range r.orders {
		list = append(list, o)
	}
	return list, nil
}
