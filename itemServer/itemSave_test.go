package itemServer

import (
	"github.com/cy422396350/crowller/engine"
	"github.com/cy422396350/crowller/rpc/supportRpc"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {
	const host = ":8888"
	newClient, err2 := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://47.104.141.245:9200"))
	if err2 != nil {
		panic(err2)
	}
	item := engine.Item{
		Url:     "123ssss",
		Id:      "456",
		Type:    "789",
		Profile: nil,
	}
	turnOn := make(chan bool)
	// 启动一个服务
	go supportRpc.Serve(host, &supportRpc.SaveItem{newClient, item}, turnOn)
	<-turnOn
	// 用客户端连服务
	client, err := supportRpc.GetClient(host)
	if err != nil {
		t.Errorf("get client err is %v", err)
	}
	result := ""
	// call 存进去
	err = client.Call("SaveItem.Save", "test_data", &result)
	if err != nil || result != "ok" {
		t.Errorf("err is %v", err)
	}

}
