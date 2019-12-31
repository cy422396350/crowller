package main

import (
	"fmt"
	"github.com/cy422396350/crowller/config"
	"github.com/cy422396350/crowller/rpc/supportRpc"
	"gopkg.in/olivere/elastic.v5"
)

func main() {
	newClient, err2 := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://47.104.141.245:9200"))
	if err2 != nil {
		panic(err2)
	}
	// 启动一个服务
	supportRpc.Serve(fmt.Sprintf(":%d", config.Host), &supportRpc.SaveItem{newClient, config.Elastic_index})
}
