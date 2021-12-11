/*
 * Copyright (c) 2021 The Clymene Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package es

import (
	"github.com/Clymene-project/Clymene/pkg/es"
	"github.com/Clymene-project/Clymene/plugin/storage/es/metricstore/dbmodel"
	"github.com/Clymene-project/Clymene/prompb"
	"go.uber.org/zap"
)

const (
	clymeneIndex = "clymene-metrics"
	metricType   = "metric"
)

type MetricWriter struct {
	client      es.Client
	logger      *zap.Logger
	metricIndex string
	converter   dbmodel.Converter
}

// MetricWriterParams holds constructor parameters for NewMetricWriter
type MetricWriterParams struct {
	Client      es.Client
	Logger      *zap.Logger
	IndexPrefix string
	Archive     bool
}

// NewMetricWriter creates a new MetricWriter for use
func NewMetricWriter(p MetricWriterParams) *MetricWriter {
	prefix := ""
	if p.IndexPrefix != "" {
		prefix = p.IndexPrefix + "-"
	}
	return &MetricWriter{
		client:      p.Client,
		logger:      p.Logger,
		metricIndex: prefix + clymeneIndex,
		converter:   dbmodel.Converter{},
	}
}

func (s *MetricWriter) WriteMetric(metrics []prompb.TimeSeries) error {
	for _, metric := range metrics {
		jsonTimeSeries := s.converter.ConvertTsToJSON(metric)
		s.writeMetric(&jsonTimeSeries)
	}
	return nil
}

// bulk insert
func (s *MetricWriter) writeMetric(metric *map[string]interface{}) {
	s.client.Index().Index(s.metricIndex).Type(metricType).BodyJson(&metric).Add()
}