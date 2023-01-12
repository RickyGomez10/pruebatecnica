package repository

import (
	"github.com/stretchr/testify/mock"
	"pruebatecnica.com/diasFestivos"
)

type (
	RepositoryMock struct {
		mock.Mock
	}
)

func (rm *RepositoryMock) GetData() diasFestivos.DiasFestivosResponse {
	args := rm.Called()
	return args.Get(0).(diasFestivos.DiasFestivosResponse)
}
