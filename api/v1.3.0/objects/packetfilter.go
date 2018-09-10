package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Packetfilter is a generated struct representing the Sophos Packetfilter Endpoint
// GET /api/nodes/packetfilter
type Packetfilter struct {
	Advanced struct {
		BlockInvalidCtPackets int64    `json:"block_invalid_ct_packets"`
		CheckPacketLength     int64    `json:"check_packet_length"`
		ConntrackHelpers      []string `json:"conntrack_helpers"`
		FtpPorts              []int64  `json:"ftp_ports"`
		LogBroadcasts         int64    `json:"log_broadcasts"`
		LogDNSRequests        int64    `json:"log_dns_requests"`
		LogFtpData            int64    `json:"log_ftp_data"`
		LogInvalid            int64    `json:"log_invalid"`
		LogMcast              int64    `json:"log_mcast"`
		LogStrictTcpState     string   `json:"log_strict_tcp_state"`
		NoErrorReplay         int64    `json:"no_error_replay"`
		Optimize              struct {
			Ipset string `json:"ipset"`
			Ports int64  `json:"ports"`
		} `json:"optimize"`
		SpoofProtection  string `json:"spoof_protection"`
		StrictTcpState   int64  `json:"strict_tcp_state"`
		TcpMaxRetrans    int64  `json:"tcp_max_retrans"`
		TcpWindowScaling int64  `json:"tcp_window_scaling"`
	} `json:"advanced"`
	Rules      []string `json:"rules"`
	RulesAuto  []string `json:"rules_auto"`
	RulesBack  string   `json:"rules_back"`
	RulesFront string   `json:"rules_front"`
	Timeouts   struct {
		IPConntrackGenericTimeout        int64 `json:"ip_conntrack_generic_timeout"`
		IPConntrackIcmpTimeout           int64 `json:"ip_conntrack_icmp_timeout"`
		IPConntrackTcpTimeoutClose       int64 `json:"ip_conntrack_tcp_timeout_close"`
		IPConntrackTcpTimeoutCloseWait   int64 `json:"ip_conntrack_tcp_timeout_close_wait"`
		IPConntrackTcpTimeoutEstablished int64 `json:"ip_conntrack_tcp_timeout_established"`
		IPConntrackTcpTimeoutFinWait     int64 `json:"ip_conntrack_tcp_timeout_fin_wait"`
		IPConntrackTcpTimeoutLastAck     int64 `json:"ip_conntrack_tcp_timeout_last_ack"`
		IPConntrackTcpTimeoutMaxRetrans  int64 `json:"ip_conntrack_tcp_timeout_max_retrans"`
		IPConntrackTcpTimeoutSynRecv     int64 `json:"ip_conntrack_tcp_timeout_syn_recv"`
		IPConntrackTcpTimeoutSynSent     int64 `json:"ip_conntrack_tcp_timeout_syn_sent"`
		IPConntrackTcpTimeoutTimeWait    int64 `json:"ip_conntrack_tcp_timeout_time_wait"`
		IPConntrackUdpTimeout            int64 `json:"ip_conntrack_udp_timeout"`
		IPConntrackUdpTimeoutStream      int64 `json:"ip_conntrack_udp_timeout_stream"`
	} `json:"timeouts"`
}

var _ sophos.Endpoint = &Packetfilter{}

var defsPacketfilter = map[string]sophos.RestObject{
	"Packetfilter1to1Nat":      &Packetfilter1to1Nat{},
	"PacketfilterGenericProxy": &PacketfilterGenericProxy{},
	"PacketfilterGroup":        &PacketfilterGroup{},
	"PacketfilterLoadbalance":  &PacketfilterLoadbalance{},
	"PacketfilterMangle":       &PacketfilterMangle{},
	"PacketfilterMasq":         &PacketfilterMasq{},
	"PacketfilterNat":          &PacketfilterNat{},
	"PacketfilterPacketfilter": &PacketfilterPacketfilter{},
	"PacketfilterRuleset":      &PacketfilterRuleset{},
}

