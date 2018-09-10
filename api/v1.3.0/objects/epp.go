package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Epp is a generated struct representing the Sophos Epp Endpoint
// GET /api/nodes/epp
type Epp struct {
	AllowedNetworks       []string      `json:"allowed_networks"`
	Certificate           string        `json:"certificate"`
	City                  string        `json:"city"`
	Country               string        `json:"country"`
	DefaultEndpointsGroup string        `json:"default_endpoints_group"`
	Devices               []interface{} `json:"devices"`
	Email                 string        `json:"email"`
	Endpoints             []interface{} `json:"endpoints"`
	EndpointsGroups       []string      `json:"endpoints_groups"`
	Exceptions            struct {
		Av []interface{} `json:"av"`
		Dc []interface{} `json:"dc"`
	} `json:"exceptions"`
	FallbackURL       string `json:"fallback_url"`
	MagnetPassword    string `json:"magnet_password"`
	MagnetUsername    string `json:"magnet_username"`
	Organization      string `json:"organization"`
	ParentProxyHost   string `json:"parent_proxy_host"`
	ParentProxyPort   int64  `json:"parent_proxy_port"`
	ParentProxyStatus int64  `json:"parent_proxy_status"`
	Policies          struct {
		Av []string `json:"av"`
		Dc []string `json:"dc"`
	} `json:"policies"`
	Port              int64  `json:"port"`
	PrivateKey        string `json:"private_key"`
	RegistrationToken string `json:"registration_token"`
	Status            struct {
		Av     int64 `json:"av"`
		Broker int64 `json:"broker"`
		Dc     int64 `json:"dc"`
		Epp    int64 `json:"epp"`
		Wc     int64 `json:"wc"`
	} `json:"status"`
	TamperPassword string `json:"tamper_password"`
	Version        string `json:"version"`
	WdxToken       string `json:"wdx_token"`
}

var _ sophos.Endpoint = &Epp{}

var defsEpp = map[string]sophos.RestObject{
	"EppAvException":    &EppAvException{},
	"EppAvPolicy":       &EppAvPolicy{},
	"EppDcException":    &EppDcException{},
	"EppDcPolicy":       &EppDcPolicy{},
	"EppDevice":         &EppDevice{},
	"EppEndpoint":       &EppEndpoint{},
	"EppEndpointsGroup": &EppEndpointsGroup{},
	"EppGroup":          &EppGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of Epp's Objects
func (Epp) RestObjects() map[string]sophos.RestObject { return defsEpp }

// GetPath implements sophos.RestGetter
func (*Epp) GetPath() string { return "/api/nodes/epp" }

// RefRequired implements sophos.RestGetter
func (*Epp) RefRequired() (string, bool) { return "", false }

var defEpp = &sophos.Definition{Description: "epp", Name: "epp", Link: "/api/definitions/epp"}

// Definition returns the /api/definitions struct of Epp
func (Epp) Definition() sophos.Definition { return *defEpp }

// ApiRoutes returns all known Epp Paths
func (Epp) ApiRoutes() []string {
	return []string{
		"/api/objects/epp/av_exception/",
		"/api/objects/epp/av_exception/{ref}",
		"/api/objects/epp/av_exception/{ref}/usedby",
		"/api/objects/epp/av_policy/",
		"/api/objects/epp/av_policy/{ref}",
		"/api/objects/epp/av_policy/{ref}/usedby",
		"/api/objects/epp/dc_exception/",
		"/api/objects/epp/dc_exception/{ref}",
		"/api/objects/epp/dc_exception/{ref}/usedby",
		"/api/objects/epp/dc_policy/",
		"/api/objects/epp/dc_policy/{ref}",
		"/api/objects/epp/dc_policy/{ref}/usedby",
		"/api/objects/epp/device/",
		"/api/objects/epp/device/{ref}",
		"/api/objects/epp/device/{ref}/usedby",
		"/api/objects/epp/endpoint/",
		"/api/objects/epp/endpoint/{ref}",
		"/api/objects/epp/endpoint/{ref}/usedby",
		"/api/objects/epp/endpoints_group/",
		"/api/objects/epp/endpoints_group/{ref}",
		"/api/objects/epp/endpoints_group/{ref}/usedby",
		"/api/objects/epp/group/",
		"/api/objects/epp/group/{ref}",
		"/api/objects/epp/group/{ref}/usedby",
	}
}

// References returns the Epp's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Epp) References() []string {
	return []string{
		"REF_EppAvException",
		"REF_EppAvPolicy",
		"REF_EppDcException",
		"REF_EppDcPolicy",
		"REF_EppDevice",
		"REF_EppEndpoint",
		"REF_EppEndpointsGroup",
		"REF_EppGroup",
	}
}

