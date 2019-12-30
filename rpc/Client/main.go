package main

import (
	"fmt"
	rpcDemo "github.com/cy422396350/crowller/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("ServiceName.Div", rpcDemo.Args{10, 3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	err = client.Call("ServiceName.Div", rpcDemo.Args{10, 0}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
