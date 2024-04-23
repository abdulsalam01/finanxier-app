package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// authenticate is an interceptor that checks for the presence of an authorization token
func JWTAuthMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "retrieving metadata failed")
	}

	token, ok := md["authorization"]
	if !ok || len(token) < 2 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	// Continue handling the request.
	return handler(ctx, req)
}
