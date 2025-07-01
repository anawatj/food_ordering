package memory

import (
	"errors"
	"sync"

	"github.com/anawatj/food-ordering/internal/domain/model"
)

type FoodRepo struct {
	mu    sync.RWMutex
	foods map[string]*model.Food
}

func NewFoodRepo() *FoodRepo {
	return &FoodRepo{
		foods: make(map[string]*model.Food),
	}
}

func (r *FoodRepo) Seed(foods []*model.Food) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, f := range foods {
		r.foods[f.ID] = f
	}
}

func (r *FoodRepo) GetAll() ([]*model.Food, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*model.Food, 0, len(r.foods))
	for _, f := range r.foods {
		list = append(list, f)
	}
	return list, nil
}

func (r *FoodRepo) GetByID(id string) (*model.Food, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	f, ok := r.foods[id]
	if !ok {
		return nil, errors.New("food not found")
	}
	return f, nil
}

func (r *FoodRepo) UpdateQuantity(id string, qty int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	f, ok := r.foods[id]
	if !ok {
		return errors.New("food not found")
	}
	f.Quantity = qty
	return nil
}
