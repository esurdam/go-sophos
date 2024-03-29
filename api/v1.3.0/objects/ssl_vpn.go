package objects

// THIS FILE IS AUTO GENERATED BY bin/gen.go! DO NOT EDIT!

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// SslVpn is a generated struct representing the Sophos SslVpn Endpoint
// GET /api/nodes/ssl_vpn
type SslVpn struct {
	AuthenticationAlgorithm string `json:"authentication_algorithm"`
	Certificate             string `json:"certificate"`
	Compression             int64  `json:"compression"`
	DatachannelKeyLifetime  int64  `json:"datachannel_key_lifetime"`
	Debug                   int64  `json:"debug"`
	DhKeySize               string `json:"dh_key_size"`
	DuplicateCn             int64  `json:"duplicate_cn"`
	EncryptionAlgorithm     string `json:"encryption_algorithm"`
	FallbackDisable         int64  `json:"fallback_disable"`
	Hostname                string `json:"hostname"`
	Interface               string `json:"interface"`
	InterfaceAddress        string `json:"interface_address"`
	IPAssignmentPool        string `json:"ip_assignment_pool"`
	Port                    int64  `json:"port"`
	Protocol                string `json:"protocol"`
	UserAuthOptional        int64  `json:"user_auth_optional"`
}

var _ sophos.Endpoint = &SslVpn{}

var defsSslVpn = map[string]sophos.RestObject{
	"SslVpnClientConnection":    &SslVpnClientConnection{},
	"SslVpnGroup":               &SslVpnGroup{},
	"SslVpnRemoteAccessProfile": &SslVpnRemoteAccessProfile{},
	"SslVpnServerConnection":    &SslVpnServerConnection{},
}

// RestObjects implements the sophos.Node interface and returns a map of SslVpn's Objects
func (SslVpn) RestObjects() map[string]sophos.RestObject { return defsSslVpn }

// GetPath implements sophos.RestGetter
func (*SslVpn) GetPath() string { return "/api/nodes/ssl_vpn" }

// RefRequired implements sophos.RestGetter
func (*SslVpn) RefRequired() (string, bool) { return "", false }

var defSslVpn = &sophos.Definition{Description: "ssl_vpn", Name: "ssl_vpn", Link: "/api/definitions/ssl_vpn"}

// Definition returns the /api/definitions struct of SslVpn
func (SslVpn) Definition() sophos.Definition { return *defSslVpn }

// ApiRoutes returns all known SslVpn Paths
func (SslVpn) ApiRoutes() []string {
	return []string{
		"/api/objects/ssl_vpn/client_connection/",
		"/api/objects/ssl_vpn/client_connection/{ref}",
		"/api/objects/ssl_vpn/client_connection/{ref}/usedby",
		"/api/objects/ssl_vpn/group/",
		"/api/objects/ssl_vpn/group/{ref}",
		"/api/objects/ssl_vpn/group/{ref}/usedby",
		"/api/objects/ssl_vpn/remote_access_profile/",
		"/api/objects/ssl_vpn/remote_access_profile/{ref}",
		"/api/objects/ssl_vpn/remote_access_profile/{ref}/usedby",
		"/api/objects/ssl_vpn/server_connection/",
		"/api/objects/ssl_vpn/server_connection/{ref}",
		"/api/objects/ssl_vpn/server_connection/{ref}/usedby",
	}
}

// References returns the SslVpn's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (SslVpn) References() []string {
	return []string{
		"REF_SslVpnClientConnection",
		"REF_SslVpnGroup",
		"REF_SslVpnRemoteAccessProfile",
		"REF_SslVpnServerConnection",
	}
}

// SslVpnClientConnections is an Sophos Endpoint subType and implements sophos.RestObject
type SslVpnClientConnections []SslVpnClientConnection

