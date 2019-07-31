package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	plexSessionsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "plex_media_server",
			Subsystem: "sessions",
			Name:      "active_total",
			Help:      "Total of actives sessions on remote plex media server",
		},
	)
)

type CollectorPlex interface {
	CurrentSessionsCount() (int, error)
}

type PlexExporter struct {
	PlexServer CollectorPlex
}

func (pe *PlexExporter) Describe(ch chan<- *prometheus.Desc) {
}

func (pe *PlexExporter) Collect(ch chan<- prometheus.Metric) {
	sessions, _ := pe.PlexServer.CurrentSessionsCount()

	plexSessionsGauge.Set(float64(sessions))
	ch <- plexSessionsGauge
}