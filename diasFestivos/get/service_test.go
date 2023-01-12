package get

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"pruebatecnica.com/diasFestivos"
	"pruebatecnica.com/diasFestivos/platform/storage/repository"
	"testing"
)

func Test_NewService(t *testing.T) {
	logger, _ := zap.NewProduction()
	repositoryMock := repository.RepositoryMock{}

	type (
		test struct {
			name     string
			expected *DiasFestivosService
		}
	)

	testCases := []test{
		{
			name: "test new service",
			expected: &DiasFestivosService{
				repository: &repositoryMock,
				logger:     logger,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := NewDiasFestivosService(&repositoryMock, logger)
			assert.Equal(t, testCase.expected, got)
		})
	}
}

func TestDiasFestivosService_GetByTypeAndDateRange(t *testing.T) {
	logger, _ := zap.NewProduction()
	repositoryMock := repository.RepositoryMock{}
	service := NewDiasFestivosService(&repositoryMock, logger)

	type (
		test struct {
			name     string
			expected []diasFestivos.DiasFestivos
		}
	)

	testCases := []test{
		{
			name: "test GetByTypeAndRange example",
			expected: []diasFestivos.DiasFestivos{{
				Date:        "2023-02-02",
				Title:       "test",
				Type:        "test",
				Inalienable: false,
				Extra:       "test",
			}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repositoryMock.ExpectedCalls = nil
			repositoryMock.On("GetData").Return(createDummyDiasFestivosResponse())
			got := service.GetByTypeAndDateRange(diasFestivos.DiasFestivosRequest{
				Type:      "test",
				StartDate: "2023-01-01",
				EndDate:   "2023-12-01",
			})
			assert.Equal(t, testCase.expected, got)
		})
	}
}

func createDummyDiasFestivosResponse() diasFestivos.DiasFestivosResponse {
	return diasFestivos.DiasFestivosResponse{
		Status: "success",
		Data: []diasFestivos.DiasFestivos{{
			Date:        "2023-02-02",
			Title:       "test",
			Type:        "test",
			Inalienable: false,
			Extra:       "test",
		},
		},
	}
}
