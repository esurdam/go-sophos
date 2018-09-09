package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// ReverseProxy is a generated struct representing the Sophos ReverseProxy Endpoint
// GET /api/nodes/reverse_proxy
type ReverseProxy struct {
	AuaRefreshEnabled  int64 `json:"aua_refresh_enabled"`
	AuaRefreshInterval int64 `json:"aua_refresh_interval"`
	Blacklist          struct {
		DnsrblZones []string `json:"dnsrbl_zones"`
		GeoipCodes  []string `json:"geoip_codes"`
	} `json:"blacklist"`
	Cookiesignkey                       string        `json:"cookiesignkey"`
	CssdHostname                        string        `json:"cssd_hostname"`
	CssdPort                            int64         `json:"cssd_port"`
	CustomThreatFiltersEnabled          int64         `json:"custom_threat_filters_enabled"`
	FormhardeningSecret                 string        `json:"formhardening_secret"`
	Keepalive                           int64         `json:"keepalive"`
	Manualmode                          int64         `json:"manualmode"`
	MaxConnectionsPerChild              int64         `json:"max_connections_per_child"`
	MaxPreforkProcesses                 int64         `json:"max_prefork_processes"`
	MaxProcesses                        int64         `json:"max_processes"`
	MaxSessionFiles                     int64         `json:"max_session_files"`
	MaxSpareProcesses                   int64         `json:"max_spare_processes"`
	MaxSpareThreads                     int64         `json:"max_spare_threads"`
	MaxThreadsPerProcess                int64         `json:"max_threads_per_process"`
	MinSpareProcesses                   int64         `json:"min_spare_processes"`
	MinSpareThreads                     int64         `json:"min_spare_threads"`
	MinTLS                              string        `json:"min_tls"`
	ModsecurityBeta                     int64         `json:"modsecurity_beta"`
	MpmMode                             string        `json:"mpm_mode"`
	Patternversion                      string        `json:"patternversion"`
	Port                                int64         `json:"port"`
	Proxyprotocol                       int64         `json:"proxyprotocol"`
	RequestLineLimit                    int64         `json:"request_line_limit"`
	SlowhttpExceptions                  []interface{} `json:"slowhttp_exceptions"`
	SlowhttpRequestHeaderTimeoutBase    int64         `json:"slowhttp_request_header_timeout_base"`
	SlowhttpRequestHeaderTimeoutEnabled int64         `json:"slowhttp_request_header_timeout_enabled"`
	SlowhttpRequestHeaderTimeoutMax     int64         `json:"slowhttp_request_header_timeout_max"`
	SlowhttpRequestHeaderTimeoutRate    int64         `json:"slowhttp_request_header_timeout_rate"`
	Status                              int64         `json:"status"`
	TraceEnabled                        int64         `json:"trace_enabled"`
	Urlhardeningsignkey                 string        `json:"urlhardeningsignkey"`
	Whatkilledus                        int64         `json:"whatkilledus"`
}

var _ sophos.Endpoint = &ReverseProxy{}

var defsReverseProxy = map[string]sophos.RestObject{
	"ReverseProxyAuthProfile":   &ReverseProxyAuthProfile{},
	"ReverseProxyBackend":       &ReverseProxyBackend{},
	"ReverseProxyException":     &ReverseProxyException{},
	"ReverseProxyFilter":        &ReverseProxyFilter{},
	"ReverseProxyFormTemplate":  &ReverseProxyFormTemplate{},
	"ReverseProxyFrontend":      &ReverseProxyFrontend{},
	"ReverseProxyGroup":         &ReverseProxyGroup{},
	"ReverseProxyLocation":      &ReverseProxyLocation{},
	"ReverseProxyProfile":       &ReverseProxyProfile{},
	"ReverseProxyRedirection":   &ReverseProxyRedirection{},
	"ReverseProxyThreatsFilter": &ReverseProxyThreatsFilter{},
}

// RestObjects implements the sophos.Node interface and returns a map of ReverseProxy's Objects
func (ReverseProxy) RestObjects() map[string]sophos.RestObject { return defsReverseProxy }

