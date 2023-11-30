package main

import (
	context "context"

	newscratch "gitlab.ozon.ru/newscratch/internal/app/ozon/newscratch"
	_ "gitlab.ozon.ru/newscratch/internal/config"
	scratch "gitlab.ozon.ru/platform/scratch"
	logger "gitlab.ozon.ru/platform/tracer-go/logger"
)

func main() {
	a, err := scratch.New()
	if err != nil {
		logger.Fatalf(context.Background(), "can't create app: %v", err)
	}

	if err := a.Run(newscratch.NewDemo()); err != nil {
		logger.Fatalf(context.Background(), "can't run app: %v", err)
	}
}
