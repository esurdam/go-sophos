package objects

// THIS FILE IS AUTO GENERATED BY bin/gen.go! DO NOT EDIT!

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Qos is a generated struct representing the Sophos Qos Endpoint
// GET /api/nodes/qos
type Qos struct {
	Advanced struct {
		Ecn                 int64 `json:"ecn"`
		KeepClassAfterEncap int64 `json:"keep_class_after_encap"`
	} `json:"advanced"`
	Interfaces []string `json:"interfaces"`
}

var _ sophos.Endpoint = &Qos{}

var defsQos = map[string]sophos.RestObject{
	"QosApplicationSelector":  &QosApplicationSelector{},
	"QosGroup":                &QosGroup{},
	"QosIngressRule":          &QosIngressRule{},
	"QosInterface":            &QosInterface{},
	"QosRule":                 &QosRule{},
	"QosTrafficSelector":      &QosTrafficSelector{},
	"QosTrafficSelectorGroup": &QosTrafficSelectorGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of Qos's Objects
func (Qos) RestObjects() map[string]sophos.RestObject { return defsQos }

// GetPath implements sophos.RestGetter
func (*Qos) GetPath() string { return "/api/nodes/qos" }

// RefRequired implements sophos.RestGetter
func (*Qos) RefRequired() (string, bool) { return "", false }

var defQos = &sophos.Definition{Description: "qos", Name: "qos", Link: "/api/definitions/qos"}

// Definition returns the /api/definitions struct of Qos
func (Qos) Definition() sophos.Definition { return *defQos }

// ApiRoutes returns all known Qos Paths
func (Qos) ApiRoutes() []string {
	return []string{
		"/api/objects/qos/application_selector/",
		"/api/objects/qos/application_selector/{ref}",
		"/api/objects/qos/application_selector/{ref}/usedby",
		"/api/objects/qos/group/",
		"/api/objects/qos/group/{ref}",
		"/api/objects/qos/group/{ref}/usedby",
		"/api/objects/qos/ingress_rule/",
		"/api/objects/qos/ingress_rule/{ref}",
		"/api/objects/qos/ingress_rule/{ref}/usedby",
		"/api/objects/qos/interface/",
		"/api/objects/qos/interface/{ref}",
		"/api/objects/qos/interface/{ref}/usedby",
		"/api/objects/qos/rule/",
		"/api/objects/qos/rule/{ref}",
		"/api/objects/qos/rule/{ref}/usedby",
		"/api/objects/qos/traffic_selector/",
		"/api/objects/qos/traffic_selector/{ref}",
		"/api/objects/qos/traffic_selector/{ref}/usedby",
		"/api/objects/qos/traffic_selector_group/",
		"/api/objects/qos/traffic_selector_group/{ref}",
		"/api/objects/qos/traffic_selector_group/{ref}/usedby",
	}
}

// References returns the Qos's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Qos) References() []string {
	return []string{
		"REF_QosApplicationSelector",
		"REF_QosGroup",
		"REF_QosIngressRule",
		"REF_QosInterface",
		"REF_QosRule",
		"REF_QosTrafficSelector",
		"REF_QosTrafficSelectorGroup",
	}
}

// QosApplicationSelectors is an Sophos Endpoint subType and implements sophos.RestObject
type QosApplicationSelectors []QosApplicationSelector

// QosApplicationSelector represents a UTM QoS application selector
type QosApplicationSelector struct {
	Locked       string   `json:"_locked"`
	ObjectType   string   `json:"_type"`
	Reference    string   `json:"_ref"`
	Applications []string `json:"applications"`
	Comment      string   `json:"comment"`
	Connbytes    int      `json:"connbytes"`
	// ConnbytesUpperlimit default value is false
	ConnbytesUpperlimit bool `json:"connbytes_upperlimit"`
	// Destination description: REF(network/*)
	Destination             string   `json:"destination"`
	GroupFilterProductivity int      `json:"group_filter_productivity"`
	GroupFilterRisk         int      `json:"group_filter_risk"`
	Groups                  []string `json:"groups"`
	Name                    string   `json:"name"`
	// Source description: REF(network/*)
	Source string `json:"source"`
}

var _ sophos.RestGetter = &QosApplicationSelector{}

// GetPath implements sophos.RestObject and returns the QosApplicationSelectors GET path
// Returns all available qos/application_selector objects
func (*QosApplicationSelectors) GetPath() string { return "/api/objects/qos/application_selector/" }

