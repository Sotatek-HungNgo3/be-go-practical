package service

import (
	"context"
	"github.com/Sotatek-HungNgo3/be-practical-order/config"
	"github.com/Sotatek-HungNgo3/be-practical-order/dto"
	"github.com/Sotatek-HungNgo3/be-practical-order/enum"
	"github.com/Sotatek-HungNgo3/be-practical-order/model"
	"github.com/Sotatek-HungNgo3/be-practical-order/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"time"
)

type OrderService struct {
	paymentClient pb.PaymentServiceClient
	db            *gorm.DB
}

func NewOrderService() *OrderService {
	conn, err := grpc.DialContext(context.Background(), config.GetPaymentConfig(), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	dbConnection, dbConnectionError := model.InitDb()
	if dbConnectionError != nil {
		log.Fatalf("DB did not connect: %v", dbConnectionError)
	}

	return &OrderService{db: dbConnection, paymentClient: pb.NewPaymentServiceClient(conn)}
}

func (orderService *OrderService) MakeExternalPayment(ctx context.Context, orderDto dto.OrderDto) (*pb.MakePaymentResponse, error) {
	var method pb.PaymentMethod
	switch orderDto.PaymentMethod {
	case enum.PaymentMethodCash:
		method = pb.PaymentMethod_Cash
		break
	case enum.PaymentMethodCard:
		method = pb.PaymentMethod_CreditCard
		break
	}

	orderAmount := orderDto.UnitPrice * float64(orderDto.Quantity)
	req := &pb.MakePaymentRequest{
		ProductName: orderDto.ProductName,
		OrderAmount: orderAmount,
		Method:      method,
	}
	return orderService.paymentClient.MakePayment(ctx, req)
}

func (orderService *OrderService) GetAllOrder() []model.Order {
	var allOrders []model.Order
	orderService.db.Find(&allOrders)
	return allOrders
}

func (orderService *OrderService) GetOrderById(orderId uint) (model.Order, error) {
	var orderDetails model.Order
	err := orderService.db.First(&orderDetails, orderId).Error
	return orderDetails, err
}

func (orderService *OrderService) CreateOrder(ctx *gin.Context, orderDto dto.OrderDto) (model.Order, error) {
	// Draft created order
	order := model.Order{
		ProductName:   orderDto.ProductName,
		UnitPrice:     orderDto.UnitPrice,
		Quantity:      orderDto.Quantity,
		Status:        enum.OrderStatusCreated,
		PaymentMethod: orderDto.PaymentMethod.ToString(),
	}
	orderService.db.Create(&order)

	// Call to payment to makePayment
	paymentRes, paymentErr := orderService.MakeExternalPayment(ctx, orderDto)
	if paymentErr != nil {
		orderService.db.Model(&order).Updates(model.Order{Status: enum.OrderStatusFailed, ErrorMessage: "Failed to make the payment"})
		return order, nil
	}
	if paymentRes.Status == pb.PaymentStatus_Confirmed {
		orderService.db.Model(&order).Updates(model.Order{Status: enum.OrderStatusConfirmed})
	} else {
		orderService.db.Model(&order).Updates(model.Order{Status: enum.OrderStatusCancelled})
	}

	// Add sleep logic in goroutine
	if order.Status == enum.OrderStatusConfirmed {
		go func(orderId uint) {
			time.Sleep(10 * time.Second)
			orderService.db.Model(&model.Order{}).Where("id = ?", orderId).Updates(model.Order{Status: enum.OrderStatusDelivered})
		}(order.ID)
	}

	return order, nil // return immediately
}

func (orderService *OrderService) CancelOrder(orderId uint) (model.Order, error) {
	orderDetails, err := orderService.GetOrderById(orderId)
	orderService.db.Model(&orderDetails).Updates(model.Order{Status: enum.OrderStatusCancelled})
	return orderDetails, err
}
