package abilities

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/endpoint"
	"context"
)

func loggingMiddleware(logger log.Logger) func(endpoint.Endpoint, string) endpoint.Endpoint {
    return func(next endpoint.Endpoint, endpointName string) endpoint.Endpoint {
    	return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling " + endpointName)
			defer logger.Log("msg", "called " + endpointName)
			return next(ctx, request)
		}
	}
}
