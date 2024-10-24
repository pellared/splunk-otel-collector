package main

import (
	"context"
	"net"
	"sync/atomic"

	clpb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	lpb "go.opentelemetry.io/proto/otlp/logs/v1"
	"google.golang.org/grpc"
)

type (
	collector struct {
		Endpoint string

		logsService *logsService
		grpcSrv     *grpc.Server
		errCh       chan error
	}

	logsService struct {
		clpb.UnimplementedLogsServiceServer

		logs *atomic.Pointer[[]*lpb.ResourceLogs]
	}
)

func (coll *collector) Start() error {
	if coll.Endpoint == "" {
		coll.Endpoint = "localhost:0"
	}
	ln, err := net.Listen("tcp", coll.Endpoint)
	if err != nil {
		return err
	}
	coll.Endpoint = ln.Addr().String() // set actual endpoint

	coll.logsService = &logsService{logs: &atomic.Pointer[[]*lpb.ResourceLogs]{}}

	coll.grpcSrv = grpc.NewServer()
	clpb.RegisterLogsServiceServer(coll.grpcSrv, coll.logsService)
	coll.errCh = make(chan error, 1)
	go func() { coll.errCh <- coll.grpcSrv.Serve(ln) }()

	return nil
}

func (coll *collector) Close() error {
	coll.grpcSrv.GracefulStop()
	if err := <-coll.errCh; err != nil && err != grpc.ErrServerStopped {
		return err
	}
	return nil
}

func (coll *collector) ExportedLogs() []*lpb.ResourceLogs {
	return *coll.logsService.logs.Load()
}

func (s *logsService) Export(_ context.Context, req *clpb.ExportLogsServiceRequest) (*clpb.ExportLogsServiceResponse, error) {
	s.logs.Store(&req.ResourceLogs)
	return &clpb.ExportLogsServiceResponse{}, nil
}
