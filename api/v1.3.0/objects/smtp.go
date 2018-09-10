package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Smtp is a generated struct representing the Sophos Smtp Endpoint
// GET /api/nodes/smtp
type Smtp struct {
	AuthAaa                   []interface{} `json:"auth_aaa"`
	AuthStatus                int64         `json:"auth_status"`
	AvFooter                  string        `json:"av_footer"`
	AvFooterStatus            int64         `json:"av_footer_status"`
	BatvSecret                string        `json:"batv_secret"`
	BlockerService            string        `json:"blocker_service"`
	CffAsMarker               string        `json:"cff_as_marker"`
	DkimDomains               []interface{} `json:"dkim_domains"`
	DkimPrivateKey            string        `json:"dkim_private_key"`
	DkimSelector              string        `json:"dkim_selector"`
	EnableOldExpressionFilter int64         `json:"enable_old_expression_filter"`
	EncryptionUtility         string        `json:"encryption_utility"`
	Exceptions                []interface{} `json:"exceptions"`
	FootersMode               string        `json:"footers_mode"`
	GlobalAsReject            string        `json:"global_as_reject"`
	GlobalAvReject            int64         `json:"global_av_reject"`
	GlobalProfile             string        `json:"global_profile"`
	HostBlacklist             []interface{} `json:"host_blacklist"`
	Hostname                  string        `json:"hostname"`
	MaxMessageSize            int64         `json:"max_message_size"`
	Mode                      string        `json:"mode"`
	ParentProxyAuthPass       string        `json:"parent_proxy_auth_pass"`
	ParentProxyAuthStatus     int64         `json:"parent_proxy_auth_status"`
	ParentProxyAuthUser       string        `json:"parent_proxy_auth_user"`
	ParentProxyHost           string        `json:"parent_proxy_host"`
	ParentProxyPort           int64         `json:"parent_proxy_port"`
	ParentProxyStatus         int64         `json:"parent_proxy_status"`
	Postmaster                string        `json:"postmaster"`
	Profiles                  []interface{} `json:"profiles"`
	RecipientsMax             int64         `json:"recipients_max"`
	Relays                    []interface{} `json:"relays"`
	ScanOutgoingEmails        int64         `json:"scan_outgoing_emails"`
	ScannerPool               struct {
		MaxPool    int64   `json:"max_pool"`
		Thresholds []int64 `json:"thresholds"`
	} `json:"scanner_pool"`
	ScannerTimeout             int64         `json:"scanner_timeout"`
	SmarthostAuth              int64         `json:"smarthost_auth"`
	SmarthostHostname          string        `json:"smarthost_hostname"`
	SmarthostPass              string        `json:"smarthost_pass"`
	SmarthostPort              int64         `json:"smarthost_port"`
	SmarthostStatus            int64         `json:"smarthost_status"`
	SmarthostUser              string        `json:"smarthost_user"`
	SMTPAcceptMax              int64         `json:"smtp_accept_max"`
	SMTPAcceptPerConnectionMax int64         `json:"smtp_accept_per_connection_max"`
	SMTPAcceptPerHostMax       int64         `json:"smtp_accept_per_host_max"`
	Status                     int64         `json:"status"`
	TLSAvoid                   []interface{} `json:"tls_avoid"`
	TLSCert                    string        `json:"tls_cert"`
	TLSRequire                 []interface{} `json:"tls_require"`
	TLSRequireSenderDomains    []interface{} `json:"tls_require_sender_domains"`
	TLSVersion                 string        `json:"tls_version"`
	Transparent                int64         `json:"transparent"`
	TransparentSkip            []interface{} `json:"transparent_skip"`
	TransparentSkipAutoPf      int64         `json:"transparent_skip_auto_pf"`
	UpstreamHosts              []interface{} `json:"upstream_hosts"`
	UpstreamHostsOnly          int64         `json:"upstream_hosts_only"`
}

var _ sophos.Endpoint = &Smtp{}

var defsSmtp = map[string]sophos.RestObject{
	"SmtpException":       &SmtpException{},
	"SmtpGroup":           &SmtpGroup{},
	"SmtpHeaderOperation": &SmtpHeaderOperation{},
	"SmtpProfile":         &SmtpProfile{},
}

// RestObjects implements the sophos.Node interface and returns a map of Smtp's Objects
func (Smtp) RestObjects() map[string]sophos.RestObject { return defsSmtp }

// GetPath implements sophos.RestGetter
func (*Smtp) GetPath() string { return "/api/nodes/smtp" }

// RefRequired implements sophos.RestGetter
func (*Smtp) RefRequired() (string, bool) { return "", false }

var defSmtp = &sophos.Definition{Description: "smtp", Name: "smtp", Link: "/api/definitions/smtp"}

// Definition returns the /api/definitions struct of Smtp
func (Smtp) Definition() sophos.Definition { return *defSmtp }

