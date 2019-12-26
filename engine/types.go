package engine

type Request struct {
	Url    string
	Parser func([]byte) Result
}

type Result struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Profile interface{}
}

func NilParserfunc([]byte) Result {
	return Result{}
}
