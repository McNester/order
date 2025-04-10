package handlers

import (
	"net/http"
	"orders/models"
	"orders/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{service: services.NewOrderService()}
}

func (h *OrderHandler) SaveOrder(c *gin.Context) {

	order := models.Order{}

	if err := c.ShouldBind(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong body for order",
		})
		return
	}

	savedOrder, err := h.service.SaveOrder(&order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, savedOrder)

}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Wrong id!",
			})
		return
	}

	order := models.Order{}

	if err := c.ShouldBind(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong body for order",
		})
		return
	}

	updatedOrder, err := h.service.UpdateOrder(id, &order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedOrder)

}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Wrong id!",
			})
		return
	}

	order, err := h.service.GetOrder(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, order)

}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.service.ListOrder()

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, orders)

}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{
				"error": "Wrong id!",
			})
		return
	}

	err = h.service.DeleteOrder(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted the order",
	})

}
