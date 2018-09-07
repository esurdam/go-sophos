package main

import (
	"fmt"
	"log"
	"os"

	"github.com/esurdam/go-sophos"
	"github.com/esurdam/go-sophos/types"
)

var (
	endpoint, token string
)

func init() {
	if len(os.Args) == 2 {
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

	// Get all the DHCP Servers
	var dss types.DhcpServers
	err = client.GetObject(&dss)
	if err != nil {
		log.Fatal(err)
	}

	// Use any ref as a sample
	var sampleRef string
	for _, ds := range dss {
		fmt.Println(ds.Name)
		fmt.Println(ds.Reference)
		sampleRef = ds.Reference
	}

	// Use the ref to fetch a single Server, as Reference is required
	ds := types.DhcpServer{Reference: sampleRef}
	err = client.GetObject(&ds)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ds.Name)

	// This will fail since single instances of objects require a Ref to query
	ds = types.DhcpServer{}
	err = client.GetObject(&ds)
	if err == sophos.ErrRefRequired {
		fmt.Printf("This will fail since ref is required: %s", err.Error())
	} else {
		panic("should not have gotten here")
	}
}
