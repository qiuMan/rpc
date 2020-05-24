// Title: rpc 客户端
// Description: 测试
// Author: longerQiu
// Date: 2020-03-28 20:12:58
// LastEditTime: 2020-03-28 20:12:59
//

package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1235")

	if err != nil {
		log.Fatal("dialing error: ", err)
	}

	client, err := rpc.NewClientWithCodec(jsonrpc.NewClient(conn))

	if err != nil {
		log.Fatal("NewClientWithCodec error: ", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
