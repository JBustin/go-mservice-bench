package main

import (
	"github.com/go-mservice-bench/lib/broker"
	"github.com/go-mservice-bench/lib/db"
	"github.com/go-mservice-bench/lib/handlers"
)

func worker(d *db.DB, q *broker.Queue) {
	w := broker.NewWorker(
		*q,
		d.Config.RedisWorkerDelayMs,
		q.Client,
		jobInjector(d, q, handlers.TransactionJob),
	)
	w.Start()
}

func jobInjector(
	d *db.DB,
	q *broker.Queue,
	fn func(*db.DB, *broker.Queue, string) error,
) func(string) error {
	return func(data string) error {
		return fn(d, q, data)
	}
}
