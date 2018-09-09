package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Http is a generated struct representing the Sophos Http Endpoint
// GET /api/nodes/http
type Http struct {
	AdSsoInterfaces                       []interface{} `json:"ad_sso_interfaces"`
	AdssoRedirectUseHostname              int64         `json:"adsso_redirect_use_hostname"`
	AllowSsl3                             int64         `json:"allow_ssl3"`
	AllowTLS1_2                           int64         `json:"allow_tls_1_2"`
	AllowedPuas                           []interface{} `json:"allowed_puas"`
	AllowedTargetServices                 []string      `json:"allowed_target_services"`
	AuaMaxconns                           int64         `json:"aua_maxconns"`
	AuaTimeout                            int64         `json:"aua_timeout"`
	AuthCacheSize                         int64         `json:"auth_cache_size"`
	AuthCacheTTL                          int64         `json:"auth_cache_ttl"`
	AuthRealm                             string        `json:"auth_realm"`
	AuthUsercacheTTL                      int64         `json:"auth_usercache_ttl"`
	BlockUnscannable                      int64         `json:"block_unscannable"`
	BypassStreaming                       int64         `json:"bypass_streaming"`
	CaList                                []interface{} `json:"ca_list"`
	CacheIgnoresCookies                   int64         `json:"cache_ignores_cookies"`
	Cachessl                              int64         `json:"cachessl"`
	Caching                               int64         `json:"caching"`
	Certcache                             string        `json:"certcache"`
	Certstore                             string        `json:"certstore"`
	CffOverrideUsers                      []interface{} `json:"cff_override_users"`
	ClientTimeout                         int64         `json:"client_timeout"`
	ConfLockWorkaround                    int64         `json:"conf_lock_workaround"`
	ConnectTimeout                        int64         `json:"connect_timeout"`
	ConnectV6Timeout                      int64         `json:"connect_v6_timeout"`
	Connlimit                             int64         `json:"connlimit"`
	CtypeInspectBody                      int64         `json:"ctype_inspect_body"`
	CtypeUnpackArchive                    int64         `json:"ctype_unpack_archive"`
	Debug                                 []interface{} `json:"debug"`
	Defaultblockaction                    string        `json:"defaultblockaction"`
	Deferagents                           []string      `json:"deferagents"`
	Deferlength                           int64         `json:"deferlength"`
	DisplayHTTPBlockpageExplicitMode      int64         `json:"display_http_blockpage_explicit_mode"`
	DisplayIntro                          int64         `json:"display_intro"`
	DownloadManagerDefaultCharset         string        `json:"download_manager_default_charset"`
	EdirDelayBasicAuth                    int64         `json:"edir_delay_basic_auth"`
	EnableOutInterface                    int64         `json:"enable_out_interface"`
	EppQuotaAction                        string        `json:"epp_quota_action"`
	Exceptions                            []string      `json:"exceptions"`
	ForcedCachingExtension                []string      `json:"forced_caching_extension"`
	ForcedCachingNeverCachePrefix         []string      `json:"forced_caching_never_cache_prefix"`
	ForcedCachingStatus                   int64         `json:"forced_caching_status"`
	ForcedCachingTTL                      int64         `json:"forced_caching_ttl"`
	ForcedCachingUserAgentPrefix          []string      `json:"forced_caching_user_agent_prefix"`
	HTTPLoopbackDetect                    int64         `json:"http_loopback_detect"`
	IeSslBlockpageWorkaround              int64         `json:"ie_ssl_blockpage_workaround"`
	LimitAdSsoInterfaces                  int64         `json:"limit_ad_sso_interfaces"`
	LocalSiteList                         []interface{} `json:"local_site_list"`
	MaxContentEncoding                    int64         `json:"max_content_encoding"`
	MaxTempfileSize                       int64         `json:"max_tempfile_size"`
	Maxthreads                            int64         `json:"maxthreads"`
	MaxthreadsUnused                      int64         `json:"maxthreads_unused"`
	Modulepath                            string        `json:"modulepath"`
	Modules                               []string      `json:"modules"`
	Noscancontent                         []string      `json:"noscancontent"`
	OpendirectoryKeytab                   string        `json:"opendirectory_keytab"`
	PacFile                               string        `json:"pac_file"`
	ParentProxyHost                       string        `json:"parent_proxy_host"`
	ParentProxyPort                       int64         `json:"parent_proxy_port"`
	ParentProxyStatus                     int64         `json:"parent_proxy_status"`
	PassthroughID                         string        `json:"passthrough_id"`
	PharmingProtection                    int64         `json:"pharming_protection"`
	Port                                  int64         `json:"port"`
	PortalCert                            string        `json:"portal_cert"`
	PortalCertChain                       []interface{} `json:"portal_cert_chain"`
	PortalDomain                          string        `json:"portal_domain"`
	PortalHosts                           []interface{} `json:"portal_hosts"`
	PortalUseCert                         int64         `json:"portal_use_cert"`
	ProceedCacheTimeout                   int64         `json:"proceed_cache_timeout"`
	Profiles                              []string      `json:"profiles"`
	QuotaSliceTime                        int64         `json:"quota_slice_time"`
	RemoveRequest                         []interface{} `json:"remove_request"`
	RemoveResponse                        []string      `json:"remove_response"`
	ResponseTimeout                       int64         `json:"response_timeout"`
	ScLocalDB                             string        `json:"sc_local_db"`
	ScanEppTraffic                        int64         `json:"scan_epp_traffic"`
	Searchdomain                          string        `json:"searchdomain"`
	StrictHTTP                            int64         `json:"strict_http"`
	TlsciphersClient                      string        `json:"tlsciphers_client"`
	TlsciphersServer                      string        `json:"tlsciphers_server"`
	TmpfsUsageMinMemsize                  int64         `json:"tmpfs_usage_min_memsize"`
	TransparentAuthTimeout                int64         `json:"transparent_auth_timeout"`
	TransparentDstSkip                    []interface{} `json:"transparent_dst_skip"`
	TransparentSkipAutoPf                 int64         `json:"transparent_skip_auto_pf"`
	TransparentSrcSkip                    []interface{} `json:"transparent_src_skip"`
	TunnelTimeout                         int64         `json:"tunnel_timeout"`
	TunnelV6Timeout                       int64         `json:"tunnel_v6_timeout"`
	Undefercontent                        []string      `json:"undefercontent"`
	Undeferextension                      []string      `json:"undeferextension"`
	URLFilteringRedirectURL               string        `json:"url_filtering_redirect_url"`
	UseConnectionInsteadofProxyconnection int64         `json:"use_connection_insteadof_proxyconnection"`
	UseDstaddrForGeopiplookup             int64         `json:"use_dstaddr_for_geopiplookup"`
	UseKrb5Adsso                          int64         `json:"use_krb5_adsso"`
	UseSni                                int64         `json:"use_sni"`
	UseSxlUrid                            int64         `json:"use_sxl_urid"`
}

