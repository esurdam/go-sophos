// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Ospf is a generated struct representing the Sophos Ospf Endpoint
// GET /api/nodes/ospf
type Ospf struct {
	OspfMessageDigestKey OspfMessageDigestKey `json:"ospf_message_digest_key"`
	OspfArea             OspfArea             `json:"ospf_area"`
	OspfGroup            OspfGroup            `json:"ospf_group"`
	OspfInterface        OspfInterface        `json:"ospf_interface"`
}

var defsOspf = map[string]sophos.RestObject{
	"OspfMessageDigestKey": &OspfMessageDigestKey{},
	"OspfArea":             &OspfArea{},
	"OspfGroup":            &OspfGroup{},
	"OspfInterface":        &OspfInterface{},
}

// RestObjects implements the sophos.Node interface and returns a map of Ospf's Objects
func (Ospf) RestObjects() map[string]sophos.RestObject {
	return defsOspf
}

// GetPath implements sophos.RestGetter
func (*Ospf) GetPath() string { return "/api/nodes/ospf" }

// RefRequired implements sophos.RestGetter
func (*Ospf) RefRequired() (string, bool) { return "", false }

var defOspf = &sophos.Definition{Description: "ospf", Name: "ospf", Link: "/api/definitions/ospf", Swag: map[string]sophos.MethodMap{"/objects/ospf/message_digest_key/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object message_digest_key", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"ospf/message_digest_key"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 204: {Description: "OK"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available message_digest_key types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ospf/message_digest_key"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object message_digest_key types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ospf/message_digest_key that will be changes", Type: "", Required: true}}, Tags: []string{"ospf/message_digest_key"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object message_digest_key", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ospf/message_digest_key that will be updated", Type: "", Required: true}}, Tags: []string{"ospf/message_digest_key"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}}}}, "/objects/ospf/area/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object area", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"ospf/area"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available area types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ospf/area"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object area types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ospf/area that will be changes", Type: "", Required: true}}, Tags: []string{"ospf/area"}, Responses: map[int]struct{ Description string }{404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object area", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ospf/area that will be updated", Type: "", Required: true}}, Tags: []string{"ospf/area"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}, "/objects/ospf/group/": {"get": sophos.MethodDescriptions{Description: "Returns all available ospf/group objects", Parameters: []sophos.Parameter(nil), Tags: []string{"ospf/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new ospf/group object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ospf/group that will be created", Type: "", Required: true}}, Tags: []string{"ospf/group"}, Responses: map[int]struct{ Description string }{201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/ospf/group/{ref}": {"get": sophos.MethodDescriptions{Description: "Returns all available group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ospf/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ospf/group that will be changes", Type: "", Required: true}}, Tags: []string{"ospf/group"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ospf/group that will be updated", Type: "", Required: true}}, Tags: []string{"ospf/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"ospf/group"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/ospf/interface/{ref}": {"patch": sophos.MethodDescriptions{Description: "Changes to parts of the object interface types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ospf/interface that will be changes", Type: "", Required: true}}, Tags: []string{"ospf/interface"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object interface", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ospf/interface that will be updated", Type: "", Required: true}}, Tags: []string{"ospf/interface"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object interface", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"ospf/interface"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available interface types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ospf/interface"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}}, "/objects/ospf/interface/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ospf/interface"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 200: {Description: "OK"}, 401: {Description: "Unauthorized"}}}}, "/objects/ospf/message_digest_key/": {"get": sophos.MethodDescriptions{Description: "Returns all available ospf/message_digest_key objects", Parameters: []sophos.Parameter(nil), Tags: []string{"ospf/message_digest_key"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "post": sophos.MethodDescriptions{Description: "Create a new ospf/message_digest_key object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ospf/message_digest_key that will be created", Type: "", Required: true}}, Tags: []string{"ospf/message_digest_key"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}}}}, "/objects/ospf/message_digest_key/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ospf/message_digest_key"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/ospf/area/": {"get": sophos.MethodDescriptions{Description: "Returns all available ospf/area objects", Parameters: []sophos.Parameter(nil), Tags: []string{"ospf/area"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new ospf/area object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ospf/area that will be created", Type: "", Required: true}}, Tags: []string{"ospf/area"}, Responses: map[int]struct{ Description string }{201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/ospf/area/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ospf/area"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/ospf/group/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"ospf/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/ospf/interface/": {"get": sophos.MethodDescriptions{Description: "Returns all available ospf/interface objects", Parameters: []sophos.Parameter(nil), Tags: []string{"ospf/interface"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new ospf/interface object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "ospf/interface that will be created", Type: "", Required: true}}, Tags: []string{"ospf/interface"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 201: {Description: "OK"}}}}}}

// Definition returns the /api/definitions struct of Ospf
func (Ospf) Definition() sophos.Definition { return *defOspf }

// ApiRoutes returns all known Ospf Paths
func (Ospf) ApiRoutes() []string {
	return []string{
		"/api/objects/ospf/area/",
		"/api/objects/ospf/area/{ref}",
		"/api/objects/ospf/area/{ref}/usedby",
		"/api/objects/ospf/group/",
		"/api/objects/ospf/group/{ref}",
		"/api/objects/ospf/group/{ref}/usedby",
		"/api/objects/ospf/interface/",
		"/api/objects/ospf/interface/{ref}",
		"/api/objects/ospf/interface/{ref}/usedby",
		"/api/objects/ospf/message_digest_key/",
		"/api/objects/ospf/message_digest_key/{ref}",
		"/api/objects/ospf/message_digest_key/{ref}/usedby",
	}
}

// References returns the Ospf's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Ospf) References() []string {
	return []string{
		"REF_OspfArea",
		"REF_OspfGroup",
		"REF_OspfInterface",
		"REF_OspfMessageDigestKey",
	}
}

