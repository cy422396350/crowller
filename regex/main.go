package main

import (
	"fmt"
	"regexp"
)

const text = `my email is 411396350@qq.com,my email is 411396350@qq.com,my email is 411396350@qq.com`


//正则 . 任何字符+一个或多个  *零个或多个
func main() {
	compile := regexp.MustCompile(`([\w]+)@([\w]+)(\.[\w]+)`)
	findString := compile.FindAllSubmatch([]byte(text),-1)
	for _,match := range findString{
		fmt.Printf("111%s\n",match)
		for _,sli := range match{
			fmt.Printf("%s\n",sli)
		}
	}
}
