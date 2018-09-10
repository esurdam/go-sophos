package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Ips is a generated struct representing the Sophos Ips Endpoint
// GET /api/nodes/ips
type Ips struct {
	DNSServers     []interface{} `json:"dns_servers"`
	Engine         string        `json:"engine"`
	Exceptions     []string      `json:"exceptions"`
	Failopen       int64         `json:"failopen"`
	FileBasedRules int64         `json:"file_based_rules"`
	Groups         []string      `json:"groups"`
	HTTPServers    []interface{} `json:"http_servers"`
	Ipsfb          struct {
		AlertInterval  int64 `json:"alert_interval"`
		ConfigInterval int64 `json:"config_interval"`
		Debug          int64 `json:"debug"`
	} `json:"ipsfb"`
	LocalNetworks  []string      `json:"local_networks"`
	NumInstances   int64         `json:"num_instances"`
	PatternChannel string        `json:"pattern_channel"`
	Policy         string        `json:"policy"`
	QueueLength    int64         `json:"queue_length"`
	QueueThreshold int64         `json:"queue_threshold"`
	ReloadMethod   string        `json:"reload_method"`
	RestartPolicy  string        `json:"restart_policy"`
	RuleModifiers  []interface{} `json:"rule_modifiers"`
	Rules          []interface{} `json:"rules"`
	SkipAcks       int64         `json:"skip_acks"`
	SMTPServers    []interface{} `json:"smtp_servers"`
	Snortsettings  struct {
		MaxQueuedBytes int64  `json:"max_queued_bytes"`
		MaxQueuedSegs  int64  `json:"max_queued_segs"`
		MaxTcp         int64  `json:"max_tcp"`
		MaxUdp         int64  `json:"max_udp"`
		Memcap         int64  `json:"memcap"`
		SearchMethod   string `json:"search_method"`
	} `json:"snortsettings"`
	SqlServers []interface{} `json:"sql_servers"`
	Status     int64         `json:"status"`
}

var _ sophos.Endpoint = &Ips{}

var defsIps = map[string]sophos.RestObject{
	"IpsException":    &IpsException{},
	"IpsGroup":        &IpsGroup{},
	"IpsRule":         &IpsRule{},
	"IpsRuleModifier": &IpsRuleModifier{},
}

// RestObjects implements the sophos.Node interface and returns a map of Ips's Objects
func (Ips) RestObjects() map[string]sophos.RestObject { return defsIps }

// GetPath implements sophos.RestGetter
func (*Ips) GetPath() string { return "/api/nodes/ips" }

// RefRequired implements sophos.RestGetter
func (*Ips) RefRequired() (string, bool) { return "", false }

var defIps = &sophos.Definition{Description: "ips", Name: "ips", Link: "/api/definitions/ips"}

// Definition returns the /api/definitions struct of Ips
func (Ips) Definition() sophos.Definition { return *defIps }

// ApiRoutes returns all known Ips Paths
func (Ips) ApiRoutes() []string {
	return []string{
		"/api/objects/ips/exception/",
		"/api/objects/ips/exception/{ref}",
		"/api/objects/ips/exception/{ref}/usedby",
		"/api/objects/ips/group/",
		"/api/objects/ips/group/{ref}",
		"/api/objects/ips/group/{ref}/usedby",
		"/api/objects/ips/rule/",
		"/api/objects/ips/rule/{ref}",
		"/api/objects/ips/rule/{ref}/usedby",
		"/api/objects/ips/rule_modifier/",
		"/api/objects/ips/rule_modifier/{ref}",
		"/api/objects/ips/rule_modifier/{ref}/usedby",
	}
}

// References returns the Ips's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Ips) References() []string {
	return []string{
		"REF_IpsException",
		"REF_IpsGroup",
		"REF_IpsRule",
		"REF_IpsRuleModifier",
	}
}

// IpsExceptions is an Sophos Endpoint subType and implements sophos.RestObject
type IpsExceptions []IpsException

// IpsException is a generated Sophos object
type IpsException struct {
	Locked              string   `json:"_locked"`
	Reference           string   `json:"_ref"`
	ObjectType          string   `json:"_type"`
	Comment             string   `json:"comment"`
	DestinationNetworks []string `json:"destination_networks"`
	Name                string   `json:"name"`
	Operator            string   `json:"operator"`
	Services            []string `json:"services"`
	Skiplist            []string `json:"skiplist"`
	SourceNetworks      []string `json:"source_networks"`
	Status              bool     `json:"status"`
}

var _ sophos.RestGetter = &IpsException{}

// GetPath implements sophos.RestObject and returns the IpsExceptions GET path
// Returns all available ips/exception objects
func (*IpsExceptions) GetPath() string { return "/api/objects/ips/exception/" }

// RefRequired implements sophos.RestObject
func (*IpsExceptions) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsExceptions GET path
// Returns all available exception types
func (i *IpsException) GetPath() string {
	return fmt.Sprintf("/api/objects/ips/exception/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsException) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsException DELETE path
// Creates or updates the complete object exception
func (*IpsException) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/exception/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsException PATCH path
// Changes to parts of the object exception types
func (*IpsException) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/exception/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsException POST path
// Create a new ips/exception object
func (*IpsException) PostPath() string {
	return "/api/objects/ips/exception/"
}

// PutPath implements sophos.RestObject and returns the IpsException PUT path
// Creates or updates the complete object exception
func (*IpsException) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/exception/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*IpsException) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/exception/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *IpsException) GetType() string { return i.ObjectType }

