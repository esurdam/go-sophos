package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Ca is a generated struct representing the Sophos Ca Endpoint
// GET /api/nodes/ca
type Ca struct {
	CaGost     string `json:"ca_gost"`
	CaIpsec    string `json:"ca_ipsec"`
	CaProxies  string `json:"ca_proxies"`
	CaRed      string `json:"ca_red"`
	DefKeysize int64  `json:"def_keysize"`
	GlobalCas  struct {
		EmailEncryption struct {
			TrustNewCas int64         `json:"trust_new_cas"`
			Trusted     []interface{} `json:"trusted"`
			Untrusted   []interface{} `json:"untrusted"`
		} `json:"email_encryption"`
		HTTPProxy struct {
			TrustNewCas int64         `json:"trust_new_cas"`
			Trusted     []interface{} `json:"trusted"`
			Untrusted   []interface{} `json:"untrusted"`
		} `json:"http_proxy"`
	} `json:"global_cas"`
}

var defsCa = map[string]sophos.RestObject{
	"CaCrl":                &CaCrl{},
	"CaGroup":              &CaGroup{},
	"CaHostCert":           &CaHostCert{},
	"CaHostKeyCert":        &CaHostKeyCert{},
	"CaHttpVerificationCa": &CaHttpVerificationCa{},
	"CaMetaCrl":            &CaMetaCrl{},
	"CaMetaX509":           &CaMetaX509{},
	"CaRsa":                &CaRsa{},
	"CaSigningCa":          &CaSigningCa{},
	"CaVerificationCa":     &CaVerificationCa{},
}

// RestObjects implements the sophos.Node interface and returns a map of Ca's Objects
func (Ca) RestObjects() map[string]sophos.RestObject { return defsCa }

// GetPath implements sophos.RestGetter
func (*Ca) GetPath() string { return "/api/nodes/ca" }

// RefRequired implements sophos.RestGetter
func (*Ca) RefRequired() (string, bool) { return "", false }

var defCa = &sophos.Definition{Description: "ca", Name: "ca", Link: "/api/definitions/ca"}

// Definition returns the /api/definitions struct of Ca
func (Ca) Definition() sophos.Definition { return *defCa }

// ApiRoutes returns all known Ca Paths
func (Ca) ApiRoutes() []string {
	return []string{
		"/api/objects/ca/crl/",
		"/api/objects/ca/crl/{ref}",
		"/api/objects/ca/crl/{ref}/usedby",
		"/api/objects/ca/group/",
		"/api/objects/ca/group/{ref}",
		"/api/objects/ca/group/{ref}/usedby",
		"/api/objects/ca/host_cert/",
		"/api/objects/ca/host_cert/{ref}",
		"/api/objects/ca/host_cert/{ref}/usedby",
		"/api/objects/ca/host_key_cert/",
		"/api/objects/ca/host_key_cert/{ref}",
		"/api/objects/ca/host_key_cert/{ref}/usedby",
		"/api/objects/ca/http_verification_ca/",
		"/api/objects/ca/http_verification_ca/{ref}",
		"/api/objects/ca/http_verification_ca/{ref}/usedby",
		"/api/objects/ca/meta_crl/",
		"/api/objects/ca/meta_crl/{ref}",
		"/api/objects/ca/meta_crl/{ref}/usedby",
		"/api/objects/ca/meta_x509/",
		"/api/objects/ca/meta_x509/{ref}",
		"/api/objects/ca/meta_x509/{ref}/usedby",
		"/api/objects/ca/rsa/",
		"/api/objects/ca/rsa/{ref}",
		"/api/objects/ca/rsa/{ref}/usedby",
		"/api/objects/ca/signing_ca/",
		"/api/objects/ca/signing_ca/{ref}",
		"/api/objects/ca/signing_ca/{ref}/usedby",
		"/api/objects/ca/verification_ca/",
		"/api/objects/ca/verification_ca/{ref}",
		"/api/objects/ca/verification_ca/{ref}/usedby",
	}
}

