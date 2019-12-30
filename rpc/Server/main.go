package main

import (
	rpcDemo "github.com/cy422396350/crowller/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcDemo.ServiceName{})
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Println("error is ", err)
			continue
		}
		go jsonrpc.ServeConn(accept)
	}

}
