package engine

import (
	"fmt"
	"github.com/cy422396350/crowller/fetch"
	"log"
)

type SimpleEngine struct {
}

//简单的实现:通过队列实现添加request实现递归来连续的爬
func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		result, e := worker(r)
		if e != nil {
			log.Printf("%v,%v", e, r.Url)
		}
		requests = append(requests, result.Requests...)
		for _, v := range result.Items {
			fmt.Printf("%v \n", v)
		}
	}
}

func worker(r Request) (Result, error) {
	log.Println(r.Url)
	bytes, e := fetch.Fetch(r.Url)
	if e != nil {
		log.Printf("出错了%v%v", e, r.Url)
		return Result{}, e
	}
	return r.Parser(bytes), nil
}
