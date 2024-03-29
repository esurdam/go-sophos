package objects

// THIS FILE IS AUTO GENERATED BY bin/gen.go! DO NOT EDIT!

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Stas is a generated struct representing the Sophos Stas Endpoint
// GET /api/nodes/stas
type Stas struct {
	StasCollector StasCollector `json:"stas_collector"`
	StasGroup     StasGroup     `json:"stas_group"`
}

var _ sophos.Endpoint = &Stas{}

var defsStas = map[string]sophos.RestObject{
	"StasCollector": &StasCollector{},
	"StasGroup":     &StasGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of Stas's Objects
func (Stas) RestObjects() map[string]sophos.RestObject { return defsStas }

// GetPath implements sophos.RestGetter
func (*Stas) GetPath() string { return "/api/nodes/stas" }

// RefRequired implements sophos.RestGetter
func (*Stas) RefRequired() (string, bool) { return "", false }

var defStas = &sophos.Definition{Description: "stas", Name: "stas", Link: "/api/definitions/stas"}

// Definition returns the /api/definitions struct of Stas
func (Stas) Definition() sophos.Definition { return *defStas }

// ApiRoutes returns all known Stas Paths
func (Stas) ApiRoutes() []string {
	return []string{
		"/api/objects/stas/collector/",
		"/api/objects/stas/collector/{ref}",
		"/api/objects/stas/collector/{ref}/usedby",
		"/api/objects/stas/group/",
		"/api/objects/stas/group/{ref}",
		"/api/objects/stas/group/{ref}/usedby",
	}
}

// References returns the Stas's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Stas) References() []string {
	return []string{
		"REF_StasCollector",
		"REF_StasGroup",
	}
}

// StasCollectors is an Sophos Endpoint subType and implements sophos.RestObject
type StasCollectors []StasCollector

// StasCollector represents a UTM STAS Collector
type StasCollector struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	// Host description: REF(network/host), REF(network/dns_host)
	Host string `json:"host"`
	Name string `json:"name"`
	// Port description: REF(service/udp)
	// Port default value is "REF_ServiceSTASCollector"
	Port string `json:"port"`
}

var _ sophos.RestGetter = &StasCollector{}

// GetPath implements sophos.RestObject and returns the StasCollectors GET path
// Returns all available stas/collector objects
func (*StasCollectors) GetPath() string { return "/api/objects/stas/collector/" }

// RefRequired implements sophos.RestObject
func (*StasCollectors) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the StasCollectors GET path
// Returns all available collector types
func (s *StasCollector) GetPath() string {
	return fmt.Sprintf("/api/objects/stas/collector/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *StasCollector) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the StasCollector DELETE path
// Creates or updates the complete object collector
func (*StasCollector) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/stas/collector/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the StasCollector PATCH path
// Changes to parts of the object collector types
func (*StasCollector) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/stas/collector/%s", ref)
}

// PostPath implements sophos.RestObject and returns the StasCollector POST path
// Create a new stas/collector object
func (*StasCollector) PostPath() string {
	return "/api/objects/stas/collector/"
}

// PutPath implements sophos.RestObject and returns the StasCollector PUT path
// Creates or updates the complete object collector
func (*StasCollector) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/stas/collector/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*StasCollector) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/stas/collector/%s/usedby", ref)
}

// StasGroups is an Sophos Endpoint subType and implements sophos.RestObject
type StasGroups []StasGroup

// StasGroup represents a UTM stas->group
type StasGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &StasGroup{}

// GetPath implements sophos.RestObject and returns the StasGroups GET path
// Returns all available stas/group objects
func (*StasGroups) GetPath() string { return "/api/objects/stas/group/" }

// RefRequired implements sophos.RestObject
func (*StasGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the StasGroups GET path
// Returns all available group types
func (s *StasGroup) GetPath() string { return fmt.Sprintf("/api/objects/stas/group/%s", s.Reference) }

// RefRequired implements sophos.RestObject
func (s *StasGroup) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the StasGroup DELETE path
// Creates or updates the complete object group
func (*StasGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/stas/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the StasGroup PATCH path
// Changes to parts of the object group types
func (*StasGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/stas/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the StasGroup POST path
// Create a new stas/group object
func (*StasGroup) PostPath() string {
	return "/api/objects/stas/group/"
}

// PutPath implements sophos.RestObject and returns the StasGroup PUT path
// Creates or updates the complete object group
func (*StasGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/stas/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*StasGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/stas/group/%s/usedby", ref)
}
