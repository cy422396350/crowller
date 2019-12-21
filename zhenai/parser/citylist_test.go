package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {

	bytes, e := ioutil.ReadFile("citylist_test_data.html")

	if e != nil {
		panic(e)
	}

	result := ParserCityList(bytes)
	resultUrls := []string{"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng"}
	resultCity := []string{"阿坝", "阿克苏", "阿拉善盟"}
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("size is wrong , size is %d", len(result.Requests))
	}
	for i, url := range resultUrls {
		if result.Requests[i].Url != url {
			t.Errorf("url is wrong , url is %s", url)
		}
	}
	for i, city := range resultCity {
		if result.Items[i] != city {
			t.Errorf("city is wrong , city is %s", city)
		}
	}
}