// References returns the Ca's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Ca) References() []string {
	return []string{
		"REF_CaCrl",
		"REF_CaGroup",
		"REF_CaHostCert",
		"REF_CaHostKeyCert",
		"REF_CaHttpVerificationCa",
		"REF_CaMetaCrl",
		"REF_CaMetaX509",
		"REF_CaRsa",
		"REF_CaSigningCa",
		"REF_CaVerificationCa",
	}
}

// CaCrl is an Sophos Endpoint subType and implements sophos.RestObject
type CaCrl []interface{}

// GetPath implements sophos.RestObject and returns the CaCrl GET path
// Returns all available ca/crl objects
func (*CaCrl) GetPath() string { return "/api/objects/ca/crl/" }

// RefRequired implements sophos.RestObject
func (*CaCrl) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the CaCrl DELETE path
// Creates or updates the complete object crl
func (*CaCrl) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/crl/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the CaCrl PATCH path
// Changes to parts of the object crl types
func (*CaCrl) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/crl/%s", ref)
}

// PostPath implements sophos.RestObject and returns the CaCrl POST path
// Create a new ca/crl object
func (*CaCrl) PostPath() string {
	return "/api/objects/ca/crl/"
}

// PutPath implements sophos.RestObject and returns the CaCrl PUT path
// Creates or updates the complete object crl
func (*CaCrl) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/crl/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*CaCrl) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/crl/%s/usedby", ref)
}

// CaGroup is an Sophos Endpoint subType and implements sophos.RestObject
type CaGroup []interface{}

// GetPath implements sophos.RestObject and returns the CaGroup GET path
// Returns all available ca/group objects
func (*CaGroup) GetPath() string { return "/api/objects/ca/group/" }

// RefRequired implements sophos.RestObject
func (*CaGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the CaGroup DELETE path
// Creates or updates the complete object group
func (*CaGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the CaGroup PATCH path
// Changes to parts of the object group types
func (*CaGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the CaGroup POST path
// Create a new ca/group object
func (*CaGroup) PostPath() string {
	return "/api/objects/ca/group/"
}

// PutPath implements sophos.RestObject and returns the CaGroup PUT path
// Creates or updates the complete object group
func (*CaGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*CaGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/group/%s/usedby", ref)
}

// CaHostCerts is an Sophos Endpoint subType and implements sophos.RestObject
type CaHostCerts []CaHostCert

// CaHostCert is a generated Sophos object
type CaHostCert struct {
	Locked      string `json:"_locked"`
	Reference   string `json:"_ref"`
	_type       string `json:"_type"`
	Certificate string `json:"certificate"`
	Comment     string `json:"comment"`
	Meta        string `json:"meta"`
	Name        string `json:"name"`
}

// GetPath implements sophos.RestObject and returns the CaHostCerts GET path
// Returns all available ca/host_cert objects
func (*CaHostCerts) GetPath() string { return "/api/objects/ca/host_cert/" }

// RefRequired implements sophos.RestObject
func (*CaHostCerts) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the CaHostCerts GET path
// Returns all available host_cert types
func (c *CaHostCert) GetPath() string { return fmt.Sprintf("/api/objects/ca/host_cert/%s", c.Reference) }

// RefRequired implements sophos.RestObject
func (c *CaHostCert) RefRequired() (string, bool) { return c.Reference, true }

// DeletePath implements sophos.RestObject and returns the CaHostCert DELETE path
// Creates or updates the complete object host_cert
func (*CaHostCert) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/host_cert/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the CaHostCert PATCH path
// Changes to parts of the object host_cert types
func (*CaHostCert) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/host_cert/%s", ref)
}

// PostPath implements sophos.RestObject and returns the CaHostCert POST path
// Create a new ca/host_cert object
func (*CaHostCert) PostPath() string {
	return "/api/objects/ca/host_cert/"
}

// PutPath implements sophos.RestObject and returns the CaHostCert PUT path
// Creates or updates the complete object host_cert
func (*CaHostCert) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/host_cert/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*CaHostCert) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/host_cert/%s/usedby", ref)
}

// GetType implements sophos.Object
func (c *CaHostCert) GetType() string { return c._type }

// CaHostKeyCerts is an Sophos Endpoint subType and implements sophos.RestObject
type CaHostKeyCerts []CaHostKeyCert

