package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Authentication is a generated struct representing the Sophos Authentication Endpoint
// GET /api/nodes/authentication
type Authentication struct {
	AuthenticationAdirectory AuthenticationAdirectory `json:"authentication_adirectory"`
	AuthenticationEdirectory AuthenticationEdirectory `json:"authentication_edirectory"`
	AuthenticationGroup      AuthenticationGroup      `json:"authentication_group"`
	AuthenticationLdap       AuthenticationLdap       `json:"authentication_ldap"`
	AuthenticationOtpToken   AuthenticationOtpToken   `json:"authentication_otp_token"`
	AuthenticationRadius     AuthenticationRadius     `json:"authentication_radius"`
	AuthenticationTacacs     AuthenticationTacacs     `json:"authentication_tacacs"`
}

var _ sophos.Endpoint = &Authentication{}

var defsAuthentication = map[string]sophos.RestObject{
	"AuthenticationAdirectory": &AuthenticationAdirectory{},
	"AuthenticationEdirectory": &AuthenticationEdirectory{},
	"AuthenticationGroup":      &AuthenticationGroup{},
	"AuthenticationLdap":       &AuthenticationLdap{},
	"AuthenticationOtpToken":   &AuthenticationOtpToken{},
	"AuthenticationRadius":     &AuthenticationRadius{},
	"AuthenticationTacacs":     &AuthenticationTacacs{},
}

// RestObjects implements the sophos.Node interface and returns a map of Authentication's Objects
func (Authentication) RestObjects() map[string]sophos.RestObject { return defsAuthentication }

// GetPath implements sophos.RestGetter
func (*Authentication) GetPath() string { return "/api/nodes/authentication" }

// RefRequired implements sophos.RestGetter
func (*Authentication) RefRequired() (string, bool) { return "", false }

var defAuthentication = &sophos.Definition{Description: "authentication", Name: "authentication", Link: "/api/definitions/authentication"}

// Definition returns the /api/definitions struct of Authentication
func (Authentication) Definition() sophos.Definition { return *defAuthentication }

// ApiRoutes returns all known Authentication Paths
func (Authentication) ApiRoutes() []string {
	return []string{
		"/api/objects/authentication/adirectory/",
		"/api/objects/authentication/adirectory/{ref}",
		"/api/objects/authentication/adirectory/{ref}/usedby",
		"/api/objects/authentication/edirectory/",
		"/api/objects/authentication/edirectory/{ref}",
		"/api/objects/authentication/edirectory/{ref}/usedby",
		"/api/objects/authentication/group/",
		"/api/objects/authentication/group/{ref}",
		"/api/objects/authentication/group/{ref}/usedby",
		"/api/objects/authentication/ldap/",
		"/api/objects/authentication/ldap/{ref}",
		"/api/objects/authentication/ldap/{ref}/usedby",
		"/api/objects/authentication/otp_token/",
		"/api/objects/authentication/otp_token/{ref}",
		"/api/objects/authentication/otp_token/{ref}/usedby",
		"/api/objects/authentication/radius/",
		"/api/objects/authentication/radius/{ref}",
		"/api/objects/authentication/radius/{ref}/usedby",
		"/api/objects/authentication/tacacs/",
		"/api/objects/authentication/tacacs/{ref}",
		"/api/objects/authentication/tacacs/{ref}/usedby",
	}
}

// References returns the Authentication's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Authentication) References() []string {
	return []string{
		"REF_AuthenticationAdirectory",
		"REF_AuthenticationEdirectory",
		"REF_AuthenticationGroup",
		"REF_AuthenticationLdap",
		"REF_AuthenticationOtpToken",
		"REF_AuthenticationRadius",
		"REF_AuthenticationTacacs",
	}
}

// AuthenticationAdirectory is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationAdirectory []interface{}

var _ sophos.RestObject = &AuthenticationAdirectory{}

// GetPath implements sophos.RestObject and returns the AuthenticationAdirectory GET path
// Returns all available authentication/adirectory objects
func (*AuthenticationAdirectory) GetPath() string { return "/api/objects/authentication/adirectory/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationAdirectory) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AuthenticationAdirectory DELETE path
// Creates or updates the complete object adirectory
func (*AuthenticationAdirectory) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/adirectory/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AuthenticationAdirectory PATCH path
// Changes to parts of the object adirectory types
func (*AuthenticationAdirectory) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/adirectory/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AuthenticationAdirectory POST path
// Create a new authentication/adirectory object
func (*AuthenticationAdirectory) PostPath() string {
	return "/api/objects/authentication/adirectory/"
}

// PutPath implements sophos.RestObject and returns the AuthenticationAdirectory PUT path
// Creates or updates the complete object adirectory
func (*AuthenticationAdirectory) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/adirectory/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationAdirectory) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/adirectory/%s/usedby", ref)
}

// AuthenticationEdirectory is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationEdirectory []interface{}

var _ sophos.RestObject = &AuthenticationEdirectory{}

// GetPath implements sophos.RestObject and returns the AuthenticationEdirectory GET path
// Returns all available authentication/edirectory objects
func (*AuthenticationEdirectory) GetPath() string { return "/api/objects/authentication/edirectory/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationEdirectory) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AuthenticationEdirectory DELETE path
// Creates or updates the complete object edirectory
func (*AuthenticationEdirectory) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/edirectory/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AuthenticationEdirectory PATCH path
// Changes to parts of the object edirectory types
func (*AuthenticationEdirectory) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/edirectory/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AuthenticationEdirectory POST path
// Create a new authentication/edirectory object
func (*AuthenticationEdirectory) PostPath() string {
	return "/api/objects/authentication/edirectory/"
}

// PutPath implements sophos.RestObject and returns the AuthenticationEdirectory PUT path
// Creates or updates the complete object edirectory
func (*AuthenticationEdirectory) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/edirectory/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationEdirectory) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/edirectory/%s/usedby", ref)
}