// SslVpnClientConnection represents a UTM SSL VPN site-to-site client connection
type SslVpnClientConnection struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// AuthenticationAlgorithm can be one of: []string{"SHA1", "SHA256", "SHA384", "SHA512", "MD5", "gost-mac"}
	AuthenticationAlgorithm string `json:"authentication_algorithm"`
	// AutoPfIn description: REF(packetfilter/packetfilter)
	// AutoPfIn default value is ""
	AutoPfIn string `json:"auto_pf_in"`
	// AutoPfOut description: REF(packetfilter/packetfilter)
	// AutoPfOut default value is ""
	AutoPfOut string `json:"auto_pf_out"`
	// AutoPfrule default value is false
	AutoPfrule  bool   `json:"auto_pfrule"`
	CaCert      string `json:"ca_cert"`
	Certificate string `json:"certificate"`
	Comment     string `json:"comment"`
	// Compression default value is false
	Compression bool `json:"compression"`
	// EncryptionAlgorithm can be one of: []string{"DES-EDE3-CBC", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "BF-CBC", "gost89"}
	EncryptionAlgorithm string `json:"encryption_algorithm"`
	// Engine can be one of: []string{"gost"}
	// Engine default value is ""
	Engine string `json:"engine"`
	// Interface default value is ""
	Interface          string   `json:"interface"`
	Key                string   `json:"key"`
	LocalNetworks      []string `json:"local_networks"`
	Name               string   `json:"name"`
	Password           string   `json:"password"`
	PlainServerAddress string   `json:"plain_server_address"`
	// Protocol can be one of: []string{"tcp", "udp"}
	Protocol string `json:"protocol"`
	// ProxyAuthPass default value is ""
	ProxyAuthPass string `json:"proxy_auth_pass"`
	// ProxyAuthStatus default value is false
	ProxyAuthStatus bool `json:"proxy_auth_status"`
	// ProxyAuthUser default value is ""
	ProxyAuthUser string `json:"proxy_auth_user"`
	// ProxyHost description: REF(network/host), REF(network/dns_host), REF(network/availability_group)
	// ProxyHost default value is ""
	ProxyHost string `json:"proxy_host"`
	ProxyPort int    `json:"proxy_port"`
	// ProxyStatus default value is false
	ProxyStatus    bool     `json:"proxy_status"`
	RemoteNetworks []string `json:"remote_networks"`
	// ServerAddress description: REF(network/host), REF(network/dns_host)
	// ServerAddress default value is ""
	ServerAddress string `json:"server_address"`
	ServerDn      string `json:"server_dn"`
	// ServerOverrideHostname description: REF(network/host), REF(network/dns_host)
	// ServerOverrideHostname default value is ""
	ServerOverrideHostname string `json:"server_override_hostname"`
	// ServerOverrideStatus default value is false
	ServerOverrideStatus bool `json:"server_override_status"`
	ServerPort           int  `json:"server_port"`
	// Status default value is false
	Status   bool   `json:"status"`
	Username string `json:"username"`
}

var _ sophos.RestGetter = &SslVpnClientConnection{}

// GetPath implements sophos.RestObject and returns the SslVpnClientConnections GET path
// Returns all available ssl_vpn/client_connection objects
func (*SslVpnClientConnections) GetPath() string { return "/api/objects/ssl_vpn/client_connection/" }

// RefRequired implements sophos.RestObject
func (*SslVpnClientConnections) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SslVpnClientConnections GET path
// Returns all available client_connection types
func (s *SslVpnClientConnection) GetPath() string {
	return fmt.Sprintf("/api/objects/ssl_vpn/client_connection/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *SslVpnClientConnection) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SslVpnClientConnection DELETE path
// Creates or updates the complete object client_connection
func (*SslVpnClientConnection) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/client_connection/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SslVpnClientConnection PATCH path
// Changes to parts of the object client_connection types
func (*SslVpnClientConnection) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/client_connection/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SslVpnClientConnection POST path
// Create a new ssl_vpn/client_connection object
func (*SslVpnClientConnection) PostPath() string {
	return "/api/objects/ssl_vpn/client_connection/"
}

// PutPath implements sophos.RestObject and returns the SslVpnClientConnection PUT path
// Creates or updates the complete object client_connection
func (*SslVpnClientConnection) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/client_connection/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SslVpnClientConnection) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/client_connection/%s/usedby", ref)
}

// SslVpnGroups is an Sophos Endpoint subType and implements sophos.RestObject
type SslVpnGroups []SslVpnGroup

// SslVpnGroup represents a UTM group
type SslVpnGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &SslVpnGroup{}

// GetPath implements sophos.RestObject and returns the SslVpnGroups GET path
// Returns all available ssl_vpn/group objects
func (*SslVpnGroups) GetPath() string { return "/api/objects/ssl_vpn/group/" }