// GetPath implements sophos.RestGetter
func (*ReverseProxy) GetPath() string { return "/api/nodes/reverse_proxy" }

// RefRequired implements sophos.RestGetter
func (*ReverseProxy) RefRequired() (string, bool) { return "", false }

var defReverseProxy = &sophos.Definition{Description: "reverse_proxy", Name: "reverse_proxy", Link: "/api/definitions/reverse_proxy"}

// Definition returns the /api/definitions struct of ReverseProxy
func (ReverseProxy) Definition() sophos.Definition { return *defReverseProxy }

// ApiRoutes returns all known ReverseProxy Paths
func (ReverseProxy) ApiRoutes() []string {
	return []string{
		"/api/objects/reverse_proxy/auth_profile/",
		"/api/objects/reverse_proxy/auth_profile/{ref}",
		"/api/objects/reverse_proxy/auth_profile/{ref}/usedby",
		"/api/objects/reverse_proxy/backend/",
		"/api/objects/reverse_proxy/backend/{ref}",
		"/api/objects/reverse_proxy/backend/{ref}/usedby",
		"/api/objects/reverse_proxy/exception/",
		"/api/objects/reverse_proxy/exception/{ref}",
		"/api/objects/reverse_proxy/exception/{ref}/usedby",
		"/api/objects/reverse_proxy/filter/",
		"/api/objects/reverse_proxy/filter/{ref}",
		"/api/objects/reverse_proxy/filter/{ref}/usedby",
		"/api/objects/reverse_proxy/form_template/",
		"/api/objects/reverse_proxy/form_template/{ref}",
		"/api/objects/reverse_proxy/form_template/{ref}/usedby",
		"/api/objects/reverse_proxy/frontend/",
		"/api/objects/reverse_proxy/frontend/{ref}",
		"/api/objects/reverse_proxy/frontend/{ref}/usedby",
		"/api/objects/reverse_proxy/group/",
		"/api/objects/reverse_proxy/group/{ref}",
		"/api/objects/reverse_proxy/group/{ref}/usedby",
		"/api/objects/reverse_proxy/location/",
		"/api/objects/reverse_proxy/location/{ref}",
		"/api/objects/reverse_proxy/location/{ref}/usedby",
		"/api/objects/reverse_proxy/profile/",
		"/api/objects/reverse_proxy/profile/{ref}",
		"/api/objects/reverse_proxy/profile/{ref}/usedby",
		"/api/objects/reverse_proxy/redirection/",
		"/api/objects/reverse_proxy/redirection/{ref}",
		"/api/objects/reverse_proxy/redirection/{ref}/usedby",
		"/api/objects/reverse_proxy/threats_filter/",
		"/api/objects/reverse_proxy/threats_filter/{ref}",
		"/api/objects/reverse_proxy/threats_filter/{ref}/usedby",
	}
}

// References returns the ReverseProxy's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (ReverseProxy) References() []string {
	return []string{
		"REF_ReverseProxyAuthProfile",
		"REF_ReverseProxyBackend",
		"REF_ReverseProxyException",
		"REF_ReverseProxyFilter",
		"REF_ReverseProxyFormTemplate",
		"REF_ReverseProxyFrontend",
		"REF_ReverseProxyGroup",
		"REF_ReverseProxyLocation",
		"REF_ReverseProxyProfile",
		"REF_ReverseProxyRedirection",
		"REF_ReverseProxyThreatsFilter",
	}
}

// ReverseProxyAuthProfiles is an Sophos Endpoint subType and implements sophos.RestObject
type ReverseProxyAuthProfiles []ReverseProxyAuthProfile

