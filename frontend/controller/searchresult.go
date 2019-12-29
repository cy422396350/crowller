package controller

import (
	"context"
	"fmt"
	"github.com/cy422396350/crowller/engine"
	"github.com/cy422396350/crowller/frontend/view/model"
	"gopkg.in/olivere/elastic.v5"
	"html/template"
	"net/http"
	"reflect"
)

type SearchController struct {
	view   *template.Template
	client *elastic.Client
}

func (s SearchController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var result model.SearchResult
	resp, _ := s.client.
		Search("dating").Type("zhenai").
		//对q 进行正则替换，只在给elasticSearch时rewrite
		Query(elastic.NewQueryStringQuery(r.FormValue("q"))).
		From(10).
		Do(context.Background())

	result.Hits = resp.TotalHits()
	result.Start = 10
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	fmt.Println(result)
	err := s.view.Execute(w, result)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, r.FormValue("q"))
}

func CreateController(filename string) SearchController {
	client, err := elastic.NewClient(
		elastic.SetURL("http://47.104.141.245:9200"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchController{
		view:   template.Must(template.ParseFiles(filename)),
		client: client,
	}
}

//type SearchResultHandler struct {
//	//从client拿数据，送给view,去render
//	view   view.SearchResultView
//	client *elastic.Client
//}
//
//func CreateSearchResultHandler(template string) SearchResultHandler {
//
//	client, err := elastic.NewClient(elastic.SetSniff(false))
//	if err != nil {
//		panic(err)
//	}
//	return SearchResultHandler{
//		view:   view.CreateSearchResultView(template),
//		client: client,
//	}
//}
//
////实现一个interface ,from用于分页
////localhost:9200/search?q=男 已购房&from=20
//func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	//去掉空格
//	q := strings.TrimSpace(req.FormValue("q"))
//	from, err := strconv.Atoi(req.FormValue("from"))
//	if err != nil {
//		from = 0
//	}
//	//fmt.Fprintf(w,"q=%s, from=%d",q,from)  //响应结果，写入w
//	page, err := h.getSearchResult(q, from)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//	}
//	err = h.view.Render(w, page)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//	}
//
//}
//
//func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
//	var result model.SearchResult
//	result.Query = q
//	resp, err := h.client.
//		Search("dating_profile").
//		//对q 进行正则替换，只在给elasticSearch时rewrite
//		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
//		From(from).
//		Do(context.Background())
//	if err != nil {
//		return result, err
//	}
//	//己有的包装
//	result.Hits = resp.TotalHits()
//	result.Start = from
//	//用到了反射
//	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
//
//	//支持分页
//	result.PrevFrom = result.Start - len(result.Items)
//	result.NextFrom = result.Start + len(result.Items)
//
//	return result, nil
//}
//
////前端搜索没有输入Payload.Age<30，而是输入的Age<30,因此为了保证能在elastic中正常使用，重新添加上Payload.
//func rewriteQueryString(q string) string {
//	re := regexp.MustCompile(`([A-Z][a-z]*):`) //([A-Z][a-z]*)表示Height: Age:
//	//$1 代表（）中的部分
//	return re.ReplaceAllString(q, "Payload.$1:")
//}
