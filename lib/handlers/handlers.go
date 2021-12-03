package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-mservice-bench/lib/account"
	"github.com/go-mservice-bench/lib/broker"
	"github.com/go-mservice-bench/lib/db"
	"github.com/go-mservice-bench/lib/transaction"
)

func getId(c *gin.Context) (uint, error) {
	u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return 0, fmt.Errorf("%v", "Id format invalid")
	}
	return uint(u64), nil
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func GetAllAccount(db *db.DB, q *broker.Queue, c *gin.Context) {
	var a []account.Account
	if err := db.Account.FindAllLimit(&a, db.Config.DbLimit); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": a})
}

func GetAccount(db *db.DB, q *broker.Queue, c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	var a account.Account
	if err := db.Account.FindById(&a, id); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": a})
}

func CreateAccount(db *db.DB, q *broker.Queue, c *gin.Context) {
	var input account.CreateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	a, err := db.Account.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": a})
}

func DeleteAccount(db *db.DB, q *broker.Queue, c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	if err := db.Account.DeleteById(id); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdateAccountById(db *db.DB, q *broker.Queue, c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	var input account.UpdateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	a, err := db.Account.UpdateById(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": a})
}

func CreateTransaction(db *db.DB, q *broker.Queue, c *gin.Context) {
	var t transaction.Transaction
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := q.Push(t.String()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
