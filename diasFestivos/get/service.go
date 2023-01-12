package get

import (
	"go.uber.org/zap"
	"pruebatecnica.com/diasFestivos"
	"pruebatecnica.com/diasFestivos/platform/storage/repository"
	"time"
)

type (
	DiasFestivosService struct {
		repository *repository.Repository
		logger     *zap.Logger
	}
)

func NewDiasFestivosService(repository *repository.Repository, logger *zap.Logger) *DiasFestivosService {

	return &DiasFestivosService{
		repository: repository,
		logger:     logger,
	}
}

func (dfs *DiasFestivosService) GetByTypeAndDateRange(request diasFestivos.DiasFestivosRequest) []diasFestivos.DiasFestivos {
	diasFeriadosData := dfs.repository.GetData()
	var diasFeriadosFiltrados []diasFestivos.DiasFestivos
	for _, diaFeriado := range diasFeriadosData.Data {

		diaFeriadoDate, _ := time.Parse("2006-01-02", diaFeriado.Date)
		date := diaFeriadoDate.Unix()
		requestStartingDate, _ := time.Parse("2006-01-02", request.StartDate)
		date2 := requestStartingDate.Unix()
		requestEndingDate, _ := time.Parse("2006-01-02", request.EndDate)
		date3 := requestEndingDate.Unix()

		if diaFeriado.Type == request.Type && insideDateRange(date2, date3, date) {
			diasFeriadosFiltrados = append(diasFeriadosFiltrados, diaFeriado)
		}
	}
	return diasFeriadosFiltrados
}

func insideDateRange(startingDate int64, endingDate int64, diaFeriadoDate int64) bool {
	if diaFeriadoDate > startingDate && diaFeriadoDate < endingDate {
		return true
	}
	return false
}