// AuthenticationGroup is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationGroup []interface{}

var _ sophos.RestObject = &AuthenticationGroup{}

// GetPath implements sophos.RestObject and returns the AuthenticationGroup GET path
// Returns all available authentication/group objects
func (*AuthenticationGroup) GetPath() string { return "/api/objects/authentication/group/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AuthenticationGroup DELETE path
// Creates or updates the complete object group
func (*AuthenticationGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AuthenticationGroup PATCH path
// Changes to parts of the object group types
func (*AuthenticationGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AuthenticationGroup POST path
// Create a new authentication/group object
func (*AuthenticationGroup) PostPath() string {
	return "/api/objects/authentication/group/"
}

// PutPath implements sophos.RestObject and returns the AuthenticationGroup PUT path
// Creates or updates the complete object group
func (*AuthenticationGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/group/%s/usedby", ref)
}

// AuthenticationLdap is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationLdap []interface{}

var _ sophos.RestObject = &AuthenticationLdap{}

// GetPath implements sophos.RestObject and returns the AuthenticationLdap GET path
// Returns all available authentication/ldap objects
func (*AuthenticationLdap) GetPath() string { return "/api/objects/authentication/ldap/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationLdap) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AuthenticationLdap DELETE path
// Creates or updates the complete object ldap
func (*AuthenticationLdap) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/ldap/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AuthenticationLdap PATCH path
// Changes to parts of the object ldap types
func (*AuthenticationLdap) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/ldap/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AuthenticationLdap POST path
// Create a new authentication/ldap object
func (*AuthenticationLdap) PostPath() string {
	return "/api/objects/authentication/ldap/"
}

// PutPath implements sophos.RestObject and returns the AuthenticationLdap PUT path
// Creates or updates the complete object ldap
func (*AuthenticationLdap) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/ldap/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationLdap) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/ldap/%s/usedby", ref)
}

// AuthenticationOtpToken is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationOtpToken []interface{}

var _ sophos.RestObject = &AuthenticationOtpToken{}

// GetPath implements sophos.RestObject and returns the AuthenticationOtpToken GET path
// Returns all available authentication/otp_token objects
func (*AuthenticationOtpToken) GetPath() string { return "/api/objects/authentication/otp_token/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationOtpToken) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AuthenticationOtpToken DELETE path
// Creates or updates the complete object otp_token
func (*AuthenticationOtpToken) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/otp_token/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AuthenticationOtpToken PATCH path
// Changes to parts of the object otp_token types
func (*AuthenticationOtpToken) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/otp_token/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AuthenticationOtpToken POST path
// Create a new authentication/otp_token object
func (*AuthenticationOtpToken) PostPath() string {
	return "/api/objects/authentication/otp_token/"
}

// PutPath implements sophos.RestObject and returns the AuthenticationOtpToken PUT path
// Creates or updates the complete object otp_token
func (*AuthenticationOtpToken) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/otp_token/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationOtpToken) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/otp_token/%s/usedby", ref)
}

// AuthenticationRadius is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationRadius []interface{}

var _ sophos.RestObject = &AuthenticationRadius{}

// GetPath implements sophos.RestObject and returns the AuthenticationRadius GET path
// Returns all available authentication/radius objects
func (*AuthenticationRadius) GetPath() string { return "/api/objects/authentication/radius/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationRadius) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AuthenticationRadius DELETE path
// Creates or updates the complete object radius
func (*AuthenticationRadius) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/radius/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AuthenticationRadius PATCH path
// Changes to parts of the object radius types
func (*AuthenticationRadius) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/radius/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AuthenticationRadius POST path
// Create a new authentication/radius object
func (*AuthenticationRadius) PostPath() string {
	return "/api/objects/authentication/radius/"
}

// PutPath implements sophos.RestObject and returns the AuthenticationRadius PUT path
// Creates or updates the complete object radius
func (*AuthenticationRadius) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/radius/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationRadius) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/radius/%s/usedby", ref)
}

// AuthenticationTacacs is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationTacacs []interface{}

var _ sophos.RestObject = &AuthenticationTacacs{}

// GetPath implements sophos.RestObject and returns the AuthenticationTacacs GET path
// Returns all available authentication/tacacs objects
func (*AuthenticationTacacs) GetPath() string { return "/api/objects/authentication/tacacs/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationTacacs) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the AuthenticationTacacs DELETE path
// Creates or updates the complete object tacacs
func (*AuthenticationTacacs) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/tacacs/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AuthenticationTacacs PATCH path
// Changes to parts of the object tacacs types
func (*AuthenticationTacacs) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/tacacs/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AuthenticationTacacs POST path
// Create a new authentication/tacacs object
func (*AuthenticationTacacs) PostPath() string {
	return "/api/objects/authentication/tacacs/"
}

// PutPath implements sophos.RestObject and returns the AuthenticationTacacs PUT path
// Creates or updates the complete object tacacs
func (*AuthenticationTacacs) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/tacacs/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationTacacs) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/tacacs/%s/usedby", ref)
}
