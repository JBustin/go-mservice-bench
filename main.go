package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-mservice-bench/lib/broker"
	"github.com/go-mservice-bench/lib/config"
	"github.com/go-mservice-bench/lib/db"
	"github.com/go-mservice-bench/lib/env"
	"github.com/go-mservice-bench/lib/fs"
	"github.com/go-mservice-bench/lib/handlers"
	"github.com/go-mservice-bench/lib/redis"
)

func main() {
	fmt.Println("\t- go mservice bench -")

	action := flag.String("x", "worker", "Action")
	flag.Parse()

	fs := fs.New()
	env, err := env.New(fs)
	handlerErr(err)

	config, err := config.New(env)
	handlerErr(err)
	fmt.Println(config)

	redisClient := redis.New(config)

	d, err := db.Init(&config)
	handlerErr(err)
	// defer d.Stop()

	q := broker.NewQueue(config.RedisTransactionQueueName, redisClient)

	// TODO
	// make two endpoints, one for the server, one for the workers

	switch *action {
	case "server":
		r := gin.Default()
		r.GET("/ping", handlers.Ping)
		r.GET("/account", ginInjector(d, &q, handlers.GetAllAccount))
		r.GET("/account/:id", ginInjector(d, &q, handlers.GetAccount))
		r.POST("/account", ginInjector(d, &q, handlers.CreateAccount))
		r.PATCH("/account/:id", ginInjector(d, &q, handlers.UpdateAccountById))
		r.DELETE("/account/:id", ginInjector(d, &q, handlers.DeleteAccount))
		r.POST("/transaction", ginInjector(d, &q, handlers.CreateTransaction))
		r.Run()
	case "worker":
		w := broker.NewWorker(
			q,
			config.RedisWorkerDelayMs,
			redisClient,
			jobInjector(d, &q, handlers.TransactionJob),
		)
		w.Start()
	default:
		fmt.Printf("Unexpected action [server, worker] %v\n", action)
	}
}

func handlerErr(err error) {
	if err != nil {
		log.Panic(err)
	}
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

func ginInjector(
	d *db.DB,
	q *broker.Queue,
	fn func(*db.DB, *broker.Queue, *gin.Context),
) func(*gin.Context) {
	return func(c *gin.Context) {
		fn(d, q, c)
	}
}
