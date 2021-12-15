package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-mservice-bench/lib/broker"
	"github.com/go-mservice-bench/lib/db"
	"github.com/go-mservice-bench/lib/injectors"
	"github.com/stretchr/testify/assert"
)

func Test_Ping(t *testing.T) {
	d := injectors.DI{
		Db:    &db.DB{},
		Queue: &broker.Queue{},
	}
	router := server(&d)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}
