package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Reporting is a generated struct representing the Sophos Reporting Endpoint
// GET /api/nodes/reporting
type Reporting struct {
	AccountingKeepdays     int64         `json:"accounting_keepdays"`
	AccountingStatus       int64         `json:"accounting_status"`
	Anonymizing            int64         `json:"anonymizing"`
	AppctrlKeepdays        int64         `json:"appctrl_keepdays"`
	AppctrlStatus          int64         `json:"appctrl_status"`
	AtpKeepdays            int64         `json:"atp_keepdays"`
	AtpReset               int64         `json:"atp_reset"`
	AtpStatus              int64         `json:"atp_status"`
	AuthenticationKeepdays int64         `json:"authentication_keepdays"`
	AuthenticationStatus   int64         `json:"authentication_status"`
	CsvSeparator           string        `json:"csv_separator"`
	EmailsecurityImport    []interface{} `json:"emailsecurity_import"`
	EmailsecurityKeepdays  int64         `json:"emailsecurity_keepdays"`
	EmailsecurityStatus    int64         `json:"emailsecurity_status"`
	EnableVpnAccounting    int64         `json:"enable_vpn_accounting"`
	HideAccountingips      []interface{} `json:"hide_accountingips"`
	HideMailaddresses      []interface{} `json:"hide_mailaddresses"`
	HideMaildomains        []interface{} `json:"hide_maildomains"`
	HideNetsecips          []interface{} `json:"hide_netsecips"`
	HideWebdomains         []interface{} `json:"hide_webdomains"`
	IpsImport              []interface{} `json:"ips_import"`
	IpsKeepdays            int64         `json:"ips_keepdays"`
	IpsStatus              int64         `json:"ips_status"`
	PacketfilterImport     []interface{} `json:"packetfilter_import"`
	PacketfilterKeepdays   int64         `json:"packetfilter_keepdays"`
	PacketfilterStatus     int64         `json:"packetfilter_status"`
	Password1              string        `json:"password1"`
	Password2              string        `json:"password2"`
	SandboxKeepdays        int64         `json:"sandbox_keepdays"`
	UserlogFromLogs        int64         `json:"userlog_from_logs"`
	VpnKeepdays            int64         `json:"vpn_keepdays"`
	VpnStatus              int64         `json:"vpn_status"`
	WafKeepdays            int64         `json:"waf_keepdays"`
	WafStatus              int64         `json:"waf_status"`
	WebsecurityDetail      int64         `json:"websecurity_detail"`
	WebsecurityImport      []interface{} `json:"websecurity_import"`
	WebsecurityKeepdays    int64         `json:"websecurity_keepdays"`
	WebsecurityStatus      int64         `json:"websecurity_status"`
}

var _ sophos.Endpoint = &Reporting{}

var defsReporting = map[string]sophos.RestObject{
	"ReportingDepartment": &ReportingDepartment{},
	"ReportingFilter":     &ReportingFilter{},
	"ReportingGroup":      &ReportingGroup{},
	"ReportingMail":       &ReportingMail{},
}

// RestObjects implements the sophos.Node interface and returns a map of Reporting's Objects
func (Reporting) RestObjects() map[string]sophos.RestObject { return defsReporting }

// GetPath implements sophos.RestGetter
func (*Reporting) GetPath() string { return "/api/nodes/reporting" }

// RefRequired implements sophos.RestGetter
func (*Reporting) RefRequired() (string, bool) { return "", false }

var defReporting = &sophos.Definition{Description: "reporting", Name: "reporting", Link: "/api/definitions/reporting"}

// Definition returns the /api/definitions struct of Reporting
func (Reporting) Definition() sophos.Definition { return *defReporting }

// ApiRoutes returns all known Reporting Paths
func (Reporting) ApiRoutes() []string {
	return []string{
		"/api/objects/reporting/department/",
		"/api/objects/reporting/department/{ref}",
		"/api/objects/reporting/department/{ref}/usedby",
		"/api/objects/reporting/filter/",
		"/api/objects/reporting/filter/{ref}",
		"/api/objects/reporting/filter/{ref}/usedby",
		"/api/objects/reporting/group/",
		"/api/objects/reporting/group/{ref}",
		"/api/objects/reporting/group/{ref}/usedby",
		"/api/objects/reporting/mail/",
		"/api/objects/reporting/mail/{ref}",
		"/api/objects/reporting/mail/{ref}/usedby",
	}
}

// References returns the Reporting's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Reporting) References() []string {
	return []string{
		"REF_ReportingDepartment",
		"REF_ReportingFilter",
		"REF_ReportingGroup",
		"REF_ReportingMail",
	}
}

// ReportingDepartments is an Sophos Endpoint subType and implements sophos.RestObject
type ReportingDepartments []ReportingDepartment

// ReportingDepartment represents a UTM department
type ReportingDepartment struct {
	Locked     string        `json:"_locked"`
	Reference  string        `json:"_ref"`
	ObjectType string        `json:"_type"`
	Comment    string        `json:"comment"`
	Name       string        `json:"name"`
	Networks   []interface{} `json:"networks"`
	Users      []interface{} `json:"users"`
}

var _ sophos.RestGetter = &ReportingDepartment{}

// GetPath implements sophos.RestObject and returns the ReportingDepartments GET path
// Returns all available reporting/department objects
func (*ReportingDepartments) GetPath() string { return "/api/objects/reporting/department/" }

// RefRequired implements sophos.RestObject
func (*ReportingDepartments) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ReportingDepartments GET path
// Returns all available department types
func (r *ReportingDepartment) GetPath() string {
	return fmt.Sprintf("/api/objects/reporting/department/%s", r.Reference)
}