// EppAvExceptions is an Sophos Endpoint subType and implements sophos.RestObject
type EppAvExceptions []EppAvException

// EppAvException represents a UTM antivirus exception
type EppAvException struct {
	Locked          string        `json:"_locked"`
	ObjectType      string        `json:"_type"`
	Reference       string        `json:"_ref"`
	EndpointsGroups []interface{} `json:"endpoints_groups"`
	HipsName        string        `json:"hips_name"`
	// IpAddress description: (IPADDR)
	// IpAddress default value is ""
	IpAddress string `json:"ip_address"`
	// Timeline default value is ""
	Timeline string `json:"timeline"`
	// Type can be one of: []string{"adware_pua", "scanning_exclusions", "scanning_extensions", "buffer_overflow", "suspicious_files", "suspicious_behaviours", "websites"}
	// Type default value is "websites"
	Type string `json:"type"`
	// Checksum default value is ""
	Checksum      string `json:"checksum"`
	Comment       string `json:"comment"`
	IpAddressMask int    `json:"ip_address_mask"`
	Name          string `json:"name"`
	// WebFormat can be one of: []string{"domain_name", "ip_address", "ip_address_mask"}
	// WebFormat default value is "domain_name"
	WebFormat string `json:"web_format"`
}

var _ sophos.RestGetter = &EppAvException{}

// GetPath implements sophos.RestObject and returns the EppAvExceptions GET path
// Returns all available epp/av_exception objects
func (*EppAvExceptions) GetPath() string { return "/api/objects/epp/av_exception/" }

// RefRequired implements sophos.RestObject
func (*EppAvExceptions) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EppAvExceptions GET path
// Returns all available av_exception types
func (e *EppAvException) GetPath() string {
	return fmt.Sprintf("/api/objects/epp/av_exception/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EppAvException) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EppAvException DELETE path
// Creates or updates the complete object av_exception
func (*EppAvException) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_exception/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppAvException PATCH path
// Changes to parts of the object av_exception types
func (*EppAvException) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_exception/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppAvException POST path
// Create a new epp/av_exception object
func (*EppAvException) PostPath() string {
	return "/api/objects/epp/av_exception/"
}

// PutPath implements sophos.RestObject and returns the EppAvException PUT path
// Creates or updates the complete object av_exception
func (*EppAvException) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_exception/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*EppAvException) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_exception/%s/usedby", ref)
}

// EppAvPolicys is an Sophos Endpoint subType and implements sophos.RestObject
type EppAvPolicys []EppAvPolicy

// EppAvPolicy is a generated Sophos object
type EppAvPolicy struct {
	Locked                    string `json:"_locked"`
	Reference                 string `json:"_ref"`
	ObjectType                string `json:"_type"`
	AlertOnly                 bool   `json:"alert_only"`
	AutoCleanup               bool   `json:"auto_cleanup"`
	BlockMaliciousSites       bool   `json:"block_malicious_sites"`
	Comment                   string `json:"comment"`
	DetectBufferOverflow      bool   `json:"detect_buffer_overflow"`
	DetectMaliciousFiles      bool   `json:"detect_malicious_files"`
	DetectSuspiciousBehaviour bool   `json:"detect_suspicious_behaviour"`
	DownloadScanning          bool   `json:"download_scanning"`
	Hips                      bool   `json:"hips"`
	LowPriorityScan           bool   `json:"low_priority_scan"`
	Name                      string `json:"name"`
	OnAccessScanning          bool   `json:"on_access_scanning"`
	OnRead                    bool   `json:"on_read"`
	OnRename                  bool   `json:"on_rename"`
	OnWrite                   bool   `json:"on_write"`
	RootKitScan               bool   `json:"root_kit_scan"`
	ScanForPua                bool   `json:"scan_for_pua"`
	ScanForSuspiciousFiles    bool   `json:"scan_for_suspicious_files"`
	ScanInsideArchive         bool   `json:"scan_inside_archive"`
	ScanSystemMemory          bool   `json:"scan_system_memory"`
	ScheduledScanning         bool   `json:"scheduled_scanning"`
	SendSampleFile            bool   `json:"send_sample_file"`
	SophosLiveProtection      bool   `json:"sophos_live_protection"`
	TimeEvent                 string `json:"time_event"`
	WebProtection             bool   `json:"web_protection"`
}