var _ sophos.Endpoint = &Http{}

var defsHttp = map[string]sophos.RestObject{
	"HttpCffAction":   &HttpCffAction{},
	"HttpCffProfile":  &HttpCffProfile{},
	"HttpDeviceAuth":  &HttpDeviceAuth{},
	"HttpDomainRegex": &HttpDomainRegex{},
	"HttpException":   &HttpException{},
	"HttpGroup":       &HttpGroup{},
	"HttpLocalSite":   &HttpLocalSite{},
	"HttpLslTag":      &HttpLslTag{},
	"HttpPacFile":     &HttpPacFile{},
	"HttpParentProxy": &HttpParentProxy{},
	"HttpProfile":     &HttpProfile{},
	"HttpSpCategory":  &HttpSpCategory{},
	"HttpSpSubcat":    &HttpSpSubcat{},
}

// RestObjects implements the sophos.Node interface and returns a map of Http's Objects
func (Http) RestObjects() map[string]sophos.RestObject { return defsHttp }

// GetPath implements sophos.RestGetter
func (*Http) GetPath() string { return "/api/nodes/http" }

// RefRequired implements sophos.RestGetter
func (*Http) RefRequired() (string, bool) { return "", false }

var defHttp = &sophos.Definition{Description: "http", Name: "http", Link: "/api/definitions/http"}

// Definition returns the /api/definitions struct of Http
func (Http) Definition() sophos.Definition { return *defHttp }