// CaHostKeyCert is a generated Sophos object
type CaHostKeyCert struct {
	Locked      string `json:"_locked"`
	Reference   string `json:"_ref"`
	_type       string `json:"_type"`
	Ca          string `json:"ca"`
	Certificate string `json:"certificate"`
	Comment     string `json:"comment"`
	Encrypted   bool   `json:"encrypted"`
	Key         string `json:"key"`
	Meta        string `json:"meta"`
	Name        string `json:"name"`
}

// GetPath implements sophos.RestObject and returns the CaHostKeyCerts GET path
// Returns all available ca/host_key_cert objects
func (*CaHostKeyCerts) GetPath() string { return "/api/objects/ca/host_key_cert/" }

// RefRequired implements sophos.RestObject
func (*CaHostKeyCerts) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the CaHostKeyCerts GET path
// Returns all available host_key_cert types
func (c *CaHostKeyCert) GetPath() string {
	return fmt.Sprintf("/api/objects/ca/host_key_cert/%s", c.Reference)
}

// RefRequired implements sophos.RestObject
func (c *CaHostKeyCert) RefRequired() (string, bool) { return c.Reference, true }

// DeletePath implements sophos.RestObject and returns the CaHostKeyCert DELETE path
// Creates or updates the complete object host_key_cert
func (*CaHostKeyCert) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/host_key_cert/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the CaHostKeyCert PATCH path
// Changes to parts of the object host_key_cert types
func (*CaHostKeyCert) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/host_key_cert/%s", ref)
}

// PostPath implements sophos.RestObject and returns the CaHostKeyCert POST path
// Create a new ca/host_key_cert object
func (*CaHostKeyCert) PostPath() string {
	return "/api/objects/ca/host_key_cert/"
}

// PutPath implements sophos.RestObject and returns the CaHostKeyCert PUT path
// Creates or updates the complete object host_key_cert
func (*CaHostKeyCert) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/host_key_cert/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*CaHostKeyCert) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/host_key_cert/%s/usedby", ref)
}

// GetType implements sophos.Object
func (c *CaHostKeyCert) GetType() string { return c._type }

// CaHttpVerificationCa is an Sophos Endpoint subType and implements sophos.RestObject
type CaHttpVerificationCa []interface{}

// GetPath implements sophos.RestObject and returns the CaHttpVerificationCa GET path
// Returns all available ca/http_verification_ca objects
func (*CaHttpVerificationCa) GetPath() string { return "/api/objects/ca/http_verification_ca/" }

// RefRequired implements sophos.RestObject
func (*CaHttpVerificationCa) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the CaHttpVerificationCa DELETE path
// Creates or updates the complete object http_verification_ca
func (*CaHttpVerificationCa) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/http_verification_ca/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the CaHttpVerificationCa PATCH path
// Changes to parts of the object http_verification_ca types
func (*CaHttpVerificationCa) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/http_verification_ca/%s", ref)
}

// PostPath implements sophos.RestObject and returns the CaHttpVerificationCa POST path
// Create a new ca/http_verification_ca object
func (*CaHttpVerificationCa) PostPath() string {
	return "/api/objects/ca/http_verification_ca/"
}

// PutPath implements sophos.RestObject and returns the CaHttpVerificationCa PUT path
// Creates or updates the complete object http_verification_ca
func (*CaHttpVerificationCa) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/http_verification_ca/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*CaHttpVerificationCa) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/http_verification_ca/%s/usedby", ref)
}

// CaMetaCrl is an Sophos Endpoint subType and implements sophos.RestObject
type CaMetaCrl []interface{}

// GetPath implements sophos.RestObject and returns the CaMetaCrl GET path
// Returns all available ca/meta_crl objects
func (*CaMetaCrl) GetPath() string { return "/api/objects/ca/meta_crl/" }

// RefRequired implements sophos.RestObject
func (*CaMetaCrl) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the CaMetaCrl DELETE path
// Creates or updates the complete object meta_crl
func (*CaMetaCrl) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/meta_crl/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the CaMetaCrl PATCH path
// Changes to parts of the object meta_crl types
func (*CaMetaCrl) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/meta_crl/%s", ref)
}

