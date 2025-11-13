package handler

import (
	GRPCInterface "GRPCX/grpc"
	ConfigService "GRPCX/service"
	"context"
	// "strings"
)

type GrpcHandler struct {
	GRPCInterface.UnimplementedConfigServiceServer
	configService ConfigService.ConfigService
}

func NewGrpcHandler(configService ConfigService.ConfigService) *GrpcHandler {
	return &GrpcHandler{configService: configService}
}
func (G *GrpcHandler) GetServiceList(ctx context.Context, input *GRPCInterface.Empty) (*GRPCInterface.ServiceLists, error) {
	services, err := G.configService.GetServiceList(ctx)
	if err != nil {
		return nil, err
	}

	ServiceList := &GRPCInterface.ServiceLists{}
	for _, service := range services {
		ServiceList.Services = append(ServiceList.Services, &GRPCInterface.ServiceInfo{Id: string(service.ID), Name: service.Name, ServiceCode: service.Code})
	}

	return ServiceList, nil
}
func (G *GrpcHandler) GetServiceDetail(ctx context.Context, service *GRPCInterface.ServiceDetailRequest) (*GRPCInterface.ServiceDetialResponse, error) {
	services, err := G.configService.GetServiceDetail(ctx, service.Id)
	if err != nil {
		return nil, err
	}
	// serviceDetial := &GRPCInterface.ServiceDetialResponse{Service: &GRPCInterface.ServiceInfo{services.ID: ID, services.Name, services.Code}}

	serviceDetail := &GRPCInterface.ServiceDetialResponse{
		Service: &GRPCInterface.ServiceInfo{
			Id:          string(services.ID),
			Name:        services.Name,
			ServiceCode: services.Code,
		},
	}

	return serviceDetail, nil
}
