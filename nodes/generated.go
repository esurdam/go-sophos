package nodes

import (
	"bytes"
	"encoding/json"

	"github.com/esurdam/go-sophos"
)

// GetAccServer1AuthSecret gets the acc.server1.auth.secret node value
func GetAccServer1AuthSecret(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/acc.server1.auth.secret")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer1AuthSecret updates the acc.server1.auth.secret node value
func UpdateAccServer1AuthSecret(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server1.auth.secret", bytes.NewReader(byt))
	return
}

// GetAccServer1AuthStatus gets the acc.server1.auth.status node value
func GetAccServer1AuthStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/acc.server1.auth.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer1AuthStatus updates the acc.server1.auth.status node value
func UpdateAccServer1AuthStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server1.auth.status", bytes.NewReader(byt))
	return
}

// GetAccServer1Port gets the acc.server1.port node value
func GetAccServer1Port(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/acc.server1.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer1Port updates the acc.server1.port node value
func UpdateAccServer1Port(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server1.port", bytes.NewReader(byt))
	return
}

// GetAccServer1Roles gets the acc.server1.roles node value
func GetAccServer1Roles(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/acc.server1.roles")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer1Roles updates the acc.server1.roles node value
func UpdateAccServer1Roles(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server1.roles", bytes.NewReader(byt))
	return
}

// GetAccServer1Server gets the acc.server1.server node value
func GetAccServer1Server(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/acc.server1.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer1Server updates the acc.server1.server node value
func UpdateAccServer1Server(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server1.server", bytes.NewReader(byt))
	return
}

// GetAccServer2AuthSecret gets the acc.server2.auth.secret node value
func GetAccServer2AuthSecret(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/acc.server2.auth.secret")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer2AuthSecret updates the acc.server2.auth.secret node value
func UpdateAccServer2AuthSecret(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server2.auth.secret", bytes.NewReader(byt))
	return
}

// GetAccServer2AuthStatus gets the acc.server2.auth.status node value
func GetAccServer2AuthStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/acc.server2.auth.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer2AuthStatus updates the acc.server2.auth.status node value
func UpdateAccServer2AuthStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server2.auth.status", bytes.NewReader(byt))
	return
}

// GetAccServer2Port gets the acc.server2.port node value
func GetAccServer2Port(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/acc.server2.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer2Port updates the acc.server2.port node value
func UpdateAccServer2Port(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server2.port", bytes.NewReader(byt))
	return
}

// GetAccServer2Roles gets the acc.server2.roles node value
func GetAccServer2Roles(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/acc.server2.roles")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer2Roles updates the acc.server2.roles node value
func UpdateAccServer2Roles(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server2.roles", bytes.NewReader(byt))
	return
}

// GetAccServer2Server gets the acc.server2.server node value
func GetAccServer2Server(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/acc.server2.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer2Server updates the acc.server2.server node value
func UpdateAccServer2Server(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server2.server", bytes.NewReader(byt))
	return
}

// GetAccServer2Status gets the acc.server2.status node value
func GetAccServer2Status(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/acc.server2.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer2Status updates the acc.server2.status node value
func UpdateAccServer2Status(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server2.status", bytes.NewReader(byt))
	return
}

// GetAccSsoAdminGroup gets the acc.sso_admin_group node value
func GetAccSsoAdminGroup(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/acc.sso_admin_group")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccSsoAdminGroup updates the acc.sso_admin_group node value
func UpdateAccSsoAdminGroup(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.sso_admin_group", bytes.NewReader(byt))
	return
}

// GetAccSsoAuditorGroup gets the acc.sso_auditor_group node value
func GetAccSsoAuditorGroup(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/acc.sso_auditor_group")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccSsoAuditorGroup updates the acc.sso_auditor_group node value
func UpdateAccSsoAuditorGroup(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.sso_auditor_group", bytes.NewReader(byt))
	return
}

// GetAccStatus gets the acc.status node value
func GetAccStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/acc.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccStatus updates the acc.status node value
func UpdateAccStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.status", bytes.NewReader(byt))
	return
}

// GetAccdAccessAllowedAdmins gets the accd.access.allowed_admins node value
func GetAccdAccessAllowedAdmins(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/accd.access.allowed_admins")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdAccessAllowedAdmins updates the accd.access.allowed_admins node value
func UpdateAccdAccessAllowedAdmins(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.access.allowed_admins", bytes.NewReader(byt))
	return
}

// GetAccdAccessAllowedNetworks gets the accd.access.allowed_networks node value
func GetAccdAccessAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/accd.access.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdAccessAllowedNetworks updates the accd.access.allowed_networks node value
func UpdateAccdAccessAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.access.allowed_networks", bytes.NewReader(byt))
	return
}

// GetAccdAccessAllowedUsers gets the accd.access.allowed_users node value
func GetAccdAccessAllowedUsers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/accd.access.allowed_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdAccessAllowedUsers updates the accd.access.allowed_users node value
func UpdateAccdAccessAllowedUsers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.access.allowed_users", bytes.NewReader(byt))
	return
}

// GetAccdAccessCert gets the accd.access.cert node value
func GetAccdAccessCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/accd.access.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdAccessCert updates the accd.access.cert node value
func UpdateAccdAccessCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.access.cert", bytes.NewReader(byt))
	return
}

// GetAccdAccessPort gets the accd.access.port node value
func GetAccdAccessPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/accd.access.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdAccessPort updates the accd.access.port node value
func UpdateAccdAccessPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.access.port", bytes.NewReader(byt))
	return
}

// GetAccdDevicesAllowedNetworks gets the accd.devices.allowed_networks node value
func GetAccdDevicesAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/accd.devices.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdDevicesAllowedNetworks updates the accd.devices.allowed_networks node value
func UpdateAccdDevicesAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.devices.allowed_networks", bytes.NewReader(byt))
	return
}

// GetAccdDevicesAuthAuto gets the accd.devices.auth.auto node value
func GetAccdDevicesAuthAuto(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/accd.devices.auth.auto")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdDevicesAuthAuto updates the accd.devices.auth.auto node value
func UpdateAccdDevicesAuthAuto(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.devices.auth.auto", bytes.NewReader(byt))
	return
}

// GetAccdDevicesAuthSecret gets the accd.devices.auth.secret node value
func GetAccdDevicesAuthSecret(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/accd.devices.auth.secret")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdDevicesAuthSecret updates the accd.devices.auth.secret node value
func UpdateAccdDevicesAuthSecret(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.devices.auth.secret", bytes.NewReader(byt))
	return
}

// GetAccdDevicesAuthStatus gets the accd.devices.auth.status node value
func GetAccdDevicesAuthStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/accd.devices.auth.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdDevicesAuthStatus updates the accd.devices.auth.status node value
func UpdateAccdDevicesAuthStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.devices.auth.status", bytes.NewReader(byt))
	return
}

// GetAccdDevicesCert gets the accd.devices.cert node value
func GetAccdDevicesCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/accd.devices.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdDevicesCert updates the accd.devices.cert node value
func UpdateAccdDevicesCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.devices.cert", bytes.NewReader(byt))
	return
}

// GetAccdDevicesPort gets the accd.devices.port node value
func GetAccdDevicesPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/accd.devices.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdDevicesPort updates the accd.devices.port node value
func UpdateAccdDevicesPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.devices.port", bytes.NewReader(byt))
	return
}

// GetAccdGeneralAllowedNetworks gets the accd.general.allowed_networks node value
func GetAccdGeneralAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/accd.general.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdGeneralAllowedNetworks updates the accd.general.allowed_networks node value
func UpdateAccdGeneralAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.general.allowed_networks", bytes.NewReader(byt))
	return
}

// GetAccdGeneralCert gets the accd.general.cert node value
func GetAccdGeneralCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/accd.general.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdGeneralCert updates the accd.general.cert node value
func UpdateAccdGeneralCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.general.cert", bytes.NewReader(byt))
	return
}

// GetAccdGeneralLanguage gets the accd.general.language node value
func GetAccdGeneralLanguage(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/accd.general.language")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdGeneralLanguage updates the accd.general.language node value
func UpdateAccdGeneralLanguage(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.general.language", bytes.NewReader(byt))
	return
}

// GetAccdGeneralPort gets the accd.general.port node value
func GetAccdGeneralPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/accd.general.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdGeneralPort updates the accd.general.port node value
func UpdateAccdGeneralPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.general.port", bytes.NewReader(byt))
	return
}

// GetAccdGeneralTimeout gets the accd.general.timeout node value
func GetAccdGeneralTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/accd.general.timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdGeneralTimeout updates the accd.general.timeout node value
func UpdateAccdGeneralTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.general.timeout", bytes.NewReader(byt))
	return
}

// GetAccountingIpfixConnections gets the accounting.ipfix.connections node value
func GetAccountingIpfixConnections(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/accounting.ipfix.connections")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccountingIpfixConnections updates the accounting.ipfix.connections node value
func UpdateAccountingIpfixConnections(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accounting.ipfix.connections", bytes.NewReader(byt))
	return
}

// GetAccountingIpfixStatus gets the accounting.ipfix.status node value
func GetAccountingIpfixStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/accounting.ipfix.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccountingIpfixStatus updates the accounting.ipfix.status node value
func UpdateAccountingIpfixStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accounting.ipfix.status", bytes.NewReader(byt))
	return
}

// GetAfcControlledNetworks gets the afc.controlled_networks node value
func GetAfcControlledNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/afc.controlled_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcControlledNetworks updates the afc.controlled_networks node value
func UpdateAfcControlledNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.controlled_networks", bytes.NewReader(byt))
	return
}

// GetAfcHiddenSkip gets the afc.hidden_skip node value
func GetAfcHiddenSkip(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/afc.hidden_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcHiddenSkip updates the afc.hidden_skip node value
func UpdateAfcHiddenSkip(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.hidden_skip", bytes.NewReader(byt))
	return
}

// GetAfcHttpRedirectUrl gets the afc.http_redirect_url node value
func GetAfcHttpRedirectUrl(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/afc.http_redirect_url")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcHttpRedirectUrl updates the afc.http_redirect_url node value
func UpdateAfcHttpRedirectUrl(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.http_redirect_url", bytes.NewReader(byt))
	return
}

// GetAfcLog gets the afc.log node value
func GetAfcLog(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/afc.log")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcLog updates the afc.log node value
func UpdateAfcLog(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.log", bytes.NewReader(byt))
	return
}

// GetAfcNfqueueLength gets the afc.nfqueue_length node value
func GetAfcNfqueueLength(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/afc.nfqueue_length")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcNfqueueLength updates the afc.nfqueue_length node value
func UpdateAfcNfqueueLength(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.nfqueue_length", bytes.NewReader(byt))
	return
}

// GetAfcNumQueues gets the afc.num_queues node value
func GetAfcNumQueues(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/afc.num_queues")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcNumQueues updates the afc.num_queues node value
func UpdateAfcNumQueues(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.num_queues", bytes.NewReader(byt))
	return
}

// GetAfcRules gets the afc.rules node value
func GetAfcRules(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/afc.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcRules updates the afc.rules node value
func UpdateAfcRules(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.rules", bytes.NewReader(byt))
	return
}

// GetAfcStatus gets the afc.status node value
func GetAfcStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/afc.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcStatus updates the afc.status node value
func UpdateAfcStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.status", bytes.NewReader(byt))
	return
}

// GetAfcSubmitUnknownTrafficData gets the afc.submit_unknown_traffic_data node value
func GetAfcSubmitUnknownTrafficData(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/afc.submit_unknown_traffic_data")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcSubmitUnknownTrafficData updates the afc.submit_unknown_traffic_data node value
func UpdateAfcSubmitUnknownTrafficData(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.submit_unknown_traffic_data", bytes.NewReader(byt))
	return
}

// GetAfcTransparentSkip gets the afc.transparent_skip node value
func GetAfcTransparentSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/afc.transparent_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcTransparentSkip updates the afc.transparent_skip node value
func UpdateAfcTransparentSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.transparent_skip", bytes.NewReader(byt))
	return
}

// GetAmazonVpcAutoPfrule gets the amazon_vpc.auto_pfrule node value
func GetAmazonVpcAutoPfrule(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/amazon_vpc.auto_pfrule")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAmazonVpcAutoPfrule updates the amazon_vpc.auto_pfrule node value
func UpdateAmazonVpcAutoPfrule(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/amazon_vpc.auto_pfrule", bytes.NewReader(byt))
	return
}

// GetAmazonVpcConnections gets the amazon_vpc.connections node value
func GetAmazonVpcConnections(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/amazon_vpc.connections")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAmazonVpcConnections updates the amazon_vpc.connections node value
func UpdateAmazonVpcConnections(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/amazon_vpc.connections", bytes.NewReader(byt))
	return
}

// GetAmazonVpcNetworks gets the amazon_vpc.networks node value
func GetAmazonVpcNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/amazon_vpc.networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAmazonVpcNetworks updates the amazon_vpc.networks node value
func UpdateAmazonVpcNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/amazon_vpc.networks", bytes.NewReader(byt))
	return
}

// GetAmazonVpcStatus gets the amazon_vpc.status node value
func GetAmazonVpcStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/amazon_vpc.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAmazonVpcStatus updates the amazon_vpc.status node value
func UpdateAmazonVpcStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/amazon_vpc.status", bytes.NewReader(byt))
	return
}

// GetAptpPolicy gets the aptp.policy node value
func GetAptpPolicy(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/aptp.policy")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAptpPolicy updates the aptp.policy node value
func UpdateAptpPolicy(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/aptp.policy", bytes.NewReader(byt))
	return
}

// GetAptpRuleModifiers gets the aptp.rule_modifiers node value
func GetAptpRuleModifiers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/aptp.rule_modifiers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAptpRuleModifiers updates the aptp.rule_modifiers node value
func UpdateAptpRuleModifiers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/aptp.rule_modifiers", bytes.NewReader(byt))
	return
}

// GetAptpStatus gets the aptp.status node value
func GetAptpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/aptp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAptpStatus updates the aptp.status node value
func UpdateAptpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/aptp.status", bytes.NewReader(byt))
	return
}

// GetAptpTransparentSkip gets the aptp.transparent_skip node value
func GetAptpTransparentSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/aptp.transparent_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAptpTransparentSkip updates the aptp.transparent_skip node value
func UpdateAptpTransparentSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/aptp.transparent_skip", bytes.NewReader(byt))
	return
}

// GetArmLicensedIp gets the arm.licensed_ip node value
func GetArmLicensedIp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.licensed_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmLicensedIp updates the arm.licensed_ip node value
func UpdateArmLicensedIp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.licensed_ip", bytes.NewReader(byt))
	return
}

// GetArmRemoteHost gets the arm.remote.host node value
func GetArmRemoteHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.remote.host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteHost updates the arm.remote.host node value
func UpdateArmRemoteHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.host", bytes.NewReader(byt))
	return
}

// GetArmRemoteMethod gets the arm.remote.method node value
func GetArmRemoteMethod(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.remote.method")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteMethod updates the arm.remote.method node value
func UpdateArmRemoteMethod(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.method", bytes.NewReader(byt))
	return
}

// GetArmRemoteSmbPassword gets the arm.remote.smb_password node value
func GetArmRemoteSmbPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.remote.smb_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteSmbPassword updates the arm.remote.smb_password node value
func UpdateArmRemoteSmbPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.smb_password", bytes.NewReader(byt))
	return
}

// GetArmRemoteSmbShare gets the arm.remote.smb_share node value
func GetArmRemoteSmbShare(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.remote.smb_share")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteSmbShare updates the arm.remote.smb_share node value
func UpdateArmRemoteSmbShare(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.smb_share", bytes.NewReader(byt))
	return
}

// GetArmRemoteSmbUser gets the arm.remote.smb_user node value
func GetArmRemoteSmbUser(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.remote.smb_user")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteSmbUser updates the arm.remote.smb_user node value
func UpdateArmRemoteSmbUser(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.smb_user", bytes.NewReader(byt))
	return
}

// GetArmRemoteStatus gets the arm.remote.status node value
func GetArmRemoteStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/arm.remote.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteStatus updates the arm.remote.status node value
func UpdateArmRemoteStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.status", bytes.NewReader(byt))
	return
}

// GetArmRemoteSyslogService gets the arm.remote.syslog_service node value
func GetArmRemoteSyslogService(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.remote.syslog_service")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteSyslogService updates the arm.remote.syslog_service node value
func UpdateArmRemoteSyslogService(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.syslog_service", bytes.NewReader(byt))
	return
}

// GetArmStatus gets the arm.status node value
func GetArmStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/arm.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmStatus updates the arm.status node value
func UpdateArmStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.status", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoForceUtf8Sync gets the auth.ad_sso.force_utf8_sync node value
func GetAuthAdSsoForceUtf8Sync(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.force_utf8_sync")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoForceUtf8Sync updates the auth.ad_sso.force_utf8_sync node value
func UpdateAuthAdSsoForceUtf8Sync(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.force_utf8_sync", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoJoinresult gets the auth.ad_sso.joinresult node value
func GetAuthAdSsoJoinresult(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.joinresult")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoJoinresult updates the auth.ad_sso.joinresult node value
func UpdateAuthAdSsoJoinresult(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.joinresult", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoLoadbalancerFqdn gets the auth.ad_sso.loadbalancer_fqdn node value
func GetAuthAdSsoLoadbalancerFqdn(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.loadbalancer_fqdn")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoLoadbalancerFqdn updates the auth.ad_sso.loadbalancer_fqdn node value
func UpdateAuthAdSsoLoadbalancerFqdn(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.loadbalancer_fqdn", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoNtlmv2Auth gets the auth.ad_sso.ntlmv2_auth node value
func GetAuthAdSsoNtlmv2Auth(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.ntlmv2_auth")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoNtlmv2Auth updates the auth.ad_sso.ntlmv2_auth node value
func UpdateAuthAdSsoNtlmv2Auth(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.ntlmv2_auth", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSecrets gets the auth.ad_sso.secrets node value
func GetAuthAdSsoSecrets(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.secrets")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSecrets updates the auth.ad_sso.secrets node value
func UpdateAuthAdSsoSecrets(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.secrets", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSmbconf gets the auth.ad_sso.smbconf node value
func GetAuthAdSsoSmbconf(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.smbconf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSmbconf updates the auth.ad_sso.smbconf node value
func UpdateAuthAdSsoSmbconf(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.smbconf", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoDomain gets the auth.ad_sso.sso_domain node value
func GetAuthAdSsoSsoDomain(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_domain")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoDomain updates the auth.ad_sso.sso_domain node value
func UpdateAuthAdSsoSsoDomain(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_domain", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoNetbiosDomain gets the auth.ad_sso.sso_netbios_domain node value
func GetAuthAdSsoSsoNetbiosDomain(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_netbios_domain")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoNetbiosDomain updates the auth.ad_sso.sso_netbios_domain node value
func UpdateAuthAdSsoSsoNetbiosDomain(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_netbios_domain", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoNetbiosHost gets the auth.ad_sso.sso_netbios_host node value
func GetAuthAdSsoSsoNetbiosHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_netbios_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoNetbiosHost updates the auth.ad_sso.sso_netbios_host node value
func UpdateAuthAdSsoSsoNetbiosHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_netbios_host", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoPassword gets the auth.ad_sso.sso_password node value
func GetAuthAdSsoSsoPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoPassword updates the auth.ad_sso.sso_password node value
func UpdateAuthAdSsoSsoPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_password", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoServer gets the auth.ad_sso.sso_server node value
func GetAuthAdSsoSsoServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoServer updates the auth.ad_sso.sso_server node value
func UpdateAuthAdSsoSsoServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_server", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoStatus gets the auth.ad_sso.sso_status node value
func GetAuthAdSsoSsoStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoStatus updates the auth.ad_sso.sso_status node value
func UpdateAuthAdSsoSsoStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_status", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoSync gets the auth.ad_sso.sso_sync node value
func GetAuthAdSsoSsoSync(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_sync")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoSync updates the auth.ad_sso.sso_sync node value
func UpdateAuthAdSsoSsoSync(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_sync", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoUsername gets the auth.ad_sso.sso_username node value
func GetAuthAdSsoSsoUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoUsername updates the auth.ad_sso.sso_username node value
func UpdateAuthAdSsoSsoUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_username", bytes.NewReader(byt))
	return
}

// GetAuthApiTokens gets the auth.api_tokens node value
func GetAuthApiTokens(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/auth.api_tokens")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthApiTokens updates the auth.api_tokens node value
func UpdateAuthApiTokens(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.api_tokens", bytes.NewReader(byt))
	return
}

// GetAuthAutoAddToFacility gets the auth.auto_add_to_facility node value
func GetAuthAutoAddToFacility(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/auth.auto_add_to_facility")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAutoAddToFacility updates the auth.auto_add_to_facility node value
func UpdateAuthAutoAddToFacility(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.auto_add_to_facility", bytes.NewReader(byt))
	return
}

// GetAuthAutoAddUsers gets the auth.auto_add_users node value
func GetAuthAutoAddUsers(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.auto_add_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAutoAddUsers updates the auth.auto_add_users node value
func UpdateAuthAutoAddUsers(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.auto_add_users", bytes.NewReader(byt))
	return
}

// GetAuthBlockAttempts gets the auth.block.attempts node value
func GetAuthBlockAttempts(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.block.attempts")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthBlockAttempts updates the auth.block.attempts node value
func UpdateAuthBlockAttempts(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.block.attempts", bytes.NewReader(byt))
	return
}

// GetAuthBlockFacilities gets the auth.block.facilities node value
func GetAuthBlockFacilities(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/auth.block.facilities")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthBlockFacilities updates the auth.block.facilities node value
func UpdateAuthBlockFacilities(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.block.facilities", bytes.NewReader(byt))
	return
}

// GetAuthBlockLockout gets the auth.block.lockout node value
func GetAuthBlockLockout(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.block.lockout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthBlockLockout updates the auth.block.lockout node value
func UpdateAuthBlockLockout(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.block.lockout", bytes.NewReader(byt))
	return
}

// GetAuthBlockNever gets the auth.block.never node value
func GetAuthBlockNever(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/auth.block.never")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthBlockNever updates the auth.block.never node value
func UpdateAuthBlockNever(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.block.never", bytes.NewReader(byt))
	return
}

// GetAuthBlockSeconds gets the auth.block.seconds node value
func GetAuthBlockSeconds(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.block.seconds")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthBlockSeconds updates the auth.block.seconds node value
func UpdateAuthBlockSeconds(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.block.seconds", bytes.NewReader(byt))
	return
}

// GetAuthCacheLifetime gets the auth.cache_lifetime node value
func GetAuthCacheLifetime(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.cache_lifetime")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthCacheLifetime updates the auth.cache_lifetime node value
func UpdateAuthCacheLifetime(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.cache_lifetime", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoEmConflict gets the auth.edir_sso.em_conflict node value
func GetAuthEdirSsoEmConflict(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.em_conflict")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoEmConflict updates the auth.edir_sso.em_conflict node value
func UpdateAuthEdirSsoEmConflict(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.em_conflict", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoEmSocketTimeout gets the auth.edir_sso.em_socket_timeout node value
func GetAuthEdirSsoEmSocketTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.em_socket_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoEmSocketTimeout updates the auth.edir_sso.em_socket_timeout node value
func UpdateAuthEdirSsoEmSocketTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.em_socket_timeout", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoEmVerifyLogout gets the auth.edir_sso.em_verify_logout node value
func GetAuthEdirSsoEmVerifyLogout(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.em_verify_logout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoEmVerifyLogout updates the auth.edir_sso.em_verify_logout node value
func UpdateAuthEdirSsoEmVerifyLogout(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.em_verify_logout", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoSsoAuaSearchIp gets the auth.edir_sso.sso_aua_search_ip node value
func GetAuthEdirSsoSsoAuaSearchIp(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.sso_aua_search_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoSsoAuaSearchIp updates the auth.edir_sso.sso_aua_search_ip node value
func UpdateAuthEdirSsoSsoAuaSearchIp(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.sso_aua_search_ip", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoSsoMode gets the auth.edir_sso.sso_mode node value
func GetAuthEdirSsoSsoMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.sso_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoSsoMode updates the auth.edir_sso.sso_mode node value
func UpdateAuthEdirSsoSsoMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.sso_mode", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoSsoServer gets the auth.edir_sso.sso_server node value
func GetAuthEdirSsoSsoServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.sso_server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoSsoServer updates the auth.edir_sso.sso_server node value
func UpdateAuthEdirSsoSsoServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.sso_server", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoSyncInterval gets the auth.edir_sso.sync_interval node value
func GetAuthEdirSsoSyncInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.sync_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoSyncInterval updates the auth.edir_sso.sync_interval node value
func UpdateAuthEdirSsoSyncInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.sync_interval", bytes.NewReader(byt))
	return
}

// GetAuthOtpAutoCreateToken gets the auth.otp.auto_create_token node value
func GetAuthOtpAutoCreateToken(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.otp.auto_create_token")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpAutoCreateToken updates the auth.otp.auto_create_token node value
func UpdateAuthOtpAutoCreateToken(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.auto_create_token", bytes.NewReader(byt))
	return
}

// GetAuthOtpAutoTokenDigest gets the auth.otp.auto_token_digest node value
func GetAuthOtpAutoTokenDigest(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.otp.auto_token_digest")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpAutoTokenDigest updates the auth.otp.auto_token_digest node value
func UpdateAuthOtpAutoTokenDigest(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.auto_token_digest", bytes.NewReader(byt))
	return
}

// GetAuthOtpDefaultTimestep gets the auth.otp.default_timestep node value
func GetAuthOtpDefaultTimestep(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.otp.default_timestep")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpDefaultTimestep updates the auth.otp.default_timestep node value
func UpdateAuthOtpDefaultTimestep(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.default_timestep", bytes.NewReader(byt))
	return
}

// GetAuthOtpFacilities gets the auth.otp.facilities node value
func GetAuthOtpFacilities(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/auth.otp.facilities")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpFacilities updates the auth.otp.facilities node value
func UpdateAuthOtpFacilities(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.facilities", bytes.NewReader(byt))
	return
}

// GetAuthOtpMaxInitTimestepDiff gets the auth.otp.max_init_timestep_diff node value
func GetAuthOtpMaxInitTimestepDiff(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.otp.max_init_timestep_diff")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpMaxInitTimestepDiff updates the auth.otp.max_init_timestep_diff node value
func UpdateAuthOtpMaxInitTimestepDiff(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.max_init_timestep_diff", bytes.NewReader(byt))
	return
}

// GetAuthOtpMaxTimestepDiff gets the auth.otp.max_timestep_diff node value
func GetAuthOtpMaxTimestepDiff(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.otp.max_timestep_diff")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpMaxTimestepDiff updates the auth.otp.max_timestep_diff node value
func UpdateAuthOtpMaxTimestepDiff(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.max_timestep_diff", bytes.NewReader(byt))
	return
}

// GetAuthOtpRequireAllUsers gets the auth.otp.require_all_users node value
func GetAuthOtpRequireAllUsers(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.otp.require_all_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpRequireAllUsers updates the auth.otp.require_all_users node value
func UpdateAuthOtpRequireAllUsers(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.require_all_users", bytes.NewReader(byt))
	return
}

// GetAuthOtpRequiredUsers gets the auth.otp.required_users node value
func GetAuthOtpRequiredUsers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/auth.otp.required_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpRequiredUsers updates the auth.otp.required_users node value
func UpdateAuthOtpRequiredUsers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.required_users", bytes.NewReader(byt))
	return
}

// GetAuthOtpStatus gets the auth.otp.status node value
func GetAuthOtpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.otp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpStatus updates the auth.otp.status node value
func UpdateAuthOtpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.status", bytes.NewReader(byt))
	return
}

// GetAuthServers gets the auth.servers node value
func GetAuthServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/auth.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthServers updates the auth.servers node value
func UpdateAuthServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.servers", bytes.NewReader(byt))
	return
}

// GetAuthUpdateBackendGroupMembersDebug gets the auth.update_backend_group_members.debug node value
func GetAuthUpdateBackendGroupMembersDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.update_backend_group_members.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthUpdateBackendGroupMembersDebug updates the auth.update_backend_group_members.debug node value
func UpdateAuthUpdateBackendGroupMembersDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.update_backend_group_members.debug", bytes.NewReader(byt))
	return
}

// GetAuthUpdateBackendGroupMembersStatus gets the auth.update_backend_group_members.status node value
func GetAuthUpdateBackendGroupMembersStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.update_backend_group_members.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthUpdateBackendGroupMembersStatus updates the auth.update_backend_group_members.status node value
func UpdateAuthUpdateBackendGroupMembersStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.update_backend_group_members.status", bytes.NewReader(byt))
	return
}

// GetAweAllowedInterfaces gets the awe.allowed_interfaces node value
func GetAweAllowedInterfaces(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/awe.allowed_interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweAllowedInterfaces updates the awe.allowed_interfaces node value
func UpdateAweAllowedInterfaces(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.allowed_interfaces", bytes.NewReader(byt))
	return
}

// GetAweClients gets the awe.clients node value
func GetAweClients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/awe.clients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweClients updates the awe.clients node value
func UpdateAweClients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.clients", bytes.NewReader(byt))
	return
}

// GetAweDevices gets the awe.devices node value
func GetAweDevices(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/awe.devices")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweDevices updates the awe.devices node value
func UpdateAweDevices(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.devices", bytes.NewReader(byt))
	return
}

// GetAweGlobalApAutoaccept gets the awe.global.ap_autoaccept node value
func GetAweGlobalApAutoaccept(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/awe.global.ap_autoaccept")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalApAutoaccept updates the awe.global.ap_autoaccept node value
func UpdateAweGlobalApAutoaccept(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.ap_autoaccept", bytes.NewReader(byt))
	return
}

// GetAweGlobalApDebuglevel gets the awe.global.ap_debuglevel node value
func GetAweGlobalApDebuglevel(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/awe.global.ap_debuglevel")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalApDebuglevel updates the awe.global.ap_debuglevel node value
func UpdateAweGlobalApDebuglevel(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.ap_debuglevel", bytes.NewReader(byt))
	return
}

// GetAweGlobalApSoftlimit gets the awe.global.ap_softlimit node value
func GetAweGlobalApSoftlimit(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/awe.global.ap_softlimit")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalApSoftlimit updates the awe.global.ap_softlimit node value
func UpdateAweGlobalApSoftlimit(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.ap_softlimit", bytes.NewReader(byt))
	return
}

// GetAweGlobalApVlantag gets the awe.global.ap_vlantag node value
func GetAweGlobalApVlantag(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/awe.global.ap_vlantag")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalApVlantag updates the awe.global.ap_vlantag node value
func UpdateAweGlobalApVlantag(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.ap_vlantag", bytes.NewReader(byt))
	return
}

// GetAweGlobalAweStatus gets the awe.global.awe_status node value
func GetAweGlobalAweStatus(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/awe.global.awe_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalAweStatus updates the awe.global.awe_status node value
func UpdateAweGlobalAweStatus(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.awe_status", bytes.NewReader(byt))
	return
}

// GetAweGlobalBridgeUpdateKickout gets the awe.global.bridge_update_kickout node value
func GetAweGlobalBridgeUpdateKickout(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/awe.global.bridge_update_kickout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalBridgeUpdateKickout updates the awe.global.bridge_update_kickout node value
func UpdateAweGlobalBridgeUpdateKickout(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.bridge_update_kickout", bytes.NewReader(byt))
	return
}

// GetAweGlobalInitialSetup gets the awe.global.initial_setup node value
func GetAweGlobalInitialSetup(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/awe.global.initial_setup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalInitialSetup updates the awe.global.initial_setup node value
func UpdateAweGlobalInitialSetup(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.initial_setup", bytes.NewReader(byt))
	return
}

// GetAweGlobalLogLevel gets the awe.global.log_level node value
func GetAweGlobalLogLevel(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/awe.global.log_level")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalLogLevel updates the awe.global.log_level node value
func UpdateAweGlobalLogLevel(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.log_level", bytes.NewReader(byt))
	return
}

// GetAweGlobalMagicIp gets the awe.global.magic_ip node value
func GetAweGlobalMagicIp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/awe.global.magic_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalMagicIp updates the awe.global.magic_ip node value
func UpdateAweGlobalMagicIp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.magic_ip", bytes.NewReader(byt))
	return
}

// GetAweGlobalNotificationTimeout gets the awe.global.notification_timeout node value
func GetAweGlobalNotificationTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/awe.global.notification_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalNotificationTimeout updates the awe.global.notification_timeout node value
func UpdateAweGlobalNotificationTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.notification_timeout", bytes.NewReader(byt))
	return
}

// GetAweGlobalRadiusConf gets the awe.global.radius_conf node value
func GetAweGlobalRadiusConf(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/awe.global.radius_conf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalRadiusConf updates the awe.global.radius_conf node value
func UpdateAweGlobalRadiusConf(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.radius_conf", bytes.NewReader(byt))
	return
}

// GetAweGlobalRootpw gets the awe.global.rootpw node value
func GetAweGlobalRootpw(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/awe.global.rootpw")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalRootpw updates the awe.global.rootpw node value
func UpdateAweGlobalRootpw(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.rootpw", bytes.NewReader(byt))
	return
}

// GetAweGlobalStayOnline gets the awe.global.stay_online node value
func GetAweGlobalStayOnline(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/awe.global.stay_online")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalStayOnline updates the awe.global.stay_online node value
func UpdateAweGlobalStayOnline(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.stay_online", bytes.NewReader(byt))
	return
}

// GetAweGlobalStoreBssStats gets the awe.global.store_bss_stats node value
func GetAweGlobalStoreBssStats(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/awe.global.store_bss_stats")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalStoreBssStats updates the awe.global.store_bss_stats node value
func UpdateAweGlobalStoreBssStats(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.store_bss_stats", bytes.NewReader(byt))
	return
}

// GetAweGlobalTunnelIdOffset gets the awe.global.tunnel_id_offset node value
func GetAweGlobalTunnelIdOffset(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/awe.global.tunnel_id_offset")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalTunnelIdOffset updates the awe.global.tunnel_id_offset node value
func UpdateAweGlobalTunnelIdOffset(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.tunnel_id_offset", bytes.NewReader(byt))
	return
}

// GetAweGlobalVlantagging gets the awe.global.vlantagging node value
func GetAweGlobalVlantagging(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/awe.global.vlantagging")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalVlantagging updates the awe.global.vlantagging node value
func UpdateAweGlobalVlantagging(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.vlantagging", bytes.NewReader(byt))
	return
}

// GetAweNetworks gets the awe.networks node value
func GetAweNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/awe.networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweNetworks updates the awe.networks node value
func UpdateAweNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.networks", bytes.NewReader(byt))
	return
}

// GetAwscliProfiles gets the awscli.profiles node value
func GetAwscliProfiles(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/awscli.profiles")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAwscliProfiles updates the awscli.profiles node value
func UpdateAwscliProfiles(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awscli.profiles", bytes.NewReader(byt))
	return
}

// GetBackupEncryption gets the backup.encryption node value
func GetBackupEncryption(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/backup.encryption")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateBackupEncryption updates the backup.encryption node value
func UpdateBackupEncryption(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/backup.encryption", bytes.NewReader(byt))
	return
}

// GetBackupInterval gets the backup.interval node value
func GetBackupInterval(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/backup.interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateBackupInterval updates the backup.interval node value
func UpdateBackupInterval(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/backup.interval", bytes.NewReader(byt))
	return
}

// GetBackupMaxBackups gets the backup.max_backups node value
func GetBackupMaxBackups(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/backup.max_backups")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateBackupMaxBackups updates the backup.max_backups node value
func UpdateBackupMaxBackups(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/backup.max_backups", bytes.NewReader(byt))
	return
}

// GetBackupPassword gets the backup.password node value
func GetBackupPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/backup.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateBackupPassword updates the backup.password node value
func UpdateBackupPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/backup.password", bytes.NewReader(byt))
	return
}

// GetBackupRecipients gets the backup.recipients node value
func GetBackupRecipients(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/backup.recipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateBackupRecipients updates the backup.recipients node value
func UpdateBackupRecipients(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/backup.recipients", bytes.NewReader(byt))
	return
}

// GetBackupStatus gets the backup.status node value
func GetBackupStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/backup.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateBackupStatus updates the backup.status node value
func UpdateBackupStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/backup.status", bytes.NewReader(byt))
	return
}

// GetCaCaGost gets the ca.ca_gost node value
func GetCaCaGost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ca.ca_gost")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaCaGost updates the ca.ca_gost node value
func UpdateCaCaGost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.ca_gost", bytes.NewReader(byt))
	return
}

// GetCaCaIpsec gets the ca.ca_ipsec node value
func GetCaCaIpsec(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ca.ca_ipsec")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaCaIpsec updates the ca.ca_ipsec node value
func UpdateCaCaIpsec(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.ca_ipsec", bytes.NewReader(byt))
	return
}

// GetCaCaProxies gets the ca.ca_proxies node value
func GetCaCaProxies(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ca.ca_proxies")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaCaProxies updates the ca.ca_proxies node value
func UpdateCaCaProxies(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.ca_proxies", bytes.NewReader(byt))
	return
}

// GetCaCaRed gets the ca.ca_red node value
func GetCaCaRed(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ca.ca_red")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaCaRed updates the ca.ca_red node value
func UpdateCaCaRed(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.ca_red", bytes.NewReader(byt))
	return
}

// GetCaDefKeysize gets the ca.def_keysize node value
func GetCaDefKeysize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ca.def_keysize")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaDefKeysize updates the ca.def_keysize node value
func UpdateCaDefKeysize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.def_keysize", bytes.NewReader(byt))
	return
}

// GetCaGlobalCasEmailEncryptionTrustNewCas gets the ca.global_cas.email_encryption.trust_new_cas node value
func GetCaGlobalCasEmailEncryptionTrustNewCas(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ca.global_cas.email_encryption.trust_new_cas")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaGlobalCasEmailEncryptionTrustNewCas updates the ca.global_cas.email_encryption.trust_new_cas node value
func UpdateCaGlobalCasEmailEncryptionTrustNewCas(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.global_cas.email_encryption.trust_new_cas", bytes.NewReader(byt))
	return
}

// GetCaGlobalCasEmailEncryptionTrusted gets the ca.global_cas.email_encryption.trusted node value
func GetCaGlobalCasEmailEncryptionTrusted(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ca.global_cas.email_encryption.trusted")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaGlobalCasEmailEncryptionTrusted updates the ca.global_cas.email_encryption.trusted node value
func UpdateCaGlobalCasEmailEncryptionTrusted(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.global_cas.email_encryption.trusted", bytes.NewReader(byt))
	return
}

// GetCaGlobalCasEmailEncryptionUntrusted gets the ca.global_cas.email_encryption.untrusted node value
func GetCaGlobalCasEmailEncryptionUntrusted(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ca.global_cas.email_encryption.untrusted")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaGlobalCasEmailEncryptionUntrusted updates the ca.global_cas.email_encryption.untrusted node value
func UpdateCaGlobalCasEmailEncryptionUntrusted(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.global_cas.email_encryption.untrusted", bytes.NewReader(byt))
	return
}

// GetCaGlobalCasHttpProxyTrustNewCas gets the ca.global_cas.http_proxy.trust_new_cas node value
func GetCaGlobalCasHttpProxyTrustNewCas(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ca.global_cas.http_proxy.trust_new_cas")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaGlobalCasHttpProxyTrustNewCas updates the ca.global_cas.http_proxy.trust_new_cas node value
func UpdateCaGlobalCasHttpProxyTrustNewCas(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.global_cas.http_proxy.trust_new_cas", bytes.NewReader(byt))
	return
}

// GetCaGlobalCasHttpProxyTrusted gets the ca.global_cas.http_proxy.trusted node value
func GetCaGlobalCasHttpProxyTrusted(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ca.global_cas.http_proxy.trusted")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaGlobalCasHttpProxyTrusted updates the ca.global_cas.http_proxy.trusted node value
func UpdateCaGlobalCasHttpProxyTrusted(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.global_cas.http_proxy.trusted", bytes.NewReader(byt))
	return
}

// GetCaGlobalCasHttpProxyUntrusted gets the ca.global_cas.http_proxy.untrusted node value
func GetCaGlobalCasHttpProxyUntrusted(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ca.global_cas.http_proxy.untrusted")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaGlobalCasHttpProxyUntrusted updates the ca.global_cas.http_proxy.untrusted node value
func UpdateCaGlobalCasHttpProxyUntrusted(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.global_cas.http_proxy.untrusted", bytes.NewReader(byt))
	return
}

// GetCrlsCrls gets the crls.crls node value
func GetCrlsCrls(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/crls.crls")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCrlsCrls updates the crls.crls node value
func UpdateCrlsCrls(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/crls.crls", bytes.NewReader(byt))
	return
}

// GetCssAvPrimaryEngine gets the css.av_primary_engine node value
func GetCssAvPrimaryEngine(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/css.av_primary_engine")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCssAvPrimaryEngine updates the css.av_primary_engine node value
func UpdateCssAvPrimaryEngine(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/css.av_primary_engine", bytes.NewReader(byt))
	return
}

// GetCssSxlLiveprotection gets the css.sxl_liveprotection node value
func GetCssSxlLiveprotection(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/css.sxl_liveprotection")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCssSxlLiveprotection updates the css.sxl_liveprotection node value
func UpdateCssSxlLiveprotection(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/css.sxl_liveprotection", bytes.NewReader(byt))
	return
}

// GetCssSxlSampleSubmit gets the css.sxl_sample_submit node value
func GetCssSxlSampleSubmit(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/css.sxl_sample_submit")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCssSxlSampleSubmit updates the css.sxl_sample_submit node value
func UpdateCssSxlSampleSubmit(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/css.sxl_sample_submit", bytes.NewReader(byt))
	return
}

// GetCustomizationEppLastUpdated gets the customization.epp.last_updated node value
func GetCustomizationEppLastUpdated(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/customization.epp.last_updated")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCustomizationEppLastUpdated updates the customization.epp.last_updated node value
func UpdateCustomizationEppLastUpdated(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/customization.epp.last_updated", bytes.NewReader(byt))
	return
}

// GetCustomizationEppResourcesRoot gets the customization.epp.resources_root node value
func GetCustomizationEppResourcesRoot(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/customization.epp.resources_root")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCustomizationEppResourcesRoot updates the customization.epp.resources_root node value
func UpdateCustomizationEppResourcesRoot(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/customization.epp.resources_root", bytes.NewReader(byt))
	return
}

// GetCustomizationHttpCustomAssets gets the customization.http.custom_assets node value
func GetCustomizationHttpCustomAssets(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/customization.http.custom_assets")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCustomizationHttpCustomAssets updates the customization.http.custom_assets node value
func UpdateCustomizationHttpCustomAssets(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/customization.http.custom_assets", bytes.NewReader(byt))
	return
}

// GetCustomizationHttpCustomTemplates gets the customization.http.custom_templates node value
func GetCustomizationHttpCustomTemplates(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/customization.http.custom_templates")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCustomizationHttpCustomTemplates updates the customization.http.custom_templates node value
func UpdateCustomizationHttpCustomTemplates(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/customization.http.custom_templates", bytes.NewReader(byt))
	return
}

// GetCustomizationHttpLastUpdated gets the customization.http.last_updated node value
func GetCustomizationHttpLastUpdated(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/customization.http.last_updated")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCustomizationHttpLastUpdated updates the customization.http.last_updated node value
func UpdateCustomizationHttpLastUpdated(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/customization.http.last_updated", bytes.NewReader(byt))
	return
}

// GetDebugmodeCrashReport gets the debugmode.crash_report node value
func GetDebugmodeCrashReport(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/debugmode.crash_report")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDebugmodeCrashReport updates the debugmode.crash_report node value
func UpdateDebugmodeCrashReport(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/debugmode.crash_report", bytes.NewReader(byt))
	return
}

// GetDebugmodeEnabled gets the debugmode.enabled node value
func GetDebugmodeEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/debugmode.enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDebugmodeEnabled updates the debugmode.enabled node value
func UpdateDebugmodeEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/debugmode.enabled", bytes.NewReader(byt))
	return
}

// GetDhcpRelayDhcpServer gets the dhcp.relay.dhcp_server node value
func GetDhcpRelayDhcpServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/dhcp.relay.dhcp_server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpRelayDhcpServer updates the dhcp.relay.dhcp_server node value
func UpdateDhcpRelayDhcpServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.relay.dhcp_server", bytes.NewReader(byt))
	return
}

// GetDhcpRelayInterfaces gets the dhcp.relay.interfaces node value
func GetDhcpRelayInterfaces(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/dhcp.relay.interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpRelayInterfaces updates the dhcp.relay.interfaces node value
func UpdateDhcpRelayInterfaces(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.relay.interfaces", bytes.NewReader(byt))
	return
}

// GetDhcpRelayStatus gets the dhcp.relay.status node value
func GetDhcpRelayStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/dhcp.relay.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpRelayStatus updates the dhcp.relay.status node value
func UpdateDhcpRelayStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.relay.status", bytes.NewReader(byt))
	return
}

// GetDhcpRelay6ItfsFacingClients gets the dhcp.relay6.itfs_facing_clients node value
func GetDhcpRelay6ItfsFacingClients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/dhcp.relay6.itfs_facing_clients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpRelay6ItfsFacingClients updates the dhcp.relay6.itfs_facing_clients node value
func UpdateDhcpRelay6ItfsFacingClients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.relay6.itfs_facing_clients", bytes.NewReader(byt))
	return
}

// GetDhcpRelay6ItfsFacingServer6 gets the dhcp.relay6.itfs_facing_server6 node value
func GetDhcpRelay6ItfsFacingServer6(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/dhcp.relay6.itfs_facing_server6")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpRelay6ItfsFacingServer6 updates the dhcp.relay6.itfs_facing_server6 node value
func UpdateDhcpRelay6ItfsFacingServer6(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.relay6.itfs_facing_server6", bytes.NewReader(byt))
	return
}

// GetDhcpRelay6Status gets the dhcp.relay6.status node value
func GetDhcpRelay6Status(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/dhcp.relay6.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpRelay6Status updates the dhcp.relay6.status node value
func UpdateDhcpRelay6Status(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.relay6.status", bytes.NewReader(byt))
	return
}

// GetDhcpServerCustom4 gets the dhcp.server.custom4 node value
func GetDhcpServerCustom4(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/dhcp.server.custom4")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpServerCustom4 updates the dhcp.server.custom4 node value
func UpdateDhcpServerCustom4(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.server.custom4", bytes.NewReader(byt))
	return
}

// GetDhcpServerCustom6 gets the dhcp.server.custom6 node value
func GetDhcpServerCustom6(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/dhcp.server.custom6")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpServerCustom6 updates the dhcp.server.custom6 node value
func UpdateDhcpServerCustom6(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.server.custom6", bytes.NewReader(byt))
	return
}

// GetDhcpServerServers gets the dhcp.server.servers node value
func GetDhcpServerServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/dhcp.server.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpServerServers updates the dhcp.server.servers node value
func UpdateDhcpServerServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.server.servers", bytes.NewReader(byt))
	return
}

// GetDigestAllowedNetworks gets the digest.allowed_networks node value
func GetDigestAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/digest.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestAllowedNetworks updates the digest.allowed_networks node value
func UpdateDigestAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.allowed_networks", bytes.NewReader(byt))
	return
}

// GetDigestCustomText gets the digest.custom_text node value
func GetDigestCustomText(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/digest.custom_text")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestCustomText updates the digest.custom_text node value
func UpdateDigestCustomText(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.custom_text", bytes.NewReader(byt))
	return
}

// GetDigestDomains gets the digest.domains node value
func GetDigestDomains(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/digest.domains")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestDomains updates the digest.domains node value
func UpdateDigestDomains(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.domains", bytes.NewReader(byt))
	return
}

// GetDigestHostname gets the digest.hostname node value
func GetDigestHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/digest.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestHostname updates the digest.hostname node value
func UpdateDigestHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.hostname", bytes.NewReader(byt))
	return
}

// GetDigestMailinglists gets the digest.mailinglists node value
func GetDigestMailinglists(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/digest.mailinglists")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestMailinglists updates the digest.mailinglists node value
func UpdateDigestMailinglists(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.mailinglists", bytes.NewReader(byt))
	return
}

// GetDigestPort gets the digest.port node value
func GetDigestPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/digest.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestPort updates the digest.port node value
func UpdateDigestPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.port", bytes.NewReader(byt))
	return
}

// GetDigestSendTimeOne gets the digest.send_time_one node value
func GetDigestSendTimeOne(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/digest.send_time_one")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestSendTimeOne updates the digest.send_time_one node value
func UpdateDigestSendTimeOne(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.send_time_one", bytes.NewReader(byt))
	return
}

// GetDigestSendTimeTwo gets the digest.send_time_two node value
func GetDigestSendTimeTwo(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/digest.send_time_two")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestSendTimeTwo updates the digest.send_time_two node value
func UpdateDigestSendTimeTwo(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.send_time_two", bytes.NewReader(byt))
	return
}

// GetDigestSkiplist gets the digest.skiplist node value
func GetDigestSkiplist(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/digest.skiplist")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestSkiplist updates the digest.skiplist node value
func UpdateDigestSkiplist(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.skiplist", bytes.NewReader(byt))
	return
}

// GetDigestStatus gets the digest.status node value
func GetDigestStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/digest.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestStatus updates the digest.status node value
func UpdateDigestStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.status", bytes.NewReader(byt))
	return
}

// GetDigestUserRelease gets the digest.user_release node value
func GetDigestUserRelease(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/digest.user_release")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestUserRelease updates the digest.user_release node value
func UpdateDigestUserRelease(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.user_release", bytes.NewReader(byt))
	return
}

// GetDnsAllowedNetworks gets the dns.allowed_networks node value
func GetDnsAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/dns.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsAllowedNetworks updates the dns.allowed_networks node value
func UpdateDnsAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.allowed_networks", bytes.NewReader(byt))
	return
}

// GetDnsAxfr gets the dns.axfr node value
func GetDnsAxfr(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/dns.axfr")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsAxfr updates the dns.axfr node value
func UpdateDnsAxfr(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.axfr", bytes.NewReader(byt))
	return
}

// GetDnsDnssec gets the dns.dnssec node value
func GetDnsDnssec(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/dns.dnssec")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsDnssec updates the dns.dnssec node value
func UpdateDnsDnssec(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.dnssec", bytes.NewReader(byt))
	return
}

// GetDnsEmail gets the dns.email node value
func GetDnsEmail(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/dns.email")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsEmail updates the dns.email node value
func UpdateDnsEmail(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.email", bytes.NewReader(byt))
	return
}

// GetDnsEmptyZones gets the dns.empty_zones node value
func GetDnsEmptyZones(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/dns.empty_zones")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsEmptyZones updates the dns.empty_zones node value
func UpdateDnsEmptyZones(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.empty_zones", bytes.NewReader(byt))
	return
}

// GetDnsFwdDynamic gets the dns.fwd_dynamic node value
func GetDnsFwdDynamic(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/dns.fwd_dynamic")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsFwdDynamic updates the dns.fwd_dynamic node value
func UpdateDnsFwdDynamic(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.fwd_dynamic", bytes.NewReader(byt))
	return
}

// GetDnsFwdStatic gets the dns.fwd_static node value
func GetDnsFwdStatic(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/dns.fwd_static")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsFwdStatic updates the dns.fwd_static node value
func UpdateDnsFwdStatic(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.fwd_static", bytes.NewReader(byt))
	return
}

// GetDnsRecheckInterval gets the dns.recheck_interval node value
func GetDnsRecheckInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/dns.recheck_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsRecheckInterval updates the dns.recheck_interval node value
func UpdateDnsRecheckInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.recheck_interval", bytes.NewReader(byt))
	return
}

// GetDnsRoutes gets the dns.routes node value
func GetDnsRoutes(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/dns.routes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsRoutes updates the dns.routes node value
func UpdateDnsRoutes(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.routes", bytes.NewReader(byt))
	return
}

// GetDyndnsRules gets the dyndns.rules node value
func GetDyndnsRules(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/dyndns.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDyndnsRules updates the dyndns.rules node value
func UpdateDyndnsRules(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dyndns.rules", bytes.NewReader(byt))
	return
}

// GetEmailpkiAuthorityCert gets the emailpki.authority.cert node value
func GetEmailpkiAuthorityCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.authority.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiAuthorityCert updates the emailpki.authority.cert node value
func UpdateEmailpkiAuthorityCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.authority.cert", bytes.NewReader(byt))
	return
}

// GetEmailpkiAuthorityFingerprint gets the emailpki.authority.fingerprint node value
func GetEmailpkiAuthorityFingerprint(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.authority.fingerprint")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiAuthorityFingerprint updates the emailpki.authority.fingerprint node value
func UpdateEmailpkiAuthorityFingerprint(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.authority.fingerprint", bytes.NewReader(byt))
	return
}

// GetEmailpkiAuthorityKey gets the emailpki.authority.key node value
func GetEmailpkiAuthorityKey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.authority.key")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiAuthorityKey updates the emailpki.authority.key node value
func UpdateEmailpkiAuthorityKey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.authority.key", bytes.NewReader(byt))
	return
}

// GetEmailpkiAuthorityPostmasterFingerprint gets the emailpki.authority.postmaster_fingerprint node value
func GetEmailpkiAuthorityPostmasterFingerprint(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.authority.postmaster_fingerprint")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiAuthorityPostmasterFingerprint updates the emailpki.authority.postmaster_fingerprint node value
func UpdateEmailpkiAuthorityPostmasterFingerprint(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.authority.postmaster_fingerprint", bytes.NewReader(byt))
	return
}

// GetEmailpkiAuthorityPostmasterPrivkey gets the emailpki.authority.postmaster_privkey node value
func GetEmailpkiAuthorityPostmasterPrivkey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.authority.postmaster_privkey")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiAuthorityPostmasterPrivkey updates the emailpki.authority.postmaster_privkey node value
func UpdateEmailpkiAuthorityPostmasterPrivkey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.authority.postmaster_privkey", bytes.NewReader(byt))
	return
}

// GetEmailpkiAuthorityPostmasterPubkey gets the emailpki.authority.postmaster_pubkey node value
func GetEmailpkiAuthorityPostmasterPubkey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.authority.postmaster_pubkey")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiAuthorityPostmasterPubkey updates the emailpki.authority.postmaster_pubkey node value
func UpdateEmailpkiAuthorityPostmasterPubkey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.authority.postmaster_pubkey", bytes.NewReader(byt))
	return
}

// GetEmailpkiGlobalCity gets the emailpki.global.city node value
func GetEmailpkiGlobalCity(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.global.city")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiGlobalCity updates the emailpki.global.city node value
func UpdateEmailpkiGlobalCity(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.global.city", bytes.NewReader(byt))
	return
}

// GetEmailpkiGlobalCountry gets the emailpki.global.country node value
func GetEmailpkiGlobalCountry(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.global.country")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiGlobalCountry updates the emailpki.global.country node value
func UpdateEmailpkiGlobalCountry(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.global.country", bytes.NewReader(byt))
	return
}

// GetEmailpkiGlobalOrganization gets the emailpki.global.organization node value
func GetEmailpkiGlobalOrganization(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.global.organization")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiGlobalOrganization updates the emailpki.global.organization node value
func UpdateEmailpkiGlobalOrganization(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.global.organization", bytes.NewReader(byt))
	return
}

// GetEmailpkiGlobalPostmaster gets the emailpki.global.postmaster node value
func GetEmailpkiGlobalPostmaster(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.global.postmaster")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiGlobalPostmaster updates the emailpki.global.postmaster node value
func UpdateEmailpkiGlobalPostmaster(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.global.postmaster", bytes.NewReader(byt))
	return
}

// GetEmailpkiGlobalStatus gets the emailpki.global.status node value
func GetEmailpkiGlobalStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/emailpki.global.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiGlobalStatus updates the emailpki.global.status node value
func UpdateEmailpkiGlobalStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.global.status", bytes.NewReader(byt))
	return
}

// GetEmailpkiObjectsCas gets the emailpki.objects.cas node value
func GetEmailpkiObjectsCas(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/emailpki.objects.cas")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiObjectsCas updates the emailpki.objects.cas node value
func UpdateEmailpkiObjectsCas(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.objects.cas", bytes.NewReader(byt))
	return
}

// GetEmailpkiObjectsOpenpgp gets the emailpki.objects.openpgp node value
func GetEmailpkiObjectsOpenpgp(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/emailpki.objects.openpgp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiObjectsOpenpgp updates the emailpki.objects.openpgp node value
func UpdateEmailpkiObjectsOpenpgp(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.objects.openpgp", bytes.NewReader(byt))
	return
}

// GetEmailpkiObjectsSmime gets the emailpki.objects.smime node value
func GetEmailpkiObjectsSmime(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/emailpki.objects.smime")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiObjectsSmime updates the emailpki.objects.smime node value
func UpdateEmailpkiObjectsSmime(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.objects.smime", bytes.NewReader(byt))
	return
}

// GetEmailpkiObjectsUsers gets the emailpki.objects.users node value
func GetEmailpkiObjectsUsers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/emailpki.objects.users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiObjectsUsers updates the emailpki.objects.users node value
func UpdateEmailpkiObjectsUsers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.objects.users", bytes.NewReader(byt))
	return
}

// GetEmailpkiOpenpgpMainKeysize gets the emailpki.openpgp.main_keysize node value
func GetEmailpkiOpenpgpMainKeysize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/emailpki.openpgp.main_keysize")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOpenpgpMainKeysize updates the emailpki.openpgp.main_keysize node value
func UpdateEmailpkiOpenpgpMainKeysize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.openpgp.main_keysize", bytes.NewReader(byt))
	return
}

// GetEmailpkiOpenpgpSubKeysize gets the emailpki.openpgp.sub_keysize node value
func GetEmailpkiOpenpgpSubKeysize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/emailpki.openpgp.sub_keysize")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOpenpgpSubKeysize updates the emailpki.openpgp.sub_keysize node value
func UpdateEmailpkiOpenpgpSubKeysize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.openpgp.sub_keysize", bytes.NewReader(byt))
	return
}

// GetEmailpkiOptionsExternalAuto gets the emailpki.options.external_auto node value
func GetEmailpkiOptionsExternalAuto(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/emailpki.options.external_auto")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOptionsExternalAuto updates the emailpki.options.external_auto node value
func UpdateEmailpkiOptionsExternalAuto(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.options.external_auto", bytes.NewReader(byt))
	return
}

// GetEmailpkiOptionsKeyserver gets the emailpki.options.keyserver node value
func GetEmailpkiOptionsKeyserver(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.options.keyserver")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOptionsKeyserver updates the emailpki.options.keyserver node value
func UpdateEmailpkiOptionsKeyserver(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.options.keyserver", bytes.NewReader(byt))
	return
}

// GetEmailpkiOptionsPolicyDecryption gets the emailpki.options.policy_decryption node value
func GetEmailpkiOptionsPolicyDecryption(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/emailpki.options.policy_decryption")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOptionsPolicyDecryption updates the emailpki.options.policy_decryption node value
func UpdateEmailpkiOptionsPolicyDecryption(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.options.policy_decryption", bytes.NewReader(byt))
	return
}

// GetEmailpkiOptionsPolicyEncryption gets the emailpki.options.policy_encryption node value
func GetEmailpkiOptionsPolicyEncryption(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/emailpki.options.policy_encryption")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOptionsPolicyEncryption updates the emailpki.options.policy_encryption node value
func UpdateEmailpkiOptionsPolicyEncryption(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.options.policy_encryption", bytes.NewReader(byt))
	return
}

// GetEmailpkiOptionsPolicySign gets the emailpki.options.policy_sign node value
func GetEmailpkiOptionsPolicySign(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/emailpki.options.policy_sign")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOptionsPolicySign updates the emailpki.options.policy_sign node value
func UpdateEmailpkiOptionsPolicySign(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.options.policy_sign", bytes.NewReader(byt))
	return
}

// GetEmailpkiOptionsPolicyVerify gets the emailpki.options.policy_verify node value
func GetEmailpkiOptionsPolicyVerify(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/emailpki.options.policy_verify")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOptionsPolicyVerify updates the emailpki.options.policy_verify node value
func UpdateEmailpkiOptionsPolicyVerify(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.options.policy_verify", bytes.NewReader(byt))
	return
}

// GetEndpointAacAllowedNetworks gets the endpoint.aac.allowed_networks node value
func GetEndpointAacAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacAllowedNetworks updates the endpoint.aac.allowed_networks node value
func UpdateEndpointAacAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.allowed_networks", bytes.NewReader(byt))
	return
}

// GetEndpointAacAllowedUsers gets the endpoint.aac.allowed_users node value
func GetEndpointAacAllowedUsers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.allowed_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacAllowedUsers updates the endpoint.aac.allowed_users node value
func UpdateEndpointAacAllowedUsers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.allowed_users", bytes.NewReader(byt))
	return
}

// GetEndpointAacCa gets the endpoint.aac.ca node value
func GetEndpointAacCa(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.ca")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacCa updates the endpoint.aac.ca node value
func UpdateEndpointAacCa(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.ca", bytes.NewReader(byt))
	return
}

// GetEndpointAacCert gets the endpoint.aac.cert node value
func GetEndpointAacCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacCert updates the endpoint.aac.cert node value
func UpdateEndpointAacCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.cert", bytes.NewReader(byt))
	return
}

// GetEndpointAacMagicIp gets the endpoint.aac.magic_ip node value
func GetEndpointAacMagicIp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.magic_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacMagicIp updates the endpoint.aac.magic_ip node value
func UpdateEndpointAacMagicIp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.magic_ip", bytes.NewReader(byt))
	return
}

// GetEndpointAacMaxUserLogins gets the endpoint.aac.max_user_logins node value
func GetEndpointAacMaxUserLogins(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.max_user_logins")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacMaxUserLogins updates the endpoint.aac.max_user_logins node value
func UpdateEndpointAacMaxUserLogins(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.max_user_logins", bytes.NewReader(byt))
	return
}

// GetEndpointAacStatus gets the endpoint.aac.status node value
func GetEndpointAacStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacStatus updates the endpoint.aac.status node value
func UpdateEndpointAacStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.status", bytes.NewReader(byt))
	return
}

// GetEndpointStasCollectors gets the endpoint.stas.collectors node value
func GetEndpointStasCollectors(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/endpoint.stas.collectors")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointStasCollectors updates the endpoint.stas.collectors node value
func UpdateEndpointStasCollectors(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.stas.collectors", bytes.NewReader(byt))
	return
}

// GetEndpointStasStatus gets the endpoint.stas.status node value
func GetEndpointStasStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/endpoint.stas.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointStasStatus updates the endpoint.stas.status node value
func UpdateEndpointStasStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.stas.status", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesCompanyLogo gets the enduser_messages.company_logo node value
func GetEnduserMessagesCompanyLogo(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.company_logo")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesCompanyLogo updates the enduser_messages.company_logo node value
func UpdateEnduserMessagesCompanyLogo(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.company_logo", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesCompanyText gets the enduser_messages.company_text node value
func GetEnduserMessagesCompanyText(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.company_text")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesCompanyText updates the enduser_messages.company_text node value
func UpdateEnduserMessagesCompanyText(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.company_text", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesDlpBlackholePart gets the enduser_messages.dlp.blackhole_part node value
func GetEnduserMessagesDlpBlackholePart(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.dlp.blackhole_part")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesDlpBlackholePart updates the enduser_messages.dlp.blackhole_part node value
func UpdateEnduserMessagesDlpBlackholePart(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.dlp.blackhole_part", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesDlpFooterPart gets the enduser_messages.dlp.footer_part node value
func GetEnduserMessagesDlpFooterPart(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.dlp.footer_part")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesDlpFooterPart updates the enduser_messages.dlp.footer_part node value
func UpdateEnduserMessagesDlpFooterPart(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.dlp.footer_part", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesDlpHeaderPart gets the enduser_messages.dlp.header_part node value
func GetEnduserMessagesDlpHeaderPart(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.dlp.header_part")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesDlpHeaderPart updates the enduser_messages.dlp.header_part node value
func UpdateEnduserMessagesDlpHeaderPart(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.dlp.header_part", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesDlpOriginalPart gets the enduser_messages.dlp.original_part node value
func GetEnduserMessagesDlpOriginalPart(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.dlp.original_part")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesDlpOriginalPart updates the enduser_messages.dlp.original_part node value
func UpdateEnduserMessagesDlpOriginalPart(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.dlp.original_part", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesDlpSpxPart gets the enduser_messages.dlp.spx_part node value
func GetEnduserMessagesDlpSpxPart(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.dlp.spx_part")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesDlpSpxPart updates the enduser_messages.dlp.spx_part node value
func UpdateEnduserMessagesDlpSpxPart(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.dlp.spx_part", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesDlpSubject gets the enduser_messages.dlp.subject node value
func GetEnduserMessagesDlpSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.dlp.subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesDlpSubject updates the enduser_messages.dlp.subject node value
func UpdateEnduserMessagesDlpSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.dlp.subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpAppDesc gets the enduser_messages.http.app_desc node value
func GetEnduserMessagesHttpAppDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.app_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpAppDesc updates the enduser_messages.http.app_desc node value
func UpdateEnduserMessagesHttpAppDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.app_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpAppSubject gets the enduser_messages.http.app_subject node value
func GetEnduserMessagesHttpAppSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.app_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpAppSubject updates the enduser_messages.http.app_subject node value
func UpdateEnduserMessagesHttpAppSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.app_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpBlacklistDesc gets the enduser_messages.http.blacklist_desc node value
func GetEnduserMessagesHttpBlacklistDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.blacklist_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpBlacklistDesc updates the enduser_messages.http.blacklist_desc node value
func UpdateEnduserMessagesHttpBlacklistDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.blacklist_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpBlacklistSubject gets the enduser_messages.http.blacklist_subject node value
func GetEnduserMessagesHttpBlacklistSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.blacklist_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpBlacklistSubject updates the enduser_messages.http.blacklist_subject node value
func UpdateEnduserMessagesHttpBlacklistSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.blacklist_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpCertfailSubject gets the enduser_messages.http.certfail_subject node value
func GetEnduserMessagesHttpCertfailSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.certfail_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpCertfailSubject updates the enduser_messages.http.certfail_subject node value
func UpdateEnduserMessagesHttpCertfailSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.certfail_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpCffOverrideDesc gets the enduser_messages.http.cff_override_desc node value
func GetEnduserMessagesHttpCffOverrideDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.cff_override_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpCffOverrideDesc updates the enduser_messages.http.cff_override_desc node value
func UpdateEnduserMessagesHttpCffOverrideDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.cff_override_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpCffOverrideSubject gets the enduser_messages.http.cff_override_subject node value
func GetEnduserMessagesHttpCffOverrideSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.cff_override_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpCffOverrideSubject updates the enduser_messages.http.cff_override_subject node value
func UpdateEnduserMessagesHttpCffOverrideSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.cff_override_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpCffOverrideTerms gets the enduser_messages.http.cff_override_terms node value
func GetEnduserMessagesHttpCffOverrideTerms(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.cff_override_terms")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpCffOverrideTerms updates the enduser_messages.http.cff_override_terms node value
func UpdateEnduserMessagesHttpCffOverrideTerms(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.cff_override_terms", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpDownloadCompleteDesc gets the enduser_messages.http.download_complete_desc node value
func GetEnduserMessagesHttpDownloadCompleteDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.download_complete_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpDownloadCompleteDesc updates the enduser_messages.http.download_complete_desc node value
func UpdateEnduserMessagesHttpDownloadCompleteDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.download_complete_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpDownloadCompleteSubject gets the enduser_messages.http.download_complete_subject node value
func GetEnduserMessagesHttpDownloadCompleteSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.download_complete_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpDownloadCompleteSubject updates the enduser_messages.http.download_complete_subject node value
func UpdateEnduserMessagesHttpDownloadCompleteSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.download_complete_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpDownloadDesc gets the enduser_messages.http.download_desc node value
func GetEnduserMessagesHttpDownloadDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.download_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpDownloadDesc updates the enduser_messages.http.download_desc node value
func UpdateEnduserMessagesHttpDownloadDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.download_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpDownloadSubject gets the enduser_messages.http.download_subject node value
func GetEnduserMessagesHttpDownloadSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.download_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpDownloadSubject updates the enduser_messages.http.download_subject node value
func UpdateEnduserMessagesHttpDownloadSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.download_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpErrorDesc gets the enduser_messages.http.error_desc node value
func GetEnduserMessagesHttpErrorDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.error_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpErrorDesc updates the enduser_messages.http.error_desc node value
func UpdateEnduserMessagesHttpErrorDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.error_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpErrorSubject gets the enduser_messages.http.error_subject node value
func GetEnduserMessagesHttpErrorSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.error_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpErrorSubject updates the enduser_messages.http.error_subject node value
func UpdateEnduserMessagesHttpErrorSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.error_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpFileextensionDesc gets the enduser_messages.http.fileextension_desc node value
func GetEnduserMessagesHttpFileextensionDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.fileextension_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpFileextensionDesc updates the enduser_messages.http.fileextension_desc node value
func UpdateEnduserMessagesHttpFileextensionDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.fileextension_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpFileextensionSubject gets the enduser_messages.http.fileextension_subject node value
func GetEnduserMessagesHttpFileextensionSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.fileextension_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpFileextensionSubject updates the enduser_messages.http.fileextension_subject node value
func UpdateEnduserMessagesHttpFileextensionSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.fileextension_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpFileextensionWarnDesc gets the enduser_messages.http.fileextension_warn_desc node value
func GetEnduserMessagesHttpFileextensionWarnDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.fileextension_warn_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpFileextensionWarnDesc updates the enduser_messages.http.fileextension_warn_desc node value
func UpdateEnduserMessagesHttpFileextensionWarnDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.fileextension_warn_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpFileextensionWarnSubject gets the enduser_messages.http.fileextension_warn_subject node value
func GetEnduserMessagesHttpFileextensionWarnSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.fileextension_warn_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpFileextensionWarnSubject updates the enduser_messages.http.fileextension_warn_subject node value
func UpdateEnduserMessagesHttpFileextensionWarnSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.fileextension_warn_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpFilesizeDesc gets the enduser_messages.http.filesize_desc node value
func GetEnduserMessagesHttpFilesizeDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.filesize_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpFilesizeDesc updates the enduser_messages.http.filesize_desc node value
func UpdateEnduserMessagesHttpFilesizeDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.filesize_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpFilesizeSubject gets the enduser_messages.http.filesize_subject node value
func GetEnduserMessagesHttpFilesizeSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.filesize_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpFilesizeSubject updates the enduser_messages.http.filesize_subject node value
func UpdateEnduserMessagesHttpFilesizeSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.filesize_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpGeoipDesc gets the enduser_messages.http.geoip_desc node value
func GetEnduserMessagesHttpGeoipDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.geoip_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpGeoipDesc updates the enduser_messages.http.geoip_desc node value
func UpdateEnduserMessagesHttpGeoipDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.geoip_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpGeoipSubject gets the enduser_messages.http.geoip_subject node value
func GetEnduserMessagesHttpGeoipSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.geoip_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpGeoipSubject updates the enduser_messages.http.geoip_subject node value
func UpdateEnduserMessagesHttpGeoipSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.geoip_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpMimetypeDesc gets the enduser_messages.http.mimetype_desc node value
func GetEnduserMessagesHttpMimetypeDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.mimetype_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpMimetypeDesc updates the enduser_messages.http.mimetype_desc node value
func UpdateEnduserMessagesHttpMimetypeDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.mimetype_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpMimetypeSubject gets the enduser_messages.http.mimetype_subject node value
func GetEnduserMessagesHttpMimetypeSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.mimetype_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpMimetypeSubject updates the enduser_messages.http.mimetype_subject node value
func UpdateEnduserMessagesHttpMimetypeSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.mimetype_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpMimetypeWarnDesc gets the enduser_messages.http.mimetype_warn_desc node value
func GetEnduserMessagesHttpMimetypeWarnDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.mimetype_warn_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpMimetypeWarnDesc updates the enduser_messages.http.mimetype_warn_desc node value
func UpdateEnduserMessagesHttpMimetypeWarnDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.mimetype_warn_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpMimetypeWarnSubject gets the enduser_messages.http.mimetype_warn_subject node value
func GetEnduserMessagesHttpMimetypeWarnSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.mimetype_warn_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpMimetypeWarnSubject updates the enduser_messages.http.mimetype_warn_subject node value
func UpdateEnduserMessagesHttpMimetypeWarnSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.mimetype_warn_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpPuaDesc gets the enduser_messages.http.pua_desc node value
func GetEnduserMessagesHttpPuaDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.pua_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpPuaDesc updates the enduser_messages.http.pua_desc node value
func UpdateEnduserMessagesHttpPuaDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.pua_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpPuaSubject gets the enduser_messages.http.pua_subject node value
func GetEnduserMessagesHttpPuaSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.pua_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpPuaSubject updates the enduser_messages.http.pua_subject node value
func UpdateEnduserMessagesHttpPuaSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.pua_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpQuotaBlockDesc gets the enduser_messages.http.quota_block_desc node value
func GetEnduserMessagesHttpQuotaBlockDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.quota_block_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpQuotaBlockDesc updates the enduser_messages.http.quota_block_desc node value
func UpdateEnduserMessagesHttpQuotaBlockDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.quota_block_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpQuotaBlockSubject gets the enduser_messages.http.quota_block_subject node value
func GetEnduserMessagesHttpQuotaBlockSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.quota_block_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpQuotaBlockSubject updates the enduser_messages.http.quota_block_subject node value
func UpdateEnduserMessagesHttpQuotaBlockSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.quota_block_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpQuotaWarnDesc gets the enduser_messages.http.quota_warn_desc node value
func GetEnduserMessagesHttpQuotaWarnDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.quota_warn_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpQuotaWarnDesc updates the enduser_messages.http.quota_warn_desc node value
func UpdateEnduserMessagesHttpQuotaWarnDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.quota_warn_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpQuotaWarnSubject gets the enduser_messages.http.quota_warn_subject node value
func GetEnduserMessagesHttpQuotaWarnSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.quota_warn_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpQuotaWarnSubject updates the enduser_messages.http.quota_warn_subject node value
func UpdateEnduserMessagesHttpQuotaWarnSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.quota_warn_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSpDesc gets the enduser_messages.http.sp_desc node value
func GetEnduserMessagesHttpSpDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.sp_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSpDesc updates the enduser_messages.http.sp_desc node value
func UpdateEnduserMessagesHttpSpDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.sp_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSpFrameSubject gets the enduser_messages.http.sp_frame_subject node value
func GetEnduserMessagesHttpSpFrameSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.sp_frame_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSpFrameSubject updates the enduser_messages.http.sp_frame_subject node value
func UpdateEnduserMessagesHttpSpFrameSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.sp_frame_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSpSubject gets the enduser_messages.http.sp_subject node value
func GetEnduserMessagesHttpSpSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.sp_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSpSubject updates the enduser_messages.http.sp_subject node value
func UpdateEnduserMessagesHttpSpSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.sp_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSpWarnDesc gets the enduser_messages.http.sp_warn_desc node value
func GetEnduserMessagesHttpSpWarnDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.sp_warn_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSpWarnDesc updates the enduser_messages.http.sp_warn_desc node value
func UpdateEnduserMessagesHttpSpWarnDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.sp_warn_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSpWarnSubject gets the enduser_messages.http.sp_warn_subject node value
func GetEnduserMessagesHttpSpWarnSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.sp_warn_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSpWarnSubject updates the enduser_messages.http.sp_warn_subject node value
func UpdateEnduserMessagesHttpSpWarnSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.sp_warn_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslCertraw gets the enduser_messages.http.ssl_certraw node value
func GetEnduserMessagesHttpSslCertraw(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_certraw")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslCertraw updates the enduser_messages.http.ssl_certraw node value
func UpdateEnduserMessagesHttpSslCertraw(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_certraw", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslCertstatus gets the enduser_messages.http.ssl_certstatus node value
func GetEnduserMessagesHttpSslCertstatus(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_certstatus")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslCertstatus updates the enduser_messages.http.ssl_certstatus node value
func UpdateEnduserMessagesHttpSslCertstatus(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_certstatus", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslIssuer gets the enduser_messages.http.ssl_issuer node value
func GetEnduserMessagesHttpSslIssuer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_issuer")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslIssuer updates the enduser_messages.http.ssl_issuer node value
func UpdateEnduserMessagesHttpSslIssuer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_issuer", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslMd5Fp gets the enduser_messages.http.ssl_md5fp node value
func GetEnduserMessagesHttpSslMd5Fp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_md5fp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslMd5Fp updates the enduser_messages.http.ssl_md5fp node value
func UpdateEnduserMessagesHttpSslMd5Fp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_md5fp", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslSha1Fp gets the enduser_messages.http.ssl_sha1fp node value
func GetEnduserMessagesHttpSslSha1Fp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_sha1fp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslSha1Fp updates the enduser_messages.http.ssl_sha1fp node value
func UpdateEnduserMessagesHttpSslSha1Fp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_sha1fp", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslSubject gets the enduser_messages.http.ssl_subject node value
func GetEnduserMessagesHttpSslSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslSubject updates the enduser_messages.http.ssl_subject node value
func UpdateEnduserMessagesHttpSslSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslValidfrom gets the enduser_messages.http.ssl_validfrom node value
func GetEnduserMessagesHttpSslValidfrom(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_validfrom")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslValidfrom updates the enduser_messages.http.ssl_validfrom node value
func UpdateEnduserMessagesHttpSslValidfrom(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_validfrom", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslValiduntil gets the enduser_messages.http.ssl_validuntil node value
func GetEnduserMessagesHttpSslValiduntil(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_validuntil")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslValiduntil updates the enduser_messages.http.ssl_validuntil node value
func UpdateEnduserMessagesHttpSslValiduntil(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_validuntil", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpThreatDesc gets the enduser_messages.http.threat_desc node value
func GetEnduserMessagesHttpThreatDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.threat_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpThreatDesc updates the enduser_messages.http.threat_desc node value
func UpdateEnduserMessagesHttpThreatDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.threat_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpThreatSubject gets the enduser_messages.http.threat_subject node value
func GetEnduserMessagesHttpThreatSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.threat_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpThreatSubject updates the enduser_messages.http.threat_subject node value
func UpdateEnduserMessagesHttpThreatSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.threat_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpTransparentAuthDesc gets the enduser_messages.http.transparent_auth_desc node value
func GetEnduserMessagesHttpTransparentAuthDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.transparent_auth_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpTransparentAuthDesc updates the enduser_messages.http.transparent_auth_desc node value
func UpdateEnduserMessagesHttpTransparentAuthDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.transparent_auth_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpTransparentAuthSubject gets the enduser_messages.http.transparent_auth_subject node value
func GetEnduserMessagesHttpTransparentAuthSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.transparent_auth_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpTransparentAuthSubject updates the enduser_messages.http.transparent_auth_subject node value
func UpdateEnduserMessagesHttpTransparentAuthSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.transparent_auth_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpTransparentAuthTerms gets the enduser_messages.http.transparent_auth_terms node value
func GetEnduserMessagesHttpTransparentAuthTerms(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.transparent_auth_terms")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpTransparentAuthTerms updates the enduser_messages.http.transparent_auth_terms node value
func UpdateEnduserMessagesHttpTransparentAuthTerms(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.transparent_auth_terms", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpVirusDesc gets the enduser_messages.http.virus_desc node value
func GetEnduserMessagesHttpVirusDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.virus_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpVirusDesc updates the enduser_messages.http.virus_desc node value
func UpdateEnduserMessagesHttpVirusDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.virus_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpVirusSubject gets the enduser_messages.http.virus_subject node value
func GetEnduserMessagesHttpVirusSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.virus_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpVirusSubject updates the enduser_messages.http.virus_subject node value
func UpdateEnduserMessagesHttpVirusSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.virus_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpVirusscanDesc gets the enduser_messages.http.virusscan_desc node value
func GetEnduserMessagesHttpVirusscanDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.virusscan_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpVirusscanDesc updates the enduser_messages.http.virusscan_desc node value
func UpdateEnduserMessagesHttpVirusscanDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.virusscan_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpVirusscanSubject gets the enduser_messages.http.virusscan_subject node value
func GetEnduserMessagesHttpVirusscanSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.virusscan_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpVirusscanSubject updates the enduser_messages.http.virusscan_subject node value
func UpdateEnduserMessagesHttpVirusscanSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.virusscan_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesMailReleaseErrDesc gets the enduser_messages.mail.release_err_desc node value
func GetEnduserMessagesMailReleaseErrDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.mail.release_err_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesMailReleaseErrDesc updates the enduser_messages.mail.release_err_desc node value
func UpdateEnduserMessagesMailReleaseErrDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.mail.release_err_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesMailReleaseErrSubject gets the enduser_messages.mail.release_err_subject node value
func GetEnduserMessagesMailReleaseErrSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.mail.release_err_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesMailReleaseErrSubject updates the enduser_messages.mail.release_err_subject node value
func UpdateEnduserMessagesMailReleaseErrSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.mail.release_err_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesMailReleasedDesc gets the enduser_messages.mail.released_desc node value
func GetEnduserMessagesMailReleasedDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.mail.released_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesMailReleasedDesc updates the enduser_messages.mail.released_desc node value
func UpdateEnduserMessagesMailReleasedDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.mail.released_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesMailReleasedSubject gets the enduser_messages.mail.released_subject node value
func GetEnduserMessagesMailReleasedSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.mail.released_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesMailReleasedSubject updates the enduser_messages.mail.released_subject node value
func UpdateEnduserMessagesMailReleasedSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.mail.released_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesPop3BlockedDesc gets the enduser_messages.pop3.blocked_desc node value
func GetEnduserMessagesPop3BlockedDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.pop3.blocked_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesPop3BlockedDesc updates the enduser_messages.pop3.blocked_desc node value
func UpdateEnduserMessagesPop3BlockedDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.pop3.blocked_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesPop3BlockedSubject gets the enduser_messages.pop3.blocked_subject node value
func GetEnduserMessagesPop3BlockedSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.pop3.blocked_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesPop3BlockedSubject updates the enduser_messages.pop3.blocked_subject node value
func UpdateEnduserMessagesPop3BlockedSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.pop3.blocked_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxInternalErrorBody gets the enduser_messages.spx.internal_error.body node value
func GetEnduserMessagesSpxInternalErrorBody(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.internal_error.body")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxInternalErrorBody updates the enduser_messages.spx.internal_error.body node value
func UpdateEnduserMessagesSpxInternalErrorBody(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.internal_error.body", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxInternalErrorSubject gets the enduser_messages.spx.internal_error.subject node value
func GetEnduserMessagesSpxInternalErrorSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.internal_error.subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxInternalErrorSubject updates the enduser_messages.spx.internal_error.subject node value
func UpdateEnduserMessagesSpxInternalErrorSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.internal_error.subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxInternalErrorSenderBody gets the enduser_messages.spx.internal_error_sender.body node value
func GetEnduserMessagesSpxInternalErrorSenderBody(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.internal_error_sender.body")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxInternalErrorSenderBody updates the enduser_messages.spx.internal_error_sender.body node value
func UpdateEnduserMessagesSpxInternalErrorSenderBody(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.internal_error_sender.body", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxInternalErrorSenderSubject gets the enduser_messages.spx.internal_error_sender.subject node value
func GetEnduserMessagesSpxInternalErrorSenderSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.internal_error_sender.subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxInternalErrorSenderSubject updates the enduser_messages.spx.internal_error_sender.subject node value
func UpdateEnduserMessagesSpxInternalErrorSenderSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.internal_error_sender.subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxPasswordNoSpecCharsBody gets the enduser_messages.spx.password_no_spec_chars.body node value
func GetEnduserMessagesSpxPasswordNoSpecCharsBody(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.password_no_spec_chars.body")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxPasswordNoSpecCharsBody updates the enduser_messages.spx.password_no_spec_chars.body node value
func UpdateEnduserMessagesSpxPasswordNoSpecCharsBody(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.password_no_spec_chars.body", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxPasswordNoSpecCharsSubject gets the enduser_messages.spx.password_no_spec_chars.subject node value
func GetEnduserMessagesSpxPasswordNoSpecCharsSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.password_no_spec_chars.subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxPasswordNoSpecCharsSubject updates the enduser_messages.spx.password_no_spec_chars.subject node value
func UpdateEnduserMessagesSpxPasswordNoSpecCharsSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.password_no_spec_chars.subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxPasswordNotLongEnoughBody gets the enduser_messages.spx.password_not_long_enough.body node value
func GetEnduserMessagesSpxPasswordNotLongEnoughBody(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.password_not_long_enough.body")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxPasswordNotLongEnoughBody updates the enduser_messages.spx.password_not_long_enough.body node value
func UpdateEnduserMessagesSpxPasswordNotLongEnoughBody(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.password_not_long_enough.body", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxPasswordNotLongEnoughSubject gets the enduser_messages.spx.password_not_long_enough.subject node value
func GetEnduserMessagesSpxPasswordNotLongEnoughSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.password_not_long_enough.subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxPasswordNotLongEnoughSubject updates the enduser_messages.spx.password_not_long_enough.subject node value
func UpdateEnduserMessagesSpxPasswordNotLongEnoughSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.password_not_long_enough.subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxPasswordNotPresentedBody gets the enduser_messages.spx.password_not_presented.body node value
func GetEnduserMessagesSpxPasswordNotPresentedBody(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.password_not_presented.body")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxPasswordNotPresentedBody updates the enduser_messages.spx.password_not_presented.body node value
func UpdateEnduserMessagesSpxPasswordNotPresentedBody(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.password_not_presented.body", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxPasswordNotPresentedSubject gets the enduser_messages.spx.password_not_presented.subject node value
func GetEnduserMessagesSpxPasswordNotPresentedSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.password_not_presented.subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxPasswordNotPresentedSubject updates the enduser_messages.spx.password_not_presented.subject node value
func UpdateEnduserMessagesSpxPasswordNotPresentedSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.password_not_presented.subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxUrlNotFoundMessage gets the enduser_messages.spx.url_not_found.message node value
func GetEnduserMessagesSpxUrlNotFoundMessage(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.url_not_found.message")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxUrlNotFoundMessage updates the enduser_messages.spx.url_not_found.message node value
func UpdateEnduserMessagesSpxUrlNotFoundMessage(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.url_not_found.message", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSquidCacheAdmin gets the enduser_messages.squid.cache_admin node value
func GetEnduserMessagesSquidCacheAdmin(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.squid.cache_admin")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSquidCacheAdmin updates the enduser_messages.squid.cache_admin node value
func UpdateEnduserMessagesSquidCacheAdmin(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.squid.cache_admin", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSquidCacheAdminMessage gets the enduser_messages.squid.cache_admin_message node value
func GetEnduserMessagesSquidCacheAdminMessage(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.squid.cache_admin_message")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSquidCacheAdminMessage updates the enduser_messages.squid.cache_admin_message node value
func UpdateEnduserMessagesSquidCacheAdminMessage(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.squid.cache_admin_message", bytes.NewReader(byt))
	return
}

// GetEppAllowedNetworks gets the epp.allowed_networks node value
func GetEppAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/epp.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppAllowedNetworks updates the epp.allowed_networks node value
func UpdateEppAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.allowed_networks", bytes.NewReader(byt))
	return
}

// GetEppCertificate gets the epp.certificate node value
func GetEppCertificate(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.certificate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppCertificate updates the epp.certificate node value
func UpdateEppCertificate(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.certificate", bytes.NewReader(byt))
	return
}

// GetEppCity gets the epp.city node value
func GetEppCity(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.city")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppCity updates the epp.city node value
func UpdateEppCity(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.city", bytes.NewReader(byt))
	return
}

// GetEppCountry gets the epp.country node value
func GetEppCountry(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.country")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppCountry updates the epp.country node value
func UpdateEppCountry(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.country", bytes.NewReader(byt))
	return
}

// GetEppDefaultEndpointsGroup gets the epp.default_endpoints_group node value
func GetEppDefaultEndpointsGroup(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.default_endpoints_group")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppDefaultEndpointsGroup updates the epp.default_endpoints_group node value
func UpdateEppDefaultEndpointsGroup(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.default_endpoints_group", bytes.NewReader(byt))
	return
}

// GetEppDevices gets the epp.devices node value
func GetEppDevices(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/epp.devices")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppDevices updates the epp.devices node value
func UpdateEppDevices(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.devices", bytes.NewReader(byt))
	return
}

// GetEppEmail gets the epp.email node value
func GetEppEmail(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.email")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppEmail updates the epp.email node value
func UpdateEppEmail(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.email", bytes.NewReader(byt))
	return
}

// GetEppEndpoints gets the epp.endpoints node value
func GetEppEndpoints(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/epp.endpoints")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppEndpoints updates the epp.endpoints node value
func UpdateEppEndpoints(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.endpoints", bytes.NewReader(byt))
	return
}

// GetEppEndpointsGroups gets the epp.endpoints_groups node value
func GetEppEndpointsGroups(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/epp.endpoints_groups")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppEndpointsGroups updates the epp.endpoints_groups node value
func UpdateEppEndpointsGroups(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.endpoints_groups", bytes.NewReader(byt))
	return
}

// GetEppExceptionsAv gets the epp.exceptions.av node value
func GetEppExceptionsAv(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/epp.exceptions.av")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppExceptionsAv updates the epp.exceptions.av node value
func UpdateEppExceptionsAv(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.exceptions.av", bytes.NewReader(byt))
	return
}

// GetEppExceptionsDc gets the epp.exceptions.dc node value
func GetEppExceptionsDc(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/epp.exceptions.dc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppExceptionsDc updates the epp.exceptions.dc node value
func UpdateEppExceptionsDc(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.exceptions.dc", bytes.NewReader(byt))
	return
}

// GetEppFallbackUrl gets the epp.fallback_url node value
func GetEppFallbackUrl(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.fallback_url")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppFallbackUrl updates the epp.fallback_url node value
func UpdateEppFallbackUrl(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.fallback_url", bytes.NewReader(byt))
	return
}

// GetEppMagnetPassword gets the epp.magnet_password node value
func GetEppMagnetPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.magnet_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppMagnetPassword updates the epp.magnet_password node value
func UpdateEppMagnetPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.magnet_password", bytes.NewReader(byt))
	return
}

// GetEppMagnetUsername gets the epp.magnet_username node value
func GetEppMagnetUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.magnet_username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppMagnetUsername updates the epp.magnet_username node value
func UpdateEppMagnetUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.magnet_username", bytes.NewReader(byt))
	return
}

// GetEppOrganization gets the epp.organization node value
func GetEppOrganization(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.organization")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppOrganization updates the epp.organization node value
func UpdateEppOrganization(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.organization", bytes.NewReader(byt))
	return
}

// GetEppParentProxyHost gets the epp.parent_proxy_host node value
func GetEppParentProxyHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.parent_proxy_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppParentProxyHost updates the epp.parent_proxy_host node value
func UpdateEppParentProxyHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.parent_proxy_host", bytes.NewReader(byt))
	return
}

// GetEppParentProxyPort gets the epp.parent_proxy_port node value
func GetEppParentProxyPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/epp.parent_proxy_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppParentProxyPort updates the epp.parent_proxy_port node value
func UpdateEppParentProxyPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.parent_proxy_port", bytes.NewReader(byt))
	return
}

// GetEppParentProxyStatus gets the epp.parent_proxy_status node value
func GetEppParentProxyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/epp.parent_proxy_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppParentProxyStatus updates the epp.parent_proxy_status node value
func UpdateEppParentProxyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.parent_proxy_status", bytes.NewReader(byt))
	return
}

// GetEppPoliciesAv gets the epp.policies.av node value
func GetEppPoliciesAv(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/epp.policies.av")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppPoliciesAv updates the epp.policies.av node value
func UpdateEppPoliciesAv(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.policies.av", bytes.NewReader(byt))
	return
}

// GetEppPoliciesDc gets the epp.policies.dc node value
func GetEppPoliciesDc(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/epp.policies.dc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppPoliciesDc updates the epp.policies.dc node value
func UpdateEppPoliciesDc(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.policies.dc", bytes.NewReader(byt))
	return
}

// GetEppPort gets the epp.port node value
func GetEppPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/epp.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppPort updates the epp.port node value
func UpdateEppPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.port", bytes.NewReader(byt))
	return
}

// GetEppPrivateKey gets the epp.private_key node value
func GetEppPrivateKey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.private_key")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppPrivateKey updates the epp.private_key node value
func UpdateEppPrivateKey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.private_key", bytes.NewReader(byt))
	return
}

// GetEppRegistrationToken gets the epp.registration_token node value
func GetEppRegistrationToken(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.registration_token")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppRegistrationToken updates the epp.registration_token node value
func UpdateEppRegistrationToken(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.registration_token", bytes.NewReader(byt))
	return
}

// GetEppStatusAv gets the epp.status.av node value
func GetEppStatusAv(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/epp.status.av")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppStatusAv updates the epp.status.av node value
func UpdateEppStatusAv(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.status.av", bytes.NewReader(byt))
	return
}

// GetEppStatusBroker gets the epp.status.broker node value
func GetEppStatusBroker(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/epp.status.broker")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppStatusBroker updates the epp.status.broker node value
func UpdateEppStatusBroker(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.status.broker", bytes.NewReader(byt))
	return
}

// GetEppStatusDc gets the epp.status.dc node value
func GetEppStatusDc(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/epp.status.dc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppStatusDc updates the epp.status.dc node value
func UpdateEppStatusDc(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.status.dc", bytes.NewReader(byt))
	return
}

// GetEppStatusEpp gets the epp.status.epp node value
func GetEppStatusEpp(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/epp.status.epp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppStatusEpp updates the epp.status.epp node value
func UpdateEppStatusEpp(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.status.epp", bytes.NewReader(byt))
	return
}

// GetEppStatusWc gets the epp.status.wc node value
func GetEppStatusWc(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/epp.status.wc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppStatusWc updates the epp.status.wc node value
func UpdateEppStatusWc(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.status.wc", bytes.NewReader(byt))
	return
}

// GetEppTamperPassword gets the epp.tamper_password node value
func GetEppTamperPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.tamper_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppTamperPassword updates the epp.tamper_password node value
func UpdateEppTamperPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.tamper_password", bytes.NewReader(byt))
	return
}

// GetEppVersion gets the epp.version node value
func GetEppVersion(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.version")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppVersion updates the epp.version node value
func UpdateEppVersion(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.version", bytes.NewReader(byt))
	return
}

// GetEppWdxToken gets the epp.wdx_token node value
func GetEppWdxToken(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.wdx_token")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppWdxToken updates the epp.wdx_token node value
func UpdateEppWdxToken(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.wdx_token", bytes.NewReader(byt))
	return
}

// GetExecutiveReportDailyArchive gets the executive_report.daily.archive node value
func GetExecutiveReportDailyArchive(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/executive_report.daily.archive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportDailyArchive updates the executive_report.daily.archive node value
func UpdateExecutiveReportDailyArchive(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.daily.archive", bytes.NewReader(byt))
	return
}

// GetExecutiveReportDailyKeep gets the executive_report.daily.keep node value
func GetExecutiveReportDailyKeep(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/executive_report.daily.keep")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportDailyKeep updates the executive_report.daily.keep node value
func UpdateExecutiveReportDailyKeep(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.daily.keep", bytes.NewReader(byt))
	return
}

// GetExecutiveReportDailyPdfrecipients gets the executive_report.daily.pdfrecipients node value
func GetExecutiveReportDailyPdfrecipients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/executive_report.daily.pdfrecipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportDailyPdfrecipients updates the executive_report.daily.pdfrecipients node value
func UpdateExecutiveReportDailyPdfrecipients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.daily.pdfrecipients", bytes.NewReader(byt))
	return
}

// GetExecutiveReportDailyRecipients gets the executive_report.daily.recipients node value
func GetExecutiveReportDailyRecipients(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/executive_report.daily.recipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportDailyRecipients updates the executive_report.daily.recipients node value
func UpdateExecutiveReportDailyRecipients(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.daily.recipients", bytes.NewReader(byt))
	return
}

// GetExecutiveReportDailyStatus gets the executive_report.daily.status node value
func GetExecutiveReportDailyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/executive_report.daily.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportDailyStatus updates the executive_report.daily.status node value
func UpdateExecutiveReportDailyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.daily.status", bytes.NewReader(byt))
	return
}

// GetExecutiveReportMonthlyArchive gets the executive_report.monthly.archive node value
func GetExecutiveReportMonthlyArchive(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/executive_report.monthly.archive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportMonthlyArchive updates the executive_report.monthly.archive node value
func UpdateExecutiveReportMonthlyArchive(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.monthly.archive", bytes.NewReader(byt))
	return
}

// GetExecutiveReportMonthlyKeep gets the executive_report.monthly.keep node value
func GetExecutiveReportMonthlyKeep(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/executive_report.monthly.keep")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportMonthlyKeep updates the executive_report.monthly.keep node value
func UpdateExecutiveReportMonthlyKeep(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.monthly.keep", bytes.NewReader(byt))
	return
}

// GetExecutiveReportMonthlyPdfrecipients gets the executive_report.monthly.pdfrecipients node value
func GetExecutiveReportMonthlyPdfrecipients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/executive_report.monthly.pdfrecipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportMonthlyPdfrecipients updates the executive_report.monthly.pdfrecipients node value
func UpdateExecutiveReportMonthlyPdfrecipients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.monthly.pdfrecipients", bytes.NewReader(byt))
	return
}

// GetExecutiveReportMonthlyRecipients gets the executive_report.monthly.recipients node value
func GetExecutiveReportMonthlyRecipients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/executive_report.monthly.recipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportMonthlyRecipients updates the executive_report.monthly.recipients node value
func UpdateExecutiveReportMonthlyRecipients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.monthly.recipients", bytes.NewReader(byt))
	return
}

// GetExecutiveReportMonthlyStatus gets the executive_report.monthly.status node value
func GetExecutiveReportMonthlyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/executive_report.monthly.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportMonthlyStatus updates the executive_report.monthly.status node value
func UpdateExecutiveReportMonthlyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.monthly.status", bytes.NewReader(byt))
	return
}

// GetExecutiveReportWeeklyArchive gets the executive_report.weekly.archive node value
func GetExecutiveReportWeeklyArchive(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/executive_report.weekly.archive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportWeeklyArchive updates the executive_report.weekly.archive node value
func UpdateExecutiveReportWeeklyArchive(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.weekly.archive", bytes.NewReader(byt))
	return
}

// GetExecutiveReportWeeklyFirstDayOfWeek gets the executive_report.weekly.first_day_of_week node value
func GetExecutiveReportWeeklyFirstDayOfWeek(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/executive_report.weekly.first_day_of_week")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportWeeklyFirstDayOfWeek updates the executive_report.weekly.first_day_of_week node value
func UpdateExecutiveReportWeeklyFirstDayOfWeek(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.weekly.first_day_of_week", bytes.NewReader(byt))
	return
}

// GetExecutiveReportWeeklyKeep gets the executive_report.weekly.keep node value
func GetExecutiveReportWeeklyKeep(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/executive_report.weekly.keep")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportWeeklyKeep updates the executive_report.weekly.keep node value
func UpdateExecutiveReportWeeklyKeep(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.weekly.keep", bytes.NewReader(byt))
	return
}

// GetExecutiveReportWeeklyPdfrecipients gets the executive_report.weekly.pdfrecipients node value
func GetExecutiveReportWeeklyPdfrecipients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/executive_report.weekly.pdfrecipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportWeeklyPdfrecipients updates the executive_report.weekly.pdfrecipients node value
func UpdateExecutiveReportWeeklyPdfrecipients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.weekly.pdfrecipients", bytes.NewReader(byt))
	return
}

// GetExecutiveReportWeeklyRecipients gets the executive_report.weekly.recipients node value
func GetExecutiveReportWeeklyRecipients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/executive_report.weekly.recipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportWeeklyRecipients updates the executive_report.weekly.recipients node value
func UpdateExecutiveReportWeeklyRecipients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.weekly.recipients", bytes.NewReader(byt))
	return
}

// GetExecutiveReportWeeklyStatus gets the executive_report.weekly.status node value
func GetExecutiveReportWeeklyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/executive_report.weekly.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportWeeklyStatus updates the executive_report.weekly.status node value
func UpdateExecutiveReportWeeklyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.weekly.status", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpDstBurst gets the flood_protection.icmp.dst_burst node value
func GetFloodProtectionIcmpDstBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.dst_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpDstBurst updates the flood_protection.icmp.dst_burst node value
func UpdateFloodProtectionIcmpDstBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.dst_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpDstExpire gets the flood_protection.icmp.dst_expire node value
func GetFloodProtectionIcmpDstExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.dst_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpDstExpire updates the flood_protection.icmp.dst_expire node value
func UpdateFloodProtectionIcmpDstExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.dst_expire", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpDstGcInterval gets the flood_protection.icmp.dst_gc_interval node value
func GetFloodProtectionIcmpDstGcInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.dst_gc_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpDstGcInterval updates the flood_protection.icmp.dst_gc_interval node value
func UpdateFloodProtectionIcmpDstGcInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.dst_gc_interval", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpDstRate gets the flood_protection.icmp.dst_rate node value
func GetFloodProtectionIcmpDstRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.dst_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpDstRate updates the flood_protection.icmp.dst_rate node value
func UpdateFloodProtectionIcmpDstRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.dst_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpLog gets the flood_protection.icmp.log node value
func GetFloodProtectionIcmpLog(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.log")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpLog updates the flood_protection.icmp.log node value
func UpdateFloodProtectionIcmpLog(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.log", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpLogLimitBurst gets the flood_protection.icmp.log_limit_burst node value
func GetFloodProtectionIcmpLogLimitBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.log_limit_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpLogLimitBurst updates the flood_protection.icmp.log_limit_burst node value
func UpdateFloodProtectionIcmpLogLimitBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.log_limit_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpLogLimitRate gets the flood_protection.icmp.log_limit_rate node value
func GetFloodProtectionIcmpLogLimitRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.log_limit_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpLogLimitRate updates the flood_protection.icmp.log_limit_rate node value
func UpdateFloodProtectionIcmpLogLimitRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.log_limit_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpMode gets the flood_protection.icmp.mode node value
func GetFloodProtectionIcmpMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpMode updates the flood_protection.icmp.mode node value
func UpdateFloodProtectionIcmpMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.mode", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpSrcBurst gets the flood_protection.icmp.src_burst node value
func GetFloodProtectionIcmpSrcBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.src_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpSrcBurst updates the flood_protection.icmp.src_burst node value
func UpdateFloodProtectionIcmpSrcBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.src_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpSrcExpire gets the flood_protection.icmp.src_expire node value
func GetFloodProtectionIcmpSrcExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.src_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpSrcExpire updates the flood_protection.icmp.src_expire node value
func UpdateFloodProtectionIcmpSrcExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.src_expire", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpSrcGcInterval gets the flood_protection.icmp.src_gc_interval node value
func GetFloodProtectionIcmpSrcGcInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.src_gc_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpSrcGcInterval updates the flood_protection.icmp.src_gc_interval node value
func UpdateFloodProtectionIcmpSrcGcInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.src_gc_interval", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpSrcRate gets the flood_protection.icmp.src_rate node value
func GetFloodProtectionIcmpSrcRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.src_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpSrcRate updates the flood_protection.icmp.src_rate node value
func UpdateFloodProtectionIcmpSrcRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.src_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpStatus gets the flood_protection.icmp.status node value
func GetFloodProtectionIcmpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpStatus updates the flood_protection.icmp.status node value
func UpdateFloodProtectionIcmpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.status", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynDstBurst gets the flood_protection.syn.dst_burst node value
func GetFloodProtectionSynDstBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.dst_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynDstBurst updates the flood_protection.syn.dst_burst node value
func UpdateFloodProtectionSynDstBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.dst_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynDstExpire gets the flood_protection.syn.dst_expire node value
func GetFloodProtectionSynDstExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.dst_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynDstExpire updates the flood_protection.syn.dst_expire node value
func UpdateFloodProtectionSynDstExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.dst_expire", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynDstGcInterval gets the flood_protection.syn.dst_gc_interval node value
func GetFloodProtectionSynDstGcInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.dst_gc_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynDstGcInterval updates the flood_protection.syn.dst_gc_interval node value
func UpdateFloodProtectionSynDstGcInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.dst_gc_interval", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynDstRate gets the flood_protection.syn.dst_rate node value
func GetFloodProtectionSynDstRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.dst_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynDstRate updates the flood_protection.syn.dst_rate node value
func UpdateFloodProtectionSynDstRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.dst_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynLog gets the flood_protection.syn.log node value
func GetFloodProtectionSynLog(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.log")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynLog updates the flood_protection.syn.log node value
func UpdateFloodProtectionSynLog(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.log", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynLogLimitBurst gets the flood_protection.syn.log_limit_burst node value
func GetFloodProtectionSynLogLimitBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.log_limit_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynLogLimitBurst updates the flood_protection.syn.log_limit_burst node value
func UpdateFloodProtectionSynLogLimitBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.log_limit_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynLogLimitRate gets the flood_protection.syn.log_limit_rate node value
func GetFloodProtectionSynLogLimitRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.log_limit_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynLogLimitRate updates the flood_protection.syn.log_limit_rate node value
func UpdateFloodProtectionSynLogLimitRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.log_limit_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynMode gets the flood_protection.syn.mode node value
func GetFloodProtectionSynMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynMode updates the flood_protection.syn.mode node value
func UpdateFloodProtectionSynMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.mode", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynSrcBurst gets the flood_protection.syn.src_burst node value
func GetFloodProtectionSynSrcBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.src_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynSrcBurst updates the flood_protection.syn.src_burst node value
func UpdateFloodProtectionSynSrcBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.src_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynSrcExpire gets the flood_protection.syn.src_expire node value
func GetFloodProtectionSynSrcExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.src_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynSrcExpire updates the flood_protection.syn.src_expire node value
func UpdateFloodProtectionSynSrcExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.src_expire", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynSrcGcInterval gets the flood_protection.syn.src_gc_interval node value
func GetFloodProtectionSynSrcGcInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.src_gc_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynSrcGcInterval updates the flood_protection.syn.src_gc_interval node value
func UpdateFloodProtectionSynSrcGcInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.src_gc_interval", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynSrcRate gets the flood_protection.syn.src_rate node value
func GetFloodProtectionSynSrcRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.src_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynSrcRate updates the flood_protection.syn.src_rate node value
func UpdateFloodProtectionSynSrcRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.src_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynStatus gets the flood_protection.syn.status node value
func GetFloodProtectionSynStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynStatus updates the flood_protection.syn.status node value
func UpdateFloodProtectionSynStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.status", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpDstBurst gets the flood_protection.udp.dst_burst node value
func GetFloodProtectionUdpDstBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.dst_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpDstBurst updates the flood_protection.udp.dst_burst node value
func UpdateFloodProtectionUdpDstBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.dst_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpDstExpire gets the flood_protection.udp.dst_expire node value
func GetFloodProtectionUdpDstExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.dst_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpDstExpire updates the flood_protection.udp.dst_expire node value
func UpdateFloodProtectionUdpDstExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.dst_expire", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpDstGcInterval gets the flood_protection.udp.dst_gc_interval node value
func GetFloodProtectionUdpDstGcInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.dst_gc_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpDstGcInterval updates the flood_protection.udp.dst_gc_interval node value
func UpdateFloodProtectionUdpDstGcInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.dst_gc_interval", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpDstRate gets the flood_protection.udp.dst_rate node value
func GetFloodProtectionUdpDstRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.dst_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpDstRate updates the flood_protection.udp.dst_rate node value
func UpdateFloodProtectionUdpDstRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.dst_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpLog gets the flood_protection.udp.log node value
func GetFloodProtectionUdpLog(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.log")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpLog updates the flood_protection.udp.log node value
func UpdateFloodProtectionUdpLog(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.log", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpLogLimitBurst gets the flood_protection.udp.log_limit_burst node value
func GetFloodProtectionUdpLogLimitBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.log_limit_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpLogLimitBurst updates the flood_protection.udp.log_limit_burst node value
func UpdateFloodProtectionUdpLogLimitBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.log_limit_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpLogLimitRate gets the flood_protection.udp.log_limit_rate node value
func GetFloodProtectionUdpLogLimitRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.log_limit_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpLogLimitRate updates the flood_protection.udp.log_limit_rate node value
func UpdateFloodProtectionUdpLogLimitRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.log_limit_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpMode gets the flood_protection.udp.mode node value
func GetFloodProtectionUdpMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpMode updates the flood_protection.udp.mode node value
func UpdateFloodProtectionUdpMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.mode", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpSrcBurst gets the flood_protection.udp.src_burst node value
func GetFloodProtectionUdpSrcBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.src_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpSrcBurst updates the flood_protection.udp.src_burst node value
func UpdateFloodProtectionUdpSrcBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.src_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpSrcExpire gets the flood_protection.udp.src_expire node value
func GetFloodProtectionUdpSrcExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.src_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpSrcExpire updates the flood_protection.udp.src_expire node value
func UpdateFloodProtectionUdpSrcExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.src_expire", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpSrcGcInterval gets the flood_protection.udp.src_gc_interval node value
func GetFloodProtectionUdpSrcGcInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.src_gc_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpSrcGcInterval updates the flood_protection.udp.src_gc_interval node value
func UpdateFloodProtectionUdpSrcGcInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.src_gc_interval", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpSrcRate gets the flood_protection.udp.src_rate node value
func GetFloodProtectionUdpSrcRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.src_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpSrcRate updates the flood_protection.udp.src_rate node value
func UpdateFloodProtectionUdpSrcRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.src_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpStatus gets the flood_protection.udp.status node value
func GetFloodProtectionUdpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpStatus updates the flood_protection.udp.status node value
func UpdateFloodProtectionUdpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.status", bytes.NewReader(byt))
	return
}

// GetFtpAllowedClients gets the ftp.allowed_clients node value
func GetFtpAllowedClients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ftp.allowed_clients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpAllowedClients updates the ftp.allowed_clients node value
func UpdateFtpAllowedClients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.allowed_clients", bytes.NewReader(byt))
	return
}

// GetFtpAllowedServers gets the ftp.allowed_servers node value
func GetFtpAllowedServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ftp.allowed_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpAllowedServers updates the ftp.allowed_servers node value
func UpdateFtpAllowedServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.allowed_servers", bytes.NewReader(byt))
	return
}

// GetFtpCffAv gets the ftp.cff_av node value
func GetFtpCffAv(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ftp.cff_av")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpCffAv updates the ftp.cff_av node value
func UpdateFtpCffAv(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.cff_av", bytes.NewReader(byt))
	return
}

// GetFtpCffAvEngines gets the ftp.cff_av_engines node value
func GetFtpCffAvEngines(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ftp.cff_av_engines")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpCffAvEngines updates the ftp.cff_av_engines node value
func UpdateFtpCffAvEngines(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.cff_av_engines", bytes.NewReader(byt))
	return
}

// GetFtpCffFileExtensions gets the ftp.cff_file_extensions node value
func GetFtpCffFileExtensions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ftp.cff_file_extensions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpCffFileExtensions updates the ftp.cff_file_extensions node value
func UpdateFtpCffFileExtensions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.cff_file_extensions", bytes.NewReader(byt))
	return
}

// GetFtpExceptions gets the ftp.exceptions node value
func GetFtpExceptions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ftp.exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpExceptions updates the ftp.exceptions node value
func UpdateFtpExceptions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.exceptions", bytes.NewReader(byt))
	return
}

// GetFtpMaxFileSize gets the ftp.max_file_size node value
func GetFtpMaxFileSize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ftp.max_file_size")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpMaxFileSize updates the ftp.max_file_size node value
func UpdateFtpMaxFileSize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.max_file_size", bytes.NewReader(byt))
	return
}

// GetFtpMsWinMode gets the ftp.ms_win_mode node value
func GetFtpMsWinMode(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ftp.ms_win_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpMsWinMode updates the ftp.ms_win_mode node value
func UpdateFtpMsWinMode(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.ms_win_mode", bytes.NewReader(byt))
	return
}

// GetFtpOperationMode gets the ftp.operation_mode node value
func GetFtpOperationMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ftp.operation_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpOperationMode updates the ftp.operation_mode node value
func UpdateFtpOperationMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.operation_mode", bytes.NewReader(byt))
	return
}

// GetFtpRestrictedServers gets the ftp.restricted_servers node value
func GetFtpRestrictedServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ftp.restricted_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpRestrictedServers updates the ftp.restricted_servers node value
func UpdateFtpRestrictedServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.restricted_servers", bytes.NewReader(byt))
	return
}

// GetFtpStatus gets the ftp.status node value
func GetFtpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ftp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpStatus updates the ftp.status node value
func UpdateFtpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.status", bytes.NewReader(byt))
	return
}

// GetFtpTransparentSkip gets the ftp.transparent_skip node value
func GetFtpTransparentSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ftp.transparent_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpTransparentSkip updates the ftp.transparent_skip node value
func UpdateFtpTransparentSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.transparent_skip", bytes.NewReader(byt))
	return
}

// GetFtpTransparentSkipAutoPf gets the ftp.transparent_skip_auto_pf node value
func GetFtpTransparentSkipAutoPf(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ftp.transparent_skip_auto_pf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpTransparentSkipAutoPf updates the ftp.transparent_skip_auto_pf node value
func UpdateFtpTransparentSkipAutoPf(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.transparent_skip_auto_pf", bytes.NewReader(byt))
	return
}

// GetGenericProxyRules gets the generic_proxy.rules node value
func GetGenericProxyRules(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/generic_proxy.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateGenericProxyRules updates the generic_proxy.rules node value
func UpdateGenericProxyRules(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/generic_proxy.rules", bytes.NewReader(byt))
	return
}

// GetGeoipCountriesDst gets the geoip.countries_dst node value
func GetGeoipCountriesDst(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/geoip.countries_dst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateGeoipCountriesDst updates the geoip.countries_dst node value
func UpdateGeoipCountriesDst(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/geoip.countries_dst", bytes.NewReader(byt))
	return
}

// GetGeoipCountriesSrc gets the geoip.countries_src node value
func GetGeoipCountriesSrc(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/geoip.countries_src")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateGeoipCountriesSrc updates the geoip.countries_src node value
func UpdateGeoipCountriesSrc(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/geoip.countries_src", bytes.NewReader(byt))
	return
}

// GetGeoipExceptions gets the geoip.exceptions node value
func GetGeoipExceptions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/geoip.exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateGeoipExceptions updates the geoip.exceptions node value
func UpdateGeoipExceptions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/geoip.exceptions", bytes.NewReader(byt))
	return
}

// GetGeoipLog gets the geoip.log node value
func GetGeoipLog(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/geoip.log")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateGeoipLog updates the geoip.log node value
func UpdateGeoipLog(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/geoip.log", bytes.NewReader(byt))
	return
}

// GetGeoipStatus gets the geoip.status node value
func GetGeoipStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/geoip.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateGeoipStatus updates the geoip.status node value
func UpdateGeoipStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/geoip.status", bytes.NewReader(byt))
	return
}

// GetH323AllowedNetworks gets the h323.allowed_networks node value
func GetH323AllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/h323.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateH323AllowedNetworks updates the h323.allowed_networks node value
func UpdateH323AllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/h323.allowed_networks", bytes.NewReader(byt))
	return
}

// GetH323LogRelated gets the h323.log_related node value
func GetH323LogRelated(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/h323.log_related")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateH323LogRelated updates the h323.log_related node value
func UpdateH323LogRelated(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/h323.log_related", bytes.NewReader(byt))
	return
}

// GetH323Servers gets the h323.servers node value
func GetH323Servers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/h323.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateH323Servers updates the h323.servers node value
func UpdateH323Servers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/h323.servers", bytes.NewReader(byt))
	return
}

// GetH323Status gets the h323.status node value
func GetH323Status(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/h323.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateH323Status updates the h323.status node value
func UpdateH323Status(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/h323.status", bytes.NewReader(byt))
	return
}

// GetHaAdvancedAutojoin gets the ha.advanced.autojoin node value
func GetHaAdvancedAutojoin(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.autojoin")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedAutojoin updates the ha.advanced.autojoin node value
func UpdateHaAdvancedAutojoin(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.autojoin", bytes.NewReader(byt))
	return
}

// GetHaAdvancedColdRollback gets the ha.advanced.cold_rollback node value
func GetHaAdvancedColdRollback(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.cold_rollback")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedColdRollback updates the ha.advanced.cold_rollback node value
func UpdateHaAdvancedColdRollback(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.cold_rollback", bytes.NewReader(byt))
	return
}

// GetHaAdvancedHttpPersistenceTime gets the ha.advanced.http_persistence_time node value
func GetHaAdvancedHttpPersistenceTime(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.http_persistence_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedHttpPersistenceTime updates the ha.advanced.http_persistence_time node value
func UpdateHaAdvancedHttpPersistenceTime(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.http_persistence_time", bytes.NewReader(byt))
	return
}

// GetHaAdvancedLoadTakeover gets the ha.advanced.load_takeover node value
func GetHaAdvancedLoadTakeover(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.load_takeover")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedLoadTakeover updates the ha.advanced.load_takeover node value
func UpdateHaAdvancedLoadTakeover(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.load_takeover", bytes.NewReader(byt))
	return
}

// GetHaAdvancedLoadWarn gets the ha.advanced.load_warn node value
func GetHaAdvancedLoadWarn(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.load_warn")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedLoadWarn updates the ha.advanced.load_warn node value
func UpdateHaAdvancedLoadWarn(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.load_warn", bytes.NewReader(byt))
	return
}

// GetHaAdvancedMaxNodes gets the ha.advanced.max_nodes node value
func GetHaAdvancedMaxNodes(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.max_nodes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedMaxNodes updates the ha.advanced.max_nodes node value
func UpdateHaAdvancedMaxNodes(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.max_nodes", bytes.NewReader(byt))
	return
}

// GetHaAdvancedMtu gets the ha.advanced.mtu node value
func GetHaAdvancedMtu(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.mtu")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedMtu updates the ha.advanced.mtu node value
func UpdateHaAdvancedMtu(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.mtu", bytes.NewReader(byt))
	return
}

// GetHaAdvancedNetconsole gets the ha.advanced.netconsole node value
func GetHaAdvancedNetconsole(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.netconsole")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedNetconsole updates the ha.advanced.netconsole node value
func UpdateHaAdvancedNetconsole(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.netconsole", bytes.NewReader(byt))
	return
}

// GetHaAdvancedPreempt gets the ha.advanced.preempt node value
func GetHaAdvancedPreempt(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.preempt")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedPreempt updates the ha.advanced.preempt node value
func UpdateHaAdvancedPreempt(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.preempt", bytes.NewReader(byt))
	return
}

// GetHaAdvancedUniqueId gets the ha.advanced.unique_id node value
func GetHaAdvancedUniqueId(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.unique_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedUniqueId updates the ha.advanced.unique_id node value
func UpdateHaAdvancedUniqueId(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.unique_id", bytes.NewReader(byt))
	return
}

// GetHaAdvancedVirtualMac gets the ha.advanced.virtual_mac node value
func GetHaAdvancedVirtualMac(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.virtual_mac")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedVirtualMac updates the ha.advanced.virtual_mac node value
func UpdateHaAdvancedVirtualMac(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.virtual_mac", bytes.NewReader(byt))
	return
}

// GetHaAwsCloudwatchProfile gets the ha.aws.cloudwatch.profile node value
func GetHaAwsCloudwatchProfile(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.aws.cloudwatch.profile")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsCloudwatchProfile updates the ha.aws.cloudwatch.profile node value
func UpdateHaAwsCloudwatchProfile(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.cloudwatch.profile", bytes.NewReader(byt))
	return
}

// GetHaAwsCloudwatchStatus gets the ha.aws.cloudwatch.status node value
func GetHaAwsCloudwatchStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.cloudwatch.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsCloudwatchStatus updates the ha.aws.cloudwatch.status node value
func UpdateHaAwsCloudwatchStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.cloudwatch.status", bytes.NewReader(byt))
	return
}

// GetHaAwsConfdBackup gets the ha.aws.confd.backup node value
func GetHaAwsConfdBackup(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.confd.backup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsConfdBackup updates the ha.aws.confd.backup node value
func UpdateHaAwsConfdBackup(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.confd.backup", bytes.NewReader(byt))
	return
}

// GetHaAwsConfdBackupInterval gets the ha.aws.confd.backup_interval node value
func GetHaAwsConfdBackupInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.aws.confd.backup_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsConfdBackupInterval updates the ha.aws.confd.backup_interval node value
func UpdateHaAwsConfdBackupInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.confd.backup_interval", bytes.NewReader(byt))
	return
}

// GetHaAwsConfdRestore gets the ha.aws.confd.restore node value
func GetHaAwsConfdRestore(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.confd.restore")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsConfdRestore updates the ha.aws.confd.restore node value
func UpdateHaAwsConfdRestore(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.confd.restore", bytes.NewReader(byt))
	return
}

// GetHaAwsConfdRestoreDone gets the ha.aws.confd.restore_done node value
func GetHaAwsConfdRestoreDone(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.confd.restore_done")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsConfdRestoreDone updates the ha.aws.confd.restore_done node value
func UpdateHaAwsConfdRestoreDone(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.confd.restore_done", bytes.NewReader(byt))
	return
}

// GetHaAwsElasticIp gets the ha.aws.elastic_ip node value
func GetHaAwsElasticIp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.aws.elastic_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsElasticIp updates the ha.aws.elastic_ip node value
func UpdateHaAwsElasticIp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.elastic_ip", bytes.NewReader(byt))
	return
}

// GetHaAwsPostgresArchiveTimeout gets the ha.aws.postgres.archive_timeout node value
func GetHaAwsPostgresArchiveTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.aws.postgres.archive_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsPostgresArchiveTimeout updates the ha.aws.postgres.archive_timeout node value
func UpdateHaAwsPostgresArchiveTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.postgres.archive_timeout", bytes.NewReader(byt))
	return
}

// GetHaAwsPostgresBackup gets the ha.aws.postgres.backup node value
func GetHaAwsPostgresBackup(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.postgres.backup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsPostgresBackup updates the ha.aws.postgres.backup node value
func UpdateHaAwsPostgresBackup(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.postgres.backup", bytes.NewReader(byt))
	return
}

// GetHaAwsPostgresBaseBackupInterval gets the ha.aws.postgres.base_backup_interval node value
func GetHaAwsPostgresBaseBackupInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.aws.postgres.base_backup_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsPostgresBaseBackupInterval updates the ha.aws.postgres.base_backup_interval node value
func UpdateHaAwsPostgresBaseBackupInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.postgres.base_backup_interval", bytes.NewReader(byt))
	return
}

// GetHaAwsPostgresRestore gets the ha.aws.postgres.restore node value
func GetHaAwsPostgresRestore(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.postgres.restore")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsPostgresRestore updates the ha.aws.postgres.restore node value
func UpdateHaAwsPostgresRestore(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.postgres.restore", bytes.NewReader(byt))
	return
}

// GetHaAwsS3Bucket gets the ha.aws.s3_bucket node value
func GetHaAwsS3Bucket(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.aws.s3_bucket")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsS3Bucket updates the ha.aws.s3_bucket node value
func UpdateHaAwsS3Bucket(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.s3_bucket", bytes.NewReader(byt))
	return
}

// GetHaAwsStackName gets the ha.aws.stack_name node value
func GetHaAwsStackName(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.aws.stack_name")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsStackName updates the ha.aws.stack_name node value
func UpdateHaAwsStackName(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.stack_name", bytes.NewReader(byt))
	return
}

// GetHaAwsSyslogBackup gets the ha.aws.syslog.backup node value
func GetHaAwsSyslogBackup(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.syslog.backup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsSyslogBackup updates the ha.aws.syslog.backup node value
func UpdateHaAwsSyslogBackup(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.syslog.backup", bytes.NewReader(byt))
	return
}

// GetHaAwsSyslogRestore gets the ha.aws.syslog.restore node value
func GetHaAwsSyslogRestore(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.syslog.restore")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsSyslogRestore updates the ha.aws.syslog.restore node value
func UpdateHaAwsSyslogRestore(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.syslog.restore", bytes.NewReader(byt))
	return
}

// GetHaAwsSyslogRestorePeriod gets the ha.aws.syslog.restore_period node value
func GetHaAwsSyslogRestorePeriod(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.aws.syslog.restore_period")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsSyslogRestorePeriod updates the ha.aws.syslog.restore_period node value
func UpdateHaAwsSyslogRestorePeriod(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.syslog.restore_period", bytes.NewReader(byt))
	return
}

// GetHaAwsTrustedNetwork gets the ha.aws.trusted_network node value
func GetHaAwsTrustedNetwork(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.aws.trusted_network")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsTrustedNetwork updates the ha.aws.trusted_network node value
func UpdateHaAwsTrustedNetwork(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.trusted_network", bytes.NewReader(byt))
	return
}

// GetHaClusterFtp gets the ha.cluster.ftp node value
func GetHaClusterFtp(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.ftp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterFtp updates the ha.cluster.ftp node value
func UpdateHaClusterFtp(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.ftp", bytes.NewReader(byt))
	return
}

// GetHaClusterHttp gets the ha.cluster.http node value
func GetHaClusterHttp(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.http")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterHttp updates the ha.cluster.http node value
func UpdateHaClusterHttp(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.http", bytes.NewReader(byt))
	return
}

// GetHaClusterIpsec gets the ha.cluster.ipsec node value
func GetHaClusterIpsec(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.ipsec")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterIpsec updates the ha.cluster.ipsec node value
func UpdateHaClusterIpsec(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.ipsec", bytes.NewReader(byt))
	return
}

// GetHaClusterPop3 gets the ha.cluster.pop3 node value
func GetHaClusterPop3(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.pop3")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterPop3 updates the ha.cluster.pop3 node value
func UpdateHaClusterPop3(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.pop3", bytes.NewReader(byt))
	return
}

// GetHaClusterSmtp gets the ha.cluster.smtp node value
func GetHaClusterSmtp(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.smtp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterSmtp updates the ha.cluster.smtp node value
func UpdateHaClusterSmtp(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.smtp", bytes.NewReader(byt))
	return
}

// GetHaClusterSnort gets the ha.cluster.snort node value
func GetHaClusterSnort(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.snort")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterSnort updates the ha.cluster.snort node value
func UpdateHaClusterSnort(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.snort", bytes.NewReader(byt))
	return
}

// GetHaClusterWaf gets the ha.cluster.waf node value
func GetHaClusterWaf(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.waf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterWaf updates the ha.cluster.waf node value
func UpdateHaClusterWaf(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.waf", bytes.NewReader(byt))
	return
}

// GetHaDeviceName gets the ha.device_name node value
func GetHaDeviceName(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.device_name")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaDeviceName updates the ha.device_name node value
func UpdateHaDeviceName(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.device_name", bytes.NewReader(byt))
	return
}

// GetHaItfhw gets the ha.itfhw node value
func GetHaItfhw(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.itfhw")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaItfhw updates the ha.itfhw node value
func UpdateHaItfhw(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.itfhw", bytes.NewReader(byt))
	return
}

// GetHaItfhwBackup gets the ha.itfhw_backup node value
func GetHaItfhwBackup(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.itfhw_backup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaItfhwBackup updates the ha.itfhw_backup node value
func UpdateHaItfhwBackup(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.itfhw_backup", bytes.NewReader(byt))
	return
}

// GetHaMasterIp gets the ha.master_ip node value
func GetHaMasterIp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.master_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaMasterIp updates the ha.master_ip node value
func UpdateHaMasterIp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.master_ip", bytes.NewReader(byt))
	return
}

// GetHaMode gets the ha.mode node value
func GetHaMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaMode updates the ha.mode node value
func UpdateHaMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.mode", bytes.NewReader(byt))
	return
}

// GetHaNodeId gets the ha.node_id node value
func GetHaNodeId(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.node_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaNodeId updates the ha.node_id node value
func UpdateHaNodeId(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.node_id", bytes.NewReader(byt))
	return
}

// GetHaPassword gets the ha.password node value
func GetHaPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaPassword updates the ha.password node value
func UpdateHaPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.password", bytes.NewReader(byt))
	return
}

// GetHaPostgresSecret gets the ha.postgres_secret node value
func GetHaPostgresSecret(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.postgres_secret")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaPostgresSecret updates the ha.postgres_secret node value
func UpdateHaPostgresSecret(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.postgres_secret", bytes.NewReader(byt))
	return
}

// GetHaSlaveIp gets the ha.slave_ip node value
func GetHaSlaveIp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.slave_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaSlaveIp updates the ha.slave_ip node value
func UpdateHaSlaveIp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.slave_ip", bytes.NewReader(byt))
	return
}

// GetHaStatus gets the ha.status node value
func GetHaStatus(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaStatus updates the ha.status node value
func UpdateHaStatus(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.status", bytes.NewReader(byt))
	return
}

// GetHaSyncConntrack gets the ha.sync.conntrack node value
func GetHaSyncConntrack(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.sync.conntrack")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaSyncConntrack updates the ha.sync.conntrack node value
func UpdateHaSyncConntrack(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.sync.conntrack", bytes.NewReader(byt))
	return
}

// GetHaSyncDatabase gets the ha.sync.database node value
func GetHaSyncDatabase(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.sync.database")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaSyncDatabase updates the ha.sync.database node value
func UpdateHaSyncDatabase(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.sync.database", bytes.NewReader(byt))
	return
}

// GetHaSyncFiles gets the ha.sync.files node value
func GetHaSyncFiles(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.sync.files")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaSyncFiles updates the ha.sync.files node value
func UpdateHaSyncFiles(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.sync.files", bytes.NewReader(byt))
	return
}

// GetHaSyncIpsec gets the ha.sync.ipsec node value
func GetHaSyncIpsec(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.sync.ipsec")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaSyncIpsec updates the ha.sync.ipsec node value
func UpdateHaSyncIpsec(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.sync.ipsec", bytes.NewReader(byt))
	return
}

// GetHaSyncSyslog gets the ha.sync.syslog node value
func GetHaSyncSyslog(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.sync.syslog")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaSyncSyslog updates the ha.sync.syslog node value
func UpdateHaSyncSyslog(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.sync.syslog", bytes.NewReader(byt))
	return
}

// GetHaTimesDeadTime gets the ha.times.dead_time node value
func GetHaTimesDeadTime(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.times.dead_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaTimesDeadTime updates the ha.times.dead_time node value
func UpdateHaTimesDeadTime(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.times.dead_time", bytes.NewReader(byt))
	return
}

// GetHaTimesLoadTime gets the ha.times.load_time node value
func GetHaTimesLoadTime(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.times.load_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaTimesLoadTime updates the ha.times.load_time node value
func UpdateHaTimesLoadTime(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.times.load_time", bytes.NewReader(byt))
	return
}

// GetHotspotCert gets the hotspot.cert node value
func GetHotspotCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/hotspot.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHotspotCert updates the hotspot.cert node value
func UpdateHotspotCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/hotspot.cert", bytes.NewReader(byt))
	return
}

// GetHotspotDeleteDays gets the hotspot.delete_days node value
func GetHotspotDeleteDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/hotspot.delete_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHotspotDeleteDays updates the hotspot.delete_days node value
func UpdateHotspotDeleteDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/hotspot.delete_days", bytes.NewReader(byt))
	return
}

// GetHotspotSslPortal gets the hotspot.ssl_portal node value
func GetHotspotSslPortal(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/hotspot.ssl_portal")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHotspotSslPortal updates the hotspot.ssl_portal node value
func UpdateHotspotSslPortal(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/hotspot.ssl_portal", bytes.NewReader(byt))
	return
}

// GetHotspotStatus gets the hotspot.status node value
func GetHotspotStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/hotspot.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHotspotStatus updates the hotspot.status node value
func UpdateHotspotStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/hotspot.status", bytes.NewReader(byt))
	return
}

// GetHotspotTransparentSkip gets the hotspot.transparent_skip node value
func GetHotspotTransparentSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/hotspot.transparent_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHotspotTransparentSkip updates the hotspot.transparent_skip node value
func UpdateHotspotTransparentSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/hotspot.transparent_skip", bytes.NewReader(byt))
	return
}

// GetHttpAdSsoInterfaces gets the http.ad_sso_interfaces node value
func GetHttpAdSsoInterfaces(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.ad_sso_interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAdSsoInterfaces updates the http.ad_sso_interfaces node value
func UpdateHttpAdSsoInterfaces(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.ad_sso_interfaces", bytes.NewReader(byt))
	return
}

// GetHttpAdssoRedirectUseHostname gets the http.adsso_redirect_use_hostname node value
func GetHttpAdssoRedirectUseHostname(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.adsso_redirect_use_hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAdssoRedirectUseHostname updates the http.adsso_redirect_use_hostname node value
func UpdateHttpAdssoRedirectUseHostname(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.adsso_redirect_use_hostname", bytes.NewReader(byt))
	return
}

// GetHttpAllowSsl3 gets the http.allow_ssl3 node value
func GetHttpAllowSsl3(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.allow_ssl3")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAllowSsl3 updates the http.allow_ssl3 node value
func UpdateHttpAllowSsl3(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.allow_ssl3", bytes.NewReader(byt))
	return
}

// GetHttpAllowTls12 gets the http.allow_tls_1_2 node value
func GetHttpAllowTls12(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.allow_tls_1_2")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAllowTls12 updates the http.allow_tls_1_2 node value
func UpdateHttpAllowTls12(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.allow_tls_1_2", bytes.NewReader(byt))
	return
}

// GetHttpAllowedPuas gets the http.allowed_puas node value
func GetHttpAllowedPuas(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.allowed_puas")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAllowedPuas updates the http.allowed_puas node value
func UpdateHttpAllowedPuas(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.allowed_puas", bytes.NewReader(byt))
	return
}

// GetHttpAllowedTargetServices gets the http.allowed_target_services node value
func GetHttpAllowedTargetServices(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.allowed_target_services")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAllowedTargetServices updates the http.allowed_target_services node value
func UpdateHttpAllowedTargetServices(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.allowed_target_services", bytes.NewReader(byt))
	return
}

// GetHttpAuaMaxconns gets the http.aua_maxconns node value
func GetHttpAuaMaxconns(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.aua_maxconns")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAuaMaxconns updates the http.aua_maxconns node value
func UpdateHttpAuaMaxconns(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.aua_maxconns", bytes.NewReader(byt))
	return
}

// GetHttpAuaTimeout gets the http.aua_timeout node value
func GetHttpAuaTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.aua_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAuaTimeout updates the http.aua_timeout node value
func UpdateHttpAuaTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.aua_timeout", bytes.NewReader(byt))
	return
}

// GetHttpAuthCacheSize gets the http.auth_cache_size node value
func GetHttpAuthCacheSize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.auth_cache_size")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAuthCacheSize updates the http.auth_cache_size node value
func UpdateHttpAuthCacheSize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.auth_cache_size", bytes.NewReader(byt))
	return
}

// GetHttpAuthCacheTtl gets the http.auth_cache_ttl node value
func GetHttpAuthCacheTtl(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.auth_cache_ttl")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAuthCacheTtl updates the http.auth_cache_ttl node value
func UpdateHttpAuthCacheTtl(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.auth_cache_ttl", bytes.NewReader(byt))
	return
}

// GetHttpAuthRealm gets the http.auth_realm node value
func GetHttpAuthRealm(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.auth_realm")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAuthRealm updates the http.auth_realm node value
func UpdateHttpAuthRealm(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.auth_realm", bytes.NewReader(byt))
	return
}

// GetHttpAuthUsercacheTtl gets the http.auth_usercache_ttl node value
func GetHttpAuthUsercacheTtl(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.auth_usercache_ttl")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAuthUsercacheTtl updates the http.auth_usercache_ttl node value
func UpdateHttpAuthUsercacheTtl(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.auth_usercache_ttl", bytes.NewReader(byt))
	return
}

// GetHttpBlockUnscannable gets the http.block_unscannable node value
func GetHttpBlockUnscannable(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.block_unscannable")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpBlockUnscannable updates the http.block_unscannable node value
func UpdateHttpBlockUnscannable(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.block_unscannable", bytes.NewReader(byt))
	return
}

// GetHttpBypassStreaming gets the http.bypass_streaming node value
func GetHttpBypassStreaming(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.bypass_streaming")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpBypassStreaming updates the http.bypass_streaming node value
func UpdateHttpBypassStreaming(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.bypass_streaming", bytes.NewReader(byt))
	return
}

// GetHttpCaList gets the http.ca_list node value
func GetHttpCaList(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.ca_list")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCaList updates the http.ca_list node value
func UpdateHttpCaList(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.ca_list", bytes.NewReader(byt))
	return
}

// GetHttpCacheIgnoresCookies gets the http.cache_ignores_cookies node value
func GetHttpCacheIgnoresCookies(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.cache_ignores_cookies")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCacheIgnoresCookies updates the http.cache_ignores_cookies node value
func UpdateHttpCacheIgnoresCookies(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.cache_ignores_cookies", bytes.NewReader(byt))
	return
}

// GetHttpCachessl gets the http.cachessl node value
func GetHttpCachessl(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.cachessl")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCachessl updates the http.cachessl node value
func UpdateHttpCachessl(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.cachessl", bytes.NewReader(byt))
	return
}

// GetHttpCaching gets the http.caching node value
func GetHttpCaching(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.caching")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCaching updates the http.caching node value
func UpdateHttpCaching(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.caching", bytes.NewReader(byt))
	return
}

// GetHttpCertcache gets the http.certcache node value
func GetHttpCertcache(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.certcache")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCertcache updates the http.certcache node value
func UpdateHttpCertcache(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.certcache", bytes.NewReader(byt))
	return
}

// GetHttpCertstore gets the http.certstore node value
func GetHttpCertstore(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.certstore")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCertstore updates the http.certstore node value
func UpdateHttpCertstore(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.certstore", bytes.NewReader(byt))
	return
}

// GetHttpCffOverrideUsers gets the http.cff_override_users node value
func GetHttpCffOverrideUsers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.cff_override_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCffOverrideUsers updates the http.cff_override_users node value
func UpdateHttpCffOverrideUsers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.cff_override_users", bytes.NewReader(byt))
	return
}

// GetHttpClientTimeout gets the http.client_timeout node value
func GetHttpClientTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.client_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpClientTimeout updates the http.client_timeout node value
func UpdateHttpClientTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.client_timeout", bytes.NewReader(byt))
	return
}

// GetHttpConfLockWorkaround gets the http.conf_lock_workaround node value
func GetHttpConfLockWorkaround(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.conf_lock_workaround")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpConfLockWorkaround updates the http.conf_lock_workaround node value
func UpdateHttpConfLockWorkaround(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.conf_lock_workaround", bytes.NewReader(byt))
	return
}

// GetHttpConnectTimeout gets the http.connect_timeout node value
func GetHttpConnectTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.connect_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpConnectTimeout updates the http.connect_timeout node value
func UpdateHttpConnectTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.connect_timeout", bytes.NewReader(byt))
	return
}

// GetHttpConnectV6Timeout gets the http.connect_v6_timeout node value
func GetHttpConnectV6Timeout(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/http.connect_v6_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpConnectV6Timeout updates the http.connect_v6_timeout node value
func UpdateHttpConnectV6Timeout(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.connect_v6_timeout", bytes.NewReader(byt))
	return
}

// GetHttpConnlimit gets the http.connlimit node value
func GetHttpConnlimit(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.connlimit")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpConnlimit updates the http.connlimit node value
func UpdateHttpConnlimit(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.connlimit", bytes.NewReader(byt))
	return
}

// GetHttpCtypeInspectBody gets the http.ctype_inspect_body node value
func GetHttpCtypeInspectBody(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.ctype_inspect_body")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCtypeInspectBody updates the http.ctype_inspect_body node value
func UpdateHttpCtypeInspectBody(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.ctype_inspect_body", bytes.NewReader(byt))
	return
}

// GetHttpCtypeUnpackArchive gets the http.ctype_unpack_archive node value
func GetHttpCtypeUnpackArchive(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.ctype_unpack_archive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCtypeUnpackArchive updates the http.ctype_unpack_archive node value
func UpdateHttpCtypeUnpackArchive(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.ctype_unpack_archive", bytes.NewReader(byt))
	return
}

// GetHttpDebug gets the http.debug node value
func GetHttpDebug(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDebug updates the http.debug node value
func UpdateHttpDebug(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.debug", bytes.NewReader(byt))
	return
}

// GetHttpDefaultblockaction gets the http.defaultblockaction node value
func GetHttpDefaultblockaction(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.defaultblockaction")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDefaultblockaction updates the http.defaultblockaction node value
func UpdateHttpDefaultblockaction(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.defaultblockaction", bytes.NewReader(byt))
	return
}

// GetHttpDeferagents gets the http.deferagents node value
func GetHttpDeferagents(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.deferagents")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDeferagents updates the http.deferagents node value
func UpdateHttpDeferagents(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.deferagents", bytes.NewReader(byt))
	return
}

// GetHttpDeferlength gets the http.deferlength node value
func GetHttpDeferlength(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.deferlength")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDeferlength updates the http.deferlength node value
func UpdateHttpDeferlength(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.deferlength", bytes.NewReader(byt))
	return
}

// GetHttpDisplayHttpBlockpageExplicitMode gets the http.display_http_blockpage_explicit_mode node value
func GetHttpDisplayHttpBlockpageExplicitMode(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.display_http_blockpage_explicit_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDisplayHttpBlockpageExplicitMode updates the http.display_http_blockpage_explicit_mode node value
func UpdateHttpDisplayHttpBlockpageExplicitMode(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.display_http_blockpage_explicit_mode", bytes.NewReader(byt))
	return
}

// GetHttpDisplayIntro gets the http.display_intro node value
func GetHttpDisplayIntro(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.display_intro")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDisplayIntro updates the http.display_intro node value
func UpdateHttpDisplayIntro(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.display_intro", bytes.NewReader(byt))
	return
}

// GetHttpDownloadManagerDefaultCharset gets the http.download_manager_default_charset node value
func GetHttpDownloadManagerDefaultCharset(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.download_manager_default_charset")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDownloadManagerDefaultCharset updates the http.download_manager_default_charset node value
func UpdateHttpDownloadManagerDefaultCharset(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.download_manager_default_charset", bytes.NewReader(byt))
	return
}

// GetHttpEdirDelayBasicAuth gets the http.edir_delay_basic_auth node value
func GetHttpEdirDelayBasicAuth(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.edir_delay_basic_auth")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpEdirDelayBasicAuth updates the http.edir_delay_basic_auth node value
func UpdateHttpEdirDelayBasicAuth(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.edir_delay_basic_auth", bytes.NewReader(byt))
	return
}

// GetHttpEnableOutInterface gets the http.enable_out_interface node value
func GetHttpEnableOutInterface(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.enable_out_interface")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpEnableOutInterface updates the http.enable_out_interface node value
func UpdateHttpEnableOutInterface(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.enable_out_interface", bytes.NewReader(byt))
	return
}

// GetHttpEppQuotaAction gets the http.epp_quota_action node value
func GetHttpEppQuotaAction(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.epp_quota_action")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpEppQuotaAction updates the http.epp_quota_action node value
func UpdateHttpEppQuotaAction(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.epp_quota_action", bytes.NewReader(byt))
	return
}

// GetHttpExceptions gets the http.exceptions node value
func GetHttpExceptions(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpExceptions updates the http.exceptions node value
func UpdateHttpExceptions(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.exceptions", bytes.NewReader(byt))
	return
}

// GetHttpForcedCachingExtension gets the http.forced_caching_extension node value
func GetHttpForcedCachingExtension(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.forced_caching_extension")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpForcedCachingExtension updates the http.forced_caching_extension node value
func UpdateHttpForcedCachingExtension(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.forced_caching_extension", bytes.NewReader(byt))
	return
}

// GetHttpForcedCachingNeverCachePrefix gets the http.forced_caching_never_cache_prefix node value
func GetHttpForcedCachingNeverCachePrefix(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.forced_caching_never_cache_prefix")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpForcedCachingNeverCachePrefix updates the http.forced_caching_never_cache_prefix node value
func UpdateHttpForcedCachingNeverCachePrefix(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.forced_caching_never_cache_prefix", bytes.NewReader(byt))
	return
}

// GetHttpForcedCachingStatus gets the http.forced_caching_status node value
func GetHttpForcedCachingStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.forced_caching_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpForcedCachingStatus updates the http.forced_caching_status node value
func UpdateHttpForcedCachingStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.forced_caching_status", bytes.NewReader(byt))
	return
}

// GetHttpForcedCachingTtl gets the http.forced_caching_ttl node value
func GetHttpForcedCachingTtl(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.forced_caching_ttl")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpForcedCachingTtl updates the http.forced_caching_ttl node value
func UpdateHttpForcedCachingTtl(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.forced_caching_ttl", bytes.NewReader(byt))
	return
}

// GetHttpForcedCachingUserAgentPrefix gets the http.forced_caching_user_agent_prefix node value
func GetHttpForcedCachingUserAgentPrefix(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.forced_caching_user_agent_prefix")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpForcedCachingUserAgentPrefix updates the http.forced_caching_user_agent_prefix node value
func UpdateHttpForcedCachingUserAgentPrefix(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.forced_caching_user_agent_prefix", bytes.NewReader(byt))
	return
}

// GetHttpHttpLoopbackDetect gets the http.http_loopback_detect node value
func GetHttpHttpLoopbackDetect(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.http_loopback_detect")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpHttpLoopbackDetect updates the http.http_loopback_detect node value
func UpdateHttpHttpLoopbackDetect(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.http_loopback_detect", bytes.NewReader(byt))
	return
}

// GetHttpIeSslBlockpageWorkaround gets the http.ie_ssl_blockpage_workaround node value
func GetHttpIeSslBlockpageWorkaround(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.ie_ssl_blockpage_workaround")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpIeSslBlockpageWorkaround updates the http.ie_ssl_blockpage_workaround node value
func UpdateHttpIeSslBlockpageWorkaround(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.ie_ssl_blockpage_workaround", bytes.NewReader(byt))
	return
}

// GetHttpLimitAdSsoInterfaces gets the http.limit_ad_sso_interfaces node value
func GetHttpLimitAdSsoInterfaces(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.limit_ad_sso_interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpLimitAdSsoInterfaces updates the http.limit_ad_sso_interfaces node value
func UpdateHttpLimitAdSsoInterfaces(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.limit_ad_sso_interfaces", bytes.NewReader(byt))
	return
}

// GetHttpLocalSiteList gets the http.local_site_list node value
func GetHttpLocalSiteList(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.local_site_list")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpLocalSiteList updates the http.local_site_list node value
func UpdateHttpLocalSiteList(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.local_site_list", bytes.NewReader(byt))
	return
}

// GetHttpMaxContentEncoding gets the http.max_content_encoding node value
func GetHttpMaxContentEncoding(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.max_content_encoding")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpMaxContentEncoding updates the http.max_content_encoding node value
func UpdateHttpMaxContentEncoding(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.max_content_encoding", bytes.NewReader(byt))
	return
}

// GetHttpMaxTempfileSize gets the http.max_tempfile_size node value
func GetHttpMaxTempfileSize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.max_tempfile_size")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpMaxTempfileSize updates the http.max_tempfile_size node value
func UpdateHttpMaxTempfileSize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.max_tempfile_size", bytes.NewReader(byt))
	return
}

// GetHttpMaxthreads gets the http.maxthreads node value
func GetHttpMaxthreads(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.maxthreads")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpMaxthreads updates the http.maxthreads node value
func UpdateHttpMaxthreads(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.maxthreads", bytes.NewReader(byt))
	return
}

// GetHttpMaxthreadsUnused gets the http.maxthreads_unused node value
func GetHttpMaxthreadsUnused(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.maxthreads_unused")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpMaxthreadsUnused updates the http.maxthreads_unused node value
func UpdateHttpMaxthreadsUnused(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.maxthreads_unused", bytes.NewReader(byt))
	return
}

// GetHttpModulepath gets the http.modulepath node value
func GetHttpModulepath(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.modulepath")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpModulepath updates the http.modulepath node value
func UpdateHttpModulepath(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.modulepath", bytes.NewReader(byt))
	return
}

// GetHttpModules gets the http.modules node value
func GetHttpModules(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.modules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpModules updates the http.modules node value
func UpdateHttpModules(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.modules", bytes.NewReader(byt))
	return
}

// GetHttpNoscancontent gets the http.noscancontent node value
func GetHttpNoscancontent(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.noscancontent")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpNoscancontent updates the http.noscancontent node value
func UpdateHttpNoscancontent(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.noscancontent", bytes.NewReader(byt))
	return
}

// GetHttpOpendirectoryKeytab gets the http.opendirectory_keytab node value
func GetHttpOpendirectoryKeytab(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.opendirectory_keytab")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpOpendirectoryKeytab updates the http.opendirectory_keytab node value
func UpdateHttpOpendirectoryKeytab(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.opendirectory_keytab", bytes.NewReader(byt))
	return
}

// GetHttpPacFile gets the http.pac_file node value
func GetHttpPacFile(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.pac_file")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPacFile updates the http.pac_file node value
func UpdateHttpPacFile(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.pac_file", bytes.NewReader(byt))
	return
}

// GetHttpParentProxyHost gets the http.parent_proxy_host node value
func GetHttpParentProxyHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.parent_proxy_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpParentProxyHost updates the http.parent_proxy_host node value
func UpdateHttpParentProxyHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.parent_proxy_host", bytes.NewReader(byt))
	return
}

// GetHttpParentProxyPort gets the http.parent_proxy_port node value
func GetHttpParentProxyPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.parent_proxy_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpParentProxyPort updates the http.parent_proxy_port node value
func UpdateHttpParentProxyPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.parent_proxy_port", bytes.NewReader(byt))
	return
}

// GetHttpParentProxyStatus gets the http.parent_proxy_status node value
func GetHttpParentProxyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.parent_proxy_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpParentProxyStatus updates the http.parent_proxy_status node value
func UpdateHttpParentProxyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.parent_proxy_status", bytes.NewReader(byt))
	return
}

// GetHttpPassthroughId gets the http.passthrough_id node value
func GetHttpPassthroughId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.passthrough_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPassthroughId updates the http.passthrough_id node value
func UpdateHttpPassthroughId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.passthrough_id", bytes.NewReader(byt))
	return
}

// GetHttpPharmingProtection gets the http.pharming_protection node value
func GetHttpPharmingProtection(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.pharming_protection")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPharmingProtection updates the http.pharming_protection node value
func UpdateHttpPharmingProtection(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.pharming_protection", bytes.NewReader(byt))
	return
}

// GetHttpPort gets the http.port node value
func GetHttpPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPort updates the http.port node value
func UpdateHttpPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.port", bytes.NewReader(byt))
	return
}

// GetHttpPortalCert gets the http.portal_cert node value
func GetHttpPortalCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.portal_cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPortalCert updates the http.portal_cert node value
func UpdateHttpPortalCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.portal_cert", bytes.NewReader(byt))
	return
}

// GetHttpPortalCertChain gets the http.portal_cert_chain node value
func GetHttpPortalCertChain(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.portal_cert_chain")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPortalCertChain updates the http.portal_cert_chain node value
func UpdateHttpPortalCertChain(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.portal_cert_chain", bytes.NewReader(byt))
	return
}

// GetHttpPortalDomain gets the http.portal_domain node value
func GetHttpPortalDomain(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.portal_domain")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPortalDomain updates the http.portal_domain node value
func UpdateHttpPortalDomain(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.portal_domain", bytes.NewReader(byt))
	return
}

// GetHttpPortalHosts gets the http.portal_hosts node value
func GetHttpPortalHosts(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.portal_hosts")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPortalHosts updates the http.portal_hosts node value
func UpdateHttpPortalHosts(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.portal_hosts", bytes.NewReader(byt))
	return
}

// GetHttpPortalUseCert gets the http.portal_use_cert node value
func GetHttpPortalUseCert(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.portal_use_cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPortalUseCert updates the http.portal_use_cert node value
func UpdateHttpPortalUseCert(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.portal_use_cert", bytes.NewReader(byt))
	return
}

// GetHttpProceedCacheTimeout gets the http.proceed_cache_timeout node value
func GetHttpProceedCacheTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.proceed_cache_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpProceedCacheTimeout updates the http.proceed_cache_timeout node value
func UpdateHttpProceedCacheTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.proceed_cache_timeout", bytes.NewReader(byt))
	return
}

// GetHttpProfiles gets the http.profiles node value
func GetHttpProfiles(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.profiles")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpProfiles updates the http.profiles node value
func UpdateHttpProfiles(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.profiles", bytes.NewReader(byt))
	return
}

// GetHttpQuotaSliceTime gets the http.quota_slice_time node value
func GetHttpQuotaSliceTime(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.quota_slice_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpQuotaSliceTime updates the http.quota_slice_time node value
func UpdateHttpQuotaSliceTime(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.quota_slice_time", bytes.NewReader(byt))
	return
}

// GetHttpRemoveRequest gets the http.remove_request node value
func GetHttpRemoveRequest(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.remove_request")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpRemoveRequest updates the http.remove_request node value
func UpdateHttpRemoveRequest(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.remove_request", bytes.NewReader(byt))
	return
}

// GetHttpRemoveResponse gets the http.remove_response node value
func GetHttpRemoveResponse(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.remove_response")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpRemoveResponse updates the http.remove_response node value
func UpdateHttpRemoveResponse(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.remove_response", bytes.NewReader(byt))
	return
}

// GetHttpResponseTimeout gets the http.response_timeout node value
func GetHttpResponseTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.response_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpResponseTimeout updates the http.response_timeout node value
func UpdateHttpResponseTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.response_timeout", bytes.NewReader(byt))
	return
}

// GetHttpScLocalDb gets the http.sc_local_db node value
func GetHttpScLocalDb(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.sc_local_db")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpScLocalDb updates the http.sc_local_db node value
func UpdateHttpScLocalDb(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.sc_local_db", bytes.NewReader(byt))
	return
}

// GetHttpScanEppTraffic gets the http.scan_epp_traffic node value
func GetHttpScanEppTraffic(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.scan_epp_traffic")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpScanEppTraffic updates the http.scan_epp_traffic node value
func UpdateHttpScanEppTraffic(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.scan_epp_traffic", bytes.NewReader(byt))
	return
}

// GetHttpSearchdomain gets the http.searchdomain node value
func GetHttpSearchdomain(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.searchdomain")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpSearchdomain updates the http.searchdomain node value
func UpdateHttpSearchdomain(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.searchdomain", bytes.NewReader(byt))
	return
}

// GetHttpStrictHttp gets the http.strict_http node value
func GetHttpStrictHttp(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.strict_http")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpStrictHttp updates the http.strict_http node value
func UpdateHttpStrictHttp(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.strict_http", bytes.NewReader(byt))
	return
}

// GetHttpTlsciphersClient gets the http.tlsciphers_client node value
func GetHttpTlsciphersClient(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.tlsciphers_client")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTlsciphersClient updates the http.tlsciphers_client node value
func UpdateHttpTlsciphersClient(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.tlsciphers_client", bytes.NewReader(byt))
	return
}

// GetHttpTlsciphersServer gets the http.tlsciphers_server node value
func GetHttpTlsciphersServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.tlsciphers_server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTlsciphersServer updates the http.tlsciphers_server node value
func UpdateHttpTlsciphersServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.tlsciphers_server", bytes.NewReader(byt))
	return
}

// GetHttpTmpfsUsageMinMemsize gets the http.tmpfs_usage_min_memsize node value
func GetHttpTmpfsUsageMinMemsize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.tmpfs_usage_min_memsize")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTmpfsUsageMinMemsize updates the http.tmpfs_usage_min_memsize node value
func UpdateHttpTmpfsUsageMinMemsize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.tmpfs_usage_min_memsize", bytes.NewReader(byt))
	return
}

// GetHttpTransparentAuthTimeout gets the http.transparent_auth_timeout node value
func GetHttpTransparentAuthTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.transparent_auth_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTransparentAuthTimeout updates the http.transparent_auth_timeout node value
func UpdateHttpTransparentAuthTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.transparent_auth_timeout", bytes.NewReader(byt))
	return
}

// GetHttpTransparentDstSkip gets the http.transparent_dst_skip node value
func GetHttpTransparentDstSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.transparent_dst_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTransparentDstSkip updates the http.transparent_dst_skip node value
func UpdateHttpTransparentDstSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.transparent_dst_skip", bytes.NewReader(byt))
	return
}

// GetHttpTransparentSkipAutoPf gets the http.transparent_skip_auto_pf node value
func GetHttpTransparentSkipAutoPf(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.transparent_skip_auto_pf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTransparentSkipAutoPf updates the http.transparent_skip_auto_pf node value
func UpdateHttpTransparentSkipAutoPf(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.transparent_skip_auto_pf", bytes.NewReader(byt))
	return
}

// GetHttpTransparentSrcSkip gets the http.transparent_src_skip node value
func GetHttpTransparentSrcSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.transparent_src_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTransparentSrcSkip updates the http.transparent_src_skip node value
func UpdateHttpTransparentSrcSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.transparent_src_skip", bytes.NewReader(byt))
	return
}

// GetHttpTunnelTimeout gets the http.tunnel_timeout node value
func GetHttpTunnelTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.tunnel_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTunnelTimeout updates the http.tunnel_timeout node value
func UpdateHttpTunnelTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.tunnel_timeout", bytes.NewReader(byt))
	return
}

// GetHttpTunnelV6Timeout gets the http.tunnel_v6_timeout node value
func GetHttpTunnelV6Timeout(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/http.tunnel_v6_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTunnelV6Timeout updates the http.tunnel_v6_timeout node value
func UpdateHttpTunnelV6Timeout(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.tunnel_v6_timeout", bytes.NewReader(byt))
	return
}

// GetHttpUndefercontent gets the http.undefercontent node value
func GetHttpUndefercontent(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.undefercontent")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUndefercontent updates the http.undefercontent node value
func UpdateHttpUndefercontent(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.undefercontent", bytes.NewReader(byt))
	return
}

// GetHttpUndeferextension gets the http.undeferextension node value
func GetHttpUndeferextension(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.undeferextension")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUndeferextension updates the http.undeferextension node value
func UpdateHttpUndeferextension(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.undeferextension", bytes.NewReader(byt))
	return
}

// GetHttpUrlFilteringRedirectUrl gets the http.url_filtering_redirect_url node value
func GetHttpUrlFilteringRedirectUrl(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.url_filtering_redirect_url")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUrlFilteringRedirectUrl updates the http.url_filtering_redirect_url node value
func UpdateHttpUrlFilteringRedirectUrl(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.url_filtering_redirect_url", bytes.NewReader(byt))
	return
}

// GetHttpUseDstaddrForGeopiplookup gets the http.use_dstaddr_for_geopiplookup node value
func GetHttpUseDstaddrForGeopiplookup(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.use_dstaddr_for_geopiplookup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUseDstaddrForGeopiplookup updates the http.use_dstaddr_for_geopiplookup node value
func UpdateHttpUseDstaddrForGeopiplookup(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.use_dstaddr_for_geopiplookup", bytes.NewReader(byt))
	return
}

// GetHttpUseKrb5Adsso gets the http.use_krb5_adsso node value
func GetHttpUseKrb5Adsso(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.use_krb5_adsso")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUseKrb5Adsso updates the http.use_krb5_adsso node value
func UpdateHttpUseKrb5Adsso(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.use_krb5_adsso", bytes.NewReader(byt))
	return
}

// GetHttpUseSni gets the http.use_sni node value
func GetHttpUseSni(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.use_sni")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUseSni updates the http.use_sni node value
func UpdateHttpUseSni(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.use_sni", bytes.NewReader(byt))
	return
}

// GetHttpUseSxlUrid gets the http.use_sxl_urid node value
func GetHttpUseSxlUrid(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.use_sxl_urid")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUseSxlUrid updates the http.use_sxl_urid node value
func UpdateHttpUseSxlUrid(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.use_sxl_urid", bytes.NewReader(byt))
	return
}

// GetIcmpForward gets the icmp.forward node value
func GetIcmpForward(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.forward")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpForward updates the icmp.forward node value
func UpdateIcmpForward(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.forward", bytes.NewReader(byt))
	return
}

// GetIcmpInput gets the icmp.input node value
func GetIcmpInput(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.input")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpInput updates the icmp.input node value
func UpdateIcmpInput(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.input", bytes.NewReader(byt))
	return
}

// GetIcmpLogRedirect gets the icmp.log_redirect node value
func GetIcmpLogRedirect(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.log_redirect")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpLogRedirect updates the icmp.log_redirect node value
func UpdateIcmpLogRedirect(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.log_redirect", bytes.NewReader(byt))
	return
}

// GetIcmpPingForward gets the icmp.ping.forward node value
func GetIcmpPingForward(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.ping.forward")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpPingForward updates the icmp.ping.forward node value
func UpdateIcmpPingForward(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.ping.forward", bytes.NewReader(byt))
	return
}

// GetIcmpPingInput gets the icmp.ping.input node value
func GetIcmpPingInput(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.ping.input")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpPingInput updates the icmp.ping.input node value
func UpdateIcmpPingInput(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.ping.input", bytes.NewReader(byt))
	return
}

// GetIcmpPingOutput gets the icmp.ping.output node value
func GetIcmpPingOutput(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.ping.output")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpPingOutput updates the icmp.ping.output node value
func UpdateIcmpPingOutput(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.ping.output", bytes.NewReader(byt))
	return
}

// GetIcmpSecure gets the icmp.secure node value
func GetIcmpSecure(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.secure")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpSecure updates the icmp.secure node value
func UpdateIcmpSecure(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.secure", bytes.NewReader(byt))
	return
}

// GetIcmpTracerouteForward gets the icmp.traceroute.forward node value
func GetIcmpTracerouteForward(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.traceroute.forward")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpTracerouteForward updates the icmp.traceroute.forward node value
func UpdateIcmpTracerouteForward(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.traceroute.forward", bytes.NewReader(byt))
	return
}

// GetIcmpTracerouteInput gets the icmp.traceroute.input node value
func GetIcmpTracerouteInput(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.traceroute.input")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpTracerouteInput updates the icmp.traceroute.input node value
func UpdateIcmpTracerouteInput(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.traceroute.input", bytes.NewReader(byt))
	return
}

// GetIdentForward gets the ident.forward node value
func GetIdentForward(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ident.forward")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIdentForward updates the ident.forward node value
func UpdateIdentForward(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ident.forward", bytes.NewReader(byt))
	return
}

// GetIdentResponse gets the ident.response node value
func GetIdentResponse(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ident.response")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIdentResponse updates the ident.response node value
func UpdateIdentResponse(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ident.response", bytes.NewReader(byt))
	return
}

// GetIdentStatus gets the ident.status node value
func GetIdentStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ident.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIdentStatus updates the ident.status node value
func UpdateIdentStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ident.status", bytes.NewReader(byt))
	return
}

// GetInterfacesAdvancedArpAnnounce gets the interfaces.advanced.arp_announce node value
func GetInterfacesAdvancedArpAnnounce(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/interfaces.advanced.arp_announce")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateInterfacesAdvancedArpAnnounce updates the interfaces.advanced.arp_announce node value
func UpdateInterfacesAdvancedArpAnnounce(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/interfaces.advanced.arp_announce", bytes.NewReader(byt))
	return
}

// GetInterfacesAdvancedArpIgnore gets the interfaces.advanced.arp_ignore node value
func GetInterfacesAdvancedArpIgnore(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/interfaces.advanced.arp_ignore")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateInterfacesAdvancedArpIgnore updates the interfaces.advanced.arp_ignore node value
func UpdateInterfacesAdvancedArpIgnore(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/interfaces.advanced.arp_ignore", bytes.NewReader(byt))
	return
}

// GetInterfacesAdvancedDefaultMetric gets the interfaces.advanced.default_metric node value
func GetInterfacesAdvancedDefaultMetric(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/interfaces.advanced.default_metric")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateInterfacesAdvancedDefaultMetric updates the interfaces.advanced.default_metric node value
func UpdateInterfacesAdvancedDefaultMetric(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/interfaces.advanced.default_metric", bytes.NewReader(byt))
	return
}

// GetInterfacesInterfaces gets the interfaces.interfaces node value
func GetInterfacesInterfaces(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/interfaces.interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateInterfacesInterfaces updates the interfaces.interfaces node value
func UpdateInterfacesInterfaces(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/interfaces.interfaces", bytes.NewReader(byt))
	return
}

// GetIpsDnsServers gets the ips.dns_servers node value
func GetIpsDnsServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.dns_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsDnsServers updates the ips.dns_servers node value
func UpdateIpsDnsServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.dns_servers", bytes.NewReader(byt))
	return
}

// GetIpsEngine gets the ips.engine node value
func GetIpsEngine(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ips.engine")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsEngine updates the ips.engine node value
func UpdateIpsEngine(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.engine", bytes.NewReader(byt))
	return
}

// GetIpsExceptions gets the ips.exceptions node value
func GetIpsExceptions(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ips.exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsExceptions updates the ips.exceptions node value
func UpdateIpsExceptions(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.exceptions", bytes.NewReader(byt))
	return
}

// GetIpsFailopen gets the ips.failopen node value
func GetIpsFailopen(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ips.failopen")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsFailopen updates the ips.failopen node value
func UpdateIpsFailopen(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.failopen", bytes.NewReader(byt))
	return
}

// GetIpsFileBasedRules gets the ips.file_based_rules node value
func GetIpsFileBasedRules(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ips.file_based_rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsFileBasedRules updates the ips.file_based_rules node value
func UpdateIpsFileBasedRules(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.file_based_rules", bytes.NewReader(byt))
	return
}

// GetIpsGroups gets the ips.groups node value
func GetIpsGroups(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ips.groups")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsGroups updates the ips.groups node value
func UpdateIpsGroups(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.groups", bytes.NewReader(byt))
	return
}

// GetIpsHttpServers gets the ips.http_servers node value
func GetIpsHttpServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.http_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsHttpServers updates the ips.http_servers node value
func UpdateIpsHttpServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.http_servers", bytes.NewReader(byt))
	return
}

// GetIpsIpsfbAlertInterval gets the ips.ipsfb.alert_interval node value
func GetIpsIpsfbAlertInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ips.ipsfb.alert_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsIpsfbAlertInterval updates the ips.ipsfb.alert_interval node value
func UpdateIpsIpsfbAlertInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.ipsfb.alert_interval", bytes.NewReader(byt))
	return
}

// GetIpsIpsfbConfigInterval gets the ips.ipsfb.config_interval node value
func GetIpsIpsfbConfigInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ips.ipsfb.config_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsIpsfbConfigInterval updates the ips.ipsfb.config_interval node value
func UpdateIpsIpsfbConfigInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.ipsfb.config_interval", bytes.NewReader(byt))
	return
}

// GetIpsIpsfbDebug gets the ips.ipsfb.debug node value
func GetIpsIpsfbDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ips.ipsfb.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsIpsfbDebug updates the ips.ipsfb.debug node value
func UpdateIpsIpsfbDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.ipsfb.debug", bytes.NewReader(byt))
	return
}

// GetIpsLocalNetworks gets the ips.local_networks node value
func GetIpsLocalNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ips.local_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsLocalNetworks updates the ips.local_networks node value
func UpdateIpsLocalNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.local_networks", bytes.NewReader(byt))
	return
}

// GetIpsNumInstances gets the ips.num_instances node value
func GetIpsNumInstances(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.num_instances")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsNumInstances updates the ips.num_instances node value
func UpdateIpsNumInstances(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.num_instances", bytes.NewReader(byt))
	return
}

// GetIpsPatternChannel gets the ips.pattern_channel node value
func GetIpsPatternChannel(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ips.pattern_channel")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsPatternChannel updates the ips.pattern_channel node value
func UpdateIpsPatternChannel(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.pattern_channel", bytes.NewReader(byt))
	return
}

// GetIpsPolicy gets the ips.policy node value
func GetIpsPolicy(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ips.policy")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsPolicy updates the ips.policy node value
func UpdateIpsPolicy(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.policy", bytes.NewReader(byt))
	return
}

// GetIpsQueueLength gets the ips.queue_length node value
func GetIpsQueueLength(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ips.queue_length")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsQueueLength updates the ips.queue_length node value
func UpdateIpsQueueLength(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.queue_length", bytes.NewReader(byt))
	return
}

// GetIpsQueueThreshold gets the ips.queue_threshold node value
func GetIpsQueueThreshold(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.queue_threshold")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsQueueThreshold updates the ips.queue_threshold node value
func UpdateIpsQueueThreshold(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.queue_threshold", bytes.NewReader(byt))
	return
}

// GetIpsReloadMethod gets the ips.reload_method node value
func GetIpsReloadMethod(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ips.reload_method")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsReloadMethod updates the ips.reload_method node value
func UpdateIpsReloadMethod(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.reload_method", bytes.NewReader(byt))
	return
}

// GetIpsRestartPolicy gets the ips.restart_policy node value
func GetIpsRestartPolicy(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ips.restart_policy")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsRestartPolicy updates the ips.restart_policy node value
func UpdateIpsRestartPolicy(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.restart_policy", bytes.NewReader(byt))
	return
}

// GetIpsRuleModifiers gets the ips.rule_modifiers node value
func GetIpsRuleModifiers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.rule_modifiers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsRuleModifiers updates the ips.rule_modifiers node value
func UpdateIpsRuleModifiers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.rule_modifiers", bytes.NewReader(byt))
	return
}

// GetIpsRules gets the ips.rules node value
func GetIpsRules(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsRules updates the ips.rules node value
func UpdateIpsRules(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.rules", bytes.NewReader(byt))
	return
}

// GetIpsSkipAcks gets the ips.skip_acks node value
func GetIpsSkipAcks(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ips.skip_acks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSkipAcks updates the ips.skip_acks node value
func UpdateIpsSkipAcks(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.skip_acks", bytes.NewReader(byt))
	return
}

// GetIpsSmtpServers gets the ips.smtp_servers node value
func GetIpsSmtpServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.smtp_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSmtpServers updates the ips.smtp_servers node value
func UpdateIpsSmtpServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.smtp_servers", bytes.NewReader(byt))
	return
}

// GetIpsSnortsettingsMaxQueuedBytes gets the ips.snortsettings.max_queued_bytes node value
func GetIpsSnortsettingsMaxQueuedBytes(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.snortsettings.max_queued_bytes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSnortsettingsMaxQueuedBytes updates the ips.snortsettings.max_queued_bytes node value
func UpdateIpsSnortsettingsMaxQueuedBytes(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.snortsettings.max_queued_bytes", bytes.NewReader(byt))
	return
}

// GetIpsSnortsettingsMaxQueuedSegs gets the ips.snortsettings.max_queued_segs node value
func GetIpsSnortsettingsMaxQueuedSegs(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.snortsettings.max_queued_segs")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSnortsettingsMaxQueuedSegs updates the ips.snortsettings.max_queued_segs node value
func UpdateIpsSnortsettingsMaxQueuedSegs(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.snortsettings.max_queued_segs", bytes.NewReader(byt))
	return
}

// GetIpsSnortsettingsMaxTcp gets the ips.snortsettings.max_tcp node value
func GetIpsSnortsettingsMaxTcp(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.snortsettings.max_tcp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSnortsettingsMaxTcp updates the ips.snortsettings.max_tcp node value
func UpdateIpsSnortsettingsMaxTcp(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.snortsettings.max_tcp", bytes.NewReader(byt))
	return
}

// GetIpsSnortsettingsMaxUdp gets the ips.snortsettings.max_udp node value
func GetIpsSnortsettingsMaxUdp(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.snortsettings.max_udp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSnortsettingsMaxUdp updates the ips.snortsettings.max_udp node value
func UpdateIpsSnortsettingsMaxUdp(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.snortsettings.max_udp", bytes.NewReader(byt))
	return
}

// GetIpsSnortsettingsMemcap gets the ips.snortsettings.memcap node value
func GetIpsSnortsettingsMemcap(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.snortsettings.memcap")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSnortsettingsMemcap updates the ips.snortsettings.memcap node value
func UpdateIpsSnortsettingsMemcap(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.snortsettings.memcap", bytes.NewReader(byt))
	return
}

// GetIpsSnortsettingsSearchMethod gets the ips.snortsettings.search_method node value
func GetIpsSnortsettingsSearchMethod(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ips.snortsettings.search_method")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSnortsettingsSearchMethod updates the ips.snortsettings.search_method node value
func UpdateIpsSnortsettingsSearchMethod(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.snortsettings.search_method", bytes.NewReader(byt))
	return
}

// GetIpsSqlServers gets the ips.sql_servers node value
func GetIpsSqlServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.sql_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSqlServers updates the ips.sql_servers node value
func UpdateIpsSqlServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.sql_servers", bytes.NewReader(byt))
	return
}

// GetIpsStatus gets the ips.status node value
func GetIpsStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ips.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsStatus updates the ips.status node value
func UpdateIpsStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.status", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedCrlAutoFetching gets the ipsec.advanced.crl_auto_fetching node value
func GetIpsecAdvancedCrlAutoFetching(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.crl_auto_fetching")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedCrlAutoFetching updates the ipsec.advanced.crl_auto_fetching node value
func UpdateIpsecAdvancedCrlAutoFetching(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.crl_auto_fetching", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedCrlStrictPolicy gets the ipsec.advanced.crl_strict_policy node value
func GetIpsecAdvancedCrlStrictPolicy(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.crl_strict_policy")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedCrlStrictPolicy updates the ipsec.advanced.crl_strict_policy node value
func UpdateIpsecAdvancedCrlStrictPolicy(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.crl_strict_policy", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedDeadPeerDetection gets the ipsec.advanced.dead_peer_detection node value
func GetIpsecAdvancedDeadPeerDetection(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.dead_peer_detection")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedDeadPeerDetection updates the ipsec.advanced.dead_peer_detection node value
func UpdateIpsecAdvancedDeadPeerDetection(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.dead_peer_detection", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedIkeDebug gets the ipsec.advanced.ike_debug node value
func GetIpsecAdvancedIkeDebug(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.ike_debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedIkeDebug updates the ipsec.advanced.ike_debug node value
func UpdateIpsecAdvancedIkeDebug(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.ike_debug", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedIkePort gets the ipsec.advanced.ike_port node value
func GetIpsecAdvancedIkePort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.ike_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedIkePort updates the ipsec.advanced.ike_port node value
func UpdateIpsecAdvancedIkePort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.ike_port", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedMetric gets the ipsec.advanced.metric node value
func GetIpsecAdvancedMetric(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.metric")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedMetric updates the ipsec.advanced.metric node value
func UpdateIpsecAdvancedMetric(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.metric", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedNatTraversal gets the ipsec.advanced.nat_traversal node value
func GetIpsecAdvancedNatTraversal(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.nat_traversal")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedNatTraversal updates the ipsec.advanced.nat_traversal node value
func UpdateIpsecAdvancedNatTraversal(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.nat_traversal", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedNatTraversalKeepalive gets the ipsec.advanced.nat_traversal_keepalive node value
func GetIpsecAdvancedNatTraversalKeepalive(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.nat_traversal_keepalive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedNatTraversalKeepalive updates the ipsec.advanced.nat_traversal_keepalive node value
func UpdateIpsecAdvancedNatTraversalKeepalive(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.nat_traversal_keepalive", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedProbePsk gets the ipsec.advanced.probe_psk node value
func GetIpsecAdvancedProbePsk(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.probe_psk")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedProbePsk updates the ipsec.advanced.probe_psk node value
func UpdateIpsecAdvancedProbePsk(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.probe_psk", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedPskVpnId gets the ipsec.advanced.psk_vpn_id node value
func GetIpsecAdvancedPskVpnId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.psk_vpn_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedPskVpnId updates the ipsec.advanced.psk_vpn_id node value
func UpdateIpsecAdvancedPskVpnId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.psk_vpn_id", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedPskVpnIdType gets the ipsec.advanced.psk_vpn_id_type node value
func GetIpsecAdvancedPskVpnIdType(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.psk_vpn_id_type")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedPskVpnIdType updates the ipsec.advanced.psk_vpn_id_type node value
func UpdateIpsecAdvancedPskVpnIdType(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.psk_vpn_id_type", bytes.NewReader(byt))
	return
}

// GetIpsecConnections gets the ipsec.connections node value
func GetIpsecConnections(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ipsec.connections")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecConnections updates the ipsec.connections node value
func UpdateIpsecConnections(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.connections", bytes.NewReader(byt))
	return
}

// GetIpsecLocalRsa gets the ipsec.local_rsa node value
func GetIpsecLocalRsa(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipsec.local_rsa")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecLocalRsa updates the ipsec.local_rsa node value
func UpdateIpsecLocalRsa(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.local_rsa", bytes.NewReader(byt))
	return
}

// GetIpsecLocalX509 gets the ipsec.local_x509 node value
func GetIpsecLocalX509(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipsec.local_x509")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecLocalX509 updates the ipsec.local_x509 node value
func UpdateIpsecLocalX509(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.local_x509", bytes.NewReader(byt))
	return
}

// GetIpsecStatus gets the ipsec.status node value
func GetIpsecStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipsec.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecStatus updates the ipsec.status node value
func UpdateIpsecStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.status", bytes.NewReader(byt))
	return
}

// GetIpv6AdvancedHopLimit gets the ipv6.advanced.hop_limit node value
func GetIpv6AdvancedHopLimit(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ipv6.advanced.hop_limit")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6AdvancedHopLimit updates the ipv6.advanced.hop_limit node value
func UpdateIpv6AdvancedHopLimit(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.advanced.hop_limit", bytes.NewReader(byt))
	return
}

// GetIpv6AdvancedMaxInterval gets the ipv6.advanced.max_interval node value
func GetIpv6AdvancedMaxInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ipv6.advanced.max_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6AdvancedMaxInterval updates the ipv6.advanced.max_interval node value
func UpdateIpv6AdvancedMaxInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.advanced.max_interval", bytes.NewReader(byt))
	return
}

// GetIpv6AdvancedMinInterval gets the ipv6.advanced.min_interval node value
func GetIpv6AdvancedMinInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ipv6.advanced.min_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6AdvancedMinInterval updates the ipv6.advanced.min_interval node value
func UpdateIpv6AdvancedMinInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.advanced.min_interval", bytes.NewReader(byt))
	return
}

// GetIpv6AdvancedPreference gets the ipv6.advanced.preference node value
func GetIpv6AdvancedPreference(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.advanced.preference")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6AdvancedPreference updates the ipv6.advanced.preference node value
func UpdateIpv6AdvancedPreference(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.advanced.preference", bytes.NewReader(byt))
	return
}

// GetIpv6AdvancedReachableTime gets the ipv6.advanced.reachable_time node value
func GetIpv6AdvancedReachableTime(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ipv6.advanced.reachable_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6AdvancedReachableTime updates the ipv6.advanced.reachable_time node value
func UpdateIpv6AdvancedReachableTime(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.advanced.reachable_time", bytes.NewReader(byt))
	return
}

// GetIpv6AdvancedRetransTime gets the ipv6.advanced.retrans_time node value
func GetIpv6AdvancedRetransTime(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ipv6.advanced.retrans_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6AdvancedRetransTime updates the ipv6.advanced.retrans_time node value
func UpdateIpv6AdvancedRetransTime(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.advanced.retrans_time", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerAuthentication gets the ipv6.broker.authentication node value
func GetIpv6BrokerAuthentication(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.authentication")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerAuthentication updates the ipv6.broker.authentication node value
func UpdateIpv6BrokerAuthentication(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.authentication", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerInterface gets the ipv6.broker.interface node value
func GetIpv6BrokerInterface(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.interface")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerInterface updates the ipv6.broker.interface node value
func UpdateIpv6BrokerInterface(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.interface", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerPassword gets the ipv6.broker.password node value
func GetIpv6BrokerPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerPassword updates the ipv6.broker.password node value
func UpdateIpv6BrokerPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.password", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerProtocol gets the ipv6.broker.protocol node value
func GetIpv6BrokerProtocol(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.protocol")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerProtocol updates the ipv6.broker.protocol node value
func UpdateIpv6BrokerProtocol(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.protocol", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerServer gets the ipv6.broker.server node value
func GetIpv6BrokerServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerServer updates the ipv6.broker.server node value
func UpdateIpv6BrokerServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.server", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerStatus gets the ipv6.broker.status node value
func GetIpv6BrokerStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerStatus updates the ipv6.broker.status node value
func UpdateIpv6BrokerStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.status", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerTunnelId gets the ipv6.broker.tunnel_id node value
func GetIpv6BrokerTunnelId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.tunnel_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerTunnelId updates the ipv6.broker.tunnel_id node value
func UpdateIpv6BrokerTunnelId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.tunnel_id", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerUsername gets the ipv6.broker.username node value
func GetIpv6BrokerUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerUsername updates the ipv6.broker.username node value
func UpdateIpv6BrokerUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.username", bytes.NewReader(byt))
	return
}

// GetIpv6Nat64Address gets the ipv6.nat64.address node value
func GetIpv6Nat64Address(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.nat64.address")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Nat64Address updates the ipv6.nat64.address node value
func UpdateIpv6Nat64Address(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.nat64.address", bytes.NewReader(byt))
	return
}

// GetIpv6Nat64Dns64V6Only gets the ipv6.nat64.dns64_v6only node value
func GetIpv6Nat64Dns64V6Only(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.nat64.dns64_v6only")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Nat64Dns64V6Only updates the ipv6.nat64.dns64_v6only node value
func UpdateIpv6Nat64Dns64V6Only(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.nat64.dns64_v6only", bytes.NewReader(byt))
	return
}

// GetIpv6Nat64Prefix gets the ipv6.nat64.prefix node value
func GetIpv6Nat64Prefix(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.nat64.prefix")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Nat64Prefix updates the ipv6.nat64.prefix node value
func UpdateIpv6Nat64Prefix(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.nat64.prefix", bytes.NewReader(byt))
	return
}

// GetIpv6Nat64Status gets the ipv6.nat64.status node value
func GetIpv6Nat64Status(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.nat64.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Nat64Status updates the ipv6.nat64.status node value
func UpdateIpv6Nat64Status(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.nat64.status", bytes.NewReader(byt))
	return
}

// GetIpv6Prefer gets the ipv6.prefer node value
func GetIpv6Prefer(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.prefer")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Prefer updates the ipv6.prefer node value
func UpdateIpv6Prefer(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.prefer", bytes.NewReader(byt))
	return
}

// GetIpv6Prefixes gets the ipv6.prefixes node value
func GetIpv6Prefixes(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ipv6.prefixes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Prefixes updates the ipv6.prefixes node value
func UpdateIpv6Prefixes(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.prefixes", bytes.NewReader(byt))
	return
}

// GetIpv6Renumbering gets the ipv6.renumbering node value
func GetIpv6Renumbering(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.renumbering")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Renumbering updates the ipv6.renumbering node value
func UpdateIpv6Renumbering(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.renumbering", bytes.NewReader(byt))
	return
}

// GetIpv6Six2FourInterface gets the ipv6.six2four.interface node value
func GetIpv6Six2FourInterface(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.six2four.interface")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Six2FourInterface updates the ipv6.six2four.interface node value
func UpdateIpv6Six2FourInterface(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.six2four.interface", bytes.NewReader(byt))
	return
}

// GetIpv6Six2FourServer gets the ipv6.six2four.server node value
func GetIpv6Six2FourServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.six2four.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Six2FourServer updates the ipv6.six2four.server node value
func UpdateIpv6Six2FourServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.six2four.server", bytes.NewReader(byt))
	return
}

// GetIpv6Six2FourStatus gets the ipv6.six2four.status node value
func GetIpv6Six2FourStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.six2four.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Six2FourStatus updates the ipv6.six2four.status node value
func UpdateIpv6Six2FourStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.six2four.status", bytes.NewReader(byt))
	return
}

// GetIpv6Status gets the ipv6.status node value
func GetIpv6Status(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Status updates the ipv6.status node value
func UpdateIpv6Status(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.status", bytes.NewReader(byt))
	return
}

// GetLicensingActiveIps gets the licensing.active_ips node value
func GetLicensingActiveIps(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/licensing.active_ips")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLicensingActiveIps updates the licensing.active_ips node value
func UpdateLicensingActiveIps(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/licensing.active_ips", bytes.NewReader(byt))
	return
}

// GetLicensingLicense gets the licensing.license node value
func GetLicensingLicense(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/licensing.license")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLicensingLicense updates the licensing.license node value
func UpdateLicensingLicense(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/licensing.license", bytes.NewReader(byt))
	return
}

// GetLicensingUserLimitExceeded gets the licensing.user_limit_exceeded node value
func GetLicensingUserLimitExceeded(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/licensing.user_limit_exceeded")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLicensingUserLimitExceeded updates the licensing.user_limit_exceeded node value
func UpdateLicensingUserLimitExceeded(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/licensing.user_limit_exceeded", bytes.NewReader(byt))
	return
}

// GetLinkAggregationGroups gets the link_aggregation.groups node value
func GetLinkAggregationGroups(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/link_aggregation.groups")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLinkAggregationGroups updates the link_aggregation.groups node value
func UpdateLinkAggregationGroups(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/link_aggregation.groups", bytes.NewReader(byt))
	return
}

// GetLoadbalanceHttpErrorCode gets the loadbalance.http_error_code node value
func GetLoadbalanceHttpErrorCode(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/loadbalance.http_error_code")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLoadbalanceHttpErrorCode updates the loadbalance.http_error_code node value
func UpdateLoadbalanceHttpErrorCode(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/loadbalance.http_error_code", bytes.NewReader(byt))
	return
}

// GetLoadbalanceRules gets the loadbalance.rules node value
func GetLoadbalanceRules(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/loadbalance.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLoadbalanceRules updates the loadbalance.rules node value
func UpdateLoadbalanceRules(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/loadbalance.rules", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalActionOne gets the logfiles.local.action_one node value
func GetLogfilesLocalActionOne(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.action_one")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalActionOne updates the logfiles.local.action_one node value
func UpdateLogfilesLocalActionOne(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.action_one", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalActionThree gets the logfiles.local.action_three node value
func GetLogfilesLocalActionThree(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.action_three")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalActionThree updates the logfiles.local.action_three node value
func UpdateLogfilesLocalActionThree(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.action_three", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalActionTwo gets the logfiles.local.action_two node value
func GetLogfilesLocalActionTwo(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.action_two")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalActionTwo updates the logfiles.local.action_two node value
func UpdateLogfilesLocalActionTwo(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.action_two", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalDeleteAfterDays gets the logfiles.local.delete_after_days node value
func GetLogfilesLocalDeleteAfterDays(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.delete_after_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalDeleteAfterDays updates the logfiles.local.delete_after_days node value
func UpdateLogfilesLocalDeleteAfterDays(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.delete_after_days", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalPercentageOne gets the logfiles.local.percentage_one node value
func GetLogfilesLocalPercentageOne(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.percentage_one")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalPercentageOne updates the logfiles.local.percentage_one node value
func UpdateLogfilesLocalPercentageOne(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.percentage_one", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalPercentageThree gets the logfiles.local.percentage_three node value
func GetLogfilesLocalPercentageThree(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.percentage_three")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalPercentageThree updates the logfiles.local.percentage_three node value
func UpdateLogfilesLocalPercentageThree(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.percentage_three", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalPercentageTwo gets the logfiles.local.percentage_two node value
func GetLogfilesLocalPercentageTwo(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.percentage_two")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalPercentageTwo updates the logfiles.local.percentage_two node value
func UpdateLogfilesLocalPercentageTwo(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.percentage_two", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalStatus gets the logfiles.local.status node value
func GetLogfilesLocalStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalStatus updates the logfiles.local.status node value
func UpdateLogfilesLocalStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.status", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteFtpService gets the logfiles.remote.ftp_service node value
func GetLogfilesRemoteFtpService(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.ftp_service")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteFtpService updates the logfiles.remote.ftp_service node value
func UpdateLogfilesRemoteFtpService(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.ftp_service", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteHost gets the logfiles.remote.host node value
func GetLogfilesRemoteHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteHost updates the logfiles.remote.host node value
func UpdateLogfilesRemoteHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.host", bytes.NewReader(byt))
	return
}

// GetLogfilesRemotePass gets the logfiles.remote.pass node value
func GetLogfilesRemotePass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.pass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemotePass updates the logfiles.remote.pass node value
func UpdateLogfilesRemotePass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.pass", bytes.NewReader(byt))
	return
}

// GetLogfilesRemotePath gets the logfiles.remote.path node value
func GetLogfilesRemotePath(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.path")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemotePath updates the logfiles.remote.path node value
func UpdateLogfilesRemotePath(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.path", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteSmbWorkgroup gets the logfiles.remote.smb_workgroup node value
func GetLogfilesRemoteSmbWorkgroup(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.smb_workgroup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteSmbWorkgroup updates the logfiles.remote.smb_workgroup node value
func UpdateLogfilesRemoteSmbWorkgroup(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.smb_workgroup", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteSmtpAddress gets the logfiles.remote.smtp_address node value
func GetLogfilesRemoteSmtpAddress(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.smtp_address")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteSmtpAddress updates the logfiles.remote.smtp_address node value
func UpdateLogfilesRemoteSmtpAddress(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.smtp_address", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteStatus gets the logfiles.remote.status node value
func GetLogfilesRemoteStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteStatus updates the logfiles.remote.status node value
func UpdateLogfilesRemoteStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.status", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteType gets the logfiles.remote.type node value
func GetLogfilesRemoteType(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.type")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteType updates the logfiles.remote.type node value
func UpdateLogfilesRemoteType(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.type", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteUser gets the logfiles.remote.user node value
func GetLogfilesRemoteUser(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.user")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteUser updates the logfiles.remote.user node value
func UpdateLogfilesRemoteUser(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.user", bytes.NewReader(byt))
	return
}

// GetMasqRules gets the masq.rules node value
func GetMasqRules(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/masq.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMasqRules updates the masq.rules node value
func UpdateMasqRules(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/masq.rules", bytes.NewReader(byt))
	return
}

// GetMigrationAccessToken gets the migration.access_token node value
func GetMigrationAccessToken(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/migration.access_token")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMigrationAccessToken updates the migration.access_token node value
func UpdateMigrationAccessToken(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/migration.access_token", bytes.NewReader(byt))
	return
}

// GetMigrationLocalOverride gets the migration.local_override node value
func GetMigrationLocalOverride(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/migration.local_override")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMigrationLocalOverride updates the migration.local_override node value
func UpdateMigrationLocalOverride(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/migration.local_override", bytes.NewReader(byt))
	return
}

// GetMigrationRefreshToken gets the migration.refresh_token node value
func GetMigrationRefreshToken(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/migration.refresh_token")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMigrationRefreshToken updates the migration.refresh_token node value
func UpdateMigrationRefreshToken(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/migration.refresh_token", bytes.NewReader(byt))
	return
}

// GetMigrationTabVisibility gets the migration.tab_visibility node value
func GetMigrationTabVisibility(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/migration.tab_visibility")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMigrationTabVisibility updates the migration.tab_visibility node value
func UpdateMigrationTabVisibility(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/migration.tab_visibility", bytes.NewReader(byt))
	return
}

// GetMigrationToolsetVersion gets the migration.toolset_version node value
func GetMigrationToolsetVersion(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/migration.toolset_version")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMigrationToolsetVersion updates the migration.toolset_version node value
func UpdateMigrationToolsetVersion(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/migration.toolset_version", bytes.NewReader(byt))
	return
}

// GetMigrationUtmVersion gets the migration.utm_version node value
func GetMigrationUtmVersion(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/migration.utm_version")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMigrationUtmVersion updates the migration.utm_version node value
func UpdateMigrationUtmVersion(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/migration.utm_version", bytes.NewReader(byt))
	return
}

// GetMobileControlCa gets the mobile_control.ca node value
func GetMobileControlCa(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.ca")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlCa updates the mobile_control.ca node value
func UpdateMobileControlCa(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.ca", bytes.NewReader(byt))
	return
}

// GetMobileControlConfigCisco gets the mobile_control.config.cisco node value
func GetMobileControlConfigCisco(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.config.cisco")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlConfigCisco updates the mobile_control.config.cisco node value
func UpdateMobileControlConfigCisco(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.config.cisco", bytes.NewReader(byt))
	return
}

// GetMobileControlConfigEapMethod gets the mobile_control.config.eap_method node value
func GetMobileControlConfigEapMethod(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.config.eap_method")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlConfigEapMethod updates the mobile_control.config.eap_method node value
func UpdateMobileControlConfigEapMethod(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.config.eap_method", bytes.NewReader(byt))
	return
}

// GetMobileControlConfigForcePush gets the mobile_control.config.force_push node value
func GetMobileControlConfigForcePush(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.config.force_push")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlConfigForcePush updates the mobile_control.config.force_push node value
func UpdateMobileControlConfigForcePush(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.config.force_push", bytes.NewReader(byt))
	return
}

// GetMobileControlConfigL2Tp gets the mobile_control.config.l2tp node value
func GetMobileControlConfigL2Tp(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.config.l2tp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlConfigL2Tp updates the mobile_control.config.l2tp node value
func UpdateMobileControlConfigL2Tp(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.config.l2tp", bytes.NewReader(byt))
	return
}

// GetMobileControlConfigWifiNetworks gets the mobile_control.config.wifi_networks node value
func GetMobileControlConfigWifiNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/mobile_control.config.wifi_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlConfigWifiNetworks updates the mobile_control.config.wifi_networks node value
func UpdateMobileControlConfigWifiNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.config.wifi_networks", bytes.NewReader(byt))
	return
}

// GetMobileControlCustomer gets the mobile_control.customer node value
func GetMobileControlCustomer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.customer")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlCustomer updates the mobile_control.customer node value
func UpdateMobileControlCustomer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.customer", bytes.NewReader(byt))
	return
}

// GetMobileControlDebug gets the mobile_control.debug node value
func GetMobileControlDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlDebug updates the mobile_control.debug node value
func UpdateMobileControlDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.debug", bytes.NewReader(byt))
	return
}

// GetMobileControlNacCisco gets the mobile_control.nac.cisco node value
func GetMobileControlNacCisco(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.cisco")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacCisco updates the mobile_control.nac.cisco node value
func UpdateMobileControlNacCisco(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.cisco", bytes.NewReader(byt))
	return
}

// GetMobileControlNacDenyAllVpn gets the mobile_control.nac.deny_all_vpn node value
func GetMobileControlNacDenyAllVpn(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.deny_all_vpn")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacDenyAllVpn updates the mobile_control.nac.deny_all_vpn node value
func UpdateMobileControlNacDenyAllVpn(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.deny_all_vpn", bytes.NewReader(byt))
	return
}

// GetMobileControlNacL2Tp gets the mobile_control.nac.l2tp node value
func GetMobileControlNacL2Tp(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.l2tp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacL2Tp updates the mobile_control.nac.l2tp node value
func UpdateMobileControlNacL2Tp(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.l2tp", bytes.NewReader(byt))
	return
}

// GetMobileControlNacMacsAllowed gets the mobile_control.nac.macs_allowed node value
func GetMobileControlNacMacsAllowed(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.macs_allowed")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacMacsAllowed updates the mobile_control.nac.macs_allowed node value
func UpdateMobileControlNacMacsAllowed(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.macs_allowed", bytes.NewReader(byt))
	return
}

// GetMobileControlNacMacsDenied gets the mobile_control.nac.macs_denied node value
func GetMobileControlNacMacsDenied(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.macs_denied")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacMacsDenied updates the mobile_control.nac.macs_denied node value
func UpdateMobileControlNacMacsDenied(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.macs_denied", bytes.NewReader(byt))
	return
}

// GetMobileControlNacPollInterval gets the mobile_control.nac.poll_interval node value
func GetMobileControlNacPollInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.poll_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacPollInterval updates the mobile_control.nac.poll_interval node value
func UpdateMobileControlNacPollInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.poll_interval", bytes.NewReader(byt))
	return
}

// GetMobileControlNacUsersDenied gets the mobile_control.nac.users_denied node value
func GetMobileControlNacUsersDenied(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.users_denied")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacUsersDenied updates the mobile_control.nac.users_denied node value
func UpdateMobileControlNacUsersDenied(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.users_denied", bytes.NewReader(byt))
	return
}

// GetMobileControlNacWifiNetworks gets the mobile_control.nac.wifi_networks node value
func GetMobileControlNacWifiNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.wifi_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacWifiNetworks updates the mobile_control.nac.wifi_networks node value
func UpdateMobileControlNacWifiNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.wifi_networks", bytes.NewReader(byt))
	return
}

// GetMobileControlPassword gets the mobile_control.password node value
func GetMobileControlPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlPassword updates the mobile_control.password node value
func UpdateMobileControlPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.password", bytes.NewReader(byt))
	return
}

// GetMobileControlServer gets the mobile_control.server node value
func GetMobileControlServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlServer updates the mobile_control.server node value
func UpdateMobileControlServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.server", bytes.NewReader(byt))
	return
}

// GetMobileControlStatus gets the mobile_control.status node value
func GetMobileControlStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlStatus updates the mobile_control.status node value
func UpdateMobileControlStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.status", bytes.NewReader(byt))
	return
}

// GetMobileControlUsername gets the mobile_control.username node value
func GetMobileControlUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlUsername updates the mobile_control.username node value
func UpdateMobileControlUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.username", bytes.NewReader(byt))
	return
}

// GetNatRules gets the nat.rules node value
func GetNatRules(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/nat.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNatRules updates the nat.rules node value
func UpdateNatRules(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/nat.rules", bytes.NewReader(byt))
	return
}

// GetNotificationsDeviceInfo gets the notifications.device_info node value
func GetNotificationsDeviceInfo(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/notifications.device_info")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsDeviceInfo updates the notifications.device_info node value
func UpdateNotificationsDeviceInfo(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.device_info", bytes.NewReader(byt))
	return
}

// GetNotificationsLimiting gets the notifications.limiting node value
func GetNotificationsLimiting(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/notifications.limiting")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsLimiting updates the notifications.limiting node value
func UpdateNotificationsLimiting(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.limiting", bytes.NewReader(byt))
	return
}

// GetNotificationsOverlay gets the notifications.overlay node value
func GetNotificationsOverlay(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/notifications.overlay")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsOverlay updates the notifications.overlay node value
func UpdateNotificationsOverlay(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.overlay", bytes.NewReader(byt))
	return
}

// GetNotificationsRebootReason gets the notifications.reboot_reason node value
func GetNotificationsRebootReason(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/notifications.reboot_reason")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsRebootReason updates the notifications.reboot_reason node value
func UpdateNotificationsRebootReason(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.reboot_reason", bytes.NewReader(byt))
	return
}

// GetNotificationsRecipients gets the notifications.recipients node value
func GetNotificationsRecipients(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/notifications.recipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsRecipients updates the notifications.recipients node value
func UpdateNotificationsRecipients(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.recipients", bytes.NewReader(byt))
	return
}

// GetNotificationsSender gets the notifications.sender node value
func GetNotificationsSender(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/notifications.sender")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSender updates the notifications.sender node value
func UpdateNotificationsSender(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.sender", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpAuthentication gets the notifications.smtp.authentication node value
func GetNotificationsSmtpAuthentication(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.authentication")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpAuthentication updates the notifications.smtp.authentication node value
func UpdateNotificationsSmtpAuthentication(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.authentication", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpPassword gets the notifications.smtp.password node value
func GetNotificationsSmtpPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpPassword updates the notifications.smtp.password node value
func UpdateNotificationsSmtpPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.password", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpPort gets the notifications.smtp.port node value
func GetNotificationsSmtpPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpPort updates the notifications.smtp.port node value
func UpdateNotificationsSmtpPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.port", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpServer gets the notifications.smtp.server node value
func GetNotificationsSmtpServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpServer updates the notifications.smtp.server node value
func UpdateNotificationsSmtpServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.server", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpStatus gets the notifications.smtp.status node value
func GetNotificationsSmtpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpStatus updates the notifications.smtp.status node value
func UpdateNotificationsSmtpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.status", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpTls gets the notifications.smtp.tls node value
func GetNotificationsSmtpTls(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.tls")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpTls updates the notifications.smtp.tls node value
func UpdateNotificationsSmtpTls(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.tls", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpUsername gets the notifications.smtp.username node value
func GetNotificationsSmtpUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpUsername updates the notifications.smtp.username node value
func UpdateNotificationsSmtpUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.username", bytes.NewReader(byt))
	return
}

// GetNtpAllowedNetworks gets the ntp.allowed_networks node value
func GetNtpAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ntp.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNtpAllowedNetworks updates the ntp.allowed_networks node value
func UpdateNtpAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ntp.allowed_networks", bytes.NewReader(byt))
	return
}

// GetNtpServers gets the ntp.servers node value
func GetNtpServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ntp.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNtpServers updates the ntp.servers node value
func UpdateNtpServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ntp.servers", bytes.NewReader(byt))
	return
}

// GetNtpStatus gets the ntp.status node value
func GetNtpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ntp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNtpStatus updates the ntp.status node value
func UpdateNtpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ntp.status", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedBlockInvalidCtPackets gets the packetfilter.advanced.block_invalid_ct_packets node value
func GetPacketfilterAdvancedBlockInvalidCtPackets(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.block_invalid_ct_packets")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedBlockInvalidCtPackets updates the packetfilter.advanced.block_invalid_ct_packets node value
func UpdatePacketfilterAdvancedBlockInvalidCtPackets(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.block_invalid_ct_packets", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedCheckPacketLength gets the packetfilter.advanced.check_packet_length node value
func GetPacketfilterAdvancedCheckPacketLength(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.check_packet_length")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedCheckPacketLength updates the packetfilter.advanced.check_packet_length node value
func UpdatePacketfilterAdvancedCheckPacketLength(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.check_packet_length", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedConntrackHelpers gets the packetfilter.advanced.conntrack_helpers node value
func GetPacketfilterAdvancedConntrackHelpers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.conntrack_helpers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedConntrackHelpers updates the packetfilter.advanced.conntrack_helpers node value
func UpdatePacketfilterAdvancedConntrackHelpers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.conntrack_helpers", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedFtpPorts gets the packetfilter.advanced.ftp_ports node value
func GetPacketfilterAdvancedFtpPorts(c sophos.ClientInterface) (val []int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.ftp_ports")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedFtpPorts updates the packetfilter.advanced.ftp_ports node value
func UpdatePacketfilterAdvancedFtpPorts(c sophos.ClientInterface, val []int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.ftp_ports", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedLogBroadcasts gets the packetfilter.advanced.log_broadcasts node value
func GetPacketfilterAdvancedLogBroadcasts(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.log_broadcasts")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedLogBroadcasts updates the packetfilter.advanced.log_broadcasts node value
func UpdatePacketfilterAdvancedLogBroadcasts(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.log_broadcasts", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedLogDnsRequests gets the packetfilter.advanced.log_dns_requests node value
func GetPacketfilterAdvancedLogDnsRequests(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.log_dns_requests")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedLogDnsRequests updates the packetfilter.advanced.log_dns_requests node value
func UpdatePacketfilterAdvancedLogDnsRequests(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.log_dns_requests", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedLogFtpData gets the packetfilter.advanced.log_ftp_data node value
func GetPacketfilterAdvancedLogFtpData(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.log_ftp_data")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedLogFtpData updates the packetfilter.advanced.log_ftp_data node value
func UpdatePacketfilterAdvancedLogFtpData(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.log_ftp_data", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedLogInvalid gets the packetfilter.advanced.log_invalid node value
func GetPacketfilterAdvancedLogInvalid(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.log_invalid")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedLogInvalid updates the packetfilter.advanced.log_invalid node value
func UpdatePacketfilterAdvancedLogInvalid(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.log_invalid", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedLogMcast gets the packetfilter.advanced.log_mcast node value
func GetPacketfilterAdvancedLogMcast(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.log_mcast")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedLogMcast updates the packetfilter.advanced.log_mcast node value
func UpdatePacketfilterAdvancedLogMcast(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.log_mcast", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedLogStrictTcpState gets the packetfilter.advanced.log_strict_tcp_state node value
func GetPacketfilterAdvancedLogStrictTcpState(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.log_strict_tcp_state")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedLogStrictTcpState updates the packetfilter.advanced.log_strict_tcp_state node value
func UpdatePacketfilterAdvancedLogStrictTcpState(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.log_strict_tcp_state", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedNoErrorReplay gets the packetfilter.advanced.no_error_replay node value
func GetPacketfilterAdvancedNoErrorReplay(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.no_error_replay")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedNoErrorReplay updates the packetfilter.advanced.no_error_replay node value
func UpdatePacketfilterAdvancedNoErrorReplay(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.no_error_replay", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedOptimizeIpset gets the packetfilter.advanced.optimize.ipset node value
func GetPacketfilterAdvancedOptimizeIpset(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.optimize.ipset")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedOptimizeIpset updates the packetfilter.advanced.optimize.ipset node value
func UpdatePacketfilterAdvancedOptimizeIpset(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.optimize.ipset", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedOptimizePorts gets the packetfilter.advanced.optimize.ports node value
func GetPacketfilterAdvancedOptimizePorts(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.optimize.ports")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedOptimizePorts updates the packetfilter.advanced.optimize.ports node value
func UpdatePacketfilterAdvancedOptimizePorts(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.optimize.ports", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedSpoofProtection gets the packetfilter.advanced.spoof_protection node value
func GetPacketfilterAdvancedSpoofProtection(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.spoof_protection")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedSpoofProtection updates the packetfilter.advanced.spoof_protection node value
func UpdatePacketfilterAdvancedSpoofProtection(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.spoof_protection", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedStrictTcpState gets the packetfilter.advanced.strict_tcp_state node value
func GetPacketfilterAdvancedStrictTcpState(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.strict_tcp_state")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedStrictTcpState updates the packetfilter.advanced.strict_tcp_state node value
func UpdatePacketfilterAdvancedStrictTcpState(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.strict_tcp_state", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedTcpMaxRetrans gets the packetfilter.advanced.tcp_max_retrans node value
func GetPacketfilterAdvancedTcpMaxRetrans(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.tcp_max_retrans")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedTcpMaxRetrans updates the packetfilter.advanced.tcp_max_retrans node value
func UpdatePacketfilterAdvancedTcpMaxRetrans(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.tcp_max_retrans", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedTcpWindowScaling gets the packetfilter.advanced.tcp_window_scaling node value
func GetPacketfilterAdvancedTcpWindowScaling(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.tcp_window_scaling")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedTcpWindowScaling updates the packetfilter.advanced.tcp_window_scaling node value
func UpdatePacketfilterAdvancedTcpWindowScaling(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.tcp_window_scaling", bytes.NewReader(byt))
	return
}

// GetPacketfilterRules gets the packetfilter.rules node value
func GetPacketfilterRules(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterRules updates the packetfilter.rules node value
func UpdatePacketfilterRules(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.rules", bytes.NewReader(byt))
	return
}

// GetPacketfilterRulesAuto gets the packetfilter.rules_auto node value
func GetPacketfilterRulesAuto(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.rules_auto")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterRulesAuto updates the packetfilter.rules_auto node value
func UpdatePacketfilterRulesAuto(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.rules_auto", bytes.NewReader(byt))
	return
}

// GetPacketfilterRulesBack gets the packetfilter.rules_back node value
func GetPacketfilterRulesBack(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.rules_back")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterRulesBack updates the packetfilter.rules_back node value
func UpdatePacketfilterRulesBack(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.rules_back", bytes.NewReader(byt))
	return
}

// GetPacketfilterRulesFront gets the packetfilter.rules_front node value
func GetPacketfilterRulesFront(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.rules_front")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterRulesFront updates the packetfilter.rules_front node value
func UpdatePacketfilterRulesFront(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.rules_front", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackGenericTimeout gets the packetfilter.timeouts.ip_conntrack_generic_timeout node value
func GetPacketfilterTimeoutsIpConntrackGenericTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_generic_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackGenericTimeout updates the packetfilter.timeouts.ip_conntrack_generic_timeout node value
func UpdatePacketfilterTimeoutsIpConntrackGenericTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_generic_timeout", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackIcmpTimeout gets the packetfilter.timeouts.ip_conntrack_icmp_timeout node value
func GetPacketfilterTimeoutsIpConntrackIcmpTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_icmp_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackIcmpTimeout updates the packetfilter.timeouts.ip_conntrack_icmp_timeout node value
func UpdatePacketfilterTimeoutsIpConntrackIcmpTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_icmp_timeout", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutClose gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_close node value
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutClose(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_close")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutClose updates the packetfilter.timeouts.ip_conntrack_tcp_timeout_close node value
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutClose(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_close", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutCloseWait gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_close_wait node value
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutCloseWait(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_close_wait")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutCloseWait updates the packetfilter.timeouts.ip_conntrack_tcp_timeout_close_wait node value
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutCloseWait(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_close_wait", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutEstablished gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_established node value
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutEstablished(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_established")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutEstablished updates the packetfilter.timeouts.ip_conntrack_tcp_timeout_established node value
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutEstablished(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_established", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutFinWait gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_fin_wait node value
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutFinWait(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_fin_wait")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutFinWait updates the packetfilter.timeouts.ip_conntrack_tcp_timeout_fin_wait node value
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutFinWait(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_fin_wait", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutLastAck gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_last_ack node value
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutLastAck(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_last_ack")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutLastAck updates the packetfilter.timeouts.ip_conntrack_tcp_timeout_last_ack node value
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutLastAck(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_last_ack", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutMaxRetrans gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_max_retrans node value
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutMaxRetrans(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_max_retrans")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutMaxRetrans updates the packetfilter.timeouts.ip_conntrack_tcp_timeout_max_retrans node value
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutMaxRetrans(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_max_retrans", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutSynRecv gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_recv node value
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutSynRecv(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_recv")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutSynRecv updates the packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_recv node value
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutSynRecv(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_recv", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutSynSent gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_sent node value
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutSynSent(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_sent")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutSynSent updates the packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_sent node value
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutSynSent(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_sent", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutTimeWait gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_time_wait node value
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutTimeWait(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_time_wait")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutTimeWait updates the packetfilter.timeouts.ip_conntrack_tcp_timeout_time_wait node value
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutTimeWait(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_time_wait", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackUdpTimeout gets the packetfilter.timeouts.ip_conntrack_udp_timeout node value
func GetPacketfilterTimeoutsIpConntrackUdpTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_udp_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackUdpTimeout updates the packetfilter.timeouts.ip_conntrack_udp_timeout node value
func UpdatePacketfilterTimeoutsIpConntrackUdpTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_udp_timeout", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackUdpTimeoutStream gets the packetfilter.timeouts.ip_conntrack_udp_timeout_stream node value
func GetPacketfilterTimeoutsIpConntrackUdpTimeoutStream(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_udp_timeout_stream")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackUdpTimeoutStream updates the packetfilter.timeouts.ip_conntrack_udp_timeout_stream node value
func UpdatePacketfilterTimeoutsIpConntrackUdpTimeoutStream(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_udp_timeout_stream", bytes.NewReader(byt))
	return
}

// GetPasswdLoginuserClearpass gets the passwd.loginuser.clearpass node value
func GetPasswdLoginuserClearpass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/passwd.loginuser.clearpass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePasswdLoginuserClearpass updates the passwd.loginuser.clearpass node value
func UpdatePasswdLoginuserClearpass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/passwd.loginuser.clearpass", bytes.NewReader(byt))
	return
}

// GetPasswdLoginuserCryptpass gets the passwd.loginuser.cryptpass node value
func GetPasswdLoginuserCryptpass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/passwd.loginuser.cryptpass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePasswdLoginuserCryptpass updates the passwd.loginuser.cryptpass node value
func UpdatePasswdLoginuserCryptpass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/passwd.loginuser.cryptpass", bytes.NewReader(byt))
	return
}

// GetPasswdRootClearpass gets the passwd.root.clearpass node value
func GetPasswdRootClearpass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/passwd.root.clearpass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePasswdRootClearpass updates the passwd.root.clearpass node value
func UpdatePasswdRootClearpass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/passwd.root.clearpass", bytes.NewReader(byt))
	return
}

// GetPasswdRootCryptpass gets the passwd.root.cryptpass node value
func GetPasswdRootCryptpass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/passwd.root.cryptpass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePasswdRootCryptpass updates the passwd.root.cryptpass node value
func UpdatePasswdRootCryptpass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/passwd.root.cryptpass", bytes.NewReader(byt))
	return
}

// GetPdfPaper gets the pdf.paper node value
func GetPdfPaper(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pdf.paper")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePdfPaper updates the pdf.paper node value
func UpdatePdfPaper(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pdf.paper", bytes.NewReader(byt))
	return
}

// GetPimSmAutoPfOut gets the pim_sm.auto_pf_out node value
func GetPimSmAutoPfOut(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pim_sm.auto_pf_out")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmAutoPfOut updates the pim_sm.auto_pf_out node value
func UpdatePimSmAutoPfOut(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.auto_pf_out", bytes.NewReader(byt))
	return
}

// GetPimSmAutoPfrule gets the pim_sm.auto_pfrule node value
func GetPimSmAutoPfrule(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pim_sm.auto_pfrule")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmAutoPfrule updates the pim_sm.auto_pfrule node value
func UpdatePimSmAutoPfrule(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.auto_pfrule", bytes.NewReader(byt))
	return
}

// GetPimSmDebug gets the pim_sm.debug node value
func GetPimSmDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pim_sm.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmDebug updates the pim_sm.debug node value
func UpdatePimSmDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.debug", bytes.NewReader(byt))
	return
}

// GetPimSmInterfaces gets the pim_sm.interfaces node value
func GetPimSmInterfaces(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pim_sm.interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmInterfaces updates the pim_sm.interfaces node value
func UpdatePimSmInterfaces(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.interfaces", bytes.NewReader(byt))
	return
}

// GetPimSmRpRouters gets the pim_sm.rp_routers node value
func GetPimSmRpRouters(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pim_sm.rp_routers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmRpRouters updates the pim_sm.rp_routers node value
func UpdatePimSmRpRouters(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.rp_routers", bytes.NewReader(byt))
	return
}

// GetPimSmSptSwitchBytes gets the pim_sm.spt_switch_bytes node value
func GetPimSmSptSwitchBytes(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/pim_sm.spt_switch_bytes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmSptSwitchBytes updates the pim_sm.spt_switch_bytes node value
func UpdatePimSmSptSwitchBytes(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.spt_switch_bytes", bytes.NewReader(byt))
	return
}

// GetPimSmSptSwitchStatus gets the pim_sm.spt_switch_status node value
func GetPimSmSptSwitchStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pim_sm.spt_switch_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmSptSwitchStatus updates the pim_sm.spt_switch_status node value
func UpdatePimSmSptSwitchStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.spt_switch_status", bytes.NewReader(byt))
	return
}

// GetPimSmStatus gets the pim_sm.status node value
func GetPimSmStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pim_sm.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmStatus updates the pim_sm.status node value
func UpdatePimSmStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.status", bytes.NewReader(byt))
	return
}

// GetPop3AllowedClients gets the pop3.allowed_clients node value
func GetPop3AllowedClients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pop3.allowed_clients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3AllowedClients updates the pop3.allowed_clients node value
func UpdatePop3AllowedClients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.allowed_clients", bytes.NewReader(byt))
	return
}

// GetPop3AllowedServers gets the pop3.allowed_servers node value
func GetPop3AllowedServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/pop3.allowed_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3AllowedServers updates the pop3.allowed_servers node value
func UpdatePop3AllowedServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.allowed_servers", bytes.NewReader(byt))
	return
}

// GetPop3CffAsMarker gets the pop3.cff_as_marker node value
func GetPop3CffAsMarker(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.cff_as_marker")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3CffAsMarker updates the pop3.cff_as_marker node value
func UpdatePop3CffAsMarker(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.cff_as_marker", bytes.NewReader(byt))
	return
}

// GetPop3CffAv gets the pop3.cff_av node value
func GetPop3CffAv(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.cff_av")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3CffAv updates the pop3.cff_av node value
func UpdatePop3CffAv(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.cff_av", bytes.NewReader(byt))
	return
}

// GetPop3CffAvAction gets the pop3.cff_av_action node value
func GetPop3CffAvAction(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.cff_av_action")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3CffAvAction updates the pop3.cff_av_action node value
func UpdatePop3CffAvAction(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.cff_av_action", bytes.NewReader(byt))
	return
}

// GetPop3CffAvEngines gets the pop3.cff_av_engines node value
func GetPop3CffAvEngines(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.cff_av_engines")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3CffAvEngines updates the pop3.cff_av_engines node value
func UpdatePop3CffAvEngines(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.cff_av_engines", bytes.NewReader(byt))
	return
}

// GetPop3CffFileExtensions gets the pop3.cff_file_extensions node value
func GetPop3CffFileExtensions(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/pop3.cff_file_extensions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3CffFileExtensions updates the pop3.cff_file_extensions node value
func UpdatePop3CffFileExtensions(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.cff_file_extensions", bytes.NewReader(byt))
	return
}

// GetPop3DirectlyDeleteQuarantined gets the pop3.directly_delete_quarantined node value
func GetPop3DirectlyDeleteQuarantined(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.directly_delete_quarantined")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3DirectlyDeleteQuarantined updates the pop3.directly_delete_quarantined node value
func UpdatePop3DirectlyDeleteQuarantined(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.directly_delete_quarantined", bytes.NewReader(byt))
	return
}

// GetPop3Exceptions gets the pop3.exceptions node value
func GetPop3Exceptions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pop3.exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3Exceptions updates the pop3.exceptions node value
func UpdatePop3Exceptions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.exceptions", bytes.NewReader(byt))
	return
}

// GetPop3KnownServers gets the pop3.known_servers node value
func GetPop3KnownServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pop3.known_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3KnownServers updates the pop3.known_servers node value
func UpdatePop3KnownServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.known_servers", bytes.NewReader(byt))
	return
}

// GetPop3MaxMessageSize gets the pop3.max_message_size node value
func GetPop3MaxMessageSize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/pop3.max_message_size")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3MaxMessageSize updates the pop3.max_message_size node value
func UpdatePop3MaxMessageSize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.max_message_size", bytes.NewReader(byt))
	return
}

// GetPop3PrefetchingInterval gets the pop3.prefetching.interval node value
func GetPop3PrefetchingInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/pop3.prefetching.interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3PrefetchingInterval updates the pop3.prefetching.interval node value
func UpdatePop3PrefetchingInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.prefetching.interval", bytes.NewReader(byt))
	return
}

// GetPop3PrefetchingOptimizeStorage gets the pop3.prefetching.optimize_storage node value
func GetPop3PrefetchingOptimizeStorage(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.prefetching.optimize_storage")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3PrefetchingOptimizeStorage updates the pop3.prefetching.optimize_storage node value
func UpdatePop3PrefetchingOptimizeStorage(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.prefetching.optimize_storage", bytes.NewReader(byt))
	return
}

// GetPop3PrefetchingStatus gets the pop3.prefetching.status node value
func GetPop3PrefetchingStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.prefetching.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3PrefetchingStatus updates the pop3.prefetching.status node value
func UpdatePop3PrefetchingStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.prefetching.status", bytes.NewReader(byt))
	return
}

// GetPop3PrefetchingStorageMinHoldDays gets the pop3.prefetching.storage_min_hold_days node value
func GetPop3PrefetchingStorageMinHoldDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/pop3.prefetching.storage_min_hold_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3PrefetchingStorageMinHoldDays updates the pop3.prefetching.storage_min_hold_days node value
func UpdatePop3PrefetchingStorageMinHoldDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.prefetching.storage_min_hold_days", bytes.NewReader(byt))
	return
}

// GetPop3QuarantineUnscannable gets the pop3.quarantine_unscannable node value
func GetPop3QuarantineUnscannable(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.quarantine_unscannable")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3QuarantineUnscannable updates the pop3.quarantine_unscannable node value
func UpdatePop3QuarantineUnscannable(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.quarantine_unscannable", bytes.NewReader(byt))
	return
}

// GetPop3SandboxMaxFilesizeMb gets the pop3.sandbox_max_filesize_mb node value
func GetPop3SandboxMaxFilesizeMb(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/pop3.sandbox_max_filesize_mb")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3SandboxMaxFilesizeMb updates the pop3.sandbox_max_filesize_mb node value
func UpdatePop3SandboxMaxFilesizeMb(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.sandbox_max_filesize_mb", bytes.NewReader(byt))
	return
}

// GetPop3SandboxScanStatus gets the pop3.sandbox_scan_status node value
func GetPop3SandboxScanStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.sandbox_scan_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3SandboxScanStatus updates the pop3.sandbox_scan_status node value
func UpdatePop3SandboxScanStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.sandbox_scan_status", bytes.NewReader(byt))
	return
}

// GetPop3ScanTls gets the pop3.scan_tls node value
func GetPop3ScanTls(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.scan_tls")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3ScanTls updates the pop3.scan_tls node value
func UpdatePop3ScanTls(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.scan_tls", bytes.NewReader(byt))
	return
}

// GetPop3SenderBlacklist gets the pop3.sender_blacklist node value
func GetPop3SenderBlacklist(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pop3.sender_blacklist")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3SenderBlacklist updates the pop3.sender_blacklist node value
func UpdatePop3SenderBlacklist(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.sender_blacklist", bytes.NewReader(byt))
	return
}

// GetPop3Spam gets the pop3.spam node value
func GetPop3Spam(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.spam")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3Spam updates the pop3.spam node value
func UpdatePop3Spam(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.spam", bytes.NewReader(byt))
	return
}

// GetPop3SpamExpressions gets the pop3.spam_expressions node value
func GetPop3SpamExpressions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pop3.spam_expressions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3SpamExpressions updates the pop3.spam_expressions node value
func UpdatePop3SpamExpressions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.spam_expressions", bytes.NewReader(byt))
	return
}

// GetPop3Spamplus gets the pop3.spamplus node value
func GetPop3Spamplus(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.spamplus")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3Spamplus updates the pop3.spamplus node value
func UpdatePop3Spamplus(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.spamplus", bytes.NewReader(byt))
	return
}

// GetPop3Spamstatus gets the pop3.spamstatus node value
func GetPop3Spamstatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.spamstatus")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3Spamstatus updates the pop3.spamstatus node value
func UpdatePop3Spamstatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.spamstatus", bytes.NewReader(byt))
	return
}

// GetPop3Status gets the pop3.status node value
func GetPop3Status(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3Status updates the pop3.status node value
func UpdatePop3Status(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.status", bytes.NewReader(byt))
	return
}

// GetPop3TlsCert gets the pop3.tls_cert node value
func GetPop3TlsCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.tls_cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3TlsCert updates the pop3.tls_cert node value
func UpdatePop3TlsCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.tls_cert", bytes.NewReader(byt))
	return
}

// GetPop3TransparentSkip gets the pop3.transparent_skip node value
func GetPop3TransparentSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pop3.transparent_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3TransparentSkip updates the pop3.transparent_skip node value
func UpdatePop3TransparentSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.transparent_skip", bytes.NewReader(byt))
	return
}

// GetPop3TransparentSkipAutoPf gets the pop3.transparent_skip_auto_pf node value
func GetPop3TransparentSkipAutoPf(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.transparent_skip_auto_pf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3TransparentSkipAutoPf updates the pop3.transparent_skip_auto_pf node value
func UpdatePop3TransparentSkipAutoPf(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.transparent_skip_auto_pf", bytes.NewReader(byt))
	return
}

// GetPop3UserCharset gets the pop3.user_charset node value
func GetPop3UserCharset(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.user_charset")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3UserCharset updates the pop3.user_charset node value
func UpdatePop3UserCharset(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.user_charset", bytes.NewReader(byt))
	return
}

// GetPortalAllowAnyUser gets the portal.allow_any_user node value
func GetPortalAllowAnyUser(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/portal.allow_any_user")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalAllowAnyUser updates the portal.allow_any_user node value
func UpdatePortalAllowAnyUser(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.allow_any_user", bytes.NewReader(byt))
	return
}

// GetPortalAllowedNetworks gets the portal.allowed_networks node value
func GetPortalAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/portal.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalAllowedNetworks updates the portal.allowed_networks node value
func UpdatePortalAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.allowed_networks", bytes.NewReader(byt))
	return
}

// GetPortalAllowedUsers gets the portal.allowed_users node value
func GetPortalAllowedUsers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/portal.allowed_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalAllowedUsers updates the portal.allowed_users node value
func UpdatePortalAllowedUsers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.allowed_users", bytes.NewReader(byt))
	return
}

// GetPortalHideItems gets the portal.hide_items node value
func GetPortalHideItems(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/portal.hide_items")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalHideItems updates the portal.hide_items node value
func UpdatePortalHideItems(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.hide_items", bytes.NewReader(byt))
	return
}

// GetPortalHostname gets the portal.hostname node value
func GetPortalHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/portal.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalHostname updates the portal.hostname node value
func UpdatePortalHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.hostname", bytes.NewReader(byt))
	return
}

// GetPortalInterfaceAddress gets the portal.interface_address node value
func GetPortalInterfaceAddress(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/portal.interface_address")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalInterfaceAddress updates the portal.interface_address node value
func UpdatePortalInterfaceAddress(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.interface_address", bytes.NewReader(byt))
	return
}

// GetPortalLanguage gets the portal.language node value
func GetPortalLanguage(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/portal.language")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalLanguage updates the portal.language node value
func UpdatePortalLanguage(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.language", bytes.NewReader(byt))
	return
}

// GetPortalPersistentCookies gets the portal.persistent_cookies node value
func GetPortalPersistentCookies(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/portal.persistent_cookies")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalPersistentCookies updates the portal.persistent_cookies node value
func UpdatePortalPersistentCookies(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.persistent_cookies", bytes.NewReader(byt))
	return
}

// GetPortalPort gets the portal.port node value
func GetPortalPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/portal.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalPort updates the portal.port node value
func UpdatePortalPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.port", bytes.NewReader(byt))
	return
}

// GetPortalStatus gets the portal.status node value
func GetPortalStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/portal.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalStatus updates the portal.status node value
func UpdatePortalStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.status", bytes.NewReader(byt))
	return
}

// GetPortalWelcomeMsg gets the portal.welcome_msg node value
func GetPortalWelcomeMsg(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/portal.welcome_msg")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalWelcomeMsg updates the portal.welcome_msg node value
func UpdatePortalWelcomeMsg(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.welcome_msg", bytes.NewReader(byt))
	return
}

// GetPsdAction gets the psd.action node value
func GetPsdAction(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/psd.action")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdAction updates the psd.action node value
func UpdatePsdAction(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.action", bytes.NewReader(byt))
	return
}

// GetPsdDelayThreshold gets the psd.delay_threshold node value
func GetPsdDelayThreshold(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/psd.delay_threshold")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdDelayThreshold updates the psd.delay_threshold node value
func UpdatePsdDelayThreshold(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.delay_threshold", bytes.NewReader(byt))
	return
}

// GetPsdHiPortsWeight gets the psd.hi_ports_weight node value
func GetPsdHiPortsWeight(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/psd.hi_ports_weight")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdHiPortsWeight updates the psd.hi_ports_weight node value
func UpdatePsdHiPortsWeight(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.hi_ports_weight", bytes.NewReader(byt))
	return
}

// GetPsdLoPortsWeight gets the psd.lo_ports_weight node value
func GetPsdLoPortsWeight(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/psd.lo_ports_weight")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdLoPortsWeight updates the psd.lo_ports_weight node value
func UpdatePsdLoPortsWeight(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.lo_ports_weight", bytes.NewReader(byt))
	return
}

// GetPsdLogLimiterBurst gets the psd.log_limiter.burst node value
func GetPsdLogLimiterBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/psd.log_limiter.burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdLogLimiterBurst updates the psd.log_limiter.burst node value
func UpdatePsdLogLimiterBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.log_limiter.burst", bytes.NewReader(byt))
	return
}

// GetPsdLogLimiterRate gets the psd.log_limiter.rate node value
func GetPsdLogLimiterRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/psd.log_limiter.rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdLogLimiterRate updates the psd.log_limiter.rate node value
func UpdatePsdLogLimiterRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.log_limiter.rate", bytes.NewReader(byt))
	return
}

// GetPsdLogLimiterStatus gets the psd.log_limiter.status node value
func GetPsdLogLimiterStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/psd.log_limiter.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdLogLimiterStatus updates the psd.log_limiter.status node value
func UpdatePsdLogLimiterStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.log_limiter.status", bytes.NewReader(byt))
	return
}

// GetPsdStatus gets the psd.status node value
func GetPsdStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/psd.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdStatus updates the psd.status node value
func UpdatePsdStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.status", bytes.NewReader(byt))
	return
}

// GetPsdWeightThreshold gets the psd.weight_threshold node value
func GetPsdWeightThreshold(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/psd.weight_threshold")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdWeightThreshold updates the psd.weight_threshold node value
func UpdatePsdWeightThreshold(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.weight_threshold", bytes.NewReader(byt))
	return
}

// GetQosAdvancedEcn gets the qos.advanced.ecn node value
func GetQosAdvancedEcn(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/qos.advanced.ecn")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateQosAdvancedEcn updates the qos.advanced.ecn node value
func UpdateQosAdvancedEcn(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/qos.advanced.ecn", bytes.NewReader(byt))
	return
}

// GetQosAdvancedKeepClassAfterEncap gets the qos.advanced.keep_class_after_encap node value
func GetQosAdvancedKeepClassAfterEncap(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/qos.advanced.keep_class_after_encap")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateQosAdvancedKeepClassAfterEncap updates the qos.advanced.keep_class_after_encap node value
func UpdateQosAdvancedKeepClassAfterEncap(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/qos.advanced.keep_class_after_encap", bytes.NewReader(byt))
	return
}

// GetQosInterfaces gets the qos.interfaces node value
func GetQosInterfaces(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/qos.interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateQosInterfaces updates the qos.interfaces node value
func UpdateQosInterfaces(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/qos.interfaces", bytes.NewReader(byt))
	return
}

// GetQuarantineKeepDbLogDays gets the quarantine.keep_db_log_days node value
func GetQuarantineKeepDbLogDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/quarantine.keep_db_log_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateQuarantineKeepDbLogDays updates the quarantine.keep_db_log_days node value
func UpdateQuarantineKeepDbLogDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/quarantine.keep_db_log_days", bytes.NewReader(byt))
	return
}

// GetQuarantineKeepQuarantineDays gets the quarantine.keep_quarantine_days node value
func GetQuarantineKeepQuarantineDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/quarantine.keep_quarantine_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateQuarantineKeepQuarantineDays updates the quarantine.keep_quarantine_days node value
func UpdateQuarantineKeepQuarantineDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/quarantine.keep_quarantine_days", bytes.NewReader(byt))
	return
}

// GetRedActivateProvFw gets the red.activate_prov_fw node value
func GetRedActivateProvFw(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/red.activate_prov_fw")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedActivateProvFw updates the red.activate_prov_fw node value
func UpdateRedActivateProvFw(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.activate_prov_fw", bytes.NewReader(byt))
	return
}

// GetRedAuthorization gets the red.authorization node value
func GetRedAuthorization(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/red.authorization")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedAuthorization updates the red.authorization node value
func UpdateRedAuthorization(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.authorization", bytes.NewReader(byt))
	return
}

// GetRedCaSettingsCity gets the red.ca_settings.city node value
func GetRedCaSettingsCity(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.ca_settings.city")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedCaSettingsCity updates the red.ca_settings.city node value
func UpdateRedCaSettingsCity(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.ca_settings.city", bytes.NewReader(byt))
	return
}

// GetRedCaSettingsCountry gets the red.ca_settings.country node value
func GetRedCaSettingsCountry(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.ca_settings.country")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedCaSettingsCountry updates the red.ca_settings.country node value
func UpdateRedCaSettingsCountry(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.ca_settings.country", bytes.NewReader(byt))
	return
}

// GetRedCaSettingsEmail gets the red.ca_settings.email node value
func GetRedCaSettingsEmail(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.ca_settings.email")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedCaSettingsEmail updates the red.ca_settings.email node value
func UpdateRedCaSettingsEmail(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.ca_settings.email", bytes.NewReader(byt))
	return
}

// GetRedCaSettingsOrganization gets the red.ca_settings.organization node value
func GetRedCaSettingsOrganization(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.ca_settings.organization")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedCaSettingsOrganization updates the red.ca_settings.organization node value
func UpdateRedCaSettingsOrganization(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.ca_settings.organization", bytes.NewReader(byt))
	return
}

// GetRedClients gets the red.clients node value
func GetRedClients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/red.clients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedClients updates the red.clients node value
func UpdateRedClients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.clients", bytes.NewReader(byt))
	return
}

// GetRedDeauthTimeout gets the red.deauth_timeout node value
func GetRedDeauthTimeout(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.deauth_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedDeauthTimeout updates the red.deauth_timeout node value
func UpdateRedDeauthTimeout(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.deauth_timeout", bytes.NewReader(byt))
	return
}

// GetRedOverlayFwEnabled gets the red.overlay_fw_enabled node value
func GetRedOverlayFwEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/red.overlay_fw_enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedOverlayFwEnabled updates the red.overlay_fw_enabled node value
func UpdateRedOverlayFwEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.overlay_fw_enabled", bytes.NewReader(byt))
	return
}

// GetRedRegistryCert gets the red.registry_cert node value
func GetRedRegistryCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.registry_cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedRegistryCert updates the red.registry_cert node value
func UpdateRedRegistryCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.registry_cert", bytes.NewReader(byt))
	return
}

// GetRedRegistryId gets the red.registry_id node value
func GetRedRegistryId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.registry_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedRegistryId updates the red.registry_id node value
func UpdateRedRegistryId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.registry_id", bytes.NewReader(byt))
	return
}

// GetRedRegistryKey gets the red.registry_key node value
func GetRedRegistryKey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.registry_key")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedRegistryKey updates the red.registry_key node value
func UpdateRedRegistryKey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.registry_key", bytes.NewReader(byt))
	return
}

// GetRedServerCert gets the red.server_cert node value
func GetRedServerCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.server_cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedServerCert updates the red.server_cert node value
func UpdateRedServerCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.server_cert", bytes.NewReader(byt))
	return
}

// GetRedServers gets the red.servers node value
func GetRedServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/red.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedServers updates the red.servers node value
func UpdateRedServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.servers", bytes.NewReader(byt))
	return
}

// GetRedStatus gets the red.status node value
func GetRedStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/red.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedStatus updates the red.status node value
func UpdateRedStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.status", bytes.NewReader(byt))
	return
}

// GetRedTls12Only gets the red.tls_1_2_only node value
func GetRedTls12Only(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/red.tls_1_2_only")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedTls12Only updates the red.tls_1_2_only node value
func UpdateRedTls12Only(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.tls_1_2_only", bytes.NewReader(byt))
	return
}

// GetRemoteAccessAdvancedMsdns1 gets the remote_access.advanced.msdns1 node value
func GetRemoteAccessAdvancedMsdns1(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.advanced.msdns1")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessAdvancedMsdns1 updates the remote_access.advanced.msdns1 node value
func UpdateRemoteAccessAdvancedMsdns1(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.advanced.msdns1", bytes.NewReader(byt))
	return
}

// GetRemoteAccessAdvancedMsdns2 gets the remote_access.advanced.msdns2 node value
func GetRemoteAccessAdvancedMsdns2(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.advanced.msdns2")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessAdvancedMsdns2 updates the remote_access.advanced.msdns2 node value
func UpdateRemoteAccessAdvancedMsdns2(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.advanced.msdns2", bytes.NewReader(byt))
	return
}

// GetRemoteAccessAdvancedMsdomain gets the remote_access.advanced.msdomain node value
func GetRemoteAccessAdvancedMsdomain(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.advanced.msdomain")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessAdvancedMsdomain updates the remote_access.advanced.msdomain node value
func UpdateRemoteAccessAdvancedMsdomain(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.advanced.msdomain", bytes.NewReader(byt))
	return
}

// GetRemoteAccessAdvancedMswins1 gets the remote_access.advanced.mswins1 node value
func GetRemoteAccessAdvancedMswins1(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.advanced.mswins1")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessAdvancedMswins1 updates the remote_access.advanced.mswins1 node value
func UpdateRemoteAccessAdvancedMswins1(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.advanced.mswins1", bytes.NewReader(byt))
	return
}

// GetRemoteAccessAdvancedMswins2 gets the remote_access.advanced.mswins2 node value
func GetRemoteAccessAdvancedMswins2(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.advanced.mswins2")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessAdvancedMswins2 updates the remote_access.advanced.mswins2 node value
func UpdateRemoteAccessAdvancedMswins2(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.advanced.mswins2", bytes.NewReader(byt))
	return
}

// GetRemoteAccessCisco gets the remote_access.cisco node value
func GetRemoteAccessCisco(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.cisco")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessCisco updates the remote_access.cisco node value
func UpdateRemoteAccessCisco(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.cisco", bytes.NewReader(byt))
	return
}

// GetRemoteAccessClientlessVpnDebug gets the remote_access.clientless_vpn.debug node value
func GetRemoteAccessClientlessVpnDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/remote_access.clientless_vpn.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessClientlessVpnDebug updates the remote_access.clientless_vpn.debug node value
func UpdateRemoteAccessClientlessVpnDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.clientless_vpn.debug", bytes.NewReader(byt))
	return
}

// GetRemoteAccessClientlessVpnStatus gets the remote_access.clientless_vpn.status node value
func GetRemoteAccessClientlessVpnStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/remote_access.clientless_vpn.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessClientlessVpnStatus updates the remote_access.clientless_vpn.status node value
func UpdateRemoteAccessClientlessVpnStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.clientless_vpn.status", bytes.NewReader(byt))
	return
}

// GetRemoteAccessL2Tp gets the remote_access.l2tp node value
func GetRemoteAccessL2Tp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.l2tp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessL2Tp updates the remote_access.l2tp node value
func UpdateRemoteAccessL2Tp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.l2tp", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpAaa gets the remote_access.pptp.aaa node value
func GetRemoteAccessPptpAaa(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.aaa")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpAaa updates the remote_access.pptp.aaa node value
func UpdateRemoteAccessPptpAaa(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.aaa", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpAuthentication gets the remote_access.pptp.authentication node value
func GetRemoteAccessPptpAuthentication(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.authentication")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpAuthentication updates the remote_access.pptp.authentication node value
func UpdateRemoteAccessPptpAuthentication(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.authentication", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpDebug gets the remote_access.pptp.debug node value
func GetRemoteAccessPptpDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpDebug updates the remote_access.pptp.debug node value
func UpdateRemoteAccessPptpDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.debug", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpEncryption gets the remote_access.pptp.encryption node value
func GetRemoteAccessPptpEncryption(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.encryption")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpEncryption updates the remote_access.pptp.encryption node value
func UpdateRemoteAccessPptpEncryption(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.encryption", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIpAssignmentDhcp gets the remote_access.pptp.ip_assignment_dhcp node value
func GetRemoteAccessPptpIpAssignmentDhcp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.ip_assignment_dhcp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIpAssignmentDhcp updates the remote_access.pptp.ip_assignment_dhcp node value
func UpdateRemoteAccessPptpIpAssignmentDhcp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.ip_assignment_dhcp", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIpAssignmentDhcpInterface gets the remote_access.pptp.ip_assignment_dhcp_interface node value
func GetRemoteAccessPptpIpAssignmentDhcpInterface(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.ip_assignment_dhcp_interface")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIpAssignmentDhcpInterface updates the remote_access.pptp.ip_assignment_dhcp_interface node value
func UpdateRemoteAccessPptpIpAssignmentDhcpInterface(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.ip_assignment_dhcp_interface", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIpAssignmentMode gets the remote_access.pptp.ip_assignment_mode node value
func GetRemoteAccessPptpIpAssignmentMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.ip_assignment_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIpAssignmentMode updates the remote_access.pptp.ip_assignment_mode node value
func UpdateRemoteAccessPptpIpAssignmentMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.ip_assignment_mode", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIpAssignmentPool gets the remote_access.pptp.ip_assignment_pool node value
func GetRemoteAccessPptpIpAssignmentPool(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.ip_assignment_pool")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIpAssignmentPool updates the remote_access.pptp.ip_assignment_pool node value
func UpdateRemoteAccessPptpIpAssignmentPool(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.ip_assignment_pool", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIphoneConnectionName gets the remote_access.pptp.iphone_connection_name node value
func GetRemoteAccessPptpIphoneConnectionName(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.iphone_connection_name")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIphoneConnectionName updates the remote_access.pptp.iphone_connection_name node value
func UpdateRemoteAccessPptpIphoneConnectionName(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.iphone_connection_name", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIphoneHostname gets the remote_access.pptp.iphone_hostname node value
func GetRemoteAccessPptpIphoneHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.iphone_hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIphoneHostname updates the remote_access.pptp.iphone_hostname node value
func UpdateRemoteAccessPptpIphoneHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.iphone_hostname", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIphoneStatus gets the remote_access.pptp.iphone_status node value
func GetRemoteAccessPptpIphoneStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.iphone_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIphoneStatus updates the remote_access.pptp.iphone_status node value
func UpdateRemoteAccessPptpIphoneStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.iphone_status", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpMtu gets the remote_access.pptp.mtu node value
func GetRemoteAccessPptpMtu(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.mtu")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpMtu updates the remote_access.pptp.mtu node value
func UpdateRemoteAccessPptpMtu(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.mtu", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpStatus gets the remote_access.pptp.status node value
func GetRemoteAccessPptpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpStatus updates the remote_access.pptp.status node value
func UpdateRemoteAccessPptpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.status", bytes.NewReader(byt))
	return
}

// GetRemoteSyslogBuffer gets the remote_syslog.buffer node value
func GetRemoteSyslogBuffer(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/remote_syslog.buffer")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteSyslogBuffer updates the remote_syslog.buffer node value
func UpdateRemoteSyslogBuffer(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_syslog.buffer", bytes.NewReader(byt))
	return
}

// GetRemoteSyslogLogs gets the remote_syslog.logs node value
func GetRemoteSyslogLogs(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/remote_syslog.logs")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteSyslogLogs updates the remote_syslog.logs node value
func UpdateRemoteSyslogLogs(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_syslog.logs", bytes.NewReader(byt))
	return
}

// GetRemoteSyslogStatus gets the remote_syslog.status node value
func GetRemoteSyslogStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/remote_syslog.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteSyslogStatus updates the remote_syslog.status node value
func UpdateRemoteSyslogStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_syslog.status", bytes.NewReader(byt))
	return
}

// GetRemoteSyslogTarget gets the remote_syslog.target node value
func GetRemoteSyslogTarget(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/remote_syslog.target")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteSyslogTarget updates the remote_syslog.target node value
func UpdateRemoteSyslogTarget(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_syslog.target", bytes.NewReader(byt))
	return
}

// GetReportingAccountingKeepdays gets the reporting.accounting_keepdays node value
func GetReportingAccountingKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.accounting_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAccountingKeepdays updates the reporting.accounting_keepdays node value
func UpdateReportingAccountingKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.accounting_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingAccountingStatus gets the reporting.accounting_status node value
func GetReportingAccountingStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.accounting_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAccountingStatus updates the reporting.accounting_status node value
func UpdateReportingAccountingStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.accounting_status", bytes.NewReader(byt))
	return
}

// GetReportingAnonymizing gets the reporting.anonymizing node value
func GetReportingAnonymizing(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.anonymizing")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAnonymizing updates the reporting.anonymizing node value
func UpdateReportingAnonymizing(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.anonymizing", bytes.NewReader(byt))
	return
}

// GetReportingAppctrlKeepdays gets the reporting.appctrl_keepdays node value
func GetReportingAppctrlKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.appctrl_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAppctrlKeepdays updates the reporting.appctrl_keepdays node value
func UpdateReportingAppctrlKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.appctrl_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingAppctrlStatus gets the reporting.appctrl_status node value
func GetReportingAppctrlStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.appctrl_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAppctrlStatus updates the reporting.appctrl_status node value
func UpdateReportingAppctrlStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.appctrl_status", bytes.NewReader(byt))
	return
}

// GetReportingAtpKeepdays gets the reporting.atp_keepdays node value
func GetReportingAtpKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.atp_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAtpKeepdays updates the reporting.atp_keepdays node value
func UpdateReportingAtpKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.atp_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingAtpReset gets the reporting.atp_reset node value
func GetReportingAtpReset(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.atp_reset")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAtpReset updates the reporting.atp_reset node value
func UpdateReportingAtpReset(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.atp_reset", bytes.NewReader(byt))
	return
}

// GetReportingAtpStatus gets the reporting.atp_status node value
func GetReportingAtpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.atp_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAtpStatus updates the reporting.atp_status node value
func UpdateReportingAtpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.atp_status", bytes.NewReader(byt))
	return
}

// GetReportingAuthenticationKeepdays gets the reporting.authentication_keepdays node value
func GetReportingAuthenticationKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.authentication_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAuthenticationKeepdays updates the reporting.authentication_keepdays node value
func UpdateReportingAuthenticationKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.authentication_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingAuthenticationStatus gets the reporting.authentication_status node value
func GetReportingAuthenticationStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.authentication_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAuthenticationStatus updates the reporting.authentication_status node value
func UpdateReportingAuthenticationStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.authentication_status", bytes.NewReader(byt))
	return
}

// GetReportingCsvSeparator gets the reporting.csv_separator node value
func GetReportingCsvSeparator(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reporting.csv_separator")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingCsvSeparator updates the reporting.csv_separator node value
func UpdateReportingCsvSeparator(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.csv_separator", bytes.NewReader(byt))
	return
}

// GetReportingEmailsecurityImport gets the reporting.emailsecurity_import node value
func GetReportingEmailsecurityImport(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.emailsecurity_import")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingEmailsecurityImport updates the reporting.emailsecurity_import node value
func UpdateReportingEmailsecurityImport(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.emailsecurity_import", bytes.NewReader(byt))
	return
}

// GetReportingEmailsecurityKeepdays gets the reporting.emailsecurity_keepdays node value
func GetReportingEmailsecurityKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.emailsecurity_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingEmailsecurityKeepdays updates the reporting.emailsecurity_keepdays node value
func UpdateReportingEmailsecurityKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.emailsecurity_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingEmailsecurityStatus gets the reporting.emailsecurity_status node value
func GetReportingEmailsecurityStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.emailsecurity_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingEmailsecurityStatus updates the reporting.emailsecurity_status node value
func UpdateReportingEmailsecurityStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.emailsecurity_status", bytes.NewReader(byt))
	return
}

// GetReportingEnableVpnAccounting gets the reporting.enable_vpn_accounting node value
func GetReportingEnableVpnAccounting(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.enable_vpn_accounting")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingEnableVpnAccounting updates the reporting.enable_vpn_accounting node value
func UpdateReportingEnableVpnAccounting(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.enable_vpn_accounting", bytes.NewReader(byt))
	return
}

// GetReportingHideAccountingips gets the reporting.hide_accountingips node value
func GetReportingHideAccountingips(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.hide_accountingips")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingHideAccountingips updates the reporting.hide_accountingips node value
func UpdateReportingHideAccountingips(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.hide_accountingips", bytes.NewReader(byt))
	return
}

// GetReportingHideMailaddresses gets the reporting.hide_mailaddresses node value
func GetReportingHideMailaddresses(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.hide_mailaddresses")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingHideMailaddresses updates the reporting.hide_mailaddresses node value
func UpdateReportingHideMailaddresses(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.hide_mailaddresses", bytes.NewReader(byt))
	return
}

// GetReportingHideMaildomains gets the reporting.hide_maildomains node value
func GetReportingHideMaildomains(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.hide_maildomains")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingHideMaildomains updates the reporting.hide_maildomains node value
func UpdateReportingHideMaildomains(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.hide_maildomains", bytes.NewReader(byt))
	return
}

// GetReportingHideNetsecips gets the reporting.hide_netsecips node value
func GetReportingHideNetsecips(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.hide_netsecips")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingHideNetsecips updates the reporting.hide_netsecips node value
func UpdateReportingHideNetsecips(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.hide_netsecips", bytes.NewReader(byt))
	return
}

// GetReportingHideWebdomains gets the reporting.hide_webdomains node value
func GetReportingHideWebdomains(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.hide_webdomains")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingHideWebdomains updates the reporting.hide_webdomains node value
func UpdateReportingHideWebdomains(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.hide_webdomains", bytes.NewReader(byt))
	return
}

// GetReportingIpsImport gets the reporting.ips_import node value
func GetReportingIpsImport(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.ips_import")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingIpsImport updates the reporting.ips_import node value
func UpdateReportingIpsImport(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.ips_import", bytes.NewReader(byt))
	return
}

// GetReportingIpsKeepdays gets the reporting.ips_keepdays node value
func GetReportingIpsKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.ips_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingIpsKeepdays updates the reporting.ips_keepdays node value
func UpdateReportingIpsKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.ips_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingIpsStatus gets the reporting.ips_status node value
func GetReportingIpsStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.ips_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingIpsStatus updates the reporting.ips_status node value
func UpdateReportingIpsStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.ips_status", bytes.NewReader(byt))
	return
}

// GetReportingPacketfilterImport gets the reporting.packetfilter_import node value
func GetReportingPacketfilterImport(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.packetfilter_import")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingPacketfilterImport updates the reporting.packetfilter_import node value
func UpdateReportingPacketfilterImport(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.packetfilter_import", bytes.NewReader(byt))
	return
}

// GetReportingPacketfilterKeepdays gets the reporting.packetfilter_keepdays node value
func GetReportingPacketfilterKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.packetfilter_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingPacketfilterKeepdays updates the reporting.packetfilter_keepdays node value
func UpdateReportingPacketfilterKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.packetfilter_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingPacketfilterStatus gets the reporting.packetfilter_status node value
func GetReportingPacketfilterStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.packetfilter_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingPacketfilterStatus updates the reporting.packetfilter_status node value
func UpdateReportingPacketfilterStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.packetfilter_status", bytes.NewReader(byt))
	return
}

// GetReportingPassword1 gets the reporting.password1 node value
func GetReportingPassword1(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reporting.password1")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingPassword1 updates the reporting.password1 node value
func UpdateReportingPassword1(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.password1", bytes.NewReader(byt))
	return
}

// GetReportingPassword2 gets the reporting.password2 node value
func GetReportingPassword2(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reporting.password2")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingPassword2 updates the reporting.password2 node value
func UpdateReportingPassword2(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.password2", bytes.NewReader(byt))
	return
}

// GetReportingSandboxKeepdays gets the reporting.sandbox_keepdays node value
func GetReportingSandboxKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.sandbox_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingSandboxKeepdays updates the reporting.sandbox_keepdays node value
func UpdateReportingSandboxKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.sandbox_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingUserlogFromLogs gets the reporting.userlog_from_logs node value
func GetReportingUserlogFromLogs(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.userlog_from_logs")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingUserlogFromLogs updates the reporting.userlog_from_logs node value
func UpdateReportingUserlogFromLogs(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.userlog_from_logs", bytes.NewReader(byt))
	return
}

// GetReportingVpnKeepdays gets the reporting.vpn_keepdays node value
func GetReportingVpnKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.vpn_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingVpnKeepdays updates the reporting.vpn_keepdays node value
func UpdateReportingVpnKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.vpn_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingVpnStatus gets the reporting.vpn_status node value
func GetReportingVpnStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.vpn_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingVpnStatus updates the reporting.vpn_status node value
func UpdateReportingVpnStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.vpn_status", bytes.NewReader(byt))
	return
}

// GetReportingWafKeepdays gets the reporting.waf_keepdays node value
func GetReportingWafKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.waf_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingWafKeepdays updates the reporting.waf_keepdays node value
func UpdateReportingWafKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.waf_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingWafStatus gets the reporting.waf_status node value
func GetReportingWafStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.waf_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingWafStatus updates the reporting.waf_status node value
func UpdateReportingWafStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.waf_status", bytes.NewReader(byt))
	return
}

// GetReportingWebsecurityDetail gets the reporting.websecurity_detail node value
func GetReportingWebsecurityDetail(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.websecurity_detail")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingWebsecurityDetail updates the reporting.websecurity_detail node value
func UpdateReportingWebsecurityDetail(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.websecurity_detail", bytes.NewReader(byt))
	return
}

// GetReportingWebsecurityImport gets the reporting.websecurity_import node value
func GetReportingWebsecurityImport(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.websecurity_import")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingWebsecurityImport updates the reporting.websecurity_import node value
func UpdateReportingWebsecurityImport(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.websecurity_import", bytes.NewReader(byt))
	return
}

// GetReportingWebsecurityKeepdays gets the reporting.websecurity_keepdays node value
func GetReportingWebsecurityKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.websecurity_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingWebsecurityKeepdays updates the reporting.websecurity_keepdays node value
func UpdateReportingWebsecurityKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.websecurity_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingWebsecurityStatus gets the reporting.websecurity_status node value
func GetReportingWebsecurityStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.websecurity_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingWebsecurityStatus updates the reporting.websecurity_status node value
func UpdateReportingWebsecurityStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.websecurity_status", bytes.NewReader(byt))
	return
}

// GetReverseProxyAuaRefreshEnabled gets the reverse_proxy.aua_refresh_enabled node value
func GetReverseProxyAuaRefreshEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.aua_refresh_enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyAuaRefreshEnabled updates the reverse_proxy.aua_refresh_enabled node value
func UpdateReverseProxyAuaRefreshEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.aua_refresh_enabled", bytes.NewReader(byt))
	return
}

// GetReverseProxyAuaRefreshInterval gets the reverse_proxy.aua_refresh_interval node value
func GetReverseProxyAuaRefreshInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.aua_refresh_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyAuaRefreshInterval updates the reverse_proxy.aua_refresh_interval node value
func UpdateReverseProxyAuaRefreshInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.aua_refresh_interval", bytes.NewReader(byt))
	return
}

// GetReverseProxyBlacklistDnsrblZones gets the reverse_proxy.blacklist.dnsrbl_zones node value
func GetReverseProxyBlacklistDnsrblZones(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.blacklist.dnsrbl_zones")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyBlacklistDnsrblZones updates the reverse_proxy.blacklist.dnsrbl_zones node value
func UpdateReverseProxyBlacklistDnsrblZones(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.blacklist.dnsrbl_zones", bytes.NewReader(byt))
	return
}

// GetReverseProxyBlacklistGeoipCodes gets the reverse_proxy.blacklist.geoip_codes node value
func GetReverseProxyBlacklistGeoipCodes(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.blacklist.geoip_codes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyBlacklistGeoipCodes updates the reverse_proxy.blacklist.geoip_codes node value
func UpdateReverseProxyBlacklistGeoipCodes(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.blacklist.geoip_codes", bytes.NewReader(byt))
	return
}

// GetReverseProxyCookiesignkey gets the reverse_proxy.cookiesignkey node value
func GetReverseProxyCookiesignkey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.cookiesignkey")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyCookiesignkey updates the reverse_proxy.cookiesignkey node value
func UpdateReverseProxyCookiesignkey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.cookiesignkey", bytes.NewReader(byt))
	return
}

// GetReverseProxyCssdHostname gets the reverse_proxy.cssd_hostname node value
func GetReverseProxyCssdHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.cssd_hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyCssdHostname updates the reverse_proxy.cssd_hostname node value
func UpdateReverseProxyCssdHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.cssd_hostname", bytes.NewReader(byt))
	return
}

// GetReverseProxyCssdPort gets the reverse_proxy.cssd_port node value
func GetReverseProxyCssdPort(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.cssd_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyCssdPort updates the reverse_proxy.cssd_port node value
func UpdateReverseProxyCssdPort(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.cssd_port", bytes.NewReader(byt))
	return
}

// GetReverseProxyCustomThreatFiltersEnabled gets the reverse_proxy.custom_threat_filters_enabled node value
func GetReverseProxyCustomThreatFiltersEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.custom_threat_filters_enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyCustomThreatFiltersEnabled updates the reverse_proxy.custom_threat_filters_enabled node value
func UpdateReverseProxyCustomThreatFiltersEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.custom_threat_filters_enabled", bytes.NewReader(byt))
	return
}

// GetReverseProxyFormhardeningSecret gets the reverse_proxy.formhardening_secret node value
func GetReverseProxyFormhardeningSecret(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.formhardening_secret")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyFormhardeningSecret updates the reverse_proxy.formhardening_secret node value
func UpdateReverseProxyFormhardeningSecret(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.formhardening_secret", bytes.NewReader(byt))
	return
}

// GetReverseProxyKeepalive gets the reverse_proxy.keepalive node value
func GetReverseProxyKeepalive(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.keepalive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyKeepalive updates the reverse_proxy.keepalive node value
func UpdateReverseProxyKeepalive(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.keepalive", bytes.NewReader(byt))
	return
}

// GetReverseProxyManualmode gets the reverse_proxy.manualmode node value
func GetReverseProxyManualmode(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.manualmode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyManualmode updates the reverse_proxy.manualmode node value
func UpdateReverseProxyManualmode(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.manualmode", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxConnectionsPerChild gets the reverse_proxy.max_connections_per_child node value
func GetReverseProxyMaxConnectionsPerChild(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_connections_per_child")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxConnectionsPerChild updates the reverse_proxy.max_connections_per_child node value
func UpdateReverseProxyMaxConnectionsPerChild(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_connections_per_child", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxPreforkProcesses gets the reverse_proxy.max_prefork_processes node value
func GetReverseProxyMaxPreforkProcesses(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_prefork_processes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxPreforkProcesses updates the reverse_proxy.max_prefork_processes node value
func UpdateReverseProxyMaxPreforkProcesses(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_prefork_processes", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxProcesses gets the reverse_proxy.max_processes node value
func GetReverseProxyMaxProcesses(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_processes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxProcesses updates the reverse_proxy.max_processes node value
func UpdateReverseProxyMaxProcesses(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_processes", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxSessionFiles gets the reverse_proxy.max_session_files node value
func GetReverseProxyMaxSessionFiles(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_session_files")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxSessionFiles updates the reverse_proxy.max_session_files node value
func UpdateReverseProxyMaxSessionFiles(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_session_files", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxSpareProcesses gets the reverse_proxy.max_spare_processes node value
func GetReverseProxyMaxSpareProcesses(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_spare_processes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxSpareProcesses updates the reverse_proxy.max_spare_processes node value
func UpdateReverseProxyMaxSpareProcesses(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_spare_processes", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxSpareThreads gets the reverse_proxy.max_spare_threads node value
func GetReverseProxyMaxSpareThreads(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_spare_threads")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxSpareThreads updates the reverse_proxy.max_spare_threads node value
func UpdateReverseProxyMaxSpareThreads(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_spare_threads", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxThreadsPerProcess gets the reverse_proxy.max_threads_per_process node value
func GetReverseProxyMaxThreadsPerProcess(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_threads_per_process")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxThreadsPerProcess updates the reverse_proxy.max_threads_per_process node value
func UpdateReverseProxyMaxThreadsPerProcess(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_threads_per_process", bytes.NewReader(byt))
	return
}

// GetReverseProxyMinSpareProcesses gets the reverse_proxy.min_spare_processes node value
func GetReverseProxyMinSpareProcesses(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.min_spare_processes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMinSpareProcesses updates the reverse_proxy.min_spare_processes node value
func UpdateReverseProxyMinSpareProcesses(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.min_spare_processes", bytes.NewReader(byt))
	return
}

// GetReverseProxyMinSpareThreads gets the reverse_proxy.min_spare_threads node value
func GetReverseProxyMinSpareThreads(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.min_spare_threads")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMinSpareThreads updates the reverse_proxy.min_spare_threads node value
func UpdateReverseProxyMinSpareThreads(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.min_spare_threads", bytes.NewReader(byt))
	return
}

// GetReverseProxyMinTls gets the reverse_proxy.min_tls node value
func GetReverseProxyMinTls(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.min_tls")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMinTls updates the reverse_proxy.min_tls node value
func UpdateReverseProxyMinTls(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.min_tls", bytes.NewReader(byt))
	return
}

// GetReverseProxyModsecurityBeta gets the reverse_proxy.modsecurity_beta node value
func GetReverseProxyModsecurityBeta(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.modsecurity_beta")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyModsecurityBeta updates the reverse_proxy.modsecurity_beta node value
func UpdateReverseProxyModsecurityBeta(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.modsecurity_beta", bytes.NewReader(byt))
	return
}

// GetReverseProxyMpmMode gets the reverse_proxy.mpm_mode node value
func GetReverseProxyMpmMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.mpm_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMpmMode updates the reverse_proxy.mpm_mode node value
func UpdateReverseProxyMpmMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.mpm_mode", bytes.NewReader(byt))
	return
}

// GetReverseProxyPatternversion gets the reverse_proxy.patternversion node value
func GetReverseProxyPatternversion(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.patternversion")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyPatternversion updates the reverse_proxy.patternversion node value
func UpdateReverseProxyPatternversion(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.patternversion", bytes.NewReader(byt))
	return
}

// GetReverseProxyPort gets the reverse_proxy.port node value
func GetReverseProxyPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyPort updates the reverse_proxy.port node value
func UpdateReverseProxyPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.port", bytes.NewReader(byt))
	return
}

// GetReverseProxyProxyprotocol gets the reverse_proxy.proxyprotocol node value
func GetReverseProxyProxyprotocol(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.proxyprotocol")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyProxyprotocol updates the reverse_proxy.proxyprotocol node value
func UpdateReverseProxyProxyprotocol(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.proxyprotocol", bytes.NewReader(byt))
	return
}

// GetReverseProxyRequestLineLimit gets the reverse_proxy.request_line_limit node value
func GetReverseProxyRequestLineLimit(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.request_line_limit")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyRequestLineLimit updates the reverse_proxy.request_line_limit node value
func UpdateReverseProxyRequestLineLimit(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.request_line_limit", bytes.NewReader(byt))
	return
}

// GetReverseProxySlowhttpExceptions gets the reverse_proxy.slowhttp_exceptions node value
func GetReverseProxySlowhttpExceptions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.slowhttp_exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxySlowhttpExceptions updates the reverse_proxy.slowhttp_exceptions node value
func UpdateReverseProxySlowhttpExceptions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.slowhttp_exceptions", bytes.NewReader(byt))
	return
}

// GetReverseProxySlowhttpRequestHeaderTimeoutBase gets the reverse_proxy.slowhttp_request_header_timeout_base node value
func GetReverseProxySlowhttpRequestHeaderTimeoutBase(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_base")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxySlowhttpRequestHeaderTimeoutBase updates the reverse_proxy.slowhttp_request_header_timeout_base node value
func UpdateReverseProxySlowhttpRequestHeaderTimeoutBase(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_base", bytes.NewReader(byt))
	return
}

// GetReverseProxySlowhttpRequestHeaderTimeoutEnabled gets the reverse_proxy.slowhttp_request_header_timeout_enabled node value
func GetReverseProxySlowhttpRequestHeaderTimeoutEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxySlowhttpRequestHeaderTimeoutEnabled updates the reverse_proxy.slowhttp_request_header_timeout_enabled node value
func UpdateReverseProxySlowhttpRequestHeaderTimeoutEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_enabled", bytes.NewReader(byt))
	return
}

// GetReverseProxySlowhttpRequestHeaderTimeoutMax gets the reverse_proxy.slowhttp_request_header_timeout_max node value
func GetReverseProxySlowhttpRequestHeaderTimeoutMax(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_max")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxySlowhttpRequestHeaderTimeoutMax updates the reverse_proxy.slowhttp_request_header_timeout_max node value
func UpdateReverseProxySlowhttpRequestHeaderTimeoutMax(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_max", bytes.NewReader(byt))
	return
}

// GetReverseProxySlowhttpRequestHeaderTimeoutRate gets the reverse_proxy.slowhttp_request_header_timeout_rate node value
func GetReverseProxySlowhttpRequestHeaderTimeoutRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxySlowhttpRequestHeaderTimeoutRate updates the reverse_proxy.slowhttp_request_header_timeout_rate node value
func UpdateReverseProxySlowhttpRequestHeaderTimeoutRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_rate", bytes.NewReader(byt))
	return
}

// GetReverseProxyStatus gets the reverse_proxy.status node value
func GetReverseProxyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyStatus updates the reverse_proxy.status node value
func UpdateReverseProxyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.status", bytes.NewReader(byt))
	return
}

// GetReverseProxyTraceEnabled gets the reverse_proxy.trace_enabled node value
func GetReverseProxyTraceEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.trace_enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyTraceEnabled updates the reverse_proxy.trace_enabled node value
func UpdateReverseProxyTraceEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.trace_enabled", bytes.NewReader(byt))
	return
}

// GetReverseProxyUrlhardeningsignkey gets the reverse_proxy.urlhardeningsignkey node value
func GetReverseProxyUrlhardeningsignkey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.urlhardeningsignkey")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyUrlhardeningsignkey updates the reverse_proxy.urlhardeningsignkey node value
func UpdateReverseProxyUrlhardeningsignkey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.urlhardeningsignkey", bytes.NewReader(byt))
	return
}

// GetReverseProxyWhatkilledus gets the reverse_proxy.whatkilledus node value
func GetReverseProxyWhatkilledus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.whatkilledus")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyWhatkilledus updates the reverse_proxy.whatkilledus node value
func UpdateReverseProxyWhatkilledus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.whatkilledus", bytes.NewReader(byt))
	return
}

// GetRoutesPolicy gets the routes.policy node value
func GetRoutesPolicy(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/routes.policy")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutesPolicy updates the routes.policy node value
func UpdateRoutesPolicy(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routes.policy", bytes.NewReader(byt))
	return
}

// GetRoutesStatic gets the routes.static node value
func GetRoutesStatic(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/routes.static")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutesStatic updates the routes.static node value
func UpdateRoutesStatic(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routes.static", bytes.NewReader(byt))
	return
}

// GetRoutingBgpMaximumPaths gets the routing.bgp.maximum_paths node value
func GetRoutingBgpMaximumPaths(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/routing.bgp.maximum_paths")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingBgpMaximumPaths updates the routing.bgp.maximum_paths node value
func UpdateRoutingBgpMaximumPaths(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.bgp.maximum_paths", bytes.NewReader(byt))
	return
}

// GetRoutingBgpMaximumPathsIbgp gets the routing.bgp.maximum_paths_ibgp node value
func GetRoutingBgpMaximumPathsIbgp(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/routing.bgp.maximum_paths_ibgp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingBgpMaximumPathsIbgp updates the routing.bgp.maximum_paths_ibgp node value
func UpdateRoutingBgpMaximumPathsIbgp(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.bgp.maximum_paths_ibgp", bytes.NewReader(byt))
	return
}

// GetRoutingBgpMultipleAs gets the routing.bgp.multiple_as node value
func GetRoutingBgpMultipleAs(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.bgp.multiple_as")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingBgpMultipleAs updates the routing.bgp.multiple_as node value
func UpdateRoutingBgpMultipleAs(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.bgp.multiple_as", bytes.NewReader(byt))
	return
}

// GetRoutingBgpStatus gets the routing.bgp.status node value
func GetRoutingBgpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.bgp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingBgpStatus updates the routing.bgp.status node value
func UpdateRoutingBgpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.bgp.status", bytes.NewReader(byt))
	return
}

// GetRoutingBgpStrictMatch gets the routing.bgp.strict_match node value
func GetRoutingBgpStrictMatch(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.bgp.strict_match")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingBgpStrictMatch updates the routing.bgp.strict_match node value
func UpdateRoutingBgpStrictMatch(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.bgp.strict_match", bytes.NewReader(byt))
	return
}

// GetRoutingBgpSystems gets the routing.bgp.systems node value
func GetRoutingBgpSystems(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/routing.bgp.systems")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingBgpSystems updates the routing.bgp.systems node value
func UpdateRoutingBgpSystems(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.bgp.systems", bytes.NewReader(byt))
	return
}

// GetRoutingOspfAbrType gets the routing.ospf.abr_type node value
func GetRoutingOspfAbrType(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.abr_type")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfAbrType updates the routing.ospf.abr_type node value
func UpdateRoutingOspfAbrType(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.abr_type", bytes.NewReader(byt))
	return
}

// GetRoutingOspfAreas gets the routing.ospf.areas node value
func GetRoutingOspfAreas(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.areas")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfAreas updates the routing.ospf.areas node value
func UpdateRoutingOspfAreas(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.areas", bytes.NewReader(byt))
	return
}

// GetRoutingOspfDefaultInformation gets the routing.ospf.default_information node value
func GetRoutingOspfDefaultInformation(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.default_information")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfDefaultInformation updates the routing.ospf.default_information node value
func UpdateRoutingOspfDefaultInformation(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.default_information", bytes.NewReader(byt))
	return
}

// GetRoutingOspfDefaultRouteMetric gets the routing.ospf.default_route_metric node value
func GetRoutingOspfDefaultRouteMetric(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.default_route_metric")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfDefaultRouteMetric updates the routing.ospf.default_route_metric node value
func UpdateRoutingOspfDefaultRouteMetric(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.default_route_metric", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeBgpMetric gets the routing.ospf.redistribute.bgp.metric node value
func GetRoutingOspfRedistributeBgpMetric(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.bgp.metric")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeBgpMetric updates the routing.ospf.redistribute.bgp.metric node value
func UpdateRoutingOspfRedistributeBgpMetric(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.bgp.metric", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeBgpStatus gets the routing.ospf.redistribute.bgp.status node value
func GetRoutingOspfRedistributeBgpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.bgp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeBgpStatus updates the routing.ospf.redistribute.bgp.status node value
func UpdateRoutingOspfRedistributeBgpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.bgp.status", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeConnectedMetric gets the routing.ospf.redistribute.connected.metric node value
func GetRoutingOspfRedistributeConnectedMetric(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.connected.metric")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeConnectedMetric updates the routing.ospf.redistribute.connected.metric node value
func UpdateRoutingOspfRedistributeConnectedMetric(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.connected.metric", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeConnectedStatus gets the routing.ospf.redistribute.connected.status node value
func GetRoutingOspfRedistributeConnectedStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.connected.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeConnectedStatus updates the routing.ospf.redistribute.connected.status node value
func UpdateRoutingOspfRedistributeConnectedStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.connected.status", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeIpsecStatus gets the routing.ospf.redistribute.ipsec.status node value
func GetRoutingOspfRedistributeIpsecStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.ipsec.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeIpsecStatus updates the routing.ospf.redistribute.ipsec.status node value
func UpdateRoutingOspfRedistributeIpsecStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.ipsec.status", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeSslVpnStatus gets the routing.ospf.redistribute.ssl_vpn.status node value
func GetRoutingOspfRedistributeSslVpnStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.ssl_vpn.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeSslVpnStatus updates the routing.ospf.redistribute.ssl_vpn.status node value
func UpdateRoutingOspfRedistributeSslVpnStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.ssl_vpn.status", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeStaticMetric gets the routing.ospf.redistribute.static.metric node value
func GetRoutingOspfRedistributeStaticMetric(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.static.metric")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeStaticMetric updates the routing.ospf.redistribute.static.metric node value
func UpdateRoutingOspfRedistributeStaticMetric(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.static.metric", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeStaticStatus gets the routing.ospf.redistribute.static.status node value
func GetRoutingOspfRedistributeStaticStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.static.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeStaticStatus updates the routing.ospf.redistribute.static.status node value
func UpdateRoutingOspfRedistributeStaticStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.static.status", bytes.NewReader(byt))
	return
}

// GetRoutingOspfReferenceBandwidth gets the routing.ospf.reference_bandwidth node value
func GetRoutingOspfReferenceBandwidth(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.reference_bandwidth")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfReferenceBandwidth updates the routing.ospf.reference_bandwidth node value
func UpdateRoutingOspfReferenceBandwidth(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.reference_bandwidth", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRouterId gets the routing.ospf.router_id node value
func GetRoutingOspfRouterId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.router_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRouterId updates the routing.ospf.router_id node value
func UpdateRoutingOspfRouterId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.router_id", bytes.NewReader(byt))
	return
}

// GetRoutingOspfStatus gets the routing.ospf.status node value
func GetRoutingOspfStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfStatus updates the routing.ospf.status node value
func UpdateRoutingOspfStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.status", bytes.NewReader(byt))
	return
}

// GetRoutingQuaggaAllowedNetworks gets the routing.quagga.allowed_networks node value
func GetRoutingQuaggaAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/routing.quagga.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingQuaggaAllowedNetworks updates the routing.quagga.allowed_networks node value
func UpdateRoutingQuaggaAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.quagga.allowed_networks", bytes.NewReader(byt))
	return
}

// GetRoutingQuaggaEnablePassword gets the routing.quagga.enable_password node value
func GetRoutingQuaggaEnablePassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/routing.quagga.enable_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingQuaggaEnablePassword updates the routing.quagga.enable_password node value
func UpdateRoutingQuaggaEnablePassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.quagga.enable_password", bytes.NewReader(byt))
	return
}

// GetRoutingQuaggaLinkDetect gets the routing.quagga.link_detect node value
func GetRoutingQuaggaLinkDetect(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.quagga.link_detect")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingQuaggaLinkDetect updates the routing.quagga.link_detect node value
func UpdateRoutingQuaggaLinkDetect(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.quagga.link_detect", bytes.NewReader(byt))
	return
}

// GetRoutingQuaggaPassword gets the routing.quagga.password node value
func GetRoutingQuaggaPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/routing.quagga.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingQuaggaPassword updates the routing.quagga.password node value
func UpdateRoutingQuaggaPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.quagga.password", bytes.NewReader(byt))
	return
}

// GetSandboxReportdDebug gets the sandbox_reportd.debug node value
func GetSandboxReportdDebug(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandbox_reportd.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxReportdDebug updates the sandbox_reportd.debug node value
func UpdateSandboxReportdDebug(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandbox_reportd.debug", bytes.NewReader(byt))
	return
}

// GetSandboxReportdPollInterval gets the sandbox_reportd.poll_interval node value
func GetSandboxReportdPollInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandbox_reportd.poll_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxReportdPollInterval updates the sandbox_reportd.poll_interval node value
func UpdateSandboxReportdPollInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandbox_reportd.poll_interval", bytes.NewReader(byt))
	return
}

// GetSandboxReportdRequestTimeout gets the sandbox_reportd.request_timeout node value
func GetSandboxReportdRequestTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandbox_reportd.request_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxReportdRequestTimeout updates the sandbox_reportd.request_timeout node value
func UpdateSandboxReportdRequestTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandbox_reportd.request_timeout", bytes.NewReader(byt))
	return
}

// GetSandboxReportdRetryTimes gets the sandbox_reportd.retry_times node value
func GetSandboxReportdRetryTimes(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandbox_reportd.retry_times")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxReportdRetryTimes updates the sandbox_reportd.retry_times node value
func UpdateSandboxReportdRetryTimes(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandbox_reportd.retry_times", bytes.NewReader(byt))
	return
}

// GetSandboxdBypassSandboxLimit gets the sandboxd.bypass_sandbox_limit node value
func GetSandboxdBypassSandboxLimit(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.bypass_sandbox_limit")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdBypassSandboxLimit updates the sandboxd.bypass_sandbox_limit node value
func UpdateSandboxdBypassSandboxLimit(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.bypass_sandbox_limit", bytes.NewReader(byt))
	return
}

// GetSandboxdCacheExpire gets the sandboxd.cache_expire node value
func GetSandboxdCacheExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.cache_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdCacheExpire updates the sandboxd.cache_expire node value
func UpdateSandboxdCacheExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.cache_expire", bytes.NewReader(byt))
	return
}

// GetSandboxdCertstore gets the sandboxd.certstore node value
func GetSandboxdCertstore(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.certstore")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdCertstore updates the sandboxd.certstore node value
func UpdateSandboxdCertstore(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.certstore", bytes.NewReader(byt))
	return
}

// GetSandboxdClientTimeout gets the sandboxd.client_timeout node value
func GetSandboxdClientTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.client_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdClientTimeout updates the sandboxd.client_timeout node value
func UpdateSandboxdClientTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.client_timeout", bytes.NewReader(byt))
	return
}

// GetSandboxdCloudPollInterval gets the sandboxd.cloud_poll_interval node value
func GetSandboxdCloudPollInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.cloud_poll_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdCloudPollInterval updates the sandboxd.cloud_poll_interval node value
func UpdateSandboxdCloudPollInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.cloud_poll_interval", bytes.NewReader(byt))
	return
}

// GetSandboxdCloudPollRequestMaximum gets the sandboxd.cloud_poll_request_maximum node value
func GetSandboxdCloudPollRequestMaximum(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.cloud_poll_request_maximum")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdCloudPollRequestMaximum updates the sandboxd.cloud_poll_request_maximum node value
func UpdateSandboxdCloudPollRequestMaximum(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.cloud_poll_request_maximum", bytes.NewReader(byt))
	return
}

// GetSandboxdCloudPollTimeout gets the sandboxd.cloud_poll_timeout node value
func GetSandboxdCloudPollTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.cloud_poll_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdCloudPollTimeout updates the sandboxd.cloud_poll_timeout node value
func UpdateSandboxdCloudPollTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.cloud_poll_timeout", bytes.NewReader(byt))
	return
}

// GetSandboxdConnectTimeout gets the sandboxd.connect_timeout node value
func GetSandboxdConnectTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.connect_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdConnectTimeout updates the sandboxd.connect_timeout node value
func UpdateSandboxdConnectTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.connect_timeout", bytes.NewReader(byt))
	return
}

// GetSandboxdDebug gets the sandboxd.debug node value
func GetSandboxdDebug(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/sandboxd.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdDebug updates the sandboxd.debug node value
func UpdateSandboxdDebug(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.debug", bytes.NewReader(byt))
	return
}

// GetSandboxdDhparams2048 gets the sandboxd.dhparams_2048 node value
func GetSandboxdDhparams2048(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.dhparams_2048")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdDhparams2048 updates the sandboxd.dhparams_2048 node value
func UpdateSandboxdDhparams2048(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.dhparams_2048", bytes.NewReader(byt))
	return
}

// GetSandboxdEdgeServerCertValidation gets the sandboxd.edge_server_cert_validation node value
func GetSandboxdEdgeServerCertValidation(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/sandboxd.edge_server_cert_validation")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdEdgeServerCertValidation updates the sandboxd.edge_server_cert_validation node value
func UpdateSandboxdEdgeServerCertValidation(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.edge_server_cert_validation", bytes.NewReader(byt))
	return
}

// GetSandboxdEdgeServerHost gets the sandboxd.edge_server_host node value
func GetSandboxdEdgeServerHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.edge_server_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdEdgeServerHost updates the sandboxd.edge_server_host node value
func UpdateSandboxdEdgeServerHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.edge_server_host", bytes.NewReader(byt))
	return
}

// GetSandboxdEdgeServerPort gets the sandboxd.edge_server_port node value
func GetSandboxdEdgeServerPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.edge_server_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdEdgeServerPort updates the sandboxd.edge_server_port node value
func UpdateSandboxdEdgeServerPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.edge_server_port", bytes.NewReader(byt))
	return
}

// GetSandboxdFiletypeSkiplist gets the sandboxd.filetype_skiplist node value
func GetSandboxdFiletypeSkiplist(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/sandboxd.filetype_skiplist")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdFiletypeSkiplist updates the sandboxd.filetype_skiplist node value
func UpdateSandboxdFiletypeSkiplist(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.filetype_skiplist", bytes.NewReader(byt))
	return
}

// GetSandboxdNumOfEventThreads gets the sandboxd.num_of_event_threads node value
func GetSandboxdNumOfEventThreads(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.num_of_event_threads")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdNumOfEventThreads updates the sandboxd.num_of_event_threads node value
func UpdateSandboxdNumOfEventThreads(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.num_of_event_threads", bytes.NewReader(byt))
	return
}

// GetSandboxdNumOfWorkerThreads gets the sandboxd.num_of_worker_threads node value
func GetSandboxdNumOfWorkerThreads(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.num_of_worker_threads")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdNumOfWorkerThreads updates the sandboxd.num_of_worker_threads node value
func UpdateSandboxdNumOfWorkerThreads(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.num_of_worker_threads", bytes.NewReader(byt))
	return
}

// GetSandboxdResponseTimeout gets the sandboxd.response_timeout node value
func GetSandboxdResponseTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.response_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdResponseTimeout updates the sandboxd.response_timeout node value
func UpdateSandboxdResponseTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.response_timeout", bytes.NewReader(byt))
	return
}

// GetSandboxdRetryIntervals gets the sandboxd.retry_intervals node value
func GetSandboxdRetryIntervals(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.retry_intervals")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdRetryIntervals updates the sandboxd.retry_intervals node value
func UpdateSandboxdRetryIntervals(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.retry_intervals", bytes.NewReader(byt))
	return
}

// GetSandboxdSandboxEnabled gets the sandboxd.sandbox_enabled node value
func GetSandboxdSandboxEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/sandboxd.sandbox_enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSandboxEnabled updates the sandboxd.sandbox_enabled node value
func UpdateSandboxdSandboxEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.sandbox_enabled", bytes.NewReader(byt))
	return
}

// GetSandboxdSandboxdScoreThreshold gets the sandboxd.sandboxd_score_threshold node value
func GetSandboxdSandboxdScoreThreshold(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.sandboxd_score_threshold")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSandboxdScoreThreshold updates the sandboxd.sandboxd_score_threshold node value
func UpdateSandboxdSandboxdScoreThreshold(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.sandboxd_score_threshold", bytes.NewReader(byt))
	return
}

// GetSandboxdSbxVersion gets the sandboxd.sbx_version node value
func GetSandboxdSbxVersion(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.sbx_version")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSbxVersion updates the sandboxd.sbx_version node value
func UpdateSandboxdSbxVersion(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.sbx_version", bytes.NewReader(byt))
	return
}

// GetSandboxdSqliteBusyTimeout gets the sandboxd.sqlite_busy_timeout node value
func GetSandboxdSqliteBusyTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.sqlite_busy_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSqliteBusyTimeout updates the sandboxd.sqlite_busy_timeout node value
func UpdateSandboxdSqliteBusyTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.sqlite_busy_timeout", bytes.NewReader(byt))
	return
}

// GetSandboxdSslCertFile gets the sandboxd.ssl_cert_file node value
func GetSandboxdSslCertFile(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.ssl_cert_file")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSslCertFile updates the sandboxd.ssl_cert_file node value
func UpdateSandboxdSslCertFile(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.ssl_cert_file", bytes.NewReader(byt))
	return
}

// GetSandboxdSslKeyFile gets the sandboxd.ssl_key_file node value
func GetSandboxdSslKeyFile(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.ssl_key_file")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSslKeyFile updates the sandboxd.ssl_key_file node value
func UpdateSandboxdSslKeyFile(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.ssl_key_file", bytes.NewReader(byt))
	return
}

// GetSandboxdSslciphers gets the sandboxd.sslciphers node value
func GetSandboxdSslciphers(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.sslciphers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSslciphers updates the sandboxd.sslciphers node value
func UpdateSandboxdSslciphers(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.sslciphers", bytes.NewReader(byt))
	return
}

// GetSettingsAdminEmail gets the settings.admin_email node value
func GetSettingsAdminEmail(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.admin_email")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsAdminEmail updates the settings.admin_email node value
func UpdateSettingsAdminEmail(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.admin_email", bytes.NewReader(byt))
	return
}

// GetSettingsBasicSetup gets the settings.basic_setup node value
func GetSettingsBasicSetup(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/settings.basic_setup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsBasicSetup updates the settings.basic_setup node value
func UpdateSettingsBasicSetup(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.basic_setup", bytes.NewReader(byt))
	return
}

// GetSettingsCcMode gets the settings.cc_mode node value
func GetSettingsCcMode(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/settings.cc_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsCcMode updates the settings.cc_mode node value
func UpdateSettingsCcMode(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.cc_mode", bytes.NewReader(byt))
	return
}

// GetSettingsCity gets the settings.city node value
func GetSettingsCity(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.city")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsCity updates the settings.city node value
func UpdateSettingsCity(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.city", bytes.NewReader(byt))
	return
}

// GetSettingsCountry gets the settings.country node value
func GetSettingsCountry(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.country")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsCountry updates the settings.country node value
func UpdateSettingsCountry(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.country", bytes.NewReader(byt))
	return
}

// GetSettingsExtraSwap gets the settings.extra_swap node value
func GetSettingsExtraSwap(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/settings.extra_swap")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsExtraSwap updates the settings.extra_swap node value
func UpdateSettingsExtraSwap(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.extra_swap", bytes.NewReader(byt))
	return
}

// GetSettingsHostname gets the settings.hostname node value
func GetSettingsHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsHostname updates the settings.hostname node value
func UpdateSettingsHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.hostname", bytes.NewReader(byt))
	return
}

// GetSettingsIcsaMode gets the settings.icsa_mode node value
func GetSettingsIcsaMode(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/settings.icsa_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsIcsaMode updates the settings.icsa_mode node value
func UpdateSettingsIcsaMode(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.icsa_mode", bytes.NewReader(byt))
	return
}

// GetSettingsOrganization gets the settings.organization node value
func GetSettingsOrganization(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.organization")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsOrganization updates the settings.organization node value
func UpdateSettingsOrganization(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.organization", bytes.NewReader(byt))
	return
}

// GetSettingsPasswordComplexityMinDigits gets the settings.password_complexity.min_digits node value
func GetSettingsPasswordComplexityMinDigits(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/settings.password_complexity.min_digits")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPasswordComplexityMinDigits updates the settings.password_complexity.min_digits node value
func UpdateSettingsPasswordComplexityMinDigits(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.password_complexity.min_digits", bytes.NewReader(byt))
	return
}

// GetSettingsPasswordComplexityMinLength gets the settings.password_complexity.min_length node value
func GetSettingsPasswordComplexityMinLength(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/settings.password_complexity.min_length")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPasswordComplexityMinLength updates the settings.password_complexity.min_length node value
func UpdateSettingsPasswordComplexityMinLength(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.password_complexity.min_length", bytes.NewReader(byt))
	return
}

// GetSettingsPasswordComplexityMinLowerChars gets the settings.password_complexity.min_lower_chars node value
func GetSettingsPasswordComplexityMinLowerChars(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/settings.password_complexity.min_lower_chars")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPasswordComplexityMinLowerChars updates the settings.password_complexity.min_lower_chars node value
func UpdateSettingsPasswordComplexityMinLowerChars(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.password_complexity.min_lower_chars", bytes.NewReader(byt))
	return
}

// GetSettingsPasswordComplexityMinSpecialChars gets the settings.password_complexity.min_special_chars node value
func GetSettingsPasswordComplexityMinSpecialChars(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/settings.password_complexity.min_special_chars")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPasswordComplexityMinSpecialChars updates the settings.password_complexity.min_special_chars node value
func UpdateSettingsPasswordComplexityMinSpecialChars(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.password_complexity.min_special_chars", bytes.NewReader(byt))
	return
}

// GetSettingsPasswordComplexityMinUpperChars gets the settings.password_complexity.min_upper_chars node value
func GetSettingsPasswordComplexityMinUpperChars(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/settings.password_complexity.min_upper_chars")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPasswordComplexityMinUpperChars updates the settings.password_complexity.min_upper_chars node value
func UpdateSettingsPasswordComplexityMinUpperChars(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.password_complexity.min_upper_chars", bytes.NewReader(byt))
	return
}

// GetSettingsPasswordComplexityStatus gets the settings.password_complexity.status node value
func GetSettingsPasswordComplexityStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/settings.password_complexity.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPasswordComplexityStatus updates the settings.password_complexity.status node value
func UpdateSettingsPasswordComplexityStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.password_complexity.status", bytes.NewReader(byt))
	return
}

// GetSettingsPopularity gets the settings.popularity node value
func GetSettingsPopularity(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.popularity")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPopularity updates the settings.popularity node value
func UpdateSettingsPopularity(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.popularity", bytes.NewReader(byt))
	return
}

// GetSettingsRasUpdate gets the settings.ras_update node value
func GetSettingsRasUpdate(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.ras_update")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsRasUpdate updates the settings.ras_update node value
func UpdateSettingsRasUpdate(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.ras_update", bytes.NewReader(byt))
	return
}

// GetSettingsSystemId gets the settings.system_id node value
func GetSettingsSystemId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.system_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsSystemId updates the settings.system_id node value
func UpdateSettingsSystemId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.system_id", bytes.NewReader(byt))
	return
}

// GetSettingsTimezone gets the settings.timezone node value
func GetSettingsTimezone(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.timezone")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsTimezone updates the settings.timezone node value
func UpdateSettingsTimezone(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.timezone", bytes.NewReader(byt))
	return
}

// GetSipAllowedNetworks gets the sip.allowed_networks node value
func GetSipAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/sip.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSipAllowedNetworks updates the sip.allowed_networks node value
func UpdateSipAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sip.allowed_networks", bytes.NewReader(byt))
	return
}

// GetSipExpectMode gets the sip.expect_mode node value
func GetSipExpectMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sip.expect_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSipExpectMode updates the sip.expect_mode node value
func UpdateSipExpectMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sip.expect_mode", bytes.NewReader(byt))
	return
}

// GetSipLogRelated gets the sip.log_related node value
func GetSipLogRelated(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/sip.log_related")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSipLogRelated updates the sip.log_related node value
func UpdateSipLogRelated(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sip.log_related", bytes.NewReader(byt))
	return
}

// GetSipServers gets the sip.servers node value
func GetSipServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/sip.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSipServers updates the sip.servers node value
func UpdateSipServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sip.servers", bytes.NewReader(byt))
	return
}

// GetSipStatus gets the sip.status node value
func GetSipStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/sip.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSipStatus updates the sip.status node value
func UpdateSipStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sip.status", bytes.NewReader(byt))
	return
}

// GetSmsClientHostname gets the sms_client.hostname node value
func GetSmsClientHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sms_client.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmsClientHostname updates the sms_client.hostname node value
func UpdateSmsClientHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sms_client.hostname", bytes.NewReader(byt))
	return
}

// GetSmsClientName gets the sms_client.name node value
func GetSmsClientName(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sms_client.name")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmsClientName updates the sms_client.name node value
func UpdateSmsClientName(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sms_client.name", bytes.NewReader(byt))
	return
}

// GetSmsClientPassword gets the sms_client.password node value
func GetSmsClientPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sms_client.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmsClientPassword updates the sms_client.password node value
func UpdateSmsClientPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sms_client.password", bytes.NewReader(byt))
	return
}

// GetSmsClientPort gets the sms_client.port node value
func GetSmsClientPort(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/sms_client.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmsClientPort updates the sms_client.port node value
func UpdateSmsClientPort(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sms_client.port", bytes.NewReader(byt))
	return
}

// GetSmsClientStatus gets the sms_client.status node value
func GetSmsClientStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/sms_client.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmsClientStatus updates the sms_client.status node value
func UpdateSmsClientStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sms_client.status", bytes.NewReader(byt))
	return
}

// GetSmsClientUsername gets the sms_client.username node value
func GetSmsClientUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sms_client.username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmsClientUsername updates the sms_client.username node value
func UpdateSmsClientUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sms_client.username", bytes.NewReader(byt))
	return
}

// GetSmtpAuthAaa gets the smtp.auth_aaa node value
func GetSmtpAuthAaa(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.auth_aaa")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpAuthAaa updates the smtp.auth_aaa node value
func UpdateSmtpAuthAaa(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.auth_aaa", bytes.NewReader(byt))
	return
}

// GetSmtpAuthStatus gets the smtp.auth_status node value
func GetSmtpAuthStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.auth_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpAuthStatus updates the smtp.auth_status node value
func UpdateSmtpAuthStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.auth_status", bytes.NewReader(byt))
	return
}

// GetSmtpAvFooter gets the smtp.av_footer node value
func GetSmtpAvFooter(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.av_footer")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpAvFooter updates the smtp.av_footer node value
func UpdateSmtpAvFooter(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.av_footer", bytes.NewReader(byt))
	return
}

// GetSmtpAvFooterStatus gets the smtp.av_footer_status node value
func GetSmtpAvFooterStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.av_footer_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpAvFooterStatus updates the smtp.av_footer_status node value
func UpdateSmtpAvFooterStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.av_footer_status", bytes.NewReader(byt))
	return
}

// GetSmtpBatvSecret gets the smtp.batv_secret node value
func GetSmtpBatvSecret(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.batv_secret")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpBatvSecret updates the smtp.batv_secret node value
func UpdateSmtpBatvSecret(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.batv_secret", bytes.NewReader(byt))
	return
}

// GetSmtpBlockerService gets the smtp.blocker_service node value
func GetSmtpBlockerService(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.blocker_service")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpBlockerService updates the smtp.blocker_service node value
func UpdateSmtpBlockerService(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.blocker_service", bytes.NewReader(byt))
	return
}

// GetSmtpCffAsMarker gets the smtp.cff_as_marker node value
func GetSmtpCffAsMarker(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.cff_as_marker")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpCffAsMarker updates the smtp.cff_as_marker node value
func UpdateSmtpCffAsMarker(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.cff_as_marker", bytes.NewReader(byt))
	return
}

// GetSmtpDkimDomains gets the smtp.dkim_domains node value
func GetSmtpDkimDomains(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.dkim_domains")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpDkimDomains updates the smtp.dkim_domains node value
func UpdateSmtpDkimDomains(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.dkim_domains", bytes.NewReader(byt))
	return
}

// GetSmtpDkimPrivateKey gets the smtp.dkim_private_key node value
func GetSmtpDkimPrivateKey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.dkim_private_key")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpDkimPrivateKey updates the smtp.dkim_private_key node value
func UpdateSmtpDkimPrivateKey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.dkim_private_key", bytes.NewReader(byt))
	return
}

// GetSmtpDkimSelector gets the smtp.dkim_selector node value
func GetSmtpDkimSelector(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.dkim_selector")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpDkimSelector updates the smtp.dkim_selector node value
func UpdateSmtpDkimSelector(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.dkim_selector", bytes.NewReader(byt))
	return
}

// GetSmtpEnableOldExpressionFilter gets the smtp.enable_old_expression_filter node value
func GetSmtpEnableOldExpressionFilter(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.enable_old_expression_filter")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpEnableOldExpressionFilter updates the smtp.enable_old_expression_filter node value
func UpdateSmtpEnableOldExpressionFilter(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.enable_old_expression_filter", bytes.NewReader(byt))
	return
}

// GetSmtpEncryptionUtility gets the smtp.encryption_utility node value
func GetSmtpEncryptionUtility(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.encryption_utility")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpEncryptionUtility updates the smtp.encryption_utility node value
func UpdateSmtpEncryptionUtility(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.encryption_utility", bytes.NewReader(byt))
	return
}

// GetSmtpExceptions gets the smtp.exceptions node value
func GetSmtpExceptions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpExceptions updates the smtp.exceptions node value
func UpdateSmtpExceptions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.exceptions", bytes.NewReader(byt))
	return
}

// GetSmtpFootersMode gets the smtp.footers_mode node value
func GetSmtpFootersMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.footers_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpFootersMode updates the smtp.footers_mode node value
func UpdateSmtpFootersMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.footers_mode", bytes.NewReader(byt))
	return
}

// GetSmtpGlobalAsReject gets the smtp.global_as_reject node value
func GetSmtpGlobalAsReject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.global_as_reject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpGlobalAsReject updates the smtp.global_as_reject node value
func UpdateSmtpGlobalAsReject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.global_as_reject", bytes.NewReader(byt))
	return
}

// GetSmtpGlobalAvReject gets the smtp.global_av_reject node value
func GetSmtpGlobalAvReject(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.global_av_reject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpGlobalAvReject updates the smtp.global_av_reject node value
func UpdateSmtpGlobalAvReject(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.global_av_reject", bytes.NewReader(byt))
	return
}

// GetSmtpGlobalProfile gets the smtp.global_profile node value
func GetSmtpGlobalProfile(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.global_profile")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpGlobalProfile updates the smtp.global_profile node value
func UpdateSmtpGlobalProfile(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.global_profile", bytes.NewReader(byt))
	return
}

// GetSmtpHostBlacklist gets the smtp.host_blacklist node value
func GetSmtpHostBlacklist(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.host_blacklist")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpHostBlacklist updates the smtp.host_blacklist node value
func UpdateSmtpHostBlacklist(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.host_blacklist", bytes.NewReader(byt))
	return
}

// GetSmtpHostname gets the smtp.hostname node value
func GetSmtpHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpHostname updates the smtp.hostname node value
func UpdateSmtpHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.hostname", bytes.NewReader(byt))
	return
}

// GetSmtpMaxMessageSize gets the smtp.max_message_size node value
func GetSmtpMaxMessageSize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.max_message_size")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpMaxMessageSize updates the smtp.max_message_size node value
func UpdateSmtpMaxMessageSize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.max_message_size", bytes.NewReader(byt))
	return
}

// GetSmtpMode gets the smtp.mode node value
func GetSmtpMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpMode updates the smtp.mode node value
func UpdateSmtpMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.mode", bytes.NewReader(byt))
	return
}

// GetSmtpParentProxyAuthPass gets the smtp.parent_proxy_auth_pass node value
func GetSmtpParentProxyAuthPass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.parent_proxy_auth_pass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpParentProxyAuthPass updates the smtp.parent_proxy_auth_pass node value
func UpdateSmtpParentProxyAuthPass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.parent_proxy_auth_pass", bytes.NewReader(byt))
	return
}

// GetSmtpParentProxyAuthStatus gets the smtp.parent_proxy_auth_status node value
func GetSmtpParentProxyAuthStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.parent_proxy_auth_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpParentProxyAuthStatus updates the smtp.parent_proxy_auth_status node value
func UpdateSmtpParentProxyAuthStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.parent_proxy_auth_status", bytes.NewReader(byt))
	return
}

// GetSmtpParentProxyAuthUser gets the smtp.parent_proxy_auth_user node value
func GetSmtpParentProxyAuthUser(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.parent_proxy_auth_user")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpParentProxyAuthUser updates the smtp.parent_proxy_auth_user node value
func UpdateSmtpParentProxyAuthUser(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.parent_proxy_auth_user", bytes.NewReader(byt))
	return
}

// GetSmtpParentProxyHost gets the smtp.parent_proxy_host node value
func GetSmtpParentProxyHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.parent_proxy_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpParentProxyHost updates the smtp.parent_proxy_host node value
func UpdateSmtpParentProxyHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.parent_proxy_host", bytes.NewReader(byt))
	return
}

// GetSmtpParentProxyPort gets the smtp.parent_proxy_port node value
func GetSmtpParentProxyPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.parent_proxy_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpParentProxyPort updates the smtp.parent_proxy_port node value
func UpdateSmtpParentProxyPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.parent_proxy_port", bytes.NewReader(byt))
	return
}

// GetSmtpParentProxyStatus gets the smtp.parent_proxy_status node value
func GetSmtpParentProxyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.parent_proxy_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpParentProxyStatus updates the smtp.parent_proxy_status node value
func UpdateSmtpParentProxyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.parent_proxy_status", bytes.NewReader(byt))
	return
}

// GetSmtpPostmaster gets the smtp.postmaster node value
func GetSmtpPostmaster(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.postmaster")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpPostmaster updates the smtp.postmaster node value
func UpdateSmtpPostmaster(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.postmaster", bytes.NewReader(byt))
	return
}

// GetSmtpProfiles gets the smtp.profiles node value
func GetSmtpProfiles(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.profiles")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpProfiles updates the smtp.profiles node value
func UpdateSmtpProfiles(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.profiles", bytes.NewReader(byt))
	return
}

// GetSmtpRecipientsMax gets the smtp.recipients_max node value
func GetSmtpRecipientsMax(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.recipients_max")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpRecipientsMax updates the smtp.recipients_max node value
func UpdateSmtpRecipientsMax(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.recipients_max", bytes.NewReader(byt))
	return
}

// GetSmtpRelays gets the smtp.relays node value
func GetSmtpRelays(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.relays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpRelays updates the smtp.relays node value
func UpdateSmtpRelays(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.relays", bytes.NewReader(byt))
	return
}

// GetSmtpScanOutgoingEmails gets the smtp.scan_outgoing_emails node value
func GetSmtpScanOutgoingEmails(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.scan_outgoing_emails")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpScanOutgoingEmails updates the smtp.scan_outgoing_emails node value
func UpdateSmtpScanOutgoingEmails(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.scan_outgoing_emails", bytes.NewReader(byt))
	return
}

// GetSmtpScannerPoolMaxPool gets the smtp.scanner_pool.max_pool node value
func GetSmtpScannerPoolMaxPool(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.scanner_pool.max_pool")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpScannerPoolMaxPool updates the smtp.scanner_pool.max_pool node value
func UpdateSmtpScannerPoolMaxPool(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.scanner_pool.max_pool", bytes.NewReader(byt))
	return
}

// GetSmtpScannerPoolThresholds gets the smtp.scanner_pool.thresholds node value
func GetSmtpScannerPoolThresholds(c sophos.ClientInterface) (val []int64, err error) {
	res, err := c.Get("/api/nodes/smtp.scanner_pool.thresholds")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpScannerPoolThresholds updates the smtp.scanner_pool.thresholds node value
func UpdateSmtpScannerPoolThresholds(c sophos.ClientInterface, val []int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.scanner_pool.thresholds", bytes.NewReader(byt))
	return
}

// GetSmtpScannerTimeout gets the smtp.scanner_timeout node value
func GetSmtpScannerTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.scanner_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpScannerTimeout updates the smtp.scanner_timeout node value
func UpdateSmtpScannerTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.scanner_timeout", bytes.NewReader(byt))
	return
}

// GetSmtpSmarthostAuth gets the smtp.smarthost_auth node value
func GetSmtpSmarthostAuth(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.smarthost_auth")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmarthostAuth updates the smtp.smarthost_auth node value
func UpdateSmtpSmarthostAuth(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smarthost_auth", bytes.NewReader(byt))
	return
}

// GetSmtpSmarthostHostname gets the smtp.smarthost_hostname node value
func GetSmtpSmarthostHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.smarthost_hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmarthostHostname updates the smtp.smarthost_hostname node value
func UpdateSmtpSmarthostHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smarthost_hostname", bytes.NewReader(byt))
	return
}

// GetSmtpSmarthostPass gets the smtp.smarthost_pass node value
func GetSmtpSmarthostPass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.smarthost_pass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmarthostPass updates the smtp.smarthost_pass node value
func UpdateSmtpSmarthostPass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smarthost_pass", bytes.NewReader(byt))
	return
}

// GetSmtpSmarthostPort gets the smtp.smarthost_port node value
func GetSmtpSmarthostPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.smarthost_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmarthostPort updates the smtp.smarthost_port node value
func UpdateSmtpSmarthostPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smarthost_port", bytes.NewReader(byt))
	return
}

// GetSmtpSmarthostStatus gets the smtp.smarthost_status node value
func GetSmtpSmarthostStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.smarthost_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmarthostStatus updates the smtp.smarthost_status node value
func UpdateSmtpSmarthostStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smarthost_status", bytes.NewReader(byt))
	return
}

// GetSmtpSmarthostUser gets the smtp.smarthost_user node value
func GetSmtpSmarthostUser(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.smarthost_user")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmarthostUser updates the smtp.smarthost_user node value
func UpdateSmtpSmarthostUser(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smarthost_user", bytes.NewReader(byt))
	return
}

// GetSmtpSmtpAcceptMax gets the smtp.smtp_accept_max node value
func GetSmtpSmtpAcceptMax(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.smtp_accept_max")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmtpAcceptMax updates the smtp.smtp_accept_max node value
func UpdateSmtpSmtpAcceptMax(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smtp_accept_max", bytes.NewReader(byt))
	return
}

// GetSmtpSmtpAcceptPerConnectionMax gets the smtp.smtp_accept_per_connection_max node value
func GetSmtpSmtpAcceptPerConnectionMax(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.smtp_accept_per_connection_max")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmtpAcceptPerConnectionMax updates the smtp.smtp_accept_per_connection_max node value
func UpdateSmtpSmtpAcceptPerConnectionMax(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smtp_accept_per_connection_max", bytes.NewReader(byt))
	return
}

// GetSmtpSmtpAcceptPerHostMax gets the smtp.smtp_accept_per_host_max node value
func GetSmtpSmtpAcceptPerHostMax(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.smtp_accept_per_host_max")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmtpAcceptPerHostMax updates the smtp.smtp_accept_per_host_max node value
func UpdateSmtpSmtpAcceptPerHostMax(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smtp_accept_per_host_max", bytes.NewReader(byt))
	return
}

// GetSmtpStatus gets the smtp.status node value
func GetSmtpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpStatus updates the smtp.status node value
func UpdateSmtpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.status", bytes.NewReader(byt))
	return
}

// GetSmtpTlsAvoid gets the smtp.tls_avoid node value
func GetSmtpTlsAvoid(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.tls_avoid")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTlsAvoid updates the smtp.tls_avoid node value
func UpdateSmtpTlsAvoid(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.tls_avoid", bytes.NewReader(byt))
	return
}

// GetSmtpTlsCert gets the smtp.tls_cert node value
func GetSmtpTlsCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.tls_cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTlsCert updates the smtp.tls_cert node value
func UpdateSmtpTlsCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.tls_cert", bytes.NewReader(byt))
	return
}

// GetSmtpTlsRequire gets the smtp.tls_require node value
func GetSmtpTlsRequire(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.tls_require")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTlsRequire updates the smtp.tls_require node value
func UpdateSmtpTlsRequire(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.tls_require", bytes.NewReader(byt))
	return
}

// GetSmtpTlsRequireSenderDomains gets the smtp.tls_require_sender_domains node value
func GetSmtpTlsRequireSenderDomains(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.tls_require_sender_domains")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTlsRequireSenderDomains updates the smtp.tls_require_sender_domains node value
func UpdateSmtpTlsRequireSenderDomains(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.tls_require_sender_domains", bytes.NewReader(byt))
	return
}

// GetSmtpTransparent gets the smtp.transparent node value
func GetSmtpTransparent(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.transparent")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTransparent updates the smtp.transparent node value
func UpdateSmtpTransparent(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.transparent", bytes.NewReader(byt))
	return
}

// GetSmtpTransparentSkip gets the smtp.transparent_skip node value
func GetSmtpTransparentSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.transparent_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTransparentSkip updates the smtp.transparent_skip node value
func UpdateSmtpTransparentSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.transparent_skip", bytes.NewReader(byt))
	return
}

// GetSmtpTransparentSkipAutoPf gets the smtp.transparent_skip_auto_pf node value
func GetSmtpTransparentSkipAutoPf(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.transparent_skip_auto_pf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTransparentSkipAutoPf updates the smtp.transparent_skip_auto_pf node value
func UpdateSmtpTransparentSkipAutoPf(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.transparent_skip_auto_pf", bytes.NewReader(byt))
	return
}

// GetSmtpUpstreamHosts gets the smtp.upstream_hosts node value
func GetSmtpUpstreamHosts(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.upstream_hosts")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpUpstreamHosts updates the smtp.upstream_hosts node value
func UpdateSmtpUpstreamHosts(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.upstream_hosts", bytes.NewReader(byt))
	return
}

// GetSmtpUpstreamHostsOnly gets the smtp.upstream_hosts_only node value
func GetSmtpUpstreamHostsOnly(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.upstream_hosts_only")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpUpstreamHostsOnly updates the smtp.upstream_hosts_only node value
func UpdateSmtpUpstreamHostsOnly(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.upstream_hosts_only", bytes.NewReader(byt))
	return
}

// GetSnmpAllowedNetworks gets the snmp.allowed_networks node value
func GetSnmpAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/snmp.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpAllowedNetworks updates the snmp.allowed_networks node value
func UpdateSnmpAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.allowed_networks", bytes.NewReader(byt))
	return
}

// GetSnmpAuthPassword gets the snmp.auth_password node value
func GetSnmpAuthPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.auth_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpAuthPassword updates the snmp.auth_password node value
func UpdateSnmpAuthPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.auth_password", bytes.NewReader(byt))
	return
}

// GetSnmpAuthType gets the snmp.auth_type node value
func GetSnmpAuthType(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.auth_type")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpAuthType updates the snmp.auth_type node value
func UpdateSnmpAuthType(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.auth_type", bytes.NewReader(byt))
	return
}

// GetSnmpCommunity gets the snmp.community node value
func GetSnmpCommunity(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.community")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpCommunity updates the snmp.community node value
func UpdateSnmpCommunity(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.community", bytes.NewReader(byt))
	return
}

// GetSnmpDeviceAdmin gets the snmp.device_admin node value
func GetSnmpDeviceAdmin(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.device_admin")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpDeviceAdmin updates the snmp.device_admin node value
func UpdateSnmpDeviceAdmin(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.device_admin", bytes.NewReader(byt))
	return
}

// GetSnmpDeviceLocation gets the snmp.device_location node value
func GetSnmpDeviceLocation(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.device_location")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpDeviceLocation updates the snmp.device_location node value
func UpdateSnmpDeviceLocation(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.device_location", bytes.NewReader(byt))
	return
}

// GetSnmpDeviceName gets the snmp.device_name node value
func GetSnmpDeviceName(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.device_name")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpDeviceName updates the snmp.device_name node value
func UpdateSnmpDeviceName(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.device_name", bytes.NewReader(byt))
	return
}

// GetSnmpEncryptPassword gets the snmp.encrypt_password node value
func GetSnmpEncryptPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.encrypt_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpEncryptPassword updates the snmp.encrypt_password node value
func UpdateSnmpEncryptPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.encrypt_password", bytes.NewReader(byt))
	return
}

// GetSnmpEncryptType gets the snmp.encrypt_type node value
func GetSnmpEncryptType(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.encrypt_type")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpEncryptType updates the snmp.encrypt_type node value
func UpdateSnmpEncryptType(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.encrypt_type", bytes.NewReader(byt))
	return
}

// GetSnmpStatus gets the snmp.status node value
func GetSnmpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/snmp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpStatus updates the snmp.status node value
func UpdateSnmpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.status", bytes.NewReader(byt))
	return
}

// GetSnmpTraps gets the snmp.traps node value
func GetSnmpTraps(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/snmp.traps")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpTraps updates the snmp.traps node value
func UpdateSnmpTraps(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.traps", bytes.NewReader(byt))
	return
}

// GetSnmpTrapsUseOldOids gets the snmp.traps_use_old_oids node value
func GetSnmpTrapsUseOldOids(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/snmp.traps_use_old_oids")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpTrapsUseOldOids updates the snmp.traps_use_old_oids node value
func UpdateSnmpTrapsUseOldOids(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.traps_use_old_oids", bytes.NewReader(byt))
	return
}

// GetSnmpUsername gets the snmp.username node value
func GetSnmpUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpUsername updates the snmp.username node value
func UpdateSnmpUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.username", bytes.NewReader(byt))
	return
}

// GetSnmpVersion gets the snmp.version node value
func GetSnmpVersion(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.version")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpVersion updates the snmp.version node value
func UpdateSnmpVersion(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.version", bytes.NewReader(byt))
	return
}

// GetSocksAaa gets the socks.aaa node value
func GetSocksAaa(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/socks.aaa")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSocksAaa updates the socks.aaa node value
func UpdateSocksAaa(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/socks.aaa", bytes.NewReader(byt))
	return
}

// GetSocksAllowedNetworks gets the socks.allowed_networks node value
func GetSocksAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/socks.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSocksAllowedNetworks updates the socks.allowed_networks node value
func UpdateSocksAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/socks.allowed_networks", bytes.NewReader(byt))
	return
}

// GetSocksAuthentication gets the socks.authentication node value
func GetSocksAuthentication(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/socks.authentication")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSocksAuthentication updates the socks.authentication node value
func UpdateSocksAuthentication(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/socks.authentication", bytes.NewReader(byt))
	return
}

// GetSocksStatus gets the socks.status node value
func GetSocksStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/socks.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSocksStatus updates the socks.status node value
func UpdateSocksStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/socks.status", bytes.NewReader(byt))
	return
}

// GetSpxGlobalErrorNotificationTarget gets the spx.global.error_notification_target node value
func GetSpxGlobalErrorNotificationTarget(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.global.error_notification_target")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalErrorNotificationTarget updates the spx.global.error_notification_target node value
func UpdateSpxGlobalErrorNotificationTarget(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.error_notification_target", bytes.NewReader(byt))
	return
}

// GetSpxGlobalExpirySettingsAllowSecureReplyDays gets the spx.global.expiry_settings.allow_secure_reply_days node value
func GetSpxGlobalExpirySettingsAllowSecureReplyDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.global.expiry_settings.allow_secure_reply_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalExpirySettingsAllowSecureReplyDays updates the spx.global.expiry_settings.allow_secure_reply_days node value
func UpdateSpxGlobalExpirySettingsAllowSecureReplyDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.expiry_settings.allow_secure_reply_days", bytes.NewReader(byt))
	return
}

// GetSpxGlobalExpirySettingsKeepDelayedMailDays gets the spx.global.expiry_settings.keep_delayed_mail_days node value
func GetSpxGlobalExpirySettingsKeepDelayedMailDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.global.expiry_settings.keep_delayed_mail_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalExpirySettingsKeepDelayedMailDays updates the spx.global.expiry_settings.keep_delayed_mail_days node value
func UpdateSpxGlobalExpirySettingsKeepDelayedMailDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.expiry_settings.keep_delayed_mail_days", bytes.NewReader(byt))
	return
}

// GetSpxGlobalExpirySettingsKeepUnusedPwdDays gets the spx.global.expiry_settings.keep_unused_pwd_days node value
func GetSpxGlobalExpirySettingsKeepUnusedPwdDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.global.expiry_settings.keep_unused_pwd_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalExpirySettingsKeepUnusedPwdDays updates the spx.global.expiry_settings.keep_unused_pwd_days node value
func UpdateSpxGlobalExpirySettingsKeepUnusedPwdDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.expiry_settings.keep_unused_pwd_days", bytes.NewReader(byt))
	return
}

// GetSpxGlobalExpirySettingsRegistrationReminderDays gets the spx.global.expiry_settings.registration_reminder_days node value
func GetSpxGlobalExpirySettingsRegistrationReminderDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.global.expiry_settings.registration_reminder_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalExpirySettingsRegistrationReminderDays updates the spx.global.expiry_settings.registration_reminder_days node value
func UpdateSpxGlobalExpirySettingsRegistrationReminderDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.expiry_settings.registration_reminder_days", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPasswordStrengthMinLength gets the spx.global.password_strength.min_length node value
func GetSpxGlobalPasswordStrengthMinLength(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.global.password_strength.min_length")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPasswordStrengthMinLength updates the spx.global.password_strength.min_length node value
func UpdateSpxGlobalPasswordStrengthMinLength(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.password_strength.min_length", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPasswordStrengthRequireSpecChars gets the spx.global.password_strength.require_spec_chars node value
func GetSpxGlobalPasswordStrengthRequireSpecChars(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/spx.global.password_strength.require_spec_chars")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPasswordStrengthRequireSpecChars updates the spx.global.password_strength.require_spec_chars node value
func UpdateSpxGlobalPasswordStrengthRequireSpecChars(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.password_strength.require_spec_chars", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPortalSettingsAllowedNetworks gets the spx.global.portal_settings.allowed_networks node value
func GetSpxGlobalPortalSettingsAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/spx.global.portal_settings.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPortalSettingsAllowedNetworks updates the spx.global.portal_settings.allowed_networks node value
func UpdateSpxGlobalPortalSettingsAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.portal_settings.allowed_networks", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPortalSettingsHostname gets the spx.global.portal_settings.hostname node value
func GetSpxGlobalPortalSettingsHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.global.portal_settings.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPortalSettingsHostname updates the spx.global.portal_settings.hostname node value
func UpdateSpxGlobalPortalSettingsHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.portal_settings.hostname", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPortalSettingsInterfaceAddress gets the spx.global.portal_settings.interface_address node value
func GetSpxGlobalPortalSettingsInterfaceAddress(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.global.portal_settings.interface_address")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPortalSettingsInterfaceAddress updates the spx.global.portal_settings.interface_address node value
func UpdateSpxGlobalPortalSettingsInterfaceAddress(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.portal_settings.interface_address", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPortalSettingsPort gets the spx.global.portal_settings.port node value
func GetSpxGlobalPortalSettingsPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.global.portal_settings.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPortalSettingsPort updates the spx.global.portal_settings.port node value
func UpdateSpxGlobalPortalSettingsPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.portal_settings.port", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPoweredByLogo gets the spx.global.powered_by_logo node value
func GetSpxGlobalPoweredByLogo(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.global.powered_by_logo")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPoweredByLogo updates the spx.global.powered_by_logo node value
func UpdateSpxGlobalPoweredByLogo(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.powered_by_logo", bytes.NewReader(byt))
	return
}

// GetSpxGlobalSpxPriority gets the spx.global.spx_priority node value
func GetSpxGlobalSpxPriority(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/spx.global.spx_priority")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalSpxPriority updates the spx.global.spx_priority node value
func UpdateSpxGlobalSpxPriority(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.spx_priority", bytes.NewReader(byt))
	return
}

// GetSpxGlobalStatus gets the spx.global.status node value
func GetSpxGlobalStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/spx.global.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalStatus updates the spx.global.status node value
func UpdateSpxGlobalStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.status", bytes.NewReader(byt))
	return
}

// GetSpxSpxAuthPassword gets the spx.spx_auth.password node value
func GetSpxSpxAuthPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.spx_auth.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxSpxAuthPassword updates the spx.spx_auth.password node value
func UpdateSpxSpxAuthPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.spx_auth.password", bytes.NewReader(byt))
	return
}

// GetSpxSpxAuthPort gets the spx.spx_auth.port node value
func GetSpxSpxAuthPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.spx_auth.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxSpxAuthPort updates the spx.spx_auth.port node value
func UpdateSpxSpxAuthPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.spx_auth.port", bytes.NewReader(byt))
	return
}

// GetSpxSpxAuthServer gets the spx.spx_auth.server node value
func GetSpxSpxAuthServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.spx_auth.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxSpxAuthServer updates the spx.spx_auth.server node value
func UpdateSpxSpxAuthServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.spx_auth.server", bytes.NewReader(byt))
	return
}

// GetSpxSpxAuthUrl gets the spx.spx_auth.url node value
func GetSpxSpxAuthUrl(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.spx_auth.url")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxSpxAuthUrl updates the spx.spx_auth.url node value
func UpdateSpxSpxAuthUrl(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.spx_auth.url", bytes.NewReader(byt))
	return
}

// GetSpxTemplates gets the spx.templates node value
func GetSpxTemplates(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/spx.templates")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxTemplates updates the spx.templates node value
func UpdateSpxTemplates(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.templates", bytes.NewReader(byt))
	return
}

// GetSshAllowedNetworks gets the ssh.allowed_networks node value
func GetSshAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ssh.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshAllowedNetworks updates the ssh.allowed_networks node value
func UpdateSshAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.allowed_networks", bytes.NewReader(byt))
	return
}

// GetSshLoginKeys gets the ssh.login_keys node value
func GetSshLoginKeys(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ssh.login_keys")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshLoginKeys updates the ssh.login_keys node value
func UpdateSshLoginKeys(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.login_keys", bytes.NewReader(byt))
	return
}

// GetSshPasswordAuth gets the ssh.password_auth node value
func GetSshPasswordAuth(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssh.password_auth")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshPasswordAuth updates the ssh.password_auth node value
func UpdateSshPasswordAuth(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.password_auth", bytes.NewReader(byt))
	return
}

// GetSshPort gets the ssh.port node value
func GetSshPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ssh.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshPort updates the ssh.port node value
func UpdateSshPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.port", bytes.NewReader(byt))
	return
}

// GetSshPubkeyAuth gets the ssh.pubkey_auth node value
func GetSshPubkeyAuth(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssh.pubkey_auth")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshPubkeyAuth updates the ssh.pubkey_auth node value
func UpdateSshPubkeyAuth(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.pubkey_auth", bytes.NewReader(byt))
	return
}

// GetSshRootKeys gets the ssh.root_keys node value
func GetSshRootKeys(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ssh.root_keys")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshRootKeys updates the ssh.root_keys node value
func UpdateSshRootKeys(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.root_keys", bytes.NewReader(byt))
	return
}

// GetSshRootLogin gets the ssh.root_login node value
func GetSshRootLogin(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssh.root_login")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshRootLogin updates the ssh.root_login node value
func UpdateSshRootLogin(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.root_login", bytes.NewReader(byt))
	return
}

// GetSshStatus gets the ssh.status node value
func GetSshStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssh.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshStatus updates the ssh.status node value
func UpdateSshStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.status", bytes.NewReader(byt))
	return
}

// GetSslVpnAuthenticationAlgorithm gets the ssl_vpn.authentication_algorithm node value
func GetSslVpnAuthenticationAlgorithm(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.authentication_algorithm")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnAuthenticationAlgorithm updates the ssl_vpn.authentication_algorithm node value
func UpdateSslVpnAuthenticationAlgorithm(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.authentication_algorithm", bytes.NewReader(byt))
	return
}

// GetSslVpnCertificate gets the ssl_vpn.certificate node value
func GetSslVpnCertificate(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.certificate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnCertificate updates the ssl_vpn.certificate node value
func UpdateSslVpnCertificate(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.certificate", bytes.NewReader(byt))
	return
}

// GetSslVpnCompression gets the ssl_vpn.compression node value
func GetSslVpnCompression(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.compression")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnCompression updates the ssl_vpn.compression node value
func UpdateSslVpnCompression(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.compression", bytes.NewReader(byt))
	return
}

// GetSslVpnDatachannelKeyLifetime gets the ssl_vpn.datachannel_key_lifetime node value
func GetSslVpnDatachannelKeyLifetime(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.datachannel_key_lifetime")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnDatachannelKeyLifetime updates the ssl_vpn.datachannel_key_lifetime node value
func UpdateSslVpnDatachannelKeyLifetime(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.datachannel_key_lifetime", bytes.NewReader(byt))
	return
}

// GetSslVpnDebug gets the ssl_vpn.debug node value
func GetSslVpnDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnDebug updates the ssl_vpn.debug node value
func UpdateSslVpnDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.debug", bytes.NewReader(byt))
	return
}

// GetSslVpnDhKeySize gets the ssl_vpn.dh_key_size node value
func GetSslVpnDhKeySize(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.dh_key_size")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnDhKeySize updates the ssl_vpn.dh_key_size node value
func UpdateSslVpnDhKeySize(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.dh_key_size", bytes.NewReader(byt))
	return
}

// GetSslVpnDuplicateCn gets the ssl_vpn.duplicate_cn node value
func GetSslVpnDuplicateCn(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.duplicate_cn")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnDuplicateCn updates the ssl_vpn.duplicate_cn node value
func UpdateSslVpnDuplicateCn(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.duplicate_cn", bytes.NewReader(byt))
	return
}

// GetSslVpnEncryptionAlgorithm gets the ssl_vpn.encryption_algorithm node value
func GetSslVpnEncryptionAlgorithm(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.encryption_algorithm")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnEncryptionAlgorithm updates the ssl_vpn.encryption_algorithm node value
func UpdateSslVpnEncryptionAlgorithm(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.encryption_algorithm", bytes.NewReader(byt))
	return
}

// GetSslVpnHostname gets the ssl_vpn.hostname node value
func GetSslVpnHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnHostname updates the ssl_vpn.hostname node value
func UpdateSslVpnHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.hostname", bytes.NewReader(byt))
	return
}

// GetSslVpnInterface gets the ssl_vpn.interface node value
func GetSslVpnInterface(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.interface")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnInterface updates the ssl_vpn.interface node value
func UpdateSslVpnInterface(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.interface", bytes.NewReader(byt))
	return
}

// GetSslVpnInterfaceAddress gets the ssl_vpn.interface_address node value
func GetSslVpnInterfaceAddress(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.interface_address")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnInterfaceAddress updates the ssl_vpn.interface_address node value
func UpdateSslVpnInterfaceAddress(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.interface_address", bytes.NewReader(byt))
	return
}

// GetSslVpnIpAssignmentPool gets the ssl_vpn.ip_assignment_pool node value
func GetSslVpnIpAssignmentPool(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.ip_assignment_pool")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnIpAssignmentPool updates the ssl_vpn.ip_assignment_pool node value
func UpdateSslVpnIpAssignmentPool(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.ip_assignment_pool", bytes.NewReader(byt))
	return
}

// GetSslVpnPort gets the ssl_vpn.port node value
func GetSslVpnPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnPort updates the ssl_vpn.port node value
func UpdateSslVpnPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.port", bytes.NewReader(byt))
	return
}

// GetSslVpnProtocol gets the ssl_vpn.protocol node value
func GetSslVpnProtocol(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.protocol")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnProtocol updates the ssl_vpn.protocol node value
func UpdateSslVpnProtocol(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.protocol", bytes.NewReader(byt))
	return
}

// GetSslVpnUserAuthOptional gets the ssl_vpn.user_auth_optional node value
func GetSslVpnUserAuthOptional(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.user_auth_optional")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnUserAuthOptional updates the ssl_vpn.user_auth_optional node value
func UpdateSslVpnUserAuthOptional(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.user_auth_optional", bytes.NewReader(byt))
	return
}

// GetSupportAccessAccessId gets the support_access.access_id node value
func GetSupportAccessAccessId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/support_access.access_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSupportAccessAccessId updates the support_access.access_id node value
func UpdateSupportAccessAccessId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/support_access.access_id", bytes.NewReader(byt))
	return
}

// GetSupportAccessApuServer gets the support_access.apu_server node value
func GetSupportAccessApuServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/support_access.apu_server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSupportAccessApuServer updates the support_access.apu_server node value
func UpdateSupportAccessApuServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/support_access.apu_server", bytes.NewReader(byt))
	return
}

// GetSupportAccessLifetimeDuration gets the support_access.lifetime_duration node value
func GetSupportAccessLifetimeDuration(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/support_access.lifetime_duration")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSupportAccessLifetimeDuration updates the support_access.lifetime_duration node value
func UpdateSupportAccessLifetimeDuration(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/support_access.lifetime_duration", bytes.NewReader(byt))
	return
}

// GetSupportAccessLifetimeEnd gets the support_access.lifetime_end node value
func GetSupportAccessLifetimeEnd(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/support_access.lifetime_end")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSupportAccessLifetimeEnd updates the support_access.lifetime_end node value
func UpdateSupportAccessLifetimeEnd(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/support_access.lifetime_end", bytes.NewReader(byt))
	return
}

// GetSupportAccessSshKeys gets the support_access.ssh_keys node value
func GetSupportAccessSshKeys(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/support_access.ssh_keys")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSupportAccessSshKeys updates the support_access.ssh_keys node value
func UpdateSupportAccessSshKeys(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/support_access.ssh_keys", bytes.NewReader(byt))
	return
}

// GetSupportAccessStatus gets the support_access.status node value
func GetSupportAccessStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/support_access.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSupportAccessStatus updates the support_access.status node value
func UpdateSupportAccessStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/support_access.status", bytes.NewReader(byt))
	return
}

// GetU2DcacheAllowedNetworks gets the u2dcache.allowed_networks node value
func GetU2DcacheAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/u2dcache.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateU2DcacheAllowedNetworks updates the u2dcache.allowed_networks node value
func UpdateU2DcacheAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/u2dcache.allowed_networks", bytes.NewReader(byt))
	return
}

// GetU2DcachePort gets the u2dcache.port node value
func GetU2DcachePort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/u2dcache.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateU2DcachePort updates the u2dcache.port node value
func UpdateU2DcachePort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/u2dcache.port", bytes.NewReader(byt))
	return
}

// GetU2DcacheStatus gets the u2dcache.status node value
func GetU2DcacheStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/u2dcache.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateU2DcacheStatus updates the u2dcache.status node value
func UpdateU2DcacheStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/u2dcache.status", bytes.NewReader(byt))
	return
}

// GetUp2DateCacheHost gets the up2date.cache_host node value
func GetUp2DateCacheHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/up2date.cache_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateCacheHost updates the up2date.cache_host node value
func UpdateUp2DateCacheHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.cache_host", bytes.NewReader(byt))
	return
}

// GetUp2DateCachePort gets the up2date.cache_port node value
func GetUp2DateCachePort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/up2date.cache_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateCachePort updates the up2date.cache_port node value
func UpdateUp2DateCachePort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.cache_port", bytes.NewReader(byt))
	return
}

// GetUp2DateCacheStatus gets the up2date.cache_status node value
func GetUp2DateCacheStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.cache_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateCacheStatus updates the up2date.cache_status node value
func UpdateUp2DateCacheStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.cache_status", bytes.NewReader(byt))
	return
}

// GetUp2DateCacheUseAcc gets the up2date.cache_use_acc node value
func GetUp2DateCacheUseAcc(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.cache_use_acc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateCacheUseAcc updates the up2date.cache_use_acc node value
func UpdateUp2DateCacheUseAcc(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.cache_use_acc", bytes.NewReader(byt))
	return
}

// GetUp2DateParentProxyAuthPass gets the up2date.parent_proxy_auth_pass node value
func GetUp2DateParentProxyAuthPass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/up2date.parent_proxy_auth_pass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateParentProxyAuthPass updates the up2date.parent_proxy_auth_pass node value
func UpdateUp2DateParentProxyAuthPass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.parent_proxy_auth_pass", bytes.NewReader(byt))
	return
}

// GetUp2DateParentProxyAuthStatus gets the up2date.parent_proxy_auth_status node value
func GetUp2DateParentProxyAuthStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.parent_proxy_auth_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateParentProxyAuthStatus updates the up2date.parent_proxy_auth_status node value
func UpdateUp2DateParentProxyAuthStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.parent_proxy_auth_status", bytes.NewReader(byt))
	return
}

// GetUp2DateParentProxyAuthUser gets the up2date.parent_proxy_auth_user node value
func GetUp2DateParentProxyAuthUser(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/up2date.parent_proxy_auth_user")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateParentProxyAuthUser updates the up2date.parent_proxy_auth_user node value
func UpdateUp2DateParentProxyAuthUser(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.parent_proxy_auth_user", bytes.NewReader(byt))
	return
}

// GetUp2DateParentProxyHost gets the up2date.parent_proxy_host node value
func GetUp2DateParentProxyHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/up2date.parent_proxy_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateParentProxyHost updates the up2date.parent_proxy_host node value
func UpdateUp2DateParentProxyHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.parent_proxy_host", bytes.NewReader(byt))
	return
}

// GetUp2DateParentProxyPort gets the up2date.parent_proxy_port node value
func GetUp2DateParentProxyPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/up2date.parent_proxy_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateParentProxyPort updates the up2date.parent_proxy_port node value
func UpdateUp2DateParentProxyPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.parent_proxy_port", bytes.NewReader(byt))
	return
}

// GetUp2DateParentProxyStatus gets the up2date.parent_proxy_status node value
func GetUp2DateParentProxyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.parent_proxy_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateParentProxyStatus updates the up2date.parent_proxy_status node value
func UpdateUp2DateParentProxyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.parent_proxy_status", bytes.NewReader(byt))
	return
}

// GetUp2DatePatternDownloadInterval gets the up2date.pattern_download_interval node value
func GetUp2DatePatternDownloadInterval(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/up2date.pattern_download_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DatePatternDownloadInterval updates the up2date.pattern_download_interval node value
func UpdateUp2DatePatternDownloadInterval(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.pattern_download_interval", bytes.NewReader(byt))
	return
}

// GetUp2DatePatternDownloadStatus gets the up2date.pattern_download_status node value
func GetUp2DatePatternDownloadStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.pattern_download_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DatePatternDownloadStatus updates the up2date.pattern_download_status node value
func UpdateUp2DatePatternDownloadStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.pattern_download_status", bytes.NewReader(byt))
	return
}

// GetUp2DateScheduledUp2Date gets the up2date.scheduled_up2date node value
func GetUp2DateScheduledUp2Date(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/up2date.scheduled_up2date")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateScheduledUp2Date updates the up2date.scheduled_up2date node value
func UpdateUp2DateScheduledUp2Date(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.scheduled_up2date", bytes.NewReader(byt))
	return
}

// GetUp2DateServers gets the up2date.servers node value
func GetUp2DateServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/up2date.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateServers updates the up2date.servers node value
func UpdateUp2DateServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.servers", bytes.NewReader(byt))
	return
}

// GetUp2DateStatus gets the up2date.status node value
func GetUp2DateStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateStatus updates the up2date.status node value
func UpdateUp2DateStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.status", bytes.NewReader(byt))
	return
}

// GetUp2DateSystemAutoinstallTime gets the up2date.system_autoinstall_time node value
func GetUp2DateSystemAutoinstallTime(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/up2date.system_autoinstall_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateSystemAutoinstallTime updates the up2date.system_autoinstall_time node value
func UpdateUp2DateSystemAutoinstallTime(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.system_autoinstall_time", bytes.NewReader(byt))
	return
}

// GetUp2DateSystemDownloadInterval gets the up2date.system_download_interval node value
func GetUp2DateSystemDownloadInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/up2date.system_download_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateSystemDownloadInterval updates the up2date.system_download_interval node value
func UpdateUp2DateSystemDownloadInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.system_download_interval", bytes.NewReader(byt))
	return
}

// GetUp2DateSystemDownloadStatus gets the up2date.system_download_status node value
func GetUp2DateSystemDownloadStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.system_download_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateSystemDownloadStatus updates the up2date.system_download_status node value
func UpdateUp2DateSystemDownloadStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.system_download_status", bytes.NewReader(byt))
	return
}

// GetUplinkActions gets the uplink.actions node value
func GetUplinkActions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/uplink.actions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkActions updates the uplink.actions node value
func UpdateUplinkActions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.actions", bytes.NewReader(byt))
	return
}

// GetUplinkActive gets the uplink.active node value
func GetUplinkActive(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/uplink.active")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkActive updates the uplink.active node value
func UpdateUplinkActive(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.active", bytes.NewReader(byt))
	return
}

// GetUplinkCondition gets the uplink.condition node value
func GetUplinkCondition(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/uplink.condition")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkCondition updates the uplink.condition node value
func UpdateUplinkCondition(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.condition", bytes.NewReader(byt))
	return
}

// GetUplinkInterfaces gets the uplink.interfaces node value
func GetUplinkInterfaces(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/uplink.interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkInterfaces updates the uplink.interfaces node value
func UpdateUplinkInterfaces(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.interfaces", bytes.NewReader(byt))
	return
}

// GetUplinkMonitoring gets the uplink.monitoring node value
func GetUplinkMonitoring(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/uplink.monitoring")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkMonitoring updates the uplink.monitoring node value
func UpdateUplinkMonitoring(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.monitoring", bytes.NewReader(byt))
	return
}

// GetUplinkPassive gets the uplink.passive node value
func GetUplinkPassive(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/uplink.passive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkPassive updates the uplink.passive node value
func UpdateUplinkPassive(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.passive", bytes.NewReader(byt))
	return
}

// GetUplinkPrimary gets the uplink.primary node value
func GetUplinkPrimary(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/uplink.primary")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkPrimary updates the uplink.primary node value
func UpdateUplinkPrimary(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.primary", bytes.NewReader(byt))
	return
}

// GetUplinkRules gets the uplink.rules node value
func GetUplinkRules(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/uplink.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkRules updates the uplink.rules node value
func UpdateUplinkRules(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.rules", bytes.NewReader(byt))
	return
}

// GetUplinkScheduler gets the uplink.scheduler node value
func GetUplinkScheduler(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/uplink.scheduler")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkScheduler updates the uplink.scheduler node value
func UpdateUplinkScheduler(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.scheduler", bytes.NewReader(byt))
	return
}

// GetUplinkStatus gets the uplink.status node value
func GetUplinkStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/uplink.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkStatus updates the uplink.status node value
func UpdateUplinkStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.status", bytes.NewReader(byt))
	return
}

// GetWebadminAllowedNetworks gets the webadmin.allowed_networks node value
func GetWebadminAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/webadmin.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminAllowedNetworks updates the webadmin.allowed_networks node value
func UpdateWebadminAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.allowed_networks", bytes.NewReader(byt))
	return
}

// GetWebadminCa gets the webadmin.ca node value
func GetWebadminCa(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.ca")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminCa updates the webadmin.ca node value
func UpdateWebadminCa(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.ca", bytes.NewReader(byt))
	return
}

// GetWebadminCert gets the webadmin.cert node value
func GetWebadminCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminCert updates the webadmin.cert node value
func UpdateWebadminCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.cert", bytes.NewReader(byt))
	return
}

// GetWebadminDashboardRefresh gets the webadmin.dashboard_refresh node value
func GetWebadminDashboardRefresh(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.dashboard_refresh")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminDashboardRefresh updates the webadmin.dashboard_refresh node value
func UpdateWebadminDashboardRefresh(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.dashboard_refresh", bytes.NewReader(byt))
	return
}

// GetWebadminLanguage gets the webadmin.language node value
func GetWebadminLanguage(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.language")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminLanguage updates the webadmin.language node value
func UpdateWebadminLanguage(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.language", bytes.NewReader(byt))
	return
}

// GetWebadminLogAdminConnect gets the webadmin.log_admin_connect node value
func GetWebadminLogAdminConnect(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/webadmin.log_admin_connect")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminLogAdminConnect updates the webadmin.log_admin_connect node value
func UpdateWebadminLogAdminConnect(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.log_admin_connect", bytes.NewReader(byt))
	return
}

// GetWebadminOfferWizard gets the webadmin.offer_wizard node value
func GetWebadminOfferWizard(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/webadmin.offer_wizard")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminOfferWizard updates the webadmin.offer_wizard node value
func UpdateWebadminOfferWizard(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.offer_wizard", bytes.NewReader(byt))
	return
}

// GetWebadminPort gets the webadmin.port node value
func GetWebadminPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/webadmin.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminPort updates the webadmin.port node value
func UpdateWebadminPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.port", bytes.NewReader(byt))
	return
}

// GetWebadminRestApi gets the webadmin.rest_api node value
func GetWebadminRestApi(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/webadmin.rest_api")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminRestApi updates the webadmin.rest_api node value
func UpdateWebadminRestApi(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.rest_api", bytes.NewReader(byt))
	return
}

// GetWebadminTermsOfUseStatus gets the webadmin.terms_of_use.status node value
func GetWebadminTermsOfUseStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/webadmin.terms_of_use.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminTermsOfUseStatus updates the webadmin.terms_of_use.status node value
func UpdateWebadminTermsOfUseStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.terms_of_use.status", bytes.NewReader(byt))
	return
}

// GetWebadminTermsOfUseText gets the webadmin.terms_of_use.text node value
func GetWebadminTermsOfUseText(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.terms_of_use.text")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminTermsOfUseText updates the webadmin.terms_of_use.text node value
func UpdateWebadminTermsOfUseText(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.terms_of_use.text", bytes.NewReader(byt))
	return
}

// GetWebadminTimeout gets the webadmin.timeout node value
func GetWebadminTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/webadmin.timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminTimeout updates the webadmin.timeout node value
func UpdateWebadminTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.timeout", bytes.NewReader(byt))
	return
}

// GetWebadminTimeoutOnDashboard gets the webadmin.timeout_on_dashboard node value
func GetWebadminTimeoutOnDashboard(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/webadmin.timeout_on_dashboard")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminTimeoutOnDashboard updates the webadmin.timeout_on_dashboard node value
func UpdateWebadminTimeoutOnDashboard(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.timeout_on_dashboard", bytes.NewReader(byt))
	return
}

// GetWebadminTlsCiphers gets the webadmin.tls_ciphers node value
func GetWebadminTlsCiphers(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.tls_ciphers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminTlsCiphers updates the webadmin.tls_ciphers node value
func UpdateWebadminTlsCiphers(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.tls_ciphers", bytes.NewReader(byt))
	return
}

// GetWebadminTlsProtocols gets the webadmin.tls_protocols node value
func GetWebadminTlsProtocols(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.tls_protocols")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminTlsProtocols updates the webadmin.tls_protocols node value
func UpdateWebadminTlsProtocols(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.tls_protocols", bytes.NewReader(byt))
	return
}
