package itemServer

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
)

func CreateItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			save(item)
		}
	}()
	return out
}

func save(item interface{}) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://47.104.141.245:9200"),
		//false in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	response, err := client.Index().Index("data_profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
