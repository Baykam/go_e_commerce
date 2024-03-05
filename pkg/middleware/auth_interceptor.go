package middleware

import (
	"context"
	"encoding/json"
	jwt "golang_testing_grpc/pkg/jtoken"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	ignoreMethods []string
}

func NewAuthInterceptor(ignoreMethods []string) *AuthInterceptor {
	return &AuthInterceptor{
		ignoreMethods: ignoreMethods,
	}
}

func (ai *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp any, err error) {
		for _, m := range ai.ignoreMethods {
			if info.FullMethod == m {
				return handler(ctx, req)
			}
		}

		ctx, userId, err := ai.authorize(ctx)
		if err != nil {
			return nil, status.New(codes.Internal, err.Error()).Err()
		}

		// attach "userId" to context
		ctx = context.WithValue(ctx, "userId", userId)
		return handler(ctx, req)
	}
}

func (ai *AuthInterceptor) authorize(ctx context.Context) (context.Context, string, error) {
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(m["token"]) == 0 {
		return ctx, "", status.New(codes.Unauthenticated, "missing token").Err()
	}

	payload, err := jwt.ValidateToken(m["token"][0])
	if err != nil {
		return ctx, "", status.New(codes.Unauthenticated, "unauthorized").Err()
	}

	var meta map[string]interface{}
	b, err := json.Marshal(payload)
	if err != nil {
		return ctx, "", status.New(codes.Unauthenticated, "unauthorized").Err()
	} else {
		if err := json.Unmarshal(b, &meta); err != nil {
			log.Println("Error while unmarshalling auth data", err)
		}
	}

	return ctx, payload["id"].(string), nil

}
