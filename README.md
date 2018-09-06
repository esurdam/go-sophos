# go-sophos

A [Sophos UTM REST API client](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en). 


Sophos types are automatically generated using `bin/gen.go` which queries a UTM's Gateway  `/api/definitions`  path

No foreign deps.

### Prerequisites

The Sophos UTM REST API must be enabled in Administrator settings.

### Usage

```
go get github.com/esurdam/go-sophos
```

Using the client

```
// Query the gateway
client, _ := sophos.New("192.168.0.1:4848", sophos.WithApiToken("abCDEFghIjklMNOPQkwSnwbutCpHdjQz"))
res, _ := client.Get("/api/nodes")

// Marshal to Nodes struct
var nodes sophos.Nodes
_ = res.MarshalTo(&nodes)
```

Requesting the current port of the WebAdmin:

```
client, _ := New("192.168.0.1:4848", sophos.WithBasicAuth("user", "pass))
res, _ := client.Get("/api/nodes/webadmin.port")

var port int
res.MarshalTo(&port)
fmt.Println(port)
// Output: 4444
```

Requesting a REST Object:

```
var dns sophos.Dns
res, _ := client.Get(dns.GetPath())
err := res.MarshalTo(&dns)
```


Deleting a packet filter rule with reference REF_PacPacXYZ.

Note – The example uses the X-RESTD-ERR-ACK: all to automatically approve the deletion of the object.
```
client, _ := sophos.New("192.168.0.1:4848", 
    sophos.WithApiToken("abCDEFghIjklMNOPQkwSnwbutCpHdjQz"),
    sophos.WithSessionClose,
    sophos.AutoResolveErrsMode,
)

res, err := client.Delete("api/objects/packetfilter/packetfilter/REF_PacPacXYZ")
```

## Generating Types

`bin/gen.go` will query the UTM `api/definitions` path to generate the`generated.go` file which contains structs corresponding to UTM API definitions.


```
export ENDPOINT=192.168.0.1:4444
export TOKEN=abcde1234

make
```

## Testing

Testing requres a valid endpoint and token
```
export ENDPOINT=192.168.0.1:4444
export TOKEN=abcde1234

make test
```

## Todo

- Expand generated types to include Resources and Node edge types (e.g. Licensing and Packetfilter(ediit) types)

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
