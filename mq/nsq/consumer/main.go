package main

import (
	"fmt"

	nsq "github.com/nsqio/go-nsq"
)

func main() {
	cfg := nsq.NewConfig()

	consumer, _ := nsq.NewConsumer("test", "ch3", cfg)
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error {
		fmt.Println("nsqd:", m.NSQDAddress)
		fmt.Println("timestamp:", m.Timestamp)
		fmt.Println("msgid:", m.ID)
		fmt.Println("body:", string(m.Body))
		return nil
	}), 10)
	consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	<-consumer.StopChan
}