// ReverseProxyAuthProfile is a generated Sophos object
type ReverseProxyAuthProfile struct {
	Locked                          string        `json:"_locked"`
	Reference                       string        `json:"_ref"`
	_type                           string        `json:"_type"`
	Aaa                             []string      `json:"aaa"`
	BackendMode                     string        `json:"backend_mode"`
	BackendStripBasicAuth           bool          `json:"backend_strip_basic_auth"`
	BackendUserPrefix               string        `json:"backend_user_prefix"`
	BackendUserSuffix               string        `json:"backend_user_suffix"`
	BasicPrompt                     string        `json:"basic_prompt"`
	Comment                         string        `json:"comment"`
	FrontendCookie                  string        `json:"frontend_cookie"`
	FrontendCookieSecret            string        `json:"frontend_cookie_secret"`
	FrontendForm                    string        `json:"frontend_form"`
	FrontendFormTemplate            string        `json:"frontend_form_template"`
	FrontendLogin                   string        `json:"frontend_login"`
	FrontendLogout                  string        `json:"frontend_logout"`
	FrontendMode                    string        `json:"frontend_mode"`
	FrontendRealm                   string        `json:"frontend_realm"`
	FrontendSessionAllowPersistency bool          `json:"frontend_session_allow_persistency"`
	FrontendSessionLifetime         int64         `json:"frontend_session_lifetime"`
	FrontendSessionLifetimeLimited  bool          `json:"frontend_session_lifetime_limited"`
	FrontendSessionLifetimeScope    string        `json:"frontend_session_lifetime_scope"`
	FrontendSessionTimeout          int64         `json:"frontend_session_timeout"`
	FrontendSessionTimeoutEnabled   bool          `json:"frontend_session_timeout_enabled"`
	FrontendSessionTimeoutScope     string        `json:"frontend_session_timeout_scope"`
	LogoutDelegationUrls            []interface{} `json:"logout_delegation_urls"`
	LogoutMode                      string        `json:"logout_mode"`
	Name                            string        `json:"name"`
	RedirectToRequestedURL          bool          `json:"redirect_to_requested_url"`
}

var _ sophos.RestGetter = &ReverseProxyAuthProfile{}

// GetPath implements sophos.RestObject and returns the ReverseProxyAuthProfiles GET path
// Returns all available reverse_proxy/auth_profile objects
func (*ReverseProxyAuthProfiles) GetPath() string { return "/api/objects/reverse_proxy/auth_profile/" }

// RefRequired implements sophos.RestObject
func (*ReverseProxyAuthProfiles) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ReverseProxyAuthProfiles GET path
// Returns all available auth_profile types
func (r *ReverseProxyAuthProfile) GetPath() string {
	return fmt.Sprintf("/api/objects/reverse_proxy/auth_profile/%s", r.Reference)
}

// RefRequired implements sophos.RestObject
func (r *ReverseProxyAuthProfile) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the ReverseProxyAuthProfile DELETE path
// Creates or updates the complete object auth_profile
func (*ReverseProxyAuthProfile) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/auth_profile/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReverseProxyAuthProfile PATCH path
// Changes to parts of the object auth_profile types
func (*ReverseProxyAuthProfile) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/auth_profile/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReverseProxyAuthProfile POST path
// Create a new reverse_proxy/auth_profile object
func (*ReverseProxyAuthProfile) PostPath() string {
	return "/api/objects/reverse_proxy/auth_profile/"
}

// PutPath implements sophos.RestObject and returns the ReverseProxyAuthProfile PUT path
// Creates or updates the complete object auth_profile
func (*ReverseProxyAuthProfile) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/auth_profile/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReverseProxyAuthProfile) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/auth_profile/%s/usedby", ref)
}

// GetType implements sophos.Object
func (r *ReverseProxyAuthProfile) GetType() string { return r._type }

// ReverseProxyBackends is an Sophos Endpoint subType and implements sophos.RestObject
type ReverseProxyBackends []ReverseProxyBackend