// ApiRoutes returns all known Http Paths
func (Http) ApiRoutes() []string {
	return []string{
		"/api/objects/http/cff_action/",
		"/api/objects/http/cff_action/{ref}",
		"/api/objects/http/cff_action/{ref}/usedby",
		"/api/objects/http/cff_profile/",
		"/api/objects/http/cff_profile/{ref}",
		"/api/objects/http/cff_profile/{ref}/usedby",
		"/api/objects/http/device_auth/",
		"/api/objects/http/device_auth/{ref}",
		"/api/objects/http/device_auth/{ref}/usedby",
		"/api/objects/http/domain_regex/",
		"/api/objects/http/domain_regex/{ref}",
		"/api/objects/http/domain_regex/{ref}/usedby",
		"/api/objects/http/exception/",
		"/api/objects/http/exception/{ref}",
		"/api/objects/http/exception/{ref}/usedby",
		"/api/objects/http/group/",
		"/api/objects/http/group/{ref}",
		"/api/objects/http/group/{ref}/usedby",
		"/api/objects/http/local_site/",
		"/api/objects/http/local_site/{ref}",
		"/api/objects/http/local_site/{ref}/usedby",
		"/api/objects/http/lsl_tag/",
		"/api/objects/http/lsl_tag/{ref}",
		"/api/objects/http/lsl_tag/{ref}/usedby",
		"/api/objects/http/pac_file/",
		"/api/objects/http/pac_file/{ref}",
		"/api/objects/http/pac_file/{ref}/usedby",
		"/api/objects/http/parent_proxy/",
		"/api/objects/http/parent_proxy/{ref}",
		"/api/objects/http/parent_proxy/{ref}/usedby",
		"/api/objects/http/profile/",
		"/api/objects/http/profile/{ref}",
		"/api/objects/http/profile/{ref}/usedby",
		"/api/objects/http/sp_category/",
		"/api/objects/http/sp_category/{ref}",
		"/api/objects/http/sp_category/{ref}/usedby",
		"/api/objects/http/sp_subcat/",
		"/api/objects/http/sp_subcat/{ref}",
		"/api/objects/http/sp_subcat/{ref}/usedby",
	}
}

// References returns the Http's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Http) References() []string {
	return []string{
		"REF_HttpCffAction",
		"REF_HttpCffProfile",
		"REF_HttpDeviceAuth",
		"REF_HttpDomainRegex",
		"REF_HttpException",
		"REF_HttpGroup",
		"REF_HttpLocalSite",
		"REF_HttpLslTag",
		"REF_HttpPacFile",
		"REF_HttpParentProxy",
		"REF_HttpProfile",
		"REF_HttpSpCategory",
		"REF_HttpSpSubcat",
	}
}

// HttpCffActions is an Sophos Endpoint subType and implements sophos.RestObject
type HttpCffActions []HttpCffAction

// HttpCffAction is a generated Sophos object
type HttpCffAction struct {
	Locked                   string        `json:"_locked"`
	Reference                string        `json:"_ref"`
	_type                    string        `json:"_type"`
	AllowTags                []interface{} `json:"allow_tags"`
	Av                       bool          `json:"av"`
	AvEngines                string        `json:"av_engines"`
	BingSafesearch           string        `json:"bing_safesearch"`
	BlockTags                []interface{} `json:"block_tags"`
	CheckMaxDownload         bool          `json:"check_max_download"`
	Comment                  string        `json:"comment"`
	ContenttypeBlacklist     []interface{} `json:"contenttype_blacklist"`
	ContenttypeBlacklistWarn []interface{} `json:"contenttype_blacklist_warn"`
	CreativeCommonsFilter    bool          `json:"creative_commons_filter"`
	EmbeddedRemoval          bool          `json:"embedded_removal"`
	Extensions               []string      `json:"extensions"`
	ExtensionsWarn           []interface{} `json:"extensions_warn"`
	GoogleSafesearch         string        `json:"google_safesearch"`
	Googleappdomains         []interface{} `json:"googleappdomains"`
	GoogleappdomainsEnabled  bool          `json:"googleappdomains_enabled"`
	LogAccess                bool          `json:"log_access"`
	LogBlocked               bool          `json:"log_blocked"`
	MaxDownloadSize          int64         `json:"max_download_size"`
	MaxFilesize              int64         `json:"max_filesize"`
	Mode                     string        `json:"mode"`
	Name                     string        `json:"name"`
	ParentProxies            []interface{} `json:"parent_proxies"`
	Pua                      bool          `json:"pua"`
	QuotaTags                []interface{} `json:"quota_tags"`
	QuotaTime                int64         `json:"quota_time"`
	Sandbox                  bool          `json:"sandbox"`
	ScriptRemoval            bool          `json:"script_removal"`
	SpCategories             []interface{} `json:"sp_categories"`
	SpCategoriesQuota        []interface{} `json:"sp_categories_quota"`
	SpCategoriesWarn         []interface{} `json:"sp_categories_warn"`
	SpMinreputation          string        `json:"sp_minreputation"`
	Spyware                  bool          `json:"spyware"`
	UncategorizedWebsites    string        `json:"uncategorized_websites"`
	URLBlacklist             []interface{} `json:"url_blacklist"`
	URLWhitelist             []interface{} `json:"url_whitelist"`
	WarnTags                 []interface{} `json:"warn_tags"`
	YahooSafesearch          string        `json:"yahoo_safesearch"`
}

