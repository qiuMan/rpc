package service

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

//Hello ...
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listenner, err := net.Listen("tcp", ":1235")
	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}

	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
