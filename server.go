package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mservice-bench/lib/broker"
	"github.com/go-mservice-bench/lib/db"
	"github.com/go-mservice-bench/lib/handlers"
)

func server(d *db.DB, q *broker.Queue) {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.GET("/account", ginInjector(d, q, handlers.GetAllAccount))
	r.GET("/account/:id", ginInjector(d, q, handlers.GetAccount))
	r.POST("/account", ginInjector(d, q, handlers.CreateAccount))
	r.PATCH("/account/:id", ginInjector(d, q, handlers.UpdateAccountById))
	r.DELETE("/account/:id", ginInjector(d, q, handlers.DeleteAccount))
	r.POST("/transaction", ginInjector(d, q, handlers.CreateTransaction))
	r.Run()
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