// PostPath implements sophos.RestObject and returns the CaMetaCrl POST path
// Create a new ca/meta_crl object
func (*CaMetaCrl) PostPath() string {
	return "/api/objects/ca/meta_crl/"
}

// PutPath implements sophos.RestObject and returns the CaMetaCrl PUT path
// Creates or updates the complete object meta_crl
func (*CaMetaCrl) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/meta_crl/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*CaMetaCrl) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/meta_crl/%s/usedby", ref)
}

// CaMetaX509s is an Sophos Endpoint subType and implements sophos.RestObject
type CaMetaX509s []CaMetaX509

// CaMetaX509 is a generated Sophos object
type CaMetaX509 struct {
	Locked             string   `json:"_locked"`
	Reference          string   `json:"_ref"`
	_type              string   `json:"_type"`
	Comment            string   `json:"comment"`
	Enddate            string   `json:"enddate"`
	Fingerprint        string   `json:"fingerprint"`
	Issuer             string   `json:"issuer"`
	IssuerHash         string   `json:"issuer_hash"`
	Name               string   `json:"name"`
	PublicKeyAlgorithm string   `json:"public_key_algorithm"`
	Serial             string   `json:"serial"`
	Startdate          string   `json:"startdate"`
	Subject            string   `json:"subject"`
	SubjectAltNames    []string `json:"subject_alt_names"`
	SubjectHash        string   `json:"subject_hash"`
	VpnID              string   `json:"vpn_id"`
	VpnIDType          string   `json:"vpn_id_type"`
}

// GetPath implements sophos.RestObject and returns the CaMetaX509s GET path
// Returns all available ca/meta_x509 objects
func (*CaMetaX509s) GetPath() string { return "/api/objects/ca/meta_x509/" }

// RefRequired implements sophos.RestObject
func (*CaMetaX509s) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the CaMetaX509s GET path
// Returns all available meta_x509 types
func (c *CaMetaX509) GetPath() string { return fmt.Sprintf("/api/objects/ca/meta_x509/%s", c.Reference) }

// RefRequired implements sophos.RestObject
func (c *CaMetaX509) RefRequired() (string, bool) { return c.Reference, true }

// DeletePath implements sophos.RestObject and returns the CaMetaX509 DELETE path
// Creates or updates the complete object meta_x509
func (*CaMetaX509) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/meta_x509/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the CaMetaX509 PATCH path
// Changes to parts of the object meta_x509 types
func (*CaMetaX509) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/meta_x509/%s", ref)
}

// PostPath implements sophos.RestObject and returns the CaMetaX509 POST path
// Create a new ca/meta_x509 object
func (*CaMetaX509) PostPath() string {
	return "/api/objects/ca/meta_x509/"
}

// PutPath implements sophos.RestObject and returns the CaMetaX509 PUT path
// Creates or updates the complete object meta_x509
func (*CaMetaX509) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/meta_x509/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*CaMetaX509) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/meta_x509/%s/usedby", ref)
}

// GetType implements sophos.Object
func (c *CaMetaX509) GetType() string { return c._type }

// CaRsas is an Sophos Endpoint subType and implements sophos.RestObject
type CaRsas []CaRsa

// CaRsa is a generated Sophos object
type CaRsa struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Key       string `json:"key"`
	KeySize   int64  `json:"key_size"`
	Name      string `json:"name"`
	Pubkey    string `json:"pubkey"`
	VpnID     string `json:"vpn_id"`
	VpnIDType string `json:"vpn_id_type"`
}

// GetPath implements sophos.RestObject and returns the CaRsas GET path
// Returns all available ca/rsa objects
func (*CaRsas) GetPath() string { return "/api/objects/ca/rsa/" }

// RefRequired implements sophos.RestObject
func (*CaRsas) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the CaRsas GET path
// Returns all available rsa types
func (c *CaRsa) GetPath() string { return fmt.Sprintf("/api/objects/ca/rsa/%s", c.Reference) }

// RefRequired implements sophos.RestObject
func (c *CaRsa) RefRequired() (string, bool) { return c.Reference, true }