var _ sophos.RestGetter = &HttpCffAction{}

// GetPath implements sophos.RestObject and returns the HttpCffActions GET path
// Returns all available http/cff_action objects
func (*HttpCffActions) GetPath() string { return "/api/objects/http/cff_action/" }

// RefRequired implements sophos.RestObject
func (*HttpCffActions) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the HttpCffActions GET path
// Returns all available cff_action types
func (h *HttpCffAction) GetPath() string {
	return fmt.Sprintf("/api/objects/http/cff_action/%s", h.Reference)
}

// RefRequired implements sophos.RestObject
func (h *HttpCffAction) RefRequired() (string, bool) { return h.Reference, true }

// DeletePath implements sophos.RestObject and returns the HttpCffAction DELETE path
// Creates or updates the complete object cff_action
func (*HttpCffAction) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/cff_action/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpCffAction PATCH path
// Changes to parts of the object cff_action types
func (*HttpCffAction) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/cff_action/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpCffAction POST path
// Create a new http/cff_action object
func (*HttpCffAction) PostPath() string {
	return "/api/objects/http/cff_action/"
}

// PutPath implements sophos.RestObject and returns the HttpCffAction PUT path
// Creates or updates the complete object cff_action
func (*HttpCffAction) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/cff_action/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpCffAction) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/cff_action/%s/usedby", ref)
}

// GetType implements sophos.Object
func (h *HttpCffAction) GetType() string { return h._type }

// HttpCffProfiles is an Sophos Endpoint subType and implements sophos.RestObject
type HttpCffProfiles []HttpCffProfile

// HttpCffProfile is a generated Sophos object
type HttpCffProfile struct {
	Locked         string   `json:"_locked"`
	Reference      string   `json:"_ref"`
	_type          string   `json:"_type"`
	Aaa            []string `json:"aaa"`
	Action         string   `json:"action"`
	CffProfileName string   `json:"cff_profile_name"`
	Comment        string   `json:"comment"`
	InProgress     string   `json:"in_progress"`
	Name           string   `json:"name"`
	SkipAuth       bool     `json:"skip_auth"`
	TimeEvent      string   `json:"time_event"`
}

var _ sophos.RestGetter = &HttpCffProfile{}

// GetPath implements sophos.RestObject and returns the HttpCffProfiles GET path
// Returns all available http/cff_profile objects
func (*HttpCffProfiles) GetPath() string { return "/api/objects/http/cff_profile/" }

// RefRequired implements sophos.RestObject
func (*HttpCffProfiles) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the HttpCffProfiles GET path
// Returns all available cff_profile types
func (h *HttpCffProfile) GetPath() string {
	return fmt.Sprintf("/api/objects/http/cff_profile/%s", h.Reference)
}

// RefRequired implements sophos.RestObject
func (h *HttpCffProfile) RefRequired() (string, bool) { return h.Reference, true }

// DeletePath implements sophos.RestObject and returns the HttpCffProfile DELETE path
// Creates or updates the complete object cff_profile
func (*HttpCffProfile) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/cff_profile/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpCffProfile PATCH path
// Changes to parts of the object cff_profile types
func (*HttpCffProfile) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/cff_profile/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpCffProfile POST path
// Create a new http/cff_profile object
func (*HttpCffProfile) PostPath() string {
	return "/api/objects/http/cff_profile/"
}

// PutPath implements sophos.RestObject and returns the HttpCffProfile PUT path
// Creates or updates the complete object cff_profile
func (*HttpCffProfile) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/cff_profile/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpCffProfile) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/cff_profile/%s/usedby", ref)
}

// GetType implements sophos.Object
func (h *HttpCffProfile) GetType() string { return h._type }

