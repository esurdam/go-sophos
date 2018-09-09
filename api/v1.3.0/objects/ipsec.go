package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Ipsec is a generated struct representing the Sophos Ipsec Endpoint
// GET /api/nodes/ipsec
type Ipsec struct {
	Advanced struct {
		CrlAutoFetching       int64         `json:"crl_auto_fetching"`
		CrlStrictPolicy       int64         `json:"crl_strict_policy"`
		DeadPeerDetection     int64         `json:"dead_peer_detection"`
		IkeDebug              []interface{} `json:"ike_debug"`
		IkePort               int64         `json:"ike_port"`
		Metric                int64         `json:"metric"`
		NatTraversal          int64         `json:"nat_traversal"`
		NatTraversalKeepalive int64         `json:"nat_traversal_keepalive"`
		ProbePsk              int64         `json:"probe_psk"`
		PskVpnID              string        `json:"psk_vpn_id"`
		PskVpnIDType          string        `json:"psk_vpn_id_type"`
	} `json:"advanced"`
	Connections []string `json:"connections"`
	LocalRsa    string   `json:"local_rsa"`
	LocalX509   string   `json:"local_x509"`
	Status      int64    `json:"status"`
}

var defsIpsec = map[string]sophos.RestObject{
	"IpsecGroup":         &IpsecGroup{},
	"IpsecPolicy":        &IpsecPolicy{},
	"IpsecRemoteGateway": &IpsecRemoteGateway{},
}

// RestObjects implements the sophos.Node interface and returns a map of Ipsec's Objects
func (Ipsec) RestObjects() map[string]sophos.RestObject { return defsIpsec }

// GetPath implements sophos.RestGetter
func (*Ipsec) GetPath() string { return "/api/nodes/ipsec" }

// RefRequired implements sophos.RestGetter
func (*Ipsec) RefRequired() (string, bool) { return "", false }

var defIpsec = &sophos.Definition{Description: "ipsec", Name: "ipsec", Link: "/api/definitions/ipsec"}

// Definition returns the /api/definitions struct of Ipsec
func (Ipsec) Definition() sophos.Definition { return *defIpsec }

// ApiRoutes returns all known Ipsec Paths
func (Ipsec) ApiRoutes() []string {
	return []string{
		"/api/objects/ipsec/group/",
		"/api/objects/ipsec/group/{ref}",
		"/api/objects/ipsec/group/{ref}/usedby",
		"/api/objects/ipsec/policy/",
		"/api/objects/ipsec/policy/{ref}",
		"/api/objects/ipsec/policy/{ref}/usedby",
		"/api/objects/ipsec/remote_gateway/",
		"/api/objects/ipsec/remote_gateway/{ref}",
		"/api/objects/ipsec/remote_gateway/{ref}/usedby",
	}
}

// References returns the Ipsec's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Ipsec) References() []string {
	return []string{
		"REF_IpsecGroup",
		"REF_IpsecPolicy",
		"REF_IpsecRemoteGateway",
	}
}

// IpsecGroup is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecGroup []interface{}

// GetPath implements sophos.RestObject and returns the IpsecGroup GET path
// Returns all available ipsec/group objects
func (*IpsecGroup) GetPath() string { return "/api/objects/ipsec/group/" }

