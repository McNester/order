package main

import (
	"orders/handlers"

	"github.com/gin-gonic/gin"
)

var (
	order_handler *handlers.OrderHandler
)

func main() {
	router := gin.Default()

	order_handler = handlers.NewOrderHandler()

	router.POST("/orders", order_handler.SaveOrder)

	router.GET("/orders/:id", order_handler.GetOrder)

	router.PATCH("/orders/:id", order_handler.UpdateOrder)

	router.DELETE("/orders/:id", order_handler.DeleteOrder)

	router.GET("/orders", order_handler.ListOrders)

	router.Run("0.0.0.0:8083")
}
