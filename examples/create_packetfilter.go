package main

import (
	"fmt"
	"log"
	"os"

	"github.com/esurdam/go-sophos"
	types "github.com/esurdam/go-sophos/api/v1.3.0/objects"
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

	p := types.PacketfilterPacketfilter{
		Action:       "accept",
		Destinations: []string{sophos.RefNetworkAny},
		Direction:    "in",
		Log:          true,
		Services:     []string{sophos.RefServiceAny},
		Sources:      []string{sophos.RefNetworkAny},
		Status:       true,
	}

	err = client.PostObject(&p, sophos.WithRestdInsert("packetfilter.rules", 0), sophos.WithSessionClose)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.Reference)
}
