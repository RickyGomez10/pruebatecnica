package handler

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"pruebatecnica.com/diasFestivos"
)

func NewHttpDiasFestivosHandler(path string, endpoints endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()

	r.Handle(path,
		httpTransport.NewServer(endpoints,
			DecodeGetRequest,
			EncodeGetResponse,
		)).Methods(http.MethodPost)
	return r
}

func DecodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	//contentType := r.Header.Get("Content-Type")
	var DiasFestivosRequest diasFestivos.DiasFestivosRequest
	decodingError := json.NewDecoder(r.Body).Decode(&DiasFestivosRequest)
	return DiasFestivosRequest, decodingError
}

func EncodeGetResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	respons, _ := response.([]diasFestivos.DiasFestivos)
	return json.NewEncoder(w).Encode(respons)
}