// ReverseProxyBackend is a generated Sophos object
type ReverseProxyBackend struct {
	Locked                          string `json:"_locked"`
	Reference                       string `json:"_ref"`
	_type                           string `json:"_type"`
	Comment                         string `json:"comment"`
	DisableBackendConnectionPooling bool   `json:"disable_backend_connection_pooling"`
	Host                            string `json:"host"`
	Keepalive                       bool   `json:"keepalive"`
	Name                            string `json:"name"`
	Path                            string `json:"path"`
	Port                            int64  `json:"port"`
	Ssl                             bool   `json:"ssl"`
	Status                          bool   `json:"status"`
	Timeout                         int64  `json:"timeout"`
}

var _ sophos.RestGetter = &ReverseProxyBackend{}

// GetPath implements sophos.RestObject and returns the ReverseProxyBackends GET path
// Returns all available reverse_proxy/backend objects
func (*ReverseProxyBackends) GetPath() string { return "/api/objects/reverse_proxy/backend/" }

// RefRequired implements sophos.RestObject
func (*ReverseProxyBackends) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ReverseProxyBackends GET path
// Returns all available backend types
func (r *ReverseProxyBackend) GetPath() string {
	return fmt.Sprintf("/api/objects/reverse_proxy/backend/%s", r.Reference)
}

// RefRequired implements sophos.RestObject
func (r *ReverseProxyBackend) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the ReverseProxyBackend DELETE path
// Creates or updates the complete object backend
func (*ReverseProxyBackend) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/backend/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReverseProxyBackend PATCH path
// Changes to parts of the object backend types
func (*ReverseProxyBackend) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/backend/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReverseProxyBackend POST path
// Create a new reverse_proxy/backend object
func (*ReverseProxyBackend) PostPath() string {
	return "/api/objects/reverse_proxy/backend/"
}

// PutPath implements sophos.RestObject and returns the ReverseProxyBackend PUT path
// Creates or updates the complete object backend
func (*ReverseProxyBackend) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/backend/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReverseProxyBackend) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/backend/%s/usedby", ref)
}

// GetType implements sophos.Object
func (r *ReverseProxyBackend) GetType() string { return r._type }

// ReverseProxyException is an Sophos Endpoint subType and implements sophos.RestObject
type ReverseProxyException []interface{}

var _ sophos.RestObject = &ReverseProxyException{}

// GetPath implements sophos.RestObject and returns the ReverseProxyException GET path
// Returns all available reverse_proxy/exception objects
func (*ReverseProxyException) GetPath() string { return "/api/objects/reverse_proxy/exception/" }

// RefRequired implements sophos.RestObject
func (*ReverseProxyException) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ReverseProxyException DELETE path
// Creates or updates the complete object exception
func (*ReverseProxyException) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/exception/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReverseProxyException PATCH path
// Changes to parts of the object exception types
func (*ReverseProxyException) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/exception/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReverseProxyException POST path
// Create a new reverse_proxy/exception object
func (*ReverseProxyException) PostPath() string {
	return "/api/objects/reverse_proxy/exception/"
}

// PutPath implements sophos.RestObject and returns the ReverseProxyException PUT path
// Creates or updates the complete object exception
func (*ReverseProxyException) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/exception/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReverseProxyException) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/exception/%s/usedby", ref)
}

// ReverseProxyFilter is an Sophos Endpoint subType and implements sophos.RestObject
type ReverseProxyFilter []interface{}

var _ sophos.RestObject = &ReverseProxyFilter{}

// GetPath implements sophos.RestObject and returns the ReverseProxyFilter GET path
// Returns all available reverse_proxy/filter objects
func (*ReverseProxyFilter) GetPath() string { return "/api/objects/reverse_proxy/filter/" }

// RefRequired implements sophos.RestObject
func (*ReverseProxyFilter) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ReverseProxyFilter DELETE path
// Creates or updates the complete object filter
func (*ReverseProxyFilter) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/filter/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReverseProxyFilter PATCH path
// Changes to parts of the object filter types
func (*ReverseProxyFilter) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/filter/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReverseProxyFilter POST path
// Create a new reverse_proxy/filter object
func (*ReverseProxyFilter) PostPath() string {
	return "/api/objects/reverse_proxy/filter/"
}

