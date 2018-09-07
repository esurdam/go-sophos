// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Route is a generated struct representing the Sophos Route Endpoint
// GET /api/nodes/route
type Route struct {
	RoutePolicy RoutePolicy `json:"route_policy"`
	RouteStatic RouteStatic `json:"route_static"`
	RouteGroup  RouteGroup  `json:"route_group"`
}

var defsRoute = map[string]sophos.RestObject{
	"RoutePolicy": &RoutePolicy{},
	"RouteStatic": &RouteStatic{},
	"RouteGroup":  &RouteGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of Route's Objects
func (Route) RestObjects() map[string]sophos.RestObject {
	return defsRoute
}

// GetPath implements sophos.RestGetter
func (*Route) GetPath() string { return "/api/nodes/route" }

// RefRequired implements sophos.RestGetter
func (*Route) RefRequired() (string, bool) { return "", false }

var defRoute = &sophos.Definition{Description: "route", Name: "route", Link: "/api/definitions/route", Swag: map[string]sophos.MethodMap{"/objects/route/group/": {"post": sophos.MethodDescriptions{Description: "Create a new route/group object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "route/group that will be created", Type: "", Required: true}}, Tags: []string{"route/group"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available route/group objects", Parameters: []sophos.Parameter(nil), Tags: []string{"route/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/route/group/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"route/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/route/policy/": {"get": sophos.MethodDescriptions{Description: "Returns all available route/policy objects", Parameters: []sophos.Parameter(nil), Tags: []string{"route/policy"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "post": sophos.MethodDescriptions{Description: "Create a new route/policy object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "route/policy that will be created", Type: "", Required: true}}, Tags: []string{"route/policy"}, Responses: map[int]struct{ Description string }{201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/route/policy/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object policy", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"route/policy"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 204: {Description: "OK"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available policy types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"route/policy"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object policy types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "route/policy that will be changes", Type: "", Required: true}}, Tags: []string{"route/policy"}, Responses: map[int]struct{ Description string }{404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object policy", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "route/policy that will be updated", Type: "", Required: true}}, Tags: []string{"route/policy"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}, "/objects/route/group/{ref}": {"patch": sophos.MethodDescriptions{Description: "Changes to parts of the object group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "route/group that will be changes", Type: "", Required: true}}, Tags: []string{"route/group"}, Responses: map[int]struct{ Description string }{404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "route/group that will be updated", Type: "", Required: true}}, Tags: []string{"route/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"route/group"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"route/group"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}}, "/objects/route/policy/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"route/policy"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}}}}, "/objects/route/static/": {"get": sophos.MethodDescriptions{Description: "Returns all available route/static objects", Parameters: []sophos.Parameter(nil), Tags: []string{"route/static"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new route/static object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "route/static that will be created", Type: "", Required: true}}, Tags: []string{"route/static"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}}}}, "/objects/route/static/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object static", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"route/static"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available static types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"route/static"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object static types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "route/static that will be changes", Type: "", Required: true}}, Tags: []string{"route/static"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object static", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "route/static that will be updated", Type: "", Required: true}}, Tags: []string{"route/static"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}}}}, "/objects/route/static/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"route/static"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}}}

// Definition returns the /api/definitions struct of Route
func (Route) Definition() sophos.Definition { return *defRoute }

// ApiRoutes returns all known Route Paths
func (Route) ApiRoutes() []string {
	return []string{
		"/api/objects/route/group/",
		"/api/objects/route/group/{ref}",
		"/api/objects/route/group/{ref}/usedby",
		"/api/objects/route/policy/",
		"/api/objects/route/policy/{ref}",
		"/api/objects/route/policy/{ref}/usedby",
		"/api/objects/route/static/",
		"/api/objects/route/static/{ref}",
		"/api/objects/route/static/{ref}/usedby",
	}
}

// References returns the Route's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Route) References() []string {
	return []string{
		"REF_RouteGroup",
		"REF_RoutePolicy",
		"REF_RouteStatic",
	}
}

// RoutePolicy is an Sophos Endpoint subType and implements sophos.RestObject
type RoutePolicy []interface{}

// GetPath implements sophos.RestObject and returns the RoutePolicy GET path
// Returns all available route/policy objects
func (*RoutePolicy) GetPath() string { return "/api/objects/route/policy/" }

// RefRequired implements sophos.RestObject
func (*RoutePolicy) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the RoutePolicy DELETE path
// Creates or updates the complete object policy
func (*RoutePolicy) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/route/policy/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the RoutePolicy PATCH path
// Changes to parts of the object policy types
func (*RoutePolicy) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/route/policy/%s", ref)
}

// PostPath implements sophos.RestObject and returns the RoutePolicy POST path
// Create a new route/policy object
func (*RoutePolicy) PostPath() string {
	return "/api/objects/route/policy/"
}

// PutPath implements sophos.RestObject and returns the RoutePolicy PUT path
// Creates or updates the complete object policy
func (*RoutePolicy) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/route/policy/%s", ref)
}

// RouteStatic is an Sophos Endpoint subType and implements sophos.RestObject
type RouteStatics []RouteStatic
type RouteStatic struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Metric    int64  `json:"metric"`
	Name      string `json:"name"`
	Network   string `json:"network"`
	Status    bool   `json:"status"`
	Target    string `json:"target"`
	Type      string `json:"type"`
}

// GetPath implements sophos.RestObject and returns the RouteStatics GET path
// Returns all available route/static objects
func (*RouteStatics) GetPath() string { return "/api/objects/route/static/" }

// RefRequired implements sophos.RestObject
func (*RouteStatics) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the RouteStatics GET path
// Returns all available static types
func (r *RouteStatic) GetPath() string {
	return fmt.Sprintf("/api/objects/route/static/%s", r.Reference)
}

// RefRequired implements sophos.RestObject
func (r *RouteStatic) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the RouteStatic DELETE path
// Creates or updates the complete object static
func (*RouteStatic) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/route/static/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the RouteStatic PATCH path
// Changes to parts of the object static types
func (*RouteStatic) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/route/static/%s", ref)
}

// PostPath implements sophos.RestObject and returns the RouteStatic POST path
// Create a new route/static object
func (*RouteStatic) PostPath() string {
	return "/api/objects/route/static/"
}

// PutPath implements sophos.RestObject and returns the RouteStatic PUT path
// Creates or updates the complete object static
func (*RouteStatic) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/route/static/%s", ref)
}

// Type implements sophos.Object
func (r *RouteStatic) GetType() string { return r._type }

// RouteGroup is an Sophos Endpoint subType and implements sophos.RestObject
type RouteGroup []interface{}

// GetPath implements sophos.RestObject and returns the RouteGroup GET path
// Returns all available route/group objects
func (*RouteGroup) GetPath() string { return "/api/objects/route/group/" }

// RefRequired implements sophos.RestObject
func (*RouteGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the RouteGroup DELETE path
// Creates or updates the complete object group
func (*RouteGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/route/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the RouteGroup PATCH path
// Changes to parts of the object group types
func (*RouteGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/route/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the RouteGroup POST path
// Create a new route/group object
func (*RouteGroup) PostPath() string {
	return "/api/objects/route/group/"
}

// PutPath implements sophos.RestObject and returns the RouteGroup PUT path
// Creates or updates the complete object group
func (*RouteGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/route/group/%s", ref)
}
