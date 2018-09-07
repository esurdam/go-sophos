// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// ClientlessVpn is a generated struct representing the Sophos ClientlessVpn Endpoint
// GET /api/nodes/clientless_vpn
type ClientlessVpn struct {
	ClientlessVpnConnection ClientlessVpnConnection `json:"clientless_vpn_connection"`
	ClientlessVpnGroup      ClientlessVpnGroup      `json:"clientless_vpn_group"`
}

var defsClientlessVpn = map[string]sophos.RestObject{
	"ClientlessVpnConnection": &ClientlessVpnConnection{},
	"ClientlessVpnGroup":      &ClientlessVpnGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of ClientlessVpn's Objects
func (ClientlessVpn) RestObjects() map[string]sophos.RestObject {
	return defsClientlessVpn
}

// GetPath implements sophos.RestGetter
func (*ClientlessVpn) GetPath() string { return "/api/nodes/clientless_vpn" }

// RefRequired implements sophos.RestGetter
func (*ClientlessVpn) RefRequired() (string, bool) { return "", false }

var defClientlessVpn = &sophos.Definition{Description: "clientless_vpn", Name: "clientless_vpn", Link: "/api/definitions/clientless_vpn", Swag: map[string]sophos.MethodMap{"/objects/clientless_vpn/connection/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"clientless_vpn/connection"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/clientless_vpn/group/": {"get": sophos.MethodDescriptions{Description: "Returns all available clientless_vpn/group objects", Parameters: []sophos.Parameter(nil), Tags: []string{"clientless_vpn/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new clientless_vpn/group object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "clientless_vpn/group that will be created", Type: "", Required: true}}, Tags: []string{"clientless_vpn/group"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 201: {Description: "OK"}}}}, "/objects/clientless_vpn/group/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"clientless_vpn/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 204: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"clientless_vpn/group"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "clientless_vpn/group that will be changes", Type: "", Required: true}}, Tags: []string{"clientless_vpn/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "clientless_vpn/group that will be updated", Type: "", Required: true}}, Tags: []string{"clientless_vpn/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}, "/objects/clientless_vpn/group/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"clientless_vpn/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}}}}, "/objects/clientless_vpn/connection/": {"get": sophos.MethodDescriptions{Description: "Returns all available clientless_vpn/connection objects", Parameters: []sophos.Parameter(nil), Tags: []string{"clientless_vpn/connection"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new clientless_vpn/connection object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "clientless_vpn/connection that will be created", Type: "", Required: true}}, Tags: []string{"clientless_vpn/connection"}, Responses: map[int]struct{ Description string }{201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/clientless_vpn/connection/{ref}": {"patch": sophos.MethodDescriptions{Description: "Changes to parts of the object connection types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "clientless_vpn/connection that will be changes", Type: "", Required: true}}, Tags: []string{"clientless_vpn/connection"}, Responses: map[int]struct{ Description string }{404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object connection", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "clientless_vpn/connection that will be updated", Type: "", Required: true}}, Tags: []string{"clientless_vpn/connection"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object connection", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"clientless_vpn/connection"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available connection types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"clientless_vpn/connection"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}}}

// Definition returns the /api/definitions struct of ClientlessVpn
func (ClientlessVpn) Definition() sophos.Definition { return *defClientlessVpn }

// ApiRoutes returns all known ClientlessVpn Paths
func (ClientlessVpn) ApiRoutes() []string {
	return []string{
		"/api/objects/clientless_vpn/connection/",
		"/api/objects/clientless_vpn/connection/{ref}",
		"/api/objects/clientless_vpn/connection/{ref}/usedby",
		"/api/objects/clientless_vpn/group/",
		"/api/objects/clientless_vpn/group/{ref}",
		"/api/objects/clientless_vpn/group/{ref}/usedby",
	}
}

// References returns the ClientlessVpn's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (ClientlessVpn) References() []string {
	return []string{
		"REF_ClientlessVpnConnection",
		"REF_ClientlessVpnGroup",
	}
}

// ClientlessVpnConnection is an Sophos Endpoint subType and implements sophos.RestObject
type ClientlessVpnConnections []ClientlessVpnConnection
type ClientlessVpnConnection struct {
	Locked        string        `json:"_locked"`
	Reference     string        `json:"_ref"`
	_type         string        `json:"_type"`
	AllowedUsers  []string      `json:"allowed_users"`
	AutoLogin     bool          `json:"auto_login"`
	Comment       string        `json:"comment"`
	Destination   string        `json:"destination"`
	HostKeyCert   string        `json:"host_key_cert"`
	Login         string        `json:"login"`
	Name          string        `json:"name"`
	Password      string        `json:"password"`
	PfExceptions  []interface{} `json:"pf_exceptions"`
	Port          int64         `json:"port"`
	PrivateKey    string        `json:"private_key"`
	RdpSecurity   string        `json:"rdp_security"`
	RecordSession bool          `json:"record_session"`
	Service       string        `json:"service"`
	ShareSession  bool          `json:"share_session"`
	Status        bool          `json:"status"`
	UID           int64         `json:"uid"`
	WebPath       string        `json:"web_path"`
}

// GetPath implements sophos.RestObject and returns the ClientlessVpnConnections GET path
// Returns all available clientless_vpn/connection objects
func (*ClientlessVpnConnections) GetPath() string { return "/api/objects/clientless_vpn/connection/" }

// RefRequired implements sophos.RestObject
func (*ClientlessVpnConnections) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ClientlessVpnConnections GET path
// Returns all available connection types
func (c *ClientlessVpnConnection) GetPath() string {
	return fmt.Sprintf("/api/objects/clientless_vpn/connection/%s", c.Reference)
}

// RefRequired implements sophos.RestObject
func (c *ClientlessVpnConnection) RefRequired() (string, bool) { return c.Reference, true }

// DeletePath implements sophos.RestObject and returns the ClientlessVpnConnection DELETE path
// Creates or updates the complete object connection
func (*ClientlessVpnConnection) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/connection/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ClientlessVpnConnection PATCH path
// Changes to parts of the object connection types
func (*ClientlessVpnConnection) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/connection/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ClientlessVpnConnection POST path
// Create a new clientless_vpn/connection object
func (*ClientlessVpnConnection) PostPath() string {
	return "/api/objects/clientless_vpn/connection/"
}

// PutPath implements sophos.RestObject and returns the ClientlessVpnConnection PUT path
// Creates or updates the complete object connection
func (*ClientlessVpnConnection) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/connection/%s", ref)
}

// Type implements sophos.Object
func (c *ClientlessVpnConnection) GetType() string { return c._type }

// ClientlessVpnGroup is an Sophos Endpoint subType and implements sophos.RestObject
type ClientlessVpnGroup []interface{}

// GetPath implements sophos.RestObject and returns the ClientlessVpnGroup GET path
// Returns all available clientless_vpn/group objects
func (*ClientlessVpnGroup) GetPath() string { return "/api/objects/clientless_vpn/group/" }

// RefRequired implements sophos.RestObject
func (*ClientlessVpnGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ClientlessVpnGroup DELETE path
// Creates or updates the complete object group
func (*ClientlessVpnGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ClientlessVpnGroup PATCH path
// Changes to parts of the object group types
func (*ClientlessVpnGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ClientlessVpnGroup POST path
// Create a new clientless_vpn/group object
func (*ClientlessVpnGroup) PostPath() string {
	return "/api/objects/clientless_vpn/group/"
}

// PutPath implements sophos.RestObject and returns the ClientlessVpnGroup PUT path
// Creates or updates the complete object group
func (*ClientlessVpnGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/clientless_vpn/group/%s", ref)
}