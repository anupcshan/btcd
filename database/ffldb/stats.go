package ffldb

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ioWrite = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "btcd",
		Name:      "leveldb_iowrite",
	})

	ioRead = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "btcd",
		Name:      "leveldb_ioread",
	})

	openTables = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "btcd",
		Name:      "leveldb_opentables",
	})

	levelTableSizes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "btcd",
		Name:      "leveldb_leveltablesizes",
	}, []string{"level"})

	levelTableCounts = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "btcd",
		Name:      "leveldb_leveltablecounts",
	}, []string{"level"})

	levelRead = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "btcd",
		Name:      "leveldb_levelread",
	}, []string{"level"})

	levelWrite = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "btcd",
		Name:      "leveldb_levelwrite",
	}, []string{"level"})

	levelCompactionDuration = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "btcd",
		Name:      "leveldb_levelcompactionduration",
	}, []string{"level"})

	compactions = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "btcd",
		Name:      "leveldb_compactions",
	}, []string{"type"})

	seekCompaction = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "btcd",
		Name:      "leveldb_seekcompactionenabled",
	})
)
