package parser

import (
	"github.com/cy422396350/crowller/engine"
	"regexp"
)

var (
	compile     = regexp.MustCompile(`<div class="photo"><a href="(http://album.zhenai.com/u/[\w]+)" target="_blank"><img src="[^"]+" alt="([^"]+)"></a></div>`)
	otherPerson = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)
)

func FindPeople(contents []byte) engine.Result {

	people := compile.FindAllStringSubmatch(string(contents), -1)
	otherPeople := otherPerson.FindAllStringSubmatch(string(contents), -1)
	result := engine.Result{}
	for _, man := range people {
		name := man[2]
		url := man[1]
		result.Requests = append(result.Requests, engine.Request{Url: url, Parser: func(bytes []byte) engine.Result {
			return ParserProfile(bytes, name, url)
		}})
	}
	for _, man := range otherPeople {
		result.Requests = append(result.Requests, engine.Request{Url: man[1], Parser: func(bytes []byte) engine.Result {
			return FindPeople(bytes)
		}})
	}

	return result
}
