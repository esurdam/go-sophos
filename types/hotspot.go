// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Hotspot is a generated struct representing the Sophos Hotspot Endpoint
// GET /api/nodes/hotspot
type Hotspot struct {
	Cert            string        `json:"cert"`
	DeleteDays      int64         `json:"delete_days"`
	SslPortal       int64         `json:"ssl_portal"`
	Status          int64         `json:"status"`
	TransparentSkip []interface{} `json:"transparent_skip"`
}

var defsHotspot = map[string]sophos.RestObject{
	"HotspotVoucher": &HotspotVoucher{},
	"HotspotGroup":   &HotspotGroup{},
	"HotspotPortal":  &HotspotPortal{},
}

// RestObjects implements the sophos.Node interface and returns a map of Hotspot's Objects
func (Hotspot) RestObjects() map[string]sophos.RestObject {
	return defsHotspot
}

// GetPath implements sophos.RestGetter
func (*Hotspot) GetPath() string { return "/api/nodes/hotspot" }

// RefRequired implements sophos.RestGetter
func (*Hotspot) RefRequired() (string, bool) { return "", false }

var defHotspot = &sophos.Definition{Description: "hotspot", Name: "hotspot", Link: "/api/definitions/hotspot", Swag: map[string]sophos.MethodMap{"/objects/hotspot/group/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"hotspot/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/hotspot/portal/": {"get": sophos.MethodDescriptions{Description: "Returns all available hotspot/portal objects", Parameters: []sophos.Parameter(nil), Tags: []string{"hotspot/portal"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new hotspot/portal object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "hotspot/portal that will be created", Type: "", Required: true}}, Tags: []string{"hotspot/portal"}, Responses: map[int]struct{ Description string }{201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/hotspot/voucher/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"hotspot/voucher"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 200: {Description: "OK"}, 401: {Description: "Unauthorized"}}}}, "/objects/hotspot/voucher/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object voucher", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"hotspot/voucher"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 204: {Description: "OK"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available voucher types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"hotspot/voucher"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object voucher types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "hotspot/voucher that will be changes", Type: "", Required: true}}, Tags: []string{"hotspot/voucher"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object voucher", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "hotspot/voucher that will be updated", Type: "", Required: true}}, Tags: []string{"hotspot/voucher"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}, "/objects/hotspot/group/": {"post": sophos.MethodDescriptions{Description: "Create a new hotspot/group object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "hotspot/group that will be created", Type: "", Required: true}}, Tags: []string{"hotspot/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available hotspot/group objects", Parameters: []sophos.Parameter(nil), Tags: []string{"hotspot/group"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}}}}, "/objects/hotspot/group/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"hotspot/group"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"hotspot/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "hotspot/group that will be changes", Type: "", Required: true}}, Tags: []string{"hotspot/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "hotspot/group that will be updated", Type: "", Required: true}}, Tags: []string{"hotspot/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}, "/objects/hotspot/portal/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object portal", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"hotspot/portal"}, Responses: map[int]struct{ Description string }{400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 204: {Description: "OK"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available portal types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"hotspot/portal"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object portal types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "hotspot/portal that will be changes", Type: "", Required: true}}, Tags: []string{"hotspot/portal"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object portal", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "hotspot/portal that will be updated", Type: "", Required: true}}, Tags: []string{"hotspot/portal"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}}, "/objects/hotspot/portal/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"hotspot/portal"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/hotspot/voucher/": {"get": sophos.MethodDescriptions{Description: "Returns all available hotspot/voucher objects", Parameters: []sophos.Parameter(nil), Tags: []string{"hotspot/voucher"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new hotspot/voucher object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "hotspot/voucher that will be created", Type: "", Required: true}}, Tags: []string{"hotspot/voucher"}, Responses: map[int]struct{ Description string }{201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}}}

// Definition returns the /api/definitions struct of Hotspot
func (Hotspot) Definition() sophos.Definition { return *defHotspot }

// ApiRoutes returns all known Hotspot Paths
func (Hotspot) ApiRoutes() []string {
	return []string{
		"/api/objects/hotspot/group/",
		"/api/objects/hotspot/group/{ref}",
		"/api/objects/hotspot/group/{ref}/usedby",
		"/api/objects/hotspot/portal/",
		"/api/objects/hotspot/portal/{ref}",
		"/api/objects/hotspot/portal/{ref}/usedby",
		"/api/objects/hotspot/voucher/",
		"/api/objects/hotspot/voucher/{ref}",
		"/api/objects/hotspot/voucher/{ref}/usedby",
	}
}

// References returns the Hotspot's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Hotspot) References() []string {
	return []string{
		"REF_HotspotGroup",
		"REF_HotspotPortal",
		"REF_HotspotVoucher",
	}
}

// HotspotVoucher is an Sophos Endpoint subType and implements sophos.RestObject
type HotspotVoucher []interface{}

// GetPath implements sophos.RestObject and returns the HotspotVoucher GET path
// Returns all available hotspot/voucher objects
func (*HotspotVoucher) GetPath() string { return "/api/objects/hotspot/voucher/" }

// RefRequired implements sophos.RestObject
func (*HotspotVoucher) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the HotspotVoucher DELETE path
// Creates or updates the complete object voucher
func (*HotspotVoucher) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/voucher/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HotspotVoucher PATCH path
// Changes to parts of the object voucher types
func (*HotspotVoucher) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/voucher/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HotspotVoucher POST path
// Create a new hotspot/voucher object
func (*HotspotVoucher) PostPath() string {
	return "/api/objects/hotspot/voucher/"
}

// PutPath implements sophos.RestObject and returns the HotspotVoucher PUT path
// Creates or updates the complete object voucher
func (*HotspotVoucher) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/voucher/%s", ref)
}

// HotspotGroup is an Sophos Endpoint subType and implements sophos.RestObject
type HotspotGroup []interface{}

// GetPath implements sophos.RestObject and returns the HotspotGroup GET path
// Returns all available hotspot/group objects
func (*HotspotGroup) GetPath() string { return "/api/objects/hotspot/group/" }

// RefRequired implements sophos.RestObject
func (*HotspotGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the HotspotGroup DELETE path
// Creates or updates the complete object group
func (*HotspotGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HotspotGroup PATCH path
// Changes to parts of the object group types
func (*HotspotGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HotspotGroup POST path
// Create a new hotspot/group object
func (*HotspotGroup) PostPath() string {
	return "/api/objects/hotspot/group/"
}

// PutPath implements sophos.RestObject and returns the HotspotGroup PUT path
// Creates or updates the complete object group
func (*HotspotGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/group/%s", ref)
}

// HotspotPortal is an Sophos Endpoint subType and implements sophos.RestObject
type HotspotPortal []interface{}

// GetPath implements sophos.RestObject and returns the HotspotPortal GET path
// Returns all available hotspot/portal objects
func (*HotspotPortal) GetPath() string { return "/api/objects/hotspot/portal/" }

// RefRequired implements sophos.RestObject
func (*HotspotPortal) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the HotspotPortal DELETE path
// Creates or updates the complete object portal
func (*HotspotPortal) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/portal/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HotspotPortal PATCH path
// Changes to parts of the object portal types
func (*HotspotPortal) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/portal/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HotspotPortal POST path
// Create a new hotspot/portal object
func (*HotspotPortal) PostPath() string {
	return "/api/objects/hotspot/portal/"
}

// PutPath implements sophos.RestObject and returns the HotspotPortal PUT path
// Creates or updates the complete object portal
func (*HotspotPortal) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/hotspot/portal/%s", ref)
}
