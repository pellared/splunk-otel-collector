package nmapreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type nmapReceiver struct {
	host         component.Host
	nextConsumer consumer.Logs
	scanner      *Scanner

	config *Config
	cancel context.CancelFunc
	logger *zap.Logger
}

func (nr *nmapReceiver) Start(ctx context.Context, host component.Host) error {
	nr.host = host
	if ctx == nil {
		ctx = context.Background()
	}
	_, nr.cancel = context.WithCancel(ctx)

	//TODO: run scanner immediately.
	//TODO: base intervals on the last scan time.
	interval, _ := time.ParseDuration(nr.config.Interval)
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				nr.logger.Info("Processing nmap logs now!")
				scan, err := nr.scanner.Run()
				if err != nil {
					nr.logger.Error("Error running nmap scan", zap.Error(err))
					continue
				}
				pLogs := plog.NewLogs()
				lr := pLogs.ResourceLogs().AppendEmpty().ScopeLogs().AppendEmpty().LogRecords().AppendEmpty()
				scan.ToLogRecord(lr)
				nr.nextConsumer.ConsumeLogs(ctx, pLogs)
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (nr *nmapReceiver) Shutdown(_ context.Context) error {
	if nr.cancel != nil {
		nr.cancel()
	}
	return nil
}
