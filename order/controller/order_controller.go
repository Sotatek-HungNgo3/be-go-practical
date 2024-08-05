package controller

import (
	"github.com/Sotatek-HungNgo3/be-practical-order/dto"
	"github.com/Sotatek-HungNgo3/be-practical-order/enum"
	"github.com/Sotatek-HungNgo3/be-practical-order/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{orderService: service.NewOrderService()}
}

// GetAllOrder
// @Summary Get a list of orders
// @Description Get a list of orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} model.Order
// @Router /orders [get]
func (orderController *OrderController) GetAllOrder(ctx *gin.Context) {
	allOrders := orderController.orderService.GetAllOrder()
	ctx.JSON(http.StatusOK, allOrders)
}

// GetOrderById
// @Summary Get order by id
// @Description Get order by id
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} model.Order
// @Router /orders/{id} [get]
func (orderController *OrderController) GetOrderById(ctx *gin.Context) {
	orderId := PrepareOrderId(ctx)
	orderDetails, err := orderController.orderService.GetOrderById(orderId)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, orderDetails)
}

// CreateOrder
// @Summary Create order
// @Description Create order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body dto.OrderDto true "New order"
// @Success 200 {object} model.Order
// @Router /orders [post]
func (orderController *OrderController) CreateOrder(ctx *gin.Context) {
	var orderDto dto.OrderDto

	// Validate dto
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("paymentMethod", enum.ValidatePaymentMethod)
	}
	if err := ctx.ShouldBindJSON(&orderDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderDetails, err := orderController.orderService.CreateOrder(ctx, orderDto)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, orderDetails)
}

// CancelOrder
// @Summary Cancel order
// @Description Cancel order
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} model.Order
// @Router /orders/{id} [patch]
func (orderController *OrderController) CancelOrder(ctx *gin.Context) {
	orderId := PrepareOrderId(ctx)
	orderDetails, err := orderController.orderService.CancelOrder(orderId)
	if err != nil {
		handleError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, orderDetails)
}

func PrepareOrderId(ctx *gin.Context) uint {
	orderIdParam := ctx.Param("id")
	orderId, err := strconv.ParseUint(orderIdParam, 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Order ID is not valid"})
	}
	return uint(orderId)
}

func handleError(context *gin.Context, err error) {
	switch err {
	case gorm.ErrRecordNotFound:
		context.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
	default:
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
