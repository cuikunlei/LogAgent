package conf

type AppConf struct {
	NsqConf  `ini:"nsq"`
	EtcdConf `ini:"etcd"`
}
type NsqConf struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:"chan_max_size"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Key     string `ini:"collect_log_key"`
	Timeout int    `ini:"timeout"`
}
