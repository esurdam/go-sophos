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

// AuthenticationAdirectorys is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationAdirectorys []AuthenticationAdirectory

// AuthenticationAdirectory represents a UTM Microsoft Active Directory server
type AuthenticationAdirectory struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Ssl default value is false
	Ssl     bool `json:"ssl"`
	Timeout int  `json:"timeout"`
	// BaseDn default value is ""
	BaseDn string `json:"base_dn"`
	// BindPw default value is ""
	BindPw string `json:"bind_pw"`
	// Server description: REF(network/host), REF(network/dns_host), REF(network/availability_group)
	Server string `json:"server"`
	// PrefetchBackendSync default value is false
	PrefetchBackendSync bool `json:"prefetch_backend_sync"`
	// BindDn default value is ""
	BindDn  string `json:"bind_dn"`
	Comment string `json:"comment"`
	Port    int    `json:"port"`
	// Status default value is false
	Status bool `json:"status"`
	// Backend default value is ""
	Backend          string        `json:"backend"`
	Name             string        `json:"name"`
	PrefetchInterval []interface{} `json:"prefetch_interval"`
	PrefetchContexts []interface{} `json:"prefetch_contexts"`
	// Sasl default value is false
	Sasl bool `json:"sasl"`
}

var _ sophos.RestGetter = &AuthenticationAdirectory{}

