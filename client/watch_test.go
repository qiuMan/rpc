package client

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"
	"time"
)

func TestWacth(t *testing.T) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing error: ", err)
	}

	doWatch(client)
}

func doWatch(client *rpc.Client) {
	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("wacth: ", keyChanged)
	}()

	err := client.Call("KVStoreService.Set", [2]string{"abc", "abc-value"}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)

	err = client.Call("KVStoreService.Set", [2]string{"abc", "abc-value1"}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)
}
