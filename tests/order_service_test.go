package tests

import (
	"testing"

	"github.com/anawatj/food-ordering/internal/domain/model"
	"github.com/anawatj/food-ordering/internal/domain/service"
	"github.com/anawatj/food-ordering/internal/infrastructure/memory"

	"github.com/stretchr/testify/assert"
)

func setupServiceWithFoods() (service.OrderService, *memory.FoodRepo, *memory.OrderRepo) {
	foodRepo := memory.NewFoodRepo()
	orderRepo := memory.NewOrderRepo()

	foodRepo.Seed([]*model.Food{
		{ID: "1", Name: "ข้าวผัด", Quantity: 10, Price: 50},
		{ID: "2", Name: "ผัดไทย", Quantity: 1, Price: 60},
	})

	svc := service.NewOrderService(foodRepo, orderRepo)
	return svc, foodRepo, orderRepo
}

func TestCreateOrder_Success(t *testing.T) {
	svc, foodRepo, _ := setupServiceWithFoods()

	order, err := svc.CreateOrder("o1", []model.CartItem{{FoodID: "1", Qty: 2}}, "4242-4242-4242-4242")
	assert.NoError(t, err)
	assert.Equal(t, "paid", order.Status)
	assert.Equal(t, float64(100), order.Total)

	food, _ := foodRepo.GetByID("1")
	assert.Equal(t, 8, food.Quantity)
}

func TestCreateOrder_StockNotEnough(t *testing.T) {
	svc, _, _ := setupServiceWithFoods()

	order, err := svc.CreateOrder("o2", []model.CartItem{{FoodID: "2", Qty: 2}}, "4242-4242-4242-4242")
	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Contains(t, err.Error(), "ไม่เพียงพอในคลัง")
}

func TestCreateOrder_FakeCard_Fail(t *testing.T) {
	svc, _, _ := setupServiceWithFoods()

	order, err := svc.CreateOrder("o3", []model.CartItem{{FoodID: "1", Qty: 1}}, "2222-2222-2222-2222")
	assert.NoError(t, err)
	assert.Equal(t, "failed", order.Status)
	assert.Equal(t, float64(50), order.Total)
}

func TestCreateOrder_FakeCard_Invalid(t *testing.T) {
	svc, _, _ := setupServiceWithFoods()

	order, err := svc.CreateOrder("o3", []model.CartItem{{FoodID: "1", Qty: 1}}, "2222-2222-2222-2223")
	assert.NoError(t, err)
	assert.Equal(t, "failed", order.Status)
	assert.Equal(t, float64(50), order.Total)
}

func TestGetOrder_Success(t *testing.T) {
	svc, _, _ := setupServiceWithFoods()

	order, _ := svc.CreateOrder("o4", []model.CartItem{{FoodID: "1", Qty: 1}}, "4242-4242-4242-4242")
	fetched, err := svc.GetOrder("o4")

	assert.NoError(t, err)
	assert.Equal(t, order.ID, fetched.ID)
}

func TestGetOrder_NotFound(t *testing.T) {
	svc, _, _ := setupServiceWithFoods()

	order, err := svc.GetOrder("not-exist")
	assert.Nil(t, order)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "order not found")
}

func TestGetAllOrders(t *testing.T) {
	svc, _, _ := setupServiceWithFoods()

	svc.CreateOrder("o5", []model.CartItem{{FoodID: "1", Qty: 1}}, "4242-4242-4242-4242")
	svc.CreateOrder("o6", []model.CartItem{{FoodID: "1", Qty: 1}}, "4242-4242-4242-4242")

	orders, err := svc.GetAllOrders()
	assert.NoError(t, err)
	assert.Len(t, orders, 2)
}

func TestGetFoods(t *testing.T) {
	svc, _, _ := setupServiceWithFoods()

	foods, err := svc.GetFoods()
	assert.NoError(t, err)
	assert.Len(t, foods, 2)
}