// GetPath implements sophos.RestObject and returns the AuthenticationAdirectorys GET path
// Returns all available authentication/adirectory objects
func (*AuthenticationAdirectorys) GetPath() string { return "/api/objects/authentication/adirectory/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationAdirectorys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AuthenticationAdirectorys GET path
// Returns all available adirectory types
func (a *AuthenticationAdirectory) GetPath() string {
	return fmt.Sprintf("/api/objects/authentication/adirectory/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AuthenticationAdirectory) RefRequired() (string, bool) { return a.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationAdirectory) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/adirectory/%s/usedby", ref)
}

// AuthenticationEdirectorys is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationEdirectorys []AuthenticationEdirectory

// AuthenticationEdirectory represents a UTM Novell eDirectory server
type AuthenticationEdirectory struct {
	Locked           string        `json:"_locked"`
	ObjectType       string        `json:"_type"`
	Reference        string        `json:"_ref"`
	PrefetchInterval []interface{} `json:"prefetch_interval"`
	// Ssl default value is true
	Ssl  bool   `json:"ssl"`
	Name string `json:"name"`
	// BindDn default value is ""
	BindDn           string        `json:"bind_dn"`
	Port             int           `json:"port"`
	PrefetchContexts []interface{} `json:"prefetch_contexts"`
	Timeout          int           `json:"timeout"`
	// Backend default value is ""
	Backend string `json:"backend"`
	// PrefetchBackendSync default value is false
	PrefetchBackendSync bool `json:"prefetch_backend_sync"`
	// Server description: REF(network/host), REF(network/dns_host), REF(network/availability_group)
	Server string `json:"server"`
	// BindPw default value is ""
	BindPw   string        `json:"bind_pw"`
	Contexts []interface{} `json:"contexts"`
	// Sasl default value is false
	Sasl bool `json:"sasl"`
	// Status default value is false
	Status  bool   `json:"status"`
	Comment string `json:"comment"`
}

var _ sophos.RestGetter = &AuthenticationEdirectory{}

// GetPath implements sophos.RestObject and returns the AuthenticationEdirectorys GET path
// Returns all available authentication/edirectory objects
func (*AuthenticationEdirectorys) GetPath() string { return "/api/objects/authentication/edirectory/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationEdirectorys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AuthenticationEdirectorys GET path
// Returns all available edirectory types
func (a *AuthenticationEdirectory) GetPath() string {
	return fmt.Sprintf("/api/objects/authentication/edirectory/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AuthenticationEdirectory) RefRequired() (string, bool) { return a.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationEdirectory) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/edirectory/%s/usedby", ref)
}

// AuthenticationGroups is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationGroups []AuthenticationGroup

// AuthenticationGroup represents a UTM group
type AuthenticationGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &AuthenticationGroup{}

// GetPath implements sophos.RestObject and returns the AuthenticationGroups GET path
// Returns all available authentication/group objects
func (*AuthenticationGroups) GetPath() string { return "/api/objects/authentication/group/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AuthenticationGroups GET path
// Returns all available group types
func (a *AuthenticationGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/authentication/group/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AuthenticationGroup) RefRequired() (string, bool) { return a.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/group/%s/usedby", ref)
}

// AuthenticationLdaps is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationLdaps []AuthenticationLdap

// AuthenticationLdap represents a UTM LDAP server
type AuthenticationLdap struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Sasl default value is false
	Sasl bool `json:"sasl"`
	// Ssl default value is false
	Ssl              bool          `json:"ssl"`
	PrefetchContexts []interface{} `json:"prefetch_contexts"`
	Port             int           `json:"port"`
	// PrefetchBackendSync default value is false
	PrefetchBackendSync bool `json:"prefetch_backend_sync"`
	// Server description: REF(network/host), REF(network/dns_host), REF(network/availability_group)
	Server string `json:"server"`
	// Status default value is false
	Status bool `json:"status"`
	// UserAttrib can be one of: []string{"cn", "sn", "uid", "custom"}
	// UserAttrib default value is "cn"
	UserAttrib string `json:"user_attrib"`
	// UserAttribCustom default value is ""
	UserAttribCustom string `json:"user_attrib_custom"`
	// BindDn default value is ""
	BindDn           string        `json:"bind_dn"`
	Comment          string        `json:"comment"`
	Name             string        `json:"name"`
	PrefetchInterval []interface{} `json:"prefetch_interval"`
	Timeout          int           `json:"timeout"`
	// Backend default value is ""
	Backend string `json:"backend"`
	// BindPw default value is ""
	BindPw string `json:"bind_pw"`
	// BaseDn default value is ""
	BaseDn string `json:"base_dn"`
}

var _ sophos.RestGetter = &AuthenticationLdap{}

// GetPath implements sophos.RestObject and returns the AuthenticationLdaps GET path
// Returns all available authentication/ldap objects
func (*AuthenticationLdaps) GetPath() string { return "/api/objects/authentication/ldap/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationLdaps) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AuthenticationLdaps GET path
// Returns all available ldap types
func (a *AuthenticationLdap) GetPath() string {
	return fmt.Sprintf("/api/objects/authentication/ldap/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AuthenticationLdap) RefRequired() (string, bool) { return a.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationLdap) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/ldap/%s/usedby", ref)
}

// AuthenticationOtpTokens is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationOtpTokens []AuthenticationOtpToken

// AuthenticationOtpToken represents a UTM One Time Password token
type AuthenticationOtpToken struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Status default value is false
	Status bool `json:"status"`
	// Timestep description: Constraints: 0, 10-120
	Timestep   int           `json:"timestep"`
	ExtraCodes []interface{} `json:"extra_codes"`
	// ForSsh default value is false
	ForSsh  bool   `json:"for_ssh"`
	Lastuse int    `json:"lastuse"`
	Name    string `json:"name"`
	Offset  int    `json:"offset"`
	Secret  string `json:"secret"`
	// User description: REF(aaa/user)
	// User default value is ""
	User    string `json:"user"`
	Comment string `json:"comment"`
	// Digest can be one of: []string{"sha1", "sha256", "sha512"}
	// Digest default value is "sha1"
	Digest string `json:"digest"`
	// Hide default value is false
	Hide bool `json:"hide"`
}

var _ sophos.RestGetter = &AuthenticationOtpToken{}

// GetPath implements sophos.RestObject and returns the AuthenticationOtpTokens GET path
// Returns all available authentication/otp_token objects
func (*AuthenticationOtpTokens) GetPath() string { return "/api/objects/authentication/otp_token/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationOtpTokens) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AuthenticationOtpTokens GET path
// Returns all available otp_token types
func (a *AuthenticationOtpToken) GetPath() string {
	return fmt.Sprintf("/api/objects/authentication/otp_token/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AuthenticationOtpToken) RefRequired() (string, bool) { return a.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationOtpToken) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/otp_token/%s/usedby", ref)
}

// AuthenticationRadiuss is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationRadiuss []AuthenticationRadius

// AuthenticationRadius represents a UTM RADIUS server
type AuthenticationRadius struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Secret default value is ""
	Secret string `json:"secret"`
	// Server description: REF(network/host), REF(network/dns_host), REF(network/availability_group)
	Server string `json:"server"`
	// Status default value is false
	Status  bool `json:"status"`
	Timeout int  `json:"timeout"`
	// Backend default value is ""
	Backend string `json:"backend"`
	Comment string `json:"comment"`
	Name    string `json:"name"`
	Port    int    `json:"port"`
}

var _ sophos.RestGetter = &AuthenticationRadius{}

// GetPath implements sophos.RestObject and returns the AuthenticationRadiuss GET path
// Returns all available authentication/radius objects
func (*AuthenticationRadiuss) GetPath() string { return "/api/objects/authentication/radius/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationRadiuss) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AuthenticationRadiuss GET path
// Returns all available radius types
func (a *AuthenticationRadius) GetPath() string {
	return fmt.Sprintf("/api/objects/authentication/radius/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AuthenticationRadius) RefRequired() (string, bool) { return a.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationRadius) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/radius/%s/usedby", ref)
}

// AuthenticationTacacss is an Sophos Endpoint subType and implements sophos.RestObject
type AuthenticationTacacss []AuthenticationTacacs

// AuthenticationTacacs represents a UTM TACACS+ server
type AuthenticationTacacs struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Backend default value is ""
	Backend string `json:"backend"`
	Comment string `json:"comment"`
	// Key default value is ""
	Key  string `json:"key"`
	Name string `json:"name"`
	Port int    `json:"port"`
	// Server description: REF(network/host), REF(network/dns_host), REF(network/availability_group)
	Server string `json:"server"`
	// Status default value is false
	Status  bool `json:"status"`
	Timeout int  `json:"timeout"`
}

var _ sophos.RestGetter = &AuthenticationTacacs{}

// GetPath implements sophos.RestObject and returns the AuthenticationTacacss GET path
// Returns all available authentication/tacacs objects
func (*AuthenticationTacacss) GetPath() string { return "/api/objects/authentication/tacacs/" }

// RefRequired implements sophos.RestObject
func (*AuthenticationTacacss) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AuthenticationTacacss GET path
// Returns all available tacacs types
func (a *AuthenticationTacacs) GetPath() string {
	return fmt.Sprintf("/api/objects/authentication/tacacs/%s", a.Reference)
}

// RefRequired implements sophos.RestObject
func (a *AuthenticationTacacs) RefRequired() (string, bool) { return a.Reference, true }

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

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AuthenticationTacacs) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/authentication/tacacs/%s/usedby", ref)
}
