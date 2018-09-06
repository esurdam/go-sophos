# go-sophos

A [Sophos UTM REST API client](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en). 

No foreign deps.

### Prerequisites

The Sophos UTM REST API must be enabled in Administrator settings.

Familiarity with the [Sophos docs](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en)

### Usage

```
go get github.com/esurdam/go-sophos
```

Using the client

```
// All options passed on initialize will be applied to all subsequent calls
client, _ := sophos.New(
    "192.168.0.1:4848", 
    sophos.WithApiToken("abCDEFghIjklMNOPQkwSnwbutCpHdjQz"),
)
res, _ := client.Get("/api/nodes") 

var nodes sophos.Nodes
_ = res.MarshalTo(&nodes) // Marshal to Nodes struct

nodes.Licensing_activeIps // active Ips
```

Requesting the current port of the WebAdmin:

```
client, _ := New("192.168.0.1:4848", sophos.WithBasicAuth("user", "pass))
res, _ := client.Get("/api/nodes/webadmin.port")

var port int
res.MarshalTo(&port)
fmt.Println(port)
// Output: 4848
```

Requesting a REST Object:

```
var dns sophos.Dns
res, _ := client.Get(dns.GetPath())
err := res.MarshalTo(&dns)
```


Deleting a packet filter rule with reference `REF_PacPacXYZ`.

This example uses the `X-Restd-Err-Ack: all` header to automatically approve the deletion of the object.
```
res, err := client.Delete(
    "api/objects/packetfilter/packetfilter/REF_PacPacXYZ", 
    sophos.WithSessionClose,
    sophos.AutoResolveErrsMode,
)
```

## Generating Types

Sophos types are automatically generated using [bin/gen.go](bin/gen.go) which queries the UTM `api/definitions` path to generate the [generated.go](generated.go) file which contains structs corresponding to UTM API definitions.


```
export ENDPOINT=192.168.0.1:4848
export TOKEN=abcde1234

make
```

## Testing

Testing requres a valid endpoint and token
```
export ENDPOINT=192.168.0.1:4848
export TOKEN=abcde1234

make test
```

## Todo

- [ ] Add PUT, POST, PATCH and DELETE methods to generated objects
- [ ] Create a wrapper Client for REST objects `client.Get(&nodes)` 

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details
