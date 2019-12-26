package parser

import (
	"github.com/cy422396350/crowller/engine"
	"github.com/cy422396350/crowller/zhenai/model"
	"log"
	"regexp"
)

var profile = regexp.MustCompile(`<div [^>]+> ([^|]+) \| ([^|]+) \| ([^|]+) \| ([^|]+) \| ([^|]+) \| ([^|]+)元`)

var idReg = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParserProfile(contents []byte, name string, url string) engine.Result {
	if contents == nil {
		log.Println("xxx")
	}
	id := idReg.FindStringSubmatch(url)
	res := profile.FindAllStringSubmatch(string(contents), -1)
	if len(res) < 1 {
		log.Println(res, "res len <6,url is"+url)
		return engine.Result{}
	}

	results := engine.Result{
		Requests: nil,
		Items: []engine.Item{
			{
				Url:  url,
				Type: "zhenai",
				Id:   id[1],
				Profile: model.Profile{
					Name:      name,
					Age:       res[0][2],
					Education: res[0][3],
					Marige:    res[0][4],
					Height:    res[0][5],
					Income:    res[0][6],
				},
			},
		},
	}

	return results
}