var _ sophos.RestGetter = &EppAvPolicy{}

// GetPath implements sophos.RestObject and returns the EppAvPolicys GET path
// Returns all available epp/av_policy objects
func (*EppAvPolicys) GetPath() string { return "/api/objects/epp/av_policy/" }

// RefRequired implements sophos.RestObject
func (*EppAvPolicys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EppAvPolicys GET path
// Returns all available av_policy types
func (e *EppAvPolicy) GetPath() string {
	return fmt.Sprintf("/api/objects/epp/av_policy/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EppAvPolicy) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EppAvPolicy DELETE path
// Creates or updates the complete object av_policy
func (*EppAvPolicy) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_policy/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppAvPolicy PATCH path
// Changes to parts of the object av_policy types
func (*EppAvPolicy) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_policy/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppAvPolicy POST path
// Create a new epp/av_policy object
func (*EppAvPolicy) PostPath() string {
	return "/api/objects/epp/av_policy/"
}

// PutPath implements sophos.RestObject and returns the EppAvPolicy PUT path
// Creates or updates the complete object av_policy
func (*EppAvPolicy) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_policy/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*EppAvPolicy) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_policy/%s/usedby", ref)
}

// GetType implements sophos.Object
func (e *EppAvPolicy) GetType() string { return e.ObjectType }

// EppDcExceptions is an Sophos Endpoint subType and implements sophos.RestObject
type EppDcExceptions []EppDcException

// EppDcException represents a UTM device exception
type EppDcException struct {
	Locked                       string        `json:"_locked"`
	ObjectType                   string        `json:"_type"`
	Reference                    string        `json:"_ref"`
	CustomBlockedEndpointsGroups []interface{} `json:"custom_blocked_endpoints_groups"`
	DeviceId                     string        `json:"device_id"`
	// DeviceType can be one of: []string{"floppy_drive", "optical_drive", "removable_storage", "encrypted_storage", "modem", "wireless", "bluetooth", "infrared"}
	// DeviceType default value is "removable_storage"
	DeviceType             string        `json:"device_type"`
	Name                   string        `json:"name"`
	AllowedEndpointsGroups []interface{} `json:"allowed_endpoints_groups"`
	Comment                string        `json:"comment"`
}

var _ sophos.RestGetter = &EppDcException{}

// GetPath implements sophos.RestObject and returns the EppDcExceptions GET path
// Returns all available epp/dc_exception objects
func (*EppDcExceptions) GetPath() string { return "/api/objects/epp/dc_exception/" }

// RefRequired implements sophos.RestObject
func (*EppDcExceptions) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EppDcExceptions GET path
// Returns all available dc_exception types
func (e *EppDcException) GetPath() string {
	return fmt.Sprintf("/api/objects/epp/dc_exception/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EppDcException) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EppDcException DELETE path
// Creates or updates the complete object dc_exception
func (*EppDcException) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_exception/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppDcException PATCH path
// Changes to parts of the object dc_exception types
func (*EppDcException) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_exception/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppDcException POST path
// Create a new epp/dc_exception object
func (*EppDcException) PostPath() string {
	return "/api/objects/epp/dc_exception/"
}

// PutPath implements sophos.RestObject and returns the EppDcException PUT path
// Creates or updates the complete object dc_exception
func (*EppDcException) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_exception/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*EppDcException) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_exception/%s/usedby", ref)
}

// EppDcPolicys is an Sophos Endpoint subType and implements sophos.RestObject
type EppDcPolicys []EppDcPolicy

// EppDcPolicy is a generated Sophos object
type EppDcPolicy struct {
	Locked           string `json:"_locked"`
	Reference        string `json:"_ref"`
	ObjectType       string `json:"_type"`
	Bluetooth        string `json:"bluetooth"`
	Comment          string `json:"comment"`
	EncryptedStorage string `json:"encrypted_storage"`
	FloppyDrive      string `json:"floppy_drive"`
	Infrared         string `json:"infrared"`
	Modem            string `json:"modem"`
	Name             string `json:"name"`
	OpticalDrive     string `json:"optical_drive"`
	RemovableStorage string `json:"removable_storage"`
	Wireless         string `json:"wireless"`
}

var _ sophos.RestGetter = &EppDcPolicy{}

// GetPath implements sophos.RestObject and returns the EppDcPolicys GET path
// Returns all available epp/dc_policy objects
func (*EppDcPolicys) GetPath() string { return "/api/objects/epp/dc_policy/" }

// RefRequired implements sophos.RestObject
func (*EppDcPolicys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EppDcPolicys GET path
// Returns all available dc_policy types
func (e *EppDcPolicy) GetPath() string {
	return fmt.Sprintf("/api/objects/epp/dc_policy/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EppDcPolicy) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EppDcPolicy DELETE path
// Creates or updates the complete object dc_policy
func (*EppDcPolicy) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_policy/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppDcPolicy PATCH path
// Changes to parts of the object dc_policy types
func (*EppDcPolicy) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_policy/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppDcPolicy POST path
// Create a new epp/dc_policy object
func (*EppDcPolicy) PostPath() string {
	return "/api/objects/epp/dc_policy/"
}

// PutPath implements sophos.RestObject and returns the EppDcPolicy PUT path
// Creates or updates the complete object dc_policy
func (*EppDcPolicy) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_policy/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*EppDcPolicy) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_policy/%s/usedby", ref)
}

// GetType implements sophos.Object
func (e *EppDcPolicy) GetType() string { return e.ObjectType }

// EppDevices is an Sophos Endpoint subType and implements sophos.RestObject
type EppDevices []EppDevice

// EppDevice represents a UTM device
type EppDevice struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// ProductName default value is ""
	ProductName                  string        `json:"product_name"`
	AllowedEndpointsGroups       []interface{} `json:"allowed_endpoints_groups"`
	CustomBlockedEndpointsGroups []interface{} `json:"custom_blocked_endpoints_groups"`
	DeviceId                     string        `json:"device_id"`
	// GenericFlag default value is false
	GenericFlag bool `json:"generic_flag"`
	// LastEndpoint description: REF(epp/endpoint)
	// LastEndpoint default value is ""
	LastEndpoint string `json:"last_endpoint"`
	Name         string `json:"name"`
	Comment      string `json:"comment"`
	// DeviceType can be one of: []string{"floppy_drive", "optical_drive", "removable_storage", "encrypted_storage", "modem", "wireless", "bluetooth", "infrared"}
	// DeviceType default value is "removable_storage"
	DeviceType string `json:"device_type"`
	// InstanceId default value is ""
	InstanceId string `json:"instance_id"`
}

var _ sophos.RestGetter = &EppDevice{}

// GetPath implements sophos.RestObject and returns the EppDevices GET path
// Returns all available epp/device objects
func (*EppDevices) GetPath() string { return "/api/objects/epp/device/" }

// RefRequired implements sophos.RestObject
func (*EppDevices) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EppDevices GET path
// Returns all available device types
func (e *EppDevice) GetPath() string { return fmt.Sprintf("/api/objects/epp/device/%s", e.Reference) }

// RefRequired implements sophos.RestObject
func (e *EppDevice) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EppDevice DELETE path
// Creates or updates the complete object device
func (*EppDevice) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/device/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppDevice PATCH path
// Changes to parts of the object device types
func (*EppDevice) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/device/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppDevice POST path
// Create a new epp/device object
func (*EppDevice) PostPath() string {
	return "/api/objects/epp/device/"
}

// PutPath implements sophos.RestObject and returns the EppDevice PUT path
// Creates or updates the complete object device
func (*EppDevice) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/device/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*EppDevice) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/device/%s/usedby", ref)
}

