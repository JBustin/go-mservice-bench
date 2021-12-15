package broker

import (
	"testing"
	"time"

	"github.com/go-mservice-bench/lib/logger"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

var name string = "myqueue"

func TestQueue(t *testing.T) {
	client, mock := redismock.NewClientMock()
	mock.MatchExpectationsInOrder(true)

	q := NewQueue(name, client)

	mock.ExpectRPush(name, "foo").SetVal(1)
	err := q.Push("foo")

	assert.Equal(t, nil, err, "should not raise an error")

	mock.ExpectLPop(name).SetVal("foo")
	value, err := q.Pop()

	assert.Equal(t, nil, err, "should not raise an error")
	assert.Equal(t, "foo", value, "should pop a value")

	mock.ExpectDel(name).SetVal(0)
	err = q.Flush()

	assert.Equal(t, nil, err, "should not raise an error")

	mock.ExpectLLen(name).SetVal(10)
	flag, err := q.IsEmpty()

	assert.Equal(t, nil, err, "should not raise an error")
	assert.Equal(t, false, flag, "should check if queue is empty")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestWorker(t *testing.T) {
	delayMs := 3
	client, mock := redismock.NewClientMock()
	mock.MatchExpectationsInOrder(true)

	actionArgs := []string{}
	action := func(data string) error {
		actionArgs = append(actionArgs, data)
		return nil
	}

	w := NewWorker(NewQueue(name, client), delayMs, logger.Logger{}, action)

	mock.ExpectLPop(name).SetVal("foo1")
	mock.ExpectLPop(name).SetVal("foo2")
	mock.ExpectLPop(name).SetVal("foo3")
	mock.ExpectLPop(name).SetVal("foo4")

	w.Start()
	time.Sleep(time.Duration(delayMs*4) * time.Millisecond)
	assert.Equal(t, []string{"foo1", "foo2", "foo3", "foo4"}, actionArgs, "should pop a value")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
