package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Route is a generated struct representing the Sophos Route Endpoint
// GET /api/nodes/route
type Route struct {
	RouteGroup  RouteGroup  `json:"route_group"`
	RoutePolicy RoutePolicy `json:"route_policy"`
	RouteStatic RouteStatic `json:"route_static"`
}

var _ sophos.Endpoint = &Route{}

var defsRoute = map[string]sophos.RestObject{
	"RouteGroup":  &RouteGroup{},
	"RoutePolicy": &RoutePolicy{},
	"RouteStatic": &RouteStatic{},
}

// RestObjects implements the sophos.Node interface and returns a map of Route's Objects
func (Route) RestObjects() map[string]sophos.RestObject { return defsRoute }

// GetPath implements sophos.RestGetter
func (*Route) GetPath() string { return "/api/nodes/route" }

// RefRequired implements sophos.RestGetter
func (*Route) RefRequired() (string, bool) { return "", false }

var defRoute = &sophos.Definition{Description: "route", Name: "route", Link: "/api/definitions/route"}

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

// RouteGroup is an Sophos Endpoint subType and implements sophos.RestObject
type RouteGroup []interface{}

var _ sophos.RestObject = &RouteGroup{}

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

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*RouteGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/route/group/%s/usedby", ref)
}

// RoutePolicy is an Sophos Endpoint subType and implements sophos.RestObject
type RoutePolicy []interface{}

var _ sophos.RestObject = &RoutePolicy{}

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

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*RoutePolicy) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/route/policy/%s/usedby", ref)
}

// RouteStatics is an Sophos Endpoint subType and implements sophos.RestObject
type RouteStatics []RouteStatic

// RouteStatic is a generated Sophos object
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

var _ sophos.RestGetter = &RouteStatic{}

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

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*RouteStatic) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/route/static/%s/usedby", ref)
}

// GetType implements sophos.Object
func (r *RouteStatic) GetType() string { return r._type }
