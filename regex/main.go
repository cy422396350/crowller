package main

import (
	"fmt"
	"regexp"
)

const text = `my email is 411396350@qq.com,my email is 411396350@qq.com,my email is 411396350@qq.com`
const strings = `<div class="m-content-box m-des" data-v-8b1eac0c><span data-v-8b1eac0c>希望以后得到的都是温柔，????????</span></div>`

//正则 . 任何字符+一个或多个  *零个或多个`
func main() {
	compile := regexp.MustCompile(`<div class="m-content-box m-des" data-v-[\w]+><span data-v-[\w]+>([^<]+)</span></div>`)
	findString := compile.FindAllStringSubmatch(strings, -1)
	fmt.Println(findString)
}
