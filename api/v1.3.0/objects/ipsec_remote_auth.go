package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// IpsecRemoteAuth is a generated struct representing the Sophos IpsecRemoteAuth Endpoint
// GET /api/nodes/ipsec_remote_auth
type IpsecRemoteAuth struct {
	IpsecRemoteAuthCa    IpsecRemoteAuthCa    `json:"ipsec_remote_auth_ca"`
	IpsecRemoteAuthGroup IpsecRemoteAuthGroup `json:"ipsec_remote_auth_group"`
	IpsecRemoteAuthPsk   IpsecRemoteAuthPsk   `json:"ipsec_remote_auth_psk"`
	IpsecRemoteAuthRsa   IpsecRemoteAuthRsa   `json:"ipsec_remote_auth_rsa"`
	IpsecRemoteAuthX509  IpsecRemoteAuthX509  `json:"ipsec_remote_auth_x509"`
}

var _ sophos.Endpoint = &IpsecRemoteAuth{}

var defsIpsecRemoteAuth = map[string]sophos.RestObject{
	"IpsecRemoteAuthCa":    &IpsecRemoteAuthCa{},
	"IpsecRemoteAuthGroup": &IpsecRemoteAuthGroup{},
	"IpsecRemoteAuthPsk":   &IpsecRemoteAuthPsk{},
	"IpsecRemoteAuthRsa":   &IpsecRemoteAuthRsa{},
	"IpsecRemoteAuthX509":  &IpsecRemoteAuthX509{},
}

// RestObjects implements the sophos.Node interface and returns a map of IpsecRemoteAuth's Objects
func (IpsecRemoteAuth) RestObjects() map[string]sophos.RestObject { return defsIpsecRemoteAuth }

// GetPath implements sophos.RestGetter
func (*IpsecRemoteAuth) GetPath() string { return "/api/nodes/ipsec_remote_auth" }

// RefRequired implements sophos.RestGetter
func (*IpsecRemoteAuth) RefRequired() (string, bool) { return "", false }

var defIpsecRemoteAuth = &sophos.Definition{Description: "ipsec_remote_auth", Name: "ipsec_remote_auth", Link: "/api/definitions/ipsec_remote_auth"}

// Definition returns the /api/definitions struct of IpsecRemoteAuth
func (IpsecRemoteAuth) Definition() sophos.Definition { return *defIpsecRemoteAuth }

// ApiRoutes returns all known IpsecRemoteAuth Paths
func (IpsecRemoteAuth) ApiRoutes() []string {
	return []string{
		"/api/objects/ipsec_remote_auth/ca/",
		"/api/objects/ipsec_remote_auth/ca/{ref}",
		"/api/objects/ipsec_remote_auth/ca/{ref}/usedby",
		"/api/objects/ipsec_remote_auth/group/",
		"/api/objects/ipsec_remote_auth/group/{ref}",
		"/api/objects/ipsec_remote_auth/group/{ref}/usedby",
		"/api/objects/ipsec_remote_auth/psk/",
		"/api/objects/ipsec_remote_auth/psk/{ref}",
		"/api/objects/ipsec_remote_auth/psk/{ref}/usedby",
		"/api/objects/ipsec_remote_auth/rsa/",
		"/api/objects/ipsec_remote_auth/rsa/{ref}",
		"/api/objects/ipsec_remote_auth/rsa/{ref}/usedby",
		"/api/objects/ipsec_remote_auth/x509/",
		"/api/objects/ipsec_remote_auth/x509/{ref}",
		"/api/objects/ipsec_remote_auth/x509/{ref}/usedby",
	}
}

// References returns the IpsecRemoteAuth's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (IpsecRemoteAuth) References() []string {
	return []string{
		"REF_IpsecRemoteAuthCa",
		"REF_IpsecRemoteAuthGroup",
		"REF_IpsecRemoteAuthPsk",
		"REF_IpsecRemoteAuthRsa",
		"REF_IpsecRemoteAuthX509",
	}
}

// IpsecRemoteAuthCas is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecRemoteAuthCas []IpsecRemoteAuthCa

// IpsecRemoteAuthCa represents a UTM X509 CA and DN match
type IpsecRemoteAuthCa struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	// VpnId default value is "C=*, ST=*, L=*, O=*, OU=*, CN=*, E=*"
	VpnId string `json:"vpn_id"`
	// Certificate description: REF(ca/signing_ca), REF(ca/verification_ca)
	Certificate string `json:"certificate"`
	Comment     string `json:"comment"`
	Name        string `json:"name"`
}

