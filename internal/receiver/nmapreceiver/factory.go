package nmapreceiver

import (
	"context"
	"fmt"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"
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

func createLogsReceiver(_ context.Context, params receiver.Settings, baseCfg component.Config, consumer consumer.Logs) (receiver.Logs, error) {
	logger := params.Logger
	nmapCfg, ok := baseCfg.(*Config)
	if !ok {
		return nil, fmt.Errorf("expected config type %T, but got %T", nmapCfg, baseCfg)
	}

	nmapRcvr := &nmapReceiver{
		logger:       logger,
		nextConsumer: consumer,
		config:       nmapCfg,
		converter: adapter.NewConverter(component.TelemetrySettings{
			Logger: zap.NewNop(),
		}),
	}

	return nmapRcvr, nil
}
