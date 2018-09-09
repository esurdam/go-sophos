package sophos

import "strings"

// A Reference is the connections between nodes and objects as well as between one object and another object.
// Each confd node and object has a list of attributes with pre defined types, where one of the types can be a
// reference. Please note however that you can’t create a reference in all cases. You can only make a reference
// in scenarios where nodes and objects are designed to be connected. Technically, references are strings that
// always start with the prefix “REF_”.
type Reference string

const (
	// RefServiceAny is a well-known Reference referring to any UTM Service
	RefServiceAny = "REF_ServiceAny"
	// RefNetworkAny is a well-known Reference referring to any UTM Network
	RefNetworkAny = "REF_NetworkAny"
	// RefNtpPool is a well-known Reference referring to the NTP Pool
	RefNtpPool = "REF_NtpPool"
	// RefDefaultSuperAdminGroup is a Reference to the UTM's Default Super Admin Group
	RefDefaultSuperAdminGroup = "REF_DefaultSuperAdminGroup"
	// RefDefaultServiceWebAdmin is a well-known Reference
	RefDefaultServiceWebAdmin = "REF_DefaultServiceWebadmin"
	// RefDefaultServiceSpamrelease is a well-known Reference
	RefDefaultServiceSpamrelease = "REF_DefaultServiceSpamrelease"
	// RefDefaultHTTPCFFBlockAction is a well-known Reference
	RefDefaultHTTPCFFBlockAction = "REF_DefaultHTTPCFFBlockAction"
	// RefDefaultHTTPPACFile is a well-known Reference
	RefDefaultHTTPPACFile = "REF_DefaultHTTPPACFile"
	// RefDefaultHTTPProfile is a well-known Reference
	RefDefaultHTTPProfile = "REF_DefaultHTTPProfile"
	// RefDefaultPPTPPool is a well-known Reference
	RefDefaultPPTPPool = "REF_DefaultPPTPPool"
	// RefDefaultSSLPool is a well-known Reference
	RefDefaultSSLPool = "REF_DefaultSSLPool"
)

const refPrefix = "REF_"

// IsReference returns true if the string has the prefex "REF_"
func IsReference(ref string) bool { return strings.HasPrefix(ref, refPrefix) }

// IsReference returns true if the string has the prefex "REF_"
func (r Reference) IsReference() bool { return IsReference(string(r)) }
