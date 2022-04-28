package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster" //support automatic consumer-group rebalancing and offset tracking
	"github.com/golang/glog"
)

func main() {
	groupID := "group-1"
	topicList := "topic_1"
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest //初始从最新的offset开始

	c, err := cluster.NewConsumer(strings.Split("localhost:9092", ","), groupID, strings.Split(topicList, ","), config)
	if err != nil {
		glog.Errorf("Failed open consumer: %v", err)
		return
	}
	defer c.Close()
	go func() {
		for err := range c.Errors() {
			glog.Errorf("Error: %s\n", err.Error())
		}
	}()

	go func() {
		for note := range c.Notifications() {
			glog.Infof("Rebalanced: %+v\n", note)
		}
	}()

	for msg := range c.Messages() {
		fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Value)
		c.MarkOffset(msg, "") //MarkOffset 并不是实时写入kafka，有可能在程序crash时丢掉未提交的offset
	}
}

//type backoffFunc func(retries, maxRetries int) time.Duration

func testSend() {

	config := sarama.NewConfig()
	config.Producer.Flush.Messages = 1
	config.Producer.Return.Successes = true
	config.Producer.Retry.BackoffFunc(1, 3)
	servers := []string{""}
	sarama.NewAsyncProducer(servers, config)

}
