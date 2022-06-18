package main

import (
	"fmt"
	"logagent/conf"
	"logagent/etcd"
	"logagent/nsq"
	"logagent/taillog"
	"sync"
	"time"

	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func main() {
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("Load ini failed,err:%v\n", err)
		return

	}

	err = nsq.InitNsq(string(cfg.NsqConf.Address))
	fmt.Printf("CKL-TEST2 ADD %s\n", cfg.NsqConf.Address)
	if err != nil {
		fmt.Printf("init nsq failed,err:%v\n", err)
		return
	}
	fmt.Println("init Nsq success.")

	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success")
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Printf("etcd.getconf failed,err:%v\n", err)
		return
	}
	fmt.Printf("Get conf from etcd success,%v\n", logEntryConf)

	for index, value := range logEntryConf {
		fmt.Printf("Main for index:%v value:%v\n", index, value)
	}
	taillog.Init(logEntryConf)
	newConfChan := taillog.NewContChan()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(cfg.EtcdConf.Key, newConfChan)
	wg.Wait()
}