var _ sophos.RestGetter = &IpsecRemoteAuthCa{}

// GetPath implements sophos.RestObject and returns the IpsecRemoteAuthCas GET path
// Returns all available ipsec_remote_auth/ca objects
func (*IpsecRemoteAuthCas) GetPath() string { return "/api/objects/ipsec_remote_auth/ca/" }

// RefRequired implements sophos.RestObject
func (*IpsecRemoteAuthCas) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsecRemoteAuthCas GET path
// Returns all available ca types
func (i *IpsecRemoteAuthCa) GetPath() string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/ca/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsecRemoteAuthCa) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsecRemoteAuthCa DELETE path
// Creates or updates the complete object ca
func (*IpsecRemoteAuthCa) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/ca/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecRemoteAuthCa PATCH path
// Changes to parts of the object ca types
func (*IpsecRemoteAuthCa) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/ca/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecRemoteAuthCa POST path
// Create a new ipsec_remote_auth/ca object
func (*IpsecRemoteAuthCa) PostPath() string {
	return "/api/objects/ipsec_remote_auth/ca/"
}

// PutPath implements sophos.RestObject and returns the IpsecRemoteAuthCa PUT path
// Creates or updates the complete object ca
func (*IpsecRemoteAuthCa) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/ca/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecRemoteAuthCa) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/ca/%s/usedby", ref)
}

// IpsecRemoteAuthGroups is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecRemoteAuthGroups []IpsecRemoteAuthGroup

// IpsecRemoteAuthGroup represents a UTM group
type IpsecRemoteAuthGroup struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
}

var _ sophos.RestGetter = &IpsecRemoteAuthGroup{}

// GetPath implements sophos.RestObject and returns the IpsecRemoteAuthGroups GET path
// Returns all available ipsec_remote_auth/group objects
func (*IpsecRemoteAuthGroups) GetPath() string { return "/api/objects/ipsec_remote_auth/group/" }

// RefRequired implements sophos.RestObject
func (*IpsecRemoteAuthGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsecRemoteAuthGroups GET path
// Returns all available group types
func (i *IpsecRemoteAuthGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/group/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsecRemoteAuthGroup) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsecRemoteAuthGroup DELETE path
// Creates or updates the complete object group
func (*IpsecRemoteAuthGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecRemoteAuthGroup PATCH path
// Changes to parts of the object group types
func (*IpsecRemoteAuthGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecRemoteAuthGroup POST path
// Create a new ipsec_remote_auth/group object
func (*IpsecRemoteAuthGroup) PostPath() string {
	return "/api/objects/ipsec_remote_auth/group/"
}

// PutPath implements sophos.RestObject and returns the IpsecRemoteAuthGroup PUT path
// Creates or updates the complete object group
func (*IpsecRemoteAuthGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecRemoteAuthGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/group/%s/usedby", ref)
}

// IpsecRemoteAuthPsks is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecRemoteAuthPsks []IpsecRemoteAuthPsk

// IpsecRemoteAuthPsk is a generated Sophos object
type IpsecRemoteAuthPsk struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
	Psk       string `json:"psk"`
	VpnID     string `json:"vpn_id"`
	VpnIDType string `json:"vpn_id_type"`
}

var _ sophos.RestGetter = &IpsecRemoteAuthPsk{}

// GetPath implements sophos.RestObject and returns the IpsecRemoteAuthPsks GET path
// Returns all available ipsec_remote_auth/psk objects
func (*IpsecRemoteAuthPsks) GetPath() string { return "/api/objects/ipsec_remote_auth/psk/" }

// RefRequired implements sophos.RestObject
func (*IpsecRemoteAuthPsks) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsecRemoteAuthPsks GET path
// Returns all available psk types
func (i *IpsecRemoteAuthPsk) GetPath() string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/psk/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsecRemoteAuthPsk) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsecRemoteAuthPsk DELETE path
// Creates or updates the complete object psk
func (*IpsecRemoteAuthPsk) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/psk/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecRemoteAuthPsk PATCH path
// Changes to parts of the object psk types
func (*IpsecRemoteAuthPsk) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/psk/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecRemoteAuthPsk POST path
// Create a new ipsec_remote_auth/psk object
func (*IpsecRemoteAuthPsk) PostPath() string {
	return "/api/objects/ipsec_remote_auth/psk/"
}

// PutPath implements sophos.RestObject and returns the IpsecRemoteAuthPsk PUT path
// Creates or updates the complete object psk
func (*IpsecRemoteAuthPsk) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/psk/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecRemoteAuthPsk) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/psk/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *IpsecRemoteAuthPsk) GetType() string { return i._type }

