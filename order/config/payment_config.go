package config

import "os"

func GetPaymentConfig() string {
	grpcPaymentHost := os.Getenv("GRPC_PAYMENT_HOST")
	grpcPaymentPort := os.Getenv("GRPC_PAYMENT_PORT")

	return grpcPaymentHost + ":" + grpcPaymentPort
}
