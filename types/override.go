// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Override is a generated struct representing the Sophos Override Endpoint
// GET /api/nodes/override
type Override struct {
	OverrideObjref OverrideObjref `json:"override_objref"`
	OverrideGroup  OverrideGroup  `json:"override_group"`
}

var defsOverride = map[string]sophos.RestObject{
	"OverrideObjref": &OverrideObjref{},
	"OverrideGroup":  &OverrideGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of Override's Objects
func (Override) RestObjects() map[string]sophos.RestObject {
	return defsOverride
}

// GetPath implements sophos.RestGetter
func (*Override) GetPath() string { return "/api/nodes/override" }

// RefRequired implements sophos.RestGetter
func (*Override) RefRequired() (string, bool) { return "", false }

var defOverride = &sophos.Definition{Description: "override", Name: "override", Link: "/api/definitions/override", Swag: map[string]sophos.MethodMap{"/objects/override/objref/{ref}": {"get": sophos.MethodDescriptions{Description: "Returns all available objref types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"override/objref"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object objref types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "override/objref that will be changes", Type: "", Required: true}}, Tags: []string{"override/objref"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object objref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "override/objref that will be updated", Type: "", Required: true}}, Tags: []string{"override/objref"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object objref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"override/objref"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/override/objref/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"override/objref"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/override/group/": {"get": sophos.MethodDescriptions{Description: "Returns all available override/group objects", Parameters: []sophos.Parameter(nil), Tags: []string{"override/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new override/group object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "override/group that will be created", Type: "", Required: true}}, Tags: []string{"override/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}}}}, "/objects/override/group/{ref}": {"put": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "override/group that will be updated", Type: "", Required: true}}, Tags: []string{"override/group"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}}}, "delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"override/group"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"override/group"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "override/group that will be changes", Type: "", Required: true}}, Tags: []string{"override/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}, "/objects/override/group/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"override/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/override/objref/": {"get": sophos.MethodDescriptions{Description: "Returns all available override/objref objects", Parameters: []sophos.Parameter(nil), Tags: []string{"override/objref"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new override/objref object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "override/objref that will be created", Type: "", Required: true}}, Tags: []string{"override/objref"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}}}}

// Definition returns the /api/definitions struct of Override
func (Override) Definition() sophos.Definition { return *defOverride }

// ApiRoutes returns all known Override Paths
func (Override) ApiRoutes() []string {
	return []string{
		"/api/objects/override/group/",
		"/api/objects/override/group/{ref}",
		"/api/objects/override/group/{ref}/usedby",
		"/api/objects/override/objref/",
		"/api/objects/override/objref/{ref}",
		"/api/objects/override/objref/{ref}/usedby",
	}
}

// References returns the Override's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Override) References() []string {
	return []string{
		"REF_OverrideGroup",
		"REF_OverrideObjref",
	}
}

// OverrideObjref is an Sophos Endpoint subType and implements sophos.RestObject
type OverrideObjref []interface{}

// GetPath implements sophos.RestObject and returns the OverrideObjref GET path
// Returns all available override/objref objects
func (*OverrideObjref) GetPath() string { return "/api/objects/override/objref/" }

// RefRequired implements sophos.RestObject
func (*OverrideObjref) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the OverrideObjref DELETE path
// Creates or updates the complete object objref
func (*OverrideObjref) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/override/objref/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the OverrideObjref PATCH path
// Changes to parts of the object objref types
func (*OverrideObjref) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/override/objref/%s", ref)
}

// PostPath implements sophos.RestObject and returns the OverrideObjref POST path
// Create a new override/objref object
func (*OverrideObjref) PostPath() string {
	return "/api/objects/override/objref/"
}

// PutPath implements sophos.RestObject and returns the OverrideObjref PUT path
// Creates or updates the complete object objref
func (*OverrideObjref) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/override/objref/%s", ref)
}

// OverrideGroup is an Sophos Endpoint subType and implements sophos.RestObject
type OverrideGroup []interface{}

// GetPath implements sophos.RestObject and returns the OverrideGroup GET path
// Returns all available override/group objects
func (*OverrideGroup) GetPath() string { return "/api/objects/override/group/" }

// RefRequired implements sophos.RestObject
func (*OverrideGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the OverrideGroup DELETE path
// Creates or updates the complete object group
func (*OverrideGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/override/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the OverrideGroup PATCH path
// Changes to parts of the object group types
func (*OverrideGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/override/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the OverrideGroup POST path
// Create a new override/group object
func (*OverrideGroup) PostPath() string {
	return "/api/objects/override/group/"
}

// PutPath implements sophos.RestObject and returns the OverrideGroup PUT path
// Creates or updates the complete object group
func (*OverrideGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/override/group/%s", ref)
}
