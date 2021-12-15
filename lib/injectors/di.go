package injectors

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mservice-bench/lib/broker"
	"github.com/go-mservice-bench/lib/config"
	"github.com/go-mservice-bench/lib/db"
	"github.com/go-mservice-bench/lib/logger"
)

type DI struct {
	Db     *db.DB
	Config *config.Config
	Logger *logger.Logger
	Queue  *broker.Queue
}

func NewDi(
	d *db.DB,
	c *config.Config,
	l *logger.Logger,
	q *broker.Queue) DI {
	return DI{Db: d, Config: c, Logger: l, Queue: q}
}

func GinHandler(
	d *DI,
	fn func(*DI, *gin.Context),
) func(*gin.Context) {
	return func(c *gin.Context) {
		fn(d, c)
	}
}

func JobHandler(
	d *DI,
	fn func(*DI, string) error,
) func(string) error {
	return func(data string) error {
		return fn(d, data)
	}
}
