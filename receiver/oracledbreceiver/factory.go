// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package oracledbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
)

const (
	typeStr   = "oracledb"
	stability = component.StabilityLevelInDevelopment
)

// NewFactory creates a new Oracle receiver factory.
func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		createDefaultConfig,
		component.WithMetricsReceiver(createMetricsReceiver, stability))
}

func createMetricsReceiver(ctx context.Context, settings component.ReceiverCreateSettings, receiver config.Receiver, metrics consumer.Metrics) (component.MetricsReceiver, error) {
	return &oracledbreceiver{}, nil
}

func createDefaultConfig() config.Receiver {
	return &Config{ReceiverSettings: config.NewReceiverSettings(config.NewComponentID(typeStr))}
}
