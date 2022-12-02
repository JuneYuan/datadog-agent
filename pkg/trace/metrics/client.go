// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package metrics

import (
	"time"

	"github.com/DataDog/datadog-go/v5/statsd"
)

// StatsClient represents a client capable of sending stats to some stat endpoint.
type StatsClient interface {
	Gauge(name string, value float64, tags []string, rate float64) error
	Count(name string, value int64, tags []string, rate float64) error
	Histogram(name string, value float64, tags []string, rate float64) error
	Timing(name string, value time.Duration, tags []string, rate float64) error
	Flush() error
}

// Client is a global Statsd client. When a client is configured via Configure,
// that becomes the new global Statsd client in the package.
var Client StatsClient = (*statsd.Client)(nil)

// Gauge calls Gauge on the global Client, if set.
func Gauge(name string, value float64, tags []string, rate float64) error {
	return nil
}

// Count calls Count on the global Client, if set.
func Count(name string, value int64, tags []string, rate float64) error {
	return nil
}

// Histogram calls Histogram on the global Client, if set.
func Histogram(name string, value float64, tags []string, rate float64) error {
	return nil
}

// Timing calls Timing on the global Client, if set.
func Timing(name string, value time.Duration, tags []string, rate float64) error {
	return nil
}

// Flush flushes any pending metrics to the agent.
func Flush() error {
	return nil
}
