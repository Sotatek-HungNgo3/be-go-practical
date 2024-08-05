package enum

type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "Created"
	OrderStatusFailed    OrderStatus = "Failed"
	OrderStatusConfirmed OrderStatus = "Confirmed"
	OrderStatusCancelled OrderStatus = "Cancelled"
	OrderStatusDelivered OrderStatus = "Delivered"
)

func (orderStatus OrderStatus) ToString() string {
	switch orderStatus {
	case OrderStatusCreated:
		return "Created"
	case OrderStatusFailed:
		return "Failed"
	case OrderStatusConfirmed:
		return "Confirmed"
	case OrderStatusCancelled:
		return "Cancelled"
	case OrderStatusDelivered:
		return "Delivered"
	default:
		return "Unknown"
	}
}
