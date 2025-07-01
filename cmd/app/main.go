package main

import (
	"github.com/anawatj/food-ordering/internal/domain/model"
	"github.com/anawatj/food-ordering/internal/domain/service"
	"github.com/anawatj/food-ordering/internal/infrastructure/api"
	"github.com/anawatj/food-ordering/internal/infrastructure/memory"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	foodRepo := memory.NewFoodRepo()
	orderRepo := memory.NewOrderRepo()

	foodRepo.Seed([]*model.Food{
		{ID: "1", Name: "ข้าวผัด", Quantity: 10, Price: 50},
		{ID: "2", Name: "ผัดไทย", Quantity: 5, Price: 60},
	})

	orderService := service.NewOrderService(foodRepo, orderRepo)
	orderHandler := api.NewOrderHandler(orderService)

	r.GET("/foods", orderHandler.GetFoods)
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.GET("/orders", orderHandler.GetAllOrders)

	r.Run(":8080")
}
