package config

import (
	"github.com/go-mservice-bench/lib/env"
)

const (
	defaultServerPort                = 8081
	defaultWorkerRoutines            = 0
	defaultDbProvider                = "sqlite"
	defaultDbName                    = "data.db"
	defaultDbLimit                   = 30
	defaultRedisPort                 = 6379
	defaultRedisHost                 = "localhost"
	defaultRedisPingMs               = 1000
	defaultRedisBatchJob             = 25
	defaultRedisTransactionQueueName = "transaction"
	defaultRedisWorkerDelayMs        = 10
)

type Config struct {
	ServerPort                int
	WorkerRoutines            int
	DbProvider                string
	DbName                    string
	DbLimit                   int
	RedisPort                 int
	RedisHost                 string
	RedisPingMs               int
	RedisBatchJob             int
	RedisTransactionQueueName string
	RedisWorkerDelayMs        int
}

func New(e env.Env) (Config, error) {
	serverPort, exists := e.GetInt("SERVER_PORT")
	if !exists {
		serverPort = defaultServerPort
	}
	workerRoutines, exists := e.GetInt("WORKER_ROUTINES")
	if !exists {
		workerRoutines = defaultWorkerRoutines
	}
	DbProvider, exists := e.Get("DB_PROVIDER")
	if !exists {
		DbProvider = defaultDbProvider
	}
	DbName, exists := e.Get("DB_NAME")
	if !exists {
		DbName = defaultDbName
	}
	redisPort, exists := e.GetInt("REDIS_PORT")
	if !exists {
		redisPort = defaultRedisPort
	}
	redisHost, exists := e.Get("REDIS_HOST")
	if !exists {
		redisHost = defaultRedisHost
	}
	redisPingMs, exists := e.GetInt("REDIS_PING_MS")
	if !exists {
		redisPingMs = defaultRedisPingMs
	}
	redisBatchJob, exists := e.GetInt("REDIS_BATCH_JOB")
	if !exists {
		redisBatchJob = defaultRedisBatchJob
	}
	RedisWorkerDelayMs, exists := e.GetInt("REDIS_WORKER_DELAY_MS")
	if !exists {
		RedisWorkerDelayMs = defaultRedisWorkerDelayMs
	}

	return Config{
		ServerPort:                serverPort,
		WorkerRoutines:            workerRoutines,
		DbProvider:                DbProvider,
		DbName:                    DbName,
		DbLimit:                   defaultDbLimit,
		RedisPort:                 redisPort,
		RedisHost:                 redisHost,
		RedisPingMs:               redisPingMs,
		RedisBatchJob:             redisBatchJob,
		RedisTransactionQueueName: defaultRedisTransactionQueueName,
		RedisWorkerDelayMs:        RedisWorkerDelayMs,
	}, nil
}
