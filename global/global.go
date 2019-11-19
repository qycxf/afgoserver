package global

import (
	"cxfProject/afGo/afGoface"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
存储全局参数
*/

type ConfigGlobal struct {
	TcpServer afGoface.IServerFace //server对象
	Host      string               //当前主机监听的IP
	TcpPort   int                  //当前监听的端口号
	Name      string               //当前服务器的名称

	Version        string //当前版本号
	MaxConn        int    //允许的最大连接数
	MaxPackageSize uint32 //数据包最大值
	WorkerPoolSize uint32 //当前业务工作worker池的 goroutine数量
	MaxWorkerSize  uint32 //允许用户最多开启的worker池的 goroutine数量
}

var Cfg *ConfigGlobal

func init() {

	//默认值
	Cfg = &ConfigGlobal{

		Name:           "afGoServer",
		Version:        "0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        100,
		MaxPackageSize: 4096,
		WorkerPoolSize: 10,
		MaxWorkerSize:  1024,
	}

	Cfg.Reload()
}

func (cfg *ConfigGlobal) Reload() {
	data, err := ioutil.ReadFile("config/afGo.json")

	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &cfg)

	fmt.Printf("read config success!\n")

	if err != nil {
		panic(err)
	}
}
