package logger

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	AuthMethod = "/auth.AuthService/Validate"
)

func LoggingInterceptor(logger *logrus.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		traceID := uuid.New().String()
		// Log the request body
		if info.FullMethod != AuthMethod {
			logger.Infof("%s Method: %s, Request body: %v", traceID, info.FullMethod, req)
		}

		// Invoke the gRPC handler to process the request.
		resp, err := handler(context.WithValue(ctx, "traceID", traceID), req)

		if err != nil {
			logger.Warnf("%s Error: %s ", traceID, err.Error())
			if info.FullMethod == AuthMethod {
				logger.Infof("%s Method: %s, Request body: %v", traceID, info.FullMethod, req)
			}
		} else if info.FullMethod != AuthMethod {
			// Log the response body
			logger.Infof("%s  Method: %s, Response body: %v", traceID, info.FullMethod, resp)
		}

		return resp, err
	}
}
