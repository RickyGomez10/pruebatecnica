package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"pruebatecnica.com/diasFestivos"
	"pruebatecnica.com/diasFestivos/get"
)

func CreateNewDiasFestivosEndpoint(service *get.DiasFestivosService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(diasFestivos.DiasFestivosRequest)
		resp := service.GetByTypeAndDateRange(req)
		return resp, nil
	}
}
