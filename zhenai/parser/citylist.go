package parser

import (
	"github.com/cy422396350/crowller/engine"
	"regexp"
)

func ParserCityList(contents []byte) engine.Result {
	compile := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[\w]+)"[^>]*>([^<]+)</a>`)

	citys := compile.FindAllStringSubmatch(string(contents), -1)

	result := engine.Result{}

	for _, city := range citys {
		result.Requests = append(result.Requests, engine.Request{Url: city[1], Parser: FindPeople})
		result.Items = append(result.Items, engine.Item{})
	}

	return result
}
