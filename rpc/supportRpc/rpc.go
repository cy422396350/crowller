package supportRpc

import (
	"context"
	"github.com/cy422396350/crowller/engine"
	"gopkg.in/olivere/elastic.v5"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type SaveItem struct {
	Client *elastic.Client
	Index  string
}

func (s *SaveItem) Save(item engine.Item, result *string) (err error) {
	_, err = s.Client.Index().Index(s.Index).Type(item.Type).Id(item.Id).BodyJson(item).Do(context.Background())
	if err != nil {
		return err
	}
	*result = "ok"
	return
}

// 启动一个jsonRpc的服务器
func Serve(host string, service interface{}) (err error) {
	rpc.Register(service)
	listen, err := net.Listen("tcp", host)
	if err != nil {
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}

		go jsonrpc.ServeConn(conn)

	}
}

func GetClient(host string) (client *rpc.Client, err error) {
	dial, err := net.Dial("tcp", host)
	client = jsonrpc.NewClient(dial)
	return
}