// EppEndpoints is an Sophos Endpoint subType and implements sophos.RestObject
type EppEndpoints []EppEndpoint

// EppEndpoint represents a UTM computer
type EppEndpoint struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Os default value is ""
	Os string `json:"os"`
	// SavVersion default value is ""
	SavVersion string `json:"sav_version"`
	// TamperProtection default value is false
	TamperProtection bool `json:"tamper_protection"`
	// McsId default value is ""
	McsId string `json:"mcs_id"`
	Name  string `json:"name"`
	// SavStatus default value is false
	SavStatus bool `json:"sav_status"`
	// Accepted default value is false
	Accepted bool   `json:"accepted"`
	Comment  string `json:"comment"`
	// EndpointType can be one of: []string{"laptop", "desktop", "server"}
	// EndpointType default value is "desktop"
	EndpointType string `json:"endpoint_type"`
	// InventoryNumber default value is ""
	InventoryNumber string `json:"inventory_number"`
}

var _ sophos.RestGetter = &EppEndpoint{}

// GetPath implements sophos.RestObject and returns the EppEndpoints GET path
// Returns all available epp/endpoint objects
func (*EppEndpoints) GetPath() string { return "/api/objects/epp/endpoint/" }

// RefRequired implements sophos.RestObject
func (*EppEndpoints) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EppEndpoints GET path
// Returns all available endpoint types
func (e *EppEndpoint) GetPath() string {
	return fmt.Sprintf("/api/objects/epp/endpoint/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EppEndpoint) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EppEndpoint DELETE path
// Creates or updates the complete object endpoint
func (*EppEndpoint) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoint/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppEndpoint PATCH path
// Changes to parts of the object endpoint types
func (*EppEndpoint) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoint/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppEndpoint POST path
// Create a new epp/endpoint object
func (*EppEndpoint) PostPath() string {
	return "/api/objects/epp/endpoint/"
}

// PutPath implements sophos.RestObject and returns the EppEndpoint PUT path
// Creates or updates the complete object endpoint
func (*EppEndpoint) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoint/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*EppEndpoint) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoint/%s/usedby", ref)
}

