package grpc

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func LoggerMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Handling the original request.
	resp, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}

	logrus.Infof("Caller: [%s] with request: [%v]", info.FullMethod, req)
	return resp, nil
}
