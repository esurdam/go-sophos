package objects

// THIS FILE IS AUTO GENERATED BY bin/gen.go! DO NOT EDIT!

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Role is a generated struct representing the Sophos Role Endpoint
// GET /api/nodes/role
type Role struct {
	RoleGroup RoleGroup `json:"role_group"`
	RoleRole  RoleRole  `json:"role_role"`
}

var _ sophos.Endpoint = &Role{}

var defsRole = map[string]sophos.RestObject{
	"RoleGroup": &RoleGroup{},
	"RoleRole":  &RoleRole{},
}

// RestObjects implements the sophos.Node interface and returns a map of Role's Objects
func (Role) RestObjects() map[string]sophos.RestObject { return defsRole }

// GetPath implements sophos.RestGetter
func (*Role) GetPath() string { return "/api/nodes/role" }

// RefRequired implements sophos.RestGetter
func (*Role) RefRequired() (string, bool) { return "", false }

var defRole = &sophos.Definition{Description: "role", Name: "role", Link: "/api/definitions/role"}

// Definition returns the /api/definitions struct of Role
func (Role) Definition() sophos.Definition { return *defRole }

// ApiRoutes returns all known Role Paths
func (Role) ApiRoutes() []string {
	return []string{
		"/api/objects/role/group/",
		"/api/objects/role/group/{ref}",
		"/api/objects/role/group/{ref}/usedby",
		"/api/objects/role/role/",
		"/api/objects/role/role/{ref}",
		"/api/objects/role/role/{ref}/usedby",
	}
}

// References returns the Role's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Role) References() []string {
	return []string{
		"REF_RoleGroup",
		"REF_RoleRole",
	}
}

// RoleGroups is an Sophos Endpoint subType and implements sophos.RestObject
type RoleGroups []RoleGroup

// RoleGroup represents a UTM group
type RoleGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &RoleGroup{}

// GetPath implements sophos.RestObject and returns the RoleGroups GET path
// Returns all available role/group objects
func (*RoleGroups) GetPath() string { return "/api/objects/role/group/" }

// RefRequired implements sophos.RestObject
func (*RoleGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the RoleGroups GET path
// Returns all available group types
func (r *RoleGroup) GetPath() string { return fmt.Sprintf("/api/objects/role/group/%s", r.Reference) }

// RefRequired implements sophos.RestObject
func (r *RoleGroup) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the RoleGroup DELETE path
// Creates or updates the complete object group
func (*RoleGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/role/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the RoleGroup PATCH path
// Changes to parts of the object group types
func (*RoleGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/role/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the RoleGroup POST path
// Create a new role/group object
func (*RoleGroup) PostPath() string {
	return "/api/objects/role/group/"
}

// PutPath implements sophos.RestObject and returns the RoleGroup PUT path
// Creates or updates the complete object group
func (*RoleGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/role/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*RoleGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/role/group/%s/usedby", ref)
}

// RoleRoles is an Sophos Endpoint subType and implements sophos.RestObject
type RoleRoles []RoleRole

// RoleRole represents a UTM role
type RoleRole struct {
	Locked     string   `json:"_locked"`
	ObjectType string   `json:"_type"`
	Reference  string   `json:"_ref"`
	Comment    string   `json:"comment"`
	Members    []string `json:"members"`
	Name       string   `json:"name"`
	Rights     []string `json:"rights"`
	// WebadminAccess default value is false
	WebadminAccess bool `json:"webadmin_access"`
}

var _ sophos.RestGetter = &RoleRole{}

// GetPath implements sophos.RestObject and returns the RoleRoles GET path
// Returns all available role/role objects
func (*RoleRoles) GetPath() string { return "/api/objects/role/role/" }

// RefRequired implements sophos.RestObject
func (*RoleRoles) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the RoleRoles GET path
// Returns all available role types
func (r *RoleRole) GetPath() string { return fmt.Sprintf("/api/objects/role/role/%s", r.Reference) }

// RefRequired implements sophos.RestObject
func (r *RoleRole) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the RoleRole DELETE path
// Creates or updates the complete object role
func (*RoleRole) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/role/role/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the RoleRole PATCH path
// Changes to parts of the object role types
func (*RoleRole) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/role/role/%s", ref)
}

// PostPath implements sophos.RestObject and returns the RoleRole POST path
// Create a new role/role object
func (*RoleRole) PostPath() string {
	return "/api/objects/role/role/"
}

// PutPath implements sophos.RestObject and returns the RoleRole PUT path
// Creates or updates the complete object role
func (*RoleRole) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/role/role/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*RoleRole) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/role/role/%s/usedby", ref)
}
