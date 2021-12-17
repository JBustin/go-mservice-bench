package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mservice-bench/lib/handlers"
	"github.com/go-mservice-bench/lib/injectors"
)

func router(d *injectors.DI) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.GET("/account", injectors.GinHandler(d, handlers.GetAllAccount))
	r.GET("/account/:id", injectors.GinHandler(d, handlers.GetAccount))
	r.POST("/account", injectors.GinHandler(d, handlers.CreateAccount))
	r.PATCH("/account/:id", injectors.GinHandler(d, handlers.UpdateAccountById))
	r.DELETE("/account/:id", injectors.GinHandler(d, handlers.DeleteAccount))
	r.POST("/transaction", injectors.GinHandler(d, handlers.CreateTransaction))
	return r
}