// RefRequired implements sophos.RestObject
func (*QosApplicationSelectors) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the QosApplicationSelectors GET path
// Returns all available application_selector types
func (q *QosApplicationSelector) GetPath() string {
	return fmt.Sprintf("/api/objects/qos/application_selector/%s", q.Reference)
}

// RefRequired implements sophos.RestObject
func (q *QosApplicationSelector) RefRequired() (string, bool) { return q.Reference, true }

// DeletePath implements sophos.RestObject and returns the QosApplicationSelector DELETE path
// Creates or updates the complete object application_selector
func (*QosApplicationSelector) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/application_selector/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosApplicationSelector PATCH path
// Changes to parts of the object application_selector types
func (*QosApplicationSelector) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/application_selector/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosApplicationSelector POST path
// Create a new qos/application_selector object
func (*QosApplicationSelector) PostPath() string {
	return "/api/objects/qos/application_selector/"
}

// PutPath implements sophos.RestObject and returns the QosApplicationSelector PUT path
// Creates or updates the complete object application_selector
func (*QosApplicationSelector) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/application_selector/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosApplicationSelector) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/application_selector/%s/usedby", ref)
}

// QosGroups is an Sophos Endpoint subType and implements sophos.RestObject
type QosGroups []QosGroup

// QosGroup represents a UTM group
type QosGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &QosGroup{}

// GetPath implements sophos.RestObject and returns the QosGroups GET path
// Returns all available qos/group objects
func (*QosGroups) GetPath() string { return "/api/objects/qos/group/" }

// RefRequired implements sophos.RestObject
func (*QosGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the QosGroups GET path
// Returns all available group types
func (q *QosGroup) GetPath() string { return fmt.Sprintf("/api/objects/qos/group/%s", q.Reference) }

// RefRequired implements sophos.RestObject
func (q *QosGroup) RefRequired() (string, bool) { return q.Reference, true }

// DeletePath implements sophos.RestObject and returns the QosGroup DELETE path
// Creates or updates the complete object group
func (*QosGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosGroup PATCH path
// Changes to parts of the object group types
func (*QosGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosGroup POST path
// Create a new qos/group object
func (*QosGroup) PostPath() string {
	return "/api/objects/qos/group/"
}

// PutPath implements sophos.RestObject and returns the QosGroup PUT path
// Creates or updates the complete object group
func (*QosGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/group/%s/usedby", ref)
}

// QosIngressRules is an Sophos Endpoint subType and implements sophos.RestObject
type QosIngressRules []QosIngressRule

// QosIngressRule represents a UTM QoS traffic throttling
type QosIngressRule struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Limit      int    `json:"limit"`
	// Mode can be one of: []string{",", "srcip", "dstip", "srcip,dstip"}
	// Mode default value is ""
	Mode string `json:"mode"`
	Name string `json:"name"`
	// Status default value is false
	Status           bool     `json:"status"`
	TrafficSelectors []string `json:"traffic_selectors"`
}

var _ sophos.RestGetter = &QosIngressRule{}

// GetPath implements sophos.RestObject and returns the QosIngressRules GET path
// Returns all available qos/ingress_rule objects
func (*QosIngressRules) GetPath() string { return "/api/objects/qos/ingress_rule/" }

// RefRequired implements sophos.RestObject
func (*QosIngressRules) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the QosIngressRules GET path
// Returns all available ingress_rule types
func (q *QosIngressRule) GetPath() string {
	return fmt.Sprintf("/api/objects/qos/ingress_rule/%s", q.Reference)
}

// RefRequired implements sophos.RestObject
func (q *QosIngressRule) RefRequired() (string, bool) { return q.Reference, true }

// DeletePath implements sophos.RestObject and returns the QosIngressRule DELETE path
// Creates or updates the complete object ingress_rule
func (*QosIngressRule) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/ingress_rule/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosIngressRule PATCH path
// Changes to parts of the object ingress_rule types
func (*QosIngressRule) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/ingress_rule/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosIngressRule POST path
// Create a new qos/ingress_rule object
func (*QosIngressRule) PostPath() string {
	return "/api/objects/qos/ingress_rule/"
}

// PutPath implements sophos.RestObject and returns the QosIngressRule PUT path
// Creates or updates the complete object ingress_rule
func (*QosIngressRule) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/ingress_rule/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosIngressRule) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/ingress_rule/%s/usedby", ref)
}

// QosInterfaces is an Sophos Endpoint subType and implements sophos.RestObject
type QosInterfaces []QosInterface

