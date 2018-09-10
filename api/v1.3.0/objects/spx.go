package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Spx is a generated struct representing the Sophos Spx Endpoint
// GET /api/nodes/spx
type Spx struct {
	Global struct {
		ErrorNotificationTarget string `json:"error_notification_target"`
		ExpirySettings          struct {
			AllowSecureReplyDays     int64 `json:"allow_secure_reply_days"`
			KeepDelayedMailDays      int64 `json:"keep_delayed_mail_days"`
			KeepUnusedPwdDays        int64 `json:"keep_unused_pwd_days"`
			RegistrationReminderDays int64 `json:"registration_reminder_days"`
		} `json:"expiry_settings"`
		PasswordStrength struct {
			MinLength        int64 `json:"min_length"`
			RequireSpecChars int64 `json:"require_spec_chars"`
		} `json:"password_strength"`
		PortalSettings struct {
			AllowedNetworks  []interface{} `json:"allowed_networks"`
			Hostname         string        `json:"hostname"`
			InterfaceAddress string        `json:"interface_address"`
			Port             int64         `json:"port"`
		} `json:"portal_settings"`
		PoweredByLogo string `json:"powered_by_logo"`
		SpxPriority   int64  `json:"spx_priority"`
		Status        int64  `json:"status"`
	} `json:"global"`
	SpxAuth struct {
		Password string `json:"password"`
		Port     int64  `json:"port"`
		Server   string `json:"server"`
		URL      string `json:"url"`
	} `json:"spx_auth"`
	Templates []string `json:"templates"`
}

var _ sophos.Endpoint = &Spx{}

var defsSpx = map[string]sophos.RestObject{
	"SpxGroup":    &SpxGroup{},
	"SpxTemplate": &SpxTemplate{},
}

// RestObjects implements the sophos.Node interface and returns a map of Spx's Objects
func (Spx) RestObjects() map[string]sophos.RestObject { return defsSpx }

// GetPath implements sophos.RestGetter
func (*Spx) GetPath() string { return "/api/nodes/spx" }

// RefRequired implements sophos.RestGetter
func (*Spx) RefRequired() (string, bool) { return "", false }

var defSpx = &sophos.Definition{Description: "spx", Name: "spx", Link: "/api/definitions/spx"}

// Definition returns the /api/definitions struct of Spx
func (Spx) Definition() sophos.Definition { return *defSpx }

// ApiRoutes returns all known Spx Paths
func (Spx) ApiRoutes() []string {
	return []string{
		"/api/objects/spx/group/",
		"/api/objects/spx/group/{ref}",
		"/api/objects/spx/group/{ref}/usedby",
		"/api/objects/spx/template/",
		"/api/objects/spx/template/{ref}",
		"/api/objects/spx/template/{ref}/usedby",
	}
}

// References returns the Spx's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Spx) References() []string {
	return []string{
		"REF_SpxGroup",
		"REF_SpxTemplate",
	}
}

// SpxGroups is an Sophos Endpoint subType and implements sophos.RestObject
type SpxGroups []SpxGroup

// SpxGroup represents a UTM group
type SpxGroup struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &SpxGroup{}

// GetPath implements sophos.RestObject and returns the SpxGroups GET path
// Returns all available spx/group objects
func (*SpxGroups) GetPath() string { return "/api/objects/spx/group/" }

