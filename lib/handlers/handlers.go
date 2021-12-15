package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-mservice-bench/lib/account"
	"github.com/go-mservice-bench/lib/injectors"
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

func GetAllAccount(d *injectors.DI, c *gin.Context) {
	var a []account.Account
	if err := d.Db.Account.FindAllLimit(&a, d.Config.DbLimit); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": a})
}

func GetAccount(d *injectors.DI, c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	var a account.Account
	if err := d.Db.Account.FindById(&a, id); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": a})
}

func CreateAccount(d *injectors.DI, c *gin.Context) {
	var input account.CreateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	a, err := d.Db.Account.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": a})
}

func DeleteAccount(d *injectors.DI, c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	if err := d.Db.Account.DeleteById(id); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func UpdateAccountById(d *injectors.DI, c *gin.Context) {
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
	a, err := d.Db.Account.UpdateById(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": a})
}

func CreateTransaction(d *injectors.DI, c *gin.Context) {
	var t transaction.Transaction
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := d.Queue.Push(t.String()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
