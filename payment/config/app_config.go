package config

import "os"

func GetGRPCPort() string {
	return os.Getenv("GRPC_PORT")
}
