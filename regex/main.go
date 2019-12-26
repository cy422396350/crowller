package main

import (
	"fmt"
	"regexp"
)

const text = `my email is 411396350@qq.com,my email is 411396350@qq.com,my email is 411396350@qq.com`
const strings = `
<div data-v-5b109fc3="" class="des f-cl"> 上海 | 31岁 | 大学本科 | 未婚 | 160cm | 12001-20000元
<a data-v-5b109fc3="" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1851277113" target="_self" class="online f-fr">查看最后登录时间</a></div>`
const url = `https://album.zhenai.com/u/1135668693`

var idReg = regexp.MustCompile(`https://album.zhenai.com/u/([\d]+)`)
var id2Reg = regexp.MustCompile(`<div [^>]+ class="m-content-box m-des"><span data-v-[\w]+>([^<]+)</span></div>`)
var profile = regexp.MustCompile(`<div data-v-5b109fc3="" class="des f-cl"> ([^|]+) \| ([^|]+) \| ([^|]+) \| ([^|]+) \| ([^|]+) \|([^|]+)`)

//正则 . 任何字符+一个或多个  *零个或多个`
func main() {
	findString := profile.FindAllStringSubmatch(strings, -1)
	find := id2Reg.FindStringSubmatch(`<div data-v-8b1eac0c="" class="m-content-box m-des"><span data-v-8b1eac0c="">做人做事不管怎样都要身前身后收拾的利落精致 ，才让人安然。实际年龄87年32周岁！
大学毕业苏州园区研发生物医药公司工作四年，之后创业两家高端皮肤管理中心用时两年，现在准备换个行业进入外企工作学习。</span><!----></div>`)
	fmt.Println(find)
	fmt.Println(findString[0][2])
	id := idReg.FindStringSubmatch(url)
	fmt.Println(id)
}
