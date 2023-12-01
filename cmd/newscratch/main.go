package main

import (
	context "context"

	newscratch "gitlab.ozon.ru/newscratch/internal/app/ozon/newscratch"
	_ "gitlab.ozon.ru/newscratch/internal/config"
	scratch "gitlab.ozon.ru/platform/scratch"
	logger "gitlab.ozon.ru/platform/tracer-go/logger"
	googleGrpc "google.golang.org/grpc"
)

func debug(ctx context.Context, req interface{}, info *googleGrpc.UnaryServerInfo, handler googleGrpc.UnaryHandler) (resp interface{}, err error) {
	logger.Warnf(ctx, "request: %#v", req)
	resp, err = handler(ctx, req)
	logger.Warnf(ctx, "response: %#v", resp)
	return
}

func main() {
	a, err := scratch.New(
		scratch.WithUnaryInterceptor(debug),
		scratch.WithServiceDescriptionInterceptor(debug))

	if err != nil {
		logger.Fatalf(context.Background(), "can't create app: %v", err)
	}

	if err := a.Run(newscratch.NewDemo()); err != nil {
		logger.Fatalf(context.Background(), "can't run app: %v", err)
	}
}