// ApiRoutes returns all known Smtp Paths
func (Smtp) ApiRoutes() []string {
	return []string{
		"/api/objects/smtp/exception/",
		"/api/objects/smtp/exception/{ref}",
		"/api/objects/smtp/exception/{ref}/usedby",
		"/api/objects/smtp/group/",
		"/api/objects/smtp/group/{ref}",
		"/api/objects/smtp/group/{ref}/usedby",
		"/api/objects/smtp/header_operation/",
		"/api/objects/smtp/header_operation/{ref}",
		"/api/objects/smtp/header_operation/{ref}/usedby",
		"/api/objects/smtp/profile/",
		"/api/objects/smtp/profile/{ref}",
		"/api/objects/smtp/profile/{ref}/usedby",
	}
}

// References returns the Smtp's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Smtp) References() []string {
	return []string{
		"REF_SmtpException",
		"REF_SmtpGroup",
		"REF_SmtpHeaderOperation",
		"REF_SmtpProfile",
	}
}

// SmtpExceptions is an Sophos Endpoint subType and implements sophos.RestObject
type SmtpExceptions []SmtpException

// SmtpException represents a UTM SMTP filter exception
type SmtpException struct {
	Locked     string        `json:"_locked"`
	Reference  string        `json:"_ref"`
	_type      string        `json:"_type"`
	Name       string        `json:"name"`
	Networks   []interface{} `json:"networks"`
	Recipients []interface{} `json:"recipients"`
	Senders    []interface{} `json:"senders"`
	Skiplist   []interface{} `json:"skiplist"`
	// Status default value is false
	Status  bool   `json:"status"`
	Comment string `json:"comment"`
}

var _ sophos.RestGetter = &SmtpException{}

// GetPath implements sophos.RestObject and returns the SmtpExceptions GET path
// Returns all available smtp/exception objects
func (*SmtpExceptions) GetPath() string { return "/api/objects/smtp/exception/" }

// RefRequired implements sophos.RestObject
func (*SmtpExceptions) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SmtpExceptions GET path
// Returns all available exception types
func (s *SmtpException) GetPath() string {
	return fmt.Sprintf("/api/objects/smtp/exception/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *SmtpException) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SmtpException DELETE path
// Creates or updates the complete object exception
func (*SmtpException) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/exception/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SmtpException PATCH path
// Changes to parts of the object exception types
func (*SmtpException) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/exception/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SmtpException POST path
// Create a new smtp/exception object
func (*SmtpException) PostPath() string {
	return "/api/objects/smtp/exception/"
}

// PutPath implements sophos.RestObject and returns the SmtpException PUT path
// Creates or updates the complete object exception
func (*SmtpException) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/exception/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SmtpException) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/exception/%s/usedby", ref)
}

// SmtpGroups is an Sophos Endpoint subType and implements sophos.RestObject
type SmtpGroups []SmtpGroup

// SmtpGroup represents a UTM group
type SmtpGroup struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
}

var _ sophos.RestGetter = &SmtpGroup{}

// GetPath implements sophos.RestObject and returns the SmtpGroups GET path
// Returns all available smtp/group objects
func (*SmtpGroups) GetPath() string { return "/api/objects/smtp/group/" }