// PutPath implements sophos.RestObject and returns the ReverseProxyFilter PUT path
// Creates or updates the complete object filter
func (*ReverseProxyFilter) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/filter/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReverseProxyFilter) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/filter/%s/usedby", ref)
}

// ReverseProxyFormTemplates is an Sophos Endpoint subType and implements sophos.RestObject
type ReverseProxyFormTemplates []ReverseProxyFormTemplate

// ReverseProxyFormTemplate is a generated Sophos object
type ReverseProxyFormTemplate struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Assets    struct {
		DefaultStylesheet_css string `json:"default_stylesheet.css"`
	} `json:"assets"`
	Comment  string `json:"comment"`
	Filename string `json:"filename"`
	Name     string `json:"name"`
	Template string `json:"template"`
}

var _ sophos.RestGetter = &ReverseProxyFormTemplate{}

// GetPath implements sophos.RestObject and returns the ReverseProxyFormTemplates GET path
// Returns all available reverse_proxy/form_template objects
func (*ReverseProxyFormTemplates) GetPath() string { return "/api/objects/reverse_proxy/form_template/" }

// RefRequired implements sophos.RestObject
func (*ReverseProxyFormTemplates) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ReverseProxyFormTemplates GET path
// Returns all available form_template types
func (r *ReverseProxyFormTemplate) GetPath() string {
	return fmt.Sprintf("/api/objects/reverse_proxy/form_template/%s", r.Reference)
}

// RefRequired implements sophos.RestObject
func (r *ReverseProxyFormTemplate) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the ReverseProxyFormTemplate DELETE path
// Creates or updates the complete object form_template
func (*ReverseProxyFormTemplate) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/form_template/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReverseProxyFormTemplate PATCH path
// Changes to parts of the object form_template types
func (*ReverseProxyFormTemplate) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/form_template/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReverseProxyFormTemplate POST path
// Create a new reverse_proxy/form_template object
func (*ReverseProxyFormTemplate) PostPath() string {
	return "/api/objects/reverse_proxy/form_template/"
}

// PutPath implements sophos.RestObject and returns the ReverseProxyFormTemplate PUT path
// Creates or updates the complete object form_template
func (*ReverseProxyFormTemplate) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/form_template/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReverseProxyFormTemplate) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/form_template/%s/usedby", ref)
}

// GetType implements sophos.Object
func (r *ReverseProxyFormTemplate) GetType() string { return r._type }

// ReverseProxyFrontends is an Sophos Endpoint subType and implements sophos.RestObject
type ReverseProxyFrontends []ReverseProxyFrontend

// ReverseProxyFrontend is a generated Sophos object
type ReverseProxyFrontend struct {
	Locked               string        `json:"_locked"`
	Reference            string        `json:"_ref"`
	_type                string        `json:"_type"`
	AddContentTypeHeader bool          `json:"add_content_type_header"`
	Address              string        `json:"address"`
	AllowedNetworks      []string      `json:"allowed_networks"`
	Certificate          string        `json:"certificate"`
	Comment              string        `json:"comment"`
	DisableCompression   bool          `json:"disable_compression"`
	Domain               []string      `json:"domain"`
	Exceptions           []interface{} `json:"exceptions"`
	Htmlrewrite          bool          `json:"htmlrewrite"`
	HtmlrewriteCookies   bool          `json:"htmlrewrite_cookies"`
	Implicitredirect     bool          `json:"implicitredirect"`
	Lbmethod             string        `json:"lbmethod"`
	Locations            []string      `json:"locations"`
	Name                 string        `json:"name"`
	Port                 int64         `json:"port"`
	Preservehost         bool          `json:"preservehost"`
	Profile              string        `json:"profile"`
	Status               bool          `json:"status"`
	Type                 string        `json:"type"`
	Xheaders             bool          `json:"xheaders"`
}

var _ sophos.RestGetter = &ReverseProxyFrontend{}