// IpsGroups is an Sophos Endpoint subType and implements sophos.RestObject
type IpsGroups []IpsGroup

// IpsGroup is a generated Sophos object
type IpsGroup struct {
	Locked       string   `json:"_locked"`
	Reference    string   `json:"_ref"`
	ObjectType   string   `json:"_type"`
	Action       string   `json:"action"`
	Age          int64    `json:"age"`
	Comment      string   `json:"comment"`
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Notification bool     `json:"notification"`
	Status       bool     `json:"status"`
	Subgroups    []string `json:"subgroups"`
	Warnings     bool     `json:"warnings"`
}

var _ sophos.RestGetter = &IpsGroup{}

// GetPath implements sophos.RestObject and returns the IpsGroups GET path
// Returns all available ips/group objects
func (*IpsGroups) GetPath() string { return "/api/objects/ips/group/" }

// RefRequired implements sophos.RestObject
func (*IpsGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsGroups GET path
// Returns all available group types
func (i *IpsGroup) GetPath() string { return fmt.Sprintf("/api/objects/ips/group/%s", i.Reference) }

// RefRequired implements sophos.RestObject
func (i *IpsGroup) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsGroup DELETE path
// Creates or updates the complete object group
func (*IpsGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsGroup PATCH path
// Changes to parts of the object group types
func (*IpsGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsGroup POST path
// Create a new ips/group object
func (*IpsGroup) PostPath() string {
	return "/api/objects/ips/group/"
}

// PutPath implements sophos.RestObject and returns the IpsGroup PUT path
// Creates or updates the complete object group
func (*IpsGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*IpsGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/group/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *IpsGroup) GetType() string { return i.ObjectType }

// IpsRules is an Sophos Endpoint subType and implements sophos.RestObject
type IpsRules []IpsRule

// IpsRule represents a UTM IPS rule
type IpsRule struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Status default value is false
	Status bool `json:"status"`
	// Action can be one of: []string{"alert", "drop"}
	// Action default value is "alert"
	Action  string `json:"action"`
	Comment string `json:"comment"`
	Filter1 string `json:"filter1"`
	Filter2 string `json:"filter2"`
	Msg     string `json:"msg"`
	Name    string `json:"name"`
	Sid     int    `json:"sid"`
}

var _ sophos.RestGetter = &IpsRule{}

// GetPath implements sophos.RestObject and returns the IpsRules GET path
// Returns all available ips/rule objects
func (*IpsRules) GetPath() string { return "/api/objects/ips/rule/" }

// RefRequired implements sophos.RestObject
func (*IpsRules) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsRules GET path
// Returns all available rule types
func (i *IpsRule) GetPath() string { return fmt.Sprintf("/api/objects/ips/rule/%s", i.Reference) }

// RefRequired implements sophos.RestObject
func (i *IpsRule) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsRule DELETE path
// Creates or updates the complete object rule
func (*IpsRule) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/rule/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsRule PATCH path
// Changes to parts of the object rule types
func (*IpsRule) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/rule/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsRule POST path
// Create a new ips/rule object
func (*IpsRule) PostPath() string {
	return "/api/objects/ips/rule/"
}

// PutPath implements sophos.RestObject and returns the IpsRule PUT path
// Creates or updates the complete object rule
func (*IpsRule) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/rule/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*IpsRule) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/rule/%s/usedby", ref)
}

// IpsRuleModifiers is an Sophos Endpoint subType and implements sophos.RestObject
type IpsRuleModifiers []IpsRuleModifier

// IpsRuleModifier represents a UTM IPS rule modifier
type IpsRuleModifier struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
	// Notification default value is false
	Notification bool `json:"notification"`
	Sid          int  `json:"sid"`
	// Status default value is false
	Status bool `json:"status"`
	// Action can be one of: []string{"alert", "drop"}
	// Action default value is "drop"
	Action string `json:"action"`
}

var _ sophos.RestGetter = &IpsRuleModifier{}

// GetPath implements sophos.RestObject and returns the IpsRuleModifiers GET path
// Returns all available ips/rule_modifier objects
func (*IpsRuleModifiers) GetPath() string { return "/api/objects/ips/rule_modifier/" }

// RefRequired implements sophos.RestObject
func (*IpsRuleModifiers) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the IpsRuleModifiers GET path
// Returns all available rule_modifier types
func (i *IpsRuleModifier) GetPath() string {
	return fmt.Sprintf("/api/objects/ips/rule_modifier/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *IpsRuleModifier) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the IpsRuleModifier DELETE path
// Creates or updates the complete object rule_modifier
func (*IpsRuleModifier) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/rule_modifier/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the IpsRuleModifier PATCH path
// Changes to parts of the object rule_modifier types
func (*IpsRuleModifier) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/rule_modifier/%s", ref)
}

// PostPath implements sophos.RestObject and returns the IpsRuleModifier POST path
// Create a new ips/rule_modifier object
func (*IpsRuleModifier) PostPath() string {
	return "/api/objects/ips/rule_modifier/"
}

// PutPath implements sophos.RestObject and returns the IpsRuleModifier PUT path
// Creates or updates the complete object rule_modifier
func (*IpsRuleModifier) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/rule_modifier/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*IpsRuleModifier) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/ips/rule_modifier/%s/usedby", ref)
}
