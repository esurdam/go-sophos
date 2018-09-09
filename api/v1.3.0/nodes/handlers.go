package nodes

import (
	"bytes"
	"encoding/json"

	"github.com/esurdam/go-sophos"
)

// GetAccServer1AuthSecret gets the acc.server1.auth.secret value from the UTM
func GetAccServer1AuthSecret(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/acc.server1.auth.secret")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer1AuthSecret PUTs the acc.server1.auth.secret value to the UTM
func UpdateAccServer1AuthSecret(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server1.auth.secret", bytes.NewReader(byt))
	return
}

// GetAccServer1AuthStatus gets the acc.server1.auth.status value from the UTM
func GetAccServer1AuthStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/acc.server1.auth.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer1AuthStatus PUTs the acc.server1.auth.status value to the UTM
func UpdateAccServer1AuthStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server1.auth.status", bytes.NewReader(byt))
	return
}

// GetAccServer1Port gets the acc.server1.port value from the UTM
func GetAccServer1Port(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/acc.server1.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer1Port PUTs the acc.server1.port value to the UTM
func UpdateAccServer1Port(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server1.port", bytes.NewReader(byt))
	return
}

// GetAccServer1Roles gets the acc.server1.roles value from the UTM
func GetAccServer1Roles(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/acc.server1.roles")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer1Roles PUTs the acc.server1.roles value to the UTM
func UpdateAccServer1Roles(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server1.roles", bytes.NewReader(byt))
	return
}

// GetAccServer1Server gets the acc.server1.server value from the UTM
func GetAccServer1Server(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/acc.server1.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer1Server PUTs the acc.server1.server value to the UTM
func UpdateAccServer1Server(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server1.server", bytes.NewReader(byt))
	return
}

// GetAccServer2AuthSecret gets the acc.server2.auth.secret value from the UTM
func GetAccServer2AuthSecret(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/acc.server2.auth.secret")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer2AuthSecret PUTs the acc.server2.auth.secret value to the UTM
func UpdateAccServer2AuthSecret(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server2.auth.secret", bytes.NewReader(byt))
	return
}

// GetAccServer2AuthStatus gets the acc.server2.auth.status value from the UTM
func GetAccServer2AuthStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/acc.server2.auth.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer2AuthStatus PUTs the acc.server2.auth.status value to the UTM
func UpdateAccServer2AuthStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server2.auth.status", bytes.NewReader(byt))
	return
}

// GetAccServer2Port gets the acc.server2.port value from the UTM
func GetAccServer2Port(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/acc.server2.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer2Port PUTs the acc.server2.port value to the UTM
func UpdateAccServer2Port(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server2.port", bytes.NewReader(byt))
	return
}

// GetAccServer2Roles gets the acc.server2.roles value from the UTM
func GetAccServer2Roles(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/acc.server2.roles")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer2Roles PUTs the acc.server2.roles value to the UTM
func UpdateAccServer2Roles(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server2.roles", bytes.NewReader(byt))
	return
}

// GetAccServer2Server gets the acc.server2.server value from the UTM
func GetAccServer2Server(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/acc.server2.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer2Server PUTs the acc.server2.server value to the UTM
func UpdateAccServer2Server(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server2.server", bytes.NewReader(byt))
	return
}

// GetAccServer2Status gets the acc.server2.status value from the UTM
func GetAccServer2Status(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/acc.server2.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccServer2Status PUTs the acc.server2.status value to the UTM
func UpdateAccServer2Status(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.server2.status", bytes.NewReader(byt))
	return
}

// GetAccSsoAdminGroup gets the acc.sso_admin_group value from the UTM
func GetAccSsoAdminGroup(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/acc.sso_admin_group")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccSsoAdminGroup PUTs the acc.sso_admin_group value to the UTM
func UpdateAccSsoAdminGroup(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.sso_admin_group", bytes.NewReader(byt))
	return
}

// GetAccSsoAuditorGroup gets the acc.sso_auditor_group value from the UTM
func GetAccSsoAuditorGroup(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/acc.sso_auditor_group")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccSsoAuditorGroup PUTs the acc.sso_auditor_group value to the UTM
func UpdateAccSsoAuditorGroup(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.sso_auditor_group", bytes.NewReader(byt))
	return
}

// GetAccStatus gets the acc.status value from the UTM
func GetAccStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/acc.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccStatus PUTs the acc.status value to the UTM
func UpdateAccStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/acc.status", bytes.NewReader(byt))
	return
}

// GetAccdAccessAllowedAdmins gets the accd.access.allowed_admins value from the UTM
func GetAccdAccessAllowedAdmins(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/accd.access.allowed_admins")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdAccessAllowedAdmins PUTs the accd.access.allowed_admins value to the UTM
func UpdateAccdAccessAllowedAdmins(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.access.allowed_admins", bytes.NewReader(byt))
	return
}

// GetAccdAccessAllowedNetworks gets the accd.access.allowed_networks value from the UTM
func GetAccdAccessAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/accd.access.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdAccessAllowedNetworks PUTs the accd.access.allowed_networks value to the UTM
func UpdateAccdAccessAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.access.allowed_networks", bytes.NewReader(byt))
	return
}

// GetAccdAccessAllowedUsers gets the accd.access.allowed_users value from the UTM
func GetAccdAccessAllowedUsers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/accd.access.allowed_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdAccessAllowedUsers PUTs the accd.access.allowed_users value to the UTM
func UpdateAccdAccessAllowedUsers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.access.allowed_users", bytes.NewReader(byt))
	return
}

// GetAccdAccessCert gets the accd.access.cert value from the UTM
func GetAccdAccessCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/accd.access.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdAccessCert PUTs the accd.access.cert value to the UTM
func UpdateAccdAccessCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.access.cert", bytes.NewReader(byt))
	return
}

// GetAccdAccessPort gets the accd.access.port value from the UTM
func GetAccdAccessPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/accd.access.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdAccessPort PUTs the accd.access.port value to the UTM
func UpdateAccdAccessPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.access.port", bytes.NewReader(byt))
	return
}

// GetAccdDevicesAllowedNetworks gets the accd.devices.allowed_networks value from the UTM
func GetAccdDevicesAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/accd.devices.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdDevicesAllowedNetworks PUTs the accd.devices.allowed_networks value to the UTM
func UpdateAccdDevicesAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.devices.allowed_networks", bytes.NewReader(byt))
	return
}

// GetAccdDevicesAuthAuto gets the accd.devices.auth.auto value from the UTM
func GetAccdDevicesAuthAuto(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/accd.devices.auth.auto")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdDevicesAuthAuto PUTs the accd.devices.auth.auto value to the UTM
func UpdateAccdDevicesAuthAuto(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.devices.auth.auto", bytes.NewReader(byt))
	return
}

// GetAccdDevicesAuthSecret gets the accd.devices.auth.secret value from the UTM
func GetAccdDevicesAuthSecret(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/accd.devices.auth.secret")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdDevicesAuthSecret PUTs the accd.devices.auth.secret value to the UTM
func UpdateAccdDevicesAuthSecret(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.devices.auth.secret", bytes.NewReader(byt))
	return
}

// GetAccdDevicesAuthStatus gets the accd.devices.auth.status value from the UTM
func GetAccdDevicesAuthStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/accd.devices.auth.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdDevicesAuthStatus PUTs the accd.devices.auth.status value to the UTM
func UpdateAccdDevicesAuthStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.devices.auth.status", bytes.NewReader(byt))
	return
}

// GetAccdDevicesCert gets the accd.devices.cert value from the UTM
func GetAccdDevicesCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/accd.devices.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdDevicesCert PUTs the accd.devices.cert value to the UTM
func UpdateAccdDevicesCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.devices.cert", bytes.NewReader(byt))
	return
}

// GetAccdDevicesPort gets the accd.devices.port value from the UTM
func GetAccdDevicesPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/accd.devices.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdDevicesPort PUTs the accd.devices.port value to the UTM
func UpdateAccdDevicesPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.devices.port", bytes.NewReader(byt))
	return
}

// GetAccdGeneralAllowedNetworks gets the accd.general.allowed_networks value from the UTM
func GetAccdGeneralAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/accd.general.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdGeneralAllowedNetworks PUTs the accd.general.allowed_networks value to the UTM
func UpdateAccdGeneralAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.general.allowed_networks", bytes.NewReader(byt))
	return
}

// GetAccdGeneralCert gets the accd.general.cert value from the UTM
func GetAccdGeneralCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/accd.general.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdGeneralCert PUTs the accd.general.cert value to the UTM
func UpdateAccdGeneralCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.general.cert", bytes.NewReader(byt))
	return
}

// GetAccdGeneralLanguage gets the accd.general.language value from the UTM
func GetAccdGeneralLanguage(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/accd.general.language")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdGeneralLanguage PUTs the accd.general.language value to the UTM
func UpdateAccdGeneralLanguage(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.general.language", bytes.NewReader(byt))
	return
}

// GetAccdGeneralPort gets the accd.general.port value from the UTM
func GetAccdGeneralPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/accd.general.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdGeneralPort PUTs the accd.general.port value to the UTM
func UpdateAccdGeneralPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.general.port", bytes.NewReader(byt))
	return
}

// GetAccdGeneralTimeout gets the accd.general.timeout value from the UTM
func GetAccdGeneralTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/accd.general.timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccdGeneralTimeout PUTs the accd.general.timeout value to the UTM
func UpdateAccdGeneralTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accd.general.timeout", bytes.NewReader(byt))
	return
}

// GetAccountingIpfixConnections gets the accounting.ipfix.connections value from the UTM
func GetAccountingIpfixConnections(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/accounting.ipfix.connections")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccountingIpfixConnections PUTs the accounting.ipfix.connections value to the UTM
func UpdateAccountingIpfixConnections(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accounting.ipfix.connections", bytes.NewReader(byt))
	return
}

// GetAccountingIpfixStatus gets the accounting.ipfix.status value from the UTM
func GetAccountingIpfixStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/accounting.ipfix.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAccountingIpfixStatus PUTs the accounting.ipfix.status value to the UTM
func UpdateAccountingIpfixStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/accounting.ipfix.status", bytes.NewReader(byt))
	return
}

// GetAfcControlledNetworks gets the afc.controlled_networks value from the UTM
func GetAfcControlledNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/afc.controlled_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcControlledNetworks PUTs the afc.controlled_networks value to the UTM
func UpdateAfcControlledNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.controlled_networks", bytes.NewReader(byt))
	return
}

// GetAfcHiddenSkip gets the afc.hidden_skip value from the UTM
func GetAfcHiddenSkip(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/afc.hidden_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcHiddenSkip PUTs the afc.hidden_skip value to the UTM
func UpdateAfcHiddenSkip(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.hidden_skip", bytes.NewReader(byt))
	return
}

// GetAfcHttpRedirectUrl gets the afc.http_redirect_url value from the UTM
func GetAfcHttpRedirectUrl(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/afc.http_redirect_url")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcHttpRedirectUrl PUTs the afc.http_redirect_url value to the UTM
func UpdateAfcHttpRedirectUrl(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.http_redirect_url", bytes.NewReader(byt))
	return
}

// GetAfcLog gets the afc.log value from the UTM
func GetAfcLog(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/afc.log")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcLog PUTs the afc.log value to the UTM
func UpdateAfcLog(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.log", bytes.NewReader(byt))
	return
}

// GetAfcNfqueueLength gets the afc.nfqueue_length value from the UTM
func GetAfcNfqueueLength(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/afc.nfqueue_length")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcNfqueueLength PUTs the afc.nfqueue_length value to the UTM
func UpdateAfcNfqueueLength(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.nfqueue_length", bytes.NewReader(byt))
	return
}

// GetAfcNumQueues gets the afc.num_queues value from the UTM
func GetAfcNumQueues(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/afc.num_queues")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcNumQueues PUTs the afc.num_queues value to the UTM
func UpdateAfcNumQueues(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.num_queues", bytes.NewReader(byt))
	return
}

// GetAfcRules gets the afc.rules value from the UTM
func GetAfcRules(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/afc.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcRules PUTs the afc.rules value to the UTM
func UpdateAfcRules(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.rules", bytes.NewReader(byt))
	return
}

// GetAfcStatus gets the afc.status value from the UTM
func GetAfcStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/afc.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcStatus PUTs the afc.status value to the UTM
func UpdateAfcStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.status", bytes.NewReader(byt))
	return
}

// GetAfcSubmitUnknownTrafficData gets the afc.submit_unknown_traffic_data value from the UTM
func GetAfcSubmitUnknownTrafficData(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/afc.submit_unknown_traffic_data")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcSubmitUnknownTrafficData PUTs the afc.submit_unknown_traffic_data value to the UTM
func UpdateAfcSubmitUnknownTrafficData(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.submit_unknown_traffic_data", bytes.NewReader(byt))
	return
}

// GetAfcTransparentSkip gets the afc.transparent_skip value from the UTM
func GetAfcTransparentSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/afc.transparent_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAfcTransparentSkip PUTs the afc.transparent_skip value to the UTM
func UpdateAfcTransparentSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/afc.transparent_skip", bytes.NewReader(byt))
	return
}

// GetAmazonVpcAutoPfrule gets the amazon_vpc.auto_pfrule value from the UTM
func GetAmazonVpcAutoPfrule(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/amazon_vpc.auto_pfrule")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAmazonVpcAutoPfrule PUTs the amazon_vpc.auto_pfrule value to the UTM
func UpdateAmazonVpcAutoPfrule(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/amazon_vpc.auto_pfrule", bytes.NewReader(byt))
	return
}

// GetAmazonVpcConnections gets the amazon_vpc.connections value from the UTM
func GetAmazonVpcConnections(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/amazon_vpc.connections")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAmazonVpcConnections PUTs the amazon_vpc.connections value to the UTM
func UpdateAmazonVpcConnections(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/amazon_vpc.connections", bytes.NewReader(byt))
	return
}

// GetAmazonVpcNetworks gets the amazon_vpc.networks value from the UTM
func GetAmazonVpcNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/amazon_vpc.networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAmazonVpcNetworks PUTs the amazon_vpc.networks value to the UTM
func UpdateAmazonVpcNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/amazon_vpc.networks", bytes.NewReader(byt))
	return
}

// GetAmazonVpcStatus gets the amazon_vpc.status value from the UTM
func GetAmazonVpcStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/amazon_vpc.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAmazonVpcStatus PUTs the amazon_vpc.status value to the UTM
func UpdateAmazonVpcStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/amazon_vpc.status", bytes.NewReader(byt))
	return
}

// GetAptpPolicy gets the aptp.policy value from the UTM
func GetAptpPolicy(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/aptp.policy")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAptpPolicy PUTs the aptp.policy value to the UTM
func UpdateAptpPolicy(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/aptp.policy", bytes.NewReader(byt))
	return
}

// GetAptpRuleModifiers gets the aptp.rule_modifiers value from the UTM
func GetAptpRuleModifiers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/aptp.rule_modifiers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAptpRuleModifiers PUTs the aptp.rule_modifiers value to the UTM
func UpdateAptpRuleModifiers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/aptp.rule_modifiers", bytes.NewReader(byt))
	return
}

// GetAptpStatus gets the aptp.status value from the UTM
func GetAptpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/aptp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAptpStatus PUTs the aptp.status value to the UTM
func UpdateAptpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/aptp.status", bytes.NewReader(byt))
	return
}

// GetAptpTransparentSkip gets the aptp.transparent_skip value from the UTM
func GetAptpTransparentSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/aptp.transparent_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAptpTransparentSkip PUTs the aptp.transparent_skip value to the UTM
func UpdateAptpTransparentSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/aptp.transparent_skip", bytes.NewReader(byt))
	return
}

// GetArmLicensedIp gets the arm.licensed_ip value from the UTM
func GetArmLicensedIp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.licensed_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmLicensedIp PUTs the arm.licensed_ip value to the UTM
func UpdateArmLicensedIp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.licensed_ip", bytes.NewReader(byt))
	return
}

// GetArmRemoteHost gets the arm.remote.host value from the UTM
func GetArmRemoteHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.remote.host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteHost PUTs the arm.remote.host value to the UTM
func UpdateArmRemoteHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.host", bytes.NewReader(byt))
	return
}

// GetArmRemoteMethod gets the arm.remote.method value from the UTM
func GetArmRemoteMethod(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.remote.method")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteMethod PUTs the arm.remote.method value to the UTM
func UpdateArmRemoteMethod(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.method", bytes.NewReader(byt))
	return
}

// GetArmRemoteSmbPassword gets the arm.remote.smb_password value from the UTM
func GetArmRemoteSmbPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.remote.smb_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteSmbPassword PUTs the arm.remote.smb_password value to the UTM
func UpdateArmRemoteSmbPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.smb_password", bytes.NewReader(byt))
	return
}

// GetArmRemoteSmbShare gets the arm.remote.smb_share value from the UTM
func GetArmRemoteSmbShare(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.remote.smb_share")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteSmbShare PUTs the arm.remote.smb_share value to the UTM
func UpdateArmRemoteSmbShare(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.smb_share", bytes.NewReader(byt))
	return
}

// GetArmRemoteSmbUser gets the arm.remote.smb_user value from the UTM
func GetArmRemoteSmbUser(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.remote.smb_user")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteSmbUser PUTs the arm.remote.smb_user value to the UTM
func UpdateArmRemoteSmbUser(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.smb_user", bytes.NewReader(byt))
	return
}

// GetArmRemoteStatus gets the arm.remote.status value from the UTM
func GetArmRemoteStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/arm.remote.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteStatus PUTs the arm.remote.status value to the UTM
func UpdateArmRemoteStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.status", bytes.NewReader(byt))
	return
}

// GetArmRemoteSyslogService gets the arm.remote.syslog_service value from the UTM
func GetArmRemoteSyslogService(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/arm.remote.syslog_service")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmRemoteSyslogService PUTs the arm.remote.syslog_service value to the UTM
func UpdateArmRemoteSyslogService(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.remote.syslog_service", bytes.NewReader(byt))
	return
}

// GetArmStatus gets the arm.status value from the UTM
func GetArmStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/arm.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateArmStatus PUTs the arm.status value to the UTM
func UpdateArmStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/arm.status", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoForceUtf8Sync gets the auth.ad_sso.force_utf8_sync value from the UTM
func GetAuthAdSsoForceUtf8Sync(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.force_utf8_sync")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoForceUtf8Sync PUTs the auth.ad_sso.force_utf8_sync value to the UTM
func UpdateAuthAdSsoForceUtf8Sync(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.force_utf8_sync", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoJoinresult gets the auth.ad_sso.joinresult value from the UTM
func GetAuthAdSsoJoinresult(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.joinresult")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoJoinresult PUTs the auth.ad_sso.joinresult value to the UTM
func UpdateAuthAdSsoJoinresult(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.joinresult", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoLoadbalancerFqdn gets the auth.ad_sso.loadbalancer_fqdn value from the UTM
func GetAuthAdSsoLoadbalancerFqdn(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.loadbalancer_fqdn")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoLoadbalancerFqdn PUTs the auth.ad_sso.loadbalancer_fqdn value to the UTM
func UpdateAuthAdSsoLoadbalancerFqdn(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.loadbalancer_fqdn", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoNtlmv2Auth gets the auth.ad_sso.ntlmv2_auth value from the UTM
func GetAuthAdSsoNtlmv2Auth(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.ntlmv2_auth")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoNtlmv2Auth PUTs the auth.ad_sso.ntlmv2_auth value to the UTM
func UpdateAuthAdSsoNtlmv2Auth(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.ntlmv2_auth", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSecrets gets the auth.ad_sso.secrets value from the UTM
func GetAuthAdSsoSecrets(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.secrets")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSecrets PUTs the auth.ad_sso.secrets value to the UTM
func UpdateAuthAdSsoSecrets(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.secrets", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSmbconf gets the auth.ad_sso.smbconf value from the UTM
func GetAuthAdSsoSmbconf(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.smbconf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSmbconf PUTs the auth.ad_sso.smbconf value to the UTM
func UpdateAuthAdSsoSmbconf(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.smbconf", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoDomain gets the auth.ad_sso.sso_domain value from the UTM
func GetAuthAdSsoSsoDomain(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_domain")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoDomain PUTs the auth.ad_sso.sso_domain value to the UTM
func UpdateAuthAdSsoSsoDomain(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_domain", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoNetbiosDomain gets the auth.ad_sso.sso_netbios_domain value from the UTM
func GetAuthAdSsoSsoNetbiosDomain(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_netbios_domain")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoNetbiosDomain PUTs the auth.ad_sso.sso_netbios_domain value to the UTM
func UpdateAuthAdSsoSsoNetbiosDomain(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_netbios_domain", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoNetbiosHost gets the auth.ad_sso.sso_netbios_host value from the UTM
func GetAuthAdSsoSsoNetbiosHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_netbios_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoNetbiosHost PUTs the auth.ad_sso.sso_netbios_host value to the UTM
func UpdateAuthAdSsoSsoNetbiosHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_netbios_host", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoPassword gets the auth.ad_sso.sso_password value from the UTM
func GetAuthAdSsoSsoPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoPassword PUTs the auth.ad_sso.sso_password value to the UTM
func UpdateAuthAdSsoSsoPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_password", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoServer gets the auth.ad_sso.sso_server value from the UTM
func GetAuthAdSsoSsoServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoServer PUTs the auth.ad_sso.sso_server value to the UTM
func UpdateAuthAdSsoSsoServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_server", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoStatus gets the auth.ad_sso.sso_status value from the UTM
func GetAuthAdSsoSsoStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoStatus PUTs the auth.ad_sso.sso_status value to the UTM
func UpdateAuthAdSsoSsoStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_status", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoSync gets the auth.ad_sso.sso_sync value from the UTM
func GetAuthAdSsoSsoSync(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_sync")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoSync PUTs the auth.ad_sso.sso_sync value to the UTM
func UpdateAuthAdSsoSsoSync(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_sync", bytes.NewReader(byt))
	return
}

// GetAuthAdSsoSsoUsername gets the auth.ad_sso.sso_username value from the UTM
func GetAuthAdSsoSsoUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.ad_sso.sso_username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAdSsoSsoUsername PUTs the auth.ad_sso.sso_username value to the UTM
func UpdateAuthAdSsoSsoUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.ad_sso.sso_username", bytes.NewReader(byt))
	return
}

// GetAuthApiTokens gets the auth.api_tokens value from the UTM
func GetAuthApiTokens(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/auth.api_tokens")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthApiTokens PUTs the auth.api_tokens value to the UTM
func UpdateAuthApiTokens(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.api_tokens", bytes.NewReader(byt))
	return
}

// GetAuthAutoAddToFacility gets the auth.auto_add_to_facility value from the UTM
func GetAuthAutoAddToFacility(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/auth.auto_add_to_facility")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAutoAddToFacility PUTs the auth.auto_add_to_facility value to the UTM
func UpdateAuthAutoAddToFacility(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.auto_add_to_facility", bytes.NewReader(byt))
	return
}

// GetAuthAutoAddUsers gets the auth.auto_add_users value from the UTM
func GetAuthAutoAddUsers(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.auto_add_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthAutoAddUsers PUTs the auth.auto_add_users value to the UTM
func UpdateAuthAutoAddUsers(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.auto_add_users", bytes.NewReader(byt))
	return
}

// GetAuthBlockAttempts gets the auth.block.attempts value from the UTM
func GetAuthBlockAttempts(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.block.attempts")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthBlockAttempts PUTs the auth.block.attempts value to the UTM
func UpdateAuthBlockAttempts(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.block.attempts", bytes.NewReader(byt))
	return
}

// GetAuthBlockFacilities gets the auth.block.facilities value from the UTM
func GetAuthBlockFacilities(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/auth.block.facilities")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthBlockFacilities PUTs the auth.block.facilities value to the UTM
func UpdateAuthBlockFacilities(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.block.facilities", bytes.NewReader(byt))
	return
}

// GetAuthBlockLockout gets the auth.block.lockout value from the UTM
func GetAuthBlockLockout(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.block.lockout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthBlockLockout PUTs the auth.block.lockout value to the UTM
func UpdateAuthBlockLockout(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.block.lockout", bytes.NewReader(byt))
	return
}

// GetAuthBlockNever gets the auth.block.never value from the UTM
func GetAuthBlockNever(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/auth.block.never")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthBlockNever PUTs the auth.block.never value to the UTM
func UpdateAuthBlockNever(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.block.never", bytes.NewReader(byt))
	return
}

// GetAuthBlockSeconds gets the auth.block.seconds value from the UTM
func GetAuthBlockSeconds(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.block.seconds")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthBlockSeconds PUTs the auth.block.seconds value to the UTM
func UpdateAuthBlockSeconds(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.block.seconds", bytes.NewReader(byt))
	return
}

// GetAuthCacheLifetime gets the auth.cache_lifetime value from the UTM
func GetAuthCacheLifetime(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.cache_lifetime")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthCacheLifetime PUTs the auth.cache_lifetime value to the UTM
func UpdateAuthCacheLifetime(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.cache_lifetime", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoEmConflict gets the auth.edir_sso.em_conflict value from the UTM
func GetAuthEdirSsoEmConflict(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.em_conflict")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoEmConflict PUTs the auth.edir_sso.em_conflict value to the UTM
func UpdateAuthEdirSsoEmConflict(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.em_conflict", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoEmSocketTimeout gets the auth.edir_sso.em_socket_timeout value from the UTM
func GetAuthEdirSsoEmSocketTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.em_socket_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoEmSocketTimeout PUTs the auth.edir_sso.em_socket_timeout value to the UTM
func UpdateAuthEdirSsoEmSocketTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.em_socket_timeout", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoEmVerifyLogout gets the auth.edir_sso.em_verify_logout value from the UTM
func GetAuthEdirSsoEmVerifyLogout(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.em_verify_logout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoEmVerifyLogout PUTs the auth.edir_sso.em_verify_logout value to the UTM
func UpdateAuthEdirSsoEmVerifyLogout(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.em_verify_logout", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoSsoAuaSearchIp gets the auth.edir_sso.sso_aua_search_ip value from the UTM
func GetAuthEdirSsoSsoAuaSearchIp(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.sso_aua_search_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoSsoAuaSearchIp PUTs the auth.edir_sso.sso_aua_search_ip value to the UTM
func UpdateAuthEdirSsoSsoAuaSearchIp(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.sso_aua_search_ip", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoSsoMode gets the auth.edir_sso.sso_mode value from the UTM
func GetAuthEdirSsoSsoMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.sso_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoSsoMode PUTs the auth.edir_sso.sso_mode value to the UTM
func UpdateAuthEdirSsoSsoMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.sso_mode", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoSsoServer gets the auth.edir_sso.sso_server value from the UTM
func GetAuthEdirSsoSsoServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.sso_server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoSsoServer PUTs the auth.edir_sso.sso_server value to the UTM
func UpdateAuthEdirSsoSsoServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.sso_server", bytes.NewReader(byt))
	return
}

// GetAuthEdirSsoSyncInterval gets the auth.edir_sso.sync_interval value from the UTM
func GetAuthEdirSsoSyncInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.edir_sso.sync_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthEdirSsoSyncInterval PUTs the auth.edir_sso.sync_interval value to the UTM
func UpdateAuthEdirSsoSyncInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.edir_sso.sync_interval", bytes.NewReader(byt))
	return
}

// GetAuthOtpAutoCreateToken gets the auth.otp.auto_create_token value from the UTM
func GetAuthOtpAutoCreateToken(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.otp.auto_create_token")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpAutoCreateToken PUTs the auth.otp.auto_create_token value to the UTM
func UpdateAuthOtpAutoCreateToken(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.auto_create_token", bytes.NewReader(byt))
	return
}

// GetAuthOtpAutoTokenDigest gets the auth.otp.auto_token_digest value from the UTM
func GetAuthOtpAutoTokenDigest(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/auth.otp.auto_token_digest")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpAutoTokenDigest PUTs the auth.otp.auto_token_digest value to the UTM
func UpdateAuthOtpAutoTokenDigest(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.auto_token_digest", bytes.NewReader(byt))
	return
}

// GetAuthOtpDefaultTimestep gets the auth.otp.default_timestep value from the UTM
func GetAuthOtpDefaultTimestep(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.otp.default_timestep")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpDefaultTimestep PUTs the auth.otp.default_timestep value to the UTM
func UpdateAuthOtpDefaultTimestep(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.default_timestep", bytes.NewReader(byt))
	return
}

// GetAuthOtpFacilities gets the auth.otp.facilities value from the UTM
func GetAuthOtpFacilities(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/auth.otp.facilities")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpFacilities PUTs the auth.otp.facilities value to the UTM
func UpdateAuthOtpFacilities(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.facilities", bytes.NewReader(byt))
	return
}

// GetAuthOtpMaxInitTimestepDiff gets the auth.otp.max_init_timestep_diff value from the UTM
func GetAuthOtpMaxInitTimestepDiff(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.otp.max_init_timestep_diff")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpMaxInitTimestepDiff PUTs the auth.otp.max_init_timestep_diff value to the UTM
func UpdateAuthOtpMaxInitTimestepDiff(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.max_init_timestep_diff", bytes.NewReader(byt))
	return
}

// GetAuthOtpMaxTimestepDiff gets the auth.otp.max_timestep_diff value from the UTM
func GetAuthOtpMaxTimestepDiff(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/auth.otp.max_timestep_diff")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpMaxTimestepDiff PUTs the auth.otp.max_timestep_diff value to the UTM
func UpdateAuthOtpMaxTimestepDiff(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.max_timestep_diff", bytes.NewReader(byt))
	return
}

// GetAuthOtpRequireAllUsers gets the auth.otp.require_all_users value from the UTM
func GetAuthOtpRequireAllUsers(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.otp.require_all_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpRequireAllUsers PUTs the auth.otp.require_all_users value to the UTM
func UpdateAuthOtpRequireAllUsers(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.require_all_users", bytes.NewReader(byt))
	return
}

// GetAuthOtpRequiredUsers gets the auth.otp.required_users value from the UTM
func GetAuthOtpRequiredUsers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/auth.otp.required_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpRequiredUsers PUTs the auth.otp.required_users value to the UTM
func UpdateAuthOtpRequiredUsers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.required_users", bytes.NewReader(byt))
	return
}

// GetAuthOtpStatus gets the auth.otp.status value from the UTM
func GetAuthOtpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.otp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthOtpStatus PUTs the auth.otp.status value to the UTM
func UpdateAuthOtpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.otp.status", bytes.NewReader(byt))
	return
}

// GetAuthServers gets the auth.servers value from the UTM
func GetAuthServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/auth.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthServers PUTs the auth.servers value to the UTM
func UpdateAuthServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.servers", bytes.NewReader(byt))
	return
}

// GetAuthUpdateBackendGroupMembersDebug gets the auth.update_backend_group_members.debug value from the UTM
func GetAuthUpdateBackendGroupMembersDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.update_backend_group_members.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthUpdateBackendGroupMembersDebug PUTs the auth.update_backend_group_members.debug value to the UTM
func UpdateAuthUpdateBackendGroupMembersDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.update_backend_group_members.debug", bytes.NewReader(byt))
	return
}

// GetAuthUpdateBackendGroupMembersStatus gets the auth.update_backend_group_members.status value from the UTM
func GetAuthUpdateBackendGroupMembersStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/auth.update_backend_group_members.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAuthUpdateBackendGroupMembersStatus PUTs the auth.update_backend_group_members.status value to the UTM
func UpdateAuthUpdateBackendGroupMembersStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/auth.update_backend_group_members.status", bytes.NewReader(byt))
	return
}

// GetAweAllowedInterfaces gets the awe.allowed_interfaces value from the UTM
func GetAweAllowedInterfaces(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/awe.allowed_interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweAllowedInterfaces PUTs the awe.allowed_interfaces value to the UTM
func UpdateAweAllowedInterfaces(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.allowed_interfaces", bytes.NewReader(byt))
	return
}

// GetAweClients gets the awe.clients value from the UTM
func GetAweClients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/awe.clients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweClients PUTs the awe.clients value to the UTM
func UpdateAweClients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.clients", bytes.NewReader(byt))
	return
}

// GetAweDevices gets the awe.devices value from the UTM
func GetAweDevices(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/awe.devices")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweDevices PUTs the awe.devices value to the UTM
func UpdateAweDevices(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.devices", bytes.NewReader(byt))
	return
}

// GetAweGlobalApAutoaccept gets the awe.global.ap_autoaccept value from the UTM
func GetAweGlobalApAutoaccept(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/awe.global.ap_autoaccept")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalApAutoaccept PUTs the awe.global.ap_autoaccept value to the UTM
func UpdateAweGlobalApAutoaccept(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.ap_autoaccept", bytes.NewReader(byt))
	return
}

// GetAweGlobalApDebuglevel gets the awe.global.ap_debuglevel value from the UTM
func GetAweGlobalApDebuglevel(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/awe.global.ap_debuglevel")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalApDebuglevel PUTs the awe.global.ap_debuglevel value to the UTM
func UpdateAweGlobalApDebuglevel(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.ap_debuglevel", bytes.NewReader(byt))
	return
}

// GetAweGlobalApSoftlimit gets the awe.global.ap_softlimit value from the UTM
func GetAweGlobalApSoftlimit(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/awe.global.ap_softlimit")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalApSoftlimit PUTs the awe.global.ap_softlimit value to the UTM
func UpdateAweGlobalApSoftlimit(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.ap_softlimit", bytes.NewReader(byt))
	return
}

// GetAweGlobalApVlantag gets the awe.global.ap_vlantag value from the UTM
func GetAweGlobalApVlantag(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/awe.global.ap_vlantag")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalApVlantag PUTs the awe.global.ap_vlantag value to the UTM
func UpdateAweGlobalApVlantag(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.ap_vlantag", bytes.NewReader(byt))
	return
}

// GetAweGlobalAweStatus gets the awe.global.awe_status value from the UTM
func GetAweGlobalAweStatus(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/awe.global.awe_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalAweStatus PUTs the awe.global.awe_status value to the UTM
func UpdateAweGlobalAweStatus(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.awe_status", bytes.NewReader(byt))
	return
}

// GetAweGlobalBridgeUpdateKickout gets the awe.global.bridge_update_kickout value from the UTM
func GetAweGlobalBridgeUpdateKickout(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/awe.global.bridge_update_kickout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalBridgeUpdateKickout PUTs the awe.global.bridge_update_kickout value to the UTM
func UpdateAweGlobalBridgeUpdateKickout(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.bridge_update_kickout", bytes.NewReader(byt))
	return
}

// GetAweGlobalInitialSetup gets the awe.global.initial_setup value from the UTM
func GetAweGlobalInitialSetup(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/awe.global.initial_setup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalInitialSetup PUTs the awe.global.initial_setup value to the UTM
func UpdateAweGlobalInitialSetup(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.initial_setup", bytes.NewReader(byt))
	return
}

// GetAweGlobalLogLevel gets the awe.global.log_level value from the UTM
func GetAweGlobalLogLevel(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/awe.global.log_level")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalLogLevel PUTs the awe.global.log_level value to the UTM
func UpdateAweGlobalLogLevel(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.log_level", bytes.NewReader(byt))
	return
}

// GetAweGlobalMagicIp gets the awe.global.magic_ip value from the UTM
func GetAweGlobalMagicIp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/awe.global.magic_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalMagicIp PUTs the awe.global.magic_ip value to the UTM
func UpdateAweGlobalMagicIp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.magic_ip", bytes.NewReader(byt))
	return
}

// GetAweGlobalNotificationTimeout gets the awe.global.notification_timeout value from the UTM
func GetAweGlobalNotificationTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/awe.global.notification_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalNotificationTimeout PUTs the awe.global.notification_timeout value to the UTM
func UpdateAweGlobalNotificationTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.notification_timeout", bytes.NewReader(byt))
	return
}

// GetAweGlobalRadiusConf gets the awe.global.radius_conf value from the UTM
func GetAweGlobalRadiusConf(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/awe.global.radius_conf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalRadiusConf PUTs the awe.global.radius_conf value to the UTM
func UpdateAweGlobalRadiusConf(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.radius_conf", bytes.NewReader(byt))
	return
}

// GetAweGlobalRootpw gets the awe.global.rootpw value from the UTM
func GetAweGlobalRootpw(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/awe.global.rootpw")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalRootpw PUTs the awe.global.rootpw value to the UTM
func UpdateAweGlobalRootpw(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.rootpw", bytes.NewReader(byt))
	return
}

// GetAweGlobalStayOnline gets the awe.global.stay_online value from the UTM
func GetAweGlobalStayOnline(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/awe.global.stay_online")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalStayOnline PUTs the awe.global.stay_online value to the UTM
func UpdateAweGlobalStayOnline(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.stay_online", bytes.NewReader(byt))
	return
}

// GetAweGlobalStoreBssStats gets the awe.global.store_bss_stats value from the UTM
func GetAweGlobalStoreBssStats(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/awe.global.store_bss_stats")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalStoreBssStats PUTs the awe.global.store_bss_stats value to the UTM
func UpdateAweGlobalStoreBssStats(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.store_bss_stats", bytes.NewReader(byt))
	return
}

// GetAweGlobalTunnelIdOffset gets the awe.global.tunnel_id_offset value from the UTM
func GetAweGlobalTunnelIdOffset(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/awe.global.tunnel_id_offset")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalTunnelIdOffset PUTs the awe.global.tunnel_id_offset value to the UTM
func UpdateAweGlobalTunnelIdOffset(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.tunnel_id_offset", bytes.NewReader(byt))
	return
}

// GetAweGlobalVlantagging gets the awe.global.vlantagging value from the UTM
func GetAweGlobalVlantagging(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/awe.global.vlantagging")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweGlobalVlantagging PUTs the awe.global.vlantagging value to the UTM
func UpdateAweGlobalVlantagging(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.global.vlantagging", bytes.NewReader(byt))
	return
}

// GetAweNetworks gets the awe.networks value from the UTM
func GetAweNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/awe.networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAweNetworks PUTs the awe.networks value to the UTM
func UpdateAweNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awe.networks", bytes.NewReader(byt))
	return
}

// GetAwscliProfiles gets the awscli.profiles value from the UTM
func GetAwscliProfiles(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/awscli.profiles")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateAwscliProfiles PUTs the awscli.profiles value to the UTM
func UpdateAwscliProfiles(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/awscli.profiles", bytes.NewReader(byt))
	return
}

// GetBackupEncryption gets the backup.encryption value from the UTM
func GetBackupEncryption(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/backup.encryption")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateBackupEncryption PUTs the backup.encryption value to the UTM
func UpdateBackupEncryption(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/backup.encryption", bytes.NewReader(byt))
	return
}

// GetBackupInterval gets the backup.interval value from the UTM
func GetBackupInterval(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/backup.interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateBackupInterval PUTs the backup.interval value to the UTM
func UpdateBackupInterval(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/backup.interval", bytes.NewReader(byt))
	return
}

// GetBackupMaxBackups gets the backup.max_backups value from the UTM
func GetBackupMaxBackups(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/backup.max_backups")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateBackupMaxBackups PUTs the backup.max_backups value to the UTM
func UpdateBackupMaxBackups(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/backup.max_backups", bytes.NewReader(byt))
	return
}

// GetBackupPassword gets the backup.password value from the UTM
func GetBackupPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/backup.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateBackupPassword PUTs the backup.password value to the UTM
func UpdateBackupPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/backup.password", bytes.NewReader(byt))
	return
}

// GetBackupRecipients gets the backup.recipients value from the UTM
func GetBackupRecipients(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/backup.recipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateBackupRecipients PUTs the backup.recipients value to the UTM
func UpdateBackupRecipients(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/backup.recipients", bytes.NewReader(byt))
	return
}

// GetBackupStatus gets the backup.status value from the UTM
func GetBackupStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/backup.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateBackupStatus PUTs the backup.status value to the UTM
func UpdateBackupStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/backup.status", bytes.NewReader(byt))
	return
}

// GetCaCaGost gets the ca.ca_gost value from the UTM
func GetCaCaGost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ca.ca_gost")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaCaGost PUTs the ca.ca_gost value to the UTM
func UpdateCaCaGost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.ca_gost", bytes.NewReader(byt))
	return
}

// GetCaCaIpsec gets the ca.ca_ipsec value from the UTM
func GetCaCaIpsec(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ca.ca_ipsec")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaCaIpsec PUTs the ca.ca_ipsec value to the UTM
func UpdateCaCaIpsec(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.ca_ipsec", bytes.NewReader(byt))
	return
}

// GetCaCaProxies gets the ca.ca_proxies value from the UTM
func GetCaCaProxies(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ca.ca_proxies")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaCaProxies PUTs the ca.ca_proxies value to the UTM
func UpdateCaCaProxies(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.ca_proxies", bytes.NewReader(byt))
	return
}

// GetCaCaRed gets the ca.ca_red value from the UTM
func GetCaCaRed(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ca.ca_red")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaCaRed PUTs the ca.ca_red value to the UTM
func UpdateCaCaRed(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.ca_red", bytes.NewReader(byt))
	return
}

// GetCaDefKeysize gets the ca.def_keysize value from the UTM
func GetCaDefKeysize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ca.def_keysize")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaDefKeysize PUTs the ca.def_keysize value to the UTM
func UpdateCaDefKeysize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.def_keysize", bytes.NewReader(byt))
	return
}

// GetCaGlobalCasEmailEncryptionTrustNewCas gets the ca.global_cas.email_encryption.trust_new_cas value from the UTM
func GetCaGlobalCasEmailEncryptionTrustNewCas(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ca.global_cas.email_encryption.trust_new_cas")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaGlobalCasEmailEncryptionTrustNewCas PUTs the ca.global_cas.email_encryption.trust_new_cas value to the UTM
func UpdateCaGlobalCasEmailEncryptionTrustNewCas(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.global_cas.email_encryption.trust_new_cas", bytes.NewReader(byt))
	return
}

// GetCaGlobalCasEmailEncryptionTrusted gets the ca.global_cas.email_encryption.trusted value from the UTM
func GetCaGlobalCasEmailEncryptionTrusted(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ca.global_cas.email_encryption.trusted")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaGlobalCasEmailEncryptionTrusted PUTs the ca.global_cas.email_encryption.trusted value to the UTM
func UpdateCaGlobalCasEmailEncryptionTrusted(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.global_cas.email_encryption.trusted", bytes.NewReader(byt))
	return
}

// GetCaGlobalCasEmailEncryptionUntrusted gets the ca.global_cas.email_encryption.untrusted value from the UTM
func GetCaGlobalCasEmailEncryptionUntrusted(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ca.global_cas.email_encryption.untrusted")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaGlobalCasEmailEncryptionUntrusted PUTs the ca.global_cas.email_encryption.untrusted value to the UTM
func UpdateCaGlobalCasEmailEncryptionUntrusted(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.global_cas.email_encryption.untrusted", bytes.NewReader(byt))
	return
}

// GetCaGlobalCasHttpProxyTrustNewCas gets the ca.global_cas.http_proxy.trust_new_cas value from the UTM
func GetCaGlobalCasHttpProxyTrustNewCas(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ca.global_cas.http_proxy.trust_new_cas")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaGlobalCasHttpProxyTrustNewCas PUTs the ca.global_cas.http_proxy.trust_new_cas value to the UTM
func UpdateCaGlobalCasHttpProxyTrustNewCas(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.global_cas.http_proxy.trust_new_cas", bytes.NewReader(byt))
	return
}

// GetCaGlobalCasHttpProxyTrusted gets the ca.global_cas.http_proxy.trusted value from the UTM
func GetCaGlobalCasHttpProxyTrusted(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ca.global_cas.http_proxy.trusted")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaGlobalCasHttpProxyTrusted PUTs the ca.global_cas.http_proxy.trusted value to the UTM
func UpdateCaGlobalCasHttpProxyTrusted(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.global_cas.http_proxy.trusted", bytes.NewReader(byt))
	return
}

// GetCaGlobalCasHttpProxyUntrusted gets the ca.global_cas.http_proxy.untrusted value from the UTM
func GetCaGlobalCasHttpProxyUntrusted(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ca.global_cas.http_proxy.untrusted")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCaGlobalCasHttpProxyUntrusted PUTs the ca.global_cas.http_proxy.untrusted value to the UTM
func UpdateCaGlobalCasHttpProxyUntrusted(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ca.global_cas.http_proxy.untrusted", bytes.NewReader(byt))
	return
}

// GetCrlsCrls gets the crls.crls value from the UTM
func GetCrlsCrls(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/crls.crls")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCrlsCrls PUTs the crls.crls value to the UTM
func UpdateCrlsCrls(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/crls.crls", bytes.NewReader(byt))
	return
}

// GetCssAvPrimaryEngine gets the css.av_primary_engine value from the UTM
func GetCssAvPrimaryEngine(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/css.av_primary_engine")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCssAvPrimaryEngine PUTs the css.av_primary_engine value to the UTM
func UpdateCssAvPrimaryEngine(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/css.av_primary_engine", bytes.NewReader(byt))
	return
}

// GetCssSxlLiveprotection gets the css.sxl_liveprotection value from the UTM
func GetCssSxlLiveprotection(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/css.sxl_liveprotection")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCssSxlLiveprotection PUTs the css.sxl_liveprotection value to the UTM
func UpdateCssSxlLiveprotection(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/css.sxl_liveprotection", bytes.NewReader(byt))
	return
}

// GetCssSxlSampleSubmit gets the css.sxl_sample_submit value from the UTM
func GetCssSxlSampleSubmit(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/css.sxl_sample_submit")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCssSxlSampleSubmit PUTs the css.sxl_sample_submit value to the UTM
func UpdateCssSxlSampleSubmit(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/css.sxl_sample_submit", bytes.NewReader(byt))
	return
}

// GetCustomizationEppLastUpdated gets the customization.epp.last_updated value from the UTM
func GetCustomizationEppLastUpdated(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/customization.epp.last_updated")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCustomizationEppLastUpdated PUTs the customization.epp.last_updated value to the UTM
func UpdateCustomizationEppLastUpdated(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/customization.epp.last_updated", bytes.NewReader(byt))
	return
}

// GetCustomizationEppResourcesRoot gets the customization.epp.resources_root value from the UTM
func GetCustomizationEppResourcesRoot(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/customization.epp.resources_root")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCustomizationEppResourcesRoot PUTs the customization.epp.resources_root value to the UTM
func UpdateCustomizationEppResourcesRoot(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/customization.epp.resources_root", bytes.NewReader(byt))
	return
}

// GetCustomizationHttpCustomAssets gets the customization.http.custom_assets value from the UTM
func GetCustomizationHttpCustomAssets(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/customization.http.custom_assets")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCustomizationHttpCustomAssets PUTs the customization.http.custom_assets value to the UTM
func UpdateCustomizationHttpCustomAssets(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/customization.http.custom_assets", bytes.NewReader(byt))
	return
}

// GetCustomizationHttpCustomTemplates gets the customization.http.custom_templates value from the UTM
func GetCustomizationHttpCustomTemplates(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/customization.http.custom_templates")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCustomizationHttpCustomTemplates PUTs the customization.http.custom_templates value to the UTM
func UpdateCustomizationHttpCustomTemplates(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/customization.http.custom_templates", bytes.NewReader(byt))
	return
}

// GetCustomizationHttpLastUpdated gets the customization.http.last_updated value from the UTM
func GetCustomizationHttpLastUpdated(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/customization.http.last_updated")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateCustomizationHttpLastUpdated PUTs the customization.http.last_updated value to the UTM
func UpdateCustomizationHttpLastUpdated(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/customization.http.last_updated", bytes.NewReader(byt))
	return
}

// GetDebugmodeCrashReport gets the debugmode.crash_report value from the UTM
func GetDebugmodeCrashReport(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/debugmode.crash_report")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDebugmodeCrashReport PUTs the debugmode.crash_report value to the UTM
func UpdateDebugmodeCrashReport(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/debugmode.crash_report", bytes.NewReader(byt))
	return
}

// GetDebugmodeEnabled gets the debugmode.enabled value from the UTM
func GetDebugmodeEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/debugmode.enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDebugmodeEnabled PUTs the debugmode.enabled value to the UTM
func UpdateDebugmodeEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/debugmode.enabled", bytes.NewReader(byt))
	return
}

// GetDhcpRelayDhcpServer gets the dhcp.relay.dhcp_server value from the UTM
func GetDhcpRelayDhcpServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/dhcp.relay.dhcp_server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpRelayDhcpServer PUTs the dhcp.relay.dhcp_server value to the UTM
func UpdateDhcpRelayDhcpServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.relay.dhcp_server", bytes.NewReader(byt))
	return
}

// GetDhcpRelayInterfaces gets the dhcp.relay.interfaces value from the UTM
func GetDhcpRelayInterfaces(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/dhcp.relay.interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpRelayInterfaces PUTs the dhcp.relay.interfaces value to the UTM
func UpdateDhcpRelayInterfaces(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.relay.interfaces", bytes.NewReader(byt))
	return
}

// GetDhcpRelayStatus gets the dhcp.relay.status value from the UTM
func GetDhcpRelayStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/dhcp.relay.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpRelayStatus PUTs the dhcp.relay.status value to the UTM
func UpdateDhcpRelayStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.relay.status", bytes.NewReader(byt))
	return
}

// GetDhcpRelay6ItfsFacingClients gets the dhcp.relay6.itfs_facing_clients value from the UTM
func GetDhcpRelay6ItfsFacingClients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/dhcp.relay6.itfs_facing_clients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpRelay6ItfsFacingClients PUTs the dhcp.relay6.itfs_facing_clients value to the UTM
func UpdateDhcpRelay6ItfsFacingClients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.relay6.itfs_facing_clients", bytes.NewReader(byt))
	return
}

// GetDhcpRelay6ItfsFacingServer6 gets the dhcp.relay6.itfs_facing_server6 value from the UTM
func GetDhcpRelay6ItfsFacingServer6(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/dhcp.relay6.itfs_facing_server6")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpRelay6ItfsFacingServer6 PUTs the dhcp.relay6.itfs_facing_server6 value to the UTM
func UpdateDhcpRelay6ItfsFacingServer6(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.relay6.itfs_facing_server6", bytes.NewReader(byt))
	return
}

// GetDhcpRelay6Status gets the dhcp.relay6.status value from the UTM
func GetDhcpRelay6Status(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/dhcp.relay6.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpRelay6Status PUTs the dhcp.relay6.status value to the UTM
func UpdateDhcpRelay6Status(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.relay6.status", bytes.NewReader(byt))
	return
}

// GetDhcpServerCustom4 gets the dhcp.server.custom4 value from the UTM
func GetDhcpServerCustom4(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/dhcp.server.custom4")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpServerCustom4 PUTs the dhcp.server.custom4 value to the UTM
func UpdateDhcpServerCustom4(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.server.custom4", bytes.NewReader(byt))
	return
}

// GetDhcpServerCustom6 gets the dhcp.server.custom6 value from the UTM
func GetDhcpServerCustom6(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/dhcp.server.custom6")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpServerCustom6 PUTs the dhcp.server.custom6 value to the UTM
func UpdateDhcpServerCustom6(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.server.custom6", bytes.NewReader(byt))
	return
}

// GetDhcpServerServers gets the dhcp.server.servers value from the UTM
func GetDhcpServerServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/dhcp.server.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDhcpServerServers PUTs the dhcp.server.servers value to the UTM
func UpdateDhcpServerServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dhcp.server.servers", bytes.NewReader(byt))
	return
}

// GetDigestAllowedNetworks gets the digest.allowed_networks value from the UTM
func GetDigestAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/digest.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestAllowedNetworks PUTs the digest.allowed_networks value to the UTM
func UpdateDigestAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.allowed_networks", bytes.NewReader(byt))
	return
}

// GetDigestCustomText gets the digest.custom_text value from the UTM
func GetDigestCustomText(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/digest.custom_text")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestCustomText PUTs the digest.custom_text value to the UTM
func UpdateDigestCustomText(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.custom_text", bytes.NewReader(byt))
	return
}

// GetDigestDomains gets the digest.domains value from the UTM
func GetDigestDomains(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/digest.domains")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestDomains PUTs the digest.domains value to the UTM
func UpdateDigestDomains(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.domains", bytes.NewReader(byt))
	return
}

// GetDigestHostname gets the digest.hostname value from the UTM
func GetDigestHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/digest.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestHostname PUTs the digest.hostname value to the UTM
func UpdateDigestHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.hostname", bytes.NewReader(byt))
	return
}

// GetDigestMailinglists gets the digest.mailinglists value from the UTM
func GetDigestMailinglists(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/digest.mailinglists")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestMailinglists PUTs the digest.mailinglists value to the UTM
func UpdateDigestMailinglists(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.mailinglists", bytes.NewReader(byt))
	return
}

// GetDigestPort gets the digest.port value from the UTM
func GetDigestPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/digest.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestPort PUTs the digest.port value to the UTM
func UpdateDigestPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.port", bytes.NewReader(byt))
	return
}

// GetDigestSendTimeOne gets the digest.send_time_one value from the UTM
func GetDigestSendTimeOne(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/digest.send_time_one")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestSendTimeOne PUTs the digest.send_time_one value to the UTM
func UpdateDigestSendTimeOne(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.send_time_one", bytes.NewReader(byt))
	return
}

// GetDigestSendTimeTwo gets the digest.send_time_two value from the UTM
func GetDigestSendTimeTwo(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/digest.send_time_two")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestSendTimeTwo PUTs the digest.send_time_two value to the UTM
func UpdateDigestSendTimeTwo(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.send_time_two", bytes.NewReader(byt))
	return
}

// GetDigestSkiplist gets the digest.skiplist value from the UTM
func GetDigestSkiplist(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/digest.skiplist")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestSkiplist PUTs the digest.skiplist value to the UTM
func UpdateDigestSkiplist(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.skiplist", bytes.NewReader(byt))
	return
}

// GetDigestStatus gets the digest.status value from the UTM
func GetDigestStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/digest.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestStatus PUTs the digest.status value to the UTM
func UpdateDigestStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.status", bytes.NewReader(byt))
	return
}

// GetDigestUserRelease gets the digest.user_release value from the UTM
func GetDigestUserRelease(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/digest.user_release")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDigestUserRelease PUTs the digest.user_release value to the UTM
func UpdateDigestUserRelease(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/digest.user_release", bytes.NewReader(byt))
	return
}

// GetDnsAllowedNetworks gets the dns.allowed_networks value from the UTM
func GetDnsAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/dns.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsAllowedNetworks PUTs the dns.allowed_networks value to the UTM
func UpdateDnsAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.allowed_networks", bytes.NewReader(byt))
	return
}

// GetDnsAxfr gets the dns.axfr value from the UTM
func GetDnsAxfr(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/dns.axfr")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsAxfr PUTs the dns.axfr value to the UTM
func UpdateDnsAxfr(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.axfr", bytes.NewReader(byt))
	return
}

// GetDnsDnssec gets the dns.dnssec value from the UTM
func GetDnsDnssec(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/dns.dnssec")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsDnssec PUTs the dns.dnssec value to the UTM
func UpdateDnsDnssec(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.dnssec", bytes.NewReader(byt))
	return
}

// GetDnsEmail gets the dns.email value from the UTM
func GetDnsEmail(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/dns.email")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsEmail PUTs the dns.email value to the UTM
func UpdateDnsEmail(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.email", bytes.NewReader(byt))
	return
}

// GetDnsEmptyZones gets the dns.empty_zones value from the UTM
func GetDnsEmptyZones(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/dns.empty_zones")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsEmptyZones PUTs the dns.empty_zones value to the UTM
func UpdateDnsEmptyZones(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.empty_zones", bytes.NewReader(byt))
	return
}

// GetDnsFwdDynamic gets the dns.fwd_dynamic value from the UTM
func GetDnsFwdDynamic(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/dns.fwd_dynamic")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsFwdDynamic PUTs the dns.fwd_dynamic value to the UTM
func UpdateDnsFwdDynamic(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.fwd_dynamic", bytes.NewReader(byt))
	return
}

// GetDnsFwdStatic gets the dns.fwd_static value from the UTM
func GetDnsFwdStatic(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/dns.fwd_static")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsFwdStatic PUTs the dns.fwd_static value to the UTM
func UpdateDnsFwdStatic(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.fwd_static", bytes.NewReader(byt))
	return
}

// GetDnsRecheckInterval gets the dns.recheck_interval value from the UTM
func GetDnsRecheckInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/dns.recheck_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsRecheckInterval PUTs the dns.recheck_interval value to the UTM
func UpdateDnsRecheckInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.recheck_interval", bytes.NewReader(byt))
	return
}

// GetDnsRoutes gets the dns.routes value from the UTM
func GetDnsRoutes(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/dns.routes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDnsRoutes PUTs the dns.routes value to the UTM
func UpdateDnsRoutes(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dns.routes", bytes.NewReader(byt))
	return
}

// GetDyndnsRules gets the dyndns.rules value from the UTM
func GetDyndnsRules(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/dyndns.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateDyndnsRules PUTs the dyndns.rules value to the UTM
func UpdateDyndnsRules(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/dyndns.rules", bytes.NewReader(byt))
	return
}

// GetEmailpkiAuthorityCert gets the emailpki.authority.cert value from the UTM
func GetEmailpkiAuthorityCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.authority.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiAuthorityCert PUTs the emailpki.authority.cert value to the UTM
func UpdateEmailpkiAuthorityCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.authority.cert", bytes.NewReader(byt))
	return
}

// GetEmailpkiAuthorityFingerprint gets the emailpki.authority.fingerprint value from the UTM
func GetEmailpkiAuthorityFingerprint(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.authority.fingerprint")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiAuthorityFingerprint PUTs the emailpki.authority.fingerprint value to the UTM
func UpdateEmailpkiAuthorityFingerprint(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.authority.fingerprint", bytes.NewReader(byt))
	return
}

// GetEmailpkiAuthorityKey gets the emailpki.authority.key value from the UTM
func GetEmailpkiAuthorityKey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.authority.key")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiAuthorityKey PUTs the emailpki.authority.key value to the UTM
func UpdateEmailpkiAuthorityKey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.authority.key", bytes.NewReader(byt))
	return
}

// GetEmailpkiAuthorityPostmasterFingerprint gets the emailpki.authority.postmaster_fingerprint value from the UTM
func GetEmailpkiAuthorityPostmasterFingerprint(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.authority.postmaster_fingerprint")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiAuthorityPostmasterFingerprint PUTs the emailpki.authority.postmaster_fingerprint value to the UTM
func UpdateEmailpkiAuthorityPostmasterFingerprint(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.authority.postmaster_fingerprint", bytes.NewReader(byt))
	return
}

// GetEmailpkiAuthorityPostmasterPrivkey gets the emailpki.authority.postmaster_privkey value from the UTM
func GetEmailpkiAuthorityPostmasterPrivkey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.authority.postmaster_privkey")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiAuthorityPostmasterPrivkey PUTs the emailpki.authority.postmaster_privkey value to the UTM
func UpdateEmailpkiAuthorityPostmasterPrivkey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.authority.postmaster_privkey", bytes.NewReader(byt))
	return
}

// GetEmailpkiAuthorityPostmasterPubkey gets the emailpki.authority.postmaster_pubkey value from the UTM
func GetEmailpkiAuthorityPostmasterPubkey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.authority.postmaster_pubkey")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiAuthorityPostmasterPubkey PUTs the emailpki.authority.postmaster_pubkey value to the UTM
func UpdateEmailpkiAuthorityPostmasterPubkey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.authority.postmaster_pubkey", bytes.NewReader(byt))
	return
}

// GetEmailpkiGlobalCity gets the emailpki.global.city value from the UTM
func GetEmailpkiGlobalCity(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.global.city")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiGlobalCity PUTs the emailpki.global.city value to the UTM
func UpdateEmailpkiGlobalCity(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.global.city", bytes.NewReader(byt))
	return
}

// GetEmailpkiGlobalCountry gets the emailpki.global.country value from the UTM
func GetEmailpkiGlobalCountry(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.global.country")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiGlobalCountry PUTs the emailpki.global.country value to the UTM
func UpdateEmailpkiGlobalCountry(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.global.country", bytes.NewReader(byt))
	return
}

// GetEmailpkiGlobalOrganization gets the emailpki.global.organization value from the UTM
func GetEmailpkiGlobalOrganization(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.global.organization")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiGlobalOrganization PUTs the emailpki.global.organization value to the UTM
func UpdateEmailpkiGlobalOrganization(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.global.organization", bytes.NewReader(byt))
	return
}

// GetEmailpkiGlobalPostmaster gets the emailpki.global.postmaster value from the UTM
func GetEmailpkiGlobalPostmaster(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.global.postmaster")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiGlobalPostmaster PUTs the emailpki.global.postmaster value to the UTM
func UpdateEmailpkiGlobalPostmaster(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.global.postmaster", bytes.NewReader(byt))
	return
}

// GetEmailpkiGlobalStatus gets the emailpki.global.status value from the UTM
func GetEmailpkiGlobalStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/emailpki.global.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiGlobalStatus PUTs the emailpki.global.status value to the UTM
func UpdateEmailpkiGlobalStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.global.status", bytes.NewReader(byt))
	return
}

// GetEmailpkiObjectsCas gets the emailpki.objects.cas value from the UTM
func GetEmailpkiObjectsCas(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/emailpki.objects.cas")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiObjectsCas PUTs the emailpki.objects.cas value to the UTM
func UpdateEmailpkiObjectsCas(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.objects.cas", bytes.NewReader(byt))
	return
}

// GetEmailpkiObjectsOpenpgp gets the emailpki.objects.openpgp value from the UTM
func GetEmailpkiObjectsOpenpgp(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/emailpki.objects.openpgp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiObjectsOpenpgp PUTs the emailpki.objects.openpgp value to the UTM
func UpdateEmailpkiObjectsOpenpgp(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.objects.openpgp", bytes.NewReader(byt))
	return
}

// GetEmailpkiObjectsSmime gets the emailpki.objects.smime value from the UTM
func GetEmailpkiObjectsSmime(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/emailpki.objects.smime")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiObjectsSmime PUTs the emailpki.objects.smime value to the UTM
func UpdateEmailpkiObjectsSmime(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.objects.smime", bytes.NewReader(byt))
	return
}

// GetEmailpkiObjectsUsers gets the emailpki.objects.users value from the UTM
func GetEmailpkiObjectsUsers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/emailpki.objects.users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiObjectsUsers PUTs the emailpki.objects.users value to the UTM
func UpdateEmailpkiObjectsUsers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.objects.users", bytes.NewReader(byt))
	return
}

// GetEmailpkiOpenpgpMainKeysize gets the emailpki.openpgp.main_keysize value from the UTM
func GetEmailpkiOpenpgpMainKeysize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/emailpki.openpgp.main_keysize")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOpenpgpMainKeysize PUTs the emailpki.openpgp.main_keysize value to the UTM
func UpdateEmailpkiOpenpgpMainKeysize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.openpgp.main_keysize", bytes.NewReader(byt))
	return
}

// GetEmailpkiOpenpgpSubKeysize gets the emailpki.openpgp.sub_keysize value from the UTM
func GetEmailpkiOpenpgpSubKeysize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/emailpki.openpgp.sub_keysize")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOpenpgpSubKeysize PUTs the emailpki.openpgp.sub_keysize value to the UTM
func UpdateEmailpkiOpenpgpSubKeysize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.openpgp.sub_keysize", bytes.NewReader(byt))
	return
}

// GetEmailpkiOptionsExternalAuto gets the emailpki.options.external_auto value from the UTM
func GetEmailpkiOptionsExternalAuto(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/emailpki.options.external_auto")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOptionsExternalAuto PUTs the emailpki.options.external_auto value to the UTM
func UpdateEmailpkiOptionsExternalAuto(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.options.external_auto", bytes.NewReader(byt))
	return
}

// GetEmailpkiOptionsKeyserver gets the emailpki.options.keyserver value from the UTM
func GetEmailpkiOptionsKeyserver(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/emailpki.options.keyserver")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOptionsKeyserver PUTs the emailpki.options.keyserver value to the UTM
func UpdateEmailpkiOptionsKeyserver(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.options.keyserver", bytes.NewReader(byt))
	return
}

// GetEmailpkiOptionsPolicyDecryption gets the emailpki.options.policy_decryption value from the UTM
func GetEmailpkiOptionsPolicyDecryption(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/emailpki.options.policy_decryption")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOptionsPolicyDecryption PUTs the emailpki.options.policy_decryption value to the UTM
func UpdateEmailpkiOptionsPolicyDecryption(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.options.policy_decryption", bytes.NewReader(byt))
	return
}

// GetEmailpkiOptionsPolicyEncryption gets the emailpki.options.policy_encryption value from the UTM
func GetEmailpkiOptionsPolicyEncryption(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/emailpki.options.policy_encryption")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOptionsPolicyEncryption PUTs the emailpki.options.policy_encryption value to the UTM
func UpdateEmailpkiOptionsPolicyEncryption(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.options.policy_encryption", bytes.NewReader(byt))
	return
}

// GetEmailpkiOptionsPolicySign gets the emailpki.options.policy_sign value from the UTM
func GetEmailpkiOptionsPolicySign(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/emailpki.options.policy_sign")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOptionsPolicySign PUTs the emailpki.options.policy_sign value to the UTM
func UpdateEmailpkiOptionsPolicySign(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.options.policy_sign", bytes.NewReader(byt))
	return
}

// GetEmailpkiOptionsPolicyVerify gets the emailpki.options.policy_verify value from the UTM
func GetEmailpkiOptionsPolicyVerify(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/emailpki.options.policy_verify")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEmailpkiOptionsPolicyVerify PUTs the emailpki.options.policy_verify value to the UTM
func UpdateEmailpkiOptionsPolicyVerify(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/emailpki.options.policy_verify", bytes.NewReader(byt))
	return
}

// GetEndpointAacAllowedNetworks gets the endpoint.aac.allowed_networks value from the UTM
func GetEndpointAacAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacAllowedNetworks PUTs the endpoint.aac.allowed_networks value to the UTM
func UpdateEndpointAacAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.allowed_networks", bytes.NewReader(byt))
	return
}

// GetEndpointAacAllowedUsers gets the endpoint.aac.allowed_users value from the UTM
func GetEndpointAacAllowedUsers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.allowed_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacAllowedUsers PUTs the endpoint.aac.allowed_users value to the UTM
func UpdateEndpointAacAllowedUsers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.allowed_users", bytes.NewReader(byt))
	return
}

// GetEndpointAacCa gets the endpoint.aac.ca value from the UTM
func GetEndpointAacCa(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.ca")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacCa PUTs the endpoint.aac.ca value to the UTM
func UpdateEndpointAacCa(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.ca", bytes.NewReader(byt))
	return
}

// GetEndpointAacCert gets the endpoint.aac.cert value from the UTM
func GetEndpointAacCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacCert PUTs the endpoint.aac.cert value to the UTM
func UpdateEndpointAacCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.cert", bytes.NewReader(byt))
	return
}

// GetEndpointAacMagicIp gets the endpoint.aac.magic_ip value from the UTM
func GetEndpointAacMagicIp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.magic_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacMagicIp PUTs the endpoint.aac.magic_ip value to the UTM
func UpdateEndpointAacMagicIp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.magic_ip", bytes.NewReader(byt))
	return
}

// GetEndpointAacMaxUserLogins gets the endpoint.aac.max_user_logins value from the UTM
func GetEndpointAacMaxUserLogins(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.max_user_logins")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacMaxUserLogins PUTs the endpoint.aac.max_user_logins value to the UTM
func UpdateEndpointAacMaxUserLogins(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.max_user_logins", bytes.NewReader(byt))
	return
}

// GetEndpointAacStatus gets the endpoint.aac.status value from the UTM
func GetEndpointAacStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/endpoint.aac.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointAacStatus PUTs the endpoint.aac.status value to the UTM
func UpdateEndpointAacStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.aac.status", bytes.NewReader(byt))
	return
}

// GetEndpointStasCollectors gets the endpoint.stas.collectors value from the UTM
func GetEndpointStasCollectors(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/endpoint.stas.collectors")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointStasCollectors PUTs the endpoint.stas.collectors value to the UTM
func UpdateEndpointStasCollectors(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.stas.collectors", bytes.NewReader(byt))
	return
}

// GetEndpointStasStatus gets the endpoint.stas.status value from the UTM
func GetEndpointStasStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/endpoint.stas.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEndpointStasStatus PUTs the endpoint.stas.status value to the UTM
func UpdateEndpointStasStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/endpoint.stas.status", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesCompanyLogo gets the enduser_messages.company_logo value from the UTM
func GetEnduserMessagesCompanyLogo(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.company_logo")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesCompanyLogo PUTs the enduser_messages.company_logo value to the UTM
func UpdateEnduserMessagesCompanyLogo(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.company_logo", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesCompanyText gets the enduser_messages.company_text value from the UTM
func GetEnduserMessagesCompanyText(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.company_text")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesCompanyText PUTs the enduser_messages.company_text value to the UTM
func UpdateEnduserMessagesCompanyText(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.company_text", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesDlpBlackholePart gets the enduser_messages.dlp.blackhole_part value from the UTM
func GetEnduserMessagesDlpBlackholePart(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.dlp.blackhole_part")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesDlpBlackholePart PUTs the enduser_messages.dlp.blackhole_part value to the UTM
func UpdateEnduserMessagesDlpBlackholePart(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.dlp.blackhole_part", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesDlpFooterPart gets the enduser_messages.dlp.footer_part value from the UTM
func GetEnduserMessagesDlpFooterPart(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.dlp.footer_part")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesDlpFooterPart PUTs the enduser_messages.dlp.footer_part value to the UTM
func UpdateEnduserMessagesDlpFooterPart(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.dlp.footer_part", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesDlpHeaderPart gets the enduser_messages.dlp.header_part value from the UTM
func GetEnduserMessagesDlpHeaderPart(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.dlp.header_part")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesDlpHeaderPart PUTs the enduser_messages.dlp.header_part value to the UTM
func UpdateEnduserMessagesDlpHeaderPart(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.dlp.header_part", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesDlpOriginalPart gets the enduser_messages.dlp.original_part value from the UTM
func GetEnduserMessagesDlpOriginalPart(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.dlp.original_part")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesDlpOriginalPart PUTs the enduser_messages.dlp.original_part value to the UTM
func UpdateEnduserMessagesDlpOriginalPart(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.dlp.original_part", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesDlpSpxPart gets the enduser_messages.dlp.spx_part value from the UTM
func GetEnduserMessagesDlpSpxPart(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.dlp.spx_part")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesDlpSpxPart PUTs the enduser_messages.dlp.spx_part value to the UTM
func UpdateEnduserMessagesDlpSpxPart(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.dlp.spx_part", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesDlpSubject gets the enduser_messages.dlp.subject value from the UTM
func GetEnduserMessagesDlpSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.dlp.subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesDlpSubject PUTs the enduser_messages.dlp.subject value to the UTM
func UpdateEnduserMessagesDlpSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.dlp.subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpAppDesc gets the enduser_messages.http.app_desc value from the UTM
func GetEnduserMessagesHttpAppDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.app_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpAppDesc PUTs the enduser_messages.http.app_desc value to the UTM
func UpdateEnduserMessagesHttpAppDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.app_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpAppSubject gets the enduser_messages.http.app_subject value from the UTM
func GetEnduserMessagesHttpAppSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.app_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpAppSubject PUTs the enduser_messages.http.app_subject value to the UTM
func UpdateEnduserMessagesHttpAppSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.app_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpBlacklistDesc gets the enduser_messages.http.blacklist_desc value from the UTM
func GetEnduserMessagesHttpBlacklistDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.blacklist_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpBlacklistDesc PUTs the enduser_messages.http.blacklist_desc value to the UTM
func UpdateEnduserMessagesHttpBlacklistDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.blacklist_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpBlacklistSubject gets the enduser_messages.http.blacklist_subject value from the UTM
func GetEnduserMessagesHttpBlacklistSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.blacklist_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpBlacklistSubject PUTs the enduser_messages.http.blacklist_subject value to the UTM
func UpdateEnduserMessagesHttpBlacklistSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.blacklist_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpCertfailSubject gets the enduser_messages.http.certfail_subject value from the UTM
func GetEnduserMessagesHttpCertfailSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.certfail_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpCertfailSubject PUTs the enduser_messages.http.certfail_subject value to the UTM
func UpdateEnduserMessagesHttpCertfailSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.certfail_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpCffOverrideDesc gets the enduser_messages.http.cff_override_desc value from the UTM
func GetEnduserMessagesHttpCffOverrideDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.cff_override_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpCffOverrideDesc PUTs the enduser_messages.http.cff_override_desc value to the UTM
func UpdateEnduserMessagesHttpCffOverrideDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.cff_override_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpCffOverrideSubject gets the enduser_messages.http.cff_override_subject value from the UTM
func GetEnduserMessagesHttpCffOverrideSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.cff_override_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpCffOverrideSubject PUTs the enduser_messages.http.cff_override_subject value to the UTM
func UpdateEnduserMessagesHttpCffOverrideSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.cff_override_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpCffOverrideTerms gets the enduser_messages.http.cff_override_terms value from the UTM
func GetEnduserMessagesHttpCffOverrideTerms(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.cff_override_terms")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpCffOverrideTerms PUTs the enduser_messages.http.cff_override_terms value to the UTM
func UpdateEnduserMessagesHttpCffOverrideTerms(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.cff_override_terms", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpDownloadCompleteDesc gets the enduser_messages.http.download_complete_desc value from the UTM
func GetEnduserMessagesHttpDownloadCompleteDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.download_complete_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpDownloadCompleteDesc PUTs the enduser_messages.http.download_complete_desc value to the UTM
func UpdateEnduserMessagesHttpDownloadCompleteDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.download_complete_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpDownloadCompleteSubject gets the enduser_messages.http.download_complete_subject value from the UTM
func GetEnduserMessagesHttpDownloadCompleteSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.download_complete_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpDownloadCompleteSubject PUTs the enduser_messages.http.download_complete_subject value to the UTM
func UpdateEnduserMessagesHttpDownloadCompleteSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.download_complete_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpDownloadDesc gets the enduser_messages.http.download_desc value from the UTM
func GetEnduserMessagesHttpDownloadDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.download_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpDownloadDesc PUTs the enduser_messages.http.download_desc value to the UTM
func UpdateEnduserMessagesHttpDownloadDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.download_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpDownloadSubject gets the enduser_messages.http.download_subject value from the UTM
func GetEnduserMessagesHttpDownloadSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.download_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpDownloadSubject PUTs the enduser_messages.http.download_subject value to the UTM
func UpdateEnduserMessagesHttpDownloadSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.download_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpErrorDesc gets the enduser_messages.http.error_desc value from the UTM
func GetEnduserMessagesHttpErrorDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.error_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpErrorDesc PUTs the enduser_messages.http.error_desc value to the UTM
func UpdateEnduserMessagesHttpErrorDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.error_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpErrorSubject gets the enduser_messages.http.error_subject value from the UTM
func GetEnduserMessagesHttpErrorSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.error_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpErrorSubject PUTs the enduser_messages.http.error_subject value to the UTM
func UpdateEnduserMessagesHttpErrorSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.error_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpFileextensionDesc gets the enduser_messages.http.fileextension_desc value from the UTM
func GetEnduserMessagesHttpFileextensionDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.fileextension_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpFileextensionDesc PUTs the enduser_messages.http.fileextension_desc value to the UTM
func UpdateEnduserMessagesHttpFileextensionDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.fileextension_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpFileextensionSubject gets the enduser_messages.http.fileextension_subject value from the UTM
func GetEnduserMessagesHttpFileextensionSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.fileextension_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpFileextensionSubject PUTs the enduser_messages.http.fileextension_subject value to the UTM
func UpdateEnduserMessagesHttpFileextensionSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.fileextension_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpFileextensionWarnDesc gets the enduser_messages.http.fileextension_warn_desc value from the UTM
func GetEnduserMessagesHttpFileextensionWarnDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.fileextension_warn_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpFileextensionWarnDesc PUTs the enduser_messages.http.fileextension_warn_desc value to the UTM
func UpdateEnduserMessagesHttpFileextensionWarnDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.fileextension_warn_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpFileextensionWarnSubject gets the enduser_messages.http.fileextension_warn_subject value from the UTM
func GetEnduserMessagesHttpFileextensionWarnSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.fileextension_warn_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpFileextensionWarnSubject PUTs the enduser_messages.http.fileextension_warn_subject value to the UTM
func UpdateEnduserMessagesHttpFileextensionWarnSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.fileextension_warn_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpFilesizeDesc gets the enduser_messages.http.filesize_desc value from the UTM
func GetEnduserMessagesHttpFilesizeDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.filesize_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpFilesizeDesc PUTs the enduser_messages.http.filesize_desc value to the UTM
func UpdateEnduserMessagesHttpFilesizeDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.filesize_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpFilesizeSubject gets the enduser_messages.http.filesize_subject value from the UTM
func GetEnduserMessagesHttpFilesizeSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.filesize_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpFilesizeSubject PUTs the enduser_messages.http.filesize_subject value to the UTM
func UpdateEnduserMessagesHttpFilesizeSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.filesize_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpGeoipDesc gets the enduser_messages.http.geoip_desc value from the UTM
func GetEnduserMessagesHttpGeoipDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.geoip_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpGeoipDesc PUTs the enduser_messages.http.geoip_desc value to the UTM
func UpdateEnduserMessagesHttpGeoipDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.geoip_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpGeoipSubject gets the enduser_messages.http.geoip_subject value from the UTM
func GetEnduserMessagesHttpGeoipSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.geoip_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpGeoipSubject PUTs the enduser_messages.http.geoip_subject value to the UTM
func UpdateEnduserMessagesHttpGeoipSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.geoip_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpMimetypeDesc gets the enduser_messages.http.mimetype_desc value from the UTM
func GetEnduserMessagesHttpMimetypeDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.mimetype_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpMimetypeDesc PUTs the enduser_messages.http.mimetype_desc value to the UTM
func UpdateEnduserMessagesHttpMimetypeDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.mimetype_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpMimetypeSubject gets the enduser_messages.http.mimetype_subject value from the UTM
func GetEnduserMessagesHttpMimetypeSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.mimetype_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpMimetypeSubject PUTs the enduser_messages.http.mimetype_subject value to the UTM
func UpdateEnduserMessagesHttpMimetypeSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.mimetype_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpMimetypeWarnDesc gets the enduser_messages.http.mimetype_warn_desc value from the UTM
func GetEnduserMessagesHttpMimetypeWarnDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.mimetype_warn_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpMimetypeWarnDesc PUTs the enduser_messages.http.mimetype_warn_desc value to the UTM
func UpdateEnduserMessagesHttpMimetypeWarnDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.mimetype_warn_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpMimetypeWarnSubject gets the enduser_messages.http.mimetype_warn_subject value from the UTM
func GetEnduserMessagesHttpMimetypeWarnSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.mimetype_warn_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpMimetypeWarnSubject PUTs the enduser_messages.http.mimetype_warn_subject value to the UTM
func UpdateEnduserMessagesHttpMimetypeWarnSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.mimetype_warn_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpPuaDesc gets the enduser_messages.http.pua_desc value from the UTM
func GetEnduserMessagesHttpPuaDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.pua_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpPuaDesc PUTs the enduser_messages.http.pua_desc value to the UTM
func UpdateEnduserMessagesHttpPuaDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.pua_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpPuaSubject gets the enduser_messages.http.pua_subject value from the UTM
func GetEnduserMessagesHttpPuaSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.pua_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpPuaSubject PUTs the enduser_messages.http.pua_subject value to the UTM
func UpdateEnduserMessagesHttpPuaSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.pua_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpQuotaBlockDesc gets the enduser_messages.http.quota_block_desc value from the UTM
func GetEnduserMessagesHttpQuotaBlockDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.quota_block_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpQuotaBlockDesc PUTs the enduser_messages.http.quota_block_desc value to the UTM
func UpdateEnduserMessagesHttpQuotaBlockDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.quota_block_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpQuotaBlockSubject gets the enduser_messages.http.quota_block_subject value from the UTM
func GetEnduserMessagesHttpQuotaBlockSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.quota_block_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpQuotaBlockSubject PUTs the enduser_messages.http.quota_block_subject value to the UTM
func UpdateEnduserMessagesHttpQuotaBlockSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.quota_block_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpQuotaWarnDesc gets the enduser_messages.http.quota_warn_desc value from the UTM
func GetEnduserMessagesHttpQuotaWarnDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.quota_warn_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpQuotaWarnDesc PUTs the enduser_messages.http.quota_warn_desc value to the UTM
func UpdateEnduserMessagesHttpQuotaWarnDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.quota_warn_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpQuotaWarnSubject gets the enduser_messages.http.quota_warn_subject value from the UTM
func GetEnduserMessagesHttpQuotaWarnSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.quota_warn_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpQuotaWarnSubject PUTs the enduser_messages.http.quota_warn_subject value to the UTM
func UpdateEnduserMessagesHttpQuotaWarnSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.quota_warn_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSpDesc gets the enduser_messages.http.sp_desc value from the UTM
func GetEnduserMessagesHttpSpDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.sp_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSpDesc PUTs the enduser_messages.http.sp_desc value to the UTM
func UpdateEnduserMessagesHttpSpDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.sp_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSpFrameSubject gets the enduser_messages.http.sp_frame_subject value from the UTM
func GetEnduserMessagesHttpSpFrameSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.sp_frame_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSpFrameSubject PUTs the enduser_messages.http.sp_frame_subject value to the UTM
func UpdateEnduserMessagesHttpSpFrameSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.sp_frame_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSpSubject gets the enduser_messages.http.sp_subject value from the UTM
func GetEnduserMessagesHttpSpSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.sp_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSpSubject PUTs the enduser_messages.http.sp_subject value to the UTM
func UpdateEnduserMessagesHttpSpSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.sp_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSpWarnDesc gets the enduser_messages.http.sp_warn_desc value from the UTM
func GetEnduserMessagesHttpSpWarnDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.sp_warn_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSpWarnDesc PUTs the enduser_messages.http.sp_warn_desc value to the UTM
func UpdateEnduserMessagesHttpSpWarnDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.sp_warn_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSpWarnSubject gets the enduser_messages.http.sp_warn_subject value from the UTM
func GetEnduserMessagesHttpSpWarnSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.sp_warn_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSpWarnSubject PUTs the enduser_messages.http.sp_warn_subject value to the UTM
func UpdateEnduserMessagesHttpSpWarnSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.sp_warn_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslCertraw gets the enduser_messages.http.ssl_certraw value from the UTM
func GetEnduserMessagesHttpSslCertraw(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_certraw")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslCertraw PUTs the enduser_messages.http.ssl_certraw value to the UTM
func UpdateEnduserMessagesHttpSslCertraw(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_certraw", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslCertstatus gets the enduser_messages.http.ssl_certstatus value from the UTM
func GetEnduserMessagesHttpSslCertstatus(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_certstatus")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslCertstatus PUTs the enduser_messages.http.ssl_certstatus value to the UTM
func UpdateEnduserMessagesHttpSslCertstatus(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_certstatus", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslIssuer gets the enduser_messages.http.ssl_issuer value from the UTM
func GetEnduserMessagesHttpSslIssuer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_issuer")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslIssuer PUTs the enduser_messages.http.ssl_issuer value to the UTM
func UpdateEnduserMessagesHttpSslIssuer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_issuer", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslMd5Fp gets the enduser_messages.http.ssl_md5fp value from the UTM
func GetEnduserMessagesHttpSslMd5Fp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_md5fp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslMd5Fp PUTs the enduser_messages.http.ssl_md5fp value to the UTM
func UpdateEnduserMessagesHttpSslMd5Fp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_md5fp", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslSha1Fp gets the enduser_messages.http.ssl_sha1fp value from the UTM
func GetEnduserMessagesHttpSslSha1Fp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_sha1fp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslSha1Fp PUTs the enduser_messages.http.ssl_sha1fp value to the UTM
func UpdateEnduserMessagesHttpSslSha1Fp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_sha1fp", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslSubject gets the enduser_messages.http.ssl_subject value from the UTM
func GetEnduserMessagesHttpSslSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslSubject PUTs the enduser_messages.http.ssl_subject value to the UTM
func UpdateEnduserMessagesHttpSslSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslValidfrom gets the enduser_messages.http.ssl_validfrom value from the UTM
func GetEnduserMessagesHttpSslValidfrom(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_validfrom")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslValidfrom PUTs the enduser_messages.http.ssl_validfrom value to the UTM
func UpdateEnduserMessagesHttpSslValidfrom(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_validfrom", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpSslValiduntil gets the enduser_messages.http.ssl_validuntil value from the UTM
func GetEnduserMessagesHttpSslValiduntil(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.ssl_validuntil")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpSslValiduntil PUTs the enduser_messages.http.ssl_validuntil value to the UTM
func UpdateEnduserMessagesHttpSslValiduntil(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.ssl_validuntil", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpThreatDesc gets the enduser_messages.http.threat_desc value from the UTM
func GetEnduserMessagesHttpThreatDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.threat_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpThreatDesc PUTs the enduser_messages.http.threat_desc value to the UTM
func UpdateEnduserMessagesHttpThreatDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.threat_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpThreatSubject gets the enduser_messages.http.threat_subject value from the UTM
func GetEnduserMessagesHttpThreatSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.threat_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpThreatSubject PUTs the enduser_messages.http.threat_subject value to the UTM
func UpdateEnduserMessagesHttpThreatSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.threat_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpTransparentAuthDesc gets the enduser_messages.http.transparent_auth_desc value from the UTM
func GetEnduserMessagesHttpTransparentAuthDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.transparent_auth_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpTransparentAuthDesc PUTs the enduser_messages.http.transparent_auth_desc value to the UTM
func UpdateEnduserMessagesHttpTransparentAuthDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.transparent_auth_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpTransparentAuthSubject gets the enduser_messages.http.transparent_auth_subject value from the UTM
func GetEnduserMessagesHttpTransparentAuthSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.transparent_auth_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpTransparentAuthSubject PUTs the enduser_messages.http.transparent_auth_subject value to the UTM
func UpdateEnduserMessagesHttpTransparentAuthSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.transparent_auth_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpTransparentAuthTerms gets the enduser_messages.http.transparent_auth_terms value from the UTM
func GetEnduserMessagesHttpTransparentAuthTerms(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.transparent_auth_terms")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpTransparentAuthTerms PUTs the enduser_messages.http.transparent_auth_terms value to the UTM
func UpdateEnduserMessagesHttpTransparentAuthTerms(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.transparent_auth_terms", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpVirusDesc gets the enduser_messages.http.virus_desc value from the UTM
func GetEnduserMessagesHttpVirusDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.virus_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpVirusDesc PUTs the enduser_messages.http.virus_desc value to the UTM
func UpdateEnduserMessagesHttpVirusDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.virus_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpVirusSubject gets the enduser_messages.http.virus_subject value from the UTM
func GetEnduserMessagesHttpVirusSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.virus_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpVirusSubject PUTs the enduser_messages.http.virus_subject value to the UTM
func UpdateEnduserMessagesHttpVirusSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.virus_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpVirusscanDesc gets the enduser_messages.http.virusscan_desc value from the UTM
func GetEnduserMessagesHttpVirusscanDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.virusscan_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpVirusscanDesc PUTs the enduser_messages.http.virusscan_desc value to the UTM
func UpdateEnduserMessagesHttpVirusscanDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.virusscan_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesHttpVirusscanSubject gets the enduser_messages.http.virusscan_subject value from the UTM
func GetEnduserMessagesHttpVirusscanSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.http.virusscan_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesHttpVirusscanSubject PUTs the enduser_messages.http.virusscan_subject value to the UTM
func UpdateEnduserMessagesHttpVirusscanSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.http.virusscan_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesMailReleaseErrDesc gets the enduser_messages.mail.release_err_desc value from the UTM
func GetEnduserMessagesMailReleaseErrDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.mail.release_err_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesMailReleaseErrDesc PUTs the enduser_messages.mail.release_err_desc value to the UTM
func UpdateEnduserMessagesMailReleaseErrDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.mail.release_err_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesMailReleaseErrSubject gets the enduser_messages.mail.release_err_subject value from the UTM
func GetEnduserMessagesMailReleaseErrSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.mail.release_err_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesMailReleaseErrSubject PUTs the enduser_messages.mail.release_err_subject value to the UTM
func UpdateEnduserMessagesMailReleaseErrSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.mail.release_err_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesMailReleasedDesc gets the enduser_messages.mail.released_desc value from the UTM
func GetEnduserMessagesMailReleasedDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.mail.released_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesMailReleasedDesc PUTs the enduser_messages.mail.released_desc value to the UTM
func UpdateEnduserMessagesMailReleasedDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.mail.released_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesMailReleasedSubject gets the enduser_messages.mail.released_subject value from the UTM
func GetEnduserMessagesMailReleasedSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.mail.released_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesMailReleasedSubject PUTs the enduser_messages.mail.released_subject value to the UTM
func UpdateEnduserMessagesMailReleasedSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.mail.released_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesPop3BlockedDesc gets the enduser_messages.pop3.blocked_desc value from the UTM
func GetEnduserMessagesPop3BlockedDesc(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.pop3.blocked_desc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesPop3BlockedDesc PUTs the enduser_messages.pop3.blocked_desc value to the UTM
func UpdateEnduserMessagesPop3BlockedDesc(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.pop3.blocked_desc", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesPop3BlockedSubject gets the enduser_messages.pop3.blocked_subject value from the UTM
func GetEnduserMessagesPop3BlockedSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.pop3.blocked_subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesPop3BlockedSubject PUTs the enduser_messages.pop3.blocked_subject value to the UTM
func UpdateEnduserMessagesPop3BlockedSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.pop3.blocked_subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxInternalErrorBody gets the enduser_messages.spx.internal_error.body value from the UTM
func GetEnduserMessagesSpxInternalErrorBody(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.internal_error.body")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxInternalErrorBody PUTs the enduser_messages.spx.internal_error.body value to the UTM
func UpdateEnduserMessagesSpxInternalErrorBody(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.internal_error.body", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxInternalErrorSubject gets the enduser_messages.spx.internal_error.subject value from the UTM
func GetEnduserMessagesSpxInternalErrorSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.internal_error.subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxInternalErrorSubject PUTs the enduser_messages.spx.internal_error.subject value to the UTM
func UpdateEnduserMessagesSpxInternalErrorSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.internal_error.subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxInternalErrorSenderBody gets the enduser_messages.spx.internal_error_sender.body value from the UTM
func GetEnduserMessagesSpxInternalErrorSenderBody(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.internal_error_sender.body")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxInternalErrorSenderBody PUTs the enduser_messages.spx.internal_error_sender.body value to the UTM
func UpdateEnduserMessagesSpxInternalErrorSenderBody(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.internal_error_sender.body", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxInternalErrorSenderSubject gets the enduser_messages.spx.internal_error_sender.subject value from the UTM
func GetEnduserMessagesSpxInternalErrorSenderSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.internal_error_sender.subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxInternalErrorSenderSubject PUTs the enduser_messages.spx.internal_error_sender.subject value to the UTM
func UpdateEnduserMessagesSpxInternalErrorSenderSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.internal_error_sender.subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxPasswordNoSpecCharsBody gets the enduser_messages.spx.password_no_spec_chars.body value from the UTM
func GetEnduserMessagesSpxPasswordNoSpecCharsBody(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.password_no_spec_chars.body")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxPasswordNoSpecCharsBody PUTs the enduser_messages.spx.password_no_spec_chars.body value to the UTM
func UpdateEnduserMessagesSpxPasswordNoSpecCharsBody(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.password_no_spec_chars.body", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxPasswordNoSpecCharsSubject gets the enduser_messages.spx.password_no_spec_chars.subject value from the UTM
func GetEnduserMessagesSpxPasswordNoSpecCharsSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.password_no_spec_chars.subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxPasswordNoSpecCharsSubject PUTs the enduser_messages.spx.password_no_spec_chars.subject value to the UTM
func UpdateEnduserMessagesSpxPasswordNoSpecCharsSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.password_no_spec_chars.subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxPasswordNotLongEnoughBody gets the enduser_messages.spx.password_not_long_enough.body value from the UTM
func GetEnduserMessagesSpxPasswordNotLongEnoughBody(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.password_not_long_enough.body")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxPasswordNotLongEnoughBody PUTs the enduser_messages.spx.password_not_long_enough.body value to the UTM
func UpdateEnduserMessagesSpxPasswordNotLongEnoughBody(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.password_not_long_enough.body", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxPasswordNotLongEnoughSubject gets the enduser_messages.spx.password_not_long_enough.subject value from the UTM
func GetEnduserMessagesSpxPasswordNotLongEnoughSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.password_not_long_enough.subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxPasswordNotLongEnoughSubject PUTs the enduser_messages.spx.password_not_long_enough.subject value to the UTM
func UpdateEnduserMessagesSpxPasswordNotLongEnoughSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.password_not_long_enough.subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxPasswordNotPresentedBody gets the enduser_messages.spx.password_not_presented.body value from the UTM
func GetEnduserMessagesSpxPasswordNotPresentedBody(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.password_not_presented.body")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxPasswordNotPresentedBody PUTs the enduser_messages.spx.password_not_presented.body value to the UTM
func UpdateEnduserMessagesSpxPasswordNotPresentedBody(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.password_not_presented.body", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxPasswordNotPresentedSubject gets the enduser_messages.spx.password_not_presented.subject value from the UTM
func GetEnduserMessagesSpxPasswordNotPresentedSubject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.password_not_presented.subject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxPasswordNotPresentedSubject PUTs the enduser_messages.spx.password_not_presented.subject value to the UTM
func UpdateEnduserMessagesSpxPasswordNotPresentedSubject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.password_not_presented.subject", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSpxUrlNotFoundMessage gets the enduser_messages.spx.url_not_found.message value from the UTM
func GetEnduserMessagesSpxUrlNotFoundMessage(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.spx.url_not_found.message")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSpxUrlNotFoundMessage PUTs the enduser_messages.spx.url_not_found.message value to the UTM
func UpdateEnduserMessagesSpxUrlNotFoundMessage(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.spx.url_not_found.message", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSquidCacheAdmin gets the enduser_messages.squid.cache_admin value from the UTM
func GetEnduserMessagesSquidCacheAdmin(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.squid.cache_admin")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSquidCacheAdmin PUTs the enduser_messages.squid.cache_admin value to the UTM
func UpdateEnduserMessagesSquidCacheAdmin(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.squid.cache_admin", bytes.NewReader(byt))
	return
}

// GetEnduserMessagesSquidCacheAdminMessage gets the enduser_messages.squid.cache_admin_message value from the UTM
func GetEnduserMessagesSquidCacheAdminMessage(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/enduser_messages.squid.cache_admin_message")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEnduserMessagesSquidCacheAdminMessage PUTs the enduser_messages.squid.cache_admin_message value to the UTM
func UpdateEnduserMessagesSquidCacheAdminMessage(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/enduser_messages.squid.cache_admin_message", bytes.NewReader(byt))
	return
}

// GetEppAllowedNetworks gets the epp.allowed_networks value from the UTM
func GetEppAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/epp.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppAllowedNetworks PUTs the epp.allowed_networks value to the UTM
func UpdateEppAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.allowed_networks", bytes.NewReader(byt))
	return
}

// GetEppCertificate gets the epp.certificate value from the UTM
func GetEppCertificate(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.certificate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppCertificate PUTs the epp.certificate value to the UTM
func UpdateEppCertificate(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.certificate", bytes.NewReader(byt))
	return
}

// GetEppCity gets the epp.city value from the UTM
func GetEppCity(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.city")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppCity PUTs the epp.city value to the UTM
func UpdateEppCity(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.city", bytes.NewReader(byt))
	return
}

// GetEppCountry gets the epp.country value from the UTM
func GetEppCountry(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.country")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppCountry PUTs the epp.country value to the UTM
func UpdateEppCountry(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.country", bytes.NewReader(byt))
	return
}

// GetEppDefaultEndpointsGroup gets the epp.default_endpoints_group value from the UTM
func GetEppDefaultEndpointsGroup(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.default_endpoints_group")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppDefaultEndpointsGroup PUTs the epp.default_endpoints_group value to the UTM
func UpdateEppDefaultEndpointsGroup(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.default_endpoints_group", bytes.NewReader(byt))
	return
}

// GetEppDevices gets the epp.devices value from the UTM
func GetEppDevices(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/epp.devices")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppDevices PUTs the epp.devices value to the UTM
func UpdateEppDevices(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.devices", bytes.NewReader(byt))
	return
}

// GetEppEmail gets the epp.email value from the UTM
func GetEppEmail(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.email")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppEmail PUTs the epp.email value to the UTM
func UpdateEppEmail(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.email", bytes.NewReader(byt))
	return
}

// GetEppEndpoints gets the epp.endpoints value from the UTM
func GetEppEndpoints(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/epp.endpoints")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppEndpoints PUTs the epp.endpoints value to the UTM
func UpdateEppEndpoints(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.endpoints", bytes.NewReader(byt))
	return
}

// GetEppEndpointsGroups gets the epp.endpoints_groups value from the UTM
func GetEppEndpointsGroups(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/epp.endpoints_groups")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppEndpointsGroups PUTs the epp.endpoints_groups value to the UTM
func UpdateEppEndpointsGroups(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.endpoints_groups", bytes.NewReader(byt))
	return
}

// GetEppExceptionsAv gets the epp.exceptions.av value from the UTM
func GetEppExceptionsAv(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/epp.exceptions.av")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppExceptionsAv PUTs the epp.exceptions.av value to the UTM
func UpdateEppExceptionsAv(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.exceptions.av", bytes.NewReader(byt))
	return
}

// GetEppExceptionsDc gets the epp.exceptions.dc value from the UTM
func GetEppExceptionsDc(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/epp.exceptions.dc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppExceptionsDc PUTs the epp.exceptions.dc value to the UTM
func UpdateEppExceptionsDc(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.exceptions.dc", bytes.NewReader(byt))
	return
}

// GetEppFallbackUrl gets the epp.fallback_url value from the UTM
func GetEppFallbackUrl(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.fallback_url")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppFallbackUrl PUTs the epp.fallback_url value to the UTM
func UpdateEppFallbackUrl(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.fallback_url", bytes.NewReader(byt))
	return
}

// GetEppMagnetPassword gets the epp.magnet_password value from the UTM
func GetEppMagnetPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.magnet_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppMagnetPassword PUTs the epp.magnet_password value to the UTM
func UpdateEppMagnetPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.magnet_password", bytes.NewReader(byt))
	return
}

// GetEppMagnetUsername gets the epp.magnet_username value from the UTM
func GetEppMagnetUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.magnet_username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppMagnetUsername PUTs the epp.magnet_username value to the UTM
func UpdateEppMagnetUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.magnet_username", bytes.NewReader(byt))
	return
}

// GetEppOrganization gets the epp.organization value from the UTM
func GetEppOrganization(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.organization")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppOrganization PUTs the epp.organization value to the UTM
func UpdateEppOrganization(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.organization", bytes.NewReader(byt))
	return
}

// GetEppParentProxyHost gets the epp.parent_proxy_host value from the UTM
func GetEppParentProxyHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.parent_proxy_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppParentProxyHost PUTs the epp.parent_proxy_host value to the UTM
func UpdateEppParentProxyHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.parent_proxy_host", bytes.NewReader(byt))
	return
}

// GetEppParentProxyPort gets the epp.parent_proxy_port value from the UTM
func GetEppParentProxyPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/epp.parent_proxy_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppParentProxyPort PUTs the epp.parent_proxy_port value to the UTM
func UpdateEppParentProxyPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.parent_proxy_port", bytes.NewReader(byt))
	return
}

// GetEppParentProxyStatus gets the epp.parent_proxy_status value from the UTM
func GetEppParentProxyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/epp.parent_proxy_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppParentProxyStatus PUTs the epp.parent_proxy_status value to the UTM
func UpdateEppParentProxyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.parent_proxy_status", bytes.NewReader(byt))
	return
}

// GetEppPoliciesAv gets the epp.policies.av value from the UTM
func GetEppPoliciesAv(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/epp.policies.av")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppPoliciesAv PUTs the epp.policies.av value to the UTM
func UpdateEppPoliciesAv(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.policies.av", bytes.NewReader(byt))
	return
}

// GetEppPoliciesDc gets the epp.policies.dc value from the UTM
func GetEppPoliciesDc(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/epp.policies.dc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppPoliciesDc PUTs the epp.policies.dc value to the UTM
func UpdateEppPoliciesDc(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.policies.dc", bytes.NewReader(byt))
	return
}

// GetEppPort gets the epp.port value from the UTM
func GetEppPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/epp.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppPort PUTs the epp.port value to the UTM
func UpdateEppPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.port", bytes.NewReader(byt))
	return
}

// GetEppPrivateKey gets the epp.private_key value from the UTM
func GetEppPrivateKey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.private_key")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppPrivateKey PUTs the epp.private_key value to the UTM
func UpdateEppPrivateKey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.private_key", bytes.NewReader(byt))
	return
}

// GetEppRegistrationToken gets the epp.registration_token value from the UTM
func GetEppRegistrationToken(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.registration_token")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppRegistrationToken PUTs the epp.registration_token value to the UTM
func UpdateEppRegistrationToken(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.registration_token", bytes.NewReader(byt))
	return
}

// GetEppStatusAv gets the epp.status.av value from the UTM
func GetEppStatusAv(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/epp.status.av")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppStatusAv PUTs the epp.status.av value to the UTM
func UpdateEppStatusAv(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.status.av", bytes.NewReader(byt))
	return
}

// GetEppStatusBroker gets the epp.status.broker value from the UTM
func GetEppStatusBroker(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/epp.status.broker")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppStatusBroker PUTs the epp.status.broker value to the UTM
func UpdateEppStatusBroker(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.status.broker", bytes.NewReader(byt))
	return
}

// GetEppStatusDc gets the epp.status.dc value from the UTM
func GetEppStatusDc(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/epp.status.dc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppStatusDc PUTs the epp.status.dc value to the UTM
func UpdateEppStatusDc(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.status.dc", bytes.NewReader(byt))
	return
}

// GetEppStatusEpp gets the epp.status.epp value from the UTM
func GetEppStatusEpp(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/epp.status.epp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppStatusEpp PUTs the epp.status.epp value to the UTM
func UpdateEppStatusEpp(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.status.epp", bytes.NewReader(byt))
	return
}

// GetEppStatusWc gets the epp.status.wc value from the UTM
func GetEppStatusWc(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/epp.status.wc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppStatusWc PUTs the epp.status.wc value to the UTM
func UpdateEppStatusWc(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.status.wc", bytes.NewReader(byt))
	return
}

// GetEppTamperPassword gets the epp.tamper_password value from the UTM
func GetEppTamperPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.tamper_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppTamperPassword PUTs the epp.tamper_password value to the UTM
func UpdateEppTamperPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.tamper_password", bytes.NewReader(byt))
	return
}

// GetEppVersion gets the epp.version value from the UTM
func GetEppVersion(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.version")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppVersion PUTs the epp.version value to the UTM
func UpdateEppVersion(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.version", bytes.NewReader(byt))
	return
}

// GetEppWdxToken gets the epp.wdx_token value from the UTM
func GetEppWdxToken(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/epp.wdx_token")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateEppWdxToken PUTs the epp.wdx_token value to the UTM
func UpdateEppWdxToken(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/epp.wdx_token", bytes.NewReader(byt))
	return
}

// GetExecutiveReportDailyArchive gets the executive_report.daily.archive value from the UTM
func GetExecutiveReportDailyArchive(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/executive_report.daily.archive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportDailyArchive PUTs the executive_report.daily.archive value to the UTM
func UpdateExecutiveReportDailyArchive(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.daily.archive", bytes.NewReader(byt))
	return
}

// GetExecutiveReportDailyKeep gets the executive_report.daily.keep value from the UTM
func GetExecutiveReportDailyKeep(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/executive_report.daily.keep")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportDailyKeep PUTs the executive_report.daily.keep value to the UTM
func UpdateExecutiveReportDailyKeep(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.daily.keep", bytes.NewReader(byt))
	return
}

// GetExecutiveReportDailyPdfrecipients gets the executive_report.daily.pdfrecipients value from the UTM
func GetExecutiveReportDailyPdfrecipients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/executive_report.daily.pdfrecipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportDailyPdfrecipients PUTs the executive_report.daily.pdfrecipients value to the UTM
func UpdateExecutiveReportDailyPdfrecipients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.daily.pdfrecipients", bytes.NewReader(byt))
	return
}

// GetExecutiveReportDailyRecipients gets the executive_report.daily.recipients value from the UTM
func GetExecutiveReportDailyRecipients(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/executive_report.daily.recipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportDailyRecipients PUTs the executive_report.daily.recipients value to the UTM
func UpdateExecutiveReportDailyRecipients(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.daily.recipients", bytes.NewReader(byt))
	return
}

// GetExecutiveReportDailyStatus gets the executive_report.daily.status value from the UTM
func GetExecutiveReportDailyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/executive_report.daily.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportDailyStatus PUTs the executive_report.daily.status value to the UTM
func UpdateExecutiveReportDailyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.daily.status", bytes.NewReader(byt))
	return
}

// GetExecutiveReportMonthlyArchive gets the executive_report.monthly.archive value from the UTM
func GetExecutiveReportMonthlyArchive(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/executive_report.monthly.archive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportMonthlyArchive PUTs the executive_report.monthly.archive value to the UTM
func UpdateExecutiveReportMonthlyArchive(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.monthly.archive", bytes.NewReader(byt))
	return
}

// GetExecutiveReportMonthlyKeep gets the executive_report.monthly.keep value from the UTM
func GetExecutiveReportMonthlyKeep(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/executive_report.monthly.keep")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportMonthlyKeep PUTs the executive_report.monthly.keep value to the UTM
func UpdateExecutiveReportMonthlyKeep(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.monthly.keep", bytes.NewReader(byt))
	return
}

// GetExecutiveReportMonthlyPdfrecipients gets the executive_report.monthly.pdfrecipients value from the UTM
func GetExecutiveReportMonthlyPdfrecipients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/executive_report.monthly.pdfrecipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportMonthlyPdfrecipients PUTs the executive_report.monthly.pdfrecipients value to the UTM
func UpdateExecutiveReportMonthlyPdfrecipients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.monthly.pdfrecipients", bytes.NewReader(byt))
	return
}

// GetExecutiveReportMonthlyRecipients gets the executive_report.monthly.recipients value from the UTM
func GetExecutiveReportMonthlyRecipients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/executive_report.monthly.recipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportMonthlyRecipients PUTs the executive_report.monthly.recipients value to the UTM
func UpdateExecutiveReportMonthlyRecipients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.monthly.recipients", bytes.NewReader(byt))
	return
}

// GetExecutiveReportMonthlyStatus gets the executive_report.monthly.status value from the UTM
func GetExecutiveReportMonthlyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/executive_report.monthly.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportMonthlyStatus PUTs the executive_report.monthly.status value to the UTM
func UpdateExecutiveReportMonthlyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.monthly.status", bytes.NewReader(byt))
	return
}

// GetExecutiveReportWeeklyArchive gets the executive_report.weekly.archive value from the UTM
func GetExecutiveReportWeeklyArchive(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/executive_report.weekly.archive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportWeeklyArchive PUTs the executive_report.weekly.archive value to the UTM
func UpdateExecutiveReportWeeklyArchive(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.weekly.archive", bytes.NewReader(byt))
	return
}

// GetExecutiveReportWeeklyFirstDayOfWeek gets the executive_report.weekly.first_day_of_week value from the UTM
func GetExecutiveReportWeeklyFirstDayOfWeek(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/executive_report.weekly.first_day_of_week")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportWeeklyFirstDayOfWeek PUTs the executive_report.weekly.first_day_of_week value to the UTM
func UpdateExecutiveReportWeeklyFirstDayOfWeek(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.weekly.first_day_of_week", bytes.NewReader(byt))
	return
}

// GetExecutiveReportWeeklyKeep gets the executive_report.weekly.keep value from the UTM
func GetExecutiveReportWeeklyKeep(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/executive_report.weekly.keep")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportWeeklyKeep PUTs the executive_report.weekly.keep value to the UTM
func UpdateExecutiveReportWeeklyKeep(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.weekly.keep", bytes.NewReader(byt))
	return
}

// GetExecutiveReportWeeklyPdfrecipients gets the executive_report.weekly.pdfrecipients value from the UTM
func GetExecutiveReportWeeklyPdfrecipients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/executive_report.weekly.pdfrecipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportWeeklyPdfrecipients PUTs the executive_report.weekly.pdfrecipients value to the UTM
func UpdateExecutiveReportWeeklyPdfrecipients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.weekly.pdfrecipients", bytes.NewReader(byt))
	return
}

// GetExecutiveReportWeeklyRecipients gets the executive_report.weekly.recipients value from the UTM
func GetExecutiveReportWeeklyRecipients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/executive_report.weekly.recipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportWeeklyRecipients PUTs the executive_report.weekly.recipients value to the UTM
func UpdateExecutiveReportWeeklyRecipients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.weekly.recipients", bytes.NewReader(byt))
	return
}

// GetExecutiveReportWeeklyStatus gets the executive_report.weekly.status value from the UTM
func GetExecutiveReportWeeklyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/executive_report.weekly.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateExecutiveReportWeeklyStatus PUTs the executive_report.weekly.status value to the UTM
func UpdateExecutiveReportWeeklyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/executive_report.weekly.status", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpDstBurst gets the flood_protection.icmp.dst_burst value from the UTM
func GetFloodProtectionIcmpDstBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.dst_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpDstBurst PUTs the flood_protection.icmp.dst_burst value to the UTM
func UpdateFloodProtectionIcmpDstBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.dst_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpDstExpire gets the flood_protection.icmp.dst_expire value from the UTM
func GetFloodProtectionIcmpDstExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.dst_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpDstExpire PUTs the flood_protection.icmp.dst_expire value to the UTM
func UpdateFloodProtectionIcmpDstExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.dst_expire", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpDstGcInterval gets the flood_protection.icmp.dst_gc_interval value from the UTM
func GetFloodProtectionIcmpDstGcInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.dst_gc_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpDstGcInterval PUTs the flood_protection.icmp.dst_gc_interval value to the UTM
func UpdateFloodProtectionIcmpDstGcInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.dst_gc_interval", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpDstRate gets the flood_protection.icmp.dst_rate value from the UTM
func GetFloodProtectionIcmpDstRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.dst_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpDstRate PUTs the flood_protection.icmp.dst_rate value to the UTM
func UpdateFloodProtectionIcmpDstRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.dst_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpLog gets the flood_protection.icmp.log value from the UTM
func GetFloodProtectionIcmpLog(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.log")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpLog PUTs the flood_protection.icmp.log value to the UTM
func UpdateFloodProtectionIcmpLog(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.log", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpLogLimitBurst gets the flood_protection.icmp.log_limit_burst value from the UTM
func GetFloodProtectionIcmpLogLimitBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.log_limit_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpLogLimitBurst PUTs the flood_protection.icmp.log_limit_burst value to the UTM
func UpdateFloodProtectionIcmpLogLimitBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.log_limit_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpLogLimitRate gets the flood_protection.icmp.log_limit_rate value from the UTM
func GetFloodProtectionIcmpLogLimitRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.log_limit_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpLogLimitRate PUTs the flood_protection.icmp.log_limit_rate value to the UTM
func UpdateFloodProtectionIcmpLogLimitRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.log_limit_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpMode gets the flood_protection.icmp.mode value from the UTM
func GetFloodProtectionIcmpMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpMode PUTs the flood_protection.icmp.mode value to the UTM
func UpdateFloodProtectionIcmpMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.mode", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpSrcBurst gets the flood_protection.icmp.src_burst value from the UTM
func GetFloodProtectionIcmpSrcBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.src_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpSrcBurst PUTs the flood_protection.icmp.src_burst value to the UTM
func UpdateFloodProtectionIcmpSrcBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.src_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpSrcExpire gets the flood_protection.icmp.src_expire value from the UTM
func GetFloodProtectionIcmpSrcExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.src_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpSrcExpire PUTs the flood_protection.icmp.src_expire value to the UTM
func UpdateFloodProtectionIcmpSrcExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.src_expire", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpSrcGcInterval gets the flood_protection.icmp.src_gc_interval value from the UTM
func GetFloodProtectionIcmpSrcGcInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.src_gc_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpSrcGcInterval PUTs the flood_protection.icmp.src_gc_interval value to the UTM
func UpdateFloodProtectionIcmpSrcGcInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.src_gc_interval", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpSrcRate gets the flood_protection.icmp.src_rate value from the UTM
func GetFloodProtectionIcmpSrcRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.src_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpSrcRate PUTs the flood_protection.icmp.src_rate value to the UTM
func UpdateFloodProtectionIcmpSrcRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.src_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionIcmpStatus gets the flood_protection.icmp.status value from the UTM
func GetFloodProtectionIcmpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/flood_protection.icmp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionIcmpStatus PUTs the flood_protection.icmp.status value to the UTM
func UpdateFloodProtectionIcmpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.icmp.status", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynDstBurst gets the flood_protection.syn.dst_burst value from the UTM
func GetFloodProtectionSynDstBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.dst_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynDstBurst PUTs the flood_protection.syn.dst_burst value to the UTM
func UpdateFloodProtectionSynDstBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.dst_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynDstExpire gets the flood_protection.syn.dst_expire value from the UTM
func GetFloodProtectionSynDstExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.dst_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynDstExpire PUTs the flood_protection.syn.dst_expire value to the UTM
func UpdateFloodProtectionSynDstExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.dst_expire", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynDstGcInterval gets the flood_protection.syn.dst_gc_interval value from the UTM
func GetFloodProtectionSynDstGcInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.dst_gc_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynDstGcInterval PUTs the flood_protection.syn.dst_gc_interval value to the UTM
func UpdateFloodProtectionSynDstGcInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.dst_gc_interval", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynDstRate gets the flood_protection.syn.dst_rate value from the UTM
func GetFloodProtectionSynDstRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.dst_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynDstRate PUTs the flood_protection.syn.dst_rate value to the UTM
func UpdateFloodProtectionSynDstRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.dst_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynLog gets the flood_protection.syn.log value from the UTM
func GetFloodProtectionSynLog(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.log")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynLog PUTs the flood_protection.syn.log value to the UTM
func UpdateFloodProtectionSynLog(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.log", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynLogLimitBurst gets the flood_protection.syn.log_limit_burst value from the UTM
func GetFloodProtectionSynLogLimitBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.log_limit_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynLogLimitBurst PUTs the flood_protection.syn.log_limit_burst value to the UTM
func UpdateFloodProtectionSynLogLimitBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.log_limit_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynLogLimitRate gets the flood_protection.syn.log_limit_rate value from the UTM
func GetFloodProtectionSynLogLimitRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.log_limit_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynLogLimitRate PUTs the flood_protection.syn.log_limit_rate value to the UTM
func UpdateFloodProtectionSynLogLimitRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.log_limit_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynMode gets the flood_protection.syn.mode value from the UTM
func GetFloodProtectionSynMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynMode PUTs the flood_protection.syn.mode value to the UTM
func UpdateFloodProtectionSynMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.mode", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynSrcBurst gets the flood_protection.syn.src_burst value from the UTM
func GetFloodProtectionSynSrcBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.src_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynSrcBurst PUTs the flood_protection.syn.src_burst value to the UTM
func UpdateFloodProtectionSynSrcBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.src_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynSrcExpire gets the flood_protection.syn.src_expire value from the UTM
func GetFloodProtectionSynSrcExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.src_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynSrcExpire PUTs the flood_protection.syn.src_expire value to the UTM
func UpdateFloodProtectionSynSrcExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.src_expire", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynSrcGcInterval gets the flood_protection.syn.src_gc_interval value from the UTM
func GetFloodProtectionSynSrcGcInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.src_gc_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynSrcGcInterval PUTs the flood_protection.syn.src_gc_interval value to the UTM
func UpdateFloodProtectionSynSrcGcInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.src_gc_interval", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynSrcRate gets the flood_protection.syn.src_rate value from the UTM
func GetFloodProtectionSynSrcRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.src_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynSrcRate PUTs the flood_protection.syn.src_rate value to the UTM
func UpdateFloodProtectionSynSrcRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.src_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionSynStatus gets the flood_protection.syn.status value from the UTM
func GetFloodProtectionSynStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/flood_protection.syn.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionSynStatus PUTs the flood_protection.syn.status value to the UTM
func UpdateFloodProtectionSynStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.syn.status", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpDstBurst gets the flood_protection.udp.dst_burst value from the UTM
func GetFloodProtectionUdpDstBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.dst_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpDstBurst PUTs the flood_protection.udp.dst_burst value to the UTM
func UpdateFloodProtectionUdpDstBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.dst_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpDstExpire gets the flood_protection.udp.dst_expire value from the UTM
func GetFloodProtectionUdpDstExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.dst_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpDstExpire PUTs the flood_protection.udp.dst_expire value to the UTM
func UpdateFloodProtectionUdpDstExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.dst_expire", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpDstGcInterval gets the flood_protection.udp.dst_gc_interval value from the UTM
func GetFloodProtectionUdpDstGcInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.dst_gc_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpDstGcInterval PUTs the flood_protection.udp.dst_gc_interval value to the UTM
func UpdateFloodProtectionUdpDstGcInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.dst_gc_interval", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpDstRate gets the flood_protection.udp.dst_rate value from the UTM
func GetFloodProtectionUdpDstRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.dst_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpDstRate PUTs the flood_protection.udp.dst_rate value to the UTM
func UpdateFloodProtectionUdpDstRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.dst_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpLog gets the flood_protection.udp.log value from the UTM
func GetFloodProtectionUdpLog(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.log")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpLog PUTs the flood_protection.udp.log value to the UTM
func UpdateFloodProtectionUdpLog(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.log", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpLogLimitBurst gets the flood_protection.udp.log_limit_burst value from the UTM
func GetFloodProtectionUdpLogLimitBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.log_limit_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpLogLimitBurst PUTs the flood_protection.udp.log_limit_burst value to the UTM
func UpdateFloodProtectionUdpLogLimitBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.log_limit_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpLogLimitRate gets the flood_protection.udp.log_limit_rate value from the UTM
func GetFloodProtectionUdpLogLimitRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.log_limit_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpLogLimitRate PUTs the flood_protection.udp.log_limit_rate value to the UTM
func UpdateFloodProtectionUdpLogLimitRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.log_limit_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpMode gets the flood_protection.udp.mode value from the UTM
func GetFloodProtectionUdpMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpMode PUTs the flood_protection.udp.mode value to the UTM
func UpdateFloodProtectionUdpMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.mode", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpSrcBurst gets the flood_protection.udp.src_burst value from the UTM
func GetFloodProtectionUdpSrcBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.src_burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpSrcBurst PUTs the flood_protection.udp.src_burst value to the UTM
func UpdateFloodProtectionUdpSrcBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.src_burst", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpSrcExpire gets the flood_protection.udp.src_expire value from the UTM
func GetFloodProtectionUdpSrcExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.src_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpSrcExpire PUTs the flood_protection.udp.src_expire value to the UTM
func UpdateFloodProtectionUdpSrcExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.src_expire", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpSrcGcInterval gets the flood_protection.udp.src_gc_interval value from the UTM
func GetFloodProtectionUdpSrcGcInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.src_gc_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpSrcGcInterval PUTs the flood_protection.udp.src_gc_interval value to the UTM
func UpdateFloodProtectionUdpSrcGcInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.src_gc_interval", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpSrcRate gets the flood_protection.udp.src_rate value from the UTM
func GetFloodProtectionUdpSrcRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.src_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpSrcRate PUTs the flood_protection.udp.src_rate value to the UTM
func UpdateFloodProtectionUdpSrcRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.src_rate", bytes.NewReader(byt))
	return
}

// GetFloodProtectionUdpStatus gets the flood_protection.udp.status value from the UTM
func GetFloodProtectionUdpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/flood_protection.udp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFloodProtectionUdpStatus PUTs the flood_protection.udp.status value to the UTM
func UpdateFloodProtectionUdpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/flood_protection.udp.status", bytes.NewReader(byt))
	return
}

// GetFtpAllowedClients gets the ftp.allowed_clients value from the UTM
func GetFtpAllowedClients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ftp.allowed_clients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpAllowedClients PUTs the ftp.allowed_clients value to the UTM
func UpdateFtpAllowedClients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.allowed_clients", bytes.NewReader(byt))
	return
}

// GetFtpAllowedServers gets the ftp.allowed_servers value from the UTM
func GetFtpAllowedServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ftp.allowed_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpAllowedServers PUTs the ftp.allowed_servers value to the UTM
func UpdateFtpAllowedServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.allowed_servers", bytes.NewReader(byt))
	return
}

// GetFtpCffAv gets the ftp.cff_av value from the UTM
func GetFtpCffAv(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ftp.cff_av")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpCffAv PUTs the ftp.cff_av value to the UTM
func UpdateFtpCffAv(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.cff_av", bytes.NewReader(byt))
	return
}

// GetFtpCffAvEngines gets the ftp.cff_av_engines value from the UTM
func GetFtpCffAvEngines(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ftp.cff_av_engines")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpCffAvEngines PUTs the ftp.cff_av_engines value to the UTM
func UpdateFtpCffAvEngines(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.cff_av_engines", bytes.NewReader(byt))
	return
}

// GetFtpCffFileExtensions gets the ftp.cff_file_extensions value from the UTM
func GetFtpCffFileExtensions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ftp.cff_file_extensions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpCffFileExtensions PUTs the ftp.cff_file_extensions value to the UTM
func UpdateFtpCffFileExtensions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.cff_file_extensions", bytes.NewReader(byt))
	return
}

// GetFtpExceptions gets the ftp.exceptions value from the UTM
func GetFtpExceptions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ftp.exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpExceptions PUTs the ftp.exceptions value to the UTM
func UpdateFtpExceptions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.exceptions", bytes.NewReader(byt))
	return
}

// GetFtpMaxFileSize gets the ftp.max_file_size value from the UTM
func GetFtpMaxFileSize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ftp.max_file_size")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpMaxFileSize PUTs the ftp.max_file_size value to the UTM
func UpdateFtpMaxFileSize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.max_file_size", bytes.NewReader(byt))
	return
}

// GetFtpMsWinMode gets the ftp.ms_win_mode value from the UTM
func GetFtpMsWinMode(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ftp.ms_win_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpMsWinMode PUTs the ftp.ms_win_mode value to the UTM
func UpdateFtpMsWinMode(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.ms_win_mode", bytes.NewReader(byt))
	return
}

// GetFtpOperationMode gets the ftp.operation_mode value from the UTM
func GetFtpOperationMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ftp.operation_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpOperationMode PUTs the ftp.operation_mode value to the UTM
func UpdateFtpOperationMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.operation_mode", bytes.NewReader(byt))
	return
}

// GetFtpRestrictedServers gets the ftp.restricted_servers value from the UTM
func GetFtpRestrictedServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ftp.restricted_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpRestrictedServers PUTs the ftp.restricted_servers value to the UTM
func UpdateFtpRestrictedServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.restricted_servers", bytes.NewReader(byt))
	return
}

// GetFtpStatus gets the ftp.status value from the UTM
func GetFtpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ftp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpStatus PUTs the ftp.status value to the UTM
func UpdateFtpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.status", bytes.NewReader(byt))
	return
}

// GetFtpTransparentSkip gets the ftp.transparent_skip value from the UTM
func GetFtpTransparentSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ftp.transparent_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpTransparentSkip PUTs the ftp.transparent_skip value to the UTM
func UpdateFtpTransparentSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.transparent_skip", bytes.NewReader(byt))
	return
}

// GetFtpTransparentSkipAutoPf gets the ftp.transparent_skip_auto_pf value from the UTM
func GetFtpTransparentSkipAutoPf(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ftp.transparent_skip_auto_pf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateFtpTransparentSkipAutoPf PUTs the ftp.transparent_skip_auto_pf value to the UTM
func UpdateFtpTransparentSkipAutoPf(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ftp.transparent_skip_auto_pf", bytes.NewReader(byt))
	return
}

// GetGenericProxyRules gets the generic_proxy.rules value from the UTM
func GetGenericProxyRules(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/generic_proxy.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateGenericProxyRules PUTs the generic_proxy.rules value to the UTM
func UpdateGenericProxyRules(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/generic_proxy.rules", bytes.NewReader(byt))
	return
}

// GetGeoipCountriesDst gets the geoip.countries_dst value from the UTM
func GetGeoipCountriesDst(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/geoip.countries_dst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateGeoipCountriesDst PUTs the geoip.countries_dst value to the UTM
func UpdateGeoipCountriesDst(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/geoip.countries_dst", bytes.NewReader(byt))
	return
}

// GetGeoipCountriesSrc gets the geoip.countries_src value from the UTM
func GetGeoipCountriesSrc(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/geoip.countries_src")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateGeoipCountriesSrc PUTs the geoip.countries_src value to the UTM
func UpdateGeoipCountriesSrc(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/geoip.countries_src", bytes.NewReader(byt))
	return
}

// GetGeoipExceptions gets the geoip.exceptions value from the UTM
func GetGeoipExceptions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/geoip.exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateGeoipExceptions PUTs the geoip.exceptions value to the UTM
func UpdateGeoipExceptions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/geoip.exceptions", bytes.NewReader(byt))
	return
}

// GetGeoipLog gets the geoip.log value from the UTM
func GetGeoipLog(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/geoip.log")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateGeoipLog PUTs the geoip.log value to the UTM
func UpdateGeoipLog(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/geoip.log", bytes.NewReader(byt))
	return
}

// GetGeoipStatus gets the geoip.status value from the UTM
func GetGeoipStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/geoip.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateGeoipStatus PUTs the geoip.status value to the UTM
func UpdateGeoipStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/geoip.status", bytes.NewReader(byt))
	return
}

// GetH323AllowedNetworks gets the h323.allowed_networks value from the UTM
func GetH323AllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/h323.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateH323AllowedNetworks PUTs the h323.allowed_networks value to the UTM
func UpdateH323AllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/h323.allowed_networks", bytes.NewReader(byt))
	return
}

// GetH323LogRelated gets the h323.log_related value from the UTM
func GetH323LogRelated(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/h323.log_related")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateH323LogRelated PUTs the h323.log_related value to the UTM
func UpdateH323LogRelated(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/h323.log_related", bytes.NewReader(byt))
	return
}

// GetH323Servers gets the h323.servers value from the UTM
func GetH323Servers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/h323.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateH323Servers PUTs the h323.servers value to the UTM
func UpdateH323Servers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/h323.servers", bytes.NewReader(byt))
	return
}

// GetH323Status gets the h323.status value from the UTM
func GetH323Status(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/h323.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateH323Status PUTs the h323.status value to the UTM
func UpdateH323Status(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/h323.status", bytes.NewReader(byt))
	return
}

// GetHaAdvancedAutojoin gets the ha.advanced.autojoin value from the UTM
func GetHaAdvancedAutojoin(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.autojoin")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedAutojoin PUTs the ha.advanced.autojoin value to the UTM
func UpdateHaAdvancedAutojoin(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.autojoin", bytes.NewReader(byt))
	return
}

// GetHaAdvancedColdRollback gets the ha.advanced.cold_rollback value from the UTM
func GetHaAdvancedColdRollback(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.cold_rollback")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedColdRollback PUTs the ha.advanced.cold_rollback value to the UTM
func UpdateHaAdvancedColdRollback(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.cold_rollback", bytes.NewReader(byt))
	return
}

// GetHaAdvancedHttpPersistenceTime gets the ha.advanced.http_persistence_time value from the UTM
func GetHaAdvancedHttpPersistenceTime(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.http_persistence_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedHttpPersistenceTime PUTs the ha.advanced.http_persistence_time value to the UTM
func UpdateHaAdvancedHttpPersistenceTime(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.http_persistence_time", bytes.NewReader(byt))
	return
}

// GetHaAdvancedLoadTakeover gets the ha.advanced.load_takeover value from the UTM
func GetHaAdvancedLoadTakeover(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.load_takeover")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedLoadTakeover PUTs the ha.advanced.load_takeover value to the UTM
func UpdateHaAdvancedLoadTakeover(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.load_takeover", bytes.NewReader(byt))
	return
}

// GetHaAdvancedLoadWarn gets the ha.advanced.load_warn value from the UTM
func GetHaAdvancedLoadWarn(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.load_warn")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedLoadWarn PUTs the ha.advanced.load_warn value to the UTM
func UpdateHaAdvancedLoadWarn(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.load_warn", bytes.NewReader(byt))
	return
}

// GetHaAdvancedMaxNodes gets the ha.advanced.max_nodes value from the UTM
func GetHaAdvancedMaxNodes(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.max_nodes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedMaxNodes PUTs the ha.advanced.max_nodes value to the UTM
func UpdateHaAdvancedMaxNodes(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.max_nodes", bytes.NewReader(byt))
	return
}

// GetHaAdvancedMtu gets the ha.advanced.mtu value from the UTM
func GetHaAdvancedMtu(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.mtu")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedMtu PUTs the ha.advanced.mtu value to the UTM
func UpdateHaAdvancedMtu(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.mtu", bytes.NewReader(byt))
	return
}

// GetHaAdvancedNetconsole gets the ha.advanced.netconsole value from the UTM
func GetHaAdvancedNetconsole(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.netconsole")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedNetconsole PUTs the ha.advanced.netconsole value to the UTM
func UpdateHaAdvancedNetconsole(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.netconsole", bytes.NewReader(byt))
	return
}

// GetHaAdvancedPreempt gets the ha.advanced.preempt value from the UTM
func GetHaAdvancedPreempt(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.preempt")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedPreempt PUTs the ha.advanced.preempt value to the UTM
func UpdateHaAdvancedPreempt(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.preempt", bytes.NewReader(byt))
	return
}

// GetHaAdvancedUniqueId gets the ha.advanced.unique_id value from the UTM
func GetHaAdvancedUniqueId(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.unique_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedUniqueId PUTs the ha.advanced.unique_id value to the UTM
func UpdateHaAdvancedUniqueId(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.unique_id", bytes.NewReader(byt))
	return
}

// GetHaAdvancedVirtualMac gets the ha.advanced.virtual_mac value from the UTM
func GetHaAdvancedVirtualMac(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.advanced.virtual_mac")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAdvancedVirtualMac PUTs the ha.advanced.virtual_mac value to the UTM
func UpdateHaAdvancedVirtualMac(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.advanced.virtual_mac", bytes.NewReader(byt))
	return
}

// GetHaAwsCloudwatchProfile gets the ha.aws.cloudwatch.profile value from the UTM
func GetHaAwsCloudwatchProfile(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.aws.cloudwatch.profile")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsCloudwatchProfile PUTs the ha.aws.cloudwatch.profile value to the UTM
func UpdateHaAwsCloudwatchProfile(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.cloudwatch.profile", bytes.NewReader(byt))
	return
}

// GetHaAwsCloudwatchStatus gets the ha.aws.cloudwatch.status value from the UTM
func GetHaAwsCloudwatchStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.cloudwatch.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsCloudwatchStatus PUTs the ha.aws.cloudwatch.status value to the UTM
func UpdateHaAwsCloudwatchStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.cloudwatch.status", bytes.NewReader(byt))
	return
}

// GetHaAwsConfdBackup gets the ha.aws.confd.backup value from the UTM
func GetHaAwsConfdBackup(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.confd.backup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsConfdBackup PUTs the ha.aws.confd.backup value to the UTM
func UpdateHaAwsConfdBackup(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.confd.backup", bytes.NewReader(byt))
	return
}

// GetHaAwsConfdBackupInterval gets the ha.aws.confd.backup_interval value from the UTM
func GetHaAwsConfdBackupInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.aws.confd.backup_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsConfdBackupInterval PUTs the ha.aws.confd.backup_interval value to the UTM
func UpdateHaAwsConfdBackupInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.confd.backup_interval", bytes.NewReader(byt))
	return
}

// GetHaAwsConfdRestore gets the ha.aws.confd.restore value from the UTM
func GetHaAwsConfdRestore(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.confd.restore")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsConfdRestore PUTs the ha.aws.confd.restore value to the UTM
func UpdateHaAwsConfdRestore(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.confd.restore", bytes.NewReader(byt))
	return
}

// GetHaAwsConfdRestoreDone gets the ha.aws.confd.restore_done value from the UTM
func GetHaAwsConfdRestoreDone(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.confd.restore_done")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsConfdRestoreDone PUTs the ha.aws.confd.restore_done value to the UTM
func UpdateHaAwsConfdRestoreDone(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.confd.restore_done", bytes.NewReader(byt))
	return
}

// GetHaAwsElasticIp gets the ha.aws.elastic_ip value from the UTM
func GetHaAwsElasticIp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.aws.elastic_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsElasticIp PUTs the ha.aws.elastic_ip value to the UTM
func UpdateHaAwsElasticIp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.elastic_ip", bytes.NewReader(byt))
	return
}

// GetHaAwsPostgresArchiveTimeout gets the ha.aws.postgres.archive_timeout value from the UTM
func GetHaAwsPostgresArchiveTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.aws.postgres.archive_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsPostgresArchiveTimeout PUTs the ha.aws.postgres.archive_timeout value to the UTM
func UpdateHaAwsPostgresArchiveTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.postgres.archive_timeout", bytes.NewReader(byt))
	return
}

// GetHaAwsPostgresBackup gets the ha.aws.postgres.backup value from the UTM
func GetHaAwsPostgresBackup(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.postgres.backup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsPostgresBackup PUTs the ha.aws.postgres.backup value to the UTM
func UpdateHaAwsPostgresBackup(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.postgres.backup", bytes.NewReader(byt))
	return
}

// GetHaAwsPostgresBaseBackupInterval gets the ha.aws.postgres.base_backup_interval value from the UTM
func GetHaAwsPostgresBaseBackupInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.aws.postgres.base_backup_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsPostgresBaseBackupInterval PUTs the ha.aws.postgres.base_backup_interval value to the UTM
func UpdateHaAwsPostgresBaseBackupInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.postgres.base_backup_interval", bytes.NewReader(byt))
	return
}

// GetHaAwsPostgresRestore gets the ha.aws.postgres.restore value from the UTM
func GetHaAwsPostgresRestore(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.postgres.restore")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsPostgresRestore PUTs the ha.aws.postgres.restore value to the UTM
func UpdateHaAwsPostgresRestore(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.postgres.restore", bytes.NewReader(byt))
	return
}

// GetHaAwsS3Bucket gets the ha.aws.s3_bucket value from the UTM
func GetHaAwsS3Bucket(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.aws.s3_bucket")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsS3Bucket PUTs the ha.aws.s3_bucket value to the UTM
func UpdateHaAwsS3Bucket(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.s3_bucket", bytes.NewReader(byt))
	return
}

// GetHaAwsStackName gets the ha.aws.stack_name value from the UTM
func GetHaAwsStackName(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.aws.stack_name")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsStackName PUTs the ha.aws.stack_name value to the UTM
func UpdateHaAwsStackName(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.stack_name", bytes.NewReader(byt))
	return
}

// GetHaAwsSyslogBackup gets the ha.aws.syslog.backup value from the UTM
func GetHaAwsSyslogBackup(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.syslog.backup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsSyslogBackup PUTs the ha.aws.syslog.backup value to the UTM
func UpdateHaAwsSyslogBackup(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.syslog.backup", bytes.NewReader(byt))
	return
}

// GetHaAwsSyslogRestore gets the ha.aws.syslog.restore value from the UTM
func GetHaAwsSyslogRestore(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.aws.syslog.restore")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsSyslogRestore PUTs the ha.aws.syslog.restore value to the UTM
func UpdateHaAwsSyslogRestore(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.syslog.restore", bytes.NewReader(byt))
	return
}

// GetHaAwsSyslogRestorePeriod gets the ha.aws.syslog.restore_period value from the UTM
func GetHaAwsSyslogRestorePeriod(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.aws.syslog.restore_period")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsSyslogRestorePeriod PUTs the ha.aws.syslog.restore_period value to the UTM
func UpdateHaAwsSyslogRestorePeriod(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.syslog.restore_period", bytes.NewReader(byt))
	return
}

// GetHaAwsTrustedNetwork gets the ha.aws.trusted_network value from the UTM
func GetHaAwsTrustedNetwork(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.aws.trusted_network")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaAwsTrustedNetwork PUTs the ha.aws.trusted_network value to the UTM
func UpdateHaAwsTrustedNetwork(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.aws.trusted_network", bytes.NewReader(byt))
	return
}

// GetHaClusterFtp gets the ha.cluster.ftp value from the UTM
func GetHaClusterFtp(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.ftp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterFtp PUTs the ha.cluster.ftp value to the UTM
func UpdateHaClusterFtp(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.ftp", bytes.NewReader(byt))
	return
}

// GetHaClusterHttp gets the ha.cluster.http value from the UTM
func GetHaClusterHttp(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.http")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterHttp PUTs the ha.cluster.http value to the UTM
func UpdateHaClusterHttp(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.http", bytes.NewReader(byt))
	return
}

// GetHaClusterIpsec gets the ha.cluster.ipsec value from the UTM
func GetHaClusterIpsec(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.ipsec")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterIpsec PUTs the ha.cluster.ipsec value to the UTM
func UpdateHaClusterIpsec(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.ipsec", bytes.NewReader(byt))
	return
}

// GetHaClusterPop3 gets the ha.cluster.pop3 value from the UTM
func GetHaClusterPop3(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.pop3")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterPop3 PUTs the ha.cluster.pop3 value to the UTM
func UpdateHaClusterPop3(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.pop3", bytes.NewReader(byt))
	return
}

// GetHaClusterSmtp gets the ha.cluster.smtp value from the UTM
func GetHaClusterSmtp(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.smtp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterSmtp PUTs the ha.cluster.smtp value to the UTM
func UpdateHaClusterSmtp(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.smtp", bytes.NewReader(byt))
	return
}

// GetHaClusterSnort gets the ha.cluster.snort value from the UTM
func GetHaClusterSnort(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.snort")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterSnort PUTs the ha.cluster.snort value to the UTM
func UpdateHaClusterSnort(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.snort", bytes.NewReader(byt))
	return
}

// GetHaClusterWaf gets the ha.cluster.waf value from the UTM
func GetHaClusterWaf(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.cluster.waf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaClusterWaf PUTs the ha.cluster.waf value to the UTM
func UpdateHaClusterWaf(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.cluster.waf", bytes.NewReader(byt))
	return
}

// GetHaDeviceName gets the ha.device_name value from the UTM
func GetHaDeviceName(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.device_name")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaDeviceName PUTs the ha.device_name value to the UTM
func UpdateHaDeviceName(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.device_name", bytes.NewReader(byt))
	return
}

// GetHaItfhw gets the ha.itfhw value from the UTM
func GetHaItfhw(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.itfhw")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaItfhw PUTs the ha.itfhw value to the UTM
func UpdateHaItfhw(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.itfhw", bytes.NewReader(byt))
	return
}

// GetHaItfhwBackup gets the ha.itfhw_backup value from the UTM
func GetHaItfhwBackup(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.itfhw_backup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaItfhwBackup PUTs the ha.itfhw_backup value to the UTM
func UpdateHaItfhwBackup(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.itfhw_backup", bytes.NewReader(byt))
	return
}

// GetHaMasterIp gets the ha.master_ip value from the UTM
func GetHaMasterIp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.master_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaMasterIp PUTs the ha.master_ip value to the UTM
func UpdateHaMasterIp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.master_ip", bytes.NewReader(byt))
	return
}

// GetHaMode gets the ha.mode value from the UTM
func GetHaMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaMode PUTs the ha.mode value to the UTM
func UpdateHaMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.mode", bytes.NewReader(byt))
	return
}

// GetHaNodeId gets the ha.node_id value from the UTM
func GetHaNodeId(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ha.node_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaNodeId PUTs the ha.node_id value to the UTM
func UpdateHaNodeId(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.node_id", bytes.NewReader(byt))
	return
}

// GetHaPassword gets the ha.password value from the UTM
func GetHaPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaPassword PUTs the ha.password value to the UTM
func UpdateHaPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.password", bytes.NewReader(byt))
	return
}

// GetHaPostgresSecret gets the ha.postgres_secret value from the UTM
func GetHaPostgresSecret(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.postgres_secret")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaPostgresSecret PUTs the ha.postgres_secret value to the UTM
func UpdateHaPostgresSecret(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.postgres_secret", bytes.NewReader(byt))
	return
}

// GetHaSlaveIp gets the ha.slave_ip value from the UTM
func GetHaSlaveIp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.slave_ip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaSlaveIp PUTs the ha.slave_ip value to the UTM
func UpdateHaSlaveIp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.slave_ip", bytes.NewReader(byt))
	return
}

// GetHaStatus gets the ha.status value from the UTM
func GetHaStatus(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ha.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaStatus PUTs the ha.status value to the UTM
func UpdateHaStatus(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.status", bytes.NewReader(byt))
	return
}

// GetHaSyncConntrack gets the ha.sync.conntrack value from the UTM
func GetHaSyncConntrack(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.sync.conntrack")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaSyncConntrack PUTs the ha.sync.conntrack value to the UTM
func UpdateHaSyncConntrack(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.sync.conntrack", bytes.NewReader(byt))
	return
}

// GetHaSyncDatabase gets the ha.sync.database value from the UTM
func GetHaSyncDatabase(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.sync.database")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaSyncDatabase PUTs the ha.sync.database value to the UTM
func UpdateHaSyncDatabase(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.sync.database", bytes.NewReader(byt))
	return
}

// GetHaSyncFiles gets the ha.sync.files value from the UTM
func GetHaSyncFiles(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.sync.files")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaSyncFiles PUTs the ha.sync.files value to the UTM
func UpdateHaSyncFiles(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.sync.files", bytes.NewReader(byt))
	return
}

// GetHaSyncIpsec gets the ha.sync.ipsec value from the UTM
func GetHaSyncIpsec(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.sync.ipsec")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaSyncIpsec PUTs the ha.sync.ipsec value to the UTM
func UpdateHaSyncIpsec(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.sync.ipsec", bytes.NewReader(byt))
	return
}

// GetHaSyncSyslog gets the ha.sync.syslog value from the UTM
func GetHaSyncSyslog(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ha.sync.syslog")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaSyncSyslog PUTs the ha.sync.syslog value to the UTM
func UpdateHaSyncSyslog(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.sync.syslog", bytes.NewReader(byt))
	return
}

// GetHaTimesDeadTime gets the ha.times.dead_time value from the UTM
func GetHaTimesDeadTime(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.times.dead_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaTimesDeadTime PUTs the ha.times.dead_time value to the UTM
func UpdateHaTimesDeadTime(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.times.dead_time", bytes.NewReader(byt))
	return
}

// GetHaTimesLoadTime gets the ha.times.load_time value from the UTM
func GetHaTimesLoadTime(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ha.times.load_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHaTimesLoadTime PUTs the ha.times.load_time value to the UTM
func UpdateHaTimesLoadTime(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ha.times.load_time", bytes.NewReader(byt))
	return
}

// GetHotspotCert gets the hotspot.cert value from the UTM
func GetHotspotCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/hotspot.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHotspotCert PUTs the hotspot.cert value to the UTM
func UpdateHotspotCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/hotspot.cert", bytes.NewReader(byt))
	return
}

// GetHotspotDeleteDays gets the hotspot.delete_days value from the UTM
func GetHotspotDeleteDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/hotspot.delete_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHotspotDeleteDays PUTs the hotspot.delete_days value to the UTM
func UpdateHotspotDeleteDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/hotspot.delete_days", bytes.NewReader(byt))
	return
}

// GetHotspotSslPortal gets the hotspot.ssl_portal value from the UTM
func GetHotspotSslPortal(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/hotspot.ssl_portal")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHotspotSslPortal PUTs the hotspot.ssl_portal value to the UTM
func UpdateHotspotSslPortal(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/hotspot.ssl_portal", bytes.NewReader(byt))
	return
}

// GetHotspotStatus gets the hotspot.status value from the UTM
func GetHotspotStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/hotspot.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHotspotStatus PUTs the hotspot.status value to the UTM
func UpdateHotspotStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/hotspot.status", bytes.NewReader(byt))
	return
}

// GetHotspotTransparentSkip gets the hotspot.transparent_skip value from the UTM
func GetHotspotTransparentSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/hotspot.transparent_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHotspotTransparentSkip PUTs the hotspot.transparent_skip value to the UTM
func UpdateHotspotTransparentSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/hotspot.transparent_skip", bytes.NewReader(byt))
	return
}

// GetHttpAdSsoInterfaces gets the http.ad_sso_interfaces value from the UTM
func GetHttpAdSsoInterfaces(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.ad_sso_interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAdSsoInterfaces PUTs the http.ad_sso_interfaces value to the UTM
func UpdateHttpAdSsoInterfaces(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.ad_sso_interfaces", bytes.NewReader(byt))
	return
}

// GetHttpAdssoRedirectUseHostname gets the http.adsso_redirect_use_hostname value from the UTM
func GetHttpAdssoRedirectUseHostname(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.adsso_redirect_use_hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAdssoRedirectUseHostname PUTs the http.adsso_redirect_use_hostname value to the UTM
func UpdateHttpAdssoRedirectUseHostname(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.adsso_redirect_use_hostname", bytes.NewReader(byt))
	return
}

// GetHttpAllowSsl3 gets the http.allow_ssl3 value from the UTM
func GetHttpAllowSsl3(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.allow_ssl3")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAllowSsl3 PUTs the http.allow_ssl3 value to the UTM
func UpdateHttpAllowSsl3(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.allow_ssl3", bytes.NewReader(byt))
	return
}

// GetHttpAllowTls12 gets the http.allow_tls_1_2 value from the UTM
func GetHttpAllowTls12(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.allow_tls_1_2")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAllowTls12 PUTs the http.allow_tls_1_2 value to the UTM
func UpdateHttpAllowTls12(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.allow_tls_1_2", bytes.NewReader(byt))
	return
}

// GetHttpAllowedPuas gets the http.allowed_puas value from the UTM
func GetHttpAllowedPuas(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.allowed_puas")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAllowedPuas PUTs the http.allowed_puas value to the UTM
func UpdateHttpAllowedPuas(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.allowed_puas", bytes.NewReader(byt))
	return
}

// GetHttpAllowedTargetServices gets the http.allowed_target_services value from the UTM
func GetHttpAllowedTargetServices(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.allowed_target_services")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAllowedTargetServices PUTs the http.allowed_target_services value to the UTM
func UpdateHttpAllowedTargetServices(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.allowed_target_services", bytes.NewReader(byt))
	return
}

// GetHttpAuaMaxconns gets the http.aua_maxconns value from the UTM
func GetHttpAuaMaxconns(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.aua_maxconns")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAuaMaxconns PUTs the http.aua_maxconns value to the UTM
func UpdateHttpAuaMaxconns(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.aua_maxconns", bytes.NewReader(byt))
	return
}

// GetHttpAuaTimeout gets the http.aua_timeout value from the UTM
func GetHttpAuaTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.aua_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAuaTimeout PUTs the http.aua_timeout value to the UTM
func UpdateHttpAuaTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.aua_timeout", bytes.NewReader(byt))
	return
}

// GetHttpAuthCacheSize gets the http.auth_cache_size value from the UTM
func GetHttpAuthCacheSize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.auth_cache_size")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAuthCacheSize PUTs the http.auth_cache_size value to the UTM
func UpdateHttpAuthCacheSize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.auth_cache_size", bytes.NewReader(byt))
	return
}

// GetHttpAuthCacheTtl gets the http.auth_cache_ttl value from the UTM
func GetHttpAuthCacheTtl(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.auth_cache_ttl")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAuthCacheTtl PUTs the http.auth_cache_ttl value to the UTM
func UpdateHttpAuthCacheTtl(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.auth_cache_ttl", bytes.NewReader(byt))
	return
}

// GetHttpAuthRealm gets the http.auth_realm value from the UTM
func GetHttpAuthRealm(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.auth_realm")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAuthRealm PUTs the http.auth_realm value to the UTM
func UpdateHttpAuthRealm(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.auth_realm", bytes.NewReader(byt))
	return
}

// GetHttpAuthUsercacheTtl gets the http.auth_usercache_ttl value from the UTM
func GetHttpAuthUsercacheTtl(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.auth_usercache_ttl")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpAuthUsercacheTtl PUTs the http.auth_usercache_ttl value to the UTM
func UpdateHttpAuthUsercacheTtl(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.auth_usercache_ttl", bytes.NewReader(byt))
	return
}

// GetHttpBlockUnscannable gets the http.block_unscannable value from the UTM
func GetHttpBlockUnscannable(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.block_unscannable")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpBlockUnscannable PUTs the http.block_unscannable value to the UTM
func UpdateHttpBlockUnscannable(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.block_unscannable", bytes.NewReader(byt))
	return
}

// GetHttpBypassStreaming gets the http.bypass_streaming value from the UTM
func GetHttpBypassStreaming(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.bypass_streaming")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpBypassStreaming PUTs the http.bypass_streaming value to the UTM
func UpdateHttpBypassStreaming(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.bypass_streaming", bytes.NewReader(byt))
	return
}

// GetHttpCaList gets the http.ca_list value from the UTM
func GetHttpCaList(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.ca_list")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCaList PUTs the http.ca_list value to the UTM
func UpdateHttpCaList(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.ca_list", bytes.NewReader(byt))
	return
}

// GetHttpCacheIgnoresCookies gets the http.cache_ignores_cookies value from the UTM
func GetHttpCacheIgnoresCookies(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.cache_ignores_cookies")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCacheIgnoresCookies PUTs the http.cache_ignores_cookies value to the UTM
func UpdateHttpCacheIgnoresCookies(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.cache_ignores_cookies", bytes.NewReader(byt))
	return
}

// GetHttpCachessl gets the http.cachessl value from the UTM
func GetHttpCachessl(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.cachessl")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCachessl PUTs the http.cachessl value to the UTM
func UpdateHttpCachessl(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.cachessl", bytes.NewReader(byt))
	return
}

// GetHttpCaching gets the http.caching value from the UTM
func GetHttpCaching(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.caching")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCaching PUTs the http.caching value to the UTM
func UpdateHttpCaching(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.caching", bytes.NewReader(byt))
	return
}

// GetHttpCertcache gets the http.certcache value from the UTM
func GetHttpCertcache(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.certcache")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCertcache PUTs the http.certcache value to the UTM
func UpdateHttpCertcache(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.certcache", bytes.NewReader(byt))
	return
}

// GetHttpCertstore gets the http.certstore value from the UTM
func GetHttpCertstore(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.certstore")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCertstore PUTs the http.certstore value to the UTM
func UpdateHttpCertstore(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.certstore", bytes.NewReader(byt))
	return
}

// GetHttpCffOverrideUsers gets the http.cff_override_users value from the UTM
func GetHttpCffOverrideUsers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.cff_override_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCffOverrideUsers PUTs the http.cff_override_users value to the UTM
func UpdateHttpCffOverrideUsers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.cff_override_users", bytes.NewReader(byt))
	return
}

// GetHttpClientTimeout gets the http.client_timeout value from the UTM
func GetHttpClientTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.client_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpClientTimeout PUTs the http.client_timeout value to the UTM
func UpdateHttpClientTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.client_timeout", bytes.NewReader(byt))
	return
}

// GetHttpConfLockWorkaround gets the http.conf_lock_workaround value from the UTM
func GetHttpConfLockWorkaround(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.conf_lock_workaround")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpConfLockWorkaround PUTs the http.conf_lock_workaround value to the UTM
func UpdateHttpConfLockWorkaround(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.conf_lock_workaround", bytes.NewReader(byt))
	return
}

// GetHttpConnectTimeout gets the http.connect_timeout value from the UTM
func GetHttpConnectTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.connect_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpConnectTimeout PUTs the http.connect_timeout value to the UTM
func UpdateHttpConnectTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.connect_timeout", bytes.NewReader(byt))
	return
}

// GetHttpConnectV6Timeout gets the http.connect_v6_timeout value from the UTM
func GetHttpConnectV6Timeout(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/http.connect_v6_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpConnectV6Timeout PUTs the http.connect_v6_timeout value to the UTM
func UpdateHttpConnectV6Timeout(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.connect_v6_timeout", bytes.NewReader(byt))
	return
}

// GetHttpConnlimit gets the http.connlimit value from the UTM
func GetHttpConnlimit(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.connlimit")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpConnlimit PUTs the http.connlimit value to the UTM
func UpdateHttpConnlimit(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.connlimit", bytes.NewReader(byt))
	return
}

// GetHttpCtypeInspectBody gets the http.ctype_inspect_body value from the UTM
func GetHttpCtypeInspectBody(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.ctype_inspect_body")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCtypeInspectBody PUTs the http.ctype_inspect_body value to the UTM
func UpdateHttpCtypeInspectBody(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.ctype_inspect_body", bytes.NewReader(byt))
	return
}

// GetHttpCtypeUnpackArchive gets the http.ctype_unpack_archive value from the UTM
func GetHttpCtypeUnpackArchive(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.ctype_unpack_archive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpCtypeUnpackArchive PUTs the http.ctype_unpack_archive value to the UTM
func UpdateHttpCtypeUnpackArchive(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.ctype_unpack_archive", bytes.NewReader(byt))
	return
}

// GetHttpDebug gets the http.debug value from the UTM
func GetHttpDebug(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDebug PUTs the http.debug value to the UTM
func UpdateHttpDebug(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.debug", bytes.NewReader(byt))
	return
}

// GetHttpDefaultblockaction gets the http.defaultblockaction value from the UTM
func GetHttpDefaultblockaction(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.defaultblockaction")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDefaultblockaction PUTs the http.defaultblockaction value to the UTM
func UpdateHttpDefaultblockaction(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.defaultblockaction", bytes.NewReader(byt))
	return
}

// GetHttpDeferagents gets the http.deferagents value from the UTM
func GetHttpDeferagents(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.deferagents")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDeferagents PUTs the http.deferagents value to the UTM
func UpdateHttpDeferagents(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.deferagents", bytes.NewReader(byt))
	return
}

// GetHttpDeferlength gets the http.deferlength value from the UTM
func GetHttpDeferlength(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.deferlength")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDeferlength PUTs the http.deferlength value to the UTM
func UpdateHttpDeferlength(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.deferlength", bytes.NewReader(byt))
	return
}

// GetHttpDisplayHttpBlockpageExplicitMode gets the http.display_http_blockpage_explicit_mode value from the UTM
func GetHttpDisplayHttpBlockpageExplicitMode(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.display_http_blockpage_explicit_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDisplayHttpBlockpageExplicitMode PUTs the http.display_http_blockpage_explicit_mode value to the UTM
func UpdateHttpDisplayHttpBlockpageExplicitMode(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.display_http_blockpage_explicit_mode", bytes.NewReader(byt))
	return
}

// GetHttpDisplayIntro gets the http.display_intro value from the UTM
func GetHttpDisplayIntro(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.display_intro")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDisplayIntro PUTs the http.display_intro value to the UTM
func UpdateHttpDisplayIntro(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.display_intro", bytes.NewReader(byt))
	return
}

// GetHttpDownloadManagerDefaultCharset gets the http.download_manager_default_charset value from the UTM
func GetHttpDownloadManagerDefaultCharset(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.download_manager_default_charset")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpDownloadManagerDefaultCharset PUTs the http.download_manager_default_charset value to the UTM
func UpdateHttpDownloadManagerDefaultCharset(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.download_manager_default_charset", bytes.NewReader(byt))
	return
}

// GetHttpEdirDelayBasicAuth gets the http.edir_delay_basic_auth value from the UTM
func GetHttpEdirDelayBasicAuth(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.edir_delay_basic_auth")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpEdirDelayBasicAuth PUTs the http.edir_delay_basic_auth value to the UTM
func UpdateHttpEdirDelayBasicAuth(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.edir_delay_basic_auth", bytes.NewReader(byt))
	return
}

// GetHttpEnableOutInterface gets the http.enable_out_interface value from the UTM
func GetHttpEnableOutInterface(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.enable_out_interface")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpEnableOutInterface PUTs the http.enable_out_interface value to the UTM
func UpdateHttpEnableOutInterface(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.enable_out_interface", bytes.NewReader(byt))
	return
}

// GetHttpEppQuotaAction gets the http.epp_quota_action value from the UTM
func GetHttpEppQuotaAction(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.epp_quota_action")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpEppQuotaAction PUTs the http.epp_quota_action value to the UTM
func UpdateHttpEppQuotaAction(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.epp_quota_action", bytes.NewReader(byt))
	return
}

// GetHttpExceptions gets the http.exceptions value from the UTM
func GetHttpExceptions(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpExceptions PUTs the http.exceptions value to the UTM
func UpdateHttpExceptions(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.exceptions", bytes.NewReader(byt))
	return
}

// GetHttpForcedCachingExtension gets the http.forced_caching_extension value from the UTM
func GetHttpForcedCachingExtension(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.forced_caching_extension")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpForcedCachingExtension PUTs the http.forced_caching_extension value to the UTM
func UpdateHttpForcedCachingExtension(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.forced_caching_extension", bytes.NewReader(byt))
	return
}

// GetHttpForcedCachingNeverCachePrefix gets the http.forced_caching_never_cache_prefix value from the UTM
func GetHttpForcedCachingNeverCachePrefix(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.forced_caching_never_cache_prefix")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpForcedCachingNeverCachePrefix PUTs the http.forced_caching_never_cache_prefix value to the UTM
func UpdateHttpForcedCachingNeverCachePrefix(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.forced_caching_never_cache_prefix", bytes.NewReader(byt))
	return
}

// GetHttpForcedCachingStatus gets the http.forced_caching_status value from the UTM
func GetHttpForcedCachingStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.forced_caching_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpForcedCachingStatus PUTs the http.forced_caching_status value to the UTM
func UpdateHttpForcedCachingStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.forced_caching_status", bytes.NewReader(byt))
	return
}

// GetHttpForcedCachingTtl gets the http.forced_caching_ttl value from the UTM
func GetHttpForcedCachingTtl(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.forced_caching_ttl")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpForcedCachingTtl PUTs the http.forced_caching_ttl value to the UTM
func UpdateHttpForcedCachingTtl(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.forced_caching_ttl", bytes.NewReader(byt))
	return
}

// GetHttpForcedCachingUserAgentPrefix gets the http.forced_caching_user_agent_prefix value from the UTM
func GetHttpForcedCachingUserAgentPrefix(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.forced_caching_user_agent_prefix")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpForcedCachingUserAgentPrefix PUTs the http.forced_caching_user_agent_prefix value to the UTM
func UpdateHttpForcedCachingUserAgentPrefix(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.forced_caching_user_agent_prefix", bytes.NewReader(byt))
	return
}

// GetHttpHttpLoopbackDetect gets the http.http_loopback_detect value from the UTM
func GetHttpHttpLoopbackDetect(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.http_loopback_detect")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpHttpLoopbackDetect PUTs the http.http_loopback_detect value to the UTM
func UpdateHttpHttpLoopbackDetect(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.http_loopback_detect", bytes.NewReader(byt))
	return
}

// GetHttpIeSslBlockpageWorkaround gets the http.ie_ssl_blockpage_workaround value from the UTM
func GetHttpIeSslBlockpageWorkaround(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.ie_ssl_blockpage_workaround")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpIeSslBlockpageWorkaround PUTs the http.ie_ssl_blockpage_workaround value to the UTM
func UpdateHttpIeSslBlockpageWorkaround(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.ie_ssl_blockpage_workaround", bytes.NewReader(byt))
	return
}

// GetHttpLimitAdSsoInterfaces gets the http.limit_ad_sso_interfaces value from the UTM
func GetHttpLimitAdSsoInterfaces(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.limit_ad_sso_interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpLimitAdSsoInterfaces PUTs the http.limit_ad_sso_interfaces value to the UTM
func UpdateHttpLimitAdSsoInterfaces(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.limit_ad_sso_interfaces", bytes.NewReader(byt))
	return
}

// GetHttpLocalSiteList gets the http.local_site_list value from the UTM
func GetHttpLocalSiteList(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.local_site_list")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpLocalSiteList PUTs the http.local_site_list value to the UTM
func UpdateHttpLocalSiteList(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.local_site_list", bytes.NewReader(byt))
	return
}

// GetHttpMaxContentEncoding gets the http.max_content_encoding value from the UTM
func GetHttpMaxContentEncoding(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.max_content_encoding")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpMaxContentEncoding PUTs the http.max_content_encoding value to the UTM
func UpdateHttpMaxContentEncoding(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.max_content_encoding", bytes.NewReader(byt))
	return
}

// GetHttpMaxTempfileSize gets the http.max_tempfile_size value from the UTM
func GetHttpMaxTempfileSize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.max_tempfile_size")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpMaxTempfileSize PUTs the http.max_tempfile_size value to the UTM
func UpdateHttpMaxTempfileSize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.max_tempfile_size", bytes.NewReader(byt))
	return
}

// GetHttpMaxthreads gets the http.maxthreads value from the UTM
func GetHttpMaxthreads(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.maxthreads")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpMaxthreads PUTs the http.maxthreads value to the UTM
func UpdateHttpMaxthreads(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.maxthreads", bytes.NewReader(byt))
	return
}

// GetHttpMaxthreadsUnused gets the http.maxthreads_unused value from the UTM
func GetHttpMaxthreadsUnused(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.maxthreads_unused")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpMaxthreadsUnused PUTs the http.maxthreads_unused value to the UTM
func UpdateHttpMaxthreadsUnused(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.maxthreads_unused", bytes.NewReader(byt))
	return
}

// GetHttpModulepath gets the http.modulepath value from the UTM
func GetHttpModulepath(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.modulepath")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpModulepath PUTs the http.modulepath value to the UTM
func UpdateHttpModulepath(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.modulepath", bytes.NewReader(byt))
	return
}

// GetHttpModules gets the http.modules value from the UTM
func GetHttpModules(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.modules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpModules PUTs the http.modules value to the UTM
func UpdateHttpModules(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.modules", bytes.NewReader(byt))
	return
}

// GetHttpNoscancontent gets the http.noscancontent value from the UTM
func GetHttpNoscancontent(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.noscancontent")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpNoscancontent PUTs the http.noscancontent value to the UTM
func UpdateHttpNoscancontent(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.noscancontent", bytes.NewReader(byt))
	return
}

// GetHttpOpendirectoryKeytab gets the http.opendirectory_keytab value from the UTM
func GetHttpOpendirectoryKeytab(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.opendirectory_keytab")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpOpendirectoryKeytab PUTs the http.opendirectory_keytab value to the UTM
func UpdateHttpOpendirectoryKeytab(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.opendirectory_keytab", bytes.NewReader(byt))
	return
}

// GetHttpPacFile gets the http.pac_file value from the UTM
func GetHttpPacFile(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.pac_file")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPacFile PUTs the http.pac_file value to the UTM
func UpdateHttpPacFile(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.pac_file", bytes.NewReader(byt))
	return
}

// GetHttpParentProxyHost gets the http.parent_proxy_host value from the UTM
func GetHttpParentProxyHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.parent_proxy_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpParentProxyHost PUTs the http.parent_proxy_host value to the UTM
func UpdateHttpParentProxyHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.parent_proxy_host", bytes.NewReader(byt))
	return
}

// GetHttpParentProxyPort gets the http.parent_proxy_port value from the UTM
func GetHttpParentProxyPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.parent_proxy_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpParentProxyPort PUTs the http.parent_proxy_port value to the UTM
func UpdateHttpParentProxyPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.parent_proxy_port", bytes.NewReader(byt))
	return
}

// GetHttpParentProxyStatus gets the http.parent_proxy_status value from the UTM
func GetHttpParentProxyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.parent_proxy_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpParentProxyStatus PUTs the http.parent_proxy_status value to the UTM
func UpdateHttpParentProxyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.parent_proxy_status", bytes.NewReader(byt))
	return
}

// GetHttpPassthroughId gets the http.passthrough_id value from the UTM
func GetHttpPassthroughId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.passthrough_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPassthroughId PUTs the http.passthrough_id value to the UTM
func UpdateHttpPassthroughId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.passthrough_id", bytes.NewReader(byt))
	return
}

// GetHttpPharmingProtection gets the http.pharming_protection value from the UTM
func GetHttpPharmingProtection(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.pharming_protection")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPharmingProtection PUTs the http.pharming_protection value to the UTM
func UpdateHttpPharmingProtection(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.pharming_protection", bytes.NewReader(byt))
	return
}

// GetHttpPort gets the http.port value from the UTM
func GetHttpPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPort PUTs the http.port value to the UTM
func UpdateHttpPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.port", bytes.NewReader(byt))
	return
}

// GetHttpPortalCert gets the http.portal_cert value from the UTM
func GetHttpPortalCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.portal_cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPortalCert PUTs the http.portal_cert value to the UTM
func UpdateHttpPortalCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.portal_cert", bytes.NewReader(byt))
	return
}

// GetHttpPortalCertChain gets the http.portal_cert_chain value from the UTM
func GetHttpPortalCertChain(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.portal_cert_chain")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPortalCertChain PUTs the http.portal_cert_chain value to the UTM
func UpdateHttpPortalCertChain(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.portal_cert_chain", bytes.NewReader(byt))
	return
}

// GetHttpPortalDomain gets the http.portal_domain value from the UTM
func GetHttpPortalDomain(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.portal_domain")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPortalDomain PUTs the http.portal_domain value to the UTM
func UpdateHttpPortalDomain(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.portal_domain", bytes.NewReader(byt))
	return
}

// GetHttpPortalHosts gets the http.portal_hosts value from the UTM
func GetHttpPortalHosts(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.portal_hosts")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPortalHosts PUTs the http.portal_hosts value to the UTM
func UpdateHttpPortalHosts(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.portal_hosts", bytes.NewReader(byt))
	return
}

// GetHttpPortalUseCert gets the http.portal_use_cert value from the UTM
func GetHttpPortalUseCert(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.portal_use_cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpPortalUseCert PUTs the http.portal_use_cert value to the UTM
func UpdateHttpPortalUseCert(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.portal_use_cert", bytes.NewReader(byt))
	return
}

// GetHttpProceedCacheTimeout gets the http.proceed_cache_timeout value from the UTM
func GetHttpProceedCacheTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.proceed_cache_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpProceedCacheTimeout PUTs the http.proceed_cache_timeout value to the UTM
func UpdateHttpProceedCacheTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.proceed_cache_timeout", bytes.NewReader(byt))
	return
}

// GetHttpProfiles gets the http.profiles value from the UTM
func GetHttpProfiles(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.profiles")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpProfiles PUTs the http.profiles value to the UTM
func UpdateHttpProfiles(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.profiles", bytes.NewReader(byt))
	return
}

// GetHttpQuotaSliceTime gets the http.quota_slice_time value from the UTM
func GetHttpQuotaSliceTime(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.quota_slice_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpQuotaSliceTime PUTs the http.quota_slice_time value to the UTM
func UpdateHttpQuotaSliceTime(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.quota_slice_time", bytes.NewReader(byt))
	return
}

// GetHttpRemoveRequest gets the http.remove_request value from the UTM
func GetHttpRemoveRequest(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.remove_request")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpRemoveRequest PUTs the http.remove_request value to the UTM
func UpdateHttpRemoveRequest(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.remove_request", bytes.NewReader(byt))
	return
}

// GetHttpRemoveResponse gets the http.remove_response value from the UTM
func GetHttpRemoveResponse(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.remove_response")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpRemoveResponse PUTs the http.remove_response value to the UTM
func UpdateHttpRemoveResponse(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.remove_response", bytes.NewReader(byt))
	return
}

// GetHttpResponseTimeout gets the http.response_timeout value from the UTM
func GetHttpResponseTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.response_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpResponseTimeout PUTs the http.response_timeout value to the UTM
func UpdateHttpResponseTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.response_timeout", bytes.NewReader(byt))
	return
}

// GetHttpScLocalDb gets the http.sc_local_db value from the UTM
func GetHttpScLocalDb(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.sc_local_db")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpScLocalDb PUTs the http.sc_local_db value to the UTM
func UpdateHttpScLocalDb(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.sc_local_db", bytes.NewReader(byt))
	return
}

// GetHttpScanEppTraffic gets the http.scan_epp_traffic value from the UTM
func GetHttpScanEppTraffic(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.scan_epp_traffic")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpScanEppTraffic PUTs the http.scan_epp_traffic value to the UTM
func UpdateHttpScanEppTraffic(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.scan_epp_traffic", bytes.NewReader(byt))
	return
}

// GetHttpSearchdomain gets the http.searchdomain value from the UTM
func GetHttpSearchdomain(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.searchdomain")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpSearchdomain PUTs the http.searchdomain value to the UTM
func UpdateHttpSearchdomain(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.searchdomain", bytes.NewReader(byt))
	return
}

// GetHttpStrictHttp gets the http.strict_http value from the UTM
func GetHttpStrictHttp(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.strict_http")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpStrictHttp PUTs the http.strict_http value to the UTM
func UpdateHttpStrictHttp(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.strict_http", bytes.NewReader(byt))
	return
}

// GetHttpTlsciphersClient gets the http.tlsciphers_client value from the UTM
func GetHttpTlsciphersClient(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.tlsciphers_client")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTlsciphersClient PUTs the http.tlsciphers_client value to the UTM
func UpdateHttpTlsciphersClient(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.tlsciphers_client", bytes.NewReader(byt))
	return
}

// GetHttpTlsciphersServer gets the http.tlsciphers_server value from the UTM
func GetHttpTlsciphersServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.tlsciphers_server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTlsciphersServer PUTs the http.tlsciphers_server value to the UTM
func UpdateHttpTlsciphersServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.tlsciphers_server", bytes.NewReader(byt))
	return
}

// GetHttpTmpfsUsageMinMemsize gets the http.tmpfs_usage_min_memsize value from the UTM
func GetHttpTmpfsUsageMinMemsize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.tmpfs_usage_min_memsize")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTmpfsUsageMinMemsize PUTs the http.tmpfs_usage_min_memsize value to the UTM
func UpdateHttpTmpfsUsageMinMemsize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.tmpfs_usage_min_memsize", bytes.NewReader(byt))
	return
}

// GetHttpTransparentAuthTimeout gets the http.transparent_auth_timeout value from the UTM
func GetHttpTransparentAuthTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.transparent_auth_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTransparentAuthTimeout PUTs the http.transparent_auth_timeout value to the UTM
func UpdateHttpTransparentAuthTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.transparent_auth_timeout", bytes.NewReader(byt))
	return
}

// GetHttpTransparentDstSkip gets the http.transparent_dst_skip value from the UTM
func GetHttpTransparentDstSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.transparent_dst_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTransparentDstSkip PUTs the http.transparent_dst_skip value to the UTM
func UpdateHttpTransparentDstSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.transparent_dst_skip", bytes.NewReader(byt))
	return
}

// GetHttpTransparentSkipAutoPf gets the http.transparent_skip_auto_pf value from the UTM
func GetHttpTransparentSkipAutoPf(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.transparent_skip_auto_pf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTransparentSkipAutoPf PUTs the http.transparent_skip_auto_pf value to the UTM
func UpdateHttpTransparentSkipAutoPf(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.transparent_skip_auto_pf", bytes.NewReader(byt))
	return
}

// GetHttpTransparentSrcSkip gets the http.transparent_src_skip value from the UTM
func GetHttpTransparentSrcSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/http.transparent_src_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTransparentSrcSkip PUTs the http.transparent_src_skip value to the UTM
func UpdateHttpTransparentSrcSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.transparent_src_skip", bytes.NewReader(byt))
	return
}

// GetHttpTunnelTimeout gets the http.tunnel_timeout value from the UTM
func GetHttpTunnelTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/http.tunnel_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTunnelTimeout PUTs the http.tunnel_timeout value to the UTM
func UpdateHttpTunnelTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.tunnel_timeout", bytes.NewReader(byt))
	return
}

// GetHttpTunnelV6Timeout gets the http.tunnel_v6_timeout value from the UTM
func GetHttpTunnelV6Timeout(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/http.tunnel_v6_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpTunnelV6Timeout PUTs the http.tunnel_v6_timeout value to the UTM
func UpdateHttpTunnelV6Timeout(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.tunnel_v6_timeout", bytes.NewReader(byt))
	return
}

// GetHttpUndefercontent gets the http.undefercontent value from the UTM
func GetHttpUndefercontent(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.undefercontent")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUndefercontent PUTs the http.undefercontent value to the UTM
func UpdateHttpUndefercontent(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.undefercontent", bytes.NewReader(byt))
	return
}

// GetHttpUndeferextension gets the http.undeferextension value from the UTM
func GetHttpUndeferextension(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/http.undeferextension")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUndeferextension PUTs the http.undeferextension value to the UTM
func UpdateHttpUndeferextension(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.undeferextension", bytes.NewReader(byt))
	return
}

// GetHttpUrlFilteringRedirectUrl gets the http.url_filtering_redirect_url value from the UTM
func GetHttpUrlFilteringRedirectUrl(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/http.url_filtering_redirect_url")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUrlFilteringRedirectUrl PUTs the http.url_filtering_redirect_url value to the UTM
func UpdateHttpUrlFilteringRedirectUrl(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.url_filtering_redirect_url", bytes.NewReader(byt))
	return
}

// GetHttpUseConnectionInsteadofProxyconnection gets the http.use_connection_insteadof_proxyconnection value from the UTM
func GetHttpUseConnectionInsteadofProxyconnection(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.use_connection_insteadof_proxyconnection")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUseConnectionInsteadofProxyconnection PUTs the http.use_connection_insteadof_proxyconnection value to the UTM
func UpdateHttpUseConnectionInsteadofProxyconnection(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.use_connection_insteadof_proxyconnection", bytes.NewReader(byt))
	return
}

// GetHttpUseDstaddrForGeopiplookup gets the http.use_dstaddr_for_geopiplookup value from the UTM
func GetHttpUseDstaddrForGeopiplookup(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.use_dstaddr_for_geopiplookup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUseDstaddrForGeopiplookup PUTs the http.use_dstaddr_for_geopiplookup value to the UTM
func UpdateHttpUseDstaddrForGeopiplookup(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.use_dstaddr_for_geopiplookup", bytes.NewReader(byt))
	return
}

// GetHttpUseKrb5Adsso gets the http.use_krb5_adsso value from the UTM
func GetHttpUseKrb5Adsso(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.use_krb5_adsso")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUseKrb5Adsso PUTs the http.use_krb5_adsso value to the UTM
func UpdateHttpUseKrb5Adsso(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.use_krb5_adsso", bytes.NewReader(byt))
	return
}

// GetHttpUseSni gets the http.use_sni value from the UTM
func GetHttpUseSni(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.use_sni")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUseSni PUTs the http.use_sni value to the UTM
func UpdateHttpUseSni(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.use_sni", bytes.NewReader(byt))
	return
}

// GetHttpUseSxlUrid gets the http.use_sxl_urid value from the UTM
func GetHttpUseSxlUrid(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/http.use_sxl_urid")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateHttpUseSxlUrid PUTs the http.use_sxl_urid value to the UTM
func UpdateHttpUseSxlUrid(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/http.use_sxl_urid", bytes.NewReader(byt))
	return
}

// GetIcmpForward gets the icmp.forward value from the UTM
func GetIcmpForward(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.forward")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpForward PUTs the icmp.forward value to the UTM
func UpdateIcmpForward(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.forward", bytes.NewReader(byt))
	return
}

// GetIcmpInput gets the icmp.input value from the UTM
func GetIcmpInput(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.input")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpInput PUTs the icmp.input value to the UTM
func UpdateIcmpInput(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.input", bytes.NewReader(byt))
	return
}

// GetIcmpLogRedirect gets the icmp.log_redirect value from the UTM
func GetIcmpLogRedirect(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.log_redirect")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpLogRedirect PUTs the icmp.log_redirect value to the UTM
func UpdateIcmpLogRedirect(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.log_redirect", bytes.NewReader(byt))
	return
}

// GetIcmpPingForward gets the icmp.ping.forward value from the UTM
func GetIcmpPingForward(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.ping.forward")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpPingForward PUTs the icmp.ping.forward value to the UTM
func UpdateIcmpPingForward(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.ping.forward", bytes.NewReader(byt))
	return
}

// GetIcmpPingInput gets the icmp.ping.input value from the UTM
func GetIcmpPingInput(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.ping.input")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpPingInput PUTs the icmp.ping.input value to the UTM
func UpdateIcmpPingInput(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.ping.input", bytes.NewReader(byt))
	return
}

// GetIcmpPingOutput gets the icmp.ping.output value from the UTM
func GetIcmpPingOutput(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.ping.output")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpPingOutput PUTs the icmp.ping.output value to the UTM
func UpdateIcmpPingOutput(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.ping.output", bytes.NewReader(byt))
	return
}

// GetIcmpSecure gets the icmp.secure value from the UTM
func GetIcmpSecure(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.secure")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpSecure PUTs the icmp.secure value to the UTM
func UpdateIcmpSecure(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.secure", bytes.NewReader(byt))
	return
}

// GetIcmpTracerouteForward gets the icmp.traceroute.forward value from the UTM
func GetIcmpTracerouteForward(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.traceroute.forward")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpTracerouteForward PUTs the icmp.traceroute.forward value to the UTM
func UpdateIcmpTracerouteForward(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.traceroute.forward", bytes.NewReader(byt))
	return
}

// GetIcmpTracerouteInput gets the icmp.traceroute.input value from the UTM
func GetIcmpTracerouteInput(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/icmp.traceroute.input")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIcmpTracerouteInput PUTs the icmp.traceroute.input value to the UTM
func UpdateIcmpTracerouteInput(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/icmp.traceroute.input", bytes.NewReader(byt))
	return
}

// GetIdentForward gets the ident.forward value from the UTM
func GetIdentForward(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ident.forward")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIdentForward PUTs the ident.forward value to the UTM
func UpdateIdentForward(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ident.forward", bytes.NewReader(byt))
	return
}

// GetIdentResponse gets the ident.response value from the UTM
func GetIdentResponse(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ident.response")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIdentResponse PUTs the ident.response value to the UTM
func UpdateIdentResponse(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ident.response", bytes.NewReader(byt))
	return
}

// GetIdentStatus gets the ident.status value from the UTM
func GetIdentStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ident.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIdentStatus PUTs the ident.status value to the UTM
func UpdateIdentStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ident.status", bytes.NewReader(byt))
	return
}

// GetInterfacesAdvancedArpAnnounce gets the interfaces.advanced.arp_announce value from the UTM
func GetInterfacesAdvancedArpAnnounce(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/interfaces.advanced.arp_announce")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateInterfacesAdvancedArpAnnounce PUTs the interfaces.advanced.arp_announce value to the UTM
func UpdateInterfacesAdvancedArpAnnounce(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/interfaces.advanced.arp_announce", bytes.NewReader(byt))
	return
}

// GetInterfacesAdvancedArpIgnore gets the interfaces.advanced.arp_ignore value from the UTM
func GetInterfacesAdvancedArpIgnore(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/interfaces.advanced.arp_ignore")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateInterfacesAdvancedArpIgnore PUTs the interfaces.advanced.arp_ignore value to the UTM
func UpdateInterfacesAdvancedArpIgnore(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/interfaces.advanced.arp_ignore", bytes.NewReader(byt))
	return
}

// GetInterfacesAdvancedDefaultMetric gets the interfaces.advanced.default_metric value from the UTM
func GetInterfacesAdvancedDefaultMetric(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/interfaces.advanced.default_metric")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateInterfacesAdvancedDefaultMetric PUTs the interfaces.advanced.default_metric value to the UTM
func UpdateInterfacesAdvancedDefaultMetric(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/interfaces.advanced.default_metric", bytes.NewReader(byt))
	return
}

// GetInterfacesInterfaces gets the interfaces.interfaces value from the UTM
func GetInterfacesInterfaces(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/interfaces.interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateInterfacesInterfaces PUTs the interfaces.interfaces value to the UTM
func UpdateInterfacesInterfaces(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/interfaces.interfaces", bytes.NewReader(byt))
	return
}

// GetIpsDnsServers gets the ips.dns_servers value from the UTM
func GetIpsDnsServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.dns_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsDnsServers PUTs the ips.dns_servers value to the UTM
func UpdateIpsDnsServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.dns_servers", bytes.NewReader(byt))
	return
}

// GetIpsEngine gets the ips.engine value from the UTM
func GetIpsEngine(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ips.engine")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsEngine PUTs the ips.engine value to the UTM
func UpdateIpsEngine(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.engine", bytes.NewReader(byt))
	return
}

// GetIpsExceptions gets the ips.exceptions value from the UTM
func GetIpsExceptions(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ips.exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsExceptions PUTs the ips.exceptions value to the UTM
func UpdateIpsExceptions(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.exceptions", bytes.NewReader(byt))
	return
}

// GetIpsFailopen gets the ips.failopen value from the UTM
func GetIpsFailopen(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ips.failopen")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsFailopen PUTs the ips.failopen value to the UTM
func UpdateIpsFailopen(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.failopen", bytes.NewReader(byt))
	return
}

// GetIpsFileBasedRules gets the ips.file_based_rules value from the UTM
func GetIpsFileBasedRules(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ips.file_based_rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsFileBasedRules PUTs the ips.file_based_rules value to the UTM
func UpdateIpsFileBasedRules(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.file_based_rules", bytes.NewReader(byt))
	return
}

// GetIpsGroups gets the ips.groups value from the UTM
func GetIpsGroups(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ips.groups")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsGroups PUTs the ips.groups value to the UTM
func UpdateIpsGroups(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.groups", bytes.NewReader(byt))
	return
}

// GetIpsHttpServers gets the ips.http_servers value from the UTM
func GetIpsHttpServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.http_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsHttpServers PUTs the ips.http_servers value to the UTM
func UpdateIpsHttpServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.http_servers", bytes.NewReader(byt))
	return
}

// GetIpsIpsfbAlertInterval gets the ips.ipsfb.alert_interval value from the UTM
func GetIpsIpsfbAlertInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ips.ipsfb.alert_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsIpsfbAlertInterval PUTs the ips.ipsfb.alert_interval value to the UTM
func UpdateIpsIpsfbAlertInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.ipsfb.alert_interval", bytes.NewReader(byt))
	return
}

// GetIpsIpsfbConfigInterval gets the ips.ipsfb.config_interval value from the UTM
func GetIpsIpsfbConfigInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ips.ipsfb.config_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsIpsfbConfigInterval PUTs the ips.ipsfb.config_interval value to the UTM
func UpdateIpsIpsfbConfigInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.ipsfb.config_interval", bytes.NewReader(byt))
	return
}

// GetIpsIpsfbDebug gets the ips.ipsfb.debug value from the UTM
func GetIpsIpsfbDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ips.ipsfb.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsIpsfbDebug PUTs the ips.ipsfb.debug value to the UTM
func UpdateIpsIpsfbDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.ipsfb.debug", bytes.NewReader(byt))
	return
}

// GetIpsLocalNetworks gets the ips.local_networks value from the UTM
func GetIpsLocalNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ips.local_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsLocalNetworks PUTs the ips.local_networks value to the UTM
func UpdateIpsLocalNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.local_networks", bytes.NewReader(byt))
	return
}

// GetIpsNumInstances gets the ips.num_instances value from the UTM
func GetIpsNumInstances(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.num_instances")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsNumInstances PUTs the ips.num_instances value to the UTM
func UpdateIpsNumInstances(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.num_instances", bytes.NewReader(byt))
	return
}

// GetIpsPatternChannel gets the ips.pattern_channel value from the UTM
func GetIpsPatternChannel(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ips.pattern_channel")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsPatternChannel PUTs the ips.pattern_channel value to the UTM
func UpdateIpsPatternChannel(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.pattern_channel", bytes.NewReader(byt))
	return
}

// GetIpsPolicy gets the ips.policy value from the UTM
func GetIpsPolicy(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ips.policy")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsPolicy PUTs the ips.policy value to the UTM
func UpdateIpsPolicy(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.policy", bytes.NewReader(byt))
	return
}

// GetIpsQueueLength gets the ips.queue_length value from the UTM
func GetIpsQueueLength(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ips.queue_length")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsQueueLength PUTs the ips.queue_length value to the UTM
func UpdateIpsQueueLength(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.queue_length", bytes.NewReader(byt))
	return
}

// GetIpsQueueThreshold gets the ips.queue_threshold value from the UTM
func GetIpsQueueThreshold(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.queue_threshold")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsQueueThreshold PUTs the ips.queue_threshold value to the UTM
func UpdateIpsQueueThreshold(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.queue_threshold", bytes.NewReader(byt))
	return
}

// GetIpsReloadMethod gets the ips.reload_method value from the UTM
func GetIpsReloadMethod(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ips.reload_method")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsReloadMethod PUTs the ips.reload_method value to the UTM
func UpdateIpsReloadMethod(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.reload_method", bytes.NewReader(byt))
	return
}

// GetIpsRestartPolicy gets the ips.restart_policy value from the UTM
func GetIpsRestartPolicy(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ips.restart_policy")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsRestartPolicy PUTs the ips.restart_policy value to the UTM
func UpdateIpsRestartPolicy(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.restart_policy", bytes.NewReader(byt))
	return
}

// GetIpsRuleModifiers gets the ips.rule_modifiers value from the UTM
func GetIpsRuleModifiers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.rule_modifiers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsRuleModifiers PUTs the ips.rule_modifiers value to the UTM
func UpdateIpsRuleModifiers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.rule_modifiers", bytes.NewReader(byt))
	return
}

// GetIpsRules gets the ips.rules value from the UTM
func GetIpsRules(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsRules PUTs the ips.rules value to the UTM
func UpdateIpsRules(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.rules", bytes.NewReader(byt))
	return
}

// GetIpsSkipAcks gets the ips.skip_acks value from the UTM
func GetIpsSkipAcks(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ips.skip_acks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSkipAcks PUTs the ips.skip_acks value to the UTM
func UpdateIpsSkipAcks(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.skip_acks", bytes.NewReader(byt))
	return
}

// GetIpsSmtpServers gets the ips.smtp_servers value from the UTM
func GetIpsSmtpServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.smtp_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSmtpServers PUTs the ips.smtp_servers value to the UTM
func UpdateIpsSmtpServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.smtp_servers", bytes.NewReader(byt))
	return
}

// GetIpsSnortsettingsMaxQueuedBytes gets the ips.snortsettings.max_queued_bytes value from the UTM
func GetIpsSnortsettingsMaxQueuedBytes(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.snortsettings.max_queued_bytes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSnortsettingsMaxQueuedBytes PUTs the ips.snortsettings.max_queued_bytes value to the UTM
func UpdateIpsSnortsettingsMaxQueuedBytes(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.snortsettings.max_queued_bytes", bytes.NewReader(byt))
	return
}

// GetIpsSnortsettingsMaxQueuedSegs gets the ips.snortsettings.max_queued_segs value from the UTM
func GetIpsSnortsettingsMaxQueuedSegs(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.snortsettings.max_queued_segs")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSnortsettingsMaxQueuedSegs PUTs the ips.snortsettings.max_queued_segs value to the UTM
func UpdateIpsSnortsettingsMaxQueuedSegs(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.snortsettings.max_queued_segs", bytes.NewReader(byt))
	return
}

// GetIpsSnortsettingsMaxTcp gets the ips.snortsettings.max_tcp value from the UTM
func GetIpsSnortsettingsMaxTcp(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.snortsettings.max_tcp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSnortsettingsMaxTcp PUTs the ips.snortsettings.max_tcp value to the UTM
func UpdateIpsSnortsettingsMaxTcp(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.snortsettings.max_tcp", bytes.NewReader(byt))
	return
}

// GetIpsSnortsettingsMaxUdp gets the ips.snortsettings.max_udp value from the UTM
func GetIpsSnortsettingsMaxUdp(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.snortsettings.max_udp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSnortsettingsMaxUdp PUTs the ips.snortsettings.max_udp value to the UTM
func UpdateIpsSnortsettingsMaxUdp(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.snortsettings.max_udp", bytes.NewReader(byt))
	return
}

// GetIpsSnortsettingsMemcap gets the ips.snortsettings.memcap value from the UTM
func GetIpsSnortsettingsMemcap(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.snortsettings.memcap")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSnortsettingsMemcap PUTs the ips.snortsettings.memcap value to the UTM
func UpdateIpsSnortsettingsMemcap(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.snortsettings.memcap", bytes.NewReader(byt))
	return
}

// GetIpsSnortsettingsSearchMethod gets the ips.snortsettings.search_method value from the UTM
func GetIpsSnortsettingsSearchMethod(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ips.snortsettings.search_method")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSnortsettingsSearchMethod PUTs the ips.snortsettings.search_method value to the UTM
func UpdateIpsSnortsettingsSearchMethod(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.snortsettings.search_method", bytes.NewReader(byt))
	return
}

// GetIpsSqlServers gets the ips.sql_servers value from the UTM
func GetIpsSqlServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ips.sql_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsSqlServers PUTs the ips.sql_servers value to the UTM
func UpdateIpsSqlServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.sql_servers", bytes.NewReader(byt))
	return
}

// GetIpsStatus gets the ips.status value from the UTM
func GetIpsStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ips.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsStatus PUTs the ips.status value to the UTM
func UpdateIpsStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ips.status", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedCrlAutoFetching gets the ipsec.advanced.crl_auto_fetching value from the UTM
func GetIpsecAdvancedCrlAutoFetching(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.crl_auto_fetching")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedCrlAutoFetching PUTs the ipsec.advanced.crl_auto_fetching value to the UTM
func UpdateIpsecAdvancedCrlAutoFetching(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.crl_auto_fetching", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedCrlStrictPolicy gets the ipsec.advanced.crl_strict_policy value from the UTM
func GetIpsecAdvancedCrlStrictPolicy(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.crl_strict_policy")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedCrlStrictPolicy PUTs the ipsec.advanced.crl_strict_policy value to the UTM
func UpdateIpsecAdvancedCrlStrictPolicy(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.crl_strict_policy", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedDeadPeerDetection gets the ipsec.advanced.dead_peer_detection value from the UTM
func GetIpsecAdvancedDeadPeerDetection(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.dead_peer_detection")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedDeadPeerDetection PUTs the ipsec.advanced.dead_peer_detection value to the UTM
func UpdateIpsecAdvancedDeadPeerDetection(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.dead_peer_detection", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedIkeDebug gets the ipsec.advanced.ike_debug value from the UTM
func GetIpsecAdvancedIkeDebug(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.ike_debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedIkeDebug PUTs the ipsec.advanced.ike_debug value to the UTM
func UpdateIpsecAdvancedIkeDebug(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.ike_debug", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedIkePort gets the ipsec.advanced.ike_port value from the UTM
func GetIpsecAdvancedIkePort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.ike_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedIkePort PUTs the ipsec.advanced.ike_port value to the UTM
func UpdateIpsecAdvancedIkePort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.ike_port", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedMetric gets the ipsec.advanced.metric value from the UTM
func GetIpsecAdvancedMetric(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.metric")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedMetric PUTs the ipsec.advanced.metric value to the UTM
func UpdateIpsecAdvancedMetric(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.metric", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedNatTraversal gets the ipsec.advanced.nat_traversal value from the UTM
func GetIpsecAdvancedNatTraversal(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.nat_traversal")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedNatTraversal PUTs the ipsec.advanced.nat_traversal value to the UTM
func UpdateIpsecAdvancedNatTraversal(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.nat_traversal", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedNatTraversalKeepalive gets the ipsec.advanced.nat_traversal_keepalive value from the UTM
func GetIpsecAdvancedNatTraversalKeepalive(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.nat_traversal_keepalive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedNatTraversalKeepalive PUTs the ipsec.advanced.nat_traversal_keepalive value to the UTM
func UpdateIpsecAdvancedNatTraversalKeepalive(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.nat_traversal_keepalive", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedProbePsk gets the ipsec.advanced.probe_psk value from the UTM
func GetIpsecAdvancedProbePsk(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.probe_psk")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedProbePsk PUTs the ipsec.advanced.probe_psk value to the UTM
func UpdateIpsecAdvancedProbePsk(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.probe_psk", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedPskVpnId gets the ipsec.advanced.psk_vpn_id value from the UTM
func GetIpsecAdvancedPskVpnId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.psk_vpn_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedPskVpnId PUTs the ipsec.advanced.psk_vpn_id value to the UTM
func UpdateIpsecAdvancedPskVpnId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.psk_vpn_id", bytes.NewReader(byt))
	return
}

// GetIpsecAdvancedPskVpnIdType gets the ipsec.advanced.psk_vpn_id_type value from the UTM
func GetIpsecAdvancedPskVpnIdType(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipsec.advanced.psk_vpn_id_type")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecAdvancedPskVpnIdType PUTs the ipsec.advanced.psk_vpn_id_type value to the UTM
func UpdateIpsecAdvancedPskVpnIdType(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.advanced.psk_vpn_id_type", bytes.NewReader(byt))
	return
}

// GetIpsecConnections gets the ipsec.connections value from the UTM
func GetIpsecConnections(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ipsec.connections")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecConnections PUTs the ipsec.connections value to the UTM
func UpdateIpsecConnections(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.connections", bytes.NewReader(byt))
	return
}

// GetIpsecLocalRsa gets the ipsec.local_rsa value from the UTM
func GetIpsecLocalRsa(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipsec.local_rsa")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecLocalRsa PUTs the ipsec.local_rsa value to the UTM
func UpdateIpsecLocalRsa(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.local_rsa", bytes.NewReader(byt))
	return
}

// GetIpsecLocalX509 gets the ipsec.local_x509 value from the UTM
func GetIpsecLocalX509(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipsec.local_x509")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecLocalX509 PUTs the ipsec.local_x509 value to the UTM
func UpdateIpsecLocalX509(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.local_x509", bytes.NewReader(byt))
	return
}

// GetIpsecStatus gets the ipsec.status value from the UTM
func GetIpsecStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipsec.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpsecStatus PUTs the ipsec.status value to the UTM
func UpdateIpsecStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipsec.status", bytes.NewReader(byt))
	return
}

// GetIpv6AdvancedHopLimit gets the ipv6.advanced.hop_limit value from the UTM
func GetIpv6AdvancedHopLimit(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ipv6.advanced.hop_limit")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6AdvancedHopLimit PUTs the ipv6.advanced.hop_limit value to the UTM
func UpdateIpv6AdvancedHopLimit(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.advanced.hop_limit", bytes.NewReader(byt))
	return
}

// GetIpv6AdvancedMaxInterval gets the ipv6.advanced.max_interval value from the UTM
func GetIpv6AdvancedMaxInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ipv6.advanced.max_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6AdvancedMaxInterval PUTs the ipv6.advanced.max_interval value to the UTM
func UpdateIpv6AdvancedMaxInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.advanced.max_interval", bytes.NewReader(byt))
	return
}

// GetIpv6AdvancedMinInterval gets the ipv6.advanced.min_interval value from the UTM
func GetIpv6AdvancedMinInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ipv6.advanced.min_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6AdvancedMinInterval PUTs the ipv6.advanced.min_interval value to the UTM
func UpdateIpv6AdvancedMinInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.advanced.min_interval", bytes.NewReader(byt))
	return
}

// GetIpv6AdvancedPreference gets the ipv6.advanced.preference value from the UTM
func GetIpv6AdvancedPreference(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.advanced.preference")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6AdvancedPreference PUTs the ipv6.advanced.preference value to the UTM
func UpdateIpv6AdvancedPreference(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.advanced.preference", bytes.NewReader(byt))
	return
}

// GetIpv6AdvancedReachableTime gets the ipv6.advanced.reachable_time value from the UTM
func GetIpv6AdvancedReachableTime(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ipv6.advanced.reachable_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6AdvancedReachableTime PUTs the ipv6.advanced.reachable_time value to the UTM
func UpdateIpv6AdvancedReachableTime(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.advanced.reachable_time", bytes.NewReader(byt))
	return
}

// GetIpv6AdvancedRetransTime gets the ipv6.advanced.retrans_time value from the UTM
func GetIpv6AdvancedRetransTime(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/ipv6.advanced.retrans_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6AdvancedRetransTime PUTs the ipv6.advanced.retrans_time value to the UTM
func UpdateIpv6AdvancedRetransTime(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.advanced.retrans_time", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerAuthentication gets the ipv6.broker.authentication value from the UTM
func GetIpv6BrokerAuthentication(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.authentication")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerAuthentication PUTs the ipv6.broker.authentication value to the UTM
func UpdateIpv6BrokerAuthentication(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.authentication", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerInterface gets the ipv6.broker.interface value from the UTM
func GetIpv6BrokerInterface(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.interface")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerInterface PUTs the ipv6.broker.interface value to the UTM
func UpdateIpv6BrokerInterface(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.interface", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerPassword gets the ipv6.broker.password value from the UTM
func GetIpv6BrokerPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerPassword PUTs the ipv6.broker.password value to the UTM
func UpdateIpv6BrokerPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.password", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerProtocol gets the ipv6.broker.protocol value from the UTM
func GetIpv6BrokerProtocol(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.protocol")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerProtocol PUTs the ipv6.broker.protocol value to the UTM
func UpdateIpv6BrokerProtocol(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.protocol", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerServer gets the ipv6.broker.server value from the UTM
func GetIpv6BrokerServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerServer PUTs the ipv6.broker.server value to the UTM
func UpdateIpv6BrokerServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.server", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerStatus gets the ipv6.broker.status value from the UTM
func GetIpv6BrokerStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerStatus PUTs the ipv6.broker.status value to the UTM
func UpdateIpv6BrokerStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.status", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerTunnelId gets the ipv6.broker.tunnel_id value from the UTM
func GetIpv6BrokerTunnelId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.tunnel_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerTunnelId PUTs the ipv6.broker.tunnel_id value to the UTM
func UpdateIpv6BrokerTunnelId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.tunnel_id", bytes.NewReader(byt))
	return
}

// GetIpv6BrokerUsername gets the ipv6.broker.username value from the UTM
func GetIpv6BrokerUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.broker.username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6BrokerUsername PUTs the ipv6.broker.username value to the UTM
func UpdateIpv6BrokerUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.broker.username", bytes.NewReader(byt))
	return
}

// GetIpv6Nat64Address gets the ipv6.nat64.address value from the UTM
func GetIpv6Nat64Address(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.nat64.address")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Nat64Address PUTs the ipv6.nat64.address value to the UTM
func UpdateIpv6Nat64Address(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.nat64.address", bytes.NewReader(byt))
	return
}

// GetIpv6Nat64Dns64V6Only gets the ipv6.nat64.dns64_v6only value from the UTM
func GetIpv6Nat64Dns64V6Only(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.nat64.dns64_v6only")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Nat64Dns64V6Only PUTs the ipv6.nat64.dns64_v6only value to the UTM
func UpdateIpv6Nat64Dns64V6Only(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.nat64.dns64_v6only", bytes.NewReader(byt))
	return
}

// GetIpv6Nat64Prefix gets the ipv6.nat64.prefix value from the UTM
func GetIpv6Nat64Prefix(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.nat64.prefix")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Nat64Prefix PUTs the ipv6.nat64.prefix value to the UTM
func UpdateIpv6Nat64Prefix(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.nat64.prefix", bytes.NewReader(byt))
	return
}

// GetIpv6Nat64Status gets the ipv6.nat64.status value from the UTM
func GetIpv6Nat64Status(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.nat64.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Nat64Status PUTs the ipv6.nat64.status value to the UTM
func UpdateIpv6Nat64Status(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.nat64.status", bytes.NewReader(byt))
	return
}

// GetIpv6Prefer gets the ipv6.prefer value from the UTM
func GetIpv6Prefer(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.prefer")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Prefer PUTs the ipv6.prefer value to the UTM
func UpdateIpv6Prefer(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.prefer", bytes.NewReader(byt))
	return
}

// GetIpv6Prefixes gets the ipv6.prefixes value from the UTM
func GetIpv6Prefixes(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/ipv6.prefixes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Prefixes PUTs the ipv6.prefixes value to the UTM
func UpdateIpv6Prefixes(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.prefixes", bytes.NewReader(byt))
	return
}

// GetIpv6Renumbering gets the ipv6.renumbering value from the UTM
func GetIpv6Renumbering(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.renumbering")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Renumbering PUTs the ipv6.renumbering value to the UTM
func UpdateIpv6Renumbering(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.renumbering", bytes.NewReader(byt))
	return
}

// GetIpv6Six2FourInterface gets the ipv6.six2four.interface value from the UTM
func GetIpv6Six2FourInterface(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.six2four.interface")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Six2FourInterface PUTs the ipv6.six2four.interface value to the UTM
func UpdateIpv6Six2FourInterface(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.six2four.interface", bytes.NewReader(byt))
	return
}

// GetIpv6Six2FourServer gets the ipv6.six2four.server value from the UTM
func GetIpv6Six2FourServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ipv6.six2four.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Six2FourServer PUTs the ipv6.six2four.server value to the UTM
func UpdateIpv6Six2FourServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.six2four.server", bytes.NewReader(byt))
	return
}

// GetIpv6Six2FourStatus gets the ipv6.six2four.status value from the UTM
func GetIpv6Six2FourStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.six2four.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Six2FourStatus PUTs the ipv6.six2four.status value to the UTM
func UpdateIpv6Six2FourStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.six2four.status", bytes.NewReader(byt))
	return
}

// GetIpv6Status gets the ipv6.status value from the UTM
func GetIpv6Status(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ipv6.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateIpv6Status PUTs the ipv6.status value to the UTM
func UpdateIpv6Status(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ipv6.status", bytes.NewReader(byt))
	return
}

// GetLicensingActiveIps gets the licensing.active_ips value from the UTM
func GetLicensingActiveIps(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/licensing.active_ips")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLicensingActiveIps PUTs the licensing.active_ips value to the UTM
func UpdateLicensingActiveIps(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/licensing.active_ips", bytes.NewReader(byt))
	return
}

// GetLicensingLicense gets the licensing.license value from the UTM
func GetLicensingLicense(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/licensing.license")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLicensingLicense PUTs the licensing.license value to the UTM
func UpdateLicensingLicense(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/licensing.license", bytes.NewReader(byt))
	return
}

// GetLicensingUserLimitExceeded gets the licensing.user_limit_exceeded value from the UTM
func GetLicensingUserLimitExceeded(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/licensing.user_limit_exceeded")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLicensingUserLimitExceeded PUTs the licensing.user_limit_exceeded value to the UTM
func UpdateLicensingUserLimitExceeded(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/licensing.user_limit_exceeded", bytes.NewReader(byt))
	return
}

// GetLinkAggregationGroups gets the link_aggregation.groups value from the UTM
func GetLinkAggregationGroups(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/link_aggregation.groups")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLinkAggregationGroups PUTs the link_aggregation.groups value to the UTM
func UpdateLinkAggregationGroups(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/link_aggregation.groups", bytes.NewReader(byt))
	return
}

// GetLoadbalanceHttpErrorCode gets the loadbalance.http_error_code value from the UTM
func GetLoadbalanceHttpErrorCode(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/loadbalance.http_error_code")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLoadbalanceHttpErrorCode PUTs the loadbalance.http_error_code value to the UTM
func UpdateLoadbalanceHttpErrorCode(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/loadbalance.http_error_code", bytes.NewReader(byt))
	return
}

// GetLoadbalanceRules gets the loadbalance.rules value from the UTM
func GetLoadbalanceRules(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/loadbalance.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLoadbalanceRules PUTs the loadbalance.rules value to the UTM
func UpdateLoadbalanceRules(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/loadbalance.rules", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalActionOne gets the logfiles.local.action_one value from the UTM
func GetLogfilesLocalActionOne(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.action_one")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalActionOne PUTs the logfiles.local.action_one value to the UTM
func UpdateLogfilesLocalActionOne(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.action_one", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalActionThree gets the logfiles.local.action_three value from the UTM
func GetLogfilesLocalActionThree(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.action_three")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalActionThree PUTs the logfiles.local.action_three value to the UTM
func UpdateLogfilesLocalActionThree(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.action_three", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalActionTwo gets the logfiles.local.action_two value from the UTM
func GetLogfilesLocalActionTwo(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.action_two")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalActionTwo PUTs the logfiles.local.action_two value to the UTM
func UpdateLogfilesLocalActionTwo(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.action_two", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalDeleteAfterDays gets the logfiles.local.delete_after_days value from the UTM
func GetLogfilesLocalDeleteAfterDays(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.delete_after_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalDeleteAfterDays PUTs the logfiles.local.delete_after_days value to the UTM
func UpdateLogfilesLocalDeleteAfterDays(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.delete_after_days", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalPercentageOne gets the logfiles.local.percentage_one value from the UTM
func GetLogfilesLocalPercentageOne(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.percentage_one")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalPercentageOne PUTs the logfiles.local.percentage_one value to the UTM
func UpdateLogfilesLocalPercentageOne(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.percentage_one", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalPercentageThree gets the logfiles.local.percentage_three value from the UTM
func GetLogfilesLocalPercentageThree(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.percentage_three")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalPercentageThree PUTs the logfiles.local.percentage_three value to the UTM
func UpdateLogfilesLocalPercentageThree(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.percentage_three", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalPercentageTwo gets the logfiles.local.percentage_two value from the UTM
func GetLogfilesLocalPercentageTwo(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.percentage_two")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalPercentageTwo PUTs the logfiles.local.percentage_two value to the UTM
func UpdateLogfilesLocalPercentageTwo(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.percentage_two", bytes.NewReader(byt))
	return
}

// GetLogfilesLocalStatus gets the logfiles.local.status value from the UTM
func GetLogfilesLocalStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/logfiles.local.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesLocalStatus PUTs the logfiles.local.status value to the UTM
func UpdateLogfilesLocalStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.local.status", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteFtpService gets the logfiles.remote.ftp_service value from the UTM
func GetLogfilesRemoteFtpService(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.ftp_service")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteFtpService PUTs the logfiles.remote.ftp_service value to the UTM
func UpdateLogfilesRemoteFtpService(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.ftp_service", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteHost gets the logfiles.remote.host value from the UTM
func GetLogfilesRemoteHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteHost PUTs the logfiles.remote.host value to the UTM
func UpdateLogfilesRemoteHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.host", bytes.NewReader(byt))
	return
}

// GetLogfilesRemotePass gets the logfiles.remote.pass value from the UTM
func GetLogfilesRemotePass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.pass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemotePass PUTs the logfiles.remote.pass value to the UTM
func UpdateLogfilesRemotePass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.pass", bytes.NewReader(byt))
	return
}

// GetLogfilesRemotePath gets the logfiles.remote.path value from the UTM
func GetLogfilesRemotePath(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.path")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemotePath PUTs the logfiles.remote.path value to the UTM
func UpdateLogfilesRemotePath(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.path", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteSmbWorkgroup gets the logfiles.remote.smb_workgroup value from the UTM
func GetLogfilesRemoteSmbWorkgroup(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.smb_workgroup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteSmbWorkgroup PUTs the logfiles.remote.smb_workgroup value to the UTM
func UpdateLogfilesRemoteSmbWorkgroup(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.smb_workgroup", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteSmtpAddress gets the logfiles.remote.smtp_address value from the UTM
func GetLogfilesRemoteSmtpAddress(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.smtp_address")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteSmtpAddress PUTs the logfiles.remote.smtp_address value to the UTM
func UpdateLogfilesRemoteSmtpAddress(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.smtp_address", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteStatus gets the logfiles.remote.status value from the UTM
func GetLogfilesRemoteStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteStatus PUTs the logfiles.remote.status value to the UTM
func UpdateLogfilesRemoteStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.status", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteType gets the logfiles.remote.type value from the UTM
func GetLogfilesRemoteType(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.type")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteType PUTs the logfiles.remote.type value to the UTM
func UpdateLogfilesRemoteType(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.type", bytes.NewReader(byt))
	return
}

// GetLogfilesRemoteUser gets the logfiles.remote.user value from the UTM
func GetLogfilesRemoteUser(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/logfiles.remote.user")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateLogfilesRemoteUser PUTs the logfiles.remote.user value to the UTM
func UpdateLogfilesRemoteUser(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/logfiles.remote.user", bytes.NewReader(byt))
	return
}

// GetMasqRules gets the masq.rules value from the UTM
func GetMasqRules(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/masq.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMasqRules PUTs the masq.rules value to the UTM
func UpdateMasqRules(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/masq.rules", bytes.NewReader(byt))
	return
}

// GetMigrationAccessToken gets the migration.access_token value from the UTM
func GetMigrationAccessToken(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/migration.access_token")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMigrationAccessToken PUTs the migration.access_token value to the UTM
func UpdateMigrationAccessToken(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/migration.access_token", bytes.NewReader(byt))
	return
}

// GetMigrationLocalOverride gets the migration.local_override value from the UTM
func GetMigrationLocalOverride(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/migration.local_override")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMigrationLocalOverride PUTs the migration.local_override value to the UTM
func UpdateMigrationLocalOverride(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/migration.local_override", bytes.NewReader(byt))
	return
}

// GetMigrationRefreshToken gets the migration.refresh_token value from the UTM
func GetMigrationRefreshToken(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/migration.refresh_token")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMigrationRefreshToken PUTs the migration.refresh_token value to the UTM
func UpdateMigrationRefreshToken(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/migration.refresh_token", bytes.NewReader(byt))
	return
}

// GetMigrationTabVisibility gets the migration.tab_visibility value from the UTM
func GetMigrationTabVisibility(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/migration.tab_visibility")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMigrationTabVisibility PUTs the migration.tab_visibility value to the UTM
func UpdateMigrationTabVisibility(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/migration.tab_visibility", bytes.NewReader(byt))
	return
}

// GetMigrationToolsetVersion gets the migration.toolset_version value from the UTM
func GetMigrationToolsetVersion(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/migration.toolset_version")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMigrationToolsetVersion PUTs the migration.toolset_version value to the UTM
func UpdateMigrationToolsetVersion(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/migration.toolset_version", bytes.NewReader(byt))
	return
}

// GetMigrationUtmVersion gets the migration.utm_version value from the UTM
func GetMigrationUtmVersion(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/migration.utm_version")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMigrationUtmVersion PUTs the migration.utm_version value to the UTM
func UpdateMigrationUtmVersion(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/migration.utm_version", bytes.NewReader(byt))
	return
}

// GetMobileControlCa gets the mobile_control.ca value from the UTM
func GetMobileControlCa(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.ca")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlCa PUTs the mobile_control.ca value to the UTM
func UpdateMobileControlCa(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.ca", bytes.NewReader(byt))
	return
}

// GetMobileControlConfigCisco gets the mobile_control.config.cisco value from the UTM
func GetMobileControlConfigCisco(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.config.cisco")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlConfigCisco PUTs the mobile_control.config.cisco value to the UTM
func UpdateMobileControlConfigCisco(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.config.cisco", bytes.NewReader(byt))
	return
}

// GetMobileControlConfigEapMethod gets the mobile_control.config.eap_method value from the UTM
func GetMobileControlConfigEapMethod(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.config.eap_method")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlConfigEapMethod PUTs the mobile_control.config.eap_method value to the UTM
func UpdateMobileControlConfigEapMethod(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.config.eap_method", bytes.NewReader(byt))
	return
}

// GetMobileControlConfigForcePush gets the mobile_control.config.force_push value from the UTM
func GetMobileControlConfigForcePush(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.config.force_push")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlConfigForcePush PUTs the mobile_control.config.force_push value to the UTM
func UpdateMobileControlConfigForcePush(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.config.force_push", bytes.NewReader(byt))
	return
}

// GetMobileControlConfigL2Tp gets the mobile_control.config.l2tp value from the UTM
func GetMobileControlConfigL2Tp(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.config.l2tp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlConfigL2Tp PUTs the mobile_control.config.l2tp value to the UTM
func UpdateMobileControlConfigL2Tp(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.config.l2tp", bytes.NewReader(byt))
	return
}

// GetMobileControlConfigWifiNetworks gets the mobile_control.config.wifi_networks value from the UTM
func GetMobileControlConfigWifiNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/mobile_control.config.wifi_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlConfigWifiNetworks PUTs the mobile_control.config.wifi_networks value to the UTM
func UpdateMobileControlConfigWifiNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.config.wifi_networks", bytes.NewReader(byt))
	return
}

// GetMobileControlCustomer gets the mobile_control.customer value from the UTM
func GetMobileControlCustomer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.customer")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlCustomer PUTs the mobile_control.customer value to the UTM
func UpdateMobileControlCustomer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.customer", bytes.NewReader(byt))
	return
}

// GetMobileControlDebug gets the mobile_control.debug value from the UTM
func GetMobileControlDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlDebug PUTs the mobile_control.debug value to the UTM
func UpdateMobileControlDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.debug", bytes.NewReader(byt))
	return
}

// GetMobileControlNacCisco gets the mobile_control.nac.cisco value from the UTM
func GetMobileControlNacCisco(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.cisco")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacCisco PUTs the mobile_control.nac.cisco value to the UTM
func UpdateMobileControlNacCisco(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.cisco", bytes.NewReader(byt))
	return
}

// GetMobileControlNacDenyAllVpn gets the mobile_control.nac.deny_all_vpn value from the UTM
func GetMobileControlNacDenyAllVpn(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.deny_all_vpn")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacDenyAllVpn PUTs the mobile_control.nac.deny_all_vpn value to the UTM
func UpdateMobileControlNacDenyAllVpn(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.deny_all_vpn", bytes.NewReader(byt))
	return
}

// GetMobileControlNacL2Tp gets the mobile_control.nac.l2tp value from the UTM
func GetMobileControlNacL2Tp(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.l2tp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacL2Tp PUTs the mobile_control.nac.l2tp value to the UTM
func UpdateMobileControlNacL2Tp(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.l2tp", bytes.NewReader(byt))
	return
}

// GetMobileControlNacMacsAllowed gets the mobile_control.nac.macs_allowed value from the UTM
func GetMobileControlNacMacsAllowed(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.macs_allowed")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacMacsAllowed PUTs the mobile_control.nac.macs_allowed value to the UTM
func UpdateMobileControlNacMacsAllowed(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.macs_allowed", bytes.NewReader(byt))
	return
}

// GetMobileControlNacMacsDenied gets the mobile_control.nac.macs_denied value from the UTM
func GetMobileControlNacMacsDenied(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.macs_denied")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacMacsDenied PUTs the mobile_control.nac.macs_denied value to the UTM
func UpdateMobileControlNacMacsDenied(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.macs_denied", bytes.NewReader(byt))
	return
}

// GetMobileControlNacPollInterval gets the mobile_control.nac.poll_interval value from the UTM
func GetMobileControlNacPollInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.poll_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacPollInterval PUTs the mobile_control.nac.poll_interval value to the UTM
func UpdateMobileControlNacPollInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.poll_interval", bytes.NewReader(byt))
	return
}

// GetMobileControlNacUsersDenied gets the mobile_control.nac.users_denied value from the UTM
func GetMobileControlNacUsersDenied(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.users_denied")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacUsersDenied PUTs the mobile_control.nac.users_denied value to the UTM
func UpdateMobileControlNacUsersDenied(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.users_denied", bytes.NewReader(byt))
	return
}

// GetMobileControlNacWifiNetworks gets the mobile_control.nac.wifi_networks value from the UTM
func GetMobileControlNacWifiNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/mobile_control.nac.wifi_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlNacWifiNetworks PUTs the mobile_control.nac.wifi_networks value to the UTM
func UpdateMobileControlNacWifiNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.nac.wifi_networks", bytes.NewReader(byt))
	return
}

// GetMobileControlPassword gets the mobile_control.password value from the UTM
func GetMobileControlPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlPassword PUTs the mobile_control.password value to the UTM
func UpdateMobileControlPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.password", bytes.NewReader(byt))
	return
}

// GetMobileControlServer gets the mobile_control.server value from the UTM
func GetMobileControlServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlServer PUTs the mobile_control.server value to the UTM
func UpdateMobileControlServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.server", bytes.NewReader(byt))
	return
}

// GetMobileControlStatus gets the mobile_control.status value from the UTM
func GetMobileControlStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/mobile_control.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlStatus PUTs the mobile_control.status value to the UTM
func UpdateMobileControlStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.status", bytes.NewReader(byt))
	return
}

// GetMobileControlUsername gets the mobile_control.username value from the UTM
func GetMobileControlUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/mobile_control.username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateMobileControlUsername PUTs the mobile_control.username value to the UTM
func UpdateMobileControlUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/mobile_control.username", bytes.NewReader(byt))
	return
}

// GetNatRules gets the nat.rules value from the UTM
func GetNatRules(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/nat.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNatRules PUTs the nat.rules value to the UTM
func UpdateNatRules(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/nat.rules", bytes.NewReader(byt))
	return
}

// GetNotificationsDeviceInfo gets the notifications.device_info value from the UTM
func GetNotificationsDeviceInfo(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/notifications.device_info")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsDeviceInfo PUTs the notifications.device_info value to the UTM
func UpdateNotificationsDeviceInfo(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.device_info", bytes.NewReader(byt))
	return
}

// GetNotificationsLimiting gets the notifications.limiting value from the UTM
func GetNotificationsLimiting(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/notifications.limiting")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsLimiting PUTs the notifications.limiting value to the UTM
func UpdateNotificationsLimiting(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.limiting", bytes.NewReader(byt))
	return
}

// GetNotificationsOverlay gets the notifications.overlay value from the UTM
func GetNotificationsOverlay(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/notifications.overlay")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsOverlay PUTs the notifications.overlay value to the UTM
func UpdateNotificationsOverlay(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.overlay", bytes.NewReader(byt))
	return
}

// GetNotificationsRebootReason gets the notifications.reboot_reason value from the UTM
func GetNotificationsRebootReason(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/notifications.reboot_reason")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsRebootReason PUTs the notifications.reboot_reason value to the UTM
func UpdateNotificationsRebootReason(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.reboot_reason", bytes.NewReader(byt))
	return
}

// GetNotificationsRecipients gets the notifications.recipients value from the UTM
func GetNotificationsRecipients(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/notifications.recipients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsRecipients PUTs the notifications.recipients value to the UTM
func UpdateNotificationsRecipients(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.recipients", bytes.NewReader(byt))
	return
}

// GetNotificationsSender gets the notifications.sender value from the UTM
func GetNotificationsSender(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/notifications.sender")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSender PUTs the notifications.sender value to the UTM
func UpdateNotificationsSender(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.sender", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpAuthentication gets the notifications.smtp.authentication value from the UTM
func GetNotificationsSmtpAuthentication(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.authentication")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpAuthentication PUTs the notifications.smtp.authentication value to the UTM
func UpdateNotificationsSmtpAuthentication(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.authentication", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpPassword gets the notifications.smtp.password value from the UTM
func GetNotificationsSmtpPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpPassword PUTs the notifications.smtp.password value to the UTM
func UpdateNotificationsSmtpPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.password", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpPort gets the notifications.smtp.port value from the UTM
func GetNotificationsSmtpPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpPort PUTs the notifications.smtp.port value to the UTM
func UpdateNotificationsSmtpPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.port", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpServer gets the notifications.smtp.server value from the UTM
func GetNotificationsSmtpServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpServer PUTs the notifications.smtp.server value to the UTM
func UpdateNotificationsSmtpServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.server", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpStatus gets the notifications.smtp.status value from the UTM
func GetNotificationsSmtpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpStatus PUTs the notifications.smtp.status value to the UTM
func UpdateNotificationsSmtpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.status", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpTls gets the notifications.smtp.tls value from the UTM
func GetNotificationsSmtpTls(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.tls")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpTls PUTs the notifications.smtp.tls value to the UTM
func UpdateNotificationsSmtpTls(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.tls", bytes.NewReader(byt))
	return
}

// GetNotificationsSmtpUsername gets the notifications.smtp.username value from the UTM
func GetNotificationsSmtpUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/notifications.smtp.username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNotificationsSmtpUsername PUTs the notifications.smtp.username value to the UTM
func UpdateNotificationsSmtpUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/notifications.smtp.username", bytes.NewReader(byt))
	return
}

// GetNtpAllowedNetworks gets the ntp.allowed_networks value from the UTM
func GetNtpAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ntp.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNtpAllowedNetworks PUTs the ntp.allowed_networks value to the UTM
func UpdateNtpAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ntp.allowed_networks", bytes.NewReader(byt))
	return
}

// GetNtpServers gets the ntp.servers value from the UTM
func GetNtpServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ntp.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNtpServers PUTs the ntp.servers value to the UTM
func UpdateNtpServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ntp.servers", bytes.NewReader(byt))
	return
}

// GetNtpStatus gets the ntp.status value from the UTM
func GetNtpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ntp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateNtpStatus PUTs the ntp.status value to the UTM
func UpdateNtpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ntp.status", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedBlockInvalidCtPackets gets the packetfilter.advanced.block_invalid_ct_packets value from the UTM
func GetPacketfilterAdvancedBlockInvalidCtPackets(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.block_invalid_ct_packets")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedBlockInvalidCtPackets PUTs the packetfilter.advanced.block_invalid_ct_packets value to the UTM
func UpdatePacketfilterAdvancedBlockInvalidCtPackets(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.block_invalid_ct_packets", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedCheckPacketLength gets the packetfilter.advanced.check_packet_length value from the UTM
func GetPacketfilterAdvancedCheckPacketLength(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.check_packet_length")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedCheckPacketLength PUTs the packetfilter.advanced.check_packet_length value to the UTM
func UpdatePacketfilterAdvancedCheckPacketLength(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.check_packet_length", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedConntrackHelpers gets the packetfilter.advanced.conntrack_helpers value from the UTM
func GetPacketfilterAdvancedConntrackHelpers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.conntrack_helpers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedConntrackHelpers PUTs the packetfilter.advanced.conntrack_helpers value to the UTM
func UpdatePacketfilterAdvancedConntrackHelpers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.conntrack_helpers", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedFtpPorts gets the packetfilter.advanced.ftp_ports value from the UTM
func GetPacketfilterAdvancedFtpPorts(c sophos.ClientInterface) (val []int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.ftp_ports")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedFtpPorts PUTs the packetfilter.advanced.ftp_ports value to the UTM
func UpdatePacketfilterAdvancedFtpPorts(c sophos.ClientInterface, val []int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.ftp_ports", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedLogBroadcasts gets the packetfilter.advanced.log_broadcasts value from the UTM
func GetPacketfilterAdvancedLogBroadcasts(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.log_broadcasts")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedLogBroadcasts PUTs the packetfilter.advanced.log_broadcasts value to the UTM
func UpdatePacketfilterAdvancedLogBroadcasts(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.log_broadcasts", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedLogDnsRequests gets the packetfilter.advanced.log_dns_requests value from the UTM
func GetPacketfilterAdvancedLogDnsRequests(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.log_dns_requests")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedLogDnsRequests PUTs the packetfilter.advanced.log_dns_requests value to the UTM
func UpdatePacketfilterAdvancedLogDnsRequests(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.log_dns_requests", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedLogFtpData gets the packetfilter.advanced.log_ftp_data value from the UTM
func GetPacketfilterAdvancedLogFtpData(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.log_ftp_data")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedLogFtpData PUTs the packetfilter.advanced.log_ftp_data value to the UTM
func UpdatePacketfilterAdvancedLogFtpData(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.log_ftp_data", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedLogInvalid gets the packetfilter.advanced.log_invalid value from the UTM
func GetPacketfilterAdvancedLogInvalid(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.log_invalid")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedLogInvalid PUTs the packetfilter.advanced.log_invalid value to the UTM
func UpdatePacketfilterAdvancedLogInvalid(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.log_invalid", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedLogMcast gets the packetfilter.advanced.log_mcast value from the UTM
func GetPacketfilterAdvancedLogMcast(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.log_mcast")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedLogMcast PUTs the packetfilter.advanced.log_mcast value to the UTM
func UpdatePacketfilterAdvancedLogMcast(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.log_mcast", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedLogStrictTcpState gets the packetfilter.advanced.log_strict_tcp_state value from the UTM
func GetPacketfilterAdvancedLogStrictTcpState(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.log_strict_tcp_state")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedLogStrictTcpState PUTs the packetfilter.advanced.log_strict_tcp_state value to the UTM
func UpdatePacketfilterAdvancedLogStrictTcpState(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.log_strict_tcp_state", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedNoErrorReplay gets the packetfilter.advanced.no_error_replay value from the UTM
func GetPacketfilterAdvancedNoErrorReplay(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.no_error_replay")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedNoErrorReplay PUTs the packetfilter.advanced.no_error_replay value to the UTM
func UpdatePacketfilterAdvancedNoErrorReplay(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.no_error_replay", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedOptimizeIpset gets the packetfilter.advanced.optimize.ipset value from the UTM
func GetPacketfilterAdvancedOptimizeIpset(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.optimize.ipset")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedOptimizeIpset PUTs the packetfilter.advanced.optimize.ipset value to the UTM
func UpdatePacketfilterAdvancedOptimizeIpset(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.optimize.ipset", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedOptimizePorts gets the packetfilter.advanced.optimize.ports value from the UTM
func GetPacketfilterAdvancedOptimizePorts(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.optimize.ports")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedOptimizePorts PUTs the packetfilter.advanced.optimize.ports value to the UTM
func UpdatePacketfilterAdvancedOptimizePorts(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.optimize.ports", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedSpoofProtection gets the packetfilter.advanced.spoof_protection value from the UTM
func GetPacketfilterAdvancedSpoofProtection(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.spoof_protection")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedSpoofProtection PUTs the packetfilter.advanced.spoof_protection value to the UTM
func UpdatePacketfilterAdvancedSpoofProtection(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.spoof_protection", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedStrictTcpState gets the packetfilter.advanced.strict_tcp_state value from the UTM
func GetPacketfilterAdvancedStrictTcpState(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.strict_tcp_state")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedStrictTcpState PUTs the packetfilter.advanced.strict_tcp_state value to the UTM
func UpdatePacketfilterAdvancedStrictTcpState(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.strict_tcp_state", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedTcpMaxRetrans gets the packetfilter.advanced.tcp_max_retrans value from the UTM
func GetPacketfilterAdvancedTcpMaxRetrans(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.tcp_max_retrans")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedTcpMaxRetrans PUTs the packetfilter.advanced.tcp_max_retrans value to the UTM
func UpdatePacketfilterAdvancedTcpMaxRetrans(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.tcp_max_retrans", bytes.NewReader(byt))
	return
}

// GetPacketfilterAdvancedTcpWindowScaling gets the packetfilter.advanced.tcp_window_scaling value from the UTM
func GetPacketfilterAdvancedTcpWindowScaling(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/packetfilter.advanced.tcp_window_scaling")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterAdvancedTcpWindowScaling PUTs the packetfilter.advanced.tcp_window_scaling value to the UTM
func UpdatePacketfilterAdvancedTcpWindowScaling(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.advanced.tcp_window_scaling", bytes.NewReader(byt))
	return
}

// GetPacketfilterRules gets the packetfilter.rules value from the UTM
func GetPacketfilterRules(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterRules PUTs the packetfilter.rules value to the UTM
func UpdatePacketfilterRules(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.rules", bytes.NewReader(byt))
	return
}

// GetPacketfilterRulesAuto gets the packetfilter.rules_auto value from the UTM
func GetPacketfilterRulesAuto(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.rules_auto")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterRulesAuto PUTs the packetfilter.rules_auto value to the UTM
func UpdatePacketfilterRulesAuto(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.rules_auto", bytes.NewReader(byt))
	return
}

// GetPacketfilterRulesBack gets the packetfilter.rules_back value from the UTM
func GetPacketfilterRulesBack(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.rules_back")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterRulesBack PUTs the packetfilter.rules_back value to the UTM
func UpdatePacketfilterRulesBack(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.rules_back", bytes.NewReader(byt))
	return
}

// GetPacketfilterRulesFront gets the packetfilter.rules_front value from the UTM
func GetPacketfilterRulesFront(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/packetfilter.rules_front")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterRulesFront PUTs the packetfilter.rules_front value to the UTM
func UpdatePacketfilterRulesFront(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.rules_front", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackGenericTimeout gets the packetfilter.timeouts.ip_conntrack_generic_timeout value from the UTM
func GetPacketfilterTimeoutsIpConntrackGenericTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_generic_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackGenericTimeout PUTs the packetfilter.timeouts.ip_conntrack_generic_timeout value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackGenericTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_generic_timeout", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackIcmpTimeout gets the packetfilter.timeouts.ip_conntrack_icmp_timeout value from the UTM
func GetPacketfilterTimeoutsIpConntrackIcmpTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_icmp_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackIcmpTimeout PUTs the packetfilter.timeouts.ip_conntrack_icmp_timeout value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackIcmpTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_icmp_timeout", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutClose gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_close value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutClose(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_close")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutClose PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_close value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutClose(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_close", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutCloseWait gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_close_wait value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutCloseWait(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_close_wait")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutCloseWait PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_close_wait value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutCloseWait(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_close_wait", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutEstablished gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_established value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutEstablished(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_established")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutEstablished PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_established value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutEstablished(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_established", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutFinWait gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_fin_wait value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutFinWait(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_fin_wait")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutFinWait PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_fin_wait value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutFinWait(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_fin_wait", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutLastAck gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_last_ack value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutLastAck(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_last_ack")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutLastAck PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_last_ack value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutLastAck(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_last_ack", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutMaxRetrans gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_max_retrans value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutMaxRetrans(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_max_retrans")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutMaxRetrans PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_max_retrans value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutMaxRetrans(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_max_retrans", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutSynRecv gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_recv value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutSynRecv(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_recv")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutSynRecv PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_recv value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutSynRecv(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_recv", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutSynSent gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_sent value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutSynSent(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_sent")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutSynSent PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_sent value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutSynSent(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_sent", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutTimeWait gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_time_wait value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutTimeWait(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_time_wait")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutTimeWait PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_time_wait value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutTimeWait(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_time_wait", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackUdpTimeout gets the packetfilter.timeouts.ip_conntrack_udp_timeout value from the UTM
func GetPacketfilterTimeoutsIpConntrackUdpTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_udp_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackUdpTimeout PUTs the packetfilter.timeouts.ip_conntrack_udp_timeout value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackUdpTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_udp_timeout", bytes.NewReader(byt))
	return
}

// GetPacketfilterTimeoutsIpConntrackUdpTimeoutStream gets the packetfilter.timeouts.ip_conntrack_udp_timeout_stream value from the UTM
func GetPacketfilterTimeoutsIpConntrackUdpTimeoutStream(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/packetfilter.timeouts.ip_conntrack_udp_timeout_stream")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackUdpTimeoutStream PUTs the packetfilter.timeouts.ip_conntrack_udp_timeout_stream value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackUdpTimeoutStream(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/packetfilter.timeouts.ip_conntrack_udp_timeout_stream", bytes.NewReader(byt))
	return
}

// GetPasswdLoginuserClearpass gets the passwd.loginuser.clearpass value from the UTM
func GetPasswdLoginuserClearpass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/passwd.loginuser.clearpass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePasswdLoginuserClearpass PUTs the passwd.loginuser.clearpass value to the UTM
func UpdatePasswdLoginuserClearpass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/passwd.loginuser.clearpass", bytes.NewReader(byt))
	return
}

// GetPasswdLoginuserCryptpass gets the passwd.loginuser.cryptpass value from the UTM
func GetPasswdLoginuserCryptpass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/passwd.loginuser.cryptpass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePasswdLoginuserCryptpass PUTs the passwd.loginuser.cryptpass value to the UTM
func UpdatePasswdLoginuserCryptpass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/passwd.loginuser.cryptpass", bytes.NewReader(byt))
	return
}

// GetPasswdRootClearpass gets the passwd.root.clearpass value from the UTM
func GetPasswdRootClearpass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/passwd.root.clearpass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePasswdRootClearpass PUTs the passwd.root.clearpass value to the UTM
func UpdatePasswdRootClearpass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/passwd.root.clearpass", bytes.NewReader(byt))
	return
}

// GetPasswdRootCryptpass gets the passwd.root.cryptpass value from the UTM
func GetPasswdRootCryptpass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/passwd.root.cryptpass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePasswdRootCryptpass PUTs the passwd.root.cryptpass value to the UTM
func UpdatePasswdRootCryptpass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/passwd.root.cryptpass", bytes.NewReader(byt))
	return
}

// GetPdfPaper gets the pdf.paper value from the UTM
func GetPdfPaper(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pdf.paper")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePdfPaper PUTs the pdf.paper value to the UTM
func UpdatePdfPaper(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pdf.paper", bytes.NewReader(byt))
	return
}

// GetPimSmAutoPfOut gets the pim_sm.auto_pf_out value from the UTM
func GetPimSmAutoPfOut(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pim_sm.auto_pf_out")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmAutoPfOut PUTs the pim_sm.auto_pf_out value to the UTM
func UpdatePimSmAutoPfOut(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.auto_pf_out", bytes.NewReader(byt))
	return
}

// GetPimSmAutoPfrule gets the pim_sm.auto_pfrule value from the UTM
func GetPimSmAutoPfrule(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pim_sm.auto_pfrule")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmAutoPfrule PUTs the pim_sm.auto_pfrule value to the UTM
func UpdatePimSmAutoPfrule(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.auto_pfrule", bytes.NewReader(byt))
	return
}

// GetPimSmDebug gets the pim_sm.debug value from the UTM
func GetPimSmDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pim_sm.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmDebug PUTs the pim_sm.debug value to the UTM
func UpdatePimSmDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.debug", bytes.NewReader(byt))
	return
}

// GetPimSmInterfaces gets the pim_sm.interfaces value from the UTM
func GetPimSmInterfaces(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pim_sm.interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmInterfaces PUTs the pim_sm.interfaces value to the UTM
func UpdatePimSmInterfaces(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.interfaces", bytes.NewReader(byt))
	return
}

// GetPimSmRpRouters gets the pim_sm.rp_routers value from the UTM
func GetPimSmRpRouters(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pim_sm.rp_routers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmRpRouters PUTs the pim_sm.rp_routers value to the UTM
func UpdatePimSmRpRouters(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.rp_routers", bytes.NewReader(byt))
	return
}

// GetPimSmSptSwitchBytes gets the pim_sm.spt_switch_bytes value from the UTM
func GetPimSmSptSwitchBytes(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/pim_sm.spt_switch_bytes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmSptSwitchBytes PUTs the pim_sm.spt_switch_bytes value to the UTM
func UpdatePimSmSptSwitchBytes(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.spt_switch_bytes", bytes.NewReader(byt))
	return
}

// GetPimSmSptSwitchStatus gets the pim_sm.spt_switch_status value from the UTM
func GetPimSmSptSwitchStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pim_sm.spt_switch_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmSptSwitchStatus PUTs the pim_sm.spt_switch_status value to the UTM
func UpdatePimSmSptSwitchStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.spt_switch_status", bytes.NewReader(byt))
	return
}

// GetPimSmStatus gets the pim_sm.status value from the UTM
func GetPimSmStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pim_sm.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePimSmStatus PUTs the pim_sm.status value to the UTM
func UpdatePimSmStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pim_sm.status", bytes.NewReader(byt))
	return
}

// GetPop3AllowedClients gets the pop3.allowed_clients value from the UTM
func GetPop3AllowedClients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pop3.allowed_clients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3AllowedClients PUTs the pop3.allowed_clients value to the UTM
func UpdatePop3AllowedClients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.allowed_clients", bytes.NewReader(byt))
	return
}

// GetPop3AllowedServers gets the pop3.allowed_servers value from the UTM
func GetPop3AllowedServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/pop3.allowed_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3AllowedServers PUTs the pop3.allowed_servers value to the UTM
func UpdatePop3AllowedServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.allowed_servers", bytes.NewReader(byt))
	return
}

// GetPop3CffAsMarker gets the pop3.cff_as_marker value from the UTM
func GetPop3CffAsMarker(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.cff_as_marker")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3CffAsMarker PUTs the pop3.cff_as_marker value to the UTM
func UpdatePop3CffAsMarker(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.cff_as_marker", bytes.NewReader(byt))
	return
}

// GetPop3CffAv gets the pop3.cff_av value from the UTM
func GetPop3CffAv(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.cff_av")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3CffAv PUTs the pop3.cff_av value to the UTM
func UpdatePop3CffAv(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.cff_av", bytes.NewReader(byt))
	return
}

// GetPop3CffAvAction gets the pop3.cff_av_action value from the UTM
func GetPop3CffAvAction(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.cff_av_action")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3CffAvAction PUTs the pop3.cff_av_action value to the UTM
func UpdatePop3CffAvAction(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.cff_av_action", bytes.NewReader(byt))
	return
}

// GetPop3CffAvEngines gets the pop3.cff_av_engines value from the UTM
func GetPop3CffAvEngines(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.cff_av_engines")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3CffAvEngines PUTs the pop3.cff_av_engines value to the UTM
func UpdatePop3CffAvEngines(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.cff_av_engines", bytes.NewReader(byt))
	return
}

// GetPop3CffFileExtensions gets the pop3.cff_file_extensions value from the UTM
func GetPop3CffFileExtensions(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/pop3.cff_file_extensions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3CffFileExtensions PUTs the pop3.cff_file_extensions value to the UTM
func UpdatePop3CffFileExtensions(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.cff_file_extensions", bytes.NewReader(byt))
	return
}

// GetPop3DirectlyDeleteQuarantined gets the pop3.directly_delete_quarantined value from the UTM
func GetPop3DirectlyDeleteQuarantined(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.directly_delete_quarantined")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3DirectlyDeleteQuarantined PUTs the pop3.directly_delete_quarantined value to the UTM
func UpdatePop3DirectlyDeleteQuarantined(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.directly_delete_quarantined", bytes.NewReader(byt))
	return
}

// GetPop3Exceptions gets the pop3.exceptions value from the UTM
func GetPop3Exceptions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pop3.exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3Exceptions PUTs the pop3.exceptions value to the UTM
func UpdatePop3Exceptions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.exceptions", bytes.NewReader(byt))
	return
}

// GetPop3KnownServers gets the pop3.known_servers value from the UTM
func GetPop3KnownServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pop3.known_servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3KnownServers PUTs the pop3.known_servers value to the UTM
func UpdatePop3KnownServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.known_servers", bytes.NewReader(byt))
	return
}

// GetPop3MaxMessageSize gets the pop3.max_message_size value from the UTM
func GetPop3MaxMessageSize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/pop3.max_message_size")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3MaxMessageSize PUTs the pop3.max_message_size value to the UTM
func UpdatePop3MaxMessageSize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.max_message_size", bytes.NewReader(byt))
	return
}

// GetPop3PrefetchingInterval gets the pop3.prefetching.interval value from the UTM
func GetPop3PrefetchingInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/pop3.prefetching.interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3PrefetchingInterval PUTs the pop3.prefetching.interval value to the UTM
func UpdatePop3PrefetchingInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.prefetching.interval", bytes.NewReader(byt))
	return
}

// GetPop3PrefetchingOptimizeStorage gets the pop3.prefetching.optimize_storage value from the UTM
func GetPop3PrefetchingOptimizeStorage(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.prefetching.optimize_storage")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3PrefetchingOptimizeStorage PUTs the pop3.prefetching.optimize_storage value to the UTM
func UpdatePop3PrefetchingOptimizeStorage(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.prefetching.optimize_storage", bytes.NewReader(byt))
	return
}

// GetPop3PrefetchingStatus gets the pop3.prefetching.status value from the UTM
func GetPop3PrefetchingStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.prefetching.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3PrefetchingStatus PUTs the pop3.prefetching.status value to the UTM
func UpdatePop3PrefetchingStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.prefetching.status", bytes.NewReader(byt))
	return
}

// GetPop3PrefetchingStorageMinHoldDays gets the pop3.prefetching.storage_min_hold_days value from the UTM
func GetPop3PrefetchingStorageMinHoldDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/pop3.prefetching.storage_min_hold_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3PrefetchingStorageMinHoldDays PUTs the pop3.prefetching.storage_min_hold_days value to the UTM
func UpdatePop3PrefetchingStorageMinHoldDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.prefetching.storage_min_hold_days", bytes.NewReader(byt))
	return
}

// GetPop3QuarantineUnscannable gets the pop3.quarantine_unscannable value from the UTM
func GetPop3QuarantineUnscannable(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.quarantine_unscannable")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3QuarantineUnscannable PUTs the pop3.quarantine_unscannable value to the UTM
func UpdatePop3QuarantineUnscannable(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.quarantine_unscannable", bytes.NewReader(byt))
	return
}

// GetPop3SandboxMaxFilesizeMb gets the pop3.sandbox_max_filesize_mb value from the UTM
func GetPop3SandboxMaxFilesizeMb(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/pop3.sandbox_max_filesize_mb")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3SandboxMaxFilesizeMb PUTs the pop3.sandbox_max_filesize_mb value to the UTM
func UpdatePop3SandboxMaxFilesizeMb(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.sandbox_max_filesize_mb", bytes.NewReader(byt))
	return
}

// GetPop3SandboxScanStatus gets the pop3.sandbox_scan_status value from the UTM
func GetPop3SandboxScanStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.sandbox_scan_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3SandboxScanStatus PUTs the pop3.sandbox_scan_status value to the UTM
func UpdatePop3SandboxScanStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.sandbox_scan_status", bytes.NewReader(byt))
	return
}

// GetPop3ScanTls gets the pop3.scan_tls value from the UTM
func GetPop3ScanTls(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.scan_tls")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3ScanTls PUTs the pop3.scan_tls value to the UTM
func UpdatePop3ScanTls(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.scan_tls", bytes.NewReader(byt))
	return
}

// GetPop3SenderBlacklist gets the pop3.sender_blacklist value from the UTM
func GetPop3SenderBlacklist(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pop3.sender_blacklist")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3SenderBlacklist PUTs the pop3.sender_blacklist value to the UTM
func UpdatePop3SenderBlacklist(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.sender_blacklist", bytes.NewReader(byt))
	return
}

// GetPop3Spam gets the pop3.spam value from the UTM
func GetPop3Spam(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.spam")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3Spam PUTs the pop3.spam value to the UTM
func UpdatePop3Spam(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.spam", bytes.NewReader(byt))
	return
}

// GetPop3SpamExpressions gets the pop3.spam_expressions value from the UTM
func GetPop3SpamExpressions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pop3.spam_expressions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3SpamExpressions PUTs the pop3.spam_expressions value to the UTM
func UpdatePop3SpamExpressions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.spam_expressions", bytes.NewReader(byt))
	return
}

// GetPop3Spamplus gets the pop3.spamplus value from the UTM
func GetPop3Spamplus(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.spamplus")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3Spamplus PUTs the pop3.spamplus value to the UTM
func UpdatePop3Spamplus(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.spamplus", bytes.NewReader(byt))
	return
}

// GetPop3Spamstatus gets the pop3.spamstatus value from the UTM
func GetPop3Spamstatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.spamstatus")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3Spamstatus PUTs the pop3.spamstatus value to the UTM
func UpdatePop3Spamstatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.spamstatus", bytes.NewReader(byt))
	return
}

// GetPop3Status gets the pop3.status value from the UTM
func GetPop3Status(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3Status PUTs the pop3.status value to the UTM
func UpdatePop3Status(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.status", bytes.NewReader(byt))
	return
}

// GetPop3TlsCert gets the pop3.tls_cert value from the UTM
func GetPop3TlsCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.tls_cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3TlsCert PUTs the pop3.tls_cert value to the UTM
func UpdatePop3TlsCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.tls_cert", bytes.NewReader(byt))
	return
}

// GetPop3TransparentSkip gets the pop3.transparent_skip value from the UTM
func GetPop3TransparentSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/pop3.transparent_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3TransparentSkip PUTs the pop3.transparent_skip value to the UTM
func UpdatePop3TransparentSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.transparent_skip", bytes.NewReader(byt))
	return
}

// GetPop3TransparentSkipAutoPf gets the pop3.transparent_skip_auto_pf value from the UTM
func GetPop3TransparentSkipAutoPf(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/pop3.transparent_skip_auto_pf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3TransparentSkipAutoPf PUTs the pop3.transparent_skip_auto_pf value to the UTM
func UpdatePop3TransparentSkipAutoPf(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.transparent_skip_auto_pf", bytes.NewReader(byt))
	return
}

// GetPop3UserCharset gets the pop3.user_charset value from the UTM
func GetPop3UserCharset(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/pop3.user_charset")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePop3UserCharset PUTs the pop3.user_charset value to the UTM
func UpdatePop3UserCharset(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/pop3.user_charset", bytes.NewReader(byt))
	return
}

// GetPortalAllowAnyUser gets the portal.allow_any_user value from the UTM
func GetPortalAllowAnyUser(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/portal.allow_any_user")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalAllowAnyUser PUTs the portal.allow_any_user value to the UTM
func UpdatePortalAllowAnyUser(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.allow_any_user", bytes.NewReader(byt))
	return
}

// GetPortalAllowedNetworks gets the portal.allowed_networks value from the UTM
func GetPortalAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/portal.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalAllowedNetworks PUTs the portal.allowed_networks value to the UTM
func UpdatePortalAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.allowed_networks", bytes.NewReader(byt))
	return
}

// GetPortalAllowedUsers gets the portal.allowed_users value from the UTM
func GetPortalAllowedUsers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/portal.allowed_users")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalAllowedUsers PUTs the portal.allowed_users value to the UTM
func UpdatePortalAllowedUsers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.allowed_users", bytes.NewReader(byt))
	return
}

// GetPortalHideItems gets the portal.hide_items value from the UTM
func GetPortalHideItems(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/portal.hide_items")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalHideItems PUTs the portal.hide_items value to the UTM
func UpdatePortalHideItems(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.hide_items", bytes.NewReader(byt))
	return
}

// GetPortalHostname gets the portal.hostname value from the UTM
func GetPortalHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/portal.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalHostname PUTs the portal.hostname value to the UTM
func UpdatePortalHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.hostname", bytes.NewReader(byt))
	return
}

// GetPortalInterfaceAddress gets the portal.interface_address value from the UTM
func GetPortalInterfaceAddress(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/portal.interface_address")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalInterfaceAddress PUTs the portal.interface_address value to the UTM
func UpdatePortalInterfaceAddress(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.interface_address", bytes.NewReader(byt))
	return
}

// GetPortalLanguage gets the portal.language value from the UTM
func GetPortalLanguage(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/portal.language")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalLanguage PUTs the portal.language value to the UTM
func UpdatePortalLanguage(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.language", bytes.NewReader(byt))
	return
}

// GetPortalPersistentCookies gets the portal.persistent_cookies value from the UTM
func GetPortalPersistentCookies(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/portal.persistent_cookies")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalPersistentCookies PUTs the portal.persistent_cookies value to the UTM
func UpdatePortalPersistentCookies(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.persistent_cookies", bytes.NewReader(byt))
	return
}

// GetPortalPort gets the portal.port value from the UTM
func GetPortalPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/portal.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalPort PUTs the portal.port value to the UTM
func UpdatePortalPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.port", bytes.NewReader(byt))
	return
}

// GetPortalStatus gets the portal.status value from the UTM
func GetPortalStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/portal.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalStatus PUTs the portal.status value to the UTM
func UpdatePortalStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.status", bytes.NewReader(byt))
	return
}

// GetPortalWelcomeMsg gets the portal.welcome_msg value from the UTM
func GetPortalWelcomeMsg(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/portal.welcome_msg")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePortalWelcomeMsg PUTs the portal.welcome_msg value to the UTM
func UpdatePortalWelcomeMsg(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/portal.welcome_msg", bytes.NewReader(byt))
	return
}

// GetPsdAction gets the psd.action value from the UTM
func GetPsdAction(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/psd.action")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdAction PUTs the psd.action value to the UTM
func UpdatePsdAction(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.action", bytes.NewReader(byt))
	return
}

// GetPsdDelayThreshold gets the psd.delay_threshold value from the UTM
func GetPsdDelayThreshold(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/psd.delay_threshold")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdDelayThreshold PUTs the psd.delay_threshold value to the UTM
func UpdatePsdDelayThreshold(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.delay_threshold", bytes.NewReader(byt))
	return
}

// GetPsdHiPortsWeight gets the psd.hi_ports_weight value from the UTM
func GetPsdHiPortsWeight(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/psd.hi_ports_weight")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdHiPortsWeight PUTs the psd.hi_ports_weight value to the UTM
func UpdatePsdHiPortsWeight(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.hi_ports_weight", bytes.NewReader(byt))
	return
}

// GetPsdLoPortsWeight gets the psd.lo_ports_weight value from the UTM
func GetPsdLoPortsWeight(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/psd.lo_ports_weight")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdLoPortsWeight PUTs the psd.lo_ports_weight value to the UTM
func UpdatePsdLoPortsWeight(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.lo_ports_weight", bytes.NewReader(byt))
	return
}

// GetPsdLogLimiterBurst gets the psd.log_limiter.burst value from the UTM
func GetPsdLogLimiterBurst(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/psd.log_limiter.burst")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdLogLimiterBurst PUTs the psd.log_limiter.burst value to the UTM
func UpdatePsdLogLimiterBurst(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.log_limiter.burst", bytes.NewReader(byt))
	return
}

// GetPsdLogLimiterRate gets the psd.log_limiter.rate value from the UTM
func GetPsdLogLimiterRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/psd.log_limiter.rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdLogLimiterRate PUTs the psd.log_limiter.rate value to the UTM
func UpdatePsdLogLimiterRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.log_limiter.rate", bytes.NewReader(byt))
	return
}

// GetPsdLogLimiterStatus gets the psd.log_limiter.status value from the UTM
func GetPsdLogLimiterStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/psd.log_limiter.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdLogLimiterStatus PUTs the psd.log_limiter.status value to the UTM
func UpdatePsdLogLimiterStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.log_limiter.status", bytes.NewReader(byt))
	return
}

// GetPsdStatus gets the psd.status value from the UTM
func GetPsdStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/psd.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdStatus PUTs the psd.status value to the UTM
func UpdatePsdStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.status", bytes.NewReader(byt))
	return
}

// GetPsdWeightThreshold gets the psd.weight_threshold value from the UTM
func GetPsdWeightThreshold(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/psd.weight_threshold")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdatePsdWeightThreshold PUTs the psd.weight_threshold value to the UTM
func UpdatePsdWeightThreshold(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/psd.weight_threshold", bytes.NewReader(byt))
	return
}

// GetQosAdvancedEcn gets the qos.advanced.ecn value from the UTM
func GetQosAdvancedEcn(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/qos.advanced.ecn")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateQosAdvancedEcn PUTs the qos.advanced.ecn value to the UTM
func UpdateQosAdvancedEcn(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/qos.advanced.ecn", bytes.NewReader(byt))
	return
}

// GetQosAdvancedKeepClassAfterEncap gets the qos.advanced.keep_class_after_encap value from the UTM
func GetQosAdvancedKeepClassAfterEncap(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/qos.advanced.keep_class_after_encap")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateQosAdvancedKeepClassAfterEncap PUTs the qos.advanced.keep_class_after_encap value to the UTM
func UpdateQosAdvancedKeepClassAfterEncap(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/qos.advanced.keep_class_after_encap", bytes.NewReader(byt))
	return
}

// GetQosInterfaces gets the qos.interfaces value from the UTM
func GetQosInterfaces(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/qos.interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateQosInterfaces PUTs the qos.interfaces value to the UTM
func UpdateQosInterfaces(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/qos.interfaces", bytes.NewReader(byt))
	return
}

// GetQuarantineKeepDbLogDays gets the quarantine.keep_db_log_days value from the UTM
func GetQuarantineKeepDbLogDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/quarantine.keep_db_log_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateQuarantineKeepDbLogDays PUTs the quarantine.keep_db_log_days value to the UTM
func UpdateQuarantineKeepDbLogDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/quarantine.keep_db_log_days", bytes.NewReader(byt))
	return
}

// GetQuarantineKeepQuarantineDays gets the quarantine.keep_quarantine_days value from the UTM
func GetQuarantineKeepQuarantineDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/quarantine.keep_quarantine_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateQuarantineKeepQuarantineDays PUTs the quarantine.keep_quarantine_days value to the UTM
func UpdateQuarantineKeepQuarantineDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/quarantine.keep_quarantine_days", bytes.NewReader(byt))
	return
}

// GetRedActivateProvFw gets the red.activate_prov_fw value from the UTM
func GetRedActivateProvFw(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/red.activate_prov_fw")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedActivateProvFw PUTs the red.activate_prov_fw value to the UTM
func UpdateRedActivateProvFw(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.activate_prov_fw", bytes.NewReader(byt))
	return
}

// GetRedAuthorization gets the red.authorization value from the UTM
func GetRedAuthorization(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/red.authorization")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedAuthorization PUTs the red.authorization value to the UTM
func UpdateRedAuthorization(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.authorization", bytes.NewReader(byt))
	return
}

// GetRedCaSettingsCity gets the red.ca_settings.city value from the UTM
func GetRedCaSettingsCity(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.ca_settings.city")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedCaSettingsCity PUTs the red.ca_settings.city value to the UTM
func UpdateRedCaSettingsCity(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.ca_settings.city", bytes.NewReader(byt))
	return
}

// GetRedCaSettingsCountry gets the red.ca_settings.country value from the UTM
func GetRedCaSettingsCountry(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.ca_settings.country")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedCaSettingsCountry PUTs the red.ca_settings.country value to the UTM
func UpdateRedCaSettingsCountry(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.ca_settings.country", bytes.NewReader(byt))
	return
}

// GetRedCaSettingsEmail gets the red.ca_settings.email value from the UTM
func GetRedCaSettingsEmail(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.ca_settings.email")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedCaSettingsEmail PUTs the red.ca_settings.email value to the UTM
func UpdateRedCaSettingsEmail(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.ca_settings.email", bytes.NewReader(byt))
	return
}

// GetRedCaSettingsOrganization gets the red.ca_settings.organization value from the UTM
func GetRedCaSettingsOrganization(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.ca_settings.organization")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedCaSettingsOrganization PUTs the red.ca_settings.organization value to the UTM
func UpdateRedCaSettingsOrganization(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.ca_settings.organization", bytes.NewReader(byt))
	return
}

// GetRedClients gets the red.clients value from the UTM
func GetRedClients(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/red.clients")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedClients PUTs the red.clients value to the UTM
func UpdateRedClients(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.clients", bytes.NewReader(byt))
	return
}

// GetRedDeauthTimeout gets the red.deauth_timeout value from the UTM
func GetRedDeauthTimeout(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.deauth_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedDeauthTimeout PUTs the red.deauth_timeout value to the UTM
func UpdateRedDeauthTimeout(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.deauth_timeout", bytes.NewReader(byt))
	return
}

// GetRedOverlayFwEnabled gets the red.overlay_fw_enabled value from the UTM
func GetRedOverlayFwEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/red.overlay_fw_enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedOverlayFwEnabled PUTs the red.overlay_fw_enabled value to the UTM
func UpdateRedOverlayFwEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.overlay_fw_enabled", bytes.NewReader(byt))
	return
}

// GetRedRegistryCert gets the red.registry_cert value from the UTM
func GetRedRegistryCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.registry_cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedRegistryCert PUTs the red.registry_cert value to the UTM
func UpdateRedRegistryCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.registry_cert", bytes.NewReader(byt))
	return
}

// GetRedRegistryId gets the red.registry_id value from the UTM
func GetRedRegistryId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.registry_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedRegistryId PUTs the red.registry_id value to the UTM
func UpdateRedRegistryId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.registry_id", bytes.NewReader(byt))
	return
}

// GetRedRegistryKey gets the red.registry_key value from the UTM
func GetRedRegistryKey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.registry_key")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedRegistryKey PUTs the red.registry_key value to the UTM
func UpdateRedRegistryKey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.registry_key", bytes.NewReader(byt))
	return
}

// GetRedServerCert gets the red.server_cert value from the UTM
func GetRedServerCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/red.server_cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedServerCert PUTs the red.server_cert value to the UTM
func UpdateRedServerCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.server_cert", bytes.NewReader(byt))
	return
}

// GetRedServers gets the red.servers value from the UTM
func GetRedServers(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/red.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedServers PUTs the red.servers value to the UTM
func UpdateRedServers(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.servers", bytes.NewReader(byt))
	return
}

// GetRedStatus gets the red.status value from the UTM
func GetRedStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/red.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedStatus PUTs the red.status value to the UTM
func UpdateRedStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.status", bytes.NewReader(byt))
	return
}

// GetRedTls12Only gets the red.tls_1_2_only value from the UTM
func GetRedTls12Only(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/red.tls_1_2_only")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRedTls12Only PUTs the red.tls_1_2_only value to the UTM
func UpdateRedTls12Only(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/red.tls_1_2_only", bytes.NewReader(byt))
	return
}

// GetRemoteAccessAdvancedMsdns1 gets the remote_access.advanced.msdns1 value from the UTM
func GetRemoteAccessAdvancedMsdns1(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.advanced.msdns1")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessAdvancedMsdns1 PUTs the remote_access.advanced.msdns1 value to the UTM
func UpdateRemoteAccessAdvancedMsdns1(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.advanced.msdns1", bytes.NewReader(byt))
	return
}

// GetRemoteAccessAdvancedMsdns2 gets the remote_access.advanced.msdns2 value from the UTM
func GetRemoteAccessAdvancedMsdns2(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.advanced.msdns2")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessAdvancedMsdns2 PUTs the remote_access.advanced.msdns2 value to the UTM
func UpdateRemoteAccessAdvancedMsdns2(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.advanced.msdns2", bytes.NewReader(byt))
	return
}

// GetRemoteAccessAdvancedMsdomain gets the remote_access.advanced.msdomain value from the UTM
func GetRemoteAccessAdvancedMsdomain(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.advanced.msdomain")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessAdvancedMsdomain PUTs the remote_access.advanced.msdomain value to the UTM
func UpdateRemoteAccessAdvancedMsdomain(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.advanced.msdomain", bytes.NewReader(byt))
	return
}

// GetRemoteAccessAdvancedMswins1 gets the remote_access.advanced.mswins1 value from the UTM
func GetRemoteAccessAdvancedMswins1(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.advanced.mswins1")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessAdvancedMswins1 PUTs the remote_access.advanced.mswins1 value to the UTM
func UpdateRemoteAccessAdvancedMswins1(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.advanced.mswins1", bytes.NewReader(byt))
	return
}

// GetRemoteAccessAdvancedMswins2 gets the remote_access.advanced.mswins2 value from the UTM
func GetRemoteAccessAdvancedMswins2(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.advanced.mswins2")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessAdvancedMswins2 PUTs the remote_access.advanced.mswins2 value to the UTM
func UpdateRemoteAccessAdvancedMswins2(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.advanced.mswins2", bytes.NewReader(byt))
	return
}

// GetRemoteAccessCisco gets the remote_access.cisco value from the UTM
func GetRemoteAccessCisco(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.cisco")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessCisco PUTs the remote_access.cisco value to the UTM
func UpdateRemoteAccessCisco(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.cisco", bytes.NewReader(byt))
	return
}

// GetRemoteAccessClientlessVpnDebug gets the remote_access.clientless_vpn.debug value from the UTM
func GetRemoteAccessClientlessVpnDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/remote_access.clientless_vpn.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessClientlessVpnDebug PUTs the remote_access.clientless_vpn.debug value to the UTM
func UpdateRemoteAccessClientlessVpnDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.clientless_vpn.debug", bytes.NewReader(byt))
	return
}

// GetRemoteAccessClientlessVpnStatus gets the remote_access.clientless_vpn.status value from the UTM
func GetRemoteAccessClientlessVpnStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/remote_access.clientless_vpn.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessClientlessVpnStatus PUTs the remote_access.clientless_vpn.status value to the UTM
func UpdateRemoteAccessClientlessVpnStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.clientless_vpn.status", bytes.NewReader(byt))
	return
}

// GetRemoteAccessL2Tp gets the remote_access.l2tp value from the UTM
func GetRemoteAccessL2Tp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.l2tp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessL2Tp PUTs the remote_access.l2tp value to the UTM
func UpdateRemoteAccessL2Tp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.l2tp", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpAaa gets the remote_access.pptp.aaa value from the UTM
func GetRemoteAccessPptpAaa(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.aaa")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpAaa PUTs the remote_access.pptp.aaa value to the UTM
func UpdateRemoteAccessPptpAaa(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.aaa", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpAuthentication gets the remote_access.pptp.authentication value from the UTM
func GetRemoteAccessPptpAuthentication(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.authentication")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpAuthentication PUTs the remote_access.pptp.authentication value to the UTM
func UpdateRemoteAccessPptpAuthentication(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.authentication", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpDebug gets the remote_access.pptp.debug value from the UTM
func GetRemoteAccessPptpDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpDebug PUTs the remote_access.pptp.debug value to the UTM
func UpdateRemoteAccessPptpDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.debug", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpEncryption gets the remote_access.pptp.encryption value from the UTM
func GetRemoteAccessPptpEncryption(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.encryption")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpEncryption PUTs the remote_access.pptp.encryption value to the UTM
func UpdateRemoteAccessPptpEncryption(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.encryption", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIpAssignmentDhcp gets the remote_access.pptp.ip_assignment_dhcp value from the UTM
func GetRemoteAccessPptpIpAssignmentDhcp(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.ip_assignment_dhcp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIpAssignmentDhcp PUTs the remote_access.pptp.ip_assignment_dhcp value to the UTM
func UpdateRemoteAccessPptpIpAssignmentDhcp(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.ip_assignment_dhcp", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIpAssignmentDhcpInterface gets the remote_access.pptp.ip_assignment_dhcp_interface value from the UTM
func GetRemoteAccessPptpIpAssignmentDhcpInterface(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.ip_assignment_dhcp_interface")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIpAssignmentDhcpInterface PUTs the remote_access.pptp.ip_assignment_dhcp_interface value to the UTM
func UpdateRemoteAccessPptpIpAssignmentDhcpInterface(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.ip_assignment_dhcp_interface", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIpAssignmentMode gets the remote_access.pptp.ip_assignment_mode value from the UTM
func GetRemoteAccessPptpIpAssignmentMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.ip_assignment_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIpAssignmentMode PUTs the remote_access.pptp.ip_assignment_mode value to the UTM
func UpdateRemoteAccessPptpIpAssignmentMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.ip_assignment_mode", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIpAssignmentPool gets the remote_access.pptp.ip_assignment_pool value from the UTM
func GetRemoteAccessPptpIpAssignmentPool(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.ip_assignment_pool")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIpAssignmentPool PUTs the remote_access.pptp.ip_assignment_pool value to the UTM
func UpdateRemoteAccessPptpIpAssignmentPool(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.ip_assignment_pool", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIphoneConnectionName gets the remote_access.pptp.iphone_connection_name value from the UTM
func GetRemoteAccessPptpIphoneConnectionName(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.iphone_connection_name")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIphoneConnectionName PUTs the remote_access.pptp.iphone_connection_name value to the UTM
func UpdateRemoteAccessPptpIphoneConnectionName(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.iphone_connection_name", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIphoneHostname gets the remote_access.pptp.iphone_hostname value from the UTM
func GetRemoteAccessPptpIphoneHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.iphone_hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIphoneHostname PUTs the remote_access.pptp.iphone_hostname value to the UTM
func UpdateRemoteAccessPptpIphoneHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.iphone_hostname", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpIphoneStatus gets the remote_access.pptp.iphone_status value from the UTM
func GetRemoteAccessPptpIphoneStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.iphone_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpIphoneStatus PUTs the remote_access.pptp.iphone_status value to the UTM
func UpdateRemoteAccessPptpIphoneStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.iphone_status", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpMtu gets the remote_access.pptp.mtu value from the UTM
func GetRemoteAccessPptpMtu(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.mtu")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpMtu PUTs the remote_access.pptp.mtu value to the UTM
func UpdateRemoteAccessPptpMtu(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.mtu", bytes.NewReader(byt))
	return
}

// GetRemoteAccessPptpStatus gets the remote_access.pptp.status value from the UTM
func GetRemoteAccessPptpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/remote_access.pptp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteAccessPptpStatus PUTs the remote_access.pptp.status value to the UTM
func UpdateRemoteAccessPptpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_access.pptp.status", bytes.NewReader(byt))
	return
}

// GetRemoteSyslogBuffer gets the remote_syslog.buffer value from the UTM
func GetRemoteSyslogBuffer(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/remote_syslog.buffer")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteSyslogBuffer PUTs the remote_syslog.buffer value to the UTM
func UpdateRemoteSyslogBuffer(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_syslog.buffer", bytes.NewReader(byt))
	return
}

// GetRemoteSyslogLogs gets the remote_syslog.logs value from the UTM
func GetRemoteSyslogLogs(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/remote_syslog.logs")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteSyslogLogs PUTs the remote_syslog.logs value to the UTM
func UpdateRemoteSyslogLogs(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_syslog.logs", bytes.NewReader(byt))
	return
}

// GetRemoteSyslogStatus gets the remote_syslog.status value from the UTM
func GetRemoteSyslogStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/remote_syslog.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteSyslogStatus PUTs the remote_syslog.status value to the UTM
func UpdateRemoteSyslogStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_syslog.status", bytes.NewReader(byt))
	return
}

// GetRemoteSyslogTarget gets the remote_syslog.target value from the UTM
func GetRemoteSyslogTarget(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/remote_syslog.target")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRemoteSyslogTarget PUTs the remote_syslog.target value to the UTM
func UpdateRemoteSyslogTarget(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/remote_syslog.target", bytes.NewReader(byt))
	return
}

// GetReportingAccountingKeepdays gets the reporting.accounting_keepdays value from the UTM
func GetReportingAccountingKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.accounting_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAccountingKeepdays PUTs the reporting.accounting_keepdays value to the UTM
func UpdateReportingAccountingKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.accounting_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingAccountingStatus gets the reporting.accounting_status value from the UTM
func GetReportingAccountingStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.accounting_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAccountingStatus PUTs the reporting.accounting_status value to the UTM
func UpdateReportingAccountingStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.accounting_status", bytes.NewReader(byt))
	return
}

// GetReportingAnonymizing gets the reporting.anonymizing value from the UTM
func GetReportingAnonymizing(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.anonymizing")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAnonymizing PUTs the reporting.anonymizing value to the UTM
func UpdateReportingAnonymizing(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.anonymizing", bytes.NewReader(byt))
	return
}

// GetReportingAppctrlKeepdays gets the reporting.appctrl_keepdays value from the UTM
func GetReportingAppctrlKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.appctrl_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAppctrlKeepdays PUTs the reporting.appctrl_keepdays value to the UTM
func UpdateReportingAppctrlKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.appctrl_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingAppctrlStatus gets the reporting.appctrl_status value from the UTM
func GetReportingAppctrlStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.appctrl_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAppctrlStatus PUTs the reporting.appctrl_status value to the UTM
func UpdateReportingAppctrlStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.appctrl_status", bytes.NewReader(byt))
	return
}

// GetReportingAtpKeepdays gets the reporting.atp_keepdays value from the UTM
func GetReportingAtpKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.atp_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAtpKeepdays PUTs the reporting.atp_keepdays value to the UTM
func UpdateReportingAtpKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.atp_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingAtpReset gets the reporting.atp_reset value from the UTM
func GetReportingAtpReset(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.atp_reset")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAtpReset PUTs the reporting.atp_reset value to the UTM
func UpdateReportingAtpReset(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.atp_reset", bytes.NewReader(byt))
	return
}

// GetReportingAtpStatus gets the reporting.atp_status value from the UTM
func GetReportingAtpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.atp_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAtpStatus PUTs the reporting.atp_status value to the UTM
func UpdateReportingAtpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.atp_status", bytes.NewReader(byt))
	return
}

// GetReportingAuthenticationKeepdays gets the reporting.authentication_keepdays value from the UTM
func GetReportingAuthenticationKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.authentication_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAuthenticationKeepdays PUTs the reporting.authentication_keepdays value to the UTM
func UpdateReportingAuthenticationKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.authentication_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingAuthenticationStatus gets the reporting.authentication_status value from the UTM
func GetReportingAuthenticationStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.authentication_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingAuthenticationStatus PUTs the reporting.authentication_status value to the UTM
func UpdateReportingAuthenticationStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.authentication_status", bytes.NewReader(byt))
	return
}

// GetReportingCsvSeparator gets the reporting.csv_separator value from the UTM
func GetReportingCsvSeparator(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reporting.csv_separator")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingCsvSeparator PUTs the reporting.csv_separator value to the UTM
func UpdateReportingCsvSeparator(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.csv_separator", bytes.NewReader(byt))
	return
}

// GetReportingEmailsecurityImport gets the reporting.emailsecurity_import value from the UTM
func GetReportingEmailsecurityImport(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.emailsecurity_import")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingEmailsecurityImport PUTs the reporting.emailsecurity_import value to the UTM
func UpdateReportingEmailsecurityImport(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.emailsecurity_import", bytes.NewReader(byt))
	return
}

// GetReportingEmailsecurityKeepdays gets the reporting.emailsecurity_keepdays value from the UTM
func GetReportingEmailsecurityKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.emailsecurity_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingEmailsecurityKeepdays PUTs the reporting.emailsecurity_keepdays value to the UTM
func UpdateReportingEmailsecurityKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.emailsecurity_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingEmailsecurityStatus gets the reporting.emailsecurity_status value from the UTM
func GetReportingEmailsecurityStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.emailsecurity_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingEmailsecurityStatus PUTs the reporting.emailsecurity_status value to the UTM
func UpdateReportingEmailsecurityStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.emailsecurity_status", bytes.NewReader(byt))
	return
}

// GetReportingEnableVpnAccounting gets the reporting.enable_vpn_accounting value from the UTM
func GetReportingEnableVpnAccounting(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.enable_vpn_accounting")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingEnableVpnAccounting PUTs the reporting.enable_vpn_accounting value to the UTM
func UpdateReportingEnableVpnAccounting(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.enable_vpn_accounting", bytes.NewReader(byt))
	return
}

// GetReportingHideAccountingips gets the reporting.hide_accountingips value from the UTM
func GetReportingHideAccountingips(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.hide_accountingips")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingHideAccountingips PUTs the reporting.hide_accountingips value to the UTM
func UpdateReportingHideAccountingips(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.hide_accountingips", bytes.NewReader(byt))
	return
}

// GetReportingHideMailaddresses gets the reporting.hide_mailaddresses value from the UTM
func GetReportingHideMailaddresses(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.hide_mailaddresses")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingHideMailaddresses PUTs the reporting.hide_mailaddresses value to the UTM
func UpdateReportingHideMailaddresses(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.hide_mailaddresses", bytes.NewReader(byt))
	return
}

// GetReportingHideMaildomains gets the reporting.hide_maildomains value from the UTM
func GetReportingHideMaildomains(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.hide_maildomains")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingHideMaildomains PUTs the reporting.hide_maildomains value to the UTM
func UpdateReportingHideMaildomains(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.hide_maildomains", bytes.NewReader(byt))
	return
}

// GetReportingHideNetsecips gets the reporting.hide_netsecips value from the UTM
func GetReportingHideNetsecips(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.hide_netsecips")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingHideNetsecips PUTs the reporting.hide_netsecips value to the UTM
func UpdateReportingHideNetsecips(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.hide_netsecips", bytes.NewReader(byt))
	return
}

// GetReportingHideWebdomains gets the reporting.hide_webdomains value from the UTM
func GetReportingHideWebdomains(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.hide_webdomains")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingHideWebdomains PUTs the reporting.hide_webdomains value to the UTM
func UpdateReportingHideWebdomains(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.hide_webdomains", bytes.NewReader(byt))
	return
}

// GetReportingIpsImport gets the reporting.ips_import value from the UTM
func GetReportingIpsImport(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.ips_import")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingIpsImport PUTs the reporting.ips_import value to the UTM
func UpdateReportingIpsImport(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.ips_import", bytes.NewReader(byt))
	return
}

// GetReportingIpsKeepdays gets the reporting.ips_keepdays value from the UTM
func GetReportingIpsKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.ips_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingIpsKeepdays PUTs the reporting.ips_keepdays value to the UTM
func UpdateReportingIpsKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.ips_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingIpsStatus gets the reporting.ips_status value from the UTM
func GetReportingIpsStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.ips_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingIpsStatus PUTs the reporting.ips_status value to the UTM
func UpdateReportingIpsStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.ips_status", bytes.NewReader(byt))
	return
}

// GetReportingPacketfilterImport gets the reporting.packetfilter_import value from the UTM
func GetReportingPacketfilterImport(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.packetfilter_import")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingPacketfilterImport PUTs the reporting.packetfilter_import value to the UTM
func UpdateReportingPacketfilterImport(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.packetfilter_import", bytes.NewReader(byt))
	return
}

// GetReportingPacketfilterKeepdays gets the reporting.packetfilter_keepdays value from the UTM
func GetReportingPacketfilterKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.packetfilter_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingPacketfilterKeepdays PUTs the reporting.packetfilter_keepdays value to the UTM
func UpdateReportingPacketfilterKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.packetfilter_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingPacketfilterStatus gets the reporting.packetfilter_status value from the UTM
func GetReportingPacketfilterStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.packetfilter_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingPacketfilterStatus PUTs the reporting.packetfilter_status value to the UTM
func UpdateReportingPacketfilterStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.packetfilter_status", bytes.NewReader(byt))
	return
}

// GetReportingPassword1 gets the reporting.password1 value from the UTM
func GetReportingPassword1(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reporting.password1")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingPassword1 PUTs the reporting.password1 value to the UTM
func UpdateReportingPassword1(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.password1", bytes.NewReader(byt))
	return
}

// GetReportingPassword2 gets the reporting.password2 value from the UTM
func GetReportingPassword2(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reporting.password2")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingPassword2 PUTs the reporting.password2 value to the UTM
func UpdateReportingPassword2(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.password2", bytes.NewReader(byt))
	return
}

// GetReportingSandboxKeepdays gets the reporting.sandbox_keepdays value from the UTM
func GetReportingSandboxKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.sandbox_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingSandboxKeepdays PUTs the reporting.sandbox_keepdays value to the UTM
func UpdateReportingSandboxKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.sandbox_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingUserlogFromLogs gets the reporting.userlog_from_logs value from the UTM
func GetReportingUserlogFromLogs(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.userlog_from_logs")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingUserlogFromLogs PUTs the reporting.userlog_from_logs value to the UTM
func UpdateReportingUserlogFromLogs(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.userlog_from_logs", bytes.NewReader(byt))
	return
}

// GetReportingVpnKeepdays gets the reporting.vpn_keepdays value from the UTM
func GetReportingVpnKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.vpn_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingVpnKeepdays PUTs the reporting.vpn_keepdays value to the UTM
func UpdateReportingVpnKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.vpn_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingVpnStatus gets the reporting.vpn_status value from the UTM
func GetReportingVpnStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.vpn_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingVpnStatus PUTs the reporting.vpn_status value to the UTM
func UpdateReportingVpnStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.vpn_status", bytes.NewReader(byt))
	return
}

// GetReportingWafKeepdays gets the reporting.waf_keepdays value from the UTM
func GetReportingWafKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.waf_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingWafKeepdays PUTs the reporting.waf_keepdays value to the UTM
func UpdateReportingWafKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.waf_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingWafStatus gets the reporting.waf_status value from the UTM
func GetReportingWafStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.waf_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingWafStatus PUTs the reporting.waf_status value to the UTM
func UpdateReportingWafStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.waf_status", bytes.NewReader(byt))
	return
}

// GetReportingWebsecurityDetail gets the reporting.websecurity_detail value from the UTM
func GetReportingWebsecurityDetail(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.websecurity_detail")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingWebsecurityDetail PUTs the reporting.websecurity_detail value to the UTM
func UpdateReportingWebsecurityDetail(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.websecurity_detail", bytes.NewReader(byt))
	return
}

// GetReportingWebsecurityImport gets the reporting.websecurity_import value from the UTM
func GetReportingWebsecurityImport(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reporting.websecurity_import")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingWebsecurityImport PUTs the reporting.websecurity_import value to the UTM
func UpdateReportingWebsecurityImport(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.websecurity_import", bytes.NewReader(byt))
	return
}

// GetReportingWebsecurityKeepdays gets the reporting.websecurity_keepdays value from the UTM
func GetReportingWebsecurityKeepdays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reporting.websecurity_keepdays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingWebsecurityKeepdays PUTs the reporting.websecurity_keepdays value to the UTM
func UpdateReportingWebsecurityKeepdays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.websecurity_keepdays", bytes.NewReader(byt))
	return
}

// GetReportingWebsecurityStatus gets the reporting.websecurity_status value from the UTM
func GetReportingWebsecurityStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reporting.websecurity_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReportingWebsecurityStatus PUTs the reporting.websecurity_status value to the UTM
func UpdateReportingWebsecurityStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reporting.websecurity_status", bytes.NewReader(byt))
	return
}

// GetReverseProxyAuaRefreshEnabled gets the reverse_proxy.aua_refresh_enabled value from the UTM
func GetReverseProxyAuaRefreshEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.aua_refresh_enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyAuaRefreshEnabled PUTs the reverse_proxy.aua_refresh_enabled value to the UTM
func UpdateReverseProxyAuaRefreshEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.aua_refresh_enabled", bytes.NewReader(byt))
	return
}

// GetReverseProxyAuaRefreshInterval gets the reverse_proxy.aua_refresh_interval value from the UTM
func GetReverseProxyAuaRefreshInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.aua_refresh_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyAuaRefreshInterval PUTs the reverse_proxy.aua_refresh_interval value to the UTM
func UpdateReverseProxyAuaRefreshInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.aua_refresh_interval", bytes.NewReader(byt))
	return
}

// GetReverseProxyBlacklistDnsrblZones gets the reverse_proxy.blacklist.dnsrbl_zones value from the UTM
func GetReverseProxyBlacklistDnsrblZones(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.blacklist.dnsrbl_zones")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyBlacklistDnsrblZones PUTs the reverse_proxy.blacklist.dnsrbl_zones value to the UTM
func UpdateReverseProxyBlacklistDnsrblZones(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.blacklist.dnsrbl_zones", bytes.NewReader(byt))
	return
}

// GetReverseProxyBlacklistGeoipCodes gets the reverse_proxy.blacklist.geoip_codes value from the UTM
func GetReverseProxyBlacklistGeoipCodes(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.blacklist.geoip_codes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyBlacklistGeoipCodes PUTs the reverse_proxy.blacklist.geoip_codes value to the UTM
func UpdateReverseProxyBlacklistGeoipCodes(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.blacklist.geoip_codes", bytes.NewReader(byt))
	return
}

// GetReverseProxyCookiesignkey gets the reverse_proxy.cookiesignkey value from the UTM
func GetReverseProxyCookiesignkey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.cookiesignkey")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyCookiesignkey PUTs the reverse_proxy.cookiesignkey value to the UTM
func UpdateReverseProxyCookiesignkey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.cookiesignkey", bytes.NewReader(byt))
	return
}

// GetReverseProxyCssdHostname gets the reverse_proxy.cssd_hostname value from the UTM
func GetReverseProxyCssdHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.cssd_hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyCssdHostname PUTs the reverse_proxy.cssd_hostname value to the UTM
func UpdateReverseProxyCssdHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.cssd_hostname", bytes.NewReader(byt))
	return
}

// GetReverseProxyCssdPort gets the reverse_proxy.cssd_port value from the UTM
func GetReverseProxyCssdPort(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.cssd_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyCssdPort PUTs the reverse_proxy.cssd_port value to the UTM
func UpdateReverseProxyCssdPort(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.cssd_port", bytes.NewReader(byt))
	return
}

// GetReverseProxyCustomThreatFiltersEnabled gets the reverse_proxy.custom_threat_filters_enabled value from the UTM
func GetReverseProxyCustomThreatFiltersEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.custom_threat_filters_enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyCustomThreatFiltersEnabled PUTs the reverse_proxy.custom_threat_filters_enabled value to the UTM
func UpdateReverseProxyCustomThreatFiltersEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.custom_threat_filters_enabled", bytes.NewReader(byt))
	return
}

// GetReverseProxyFormhardeningSecret gets the reverse_proxy.formhardening_secret value from the UTM
func GetReverseProxyFormhardeningSecret(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.formhardening_secret")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyFormhardeningSecret PUTs the reverse_proxy.formhardening_secret value to the UTM
func UpdateReverseProxyFormhardeningSecret(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.formhardening_secret", bytes.NewReader(byt))
	return
}

// GetReverseProxyKeepalive gets the reverse_proxy.keepalive value from the UTM
func GetReverseProxyKeepalive(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.keepalive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyKeepalive PUTs the reverse_proxy.keepalive value to the UTM
func UpdateReverseProxyKeepalive(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.keepalive", bytes.NewReader(byt))
	return
}

// GetReverseProxyManualmode gets the reverse_proxy.manualmode value from the UTM
func GetReverseProxyManualmode(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.manualmode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyManualmode PUTs the reverse_proxy.manualmode value to the UTM
func UpdateReverseProxyManualmode(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.manualmode", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxConnectionsPerChild gets the reverse_proxy.max_connections_per_child value from the UTM
func GetReverseProxyMaxConnectionsPerChild(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_connections_per_child")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxConnectionsPerChild PUTs the reverse_proxy.max_connections_per_child value to the UTM
func UpdateReverseProxyMaxConnectionsPerChild(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_connections_per_child", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxPreforkProcesses gets the reverse_proxy.max_prefork_processes value from the UTM
func GetReverseProxyMaxPreforkProcesses(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_prefork_processes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxPreforkProcesses PUTs the reverse_proxy.max_prefork_processes value to the UTM
func UpdateReverseProxyMaxPreforkProcesses(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_prefork_processes", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxProcesses gets the reverse_proxy.max_processes value from the UTM
func GetReverseProxyMaxProcesses(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_processes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxProcesses PUTs the reverse_proxy.max_processes value to the UTM
func UpdateReverseProxyMaxProcesses(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_processes", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxSessionFiles gets the reverse_proxy.max_session_files value from the UTM
func GetReverseProxyMaxSessionFiles(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_session_files")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxSessionFiles PUTs the reverse_proxy.max_session_files value to the UTM
func UpdateReverseProxyMaxSessionFiles(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_session_files", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxSpareProcesses gets the reverse_proxy.max_spare_processes value from the UTM
func GetReverseProxyMaxSpareProcesses(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_spare_processes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxSpareProcesses PUTs the reverse_proxy.max_spare_processes value to the UTM
func UpdateReverseProxyMaxSpareProcesses(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_spare_processes", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxSpareThreads gets the reverse_proxy.max_spare_threads value from the UTM
func GetReverseProxyMaxSpareThreads(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_spare_threads")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxSpareThreads PUTs the reverse_proxy.max_spare_threads value to the UTM
func UpdateReverseProxyMaxSpareThreads(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_spare_threads", bytes.NewReader(byt))
	return
}

// GetReverseProxyMaxThreadsPerProcess gets the reverse_proxy.max_threads_per_process value from the UTM
func GetReverseProxyMaxThreadsPerProcess(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.max_threads_per_process")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMaxThreadsPerProcess PUTs the reverse_proxy.max_threads_per_process value to the UTM
func UpdateReverseProxyMaxThreadsPerProcess(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.max_threads_per_process", bytes.NewReader(byt))
	return
}

// GetReverseProxyMinSpareProcesses gets the reverse_proxy.min_spare_processes value from the UTM
func GetReverseProxyMinSpareProcesses(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.min_spare_processes")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMinSpareProcesses PUTs the reverse_proxy.min_spare_processes value to the UTM
func UpdateReverseProxyMinSpareProcesses(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.min_spare_processes", bytes.NewReader(byt))
	return
}

// GetReverseProxyMinSpareThreads gets the reverse_proxy.min_spare_threads value from the UTM
func GetReverseProxyMinSpareThreads(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.min_spare_threads")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMinSpareThreads PUTs the reverse_proxy.min_spare_threads value to the UTM
func UpdateReverseProxyMinSpareThreads(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.min_spare_threads", bytes.NewReader(byt))
	return
}

// GetReverseProxyMinTls gets the reverse_proxy.min_tls value from the UTM
func GetReverseProxyMinTls(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.min_tls")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMinTls PUTs the reverse_proxy.min_tls value to the UTM
func UpdateReverseProxyMinTls(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.min_tls", bytes.NewReader(byt))
	return
}

// GetReverseProxyModsecurityBeta gets the reverse_proxy.modsecurity_beta value from the UTM
func GetReverseProxyModsecurityBeta(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.modsecurity_beta")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyModsecurityBeta PUTs the reverse_proxy.modsecurity_beta value to the UTM
func UpdateReverseProxyModsecurityBeta(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.modsecurity_beta", bytes.NewReader(byt))
	return
}

// GetReverseProxyMpmMode gets the reverse_proxy.mpm_mode value from the UTM
func GetReverseProxyMpmMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.mpm_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyMpmMode PUTs the reverse_proxy.mpm_mode value to the UTM
func UpdateReverseProxyMpmMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.mpm_mode", bytes.NewReader(byt))
	return
}

// GetReverseProxyPatternversion gets the reverse_proxy.patternversion value from the UTM
func GetReverseProxyPatternversion(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.patternversion")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyPatternversion PUTs the reverse_proxy.patternversion value to the UTM
func UpdateReverseProxyPatternversion(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.patternversion", bytes.NewReader(byt))
	return
}

// GetReverseProxyPort gets the reverse_proxy.port value from the UTM
func GetReverseProxyPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyPort PUTs the reverse_proxy.port value to the UTM
func UpdateReverseProxyPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.port", bytes.NewReader(byt))
	return
}

// GetReverseProxyProxyprotocol gets the reverse_proxy.proxyprotocol value from the UTM
func GetReverseProxyProxyprotocol(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.proxyprotocol")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyProxyprotocol PUTs the reverse_proxy.proxyprotocol value to the UTM
func UpdateReverseProxyProxyprotocol(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.proxyprotocol", bytes.NewReader(byt))
	return
}

// GetReverseProxyRequestLineLimit gets the reverse_proxy.request_line_limit value from the UTM
func GetReverseProxyRequestLineLimit(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.request_line_limit")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyRequestLineLimit PUTs the reverse_proxy.request_line_limit value to the UTM
func UpdateReverseProxyRequestLineLimit(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.request_line_limit", bytes.NewReader(byt))
	return
}

// GetReverseProxySlowhttpExceptions gets the reverse_proxy.slowhttp_exceptions value from the UTM
func GetReverseProxySlowhttpExceptions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.slowhttp_exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxySlowhttpExceptions PUTs the reverse_proxy.slowhttp_exceptions value to the UTM
func UpdateReverseProxySlowhttpExceptions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.slowhttp_exceptions", bytes.NewReader(byt))
	return
}

// GetReverseProxySlowhttpRequestHeaderTimeoutBase gets the reverse_proxy.slowhttp_request_header_timeout_base value from the UTM
func GetReverseProxySlowhttpRequestHeaderTimeoutBase(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_base")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxySlowhttpRequestHeaderTimeoutBase PUTs the reverse_proxy.slowhttp_request_header_timeout_base value to the UTM
func UpdateReverseProxySlowhttpRequestHeaderTimeoutBase(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_base", bytes.NewReader(byt))
	return
}

// GetReverseProxySlowhttpRequestHeaderTimeoutEnabled gets the reverse_proxy.slowhttp_request_header_timeout_enabled value from the UTM
func GetReverseProxySlowhttpRequestHeaderTimeoutEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxySlowhttpRequestHeaderTimeoutEnabled PUTs the reverse_proxy.slowhttp_request_header_timeout_enabled value to the UTM
func UpdateReverseProxySlowhttpRequestHeaderTimeoutEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_enabled", bytes.NewReader(byt))
	return
}

// GetReverseProxySlowhttpRequestHeaderTimeoutMax gets the reverse_proxy.slowhttp_request_header_timeout_max value from the UTM
func GetReverseProxySlowhttpRequestHeaderTimeoutMax(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_max")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxySlowhttpRequestHeaderTimeoutMax PUTs the reverse_proxy.slowhttp_request_header_timeout_max value to the UTM
func UpdateReverseProxySlowhttpRequestHeaderTimeoutMax(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_max", bytes.NewReader(byt))
	return
}

// GetReverseProxySlowhttpRequestHeaderTimeoutRate gets the reverse_proxy.slowhttp_request_header_timeout_rate value from the UTM
func GetReverseProxySlowhttpRequestHeaderTimeoutRate(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_rate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxySlowhttpRequestHeaderTimeoutRate PUTs the reverse_proxy.slowhttp_request_header_timeout_rate value to the UTM
func UpdateReverseProxySlowhttpRequestHeaderTimeoutRate(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.slowhttp_request_header_timeout_rate", bytes.NewReader(byt))
	return
}

// GetReverseProxyStatus gets the reverse_proxy.status value from the UTM
func GetReverseProxyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyStatus PUTs the reverse_proxy.status value to the UTM
func UpdateReverseProxyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.status", bytes.NewReader(byt))
	return
}

// GetReverseProxyTraceEnabled gets the reverse_proxy.trace_enabled value from the UTM
func GetReverseProxyTraceEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.trace_enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyTraceEnabled PUTs the reverse_proxy.trace_enabled value to the UTM
func UpdateReverseProxyTraceEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.trace_enabled", bytes.NewReader(byt))
	return
}

// GetReverseProxyUrlhardeningsignkey gets the reverse_proxy.urlhardeningsignkey value from the UTM
func GetReverseProxyUrlhardeningsignkey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.urlhardeningsignkey")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyUrlhardeningsignkey PUTs the reverse_proxy.urlhardeningsignkey value to the UTM
func UpdateReverseProxyUrlhardeningsignkey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.urlhardeningsignkey", bytes.NewReader(byt))
	return
}

// GetReverseProxyWhatkilledus gets the reverse_proxy.whatkilledus value from the UTM
func GetReverseProxyWhatkilledus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/reverse_proxy.whatkilledus")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateReverseProxyWhatkilledus PUTs the reverse_proxy.whatkilledus value to the UTM
func UpdateReverseProxyWhatkilledus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/reverse_proxy.whatkilledus", bytes.NewReader(byt))
	return
}

// GetRoutesPolicy gets the routes.policy value from the UTM
func GetRoutesPolicy(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/routes.policy")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutesPolicy PUTs the routes.policy value to the UTM
func UpdateRoutesPolicy(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routes.policy", bytes.NewReader(byt))
	return
}

// GetRoutesStatic gets the routes.static value from the UTM
func GetRoutesStatic(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/routes.static")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutesStatic PUTs the routes.static value to the UTM
func UpdateRoutesStatic(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routes.static", bytes.NewReader(byt))
	return
}

// GetRoutingBgpMaximumPaths gets the routing.bgp.maximum_paths value from the UTM
func GetRoutingBgpMaximumPaths(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/routing.bgp.maximum_paths")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingBgpMaximumPaths PUTs the routing.bgp.maximum_paths value to the UTM
func UpdateRoutingBgpMaximumPaths(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.bgp.maximum_paths", bytes.NewReader(byt))
	return
}

// GetRoutingBgpMaximumPathsIbgp gets the routing.bgp.maximum_paths_ibgp value from the UTM
func GetRoutingBgpMaximumPathsIbgp(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/routing.bgp.maximum_paths_ibgp")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingBgpMaximumPathsIbgp PUTs the routing.bgp.maximum_paths_ibgp value to the UTM
func UpdateRoutingBgpMaximumPathsIbgp(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.bgp.maximum_paths_ibgp", bytes.NewReader(byt))
	return
}

// GetRoutingBgpMultipleAs gets the routing.bgp.multiple_as value from the UTM
func GetRoutingBgpMultipleAs(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.bgp.multiple_as")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingBgpMultipleAs PUTs the routing.bgp.multiple_as value to the UTM
func UpdateRoutingBgpMultipleAs(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.bgp.multiple_as", bytes.NewReader(byt))
	return
}

// GetRoutingBgpStatus gets the routing.bgp.status value from the UTM
func GetRoutingBgpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.bgp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingBgpStatus PUTs the routing.bgp.status value to the UTM
func UpdateRoutingBgpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.bgp.status", bytes.NewReader(byt))
	return
}

// GetRoutingBgpStrictMatch gets the routing.bgp.strict_match value from the UTM
func GetRoutingBgpStrictMatch(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.bgp.strict_match")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingBgpStrictMatch PUTs the routing.bgp.strict_match value to the UTM
func UpdateRoutingBgpStrictMatch(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.bgp.strict_match", bytes.NewReader(byt))
	return
}

// GetRoutingBgpSystems gets the routing.bgp.systems value from the UTM
func GetRoutingBgpSystems(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/routing.bgp.systems")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingBgpSystems PUTs the routing.bgp.systems value to the UTM
func UpdateRoutingBgpSystems(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.bgp.systems", bytes.NewReader(byt))
	return
}

// GetRoutingOspfAbrType gets the routing.ospf.abr_type value from the UTM
func GetRoutingOspfAbrType(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.abr_type")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfAbrType PUTs the routing.ospf.abr_type value to the UTM
func UpdateRoutingOspfAbrType(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.abr_type", bytes.NewReader(byt))
	return
}

// GetRoutingOspfAreas gets the routing.ospf.areas value from the UTM
func GetRoutingOspfAreas(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.areas")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfAreas PUTs the routing.ospf.areas value to the UTM
func UpdateRoutingOspfAreas(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.areas", bytes.NewReader(byt))
	return
}

// GetRoutingOspfDefaultInformation gets the routing.ospf.default_information value from the UTM
func GetRoutingOspfDefaultInformation(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.default_information")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfDefaultInformation PUTs the routing.ospf.default_information value to the UTM
func UpdateRoutingOspfDefaultInformation(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.default_information", bytes.NewReader(byt))
	return
}

// GetRoutingOspfDefaultRouteMetric gets the routing.ospf.default_route_metric value from the UTM
func GetRoutingOspfDefaultRouteMetric(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.default_route_metric")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfDefaultRouteMetric PUTs the routing.ospf.default_route_metric value to the UTM
func UpdateRoutingOspfDefaultRouteMetric(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.default_route_metric", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeBgpMetric gets the routing.ospf.redistribute.bgp.metric value from the UTM
func GetRoutingOspfRedistributeBgpMetric(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.bgp.metric")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeBgpMetric PUTs the routing.ospf.redistribute.bgp.metric value to the UTM
func UpdateRoutingOspfRedistributeBgpMetric(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.bgp.metric", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeBgpStatus gets the routing.ospf.redistribute.bgp.status value from the UTM
func GetRoutingOspfRedistributeBgpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.bgp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeBgpStatus PUTs the routing.ospf.redistribute.bgp.status value to the UTM
func UpdateRoutingOspfRedistributeBgpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.bgp.status", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeConnectedMetric gets the routing.ospf.redistribute.connected.metric value from the UTM
func GetRoutingOspfRedistributeConnectedMetric(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.connected.metric")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeConnectedMetric PUTs the routing.ospf.redistribute.connected.metric value to the UTM
func UpdateRoutingOspfRedistributeConnectedMetric(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.connected.metric", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeConnectedStatus gets the routing.ospf.redistribute.connected.status value from the UTM
func GetRoutingOspfRedistributeConnectedStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.connected.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeConnectedStatus PUTs the routing.ospf.redistribute.connected.status value to the UTM
func UpdateRoutingOspfRedistributeConnectedStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.connected.status", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeIpsecStatus gets the routing.ospf.redistribute.ipsec.status value from the UTM
func GetRoutingOspfRedistributeIpsecStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.ipsec.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeIpsecStatus PUTs the routing.ospf.redistribute.ipsec.status value to the UTM
func UpdateRoutingOspfRedistributeIpsecStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.ipsec.status", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeSslVpnStatus gets the routing.ospf.redistribute.ssl_vpn.status value from the UTM
func GetRoutingOspfRedistributeSslVpnStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.ssl_vpn.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeSslVpnStatus PUTs the routing.ospf.redistribute.ssl_vpn.status value to the UTM
func UpdateRoutingOspfRedistributeSslVpnStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.ssl_vpn.status", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeStaticMetric gets the routing.ospf.redistribute.static.metric value from the UTM
func GetRoutingOspfRedistributeStaticMetric(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.static.metric")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeStaticMetric PUTs the routing.ospf.redistribute.static.metric value to the UTM
func UpdateRoutingOspfRedistributeStaticMetric(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.static.metric", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRedistributeStaticStatus gets the routing.ospf.redistribute.static.status value from the UTM
func GetRoutingOspfRedistributeStaticStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.redistribute.static.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRedistributeStaticStatus PUTs the routing.ospf.redistribute.static.status value to the UTM
func UpdateRoutingOspfRedistributeStaticStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.redistribute.static.status", bytes.NewReader(byt))
	return
}

// GetRoutingOspfReferenceBandwidth gets the routing.ospf.reference_bandwidth value from the UTM
func GetRoutingOspfReferenceBandwidth(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.reference_bandwidth")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfReferenceBandwidth PUTs the routing.ospf.reference_bandwidth value to the UTM
func UpdateRoutingOspfReferenceBandwidth(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.reference_bandwidth", bytes.NewReader(byt))
	return
}

// GetRoutingOspfRouterId gets the routing.ospf.router_id value from the UTM
func GetRoutingOspfRouterId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.router_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfRouterId PUTs the routing.ospf.router_id value to the UTM
func UpdateRoutingOspfRouterId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.router_id", bytes.NewReader(byt))
	return
}

// GetRoutingOspfStatus gets the routing.ospf.status value from the UTM
func GetRoutingOspfStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.ospf.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingOspfStatus PUTs the routing.ospf.status value to the UTM
func UpdateRoutingOspfStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.ospf.status", bytes.NewReader(byt))
	return
}

// GetRoutingQuaggaAllowedNetworks gets the routing.quagga.allowed_networks value from the UTM
func GetRoutingQuaggaAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/routing.quagga.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingQuaggaAllowedNetworks PUTs the routing.quagga.allowed_networks value to the UTM
func UpdateRoutingQuaggaAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.quagga.allowed_networks", bytes.NewReader(byt))
	return
}

// GetRoutingQuaggaEnablePassword gets the routing.quagga.enable_password value from the UTM
func GetRoutingQuaggaEnablePassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/routing.quagga.enable_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingQuaggaEnablePassword PUTs the routing.quagga.enable_password value to the UTM
func UpdateRoutingQuaggaEnablePassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.quagga.enable_password", bytes.NewReader(byt))
	return
}

// GetRoutingQuaggaLinkDetect gets the routing.quagga.link_detect value from the UTM
func GetRoutingQuaggaLinkDetect(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/routing.quagga.link_detect")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingQuaggaLinkDetect PUTs the routing.quagga.link_detect value to the UTM
func UpdateRoutingQuaggaLinkDetect(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.quagga.link_detect", bytes.NewReader(byt))
	return
}

// GetRoutingQuaggaPassword gets the routing.quagga.password value from the UTM
func GetRoutingQuaggaPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/routing.quagga.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateRoutingQuaggaPassword PUTs the routing.quagga.password value to the UTM
func UpdateRoutingQuaggaPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/routing.quagga.password", bytes.NewReader(byt))
	return
}

// GetSandboxReportdDebug gets the sandbox_reportd.debug value from the UTM
func GetSandboxReportdDebug(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandbox_reportd.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxReportdDebug PUTs the sandbox_reportd.debug value to the UTM
func UpdateSandboxReportdDebug(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandbox_reportd.debug", bytes.NewReader(byt))
	return
}

// GetSandboxReportdPollInterval gets the sandbox_reportd.poll_interval value from the UTM
func GetSandboxReportdPollInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandbox_reportd.poll_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxReportdPollInterval PUTs the sandbox_reportd.poll_interval value to the UTM
func UpdateSandboxReportdPollInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandbox_reportd.poll_interval", bytes.NewReader(byt))
	return
}

// GetSandboxReportdRequestTimeout gets the sandbox_reportd.request_timeout value from the UTM
func GetSandboxReportdRequestTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandbox_reportd.request_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxReportdRequestTimeout PUTs the sandbox_reportd.request_timeout value to the UTM
func UpdateSandboxReportdRequestTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandbox_reportd.request_timeout", bytes.NewReader(byt))
	return
}

// GetSandboxReportdRetryTimes gets the sandbox_reportd.retry_times value from the UTM
func GetSandboxReportdRetryTimes(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandbox_reportd.retry_times")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxReportdRetryTimes PUTs the sandbox_reportd.retry_times value to the UTM
func UpdateSandboxReportdRetryTimes(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandbox_reportd.retry_times", bytes.NewReader(byt))
	return
}

// GetSandboxdBypassSandboxLimit gets the sandboxd.bypass_sandbox_limit value from the UTM
func GetSandboxdBypassSandboxLimit(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.bypass_sandbox_limit")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdBypassSandboxLimit PUTs the sandboxd.bypass_sandbox_limit value to the UTM
func UpdateSandboxdBypassSandboxLimit(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.bypass_sandbox_limit", bytes.NewReader(byt))
	return
}

// GetSandboxdCacheExpire gets the sandboxd.cache_expire value from the UTM
func GetSandboxdCacheExpire(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.cache_expire")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdCacheExpire PUTs the sandboxd.cache_expire value to the UTM
func UpdateSandboxdCacheExpire(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.cache_expire", bytes.NewReader(byt))
	return
}

// GetSandboxdCertstore gets the sandboxd.certstore value from the UTM
func GetSandboxdCertstore(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.certstore")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdCertstore PUTs the sandboxd.certstore value to the UTM
func UpdateSandboxdCertstore(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.certstore", bytes.NewReader(byt))
	return
}

// GetSandboxdClientTimeout gets the sandboxd.client_timeout value from the UTM
func GetSandboxdClientTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.client_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdClientTimeout PUTs the sandboxd.client_timeout value to the UTM
func UpdateSandboxdClientTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.client_timeout", bytes.NewReader(byt))
	return
}

// GetSandboxdCloudPollInterval gets the sandboxd.cloud_poll_interval value from the UTM
func GetSandboxdCloudPollInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.cloud_poll_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdCloudPollInterval PUTs the sandboxd.cloud_poll_interval value to the UTM
func UpdateSandboxdCloudPollInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.cloud_poll_interval", bytes.NewReader(byt))
	return
}

// GetSandboxdCloudPollRequestMaximum gets the sandboxd.cloud_poll_request_maximum value from the UTM
func GetSandboxdCloudPollRequestMaximum(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.cloud_poll_request_maximum")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdCloudPollRequestMaximum PUTs the sandboxd.cloud_poll_request_maximum value to the UTM
func UpdateSandboxdCloudPollRequestMaximum(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.cloud_poll_request_maximum", bytes.NewReader(byt))
	return
}

// GetSandboxdCloudPollTimeout gets the sandboxd.cloud_poll_timeout value from the UTM
func GetSandboxdCloudPollTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.cloud_poll_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdCloudPollTimeout PUTs the sandboxd.cloud_poll_timeout value to the UTM
func UpdateSandboxdCloudPollTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.cloud_poll_timeout", bytes.NewReader(byt))
	return
}

// GetSandboxdConnectTimeout gets the sandboxd.connect_timeout value from the UTM
func GetSandboxdConnectTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.connect_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdConnectTimeout PUTs the sandboxd.connect_timeout value to the UTM
func UpdateSandboxdConnectTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.connect_timeout", bytes.NewReader(byt))
	return
}

// GetSandboxdDebug gets the sandboxd.debug value from the UTM
func GetSandboxdDebug(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/sandboxd.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdDebug PUTs the sandboxd.debug value to the UTM
func UpdateSandboxdDebug(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.debug", bytes.NewReader(byt))
	return
}

// GetSandboxdDhparams2048 gets the sandboxd.dhparams_2048 value from the UTM
func GetSandboxdDhparams2048(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.dhparams_2048")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdDhparams2048 PUTs the sandboxd.dhparams_2048 value to the UTM
func UpdateSandboxdDhparams2048(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.dhparams_2048", bytes.NewReader(byt))
	return
}

// GetSandboxdEdgeServerCertValidation gets the sandboxd.edge_server_cert_validation value from the UTM
func GetSandboxdEdgeServerCertValidation(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/sandboxd.edge_server_cert_validation")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdEdgeServerCertValidation PUTs the sandboxd.edge_server_cert_validation value to the UTM
func UpdateSandboxdEdgeServerCertValidation(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.edge_server_cert_validation", bytes.NewReader(byt))
	return
}

// GetSandboxdEdgeServerHost gets the sandboxd.edge_server_host value from the UTM
func GetSandboxdEdgeServerHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.edge_server_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdEdgeServerHost PUTs the sandboxd.edge_server_host value to the UTM
func UpdateSandboxdEdgeServerHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.edge_server_host", bytes.NewReader(byt))
	return
}

// GetSandboxdEdgeServerPort gets the sandboxd.edge_server_port value from the UTM
func GetSandboxdEdgeServerPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.edge_server_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdEdgeServerPort PUTs the sandboxd.edge_server_port value to the UTM
func UpdateSandboxdEdgeServerPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.edge_server_port", bytes.NewReader(byt))
	return
}

// GetSandboxdFiletypeSkiplist gets the sandboxd.filetype_skiplist value from the UTM
func GetSandboxdFiletypeSkiplist(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/sandboxd.filetype_skiplist")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdFiletypeSkiplist PUTs the sandboxd.filetype_skiplist value to the UTM
func UpdateSandboxdFiletypeSkiplist(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.filetype_skiplist", bytes.NewReader(byt))
	return
}

// GetSandboxdNumOfEventThreads gets the sandboxd.num_of_event_threads value from the UTM
func GetSandboxdNumOfEventThreads(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.num_of_event_threads")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdNumOfEventThreads PUTs the sandboxd.num_of_event_threads value to the UTM
func UpdateSandboxdNumOfEventThreads(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.num_of_event_threads", bytes.NewReader(byt))
	return
}

// GetSandboxdNumOfWorkerThreads gets the sandboxd.num_of_worker_threads value from the UTM
func GetSandboxdNumOfWorkerThreads(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.num_of_worker_threads")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdNumOfWorkerThreads PUTs the sandboxd.num_of_worker_threads value to the UTM
func UpdateSandboxdNumOfWorkerThreads(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.num_of_worker_threads", bytes.NewReader(byt))
	return
}

// GetSandboxdResponseTimeout gets the sandboxd.response_timeout value from the UTM
func GetSandboxdResponseTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.response_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdResponseTimeout PUTs the sandboxd.response_timeout value to the UTM
func UpdateSandboxdResponseTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.response_timeout", bytes.NewReader(byt))
	return
}

// GetSandboxdRetryIntervals gets the sandboxd.retry_intervals value from the UTM
func GetSandboxdRetryIntervals(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.retry_intervals")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdRetryIntervals PUTs the sandboxd.retry_intervals value to the UTM
func UpdateSandboxdRetryIntervals(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.retry_intervals", bytes.NewReader(byt))
	return
}

// GetSandboxdSandboxEnabled gets the sandboxd.sandbox_enabled value from the UTM
func GetSandboxdSandboxEnabled(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/sandboxd.sandbox_enabled")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSandboxEnabled PUTs the sandboxd.sandbox_enabled value to the UTM
func UpdateSandboxdSandboxEnabled(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.sandbox_enabled", bytes.NewReader(byt))
	return
}

// GetSandboxdSandboxdScoreThreshold gets the sandboxd.sandboxd_score_threshold value from the UTM
func GetSandboxdSandboxdScoreThreshold(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.sandboxd_score_threshold")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSandboxdScoreThreshold PUTs the sandboxd.sandboxd_score_threshold value to the UTM
func UpdateSandboxdSandboxdScoreThreshold(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.sandboxd_score_threshold", bytes.NewReader(byt))
	return
}

// GetSandboxdSbxVersion gets the sandboxd.sbx_version value from the UTM
func GetSandboxdSbxVersion(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.sbx_version")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSbxVersion PUTs the sandboxd.sbx_version value to the UTM
func UpdateSandboxdSbxVersion(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.sbx_version", bytes.NewReader(byt))
	return
}

// GetSandboxdSqliteBusyTimeout gets the sandboxd.sqlite_busy_timeout value from the UTM
func GetSandboxdSqliteBusyTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/sandboxd.sqlite_busy_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSqliteBusyTimeout PUTs the sandboxd.sqlite_busy_timeout value to the UTM
func UpdateSandboxdSqliteBusyTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.sqlite_busy_timeout", bytes.NewReader(byt))
	return
}

// GetSandboxdSslCertFile gets the sandboxd.ssl_cert_file value from the UTM
func GetSandboxdSslCertFile(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.ssl_cert_file")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSslCertFile PUTs the sandboxd.ssl_cert_file value to the UTM
func UpdateSandboxdSslCertFile(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.ssl_cert_file", bytes.NewReader(byt))
	return
}

// GetSandboxdSslKeyFile gets the sandboxd.ssl_key_file value from the UTM
func GetSandboxdSslKeyFile(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.ssl_key_file")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSslKeyFile PUTs the sandboxd.ssl_key_file value to the UTM
func UpdateSandboxdSslKeyFile(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.ssl_key_file", bytes.NewReader(byt))
	return
}

// GetSandboxdSslciphers gets the sandboxd.sslciphers value from the UTM
func GetSandboxdSslciphers(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sandboxd.sslciphers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSandboxdSslciphers PUTs the sandboxd.sslciphers value to the UTM
func UpdateSandboxdSslciphers(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sandboxd.sslciphers", bytes.NewReader(byt))
	return
}

// GetSettingsAdminEmail gets the settings.admin_email value from the UTM
func GetSettingsAdminEmail(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.admin_email")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsAdminEmail PUTs the settings.admin_email value to the UTM
func UpdateSettingsAdminEmail(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.admin_email", bytes.NewReader(byt))
	return
}

// GetSettingsBasicSetup gets the settings.basic_setup value from the UTM
func GetSettingsBasicSetup(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/settings.basic_setup")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsBasicSetup PUTs the settings.basic_setup value to the UTM
func UpdateSettingsBasicSetup(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.basic_setup", bytes.NewReader(byt))
	return
}

// GetSettingsCcMode gets the settings.cc_mode value from the UTM
func GetSettingsCcMode(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/settings.cc_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsCcMode PUTs the settings.cc_mode value to the UTM
func UpdateSettingsCcMode(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.cc_mode", bytes.NewReader(byt))
	return
}

// GetSettingsCity gets the settings.city value from the UTM
func GetSettingsCity(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.city")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsCity PUTs the settings.city value to the UTM
func UpdateSettingsCity(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.city", bytes.NewReader(byt))
	return
}

// GetSettingsCountry gets the settings.country value from the UTM
func GetSettingsCountry(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.country")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsCountry PUTs the settings.country value to the UTM
func UpdateSettingsCountry(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.country", bytes.NewReader(byt))
	return
}

// GetSettingsExtraSwap gets the settings.extra_swap value from the UTM
func GetSettingsExtraSwap(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/settings.extra_swap")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsExtraSwap PUTs the settings.extra_swap value to the UTM
func UpdateSettingsExtraSwap(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.extra_swap", bytes.NewReader(byt))
	return
}

// GetSettingsHostname gets the settings.hostname value from the UTM
func GetSettingsHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsHostname PUTs the settings.hostname value to the UTM
func UpdateSettingsHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.hostname", bytes.NewReader(byt))
	return
}

// GetSettingsIcsaMode gets the settings.icsa_mode value from the UTM
func GetSettingsIcsaMode(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/settings.icsa_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsIcsaMode PUTs the settings.icsa_mode value to the UTM
func UpdateSettingsIcsaMode(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.icsa_mode", bytes.NewReader(byt))
	return
}

// GetSettingsOrganization gets the settings.organization value from the UTM
func GetSettingsOrganization(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.organization")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsOrganization PUTs the settings.organization value to the UTM
func UpdateSettingsOrganization(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.organization", bytes.NewReader(byt))
	return
}

// GetSettingsPasswordComplexityMinDigits gets the settings.password_complexity.min_digits value from the UTM
func GetSettingsPasswordComplexityMinDigits(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/settings.password_complexity.min_digits")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPasswordComplexityMinDigits PUTs the settings.password_complexity.min_digits value to the UTM
func UpdateSettingsPasswordComplexityMinDigits(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.password_complexity.min_digits", bytes.NewReader(byt))
	return
}

// GetSettingsPasswordComplexityMinLength gets the settings.password_complexity.min_length value from the UTM
func GetSettingsPasswordComplexityMinLength(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/settings.password_complexity.min_length")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPasswordComplexityMinLength PUTs the settings.password_complexity.min_length value to the UTM
func UpdateSettingsPasswordComplexityMinLength(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.password_complexity.min_length", bytes.NewReader(byt))
	return
}

// GetSettingsPasswordComplexityMinLowerChars gets the settings.password_complexity.min_lower_chars value from the UTM
func GetSettingsPasswordComplexityMinLowerChars(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/settings.password_complexity.min_lower_chars")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPasswordComplexityMinLowerChars PUTs the settings.password_complexity.min_lower_chars value to the UTM
func UpdateSettingsPasswordComplexityMinLowerChars(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.password_complexity.min_lower_chars", bytes.NewReader(byt))
	return
}

// GetSettingsPasswordComplexityMinSpecialChars gets the settings.password_complexity.min_special_chars value from the UTM
func GetSettingsPasswordComplexityMinSpecialChars(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/settings.password_complexity.min_special_chars")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPasswordComplexityMinSpecialChars PUTs the settings.password_complexity.min_special_chars value to the UTM
func UpdateSettingsPasswordComplexityMinSpecialChars(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.password_complexity.min_special_chars", bytes.NewReader(byt))
	return
}

// GetSettingsPasswordComplexityMinUpperChars gets the settings.password_complexity.min_upper_chars value from the UTM
func GetSettingsPasswordComplexityMinUpperChars(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/settings.password_complexity.min_upper_chars")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPasswordComplexityMinUpperChars PUTs the settings.password_complexity.min_upper_chars value to the UTM
func UpdateSettingsPasswordComplexityMinUpperChars(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.password_complexity.min_upper_chars", bytes.NewReader(byt))
	return
}

// GetSettingsPasswordComplexityStatus gets the settings.password_complexity.status value from the UTM
func GetSettingsPasswordComplexityStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/settings.password_complexity.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPasswordComplexityStatus PUTs the settings.password_complexity.status value to the UTM
func UpdateSettingsPasswordComplexityStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.password_complexity.status", bytes.NewReader(byt))
	return
}

// GetSettingsPopularity gets the settings.popularity value from the UTM
func GetSettingsPopularity(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.popularity")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsPopularity PUTs the settings.popularity value to the UTM
func UpdateSettingsPopularity(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.popularity", bytes.NewReader(byt))
	return
}

// GetSettingsRasUpdate gets the settings.ras_update value from the UTM
func GetSettingsRasUpdate(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.ras_update")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsRasUpdate PUTs the settings.ras_update value to the UTM
func UpdateSettingsRasUpdate(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.ras_update", bytes.NewReader(byt))
	return
}

// GetSettingsSystemId gets the settings.system_id value from the UTM
func GetSettingsSystemId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.system_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsSystemId PUTs the settings.system_id value to the UTM
func UpdateSettingsSystemId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.system_id", bytes.NewReader(byt))
	return
}

// GetSettingsTimezone gets the settings.timezone value from the UTM
func GetSettingsTimezone(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/settings.timezone")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSettingsTimezone PUTs the settings.timezone value to the UTM
func UpdateSettingsTimezone(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/settings.timezone", bytes.NewReader(byt))
	return
}

// GetSipAllowedNetworks gets the sip.allowed_networks value from the UTM
func GetSipAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/sip.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSipAllowedNetworks PUTs the sip.allowed_networks value to the UTM
func UpdateSipAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sip.allowed_networks", bytes.NewReader(byt))
	return
}

// GetSipExpectMode gets the sip.expect_mode value from the UTM
func GetSipExpectMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sip.expect_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSipExpectMode PUTs the sip.expect_mode value to the UTM
func UpdateSipExpectMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sip.expect_mode", bytes.NewReader(byt))
	return
}

// GetSipLogRelated gets the sip.log_related value from the UTM
func GetSipLogRelated(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/sip.log_related")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSipLogRelated PUTs the sip.log_related value to the UTM
func UpdateSipLogRelated(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sip.log_related", bytes.NewReader(byt))
	return
}

// GetSipServers gets the sip.servers value from the UTM
func GetSipServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/sip.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSipServers PUTs the sip.servers value to the UTM
func UpdateSipServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sip.servers", bytes.NewReader(byt))
	return
}

// GetSipStatus gets the sip.status value from the UTM
func GetSipStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/sip.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSipStatus PUTs the sip.status value to the UTM
func UpdateSipStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sip.status", bytes.NewReader(byt))
	return
}

// GetSmsClientHostname gets the sms_client.hostname value from the UTM
func GetSmsClientHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sms_client.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmsClientHostname PUTs the sms_client.hostname value to the UTM
func UpdateSmsClientHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sms_client.hostname", bytes.NewReader(byt))
	return
}

// GetSmsClientName gets the sms_client.name value from the UTM
func GetSmsClientName(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sms_client.name")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmsClientName PUTs the sms_client.name value to the UTM
func UpdateSmsClientName(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sms_client.name", bytes.NewReader(byt))
	return
}

// GetSmsClientPassword gets the sms_client.password value from the UTM
func GetSmsClientPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sms_client.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmsClientPassword PUTs the sms_client.password value to the UTM
func UpdateSmsClientPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sms_client.password", bytes.NewReader(byt))
	return
}

// GetSmsClientPort gets the sms_client.port value from the UTM
func GetSmsClientPort(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/sms_client.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmsClientPort PUTs the sms_client.port value to the UTM
func UpdateSmsClientPort(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sms_client.port", bytes.NewReader(byt))
	return
}

// GetSmsClientStatus gets the sms_client.status value from the UTM
func GetSmsClientStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/sms_client.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmsClientStatus PUTs the sms_client.status value to the UTM
func UpdateSmsClientStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sms_client.status", bytes.NewReader(byt))
	return
}

// GetSmsClientUsername gets the sms_client.username value from the UTM
func GetSmsClientUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/sms_client.username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmsClientUsername PUTs the sms_client.username value to the UTM
func UpdateSmsClientUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/sms_client.username", bytes.NewReader(byt))
	return
}

// GetSmtpAuthAaa gets the smtp.auth_aaa value from the UTM
func GetSmtpAuthAaa(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.auth_aaa")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpAuthAaa PUTs the smtp.auth_aaa value to the UTM
func UpdateSmtpAuthAaa(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.auth_aaa", bytes.NewReader(byt))
	return
}

// GetSmtpAuthStatus gets the smtp.auth_status value from the UTM
func GetSmtpAuthStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.auth_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpAuthStatus PUTs the smtp.auth_status value to the UTM
func UpdateSmtpAuthStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.auth_status", bytes.NewReader(byt))
	return
}

// GetSmtpAvFooter gets the smtp.av_footer value from the UTM
func GetSmtpAvFooter(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.av_footer")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpAvFooter PUTs the smtp.av_footer value to the UTM
func UpdateSmtpAvFooter(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.av_footer", bytes.NewReader(byt))
	return
}

// GetSmtpAvFooterStatus gets the smtp.av_footer_status value from the UTM
func GetSmtpAvFooterStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.av_footer_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpAvFooterStatus PUTs the smtp.av_footer_status value to the UTM
func UpdateSmtpAvFooterStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.av_footer_status", bytes.NewReader(byt))
	return
}

// GetSmtpBatvSecret gets the smtp.batv_secret value from the UTM
func GetSmtpBatvSecret(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.batv_secret")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpBatvSecret PUTs the smtp.batv_secret value to the UTM
func UpdateSmtpBatvSecret(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.batv_secret", bytes.NewReader(byt))
	return
}

// GetSmtpBlockerService gets the smtp.blocker_service value from the UTM
func GetSmtpBlockerService(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.blocker_service")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpBlockerService PUTs the smtp.blocker_service value to the UTM
func UpdateSmtpBlockerService(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.blocker_service", bytes.NewReader(byt))
	return
}

// GetSmtpCffAsMarker gets the smtp.cff_as_marker value from the UTM
func GetSmtpCffAsMarker(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.cff_as_marker")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpCffAsMarker PUTs the smtp.cff_as_marker value to the UTM
func UpdateSmtpCffAsMarker(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.cff_as_marker", bytes.NewReader(byt))
	return
}

// GetSmtpDkimDomains gets the smtp.dkim_domains value from the UTM
func GetSmtpDkimDomains(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.dkim_domains")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpDkimDomains PUTs the smtp.dkim_domains value to the UTM
func UpdateSmtpDkimDomains(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.dkim_domains", bytes.NewReader(byt))
	return
}

// GetSmtpDkimPrivateKey gets the smtp.dkim_private_key value from the UTM
func GetSmtpDkimPrivateKey(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.dkim_private_key")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpDkimPrivateKey PUTs the smtp.dkim_private_key value to the UTM
func UpdateSmtpDkimPrivateKey(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.dkim_private_key", bytes.NewReader(byt))
	return
}

// GetSmtpDkimSelector gets the smtp.dkim_selector value from the UTM
func GetSmtpDkimSelector(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.dkim_selector")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpDkimSelector PUTs the smtp.dkim_selector value to the UTM
func UpdateSmtpDkimSelector(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.dkim_selector", bytes.NewReader(byt))
	return
}

// GetSmtpEnableOldExpressionFilter gets the smtp.enable_old_expression_filter value from the UTM
func GetSmtpEnableOldExpressionFilter(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.enable_old_expression_filter")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpEnableOldExpressionFilter PUTs the smtp.enable_old_expression_filter value to the UTM
func UpdateSmtpEnableOldExpressionFilter(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.enable_old_expression_filter", bytes.NewReader(byt))
	return
}

// GetSmtpEncryptionUtility gets the smtp.encryption_utility value from the UTM
func GetSmtpEncryptionUtility(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.encryption_utility")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpEncryptionUtility PUTs the smtp.encryption_utility value to the UTM
func UpdateSmtpEncryptionUtility(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.encryption_utility", bytes.NewReader(byt))
	return
}

// GetSmtpExceptions gets the smtp.exceptions value from the UTM
func GetSmtpExceptions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.exceptions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpExceptions PUTs the smtp.exceptions value to the UTM
func UpdateSmtpExceptions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.exceptions", bytes.NewReader(byt))
	return
}

// GetSmtpFootersMode gets the smtp.footers_mode value from the UTM
func GetSmtpFootersMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.footers_mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpFootersMode PUTs the smtp.footers_mode value to the UTM
func UpdateSmtpFootersMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.footers_mode", bytes.NewReader(byt))
	return
}

// GetSmtpGlobalAsReject gets the smtp.global_as_reject value from the UTM
func GetSmtpGlobalAsReject(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.global_as_reject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpGlobalAsReject PUTs the smtp.global_as_reject value to the UTM
func UpdateSmtpGlobalAsReject(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.global_as_reject", bytes.NewReader(byt))
	return
}

// GetSmtpGlobalAvReject gets the smtp.global_av_reject value from the UTM
func GetSmtpGlobalAvReject(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.global_av_reject")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpGlobalAvReject PUTs the smtp.global_av_reject value to the UTM
func UpdateSmtpGlobalAvReject(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.global_av_reject", bytes.NewReader(byt))
	return
}

// GetSmtpGlobalProfile gets the smtp.global_profile value from the UTM
func GetSmtpGlobalProfile(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.global_profile")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpGlobalProfile PUTs the smtp.global_profile value to the UTM
func UpdateSmtpGlobalProfile(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.global_profile", bytes.NewReader(byt))
	return
}

// GetSmtpHostBlacklist gets the smtp.host_blacklist value from the UTM
func GetSmtpHostBlacklist(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.host_blacklist")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpHostBlacklist PUTs the smtp.host_blacklist value to the UTM
func UpdateSmtpHostBlacklist(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.host_blacklist", bytes.NewReader(byt))
	return
}

// GetSmtpHostname gets the smtp.hostname value from the UTM
func GetSmtpHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpHostname PUTs the smtp.hostname value to the UTM
func UpdateSmtpHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.hostname", bytes.NewReader(byt))
	return
}

// GetSmtpMaxMessageSize gets the smtp.max_message_size value from the UTM
func GetSmtpMaxMessageSize(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.max_message_size")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpMaxMessageSize PUTs the smtp.max_message_size value to the UTM
func UpdateSmtpMaxMessageSize(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.max_message_size", bytes.NewReader(byt))
	return
}

// GetSmtpMode gets the smtp.mode value from the UTM
func GetSmtpMode(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.mode")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpMode PUTs the smtp.mode value to the UTM
func UpdateSmtpMode(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.mode", bytes.NewReader(byt))
	return
}

// GetSmtpParentProxyAuthPass gets the smtp.parent_proxy_auth_pass value from the UTM
func GetSmtpParentProxyAuthPass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.parent_proxy_auth_pass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpParentProxyAuthPass PUTs the smtp.parent_proxy_auth_pass value to the UTM
func UpdateSmtpParentProxyAuthPass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.parent_proxy_auth_pass", bytes.NewReader(byt))
	return
}

// GetSmtpParentProxyAuthStatus gets the smtp.parent_proxy_auth_status value from the UTM
func GetSmtpParentProxyAuthStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.parent_proxy_auth_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpParentProxyAuthStatus PUTs the smtp.parent_proxy_auth_status value to the UTM
func UpdateSmtpParentProxyAuthStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.parent_proxy_auth_status", bytes.NewReader(byt))
	return
}

// GetSmtpParentProxyAuthUser gets the smtp.parent_proxy_auth_user value from the UTM
func GetSmtpParentProxyAuthUser(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.parent_proxy_auth_user")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpParentProxyAuthUser PUTs the smtp.parent_proxy_auth_user value to the UTM
func UpdateSmtpParentProxyAuthUser(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.parent_proxy_auth_user", bytes.NewReader(byt))
	return
}

// GetSmtpParentProxyHost gets the smtp.parent_proxy_host value from the UTM
func GetSmtpParentProxyHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.parent_proxy_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpParentProxyHost PUTs the smtp.parent_proxy_host value to the UTM
func UpdateSmtpParentProxyHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.parent_proxy_host", bytes.NewReader(byt))
	return
}

// GetSmtpParentProxyPort gets the smtp.parent_proxy_port value from the UTM
func GetSmtpParentProxyPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.parent_proxy_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpParentProxyPort PUTs the smtp.parent_proxy_port value to the UTM
func UpdateSmtpParentProxyPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.parent_proxy_port", bytes.NewReader(byt))
	return
}

// GetSmtpParentProxyStatus gets the smtp.parent_proxy_status value from the UTM
func GetSmtpParentProxyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.parent_proxy_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpParentProxyStatus PUTs the smtp.parent_proxy_status value to the UTM
func UpdateSmtpParentProxyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.parent_proxy_status", bytes.NewReader(byt))
	return
}

// GetSmtpPostmaster gets the smtp.postmaster value from the UTM
func GetSmtpPostmaster(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.postmaster")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpPostmaster PUTs the smtp.postmaster value to the UTM
func UpdateSmtpPostmaster(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.postmaster", bytes.NewReader(byt))
	return
}

// GetSmtpProfiles gets the smtp.profiles value from the UTM
func GetSmtpProfiles(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.profiles")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpProfiles PUTs the smtp.profiles value to the UTM
func UpdateSmtpProfiles(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.profiles", bytes.NewReader(byt))
	return
}

// GetSmtpRecipientsMax gets the smtp.recipients_max value from the UTM
func GetSmtpRecipientsMax(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.recipients_max")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpRecipientsMax PUTs the smtp.recipients_max value to the UTM
func UpdateSmtpRecipientsMax(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.recipients_max", bytes.NewReader(byt))
	return
}

// GetSmtpRelays gets the smtp.relays value from the UTM
func GetSmtpRelays(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.relays")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpRelays PUTs the smtp.relays value to the UTM
func UpdateSmtpRelays(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.relays", bytes.NewReader(byt))
	return
}

// GetSmtpScanOutgoingEmails gets the smtp.scan_outgoing_emails value from the UTM
func GetSmtpScanOutgoingEmails(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.scan_outgoing_emails")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpScanOutgoingEmails PUTs the smtp.scan_outgoing_emails value to the UTM
func UpdateSmtpScanOutgoingEmails(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.scan_outgoing_emails", bytes.NewReader(byt))
	return
}

// GetSmtpScannerPoolMaxPool gets the smtp.scanner_pool.max_pool value from the UTM
func GetSmtpScannerPoolMaxPool(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.scanner_pool.max_pool")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpScannerPoolMaxPool PUTs the smtp.scanner_pool.max_pool value to the UTM
func UpdateSmtpScannerPoolMaxPool(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.scanner_pool.max_pool", bytes.NewReader(byt))
	return
}

// GetSmtpScannerPoolThresholds gets the smtp.scanner_pool.thresholds value from the UTM
func GetSmtpScannerPoolThresholds(c sophos.ClientInterface) (val []int64, err error) {
	res, err := c.Get("/api/nodes/smtp.scanner_pool.thresholds")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpScannerPoolThresholds PUTs the smtp.scanner_pool.thresholds value to the UTM
func UpdateSmtpScannerPoolThresholds(c sophos.ClientInterface, val []int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.scanner_pool.thresholds", bytes.NewReader(byt))
	return
}

// GetSmtpScannerTimeout gets the smtp.scanner_timeout value from the UTM
func GetSmtpScannerTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.scanner_timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpScannerTimeout PUTs the smtp.scanner_timeout value to the UTM
func UpdateSmtpScannerTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.scanner_timeout", bytes.NewReader(byt))
	return
}

// GetSmtpSmarthostAuth gets the smtp.smarthost_auth value from the UTM
func GetSmtpSmarthostAuth(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.smarthost_auth")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmarthostAuth PUTs the smtp.smarthost_auth value to the UTM
func UpdateSmtpSmarthostAuth(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smarthost_auth", bytes.NewReader(byt))
	return
}

// GetSmtpSmarthostHostname gets the smtp.smarthost_hostname value from the UTM
func GetSmtpSmarthostHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.smarthost_hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmarthostHostname PUTs the smtp.smarthost_hostname value to the UTM
func UpdateSmtpSmarthostHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smarthost_hostname", bytes.NewReader(byt))
	return
}

// GetSmtpSmarthostPass gets the smtp.smarthost_pass value from the UTM
func GetSmtpSmarthostPass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.smarthost_pass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmarthostPass PUTs the smtp.smarthost_pass value to the UTM
func UpdateSmtpSmarthostPass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smarthost_pass", bytes.NewReader(byt))
	return
}

// GetSmtpSmarthostPort gets the smtp.smarthost_port value from the UTM
func GetSmtpSmarthostPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.smarthost_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmarthostPort PUTs the smtp.smarthost_port value to the UTM
func UpdateSmtpSmarthostPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smarthost_port", bytes.NewReader(byt))
	return
}

// GetSmtpSmarthostStatus gets the smtp.smarthost_status value from the UTM
func GetSmtpSmarthostStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.smarthost_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmarthostStatus PUTs the smtp.smarthost_status value to the UTM
func UpdateSmtpSmarthostStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smarthost_status", bytes.NewReader(byt))
	return
}

// GetSmtpSmarthostUser gets the smtp.smarthost_user value from the UTM
func GetSmtpSmarthostUser(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.smarthost_user")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmarthostUser PUTs the smtp.smarthost_user value to the UTM
func UpdateSmtpSmarthostUser(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smarthost_user", bytes.NewReader(byt))
	return
}

// GetSmtpSmtpAcceptMax gets the smtp.smtp_accept_max value from the UTM
func GetSmtpSmtpAcceptMax(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.smtp_accept_max")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmtpAcceptMax PUTs the smtp.smtp_accept_max value to the UTM
func UpdateSmtpSmtpAcceptMax(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smtp_accept_max", bytes.NewReader(byt))
	return
}

// GetSmtpSmtpAcceptPerConnectionMax gets the smtp.smtp_accept_per_connection_max value from the UTM
func GetSmtpSmtpAcceptPerConnectionMax(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.smtp_accept_per_connection_max")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmtpAcceptPerConnectionMax PUTs the smtp.smtp_accept_per_connection_max value to the UTM
func UpdateSmtpSmtpAcceptPerConnectionMax(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smtp_accept_per_connection_max", bytes.NewReader(byt))
	return
}

// GetSmtpSmtpAcceptPerHostMax gets the smtp.smtp_accept_per_host_max value from the UTM
func GetSmtpSmtpAcceptPerHostMax(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/smtp.smtp_accept_per_host_max")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpSmtpAcceptPerHostMax PUTs the smtp.smtp_accept_per_host_max value to the UTM
func UpdateSmtpSmtpAcceptPerHostMax(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.smtp_accept_per_host_max", bytes.NewReader(byt))
	return
}

// GetSmtpStatus gets the smtp.status value from the UTM
func GetSmtpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpStatus PUTs the smtp.status value to the UTM
func UpdateSmtpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.status", bytes.NewReader(byt))
	return
}

// GetSmtpTlsAvoid gets the smtp.tls_avoid value from the UTM
func GetSmtpTlsAvoid(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.tls_avoid")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTlsAvoid PUTs the smtp.tls_avoid value to the UTM
func UpdateSmtpTlsAvoid(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.tls_avoid", bytes.NewReader(byt))
	return
}

// GetSmtpTlsCert gets the smtp.tls_cert value from the UTM
func GetSmtpTlsCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.tls_cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTlsCert PUTs the smtp.tls_cert value to the UTM
func UpdateSmtpTlsCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.tls_cert", bytes.NewReader(byt))
	return
}

// GetSmtpTlsRequire gets the smtp.tls_require value from the UTM
func GetSmtpTlsRequire(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.tls_require")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTlsRequire PUTs the smtp.tls_require value to the UTM
func UpdateSmtpTlsRequire(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.tls_require", bytes.NewReader(byt))
	return
}

// GetSmtpTlsRequireSenderDomains gets the smtp.tls_require_sender_domains value from the UTM
func GetSmtpTlsRequireSenderDomains(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.tls_require_sender_domains")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTlsRequireSenderDomains PUTs the smtp.tls_require_sender_domains value to the UTM
func UpdateSmtpTlsRequireSenderDomains(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.tls_require_sender_domains", bytes.NewReader(byt))
	return
}

// GetSmtpTlsVersion gets the smtp.tls_version value from the UTM
func GetSmtpTlsVersion(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/smtp.tls_version")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTlsVersion PUTs the smtp.tls_version value to the UTM
func UpdateSmtpTlsVersion(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.tls_version", bytes.NewReader(byt))
	return
}

// GetSmtpTransparent gets the smtp.transparent value from the UTM
func GetSmtpTransparent(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.transparent")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTransparent PUTs the smtp.transparent value to the UTM
func UpdateSmtpTransparent(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.transparent", bytes.NewReader(byt))
	return
}

// GetSmtpTransparentSkip gets the smtp.transparent_skip value from the UTM
func GetSmtpTransparentSkip(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.transparent_skip")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTransparentSkip PUTs the smtp.transparent_skip value to the UTM
func UpdateSmtpTransparentSkip(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.transparent_skip", bytes.NewReader(byt))
	return
}

// GetSmtpTransparentSkipAutoPf gets the smtp.transparent_skip_auto_pf value from the UTM
func GetSmtpTransparentSkipAutoPf(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.transparent_skip_auto_pf")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpTransparentSkipAutoPf PUTs the smtp.transparent_skip_auto_pf value to the UTM
func UpdateSmtpTransparentSkipAutoPf(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.transparent_skip_auto_pf", bytes.NewReader(byt))
	return
}

// GetSmtpUpstreamHosts gets the smtp.upstream_hosts value from the UTM
func GetSmtpUpstreamHosts(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/smtp.upstream_hosts")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpUpstreamHosts PUTs the smtp.upstream_hosts value to the UTM
func UpdateSmtpUpstreamHosts(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.upstream_hosts", bytes.NewReader(byt))
	return
}

// GetSmtpUpstreamHostsOnly gets the smtp.upstream_hosts_only value from the UTM
func GetSmtpUpstreamHostsOnly(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/smtp.upstream_hosts_only")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSmtpUpstreamHostsOnly PUTs the smtp.upstream_hosts_only value to the UTM
func UpdateSmtpUpstreamHostsOnly(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/smtp.upstream_hosts_only", bytes.NewReader(byt))
	return
}

// GetSnmpAllowedNetworks gets the snmp.allowed_networks value from the UTM
func GetSnmpAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/snmp.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpAllowedNetworks PUTs the snmp.allowed_networks value to the UTM
func UpdateSnmpAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.allowed_networks", bytes.NewReader(byt))
	return
}

// GetSnmpAuthPassword gets the snmp.auth_password value from the UTM
func GetSnmpAuthPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.auth_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpAuthPassword PUTs the snmp.auth_password value to the UTM
func UpdateSnmpAuthPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.auth_password", bytes.NewReader(byt))
	return
}

// GetSnmpAuthType gets the snmp.auth_type value from the UTM
func GetSnmpAuthType(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.auth_type")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpAuthType PUTs the snmp.auth_type value to the UTM
func UpdateSnmpAuthType(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.auth_type", bytes.NewReader(byt))
	return
}

// GetSnmpCommunity gets the snmp.community value from the UTM
func GetSnmpCommunity(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.community")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpCommunity PUTs the snmp.community value to the UTM
func UpdateSnmpCommunity(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.community", bytes.NewReader(byt))
	return
}

// GetSnmpDeviceAdmin gets the snmp.device_admin value from the UTM
func GetSnmpDeviceAdmin(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.device_admin")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpDeviceAdmin PUTs the snmp.device_admin value to the UTM
func UpdateSnmpDeviceAdmin(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.device_admin", bytes.NewReader(byt))
	return
}

// GetSnmpDeviceLocation gets the snmp.device_location value from the UTM
func GetSnmpDeviceLocation(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.device_location")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpDeviceLocation PUTs the snmp.device_location value to the UTM
func UpdateSnmpDeviceLocation(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.device_location", bytes.NewReader(byt))
	return
}

// GetSnmpDeviceName gets the snmp.device_name value from the UTM
func GetSnmpDeviceName(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.device_name")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpDeviceName PUTs the snmp.device_name value to the UTM
func UpdateSnmpDeviceName(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.device_name", bytes.NewReader(byt))
	return
}

// GetSnmpEncryptPassword gets the snmp.encrypt_password value from the UTM
func GetSnmpEncryptPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.encrypt_password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpEncryptPassword PUTs the snmp.encrypt_password value to the UTM
func UpdateSnmpEncryptPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.encrypt_password", bytes.NewReader(byt))
	return
}

// GetSnmpEncryptType gets the snmp.encrypt_type value from the UTM
func GetSnmpEncryptType(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.encrypt_type")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpEncryptType PUTs the snmp.encrypt_type value to the UTM
func UpdateSnmpEncryptType(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.encrypt_type", bytes.NewReader(byt))
	return
}

// GetSnmpStatus gets the snmp.status value from the UTM
func GetSnmpStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/snmp.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpStatus PUTs the snmp.status value to the UTM
func UpdateSnmpStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.status", bytes.NewReader(byt))
	return
}

// GetSnmpTraps gets the snmp.traps value from the UTM
func GetSnmpTraps(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/snmp.traps")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpTraps PUTs the snmp.traps value to the UTM
func UpdateSnmpTraps(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.traps", bytes.NewReader(byt))
	return
}

// GetSnmpTrapsUseOldOids gets the snmp.traps_use_old_oids value from the UTM
func GetSnmpTrapsUseOldOids(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/snmp.traps_use_old_oids")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpTrapsUseOldOids PUTs the snmp.traps_use_old_oids value to the UTM
func UpdateSnmpTrapsUseOldOids(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.traps_use_old_oids", bytes.NewReader(byt))
	return
}

// GetSnmpUsername gets the snmp.username value from the UTM
func GetSnmpUsername(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.username")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpUsername PUTs the snmp.username value to the UTM
func UpdateSnmpUsername(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.username", bytes.NewReader(byt))
	return
}

// GetSnmpVersion gets the snmp.version value from the UTM
func GetSnmpVersion(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/snmp.version")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSnmpVersion PUTs the snmp.version value to the UTM
func UpdateSnmpVersion(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/snmp.version", bytes.NewReader(byt))
	return
}

// GetSocksAaa gets the socks.aaa value from the UTM
func GetSocksAaa(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/socks.aaa")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSocksAaa PUTs the socks.aaa value to the UTM
func UpdateSocksAaa(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/socks.aaa", bytes.NewReader(byt))
	return
}

// GetSocksAllowedNetworks gets the socks.allowed_networks value from the UTM
func GetSocksAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/socks.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSocksAllowedNetworks PUTs the socks.allowed_networks value to the UTM
func UpdateSocksAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/socks.allowed_networks", bytes.NewReader(byt))
	return
}

// GetSocksAuthentication gets the socks.authentication value from the UTM
func GetSocksAuthentication(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/socks.authentication")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSocksAuthentication PUTs the socks.authentication value to the UTM
func UpdateSocksAuthentication(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/socks.authentication", bytes.NewReader(byt))
	return
}

// GetSocksStatus gets the socks.status value from the UTM
func GetSocksStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/socks.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSocksStatus PUTs the socks.status value to the UTM
func UpdateSocksStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/socks.status", bytes.NewReader(byt))
	return
}

// GetSpxGlobalErrorNotificationTarget gets the spx.global.error_notification_target value from the UTM
func GetSpxGlobalErrorNotificationTarget(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.global.error_notification_target")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalErrorNotificationTarget PUTs the spx.global.error_notification_target value to the UTM
func UpdateSpxGlobalErrorNotificationTarget(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.error_notification_target", bytes.NewReader(byt))
	return
}

// GetSpxGlobalExpirySettingsAllowSecureReplyDays gets the spx.global.expiry_settings.allow_secure_reply_days value from the UTM
func GetSpxGlobalExpirySettingsAllowSecureReplyDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.global.expiry_settings.allow_secure_reply_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalExpirySettingsAllowSecureReplyDays PUTs the spx.global.expiry_settings.allow_secure_reply_days value to the UTM
func UpdateSpxGlobalExpirySettingsAllowSecureReplyDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.expiry_settings.allow_secure_reply_days", bytes.NewReader(byt))
	return
}

// GetSpxGlobalExpirySettingsKeepDelayedMailDays gets the spx.global.expiry_settings.keep_delayed_mail_days value from the UTM
func GetSpxGlobalExpirySettingsKeepDelayedMailDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.global.expiry_settings.keep_delayed_mail_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalExpirySettingsKeepDelayedMailDays PUTs the spx.global.expiry_settings.keep_delayed_mail_days value to the UTM
func UpdateSpxGlobalExpirySettingsKeepDelayedMailDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.expiry_settings.keep_delayed_mail_days", bytes.NewReader(byt))
	return
}

// GetSpxGlobalExpirySettingsKeepUnusedPwdDays gets the spx.global.expiry_settings.keep_unused_pwd_days value from the UTM
func GetSpxGlobalExpirySettingsKeepUnusedPwdDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.global.expiry_settings.keep_unused_pwd_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalExpirySettingsKeepUnusedPwdDays PUTs the spx.global.expiry_settings.keep_unused_pwd_days value to the UTM
func UpdateSpxGlobalExpirySettingsKeepUnusedPwdDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.expiry_settings.keep_unused_pwd_days", bytes.NewReader(byt))
	return
}

// GetSpxGlobalExpirySettingsRegistrationReminderDays gets the spx.global.expiry_settings.registration_reminder_days value from the UTM
func GetSpxGlobalExpirySettingsRegistrationReminderDays(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.global.expiry_settings.registration_reminder_days")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalExpirySettingsRegistrationReminderDays PUTs the spx.global.expiry_settings.registration_reminder_days value to the UTM
func UpdateSpxGlobalExpirySettingsRegistrationReminderDays(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.expiry_settings.registration_reminder_days", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPasswordStrengthMinLength gets the spx.global.password_strength.min_length value from the UTM
func GetSpxGlobalPasswordStrengthMinLength(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.global.password_strength.min_length")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPasswordStrengthMinLength PUTs the spx.global.password_strength.min_length value to the UTM
func UpdateSpxGlobalPasswordStrengthMinLength(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.password_strength.min_length", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPasswordStrengthRequireSpecChars gets the spx.global.password_strength.require_spec_chars value from the UTM
func GetSpxGlobalPasswordStrengthRequireSpecChars(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/spx.global.password_strength.require_spec_chars")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPasswordStrengthRequireSpecChars PUTs the spx.global.password_strength.require_spec_chars value to the UTM
func UpdateSpxGlobalPasswordStrengthRequireSpecChars(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.password_strength.require_spec_chars", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPortalSettingsAllowedNetworks gets the spx.global.portal_settings.allowed_networks value from the UTM
func GetSpxGlobalPortalSettingsAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/spx.global.portal_settings.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPortalSettingsAllowedNetworks PUTs the spx.global.portal_settings.allowed_networks value to the UTM
func UpdateSpxGlobalPortalSettingsAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.portal_settings.allowed_networks", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPortalSettingsHostname gets the spx.global.portal_settings.hostname value from the UTM
func GetSpxGlobalPortalSettingsHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.global.portal_settings.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPortalSettingsHostname PUTs the spx.global.portal_settings.hostname value to the UTM
func UpdateSpxGlobalPortalSettingsHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.portal_settings.hostname", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPortalSettingsInterfaceAddress gets the spx.global.portal_settings.interface_address value from the UTM
func GetSpxGlobalPortalSettingsInterfaceAddress(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.global.portal_settings.interface_address")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPortalSettingsInterfaceAddress PUTs the spx.global.portal_settings.interface_address value to the UTM
func UpdateSpxGlobalPortalSettingsInterfaceAddress(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.portal_settings.interface_address", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPortalSettingsPort gets the spx.global.portal_settings.port value from the UTM
func GetSpxGlobalPortalSettingsPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.global.portal_settings.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPortalSettingsPort PUTs the spx.global.portal_settings.port value to the UTM
func UpdateSpxGlobalPortalSettingsPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.portal_settings.port", bytes.NewReader(byt))
	return
}

// GetSpxGlobalPoweredByLogo gets the spx.global.powered_by_logo value from the UTM
func GetSpxGlobalPoweredByLogo(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.global.powered_by_logo")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalPoweredByLogo PUTs the spx.global.powered_by_logo value to the UTM
func UpdateSpxGlobalPoweredByLogo(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.powered_by_logo", bytes.NewReader(byt))
	return
}

// GetSpxGlobalSpxPriority gets the spx.global.spx_priority value from the UTM
func GetSpxGlobalSpxPriority(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/spx.global.spx_priority")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalSpxPriority PUTs the spx.global.spx_priority value to the UTM
func UpdateSpxGlobalSpxPriority(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.spx_priority", bytes.NewReader(byt))
	return
}

// GetSpxGlobalStatus gets the spx.global.status value from the UTM
func GetSpxGlobalStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/spx.global.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxGlobalStatus PUTs the spx.global.status value to the UTM
func UpdateSpxGlobalStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.global.status", bytes.NewReader(byt))
	return
}

// GetSpxSpxAuthPassword gets the spx.spx_auth.password value from the UTM
func GetSpxSpxAuthPassword(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.spx_auth.password")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxSpxAuthPassword PUTs the spx.spx_auth.password value to the UTM
func UpdateSpxSpxAuthPassword(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.spx_auth.password", bytes.NewReader(byt))
	return
}

// GetSpxSpxAuthPort gets the spx.spx_auth.port value from the UTM
func GetSpxSpxAuthPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/spx.spx_auth.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxSpxAuthPort PUTs the spx.spx_auth.port value to the UTM
func UpdateSpxSpxAuthPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.spx_auth.port", bytes.NewReader(byt))
	return
}

// GetSpxSpxAuthServer gets the spx.spx_auth.server value from the UTM
func GetSpxSpxAuthServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.spx_auth.server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxSpxAuthServer PUTs the spx.spx_auth.server value to the UTM
func UpdateSpxSpxAuthServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.spx_auth.server", bytes.NewReader(byt))
	return
}

// GetSpxSpxAuthUrl gets the spx.spx_auth.url value from the UTM
func GetSpxSpxAuthUrl(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/spx.spx_auth.url")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxSpxAuthUrl PUTs the spx.spx_auth.url value to the UTM
func UpdateSpxSpxAuthUrl(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.spx_auth.url", bytes.NewReader(byt))
	return
}

// GetSpxTemplates gets the spx.templates value from the UTM
func GetSpxTemplates(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/spx.templates")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSpxTemplates PUTs the spx.templates value to the UTM
func UpdateSpxTemplates(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/spx.templates", bytes.NewReader(byt))
	return
}

// GetSshAllowedNetworks gets the ssh.allowed_networks value from the UTM
func GetSshAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ssh.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshAllowedNetworks PUTs the ssh.allowed_networks value to the UTM
func UpdateSshAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.allowed_networks", bytes.NewReader(byt))
	return
}

// GetSshLoginKeys gets the ssh.login_keys value from the UTM
func GetSshLoginKeys(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ssh.login_keys")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshLoginKeys PUTs the ssh.login_keys value to the UTM
func UpdateSshLoginKeys(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.login_keys", bytes.NewReader(byt))
	return
}

// GetSshPasswordAuth gets the ssh.password_auth value from the UTM
func GetSshPasswordAuth(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssh.password_auth")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshPasswordAuth PUTs the ssh.password_auth value to the UTM
func UpdateSshPasswordAuth(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.password_auth", bytes.NewReader(byt))
	return
}

// GetSshPort gets the ssh.port value from the UTM
func GetSshPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ssh.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshPort PUTs the ssh.port value to the UTM
func UpdateSshPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.port", bytes.NewReader(byt))
	return
}

// GetSshPubkeyAuth gets the ssh.pubkey_auth value from the UTM
func GetSshPubkeyAuth(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssh.pubkey_auth")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshPubkeyAuth PUTs the ssh.pubkey_auth value to the UTM
func UpdateSshPubkeyAuth(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.pubkey_auth", bytes.NewReader(byt))
	return
}

// GetSshRootKeys gets the ssh.root_keys value from the UTM
func GetSshRootKeys(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/ssh.root_keys")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshRootKeys PUTs the ssh.root_keys value to the UTM
func UpdateSshRootKeys(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.root_keys", bytes.NewReader(byt))
	return
}

// GetSshRootLogin gets the ssh.root_login value from the UTM
func GetSshRootLogin(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssh.root_login")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshRootLogin PUTs the ssh.root_login value to the UTM
func UpdateSshRootLogin(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.root_login", bytes.NewReader(byt))
	return
}

// GetSshStatus gets the ssh.status value from the UTM
func GetSshStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssh.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSshStatus PUTs the ssh.status value to the UTM
func UpdateSshStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssh.status", bytes.NewReader(byt))
	return
}

// GetSslVpnAuthenticationAlgorithm gets the ssl_vpn.authentication_algorithm value from the UTM
func GetSslVpnAuthenticationAlgorithm(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.authentication_algorithm")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnAuthenticationAlgorithm PUTs the ssl_vpn.authentication_algorithm value to the UTM
func UpdateSslVpnAuthenticationAlgorithm(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.authentication_algorithm", bytes.NewReader(byt))
	return
}

// GetSslVpnCertificate gets the ssl_vpn.certificate value from the UTM
func GetSslVpnCertificate(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.certificate")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnCertificate PUTs the ssl_vpn.certificate value to the UTM
func UpdateSslVpnCertificate(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.certificate", bytes.NewReader(byt))
	return
}

// GetSslVpnCompression gets the ssl_vpn.compression value from the UTM
func GetSslVpnCompression(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.compression")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnCompression PUTs the ssl_vpn.compression value to the UTM
func UpdateSslVpnCompression(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.compression", bytes.NewReader(byt))
	return
}

// GetSslVpnDatachannelKeyLifetime gets the ssl_vpn.datachannel_key_lifetime value from the UTM
func GetSslVpnDatachannelKeyLifetime(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.datachannel_key_lifetime")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnDatachannelKeyLifetime PUTs the ssl_vpn.datachannel_key_lifetime value to the UTM
func UpdateSslVpnDatachannelKeyLifetime(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.datachannel_key_lifetime", bytes.NewReader(byt))
	return
}

// GetSslVpnDebug gets the ssl_vpn.debug value from the UTM
func GetSslVpnDebug(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.debug")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnDebug PUTs the ssl_vpn.debug value to the UTM
func UpdateSslVpnDebug(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.debug", bytes.NewReader(byt))
	return
}

// GetSslVpnDhKeySize gets the ssl_vpn.dh_key_size value from the UTM
func GetSslVpnDhKeySize(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.dh_key_size")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnDhKeySize PUTs the ssl_vpn.dh_key_size value to the UTM
func UpdateSslVpnDhKeySize(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.dh_key_size", bytes.NewReader(byt))
	return
}

// GetSslVpnDuplicateCn gets the ssl_vpn.duplicate_cn value from the UTM
func GetSslVpnDuplicateCn(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.duplicate_cn")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnDuplicateCn PUTs the ssl_vpn.duplicate_cn value to the UTM
func UpdateSslVpnDuplicateCn(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.duplicate_cn", bytes.NewReader(byt))
	return
}

// GetSslVpnEncryptionAlgorithm gets the ssl_vpn.encryption_algorithm value from the UTM
func GetSslVpnEncryptionAlgorithm(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.encryption_algorithm")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnEncryptionAlgorithm PUTs the ssl_vpn.encryption_algorithm value to the UTM
func UpdateSslVpnEncryptionAlgorithm(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.encryption_algorithm", bytes.NewReader(byt))
	return
}

// GetSslVpnHostname gets the ssl_vpn.hostname value from the UTM
func GetSslVpnHostname(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.hostname")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnHostname PUTs the ssl_vpn.hostname value to the UTM
func UpdateSslVpnHostname(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.hostname", bytes.NewReader(byt))
	return
}

// GetSslVpnInterface gets the ssl_vpn.interface value from the UTM
func GetSslVpnInterface(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.interface")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnInterface PUTs the ssl_vpn.interface value to the UTM
func UpdateSslVpnInterface(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.interface", bytes.NewReader(byt))
	return
}

// GetSslVpnInterfaceAddress gets the ssl_vpn.interface_address value from the UTM
func GetSslVpnInterfaceAddress(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.interface_address")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnInterfaceAddress PUTs the ssl_vpn.interface_address value to the UTM
func UpdateSslVpnInterfaceAddress(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.interface_address", bytes.NewReader(byt))
	return
}

// GetSslVpnIpAssignmentPool gets the ssl_vpn.ip_assignment_pool value from the UTM
func GetSslVpnIpAssignmentPool(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.ip_assignment_pool")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnIpAssignmentPool PUTs the ssl_vpn.ip_assignment_pool value to the UTM
func UpdateSslVpnIpAssignmentPool(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.ip_assignment_pool", bytes.NewReader(byt))
	return
}

// GetSslVpnPort gets the ssl_vpn.port value from the UTM
func GetSslVpnPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnPort PUTs the ssl_vpn.port value to the UTM
func UpdateSslVpnPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.port", bytes.NewReader(byt))
	return
}

// GetSslVpnProtocol gets the ssl_vpn.protocol value from the UTM
func GetSslVpnProtocol(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.protocol")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnProtocol PUTs the ssl_vpn.protocol value to the UTM
func UpdateSslVpnProtocol(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.protocol", bytes.NewReader(byt))
	return
}

// GetSslVpnUserAuthOptional gets the ssl_vpn.user_auth_optional value from the UTM
func GetSslVpnUserAuthOptional(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/ssl_vpn.user_auth_optional")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSslVpnUserAuthOptional PUTs the ssl_vpn.user_auth_optional value to the UTM
func UpdateSslVpnUserAuthOptional(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/ssl_vpn.user_auth_optional", bytes.NewReader(byt))
	return
}

// GetSupportAccessAccessId gets the support_access.access_id value from the UTM
func GetSupportAccessAccessId(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/support_access.access_id")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSupportAccessAccessId PUTs the support_access.access_id value to the UTM
func UpdateSupportAccessAccessId(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/support_access.access_id", bytes.NewReader(byt))
	return
}

// GetSupportAccessApuServer gets the support_access.apu_server value from the UTM
func GetSupportAccessApuServer(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/support_access.apu_server")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSupportAccessApuServer PUTs the support_access.apu_server value to the UTM
func UpdateSupportAccessApuServer(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/support_access.apu_server", bytes.NewReader(byt))
	return
}

// GetSupportAccessLifetimeDuration gets the support_access.lifetime_duration value from the UTM
func GetSupportAccessLifetimeDuration(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/support_access.lifetime_duration")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSupportAccessLifetimeDuration PUTs the support_access.lifetime_duration value to the UTM
func UpdateSupportAccessLifetimeDuration(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/support_access.lifetime_duration", bytes.NewReader(byt))
	return
}

// GetSupportAccessLifetimeEnd gets the support_access.lifetime_end value from the UTM
func GetSupportAccessLifetimeEnd(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/support_access.lifetime_end")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSupportAccessLifetimeEnd PUTs the support_access.lifetime_end value to the UTM
func UpdateSupportAccessLifetimeEnd(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/support_access.lifetime_end", bytes.NewReader(byt))
	return
}

// GetSupportAccessSshKeys gets the support_access.ssh_keys value from the UTM
func GetSupportAccessSshKeys(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/support_access.ssh_keys")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSupportAccessSshKeys PUTs the support_access.ssh_keys value to the UTM
func UpdateSupportAccessSshKeys(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/support_access.ssh_keys", bytes.NewReader(byt))
	return
}

// GetSupportAccessStatus gets the support_access.status value from the UTM
func GetSupportAccessStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/support_access.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateSupportAccessStatus PUTs the support_access.status value to the UTM
func UpdateSupportAccessStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/support_access.status", bytes.NewReader(byt))
	return
}

// GetU2DcacheAllowedNetworks gets the u2dcache.allowed_networks value from the UTM
func GetU2DcacheAllowedNetworks(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/u2dcache.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateU2DcacheAllowedNetworks PUTs the u2dcache.allowed_networks value to the UTM
func UpdateU2DcacheAllowedNetworks(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/u2dcache.allowed_networks", bytes.NewReader(byt))
	return
}

// GetU2DcachePort gets the u2dcache.port value from the UTM
func GetU2DcachePort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/u2dcache.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateU2DcachePort PUTs the u2dcache.port value to the UTM
func UpdateU2DcachePort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/u2dcache.port", bytes.NewReader(byt))
	return
}

// GetU2DcacheStatus gets the u2dcache.status value from the UTM
func GetU2DcacheStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/u2dcache.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateU2DcacheStatus PUTs the u2dcache.status value to the UTM
func UpdateU2DcacheStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/u2dcache.status", bytes.NewReader(byt))
	return
}

// GetUp2DateCacheHost gets the up2date.cache_host value from the UTM
func GetUp2DateCacheHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/up2date.cache_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateCacheHost PUTs the up2date.cache_host value to the UTM
func UpdateUp2DateCacheHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.cache_host", bytes.NewReader(byt))
	return
}

// GetUp2DateCachePort gets the up2date.cache_port value from the UTM
func GetUp2DateCachePort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/up2date.cache_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateCachePort PUTs the up2date.cache_port value to the UTM
func UpdateUp2DateCachePort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.cache_port", bytes.NewReader(byt))
	return
}

// GetUp2DateCacheStatus gets the up2date.cache_status value from the UTM
func GetUp2DateCacheStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.cache_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateCacheStatus PUTs the up2date.cache_status value to the UTM
func UpdateUp2DateCacheStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.cache_status", bytes.NewReader(byt))
	return
}

// GetUp2DateCacheUseAcc gets the up2date.cache_use_acc value from the UTM
func GetUp2DateCacheUseAcc(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.cache_use_acc")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateCacheUseAcc PUTs the up2date.cache_use_acc value to the UTM
func UpdateUp2DateCacheUseAcc(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.cache_use_acc", bytes.NewReader(byt))
	return
}

// GetUp2DateParentProxyAuthPass gets the up2date.parent_proxy_auth_pass value from the UTM
func GetUp2DateParentProxyAuthPass(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/up2date.parent_proxy_auth_pass")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateParentProxyAuthPass PUTs the up2date.parent_proxy_auth_pass value to the UTM
func UpdateUp2DateParentProxyAuthPass(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.parent_proxy_auth_pass", bytes.NewReader(byt))
	return
}

// GetUp2DateParentProxyAuthStatus gets the up2date.parent_proxy_auth_status value from the UTM
func GetUp2DateParentProxyAuthStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.parent_proxy_auth_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateParentProxyAuthStatus PUTs the up2date.parent_proxy_auth_status value to the UTM
func UpdateUp2DateParentProxyAuthStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.parent_proxy_auth_status", bytes.NewReader(byt))
	return
}

// GetUp2DateParentProxyAuthUser gets the up2date.parent_proxy_auth_user value from the UTM
func GetUp2DateParentProxyAuthUser(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/up2date.parent_proxy_auth_user")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateParentProxyAuthUser PUTs the up2date.parent_proxy_auth_user value to the UTM
func UpdateUp2DateParentProxyAuthUser(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.parent_proxy_auth_user", bytes.NewReader(byt))
	return
}

// GetUp2DateParentProxyHost gets the up2date.parent_proxy_host value from the UTM
func GetUp2DateParentProxyHost(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/up2date.parent_proxy_host")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateParentProxyHost PUTs the up2date.parent_proxy_host value to the UTM
func UpdateUp2DateParentProxyHost(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.parent_proxy_host", bytes.NewReader(byt))
	return
}

// GetUp2DateParentProxyPort gets the up2date.parent_proxy_port value from the UTM
func GetUp2DateParentProxyPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/up2date.parent_proxy_port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateParentProxyPort PUTs the up2date.parent_proxy_port value to the UTM
func UpdateUp2DateParentProxyPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.parent_proxy_port", bytes.NewReader(byt))
	return
}

// GetUp2DateParentProxyStatus gets the up2date.parent_proxy_status value from the UTM
func GetUp2DateParentProxyStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.parent_proxy_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateParentProxyStatus PUTs the up2date.parent_proxy_status value to the UTM
func UpdateUp2DateParentProxyStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.parent_proxy_status", bytes.NewReader(byt))
	return
}

// GetUp2DatePatternDownloadInterval gets the up2date.pattern_download_interval value from the UTM
func GetUp2DatePatternDownloadInterval(c sophos.ClientInterface) (val map[string]interface{}, err error) {
	res, err := c.Get("/api/nodes/up2date.pattern_download_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DatePatternDownloadInterval PUTs the up2date.pattern_download_interval value to the UTM
func UpdateUp2DatePatternDownloadInterval(c sophos.ClientInterface, val map[string]interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.pattern_download_interval", bytes.NewReader(byt))
	return
}

// GetUp2DatePatternDownloadStatus gets the up2date.pattern_download_status value from the UTM
func GetUp2DatePatternDownloadStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.pattern_download_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DatePatternDownloadStatus PUTs the up2date.pattern_download_status value to the UTM
func UpdateUp2DatePatternDownloadStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.pattern_download_status", bytes.NewReader(byt))
	return
}

// GetUp2DateScheduledUp2Date gets the up2date.scheduled_up2date value from the UTM
func GetUp2DateScheduledUp2Date(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/up2date.scheduled_up2date")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateScheduledUp2Date PUTs the up2date.scheduled_up2date value to the UTM
func UpdateUp2DateScheduledUp2Date(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.scheduled_up2date", bytes.NewReader(byt))
	return
}

// GetUp2DateServers gets the up2date.servers value from the UTM
func GetUp2DateServers(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/up2date.servers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateServers PUTs the up2date.servers value to the UTM
func UpdateUp2DateServers(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.servers", bytes.NewReader(byt))
	return
}

// GetUp2DateStatus gets the up2date.status value from the UTM
func GetUp2DateStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateStatus PUTs the up2date.status value to the UTM
func UpdateUp2DateStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.status", bytes.NewReader(byt))
	return
}

// GetUp2DateSystemAutoinstallTime gets the up2date.system_autoinstall_time value from the UTM
func GetUp2DateSystemAutoinstallTime(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/up2date.system_autoinstall_time")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateSystemAutoinstallTime PUTs the up2date.system_autoinstall_time value to the UTM
func UpdateUp2DateSystemAutoinstallTime(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.system_autoinstall_time", bytes.NewReader(byt))
	return
}

// GetUp2DateSystemDownloadInterval gets the up2date.system_download_interval value from the UTM
func GetUp2DateSystemDownloadInterval(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/up2date.system_download_interval")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateSystemDownloadInterval PUTs the up2date.system_download_interval value to the UTM
func UpdateUp2DateSystemDownloadInterval(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.system_download_interval", bytes.NewReader(byt))
	return
}

// GetUp2DateSystemDownloadStatus gets the up2date.system_download_status value from the UTM
func GetUp2DateSystemDownloadStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/up2date.system_download_status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUp2DateSystemDownloadStatus PUTs the up2date.system_download_status value to the UTM
func UpdateUp2DateSystemDownloadStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/up2date.system_download_status", bytes.NewReader(byt))
	return
}

// GetUplinkActions gets the uplink.actions value from the UTM
func GetUplinkActions(c sophos.ClientInterface) (val []interface{}, err error) {
	res, err := c.Get("/api/nodes/uplink.actions")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkActions PUTs the uplink.actions value to the UTM
func UpdateUplinkActions(c sophos.ClientInterface, val []interface{}) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.actions", bytes.NewReader(byt))
	return
}

// GetUplinkActive gets the uplink.active value from the UTM
func GetUplinkActive(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/uplink.active")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkActive PUTs the uplink.active value to the UTM
func UpdateUplinkActive(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.active", bytes.NewReader(byt))
	return
}

// GetUplinkCondition gets the uplink.condition value from the UTM
func GetUplinkCondition(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/uplink.condition")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkCondition PUTs the uplink.condition value to the UTM
func UpdateUplinkCondition(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.condition", bytes.NewReader(byt))
	return
}

// GetUplinkInterfaces gets the uplink.interfaces value from the UTM
func GetUplinkInterfaces(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/uplink.interfaces")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkInterfaces PUTs the uplink.interfaces value to the UTM
func UpdateUplinkInterfaces(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.interfaces", bytes.NewReader(byt))
	return
}

// GetUplinkMonitoring gets the uplink.monitoring value from the UTM
func GetUplinkMonitoring(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/uplink.monitoring")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkMonitoring PUTs the uplink.monitoring value to the UTM
func UpdateUplinkMonitoring(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.monitoring", bytes.NewReader(byt))
	return
}

// GetUplinkPassive gets the uplink.passive value from the UTM
func GetUplinkPassive(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/uplink.passive")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkPassive PUTs the uplink.passive value to the UTM
func UpdateUplinkPassive(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.passive", bytes.NewReader(byt))
	return
}

// GetUplinkPrimary gets the uplink.primary value from the UTM
func GetUplinkPrimary(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/uplink.primary")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkPrimary PUTs the uplink.primary value to the UTM
func UpdateUplinkPrimary(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.primary", bytes.NewReader(byt))
	return
}

// GetUplinkRules gets the uplink.rules value from the UTM
func GetUplinkRules(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/uplink.rules")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkRules PUTs the uplink.rules value to the UTM
func UpdateUplinkRules(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.rules", bytes.NewReader(byt))
	return
}

// GetUplinkScheduler gets the uplink.scheduler value from the UTM
func GetUplinkScheduler(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/uplink.scheduler")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkScheduler PUTs the uplink.scheduler value to the UTM
func UpdateUplinkScheduler(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.scheduler", bytes.NewReader(byt))
	return
}

// GetUplinkStatus gets the uplink.status value from the UTM
func GetUplinkStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/uplink.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateUplinkStatus PUTs the uplink.status value to the UTM
func UpdateUplinkStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/uplink.status", bytes.NewReader(byt))
	return
}

// GetWebadminAllowedNetworks gets the webadmin.allowed_networks value from the UTM
func GetWebadminAllowedNetworks(c sophos.ClientInterface) (val []string, err error) {
	res, err := c.Get("/api/nodes/webadmin.allowed_networks")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminAllowedNetworks PUTs the webadmin.allowed_networks value to the UTM
func UpdateWebadminAllowedNetworks(c sophos.ClientInterface, val []string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.allowed_networks", bytes.NewReader(byt))
	return
}

// GetWebadminCa gets the webadmin.ca value from the UTM
func GetWebadminCa(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.ca")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminCa PUTs the webadmin.ca value to the UTM
func UpdateWebadminCa(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.ca", bytes.NewReader(byt))
	return
}

// GetWebadminCert gets the webadmin.cert value from the UTM
func GetWebadminCert(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.cert")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminCert PUTs the webadmin.cert value to the UTM
func UpdateWebadminCert(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.cert", bytes.NewReader(byt))
	return
}

// GetWebadminDashboardRefresh gets the webadmin.dashboard_refresh value from the UTM
func GetWebadminDashboardRefresh(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.dashboard_refresh")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminDashboardRefresh PUTs the webadmin.dashboard_refresh value to the UTM
func UpdateWebadminDashboardRefresh(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.dashboard_refresh", bytes.NewReader(byt))
	return
}

// GetWebadminLanguage gets the webadmin.language value from the UTM
func GetWebadminLanguage(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.language")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminLanguage PUTs the webadmin.language value to the UTM
func UpdateWebadminLanguage(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.language", bytes.NewReader(byt))
	return
}

// GetWebadminLogAdminConnect gets the webadmin.log_admin_connect value from the UTM
func GetWebadminLogAdminConnect(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/webadmin.log_admin_connect")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminLogAdminConnect PUTs the webadmin.log_admin_connect value to the UTM
func UpdateWebadminLogAdminConnect(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.log_admin_connect", bytes.NewReader(byt))
	return
}

// GetWebadminOfferWizard gets the webadmin.offer_wizard value from the UTM
func GetWebadminOfferWizard(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/webadmin.offer_wizard")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminOfferWizard PUTs the webadmin.offer_wizard value to the UTM
func UpdateWebadminOfferWizard(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.offer_wizard", bytes.NewReader(byt))
	return
}

// GetWebadminPort gets the webadmin.port value from the UTM
func GetWebadminPort(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/webadmin.port")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminPort PUTs the webadmin.port value to the UTM
func UpdateWebadminPort(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.port", bytes.NewReader(byt))
	return
}

// GetWebadminRestApi gets the webadmin.rest_api value from the UTM
func GetWebadminRestApi(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/webadmin.rest_api")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminRestApi PUTs the webadmin.rest_api value to the UTM
func UpdateWebadminRestApi(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.rest_api", bytes.NewReader(byt))
	return
}

// GetWebadminTermsOfUseStatus gets the webadmin.terms_of_use.status value from the UTM
func GetWebadminTermsOfUseStatus(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/webadmin.terms_of_use.status")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminTermsOfUseStatus PUTs the webadmin.terms_of_use.status value to the UTM
func UpdateWebadminTermsOfUseStatus(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.terms_of_use.status", bytes.NewReader(byt))
	return
}

// GetWebadminTermsOfUseText gets the webadmin.terms_of_use.text value from the UTM
func GetWebadminTermsOfUseText(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.terms_of_use.text")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminTermsOfUseText PUTs the webadmin.terms_of_use.text value to the UTM
func UpdateWebadminTermsOfUseText(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.terms_of_use.text", bytes.NewReader(byt))
	return
}

// GetWebadminTimeout gets the webadmin.timeout value from the UTM
func GetWebadminTimeout(c sophos.ClientInterface) (val int64, err error) {
	res, err := c.Get("/api/nodes/webadmin.timeout")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminTimeout PUTs the webadmin.timeout value to the UTM
func UpdateWebadminTimeout(c sophos.ClientInterface, val int64) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.timeout", bytes.NewReader(byt))
	return
}

// GetWebadminTimeoutOnDashboard gets the webadmin.timeout_on_dashboard value from the UTM
func GetWebadminTimeoutOnDashboard(c sophos.ClientInterface) (val bool, err error) {
	res, err := c.Get("/api/nodes/webadmin.timeout_on_dashboard")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminTimeoutOnDashboard PUTs the webadmin.timeout_on_dashboard value to the UTM
func UpdateWebadminTimeoutOnDashboard(c sophos.ClientInterface, val bool) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.timeout_on_dashboard", bytes.NewReader(byt))
	return
}

// GetWebadminTlsCiphers gets the webadmin.tls_ciphers value from the UTM
func GetWebadminTlsCiphers(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.tls_ciphers")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminTlsCiphers PUTs the webadmin.tls_ciphers value to the UTM
func UpdateWebadminTlsCiphers(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.tls_ciphers", bytes.NewReader(byt))
	return
}

// GetWebadminTlsProtocols gets the webadmin.tls_protocols value from the UTM
func GetWebadminTlsProtocols(c sophos.ClientInterface) (val string, err error) {
	res, err := c.Get("/api/nodes/webadmin.tls_protocols")
	if err != nil {
		return val, err
	}
	err = res.MarshalTo(&val)
	return
}

// UpdateWebadminTlsProtocols PUTs the webadmin.tls_protocols value to the UTM
func UpdateWebadminTlsProtocols(c sophos.ClientInterface, val string) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put("/api/nodes/webadmin.tls_protocols", bytes.NewReader(byt))
	return
}
