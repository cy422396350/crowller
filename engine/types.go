package engine

type Request struct {
	Url string
	Parser func([]byte) Result
}

type Result struct {
	Requests []Request
	Items []interface{}
}

func NilParserfunc([]byte) Result {
	return Result{}
}