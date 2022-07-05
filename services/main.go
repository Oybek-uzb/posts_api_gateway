package services

import (
	"fmt"
	"github.com/Oybek-uzb/posts_api_gateway/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	PostsCRUDService() bond_service.BondServiceClient
}

type grpcClients struct {
	postsCRUDService bond_service.BondServiceClient
}

func NewGrpcClients(conf *config.Config) (*grpcClients, error) {
	postsCRUDService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostsCRUDServiceHost, conf.PostsCRUDServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		postsCRUDService: bond_service.NewBondServiceClient(postsCRUDService),
	}, nil
}
