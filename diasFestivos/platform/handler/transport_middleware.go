package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func TransportMiddleware(log *zap.Logger) Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {

		return func(ctx context.Context, request interface{}) (interface{}, error) {

			defer log.Info("Request received ")
			return e(ctx, request)
		}
	}
}