// RefRequired implements sophos.RestObject
func (*SmtpGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SmtpGroups GET path
// Returns all available group types
func (s *SmtpGroup) GetPath() string { return fmt.Sprintf("/api/objects/smtp/group/%s", s.Reference) }

// RefRequired implements sophos.RestObject
func (s *SmtpGroup) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SmtpGroup DELETE path
// Creates or updates the complete object group
func (*SmtpGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SmtpGroup PATCH path
// Changes to parts of the object group types
func (*SmtpGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SmtpGroup POST path
// Create a new smtp/group object
func (*SmtpGroup) PostPath() string {
	return "/api/objects/smtp/group/"
}

// PutPath implements sophos.RestObject and returns the SmtpGroup PUT path
// Creates or updates the complete object group
func (*SmtpGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SmtpGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/group/%s/usedby", ref)
}

// SmtpHeaderOperations is an Sophos Endpoint subType and implements sophos.RestObject
type SmtpHeaderOperations []SmtpHeaderOperation

// SmtpHeaderOperation represents a UTM Header modification operation
type SmtpHeaderOperation struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	_type      string `json:"_type"`
	Comment    string `json:"comment"`
	HeaderName string `json:"header_name"`
	Name       string `json:"name"`
	// Operation can be one of: []string{"add", "delete"}
	// Operation default value is "add"
	Operation string `json:"operation"`
	// Parameter default value is ""
	Parameter string `json:"parameter"`
}

var _ sophos.RestGetter = &SmtpHeaderOperation{}

// GetPath implements sophos.RestObject and returns the SmtpHeaderOperations GET path
// Returns all available smtp/header_operation objects
func (*SmtpHeaderOperations) GetPath() string { return "/api/objects/smtp/header_operation/" }

// RefRequired implements sophos.RestObject
func (*SmtpHeaderOperations) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SmtpHeaderOperations GET path
// Returns all available header_operation types
func (s *SmtpHeaderOperation) GetPath() string {
	return fmt.Sprintf("/api/objects/smtp/header_operation/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *SmtpHeaderOperation) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SmtpHeaderOperation DELETE path
// Creates or updates the complete object header_operation
func (*SmtpHeaderOperation) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/header_operation/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SmtpHeaderOperation PATCH path
// Changes to parts of the object header_operation types
func (*SmtpHeaderOperation) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/header_operation/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SmtpHeaderOperation POST path
// Create a new smtp/header_operation object
func (*SmtpHeaderOperation) PostPath() string {
	return "/api/objects/smtp/header_operation/"
}

// PutPath implements sophos.RestObject and returns the SmtpHeaderOperation PUT path
// Creates or updates the complete object header_operation
func (*SmtpHeaderOperation) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/header_operation/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SmtpHeaderOperation) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/header_operation/%s/usedby", ref)
}

// SmtpProfiles is an Sophos Endpoint subType and implements sophos.RestObject
type SmtpProfiles []SmtpProfile

// SmtpProfile is a generated Sophos object
type SmtpProfile struct {
	Locked                      string        `json:"_locked"`
	Reference                   string        `json:"_ref"`
	_type                       string        `json:"_type"`
	AdBaseDn                    string        `json:"ad_base_dn"`
	Batv                        bool          `json:"batv"`
	CffAv                       string        `json:"cff_av"`
	CffAvEngines                string        `json:"cff_av_engines"`
	CffFileExtensions           []string      `json:"cff_file_extensions"`
	Comment                     string        `json:"comment"`
	ConfidentialFooter          string        `json:"confidential_footer"`
	ConfidentialFooterStatus    bool          `json:"confidential_footer_status"`
	DlpAction                   string        `json:"dlp_action"`
	DlpCclRules                 []interface{} `json:"dlp_ccl_rules"`
	DlpCustomExpressions        []interface{} `json:"dlp_custom_expressions"`
	DlpNotificationAdmin        bool          `json:"dlp_notification_admin"`
	DlpNotificationOther        bool          `json:"dlp_notification_other"`
	DlpNotificationOtherAddress string        `json:"dlp_notification_other_address"`
	DlpNotificationSender       bool          `json:"dlp_notification_sender"`
	DlpScanAttachments          bool          `json:"dlp_scan_attachments"`
	Domains                     []interface{} `json:"domains"`
	GlobalAdd                   []interface{} `json:"global_add"`
	GlobalCopy                  []interface{} `json:"global_copy"`
	Greylisting                 bool          `json:"greylisting"`
	HeaderModification          []interface{} `json:"header_modification"`
	MimeAudio                   bool          `json:"mime_audio"`
	MimeBlacklist               []interface{} `json:"mime_blacklist"`
	MimeExecutable              bool          `json:"mime_executable"`
	MimeVideo                   bool          `json:"mime_video"`
	MimeWhitelist               []interface{} `json:"mime_whitelist"`
	Name                        string        `json:"name"`
	Rbl                         bool          `json:"rbl"`
	RblExtra                    []interface{} `json:"rbl_extra"`
	RcptAdServer                string        `json:"rcpt_ad_server"`
	RcptVerify                  string        `json:"rcpt_verify"`
	RdnsReject                  bool          `json:"rdns_reject"`
	RdnsRejectStrict            bool          `json:"rdns_reject_strict"`
	RouteList                   []interface{} `json:"route_list"`
	RouteTarget                 string        `json:"route_target"`
	RouteTargetPort             int64         `json:"route_target_port"`
	RouteTargetType             string        `json:"route_target_type"`
	SandboxMaxFilesizeMb        int64         `json:"sandbox_max_filesize_mb"`
	SandboxScanStatus           bool          `json:"sandbox_scan_status"`
	SenderBlacklist             []interface{} `json:"sender_blacklist"`
	Spam                        string        `json:"spam"`
	SpamExpressions             []interface{} `json:"spam_expressions"`
	Spamplus                    string        `json:"spamplus"`
	Spf                         bool          `json:"spf"`
	SpxTemplate                 string        `json:"spx_template"`
	Status                      bool          `json:"status"`
	Unscannable                 string        `json:"unscannable"`
}

var _ sophos.RestGetter = &SmtpProfile{}

// GetPath implements sophos.RestObject and returns the SmtpProfiles GET path
// Returns all available smtp/profile objects
func (*SmtpProfiles) GetPath() string { return "/api/objects/smtp/profile/" }

// RefRequired implements sophos.RestObject
func (*SmtpProfiles) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SmtpProfiles GET path
// Returns all available profile types
func (s *SmtpProfile) GetPath() string {
	return fmt.Sprintf("/api/objects/smtp/profile/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *SmtpProfile) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SmtpProfile DELETE path
// Creates or updates the complete object profile
func (*SmtpProfile) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/profile/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SmtpProfile PATCH path
// Changes to parts of the object profile types
func (*SmtpProfile) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/profile/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SmtpProfile POST path
// Create a new smtp/profile object
func (*SmtpProfile) PostPath() string {
	return "/api/objects/smtp/profile/"
}

// PutPath implements sophos.RestObject and returns the SmtpProfile PUT path
// Creates or updates the complete object profile
func (*SmtpProfile) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/profile/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SmtpProfile) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/smtp/profile/%s/usedby", ref)
}

// GetType implements sophos.Object
func (s *SmtpProfile) GetType() string { return s._type }
