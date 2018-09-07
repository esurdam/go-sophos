// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Aws is a generated struct representing the Sophos Aws Endpoint
// GET /api/nodes/aws
type Aws struct {
	AwsRegion       AwsRegion       `json:"aws_region"`
	AwsGroup        AwsGroup        `json:"aws_group"`
	AwsInstanceType AwsInstanceType `json:"aws_instance_type"`
}

var defsAws = map[string]sophos.RestObject{
	"AwsRegion":       &AwsRegion{},
	"AwsGroup":        &AwsGroup{},
	"AwsInstanceType": &AwsInstanceType{},
}

// RestObjects implements the sophos.Node interface and returns a map of Aws's Objects
func (Aws) RestObjects() map[string]sophos.RestObject {
	return defsAws
}

// GetPath implements sophos.RestGetter
func (*Aws) GetPath() string { return "/api/nodes/aws" }

// RefRequired implements sophos.RestGetter
func (*Aws) RefRequired() (string, bool) { return "", false }

var defAws = &sophos.Definition{Description: "aws", Name: "aws", Link: "/api/definitions/aws", Swag: map[string]sophos.MethodMap{"/objects/aws/region/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object region", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"aws/region"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available region types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"aws/region"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object region types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "aws/region that will be changes", Type: "", Required: true}}, Tags: []string{"aws/region"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object region", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "aws/region that will be updated", Type: "", Required: true}}, Tags: []string{"aws/region"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}, "/objects/aws/region/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"aws/region"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}}}}, "/objects/aws/group/{ref}": {"get": sophos.MethodDescriptions{Description: "Returns all available group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"aws/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "aws/group that will be changes", Type: "", Required: true}}, Tags: []string{"aws/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "aws/group that will be updated", Type: "", Required: true}}, Tags: []string{"aws/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"aws/group"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/aws/group/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"aws/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/aws/instance_type/": {"get": sophos.MethodDescriptions{Description: "Returns all available aws/instance_type objects", Parameters: []sophos.Parameter(nil), Tags: []string{"aws/instance_type"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new aws/instance_type object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "aws/instance_type that will be created", Type: "", Required: true}}, Tags: []string{"aws/instance_type"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}}}}, "/objects/aws/instance_type/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object instance_type", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"aws/instance_type"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available instance_type types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"aws/instance_type"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object instance_type types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "aws/instance_type that will be changes", Type: "", Required: true}}, Tags: []string{"aws/instance_type"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object instance_type", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "aws/instance_type that will be updated", Type: "", Required: true}}, Tags: []string{"aws/instance_type"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}, "/objects/aws/instance_type/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"aws/instance_type"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/aws/region/": {"get": sophos.MethodDescriptions{Description: "Returns all available aws/region objects", Parameters: []sophos.Parameter(nil), Tags: []string{"aws/region"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "post": sophos.MethodDescriptions{Description: "Create a new aws/region object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "aws/region that will be created", Type: "", Required: true}}, Tags: []string{"aws/region"}, Responses: map[int]struct{ Description string }{201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/aws/group/": {"get": sophos.MethodDescriptions{Description: "Returns all available aws/group objects", Parameters: []sophos.Parameter(nil), Tags: []string{"aws/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "post": sophos.MethodDescriptions{Description: "Create a new aws/group object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "aws/group that will be created", Type: "", Required: true}}, Tags: []string{"aws/group"}, Responses: map[int]struct{ Description string }{201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}}}

// Definition returns the /api/definitions struct of Aws
func (Aws) Definition() sophos.Definition { return *defAws }

// ApiRoutes returns all known Aws Paths
func (Aws) ApiRoutes() []string {
	return []string{
		"/api/objects/aws/group/",
		"/api/objects/aws/group/{ref}",
		"/api/objects/aws/group/{ref}/usedby",
		"/api/objects/aws/instance_type/",
		"/api/objects/aws/instance_type/{ref}",
		"/api/objects/aws/instance_type/{ref}/usedby",
		"/api/objects/aws/region/",
		"/api/objects/aws/region/{ref}",
		"/api/objects/aws/region/{ref}/usedby",
	}
}

// References returns the Aws's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Aws) References() []string {
	return []string{
		"REF_AwsGroup",
		"REF_AwsInstanceType",
		"REF_AwsRegion",
	}
}

// AwsRegion is an Sophos Endpoint subType and implements sophos.RestObject
type AwsRegions []AwsRegion
type AwsRegion struct {
	Locked            string   `json:"_locked"`
	Reference         string   `json:"_ref"`
	_type             string   `json:"_type"`
	AvailabilityZones []string `json:"availability_zones"`
	Code              string   `json:"code"`
	Comment           string   `json:"comment"`
	InstanceTypes     []string `json:"instance_types"`
	Name              string   `json:"name"`
	Partition         string   `json:"partition"`
}

// GetPath implements sophos.RestObject and returns the AwsRegions GET path
// Returns all available aws/region objects
func (*AwsRegions) GetPath() string { return "/api/objects/aws/region/" }

// RefRequired implements sophos.RestObject
func (*AwsRegions) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AwsRegions GET path
// Returns all available region types
func (a *AwsRegion) GetPath() string { return fmt.Sprintf("/api/objects/aws/region/%s", a.Reference) }

// RefRequired implements sophos.RestObject
func (a *AwsRegion) RefRequired() (string, bool) { return a.Reference, true }

// DeletePath implements sophos.RestObject and returns the AwsRegion DELETE path
// Creates or updates the complete object region
func (*AwsRegion) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/aws/region/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AwsRegion PATCH path
// Changes to parts of the object region types
func (*AwsRegion) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/aws/region/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AwsRegion POST path
// Create a new aws/region object
func (*AwsRegion) PostPath() string {
	return "/api/objects/aws/region/"
}

// PutPath implements sophos.RestObject and returns the AwsRegion PUT path
// Creates or updates the complete object region
func (*AwsRegion) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/aws/region/%s", ref)
}

