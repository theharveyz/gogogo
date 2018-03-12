package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	nsq "github.com/nsqio/go-nsq"
)

func main() {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", cfg)
	if err != nil {
		panic(err)
	}
	ticker := time.NewTicker(1 * time.Second)
	timeout := time.After(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			err = producer.Publish("test", []byte("test msg: time:"+strconv.FormatInt(time.Now().Unix(), 10)))
			if err != nil {
				panic(err)
			}
			fmt.Println("publish successfully")
		case <-timeout:
			os.Exit(0)
			break
		}
	}
}
