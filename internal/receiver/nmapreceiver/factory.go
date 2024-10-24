package nmapreceiver

import (
	"context"
	"fmt"

	// "github.com/Ullaakut/nmap/v3"
	// "github.com/mattn/go-shellwords"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var (
	typeStr = component.MustNewType("nmap")
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithLogs(createLogsReceiver, component.StabilityLevelAlpha),
	)
}

func createLogsReceiver(ctx context.Context, params receiver.Settings, baseCfg component.Config, consumer consumer.Logs) (receiver.Logs, error) {
	logger := params.Logger
	nmapCfg, ok := baseCfg.(*Config)
	if !ok {
		return nil, fmt.Errorf("expected config type %T, but got %T", nmapCfg, baseCfg)
	}

	// args, err := shellwords.Parse(nmapCfg.ScannerArguments)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to parse scanner arguments: %w", err)
	// }

	// scanner, err := nmap.NewScanner(
	// 	ctx,
	// 	nmap.WithCustomArguments(args...),
	// )

	// if err != nil {
	// 	return nil, fmt.Errorf("unable to create nmap scanner: %w", err)
	// }

	nmapRcvr := &nmapReceiver{
		logger:       logger,
		nextConsumer: consumer,
		config:       nmapCfg,
		scanner:      NewScanner(nil, nmapCfg.Fake),
	}

	return nmapRcvr, nil
}