// RefRequired implements sophos.RestObject
func (r *ReportingDepartment) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the ReportingDepartment DELETE path
// Creates or updates the complete object department
func (*ReportingDepartment) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/department/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReportingDepartment PATCH path
// Changes to parts of the object department types
func (*ReportingDepartment) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/department/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReportingDepartment POST path
// Create a new reporting/department object
func (*ReportingDepartment) PostPath() string {
	return "/api/objects/reporting/department/"
}

// PutPath implements sophos.RestObject and returns the ReportingDepartment PUT path
// Creates or updates the complete object department
func (*ReportingDepartment) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/department/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReportingDepartment) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/department/%s/usedby", ref)
}

// ReportingFilters is an Sophos Endpoint subType and implements sophos.RestObject
type ReportingFilters []ReportingFilter

// ReportingFilter is a generated Sophos object
type ReportingFilter struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Callname   string `json:"callname"`
	Comment    string `json:"comment"`
	Department string `json:"department"`
	Filter     struct {
		Action string `json:"action"`
		Info   string `json:"info"`
	} `json:"filter"`
	Name      string `json:"name"`
	SortBy    string `json:"sort_by"`
	SortDir   string `json:"sort_dir"`
	Subsystem string `json:"subsystem"`
	Timeframe string `json:"timeframe"`
	Top       int64  `json:"top"`
}

var _ sophos.RestGetter = &ReportingFilter{}

// GetPath implements sophos.RestObject and returns the ReportingFilters GET path
// Returns all available reporting/filter objects
func (*ReportingFilters) GetPath() string { return "/api/objects/reporting/filter/" }

// RefRequired implements sophos.RestObject
func (*ReportingFilters) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ReportingFilters GET path
// Returns all available filter types
func (r *ReportingFilter) GetPath() string {
	return fmt.Sprintf("/api/objects/reporting/filter/%s", r.Reference)
}

// RefRequired implements sophos.RestObject
func (r *ReportingFilter) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the ReportingFilter DELETE path
// Creates or updates the complete object filter
func (*ReportingFilter) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/filter/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReportingFilter PATCH path
// Changes to parts of the object filter types
func (*ReportingFilter) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/filter/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReportingFilter POST path
// Create a new reporting/filter object
func (*ReportingFilter) PostPath() string {
	return "/api/objects/reporting/filter/"
}

// PutPath implements sophos.RestObject and returns the ReportingFilter PUT path
// Creates or updates the complete object filter
func (*ReportingFilter) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/filter/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReportingFilter) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/filter/%s/usedby", ref)
}

// GetType implements sophos.Object
func (r *ReportingFilter) GetType() string { return r.ObjectType }

// ReportingGroups is an Sophos Endpoint subType and implements sophos.RestObject
type ReportingGroups []ReportingGroup

// ReportingGroup represents a UTM group
type ReportingGroup struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &ReportingGroup{}

// GetPath implements sophos.RestObject and returns the ReportingGroups GET path
// Returns all available reporting/group objects
func (*ReportingGroups) GetPath() string { return "/api/objects/reporting/group/" }

// RefRequired implements sophos.RestObject
func (*ReportingGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ReportingGroups GET path
// Returns all available group types
func (r *ReportingGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/reporting/group/%s", r.Reference)
}

// RefRequired implements sophos.RestObject
func (r *ReportingGroup) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the ReportingGroup DELETE path
// Creates or updates the complete object group
func (*ReportingGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReportingGroup PATCH path
// Changes to parts of the object group types
func (*ReportingGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReportingGroup POST path
// Create a new reporting/group object
func (*ReportingGroup) PostPath() string {
	return "/api/objects/reporting/group/"
}

// PutPath implements sophos.RestObject and returns the ReportingGroup PUT path
// Creates or updates the complete object group
func (*ReportingGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReportingGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/group/%s/usedby", ref)
}

// ReportingMails is an Sophos Endpoint subType and implements sophos.RestObject
type ReportingMails []ReportingMail

// ReportingMail represents a UTM scheduled report
type ReportingMail struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	// Status default value is false
	Status  bool          `json:"status"`
	Comment string        `json:"comment"`
	Filters []interface{} `json:"filters"`
	// Interval can be one of: []string{"daily", "weekly", "monthly"}
	// Interval default value is "daily"
	Interval   string        `json:"interval"`
	Name       string        `json:"name"`
	Recipients []interface{} `json:"recipients"`
}

var _ sophos.RestGetter = &ReportingMail{}

// GetPath implements sophos.RestObject and returns the ReportingMails GET path
// Returns all available reporting/mail objects
func (*ReportingMails) GetPath() string { return "/api/objects/reporting/mail/" }

// RefRequired implements sophos.RestObject
func (*ReportingMails) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ReportingMails GET path
// Returns all available mail types
func (r *ReportingMail) GetPath() string {
	return fmt.Sprintf("/api/objects/reporting/mail/%s", r.Reference)
}

// RefRequired implements sophos.RestObject
func (r *ReportingMail) RefRequired() (string, bool) { return r.Reference, true }

// DeletePath implements sophos.RestObject and returns the ReportingMail DELETE path
// Creates or updates the complete object mail
func (*ReportingMail) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/mail/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ReportingMail PATCH path
// Changes to parts of the object mail types
func (*ReportingMail) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/mail/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ReportingMail POST path
// Create a new reporting/mail object
func (*ReportingMail) PostPath() string {
	return "/api/objects/reporting/mail/"
}

// PutPath implements sophos.RestObject and returns the ReportingMail PUT path
// Creates or updates the complete object mail
func (*ReportingMail) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/mail/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*ReportingMail) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/reporting/mail/%s/usedby", ref)
}