// RestObjects implements the sophos.Node interface and returns a map of Packetfilter's Objects
func (Packetfilter) RestObjects() map[string]sophos.RestObject { return defsPacketfilter }

// GetPath implements sophos.RestGetter
func (*Packetfilter) GetPath() string { return "/api/nodes/packetfilter" }

// RefRequired implements sophos.RestGetter
func (*Packetfilter) RefRequired() (string, bool) { return "", false }

var defPacketfilter = &sophos.Definition{Description: "packetfilter", Name: "packetfilter", Link: "/api/definitions/packetfilter"}

// Definition returns the /api/definitions struct of Packetfilter
func (Packetfilter) Definition() sophos.Definition { return *defPacketfilter }

// ApiRoutes returns all known Packetfilter Paths
func (Packetfilter) ApiRoutes() []string {
	return []string{
		"/api/objects/packetfilter/1to1nat/",
		"/api/objects/packetfilter/1to1nat/{ref}",
		"/api/objects/packetfilter/1to1nat/{ref}/usedby",
		"/api/objects/packetfilter/generic_proxy/",
		"/api/objects/packetfilter/generic_proxy/{ref}",
		"/api/objects/packetfilter/generic_proxy/{ref}/usedby",
		"/api/objects/packetfilter/group/",
		"/api/objects/packetfilter/group/{ref}",
		"/api/objects/packetfilter/group/{ref}/usedby",
		"/api/objects/packetfilter/loadbalance/",
		"/api/objects/packetfilter/loadbalance/{ref}",
		"/api/objects/packetfilter/loadbalance/{ref}/usedby",
		"/api/objects/packetfilter/mangle/",
		"/api/objects/packetfilter/mangle/{ref}",
		"/api/objects/packetfilter/mangle/{ref}/usedby",
		"/api/objects/packetfilter/masq/",
		"/api/objects/packetfilter/masq/{ref}",
		"/api/objects/packetfilter/masq/{ref}/usedby",
		"/api/objects/packetfilter/nat/",
		"/api/objects/packetfilter/nat/{ref}",
		"/api/objects/packetfilter/nat/{ref}/usedby",
		"/api/objects/packetfilter/packetfilter/",
		"/api/objects/packetfilter/packetfilter/{ref}",
		"/api/objects/packetfilter/packetfilter/{ref}/usedby",
		"/api/objects/packetfilter/ruleset/",
		"/api/objects/packetfilter/ruleset/{ref}",
		"/api/objects/packetfilter/ruleset/{ref}/usedby",
	}
}

// References returns the Packetfilter's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Packetfilter) References() []string {
	return []string{
		"REF_Packetfilter1to1Nat",
		"REF_PacketfilterGenericProxy",
		"REF_PacketfilterGroup",
		"REF_PacketfilterLoadbalance",
		"REF_PacketfilterMangle",
		"REF_PacketfilterMasq",
		"REF_PacketfilterNat",
		"REF_PacketfilterPacketfilter",
		"REF_PacketfilterRuleset",
	}
}

// Packetfilter1to1Nats is an Sophos Endpoint subType and implements sophos.RestObject
type Packetfilter1to1Nats []Packetfilter1to1Nat

