package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Emailpki is a generated struct representing the Sophos Emailpki Endpoint
// GET /api/nodes/emailpki
type Emailpki struct {
	Authority struct {
		Cert                  string `json:"cert"`
		Fingerprint           string `json:"fingerprint"`
		Key                   string `json:"key"`
		PostmasterFingerprint string `json:"postmaster_fingerprint"`
		PostmasterPrivkey     string `json:"postmaster_privkey"`
		PostmasterPubkey      string `json:"postmaster_pubkey"`
	} `json:"authority"`
	Global struct {
		City         string `json:"city"`
		Country      string `json:"country"`
		Organization string `json:"organization"`
		Postmaster   string `json:"postmaster"`
		Status       int64  `json:"status"`
	} `json:"global"`
	Objects struct {
		Cas     []string      `json:"cas"`
		Openpgp []interface{} `json:"openpgp"`
		Smime   []interface{} `json:"smime"`
		Users   []interface{} `json:"users"`
	} `json:"objects"`
	Openpgp struct {
		MainKeysize int64 `json:"main_keysize"`
		SubKeysize  int64 `json:"sub_keysize"`
	} `json:"openpgp"`
	Options struct {
		ExternalAuto     int64  `json:"external_auto"`
		Keyserver        string `json:"keyserver"`
		PolicyDecryption int64  `json:"policy_decryption"`
		PolicyEncryption int64  `json:"policy_encryption"`
		PolicySign       int64  `json:"policy_sign"`
		PolicyVerify     int64  `json:"policy_verify"`
	} `json:"options"`
}

var _ sophos.Endpoint = &Emailpki{}

var defsEmailpki = map[string]sophos.RestObject{
	"EmailpkiGroup":   &EmailpkiGroup{},
	"EmailpkiOpenpgp": &EmailpkiOpenpgp{},
	"EmailpkiSmime":   &EmailpkiSmime{},
	"EmailpkiUser":    &EmailpkiUser{},
}

// RestObjects implements the sophos.Node interface and returns a map of Emailpki's Objects
func (Emailpki) RestObjects() map[string]sophos.RestObject { return defsEmailpki }

// GetPath implements sophos.RestGetter
func (*Emailpki) GetPath() string { return "/api/nodes/emailpki" }

// RefRequired implements sophos.RestGetter
func (*Emailpki) RefRequired() (string, bool) { return "", false }

var defEmailpki = &sophos.Definition{Description: "emailpki", Name: "emailpki", Link: "/api/definitions/emailpki"}

// Definition returns the /api/definitions struct of Emailpki
func (Emailpki) Definition() sophos.Definition { return *defEmailpki }

// ApiRoutes returns all known Emailpki Paths
func (Emailpki) ApiRoutes() []string {
	return []string{
		"/api/objects/emailpki/group/",
		"/api/objects/emailpki/group/{ref}",
		"/api/objects/emailpki/group/{ref}/usedby",
		"/api/objects/emailpki/openpgp/",
		"/api/objects/emailpki/openpgp/{ref}",
		"/api/objects/emailpki/openpgp/{ref}/usedby",
		"/api/objects/emailpki/smime/",
		"/api/objects/emailpki/smime/{ref}",
		"/api/objects/emailpki/smime/{ref}/usedby",
		"/api/objects/emailpki/user/",
		"/api/objects/emailpki/user/{ref}",
		"/api/objects/emailpki/user/{ref}/usedby",
	}
}

// References returns the Emailpki's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Emailpki) References() []string {
	return []string{
		"REF_EmailpkiGroup",
		"REF_EmailpkiOpenpgp",
		"REF_EmailpkiSmime",
		"REF_EmailpkiUser",
	}
}

// EmailpkiGroups is an Sophos Endpoint subType and implements sophos.RestObject
type EmailpkiGroups []EmailpkiGroup

// EmailpkiGroup represents a UTM group
type EmailpkiGroup struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
}

var _ sophos.RestGetter = &EmailpkiGroup{}