// HttpDeviceAuth is an Sophos Endpoint subType and implements sophos.RestObject
type HttpDeviceAuth []interface{}

var _ sophos.RestObject = &HttpDeviceAuth{}

// GetPath implements sophos.RestObject and returns the HttpDeviceAuth GET path
// Returns all available http/device_auth objects
func (*HttpDeviceAuth) GetPath() string { return "/api/objects/http/device_auth/" }

// RefRequired implements sophos.RestObject
func (*HttpDeviceAuth) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the HttpDeviceAuth DELETE path
// Creates or updates the complete object device_auth
func (*HttpDeviceAuth) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/device_auth/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpDeviceAuth PATCH path
// Changes to parts of the object device_auth types
func (*HttpDeviceAuth) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/device_auth/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpDeviceAuth POST path
// Create a new http/device_auth object
func (*HttpDeviceAuth) PostPath() string {
	return "/api/objects/http/device_auth/"
}

// PutPath implements sophos.RestObject and returns the HttpDeviceAuth PUT path
// Creates or updates the complete object device_auth
func (*HttpDeviceAuth) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/device_auth/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpDeviceAuth) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/device_auth/%s/usedby", ref)
}

// HttpDomainRegex is an Sophos Endpoint subType and implements sophos.RestObject
type HttpDomainRegex []interface{}

var _ sophos.RestObject = &HttpDomainRegex{}

// GetPath implements sophos.RestObject and returns the HttpDomainRegex GET path
// Returns all available http/domain_regex objects
func (*HttpDomainRegex) GetPath() string { return "/api/objects/http/domain_regex/" }

// RefRequired implements sophos.RestObject
func (*HttpDomainRegex) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the HttpDomainRegex DELETE path
// Creates or updates the complete object domain_regex
func (*HttpDomainRegex) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/domain_regex/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpDomainRegex PATCH path
// Changes to parts of the object domain_regex types
func (*HttpDomainRegex) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/domain_regex/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpDomainRegex POST path
// Create a new http/domain_regex object
func (*HttpDomainRegex) PostPath() string {
	return "/api/objects/http/domain_regex/"
}

// PutPath implements sophos.RestObject and returns the HttpDomainRegex PUT path
// Creates or updates the complete object domain_regex
func (*HttpDomainRegex) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/domain_regex/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpDomainRegex) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/domain_regex/%s/usedby", ref)
}

// HttpExceptions is an Sophos Endpoint subType and implements sophos.RestObject
type HttpExceptions []HttpException

// HttpException is a generated Sophos object
type HttpException struct {
	Locked          string        `json:"_locked"`
	Reference       string        `json:"_ref"`
	_type           string        `json:"_type"`
	Aaa             []interface{} `json:"aaa"`
	Comment         string        `json:"comment"`
	Domains         []string      `json:"domains"`
	EndpointsGroups []interface{} `json:"endpoints_groups"`
	Name            string        `json:"name"`
	Networks        []interface{} `json:"networks"`
	Operator        string        `json:"operator"`
	Skiplist        []string      `json:"skiplist"`
	SpCategories    []interface{} `json:"sp_categories"`
	Status          bool          `json:"status"`
	Tags            []interface{} `json:"tags"`
	UserAgents      []interface{} `json:"user_agents"`
}

var _ sophos.RestGetter = &HttpException{}

// GetPath implements sophos.RestObject and returns the HttpExceptions GET path
// Returns all available http/exception objects
func (*HttpExceptions) GetPath() string { return "/api/objects/http/exception/" }

// RefRequired implements sophos.RestObject
func (*HttpExceptions) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the HttpExceptions GET path
// Returns all available exception types
func (h *HttpException) GetPath() string {
	return fmt.Sprintf("/api/objects/http/exception/%s", h.Reference)
}

// RefRequired implements sophos.RestObject
func (h *HttpException) RefRequired() (string, bool) { return h.Reference, true }

// DeletePath implements sophos.RestObject and returns the HttpException DELETE path
// Creates or updates the complete object exception
func (*HttpException) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/exception/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpException PATCH path
// Changes to parts of the object exception types
func (*HttpException) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/exception/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpException POST path
// Create a new http/exception object
func (*HttpException) PostPath() string {
	return "/api/objects/http/exception/"
}

// PutPath implements sophos.RestObject and returns the HttpException PUT path
// Creates or updates the complete object exception
func (*HttpException) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/exception/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpException) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/exception/%s/usedby", ref)
}

