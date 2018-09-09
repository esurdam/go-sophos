package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// RemoteSyslog is a generated struct representing the Sophos RemoteSyslog Endpoint
// GET /api/nodes/remote_syslog
type RemoteSyslog struct {
	Buffer int64         `json:"buffer"`
	Logs   []interface{} `json:"logs"`
	Status int64         `json:"status"`
	Target []interface{} `json:"target"`
}

var _ sophos.Endpoint = &RemoteSyslog{}

var defsRemoteSyslog = map[string]sophos.RestObject{
	"RemoteSyslogGroup":  &RemoteSyslogGroup{},
	"RemoteSyslogServer": &RemoteSyslogServer{},
}

// RestObjects implements the sophos.Node interface and returns a map of RemoteSyslog's Objects
func (RemoteSyslog) RestObjects() map[string]sophos.RestObject { return defsRemoteSyslog }

// GetPath implements sophos.RestGetter
func (*RemoteSyslog) GetPath() string { return "/api/nodes/remote_syslog" }

// RefRequired implements sophos.RestGetter
func (*RemoteSyslog) RefRequired() (string, bool) { return "", false }

var defRemoteSyslog = &sophos.Definition{Description: "remote_syslog", Name: "remote_syslog", Link: "/api/definitions/remote_syslog"}

// Definition returns the /api/definitions struct of RemoteSyslog
func (RemoteSyslog) Definition() sophos.Definition { return *defRemoteSyslog }

// ApiRoutes returns all known RemoteSyslog Paths
func (RemoteSyslog) ApiRoutes() []string {
	return []string{
		"/api/objects/remote_syslog/group/",
		"/api/objects/remote_syslog/group/{ref}",
		"/api/objects/remote_syslog/group/{ref}/usedby",
		"/api/objects/remote_syslog/server/",
		"/api/objects/remote_syslog/server/{ref}",
		"/api/objects/remote_syslog/server/{ref}/usedby",
	}
}

// References returns the RemoteSyslog's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (RemoteSyslog) References() []string {
	return []string{
		"REF_RemoteSyslogGroup",
		"REF_RemoteSyslogServer",
	}
}

// RemoteSyslogGroup is an Sophos Endpoint subType and implements sophos.RestObject
type RemoteSyslogGroup []interface{}

var _ sophos.RestObject = &RemoteSyslogGroup{}

// GetPath implements sophos.RestObject and returns the RemoteSyslogGroup GET path
// Returns all available remote_syslog/group objects
func (*RemoteSyslogGroup) GetPath() string { return "/api/objects/remote_syslog/group/" }

// RefRequired implements sophos.RestObject
func (*RemoteSyslogGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the RemoteSyslogGroup DELETE path
// Creates or updates the complete object group
func (*RemoteSyslogGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the RemoteSyslogGroup PATCH path
// Changes to parts of the object group types
func (*RemoteSyslogGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the RemoteSyslogGroup POST path
// Create a new remote_syslog/group object
func (*RemoteSyslogGroup) PostPath() string {
	return "/api/objects/remote_syslog/group/"
}

// PutPath implements sophos.RestObject and returns the RemoteSyslogGroup PUT path
// Creates or updates the complete object group
func (*RemoteSyslogGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*RemoteSyslogGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/group/%s/usedby", ref)
}

// RemoteSyslogServer is an Sophos Endpoint subType and implements sophos.RestObject
type RemoteSyslogServer []interface{}

var _ sophos.RestObject = &RemoteSyslogServer{}

// GetPath implements sophos.RestObject and returns the RemoteSyslogServer GET path
// Returns all available remote_syslog/server objects
func (*RemoteSyslogServer) GetPath() string { return "/api/objects/remote_syslog/server/" }

// RefRequired implements sophos.RestObject
func (*RemoteSyslogServer) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the RemoteSyslogServer DELETE path
// Creates or updates the complete object server
func (*RemoteSyslogServer) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/server/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the RemoteSyslogServer PATCH path
// Changes to parts of the object server types
func (*RemoteSyslogServer) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/server/%s", ref)
}

// PostPath implements sophos.RestObject and returns the RemoteSyslogServer POST path
// Create a new remote_syslog/server object
func (*RemoteSyslogServer) PostPath() string {
	return "/api/objects/remote_syslog/server/"
}

// PutPath implements sophos.RestObject and returns the RemoteSyslogServer PUT path
// Creates or updates the complete object server
func (*RemoteSyslogServer) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/server/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*RemoteSyslogServer) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/remote_syslog/server/%s/usedby", ref)
}
