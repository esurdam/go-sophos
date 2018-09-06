package sophos

// Node is a resource within Sophos UTM that you can update but cannot create, delete, or move.
// Nodes are organized hierarchically as in a filesystem. An example of a node is the “Shell Access”
// service module within Sophos UTM (identified as the “ssh” node). To enable the Shell Access service,
// you can set the ssh.status leaf node.
//
// Unlike filesystems which use a slash “/” to separate the different nodes, Sophos UTM RESTful API uses
// a dot “.” to separate different nodes. Nodes reference objects, for example “Shell Access” and have a
// leaf node of allowed_networks which is an array of references to network objects.
type Node interface{}

// An Object resides in collections, which you can create, change, or delete. The collections of objects are
// predefined into classes and types. A class describes the general purpose of the objects whereas the type
// describes the required data for the object. Some objects need to be referenced by a node, to work properly
// (e.g. packetfilter/packetfilter and packetfilter/nat).
//
// For example, one of most used class is “network”. The network class describes objects like “host” for real hosts
// (e.g. 192.168.0.1) or “network” for subnets (e.g. 192.168.0.0/24). The collection of objects is always expressed
// with a class and a type, e.g. “network/host” or “network/network”.
type Object interface {
	// ApiRoutes returns a slice of strings containing the generated structs endpoints
	ApiRoutes() []string
}

// A Reference is the connections between nodes and objects as well as between one object and another object.
// Each confd node and object has a list of attributes with pre defined types, where one of the types can be a
// reference. Please note however that you can’t create a reference in all cases. You can only make a reference
// in scenarios where nodes and objects are designed to be connected. Technically, references are strings that
// always start with the prefix “REF_”.
type Reference string
