# go-sophos

[![Documentation](https://godoc.org/github.com/esurdam/go-sophos?status.svg)](http://godoc.org/github.com/esurdam/go-sophos)
[![Go Report Card](https://goreportcard.com/badge/github.com/esurdam/go-sophos)](https://goreportcard.com/report/github.com/esurdam/go-sophos)
[![codecov](https://codecov.io/gh/esurdam/go-sophos/branch/master/graph/badge.svg)](https://codecov.io/gh/esurdam/go-sophos)
[![Build Status](https://travis-ci.com/esurdam/go-sophos.svg?branch=master)](https://travis-ci.com/esurdam/go-sophos)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/esurdam/go-sophos/blob/master/LICENSE)

A [Sophos UTM REST API client](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) for Go. 

### Prerequisites

The Sophos UTM REST API must be enabled in Administrator settings.

Familiarity with the [Sophos docs](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en).

### Usage

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

Requesting the current port of the WebAdmin:

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

Requesting a REST type:

```go
import "github.com/esurdam/go-sophos/types"

var nodes types.Nodes
_ := client.GetObject(&nodes)

// active Ips
nodes.LicensingActiveIps 
```

Deleting a packet filter rule with reference `REF_PacPacXYZ`.

This example uses the `X-Restd-Err-Ack: all` header to automatically approve the deletion of the object:
```go
_, err := client.Delete(
    "api/objects/packetfilter/packetfilter/REF_PacPacXYZ", 
    sophos.WithSessionClose,
    sophos.AutoResolveErrsMode,
)

// OR as a RestObject

pf := types.PacketfilterPacketfilter{Reference: "REF_PacPacXYZ"}
err := client.DeleteObject(&pf, sophos.WithSessionClose, sophos.AutoResolveErrsMode)
```

Requesting a Definition:

```go
import "github.com/esurdam/go-sophos/types"

var nodes types.Nodes
d := nodes.Definition()

_ := d.GetSwag(client)
fmt.Printf("%v", d.Swag)
```


## Generating Types

Sophos types are automatically generated using [bin/gen.go](bin/gen.go) which queries the UTM `api/definitions` path to generate the [generated.go](types/generated.go) file which contains structs corresponding to UTM API definitions.


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

- [x] Add PUT, POST, PATCH and DELETE methods to generated objects
- [x] Create a wrapper Client for REST objects `client.Get(&nodes)` 

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details
