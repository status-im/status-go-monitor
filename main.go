package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
)

type rcpResp map[string]interface{}

const serverAddr = "127.0.0.1"
const serverPort = 8545

func main() {
	addr := fmt.Sprintf("http://%s:%d", serverAddr, serverPort)
	c, err := rpc.DialHTTP(addr)
	if err != nil {
		log.Fatal("Failed to connec to RPC:", err)
	}
	fmt.Println("Connected: ", addr)
	var rval interface{}
	err = c.Call(&rval, "admin_peers")
	if err != nil {
		log.Fatal("Failed to make RPC call:", err)
	}
	pretty, err := json.MarshalIndent(rval, "", "  ")
	fmt.Println("Response: ", string(pretty))
}
