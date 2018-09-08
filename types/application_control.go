package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// ApplicationControl is a generated struct representing the Sophos ApplicationControl Endpoint
// GET /api/nodes/application_control
type ApplicationControl struct {
	ApplicationControlGroup ApplicationControlGroup `json:"application_control_group"`
	ApplicationControlRule  ApplicationControlRule  `json:"application_control_rule"`
}

var defsApplicationControl = map[string]sophos.RestObject{
	"ApplicationControlGroup": &ApplicationControlGroup{},
	"ApplicationControlRule":  &ApplicationControlRule{},
}

// RestObjects implements the sophos.Node interface and returns a map of ApplicationControl's Objects
func (ApplicationControl) RestObjects() map[string]sophos.RestObject { return defsApplicationControl }

// GetPath implements sophos.RestGetter
func (*ApplicationControl) GetPath() string { return "/api/nodes/application_control" }

// RefRequired implements sophos.RestGetter
func (*ApplicationControl) RefRequired() (string, bool) { return "", false }

var defApplicationControl = &sophos.Definition{Description: "application_control", Name: "application_control", Link: "/api/definitions/application_control"}

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

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*ApplicationControlGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/application_control/group/%s/usedby", ref)
}

// ApplicationControlRules is an Sophos Endpoint subType and implements sophos.RestObject
type ApplicationControlRules []ApplicationControlRule

// ApplicationControlRule is a generated Sophos object
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

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*ApplicationControlRule) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/application_control/rule/%s/usedby", ref)
}

// GetType implements sophos.Object
func (a *ApplicationControlRule) GetType() string { return a._type }