// GetType implements sophos.Object
func (h *HttpException) GetType() string { return h._type }

// HttpGroup is an Sophos Endpoint subType and implements sophos.RestObject
type HttpGroup []interface{}

var _ sophos.RestObject = &HttpGroup{}

// GetPath implements sophos.RestObject and returns the HttpGroup GET path
// Returns all available http/group objects
func (*HttpGroup) GetPath() string { return "/api/objects/http/group/" }

// RefRequired implements sophos.RestObject
func (*HttpGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the HttpGroup DELETE path
// Creates or updates the complete object group
func (*HttpGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpGroup PATCH path
// Changes to parts of the object group types
func (*HttpGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpGroup POST path
// Create a new http/group object
func (*HttpGroup) PostPath() string {
	return "/api/objects/http/group/"
}

// PutPath implements sophos.RestObject and returns the HttpGroup PUT path
// Creates or updates the complete object group
func (*HttpGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/group/%s/usedby", ref)
}

// HttpLocalSite is an Sophos Endpoint subType and implements sophos.RestObject
type HttpLocalSite []interface{}

var _ sophos.RestObject = &HttpLocalSite{}

// GetPath implements sophos.RestObject and returns the HttpLocalSite GET path
// Returns all available http/local_site objects
func (*HttpLocalSite) GetPath() string { return "/api/objects/http/local_site/" }

// RefRequired implements sophos.RestObject
func (*HttpLocalSite) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the HttpLocalSite DELETE path
// Creates or updates the complete object local_site
func (*HttpLocalSite) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/local_site/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpLocalSite PATCH path
// Changes to parts of the object local_site types
func (*HttpLocalSite) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/local_site/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpLocalSite POST path
// Create a new http/local_site object
func (*HttpLocalSite) PostPath() string {
	return "/api/objects/http/local_site/"
}

// PutPath implements sophos.RestObject and returns the HttpLocalSite PUT path
// Creates or updates the complete object local_site
func (*HttpLocalSite) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/local_site/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpLocalSite) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/local_site/%s/usedby", ref)
}

// HttpLslTag is an Sophos Endpoint subType and implements sophos.RestObject
type HttpLslTag []interface{}

var _ sophos.RestObject = &HttpLslTag{}

// GetPath implements sophos.RestObject and returns the HttpLslTag GET path
// Returns all available http/lsl_tag objects
func (*HttpLslTag) GetPath() string { return "/api/objects/http/lsl_tag/" }

// RefRequired implements sophos.RestObject
func (*HttpLslTag) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the HttpLslTag DELETE path
// Creates or updates the complete object lsl_tag
func (*HttpLslTag) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/lsl_tag/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpLslTag PATCH path
// Changes to parts of the object lsl_tag types
func (*HttpLslTag) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/lsl_tag/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpLslTag POST path
// Create a new http/lsl_tag object
func (*HttpLslTag) PostPath() string {
	return "/api/objects/http/lsl_tag/"
}

// PutPath implements sophos.RestObject and returns the HttpLslTag PUT path
// Creates or updates the complete object lsl_tag
func (*HttpLslTag) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/lsl_tag/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpLslTag) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/lsl_tag/%s/usedby", ref)
}

// HttpPacFiles is an Sophos Endpoint subType and implements sophos.RestObject
type HttpPacFiles []HttpPacFile

// HttpPacFile is a generated Sophos object
type HttpPacFile struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Content   string `json:"content"`
	Name      string `json:"name"`
	Status    bool   `json:"status"`
}

var _ sophos.RestGetter = &HttpPacFile{}

// GetPath implements sophos.RestObject and returns the HttpPacFiles GET path
// Returns all available http/pac_file objects
func (*HttpPacFiles) GetPath() string { return "/api/objects/http/pac_file/" }

// RefRequired implements sophos.RestObject
func (*HttpPacFiles) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the HttpPacFiles GET path
// Returns all available pac_file types
func (h *HttpPacFile) GetPath() string {
	return fmt.Sprintf("/api/objects/http/pac_file/%s", h.Reference)
}

// RefRequired implements sophos.RestObject
func (h *HttpPacFile) RefRequired() (string, bool) { return h.Reference, true }

