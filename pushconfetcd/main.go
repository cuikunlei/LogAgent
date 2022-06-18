package main

import (
	"context"
	"fmt"
	"time"

	client "go.etcd.io/etcd/client/v3"
)

// etcd client put/get demo
// use etcd/clientv3
func main() {
	cli, err := client.New(client.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")

	defer cli.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//value := `[{"path":"/usr/local/var/log/nginx/access.log","topic":"web_log"}]`
	value := `[{"path":"/usr/local/var/log/nginx/access.log","topic":"web_log"},{"path":"/usr/local/var/log/nginx/access.log","topic":"redis_log"}]`
	_, err = cli.Put(ctx, "/logagent/collect_config", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}