// GetPath implements sophos.RestObject and returns the ReverseProxyFrontends GET path
// Returns all available reverse_proxy/frontend objects
func (*ReverseProxyFrontends) GetPath() string { return "/api/objects/reverse_proxy/frontend/" }

// RefRequired implements sophos.RestObject
func (*ReverseProxyFrontends) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ReverseProxyFrontends GET path
// Returns all available frontend types
func (r *ReverseProxyFrontend) GetPath() string {
	return fmt.Sprintf("/api/objects/reverse_proxy/frontend/%s", r.Reference)
}

// RefRequired implements sophos.RestObject
func (r *ReverseProxyFrontend) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the ReverseProxyFrontend DELETE path
// Creates or updates the complete object frontend
func (*ReverseProxyFrontend) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/frontend/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReverseProxyFrontend PATCH path
// Changes to parts of the object frontend types
func (*ReverseProxyFrontend) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/frontend/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReverseProxyFrontend POST path
// Create a new reverse_proxy/frontend object
func (*ReverseProxyFrontend) PostPath() string {
	return "/api/objects/reverse_proxy/frontend/"
}

// PutPath implements sophos.RestObject and returns the ReverseProxyFrontend PUT path
// Creates or updates the complete object frontend
func (*ReverseProxyFrontend) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/frontend/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReverseProxyFrontend) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/frontend/%s/usedby", ref)
}

// GetType implements sophos.Object
func (r *ReverseProxyFrontend) GetType() string { return r._type }

// ReverseProxyGroup is an Sophos Endpoint subType and implements sophos.RestObject
type ReverseProxyGroup []interface{}

var _ sophos.RestObject = &ReverseProxyGroup{}

// GetPath implements sophos.RestObject and returns the ReverseProxyGroup GET path
// Returns all available reverse_proxy/group objects
func (*ReverseProxyGroup) GetPath() string { return "/api/objects/reverse_proxy/group/" }

