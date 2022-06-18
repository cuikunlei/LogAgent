package nsq

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

//初始化nsq
func InitNsq(address string) (err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(address, config)
	if err != nil {
		fmt.Printf("Create Producer Failed ERR:%v\n", err)
		return
	}
	return nil
}

//想nsq中发送数据
func SeedMeg(topic, data string) (err error) {
	err = producer.Publish(topic, []byte(data))
	if err != nil {
		fmt.Printf("publish msg to nsq failed,err:%v\n", err)
		return
	}
	return nil
}
