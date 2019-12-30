package rpcDemo

import "errors"

type ServiceName struct{}

type Args struct {
	A, B int
}

func (s ServiceName) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("div zero")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}