// RefRequired implements sophos.RestObject
func (*ReverseProxyGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ReverseProxyGroup DELETE path
// Creates or updates the complete object group
func (*ReverseProxyGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReverseProxyGroup PATCH path
// Changes to parts of the object group types
func (*ReverseProxyGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReverseProxyGroup POST path
// Create a new reverse_proxy/group object
func (*ReverseProxyGroup) PostPath() string {
	return "/api/objects/reverse_proxy/group/"
}

// PutPath implements sophos.RestObject and returns the ReverseProxyGroup PUT path
// Creates or updates the complete object group
func (*ReverseProxyGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReverseProxyGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/group/%s/usedby", ref)
}

// ReverseProxyLocations is an Sophos Endpoint subType and implements sophos.RestObject
type ReverseProxyLocations []ReverseProxyLocation

// ReverseProxyLocation is a generated Sophos object
type ReverseProxyLocation struct {
	Locked               string        `json:"_locked"`
	Reference            string        `json:"_ref"`
	_type                string        `json:"_type"`
	AccessControl        string        `json:"access_control"`
	AllowedNetworks      []string      `json:"allowed_networks"`
	AuthProfile          string        `json:"auth_profile"`
	Backend              []string      `json:"backend"`
	BePath               string        `json:"be_path"`
	Comment              string        `json:"comment"`
	DeniedNetworks       []interface{} `json:"denied_networks"`
	HotStandby           bool          `json:"hot_standby"`
	Name                 string        `json:"name"`
	Path                 string        `json:"path"`
	Status               bool          `json:"status"`
	StickysessionID      string        `json:"stickysession_id"`
	StickysessionStatus  bool          `json:"stickysession_status"`
	WebsocketPassthrough bool          `json:"websocket_passthrough"`
}

var _ sophos.RestGetter = &ReverseProxyLocation{}

// GetPath implements sophos.RestObject and returns the ReverseProxyLocations GET path
// Returns all available reverse_proxy/location objects
func (*ReverseProxyLocations) GetPath() string { return "/api/objects/reverse_proxy/location/" }

// RefRequired implements sophos.RestObject
func (*ReverseProxyLocations) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ReverseProxyLocations GET path
// Returns all available location types
func (r *ReverseProxyLocation) GetPath() string {
	return fmt.Sprintf("/api/objects/reverse_proxy/location/%s", r.Reference)
}

// RefRequired implements sophos.RestObject
func (r *ReverseProxyLocation) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the ReverseProxyLocation DELETE path
// Creates or updates the complete object location
func (*ReverseProxyLocation) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/location/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReverseProxyLocation PATCH path
// Changes to parts of the object location types
func (*ReverseProxyLocation) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/location/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReverseProxyLocation POST path
// Create a new reverse_proxy/location object
func (*ReverseProxyLocation) PostPath() string {
	return "/api/objects/reverse_proxy/location/"
}

// PutPath implements sophos.RestObject and returns the ReverseProxyLocation PUT path
// Creates or updates the complete object location
func (*ReverseProxyLocation) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/location/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReverseProxyLocation) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/location/%s/usedby", ref)
}

// GetType implements sophos.Object
func (r *ReverseProxyLocation) GetType() string { return r._type }

// ReverseProxyProfiles is an Sophos Endpoint subType and implements sophos.RestObject
type ReverseProxyProfiles []ReverseProxyProfile

// ReverseProxyProfile is a generated Sophos object
type ReverseProxyProfile struct {
	Locked                       string        `json:"_locked"`
	Reference                    string        `json:"_ref"`
	_type                        string        `json:"_type"`
	Av                           bool          `json:"av"`
	AvBlockUnscannable           bool          `json:"av_block_unscannable"`
	AvDirections                 string        `json:"av_directions"`
	AvEngines                    string        `json:"av_engines"`
	AvSizeLimit                  int64         `json:"av_size_limit"`
	AvTimeout                    int64         `json:"av_timeout"`
	BadClients                   bool          `json:"bad_clients"`
	BadClientsNoDnslookup        bool          `json:"bad_clients_no_dnslookup"`
	Comment                      string        `json:"comment"`
	Cookiesign                   bool          `json:"cookiesign"`
	CookiesignDropUnsigned       bool          `json:"cookiesign_drop_unsigned"`
	CustomThreatsFilters         []interface{} `json:"custom_threats_filters"`
	Extensions                   []interface{} `json:"extensions"`
	Filter                       []interface{} `json:"filter"`
	FilterMode                   string        `json:"filter_mode"`
	Formhardening                bool          `json:"formhardening"`
	Name                         string        `json:"name"`
	Outlookanywhere              bool          `json:"outlookanywhere"`
	SecRequestBodyNoFilesLimit   int64         `json:"sec_request_body_no_files_limit"`
	Skipwafrules                 []int64       `json:"skipwafrules"`
	Tft                          bool          `json:"tft"`
	TftBlockUnscannable          bool          `json:"tft_block_unscannable"`
	TftBlockedMimeTypes          []interface{} `json:"tft_blocked_mime_types"`
	ThreatsFilter                bool          `json:"threats_filter"`
	ThreatsFilterCategories      []string      `json:"threats_filter_categories"`
	ThreatsFilterRigid           bool          `json:"threats_filter_rigid"`
	Urlhardening                 bool          `json:"urlhardening"`
	UrlhardeningEntrypages       []string      `json:"urlhardening_entrypages"`
	UrlhardeningEntrypagesSource string        `json:"urlhardening_entrypages_source"`
	UrlhardeningSitemapUpdate    int64         `json:"urlhardening_sitemap_update"`
	UrlhardeningSitemapURL       string        `json:"urlhardening_sitemap_url"`
	Waf                          bool          `json:"waf"`
	Wafmode                      string        `json:"wafmode"`
	Wafparanoia                  bool          `json:"wafparanoia"`
}

var _ sophos.RestGetter = &ReverseProxyProfile{}

// GetPath implements sophos.RestObject and returns the ReverseProxyProfiles GET path
// Returns all available reverse_proxy/profile objects
func (*ReverseProxyProfiles) GetPath() string { return "/api/objects/reverse_proxy/profile/" }

// RefRequired implements sophos.RestObject
func (*ReverseProxyProfiles) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ReverseProxyProfiles GET path
// Returns all available profile types
func (r *ReverseProxyProfile) GetPath() string {
	return fmt.Sprintf("/api/objects/reverse_proxy/profile/%s", r.Reference)
}

// RefRequired implements sophos.RestObject
func (r *ReverseProxyProfile) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the ReverseProxyProfile DELETE path
// Creates or updates the complete object profile
func (*ReverseProxyProfile) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/profile/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReverseProxyProfile PATCH path
// Changes to parts of the object profile types
func (*ReverseProxyProfile) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/profile/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReverseProxyProfile POST path
// Create a new reverse_proxy/profile object
func (*ReverseProxyProfile) PostPath() string {
	return "/api/objects/reverse_proxy/profile/"
}

// PutPath implements sophos.RestObject and returns the ReverseProxyProfile PUT path
// Creates or updates the complete object profile
func (*ReverseProxyProfile) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/profile/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReverseProxyProfile) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/profile/%s/usedby", ref)
}

