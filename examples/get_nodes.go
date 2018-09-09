package main

import (
	"fmt"
	"log"
	"os"

	"github.com/esurdam/go-sophos"
	"github.com/esurdam/go-sophos/api/v1.3.0/nodes"
	types "github.com/esurdam/go-sophos/api/v1.3.0/objects"
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

	fmt.Println(`
	DHCP SERVERS:
	`)
	// Get all the DHCP Servers
	var dss types.DhcpServers
	err = client.GetObject(&dss)
	if err != nil {
		log.Fatal(err)
	}
	for _, ds := range dss {
		fmt.Println(ds.Name)
		fmt.Println(ds.Reference)
		fmt.Println(ds.Status)
	}

	fmt.Println(`
	DNS Definition:
	`)

	d := types.Dns{}.Definition()
	swag, err := d.GetSwag(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(swag.Paths)


	fmt.Println(`
	DNS Get:
	`)

	var dns types.Dns
	err = client.GetObject(&dns)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", dns)

	fmt.Println(`
	Node Getters:
	`)

	var ll nodes.LicensingLicense
	ll.Get(client)
	fmt.Println(ll.Value)

	v, err := nodes.GetWebadminPort(client)
	fmt.Println(v)

	err = nodes.UpdateWebadminPort(client, 4444)
	if err != nil {
		log.Fatal(err.Error())
	}

	var wa nodes.WebadminPort
	wa.Get(client)
	fmt.Println(wa.Value)

	fmt.Println(`
	AmazonVpcConnections:
	`)

	var cc types.AmazonVpcConnections
	err = client.GetObject(&cc)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range cc {
		fmt.Println(c.Name)
		r, _ := client.Get(c.UsedByPath(c.Reference))
		var usedBy map[string]interface{}
		r.MarshalTo(&usedBy)
		fmt.Println(usedBy)
	}

	// Use any ref as a sample
	var sampleRef string
	for _, ds := range dss {
		sampleRef = ds.Reference
		break
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
		fmt.Printf("This will fail since ref is required: %s\n", err.Error())
	} else {
		panic("should not have gotten here")
	}
}

func convertKeysToStrings(obj map[interface{}]interface{}) map[string]interface{} {
	res := make(map[string]interface{})

	for k, v := range obj {
		res[fmt.Sprintf("%v", k)] = v
	}

	return res
}
