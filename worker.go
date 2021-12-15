package main

import (
	"github.com/go-mservice-bench/lib/broker"
	"github.com/go-mservice-bench/lib/handlers"
	"github.com/go-mservice-bench/lib/injectors"
)

func worker(d *injectors.DI) {
	w := broker.NewWorker(
		*d.Queue,
		d.Config.RedisWorkerDelayMs,
		*d.Logger,
		injectors.JobHandler(d, handlers.TransactionJob),
	)
	w.Start()
}