// DeletePath implements sophos.RestObject and returns the CaRsa DELETE path
// Creates or updates the complete object rsa
func (*CaRsa) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/rsa/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the CaRsa PATCH path
// Changes to parts of the object rsa types
func (*CaRsa) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/rsa/%s", ref)
}

// PostPath implements sophos.RestObject and returns the CaRsa POST path
// Create a new ca/rsa object
func (*CaRsa) PostPath() string {
	return "/api/objects/ca/rsa/"
}

// PutPath implements sophos.RestObject and returns the CaRsa PUT path
// Creates or updates the complete object rsa
func (*CaRsa) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/rsa/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*CaRsa) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/rsa/%s/usedby", ref)
}

// GetType implements sophos.Object
func (c *CaRsa) GetType() string { return c._type }

// CaSigningCas is an Sophos Endpoint subType and implements sophos.RestObject
type CaSigningCas []CaSigningCa

// CaSigningCa is a generated Sophos object
type CaSigningCa struct {
	Locked      string `json:"_locked"`
	Reference   string `json:"_ref"`
	_type       string `json:"_type"`
	Certificate string `json:"certificate"`
	Comment     string `json:"comment"`
	Config      string `json:"config"`
	Encrypted   bool   `json:"encrypted"`
	Index       string `json:"index"`
	Key         string `json:"key"`
	Meta        string `json:"meta"`
	Name        string `json:"name"`
	Serial      string `json:"serial"`
}

// GetPath implements sophos.RestObject and returns the CaSigningCas GET path
// Returns all available ca/signing_ca objects
func (*CaSigningCas) GetPath() string { return "/api/objects/ca/signing_ca/" }

// RefRequired implements sophos.RestObject
func (*CaSigningCas) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the CaSigningCas GET path
// Returns all available signing_ca types
func (c *CaSigningCa) GetPath() string {
	return fmt.Sprintf("/api/objects/ca/signing_ca/%s", c.Reference)
}

// RefRequired implements sophos.RestObject
func (c *CaSigningCa) RefRequired() (string, bool) { return c.Reference, true }

// DeletePath implements sophos.RestObject and returns the CaSigningCa DELETE path
// Creates or updates the complete object signing_ca
func (*CaSigningCa) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/signing_ca/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the CaSigningCa PATCH path
// Changes to parts of the object signing_ca types
func (*CaSigningCa) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/signing_ca/%s", ref)
}

// PostPath implements sophos.RestObject and returns the CaSigningCa POST path
// Create a new ca/signing_ca object
func (*CaSigningCa) PostPath() string {
	return "/api/objects/ca/signing_ca/"
}

// PutPath implements sophos.RestObject and returns the CaSigningCa PUT path
// Creates or updates the complete object signing_ca
func (*CaSigningCa) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/signing_ca/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*CaSigningCa) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/signing_ca/%s/usedby", ref)
}

// GetType implements sophos.Object
func (c *CaSigningCa) GetType() string { return c._type }

// CaVerificationCa is an Sophos Endpoint subType and implements sophos.RestObject
type CaVerificationCa []interface{}

// GetPath implements sophos.RestObject and returns the CaVerificationCa GET path
// Returns all available ca/verification_ca objects
func (*CaVerificationCa) GetPath() string { return "/api/objects/ca/verification_ca/" }

// RefRequired implements sophos.RestObject
func (*CaVerificationCa) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the CaVerificationCa DELETE path
// Creates or updates the complete object verification_ca
func (*CaVerificationCa) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/verification_ca/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the CaVerificationCa PATCH path
// Changes to parts of the object verification_ca types
func (*CaVerificationCa) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/verification_ca/%s", ref)
}

// PostPath implements sophos.RestObject and returns the CaVerificationCa POST path
// Create a new ca/verification_ca object
func (*CaVerificationCa) PostPath() string {
	return "/api/objects/ca/verification_ca/"
}

// PutPath implements sophos.RestObject and returns the CaVerificationCa PUT path
// Creates or updates the complete object verification_ca
func (*CaVerificationCa) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/verification_ca/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*CaVerificationCa) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ca/verification_ca/%s/usedby", ref)
}
