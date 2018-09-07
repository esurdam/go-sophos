// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// ApplicationControl is a generated struct representing the Sophos ApplicationControl Endpoint
// GET /api/nodes/application_control
type ApplicationControl struct {
	ApplicationControlRule  ApplicationControlRule  `json:"application_control_rule"`
	ApplicationControlGroup ApplicationControlGroup `json:"application_control_group"`
}

var defsApplicationControl = map[string]sophos.RestObject{
	"ApplicationControlRule":  &ApplicationControlRule{},
	"ApplicationControlGroup": &ApplicationControlGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of ApplicationControl's Objects
func (ApplicationControl) RestObjects() map[string]sophos.RestObject {
	return defsApplicationControl
}

// GetPath implements sophos.RestGetter
func (*ApplicationControl) GetPath() string { return "/api/nodes/application_control" }

// RefRequired implements sophos.RestGetter
func (*ApplicationControl) RefRequired() (string, bool) { return "", false }

var defApplicationControl = &sophos.Definition{Description: "application_control", Name: "application_control", Link: "/api/definitions/application_control", Swag: map[string]sophos.MethodMap{"/objects/application_control/rule/{ref}": {"put": sophos.MethodDescriptions{Description: "Creates or updates the complete object rule", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "application_control/rule that will be updated", Type: "", Required: true}}, Tags: []string{"application_control/rule"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object rule", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"application_control/rule"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available rule types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"application_control/rule"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object rule types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "application_control/rule that will be changes", Type: "", Required: true}}, Tags: []string{"application_control/rule"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}, "/objects/application_control/rule/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"application_control/rule"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}}}}, "/objects/application_control/group/": {"get": sophos.MethodDescriptions{Description: "Returns all available application_control/group objects", Parameters: []sophos.Parameter(nil), Tags: []string{"application_control/group"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}}}, "post": sophos.MethodDescriptions{Description: "Create a new application_control/group object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "application_control/group that will be created", Type: "", Required: true}}, Tags: []string{"application_control/group"}, Responses: map[int]struct{ Description string }{201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/application_control/group/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"application_control/group"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 204: {Description: "OK"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"application_control/group"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "application_control/group that will be changes", Type: "", Required: true}}, Tags: []string{"application_control/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "application_control/group that will be updated", Type: "", Required: true}}, Tags: []string{"application_control/group"}, Responses: map[int]struct{ Description string }{404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/application_control/group/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"application_control/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/application_control/rule/": {"get": sophos.MethodDescriptions{Description: "Returns all available application_control/rule objects", Parameters: []sophos.Parameter(nil), Tags: []string{"application_control/rule"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "post": sophos.MethodDescriptions{Description: "Create a new application_control/rule object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "application_control/rule that will be created", Type: "", Required: true}}, Tags: []string{"application_control/rule"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 201: {Description: "OK"}}}}}}

// Definition returns the /api/definitions struct of ApplicationControl
func (ApplicationControl) Definition() sophos.Definition { return *defApplicationControl }

// ApiRoutes returns all known ApplicationControl Paths
func (ApplicationControl) ApiRoutes() []string {
	return []string{
		"/api/objects/application_control/group/",
		"/api/objects/application_control/group/{ref}",
		"/api/objects/application_control/group/{ref}/usedby",
		"/api/objects/application_control/rule/",
		"/api/objects/application_control/rule/{ref}",
		"/api/objects/application_control/rule/{ref}/usedby",
	}
}

// References returns the ApplicationControl's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (ApplicationControl) References() []string {
	return []string{
		"REF_ApplicationControlGroup",
		"REF_ApplicationControlRule",
	}
}

// ApplicationControlRule is an Sophos Endpoint subType and implements sophos.RestObject
type ApplicationControlRules []ApplicationControlRule
type ApplicationControlRule struct {
	Locked                  string        `json:"_locked"`
	Reference               string        `json:"_ref"`
	_type                   string        `json:"_type"`
	Action                  string        `json:"action"`
	Applications            []string      `json:"applications"`
	Comment                 string        `json:"comment"`
	DestinationNetworks     []string      `json:"destination_networks"`
	Group                   string        `json:"group"`
	GroupFilterProductivity int64         `json:"group_filter_productivity"`
	GroupFilterRisk         int64         `json:"group_filter_risk"`
	Groups                  []interface{} `json:"groups"`
	Log                     bool          `json:"log"`
	Name                    string        `json:"name"`
	SourceNetworks          []string      `json:"source_networks"`
	Status                  bool          `json:"status"`
}

// GetPath implements sophos.RestObject and returns the ApplicationControlRules GET path
// Returns all available application_control/rule objects
func (*ApplicationControlRules) GetPath() string { return "/api/objects/application_control/rule/" }

// RefRequired implements sophos.RestObject
func (*ApplicationControlRules) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ApplicationControlRules GET path
// Returns all available rule types
func (a *ApplicationControlRule) GetPath() string {
	return fmt.Sprintf("/api/objects/application_control/rule/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *ApplicationControlRule) RefRequired() (string, bool) { return a.Reference, true }

// DeletePath implements sophos.RestObject and returns the ApplicationControlRule DELETE path
// Creates or updates the complete object rule
func (*ApplicationControlRule) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/application_control/rule/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ApplicationControlRule PATCH path
// Changes to parts of the object rule types
func (*ApplicationControlRule) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/application_control/rule/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ApplicationControlRule POST path
// Create a new application_control/rule object
func (*ApplicationControlRule) PostPath() string {
	return "/api/objects/application_control/rule/"
}

// PutPath implements sophos.RestObject and returns the ApplicationControlRule PUT path
// Creates or updates the complete object rule
func (*ApplicationControlRule) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/application_control/rule/%s", ref)
}

// Type implements sophos.Object
func (a *ApplicationControlRule) GetType() string { return a._type }

// ApplicationControlGroup is an Sophos Endpoint subType and implements sophos.RestObject
type ApplicationControlGroup []interface{}

// GetPath implements sophos.RestObject and returns the ApplicationControlGroup GET path
// Returns all available application_control/group objects
func (*ApplicationControlGroup) GetPath() string { return "/api/objects/application_control/group/" }

// RefRequired implements sophos.RestObject
func (*ApplicationControlGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ApplicationControlGroup DELETE path
// Creates or updates the complete object group
func (*ApplicationControlGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/application_control/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ApplicationControlGroup PATCH path
// Changes to parts of the object group types
func (*ApplicationControlGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/application_control/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ApplicationControlGroup POST path
// Create a new application_control/group object
func (*ApplicationControlGroup) PostPath() string {
	return "/api/objects/application_control/group/"
}

// PutPath implements sophos.RestObject and returns the ApplicationControlGroup PUT path
// Creates or updates the complete object group
func (*ApplicationControlGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/application_control/group/%s", ref)
}
