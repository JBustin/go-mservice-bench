package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-mservice-bench/lib/broker"
	"github.com/go-mservice-bench/lib/config"
	"github.com/go-mservice-bench/lib/db"
	"github.com/go-mservice-bench/lib/env"
	"github.com/go-mservice-bench/lib/fs"
	"github.com/go-mservice-bench/lib/redis"
)

func main() {
	fmt.Println("\t- go mservice bench -")

	action := flag.String("x", "", "Action")
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

	switch *action {
	case "server":
		server(d, &q)
	case "worker":
		worker(d, &q)
	default:
		fmt.Printf("Unexpected action [server, worker] %v\n", *action)
	}
}

func handlerErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
