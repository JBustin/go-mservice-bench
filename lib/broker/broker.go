package broker

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Queue struct {
	name   string
	client *redis.Client
}

func NewQueue(name string, client *redis.Client) Queue {
	return Queue{name, client}
}

func (q Queue) Flush() error {
	_, err := q.client.Del(ctx, q.name).Result()
	return err
}

func (q Queue) Push(data string) error {
	_, err := q.client.RPush(ctx, q.name, data).Result()
	return err
}

func (q Queue) Pop() (string, error) {
	return q.client.LPop(ctx, q.name).Result()
}

func (q Queue) IsEmpty() (bool, error) {
	length, err := q.client.LLen(ctx, q.name).Result()
	return length == 0, err
}

type Worker struct {
	q       Queue
	qFail   Queue
	action  func(data string) error
	play    bool
	delayMs int
	client  *redis.Client
}

func NewWorker(q Queue, delayMs int, client *redis.Client, action func(data string) error) Worker {
	return Worker{
		q:       q,
		qFail:   NewQueue(fmt.Sprintf("%v%v", q.name, ":fail"), client),
		delayMs: delayMs,
		action:  action,
		client:  client,
		play:    false,
	}
}

func (w *Worker) Start() {
	stop := make(chan bool)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		for {
			select {
			case <-sig: // if cancel() execute
				stop <- true
				return
			default:
				data, err := w.q.Pop()
				if err != nil && err != redis.Nil {
					fmt.Println(err)
				}
				if data != "" {
					err = w.action(data)
					if err != nil {
						fmt.Println(err)
					}
				}
			}

			time.Sleep(time.Duration(w.delayMs) * time.Millisecond)
		}
	}()
	<-stop
}