// QosInterface represents a UTM QoS interface
type QosInterface struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Downlink   int    `json:"downlink"`
	// DownlinkOptimizer default value is true
	DownlinkOptimizer bool     `json:"downlink_optimizer"`
	IngressRules      []string `json:"ingress_rules"`
	// Interface description: REF(interface/*)
	Interface string   `json:"interface"`
	Name      string   `json:"name"`
	Rules     []string `json:"rules"`
	// Status default value is false
	Status bool `json:"status"`
	Uplink int  `json:"uplink"`
	// UplinkLimit default value is true
	UplinkLimit bool `json:"uplink_limit"`
	// UplinkOptimizer default value is true
	UplinkOptimizer bool `json:"uplink_optimizer"`
}

var _ sophos.RestGetter = &QosInterface{}

// GetPath implements sophos.RestObject and returns the QosInterfaces GET path
// Returns all available qos/interface objects
func (*QosInterfaces) GetPath() string { return "/api/objects/qos/interface/" }

// RefRequired implements sophos.RestObject
func (*QosInterfaces) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the QosInterfaces GET path
// Returns all available interface types
func (q *QosInterface) GetPath() string {
	return fmt.Sprintf("/api/objects/qos/interface/%s", q.Reference)
}

// RefRequired implements sophos.RestObject
func (q *QosInterface) RefRequired() (string, bool) { return q.Reference, true }

// DeletePath implements sophos.RestObject and returns the QosInterface DELETE path
// Creates or updates the complete object interface
func (*QosInterface) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/interface/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosInterface PATCH path
// Changes to parts of the object interface types
func (*QosInterface) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/interface/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosInterface POST path
// Create a new qos/interface object
func (*QosInterface) PostPath() string {
	return "/api/objects/qos/interface/"
}

// PutPath implements sophos.RestObject and returns the QosInterface PUT path
// Creates or updates the complete object interface
func (*QosInterface) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/interface/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosInterface) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/interface/%s/usedby", ref)
}

// QosRules is an Sophos Endpoint subType and implements sophos.RestObject
type QosRules []QosRule