// OspfMessageDigestKey is an Sophos Endpoint subType and implements sophos.RestObject
type OspfMessageDigestKey []interface{}

// GetPath implements sophos.RestObject and returns the OspfMessageDigestKey GET path
// Returns all available ospf/message_digest_key objects
func (*OspfMessageDigestKey) GetPath() string { return "/api/objects/ospf/message_digest_key/" }

// RefRequired implements sophos.RestObject
func (*OspfMessageDigestKey) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the OspfMessageDigestKey DELETE path
// Creates or updates the complete object message_digest_key
func (*OspfMessageDigestKey) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/message_digest_key/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the OspfMessageDigestKey PATCH path
// Changes to parts of the object message_digest_key types
func (*OspfMessageDigestKey) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/message_digest_key/%s", ref)
}

// PostPath implements sophos.RestObject and returns the OspfMessageDigestKey POST path
// Create a new ospf/message_digest_key object
func (*OspfMessageDigestKey) PostPath() string {
	return "/api/objects/ospf/message_digest_key/"
}

// PutPath implements sophos.RestObject and returns the OspfMessageDigestKey PUT path
// Creates or updates the complete object message_digest_key
func (*OspfMessageDigestKey) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/message_digest_key/%s", ref)
}

// OspfArea is an Sophos Endpoint subType and implements sophos.RestObject
type OspfArea []interface{}

// GetPath implements sophos.RestObject and returns the OspfArea GET path
// Returns all available ospf/area objects
func (*OspfArea) GetPath() string { return "/api/objects/ospf/area/" }

// RefRequired implements sophos.RestObject
func (*OspfArea) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the OspfArea DELETE path
// Creates or updates the complete object area
func (*OspfArea) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/area/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the OspfArea PATCH path
// Changes to parts of the object area types
func (*OspfArea) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/area/%s", ref)
}

// PostPath implements sophos.RestObject and returns the OspfArea POST path
// Create a new ospf/area object
func (*OspfArea) PostPath() string {
	return "/api/objects/ospf/area/"
}

// PutPath implements sophos.RestObject and returns the OspfArea PUT path
// Creates or updates the complete object area
func (*OspfArea) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/area/%s", ref)
}

// OspfGroup is an Sophos Endpoint subType and implements sophos.RestObject
type OspfGroup []interface{}

// GetPath implements sophos.RestObject and returns the OspfGroup GET path
// Returns all available ospf/group objects
func (*OspfGroup) GetPath() string { return "/api/objects/ospf/group/" }

// RefRequired implements sophos.RestObject
func (*OspfGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the OspfGroup DELETE path
// Creates or updates the complete object group
func (*OspfGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the OspfGroup PATCH path
// Changes to parts of the object group types
func (*OspfGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the OspfGroup POST path
// Create a new ospf/group object
func (*OspfGroup) PostPath() string {
	return "/api/objects/ospf/group/"
}

// PutPath implements sophos.RestObject and returns the OspfGroup PUT path
// Creates or updates the complete object group
func (*OspfGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/group/%s", ref)
}

// OspfInterface is an Sophos Endpoint subType and implements sophos.RestObject
type OspfInterface []interface{}

// GetPath implements sophos.RestObject and returns the OspfInterface GET path
// Returns all available ospf/interface objects
func (*OspfInterface) GetPath() string { return "/api/objects/ospf/interface/" }

// RefRequired implements sophos.RestObject
func (*OspfInterface) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the OspfInterface DELETE path
// Creates or updates the complete object interface
func (*OspfInterface) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/interface/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the OspfInterface PATCH path
// Changes to parts of the object interface types
func (*OspfInterface) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/interface/%s", ref)
}

// PostPath implements sophos.RestObject and returns the OspfInterface POST path
// Create a new ospf/interface object
func (*OspfInterface) PostPath() string {
	return "/api/objects/ospf/interface/"
}

// PutPath implements sophos.RestObject and returns the OspfInterface PUT path
// Creates or updates the complete object interface
func (*OspfInterface) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ospf/interface/%s", ref)
}
