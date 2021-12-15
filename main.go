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
	"github.com/go-mservice-bench/lib/injectors"
	"github.com/go-mservice-bench/lib/logger"
	"github.com/go-mservice-bench/lib/redis"
)

func main() {
	fmt.Println("\t- go mservice bench -")

	action := flag.String("x", "", "Action")
	flag.Parse()

	fs := fs.New()
	env, err := env.New(fs)
	handlerErr(err)

	c, err := config.New(env)
	handlerErr(err)
	fmt.Println(c)

	l := logger.NewLog(c.LogLevel)
	redisClient := redis.New(c)

	d, err := db.Init(&c)
	handlerErr(err)
	defer d.Client.Close()

	q := broker.NewQueue(c.RedisTransactionQueueName, redisClient)
	di := injectors.NewDi(d, &c, &l, &q)

	switch *action {
	case "server":
		router := server(&di)
		router.Run()
	case "worker":
		worker(&di)
	default:
		fmt.Printf("Unexpected action [server, worker] %v\n", *action)
	}
}

func handlerErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