// RefRequired implements sophos.RestObject
func (*IpsecGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the IpsecGroup DELETE path
// Creates or updates the complete object group
func (*IpsecGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecGroup PATCH path
// Changes to parts of the object group types
func (*IpsecGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecGroup POST path
// Create a new ipsec/group object
func (*IpsecGroup) PostPath() string {
	return "/api/objects/ipsec/group/"
}

// PutPath implements sophos.RestObject and returns the IpsecGroup PUT path
// Creates or updates the complete object group
func (*IpsecGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/group/%s/usedby", ref)
}

// IpsecPolicys is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecPolicys []IpsecPolicy

// IpsecPolicy is a generated Sophos object
type IpsecPolicy struct {
	Locked            string `json:"_locked"`
	Reference         string `json:"_ref"`
	_type             string `json:"_type"`
	Comment           string `json:"comment"`
	IkeAuthAlg        string `json:"ike_auth_alg"`
	IkeDhGroup        string `json:"ike_dh_group"`
	IkeEncAlg         string `json:"ike_enc_alg"`
	IkeSaLifetime     int64  `json:"ike_sa_lifetime"`
	IpsecAuthAlg      string `json:"ipsec_auth_alg"`
	IpsecCompression  bool   `json:"ipsec_compression"`
	IpsecEncAlg       string `json:"ipsec_enc_alg"`
	IpsecPfsGroup     string `json:"ipsec_pfs_group"`
	IpsecSaLifetime   int64  `json:"ipsec_sa_lifetime"`
	IpsecStrictPolicy bool   `json:"ipsec_strict_policy"`
	Name              string `json:"name"`
}

// GetPath implements sophos.RestObject and returns the IpsecPolicys GET path
// Returns all available ipsec/policy objects
func (*IpsecPolicys) GetPath() string { return "/api/objects/ipsec/policy/" }

// RefRequired implements sophos.RestObject
func (*IpsecPolicys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsecPolicys GET path
// Returns all available policy types
func (i *IpsecPolicy) GetPath() string {
	return fmt.Sprintf("/api/objects/ipsec/policy/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsecPolicy) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsecPolicy DELETE path
// Creates or updates the complete object policy
func (*IpsecPolicy) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/policy/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecPolicy PATCH path
// Changes to parts of the object policy types
func (*IpsecPolicy) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/policy/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecPolicy POST path
// Create a new ipsec/policy object
func (*IpsecPolicy) PostPath() string {
	return "/api/objects/ipsec/policy/"
}

// PutPath implements sophos.RestObject and returns the IpsecPolicy PUT path
// Creates or updates the complete object policy
func (*IpsecPolicy) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/policy/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecPolicy) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/policy/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *IpsecPolicy) GetType() string { return i._type }

// IpsecRemoteGateways is an Sophos Endpoint subType and implements sophos.RestObject
type IpsecRemoteGateways []IpsecRemoteGateway

// IpsecRemoteGateway is a generated Sophos object
type IpsecRemoteGateway struct {
	Locked         string   `json:"_locked"`
	Reference      string   `json:"_ref"`
	_type          string   `json:"_type"`
	Authentication string   `json:"authentication"`
	Comment        string   `json:"comment"`
	Ecn            bool     `json:"ecn"`
	Host           string   `json:"host"`
	Name           string   `json:"name"`
	Networks       []string `json:"networks"`
	PmtuDiscovery  bool     `json:"pmtu_discovery"`
	Xauth          bool     `json:"xauth"`
	XauthPassword  string   `json:"xauth_password"`
	XauthUsername  string   `json:"xauth_username"`
}

// GetPath implements sophos.RestObject and returns the IpsecRemoteGateways GET path
// Returns all available ipsec/remote_gateway objects
func (*IpsecRemoteGateways) GetPath() string { return "/api/objects/ipsec/remote_gateway/" }

// RefRequired implements sophos.RestObject
func (*IpsecRemoteGateways) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsecRemoteGateways GET path
// Returns all available remote_gateway types
func (i *IpsecRemoteGateway) GetPath() string {
	return fmt.Sprintf("/api/objects/ipsec/remote_gateway/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsecRemoteGateway) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsecRemoteGateway DELETE path
// Creates or updates the complete object remote_gateway
func (*IpsecRemoteGateway) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/remote_gateway/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsecRemoteGateway PATCH path
// Changes to parts of the object remote_gateway types
func (*IpsecRemoteGateway) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/remote_gateway/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsecRemoteGateway POST path
// Create a new ipsec/remote_gateway object
func (*IpsecRemoteGateway) PostPath() string {
	return "/api/objects/ipsec/remote_gateway/"
}

// PutPath implements sophos.RestObject and returns the IpsecRemoteGateway PUT path
// Creates or updates the complete object remote_gateway
func (*IpsecRemoteGateway) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/remote_gateway/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*IpsecRemoteGateway) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ipsec/remote_gateway/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *IpsecRemoteGateway) GetType() string { return i._type }
