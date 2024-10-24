package main

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/sdk/log"
)

func newLoggerProvider(ctx context.Context) (*log.LoggerProvider, error) {
	exp, _ := otlploggrpc.New(ctx)
	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(exp)),
	)
	return loggerProvider, nil
}
