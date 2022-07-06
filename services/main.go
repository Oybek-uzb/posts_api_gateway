package services

import (
	"fmt"

	"github.com/Oybek-uzb/posts_api_gateway/config"
	posts_crud_service "github.com/Oybek-uzb/posts_api_gateway/pkg/api/posts_crud_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	PostsCRUDService() posts_crud_service.PostsCRUDServiceClient
}

type grpcClients struct {
	postsCRUDService posts_crud_service.PostsCRUDServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	postsCRUDService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostsCRUDServiceHost, conf.PostsCRUDServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		postsCRUDService: posts_crud_service.NewPostsCRUDServiceClient(postsCRUDService),
	}, nil
}

func (g grpcClients) PostsCRUDService() posts_crud_service.PostsCRUDServiceClient {
	return g.postsCRUDService
}
