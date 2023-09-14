package main

import "github.com/prometheus/client_golang/prometheus"
import "github.com/prometheus/client_golang/prometheus/promauto"

var bot_processed_requests_total = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "bot_interaction_requests_total",
		Help: "The total number of interaction requests received from Discord",
	},
	[]string{"server"},
)
