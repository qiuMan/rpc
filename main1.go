package main

import (
	"log"
	"net"
	"net/rpc"
)

//HelloServiceName 服务名称
const HelloServiceName = ""

// HelloServiceInterface 服务接口
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

//RegisterHelloService 注册服务
func RegisterHelloService(rcvr HelloServiceInterface) error {
	return rpc.RegisterName(HelloServicName, rcvr)
}