// GetType implements sophos.Object
func (r *ReverseProxyProfile) GetType() string { return r._type }

// ReverseProxyRedirection is an Sophos Endpoint subType and implements sophos.RestObject
type ReverseProxyRedirection []interface{}

var _ sophos.RestObject = &ReverseProxyRedirection{}

// GetPath implements sophos.RestObject and returns the ReverseProxyRedirection GET path
// Returns all available reverse_proxy/redirection objects
func (*ReverseProxyRedirection) GetPath() string { return "/api/objects/reverse_proxy/redirection/" }

// RefRequired implements sophos.RestObject
func (*ReverseProxyRedirection) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ReverseProxyRedirection DELETE path
// Creates or updates the complete object redirection
func (*ReverseProxyRedirection) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/redirection/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReverseProxyRedirection PATCH path
// Changes to parts of the object redirection types
func (*ReverseProxyRedirection) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/redirection/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReverseProxyRedirection POST path
// Create a new reverse_proxy/redirection object
func (*ReverseProxyRedirection) PostPath() string {
	return "/api/objects/reverse_proxy/redirection/"
}

// PutPath implements sophos.RestObject and returns the ReverseProxyRedirection PUT path
// Creates or updates the complete object redirection
func (*ReverseProxyRedirection) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/redirection/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReverseProxyRedirection) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/redirection/%s/usedby", ref)
}

// ReverseProxyThreatsFilter is an Sophos Endpoint subType and implements sophos.RestObject
type ReverseProxyThreatsFilter []interface{}

var _ sophos.RestObject = &ReverseProxyThreatsFilter{}

// GetPath implements sophos.RestObject and returns the ReverseProxyThreatsFilter GET path
// Returns all available reverse_proxy/threats_filter objects
func (*ReverseProxyThreatsFilter) GetPath() string {
	return "/api/objects/reverse_proxy/threats_filter/"
}

// RefRequired implements sophos.RestObject
func (*ReverseProxyThreatsFilter) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ReverseProxyThreatsFilter DELETE path
// Creates or updates the complete object threats_filter
func (*ReverseProxyThreatsFilter) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/threats_filter/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReverseProxyThreatsFilter PATCH path
// Changes to parts of the object threats_filter types
func (*ReverseProxyThreatsFilter) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/threats_filter/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReverseProxyThreatsFilter POST path
// Create a new reverse_proxy/threats_filter object
func (*ReverseProxyThreatsFilter) PostPath() string {
	return "/api/objects/reverse_proxy/threats_filter/"
}

// PutPath implements sophos.RestObject and returns the ReverseProxyThreatsFilter PUT path
// Creates or updates the complete object threats_filter
func (*ReverseProxyThreatsFilter) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/threats_filter/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReverseProxyThreatsFilter) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reverse_proxy/threats_filter/%s/usedby", ref)
}
