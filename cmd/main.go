package main

import (
	"github.com/Oybek-uzb/posts_api_gateway/api"
	"github.com/Oybek-uzb/posts_api_gateway/config"
	"github.com/Oybek-uzb/posts_api_gateway/services"
)

func main() {

	cfg := config.Load()

	grpcClients, _ := services.NewGrpcClients(&cfg)

	server := api.NewServer(grpcClients)

	err := server.Run(cfg.HttpPort)

	if err != nil {
		return
	}
}
