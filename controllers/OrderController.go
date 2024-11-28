package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product-management-project/models"
	"product-management-project/services"
)

type OrderController struct {
	orderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (oc *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.OrderModel
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	newOrder, err := oc.orderService.CreateOrder(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tạo đơn hàng thành công", "order": newOrder})
}

func (oc *OrderController) FindAllOrders(ctx *gin.Context) {
	orders, err := oc.orderService.FindAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}
