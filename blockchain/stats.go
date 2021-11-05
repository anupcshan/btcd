package blockchain

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	blockcount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "btcd",
		Name:      "blockcount",
	}, []string{"live"})

	totalTransactions = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "btcd",
		Name:      "total_transactions",
	}, []string{"live"})
)

func logSnapshotState(s *BestState) {
	blockcount.WithLabelValues("live").Set(float64(s.Height))
	totalTransactions.WithLabelValues("live").Set(float64(s.TotalTxns))
}
