package diasFestivos

type (
	DiasFestivosService interface {
		GetByTypeAndDateRange(request DiasFestivosRequest) DiasFestivos
	}
	DiasFestivosRepository interface {
		GetData() DiasFestivosResponse
	}
	DiasFestivosResponse struct {
		Status string         `json:"status"`
		Data   []DiasFestivos `json:"data"`
	}
	DiasFestivos struct {
		Date        string `json:"date"`
		Title       string `json:"title"`
		Type        string `json:"type"`
		Inalienable bool   `json:"inalienable"`
		Extra       string `json:"extra"`
	}

	DiasFestivosRequest struct {
		Type      string `json:"type"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}
)
