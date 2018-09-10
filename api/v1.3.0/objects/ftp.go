package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Ftp is a generated struct representing the Sophos Ftp Endpoint
// GET /api/nodes/ftp
type Ftp struct {
	AllowedClients        []interface{} `json:"allowed_clients"`
	AllowedServers        []string      `json:"allowed_servers"`
	CffAv                 int64         `json:"cff_av"`
	CffAvEngines          string        `json:"cff_av_engines"`
	CffFileExtensions     []interface{} `json:"cff_file_extensions"`
	Exceptions            []interface{} `json:"exceptions"`
	MaxFileSize           int64         `json:"max_file_size"`
	MsWinMode             int64         `json:"ms_win_mode"`
	OperationMode         string        `json:"operation_mode"`
	RestrictedServers     []string      `json:"restricted_servers"`
	Status                int64         `json:"status"`
	TransparentSkip       []interface{} `json:"transparent_skip"`
	TransparentSkipAutoPf int64         `json:"transparent_skip_auto_pf"`
}

var _ sophos.Endpoint = &Ftp{}

var defsFtp = map[string]sophos.RestObject{
	"FtpException": &FtpException{},
	"FtpGroup":     &FtpGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of Ftp's Objects
func (Ftp) RestObjects() map[string]sophos.RestObject { return defsFtp }

// GetPath implements sophos.RestGetter
func (*Ftp) GetPath() string { return "/api/nodes/ftp" }

// RefRequired implements sophos.RestGetter
func (*Ftp) RefRequired() (string, bool) { return "", false }

var defFtp = &sophos.Definition{Description: "ftp", Name: "ftp", Link: "/api/definitions/ftp"}

// Definition returns the /api/definitions struct of Ftp
func (Ftp) Definition() sophos.Definition { return *defFtp }

// ApiRoutes returns all known Ftp Paths
func (Ftp) ApiRoutes() []string {
	return []string{
		"/api/objects/ftp/exception/",
		"/api/objects/ftp/exception/{ref}",
		"/api/objects/ftp/exception/{ref}/usedby",
		"/api/objects/ftp/group/",
		"/api/objects/ftp/group/{ref}",
		"/api/objects/ftp/group/{ref}/usedby",
	}
}

// References returns the Ftp's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Ftp) References() []string {
	return []string{
		"REF_FtpException",
		"REF_FtpGroup",
	}
}

// FtpExceptions is an Sophos Endpoint subType and implements sophos.RestObject
type FtpExceptions []FtpException

// FtpException represents a UTM FTP filter exception
type FtpException struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	// Status default value is false
	Status   bool          `json:"status"`
	Client   []interface{} `json:"client"`
	Comment  string        `json:"comment"`
	Name     string        `json:"name"`
	Server   []interface{} `json:"server"`
	Skiplist []interface{} `json:"skiplist"`
}

var _ sophos.RestGetter = &FtpException{}

// GetPath implements sophos.RestObject and returns the FtpExceptions GET path
// Returns all available ftp/exception objects
func (*FtpExceptions) GetPath() string { return "/api/objects/ftp/exception/" }

// RefRequired implements sophos.RestObject
func (*FtpExceptions) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the FtpExceptions GET path
// Returns all available exception types
func (f *FtpException) GetPath() string {
	return fmt.Sprintf("/api/objects/ftp/exception/%s", f.Reference)
}

// RefRequired implements sophos.RestObject
func (f *FtpException) RefRequired() (string, bool) { return f.Reference, true }

// DeletePath implements sophos.RestObject and returns the FtpException DELETE path
// Creates or updates the complete object exception
func (*FtpException) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ftp/exception/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the FtpException PATCH path
// Changes to parts of the object exception types
func (*FtpException) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ftp/exception/%s", ref)
}

// PostPath implements sophos.RestObject and returns the FtpException POST path
// Create a new ftp/exception object
func (*FtpException) PostPath() string {
	return "/api/objects/ftp/exception/"
}

// PutPath implements sophos.RestObject and returns the FtpException PUT path
// Creates or updates the complete object exception
func (*FtpException) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ftp/exception/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*FtpException) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ftp/exception/%s/usedby", ref)
}

// FtpGroups is an Sophos Endpoint subType and implements sophos.RestObject
type FtpGroups []FtpGroup

// FtpGroup represents a UTM group
type FtpGroup struct {
	Locked     string `json:"_locked"`
	Reference  string `json:"_ref"`
	ObjectType string `json:"_type"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &FtpGroup{}

// GetPath implements sophos.RestObject and returns the FtpGroups GET path
// Returns all available ftp/group objects
func (*FtpGroups) GetPath() string { return "/api/objects/ftp/group/" }

// RefRequired implements sophos.RestObject
func (*FtpGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the FtpGroups GET path
// Returns all available group types
func (f *FtpGroup) GetPath() string { return fmt.Sprintf("/api/objects/ftp/group/%s", f.Reference) }

// RefRequired implements sophos.RestObject
func (f *FtpGroup) RefRequired() (string, bool) { return f.Reference, true }

// DeletePath implements sophos.RestObject and returns the FtpGroup DELETE path
// Creates or updates the complete object group
func (*FtpGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ftp/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the FtpGroup PATCH path
// Changes to parts of the object group types
func (*FtpGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ftp/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the FtpGroup POST path
// Create a new ftp/group object
func (*FtpGroup) PostPath() string {
	return "/api/objects/ftp/group/"
}

// PutPath implements sophos.RestObject and returns the FtpGroup PUT path
// Creates or updates the complete object group
func (*FtpGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ftp/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*FtpGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ftp/group/%s/usedby", ref)
}
