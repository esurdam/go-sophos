# go-sophos

[![Documentation](https://godoc.org/github.com/esurdam/go-sophos?status.svg)](http://godoc.org/github.com/esurdam/go-sophos)
[![Go Report Card](https://goreportcard.com/badge/github.com/esurdam/go-sophos)](https://goreportcard.com/report/github.com/esurdam/go-sophos)
[![codecov](https://codecov.io/gh/esurdam/go-sophos/branch/master/graph/badge.svg)](https://codecov.io/gh/esurdam/go-sophos)
[![Build Status](https://travis-ci.com/esurdam/go-sophos.svg?branch=master)](https://travis-ci.com/esurdam/go-sophos)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/esurdam/go-sophos/blob/master/LICENSE)

A [Sophos UTM REST API client](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) for Go with zero dependencies. 

## Prerequisites

The Sophos UTM REST API must be enabled in Administrator settings.

Familiarity with the [Sophos docs](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en).

API types and functions are [generated](#generating-types) and versioned against UTM's declared Restd version.

## Usage

```bash
go get github.com/esurdam/go-sophos
```

Create a client:

```go
import "github.com/esurdam/go-sophos"

// All Options passed on initialize will be applied to all subsequent calls
client, _ := sophos.New(
    "192.168.0.1:4848", 
    sophos.WithBasicAuth("user", "pass"),
)
```

Requesting the current port of the WebAdmin (see [Nodes](#nodes) for more usage):
```go
import "github.com/esurdam/go-sophos"

client, _ := sophos.New(
    "192.168.0.1:4848", 
    sophos.WithApiToken("abCDEFghIjklMNOPQkwSnwbutCpHdjQz"),
)
res, _ := client.Get("/api/nodes/webadmin.port")

var port int
res.MarshalTo(&port)
fmt.Println(port)
// Output: 4848
```

### Nodes

Nodes are interacted with using pacakage level functions:

```go
import "github.com/esurdam/go-sophos/api/v1.3.0/nodes"

v, err := nodes.GetWebadminPort(client)
fmt.Println(v)
// Output: 4848

err = nodes.UpdateWebadminPort(client, 4444)
```

Or as struct types with syntactic sugar around the functions, as represented by the [Node](nodes.go#L24) interface:

```go
import "github.com/esurdam/go-sophos/api/v1.3.0/nodes"

var wap nodes.WebadminPort
err := wap.Get(client)
fmt.Println(wap.Value)
// Output: 4848

wap.Value = 4444
err = wap.Update(client)

```

You can get the whole UTM node tree as an [object](#objects) as well:

```go
import "github.com/esurdam/go-sophos/api/v1.3.0/objects"

var nodes objects.Nodes
_ := client.GetObject(&nodes)

// active Ips
nodes.LicensingActiveIps 
```

### Objects

Each file in the [objects](api/v1.3.0/objects) dir represents an [Endpoint](nodes.go#L7) generated from a [Definition](definition.go) and contains its generated Objects.

Objects implement the [RestObject](nodes.go) interface:

```go
import "github.com/esurdam/go-sophos/api/v1.3.0/objects"

var dns objects.Dns
err := client.GetObject(&dns)
```

Notice that some objects are pluralized and only implement the [RestGetter](nodes.go) interface:
```go
import "github.com/esurdam/go-sophos/api/v1.3.0/objects"

var ss objects.DnsRoutes
_ = client.GetObject(&ss)

// Each individual DnsRoute is therefore a RestObject
for _, s := range ss {
    ub, _ := client.GetUsedBy(&s)
    fmt.Printf("DNS ROUTE: %s [%s]\n  Used By Nodes: %v\n  Used by Objects: %v\n",s.Name, s.Reference, ub.Nodes, ub.Objects)
    // OUTPUT: DNS ROUTE: sophos.boom.local [REF_DnsRouBoomloca]
    //             Used By Nodes: [dns.routes]
    //             Used by Objects: []
}
```

Note that [Endpoint](nodes.go#L5) types contain their [Definition](definition.go#L3):

```go
import "github.com/esurdam/go-sophos/api/v1.3.0/objects"

fmt.Printf("%#v", objects.Dns{}.Definition())
// Output: sophos.Definition{
//  Description:"dns", 
//  Name:"dns", 
//  Link:"/api/definitions/dns"
// }
```

Requesting an Endpoint's [Swag](definition.go#L29):

```go
import "github.com/esurdam/go-sophos/api/v1.3.0/objects"

// with sugar
var dns objects.Dns
swag, _ := client.GetEndpointSwag(dns)

// without sweets
d := objects.Dns{}.Definition()
swag, _ := d.GetSwag(client)
```


## Examples

Examples from [Sophos docs](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en).

Deleting a packet filter rule with reference `REF_PacPacXYZ`:

This example uses the `X-Restd-Err-Ack: all` header to automatically approve the deletion of the object:
```go
import "github.com/esurdam/go-sophos"

client, _ := sophos.New(
    "192.168.0.1:4848", 
    sophos.WithBasicAuth("user", "pass"),
)

_, err := client.Delete(
    "api/objects/packetfilter/packetfilter/REF_PacPacXYZ", 
    sophos.WithSessionClose, 
    sophos.AutoResolveErrsMode,
)
```

The same as above but using objects: [[example](examples/delete_packetfilter.go)]

```go
import "github.com/esurdam/go-sophos"
import "github.com/esurdam/go-sophos/api/v1.3.0/objects"

client, _ := sophos.New(
    "192.168.0.1:4848", 
    sophos.WithBasicAuth("user", "pass"),
)

// object knows api route
pf := objects.PacketfilterPacketfilter{
	Reference: "REF_PacPacXYZ"
}

err := client.DeleteObject(&pf, 
	sophos.WithSessionClose, 
	sophos.AutoResolveErrsMode
)
```

Creating a PacketFilter: [[example](examples/create_packetfilter.go)]

```go
import "github.com/esurdam/go-sophos"
import "github.com/esurdam/go-sophos/api/v1.3.0/objects"

client, _ := sophos.New(
    "192.168.0.1:4848", 
    sophos.WithBasicAuth("user", "pass"),
)

pf := objects.PacketfilterPacketfilter{
    Action:       "accept",
    Destinations: []string{sophos.RefNetworkAny},
    Direction:    "in",
    Log:          true,
    Services:     []string{sophos.RefServiceAny},
    Sources:      []string{sophos.RefNetworkAny},
    Status:       true,
}

err := client.PostObject(&pf, 
	sophos.WithRestdInsert("packetfilter.rules", 0), 
	sophos.WithSessionClose,
)

// successful creation will have unmarshalleed the Response
pf.Reference  
```

Errors

```go
if err != nil {
    // for modifying requests (PATCH, PUT, POST, DELETE), err returned may be of type *sophos.Errors
    // see client.Do and Response type for how errors are parsed
    err.(*sophos.Errors).Error() == err.Error()
    sophos.IsFatalErr(err) == err.(*sophos.Errors).IsFatal()
    
    // view each individual error
    for _, e := range *err.(*sophos.Errors) {
    	e.Error() 
    	e.IsFatal()
    }
}
```

## Generating Types

Sophos types are automatically generated using [bin/gen.go](bin/gen.go) which queries the UTM `api/definitions` path to generate all the files in the [api](api/v1.3.0) which contain structs and helper functions corresponding to UTM API definitions.

Generated pacakages are versioned, feel free to generate against an older version and submit.

```bash
export ENDPOINT=192.168.0.1:4848
export TOKEN=abcde1234

make
```

## Testing

```bash
make test
```

## Todo
- [x] Respond with Errors to ObjectClient functions for caller inspection
- [x] Finish adding all example from [Sophos docs](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en)
- [x] Add [nodes](nodes.go) examples in README
- [x] Add PUT, POST, PATCH and DELETE methods to generated objects
- [x] Create a wrapper Client for REST objects `client.Get(&nodes)` 

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details