// EppEndpointsGroups is an Sophos Endpoint subType and implements sophos.RestObject
type EppEndpointsGroups []EppEndpointsGroup

// EppEndpointsGroup is a generated Sophos object
type EppEndpointsGroup struct {
	Locked           string        `json:"_locked"`
	Reference        string        `json:"_ref"`
	ObjectType       string        `json:"_type"`
	AvPolicy         string        `json:"av_policy"`
	Comment          string        `json:"comment"`
	DcPolicy         string        `json:"dc_policy"`
	Endpoints        []interface{} `json:"endpoints"`
	Name             string        `json:"name"`
	ProxyAddress     string        `json:"proxy_address"`
	ProxyPassword    string        `json:"proxy_password"`
	ProxyPort        int64         `json:"proxy_port"`
	ProxySupport     bool          `json:"proxy_support"`
	ProxyUser        string        `json:"proxy_user"`
	TamperProtection bool          `json:"tamper_protection"`
	WebControl       bool          `json:"web_control"`
}

var _ sophos.RestGetter = &EppEndpointsGroup{}

// GetPath implements sophos.RestObject and returns the EppEndpointsGroups GET path
// Returns all available epp/endpoints_group objects
func (*EppEndpointsGroups) GetPath() string { return "/api/objects/epp/endpoints_group/" }

// RefRequired implements sophos.RestObject
func (*EppEndpointsGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EppEndpointsGroups GET path
// Returns all available endpoints_group types
func (e *EppEndpointsGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/epp/endpoints_group/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EppEndpointsGroup) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EppEndpointsGroup DELETE path
// Creates or updates the complete object endpoints_group
func (*EppEndpointsGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoints_group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppEndpointsGroup PATCH path
// Changes to parts of the object endpoints_group types
func (*EppEndpointsGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoints_group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppEndpointsGroup POST path
// Create a new epp/endpoints_group object
func (*EppEndpointsGroup) PostPath() string {
	return "/api/objects/epp/endpoints_group/"
}

// PutPath implements sophos.RestObject and returns the EppEndpointsGroup PUT path
// Creates or updates the complete object endpoints_group
func (*EppEndpointsGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoints_group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*EppEndpointsGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoints_group/%s/usedby", ref)
}

// GetType implements sophos.Object
func (e *EppEndpointsGroup) GetType() string { return e.ObjectType }

// EppGroups is an Sophos Endpoint subType and implements sophos.RestObject
type EppGroups []EppGroup

// EppGroup represents a UTM group
type EppGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &EppGroup{}

// GetPath implements sophos.RestObject and returns the EppGroups GET path
// Returns all available epp/group objects
func (*EppGroups) GetPath() string { return "/api/objects/epp/group/" }

// RefRequired implements sophos.RestObject
func (*EppGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EppGroups GET path
// Returns all available group types
func (e *EppGroup) GetPath() string { return fmt.Sprintf("/api/objects/epp/group/%s", e.Reference) }

// RefRequired implements sophos.RestObject
func (e *EppGroup) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EppGroup DELETE path
// Creates or updates the complete object group
func (*EppGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppGroup PATCH path
// Changes to parts of the object group types
func (*EppGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppGroup POST path
// Create a new epp/group object
func (*EppGroup) PostPath() string {
	return "/api/objects/epp/group/"
}

// PutPath implements sophos.RestObject and returns the EppGroup PUT path
// Creates or updates the complete object group
func (*EppGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*EppGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/group/%s/usedby", ref)
}