// DeletePath implements sophos.RestObject and returns the HttpPacFile DELETE path
// Creates or updates the complete object pac_file
func (*HttpPacFile) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/pac_file/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpPacFile PATCH path
// Changes to parts of the object pac_file types
func (*HttpPacFile) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/pac_file/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpPacFile POST path
// Create a new http/pac_file object
func (*HttpPacFile) PostPath() string {
	return "/api/objects/http/pac_file/"
}

// PutPath implements sophos.RestObject and returns the HttpPacFile PUT path
// Creates or updates the complete object pac_file
func (*HttpPacFile) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/pac_file/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpPacFile) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/pac_file/%s/usedby", ref)
}

// GetType implements sophos.Object
func (h *HttpPacFile) GetType() string { return h._type }

// HttpParentProxy is an Sophos Endpoint subType and implements sophos.RestObject
type HttpParentProxy []interface{}

var _ sophos.RestObject = &HttpParentProxy{}

// GetPath implements sophos.RestObject and returns the HttpParentProxy GET path
// Returns all available http/parent_proxy objects
func (*HttpParentProxy) GetPath() string { return "/api/objects/http/parent_proxy/" }

// RefRequired implements sophos.RestObject
func (*HttpParentProxy) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the HttpParentProxy DELETE path
// Creates or updates the complete object parent_proxy
func (*HttpParentProxy) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/parent_proxy/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpParentProxy PATCH path
// Changes to parts of the object parent_proxy types
func (*HttpParentProxy) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/parent_proxy/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpParentProxy POST path
// Create a new http/parent_proxy object
func (*HttpParentProxy) PostPath() string {
	return "/api/objects/http/parent_proxy/"
}

// PutPath implements sophos.RestObject and returns the HttpParentProxy PUT path
// Creates or updates the complete object parent_proxy
func (*HttpParentProxy) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/parent_proxy/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpParentProxy) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/parent_proxy/%s/usedby", ref)
}

// HttpProfiles is an Sophos Endpoint subType and implements sophos.RestObject
type HttpProfiles []HttpProfile

// HttpProfile is a generated Sophos object
type HttpProfile struct {
	Locked             string        `json:"_locked"`
	Reference          string        `json:"_ref"`
	_type              string        `json:"_type"`
	Aua                bool          `json:"aua"`
	BlockOnAuthFailed  bool          `json:"block_on_auth_failed"`
	CffProfiles        []string      `json:"cff_profiles"`
	Comment            string        `json:"comment"`
	DefaultCffAction   string        `json:"default_cff_action"`
	DeviceAuth         []interface{} `json:"device_auth"`
	EdirSso            bool          `json:"edir_sso"`
	EnableDeviceAuth   bool          `json:"enable_device_auth"`
	EndpointsGroups    []interface{} `json:"endpoints_groups"`
	FullTransparent    bool          `json:"full_transparent"`
	InProgress         bool          `json:"in_progress"`
	Name               string        `json:"name"`
	Networks           []interface{} `json:"networks"`
	Ntlm               bool          `json:"ntlm"`
	OpendirectoryAuth  bool          `json:"opendirectory_auth"`
	OrderedCffProfiles []string      `json:"ordered_cff_profiles"`
	OutInterface       string        `json:"out_interface"`
	ScanSslOpt         string        `json:"scan_ssl_opt"`
	SelectiveScanCat   []string      `json:"selective_scan_cat"`
	SelectiveScanTags  []interface{} `json:"selective_scan_tags"`
	Status             bool          `json:"status"`
	Transparent        bool          `json:"transparent"`
	TransparentAac     bool          `json:"transparent_aac"`
	TransparentAuth    bool          `json:"transparent_auth"`
}

var _ sophos.RestGetter = &HttpProfile{}

// GetPath implements sophos.RestObject and returns the HttpProfiles GET path
// Returns all available http/profile objects
func (*HttpProfiles) GetPath() string { return "/api/objects/http/profile/" }

// RefRequired implements sophos.RestObject
func (*HttpProfiles) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the HttpProfiles GET path
// Returns all available profile types
func (h *HttpProfile) GetPath() string {
	return fmt.Sprintf("/api/objects/http/profile/%s", h.Reference)
}

// RefRequired implements sophos.RestObject
func (h *HttpProfile) RefRequired() (string, bool) { return h.Reference, true }

// DeletePath implements sophos.RestObject and returns the HttpProfile DELETE path
// Creates or updates the complete object profile
func (*HttpProfile) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/profile/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpProfile PATCH path
// Changes to parts of the object profile types
func (*HttpProfile) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/profile/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpProfile POST path
// Create a new http/profile object
func (*HttpProfile) PostPath() string {
	return "/api/objects/http/profile/"
}

