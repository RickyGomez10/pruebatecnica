package pruebatecnica

import (
	"go.uber.org/zap"
	"log"
	"pruebatecnica.com/diasFestivos/get"
	"pruebatecnica.com/diasFestivos/platform/handler"
	repository2 "pruebatecnica.com/diasFestivos/platform/storage/repository"
	"pruebatecnica.com/kit/server"
)

const (
	PUERTO = "90"
	PATH   = "/v1/dias-festivos"
)

func Run() {

	logger, err := zap.NewProduction()

	if err != nil {
		log.Fatal(err)
	}
	repository := repository2.NewRepository(logger)
	service := get.NewDiasFestivosService(repository, logger)
	endpnt := handler.CreateNewDiasFestivosEndpoint(service)
	endpnt = handler.TransportMiddleware(logger)(endpnt)
	handlr := handler.NewHttpDiasFestivosHandler(PATH, endpnt)

	svc := server.NewServer()
	svc.RegisterRoutes(PATH, handlr)
	svc.Run(PUERTO)
}
