package service

import (
	"fmt"
	"sync"

	"github.com/anawatj/food-ordering/internal/domain/model"
	"github.com/anawatj/food-ordering/internal/domain/repository"
)

// Fake payment function as per requirement
func FakeCharge(cardNumber string, amount float64) (bool, string) {
	switch cardNumber {
	case "4242-4242-4242-4242":
		return true, "ตัดบัตรสำเร็จ"
	case "2222-2222-2222-2222":
		return false, "วงเงินเต็ม"
	default:
		return false, "บัตรไม่ถูกต้อง"
	}
}

type OrderService interface {
	GetFoods() ([]*model.Food, error)
	CreateOrder(orderID string, cart []model.CartItem, cardNumber string) (*model.Order, error)
	GetOrder(id string) (*model.Order, error)
	GetAllOrders() ([]*model.Order, error)
}

type orderService struct {
	foodRepo  repository.FoodRepository
	orderRepo repository.OrderRepository
	mu        sync.Mutex
}

func NewOrderService(foodRepo repository.FoodRepository, orderRepo repository.OrderRepository) OrderService {
	return &orderService{
		foodRepo:  foodRepo,
		orderRepo: orderRepo,
	}
}

func (s *orderService) GetFoods() ([]*model.Food, error) {
	return s.foodRepo.GetAll()
}

func (s *orderService) CreateOrder(orderID string, cart []model.CartItem, cardNumber string) (*model.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	total := 0.0
	for _, item := range cart {
		food, _ := s.foodRepo.GetByID(item.FoodID)
		/*if err != nil {
			return nil, err
		}*/
		if food.Quantity < item.Qty {
			return nil, fmt.Errorf("สินค้า %s ไม่เพียงพอในคลัง", food.Name)
		}
		total += float64(item.Qty) * food.Price
	}

	success, _ := FakeCharge(cardNumber, total)
	order := &model.Order{
		ID:     orderID,
		Items:  cart,
		Total:  total,
		Status: "failed",
	}

	if success {
		order.Status = "paid"
		for _, item := range cart {
			food, _ := s.foodRepo.GetByID(item.FoodID)
			s.foodRepo.UpdateQuantity(item.FoodID, food.Quantity-item.Qty)
		}
	}
	s.orderRepo.Save(order)
	/*if err := s.orderRepo.Save(order); err != nil {
		return nil, err
	}*/

	return order, nil
}

func (s *orderService) GetOrder(id string) (*model.Order, error) {
	return s.orderRepo.GetByID(id)
}

func (s *orderService) GetAllOrders() ([]*model.Order, error) {
	return s.orderRepo.GetAll()
}
