package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Awscli is a generated struct representing the Sophos Awscli Endpoint
// GET /api/nodes/awscli
type Awscli struct {
	Profiles []interface{} `json:"profiles"`
}

var _ sophos.Endpoint = &Awscli{}

var defsAwscli = map[string]sophos.RestObject{
	"AwscliGroup":   &AwscliGroup{},
	"AwscliProfile": &AwscliProfile{},
}

// RestObjects implements the sophos.Node interface and returns a map of Awscli's Objects
func (Awscli) RestObjects() map[string]sophos.RestObject { return defsAwscli }

// GetPath implements sophos.RestGetter
func (*Awscli) GetPath() string { return "/api/nodes/awscli" }

// RefRequired implements sophos.RestGetter
func (*Awscli) RefRequired() (string, bool) { return "", false }

var defAwscli = &sophos.Definition{Description: "awscli", Name: "awscli", Link: "/api/definitions/awscli"}

// Definition returns the /api/definitions struct of Awscli
func (Awscli) Definition() sophos.Definition { return *defAwscli }

// ApiRoutes returns all known Awscli Paths
func (Awscli) ApiRoutes() []string {
	return []string{
		"/api/objects/awscli/group/",
		"/api/objects/awscli/group/{ref}",
		"/api/objects/awscli/group/{ref}/usedby",
		"/api/objects/awscli/profile/",
		"/api/objects/awscli/profile/{ref}",
		"/api/objects/awscli/profile/{ref}/usedby",
	}
}

// References returns the Awscli's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Awscli) References() []string {
	return []string{
		"REF_AwscliGroup",
		"REF_AwscliProfile",
	}
}

// AwscliGroups is an Sophos Endpoint subType and implements sophos.RestObject
type AwscliGroups []AwscliGroup

// AwscliGroup represents a UTM awscli->group
type AwscliGroup struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
}

var _ sophos.RestGetter = &AwscliGroup{}

// GetPath implements sophos.RestObject and returns the AwscliGroups GET path
// Returns all available awscli/group objects
func (*AwscliGroups) GetPath() string { return "/api/objects/awscli/group/" }

// RefRequired implements sophos.RestObject
func (*AwscliGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AwscliGroups GET path
// Returns all available group types
func (a *AwscliGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/awscli/group/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AwscliGroup) RefRequired() (string, bool) { return a.Reference, true }

// DeletePath implements sophos.RestObject and returns the AwscliGroup DELETE path
// Creates or updates the complete object group
func (*AwscliGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/awscli/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AwscliGroup PATCH path
// Changes to parts of the object group types
func (*AwscliGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/awscli/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AwscliGroup POST path
// Create a new awscli/group object
func (*AwscliGroup) PostPath() string {
	return "/api/objects/awscli/group/"
}

// PutPath implements sophos.RestObject and returns the AwscliGroup PUT path
// Creates or updates the complete object group
func (*AwscliGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/awscli/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AwscliGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awscli/group/%s/usedby", ref)
}

// AwscliProfiles is an Sophos Endpoint subType and implements sophos.RestObject
type AwscliProfiles []AwscliProfile

// AwscliProfile represents a UTM AWS CLI Profile
type AwscliProfile struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	// AwsSessionToken default value is ""
	AwsSessionToken string `json:"aws_session_token"`
	Comment         string `json:"comment"`
	Name            string `json:"name"`
	// Output can be one of: []string{"json", "text", "table"}
	// Output default value is "json"
	Output string `json:"output"`
	// ProfileName description: (REGEX)
	// ProfileName default value is "default"
	ProfileName string `json:"profile_name"`
	// Region description: REF(aws/region)
	Region string `json:"region"`
	// AwsAccessKeyId description: (REGEX)
	// AwsAccessKeyId default value is ""
	AwsAccessKeyId string `json:"aws_access_key_id"`
	// AwsSecretAccessKey default value is ""
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
}

var _ sophos.RestGetter = &AwscliProfile{}

// GetPath implements sophos.RestObject and returns the AwscliProfiles GET path
// Returns all available awscli/profile objects
func (*AwscliProfiles) GetPath() string { return "/api/objects/awscli/profile/" }

// RefRequired implements sophos.RestObject
func (*AwscliProfiles) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AwscliProfiles GET path
// Returns all available profile types
func (a *AwscliProfile) GetPath() string {
	return fmt.Sprintf("/api/objects/awscli/profile/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AwscliProfile) RefRequired() (string, bool) { return a.Reference, true }

// DeletePath implements sophos.RestObject and returns the AwscliProfile DELETE path
// Creates or updates the complete object profile
func (*AwscliProfile) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/awscli/profile/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AwscliProfile PATCH path
// Changes to parts of the object profile types
func (*AwscliProfile) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/awscli/profile/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AwscliProfile POST path
// Create a new awscli/profile object
func (*AwscliProfile) PostPath() string {
	return "/api/objects/awscli/profile/"
}

// PutPath implements sophos.RestObject and returns the AwscliProfile PUT path
// Creates or updates the complete object profile
func (*AwscliProfile) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/awscli/profile/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AwscliProfile) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/awscli/profile/%s/usedby", ref)
}
