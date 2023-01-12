package repository

import (
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"pruebatecnica.com/diasFestivos"
)

type (
	Repository struct {
		logger *zap.Logger
	}
)

func NewRepository(logger *zap.Logger) *Repository {
	return &Repository{
		logger: logger,
	}
}

func (r *Repository) GetData() diasFestivos.DiasFestivosResponse {
	var diasFestivosData diasFestivos.DiasFestivosResponse
	response, err := http.Get("https://api.victorsanmartin.com/feriados/en.json")

	if err != nil {
		r.logger.Error("Error while trying to get data from dias feriados endpoint" + err.Error())
		return diasFestivos.DiasFestivosResponse{}
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		r.logger.Error("Error reading data from dias feriados endpoint" + err.Error())
		return diasFestivos.DiasFestivosResponse{}
	}

	errData := json.Unmarshal(responseData, &diasFestivosData)
	if errData != nil {
		r.logger.Error("Error unmarshaling data from dias feriados endpoint: " + errData.Error())
		return diasFestivos.DiasFestivosResponse{}
	}
	return diasFestivosData

}
