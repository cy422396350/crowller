package itemServer

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func CreateItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		for {
			item := <-out
			//log.Print(item,"\n")
			_, err := save(item)
			if err != nil {
				log.Printf("item save err %v item is %v \n", err, item)
			}
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://47.104.141.245:9200"),
		//false in docker
		elastic.SetSniff(false),
	)
	if err != nil {
		return "", err
	}
	response, err := client.Index().Index("profile").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return response.Id, nil
}