// PutPath implements sophos.RestObject and returns the HttpProfile PUT path
// Creates or updates the complete object profile
func (*HttpProfile) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/profile/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpProfile) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/profile/%s/usedby", ref)
}

// GetType implements sophos.Object
func (h *HttpProfile) GetType() string { return h._type }

// HttpSpCategorys is an Sophos Endpoint subType and implements sophos.RestObject
type HttpSpCategorys []HttpSpCategory

// HttpSpCategory is a generated Sophos object
type HttpSpCategory struct {
	Locked    string   `json:"_locked"`
	Reference string   `json:"_ref"`
	_type     string   `json:"_type"`
	Comment   string   `json:"comment"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Subcats   []string `json:"subcats"`
}

var _ sophos.RestGetter = &HttpSpCategory{}

// GetPath implements sophos.RestObject and returns the HttpSpCategorys GET path
// Returns all available http/sp_category objects
func (*HttpSpCategorys) GetPath() string { return "/api/objects/http/sp_category/" }

// RefRequired implements sophos.RestObject
func (*HttpSpCategorys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the HttpSpCategorys GET path
// Returns all available sp_category types
func (h *HttpSpCategory) GetPath() string {
	return fmt.Sprintf("/api/objects/http/sp_category/%s", h.Reference)
}

// RefRequired implements sophos.RestObject
func (h *HttpSpCategory) RefRequired() (string, bool) { return h.Reference, true }

// DeletePath implements sophos.RestObject and returns the HttpSpCategory DELETE path
// Creates or updates the complete object sp_category
func (*HttpSpCategory) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/sp_category/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpSpCategory PATCH path
// Changes to parts of the object sp_category types
func (*HttpSpCategory) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/sp_category/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpSpCategory POST path
// Create a new http/sp_category object
func (*HttpSpCategory) PostPath() string {
	return "/api/objects/http/sp_category/"
}

// PutPath implements sophos.RestObject and returns the HttpSpCategory PUT path
// Creates or updates the complete object sp_category
func (*HttpSpCategory) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/sp_category/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpSpCategory) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/sp_category/%s/usedby", ref)
}

// GetType implements sophos.Object
func (h *HttpSpCategory) GetType() string { return h._type }

// HttpSpSubcats is an Sophos Endpoint subType and implements sophos.RestObject
type HttpSpSubcats []HttpSpSubcat

// HttpSpSubcat is a generated Sophos object
type HttpSpSubcat struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	ID        string `json:"id"`
	Name      string `json:"name"`
}

var _ sophos.RestGetter = &HttpSpSubcat{}

// GetPath implements sophos.RestObject and returns the HttpSpSubcats GET path
// Returns all available http/sp_subcat objects
func (*HttpSpSubcats) GetPath() string { return "/api/objects/http/sp_subcat/" }

// RefRequired implements sophos.RestObject
func (*HttpSpSubcats) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the HttpSpSubcats GET path
// Returns all available sp_subcat types
func (h *HttpSpSubcat) GetPath() string {
	return fmt.Sprintf("/api/objects/http/sp_subcat/%s", h.Reference)
}

// RefRequired implements sophos.RestObject
func (h *HttpSpSubcat) RefRequired() (string, bool) { return h.Reference, true }

// DeletePath implements sophos.RestObject and returns the HttpSpSubcat DELETE path
// Creates or updates the complete object sp_subcat
func (*HttpSpSubcat) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/http/sp_subcat/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the HttpSpSubcat PATCH path
// Changes to parts of the object sp_subcat types
func (*HttpSpSubcat) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/sp_subcat/%s", ref)
}

// PostPath implements sophos.RestObject and returns the HttpSpSubcat POST path
// Create a new http/sp_subcat object
func (*HttpSpSubcat) PostPath() string {
	return "/api/objects/http/sp_subcat/"
}

// PutPath implements sophos.RestObject and returns the HttpSpSubcat PUT path
// Creates or updates the complete object sp_subcat
func (*HttpSpSubcat) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/sp_subcat/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*HttpSpSubcat) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/http/sp_subcat/%s/usedby", ref)
}

// GetType implements sophos.Object
func (h *HttpSpSubcat) GetType() string { return h._type }
