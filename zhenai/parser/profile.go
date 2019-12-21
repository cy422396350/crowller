package parser

import (
	"github.com/cy422396350/crowller/engine"
	"github.com/cy422396350/crowller/zhenai/model"
	"log"
	"regexp"
	"strconv"
)

var AgeStr = regexp.MustCompile(`<div class="m-btn [a-z0-9A-Z]+" data-v-[\w]+>[^<]+</div><div class="m-btn [a-z0-9A-Z]+" data-v-[\w]+>([0-9]+)岁</div><div class="m-btn [a-z0-9A-Z]+" data-v-[\w]+>([^<]+)</div><div class="m-btn [a-z0-9A-Z]+" data-v-[\w]+>([^<]+)</div><div class="m-btn [a-z0-9A-Z]+" data-v-[\w]+>([^<]+)</div><div class="m-btn [a-z0-9A-Z]+" data-v-[\w]+>([^<]+)</div><div class="m-btn [a-z0-9A-Z]+" data-v-[\w]+>([^<]+)</div><div class="m-btn [a-z0-9A-Z]+" data-v-[\w]+>([^<]+)</div><div class="m-btn [a-z0-9A-Z]+" data-v-[\w]+>([^<]+)</div>`)
var Mdes = regexp.MustCompile(`<div class="m-content-box m-des" data-v-[\w]+><span data-v-[\w]+>([^<]+)</span></div>`)

func ParserProfile(contents []byte, name string) engine.Result {
	if contents == nil {
		log.Println("xxx")
	}
	f := model.Profile{}
	f.Name = name
	res := AgeStr.FindStringSubmatch(string(contents))
	mResult := Mdes.FindStringSubmatch(string(contents))
	if len(res) > 0 {
		f.Age, _ = strconv.Atoi(res[1])
		f.Education = res[3]
		f.Marige = res[4]
		f.Income = res[5]
	}
	if len(mResult) > 1 {
		f.Mmdes = mResult[1]
	}

	results := engine.Result{
		Items: []interface{}{f},
	}

	return results
}
