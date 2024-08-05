package enum

import "github.com/go-playground/validator/v10"

type PaymentMethod string

const (
	PaymentMethodCard PaymentMethod = "Card"
	PaymentMethodCash PaymentMethod = "Cash"
)

func (paymentMethod PaymentMethod) ToString() string {
	switch paymentMethod {
	case PaymentMethodCash:
		return "Cash"
	case PaymentMethodCard:
		return "Card"
	default:
		return "Unknown"
	}
}

func ValidatePaymentMethod(fl validator.FieldLevel) bool {
	paymentMethod := fl.Field().String()
	switch PaymentMethod(paymentMethod) {
	case PaymentMethodCard, PaymentMethodCash:
		return true
	}
	return false
}
