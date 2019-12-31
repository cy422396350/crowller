package itemServer

import (
	"context"
	"github.com/cy422396350/crowller/config"
	"github.com/cy422396350/crowller/engine"
	"github.com/cy422396350/crowller/rpc/supportRpc"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func CreateItemServer(host string) (chan engine.Item, error) {
	client, err := supportRpc.GetClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		for {
			item := <-out
			//log.Print(item,"\n")
			result := ""
			err2 := client.Call(config.Service_name, item, &result)
			if err2 != nil || result != "ok" {
				log.Println(err2)
			}
		}
	}()
	return out, nil
}

func save(item engine.Item, client *elastic.Client, index string) (id string, err error) {
	response, err := client.Index().Index(index).Type(item.Type).Id(item.Id).BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return response.Id, nil
}
