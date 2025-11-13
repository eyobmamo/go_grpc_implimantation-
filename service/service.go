package service

import (
	// GRPCInterface "GRPCX/grpc"
	"context"
)

type ServiceReposne struct {
	ID   int
	Name string
	Code string
}

type ConfigServiceImpl struct{}

type ConfigService interface {
	GetServiceList(ctx context.Context) ([]ServiceReposne, error)
	GetServiceDetail(ctx context.Context, serviceID string) (ServiceReposne, error)
}

func NewServiceInstance(ctx context.Context) *ConfigServiceImpl {
	return &ConfigServiceImpl{}
}

func (c ConfigServiceImpl) GetServiceList(ctx context.Context) ([]ServiceReposne, error) {
	return []ServiceReposne{
		{1, "travle", "TR"},
		{2, "work", "WR"},
	}, nil
}
func (c ConfigServiceImpl) GetServiceDetail(ctx context.Context, ServiceID string) (ServiceReposne, error) {
	return ServiceReposne{1, "travel", "TR"}, nil
}