// IpsecRemoteAuthRsas is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecRemoteAuthRsas []IpsecRemoteAuthRsa

// IpsecRemoteAuthRsa represents a UTM RSA public key
type IpsecRemoteAuthRsa struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
	Pubkey    string `json:"pubkey"`
	// VpnId default value is ""
	VpnId string `json:"vpn_id"`
	// VpnIdType can be one of: []string{"ipv4_address", "fqdn", "user_fqdn"}
	VpnIdType string `json:"vpn_id_type"`
}

var _ sophos.RestGetter = &IpsecRemoteAuthRsa{}

// GetPath implements sophos.RestObject and returns the IpsecRemoteAuthRsas GET path
// Returns all available ipsec_remote_auth/rsa objects
func (*IpsecRemoteAuthRsas) GetPath() string { return "/api/objects/ipsec_remote_auth/rsa/" }

// RefRequired implements sophos.RestObject
func (*IpsecRemoteAuthRsas) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsecRemoteAuthRsas GET path
// Returns all available rsa types
func (i *IpsecRemoteAuthRsa) GetPath() string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/rsa/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsecRemoteAuthRsa) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsecRemoteAuthRsa DELETE path
// Creates or updates the complete object rsa
func (*IpsecRemoteAuthRsa) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/rsa/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecRemoteAuthRsa PATCH path
// Changes to parts of the object rsa types
func (*IpsecRemoteAuthRsa) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/rsa/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecRemoteAuthRsa POST path
// Create a new ipsec_remote_auth/rsa object
func (*IpsecRemoteAuthRsa) PostPath() string {
	return "/api/objects/ipsec_remote_auth/rsa/"
}

// PutPath implements sophos.RestObject and returns the IpsecRemoteAuthRsa PUT path
// Creates or updates the complete object rsa
func (*IpsecRemoteAuthRsa) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/rsa/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecRemoteAuthRsa) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/rsa/%s/usedby", ref)
}

// IpsecRemoteAuthX509s is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecRemoteAuthX509s []IpsecRemoteAuthX509

// IpsecRemoteAuthX509 is a generated Sophos object
type IpsecRemoteAuthX509 struct {
	Locked      string `json:"_locked"`
	Reference   string `json:"_ref"`
	_type       string `json:"_type"`
	Certificate string `json:"certificate"`
	Comment     string `json:"comment"`
	Name        string `json:"name"`
	VpnID       string `json:"vpn_id"`
	VpnIDType   string `json:"vpn_id_type"`
}

var _ sophos.RestGetter = &IpsecRemoteAuthX509{}

// GetPath implements sophos.RestObject and returns the IpsecRemoteAuthX509s GET path
// Returns all available ipsec_remote_auth/x509 objects
func (*IpsecRemoteAuthX509s) GetPath() string { return "/api/objects/ipsec_remote_auth/x509/" }

// RefRequired implements sophos.RestObject
func (*IpsecRemoteAuthX509s) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsecRemoteAuthX509s GET path
// Returns all available x509 types
func (i *IpsecRemoteAuthX509) GetPath() string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/x509/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsecRemoteAuthX509) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsecRemoteAuthX509 DELETE path
// Creates or updates the complete object x509
func (*IpsecRemoteAuthX509) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/x509/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecRemoteAuthX509 PATCH path
// Changes to parts of the object x509 types
func (*IpsecRemoteAuthX509) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/x509/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecRemoteAuthX509 POST path
// Create a new ipsec_remote_auth/x509 object
func (*IpsecRemoteAuthX509) PostPath() string {
	return "/api/objects/ipsec_remote_auth/x509/"
}

// PutPath implements sophos.RestObject and returns the IpsecRemoteAuthX509 PUT path
// Creates or updates the complete object x509
func (*IpsecRemoteAuthX509) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/x509/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecRemoteAuthX509) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec_remote_auth/x509/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *IpsecRemoteAuthX509) GetType() string { return i._type }