// RefRequired implements sophos.RestObject
func (*SpxGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SpxGroups GET path
// Returns all available group types
func (s *SpxGroup) GetPath() string { return fmt.Sprintf("/api/objects/spx/group/%s", s.Reference) }

// RefRequired implements sophos.RestObject
func (s *SpxGroup) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SpxGroup DELETE path
// Creates or updates the complete object group
func (*SpxGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/spx/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SpxGroup PATCH path
// Changes to parts of the object group types
func (*SpxGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/spx/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SpxGroup POST path
// Create a new spx/group object
func (*SpxGroup) PostPath() string {
	return "/api/objects/spx/group/"
}

// PutPath implements sophos.RestObject and returns the SpxGroup PUT path
// Creates or updates the complete object group
func (*SpxGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/spx/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SpxGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/spx/group/%s/usedby", ref)
}

// SpxTemplates is an Sophos Endpoint subType and implements sophos.RestObject
type SpxTemplates []SpxTemplate

// SpxTemplate is a generated Sophos object
type SpxTemplate struct {
	Locked                         string `json:"_locked"`
	Reference                      string `json:"_ref"`
	ObjectType                     string `json:"_type"`
	Comment                        string `json:"comment"`
	Name                           string `json:"name"`
	OrgName                        string `json:"org_name"`
	PdfCoverPage                   string `json:"pdf_cover_page"`
	PdfCoverPageExt                string `json:"pdf_cover_page_ext"`
	PdfCoverPageFile               string `json:"pdf_cover_page_file"`
	PdfEncryption                  string `json:"pdf_encryption"`
	PdfLanguage                    string `json:"pdf_language"`
	PdfPageSize                    string `json:"pdf_page_size"`
	PortalAuthRequired             bool   `json:"portal_auth_required"`
	PortalFooterImage              string `json:"portal_footer_image"`
	PortalFooterImageExt           string `json:"portal_footer_image_ext"`
	PortalFooterImageFile          string `json:"portal_footer_image_file"`
	PortalHeaderImage              string `json:"portal_header_image"`
	PortalHeaderImageExt           string `json:"portal_header_image_ext"`
	PortalHeaderImageFile          string `json:"portal_header_image_file"`
	PortalIncludeOriginalBody      bool   `json:"portal_include_original_body"`
	PortalReplyAllButton           bool   `json:"portal_reply_all_button"`
	PortalSecureReplyEnabled       bool   `json:"portal_secure_reply_enabled"`
	PwdChallengeQuestionNumber     string `json:"pwd_challenge_question_number"`
	PwdFunctionChange              bool   `json:"pwd_function_change"`
	PwdFunctionRecover             bool   `json:"pwd_function_recover"`
	PwdFunctionReset               bool   `json:"pwd_function_reset"`
	PwdNotificationBody            string `json:"pwd_notification_body"`
	PwdNotificationFooterImage     string `json:"pwd_notification_footer_image"`
	PwdNotificationFooterImageExt  string `json:"pwd_notification_footer_image_ext"`
	PwdNotificationFooterImageFile string `json:"pwd_notification_footer_image_file"`
	PwdNotificationHeaderImage     string `json:"pwd_notification_header_image"`
	PwdNotificationHeaderImageExt  string `json:"pwd_notification_header_image_ext"`
	PwdNotificationHeaderImageFile string `json:"pwd_notification_header_image_file"`
	PwdNotificationSubject         string `json:"pwd_notification_subject"`
	PwdRcptNotificationBody        string `json:"pwd_rcpt_notification_body"`
	PwdRcptNotificationSubject     string `json:"pwd_rcpt_notification_subject"`
	PwdType                        string `json:"pwd_type"`
	RcptFooterImage                string `json:"rcpt_footer_image"`
	RcptFooterImageExt             string `json:"rcpt_footer_image_ext"`
	RcptFooterImageFile            string `json:"rcpt_footer_image_file"`
	RcptHeaderImage                string `json:"rcpt_header_image"`
	RcptHeaderImageExt             string `json:"rcpt_header_image_ext"`
	RcptHeaderImageFile            string `json:"rcpt_header_image_file"`
	RcptInstructions               string `json:"rcpt_instructions"`
	RemoveSophosLogo               bool   `json:"remove_sophos_logo"`
}

var _ sophos.RestGetter = &SpxTemplate{}

// GetPath implements sophos.RestObject and returns the SpxTemplates GET path
// Returns all available spx/template objects
func (*SpxTemplates) GetPath() string { return "/api/objects/spx/template/" }

// RefRequired implements sophos.RestObject
func (*SpxTemplates) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the SpxTemplates GET path
// Returns all available template types
func (s *SpxTemplate) GetPath() string {
	return fmt.Sprintf("/api/objects/spx/template/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *SpxTemplate) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the SpxTemplate DELETE path
// Creates or updates the complete object template
func (*SpxTemplate) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/spx/template/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the SpxTemplate PATCH path
// Changes to parts of the object template types
func (*SpxTemplate) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/spx/template/%s", ref)
}

// PostPath implements sophos.RestObject and returns the SpxTemplate POST path
// Create a new spx/template object
func (*SpxTemplate) PostPath() string {
	return "/api/objects/spx/template/"
}

// PutPath implements sophos.RestObject and returns the SpxTemplate PUT path
// Creates or updates the complete object template
func (*SpxTemplate) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/spx/template/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*SpxTemplate) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/spx/template/%s/usedby", ref)
}

// GetType implements sophos.Object
func (s *SpxTemplate) GetType() string { return s.ObjectType }