// Packetfilter1to1Nat represents a UTM 1to1nat
type Packetfilter1to1Nat struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Service description: REF(service/*)
	Service string `json:"service"`
	// Source description: REF(network/*)
	Source string `json:"source"`
	// AutoPfIn description: REF(packetfilter/packetfilter)
	// AutoPfIn default value is ""
	AutoPfIn string `json:"auto_pf_in"`
	// Destination description: REF(network/*)
	Destination string `json:"destination"`
	// Group default value is ""
	Group string `json:"group"`
	// Log default value is false
	Log bool `json:"log"`
	// MapTo description: REF(network/network)
	MapTo string `json:"map_to"`
	// AutoPfrule default value is false
	AutoPfrule bool   `json:"auto_pfrule"`
	Comment    string `json:"comment"`
	// Mode can be one of: []string{"mapsrc", "mapdst"}
	Mode string `json:"mode"`
	Name string `json:"name"`
	// Status default value is false
	Status bool `json:"status"`
}

var _ sophos.RestGetter = &Packetfilter1to1Nat{}

// GetPath implements sophos.RestObject and returns the Packetfilter1to1Nats GET path
// Returns all available packetfilter/1to1nat objects
func (*Packetfilter1to1Nats) GetPath() string { return "/api/objects/packetfilter/1to1nat/" }

// RefRequired implements sophos.RestObject
func (*Packetfilter1to1Nats) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the Packetfilter1to1Nats GET path
// Returns all available 1to1nat types
func (p *Packetfilter1to1Nat) GetPath() string {
	return fmt.Sprintf("/api/objects/packetfilter/1to1nat/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *Packetfilter1to1Nat) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the Packetfilter1to1Nat DELETE path
// Creates or updates the complete object 1to1nat
func (*Packetfilter1to1Nat) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/1to1nat/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the Packetfilter1to1Nat PATCH path
// Changes to parts of the object 1to1nat types
func (*Packetfilter1to1Nat) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/1to1nat/%s", ref)
}

// PostPath implements sophos.RestObject and returns the Packetfilter1to1Nat POST path
// Create a new packetfilter/1to1nat object
func (*Packetfilter1to1Nat) PostPath() string {
	return "/api/objects/packetfilter/1to1nat/"
}

// PutPath implements sophos.RestObject and returns the Packetfilter1to1Nat PUT path
// Creates or updates the complete object 1to1nat
func (*Packetfilter1to1Nat) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/1to1nat/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*Packetfilter1to1Nat) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/1to1nat/%s/usedby", ref)
}

// PacketfilterGenericProxys is an Sophos Endpoint subType and implements sophos.RestObject
type PacketfilterGenericProxys []PacketfilterGenericProxy

// PacketfilterGenericProxy represents a UTM generic proxy rule
type PacketfilterGenericProxy struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Toservice description: REF(service/*)
	Toservice       string        `json:"toservice"`
	AllowedNetworks []interface{} `json:"allowed_networks"`
	Comment         string        `json:"comment"`
	// Ininterface description: REF(interface/*)
	Ininterface string `json:"ininterface"`
	Name        string `json:"name"`
	// Service description: REF(service/*)
	Service string `json:"service"`
	// Status default value is false
	Status bool `json:"status"`
	// Tohost description: REF(network/host), REF(network/dns_host), REF(network/availability_group)
	Tohost string `json:"tohost"`
}

var _ sophos.RestGetter = &PacketfilterGenericProxy{}

// GetPath implements sophos.RestObject and returns the PacketfilterGenericProxys GET path
// Returns all available packetfilter/generic_proxy objects
func (*PacketfilterGenericProxys) GetPath() string { return "/api/objects/packetfilter/generic_proxy/" }

// RefRequired implements sophos.RestObject
func (*PacketfilterGenericProxys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the PacketfilterGenericProxys GET path
// Returns all available generic_proxy types
func (p *PacketfilterGenericProxy) GetPath() string {
	return fmt.Sprintf("/api/objects/packetfilter/generic_proxy/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *PacketfilterGenericProxy) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the PacketfilterGenericProxy DELETE path
// Creates or updates the complete object generic_proxy
func (*PacketfilterGenericProxy) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/generic_proxy/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the PacketfilterGenericProxy PATCH path
// Changes to parts of the object generic_proxy types
func (*PacketfilterGenericProxy) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/generic_proxy/%s", ref)
}

// PostPath implements sophos.RestObject and returns the PacketfilterGenericProxy POST path
// Create a new packetfilter/generic_proxy object
func (*PacketfilterGenericProxy) PostPath() string {
	return "/api/objects/packetfilter/generic_proxy/"
}

// PutPath implements sophos.RestObject and returns the PacketfilterGenericProxy PUT path
// Creates or updates the complete object generic_proxy
func (*PacketfilterGenericProxy) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/generic_proxy/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*PacketfilterGenericProxy) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/generic_proxy/%s/usedby", ref)
}

// PacketfilterGroups is an Sophos Endpoint subType and implements sophos.RestObject
type PacketfilterGroups []PacketfilterGroup

// PacketfilterGroup represents a UTM group
type PacketfilterGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &PacketfilterGroup{}

// GetPath implements sophos.RestObject and returns the PacketfilterGroups GET path
// Returns all available packetfilter/group objects
func (*PacketfilterGroups) GetPath() string { return "/api/objects/packetfilter/group/" }

// RefRequired implements sophos.RestObject
func (*PacketfilterGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the PacketfilterGroups GET path
// Returns all available group types
func (p *PacketfilterGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/packetfilter/group/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *PacketfilterGroup) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the PacketfilterGroup DELETE path
// Creates or updates the complete object group
func (*PacketfilterGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the PacketfilterGroup PATCH path
// Changes to parts of the object group types
func (*PacketfilterGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the PacketfilterGroup POST path
// Create a new packetfilter/group object
func (*PacketfilterGroup) PostPath() string {
	return "/api/objects/packetfilter/group/"
}

// PutPath implements sophos.RestObject and returns the PacketfilterGroup PUT path
// Creates or updates the complete object group
func (*PacketfilterGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*PacketfilterGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/group/%s/usedby", ref)
}

// PacketfilterLoadbalances is an Sophos Endpoint subType and implements sophos.RestObject
type PacketfilterLoadbalances []PacketfilterLoadbalance

// PacketfilterLoadbalance represents a UTM load balancing rule
type PacketfilterLoadbalance struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Scheduler description: REF(scheduler/loadbalance)
	// Scheduler default value is ""
	Scheduler string `json:"scheduler"`
	// Status default value is false
	Status bool `json:"status"`
	// AutoPfIn description: REF(packetfilter/packetfilter)
	// AutoPfIn default value is ""
	AutoPfIn string `json:"auto_pf_in"`
	// Destination description: REF(network/*)
	Destination         string        `json:"destination"`
	DestinationNatGroup []interface{} `json:"destination_nat_group"`
	// ShutdownAddress default value is false
	ShutdownAddress bool `json:"shutdown_address"`
	// DestinationNatStatus description: (HASH)
	DestinationNatStatus interface{} `json:"destination_nat_status"`
	// DestinationNatStatus6 description: (HASH)
	DestinationNatStatus6 interface{} `json:"destination_nat_status6"`
	Name                  string      `json:"name"`
	// Service description: REF(service/*)
	Service string `json:"service"`
	// ShutdownCondition description: REF(condition/objref)
	// ShutdownCondition default value is ""
	ShutdownCondition string `json:"shutdown_condition"`
	// ShutdownOverride description: REF(override/objref)
	// ShutdownOverride default value is ""
	ShutdownOverride string `json:"shutdown_override"`
	// AutoPfrule default value is true
	AutoPfrule bool   `json:"auto_pfrule"`
	Comment    string `json:"comment"`
}

var _ sophos.RestGetter = &PacketfilterLoadbalance{}

// GetPath implements sophos.RestObject and returns the PacketfilterLoadbalances GET path
// Returns all available packetfilter/loadbalance objects
func (*PacketfilterLoadbalances) GetPath() string { return "/api/objects/packetfilter/loadbalance/" }

// RefRequired implements sophos.RestObject
func (*PacketfilterLoadbalances) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the PacketfilterLoadbalances GET path
// Returns all available loadbalance types
func (p *PacketfilterLoadbalance) GetPath() string {
	return fmt.Sprintf("/api/objects/packetfilter/loadbalance/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *PacketfilterLoadbalance) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the PacketfilterLoadbalance DELETE path
// Creates or updates the complete object loadbalance
func (*PacketfilterLoadbalance) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/loadbalance/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the PacketfilterLoadbalance PATCH path
// Changes to parts of the object loadbalance types
func (*PacketfilterLoadbalance) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/loadbalance/%s", ref)
}

// PostPath implements sophos.RestObject and returns the PacketfilterLoadbalance POST path
// Create a new packetfilter/loadbalance object
func (*PacketfilterLoadbalance) PostPath() string {
	return "/api/objects/packetfilter/loadbalance/"
}

// PutPath implements sophos.RestObject and returns the PacketfilterLoadbalance PUT path
// Creates or updates the complete object loadbalance
func (*PacketfilterLoadbalance) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/loadbalance/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*PacketfilterLoadbalance) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/loadbalance/%s/usedby", ref)
}

// PacketfilterMangles is an Sophos Endpoint subType and implements sophos.RestObject
type PacketfilterMangles []PacketfilterMangle

// PacketfilterMangle represents a UTM packet mangle rule
type PacketfilterMangle struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Destination description: REF(network/*)
	Destination string `json:"destination"`
	// Direction can be one of: []string{"in", "out"}
	Direction string `json:"direction"`
	Name      string `json:"name"`
	// Service description: REF(service/*)
	Service string `json:"service"`
	// Source description: REF(network/*)
	Source string `json:"source"`
	// Status default value is false
	Status  bool          `json:"status"`
	Action  []interface{} `json:"action"`
	Comment string        `json:"comment"`
}

var _ sophos.RestGetter = &PacketfilterMangle{}

// GetPath implements sophos.RestObject and returns the PacketfilterMangles GET path
// Returns all available packetfilter/mangle objects
func (*PacketfilterMangles) GetPath() string { return "/api/objects/packetfilter/mangle/" }

// RefRequired implements sophos.RestObject
func (*PacketfilterMangles) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the PacketfilterMangles GET path
// Returns all available mangle types
func (p *PacketfilterMangle) GetPath() string {
	return fmt.Sprintf("/api/objects/packetfilter/mangle/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *PacketfilterMangle) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the PacketfilterMangle DELETE path
// Creates or updates the complete object mangle
func (*PacketfilterMangle) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/mangle/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the PacketfilterMangle PATCH path
// Changes to parts of the object mangle types
func (*PacketfilterMangle) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/mangle/%s", ref)
}

// PostPath implements sophos.RestObject and returns the PacketfilterMangle POST path
// Create a new packetfilter/mangle object
func (*PacketfilterMangle) PostPath() string {
	return "/api/objects/packetfilter/mangle/"
}

// PutPath implements sophos.RestObject and returns the PacketfilterMangle PUT path
// Creates or updates the complete object mangle
func (*PacketfilterMangle) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/mangle/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*PacketfilterMangle) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/mangle/%s/usedby", ref)
}

// PacketfilterMasqs is an Sophos Endpoint subType and implements sophos.RestObject
type PacketfilterMasqs []PacketfilterMasq

// PacketfilterMasq is a generated Sophos object
type PacketfilterMasq struct {
	Locked                   string `json:"_locked"`
	Reference                string `json:"_ref"`
	ObjectType               string `json:"_type"`
	AdditionalAddress        string `json:"additional_address"`
	AdditionalAddressRestore string `json:"additional_address_restore"`
	Comment                  string `json:"comment"`
	Name                     string `json:"name"`
	Source                   string `json:"source"`
	SourceNatInterface       string `json:"source_nat_interface"`
	Status                   bool   `json:"status"`
}

var _ sophos.RestGetter = &PacketfilterMasq{}

// GetPath implements sophos.RestObject and returns the PacketfilterMasqs GET path
// Returns all available packetfilter/masq objects
func (*PacketfilterMasqs) GetPath() string { return "/api/objects/packetfilter/masq/" }

// RefRequired implements sophos.RestObject
func (*PacketfilterMasqs) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the PacketfilterMasqs GET path
// Returns all available masq types
func (p *PacketfilterMasq) GetPath() string {
	return fmt.Sprintf("/api/objects/packetfilter/masq/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *PacketfilterMasq) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the PacketfilterMasq DELETE path
// Creates or updates the complete object masq
func (*PacketfilterMasq) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/masq/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the PacketfilterMasq PATCH path
// Changes to parts of the object masq types
func (*PacketfilterMasq) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/masq/%s", ref)
}

// PostPath implements sophos.RestObject and returns the PacketfilterMasq POST path
// Create a new packetfilter/masq object
func (*PacketfilterMasq) PostPath() string {
	return "/api/objects/packetfilter/masq/"
}

// PutPath implements sophos.RestObject and returns the PacketfilterMasq PUT path
// Creates or updates the complete object masq
func (*PacketfilterMasq) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/masq/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*PacketfilterMasq) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/masq/%s/usedby", ref)
}

// GetType implements sophos.Object
func (p *PacketfilterMasq) GetType() string { return p.ObjectType }

// PacketfilterNats is an Sophos Endpoint subType and implements sophos.RestObject
type PacketfilterNats []PacketfilterNat

// PacketfilterNat is a generated Sophos object
type PacketfilterNat struct {
	Locked                string `json:"_locked"`
	Reference             string `json:"_ref"`
	ObjectType            string `json:"_type"`
	AutoPfIn              string `json:"auto_pf_in"`
	AutoPfrule            bool   `json:"auto_pfrule"`
	Comment               string `json:"comment"`
	Destination           string `json:"destination"`
	DestinationNatAddress string `json:"destination_nat_address"`
	DestinationNatService string `json:"destination_nat_service"`
	Group                 string `json:"group"`
	Ipsec                 bool   `json:"ipsec"`
	Log                   bool   `json:"log"`
	Mode                  string `json:"mode"`
	Name                  string `json:"name"`
	Service               string `json:"service"`
	Source                string `json:"source"`
	SourceNatAddress      string `json:"source_nat_address"`
	SourceNatService      string `json:"source_nat_service"`
	Status                bool   `json:"status"`
}

var _ sophos.RestGetter = &PacketfilterNat{}

// GetPath implements sophos.RestObject and returns the PacketfilterNats GET path
// Returns all available packetfilter/nat objects
func (*PacketfilterNats) GetPath() string { return "/api/objects/packetfilter/nat/" }

// RefRequired implements sophos.RestObject
func (*PacketfilterNats) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the PacketfilterNats GET path
// Returns all available nat types
func (p *PacketfilterNat) GetPath() string {
	return fmt.Sprintf("/api/objects/packetfilter/nat/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *PacketfilterNat) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the PacketfilterNat DELETE path
// Creates or updates the complete object nat
func (*PacketfilterNat) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/nat/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the PacketfilterNat PATCH path
// Changes to parts of the object nat types
func (*PacketfilterNat) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/nat/%s", ref)
}

// PostPath implements sophos.RestObject and returns the PacketfilterNat POST path
// Create a new packetfilter/nat object
func (*PacketfilterNat) PostPath() string {
	return "/api/objects/packetfilter/nat/"
}

// PutPath implements sophos.RestObject and returns the PacketfilterNat PUT path
// Creates or updates the complete object nat
func (*PacketfilterNat) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/nat/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*PacketfilterNat) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/nat/%s/usedby", ref)
}

// GetType implements sophos.Object
func (p *PacketfilterNat) GetType() string { return p.ObjectType }

// PacketfilterPacketfilters is an Sophos Endpoint subType and implements sophos.RestObject
type PacketfilterPacketfilters []PacketfilterPacketfilter

// PacketfilterPacketfilter is a generated Sophos object
type PacketfilterPacketfilter struct {
	Locked             string   `json:"_locked"`
	Reference          string   `json:"_ref"`
	ObjectType         string   `json:"_type"`
	Action             string   `json:"action"`
	Auto               bool     `json:"auto"`
	AutoType           string   `json:"auto_type"`
	Comment            string   `json:"comment"`
	Destinations       []string `json:"destinations"`
	Direction          string   `json:"direction"`
	Group              string   `json:"group"`
	Interface          string   `json:"interface"`
	Log                bool     `json:"log"`
	Name               string   `json:"name"`
	Services           []string `json:"services"`
	SourceMacAddresses string   `json:"source_mac_addresses"`
	Sources            []string `json:"sources"`
	Status             bool     `json:"status"`
	Time               string   `json:"time"`
}

var _ sophos.RestGetter = &PacketfilterPacketfilter{}

// GetPath implements sophos.RestObject and returns the PacketfilterPacketfilters GET path
// Returns all available packetfilter/packetfilter objects
func (*PacketfilterPacketfilters) GetPath() string { return "/api/objects/packetfilter/packetfilter/" }

// RefRequired implements sophos.RestObject
func (*PacketfilterPacketfilters) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the PacketfilterPacketfilters GET path
// Returns all available packetfilter types
func (p *PacketfilterPacketfilter) GetPath() string {
	return fmt.Sprintf("/api/objects/packetfilter/packetfilter/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *PacketfilterPacketfilter) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the PacketfilterPacketfilter DELETE path
// Creates or updates the complete object packetfilter
func (*PacketfilterPacketfilter) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/packetfilter/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the PacketfilterPacketfilter PATCH path
// Changes to parts of the object packetfilter types
func (*PacketfilterPacketfilter) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/packetfilter/%s", ref)
}

// PostPath implements sophos.RestObject and returns the PacketfilterPacketfilter POST path
// Create a new packetfilter/packetfilter object
func (*PacketfilterPacketfilter) PostPath() string {
	return "/api/objects/packetfilter/packetfilter/"
}

// PutPath implements sophos.RestObject and returns the PacketfilterPacketfilter PUT path
// Creates or updates the complete object packetfilter
func (*PacketfilterPacketfilter) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/packetfilter/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*PacketfilterPacketfilter) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/packetfilter/%s/usedby", ref)
}

// GetType implements sophos.Object
func (p *PacketfilterPacketfilter) GetType() string { return p.ObjectType }

// PacketfilterRulesets is an Sophos Endpoint subType and implements sophos.RestObject
type PacketfilterRulesets []PacketfilterRuleset

// PacketfilterRuleset represents a UTM ruleset
type PacketfilterRuleset struct {
	Locked     string        `json:"_locked"`
	ObjectType string        `json:"_type"`
	Reference  string        `json:"_ref"`
	Rules      []interface{} `json:"rules"`
	// RulesStatus description: (HASH)
	RulesStatus interface{} `json:"rules_status"`
	// Status default value is false
	Status  bool   `json:"status"`
	Comment string `json:"comment"`
	Name    string `json:"name"`
}

var _ sophos.RestGetter = &PacketfilterRuleset{}

// GetPath implements sophos.RestObject and returns the PacketfilterRulesets GET path
// Returns all available packetfilter/ruleset objects
func (*PacketfilterRulesets) GetPath() string { return "/api/objects/packetfilter/ruleset/" }

// RefRequired implements sophos.RestObject
func (*PacketfilterRulesets) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the PacketfilterRulesets GET path
// Returns all available ruleset types
func (p *PacketfilterRuleset) GetPath() string {
	return fmt.Sprintf("/api/objects/packetfilter/ruleset/%s", p.Reference)
}

// RefRequired implements sophos.RestObject
func (p *PacketfilterRuleset) RefRequired() (string, bool) { return p.Reference, true }

// DeletePath implements sophos.RestObject and returns the PacketfilterRuleset DELETE path
// Creates or updates the complete object ruleset
func (*PacketfilterRuleset) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/ruleset/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the PacketfilterRuleset PATCH path
// Changes to parts of the object ruleset types
func (*PacketfilterRuleset) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/ruleset/%s", ref)
}

// PostPath implements sophos.RestObject and returns the PacketfilterRuleset POST path
// Create a new packetfilter/ruleset object
func (*PacketfilterRuleset) PostPath() string {
	return "/api/objects/packetfilter/ruleset/"
}

// PutPath implements sophos.RestObject and returns the PacketfilterRuleset PUT path
// Creates or updates the complete object ruleset
func (*PacketfilterRuleset) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/ruleset/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*PacketfilterRuleset) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/packetfilter/ruleset/%s/usedby", ref)
}
