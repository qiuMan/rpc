package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc-qiu/service"
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

	rpc.RegisterName("KVStoreService", service.NewKVStoreService())

	listenner, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}
	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}
		rpc.ServeConn(conn)
	}

}