// QosRule represents a UTM QoS bandwidth pool
type QosRule struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Bandwidth  int    `json:"bandwidth"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
	// Status default value is false
	Status           bool     `json:"status"`
	TrafficSelectors []string `json:"traffic_selectors"`
	// UpperLimitStatus default value is false
	UpperLimitStatus bool `json:"upper_limit_status"`
	UpperLimitValue  int  `json:"upper_limit_value"`
}

var _ sophos.RestGetter = &QosRule{}

// GetPath implements sophos.RestObject and returns the QosRules GET path
// Returns all available qos/rule objects
func (*QosRules) GetPath() string { return "/api/objects/qos/rule/" }

// RefRequired implements sophos.RestObject
func (*QosRules) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the QosRules GET path
// Returns all available rule types
func (q *QosRule) GetPath() string { return fmt.Sprintf("/api/objects/qos/rule/%s", q.Reference) }

// RefRequired implements sophos.RestObject
func (q *QosRule) RefRequired() (string, bool) { return q.Reference, true }

// DeletePath implements sophos.RestObject and returns the QosRule DELETE path
// Creates or updates the complete object rule
func (*QosRule) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/rule/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosRule PATCH path
// Changes to parts of the object rule types
func (*QosRule) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/rule/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosRule POST path
// Create a new qos/rule object
func (*QosRule) PostPath() string {
	return "/api/objects/qos/rule/"
}

// PutPath implements sophos.RestObject and returns the QosRule PUT path
// Creates or updates the complete object rule
func (*QosRule) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/rule/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosRule) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/rule/%s/usedby", ref)
}

// QosTrafficSelectors is an Sophos Endpoint subType and implements sophos.RestObject
type QosTrafficSelectors []QosTrafficSelector

// QosTrafficSelector represents a UTM QoS traffic selector
type QosTrafficSelector struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Connbytes  int    `json:"connbytes"`
	// ConnbytesUpperlimit default value is false
	ConnbytesUpperlimit bool `json:"connbytes_upperlimit"`
	// Destination description: REF(network/*)
	Destination string `json:"destination"`
	// DscpString can be one of: []string{"BE", "AF11", "AF12", "AF13", "AF21", "AF22", "AF23", "AF31", "AF32", "AF33", "AF41", "AF42", "AF43", "CS1", "CS2", "CS3", "CS4", "CS5", "CS6", "CS7", "EF"}
	// DscpString default value is "BE"
	DscpString string `json:"dscp_string"`
	// DscpType can be one of: []string{"off", "value", "class"}
	// DscpType default value is "off"
	DscpType  string `json:"dscp_type"`
	DscpValue int    `json:"dscp_value"`
	// Helper description: (REGEX)
	// Helper default value is ""
	Helper string `json:"helper"`
	Name   string `json:"name"`
	// PacketLength default value is ""
	PacketLength string `json:"packet_length"`
	// Service description: REF(service/*)
	// Service default value is "REF_ServiceAny"
	Service string `json:"service"`
	// Source description: REF(network/*)
	Source   string   `json:"source"`
	TcpFlags []string `json:"tcp_flags"`
	// Tos can be one of: []string{"off", "normal", "min_cost", "max_reliable", "max_throughput", "min_delay"}
	// Tos default value is "off"
	Tos string `json:"tos"`
}

var _ sophos.RestGetter = &QosTrafficSelector{}

// GetPath implements sophos.RestObject and returns the QosTrafficSelectors GET path
// Returns all available qos/traffic_selector objects
func (*QosTrafficSelectors) GetPath() string { return "/api/objects/qos/traffic_selector/" }

// RefRequired implements sophos.RestObject
func (*QosTrafficSelectors) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the QosTrafficSelectors GET path
// Returns all available traffic_selector types
func (q *QosTrafficSelector) GetPath() string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector/%s", q.Reference)
}

// RefRequired implements sophos.RestObject
func (q *QosTrafficSelector) RefRequired() (string, bool) { return q.Reference, true }

// DeletePath implements sophos.RestObject and returns the QosTrafficSelector DELETE path
// Creates or updates the complete object traffic_selector
func (*QosTrafficSelector) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosTrafficSelector PATCH path
// Changes to parts of the object traffic_selector types
func (*QosTrafficSelector) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosTrafficSelector POST path
// Create a new qos/traffic_selector object
func (*QosTrafficSelector) PostPath() string {
	return "/api/objects/qos/traffic_selector/"
}

// PutPath implements sophos.RestObject and returns the QosTrafficSelector PUT path
// Creates or updates the complete object traffic_selector
func (*QosTrafficSelector) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosTrafficSelector) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector/%s/usedby", ref)
}

// QosTrafficSelectorGroups is an Sophos Endpoint subType and implements sophos.RestObject
type QosTrafficSelectorGroups []QosTrafficSelectorGroup

// QosTrafficSelectorGroup represents a UTM QoS traffic selector group
type QosTrafficSelectorGroup struct {
	Locked     string   `json:"_locked"`
	ObjectType string   `json:"_type"`
	Reference  string   `json:"_ref"`
	Comment    string   `json:"comment"`
	Members    []string `json:"members"`
	Name       string   `json:"name"`
}

var _ sophos.RestGetter = &QosTrafficSelectorGroup{}

// GetPath implements sophos.RestObject and returns the QosTrafficSelectorGroups GET path
// Returns all available qos/traffic_selector_group objects
func (*QosTrafficSelectorGroups) GetPath() string { return "/api/objects/qos/traffic_selector_group/" }

// RefRequired implements sophos.RestObject
func (*QosTrafficSelectorGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the QosTrafficSelectorGroups GET path
// Returns all available traffic_selector_group types
func (q *QosTrafficSelectorGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector_group/%s", q.Reference)
}

// RefRequired implements sophos.RestObject
func (q *QosTrafficSelectorGroup) RefRequired() (string, bool) { return q.Reference, true }

// DeletePath implements sophos.RestObject and returns the QosTrafficSelectorGroup DELETE path
// Creates or updates the complete object traffic_selector_group
func (*QosTrafficSelectorGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector_group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosTrafficSelectorGroup PATCH path
// Changes to parts of the object traffic_selector_group types
func (*QosTrafficSelectorGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector_group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosTrafficSelectorGroup POST path
// Create a new qos/traffic_selector_group object
func (*QosTrafficSelectorGroup) PostPath() string {
	return "/api/objects/qos/traffic_selector_group/"
}

// PutPath implements sophos.RestObject and returns the QosTrafficSelectorGroup PUT path
// Creates or updates the complete object traffic_selector_group
func (*QosTrafficSelectorGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector_group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosTrafficSelectorGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector_group/%s/usedby", ref)
}
