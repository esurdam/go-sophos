package main

import (
	"fmt"
	"log"
	"os"

	"github.com/esurdam/go-sophos"
	"github.com/esurdam/go-sophos/api/v1.3.0/nodes"
)

//noinspection ALL
var (
	endpoint, token string
)

func init() {
	if len(os.Args) == 3 {
		endpoint = os.Args[1]
		token = os.Args[2]
	}

	if endpoint == "" {
		endpoint = os.Getenv("ENDPOINT")
	}
	if token == "" {
		token = os.Getenv("TOKEN")
	}

	if endpoint == "" || token == "" {
		panic("need endpoint and token as args or from env ($ENDPOINT, $TOKEN)")
	}
}

func main() {
	client, err := sophos.New(endpoint, sophos.WithAPIToken(token))
	if err != nil {
		log.Fatal(err)
	}

	// using package func
	port, err := nodes.GetWebadminPort(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(port)

	// using node type
	var wp nodes.WebadminPort
	err = wp.Get(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(wp.Value)
}
