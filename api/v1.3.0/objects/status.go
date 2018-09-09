package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Status is a generated struct representing the Sophos Status Endpoint
// GET /api/nodes/status
type Status struct {
	StatusVersion StatusVersion `json:"Status_status"`
}

var _ sophos.Endpoint = &Status{}

var defsStatus = map[string]sophos.RestObject{
	"StatusVersion": &StatusVersion{},
}

// RestObjects implements the sophos.Node interface and returns a map of Status's Objects
func (Status) RestObjects() map[string]sophos.RestObject { return defsStatus }

// GetPath implements sophos.RestGetter
func (*Status) GetPath() string { return "/api/nodes/status" }

// RefRequired implements sophos.RestGetter
func (*Status) RefRequired() (string, bool) { return "", false }

var defStatus = &sophos.Definition{Description: "Status", Name: "Status", Link: "/api/definitions/status"}

// Definition returns the /api/definitions struct of Status
func (Status) Definition() sophos.Definition { return *defStatus }

// ApiRoutes returns all known Status Paths
func (Status) ApiRoutes() []string {
	return []string{
		"/api/status/version",
	}
}

// References returns the Status's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Status) References() []string {
	return []string{}
}

// StatusVersion is an Sophos Endpoint subType and implements sophos.RestObject
type StatusVersion struct {
	Restd string `json:"restd"`
	Utm   string `json:"utm"`
}

var _ sophos.RestObject = &StatusVersion{}

// GetPath implements sophos.RestObject and returns the StatusVersion GET path
// Returns some version numbers of the UTM software
func (*StatusVersion) GetPath() string { return "/api/status/version" }

// RefRequired implements sophos.RestObject
func (*StatusVersion) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the StatusVersion DELETE path
func (*StatusVersion) DeletePath(ref string) string {
	return fmt.Sprintf("/api", ref)
}

// PatchPath implements sophos.RestObject and returns the StatusVersion PATCH path
func (*StatusVersion) PatchPath(ref string) string {
	return fmt.Sprintf("/api", ref)
}

// PostPath implements sophos.RestObject and returns the StatusVersion POST path
func (*StatusVersion) PostPath() string {
	return "/api"
}

// PutPath implements sophos.RestObject and returns the StatusVersion PUT path
func (*StatusVersion) PutPath(ref string) string {
	return fmt.Sprintf("/api", ref)
}

// UsedByPath implements sophos.UsedObject
func (*StatusVersion) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/usedby", ref)
}