// Type implements sophos.Object
func (a *AwsRegion) GetType() string { return a._type }

// AwsGroup is an Sophos Endpoint subType and implements sophos.RestObject
type AwsGroup []interface{}

// GetPath implements sophos.RestObject and returns the AwsGroup GET path
// Returns all available aws/group objects
func (*AwsGroup) GetPath() string { return "/api/objects/aws/group/" }

// RefRequired implements sophos.RestObject
func (*AwsGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AwsGroup DELETE path
// Creates or updates the complete object group
func (*AwsGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/aws/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AwsGroup PATCH path
// Changes to parts of the object group types
func (*AwsGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/aws/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AwsGroup POST path
// Create a new aws/group object
func (*AwsGroup) PostPath() string {
	return "/api/objects/aws/group/"
}

// PutPath implements sophos.RestObject and returns the AwsGroup PUT path
// Creates or updates the complete object group
func (*AwsGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/aws/group/%s", ref)
}

// AwsInstanceType is an Sophos Endpoint subType and implements sophos.RestObject
type AwsInstanceTypes []AwsInstanceType
type AwsInstanceType struct {
	Locked             string      `json:"_locked"`
	Reference          string      `json:"_ref"`
	_type              string      `json:"_type"`
	Comment            string      `json:"comment"`
	CPUCores           int64       `json:"cpu_cores"`
	Deprecated         bool        `json:"deprecated"`
	MemoryBytes        interface{} `json:"memory_bytes"`
	Model              string      `json:"model"`
	Name               string      `json:"name"`
	NetworkPerformance string      `json:"network_performance"`
}

// GetPath implements sophos.RestObject and returns the AwsInstanceTypes GET path
// Returns all available aws/instance_type objects
func (*AwsInstanceTypes) GetPath() string { return "/api/objects/aws/instance_type/" }

// RefRequired implements sophos.RestObject
func (*AwsInstanceTypes) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AwsInstanceTypes GET path
// Returns all available instance_type types
func (a *AwsInstanceType) GetPath() string {
	return fmt.Sprintf("/api/objects/aws/instance_type/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AwsInstanceType) RefRequired() (string, bool) { return a.Reference, true }

// DeletePath implements sophos.RestObject and returns the AwsInstanceType DELETE path
// Creates or updates the complete object instance_type
func (*AwsInstanceType) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/aws/instance_type/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AwsInstanceType PATCH path
// Changes to parts of the object instance_type types
func (*AwsInstanceType) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/aws/instance_type/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AwsInstanceType POST path
// Create a new aws/instance_type object
func (*AwsInstanceType) PostPath() string {
	return "/api/objects/aws/instance_type/"
}

// PutPath implements sophos.RestObject and returns the AwsInstanceType PUT path
// Creates or updates the complete object instance_type
func (*AwsInstanceType) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/aws/instance_type/%s", ref)
}

// Type implements sophos.Object
func (a *AwsInstanceType) GetType() string { return a._type }
