LogAgent 2022 by cuikunlei

   这个logagent 是我跟着网上教程学习，本来是tail+kafka+etcd  由于学习之用，所以将程序由kafka改为nsq ,供大家学习之用

//一下为向ETCD中添加配置。可以优化下

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
