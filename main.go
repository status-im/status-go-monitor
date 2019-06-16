package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type rcpResp map[string]interface{}

const host = "127.0.0.1"
const port = 8545

func main() {
	url := fmt.Sprintf("http://%s:%d", host, port)
	fmt.Println("Type :%t", url)
	c, err := newClient(url)
	if err != nil {
		log.Fatal("Failed to connec to RPC:", err)
	}
	fmt.Println("Connected: ", url)

	peers, err := c.getPeers()
	if err != nil {
		log.Fatal("Failed to make RPC call:", err)
	}
	pretty, err := json.MarshalIndent(peers, "", "  ")
	fmt.Println("Response: ", string(pretty))
}