// RefRequired implements sophos.RestObject
func (*SslVpnGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SslVpnGroups GET path
// Returns all available group types
func (s *SslVpnGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/ssl_vpn/group/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *SslVpnGroup) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SslVpnGroup DELETE path
// Creates or updates the complete object group
func (*SslVpnGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SslVpnGroup PATCH path
// Changes to parts of the object group types
func (*SslVpnGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SslVpnGroup POST path
// Create a new ssl_vpn/group object
func (*SslVpnGroup) PostPath() string {
	return "/api/objects/ssl_vpn/group/"
}

// PutPath implements sophos.RestObject and returns the SslVpnGroup PUT path
// Creates or updates the complete object group
func (*SslVpnGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SslVpnGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/group/%s/usedby", ref)
}

// SslVpnRemoteAccessProfiles is an Sophos Endpoint subType and implements sophos.RestObject
type SslVpnRemoteAccessProfiles []SslVpnRemoteAccessProfile

// SslVpnRemoteAccessProfile represents a UTM SSL VPN remote access profile
type SslVpnRemoteAccessProfile struct {
	Locked     string   `json:"_locked"`
	ObjectType string   `json:"_type"`
	Reference  string   `json:"_ref"`
	Aaa        []string `json:"aaa"`
	// AutoPfIn description: REF(packetfilter/packetfilter)
	// AutoPfIn default value is ""
	AutoPfIn string `json:"auto_pf_in"`
	// AutoPfrule default value is false
	AutoPfrule bool     `json:"auto_pfrule"`
	Comment    string   `json:"comment"`
	Name       string   `json:"name"`
	Networks   []string `json:"networks"`
	// Status default value is false
	Status bool `json:"status"`
}

var _ sophos.RestGetter = &SslVpnRemoteAccessProfile{}

// GetPath implements sophos.RestObject and returns the SslVpnRemoteAccessProfiles GET path
// Returns all available ssl_vpn/remote_access_profile objects
func (*SslVpnRemoteAccessProfiles) GetPath() string {
	return "/api/objects/ssl_vpn/remote_access_profile/"
}

// RefRequired implements sophos.RestObject
func (*SslVpnRemoteAccessProfiles) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SslVpnRemoteAccessProfiles GET path
// Returns all available remote_access_profile types
func (s *SslVpnRemoteAccessProfile) GetPath() string {
	return fmt.Sprintf("/api/objects/ssl_vpn/remote_access_profile/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *SslVpnRemoteAccessProfile) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SslVpnRemoteAccessProfile DELETE path
// Creates or updates the complete object remote_access_profile
func (*SslVpnRemoteAccessProfile) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/remote_access_profile/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SslVpnRemoteAccessProfile PATCH path
// Changes to parts of the object remote_access_profile types
func (*SslVpnRemoteAccessProfile) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/remote_access_profile/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SslVpnRemoteAccessProfile POST path
// Create a new ssl_vpn/remote_access_profile object
func (*SslVpnRemoteAccessProfile) PostPath() string {
	return "/api/objects/ssl_vpn/remote_access_profile/"
}

// PutPath implements sophos.RestObject and returns the SslVpnRemoteAccessProfile PUT path
// Creates or updates the complete object remote_access_profile
func (*SslVpnRemoteAccessProfile) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/remote_access_profile/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SslVpnRemoteAccessProfile) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/remote_access_profile/%s/usedby", ref)
}

// SslVpnServerConnections is an Sophos Endpoint subType and implements sophos.RestObject
type SslVpnServerConnections []SslVpnServerConnection

// SslVpnServerConnection represents a UTM SSL VPN site-to-site server connection
type SslVpnServerConnection struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// AutoPfIn description: REF(packetfilter/packetfilter)
	// AutoPfIn default value is ""
	AutoPfIn string `json:"auto_pf_in"`
	// AutoPfOut description: REF(packetfilter/packetfilter)
	// AutoPfOut default value is ""
	AutoPfOut string `json:"auto_pf_out"`
	// AutoPfrule default value is false
	AutoPfrule    bool     `json:"auto_pfrule"`
	Comment       string   `json:"comment"`
	LocalNetworks []string `json:"local_networks"`
	Name          string   `json:"name"`
	// Peer description: REF(aaa/user)
	// Peer default value is ""
	Peer           string   `json:"peer"`
	RemoteNetworks []string `json:"remote_networks"`
	// StaticIp description: (IPADDR)
	// StaticIp default value is "0.0.0.0"
	StaticIp string `json:"static_ip"`
	// StaticIp6 description: (IP6ADDR)
	// StaticIp6 default value is "::"
	StaticIp6 string `json:"static_ip6"`
	// StaticIpStatus default value is false
	StaticIpStatus bool `json:"static_ip_status"`
	// Status default value is false
	Status bool `json:"status"`
}

var _ sophos.RestGetter = &SslVpnServerConnection{}

// GetPath implements sophos.RestObject and returns the SslVpnServerConnections GET path
// Returns all available ssl_vpn/server_connection objects
func (*SslVpnServerConnections) GetPath() string { return "/api/objects/ssl_vpn/server_connection/" }

// RefRequired implements sophos.RestObject
func (*SslVpnServerConnections) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SslVpnServerConnections GET path
// Returns all available server_connection types
func (s *SslVpnServerConnection) GetPath() string {
	return fmt.Sprintf("/api/objects/ssl_vpn/server_connection/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *SslVpnServerConnection) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SslVpnServerConnection DELETE path
// Creates or updates the complete object server_connection
func (*SslVpnServerConnection) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/server_connection/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SslVpnServerConnection PATCH path
// Changes to parts of the object server_connection types
func (*SslVpnServerConnection) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/server_connection/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SslVpnServerConnection POST path
// Create a new ssl_vpn/server_connection object
func (*SslVpnServerConnection) PostPath() string {
	return "/api/objects/ssl_vpn/server_connection/"
}

// PutPath implements sophos.RestObject and returns the SslVpnServerConnection PUT path
// Creates or updates the complete object server_connection
func (*SslVpnServerConnection) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/server_connection/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SslVpnServerConnection) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ssl_vpn/server_connection/%s/usedby", ref)
}