// GetPath implements sophos.RestObject and returns the EmailpkiGroups GET path
// Returns all available emailpki/group objects
func (*EmailpkiGroups) GetPath() string { return "/api/objects/emailpki/group/" }

// RefRequired implements sophos.RestObject
func (*EmailpkiGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EmailpkiGroups GET path
// Returns all available group types
func (e *EmailpkiGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/emailpki/group/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EmailpkiGroup) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EmailpkiGroup DELETE path
// Creates or updates the complete object group
func (*EmailpkiGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EmailpkiGroup PATCH path
// Changes to parts of the object group types
func (*EmailpkiGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EmailpkiGroup POST path
// Create a new emailpki/group object
func (*EmailpkiGroup) PostPath() string {
	return "/api/objects/emailpki/group/"
}

// PutPath implements sophos.RestObject and returns the EmailpkiGroup PUT path
// Creates or updates the complete object group
func (*EmailpkiGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*EmailpkiGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/group/%s/usedby", ref)
}

// EmailpkiOpenpgps is an Sophos Endpoint subType and implements sophos.RestObject
type EmailpkiOpenpgps []EmailpkiOpenpgp

// EmailpkiOpenpgp represents a UTM Email encryption OpenPGP key
type EmailpkiOpenpgp struct {
	Locked      string `json:"_locked"`
	Reference   string `json:"_ref"`
	_type       string `json:"_type"`
	Fingerprint string `json:"fingerprint"`
	Name        string `json:"name"`
	// Privkey default value is ""
	Privkey string `json:"privkey"`
	Pubkey  string `json:"pubkey"`
	// Realname default value is ""
	Realname string        `json:"realname"`
	Comment  string        `json:"comment"`
	Emails   []interface{} `json:"emails"`
	// Expires default value is ""
	Expires string `json:"expires"`
}

var _ sophos.RestGetter = &EmailpkiOpenpgp{}

// GetPath implements sophos.RestObject and returns the EmailpkiOpenpgps GET path
// Returns all available emailpki/openpgp objects
func (*EmailpkiOpenpgps) GetPath() string { return "/api/objects/emailpki/openpgp/" }

// RefRequired implements sophos.RestObject
func (*EmailpkiOpenpgps) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EmailpkiOpenpgps GET path
// Returns all available openpgp types
func (e *EmailpkiOpenpgp) GetPath() string {
	return fmt.Sprintf("/api/objects/emailpki/openpgp/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EmailpkiOpenpgp) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EmailpkiOpenpgp DELETE path
// Creates or updates the complete object openpgp
func (*EmailpkiOpenpgp) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/openpgp/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EmailpkiOpenpgp PATCH path
// Changes to parts of the object openpgp types
func (*EmailpkiOpenpgp) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/openpgp/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EmailpkiOpenpgp POST path
// Create a new emailpki/openpgp object
func (*EmailpkiOpenpgp) PostPath() string {
	return "/api/objects/emailpki/openpgp/"
}

// PutPath implements sophos.RestObject and returns the EmailpkiOpenpgp PUT path
// Creates or updates the complete object openpgp
func (*EmailpkiOpenpgp) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/openpgp/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*EmailpkiOpenpgp) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/openpgp/%s/usedby", ref)
}

// EmailpkiSmimes is an Sophos Endpoint subType and implements sophos.RestObject
type EmailpkiSmimes []EmailpkiSmime

// EmailpkiSmime represents a UTM email encryption S/MIME certificate
type EmailpkiSmime struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	// Domaincert default value is false
	Domaincert  bool   `json:"domaincert"`
	Expires     string `json:"expires"`
	Fingerprint string `json:"fingerprint"`
	// Key default value is ""
	Key     string        `json:"key"`
	Name    string        `json:"name"`
	Cert    string        `json:"cert"`
	Comment string        `json:"comment"`
	Dn      string        `json:"dn"`
	Emails  []interface{} `json:"emails"`
	// Realname default value is ""
	Realname string `json:"realname"`
	// Trust default value is false
	Trust bool `json:"trust"`
}

var _ sophos.RestGetter = &EmailpkiSmime{}

// GetPath implements sophos.RestObject and returns the EmailpkiSmimes GET path
// Returns all available emailpki/smime objects
func (*EmailpkiSmimes) GetPath() string { return "/api/objects/emailpki/smime/" }

// RefRequired implements sophos.RestObject
func (*EmailpkiSmimes) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EmailpkiSmimes GET path
// Returns all available smime types
func (e *EmailpkiSmime) GetPath() string {
	return fmt.Sprintf("/api/objects/emailpki/smime/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EmailpkiSmime) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EmailpkiSmime DELETE path
// Creates or updates the complete object smime
func (*EmailpkiSmime) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/smime/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EmailpkiSmime PATCH path
// Changes to parts of the object smime types
func (*EmailpkiSmime) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/smime/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EmailpkiSmime POST path
// Create a new emailpki/smime object
func (*EmailpkiSmime) PostPath() string {
	return "/api/objects/emailpki/smime/"
}

// PutPath implements sophos.RestObject and returns the EmailpkiSmime PUT path
// Creates or updates the complete object smime
func (*EmailpkiSmime) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/smime/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*EmailpkiSmime) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/smime/%s/usedby", ref)
}

// EmailpkiUsers is an Sophos Endpoint subType and implements sophos.RestObject
type EmailpkiUsers []EmailpkiUser

// EmailpkiUser represents a UTM email encryption user
type EmailpkiUser struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	// Verify can be one of: []string{"global", "on", "off"}
	// Verify default value is "global"
	Verify string `json:"verify"`
	// Domaincert default value is false
	Domaincert bool   `json:"domaincert"`
	Name       string `json:"name"`
	// OpenpgpStatus default value is true
	OpenpgpStatus bool `json:"openpgp_status"`
	// SmimeStatus default value is true
	SmimeStatus bool   `json:"smime_status"`
	Realname    string `json:"realname"`
	// Sign can be one of: []string{"global", "on", "off"}
	// Sign default value is "global"
	Sign string `json:"sign"`
	// Smime description: REF(emailpki/smime)
	// Smime default value is ""
	Smime   string `json:"smime"`
	Comment string `json:"comment"`
	// Decrypt can be one of: []string{"global", "on", "off"}
	// Decrypt default value is "global"
	Decrypt string `json:"decrypt"`
	// Encrypt can be one of: []string{"global", "on", "off"}
	// Encrypt default value is "global"
	Encrypt string `json:"encrypt"`
	// Openpgp description: REF(emailpki/openpgp)
	// Openpgp default value is ""
	Openpgp string `json:"openpgp"`
}

var _ sophos.RestGetter = &EmailpkiUser{}

// GetPath implements sophos.RestObject and returns the EmailpkiUsers GET path
// Returns all available emailpki/user objects
func (*EmailpkiUsers) GetPath() string { return "/api/objects/emailpki/user/" }

// RefRequired implements sophos.RestObject
func (*EmailpkiUsers) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EmailpkiUsers GET path
// Returns all available user types
func (e *EmailpkiUser) GetPath() string {
	return fmt.Sprintf("/api/objects/emailpki/user/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EmailpkiUser) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EmailpkiUser DELETE path
// Creates or updates the complete object user
func (*EmailpkiUser) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/user/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EmailpkiUser PATCH path
// Changes to parts of the object user types
func (*EmailpkiUser) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/user/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EmailpkiUser POST path
// Create a new emailpki/user object
func (*EmailpkiUser) PostPath() string {
	return "/api/objects/emailpki/user/"
}

// PutPath implements sophos.RestObject and returns the EmailpkiUser PUT path
// Creates or updates the complete object user
func (*EmailpkiUser) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/user/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*EmailpkiUser) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/emailpki/user/%s/usedby", ref)
}
