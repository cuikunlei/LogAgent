package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	client "go.etcd.io/etcd/client/v3"
)

var (
	cli *client.Client
)

//需要收集的日志的配置信息
type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

//初始化ETCD的函数
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = client.New(client.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeout,
	})
	if err != nil {
		//handel error!!
		fmt.Printf("Connect To ETCD Failed,ERR:%v\n", err)
		return
	}
	return
}

// etcd watch 从ETCD中根据key获取配置项
func GetConf(key string) (LogEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("GET From ETCD Failed,ERR:%v\n", err)
		return
	}

	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &LogEntryConf)
		if err != nil {
			fmt.Printf("Unmarshal ETCD Value Failed,ERR:%v\n")
			return
		}
	}
	return
}

//etcd watch 不停的从etcd中获取数据
func WatchConf(key string, newConfCh chan<- []*LogEntry) {
	ch := cli.Watch(context.Background(), key)
	// 从通道尝试取值(监视的信息)
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
			// 通知taillog.tskMgr
			// 1. 先判断操作的类型
			var newConf []*LogEntry
			if evt.Type != client.EventTypeDelete {
				//如果是删除操作，手动传递一个空的配置项
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("Unmarshal Failed, ERR:%v\n", err)
					continue
				}
			}
			fmt.Printf(" Get New Conf:%v\n", newConf)
			newConfCh <- newConf
		}
	}
}
