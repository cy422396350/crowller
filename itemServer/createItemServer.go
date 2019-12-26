package itemServer

import (
	"context"
	"github.com/cy422396350/crowller/engine"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func CreateItemServer(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://47.104.141.245:9200"),
		//false in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		for {
			item := <-out
			//log.Print(item,"\n")
			_, err := save(item, client, index)
			if err != nil {
				log.Printf("item save err %v item is %v \n", err, item)
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
