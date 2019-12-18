package engine

import (
	"fmt"
	"github.com/cy422396350/crowller/fetch"
	"log"
)


type Widget interface {
	ID() string
}
type widget struct {
	id string
}

func (w widget) ID() string {
	return w.id
}

func NewWidget() widget  {
	return widget{
		id:"333",
	}
}



func Run(seeds ...Request)  {
	var requests  []Request
	 for _,seed := range seeds {
	 	requests = append(requests,seed)
	 }

	 for len(requests) > 0 {
		r:=requests[0]
		requests = requests[1:]
		log.Printf("request url is %s"+r.Url)
		bytes, e := fetch.Fetch(r.Url)
		if e!=nil {
		 log.Fatal(e,r.Url)
		}
		 result := r.Parser(bytes)
		 requests = append(requests,result.Requests...)
		 for _,v:= range result.Items {
		 	fmt.Printf("%v \n",v)
		 }
		 fmt.Println(len(result.Items))
	 }
}