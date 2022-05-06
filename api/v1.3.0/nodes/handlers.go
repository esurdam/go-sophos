package nodes

import (
	"bytes"
	"encoding/json"

	"github.com/esurdam/go-sophos"
)

func get(c sophos.ClientInterface, path string, val interface{}, options ...sophos.Option) (err error) {
	res, err := c.Get(path, options...)
	if err != nil {
		return err
	}
	err = res.MarshalTo(val)
	return
}

func put(c sophos.ClientInterface, path string, val interface{}, options ...sophos.Option) (err error) {
	byt, _ := json.Marshal(val)
	_, err = c.Put(path, bytes.NewReader(byt), options...)
	return
}

// GetAccServer1AuthSecret gets the acc.server1.auth.secret value from the UTM
func GetAccServer1AuthSecret(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/acc.server1.auth.secret", &val, options...)
	return
}

// UpdateAccServer1AuthSecret PUTs the acc.server1.auth.secret value to the UTM
func UpdateAccServer1AuthSecret(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server1.auth.secret", val, options...)
}

// GetAccServer1AuthServerCert gets the acc.server1.auth.server_cert value from the UTM
func GetAccServer1AuthServerCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/acc.server1.auth.server_cert", &val, options...)
	return
}

// UpdateAccServer1AuthServerCert PUTs the acc.server1.auth.server_cert value to the UTM
func UpdateAccServer1AuthServerCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server1.auth.server_cert", val, options...)
}

// GetAccServer1AuthStatus gets the acc.server1.auth.status value from the UTM
func GetAccServer1AuthStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/acc.server1.auth.status", &val, options...)
	return
}

// UpdateAccServer1AuthStatus PUTs the acc.server1.auth.status value to the UTM
func UpdateAccServer1AuthStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server1.auth.status", val, options...)
}

// GetAccServer1Port gets the acc.server1.port value from the UTM
func GetAccServer1Port(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/acc.server1.port", &val, options...)
	return
}

// UpdateAccServer1Port PUTs the acc.server1.port value to the UTM
func UpdateAccServer1Port(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server1.port", val, options...)
}

// GetAccServer1Roles gets the acc.server1.roles value from the UTM
func GetAccServer1Roles(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/acc.server1.roles", &val, options...)
	return
}

// UpdateAccServer1Roles PUTs the acc.server1.roles value to the UTM
func UpdateAccServer1Roles(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server1.roles", val, options...)
}

// GetAccServer1Server gets the acc.server1.server value from the UTM
func GetAccServer1Server(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/acc.server1.server", &val, options...)
	return
}

// UpdateAccServer1Server PUTs the acc.server1.server value to the UTM
func UpdateAccServer1Server(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server1.server", val, options...)
}

// GetAccServer2AuthSecret gets the acc.server2.auth.secret value from the UTM
func GetAccServer2AuthSecret(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/acc.server2.auth.secret", &val, options...)
	return
}

// UpdateAccServer2AuthSecret PUTs the acc.server2.auth.secret value to the UTM
func UpdateAccServer2AuthSecret(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server2.auth.secret", val, options...)
}

// GetAccServer2AuthServerCert gets the acc.server2.auth.server_cert value from the UTM
func GetAccServer2AuthServerCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/acc.server2.auth.server_cert", &val, options...)
	return
}

// UpdateAccServer2AuthServerCert PUTs the acc.server2.auth.server_cert value to the UTM
func UpdateAccServer2AuthServerCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server2.auth.server_cert", val, options...)
}

// GetAccServer2AuthStatus gets the acc.server2.auth.status value from the UTM
func GetAccServer2AuthStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/acc.server2.auth.status", &val, options...)
	return
}

// UpdateAccServer2AuthStatus PUTs the acc.server2.auth.status value to the UTM
func UpdateAccServer2AuthStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server2.auth.status", val, options...)
}

// GetAccServer2Port gets the acc.server2.port value from the UTM
func GetAccServer2Port(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/acc.server2.port", &val, options...)
	return
}

// UpdateAccServer2Port PUTs the acc.server2.port value to the UTM
func UpdateAccServer2Port(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server2.port", val, options...)
}

// GetAccServer2Roles gets the acc.server2.roles value from the UTM
func GetAccServer2Roles(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/acc.server2.roles", &val, options...)
	return
}

// UpdateAccServer2Roles PUTs the acc.server2.roles value to the UTM
func UpdateAccServer2Roles(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server2.roles", val, options...)
}

// GetAccServer2Server gets the acc.server2.server value from the UTM
func GetAccServer2Server(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/acc.server2.server", &val, options...)
	return
}

// UpdateAccServer2Server PUTs the acc.server2.server value to the UTM
func UpdateAccServer2Server(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server2.server", val, options...)
}

// GetAccServer2Status gets the acc.server2.status value from the UTM
func GetAccServer2Status(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/acc.server2.status", &val, options...)
	return
}

// UpdateAccServer2Status PUTs the acc.server2.status value to the UTM
func UpdateAccServer2Status(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.server2.status", val, options...)
}

// GetAccSsoAdminGroup gets the acc.sso_admin_group value from the UTM
func GetAccSsoAdminGroup(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/acc.sso_admin_group", &val, options...)
	return
}

// UpdateAccSsoAdminGroup PUTs the acc.sso_admin_group value to the UTM
func UpdateAccSsoAdminGroup(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.sso_admin_group", val, options...)
}

// GetAccSsoAuditorGroup gets the acc.sso_auditor_group value from the UTM
func GetAccSsoAuditorGroup(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/acc.sso_auditor_group", &val, options...)
	return
}

// UpdateAccSsoAuditorGroup PUTs the acc.sso_auditor_group value to the UTM
func UpdateAccSsoAuditorGroup(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.sso_auditor_group", val, options...)
}

// GetAccStatus gets the acc.status value from the UTM
func GetAccStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/acc.status", &val, options...)
	return
}

// UpdateAccStatus PUTs the acc.status value to the UTM
func UpdateAccStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/acc.status", val, options...)
}

// GetAccdAccessAllowedAdmins gets the accd.access.allowed_admins value from the UTM
func GetAccdAccessAllowedAdmins(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/accd.access.allowed_admins", &val, options...)
	return
}

// UpdateAccdAccessAllowedAdmins PUTs the accd.access.allowed_admins value to the UTM
func UpdateAccdAccessAllowedAdmins(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.access.allowed_admins", val, options...)
}

// GetAccdAccessAllowedNetworks gets the accd.access.allowed_networks value from the UTM
func GetAccdAccessAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/accd.access.allowed_networks", &val, options...)
	return
}

// UpdateAccdAccessAllowedNetworks PUTs the accd.access.allowed_networks value to the UTM
func UpdateAccdAccessAllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.access.allowed_networks", val, options...)
}

// GetAccdAccessAllowedUsers gets the accd.access.allowed_users value from the UTM
func GetAccdAccessAllowedUsers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/accd.access.allowed_users", &val, options...)
	return
}

// UpdateAccdAccessAllowedUsers PUTs the accd.access.allowed_users value to the UTM
func UpdateAccdAccessAllowedUsers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.access.allowed_users", val, options...)
}

// GetAccdAccessCert gets the accd.access.cert value from the UTM
func GetAccdAccessCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/accd.access.cert", &val, options...)
	return
}

// UpdateAccdAccessCert PUTs the accd.access.cert value to the UTM
func UpdateAccdAccessCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.access.cert", val, options...)
}

// GetAccdAccessPort gets the accd.access.port value from the UTM
func GetAccdAccessPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/accd.access.port", &val, options...)
	return
}

// UpdateAccdAccessPort PUTs the accd.access.port value to the UTM
func UpdateAccdAccessPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.access.port", val, options...)
}

// GetAccdDevicesAllowedNetworks gets the accd.devices.allowed_networks value from the UTM
func GetAccdDevicesAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/accd.devices.allowed_networks", &val, options...)
	return
}

// UpdateAccdDevicesAllowedNetworks PUTs the accd.devices.allowed_networks value to the UTM
func UpdateAccdDevicesAllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.devices.allowed_networks", val, options...)
}

// GetAccdDevicesAuthAuto gets the accd.devices.auth.auto value from the UTM
func GetAccdDevicesAuthAuto(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/accd.devices.auth.auto", &val, options...)
	return
}

// UpdateAccdDevicesAuthAuto PUTs the accd.devices.auth.auto value to the UTM
func UpdateAccdDevicesAuthAuto(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.devices.auth.auto", val, options...)
}

// GetAccdDevicesAuthSecret gets the accd.devices.auth.secret value from the UTM
func GetAccdDevicesAuthSecret(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/accd.devices.auth.secret", &val, options...)
	return
}

// UpdateAccdDevicesAuthSecret PUTs the accd.devices.auth.secret value to the UTM
func UpdateAccdDevicesAuthSecret(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.devices.auth.secret", val, options...)
}

// GetAccdDevicesAuthStatus gets the accd.devices.auth.status value from the UTM
func GetAccdDevicesAuthStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/accd.devices.auth.status", &val, options...)
	return
}

// UpdateAccdDevicesAuthStatus PUTs the accd.devices.auth.status value to the UTM
func UpdateAccdDevicesAuthStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.devices.auth.status", val, options...)
}

// GetAccdDevicesCert gets the accd.devices.cert value from the UTM
func GetAccdDevicesCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/accd.devices.cert", &val, options...)
	return
}

// UpdateAccdDevicesCert PUTs the accd.devices.cert value to the UTM
func UpdateAccdDevicesCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.devices.cert", val, options...)
}

// GetAccdDevicesPort gets the accd.devices.port value from the UTM
func GetAccdDevicesPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/accd.devices.port", &val, options...)
	return
}

// UpdateAccdDevicesPort PUTs the accd.devices.port value to the UTM
func UpdateAccdDevicesPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.devices.port", val, options...)
}

// GetAccdGeneralAllowedNetworks gets the accd.general.allowed_networks value from the UTM
func GetAccdGeneralAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/accd.general.allowed_networks", &val, options...)
	return
}

// UpdateAccdGeneralAllowedNetworks PUTs the accd.general.allowed_networks value to the UTM
func UpdateAccdGeneralAllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.general.allowed_networks", val, options...)
}

// GetAccdGeneralCert gets the accd.general.cert value from the UTM
func GetAccdGeneralCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/accd.general.cert", &val, options...)
	return
}

// UpdateAccdGeneralCert PUTs the accd.general.cert value to the UTM
func UpdateAccdGeneralCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.general.cert", val, options...)
}

// GetAccdGeneralLanguage gets the accd.general.language value from the UTM
func GetAccdGeneralLanguage(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/accd.general.language", &val, options...)
	return
}

// UpdateAccdGeneralLanguage PUTs the accd.general.language value to the UTM
func UpdateAccdGeneralLanguage(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.general.language", val, options...)
}

// GetAccdGeneralPort gets the accd.general.port value from the UTM
func GetAccdGeneralPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/accd.general.port", &val, options...)
	return
}

// UpdateAccdGeneralPort PUTs the accd.general.port value to the UTM
func UpdateAccdGeneralPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.general.port", val, options...)
}

// GetAccdGeneralTimeout gets the accd.general.timeout value from the UTM
func GetAccdGeneralTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/accd.general.timeout", &val, options...)
	return
}

// UpdateAccdGeneralTimeout PUTs the accd.general.timeout value to the UTM
func UpdateAccdGeneralTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accd.general.timeout", val, options...)
}

// GetAccountingIpfixConnections gets the accounting.ipfix.connections value from the UTM
func GetAccountingIpfixConnections(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/accounting.ipfix.connections", &val, options...)
	return
}

// UpdateAccountingIpfixConnections PUTs the accounting.ipfix.connections value to the UTM
func UpdateAccountingIpfixConnections(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accounting.ipfix.connections", val, options...)
}

// GetAccountingIpfixStatus gets the accounting.ipfix.status value from the UTM
func GetAccountingIpfixStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/accounting.ipfix.status", &val, options...)
	return
}

// UpdateAccountingIpfixStatus PUTs the accounting.ipfix.status value to the UTM
func UpdateAccountingIpfixStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/accounting.ipfix.status", val, options...)
}

// GetAfcControlledNetworks gets the afc.controlled_networks value from the UTM
func GetAfcControlledNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/afc.controlled_networks", &val, options...)
	return
}

// UpdateAfcControlledNetworks PUTs the afc.controlled_networks value to the UTM
func UpdateAfcControlledNetworks(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/afc.controlled_networks", val, options...)
}

// GetAfcHiddenSkip gets the afc.hidden_skip value from the UTM
func GetAfcHiddenSkip(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/afc.hidden_skip", &val, options...)
	return
}

// UpdateAfcHiddenSkip PUTs the afc.hidden_skip value to the UTM
func UpdateAfcHiddenSkip(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/afc.hidden_skip", val, options...)
}

// GetAfcHttpRedirectUrl gets the afc.http_redirect_url value from the UTM
func GetAfcHttpRedirectUrl(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/afc.http_redirect_url", &val, options...)
	return
}

// UpdateAfcHttpRedirectUrl PUTs the afc.http_redirect_url value to the UTM
func UpdateAfcHttpRedirectUrl(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/afc.http_redirect_url", val, options...)
}

// GetAfcLog gets the afc.log value from the UTM
func GetAfcLog(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/afc.log", &val, options...)
	return
}

// UpdateAfcLog PUTs the afc.log value to the UTM
func UpdateAfcLog(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/afc.log", val, options...)
}

// GetAfcNfqueueLength gets the afc.nfqueue_length value from the UTM
func GetAfcNfqueueLength(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/afc.nfqueue_length", &val, options...)
	return
}

// UpdateAfcNfqueueLength PUTs the afc.nfqueue_length value to the UTM
func UpdateAfcNfqueueLength(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/afc.nfqueue_length", val, options...)
}

// GetAfcNumQueues gets the afc.num_queues value from the UTM
func GetAfcNumQueues(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/afc.num_queues", &val, options...)
	return
}

// UpdateAfcNumQueues PUTs the afc.num_queues value to the UTM
func UpdateAfcNumQueues(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/afc.num_queues", val, options...)
}

// GetAfcRules gets the afc.rules value from the UTM
func GetAfcRules(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/afc.rules", &val, options...)
	return
}

// UpdateAfcRules PUTs the afc.rules value to the UTM
func UpdateAfcRules(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/afc.rules", val, options...)
}

// GetAfcStatus gets the afc.status value from the UTM
func GetAfcStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/afc.status", &val, options...)
	return
}

// UpdateAfcStatus PUTs the afc.status value to the UTM
func UpdateAfcStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/afc.status", val, options...)
}

// GetAfcSubappDetection gets the afc.subapp_detection value from the UTM
func GetAfcSubappDetection(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/afc.subapp_detection", &val, options...)
	return
}

// UpdateAfcSubappDetection PUTs the afc.subapp_detection value to the UTM
func UpdateAfcSubappDetection(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/afc.subapp_detection", val, options...)
}

// GetAfcSubmitUnknownTrafficData gets the afc.submit_unknown_traffic_data value from the UTM
func GetAfcSubmitUnknownTrafficData(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/afc.submit_unknown_traffic_data", &val, options...)
	return
}

// UpdateAfcSubmitUnknownTrafficData PUTs the afc.submit_unknown_traffic_data value to the UTM
func UpdateAfcSubmitUnknownTrafficData(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/afc.submit_unknown_traffic_data", val, options...)
}

// GetAfcTransparentSkip gets the afc.transparent_skip value from the UTM
func GetAfcTransparentSkip(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/afc.transparent_skip", &val, options...)
	return
}

// UpdateAfcTransparentSkip PUTs the afc.transparent_skip value to the UTM
func UpdateAfcTransparentSkip(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/afc.transparent_skip", val, options...)
}

// GetAmazonVpcAutoPfrule gets the amazon_vpc.auto_pfrule value from the UTM
func GetAmazonVpcAutoPfrule(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/amazon_vpc.auto_pfrule", &val, options...)
	return
}

// UpdateAmazonVpcAutoPfrule PUTs the amazon_vpc.auto_pfrule value to the UTM
func UpdateAmazonVpcAutoPfrule(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/amazon_vpc.auto_pfrule", val, options...)
}

// GetAmazonVpcConnections gets the amazon_vpc.connections value from the UTM
func GetAmazonVpcConnections(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/amazon_vpc.connections", &val, options...)
	return
}

// UpdateAmazonVpcConnections PUTs the amazon_vpc.connections value to the UTM
func UpdateAmazonVpcConnections(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/amazon_vpc.connections", val, options...)
}

// GetAmazonVpcNetworks gets the amazon_vpc.networks value from the UTM
func GetAmazonVpcNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/amazon_vpc.networks", &val, options...)
	return
}

// UpdateAmazonVpcNetworks PUTs the amazon_vpc.networks value to the UTM
func UpdateAmazonVpcNetworks(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/amazon_vpc.networks", val, options...)
}

// GetAmazonVpcStatus gets the amazon_vpc.status value from the UTM
func GetAmazonVpcStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/amazon_vpc.status", &val, options...)
	return
}

// UpdateAmazonVpcStatus PUTs the amazon_vpc.status value to the UTM
func UpdateAmazonVpcStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/amazon_vpc.status", val, options...)
}

// GetAptpDbPlugin gets the aptp.db_plugin value from the UTM
func GetAptpDbPlugin(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/aptp.db_plugin", &val, options...)
	return
}

// UpdateAptpDbPlugin PUTs the aptp.db_plugin value to the UTM
func UpdateAptpDbPlugin(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/aptp.db_plugin", val, options...)
}

// GetAptpNumServers gets the aptp.num_servers value from the UTM
func GetAptpNumServers(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/aptp.num_servers", &val, options...)
	return
}

// UpdateAptpNumServers PUTs the aptp.num_servers value to the UTM
func UpdateAptpNumServers(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/aptp.num_servers", val, options...)
}

// GetAptpNumThreads gets the aptp.num_threads value from the UTM
func GetAptpNumThreads(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/aptp.num_threads", &val, options...)
	return
}

// UpdateAptpNumThreads PUTs the aptp.num_threads value to the UTM
func UpdateAptpNumThreads(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/aptp.num_threads", val, options...)
}

// GetAptpPolicy gets the aptp.policy value from the UTM
func GetAptpPolicy(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/aptp.policy", &val, options...)
	return
}

// UpdateAptpPolicy PUTs the aptp.policy value to the UTM
func UpdateAptpPolicy(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/aptp.policy", val, options...)
}

// GetAptpPort gets the aptp.port value from the UTM
func GetAptpPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/aptp.port", &val, options...)
	return
}

// UpdateAptpPort PUTs the aptp.port value to the UTM
func UpdateAptpPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/aptp.port", val, options...)
}

// GetAptpRuleModifiers gets the aptp.rule_modifiers value from the UTM
func GetAptpRuleModifiers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/aptp.rule_modifiers", &val, options...)
	return
}

// UpdateAptpRuleModifiers PUTs the aptp.rule_modifiers value to the UTM
func UpdateAptpRuleModifiers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/aptp.rule_modifiers", val, options...)
}

// GetAptpStatus gets the aptp.status value from the UTM
func GetAptpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/aptp.status", &val, options...)
	return
}

// UpdateAptpStatus PUTs the aptp.status value to the UTM
func UpdateAptpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/aptp.status", val, options...)
}

// GetAptpTransparentSkip gets the aptp.transparent_skip value from the UTM
func GetAptpTransparentSkip(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/aptp.transparent_skip", &val, options...)
	return
}

// UpdateAptpTransparentSkip PUTs the aptp.transparent_skip value to the UTM
func UpdateAptpTransparentSkip(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/aptp.transparent_skip", val, options...)
}

// GetArmLicensedIp gets the arm.licensed_ip value from the UTM
func GetArmLicensedIp(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/arm.licensed_ip", &val, options...)
	return
}

// UpdateArmLicensedIp PUTs the arm.licensed_ip value to the UTM
func UpdateArmLicensedIp(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/arm.licensed_ip", val, options...)
}

// GetArmRemoteHost gets the arm.remote.host value from the UTM
func GetArmRemoteHost(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/arm.remote.host", &val, options...)
	return
}

// UpdateArmRemoteHost PUTs the arm.remote.host value to the UTM
func UpdateArmRemoteHost(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/arm.remote.host", val, options...)
}

// GetArmRemoteMethod gets the arm.remote.method value from the UTM
func GetArmRemoteMethod(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/arm.remote.method", &val, options...)
	return
}

// UpdateArmRemoteMethod PUTs the arm.remote.method value to the UTM
func UpdateArmRemoteMethod(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/arm.remote.method", val, options...)
}

// GetArmRemoteSmbPassword gets the arm.remote.smb_password value from the UTM
func GetArmRemoteSmbPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/arm.remote.smb_password", &val, options...)
	return
}

// UpdateArmRemoteSmbPassword PUTs the arm.remote.smb_password value to the UTM
func UpdateArmRemoteSmbPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/arm.remote.smb_password", val, options...)
}

// GetArmRemoteSmbShare gets the arm.remote.smb_share value from the UTM
func GetArmRemoteSmbShare(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/arm.remote.smb_share", &val, options...)
	return
}

// UpdateArmRemoteSmbShare PUTs the arm.remote.smb_share value to the UTM
func UpdateArmRemoteSmbShare(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/arm.remote.smb_share", val, options...)
}

// GetArmRemoteSmbUser gets the arm.remote.smb_user value from the UTM
func GetArmRemoteSmbUser(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/arm.remote.smb_user", &val, options...)
	return
}

// UpdateArmRemoteSmbUser PUTs the arm.remote.smb_user value to the UTM
func UpdateArmRemoteSmbUser(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/arm.remote.smb_user", val, options...)
}

// GetArmRemoteStatus gets the arm.remote.status value from the UTM
func GetArmRemoteStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/arm.remote.status", &val, options...)
	return
}

// UpdateArmRemoteStatus PUTs the arm.remote.status value to the UTM
func UpdateArmRemoteStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/arm.remote.status", val, options...)
}

// GetArmRemoteSyslogService gets the arm.remote.syslog_service value from the UTM
func GetArmRemoteSyslogService(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/arm.remote.syslog_service", &val, options...)
	return
}

// UpdateArmRemoteSyslogService PUTs the arm.remote.syslog_service value to the UTM
func UpdateArmRemoteSyslogService(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/arm.remote.syslog_service", val, options...)
}

// GetArmStatus gets the arm.status value from the UTM
func GetArmStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/arm.status", &val, options...)
	return
}

// UpdateArmStatus PUTs the arm.status value to the UTM
func UpdateArmStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/arm.status", val, options...)
}

// GetAuthAdSsoForceUtf8Sync gets the auth.ad_sso.force_utf8_sync value from the UTM
func GetAuthAdSsoForceUtf8Sync(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.force_utf8_sync", &val, options...)
	return
}

// UpdateAuthAdSsoForceUtf8Sync PUTs the auth.ad_sso.force_utf8_sync value to the UTM
func UpdateAuthAdSsoForceUtf8Sync(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.force_utf8_sync", val, options...)
}

// GetAuthAdSsoJoinresult gets the auth.ad_sso.joinresult value from the UTM
func GetAuthAdSsoJoinresult(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.joinresult", &val, options...)
	return
}

// UpdateAuthAdSsoJoinresult PUTs the auth.ad_sso.joinresult value to the UTM
func UpdateAuthAdSsoJoinresult(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.joinresult", val, options...)
}

// GetAuthAdSsoLoadbalancerFqdn gets the auth.ad_sso.loadbalancer_fqdn value from the UTM
func GetAuthAdSsoLoadbalancerFqdn(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.loadbalancer_fqdn", &val, options...)
	return
}

// UpdateAuthAdSsoLoadbalancerFqdn PUTs the auth.ad_sso.loadbalancer_fqdn value to the UTM
func UpdateAuthAdSsoLoadbalancerFqdn(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.loadbalancer_fqdn", val, options...)
}

// GetAuthAdSsoNtlmv2Auth gets the auth.ad_sso.ntlmv2_auth value from the UTM
func GetAuthAdSsoNtlmv2Auth(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.ntlmv2_auth", &val, options...)
	return
}

// UpdateAuthAdSsoNtlmv2Auth PUTs the auth.ad_sso.ntlmv2_auth value to the UTM
func UpdateAuthAdSsoNtlmv2Auth(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.ntlmv2_auth", val, options...)
}

// GetAuthAdSsoSecrets gets the auth.ad_sso.secrets value from the UTM
func GetAuthAdSsoSecrets(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.secrets", &val, options...)
	return
}

// UpdateAuthAdSsoSecrets PUTs the auth.ad_sso.secrets value to the UTM
func UpdateAuthAdSsoSecrets(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.secrets", val, options...)
}

// GetAuthAdSsoSmbconf gets the auth.ad_sso.smbconf value from the UTM
func GetAuthAdSsoSmbconf(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.smbconf", &val, options...)
	return
}

// UpdateAuthAdSsoSmbconf PUTs the auth.ad_sso.smbconf value to the UTM
func UpdateAuthAdSsoSmbconf(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.smbconf", val, options...)
}

// GetAuthAdSsoSsoDomain gets the auth.ad_sso.sso_domain value from the UTM
func GetAuthAdSsoSsoDomain(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.sso_domain", &val, options...)
	return
}

// UpdateAuthAdSsoSsoDomain PUTs the auth.ad_sso.sso_domain value to the UTM
func UpdateAuthAdSsoSsoDomain(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.sso_domain", val, options...)
}

// GetAuthAdSsoSsoNetbiosDomain gets the auth.ad_sso.sso_netbios_domain value from the UTM
func GetAuthAdSsoSsoNetbiosDomain(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.sso_netbios_domain", &val, options...)
	return
}

// UpdateAuthAdSsoSsoNetbiosDomain PUTs the auth.ad_sso.sso_netbios_domain value to the UTM
func UpdateAuthAdSsoSsoNetbiosDomain(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.sso_netbios_domain", val, options...)
}

// GetAuthAdSsoSsoNetbiosHost gets the auth.ad_sso.sso_netbios_host value from the UTM
func GetAuthAdSsoSsoNetbiosHost(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.sso_netbios_host", &val, options...)
	return
}

// UpdateAuthAdSsoSsoNetbiosHost PUTs the auth.ad_sso.sso_netbios_host value to the UTM
func UpdateAuthAdSsoSsoNetbiosHost(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.sso_netbios_host", val, options...)
}

// GetAuthAdSsoSsoPassword gets the auth.ad_sso.sso_password value from the UTM
func GetAuthAdSsoSsoPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.sso_password", &val, options...)
	return
}

// UpdateAuthAdSsoSsoPassword PUTs the auth.ad_sso.sso_password value to the UTM
func UpdateAuthAdSsoSsoPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.sso_password", val, options...)
}

// GetAuthAdSsoSsoServer gets the auth.ad_sso.sso_server value from the UTM
func GetAuthAdSsoSsoServer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.sso_server", &val, options...)
	return
}

// UpdateAuthAdSsoSsoServer PUTs the auth.ad_sso.sso_server value to the UTM
func UpdateAuthAdSsoSsoServer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.sso_server", val, options...)
}

// GetAuthAdSsoSsoStatus gets the auth.ad_sso.sso_status value from the UTM
func GetAuthAdSsoSsoStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.sso_status", &val, options...)
	return
}

// UpdateAuthAdSsoSsoStatus PUTs the auth.ad_sso.sso_status value to the UTM
func UpdateAuthAdSsoSsoStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.sso_status", val, options...)
}

// GetAuthAdSsoSsoSync gets the auth.ad_sso.sso_sync value from the UTM
func GetAuthAdSsoSsoSync(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.sso_sync", &val, options...)
	return
}

// UpdateAuthAdSsoSsoSync PUTs the auth.ad_sso.sso_sync value to the UTM
func UpdateAuthAdSsoSsoSync(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.sso_sync", val, options...)
}

// GetAuthAdSsoSsoUsername gets the auth.ad_sso.sso_username value from the UTM
func GetAuthAdSsoSsoUsername(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.ad_sso.sso_username", &val, options...)
	return
}

// UpdateAuthAdSsoSsoUsername PUTs the auth.ad_sso.sso_username value to the UTM
func UpdateAuthAdSsoSsoUsername(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.ad_sso.sso_username", val, options...)
}

// GetAuthApiTokens gets the auth.api_tokens value from the UTM
func GetAuthApiTokens(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/auth.api_tokens", &val, options...)
	return
}

// UpdateAuthApiTokens PUTs the auth.api_tokens value to the UTM
func UpdateAuthApiTokens(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.api_tokens", val, options...)
}

// GetAuthAutoAddToFacility gets the auth.auto_add_to_facility value from the UTM
func GetAuthAutoAddToFacility(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/auth.auto_add_to_facility", &val, options...)
	return
}

// UpdateAuthAutoAddToFacility PUTs the auth.auto_add_to_facility value to the UTM
func UpdateAuthAutoAddToFacility(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.auto_add_to_facility", val, options...)
}

// GetAuthAutoAddUsers gets the auth.auto_add_users value from the UTM
func GetAuthAutoAddUsers(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.auto_add_users", &val, options...)
	return
}

// UpdateAuthAutoAddUsers PUTs the auth.auto_add_users value to the UTM
func UpdateAuthAutoAddUsers(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.auto_add_users", val, options...)
}

// GetAuthBlockAttempts gets the auth.block.attempts value from the UTM
func GetAuthBlockAttempts(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/auth.block.attempts", &val, options...)
	return
}

// UpdateAuthBlockAttempts PUTs the auth.block.attempts value to the UTM
func UpdateAuthBlockAttempts(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.block.attempts", val, options...)
}

// GetAuthBlockFacilities gets the auth.block.facilities value from the UTM
func GetAuthBlockFacilities(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/auth.block.facilities", &val, options...)
	return
}

// UpdateAuthBlockFacilities PUTs the auth.block.facilities value to the UTM
func UpdateAuthBlockFacilities(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.block.facilities", val, options...)
}

// GetAuthBlockLockout gets the auth.block.lockout value from the UTM
func GetAuthBlockLockout(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.block.lockout", &val, options...)
	return
}

// UpdateAuthBlockLockout PUTs the auth.block.lockout value to the UTM
func UpdateAuthBlockLockout(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.block.lockout", val, options...)
}

// GetAuthBlockNever gets the auth.block.never value from the UTM
func GetAuthBlockNever(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/auth.block.never", &val, options...)
	return
}

// UpdateAuthBlockNever PUTs the auth.block.never value to the UTM
func UpdateAuthBlockNever(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.block.never", val, options...)
}

// GetAuthBlockSeconds gets the auth.block.seconds value from the UTM
func GetAuthBlockSeconds(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/auth.block.seconds", &val, options...)
	return
}

// UpdateAuthBlockSeconds PUTs the auth.block.seconds value to the UTM
func UpdateAuthBlockSeconds(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.block.seconds", val, options...)
}

// GetAuthCacheLifetime gets the auth.cache_lifetime value from the UTM
func GetAuthCacheLifetime(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/auth.cache_lifetime", &val, options...)
	return
}

// UpdateAuthCacheLifetime PUTs the auth.cache_lifetime value to the UTM
func UpdateAuthCacheLifetime(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.cache_lifetime", val, options...)
}

// GetAuthDelayedIpsetExpansion gets the auth.delayed_ipset_expansion value from the UTM
func GetAuthDelayedIpsetExpansion(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/auth.delayed_ipset_expansion", &val, options...)
	return
}

// UpdateAuthDelayedIpsetExpansion PUTs the auth.delayed_ipset_expansion value to the UTM
func UpdateAuthDelayedIpsetExpansion(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.delayed_ipset_expansion", val, options...)
}

// GetAuthEdirSsoEmConflict gets the auth.edir_sso.em_conflict value from the UTM
func GetAuthEdirSsoEmConflict(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.edir_sso.em_conflict", &val, options...)
	return
}

// UpdateAuthEdirSsoEmConflict PUTs the auth.edir_sso.em_conflict value to the UTM
func UpdateAuthEdirSsoEmConflict(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.edir_sso.em_conflict", val, options...)
}

// GetAuthEdirSsoEmSocketTimeout gets the auth.edir_sso.em_socket_timeout value from the UTM
func GetAuthEdirSsoEmSocketTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/auth.edir_sso.em_socket_timeout", &val, options...)
	return
}

// UpdateAuthEdirSsoEmSocketTimeout PUTs the auth.edir_sso.em_socket_timeout value to the UTM
func UpdateAuthEdirSsoEmSocketTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.edir_sso.em_socket_timeout", val, options...)
}

// GetAuthEdirSsoEmVerifyLogout gets the auth.edir_sso.em_verify_logout value from the UTM
func GetAuthEdirSsoEmVerifyLogout(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.edir_sso.em_verify_logout", &val, options...)
	return
}

// UpdateAuthEdirSsoEmVerifyLogout PUTs the auth.edir_sso.em_verify_logout value to the UTM
func UpdateAuthEdirSsoEmVerifyLogout(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.edir_sso.em_verify_logout", val, options...)
}

// GetAuthEdirSsoSsoAuaSearchIp gets the auth.edir_sso.sso_aua_search_ip value from the UTM
func GetAuthEdirSsoSsoAuaSearchIp(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.edir_sso.sso_aua_search_ip", &val, options...)
	return
}

// UpdateAuthEdirSsoSsoAuaSearchIp PUTs the auth.edir_sso.sso_aua_search_ip value to the UTM
func UpdateAuthEdirSsoSsoAuaSearchIp(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.edir_sso.sso_aua_search_ip", val, options...)
}

// GetAuthEdirSsoSsoMode gets the auth.edir_sso.sso_mode value from the UTM
func GetAuthEdirSsoSsoMode(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.edir_sso.sso_mode", &val, options...)
	return
}

// UpdateAuthEdirSsoSsoMode PUTs the auth.edir_sso.sso_mode value to the UTM
func UpdateAuthEdirSsoSsoMode(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.edir_sso.sso_mode", val, options...)
}

// GetAuthEdirSsoSsoServer gets the auth.edir_sso.sso_server value from the UTM
func GetAuthEdirSsoSsoServer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.edir_sso.sso_server", &val, options...)
	return
}

// UpdateAuthEdirSsoSsoServer PUTs the auth.edir_sso.sso_server value to the UTM
func UpdateAuthEdirSsoSsoServer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.edir_sso.sso_server", val, options...)
}

// GetAuthEdirSsoSyncInterval gets the auth.edir_sso.sync_interval value from the UTM
func GetAuthEdirSsoSyncInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/auth.edir_sso.sync_interval", &val, options...)
	return
}

// UpdateAuthEdirSsoSyncInterval PUTs the auth.edir_sso.sync_interval value to the UTM
func UpdateAuthEdirSsoSyncInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.edir_sso.sync_interval", val, options...)
}

// GetAuthOtpAutoCreateToken gets the auth.otp.auto_create_token value from the UTM
func GetAuthOtpAutoCreateToken(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.otp.auto_create_token", &val, options...)
	return
}

// UpdateAuthOtpAutoCreateToken PUTs the auth.otp.auto_create_token value to the UTM
func UpdateAuthOtpAutoCreateToken(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.otp.auto_create_token", val, options...)
}

// GetAuthOtpAutoTokenDigest gets the auth.otp.auto_token_digest value from the UTM
func GetAuthOtpAutoTokenDigest(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/auth.otp.auto_token_digest", &val, options...)
	return
}

// UpdateAuthOtpAutoTokenDigest PUTs the auth.otp.auto_token_digest value to the UTM
func UpdateAuthOtpAutoTokenDigest(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.otp.auto_token_digest", val, options...)
}

// GetAuthOtpDefaultTimestep gets the auth.otp.default_timestep value from the UTM
func GetAuthOtpDefaultTimestep(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/auth.otp.default_timestep", &val, options...)
	return
}

// UpdateAuthOtpDefaultTimestep PUTs the auth.otp.default_timestep value to the UTM
func UpdateAuthOtpDefaultTimestep(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.otp.default_timestep", val, options...)
}

// GetAuthOtpFacilities gets the auth.otp.facilities value from the UTM
func GetAuthOtpFacilities(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/auth.otp.facilities", &val, options...)
	return
}

// UpdateAuthOtpFacilities PUTs the auth.otp.facilities value to the UTM
func UpdateAuthOtpFacilities(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.otp.facilities", val, options...)
}

// GetAuthOtpMaxInitTimestepDiff gets the auth.otp.max_init_timestep_diff value from the UTM
func GetAuthOtpMaxInitTimestepDiff(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/auth.otp.max_init_timestep_diff", &val, options...)
	return
}

// UpdateAuthOtpMaxInitTimestepDiff PUTs the auth.otp.max_init_timestep_diff value to the UTM
func UpdateAuthOtpMaxInitTimestepDiff(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.otp.max_init_timestep_diff", val, options...)
}

// GetAuthOtpMaxTimestepDiff gets the auth.otp.max_timestep_diff value from the UTM
func GetAuthOtpMaxTimestepDiff(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/auth.otp.max_timestep_diff", &val, options...)
	return
}

// UpdateAuthOtpMaxTimestepDiff PUTs the auth.otp.max_timestep_diff value to the UTM
func UpdateAuthOtpMaxTimestepDiff(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.otp.max_timestep_diff", val, options...)
}

// GetAuthOtpRequireAllUsers gets the auth.otp.require_all_users value from the UTM
func GetAuthOtpRequireAllUsers(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.otp.require_all_users", &val, options...)
	return
}

// UpdateAuthOtpRequireAllUsers PUTs the auth.otp.require_all_users value to the UTM
func UpdateAuthOtpRequireAllUsers(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.otp.require_all_users", val, options...)
}

// GetAuthOtpRequiredUsers gets the auth.otp.required_users value from the UTM
func GetAuthOtpRequiredUsers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/auth.otp.required_users", &val, options...)
	return
}

// UpdateAuthOtpRequiredUsers PUTs the auth.otp.required_users value to the UTM
func UpdateAuthOtpRequiredUsers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.otp.required_users", val, options...)
}

// GetAuthOtpStatus gets the auth.otp.status value from the UTM
func GetAuthOtpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.otp.status", &val, options...)
	return
}

// UpdateAuthOtpStatus PUTs the auth.otp.status value to the UTM
func UpdateAuthOtpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.otp.status", val, options...)
}

// GetAuthServers gets the auth.servers value from the UTM
func GetAuthServers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/auth.servers", &val, options...)
	return
}

// UpdateAuthServers PUTs the auth.servers value to the UTM
func UpdateAuthServers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.servers", val, options...)
}

// GetAuthUpdateBackendGroupMembersDebug gets the auth.update_backend_group_members.debug value from the UTM
func GetAuthUpdateBackendGroupMembersDebug(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.update_backend_group_members.debug", &val, options...)
	return
}

// UpdateAuthUpdateBackendGroupMembersDebug PUTs the auth.update_backend_group_members.debug value to the UTM
func UpdateAuthUpdateBackendGroupMembersDebug(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.update_backend_group_members.debug", val, options...)
}

// GetAuthUpdateBackendGroupMembersStatus gets the auth.update_backend_group_members.status value from the UTM
func GetAuthUpdateBackendGroupMembersStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/auth.update_backend_group_members.status", &val, options...)
	return
}

// UpdateAuthUpdateBackendGroupMembersStatus PUTs the auth.update_backend_group_members.status value to the UTM
func UpdateAuthUpdateBackendGroupMembersStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/auth.update_backend_group_members.status", val, options...)
}

// GetAweAllowedInterfaces gets the awe.allowed_interfaces value from the UTM
func GetAweAllowedInterfaces(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/awe.allowed_interfaces", &val, options...)
	return
}

// UpdateAweAllowedInterfaces PUTs the awe.allowed_interfaces value to the UTM
func UpdateAweAllowedInterfaces(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.allowed_interfaces", val, options...)
}

// GetAweClients gets the awe.clients value from the UTM
func GetAweClients(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/awe.clients", &val, options...)
	return
}

// UpdateAweClients PUTs the awe.clients value to the UTM
func UpdateAweClients(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.clients", val, options...)
}

// GetAweDevices gets the awe.devices value from the UTM
func GetAweDevices(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/awe.devices", &val, options...)
	return
}

// UpdateAweDevices PUTs the awe.devices value to the UTM
func UpdateAweDevices(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.devices", val, options...)
}

// GetAweGlobalApAutoaccept gets the awe.global.ap_autoaccept value from the UTM
func GetAweGlobalApAutoaccept(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/awe.global.ap_autoaccept", &val, options...)
	return
}

// UpdateAweGlobalApAutoaccept PUTs the awe.global.ap_autoaccept value to the UTM
func UpdateAweGlobalApAutoaccept(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.ap_autoaccept", val, options...)
}

// GetAweGlobalApDebuglevel gets the awe.global.ap_debuglevel value from the UTM
func GetAweGlobalApDebuglevel(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/awe.global.ap_debuglevel", &val, options...)
	return
}

// UpdateAweGlobalApDebuglevel PUTs the awe.global.ap_debuglevel value to the UTM
func UpdateAweGlobalApDebuglevel(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.ap_debuglevel", val, options...)
}

// GetAweGlobalApSoftlimit gets the awe.global.ap_softlimit value from the UTM
func GetAweGlobalApSoftlimit(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/awe.global.ap_softlimit", &val, options...)
	return
}

// UpdateAweGlobalApSoftlimit PUTs the awe.global.ap_softlimit value to the UTM
func UpdateAweGlobalApSoftlimit(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.ap_softlimit", val, options...)
}

// GetAweGlobalApVlantag gets the awe.global.ap_vlantag value from the UTM
func GetAweGlobalApVlantag(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/awe.global.ap_vlantag", &val, options...)
	return
}

// UpdateAweGlobalApVlantag PUTs the awe.global.ap_vlantag value to the UTM
func UpdateAweGlobalApVlantag(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.ap_vlantag", val, options...)
}

// GetAweGlobalAweStatus gets the awe.global.awe_status value from the UTM
func GetAweGlobalAweStatus(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/awe.global.awe_status", &val, options...)
	return
}

// UpdateAweGlobalAweStatus PUTs the awe.global.awe_status value to the UTM
func UpdateAweGlobalAweStatus(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.awe_status", val, options...)
}

// GetAweGlobalBridgeUpdateKickout gets the awe.global.bridge_update_kickout value from the UTM
func GetAweGlobalBridgeUpdateKickout(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/awe.global.bridge_update_kickout", &val, options...)
	return
}

// UpdateAweGlobalBridgeUpdateKickout PUTs the awe.global.bridge_update_kickout value to the UTM
func UpdateAweGlobalBridgeUpdateKickout(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.bridge_update_kickout", val, options...)
}

// GetAweGlobalInitialSetup gets the awe.global.initial_setup value from the UTM
func GetAweGlobalInitialSetup(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/awe.global.initial_setup", &val, options...)
	return
}

// UpdateAweGlobalInitialSetup PUTs the awe.global.initial_setup value to the UTM
func UpdateAweGlobalInitialSetup(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.initial_setup", val, options...)
}

// GetAweGlobalLogLevel gets the awe.global.log_level value from the UTM
func GetAweGlobalLogLevel(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/awe.global.log_level", &val, options...)
	return
}

// UpdateAweGlobalLogLevel PUTs the awe.global.log_level value to the UTM
func UpdateAweGlobalLogLevel(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.log_level", val, options...)
}

// GetAweGlobalMagicIp gets the awe.global.magic_ip value from the UTM
func GetAweGlobalMagicIp(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/awe.global.magic_ip", &val, options...)
	return
}

// UpdateAweGlobalMagicIp PUTs the awe.global.magic_ip value to the UTM
func UpdateAweGlobalMagicIp(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.magic_ip", val, options...)
}

// GetAweGlobalMultiWifiIfaceBr gets the awe.global.multi_wifi_iface_br value from the UTM
func GetAweGlobalMultiWifiIfaceBr(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/awe.global.multi_wifi_iface_br", &val, options...)
	return
}

// UpdateAweGlobalMultiWifiIfaceBr PUTs the awe.global.multi_wifi_iface_br value to the UTM
func UpdateAweGlobalMultiWifiIfaceBr(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.multi_wifi_iface_br", val, options...)
}

// GetAweGlobalNotificationTimeout gets the awe.global.notification_timeout value from the UTM
func GetAweGlobalNotificationTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/awe.global.notification_timeout", &val, options...)
	return
}

// UpdateAweGlobalNotificationTimeout PUTs the awe.global.notification_timeout value to the UTM
func UpdateAweGlobalNotificationTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.notification_timeout", val, options...)
}

// GetAweGlobalRadiusConf gets the awe.global.radius_conf value from the UTM
func GetAweGlobalRadiusConf(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/awe.global.radius_conf", &val, options...)
	return
}

// UpdateAweGlobalRadiusConf PUTs the awe.global.radius_conf value to the UTM
func UpdateAweGlobalRadiusConf(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.radius_conf", val, options...)
}

// GetAweGlobalRootpw gets the awe.global.rootpw value from the UTM
func GetAweGlobalRootpw(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/awe.global.rootpw", &val, options...)
	return
}

// UpdateAweGlobalRootpw PUTs the awe.global.rootpw value to the UTM
func UpdateAweGlobalRootpw(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.rootpw", val, options...)
}

// GetAweGlobalStayOnline gets the awe.global.stay_online value from the UTM
func GetAweGlobalStayOnline(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/awe.global.stay_online", &val, options...)
	return
}

// UpdateAweGlobalStayOnline PUTs the awe.global.stay_online value to the UTM
func UpdateAweGlobalStayOnline(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.stay_online", val, options...)
}

// GetAweGlobalStoreBssStats gets the awe.global.store_bss_stats value from the UTM
func GetAweGlobalStoreBssStats(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/awe.global.store_bss_stats", &val, options...)
	return
}

// UpdateAweGlobalStoreBssStats PUTs the awe.global.store_bss_stats value to the UTM
func UpdateAweGlobalStoreBssStats(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.store_bss_stats", val, options...)
}

// GetAweGlobalTunnelIdOffset gets the awe.global.tunnel_id_offset value from the UTM
func GetAweGlobalTunnelIdOffset(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/awe.global.tunnel_id_offset", &val, options...)
	return
}

// UpdateAweGlobalTunnelIdOffset PUTs the awe.global.tunnel_id_offset value to the UTM
func UpdateAweGlobalTunnelIdOffset(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.tunnel_id_offset", val, options...)
}

// GetAweGlobalVlantagging gets the awe.global.vlantagging value from the UTM
func GetAweGlobalVlantagging(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/awe.global.vlantagging", &val, options...)
	return
}

// UpdateAweGlobalVlantagging PUTs the awe.global.vlantagging value to the UTM
func UpdateAweGlobalVlantagging(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.global.vlantagging", val, options...)
}

// GetAweNetworks gets the awe.networks value from the UTM
func GetAweNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/awe.networks", &val, options...)
	return
}

// UpdateAweNetworks PUTs the awe.networks value to the UTM
func UpdateAweNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awe.networks", val, options...)
}

// GetAwscliProfiles gets the awscli.profiles value from the UTM
func GetAwscliProfiles(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/awscli.profiles", &val, options...)
	return
}

// UpdateAwscliProfiles PUTs the awscli.profiles value to the UTM
func UpdateAwscliProfiles(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/awscli.profiles", val, options...)
}

// GetBackupEncryption gets the backup.encryption value from the UTM
func GetBackupEncryption(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/backup.encryption", &val, options...)
	return
}

// UpdateBackupEncryption PUTs the backup.encryption value to the UTM
func UpdateBackupEncryption(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/backup.encryption", val, options...)
}

// GetBackupInterval gets the backup.interval value from the UTM
func GetBackupInterval(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/backup.interval", &val, options...)
	return
}

// UpdateBackupInterval PUTs the backup.interval value to the UTM
func UpdateBackupInterval(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/backup.interval", val, options...)
}

// GetBackupMaxBackups gets the backup.max_backups value from the UTM
func GetBackupMaxBackups(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/backup.max_backups", &val, options...)
	return
}

// UpdateBackupMaxBackups PUTs the backup.max_backups value to the UTM
func UpdateBackupMaxBackups(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/backup.max_backups", val, options...)
}

// GetBackupPassword gets the backup.password value from the UTM
func GetBackupPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/backup.password", &val, options...)
	return
}

// UpdateBackupPassword PUTs the backup.password value to the UTM
func UpdateBackupPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/backup.password", val, options...)
}

// GetBackupRecipients gets the backup.recipients value from the UTM
func GetBackupRecipients(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/backup.recipients", &val, options...)
	return
}

// UpdateBackupRecipients PUTs the backup.recipients value to the UTM
func UpdateBackupRecipients(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/backup.recipients", val, options...)
}

// GetBackupStatus gets the backup.status value from the UTM
func GetBackupStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/backup.status", &val, options...)
	return
}

// UpdateBackupStatus PUTs the backup.status value to the UTM
func UpdateBackupStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/backup.status", val, options...)
}

// GetCaCaGost gets the ca.ca_gost value from the UTM
func GetCaCaGost(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ca.ca_gost", &val, options...)
	return
}

// UpdateCaCaGost PUTs the ca.ca_gost value to the UTM
func UpdateCaCaGost(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.ca_gost", val, options...)
}

// GetCaCaIpsec gets the ca.ca_ipsec value from the UTM
func GetCaCaIpsec(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ca.ca_ipsec", &val, options...)
	return
}

// UpdateCaCaIpsec PUTs the ca.ca_ipsec value to the UTM
func UpdateCaCaIpsec(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.ca_ipsec", val, options...)
}

// GetCaCaProxies gets the ca.ca_proxies value from the UTM
func GetCaCaProxies(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ca.ca_proxies", &val, options...)
	return
}

// UpdateCaCaProxies PUTs the ca.ca_proxies value to the UTM
func UpdateCaCaProxies(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.ca_proxies", val, options...)
}

// GetCaCaRed gets the ca.ca_red value from the UTM
func GetCaCaRed(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ca.ca_red", &val, options...)
	return
}

// UpdateCaCaRed PUTs the ca.ca_red value to the UTM
func UpdateCaCaRed(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.ca_red", val, options...)
}

// GetCaDefKeysize gets the ca.def_keysize value from the UTM
func GetCaDefKeysize(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ca.def_keysize", &val, options...)
	return
}

// UpdateCaDefKeysize PUTs the ca.def_keysize value to the UTM
func UpdateCaDefKeysize(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.def_keysize", val, options...)
}

// GetCaGlobalCasEmailEncryptionTrustNewCas gets the ca.global_cas.email_encryption.trust_new_cas value from the UTM
func GetCaGlobalCasEmailEncryptionTrustNewCas(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ca.global_cas.email_encryption.trust_new_cas", &val, options...)
	return
}

// UpdateCaGlobalCasEmailEncryptionTrustNewCas PUTs the ca.global_cas.email_encryption.trust_new_cas value to the UTM
func UpdateCaGlobalCasEmailEncryptionTrustNewCas(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.global_cas.email_encryption.trust_new_cas", val, options...)
}

// GetCaGlobalCasEmailEncryptionTrusted gets the ca.global_cas.email_encryption.trusted value from the UTM
func GetCaGlobalCasEmailEncryptionTrusted(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ca.global_cas.email_encryption.trusted", &val, options...)
	return
}

// UpdateCaGlobalCasEmailEncryptionTrusted PUTs the ca.global_cas.email_encryption.trusted value to the UTM
func UpdateCaGlobalCasEmailEncryptionTrusted(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.global_cas.email_encryption.trusted", val, options...)
}

// GetCaGlobalCasEmailEncryptionUntrusted gets the ca.global_cas.email_encryption.untrusted value from the UTM
func GetCaGlobalCasEmailEncryptionUntrusted(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ca.global_cas.email_encryption.untrusted", &val, options...)
	return
}

// UpdateCaGlobalCasEmailEncryptionUntrusted PUTs the ca.global_cas.email_encryption.untrusted value to the UTM
func UpdateCaGlobalCasEmailEncryptionUntrusted(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.global_cas.email_encryption.untrusted", val, options...)
}

// GetCaGlobalCasHttpProxyTrustNewCas gets the ca.global_cas.http_proxy.trust_new_cas value from the UTM
func GetCaGlobalCasHttpProxyTrustNewCas(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ca.global_cas.http_proxy.trust_new_cas", &val, options...)
	return
}

// UpdateCaGlobalCasHttpProxyTrustNewCas PUTs the ca.global_cas.http_proxy.trust_new_cas value to the UTM
func UpdateCaGlobalCasHttpProxyTrustNewCas(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.global_cas.http_proxy.trust_new_cas", val, options...)
}

// GetCaGlobalCasHttpProxyTrusted gets the ca.global_cas.http_proxy.trusted value from the UTM
func GetCaGlobalCasHttpProxyTrusted(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ca.global_cas.http_proxy.trusted", &val, options...)
	return
}

// UpdateCaGlobalCasHttpProxyTrusted PUTs the ca.global_cas.http_proxy.trusted value to the UTM
func UpdateCaGlobalCasHttpProxyTrusted(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.global_cas.http_proxy.trusted", val, options...)
}

// GetCaGlobalCasHttpProxyUntrusted gets the ca.global_cas.http_proxy.untrusted value from the UTM
func GetCaGlobalCasHttpProxyUntrusted(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ca.global_cas.http_proxy.untrusted", &val, options...)
	return
}

// UpdateCaGlobalCasHttpProxyUntrusted PUTs the ca.global_cas.http_proxy.untrusted value to the UTM
func UpdateCaGlobalCasHttpProxyUntrusted(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.global_cas.http_proxy.untrusted", val, options...)
}

// GetCaLetsencryptAccountId gets the ca.letsencrypt.account_id value from the UTM
func GetCaLetsencryptAccountId(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ca.letsencrypt.account_id", &val, options...)
	return
}

// UpdateCaLetsencryptAccountId PUTs the ca.letsencrypt.account_id value to the UTM
func UpdateCaLetsencryptAccountId(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.letsencrypt.account_id", val, options...)
}

// GetCaLetsencryptAccountKey gets the ca.letsencrypt.account_key value from the UTM
func GetCaLetsencryptAccountKey(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ca.letsencrypt.account_key", &val, options...)
	return
}

// UpdateCaLetsencryptAccountKey PUTs the ca.letsencrypt.account_key value to the UTM
func UpdateCaLetsencryptAccountKey(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.letsencrypt.account_key", val, options...)
}

// GetCaLetsencryptAcmeServer gets the ca.letsencrypt.acme_server value from the UTM
func GetCaLetsencryptAcmeServer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ca.letsencrypt.acme_server", &val, options...)
	return
}

// UpdateCaLetsencryptAcmeServer PUTs the ca.letsencrypt.acme_server value to the UTM
func UpdateCaLetsencryptAcmeServer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.letsencrypt.acme_server", val, options...)
}

// GetCaLetsencryptDebug gets the ca.letsencrypt.debug value from the UTM
func GetCaLetsencryptDebug(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ca.letsencrypt.debug", &val, options...)
	return
}

// UpdateCaLetsencryptDebug PUTs the ca.letsencrypt.debug value to the UTM
func UpdateCaLetsencryptDebug(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.letsencrypt.debug", val, options...)
}

// GetCaLetsencryptErrorInfo gets the ca.letsencrypt.error_info value from the UTM
func GetCaLetsencryptErrorInfo(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ca.letsencrypt.error_info", &val, options...)
	return
}

// UpdateCaLetsencryptErrorInfo PUTs the ca.letsencrypt.error_info value to the UTM
func UpdateCaLetsencryptErrorInfo(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.letsencrypt.error_info", val, options...)
}

// GetCaLetsencryptErrorMessage gets the ca.letsencrypt.error_message value from the UTM
func GetCaLetsencryptErrorMessage(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ca.letsencrypt.error_message", &val, options...)
	return
}

// UpdateCaLetsencryptErrorMessage PUTs the ca.letsencrypt.error_message value to the UTM
func UpdateCaLetsencryptErrorMessage(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.letsencrypt.error_message", val, options...)
}

// GetCaLetsencryptRegistrationInfo gets the ca.letsencrypt.registration_info value from the UTM
func GetCaLetsencryptRegistrationInfo(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ca.letsencrypt.registration_info", &val, options...)
	return
}

// UpdateCaLetsencryptRegistrationInfo PUTs the ca.letsencrypt.registration_info value to the UTM
func UpdateCaLetsencryptRegistrationInfo(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.letsencrypt.registration_info", val, options...)
}

// GetCaLetsencryptStatus gets the ca.letsencrypt.status value from the UTM
func GetCaLetsencryptStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ca.letsencrypt.status", &val, options...)
	return
}

// UpdateCaLetsencryptStatus PUTs the ca.letsencrypt.status value to the UTM
func UpdateCaLetsencryptStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.letsencrypt.status", val, options...)
}

// GetCaLetsencryptTosUrl gets the ca.letsencrypt.tos_url value from the UTM
func GetCaLetsencryptTosUrl(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ca.letsencrypt.tos_url", &val, options...)
	return
}

// UpdateCaLetsencryptTosUrl PUTs the ca.letsencrypt.tos_url value to the UTM
func UpdateCaLetsencryptTosUrl(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ca.letsencrypt.tos_url", val, options...)
}

// GetCrlsCrls gets the crls.crls value from the UTM
func GetCrlsCrls(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/crls.crls", &val, options...)
	return
}

// UpdateCrlsCrls PUTs the crls.crls value to the UTM
func UpdateCrlsCrls(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/crls.crls", val, options...)
}

// GetCssAvPrimaryEngine gets the css.av_primary_engine value from the UTM
func GetCssAvPrimaryEngine(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/css.av_primary_engine", &val, options...)
	return
}

// UpdateCssAvPrimaryEngine PUTs the css.av_primary_engine value to the UTM
func UpdateCssAvPrimaryEngine(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/css.av_primary_engine", val, options...)
}

// GetCssSxlLiveprotection gets the css.sxl_liveprotection value from the UTM
func GetCssSxlLiveprotection(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/css.sxl_liveprotection", &val, options...)
	return
}

// UpdateCssSxlLiveprotection PUTs the css.sxl_liveprotection value to the UTM
func UpdateCssSxlLiveprotection(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/css.sxl_liveprotection", val, options...)
}

// GetCssSxlSampleSubmit gets the css.sxl_sample_submit value from the UTM
func GetCssSxlSampleSubmit(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/css.sxl_sample_submit", &val, options...)
	return
}

// UpdateCssSxlSampleSubmit PUTs the css.sxl_sample_submit value to the UTM
func UpdateCssSxlSampleSubmit(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/css.sxl_sample_submit", val, options...)
}

// GetCustomizationEppLastUpdated gets the customization.epp.last_updated value from the UTM
func GetCustomizationEppLastUpdated(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/customization.epp.last_updated", &val, options...)
	return
}

// UpdateCustomizationEppLastUpdated PUTs the customization.epp.last_updated value to the UTM
func UpdateCustomizationEppLastUpdated(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/customization.epp.last_updated", val, options...)
}

// GetCustomizationEppResourcesRoot gets the customization.epp.resources_root value from the UTM
func GetCustomizationEppResourcesRoot(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/customization.epp.resources_root", &val, options...)
	return
}

// UpdateCustomizationEppResourcesRoot PUTs the customization.epp.resources_root value to the UTM
func UpdateCustomizationEppResourcesRoot(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/customization.epp.resources_root", val, options...)
}

// GetCustomizationHttpCustomAssets gets the customization.http.custom_assets value from the UTM
func GetCustomizationHttpCustomAssets(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/customization.http.custom_assets", &val, options...)
	return
}

// UpdateCustomizationHttpCustomAssets PUTs the customization.http.custom_assets value to the UTM
func UpdateCustomizationHttpCustomAssets(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/customization.http.custom_assets", val, options...)
}

// GetCustomizationHttpCustomTemplates gets the customization.http.custom_templates value from the UTM
func GetCustomizationHttpCustomTemplates(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/customization.http.custom_templates", &val, options...)
	return
}

// UpdateCustomizationHttpCustomTemplates PUTs the customization.http.custom_templates value to the UTM
func UpdateCustomizationHttpCustomTemplates(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/customization.http.custom_templates", val, options...)
}

// GetCustomizationHttpLastUpdated gets the customization.http.last_updated value from the UTM
func GetCustomizationHttpLastUpdated(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/customization.http.last_updated", &val, options...)
	return
}

// UpdateCustomizationHttpLastUpdated PUTs the customization.http.last_updated value to the UTM
func UpdateCustomizationHttpLastUpdated(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/customization.http.last_updated", val, options...)
}

// GetDebugmodeCrashReport gets the debugmode.crash_report value from the UTM
func GetDebugmodeCrashReport(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/debugmode.crash_report", &val, options...)
	return
}

// UpdateDebugmodeCrashReport PUTs the debugmode.crash_report value to the UTM
func UpdateDebugmodeCrashReport(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/debugmode.crash_report", val, options...)
}

// GetDebugmodeEnabled gets the debugmode.enabled value from the UTM
func GetDebugmodeEnabled(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/debugmode.enabled", &val, options...)
	return
}

// UpdateDebugmodeEnabled PUTs the debugmode.enabled value to the UTM
func UpdateDebugmodeEnabled(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/debugmode.enabled", val, options...)
}

// GetDhcpDhclientBindToInterface gets the dhcp.dhclient_bind_to_interface value from the UTM
func GetDhcpDhclientBindToInterface(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/dhcp.dhclient_bind_to_interface", &val, options...)
	return
}

// UpdateDhcpDhclientBindToInterface PUTs the dhcp.dhclient_bind_to_interface value to the UTM
func UpdateDhcpDhclientBindToInterface(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dhcp.dhclient_bind_to_interface", val, options...)
}

// GetDhcpRelayDhcpServer gets the dhcp.relay.dhcp_server value from the UTM
func GetDhcpRelayDhcpServer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/dhcp.relay.dhcp_server", &val, options...)
	return
}

// UpdateDhcpRelayDhcpServer PUTs the dhcp.relay.dhcp_server value to the UTM
func UpdateDhcpRelayDhcpServer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dhcp.relay.dhcp_server", val, options...)
}

// GetDhcpRelayInterfaces gets the dhcp.relay.interfaces value from the UTM
func GetDhcpRelayInterfaces(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/dhcp.relay.interfaces", &val, options...)
	return
}

// UpdateDhcpRelayInterfaces PUTs the dhcp.relay.interfaces value to the UTM
func UpdateDhcpRelayInterfaces(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dhcp.relay.interfaces", val, options...)
}

// GetDhcpRelayStatus gets the dhcp.relay.status value from the UTM
func GetDhcpRelayStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/dhcp.relay.status", &val, options...)
	return
}

// UpdateDhcpRelayStatus PUTs the dhcp.relay.status value to the UTM
func UpdateDhcpRelayStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dhcp.relay.status", val, options...)
}

// GetDhcpRelay6ItfsFacingClients gets the dhcp.relay6.itfs_facing_clients value from the UTM
func GetDhcpRelay6ItfsFacingClients(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/dhcp.relay6.itfs_facing_clients", &val, options...)
	return
}

// UpdateDhcpRelay6ItfsFacingClients PUTs the dhcp.relay6.itfs_facing_clients value to the UTM
func UpdateDhcpRelay6ItfsFacingClients(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dhcp.relay6.itfs_facing_clients", val, options...)
}

// GetDhcpRelay6ItfsFacingServer6 gets the dhcp.relay6.itfs_facing_server6 value from the UTM
func GetDhcpRelay6ItfsFacingServer6(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/dhcp.relay6.itfs_facing_server6", &val, options...)
	return
}

// UpdateDhcpRelay6ItfsFacingServer6 PUTs the dhcp.relay6.itfs_facing_server6 value to the UTM
func UpdateDhcpRelay6ItfsFacingServer6(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dhcp.relay6.itfs_facing_server6", val, options...)
}

// GetDhcpRelay6Status gets the dhcp.relay6.status value from the UTM
func GetDhcpRelay6Status(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/dhcp.relay6.status", &val, options...)
	return
}

// UpdateDhcpRelay6Status PUTs the dhcp.relay6.status value to the UTM
func UpdateDhcpRelay6Status(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dhcp.relay6.status", val, options...)
}

// GetDhcpServerCustom4 gets the dhcp.server.custom4 value from the UTM
func GetDhcpServerCustom4(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/dhcp.server.custom4", &val, options...)
	return
}

// UpdateDhcpServerCustom4 PUTs the dhcp.server.custom4 value to the UTM
func UpdateDhcpServerCustom4(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dhcp.server.custom4", val, options...)
}

// GetDhcpServerCustom6 gets the dhcp.server.custom6 value from the UTM
func GetDhcpServerCustom6(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/dhcp.server.custom6", &val, options...)
	return
}

// UpdateDhcpServerCustom6 PUTs the dhcp.server.custom6 value to the UTM
func UpdateDhcpServerCustom6(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dhcp.server.custom6", val, options...)
}

// GetDhcpServerServers gets the dhcp.server.servers value from the UTM
func GetDhcpServerServers(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/dhcp.server.servers", &val, options...)
	return
}

// UpdateDhcpServerServers PUTs the dhcp.server.servers value to the UTM
func UpdateDhcpServerServers(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dhcp.server.servers", val, options...)
}

// GetDigestAllowedNetworks gets the digest.allowed_networks value from the UTM
func GetDigestAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/digest.allowed_networks", &val, options...)
	return
}

// UpdateDigestAllowedNetworks PUTs the digest.allowed_networks value to the UTM
func UpdateDigestAllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/digest.allowed_networks", val, options...)
}

// GetDigestCustomText gets the digest.custom_text value from the UTM
func GetDigestCustomText(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/digest.custom_text", &val, options...)
	return
}

// UpdateDigestCustomText PUTs the digest.custom_text value to the UTM
func UpdateDigestCustomText(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/digest.custom_text", val, options...)
}

// GetDigestDomains gets the digest.domains value from the UTM
func GetDigestDomains(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/digest.domains", &val, options...)
	return
}

// UpdateDigestDomains PUTs the digest.domains value to the UTM
func UpdateDigestDomains(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/digest.domains", val, options...)
}

// GetDigestHostname gets the digest.hostname value from the UTM
func GetDigestHostname(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/digest.hostname", &val, options...)
	return
}

// UpdateDigestHostname PUTs the digest.hostname value to the UTM
func UpdateDigestHostname(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/digest.hostname", val, options...)
}

// GetDigestMailinglists gets the digest.mailinglists value from the UTM
func GetDigestMailinglists(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/digest.mailinglists", &val, options...)
	return
}

// UpdateDigestMailinglists PUTs the digest.mailinglists value to the UTM
func UpdateDigestMailinglists(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/digest.mailinglists", val, options...)
}

// GetDigestPort gets the digest.port value from the UTM
func GetDigestPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/digest.port", &val, options...)
	return
}

// UpdateDigestPort PUTs the digest.port value to the UTM
func UpdateDigestPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/digest.port", val, options...)
}

// GetDigestSendTimeOne gets the digest.send_time_one value from the UTM
func GetDigestSendTimeOne(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/digest.send_time_one", &val, options...)
	return
}

// UpdateDigestSendTimeOne PUTs the digest.send_time_one value to the UTM
func UpdateDigestSendTimeOne(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/digest.send_time_one", val, options...)
}

// GetDigestSendTimeTwo gets the digest.send_time_two value from the UTM
func GetDigestSendTimeTwo(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/digest.send_time_two", &val, options...)
	return
}

// UpdateDigestSendTimeTwo PUTs the digest.send_time_two value to the UTM
func UpdateDigestSendTimeTwo(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/digest.send_time_two", val, options...)
}

// GetDigestSkiplist gets the digest.skiplist value from the UTM
func GetDigestSkiplist(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/digest.skiplist", &val, options...)
	return
}

// UpdateDigestSkiplist PUTs the digest.skiplist value to the UTM
func UpdateDigestSkiplist(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/digest.skiplist", val, options...)
}

// GetDigestStatus gets the digest.status value from the UTM
func GetDigestStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/digest.status", &val, options...)
	return
}

// UpdateDigestStatus PUTs the digest.status value to the UTM
func UpdateDigestStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/digest.status", val, options...)
}

// GetDigestUserRelease gets the digest.user_release value from the UTM
func GetDigestUserRelease(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/digest.user_release", &val, options...)
	return
}

// UpdateDigestUserRelease PUTs the digest.user_release value to the UTM
func UpdateDigestUserRelease(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/digest.user_release", val, options...)
}

// GetDnsAllowedNetworks gets the dns.allowed_networks value from the UTM
func GetDnsAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/dns.allowed_networks", &val, options...)
	return
}

// UpdateDnsAllowedNetworks PUTs the dns.allowed_networks value to the UTM
func UpdateDnsAllowedNetworks(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dns.allowed_networks", val, options...)
}

// GetDnsAxfr gets the dns.axfr value from the UTM
func GetDnsAxfr(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/dns.axfr", &val, options...)
	return
}

// UpdateDnsAxfr PUTs the dns.axfr value to the UTM
func UpdateDnsAxfr(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dns.axfr", val, options...)
}

// GetDnsDnssec gets the dns.dnssec value from the UTM
func GetDnsDnssec(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/dns.dnssec", &val, options...)
	return
}

// UpdateDnsDnssec PUTs the dns.dnssec value to the UTM
func UpdateDnsDnssec(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dns.dnssec", val, options...)
}

// GetDnsEmail gets the dns.email value from the UTM
func GetDnsEmail(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/dns.email", &val, options...)
	return
}

// UpdateDnsEmail PUTs the dns.email value to the UTM
func UpdateDnsEmail(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dns.email", val, options...)
}

// GetDnsEmptyZones gets the dns.empty_zones value from the UTM
func GetDnsEmptyZones(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/dns.empty_zones", &val, options...)
	return
}

// UpdateDnsEmptyZones PUTs the dns.empty_zones value to the UTM
func UpdateDnsEmptyZones(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dns.empty_zones", val, options...)
}

// GetDnsFwdDynamic gets the dns.fwd_dynamic value from the UTM
func GetDnsFwdDynamic(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/dns.fwd_dynamic", &val, options...)
	return
}

// UpdateDnsFwdDynamic PUTs the dns.fwd_dynamic value to the UTM
func UpdateDnsFwdDynamic(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dns.fwd_dynamic", val, options...)
}

// GetDnsFwdStatic gets the dns.fwd_static value from the UTM
func GetDnsFwdStatic(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/dns.fwd_static", &val, options...)
	return
}

// UpdateDnsFwdStatic PUTs the dns.fwd_static value to the UTM
func UpdateDnsFwdStatic(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dns.fwd_static", val, options...)
}

// GetDnsRecheckInterval gets the dns.recheck_interval value from the UTM
func GetDnsRecheckInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/dns.recheck_interval", &val, options...)
	return
}

// UpdateDnsRecheckInterval PUTs the dns.recheck_interval value to the UTM
func UpdateDnsRecheckInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dns.recheck_interval", val, options...)
}

// GetDnsRetryForNonexistingNxdomain gets the dns.retry_for_nonexisting_nxdomain value from the UTM
func GetDnsRetryForNonexistingNxdomain(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/dns.retry_for_nonexisting_nxdomain", &val, options...)
	return
}

// UpdateDnsRetryForNonexistingNxdomain PUTs the dns.retry_for_nonexisting_nxdomain value to the UTM
func UpdateDnsRetryForNonexistingNxdomain(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dns.retry_for_nonexisting_nxdomain", val, options...)
}

// GetDnsRoutes gets the dns.routes value from the UTM
func GetDnsRoutes(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/dns.routes", &val, options...)
	return
}

// UpdateDnsRoutes PUTs the dns.routes value to the UTM
func UpdateDnsRoutes(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dns.routes", val, options...)
}

// GetDyndnsRules gets the dyndns.rules value from the UTM
func GetDyndnsRules(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/dyndns.rules", &val, options...)
	return
}

// UpdateDyndnsRules PUTs the dyndns.rules value to the UTM
func UpdateDyndnsRules(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/dyndns.rules", val, options...)
}

// GetEmailpkiAuthorityCert gets the emailpki.authority.cert value from the UTM
func GetEmailpkiAuthorityCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/emailpki.authority.cert", &val, options...)
	return
}

// UpdateEmailpkiAuthorityCert PUTs the emailpki.authority.cert value to the UTM
func UpdateEmailpkiAuthorityCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.authority.cert", val, options...)
}

// GetEmailpkiAuthorityFingerprint gets the emailpki.authority.fingerprint value from the UTM
func GetEmailpkiAuthorityFingerprint(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/emailpki.authority.fingerprint", &val, options...)
	return
}

// UpdateEmailpkiAuthorityFingerprint PUTs the emailpki.authority.fingerprint value to the UTM
func UpdateEmailpkiAuthorityFingerprint(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.authority.fingerprint", val, options...)
}

// GetEmailpkiAuthorityKey gets the emailpki.authority.key value from the UTM
func GetEmailpkiAuthorityKey(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/emailpki.authority.key", &val, options...)
	return
}

// UpdateEmailpkiAuthorityKey PUTs the emailpki.authority.key value to the UTM
func UpdateEmailpkiAuthorityKey(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.authority.key", val, options...)
}

// GetEmailpkiAuthorityPostmasterFingerprint gets the emailpki.authority.postmaster_fingerprint value from the UTM
func GetEmailpkiAuthorityPostmasterFingerprint(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/emailpki.authority.postmaster_fingerprint", &val, options...)
	return
}

// UpdateEmailpkiAuthorityPostmasterFingerprint PUTs the emailpki.authority.postmaster_fingerprint value to the UTM
func UpdateEmailpkiAuthorityPostmasterFingerprint(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.authority.postmaster_fingerprint", val, options...)
}

// GetEmailpkiAuthorityPostmasterPrivkey gets the emailpki.authority.postmaster_privkey value from the UTM
func GetEmailpkiAuthorityPostmasterPrivkey(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/emailpki.authority.postmaster_privkey", &val, options...)
	return
}

// UpdateEmailpkiAuthorityPostmasterPrivkey PUTs the emailpki.authority.postmaster_privkey value to the UTM
func UpdateEmailpkiAuthorityPostmasterPrivkey(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.authority.postmaster_privkey", val, options...)
}

// GetEmailpkiAuthorityPostmasterPubkey gets the emailpki.authority.postmaster_pubkey value from the UTM
func GetEmailpkiAuthorityPostmasterPubkey(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/emailpki.authority.postmaster_pubkey", &val, options...)
	return
}

// UpdateEmailpkiAuthorityPostmasterPubkey PUTs the emailpki.authority.postmaster_pubkey value to the UTM
func UpdateEmailpkiAuthorityPostmasterPubkey(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.authority.postmaster_pubkey", val, options...)
}

// GetEmailpkiGlobalCity gets the emailpki.global.city value from the UTM
func GetEmailpkiGlobalCity(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/emailpki.global.city", &val, options...)
	return
}

// UpdateEmailpkiGlobalCity PUTs the emailpki.global.city value to the UTM
func UpdateEmailpkiGlobalCity(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.global.city", val, options...)
}

// GetEmailpkiGlobalCountry gets the emailpki.global.country value from the UTM
func GetEmailpkiGlobalCountry(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/emailpki.global.country", &val, options...)
	return
}

// UpdateEmailpkiGlobalCountry PUTs the emailpki.global.country value to the UTM
func UpdateEmailpkiGlobalCountry(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.global.country", val, options...)
}

// GetEmailpkiGlobalOrganization gets the emailpki.global.organization value from the UTM
func GetEmailpkiGlobalOrganization(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/emailpki.global.organization", &val, options...)
	return
}

// UpdateEmailpkiGlobalOrganization PUTs the emailpki.global.organization value to the UTM
func UpdateEmailpkiGlobalOrganization(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.global.organization", val, options...)
}

// GetEmailpkiGlobalPostmaster gets the emailpki.global.postmaster value from the UTM
func GetEmailpkiGlobalPostmaster(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/emailpki.global.postmaster", &val, options...)
	return
}

// UpdateEmailpkiGlobalPostmaster PUTs the emailpki.global.postmaster value to the UTM
func UpdateEmailpkiGlobalPostmaster(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.global.postmaster", val, options...)
}

// GetEmailpkiGlobalStatus gets the emailpki.global.status value from the UTM
func GetEmailpkiGlobalStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/emailpki.global.status", &val, options...)
	return
}

// UpdateEmailpkiGlobalStatus PUTs the emailpki.global.status value to the UTM
func UpdateEmailpkiGlobalStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.global.status", val, options...)
}

// GetEmailpkiObjectsCas gets the emailpki.objects.cas value from the UTM
func GetEmailpkiObjectsCas(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/emailpki.objects.cas", &val, options...)
	return
}

// UpdateEmailpkiObjectsCas PUTs the emailpki.objects.cas value to the UTM
func UpdateEmailpkiObjectsCas(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.objects.cas", val, options...)
}

// GetEmailpkiObjectsOpenpgp gets the emailpki.objects.openpgp value from the UTM
func GetEmailpkiObjectsOpenpgp(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/emailpki.objects.openpgp", &val, options...)
	return
}

// UpdateEmailpkiObjectsOpenpgp PUTs the emailpki.objects.openpgp value to the UTM
func UpdateEmailpkiObjectsOpenpgp(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.objects.openpgp", val, options...)
}

// GetEmailpkiObjectsSmime gets the emailpki.objects.smime value from the UTM
func GetEmailpkiObjectsSmime(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/emailpki.objects.smime", &val, options...)
	return
}

// UpdateEmailpkiObjectsSmime PUTs the emailpki.objects.smime value to the UTM
func UpdateEmailpkiObjectsSmime(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.objects.smime", val, options...)
}

// GetEmailpkiObjectsUsers gets the emailpki.objects.users value from the UTM
func GetEmailpkiObjectsUsers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/emailpki.objects.users", &val, options...)
	return
}

// UpdateEmailpkiObjectsUsers PUTs the emailpki.objects.users value to the UTM
func UpdateEmailpkiObjectsUsers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.objects.users", val, options...)
}

// GetEmailpkiOpenpgpMainKeysize gets the emailpki.openpgp.main_keysize value from the UTM
func GetEmailpkiOpenpgpMainKeysize(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/emailpki.openpgp.main_keysize", &val, options...)
	return
}

// UpdateEmailpkiOpenpgpMainKeysize PUTs the emailpki.openpgp.main_keysize value to the UTM
func UpdateEmailpkiOpenpgpMainKeysize(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.openpgp.main_keysize", val, options...)
}

// GetEmailpkiOpenpgpSubKeysize gets the emailpki.openpgp.sub_keysize value from the UTM
func GetEmailpkiOpenpgpSubKeysize(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/emailpki.openpgp.sub_keysize", &val, options...)
	return
}

// UpdateEmailpkiOpenpgpSubKeysize PUTs the emailpki.openpgp.sub_keysize value to the UTM
func UpdateEmailpkiOpenpgpSubKeysize(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.openpgp.sub_keysize", val, options...)
}

// GetEmailpkiOptionsExternalAuto gets the emailpki.options.external_auto value from the UTM
func GetEmailpkiOptionsExternalAuto(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/emailpki.options.external_auto", &val, options...)
	return
}

// UpdateEmailpkiOptionsExternalAuto PUTs the emailpki.options.external_auto value to the UTM
func UpdateEmailpkiOptionsExternalAuto(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.options.external_auto", val, options...)
}

// GetEmailpkiOptionsKeyserver gets the emailpki.options.keyserver value from the UTM
func GetEmailpkiOptionsKeyserver(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/emailpki.options.keyserver", &val, options...)
	return
}

// UpdateEmailpkiOptionsKeyserver PUTs the emailpki.options.keyserver value to the UTM
func UpdateEmailpkiOptionsKeyserver(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.options.keyserver", val, options...)
}

// GetEmailpkiOptionsPolicyDecryption gets the emailpki.options.policy_decryption value from the UTM
func GetEmailpkiOptionsPolicyDecryption(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/emailpki.options.policy_decryption", &val, options...)
	return
}

// UpdateEmailpkiOptionsPolicyDecryption PUTs the emailpki.options.policy_decryption value to the UTM
func UpdateEmailpkiOptionsPolicyDecryption(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.options.policy_decryption", val, options...)
}

// GetEmailpkiOptionsPolicyEncryption gets the emailpki.options.policy_encryption value from the UTM
func GetEmailpkiOptionsPolicyEncryption(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/emailpki.options.policy_encryption", &val, options...)
	return
}

// UpdateEmailpkiOptionsPolicyEncryption PUTs the emailpki.options.policy_encryption value to the UTM
func UpdateEmailpkiOptionsPolicyEncryption(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.options.policy_encryption", val, options...)
}

// GetEmailpkiOptionsPolicySign gets the emailpki.options.policy_sign value from the UTM
func GetEmailpkiOptionsPolicySign(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/emailpki.options.policy_sign", &val, options...)
	return
}

// UpdateEmailpkiOptionsPolicySign PUTs the emailpki.options.policy_sign value to the UTM
func UpdateEmailpkiOptionsPolicySign(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.options.policy_sign", val, options...)
}

// GetEmailpkiOptionsPolicyVerify gets the emailpki.options.policy_verify value from the UTM
func GetEmailpkiOptionsPolicyVerify(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/emailpki.options.policy_verify", &val, options...)
	return
}

// UpdateEmailpkiOptionsPolicyVerify PUTs the emailpki.options.policy_verify value to the UTM
func UpdateEmailpkiOptionsPolicyVerify(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/emailpki.options.policy_verify", val, options...)
}

// GetEndpointAacAllowedNetworks gets the endpoint.aac.allowed_networks value from the UTM
func GetEndpointAacAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/endpoint.aac.allowed_networks", &val, options...)
	return
}

// UpdateEndpointAacAllowedNetworks PUTs the endpoint.aac.allowed_networks value to the UTM
func UpdateEndpointAacAllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/endpoint.aac.allowed_networks", val, options...)
}

// GetEndpointAacAllowedUsers gets the endpoint.aac.allowed_users value from the UTM
func GetEndpointAacAllowedUsers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/endpoint.aac.allowed_users", &val, options...)
	return
}

// UpdateEndpointAacAllowedUsers PUTs the endpoint.aac.allowed_users value to the UTM
func UpdateEndpointAacAllowedUsers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/endpoint.aac.allowed_users", val, options...)
}

// GetEndpointAacCa gets the endpoint.aac.ca value from the UTM
func GetEndpointAacCa(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/endpoint.aac.ca", &val, options...)
	return
}

// UpdateEndpointAacCa PUTs the endpoint.aac.ca value to the UTM
func UpdateEndpointAacCa(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/endpoint.aac.ca", val, options...)
}

// GetEndpointAacCert gets the endpoint.aac.cert value from the UTM
func GetEndpointAacCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/endpoint.aac.cert", &val, options...)
	return
}

// UpdateEndpointAacCert PUTs the endpoint.aac.cert value to the UTM
func UpdateEndpointAacCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/endpoint.aac.cert", val, options...)
}

// GetEndpointAacMagicIp gets the endpoint.aac.magic_ip value from the UTM
func GetEndpointAacMagicIp(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/endpoint.aac.magic_ip", &val, options...)
	return
}

// UpdateEndpointAacMagicIp PUTs the endpoint.aac.magic_ip value to the UTM
func UpdateEndpointAacMagicIp(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/endpoint.aac.magic_ip", val, options...)
}

// GetEndpointAacMaxUserLogins gets the endpoint.aac.max_user_logins value from the UTM
func GetEndpointAacMaxUserLogins(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/endpoint.aac.max_user_logins", &val, options...)
	return
}

// UpdateEndpointAacMaxUserLogins PUTs the endpoint.aac.max_user_logins value to the UTM
func UpdateEndpointAacMaxUserLogins(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/endpoint.aac.max_user_logins", val, options...)
}

// GetEndpointAacStatus gets the endpoint.aac.status value from the UTM
func GetEndpointAacStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/endpoint.aac.status", &val, options...)
	return
}

// UpdateEndpointAacStatus PUTs the endpoint.aac.status value to the UTM
func UpdateEndpointAacStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/endpoint.aac.status", val, options...)
}

// GetEndpointStasCollectors gets the endpoint.stas.collectors value from the UTM
func GetEndpointStasCollectors(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/endpoint.stas.collectors", &val, options...)
	return
}

// UpdateEndpointStasCollectors PUTs the endpoint.stas.collectors value to the UTM
func UpdateEndpointStasCollectors(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/endpoint.stas.collectors", val, options...)
}

// GetEndpointStasStatus gets the endpoint.stas.status value from the UTM
func GetEndpointStasStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/endpoint.stas.status", &val, options...)
	return
}

// UpdateEndpointStasStatus PUTs the endpoint.stas.status value to the UTM
func UpdateEndpointStasStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/endpoint.stas.status", val, options...)
}

// GetEnduserMessagesCompanyLogo gets the enduser_messages.company_logo value from the UTM
func GetEnduserMessagesCompanyLogo(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.company_logo", &val, options...)
	return
}

// UpdateEnduserMessagesCompanyLogo PUTs the enduser_messages.company_logo value to the UTM
func UpdateEnduserMessagesCompanyLogo(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.company_logo", val, options...)
}

// GetEnduserMessagesCompanyText gets the enduser_messages.company_text value from the UTM
func GetEnduserMessagesCompanyText(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.company_text", &val, options...)
	return
}

// UpdateEnduserMessagesCompanyText PUTs the enduser_messages.company_text value to the UTM
func UpdateEnduserMessagesCompanyText(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.company_text", val, options...)
}

// GetEnduserMessagesDlpBlackholePart gets the enduser_messages.dlp.blackhole_part value from the UTM
func GetEnduserMessagesDlpBlackholePart(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.dlp.blackhole_part", &val, options...)
	return
}

// UpdateEnduserMessagesDlpBlackholePart PUTs the enduser_messages.dlp.blackhole_part value to the UTM
func UpdateEnduserMessagesDlpBlackholePart(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.dlp.blackhole_part", val, options...)
}

// GetEnduserMessagesDlpFooterPart gets the enduser_messages.dlp.footer_part value from the UTM
func GetEnduserMessagesDlpFooterPart(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.dlp.footer_part", &val, options...)
	return
}

// UpdateEnduserMessagesDlpFooterPart PUTs the enduser_messages.dlp.footer_part value to the UTM
func UpdateEnduserMessagesDlpFooterPart(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.dlp.footer_part", val, options...)
}

// GetEnduserMessagesDlpHeaderPart gets the enduser_messages.dlp.header_part value from the UTM
func GetEnduserMessagesDlpHeaderPart(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.dlp.header_part", &val, options...)
	return
}

// UpdateEnduserMessagesDlpHeaderPart PUTs the enduser_messages.dlp.header_part value to the UTM
func UpdateEnduserMessagesDlpHeaderPart(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.dlp.header_part", val, options...)
}

// GetEnduserMessagesDlpOriginalPart gets the enduser_messages.dlp.original_part value from the UTM
func GetEnduserMessagesDlpOriginalPart(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.dlp.original_part", &val, options...)
	return
}

// UpdateEnduserMessagesDlpOriginalPart PUTs the enduser_messages.dlp.original_part value to the UTM
func UpdateEnduserMessagesDlpOriginalPart(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.dlp.original_part", val, options...)
}

// GetEnduserMessagesDlpSpxPart gets the enduser_messages.dlp.spx_part value from the UTM
func GetEnduserMessagesDlpSpxPart(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.dlp.spx_part", &val, options...)
	return
}

// UpdateEnduserMessagesDlpSpxPart PUTs the enduser_messages.dlp.spx_part value to the UTM
func UpdateEnduserMessagesDlpSpxPart(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.dlp.spx_part", val, options...)
}

// GetEnduserMessagesDlpSubject gets the enduser_messages.dlp.subject value from the UTM
func GetEnduserMessagesDlpSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.dlp.subject", &val, options...)
	return
}

// UpdateEnduserMessagesDlpSubject PUTs the enduser_messages.dlp.subject value to the UTM
func UpdateEnduserMessagesDlpSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.dlp.subject", val, options...)
}

// GetEnduserMessagesHttpAppDesc gets the enduser_messages.http.app_desc value from the UTM
func GetEnduserMessagesHttpAppDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.app_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpAppDesc PUTs the enduser_messages.http.app_desc value to the UTM
func UpdateEnduserMessagesHttpAppDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.app_desc", val, options...)
}

// GetEnduserMessagesHttpAppSubject gets the enduser_messages.http.app_subject value from the UTM
func GetEnduserMessagesHttpAppSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.app_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpAppSubject PUTs the enduser_messages.http.app_subject value to the UTM
func UpdateEnduserMessagesHttpAppSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.app_subject", val, options...)
}

// GetEnduserMessagesHttpBlacklistDesc gets the enduser_messages.http.blacklist_desc value from the UTM
func GetEnduserMessagesHttpBlacklistDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.blacklist_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpBlacklistDesc PUTs the enduser_messages.http.blacklist_desc value to the UTM
func UpdateEnduserMessagesHttpBlacklistDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.blacklist_desc", val, options...)
}

// GetEnduserMessagesHttpBlacklistSubject gets the enduser_messages.http.blacklist_subject value from the UTM
func GetEnduserMessagesHttpBlacklistSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.blacklist_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpBlacklistSubject PUTs the enduser_messages.http.blacklist_subject value to the UTM
func UpdateEnduserMessagesHttpBlacklistSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.blacklist_subject", val, options...)
}

// GetEnduserMessagesHttpCertfailSubject gets the enduser_messages.http.certfail_subject value from the UTM
func GetEnduserMessagesHttpCertfailSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.certfail_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpCertfailSubject PUTs the enduser_messages.http.certfail_subject value to the UTM
func UpdateEnduserMessagesHttpCertfailSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.certfail_subject", val, options...)
}

// GetEnduserMessagesHttpCffOverrideDesc gets the enduser_messages.http.cff_override_desc value from the UTM
func GetEnduserMessagesHttpCffOverrideDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.cff_override_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpCffOverrideDesc PUTs the enduser_messages.http.cff_override_desc value to the UTM
func UpdateEnduserMessagesHttpCffOverrideDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.cff_override_desc", val, options...)
}

// GetEnduserMessagesHttpCffOverrideSubject gets the enduser_messages.http.cff_override_subject value from the UTM
func GetEnduserMessagesHttpCffOverrideSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.cff_override_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpCffOverrideSubject PUTs the enduser_messages.http.cff_override_subject value to the UTM
func UpdateEnduserMessagesHttpCffOverrideSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.cff_override_subject", val, options...)
}

// GetEnduserMessagesHttpCffOverrideTerms gets the enduser_messages.http.cff_override_terms value from the UTM
func GetEnduserMessagesHttpCffOverrideTerms(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.cff_override_terms", &val, options...)
	return
}

// UpdateEnduserMessagesHttpCffOverrideTerms PUTs the enduser_messages.http.cff_override_terms value to the UTM
func UpdateEnduserMessagesHttpCffOverrideTerms(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.cff_override_terms", val, options...)
}

// GetEnduserMessagesHttpDownloadCompleteDesc gets the enduser_messages.http.download_complete_desc value from the UTM
func GetEnduserMessagesHttpDownloadCompleteDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.download_complete_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpDownloadCompleteDesc PUTs the enduser_messages.http.download_complete_desc value to the UTM
func UpdateEnduserMessagesHttpDownloadCompleteDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.download_complete_desc", val, options...)
}

// GetEnduserMessagesHttpDownloadCompleteSubject gets the enduser_messages.http.download_complete_subject value from the UTM
func GetEnduserMessagesHttpDownloadCompleteSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.download_complete_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpDownloadCompleteSubject PUTs the enduser_messages.http.download_complete_subject value to the UTM
func UpdateEnduserMessagesHttpDownloadCompleteSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.download_complete_subject", val, options...)
}

// GetEnduserMessagesHttpDownloadDesc gets the enduser_messages.http.download_desc value from the UTM
func GetEnduserMessagesHttpDownloadDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.download_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpDownloadDesc PUTs the enduser_messages.http.download_desc value to the UTM
func UpdateEnduserMessagesHttpDownloadDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.download_desc", val, options...)
}

// GetEnduserMessagesHttpDownloadSubject gets the enduser_messages.http.download_subject value from the UTM
func GetEnduserMessagesHttpDownloadSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.download_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpDownloadSubject PUTs the enduser_messages.http.download_subject value to the UTM
func UpdateEnduserMessagesHttpDownloadSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.download_subject", val, options...)
}

// GetEnduserMessagesHttpErrorDesc gets the enduser_messages.http.error_desc value from the UTM
func GetEnduserMessagesHttpErrorDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.error_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpErrorDesc PUTs the enduser_messages.http.error_desc value to the UTM
func UpdateEnduserMessagesHttpErrorDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.error_desc", val, options...)
}

// GetEnduserMessagesHttpErrorSubject gets the enduser_messages.http.error_subject value from the UTM
func GetEnduserMessagesHttpErrorSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.error_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpErrorSubject PUTs the enduser_messages.http.error_subject value to the UTM
func UpdateEnduserMessagesHttpErrorSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.error_subject", val, options...)
}

// GetEnduserMessagesHttpFileextensionDesc gets the enduser_messages.http.fileextension_desc value from the UTM
func GetEnduserMessagesHttpFileextensionDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.fileextension_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpFileextensionDesc PUTs the enduser_messages.http.fileextension_desc value to the UTM
func UpdateEnduserMessagesHttpFileextensionDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.fileextension_desc", val, options...)
}

// GetEnduserMessagesHttpFileextensionSubject gets the enduser_messages.http.fileextension_subject value from the UTM
func GetEnduserMessagesHttpFileextensionSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.fileextension_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpFileextensionSubject PUTs the enduser_messages.http.fileextension_subject value to the UTM
func UpdateEnduserMessagesHttpFileextensionSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.fileextension_subject", val, options...)
}

// GetEnduserMessagesHttpFileextensionWarnDesc gets the enduser_messages.http.fileextension_warn_desc value from the UTM
func GetEnduserMessagesHttpFileextensionWarnDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.fileextension_warn_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpFileextensionWarnDesc PUTs the enduser_messages.http.fileextension_warn_desc value to the UTM
func UpdateEnduserMessagesHttpFileextensionWarnDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.fileextension_warn_desc", val, options...)
}

// GetEnduserMessagesHttpFileextensionWarnSubject gets the enduser_messages.http.fileextension_warn_subject value from the UTM
func GetEnduserMessagesHttpFileextensionWarnSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.fileextension_warn_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpFileextensionWarnSubject PUTs the enduser_messages.http.fileextension_warn_subject value to the UTM
func UpdateEnduserMessagesHttpFileextensionWarnSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.fileextension_warn_subject", val, options...)
}

// GetEnduserMessagesHttpFilesizeDesc gets the enduser_messages.http.filesize_desc value from the UTM
func GetEnduserMessagesHttpFilesizeDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.filesize_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpFilesizeDesc PUTs the enduser_messages.http.filesize_desc value to the UTM
func UpdateEnduserMessagesHttpFilesizeDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.filesize_desc", val, options...)
}

// GetEnduserMessagesHttpFilesizeSubject gets the enduser_messages.http.filesize_subject value from the UTM
func GetEnduserMessagesHttpFilesizeSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.filesize_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpFilesizeSubject PUTs the enduser_messages.http.filesize_subject value to the UTM
func UpdateEnduserMessagesHttpFilesizeSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.filesize_subject", val, options...)
}

// GetEnduserMessagesHttpGeoipDesc gets the enduser_messages.http.geoip_desc value from the UTM
func GetEnduserMessagesHttpGeoipDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.geoip_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpGeoipDesc PUTs the enduser_messages.http.geoip_desc value to the UTM
func UpdateEnduserMessagesHttpGeoipDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.geoip_desc", val, options...)
}

// GetEnduserMessagesHttpGeoipSubject gets the enduser_messages.http.geoip_subject value from the UTM
func GetEnduserMessagesHttpGeoipSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.geoip_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpGeoipSubject PUTs the enduser_messages.http.geoip_subject value to the UTM
func UpdateEnduserMessagesHttpGeoipSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.geoip_subject", val, options...)
}

// GetEnduserMessagesHttpMimetypeDesc gets the enduser_messages.http.mimetype_desc value from the UTM
func GetEnduserMessagesHttpMimetypeDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.mimetype_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpMimetypeDesc PUTs the enduser_messages.http.mimetype_desc value to the UTM
func UpdateEnduserMessagesHttpMimetypeDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.mimetype_desc", val, options...)
}

// GetEnduserMessagesHttpMimetypeSubject gets the enduser_messages.http.mimetype_subject value from the UTM
func GetEnduserMessagesHttpMimetypeSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.mimetype_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpMimetypeSubject PUTs the enduser_messages.http.mimetype_subject value to the UTM
func UpdateEnduserMessagesHttpMimetypeSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.mimetype_subject", val, options...)
}

// GetEnduserMessagesHttpMimetypeWarnDesc gets the enduser_messages.http.mimetype_warn_desc value from the UTM
func GetEnduserMessagesHttpMimetypeWarnDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.mimetype_warn_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpMimetypeWarnDesc PUTs the enduser_messages.http.mimetype_warn_desc value to the UTM
func UpdateEnduserMessagesHttpMimetypeWarnDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.mimetype_warn_desc", val, options...)
}

// GetEnduserMessagesHttpMimetypeWarnSubject gets the enduser_messages.http.mimetype_warn_subject value from the UTM
func GetEnduserMessagesHttpMimetypeWarnSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.mimetype_warn_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpMimetypeWarnSubject PUTs the enduser_messages.http.mimetype_warn_subject value to the UTM
func UpdateEnduserMessagesHttpMimetypeWarnSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.mimetype_warn_subject", val, options...)
}

// GetEnduserMessagesHttpPuaDesc gets the enduser_messages.http.pua_desc value from the UTM
func GetEnduserMessagesHttpPuaDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.pua_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpPuaDesc PUTs the enduser_messages.http.pua_desc value to the UTM
func UpdateEnduserMessagesHttpPuaDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.pua_desc", val, options...)
}

// GetEnduserMessagesHttpPuaSubject gets the enduser_messages.http.pua_subject value from the UTM
func GetEnduserMessagesHttpPuaSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.pua_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpPuaSubject PUTs the enduser_messages.http.pua_subject value to the UTM
func UpdateEnduserMessagesHttpPuaSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.pua_subject", val, options...)
}

// GetEnduserMessagesHttpQuotaBlockDesc gets the enduser_messages.http.quota_block_desc value from the UTM
func GetEnduserMessagesHttpQuotaBlockDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.quota_block_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpQuotaBlockDesc PUTs the enduser_messages.http.quota_block_desc value to the UTM
func UpdateEnduserMessagesHttpQuotaBlockDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.quota_block_desc", val, options...)
}

// GetEnduserMessagesHttpQuotaBlockSubject gets the enduser_messages.http.quota_block_subject value from the UTM
func GetEnduserMessagesHttpQuotaBlockSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.quota_block_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpQuotaBlockSubject PUTs the enduser_messages.http.quota_block_subject value to the UTM
func UpdateEnduserMessagesHttpQuotaBlockSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.quota_block_subject", val, options...)
}

// GetEnduserMessagesHttpQuotaWarnDesc gets the enduser_messages.http.quota_warn_desc value from the UTM
func GetEnduserMessagesHttpQuotaWarnDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.quota_warn_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpQuotaWarnDesc PUTs the enduser_messages.http.quota_warn_desc value to the UTM
func UpdateEnduserMessagesHttpQuotaWarnDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.quota_warn_desc", val, options...)
}

// GetEnduserMessagesHttpQuotaWarnSubject gets the enduser_messages.http.quota_warn_subject value from the UTM
func GetEnduserMessagesHttpQuotaWarnSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.quota_warn_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpQuotaWarnSubject PUTs the enduser_messages.http.quota_warn_subject value to the UTM
func UpdateEnduserMessagesHttpQuotaWarnSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.quota_warn_subject", val, options...)
}

// GetEnduserMessagesHttpSpDesc gets the enduser_messages.http.sp_desc value from the UTM
func GetEnduserMessagesHttpSpDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.sp_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSpDesc PUTs the enduser_messages.http.sp_desc value to the UTM
func UpdateEnduserMessagesHttpSpDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.sp_desc", val, options...)
}

// GetEnduserMessagesHttpSpFrameSubject gets the enduser_messages.http.sp_frame_subject value from the UTM
func GetEnduserMessagesHttpSpFrameSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.sp_frame_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSpFrameSubject PUTs the enduser_messages.http.sp_frame_subject value to the UTM
func UpdateEnduserMessagesHttpSpFrameSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.sp_frame_subject", val, options...)
}

// GetEnduserMessagesHttpSpSubject gets the enduser_messages.http.sp_subject value from the UTM
func GetEnduserMessagesHttpSpSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.sp_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSpSubject PUTs the enduser_messages.http.sp_subject value to the UTM
func UpdateEnduserMessagesHttpSpSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.sp_subject", val, options...)
}

// GetEnduserMessagesHttpSpWarnDesc gets the enduser_messages.http.sp_warn_desc value from the UTM
func GetEnduserMessagesHttpSpWarnDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.sp_warn_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSpWarnDesc PUTs the enduser_messages.http.sp_warn_desc value to the UTM
func UpdateEnduserMessagesHttpSpWarnDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.sp_warn_desc", val, options...)
}

// GetEnduserMessagesHttpSpWarnSubject gets the enduser_messages.http.sp_warn_subject value from the UTM
func GetEnduserMessagesHttpSpWarnSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.sp_warn_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSpWarnSubject PUTs the enduser_messages.http.sp_warn_subject value to the UTM
func UpdateEnduserMessagesHttpSpWarnSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.sp_warn_subject", val, options...)
}

// GetEnduserMessagesHttpSslCertraw gets the enduser_messages.http.ssl_certraw value from the UTM
func GetEnduserMessagesHttpSslCertraw(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.ssl_certraw", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSslCertraw PUTs the enduser_messages.http.ssl_certraw value to the UTM
func UpdateEnduserMessagesHttpSslCertraw(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.ssl_certraw", val, options...)
}

// GetEnduserMessagesHttpSslCertstatus gets the enduser_messages.http.ssl_certstatus value from the UTM
func GetEnduserMessagesHttpSslCertstatus(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.ssl_certstatus", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSslCertstatus PUTs the enduser_messages.http.ssl_certstatus value to the UTM
func UpdateEnduserMessagesHttpSslCertstatus(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.ssl_certstatus", val, options...)
}

// GetEnduserMessagesHttpSslIssuer gets the enduser_messages.http.ssl_issuer value from the UTM
func GetEnduserMessagesHttpSslIssuer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.ssl_issuer", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSslIssuer PUTs the enduser_messages.http.ssl_issuer value to the UTM
func UpdateEnduserMessagesHttpSslIssuer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.ssl_issuer", val, options...)
}

// GetEnduserMessagesHttpSslMd5Fp gets the enduser_messages.http.ssl_md5fp value from the UTM
func GetEnduserMessagesHttpSslMd5Fp(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.ssl_md5fp", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSslMd5Fp PUTs the enduser_messages.http.ssl_md5fp value to the UTM
func UpdateEnduserMessagesHttpSslMd5Fp(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.ssl_md5fp", val, options...)
}

// GetEnduserMessagesHttpSslSha1Fp gets the enduser_messages.http.ssl_sha1fp value from the UTM
func GetEnduserMessagesHttpSslSha1Fp(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.ssl_sha1fp", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSslSha1Fp PUTs the enduser_messages.http.ssl_sha1fp value to the UTM
func UpdateEnduserMessagesHttpSslSha1Fp(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.ssl_sha1fp", val, options...)
}

// GetEnduserMessagesHttpSslSubject gets the enduser_messages.http.ssl_subject value from the UTM
func GetEnduserMessagesHttpSslSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.ssl_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSslSubject PUTs the enduser_messages.http.ssl_subject value to the UTM
func UpdateEnduserMessagesHttpSslSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.ssl_subject", val, options...)
}

// GetEnduserMessagesHttpSslValidfrom gets the enduser_messages.http.ssl_validfrom value from the UTM
func GetEnduserMessagesHttpSslValidfrom(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.ssl_validfrom", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSslValidfrom PUTs the enduser_messages.http.ssl_validfrom value to the UTM
func UpdateEnduserMessagesHttpSslValidfrom(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.ssl_validfrom", val, options...)
}

// GetEnduserMessagesHttpSslValiduntil gets the enduser_messages.http.ssl_validuntil value from the UTM
func GetEnduserMessagesHttpSslValiduntil(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.ssl_validuntil", &val, options...)
	return
}

// UpdateEnduserMessagesHttpSslValiduntil PUTs the enduser_messages.http.ssl_validuntil value to the UTM
func UpdateEnduserMessagesHttpSslValiduntil(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.ssl_validuntil", val, options...)
}

// GetEnduserMessagesHttpThreatDesc gets the enduser_messages.http.threat_desc value from the UTM
func GetEnduserMessagesHttpThreatDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.threat_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpThreatDesc PUTs the enduser_messages.http.threat_desc value to the UTM
func UpdateEnduserMessagesHttpThreatDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.threat_desc", val, options...)
}

// GetEnduserMessagesHttpThreatSubject gets the enduser_messages.http.threat_subject value from the UTM
func GetEnduserMessagesHttpThreatSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.threat_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpThreatSubject PUTs the enduser_messages.http.threat_subject value to the UTM
func UpdateEnduserMessagesHttpThreatSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.threat_subject", val, options...)
}

// GetEnduserMessagesHttpTransparentAuthDesc gets the enduser_messages.http.transparent_auth_desc value from the UTM
func GetEnduserMessagesHttpTransparentAuthDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.transparent_auth_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpTransparentAuthDesc PUTs the enduser_messages.http.transparent_auth_desc value to the UTM
func UpdateEnduserMessagesHttpTransparentAuthDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.transparent_auth_desc", val, options...)
}

// GetEnduserMessagesHttpTransparentAuthSubject gets the enduser_messages.http.transparent_auth_subject value from the UTM
func GetEnduserMessagesHttpTransparentAuthSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.transparent_auth_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpTransparentAuthSubject PUTs the enduser_messages.http.transparent_auth_subject value to the UTM
func UpdateEnduserMessagesHttpTransparentAuthSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.transparent_auth_subject", val, options...)
}

// GetEnduserMessagesHttpTransparentAuthTerms gets the enduser_messages.http.transparent_auth_terms value from the UTM
func GetEnduserMessagesHttpTransparentAuthTerms(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.transparent_auth_terms", &val, options...)
	return
}

// UpdateEnduserMessagesHttpTransparentAuthTerms PUTs the enduser_messages.http.transparent_auth_terms value to the UTM
func UpdateEnduserMessagesHttpTransparentAuthTerms(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.transparent_auth_terms", val, options...)
}

// GetEnduserMessagesHttpVirusDesc gets the enduser_messages.http.virus_desc value from the UTM
func GetEnduserMessagesHttpVirusDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.virus_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpVirusDesc PUTs the enduser_messages.http.virus_desc value to the UTM
func UpdateEnduserMessagesHttpVirusDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.virus_desc", val, options...)
}

// GetEnduserMessagesHttpVirusSubject gets the enduser_messages.http.virus_subject value from the UTM
func GetEnduserMessagesHttpVirusSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.virus_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpVirusSubject PUTs the enduser_messages.http.virus_subject value to the UTM
func UpdateEnduserMessagesHttpVirusSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.virus_subject", val, options...)
}

// GetEnduserMessagesHttpVirusscanDesc gets the enduser_messages.http.virusscan_desc value from the UTM
func GetEnduserMessagesHttpVirusscanDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.virusscan_desc", &val, options...)
	return
}

// UpdateEnduserMessagesHttpVirusscanDesc PUTs the enduser_messages.http.virusscan_desc value to the UTM
func UpdateEnduserMessagesHttpVirusscanDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.virusscan_desc", val, options...)
}

// GetEnduserMessagesHttpVirusscanSubject gets the enduser_messages.http.virusscan_subject value from the UTM
func GetEnduserMessagesHttpVirusscanSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.http.virusscan_subject", &val, options...)
	return
}

// UpdateEnduserMessagesHttpVirusscanSubject PUTs the enduser_messages.http.virusscan_subject value to the UTM
func UpdateEnduserMessagesHttpVirusscanSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.http.virusscan_subject", val, options...)
}

// GetEnduserMessagesMailReleaseErrDesc gets the enduser_messages.mail.release_err_desc value from the UTM
func GetEnduserMessagesMailReleaseErrDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.mail.release_err_desc", &val, options...)
	return
}

// UpdateEnduserMessagesMailReleaseErrDesc PUTs the enduser_messages.mail.release_err_desc value to the UTM
func UpdateEnduserMessagesMailReleaseErrDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.mail.release_err_desc", val, options...)
}

// GetEnduserMessagesMailReleaseErrSubject gets the enduser_messages.mail.release_err_subject value from the UTM
func GetEnduserMessagesMailReleaseErrSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.mail.release_err_subject", &val, options...)
	return
}

// UpdateEnduserMessagesMailReleaseErrSubject PUTs the enduser_messages.mail.release_err_subject value to the UTM
func UpdateEnduserMessagesMailReleaseErrSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.mail.release_err_subject", val, options...)
}

// GetEnduserMessagesMailReleasedDesc gets the enduser_messages.mail.released_desc value from the UTM
func GetEnduserMessagesMailReleasedDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.mail.released_desc", &val, options...)
	return
}

// UpdateEnduserMessagesMailReleasedDesc PUTs the enduser_messages.mail.released_desc value to the UTM
func UpdateEnduserMessagesMailReleasedDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.mail.released_desc", val, options...)
}

// GetEnduserMessagesMailReleasedSubject gets the enduser_messages.mail.released_subject value from the UTM
func GetEnduserMessagesMailReleasedSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.mail.released_subject", &val, options...)
	return
}

// UpdateEnduserMessagesMailReleasedSubject PUTs the enduser_messages.mail.released_subject value to the UTM
func UpdateEnduserMessagesMailReleasedSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.mail.released_subject", val, options...)
}

// GetEnduserMessagesPop3BlockedDesc gets the enduser_messages.pop3.blocked_desc value from the UTM
func GetEnduserMessagesPop3BlockedDesc(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.pop3.blocked_desc", &val, options...)
	return
}

// UpdateEnduserMessagesPop3BlockedDesc PUTs the enduser_messages.pop3.blocked_desc value to the UTM
func UpdateEnduserMessagesPop3BlockedDesc(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.pop3.blocked_desc", val, options...)
}

// GetEnduserMessagesPop3BlockedSubject gets the enduser_messages.pop3.blocked_subject value from the UTM
func GetEnduserMessagesPop3BlockedSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.pop3.blocked_subject", &val, options...)
	return
}

// UpdateEnduserMessagesPop3BlockedSubject PUTs the enduser_messages.pop3.blocked_subject value to the UTM
func UpdateEnduserMessagesPop3BlockedSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.pop3.blocked_subject", val, options...)
}

// GetEnduserMessagesSpxInternalErrorBody gets the enduser_messages.spx.internal_error.body value from the UTM
func GetEnduserMessagesSpxInternalErrorBody(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.spx.internal_error.body", &val, options...)
	return
}

// UpdateEnduserMessagesSpxInternalErrorBody PUTs the enduser_messages.spx.internal_error.body value to the UTM
func UpdateEnduserMessagesSpxInternalErrorBody(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.spx.internal_error.body", val, options...)
}

// GetEnduserMessagesSpxInternalErrorSubject gets the enduser_messages.spx.internal_error.subject value from the UTM
func GetEnduserMessagesSpxInternalErrorSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.spx.internal_error.subject", &val, options...)
	return
}

// UpdateEnduserMessagesSpxInternalErrorSubject PUTs the enduser_messages.spx.internal_error.subject value to the UTM
func UpdateEnduserMessagesSpxInternalErrorSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.spx.internal_error.subject", val, options...)
}

// GetEnduserMessagesSpxInternalErrorSenderBody gets the enduser_messages.spx.internal_error_sender.body value from the UTM
func GetEnduserMessagesSpxInternalErrorSenderBody(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.spx.internal_error_sender.body", &val, options...)
	return
}

// UpdateEnduserMessagesSpxInternalErrorSenderBody PUTs the enduser_messages.spx.internal_error_sender.body value to the UTM
func UpdateEnduserMessagesSpxInternalErrorSenderBody(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.spx.internal_error_sender.body", val, options...)
}

// GetEnduserMessagesSpxInternalErrorSenderSubject gets the enduser_messages.spx.internal_error_sender.subject value from the UTM
func GetEnduserMessagesSpxInternalErrorSenderSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.spx.internal_error_sender.subject", &val, options...)
	return
}

// UpdateEnduserMessagesSpxInternalErrorSenderSubject PUTs the enduser_messages.spx.internal_error_sender.subject value to the UTM
func UpdateEnduserMessagesSpxInternalErrorSenderSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.spx.internal_error_sender.subject", val, options...)
}

// GetEnduserMessagesSpxPasswordNoSpecCharsBody gets the enduser_messages.spx.password_no_spec_chars.body value from the UTM
func GetEnduserMessagesSpxPasswordNoSpecCharsBody(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.spx.password_no_spec_chars.body", &val, options...)
	return
}

// UpdateEnduserMessagesSpxPasswordNoSpecCharsBody PUTs the enduser_messages.spx.password_no_spec_chars.body value to the UTM
func UpdateEnduserMessagesSpxPasswordNoSpecCharsBody(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.spx.password_no_spec_chars.body", val, options...)
}

// GetEnduserMessagesSpxPasswordNoSpecCharsSubject gets the enduser_messages.spx.password_no_spec_chars.subject value from the UTM
func GetEnduserMessagesSpxPasswordNoSpecCharsSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.spx.password_no_spec_chars.subject", &val, options...)
	return
}

// UpdateEnduserMessagesSpxPasswordNoSpecCharsSubject PUTs the enduser_messages.spx.password_no_spec_chars.subject value to the UTM
func UpdateEnduserMessagesSpxPasswordNoSpecCharsSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.spx.password_no_spec_chars.subject", val, options...)
}

// GetEnduserMessagesSpxPasswordNotLongEnoughBody gets the enduser_messages.spx.password_not_long_enough.body value from the UTM
func GetEnduserMessagesSpxPasswordNotLongEnoughBody(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.spx.password_not_long_enough.body", &val, options...)
	return
}

// UpdateEnduserMessagesSpxPasswordNotLongEnoughBody PUTs the enduser_messages.spx.password_not_long_enough.body value to the UTM
func UpdateEnduserMessagesSpxPasswordNotLongEnoughBody(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.spx.password_not_long_enough.body", val, options...)
}

// GetEnduserMessagesSpxPasswordNotLongEnoughSubject gets the enduser_messages.spx.password_not_long_enough.subject value from the UTM
func GetEnduserMessagesSpxPasswordNotLongEnoughSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.spx.password_not_long_enough.subject", &val, options...)
	return
}

// UpdateEnduserMessagesSpxPasswordNotLongEnoughSubject PUTs the enduser_messages.spx.password_not_long_enough.subject value to the UTM
func UpdateEnduserMessagesSpxPasswordNotLongEnoughSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.spx.password_not_long_enough.subject", val, options...)
}

// GetEnduserMessagesSpxPasswordNotPresentedBody gets the enduser_messages.spx.password_not_presented.body value from the UTM
func GetEnduserMessagesSpxPasswordNotPresentedBody(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.spx.password_not_presented.body", &val, options...)
	return
}

// UpdateEnduserMessagesSpxPasswordNotPresentedBody PUTs the enduser_messages.spx.password_not_presented.body value to the UTM
func UpdateEnduserMessagesSpxPasswordNotPresentedBody(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.spx.password_not_presented.body", val, options...)
}

// GetEnduserMessagesSpxPasswordNotPresentedSubject gets the enduser_messages.spx.password_not_presented.subject value from the UTM
func GetEnduserMessagesSpxPasswordNotPresentedSubject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.spx.password_not_presented.subject", &val, options...)
	return
}

// UpdateEnduserMessagesSpxPasswordNotPresentedSubject PUTs the enduser_messages.spx.password_not_presented.subject value to the UTM
func UpdateEnduserMessagesSpxPasswordNotPresentedSubject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.spx.password_not_presented.subject", val, options...)
}

// GetEnduserMessagesSpxUrlNotFoundMessage gets the enduser_messages.spx.url_not_found.message value from the UTM
func GetEnduserMessagesSpxUrlNotFoundMessage(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.spx.url_not_found.message", &val, options...)
	return
}

// UpdateEnduserMessagesSpxUrlNotFoundMessage PUTs the enduser_messages.spx.url_not_found.message value to the UTM
func UpdateEnduserMessagesSpxUrlNotFoundMessage(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.spx.url_not_found.message", val, options...)
}

// GetEnduserMessagesSquidCacheAdmin gets the enduser_messages.squid.cache_admin value from the UTM
func GetEnduserMessagesSquidCacheAdmin(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.squid.cache_admin", &val, options...)
	return
}

// UpdateEnduserMessagesSquidCacheAdmin PUTs the enduser_messages.squid.cache_admin value to the UTM
func UpdateEnduserMessagesSquidCacheAdmin(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.squid.cache_admin", val, options...)
}

// GetEnduserMessagesSquidCacheAdminMessage gets the enduser_messages.squid.cache_admin_message value from the UTM
func GetEnduserMessagesSquidCacheAdminMessage(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/enduser_messages.squid.cache_admin_message", &val, options...)
	return
}

// UpdateEnduserMessagesSquidCacheAdminMessage PUTs the enduser_messages.squid.cache_admin_message value to the UTM
func UpdateEnduserMessagesSquidCacheAdminMessage(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/enduser_messages.squid.cache_admin_message", val, options...)
}

// GetEppAllowedNetworks gets the epp.allowed_networks value from the UTM
func GetEppAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/epp.allowed_networks", &val, options...)
	return
}

// UpdateEppAllowedNetworks PUTs the epp.allowed_networks value to the UTM
func UpdateEppAllowedNetworks(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.allowed_networks", val, options...)
}

// GetEppCertificate gets the epp.certificate value from the UTM
func GetEppCertificate(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.certificate", &val, options...)
	return
}

// UpdateEppCertificate PUTs the epp.certificate value to the UTM
func UpdateEppCertificate(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.certificate", val, options...)
}

// GetEppCity gets the epp.city value from the UTM
func GetEppCity(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.city", &val, options...)
	return
}

// UpdateEppCity PUTs the epp.city value to the UTM
func UpdateEppCity(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.city", val, options...)
}

// GetEppCountry gets the epp.country value from the UTM
func GetEppCountry(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.country", &val, options...)
	return
}

// UpdateEppCountry PUTs the epp.country value to the UTM
func UpdateEppCountry(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.country", val, options...)
}

// GetEppDefaultEndpointsGroup gets the epp.default_endpoints_group value from the UTM
func GetEppDefaultEndpointsGroup(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.default_endpoints_group", &val, options...)
	return
}

// UpdateEppDefaultEndpointsGroup PUTs the epp.default_endpoints_group value to the UTM
func UpdateEppDefaultEndpointsGroup(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.default_endpoints_group", val, options...)
}

// GetEppDevices gets the epp.devices value from the UTM
func GetEppDevices(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/epp.devices", &val, options...)
	return
}

// UpdateEppDevices PUTs the epp.devices value to the UTM
func UpdateEppDevices(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.devices", val, options...)
}

// GetEppEmail gets the epp.email value from the UTM
func GetEppEmail(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.email", &val, options...)
	return
}

// UpdateEppEmail PUTs the epp.email value to the UTM
func UpdateEppEmail(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.email", val, options...)
}

// GetEppEndpoints gets the epp.endpoints value from the UTM
func GetEppEndpoints(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/epp.endpoints", &val, options...)
	return
}

// UpdateEppEndpoints PUTs the epp.endpoints value to the UTM
func UpdateEppEndpoints(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.endpoints", val, options...)
}

// GetEppEndpointsGroups gets the epp.endpoints_groups value from the UTM
func GetEppEndpointsGroups(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/epp.endpoints_groups", &val, options...)
	return
}

// UpdateEppEndpointsGroups PUTs the epp.endpoints_groups value to the UTM
func UpdateEppEndpointsGroups(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.endpoints_groups", val, options...)
}

// GetEppExceptionsAv gets the epp.exceptions.av value from the UTM
func GetEppExceptionsAv(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/epp.exceptions.av", &val, options...)
	return
}

// UpdateEppExceptionsAv PUTs the epp.exceptions.av value to the UTM
func UpdateEppExceptionsAv(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.exceptions.av", val, options...)
}

// GetEppExceptionsDc gets the epp.exceptions.dc value from the UTM
func GetEppExceptionsDc(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/epp.exceptions.dc", &val, options...)
	return
}

// UpdateEppExceptionsDc PUTs the epp.exceptions.dc value to the UTM
func UpdateEppExceptionsDc(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.exceptions.dc", val, options...)
}

// GetEppFallbackUrl gets the epp.fallback_url value from the UTM
func GetEppFallbackUrl(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.fallback_url", &val, options...)
	return
}

// UpdateEppFallbackUrl PUTs the epp.fallback_url value to the UTM
func UpdateEppFallbackUrl(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.fallback_url", val, options...)
}

// GetEppMagnetPassword gets the epp.magnet_password value from the UTM
func GetEppMagnetPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.magnet_password", &val, options...)
	return
}

// UpdateEppMagnetPassword PUTs the epp.magnet_password value to the UTM
func UpdateEppMagnetPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.magnet_password", val, options...)
}

// GetEppMagnetUsername gets the epp.magnet_username value from the UTM
func GetEppMagnetUsername(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.magnet_username", &val, options...)
	return
}

// UpdateEppMagnetUsername PUTs the epp.magnet_username value to the UTM
func UpdateEppMagnetUsername(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.magnet_username", val, options...)
}

// GetEppOrganization gets the epp.organization value from the UTM
func GetEppOrganization(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.organization", &val, options...)
	return
}

// UpdateEppOrganization PUTs the epp.organization value to the UTM
func UpdateEppOrganization(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.organization", val, options...)
}

// GetEppParentProxyHost gets the epp.parent_proxy_host value from the UTM
func GetEppParentProxyHost(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.parent_proxy_host", &val, options...)
	return
}

// UpdateEppParentProxyHost PUTs the epp.parent_proxy_host value to the UTM
func UpdateEppParentProxyHost(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.parent_proxy_host", val, options...)
}

// GetEppParentProxyPort gets the epp.parent_proxy_port value from the UTM
func GetEppParentProxyPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/epp.parent_proxy_port", &val, options...)
	return
}

// UpdateEppParentProxyPort PUTs the epp.parent_proxy_port value to the UTM
func UpdateEppParentProxyPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.parent_proxy_port", val, options...)
}

// GetEppParentProxyStatus gets the epp.parent_proxy_status value from the UTM
func GetEppParentProxyStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/epp.parent_proxy_status", &val, options...)
	return
}

// UpdateEppParentProxyStatus PUTs the epp.parent_proxy_status value to the UTM
func UpdateEppParentProxyStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.parent_proxy_status", val, options...)
}

// GetEppPoliciesAv gets the epp.policies.av value from the UTM
func GetEppPoliciesAv(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/epp.policies.av", &val, options...)
	return
}

// UpdateEppPoliciesAv PUTs the epp.policies.av value to the UTM
func UpdateEppPoliciesAv(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.policies.av", val, options...)
}

// GetEppPoliciesDc gets the epp.policies.dc value from the UTM
func GetEppPoliciesDc(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/epp.policies.dc", &val, options...)
	return
}

// UpdateEppPoliciesDc PUTs the epp.policies.dc value to the UTM
func UpdateEppPoliciesDc(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.policies.dc", val, options...)
}

// GetEppPort gets the epp.port value from the UTM
func GetEppPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/epp.port", &val, options...)
	return
}

// UpdateEppPort PUTs the epp.port value to the UTM
func UpdateEppPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.port", val, options...)
}

// GetEppPrivateKey gets the epp.private_key value from the UTM
func GetEppPrivateKey(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.private_key", &val, options...)
	return
}

// UpdateEppPrivateKey PUTs the epp.private_key value to the UTM
func UpdateEppPrivateKey(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.private_key", val, options...)
}

// GetEppRegistrationToken gets the epp.registration_token value from the UTM
func GetEppRegistrationToken(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.registration_token", &val, options...)
	return
}

// UpdateEppRegistrationToken PUTs the epp.registration_token value to the UTM
func UpdateEppRegistrationToken(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.registration_token", val, options...)
}

// GetEppStatusAv gets the epp.status.av value from the UTM
func GetEppStatusAv(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/epp.status.av", &val, options...)
	return
}

// UpdateEppStatusAv PUTs the epp.status.av value to the UTM
func UpdateEppStatusAv(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.status.av", val, options...)
}

// GetEppStatusBroker gets the epp.status.broker value from the UTM
func GetEppStatusBroker(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/epp.status.broker", &val, options...)
	return
}

// UpdateEppStatusBroker PUTs the epp.status.broker value to the UTM
func UpdateEppStatusBroker(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.status.broker", val, options...)
}

// GetEppStatusDc gets the epp.status.dc value from the UTM
func GetEppStatusDc(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/epp.status.dc", &val, options...)
	return
}

// UpdateEppStatusDc PUTs the epp.status.dc value to the UTM
func UpdateEppStatusDc(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.status.dc", val, options...)
}

// GetEppStatusEpp gets the epp.status.epp value from the UTM
func GetEppStatusEpp(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/epp.status.epp", &val, options...)
	return
}

// UpdateEppStatusEpp PUTs the epp.status.epp value to the UTM
func UpdateEppStatusEpp(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.status.epp", val, options...)
}

// GetEppStatusWc gets the epp.status.wc value from the UTM
func GetEppStatusWc(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/epp.status.wc", &val, options...)
	return
}

// UpdateEppStatusWc PUTs the epp.status.wc value to the UTM
func UpdateEppStatusWc(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.status.wc", val, options...)
}

// GetEppTamperPassword gets the epp.tamper_password value from the UTM
func GetEppTamperPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.tamper_password", &val, options...)
	return
}

// UpdateEppTamperPassword PUTs the epp.tamper_password value to the UTM
func UpdateEppTamperPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.tamper_password", val, options...)
}

// GetEppVersion gets the epp.version value from the UTM
func GetEppVersion(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.version", &val, options...)
	return
}

// UpdateEppVersion PUTs the epp.version value to the UTM
func UpdateEppVersion(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.version", val, options...)
}

// GetEppWdxToken gets the epp.wdx_token value from the UTM
func GetEppWdxToken(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/epp.wdx_token", &val, options...)
	return
}

// UpdateEppWdxToken PUTs the epp.wdx_token value to the UTM
func UpdateEppWdxToken(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/epp.wdx_token", val, options...)
}

// GetExecutiveReportDailyArchive gets the executive_report.daily.archive value from the UTM
func GetExecutiveReportDailyArchive(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/executive_report.daily.archive", &val, options...)
	return
}

// UpdateExecutiveReportDailyArchive PUTs the executive_report.daily.archive value to the UTM
func UpdateExecutiveReportDailyArchive(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.daily.archive", val, options...)
}

// GetExecutiveReportDailyKeep gets the executive_report.daily.keep value from the UTM
func GetExecutiveReportDailyKeep(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/executive_report.daily.keep", &val, options...)
	return
}

// UpdateExecutiveReportDailyKeep PUTs the executive_report.daily.keep value to the UTM
func UpdateExecutiveReportDailyKeep(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.daily.keep", val, options...)
}

// GetExecutiveReportDailyPdfrecipients gets the executive_report.daily.pdfrecipients value from the UTM
func GetExecutiveReportDailyPdfrecipients(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/executive_report.daily.pdfrecipients", &val, options...)
	return
}

// UpdateExecutiveReportDailyPdfrecipients PUTs the executive_report.daily.pdfrecipients value to the UTM
func UpdateExecutiveReportDailyPdfrecipients(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.daily.pdfrecipients", val, options...)
}

// GetExecutiveReportDailyRecipients gets the executive_report.daily.recipients value from the UTM
func GetExecutiveReportDailyRecipients(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/executive_report.daily.recipients", &val, options...)
	return
}

// UpdateExecutiveReportDailyRecipients PUTs the executive_report.daily.recipients value to the UTM
func UpdateExecutiveReportDailyRecipients(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.daily.recipients", val, options...)
}

// GetExecutiveReportDailyStatus gets the executive_report.daily.status value from the UTM
func GetExecutiveReportDailyStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/executive_report.daily.status", &val, options...)
	return
}

// UpdateExecutiveReportDailyStatus PUTs the executive_report.daily.status value to the UTM
func UpdateExecutiveReportDailyStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.daily.status", val, options...)
}

// GetExecutiveReportMonthlyArchive gets the executive_report.monthly.archive value from the UTM
func GetExecutiveReportMonthlyArchive(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/executive_report.monthly.archive", &val, options...)
	return
}

// UpdateExecutiveReportMonthlyArchive PUTs the executive_report.monthly.archive value to the UTM
func UpdateExecutiveReportMonthlyArchive(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.monthly.archive", val, options...)
}

// GetExecutiveReportMonthlyKeep gets the executive_report.monthly.keep value from the UTM
func GetExecutiveReportMonthlyKeep(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/executive_report.monthly.keep", &val, options...)
	return
}

// UpdateExecutiveReportMonthlyKeep PUTs the executive_report.monthly.keep value to the UTM
func UpdateExecutiveReportMonthlyKeep(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.monthly.keep", val, options...)
}

// GetExecutiveReportMonthlyPdfrecipients gets the executive_report.monthly.pdfrecipients value from the UTM
func GetExecutiveReportMonthlyPdfrecipients(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/executive_report.monthly.pdfrecipients", &val, options...)
	return
}

// UpdateExecutiveReportMonthlyPdfrecipients PUTs the executive_report.monthly.pdfrecipients value to the UTM
func UpdateExecutiveReportMonthlyPdfrecipients(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.monthly.pdfrecipients", val, options...)
}

// GetExecutiveReportMonthlyRecipients gets the executive_report.monthly.recipients value from the UTM
func GetExecutiveReportMonthlyRecipients(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/executive_report.monthly.recipients", &val, options...)
	return
}

// UpdateExecutiveReportMonthlyRecipients PUTs the executive_report.monthly.recipients value to the UTM
func UpdateExecutiveReportMonthlyRecipients(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.monthly.recipients", val, options...)
}

// GetExecutiveReportMonthlyStatus gets the executive_report.monthly.status value from the UTM
func GetExecutiveReportMonthlyStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/executive_report.monthly.status", &val, options...)
	return
}

// UpdateExecutiveReportMonthlyStatus PUTs the executive_report.monthly.status value to the UTM
func UpdateExecutiveReportMonthlyStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.monthly.status", val, options...)
}

// GetExecutiveReportWeeklyArchive gets the executive_report.weekly.archive value from the UTM
func GetExecutiveReportWeeklyArchive(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/executive_report.weekly.archive", &val, options...)
	return
}

// UpdateExecutiveReportWeeklyArchive PUTs the executive_report.weekly.archive value to the UTM
func UpdateExecutiveReportWeeklyArchive(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.weekly.archive", val, options...)
}

// GetExecutiveReportWeeklyFirstDayOfWeek gets the executive_report.weekly.first_day_of_week value from the UTM
func GetExecutiveReportWeeklyFirstDayOfWeek(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/executive_report.weekly.first_day_of_week", &val, options...)
	return
}

// UpdateExecutiveReportWeeklyFirstDayOfWeek PUTs the executive_report.weekly.first_day_of_week value to the UTM
func UpdateExecutiveReportWeeklyFirstDayOfWeek(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.weekly.first_day_of_week", val, options...)
}

// GetExecutiveReportWeeklyKeep gets the executive_report.weekly.keep value from the UTM
func GetExecutiveReportWeeklyKeep(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/executive_report.weekly.keep", &val, options...)
	return
}

// UpdateExecutiveReportWeeklyKeep PUTs the executive_report.weekly.keep value to the UTM
func UpdateExecutiveReportWeeklyKeep(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.weekly.keep", val, options...)
}

// GetExecutiveReportWeeklyPdfrecipients gets the executive_report.weekly.pdfrecipients value from the UTM
func GetExecutiveReportWeeklyPdfrecipients(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/executive_report.weekly.pdfrecipients", &val, options...)
	return
}

// UpdateExecutiveReportWeeklyPdfrecipients PUTs the executive_report.weekly.pdfrecipients value to the UTM
func UpdateExecutiveReportWeeklyPdfrecipients(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.weekly.pdfrecipients", val, options...)
}

// GetExecutiveReportWeeklyRecipients gets the executive_report.weekly.recipients value from the UTM
func GetExecutiveReportWeeklyRecipients(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/executive_report.weekly.recipients", &val, options...)
	return
}

// UpdateExecutiveReportWeeklyRecipients PUTs the executive_report.weekly.recipients value to the UTM
func UpdateExecutiveReportWeeklyRecipients(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.weekly.recipients", val, options...)
}

// GetExecutiveReportWeeklyStatus gets the executive_report.weekly.status value from the UTM
func GetExecutiveReportWeeklyStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/executive_report.weekly.status", &val, options...)
	return
}

// UpdateExecutiveReportWeeklyStatus PUTs the executive_report.weekly.status value to the UTM
func UpdateExecutiveReportWeeklyStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/executive_report.weekly.status", val, options...)
}

// GetFloodProtectionIcmpDstBurst gets the flood_protection.icmp.dst_burst value from the UTM
func GetFloodProtectionIcmpDstBurst(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.dst_burst", &val, options...)
	return
}

// UpdateFloodProtectionIcmpDstBurst PUTs the flood_protection.icmp.dst_burst value to the UTM
func UpdateFloodProtectionIcmpDstBurst(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.dst_burst", val, options...)
}

// GetFloodProtectionIcmpDstExpire gets the flood_protection.icmp.dst_expire value from the UTM
func GetFloodProtectionIcmpDstExpire(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.dst_expire", &val, options...)
	return
}

// UpdateFloodProtectionIcmpDstExpire PUTs the flood_protection.icmp.dst_expire value to the UTM
func UpdateFloodProtectionIcmpDstExpire(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.dst_expire", val, options...)
}

// GetFloodProtectionIcmpDstGcInterval gets the flood_protection.icmp.dst_gc_interval value from the UTM
func GetFloodProtectionIcmpDstGcInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.dst_gc_interval", &val, options...)
	return
}

// UpdateFloodProtectionIcmpDstGcInterval PUTs the flood_protection.icmp.dst_gc_interval value to the UTM
func UpdateFloodProtectionIcmpDstGcInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.dst_gc_interval", val, options...)
}

// GetFloodProtectionIcmpDstRate gets the flood_protection.icmp.dst_rate value from the UTM
func GetFloodProtectionIcmpDstRate(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.dst_rate", &val, options...)
	return
}

// UpdateFloodProtectionIcmpDstRate PUTs the flood_protection.icmp.dst_rate value to the UTM
func UpdateFloodProtectionIcmpDstRate(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.dst_rate", val, options...)
}

// GetFloodProtectionIcmpLog gets the flood_protection.icmp.log value from the UTM
func GetFloodProtectionIcmpLog(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.log", &val, options...)
	return
}

// UpdateFloodProtectionIcmpLog PUTs the flood_protection.icmp.log value to the UTM
func UpdateFloodProtectionIcmpLog(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.log", val, options...)
}

// GetFloodProtectionIcmpLogLimitBurst gets the flood_protection.icmp.log_limit_burst value from the UTM
func GetFloodProtectionIcmpLogLimitBurst(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.log_limit_burst", &val, options...)
	return
}

// UpdateFloodProtectionIcmpLogLimitBurst PUTs the flood_protection.icmp.log_limit_burst value to the UTM
func UpdateFloodProtectionIcmpLogLimitBurst(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.log_limit_burst", val, options...)
}

// GetFloodProtectionIcmpLogLimitRate gets the flood_protection.icmp.log_limit_rate value from the UTM
func GetFloodProtectionIcmpLogLimitRate(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.log_limit_rate", &val, options...)
	return
}

// UpdateFloodProtectionIcmpLogLimitRate PUTs the flood_protection.icmp.log_limit_rate value to the UTM
func UpdateFloodProtectionIcmpLogLimitRate(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.log_limit_rate", val, options...)
}

// GetFloodProtectionIcmpMode gets the flood_protection.icmp.mode value from the UTM
func GetFloodProtectionIcmpMode(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.mode", &val, options...)
	return
}

// UpdateFloodProtectionIcmpMode PUTs the flood_protection.icmp.mode value to the UTM
func UpdateFloodProtectionIcmpMode(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.mode", val, options...)
}

// GetFloodProtectionIcmpSrcBurst gets the flood_protection.icmp.src_burst value from the UTM
func GetFloodProtectionIcmpSrcBurst(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.src_burst", &val, options...)
	return
}

// UpdateFloodProtectionIcmpSrcBurst PUTs the flood_protection.icmp.src_burst value to the UTM
func UpdateFloodProtectionIcmpSrcBurst(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.src_burst", val, options...)
}

// GetFloodProtectionIcmpSrcExpire gets the flood_protection.icmp.src_expire value from the UTM
func GetFloodProtectionIcmpSrcExpire(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.src_expire", &val, options...)
	return
}

// UpdateFloodProtectionIcmpSrcExpire PUTs the flood_protection.icmp.src_expire value to the UTM
func UpdateFloodProtectionIcmpSrcExpire(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.src_expire", val, options...)
}

// GetFloodProtectionIcmpSrcGcInterval gets the flood_protection.icmp.src_gc_interval value from the UTM
func GetFloodProtectionIcmpSrcGcInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.src_gc_interval", &val, options...)
	return
}

// UpdateFloodProtectionIcmpSrcGcInterval PUTs the flood_protection.icmp.src_gc_interval value to the UTM
func UpdateFloodProtectionIcmpSrcGcInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.src_gc_interval", val, options...)
}

// GetFloodProtectionIcmpSrcRate gets the flood_protection.icmp.src_rate value from the UTM
func GetFloodProtectionIcmpSrcRate(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.src_rate", &val, options...)
	return
}

// UpdateFloodProtectionIcmpSrcRate PUTs the flood_protection.icmp.src_rate value to the UTM
func UpdateFloodProtectionIcmpSrcRate(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.src_rate", val, options...)
}

// GetFloodProtectionIcmpStatus gets the flood_protection.icmp.status value from the UTM
func GetFloodProtectionIcmpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/flood_protection.icmp.status", &val, options...)
	return
}

// UpdateFloodProtectionIcmpStatus PUTs the flood_protection.icmp.status value to the UTM
func UpdateFloodProtectionIcmpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.icmp.status", val, options...)
}

// GetFloodProtectionSynDstBurst gets the flood_protection.syn.dst_burst value from the UTM
func GetFloodProtectionSynDstBurst(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.dst_burst", &val, options...)
	return
}

// UpdateFloodProtectionSynDstBurst PUTs the flood_protection.syn.dst_burst value to the UTM
func UpdateFloodProtectionSynDstBurst(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.dst_burst", val, options...)
}

// GetFloodProtectionSynDstExpire gets the flood_protection.syn.dst_expire value from the UTM
func GetFloodProtectionSynDstExpire(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.dst_expire", &val, options...)
	return
}

// UpdateFloodProtectionSynDstExpire PUTs the flood_protection.syn.dst_expire value to the UTM
func UpdateFloodProtectionSynDstExpire(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.dst_expire", val, options...)
}

// GetFloodProtectionSynDstGcInterval gets the flood_protection.syn.dst_gc_interval value from the UTM
func GetFloodProtectionSynDstGcInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.dst_gc_interval", &val, options...)
	return
}

// UpdateFloodProtectionSynDstGcInterval PUTs the flood_protection.syn.dst_gc_interval value to the UTM
func UpdateFloodProtectionSynDstGcInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.dst_gc_interval", val, options...)
}

// GetFloodProtectionSynDstRate gets the flood_protection.syn.dst_rate value from the UTM
func GetFloodProtectionSynDstRate(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.dst_rate", &val, options...)
	return
}

// UpdateFloodProtectionSynDstRate PUTs the flood_protection.syn.dst_rate value to the UTM
func UpdateFloodProtectionSynDstRate(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.dst_rate", val, options...)
}

// GetFloodProtectionSynLog gets the flood_protection.syn.log value from the UTM
func GetFloodProtectionSynLog(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.log", &val, options...)
	return
}

// UpdateFloodProtectionSynLog PUTs the flood_protection.syn.log value to the UTM
func UpdateFloodProtectionSynLog(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.log", val, options...)
}

// GetFloodProtectionSynLogLimitBurst gets the flood_protection.syn.log_limit_burst value from the UTM
func GetFloodProtectionSynLogLimitBurst(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.log_limit_burst", &val, options...)
	return
}

// UpdateFloodProtectionSynLogLimitBurst PUTs the flood_protection.syn.log_limit_burst value to the UTM
func UpdateFloodProtectionSynLogLimitBurst(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.log_limit_burst", val, options...)
}

// GetFloodProtectionSynLogLimitRate gets the flood_protection.syn.log_limit_rate value from the UTM
func GetFloodProtectionSynLogLimitRate(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.log_limit_rate", &val, options...)
	return
}

// UpdateFloodProtectionSynLogLimitRate PUTs the flood_protection.syn.log_limit_rate value to the UTM
func UpdateFloodProtectionSynLogLimitRate(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.log_limit_rate", val, options...)
}

// GetFloodProtectionSynMode gets the flood_protection.syn.mode value from the UTM
func GetFloodProtectionSynMode(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.mode", &val, options...)
	return
}

// UpdateFloodProtectionSynMode PUTs the flood_protection.syn.mode value to the UTM
func UpdateFloodProtectionSynMode(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.mode", val, options...)
}

// GetFloodProtectionSynSrcBurst gets the flood_protection.syn.src_burst value from the UTM
func GetFloodProtectionSynSrcBurst(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.src_burst", &val, options...)
	return
}

// UpdateFloodProtectionSynSrcBurst PUTs the flood_protection.syn.src_burst value to the UTM
func UpdateFloodProtectionSynSrcBurst(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.src_burst", val, options...)
}

// GetFloodProtectionSynSrcExpire gets the flood_protection.syn.src_expire value from the UTM
func GetFloodProtectionSynSrcExpire(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.src_expire", &val, options...)
	return
}

// UpdateFloodProtectionSynSrcExpire PUTs the flood_protection.syn.src_expire value to the UTM
func UpdateFloodProtectionSynSrcExpire(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.src_expire", val, options...)
}

// GetFloodProtectionSynSrcGcInterval gets the flood_protection.syn.src_gc_interval value from the UTM
func GetFloodProtectionSynSrcGcInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.src_gc_interval", &val, options...)
	return
}

// UpdateFloodProtectionSynSrcGcInterval PUTs the flood_protection.syn.src_gc_interval value to the UTM
func UpdateFloodProtectionSynSrcGcInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.src_gc_interval", val, options...)
}

// GetFloodProtectionSynSrcRate gets the flood_protection.syn.src_rate value from the UTM
func GetFloodProtectionSynSrcRate(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.src_rate", &val, options...)
	return
}

// UpdateFloodProtectionSynSrcRate PUTs the flood_protection.syn.src_rate value to the UTM
func UpdateFloodProtectionSynSrcRate(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.src_rate", val, options...)
}

// GetFloodProtectionSynStatus gets the flood_protection.syn.status value from the UTM
func GetFloodProtectionSynStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/flood_protection.syn.status", &val, options...)
	return
}

// UpdateFloodProtectionSynStatus PUTs the flood_protection.syn.status value to the UTM
func UpdateFloodProtectionSynStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.syn.status", val, options...)
}

// GetFloodProtectionUdpDstBurst gets the flood_protection.udp.dst_burst value from the UTM
func GetFloodProtectionUdpDstBurst(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.dst_burst", &val, options...)
	return
}

// UpdateFloodProtectionUdpDstBurst PUTs the flood_protection.udp.dst_burst value to the UTM
func UpdateFloodProtectionUdpDstBurst(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.dst_burst", val, options...)
}

// GetFloodProtectionUdpDstExpire gets the flood_protection.udp.dst_expire value from the UTM
func GetFloodProtectionUdpDstExpire(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.dst_expire", &val, options...)
	return
}

// UpdateFloodProtectionUdpDstExpire PUTs the flood_protection.udp.dst_expire value to the UTM
func UpdateFloodProtectionUdpDstExpire(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.dst_expire", val, options...)
}

// GetFloodProtectionUdpDstGcInterval gets the flood_protection.udp.dst_gc_interval value from the UTM
func GetFloodProtectionUdpDstGcInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.dst_gc_interval", &val, options...)
	return
}

// UpdateFloodProtectionUdpDstGcInterval PUTs the flood_protection.udp.dst_gc_interval value to the UTM
func UpdateFloodProtectionUdpDstGcInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.dst_gc_interval", val, options...)
}

// GetFloodProtectionUdpDstRate gets the flood_protection.udp.dst_rate value from the UTM
func GetFloodProtectionUdpDstRate(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.dst_rate", &val, options...)
	return
}

// UpdateFloodProtectionUdpDstRate PUTs the flood_protection.udp.dst_rate value to the UTM
func UpdateFloodProtectionUdpDstRate(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.dst_rate", val, options...)
}

// GetFloodProtectionUdpLog gets the flood_protection.udp.log value from the UTM
func GetFloodProtectionUdpLog(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.log", &val, options...)
	return
}

// UpdateFloodProtectionUdpLog PUTs the flood_protection.udp.log value to the UTM
func UpdateFloodProtectionUdpLog(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.log", val, options...)
}

// GetFloodProtectionUdpLogLimitBurst gets the flood_protection.udp.log_limit_burst value from the UTM
func GetFloodProtectionUdpLogLimitBurst(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.log_limit_burst", &val, options...)
	return
}

// UpdateFloodProtectionUdpLogLimitBurst PUTs the flood_protection.udp.log_limit_burst value to the UTM
func UpdateFloodProtectionUdpLogLimitBurst(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.log_limit_burst", val, options...)
}

// GetFloodProtectionUdpLogLimitRate gets the flood_protection.udp.log_limit_rate value from the UTM
func GetFloodProtectionUdpLogLimitRate(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.log_limit_rate", &val, options...)
	return
}

// UpdateFloodProtectionUdpLogLimitRate PUTs the flood_protection.udp.log_limit_rate value to the UTM
func UpdateFloodProtectionUdpLogLimitRate(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.log_limit_rate", val, options...)
}

// GetFloodProtectionUdpMode gets the flood_protection.udp.mode value from the UTM
func GetFloodProtectionUdpMode(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.mode", &val, options...)
	return
}

// UpdateFloodProtectionUdpMode PUTs the flood_protection.udp.mode value to the UTM
func UpdateFloodProtectionUdpMode(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.mode", val, options...)
}

// GetFloodProtectionUdpSrcBurst gets the flood_protection.udp.src_burst value from the UTM
func GetFloodProtectionUdpSrcBurst(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.src_burst", &val, options...)
	return
}

// UpdateFloodProtectionUdpSrcBurst PUTs the flood_protection.udp.src_burst value to the UTM
func UpdateFloodProtectionUdpSrcBurst(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.src_burst", val, options...)
}

// GetFloodProtectionUdpSrcExpire gets the flood_protection.udp.src_expire value from the UTM
func GetFloodProtectionUdpSrcExpire(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.src_expire", &val, options...)
	return
}

// UpdateFloodProtectionUdpSrcExpire PUTs the flood_protection.udp.src_expire value to the UTM
func UpdateFloodProtectionUdpSrcExpire(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.src_expire", val, options...)
}

// GetFloodProtectionUdpSrcGcInterval gets the flood_protection.udp.src_gc_interval value from the UTM
func GetFloodProtectionUdpSrcGcInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.src_gc_interval", &val, options...)
	return
}

// UpdateFloodProtectionUdpSrcGcInterval PUTs the flood_protection.udp.src_gc_interval value to the UTM
func UpdateFloodProtectionUdpSrcGcInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.src_gc_interval", val, options...)
}

// GetFloodProtectionUdpSrcRate gets the flood_protection.udp.src_rate value from the UTM
func GetFloodProtectionUdpSrcRate(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.src_rate", &val, options...)
	return
}

// UpdateFloodProtectionUdpSrcRate PUTs the flood_protection.udp.src_rate value to the UTM
func UpdateFloodProtectionUdpSrcRate(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.src_rate", val, options...)
}

// GetFloodProtectionUdpStatus gets the flood_protection.udp.status value from the UTM
func GetFloodProtectionUdpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/flood_protection.udp.status", &val, options...)
	return
}

// UpdateFloodProtectionUdpStatus PUTs the flood_protection.udp.status value to the UTM
func UpdateFloodProtectionUdpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/flood_protection.udp.status", val, options...)
}

// GetFtpAllowedClients gets the ftp.allowed_clients value from the UTM
func GetFtpAllowedClients(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ftp.allowed_clients", &val, options...)
	return
}

// UpdateFtpAllowedClients PUTs the ftp.allowed_clients value to the UTM
func UpdateFtpAllowedClients(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.allowed_clients", val, options...)
}

// GetFtpAllowedServers gets the ftp.allowed_servers value from the UTM
func GetFtpAllowedServers(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/ftp.allowed_servers", &val, options...)
	return
}

// UpdateFtpAllowedServers PUTs the ftp.allowed_servers value to the UTM
func UpdateFtpAllowedServers(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.allowed_servers", val, options...)
}

// GetFtpCffAv gets the ftp.cff_av value from the UTM
func GetFtpCffAv(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ftp.cff_av", &val, options...)
	return
}

// UpdateFtpCffAv PUTs the ftp.cff_av value to the UTM
func UpdateFtpCffAv(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.cff_av", val, options...)
}

// GetFtpCffAvEngines gets the ftp.cff_av_engines value from the UTM
func GetFtpCffAvEngines(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ftp.cff_av_engines", &val, options...)
	return
}

// UpdateFtpCffAvEngines PUTs the ftp.cff_av_engines value to the UTM
func UpdateFtpCffAvEngines(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.cff_av_engines", val, options...)
}

// GetFtpCffFileExtensions gets the ftp.cff_file_extensions value from the UTM
func GetFtpCffFileExtensions(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ftp.cff_file_extensions", &val, options...)
	return
}

// UpdateFtpCffFileExtensions PUTs the ftp.cff_file_extensions value to the UTM
func UpdateFtpCffFileExtensions(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.cff_file_extensions", val, options...)
}

// GetFtpExceptions gets the ftp.exceptions value from the UTM
func GetFtpExceptions(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ftp.exceptions", &val, options...)
	return
}

// UpdateFtpExceptions PUTs the ftp.exceptions value to the UTM
func UpdateFtpExceptions(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.exceptions", val, options...)
}

// GetFtpMaxFileSize gets the ftp.max_file_size value from the UTM
func GetFtpMaxFileSize(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ftp.max_file_size", &val, options...)
	return
}

// UpdateFtpMaxFileSize PUTs the ftp.max_file_size value to the UTM
func UpdateFtpMaxFileSize(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.max_file_size", val, options...)
}

// GetFtpMsWinMode gets the ftp.ms_win_mode value from the UTM
func GetFtpMsWinMode(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ftp.ms_win_mode", &val, options...)
	return
}

// UpdateFtpMsWinMode PUTs the ftp.ms_win_mode value to the UTM
func UpdateFtpMsWinMode(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.ms_win_mode", val, options...)
}

// GetFtpOperationMode gets the ftp.operation_mode value from the UTM
func GetFtpOperationMode(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ftp.operation_mode", &val, options...)
	return
}

// UpdateFtpOperationMode PUTs the ftp.operation_mode value to the UTM
func UpdateFtpOperationMode(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.operation_mode", val, options...)
}

// GetFtpRestrictedServers gets the ftp.restricted_servers value from the UTM
func GetFtpRestrictedServers(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/ftp.restricted_servers", &val, options...)
	return
}

// UpdateFtpRestrictedServers PUTs the ftp.restricted_servers value to the UTM
func UpdateFtpRestrictedServers(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.restricted_servers", val, options...)
}

// GetFtpStatus gets the ftp.status value from the UTM
func GetFtpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ftp.status", &val, options...)
	return
}

// UpdateFtpStatus PUTs the ftp.status value to the UTM
func UpdateFtpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.status", val, options...)
}

// GetFtpTransparentSkip gets the ftp.transparent_skip value from the UTM
func GetFtpTransparentSkip(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ftp.transparent_skip", &val, options...)
	return
}

// UpdateFtpTransparentSkip PUTs the ftp.transparent_skip value to the UTM
func UpdateFtpTransparentSkip(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.transparent_skip", val, options...)
}

// GetFtpTransparentSkipAutoPf gets the ftp.transparent_skip_auto_pf value from the UTM
func GetFtpTransparentSkipAutoPf(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ftp.transparent_skip_auto_pf", &val, options...)
	return
}

// UpdateFtpTransparentSkipAutoPf PUTs the ftp.transparent_skip_auto_pf value to the UTM
func UpdateFtpTransparentSkipAutoPf(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ftp.transparent_skip_auto_pf", val, options...)
}

// GetGenericProxyRules gets the generic_proxy.rules value from the UTM
func GetGenericProxyRules(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/generic_proxy.rules", &val, options...)
	return
}

// UpdateGenericProxyRules PUTs the generic_proxy.rules value to the UTM
func UpdateGenericProxyRules(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/generic_proxy.rules", val, options...)
}

// GetGeoipCountriesDst gets the geoip.countries_dst value from the UTM
func GetGeoipCountriesDst(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/geoip.countries_dst", &val, options...)
	return
}

// UpdateGeoipCountriesDst PUTs the geoip.countries_dst value to the UTM
func UpdateGeoipCountriesDst(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/geoip.countries_dst", val, options...)
}

// GetGeoipCountriesSrc gets the geoip.countries_src value from the UTM
func GetGeoipCountriesSrc(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/geoip.countries_src", &val, options...)
	return
}

// UpdateGeoipCountriesSrc PUTs the geoip.countries_src value to the UTM
func UpdateGeoipCountriesSrc(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/geoip.countries_src", val, options...)
}

// GetGeoipExceptions gets the geoip.exceptions value from the UTM
func GetGeoipExceptions(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/geoip.exceptions", &val, options...)
	return
}

// UpdateGeoipExceptions PUTs the geoip.exceptions value to the UTM
func UpdateGeoipExceptions(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/geoip.exceptions", val, options...)
}

// GetGeoipLog gets the geoip.log value from the UTM
func GetGeoipLog(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/geoip.log", &val, options...)
	return
}

// UpdateGeoipLog PUTs the geoip.log value to the UTM
func UpdateGeoipLog(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/geoip.log", val, options...)
}

// GetGeoipStatus gets the geoip.status value from the UTM
func GetGeoipStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/geoip.status", &val, options...)
	return
}

// UpdateGeoipStatus PUTs the geoip.status value to the UTM
func UpdateGeoipStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/geoip.status", val, options...)
}

// GetH323AllowedNetworks gets the h323.allowed_networks value from the UTM
func GetH323AllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/h323.allowed_networks", &val, options...)
	return
}

// UpdateH323AllowedNetworks PUTs the h323.allowed_networks value to the UTM
func UpdateH323AllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/h323.allowed_networks", val, options...)
}

// GetH323LogRelated gets the h323.log_related value from the UTM
func GetH323LogRelated(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/h323.log_related", &val, options...)
	return
}

// UpdateH323LogRelated PUTs the h323.log_related value to the UTM
func UpdateH323LogRelated(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/h323.log_related", val, options...)
}

// GetH323Servers gets the h323.servers value from the UTM
func GetH323Servers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/h323.servers", &val, options...)
	return
}

// UpdateH323Servers PUTs the h323.servers value to the UTM
func UpdateH323Servers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/h323.servers", val, options...)
}

// GetH323Status gets the h323.status value from the UTM
func GetH323Status(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/h323.status", &val, options...)
	return
}

// UpdateH323Status PUTs the h323.status value to the UTM
func UpdateH323Status(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/h323.status", val, options...)
}

// GetHaAdvancedAutojoin gets the ha.advanced.autojoin value from the UTM
func GetHaAdvancedAutojoin(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.advanced.autojoin", &val, options...)
	return
}

// UpdateHaAdvancedAutojoin PUTs the ha.advanced.autojoin value to the UTM
func UpdateHaAdvancedAutojoin(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.advanced.autojoin", val, options...)
}

// GetHaAdvancedColdRollback gets the ha.advanced.cold_rollback value from the UTM
func GetHaAdvancedColdRollback(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.advanced.cold_rollback", &val, options...)
	return
}

// UpdateHaAdvancedColdRollback PUTs the ha.advanced.cold_rollback value to the UTM
func UpdateHaAdvancedColdRollback(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.advanced.cold_rollback", val, options...)
}

// GetHaAdvancedHttpPersistenceTime gets the ha.advanced.http_persistence_time value from the UTM
func GetHaAdvancedHttpPersistenceTime(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/ha.advanced.http_persistence_time", &val, options...)
	return
}

// UpdateHaAdvancedHttpPersistenceTime PUTs the ha.advanced.http_persistence_time value to the UTM
func UpdateHaAdvancedHttpPersistenceTime(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.advanced.http_persistence_time", val, options...)
}

// GetHaAdvancedLoadTakeover gets the ha.advanced.load_takeover value from the UTM
func GetHaAdvancedLoadTakeover(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ha.advanced.load_takeover", &val, options...)
	return
}

// UpdateHaAdvancedLoadTakeover PUTs the ha.advanced.load_takeover value to the UTM
func UpdateHaAdvancedLoadTakeover(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.advanced.load_takeover", val, options...)
}

// GetHaAdvancedLoadWarn gets the ha.advanced.load_warn value from the UTM
func GetHaAdvancedLoadWarn(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ha.advanced.load_warn", &val, options...)
	return
}

// UpdateHaAdvancedLoadWarn PUTs the ha.advanced.load_warn value to the UTM
func UpdateHaAdvancedLoadWarn(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.advanced.load_warn", val, options...)
}

// GetHaAdvancedMaxNodes gets the ha.advanced.max_nodes value from the UTM
func GetHaAdvancedMaxNodes(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ha.advanced.max_nodes", &val, options...)
	return
}

// UpdateHaAdvancedMaxNodes PUTs the ha.advanced.max_nodes value to the UTM
func UpdateHaAdvancedMaxNodes(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.advanced.max_nodes", val, options...)
}

// GetHaAdvancedMtu gets the ha.advanced.mtu value from the UTM
func GetHaAdvancedMtu(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.advanced.mtu", &val, options...)
	return
}

// UpdateHaAdvancedMtu PUTs the ha.advanced.mtu value to the UTM
func UpdateHaAdvancedMtu(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.advanced.mtu", val, options...)
}

// GetHaAdvancedNetconsole gets the ha.advanced.netconsole value from the UTM
func GetHaAdvancedNetconsole(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.advanced.netconsole", &val, options...)
	return
}

// UpdateHaAdvancedNetconsole PUTs the ha.advanced.netconsole value to the UTM
func UpdateHaAdvancedNetconsole(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.advanced.netconsole", val, options...)
}

// GetHaAdvancedPreempt gets the ha.advanced.preempt value from the UTM
func GetHaAdvancedPreempt(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.advanced.preempt", &val, options...)
	return
}

// UpdateHaAdvancedPreempt PUTs the ha.advanced.preempt value to the UTM
func UpdateHaAdvancedPreempt(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.advanced.preempt", val, options...)
}

// GetHaAdvancedUniqueId gets the ha.advanced.unique_id value from the UTM
func GetHaAdvancedUniqueId(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ha.advanced.unique_id", &val, options...)
	return
}

// UpdateHaAdvancedUniqueId PUTs the ha.advanced.unique_id value to the UTM
func UpdateHaAdvancedUniqueId(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.advanced.unique_id", val, options...)
}

// GetHaAdvancedVirtualMac gets the ha.advanced.virtual_mac value from the UTM
func GetHaAdvancedVirtualMac(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.advanced.virtual_mac", &val, options...)
	return
}

// UpdateHaAdvancedVirtualMac PUTs the ha.advanced.virtual_mac value to the UTM
func UpdateHaAdvancedVirtualMac(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.advanced.virtual_mac", val, options...)
}

// GetHaAwsCloudwatchProfile gets the ha.aws.cloudwatch.profile value from the UTM
func GetHaAwsCloudwatchProfile(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.aws.cloudwatch.profile", &val, options...)
	return
}

// UpdateHaAwsCloudwatchProfile PUTs the ha.aws.cloudwatch.profile value to the UTM
func UpdateHaAwsCloudwatchProfile(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.cloudwatch.profile", val, options...)
}

// GetHaAwsCloudwatchStatus gets the ha.aws.cloudwatch.status value from the UTM
func GetHaAwsCloudwatchStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.aws.cloudwatch.status", &val, options...)
	return
}

// UpdateHaAwsCloudwatchStatus PUTs the ha.aws.cloudwatch.status value to the UTM
func UpdateHaAwsCloudwatchStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.cloudwatch.status", val, options...)
}

// GetHaAwsConfdBackup gets the ha.aws.confd.backup value from the UTM
func GetHaAwsConfdBackup(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.aws.confd.backup", &val, options...)
	return
}

// UpdateHaAwsConfdBackup PUTs the ha.aws.confd.backup value to the UTM
func UpdateHaAwsConfdBackup(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.confd.backup", val, options...)
}

// GetHaAwsConfdBackupInterval gets the ha.aws.confd.backup_interval value from the UTM
func GetHaAwsConfdBackupInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ha.aws.confd.backup_interval", &val, options...)
	return
}

// UpdateHaAwsConfdBackupInterval PUTs the ha.aws.confd.backup_interval value to the UTM
func UpdateHaAwsConfdBackupInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.confd.backup_interval", val, options...)
}

// GetHaAwsConfdRestore gets the ha.aws.confd.restore value from the UTM
func GetHaAwsConfdRestore(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.aws.confd.restore", &val, options...)
	return
}

// UpdateHaAwsConfdRestore PUTs the ha.aws.confd.restore value to the UTM
func UpdateHaAwsConfdRestore(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.confd.restore", val, options...)
}

// GetHaAwsConfdRestoreDone gets the ha.aws.confd.restore_done value from the UTM
func GetHaAwsConfdRestoreDone(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.aws.confd.restore_done", &val, options...)
	return
}

// UpdateHaAwsConfdRestoreDone PUTs the ha.aws.confd.restore_done value to the UTM
func UpdateHaAwsConfdRestoreDone(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.confd.restore_done", val, options...)
}

// GetHaAwsElasticIp gets the ha.aws.elastic_ip value from the UTM
func GetHaAwsElasticIp(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.aws.elastic_ip", &val, options...)
	return
}

// UpdateHaAwsElasticIp PUTs the ha.aws.elastic_ip value to the UTM
func UpdateHaAwsElasticIp(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.elastic_ip", val, options...)
}

// GetHaAwsPostgresArchiveTimeout gets the ha.aws.postgres.archive_timeout value from the UTM
func GetHaAwsPostgresArchiveTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ha.aws.postgres.archive_timeout", &val, options...)
	return
}

// UpdateHaAwsPostgresArchiveTimeout PUTs the ha.aws.postgres.archive_timeout value to the UTM
func UpdateHaAwsPostgresArchiveTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.postgres.archive_timeout", val, options...)
}

// GetHaAwsPostgresBackup gets the ha.aws.postgres.backup value from the UTM
func GetHaAwsPostgresBackup(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.aws.postgres.backup", &val, options...)
	return
}

// UpdateHaAwsPostgresBackup PUTs the ha.aws.postgres.backup value to the UTM
func UpdateHaAwsPostgresBackup(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.postgres.backup", val, options...)
}

// GetHaAwsPostgresBaseBackupInterval gets the ha.aws.postgres.base_backup_interval value from the UTM
func GetHaAwsPostgresBaseBackupInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ha.aws.postgres.base_backup_interval", &val, options...)
	return
}

// UpdateHaAwsPostgresBaseBackupInterval PUTs the ha.aws.postgres.base_backup_interval value to the UTM
func UpdateHaAwsPostgresBaseBackupInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.postgres.base_backup_interval", val, options...)
}

// GetHaAwsPostgresRestore gets the ha.aws.postgres.restore value from the UTM
func GetHaAwsPostgresRestore(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.aws.postgres.restore", &val, options...)
	return
}

// UpdateHaAwsPostgresRestore PUTs the ha.aws.postgres.restore value to the UTM
func UpdateHaAwsPostgresRestore(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.postgres.restore", val, options...)
}

// GetHaAwsS3Bucket gets the ha.aws.s3_bucket value from the UTM
func GetHaAwsS3Bucket(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.aws.s3_bucket", &val, options...)
	return
}

// UpdateHaAwsS3Bucket PUTs the ha.aws.s3_bucket value to the UTM
func UpdateHaAwsS3Bucket(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.s3_bucket", val, options...)
}

// GetHaAwsStackName gets the ha.aws.stack_name value from the UTM
func GetHaAwsStackName(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.aws.stack_name", &val, options...)
	return
}

// UpdateHaAwsStackName PUTs the ha.aws.stack_name value to the UTM
func UpdateHaAwsStackName(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.stack_name", val, options...)
}

// GetHaAwsSyslogBackup gets the ha.aws.syslog.backup value from the UTM
func GetHaAwsSyslogBackup(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.aws.syslog.backup", &val, options...)
	return
}

// UpdateHaAwsSyslogBackup PUTs the ha.aws.syslog.backup value to the UTM
func UpdateHaAwsSyslogBackup(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.syslog.backup", val, options...)
}

// GetHaAwsSyslogRestore gets the ha.aws.syslog.restore value from the UTM
func GetHaAwsSyslogRestore(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.aws.syslog.restore", &val, options...)
	return
}

// UpdateHaAwsSyslogRestore PUTs the ha.aws.syslog.restore value to the UTM
func UpdateHaAwsSyslogRestore(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.syslog.restore", val, options...)
}

// GetHaAwsSyslogRestorePeriod gets the ha.aws.syslog.restore_period value from the UTM
func GetHaAwsSyslogRestorePeriod(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ha.aws.syslog.restore_period", &val, options...)
	return
}

// UpdateHaAwsSyslogRestorePeriod PUTs the ha.aws.syslog.restore_period value to the UTM
func UpdateHaAwsSyslogRestorePeriod(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.syslog.restore_period", val, options...)
}

// GetHaAwsTrustedNetwork gets the ha.aws.trusted_network value from the UTM
func GetHaAwsTrustedNetwork(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.aws.trusted_network", &val, options...)
	return
}

// UpdateHaAwsTrustedNetwork PUTs the ha.aws.trusted_network value to the UTM
func UpdateHaAwsTrustedNetwork(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.aws.trusted_network", val, options...)
}

// GetHaClusterFtp gets the ha.cluster.ftp value from the UTM
func GetHaClusterFtp(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ha.cluster.ftp", &val, options...)
	return
}

// UpdateHaClusterFtp PUTs the ha.cluster.ftp value to the UTM
func UpdateHaClusterFtp(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.cluster.ftp", val, options...)
}

// GetHaClusterHttp gets the ha.cluster.http value from the UTM
func GetHaClusterHttp(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ha.cluster.http", &val, options...)
	return
}

// UpdateHaClusterHttp PUTs the ha.cluster.http value to the UTM
func UpdateHaClusterHttp(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.cluster.http", val, options...)
}

// GetHaClusterIpsec gets the ha.cluster.ipsec value from the UTM
func GetHaClusterIpsec(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ha.cluster.ipsec", &val, options...)
	return
}

// UpdateHaClusterIpsec PUTs the ha.cluster.ipsec value to the UTM
func UpdateHaClusterIpsec(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.cluster.ipsec", val, options...)
}

// GetHaClusterPop3 gets the ha.cluster.pop3 value from the UTM
func GetHaClusterPop3(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ha.cluster.pop3", &val, options...)
	return
}

// UpdateHaClusterPop3 PUTs the ha.cluster.pop3 value to the UTM
func UpdateHaClusterPop3(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.cluster.pop3", val, options...)
}

// GetHaClusterSmtp gets the ha.cluster.smtp value from the UTM
func GetHaClusterSmtp(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ha.cluster.smtp", &val, options...)
	return
}

// UpdateHaClusterSmtp PUTs the ha.cluster.smtp value to the UTM
func UpdateHaClusterSmtp(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.cluster.smtp", val, options...)
}

// GetHaClusterSnort gets the ha.cluster.snort value from the UTM
func GetHaClusterSnort(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ha.cluster.snort", &val, options...)
	return
}

// UpdateHaClusterSnort PUTs the ha.cluster.snort value to the UTM
func UpdateHaClusterSnort(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.cluster.snort", val, options...)
}

// GetHaClusterWaf gets the ha.cluster.waf value from the UTM
func GetHaClusterWaf(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ha.cluster.waf", &val, options...)
	return
}

// UpdateHaClusterWaf PUTs the ha.cluster.waf value to the UTM
func UpdateHaClusterWaf(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.cluster.waf", val, options...)
}

// GetHaDeviceName gets the ha.device_name value from the UTM
func GetHaDeviceName(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.device_name", &val, options...)
	return
}

// UpdateHaDeviceName PUTs the ha.device_name value to the UTM
func UpdateHaDeviceName(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.device_name", val, options...)
}

// GetHaItfhw gets the ha.itfhw value from the UTM
func GetHaItfhw(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.itfhw", &val, options...)
	return
}

// UpdateHaItfhw PUTs the ha.itfhw value to the UTM
func UpdateHaItfhw(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.itfhw", val, options...)
}

// GetHaItfhwBackup gets the ha.itfhw_backup value from the UTM
func GetHaItfhwBackup(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.itfhw_backup", &val, options...)
	return
}

// UpdateHaItfhwBackup PUTs the ha.itfhw_backup value to the UTM
func UpdateHaItfhwBackup(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.itfhw_backup", val, options...)
}

// GetHaMasterIp gets the ha.master_ip value from the UTM
func GetHaMasterIp(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.master_ip", &val, options...)
	return
}

// UpdateHaMasterIp PUTs the ha.master_ip value to the UTM
func UpdateHaMasterIp(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.master_ip", val, options...)
}

// GetHaMode gets the ha.mode value from the UTM
func GetHaMode(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.mode", &val, options...)
	return
}

// UpdateHaMode PUTs the ha.mode value to the UTM
func UpdateHaMode(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.mode", val, options...)
}

// GetHaNodeId gets the ha.node_id value from the UTM
func GetHaNodeId(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/ha.node_id", &val, options...)
	return
}

// UpdateHaNodeId PUTs the ha.node_id value to the UTM
func UpdateHaNodeId(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.node_id", val, options...)
}

// GetHaPassword gets the ha.password value from the UTM
func GetHaPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.password", &val, options...)
	return
}

// UpdateHaPassword PUTs the ha.password value to the UTM
func UpdateHaPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.password", val, options...)
}

// GetHaPostgresSecret gets the ha.postgres_secret value from the UTM
func GetHaPostgresSecret(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.postgres_secret", &val, options...)
	return
}

// UpdateHaPostgresSecret PUTs the ha.postgres_secret value to the UTM
func UpdateHaPostgresSecret(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.postgres_secret", val, options...)
}

// GetHaSlaveIp gets the ha.slave_ip value from the UTM
func GetHaSlaveIp(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.slave_ip", &val, options...)
	return
}

// UpdateHaSlaveIp PUTs the ha.slave_ip value to the UTM
func UpdateHaSlaveIp(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.slave_ip", val, options...)
}

// GetHaStatus gets the ha.status value from the UTM
func GetHaStatus(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ha.status", &val, options...)
	return
}

// UpdateHaStatus PUTs the ha.status value to the UTM
func UpdateHaStatus(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.status", val, options...)
}

// GetHaSyncConntrack gets the ha.sync.conntrack value from the UTM
func GetHaSyncConntrack(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.sync.conntrack", &val, options...)
	return
}

// UpdateHaSyncConntrack PUTs the ha.sync.conntrack value to the UTM
func UpdateHaSyncConntrack(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.sync.conntrack", val, options...)
}

// GetHaSyncDatabase gets the ha.sync.database value from the UTM
func GetHaSyncDatabase(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.sync.database", &val, options...)
	return
}

// UpdateHaSyncDatabase PUTs the ha.sync.database value to the UTM
func UpdateHaSyncDatabase(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.sync.database", val, options...)
}

// GetHaSyncFiles gets the ha.sync.files value from the UTM
func GetHaSyncFiles(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.sync.files", &val, options...)
	return
}

// UpdateHaSyncFiles PUTs the ha.sync.files value to the UTM
func UpdateHaSyncFiles(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.sync.files", val, options...)
}

// GetHaSyncIpsec gets the ha.sync.ipsec value from the UTM
func GetHaSyncIpsec(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.sync.ipsec", &val, options...)
	return
}

// UpdateHaSyncIpsec PUTs the ha.sync.ipsec value to the UTM
func UpdateHaSyncIpsec(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.sync.ipsec", val, options...)
}

// GetHaSyncSyslog gets the ha.sync.syslog value from the UTM
func GetHaSyncSyslog(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ha.sync.syslog", &val, options...)
	return
}

// UpdateHaSyncSyslog PUTs the ha.sync.syslog value to the UTM
func UpdateHaSyncSyslog(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.sync.syslog", val, options...)
}

// GetHaTimesDeadTime gets the ha.times.dead_time value from the UTM
func GetHaTimesDeadTime(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ha.times.dead_time", &val, options...)
	return
}

// UpdateHaTimesDeadTime PUTs the ha.times.dead_time value to the UTM
func UpdateHaTimesDeadTime(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.times.dead_time", val, options...)
}

// GetHaTimesLoadTime gets the ha.times.load_time value from the UTM
func GetHaTimesLoadTime(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ha.times.load_time", &val, options...)
	return
}

// UpdateHaTimesLoadTime PUTs the ha.times.load_time value to the UTM
func UpdateHaTimesLoadTime(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ha.times.load_time", val, options...)
}

// GetHotspotCert gets the hotspot.cert value from the UTM
func GetHotspotCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/hotspot.cert", &val, options...)
	return
}

// UpdateHotspotCert PUTs the hotspot.cert value to the UTM
func UpdateHotspotCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/hotspot.cert", val, options...)
}

// GetHotspotDeleteDays gets the hotspot.delete_days value from the UTM
func GetHotspotDeleteDays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/hotspot.delete_days", &val, options...)
	return
}

// UpdateHotspotDeleteDays PUTs the hotspot.delete_days value to the UTM
func UpdateHotspotDeleteDays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/hotspot.delete_days", val, options...)
}

// GetHotspotSslPortal gets the hotspot.ssl_portal value from the UTM
func GetHotspotSslPortal(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/hotspot.ssl_portal", &val, options...)
	return
}

// UpdateHotspotSslPortal PUTs the hotspot.ssl_portal value to the UTM
func UpdateHotspotSslPortal(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/hotspot.ssl_portal", val, options...)
}

// GetHotspotStatus gets the hotspot.status value from the UTM
func GetHotspotStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/hotspot.status", &val, options...)
	return
}

// UpdateHotspotStatus PUTs the hotspot.status value to the UTM
func UpdateHotspotStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/hotspot.status", val, options...)
}

// GetHotspotTransparentSkip gets the hotspot.transparent_skip value from the UTM
func GetHotspotTransparentSkip(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/hotspot.transparent_skip", &val, options...)
	return
}

// UpdateHotspotTransparentSkip PUTs the hotspot.transparent_skip value to the UTM
func UpdateHotspotTransparentSkip(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/hotspot.transparent_skip", val, options...)
}

// GetHttpAdSsoInterfaces gets the http.ad_sso_interfaces value from the UTM
func GetHttpAdSsoInterfaces(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/http.ad_sso_interfaces", &val, options...)
	return
}

// UpdateHttpAdSsoInterfaces PUTs the http.ad_sso_interfaces value to the UTM
func UpdateHttpAdSsoInterfaces(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.ad_sso_interfaces", val, options...)
}

// GetHttpAdssoRedirectUseHostname gets the http.adsso_redirect_use_hostname value from the UTM
func GetHttpAdssoRedirectUseHostname(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.adsso_redirect_use_hostname", &val, options...)
	return
}

// UpdateHttpAdssoRedirectUseHostname PUTs the http.adsso_redirect_use_hostname value to the UTM
func UpdateHttpAdssoRedirectUseHostname(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.adsso_redirect_use_hostname", val, options...)
}

// GetHttpAllowHttpsTrafficOverHttpPort gets the http.allow_https_traffic_over_http_port value from the UTM
func GetHttpAllowHttpsTrafficOverHttpPort(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.allow_https_traffic_over_http_port", &val, options...)
	return
}

// UpdateHttpAllowHttpsTrafficOverHttpPort PUTs the http.allow_https_traffic_over_http_port value to the UTM
func UpdateHttpAllowHttpsTrafficOverHttpPort(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.allow_https_traffic_over_http_port", val, options...)
}

// GetHttpAllowSsl3 gets the http.allow_ssl3 value from the UTM
func GetHttpAllowSsl3(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.allow_ssl3", &val, options...)
	return
}

// UpdateHttpAllowSsl3 PUTs the http.allow_ssl3 value to the UTM
func UpdateHttpAllowSsl3(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.allow_ssl3", val, options...)
}

// GetHttpAllowTls12 gets the http.allow_tls_1_2 value from the UTM
func GetHttpAllowTls12(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.allow_tls_1_2", &val, options...)
	return
}

// UpdateHttpAllowTls12 PUTs the http.allow_tls_1_2 value to the UTM
func UpdateHttpAllowTls12(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.allow_tls_1_2", val, options...)
}

// GetHttpAllowedPuas gets the http.allowed_puas value from the UTM
func GetHttpAllowedPuas(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/http.allowed_puas", &val, options...)
	return
}

// UpdateHttpAllowedPuas PUTs the http.allowed_puas value to the UTM
func UpdateHttpAllowedPuas(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.allowed_puas", val, options...)
}

// GetHttpAllowedTargetServices gets the http.allowed_target_services value from the UTM
func GetHttpAllowedTargetServices(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/http.allowed_target_services", &val, options...)
	return
}

// UpdateHttpAllowedTargetServices PUTs the http.allowed_target_services value to the UTM
func UpdateHttpAllowedTargetServices(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.allowed_target_services", val, options...)
}

// GetHttpAuaMaxconns gets the http.aua_maxconns value from the UTM
func GetHttpAuaMaxconns(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.aua_maxconns", &val, options...)
	return
}

// UpdateHttpAuaMaxconns PUTs the http.aua_maxconns value to the UTM
func UpdateHttpAuaMaxconns(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.aua_maxconns", val, options...)
}

// GetHttpAuaTimeout gets the http.aua_timeout value from the UTM
func GetHttpAuaTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.aua_timeout", &val, options...)
	return
}

// UpdateHttpAuaTimeout PUTs the http.aua_timeout value to the UTM
func UpdateHttpAuaTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.aua_timeout", val, options...)
}

// GetHttpAuthCacheSize gets the http.auth_cache_size value from the UTM
func GetHttpAuthCacheSize(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.auth_cache_size", &val, options...)
	return
}

// UpdateHttpAuthCacheSize PUTs the http.auth_cache_size value to the UTM
func UpdateHttpAuthCacheSize(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.auth_cache_size", val, options...)
}

// GetHttpAuthCacheTtl gets the http.auth_cache_ttl value from the UTM
func GetHttpAuthCacheTtl(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.auth_cache_ttl", &val, options...)
	return
}

// UpdateHttpAuthCacheTtl PUTs the http.auth_cache_ttl value to the UTM
func UpdateHttpAuthCacheTtl(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.auth_cache_ttl", val, options...)
}

// GetHttpAuthRealm gets the http.auth_realm value from the UTM
func GetHttpAuthRealm(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.auth_realm", &val, options...)
	return
}

// UpdateHttpAuthRealm PUTs the http.auth_realm value to the UTM
func UpdateHttpAuthRealm(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.auth_realm", val, options...)
}

// GetHttpAuthUsercacheTtl gets the http.auth_usercache_ttl value from the UTM
func GetHttpAuthUsercacheTtl(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.auth_usercache_ttl", &val, options...)
	return
}

// UpdateHttpAuthUsercacheTtl PUTs the http.auth_usercache_ttl value to the UTM
func UpdateHttpAuthUsercacheTtl(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.auth_usercache_ttl", val, options...)
}

// GetHttpBlockUnscannable gets the http.block_unscannable value from the UTM
func GetHttpBlockUnscannable(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.block_unscannable", &val, options...)
	return
}

// UpdateHttpBlockUnscannable PUTs the http.block_unscannable value to the UTM
func UpdateHttpBlockUnscannable(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.block_unscannable", val, options...)
}

// GetHttpBypassStreaming gets the http.bypass_streaming value from the UTM
func GetHttpBypassStreaming(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.bypass_streaming", &val, options...)
	return
}

// UpdateHttpBypassStreaming PUTs the http.bypass_streaming value to the UTM
func UpdateHttpBypassStreaming(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.bypass_streaming", val, options...)
}

// GetHttpCaList gets the http.ca_list value from the UTM
func GetHttpCaList(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/http.ca_list", &val, options...)
	return
}

// UpdateHttpCaList PUTs the http.ca_list value to the UTM
func UpdateHttpCaList(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.ca_list", val, options...)
}

// GetHttpCacheIgnoresCookies gets the http.cache_ignores_cookies value from the UTM
func GetHttpCacheIgnoresCookies(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.cache_ignores_cookies", &val, options...)
	return
}

// UpdateHttpCacheIgnoresCookies PUTs the http.cache_ignores_cookies value to the UTM
func UpdateHttpCacheIgnoresCookies(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.cache_ignores_cookies", val, options...)
}

// GetHttpCachessl gets the http.cachessl value from the UTM
func GetHttpCachessl(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.cachessl", &val, options...)
	return
}

// UpdateHttpCachessl PUTs the http.cachessl value to the UTM
func UpdateHttpCachessl(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.cachessl", val, options...)
}

// GetHttpCaching gets the http.caching value from the UTM
func GetHttpCaching(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.caching", &val, options...)
	return
}

// UpdateHttpCaching PUTs the http.caching value to the UTM
func UpdateHttpCaching(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.caching", val, options...)
}

// GetHttpCertcache gets the http.certcache value from the UTM
func GetHttpCertcache(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.certcache", &val, options...)
	return
}

// UpdateHttpCertcache PUTs the http.certcache value to the UTM
func UpdateHttpCertcache(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.certcache", val, options...)
}

// GetHttpCertstore gets the http.certstore value from the UTM
func GetHttpCertstore(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.certstore", &val, options...)
	return
}

// UpdateHttpCertstore PUTs the http.certstore value to the UTM
func UpdateHttpCertstore(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.certstore", val, options...)
}

// GetHttpCffOverrideUsers gets the http.cff_override_users value from the UTM
func GetHttpCffOverrideUsers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/http.cff_override_users", &val, options...)
	return
}

// UpdateHttpCffOverrideUsers PUTs the http.cff_override_users value to the UTM
func UpdateHttpCffOverrideUsers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.cff_override_users", val, options...)
}

// GetHttpClientTimeout gets the http.client_timeout value from the UTM
func GetHttpClientTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.client_timeout", &val, options...)
	return
}

// UpdateHttpClientTimeout PUTs the http.client_timeout value to the UTM
func UpdateHttpClientTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.client_timeout", val, options...)
}

// GetHttpConfLockWorkaround gets the http.conf_lock_workaround value from the UTM
func GetHttpConfLockWorkaround(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.conf_lock_workaround", &val, options...)
	return
}

// UpdateHttpConfLockWorkaround PUTs the http.conf_lock_workaround value to the UTM
func UpdateHttpConfLockWorkaround(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.conf_lock_workaround", val, options...)
}

// GetHttpConnectTimeout gets the http.connect_timeout value from the UTM
func GetHttpConnectTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.connect_timeout", &val, options...)
	return
}

// UpdateHttpConnectTimeout PUTs the http.connect_timeout value to the UTM
func UpdateHttpConnectTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.connect_timeout", val, options...)
}

// GetHttpConnectV6Timeout gets the http.connect_v6_timeout value from the UTM
func GetHttpConnectV6Timeout(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/http.connect_v6_timeout", &val, options...)
	return
}

// UpdateHttpConnectV6Timeout PUTs the http.connect_v6_timeout value to the UTM
func UpdateHttpConnectV6Timeout(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.connect_v6_timeout", val, options...)
}

// GetHttpConnlimit gets the http.connlimit value from the UTM
func GetHttpConnlimit(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.connlimit", &val, options...)
	return
}

// UpdateHttpConnlimit PUTs the http.connlimit value to the UTM
func UpdateHttpConnlimit(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.connlimit", val, options...)
}

// GetHttpCtypeInspectBody gets the http.ctype_inspect_body value from the UTM
func GetHttpCtypeInspectBody(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.ctype_inspect_body", &val, options...)
	return
}

// UpdateHttpCtypeInspectBody PUTs the http.ctype_inspect_body value to the UTM
func UpdateHttpCtypeInspectBody(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.ctype_inspect_body", val, options...)
}

// GetHttpCtypeUnpackArchive gets the http.ctype_unpack_archive value from the UTM
func GetHttpCtypeUnpackArchive(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.ctype_unpack_archive", &val, options...)
	return
}

// UpdateHttpCtypeUnpackArchive PUTs the http.ctype_unpack_archive value to the UTM
func UpdateHttpCtypeUnpackArchive(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.ctype_unpack_archive", val, options...)
}

// GetHttpDebug gets the http.debug value from the UTM
func GetHttpDebug(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/http.debug", &val, options...)
	return
}

// UpdateHttpDebug PUTs the http.debug value to the UTM
func UpdateHttpDebug(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.debug", val, options...)
}

// GetHttpDefaultblockaction gets the http.defaultblockaction value from the UTM
func GetHttpDefaultblockaction(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.defaultblockaction", &val, options...)
	return
}

// UpdateHttpDefaultblockaction PUTs the http.defaultblockaction value to the UTM
func UpdateHttpDefaultblockaction(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.defaultblockaction", val, options...)
}

// GetHttpDeferagents gets the http.deferagents value from the UTM
func GetHttpDeferagents(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/http.deferagents", &val, options...)
	return
}

// UpdateHttpDeferagents PUTs the http.deferagents value to the UTM
func UpdateHttpDeferagents(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.deferagents", val, options...)
}

// GetHttpDeferlength gets the http.deferlength value from the UTM
func GetHttpDeferlength(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.deferlength", &val, options...)
	return
}

// UpdateHttpDeferlength PUTs the http.deferlength value to the UTM
func UpdateHttpDeferlength(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.deferlength", val, options...)
}

// GetHttpDeferredDownloadReadyTimeout gets the http.deferred_download_ready_timeout value from the UTM
func GetHttpDeferredDownloadReadyTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.deferred_download_ready_timeout", &val, options...)
	return
}

// UpdateHttpDeferredDownloadReadyTimeout PUTs the http.deferred_download_ready_timeout value to the UTM
func UpdateHttpDeferredDownloadReadyTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.deferred_download_ready_timeout", val, options...)
}

// GetHttpDisplayHttpBlockpageExplicitMode gets the http.display_http_blockpage_explicit_mode value from the UTM
func GetHttpDisplayHttpBlockpageExplicitMode(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.display_http_blockpage_explicit_mode", &val, options...)
	return
}

// UpdateHttpDisplayHttpBlockpageExplicitMode PUTs the http.display_http_blockpage_explicit_mode value to the UTM
func UpdateHttpDisplayHttpBlockpageExplicitMode(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.display_http_blockpage_explicit_mode", val, options...)
}

// GetHttpDisplayIntro gets the http.display_intro value from the UTM
func GetHttpDisplayIntro(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.display_intro", &val, options...)
	return
}

// UpdateHttpDisplayIntro PUTs the http.display_intro value to the UTM
func UpdateHttpDisplayIntro(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.display_intro", val, options...)
}

// GetHttpDownloadManagerDefaultCharset gets the http.download_manager_default_charset value from the UTM
func GetHttpDownloadManagerDefaultCharset(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.download_manager_default_charset", &val, options...)
	return
}

// UpdateHttpDownloadManagerDefaultCharset PUTs the http.download_manager_default_charset value to the UTM
func UpdateHttpDownloadManagerDefaultCharset(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.download_manager_default_charset", val, options...)
}

// GetHttpEdirDelayBasicAuth gets the http.edir_delay_basic_auth value from the UTM
func GetHttpEdirDelayBasicAuth(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.edir_delay_basic_auth", &val, options...)
	return
}

// UpdateHttpEdirDelayBasicAuth PUTs the http.edir_delay_basic_auth value to the UTM
func UpdateHttpEdirDelayBasicAuth(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.edir_delay_basic_auth", val, options...)
}

// GetHttpEnableHsts gets the http.enable_hsts value from the UTM
func GetHttpEnableHsts(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.enable_hsts", &val, options...)
	return
}

// UpdateHttpEnableHsts PUTs the http.enable_hsts value to the UTM
func UpdateHttpEnableHsts(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.enable_hsts", val, options...)
}

// GetHttpEnableOutInterface gets the http.enable_out_interface value from the UTM
func GetHttpEnableOutInterface(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.enable_out_interface", &val, options...)
	return
}

// UpdateHttpEnableOutInterface PUTs the http.enable_out_interface value to the UTM
func UpdateHttpEnableOutInterface(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.enable_out_interface", val, options...)
}

// GetHttpEppQuotaAction gets the http.epp_quota_action value from the UTM
func GetHttpEppQuotaAction(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.epp_quota_action", &val, options...)
	return
}

// UpdateHttpEppQuotaAction PUTs the http.epp_quota_action value to the UTM
func UpdateHttpEppQuotaAction(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.epp_quota_action", val, options...)
}

// GetHttpExceptions gets the http.exceptions value from the UTM
func GetHttpExceptions(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/http.exceptions", &val, options...)
	return
}

// UpdateHttpExceptions PUTs the http.exceptions value to the UTM
func UpdateHttpExceptions(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.exceptions", val, options...)
}

// GetHttpForcedCachingExtension gets the http.forced_caching_extension value from the UTM
func GetHttpForcedCachingExtension(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/http.forced_caching_extension", &val, options...)
	return
}

// UpdateHttpForcedCachingExtension PUTs the http.forced_caching_extension value to the UTM
func UpdateHttpForcedCachingExtension(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.forced_caching_extension", val, options...)
}

// GetHttpForcedCachingNeverCachePrefix gets the http.forced_caching_never_cache_prefix value from the UTM
func GetHttpForcedCachingNeverCachePrefix(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/http.forced_caching_never_cache_prefix", &val, options...)
	return
}

// UpdateHttpForcedCachingNeverCachePrefix PUTs the http.forced_caching_never_cache_prefix value to the UTM
func UpdateHttpForcedCachingNeverCachePrefix(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.forced_caching_never_cache_prefix", val, options...)
}

// GetHttpForcedCachingStatus gets the http.forced_caching_status value from the UTM
func GetHttpForcedCachingStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.forced_caching_status", &val, options...)
	return
}

// UpdateHttpForcedCachingStatus PUTs the http.forced_caching_status value to the UTM
func UpdateHttpForcedCachingStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.forced_caching_status", val, options...)
}

// GetHttpForcedCachingTtl gets the http.forced_caching_ttl value from the UTM
func GetHttpForcedCachingTtl(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.forced_caching_ttl", &val, options...)
	return
}

// UpdateHttpForcedCachingTtl PUTs the http.forced_caching_ttl value to the UTM
func UpdateHttpForcedCachingTtl(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.forced_caching_ttl", val, options...)
}

// GetHttpForcedCachingUserAgentPrefix gets the http.forced_caching_user_agent_prefix value from the UTM
func GetHttpForcedCachingUserAgentPrefix(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/http.forced_caching_user_agent_prefix", &val, options...)
	return
}

// UpdateHttpForcedCachingUserAgentPrefix PUTs the http.forced_caching_user_agent_prefix value to the UTM
func UpdateHttpForcedCachingUserAgentPrefix(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.forced_caching_user_agent_prefix", val, options...)
}

// GetHttpHttpLoopbackDetect gets the http.http_loopback_detect value from the UTM
func GetHttpHttpLoopbackDetect(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.http_loopback_detect", &val, options...)
	return
}

// UpdateHttpHttpLoopbackDetect PUTs the http.http_loopback_detect value to the UTM
func UpdateHttpHttpLoopbackDetect(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.http_loopback_detect", val, options...)
}

// GetHttpIeSslBlockpageWorkaround gets the http.ie_ssl_blockpage_workaround value from the UTM
func GetHttpIeSslBlockpageWorkaround(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.ie_ssl_blockpage_workaround", &val, options...)
	return
}

// UpdateHttpIeSslBlockpageWorkaround PUTs the http.ie_ssl_blockpage_workaround value to the UTM
func UpdateHttpIeSslBlockpageWorkaround(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.ie_ssl_blockpage_workaround", val, options...)
}

// GetHttpLimitAdSsoInterfaces gets the http.limit_ad_sso_interfaces value from the UTM
func GetHttpLimitAdSsoInterfaces(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.limit_ad_sso_interfaces", &val, options...)
	return
}

// UpdateHttpLimitAdSsoInterfaces PUTs the http.limit_ad_sso_interfaces value to the UTM
func UpdateHttpLimitAdSsoInterfaces(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.limit_ad_sso_interfaces", val, options...)
}

// GetHttpLocalSiteList gets the http.local_site_list value from the UTM
func GetHttpLocalSiteList(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/http.local_site_list", &val, options...)
	return
}

// UpdateHttpLocalSiteList PUTs the http.local_site_list value to the UTM
func UpdateHttpLocalSiteList(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.local_site_list", val, options...)
}

// GetHttpMaxContentEncoding gets the http.max_content_encoding value from the UTM
func GetHttpMaxContentEncoding(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.max_content_encoding", &val, options...)
	return
}

// UpdateHttpMaxContentEncoding PUTs the http.max_content_encoding value to the UTM
func UpdateHttpMaxContentEncoding(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.max_content_encoding", val, options...)
}

// GetHttpMaxTempfileSize gets the http.max_tempfile_size value from the UTM
func GetHttpMaxTempfileSize(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.max_tempfile_size", &val, options...)
	return
}

// UpdateHttpMaxTempfileSize PUTs the http.max_tempfile_size value to the UTM
func UpdateHttpMaxTempfileSize(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.max_tempfile_size", val, options...)
}

// GetHttpMaxthreads gets the http.maxthreads value from the UTM
func GetHttpMaxthreads(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.maxthreads", &val, options...)
	return
}

// UpdateHttpMaxthreads PUTs the http.maxthreads value to the UTM
func UpdateHttpMaxthreads(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.maxthreads", val, options...)
}

// GetHttpMaxthreadsUnused gets the http.maxthreads_unused value from the UTM
func GetHttpMaxthreadsUnused(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.maxthreads_unused", &val, options...)
	return
}

// UpdateHttpMaxthreadsUnused PUTs the http.maxthreads_unused value to the UTM
func UpdateHttpMaxthreadsUnused(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.maxthreads_unused", val, options...)
}

// GetHttpModulepath gets the http.modulepath value from the UTM
func GetHttpModulepath(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.modulepath", &val, options...)
	return
}

// UpdateHttpModulepath PUTs the http.modulepath value to the UTM
func UpdateHttpModulepath(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.modulepath", val, options...)
}

// GetHttpModules gets the http.modules value from the UTM
func GetHttpModules(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/http.modules", &val, options...)
	return
}

// UpdateHttpModules PUTs the http.modules value to the UTM
func UpdateHttpModules(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.modules", val, options...)
}

// GetHttpNoscancontent gets the http.noscancontent value from the UTM
func GetHttpNoscancontent(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/http.noscancontent", &val, options...)
	return
}

// UpdateHttpNoscancontent PUTs the http.noscancontent value to the UTM
func UpdateHttpNoscancontent(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.noscancontent", val, options...)
}

// GetHttpOpendirectoryKeytab gets the http.opendirectory_keytab value from the UTM
func GetHttpOpendirectoryKeytab(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.opendirectory_keytab", &val, options...)
	return
}

// UpdateHttpOpendirectoryKeytab PUTs the http.opendirectory_keytab value to the UTM
func UpdateHttpOpendirectoryKeytab(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.opendirectory_keytab", val, options...)
}

// GetHttpOverrideProceedProtocol gets the http.override_proceed_protocol value from the UTM
func GetHttpOverrideProceedProtocol(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.override_proceed_protocol", &val, options...)
	return
}

// UpdateHttpOverrideProceedProtocol PUTs the http.override_proceed_protocol value to the UTM
func UpdateHttpOverrideProceedProtocol(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.override_proceed_protocol", val, options...)
}

// GetHttpPacFile gets the http.pac_file value from the UTM
func GetHttpPacFile(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.pac_file", &val, options...)
	return
}

// UpdateHttpPacFile PUTs the http.pac_file value to the UTM
func UpdateHttpPacFile(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.pac_file", val, options...)
}

// GetHttpParentProxyHost gets the http.parent_proxy_host value from the UTM
func GetHttpParentProxyHost(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.parent_proxy_host", &val, options...)
	return
}

// UpdateHttpParentProxyHost PUTs the http.parent_proxy_host value to the UTM
func UpdateHttpParentProxyHost(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.parent_proxy_host", val, options...)
}

// GetHttpParentProxyPort gets the http.parent_proxy_port value from the UTM
func GetHttpParentProxyPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.parent_proxy_port", &val, options...)
	return
}

// UpdateHttpParentProxyPort PUTs the http.parent_proxy_port value to the UTM
func UpdateHttpParentProxyPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.parent_proxy_port", val, options...)
}

// GetHttpParentProxyStatus gets the http.parent_proxy_status value from the UTM
func GetHttpParentProxyStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.parent_proxy_status", &val, options...)
	return
}

// UpdateHttpParentProxyStatus PUTs the http.parent_proxy_status value to the UTM
func UpdateHttpParentProxyStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.parent_proxy_status", val, options...)
}

// GetHttpPassthroughId gets the http.passthrough_id value from the UTM
func GetHttpPassthroughId(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.passthrough_id", &val, options...)
	return
}

// UpdateHttpPassthroughId PUTs the http.passthrough_id value to the UTM
func UpdateHttpPassthroughId(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.passthrough_id", val, options...)
}

// GetHttpPharmingProtection gets the http.pharming_protection value from the UTM
func GetHttpPharmingProtection(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.pharming_protection", &val, options...)
	return
}

// UpdateHttpPharmingProtection PUTs the http.pharming_protection value to the UTM
func UpdateHttpPharmingProtection(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.pharming_protection", val, options...)
}

// GetHttpPort gets the http.port value from the UTM
func GetHttpPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.port", &val, options...)
	return
}

// UpdateHttpPort PUTs the http.port value to the UTM
func UpdateHttpPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.port", val, options...)
}

// GetHttpPortalCert gets the http.portal_cert value from the UTM
func GetHttpPortalCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.portal_cert", &val, options...)
	return
}

// UpdateHttpPortalCert PUTs the http.portal_cert value to the UTM
func UpdateHttpPortalCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.portal_cert", val, options...)
}

// GetHttpPortalCertChain gets the http.portal_cert_chain value from the UTM
func GetHttpPortalCertChain(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/http.portal_cert_chain", &val, options...)
	return
}

// UpdateHttpPortalCertChain PUTs the http.portal_cert_chain value to the UTM
func UpdateHttpPortalCertChain(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.portal_cert_chain", val, options...)
}

// GetHttpPortalDomain gets the http.portal_domain value from the UTM
func GetHttpPortalDomain(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.portal_domain", &val, options...)
	return
}

// UpdateHttpPortalDomain PUTs the http.portal_domain value to the UTM
func UpdateHttpPortalDomain(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.portal_domain", val, options...)
}

// GetHttpPortalHosts gets the http.portal_hosts value from the UTM
func GetHttpPortalHosts(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/http.portal_hosts", &val, options...)
	return
}

// UpdateHttpPortalHosts PUTs the http.portal_hosts value to the UTM
func UpdateHttpPortalHosts(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.portal_hosts", val, options...)
}

// GetHttpPortalUseCert gets the http.portal_use_cert value from the UTM
func GetHttpPortalUseCert(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.portal_use_cert", &val, options...)
	return
}

// UpdateHttpPortalUseCert PUTs the http.portal_use_cert value to the UTM
func UpdateHttpPortalUseCert(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.portal_use_cert", val, options...)
}

// GetHttpProceedCacheTimeout gets the http.proceed_cache_timeout value from the UTM
func GetHttpProceedCacheTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.proceed_cache_timeout", &val, options...)
	return
}

// UpdateHttpProceedCacheTimeout PUTs the http.proceed_cache_timeout value to the UTM
func UpdateHttpProceedCacheTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.proceed_cache_timeout", val, options...)
}

// GetHttpProfiles gets the http.profiles value from the UTM
func GetHttpProfiles(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/http.profiles", &val, options...)
	return
}

// UpdateHttpProfiles PUTs the http.profiles value to the UTM
func UpdateHttpProfiles(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.profiles", val, options...)
}

// GetHttpQuotaSliceTime gets the http.quota_slice_time value from the UTM
func GetHttpQuotaSliceTime(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.quota_slice_time", &val, options...)
	return
}

// UpdateHttpQuotaSliceTime PUTs the http.quota_slice_time value to the UTM
func UpdateHttpQuotaSliceTime(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.quota_slice_time", val, options...)
}

// GetHttpRemoveRequest gets the http.remove_request value from the UTM
func GetHttpRemoveRequest(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/http.remove_request", &val, options...)
	return
}

// UpdateHttpRemoveRequest PUTs the http.remove_request value to the UTM
func UpdateHttpRemoveRequest(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.remove_request", val, options...)
}

// GetHttpRemoveResponse gets the http.remove_response value from the UTM
func GetHttpRemoveResponse(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/http.remove_response", &val, options...)
	return
}

// UpdateHttpRemoveResponse PUTs the http.remove_response value to the UTM
func UpdateHttpRemoveResponse(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.remove_response", val, options...)
}

// GetHttpResponseTimeout gets the http.response_timeout value from the UTM
func GetHttpResponseTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.response_timeout", &val, options...)
	return
}

// UpdateHttpResponseTimeout PUTs the http.response_timeout value to the UTM
func UpdateHttpResponseTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.response_timeout", val, options...)
}

// GetHttpSaviScanTimeout gets the http.savi_scan_timeout value from the UTM
func GetHttpSaviScanTimeout(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/http.savi_scan_timeout", &val, options...)
	return
}

// UpdateHttpSaviScanTimeout PUTs the http.savi_scan_timeout value to the UTM
func UpdateHttpSaviScanTimeout(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.savi_scan_timeout", val, options...)
}

// GetHttpScLocalDb gets the http.sc_local_db value from the UTM
func GetHttpScLocalDb(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.sc_local_db", &val, options...)
	return
}

// UpdateHttpScLocalDb PUTs the http.sc_local_db value to the UTM
func UpdateHttpScLocalDb(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.sc_local_db", val, options...)
}

// GetHttpScanEppTraffic gets the http.scan_epp_traffic value from the UTM
func GetHttpScanEppTraffic(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.scan_epp_traffic", &val, options...)
	return
}

// UpdateHttpScanEppTraffic PUTs the http.scan_epp_traffic value to the UTM
func UpdateHttpScanEppTraffic(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.scan_epp_traffic", val, options...)
}

// GetHttpSearchdomain gets the http.searchdomain value from the UTM
func GetHttpSearchdomain(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.searchdomain", &val, options...)
	return
}

// UpdateHttpSearchdomain PUTs the http.searchdomain value to the UTM
func UpdateHttpSearchdomain(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.searchdomain", val, options...)
}

// GetHttpStrictHttp gets the http.strict_http value from the UTM
func GetHttpStrictHttp(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.strict_http", &val, options...)
	return
}

// UpdateHttpStrictHttp PUTs the http.strict_http value to the UTM
func UpdateHttpStrictHttp(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.strict_http", val, options...)
}

// GetHttpTlsciphersClient gets the http.tlsciphers_client value from the UTM
func GetHttpTlsciphersClient(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.tlsciphers_client", &val, options...)
	return
}

// UpdateHttpTlsciphersClient PUTs the http.tlsciphers_client value to the UTM
func UpdateHttpTlsciphersClient(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.tlsciphers_client", val, options...)
}

// GetHttpTlsciphersServer gets the http.tlsciphers_server value from the UTM
func GetHttpTlsciphersServer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.tlsciphers_server", &val, options...)
	return
}

// UpdateHttpTlsciphersServer PUTs the http.tlsciphers_server value to the UTM
func UpdateHttpTlsciphersServer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.tlsciphers_server", val, options...)
}

// GetHttpTmpfsUsageMinMemsize gets the http.tmpfs_usage_min_memsize value from the UTM
func GetHttpTmpfsUsageMinMemsize(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.tmpfs_usage_min_memsize", &val, options...)
	return
}

// UpdateHttpTmpfsUsageMinMemsize PUTs the http.tmpfs_usage_min_memsize value to the UTM
func UpdateHttpTmpfsUsageMinMemsize(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.tmpfs_usage_min_memsize", val, options...)
}

// GetHttpTransparentAuthTimeout gets the http.transparent_auth_timeout value from the UTM
func GetHttpTransparentAuthTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.transparent_auth_timeout", &val, options...)
	return
}

// UpdateHttpTransparentAuthTimeout PUTs the http.transparent_auth_timeout value to the UTM
func UpdateHttpTransparentAuthTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.transparent_auth_timeout", val, options...)
}

// GetHttpTransparentDstSkip gets the http.transparent_dst_skip value from the UTM
func GetHttpTransparentDstSkip(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/http.transparent_dst_skip", &val, options...)
	return
}

// UpdateHttpTransparentDstSkip PUTs the http.transparent_dst_skip value to the UTM
func UpdateHttpTransparentDstSkip(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.transparent_dst_skip", val, options...)
}

// GetHttpTransparentSkipAutoPf gets the http.transparent_skip_auto_pf value from the UTM
func GetHttpTransparentSkipAutoPf(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.transparent_skip_auto_pf", &val, options...)
	return
}

// UpdateHttpTransparentSkipAutoPf PUTs the http.transparent_skip_auto_pf value to the UTM
func UpdateHttpTransparentSkipAutoPf(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.transparent_skip_auto_pf", val, options...)
}

// GetHttpTransparentSrcSkip gets the http.transparent_src_skip value from the UTM
func GetHttpTransparentSrcSkip(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/http.transparent_src_skip", &val, options...)
	return
}

// UpdateHttpTransparentSrcSkip PUTs the http.transparent_src_skip value to the UTM
func UpdateHttpTransparentSrcSkip(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.transparent_src_skip", val, options...)
}

// GetHttpTunnelTimeout gets the http.tunnel_timeout value from the UTM
func GetHttpTunnelTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/http.tunnel_timeout", &val, options...)
	return
}

// UpdateHttpTunnelTimeout PUTs the http.tunnel_timeout value to the UTM
func UpdateHttpTunnelTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.tunnel_timeout", val, options...)
}

// GetHttpTunnelV6Timeout gets the http.tunnel_v6_timeout value from the UTM
func GetHttpTunnelV6Timeout(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/http.tunnel_v6_timeout", &val, options...)
	return
}

// UpdateHttpTunnelV6Timeout PUTs the http.tunnel_v6_timeout value to the UTM
func UpdateHttpTunnelV6Timeout(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.tunnel_v6_timeout", val, options...)
}

// GetHttpUndefercontent gets the http.undefercontent value from the UTM
func GetHttpUndefercontent(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/http.undefercontent", &val, options...)
	return
}

// UpdateHttpUndefercontent PUTs the http.undefercontent value to the UTM
func UpdateHttpUndefercontent(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.undefercontent", val, options...)
}

// GetHttpUndeferextension gets the http.undeferextension value from the UTM
func GetHttpUndeferextension(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/http.undeferextension", &val, options...)
	return
}

// UpdateHttpUndeferextension PUTs the http.undeferextension value to the UTM
func UpdateHttpUndeferextension(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.undeferextension", val, options...)
}

// GetHttpUrlFilteringRedirectUrl gets the http.url_filtering_redirect_url value from the UTM
func GetHttpUrlFilteringRedirectUrl(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/http.url_filtering_redirect_url", &val, options...)
	return
}

// UpdateHttpUrlFilteringRedirectUrl PUTs the http.url_filtering_redirect_url value to the UTM
func UpdateHttpUrlFilteringRedirectUrl(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.url_filtering_redirect_url", val, options...)
}

// GetHttpUseConnectionInsteadofProxyconnection gets the http.use_connection_insteadof_proxyconnection value from the UTM
func GetHttpUseConnectionInsteadofProxyconnection(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.use_connection_insteadof_proxyconnection", &val, options...)
	return
}

// UpdateHttpUseConnectionInsteadofProxyconnection PUTs the http.use_connection_insteadof_proxyconnection value to the UTM
func UpdateHttpUseConnectionInsteadofProxyconnection(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.use_connection_insteadof_proxyconnection", val, options...)
}

// GetHttpUseDnsCnameSafesearch gets the http.use_dns_cname_safesearch value from the UTM
func GetHttpUseDnsCnameSafesearch(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.use_dns_cname_safesearch", &val, options...)
	return
}

// UpdateHttpUseDnsCnameSafesearch PUTs the http.use_dns_cname_safesearch value to the UTM
func UpdateHttpUseDnsCnameSafesearch(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.use_dns_cname_safesearch", val, options...)
}

// GetHttpUseDstaddrForGeopiplookup gets the http.use_dstaddr_for_geopiplookup value from the UTM
func GetHttpUseDstaddrForGeopiplookup(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.use_dstaddr_for_geopiplookup", &val, options...)
	return
}

// UpdateHttpUseDstaddrForGeopiplookup PUTs the http.use_dstaddr_for_geopiplookup value to the UTM
func UpdateHttpUseDstaddrForGeopiplookup(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.use_dstaddr_for_geopiplookup", val, options...)
}

// GetHttpUseKrb5Adsso gets the http.use_krb5_adsso value from the UTM
func GetHttpUseKrb5Adsso(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.use_krb5_adsso", &val, options...)
	return
}

// UpdateHttpUseKrb5Adsso PUTs the http.use_krb5_adsso value to the UTM
func UpdateHttpUseKrb5Adsso(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.use_krb5_adsso", val, options...)
}

// GetHttpUseSni gets the http.use_sni value from the UTM
func GetHttpUseSni(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.use_sni", &val, options...)
	return
}

// UpdateHttpUseSni PUTs the http.use_sni value to the UTM
func UpdateHttpUseSni(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.use_sni", val, options...)
}

// GetHttpUseSxlUrid gets the http.use_sxl_urid value from the UTM
func GetHttpUseSxlUrid(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/http.use_sxl_urid", &val, options...)
	return
}

// UpdateHttpUseSxlUrid PUTs the http.use_sxl_urid value to the UTM
func UpdateHttpUseSxlUrid(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/http.use_sxl_urid", val, options...)
}

// GetIcmpForward gets the icmp.forward value from the UTM
func GetIcmpForward(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/icmp.forward", &val, options...)
	return
}

// UpdateIcmpForward PUTs the icmp.forward value to the UTM
func UpdateIcmpForward(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/icmp.forward", val, options...)
}

// GetIcmpInput gets the icmp.input value from the UTM
func GetIcmpInput(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/icmp.input", &val, options...)
	return
}

// UpdateIcmpInput PUTs the icmp.input value to the UTM
func UpdateIcmpInput(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/icmp.input", val, options...)
}

// GetIcmpLogRedirect gets the icmp.log_redirect value from the UTM
func GetIcmpLogRedirect(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/icmp.log_redirect", &val, options...)
	return
}

// UpdateIcmpLogRedirect PUTs the icmp.log_redirect value to the UTM
func UpdateIcmpLogRedirect(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/icmp.log_redirect", val, options...)
}

// GetIcmpPingForward gets the icmp.ping.forward value from the UTM
func GetIcmpPingForward(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/icmp.ping.forward", &val, options...)
	return
}

// UpdateIcmpPingForward PUTs the icmp.ping.forward value to the UTM
func UpdateIcmpPingForward(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/icmp.ping.forward", val, options...)
}

// GetIcmpPingInput gets the icmp.ping.input value from the UTM
func GetIcmpPingInput(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/icmp.ping.input", &val, options...)
	return
}

// UpdateIcmpPingInput PUTs the icmp.ping.input value to the UTM
func UpdateIcmpPingInput(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/icmp.ping.input", val, options...)
}

// GetIcmpPingOutput gets the icmp.ping.output value from the UTM
func GetIcmpPingOutput(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/icmp.ping.output", &val, options...)
	return
}

// UpdateIcmpPingOutput PUTs the icmp.ping.output value to the UTM
func UpdateIcmpPingOutput(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/icmp.ping.output", val, options...)
}

// GetIcmpSecure gets the icmp.secure value from the UTM
func GetIcmpSecure(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/icmp.secure", &val, options...)
	return
}

// UpdateIcmpSecure PUTs the icmp.secure value to the UTM
func UpdateIcmpSecure(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/icmp.secure", val, options...)
}

// GetIcmpTracerouteForward gets the icmp.traceroute.forward value from the UTM
func GetIcmpTracerouteForward(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/icmp.traceroute.forward", &val, options...)
	return
}

// UpdateIcmpTracerouteForward PUTs the icmp.traceroute.forward value to the UTM
func UpdateIcmpTracerouteForward(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/icmp.traceroute.forward", val, options...)
}

// GetIcmpTracerouteInput gets the icmp.traceroute.input value from the UTM
func GetIcmpTracerouteInput(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/icmp.traceroute.input", &val, options...)
	return
}

// UpdateIcmpTracerouteInput PUTs the icmp.traceroute.input value to the UTM
func UpdateIcmpTracerouteInput(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/icmp.traceroute.input", val, options...)
}

// GetIdentForward gets the ident.forward value from the UTM
func GetIdentForward(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ident.forward", &val, options...)
	return
}

// UpdateIdentForward PUTs the ident.forward value to the UTM
func UpdateIdentForward(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ident.forward", val, options...)
}

// GetIdentResponse gets the ident.response value from the UTM
func GetIdentResponse(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ident.response", &val, options...)
	return
}

// UpdateIdentResponse PUTs the ident.response value to the UTM
func UpdateIdentResponse(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ident.response", val, options...)
}

// GetIdentStatus gets the ident.status value from the UTM
func GetIdentStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ident.status", &val, options...)
	return
}

// UpdateIdentStatus PUTs the ident.status value to the UTM
func UpdateIdentStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ident.status", val, options...)
}

// GetInterfacesAdvancedArpAnnounce gets the interfaces.advanced.arp_announce value from the UTM
func GetInterfacesAdvancedArpAnnounce(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/interfaces.advanced.arp_announce", &val, options...)
	return
}

// UpdateInterfacesAdvancedArpAnnounce PUTs the interfaces.advanced.arp_announce value to the UTM
func UpdateInterfacesAdvancedArpAnnounce(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/interfaces.advanced.arp_announce", val, options...)
}

// GetInterfacesAdvancedArpCacheGcThresh1 gets the interfaces.advanced.arp_cache_gc_thresh1 value from the UTM
func GetInterfacesAdvancedArpCacheGcThresh1(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/interfaces.advanced.arp_cache_gc_thresh1", &val, options...)
	return
}

// UpdateInterfacesAdvancedArpCacheGcThresh1 PUTs the interfaces.advanced.arp_cache_gc_thresh1 value to the UTM
func UpdateInterfacesAdvancedArpCacheGcThresh1(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/interfaces.advanced.arp_cache_gc_thresh1", val, options...)
}

// GetInterfacesAdvancedArpCacheGcThresh2 gets the interfaces.advanced.arp_cache_gc_thresh2 value from the UTM
func GetInterfacesAdvancedArpCacheGcThresh2(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/interfaces.advanced.arp_cache_gc_thresh2", &val, options...)
	return
}

// UpdateInterfacesAdvancedArpCacheGcThresh2 PUTs the interfaces.advanced.arp_cache_gc_thresh2 value to the UTM
func UpdateInterfacesAdvancedArpCacheGcThresh2(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/interfaces.advanced.arp_cache_gc_thresh2", val, options...)
}

// GetInterfacesAdvancedArpCacheGcThresh3 gets the interfaces.advanced.arp_cache_gc_thresh3 value from the UTM
func GetInterfacesAdvancedArpCacheGcThresh3(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/interfaces.advanced.arp_cache_gc_thresh3", &val, options...)
	return
}

// UpdateInterfacesAdvancedArpCacheGcThresh3 PUTs the interfaces.advanced.arp_cache_gc_thresh3 value to the UTM
func UpdateInterfacesAdvancedArpCacheGcThresh3(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/interfaces.advanced.arp_cache_gc_thresh3", val, options...)
}

// GetInterfacesAdvancedArpIgnore gets the interfaces.advanced.arp_ignore value from the UTM
func GetInterfacesAdvancedArpIgnore(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/interfaces.advanced.arp_ignore", &val, options...)
	return
}

// UpdateInterfacesAdvancedArpIgnore PUTs the interfaces.advanced.arp_ignore value to the UTM
func UpdateInterfacesAdvancedArpIgnore(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/interfaces.advanced.arp_ignore", val, options...)
}

// GetInterfacesAdvancedDefaultMetric gets the interfaces.advanced.default_metric value from the UTM
func GetInterfacesAdvancedDefaultMetric(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/interfaces.advanced.default_metric", &val, options...)
	return
}

// UpdateInterfacesAdvancedDefaultMetric PUTs the interfaces.advanced.default_metric value to the UTM
func UpdateInterfacesAdvancedDefaultMetric(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/interfaces.advanced.default_metric", val, options...)
}

// GetInterfacesInterfaces gets the interfaces.interfaces value from the UTM
func GetInterfacesInterfaces(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/interfaces.interfaces", &val, options...)
	return
}

// UpdateInterfacesInterfaces PUTs the interfaces.interfaces value to the UTM
func UpdateInterfacesInterfaces(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/interfaces.interfaces", val, options...)
}

// GetIpsDnsServers gets the ips.dns_servers value from the UTM
func GetIpsDnsServers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ips.dns_servers", &val, options...)
	return
}

// UpdateIpsDnsServers PUTs the ips.dns_servers value to the UTM
func UpdateIpsDnsServers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.dns_servers", val, options...)
}

// GetIpsEnableAcceptForAllPackets gets the ips.enable_accept_for_all_packets value from the UTM
func GetIpsEnableAcceptForAllPackets(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ips.enable_accept_for_all_packets", &val, options...)
	return
}

// UpdateIpsEnableAcceptForAllPackets PUTs the ips.enable_accept_for_all_packets value to the UTM
func UpdateIpsEnableAcceptForAllPackets(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.enable_accept_for_all_packets", val, options...)
}

// GetIpsEngine gets the ips.engine value from the UTM
func GetIpsEngine(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ips.engine", &val, options...)
	return
}

// UpdateIpsEngine PUTs the ips.engine value to the UTM
func UpdateIpsEngine(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.engine", val, options...)
}

// GetIpsExceptions gets the ips.exceptions value from the UTM
func GetIpsExceptions(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/ips.exceptions", &val, options...)
	return
}

// UpdateIpsExceptions PUTs the ips.exceptions value to the UTM
func UpdateIpsExceptions(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.exceptions", val, options...)
}

// GetIpsFailopen gets the ips.failopen value from the UTM
func GetIpsFailopen(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ips.failopen", &val, options...)
	return
}

// UpdateIpsFailopen PUTs the ips.failopen value to the UTM
func UpdateIpsFailopen(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.failopen", val, options...)
}

// GetIpsFileBasedRules gets the ips.file_based_rules value from the UTM
func GetIpsFileBasedRules(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ips.file_based_rules", &val, options...)
	return
}

// UpdateIpsFileBasedRules PUTs the ips.file_based_rules value to the UTM
func UpdateIpsFileBasedRules(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.file_based_rules", val, options...)
}

// GetIpsGroups gets the ips.groups value from the UTM
func GetIpsGroups(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/ips.groups", &val, options...)
	return
}

// UpdateIpsGroups PUTs the ips.groups value to the UTM
func UpdateIpsGroups(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.groups", val, options...)
}

// GetIpsHttpServers gets the ips.http_servers value from the UTM
func GetIpsHttpServers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ips.http_servers", &val, options...)
	return
}

// UpdateIpsHttpServers PUTs the ips.http_servers value to the UTM
func UpdateIpsHttpServers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.http_servers", val, options...)
}

// GetIpsIpsfbAlertInterval gets the ips.ipsfb.alert_interval value from the UTM
func GetIpsIpsfbAlertInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ips.ipsfb.alert_interval", &val, options...)
	return
}

// UpdateIpsIpsfbAlertInterval PUTs the ips.ipsfb.alert_interval value to the UTM
func UpdateIpsIpsfbAlertInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.ipsfb.alert_interval", val, options...)
}

// GetIpsIpsfbConfigInterval gets the ips.ipsfb.config_interval value from the UTM
func GetIpsIpsfbConfigInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ips.ipsfb.config_interval", &val, options...)
	return
}

// UpdateIpsIpsfbConfigInterval PUTs the ips.ipsfb.config_interval value to the UTM
func UpdateIpsIpsfbConfigInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.ipsfb.config_interval", val, options...)
}

// GetIpsIpsfbDebug gets the ips.ipsfb.debug value from the UTM
func GetIpsIpsfbDebug(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ips.ipsfb.debug", &val, options...)
	return
}

// UpdateIpsIpsfbDebug PUTs the ips.ipsfb.debug value to the UTM
func UpdateIpsIpsfbDebug(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.ipsfb.debug", val, options...)
}

// GetIpsLocalNetworks gets the ips.local_networks value from the UTM
func GetIpsLocalNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/ips.local_networks", &val, options...)
	return
}

// UpdateIpsLocalNetworks PUTs the ips.local_networks value to the UTM
func UpdateIpsLocalNetworks(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.local_networks", val, options...)
}

// GetIpsNumInstances gets the ips.num_instances value from the UTM
func GetIpsNumInstances(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/ips.num_instances", &val, options...)
	return
}

// UpdateIpsNumInstances PUTs the ips.num_instances value to the UTM
func UpdateIpsNumInstances(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.num_instances", val, options...)
}

// GetIpsPatternChannel gets the ips.pattern_channel value from the UTM
func GetIpsPatternChannel(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ips.pattern_channel", &val, options...)
	return
}

// UpdateIpsPatternChannel PUTs the ips.pattern_channel value to the UTM
func UpdateIpsPatternChannel(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.pattern_channel", val, options...)
}

// GetIpsPolicy gets the ips.policy value from the UTM
func GetIpsPolicy(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ips.policy", &val, options...)
	return
}

// UpdateIpsPolicy PUTs the ips.policy value to the UTM
func UpdateIpsPolicy(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.policy", val, options...)
}

// GetIpsQueueLength gets the ips.queue_length value from the UTM
func GetIpsQueueLength(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ips.queue_length", &val, options...)
	return
}

// UpdateIpsQueueLength PUTs the ips.queue_length value to the UTM
func UpdateIpsQueueLength(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.queue_length", val, options...)
}

// GetIpsQueueThreshold gets the ips.queue_threshold value from the UTM
func GetIpsQueueThreshold(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/ips.queue_threshold", &val, options...)
	return
}

// UpdateIpsQueueThreshold PUTs the ips.queue_threshold value to the UTM
func UpdateIpsQueueThreshold(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.queue_threshold", val, options...)
}

// GetIpsReloadMethod gets the ips.reload_method value from the UTM
func GetIpsReloadMethod(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ips.reload_method", &val, options...)
	return
}

// UpdateIpsReloadMethod PUTs the ips.reload_method value to the UTM
func UpdateIpsReloadMethod(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.reload_method", val, options...)
}

// GetIpsRestartPolicy gets the ips.restart_policy value from the UTM
func GetIpsRestartPolicy(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ips.restart_policy", &val, options...)
	return
}

// UpdateIpsRestartPolicy PUTs the ips.restart_policy value to the UTM
func UpdateIpsRestartPolicy(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.restart_policy", val, options...)
}

// GetIpsRuleModifiers gets the ips.rule_modifiers value from the UTM
func GetIpsRuleModifiers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ips.rule_modifiers", &val, options...)
	return
}

// UpdateIpsRuleModifiers PUTs the ips.rule_modifiers value to the UTM
func UpdateIpsRuleModifiers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.rule_modifiers", val, options...)
}

// GetIpsRules gets the ips.rules value from the UTM
func GetIpsRules(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ips.rules", &val, options...)
	return
}

// UpdateIpsRules PUTs the ips.rules value to the UTM
func UpdateIpsRules(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.rules", val, options...)
}

// GetIpsSkipAcks gets the ips.skip_acks value from the UTM
func GetIpsSkipAcks(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ips.skip_acks", &val, options...)
	return
}

// UpdateIpsSkipAcks PUTs the ips.skip_acks value to the UTM
func UpdateIpsSkipAcks(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.skip_acks", val, options...)
}

// GetIpsSmtpServers gets the ips.smtp_servers value from the UTM
func GetIpsSmtpServers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ips.smtp_servers", &val, options...)
	return
}

// UpdateIpsSmtpServers PUTs the ips.smtp_servers value to the UTM
func UpdateIpsSmtpServers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.smtp_servers", val, options...)
}

// GetIpsSnortsettingsMaxQueuedBytes gets the ips.snortsettings.max_queued_bytes value from the UTM
func GetIpsSnortsettingsMaxQueuedBytes(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/ips.snortsettings.max_queued_bytes", &val, options...)
	return
}

// UpdateIpsSnortsettingsMaxQueuedBytes PUTs the ips.snortsettings.max_queued_bytes value to the UTM
func UpdateIpsSnortsettingsMaxQueuedBytes(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.snortsettings.max_queued_bytes", val, options...)
}

// GetIpsSnortsettingsMaxQueuedSegs gets the ips.snortsettings.max_queued_segs value from the UTM
func GetIpsSnortsettingsMaxQueuedSegs(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/ips.snortsettings.max_queued_segs", &val, options...)
	return
}

// UpdateIpsSnortsettingsMaxQueuedSegs PUTs the ips.snortsettings.max_queued_segs value to the UTM
func UpdateIpsSnortsettingsMaxQueuedSegs(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.snortsettings.max_queued_segs", val, options...)
}

// GetIpsSnortsettingsMaxTcp gets the ips.snortsettings.max_tcp value from the UTM
func GetIpsSnortsettingsMaxTcp(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/ips.snortsettings.max_tcp", &val, options...)
	return
}

// UpdateIpsSnortsettingsMaxTcp PUTs the ips.snortsettings.max_tcp value to the UTM
func UpdateIpsSnortsettingsMaxTcp(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.snortsettings.max_tcp", val, options...)
}

// GetIpsSnortsettingsMaxUdp gets the ips.snortsettings.max_udp value from the UTM
func GetIpsSnortsettingsMaxUdp(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/ips.snortsettings.max_udp", &val, options...)
	return
}

// UpdateIpsSnortsettingsMaxUdp PUTs the ips.snortsettings.max_udp value to the UTM
func UpdateIpsSnortsettingsMaxUdp(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.snortsettings.max_udp", val, options...)
}

// GetIpsSnortsettingsMemcap gets the ips.snortsettings.memcap value from the UTM
func GetIpsSnortsettingsMemcap(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/ips.snortsettings.memcap", &val, options...)
	return
}

// UpdateIpsSnortsettingsMemcap PUTs the ips.snortsettings.memcap value to the UTM
func UpdateIpsSnortsettingsMemcap(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.snortsettings.memcap", val, options...)
}

// GetIpsSnortsettingsSearchMethod gets the ips.snortsettings.search_method value from the UTM
func GetIpsSnortsettingsSearchMethod(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ips.snortsettings.search_method", &val, options...)
	return
}

// UpdateIpsSnortsettingsSearchMethod PUTs the ips.snortsettings.search_method value to the UTM
func UpdateIpsSnortsettingsSearchMethod(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.snortsettings.search_method", val, options...)
}

// GetIpsSqlServers gets the ips.sql_servers value from the UTM
func GetIpsSqlServers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ips.sql_servers", &val, options...)
	return
}

// UpdateIpsSqlServers PUTs the ips.sql_servers value to the UTM
func UpdateIpsSqlServers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.sql_servers", val, options...)
}

// GetIpsStatus gets the ips.status value from the UTM
func GetIpsStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ips.status", &val, options...)
	return
}

// UpdateIpsStatus PUTs the ips.status value to the UTM
func UpdateIpsStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ips.status", val, options...)
}

// GetIpsecAdvancedCrlAutoFetching gets the ipsec.advanced.crl_auto_fetching value from the UTM
func GetIpsecAdvancedCrlAutoFetching(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipsec.advanced.crl_auto_fetching", &val, options...)
	return
}

// UpdateIpsecAdvancedCrlAutoFetching PUTs the ipsec.advanced.crl_auto_fetching value to the UTM
func UpdateIpsecAdvancedCrlAutoFetching(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.advanced.crl_auto_fetching", val, options...)
}

// GetIpsecAdvancedCrlStrictPolicy gets the ipsec.advanced.crl_strict_policy value from the UTM
func GetIpsecAdvancedCrlStrictPolicy(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipsec.advanced.crl_strict_policy", &val, options...)
	return
}

// UpdateIpsecAdvancedCrlStrictPolicy PUTs the ipsec.advanced.crl_strict_policy value to the UTM
func UpdateIpsecAdvancedCrlStrictPolicy(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.advanced.crl_strict_policy", val, options...)
}

// GetIpsecAdvancedDeadPeerDetection gets the ipsec.advanced.dead_peer_detection value from the UTM
func GetIpsecAdvancedDeadPeerDetection(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipsec.advanced.dead_peer_detection", &val, options...)
	return
}

// UpdateIpsecAdvancedDeadPeerDetection PUTs the ipsec.advanced.dead_peer_detection value to the UTM
func UpdateIpsecAdvancedDeadPeerDetection(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.advanced.dead_peer_detection", val, options...)
}

// GetIpsecAdvancedIkeDebug gets the ipsec.advanced.ike_debug value from the UTM
func GetIpsecAdvancedIkeDebug(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ipsec.advanced.ike_debug", &val, options...)
	return
}

// UpdateIpsecAdvancedIkeDebug PUTs the ipsec.advanced.ike_debug value to the UTM
func UpdateIpsecAdvancedIkeDebug(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.advanced.ike_debug", val, options...)
}

// GetIpsecAdvancedIkePort gets the ipsec.advanced.ike_port value from the UTM
func GetIpsecAdvancedIkePort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ipsec.advanced.ike_port", &val, options...)
	return
}

// UpdateIpsecAdvancedIkePort PUTs the ipsec.advanced.ike_port value to the UTM
func UpdateIpsecAdvancedIkePort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.advanced.ike_port", val, options...)
}

// GetIpsecAdvancedMetric gets the ipsec.advanced.metric value from the UTM
func GetIpsecAdvancedMetric(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/ipsec.advanced.metric", &val, options...)
	return
}

// UpdateIpsecAdvancedMetric PUTs the ipsec.advanced.metric value to the UTM
func UpdateIpsecAdvancedMetric(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.advanced.metric", val, options...)
}

// GetIpsecAdvancedNatTraversal gets the ipsec.advanced.nat_traversal value from the UTM
func GetIpsecAdvancedNatTraversal(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipsec.advanced.nat_traversal", &val, options...)
	return
}

// UpdateIpsecAdvancedNatTraversal PUTs the ipsec.advanced.nat_traversal value to the UTM
func UpdateIpsecAdvancedNatTraversal(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.advanced.nat_traversal", val, options...)
}

// GetIpsecAdvancedNatTraversalKeepalive gets the ipsec.advanced.nat_traversal_keepalive value from the UTM
func GetIpsecAdvancedNatTraversalKeepalive(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ipsec.advanced.nat_traversal_keepalive", &val, options...)
	return
}

// UpdateIpsecAdvancedNatTraversalKeepalive PUTs the ipsec.advanced.nat_traversal_keepalive value to the UTM
func UpdateIpsecAdvancedNatTraversalKeepalive(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.advanced.nat_traversal_keepalive", val, options...)
}

// GetIpsecAdvancedProbePsk gets the ipsec.advanced.probe_psk value from the UTM
func GetIpsecAdvancedProbePsk(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipsec.advanced.probe_psk", &val, options...)
	return
}

// UpdateIpsecAdvancedProbePsk PUTs the ipsec.advanced.probe_psk value to the UTM
func UpdateIpsecAdvancedProbePsk(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.advanced.probe_psk", val, options...)
}

// GetIpsecAdvancedPskVpnId gets the ipsec.advanced.psk_vpn_id value from the UTM
func GetIpsecAdvancedPskVpnId(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipsec.advanced.psk_vpn_id", &val, options...)
	return
}

// UpdateIpsecAdvancedPskVpnId PUTs the ipsec.advanced.psk_vpn_id value to the UTM
func UpdateIpsecAdvancedPskVpnId(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.advanced.psk_vpn_id", val, options...)
}

// GetIpsecAdvancedPskVpnIdType gets the ipsec.advanced.psk_vpn_id_type value from the UTM
func GetIpsecAdvancedPskVpnIdType(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipsec.advanced.psk_vpn_id_type", &val, options...)
	return
}

// UpdateIpsecAdvancedPskVpnIdType PUTs the ipsec.advanced.psk_vpn_id_type value to the UTM
func UpdateIpsecAdvancedPskVpnIdType(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.advanced.psk_vpn_id_type", val, options...)
}

// GetIpsecConnections gets the ipsec.connections value from the UTM
func GetIpsecConnections(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ipsec.connections", &val, options...)
	return
}

// UpdateIpsecConnections PUTs the ipsec.connections value to the UTM
func UpdateIpsecConnections(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.connections", val, options...)
}

// GetIpsecLocalRsa gets the ipsec.local_rsa value from the UTM
func GetIpsecLocalRsa(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipsec.local_rsa", &val, options...)
	return
}

// UpdateIpsecLocalRsa PUTs the ipsec.local_rsa value to the UTM
func UpdateIpsecLocalRsa(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.local_rsa", val, options...)
}

// GetIpsecLocalX509 gets the ipsec.local_x509 value from the UTM
func GetIpsecLocalX509(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipsec.local_x509", &val, options...)
	return
}

// UpdateIpsecLocalX509 PUTs the ipsec.local_x509 value to the UTM
func UpdateIpsecLocalX509(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.local_x509", val, options...)
}

// GetIpsecStatus gets the ipsec.status value from the UTM
func GetIpsecStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipsec.status", &val, options...)
	return
}

// UpdateIpsecStatus PUTs the ipsec.status value to the UTM
func UpdateIpsecStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipsec.status", val, options...)
}

// GetIpv6AdvancedHopLimit gets the ipv6.advanced.hop_limit value from the UTM
func GetIpv6AdvancedHopLimit(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ipv6.advanced.hop_limit", &val, options...)
	return
}

// UpdateIpv6AdvancedHopLimit PUTs the ipv6.advanced.hop_limit value to the UTM
func UpdateIpv6AdvancedHopLimit(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.advanced.hop_limit", val, options...)
}

// GetIpv6AdvancedMaxInterval gets the ipv6.advanced.max_interval value from the UTM
func GetIpv6AdvancedMaxInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ipv6.advanced.max_interval", &val, options...)
	return
}

// UpdateIpv6AdvancedMaxInterval PUTs the ipv6.advanced.max_interval value to the UTM
func UpdateIpv6AdvancedMaxInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.advanced.max_interval", val, options...)
}

// GetIpv6AdvancedMinInterval gets the ipv6.advanced.min_interval value from the UTM
func GetIpv6AdvancedMinInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ipv6.advanced.min_interval", &val, options...)
	return
}

// UpdateIpv6AdvancedMinInterval PUTs the ipv6.advanced.min_interval value to the UTM
func UpdateIpv6AdvancedMinInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.advanced.min_interval", val, options...)
}

// GetIpv6AdvancedPreference gets the ipv6.advanced.preference value from the UTM
func GetIpv6AdvancedPreference(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipv6.advanced.preference", &val, options...)
	return
}

// UpdateIpv6AdvancedPreference PUTs the ipv6.advanced.preference value to the UTM
func UpdateIpv6AdvancedPreference(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.advanced.preference", val, options...)
}

// GetIpv6AdvancedReachableTime gets the ipv6.advanced.reachable_time value from the UTM
func GetIpv6AdvancedReachableTime(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/ipv6.advanced.reachable_time", &val, options...)
	return
}

// UpdateIpv6AdvancedReachableTime PUTs the ipv6.advanced.reachable_time value to the UTM
func UpdateIpv6AdvancedReachableTime(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.advanced.reachable_time", val, options...)
}

// GetIpv6AdvancedRetransTime gets the ipv6.advanced.retrans_time value from the UTM
func GetIpv6AdvancedRetransTime(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/ipv6.advanced.retrans_time", &val, options...)
	return
}

// UpdateIpv6AdvancedRetransTime PUTs the ipv6.advanced.retrans_time value to the UTM
func UpdateIpv6AdvancedRetransTime(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.advanced.retrans_time", val, options...)
}

// GetIpv6BrokerAuthentication gets the ipv6.broker.authentication value from the UTM
func GetIpv6BrokerAuthentication(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipv6.broker.authentication", &val, options...)
	return
}

// UpdateIpv6BrokerAuthentication PUTs the ipv6.broker.authentication value to the UTM
func UpdateIpv6BrokerAuthentication(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.broker.authentication", val, options...)
}

// GetIpv6BrokerInterface gets the ipv6.broker.interface value from the UTM
func GetIpv6BrokerInterface(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipv6.broker.interface", &val, options...)
	return
}

// UpdateIpv6BrokerInterface PUTs the ipv6.broker.interface value to the UTM
func UpdateIpv6BrokerInterface(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.broker.interface", val, options...)
}

// GetIpv6BrokerPassword gets the ipv6.broker.password value from the UTM
func GetIpv6BrokerPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipv6.broker.password", &val, options...)
	return
}

// UpdateIpv6BrokerPassword PUTs the ipv6.broker.password value to the UTM
func UpdateIpv6BrokerPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.broker.password", val, options...)
}

// GetIpv6BrokerProtocol gets the ipv6.broker.protocol value from the UTM
func GetIpv6BrokerProtocol(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipv6.broker.protocol", &val, options...)
	return
}

// UpdateIpv6BrokerProtocol PUTs the ipv6.broker.protocol value to the UTM
func UpdateIpv6BrokerProtocol(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.broker.protocol", val, options...)
}

// GetIpv6BrokerServer gets the ipv6.broker.server value from the UTM
func GetIpv6BrokerServer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipv6.broker.server", &val, options...)
	return
}

// UpdateIpv6BrokerServer PUTs the ipv6.broker.server value to the UTM
func UpdateIpv6BrokerServer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.broker.server", val, options...)
}

// GetIpv6BrokerStatus gets the ipv6.broker.status value from the UTM
func GetIpv6BrokerStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipv6.broker.status", &val, options...)
	return
}

// UpdateIpv6BrokerStatus PUTs the ipv6.broker.status value to the UTM
func UpdateIpv6BrokerStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.broker.status", val, options...)
}

// GetIpv6BrokerTunnelId gets the ipv6.broker.tunnel_id value from the UTM
func GetIpv6BrokerTunnelId(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipv6.broker.tunnel_id", &val, options...)
	return
}

// UpdateIpv6BrokerTunnelId PUTs the ipv6.broker.tunnel_id value to the UTM
func UpdateIpv6BrokerTunnelId(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.broker.tunnel_id", val, options...)
}

// GetIpv6BrokerUsername gets the ipv6.broker.username value from the UTM
func GetIpv6BrokerUsername(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipv6.broker.username", &val, options...)
	return
}

// UpdateIpv6BrokerUsername PUTs the ipv6.broker.username value to the UTM
func UpdateIpv6BrokerUsername(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.broker.username", val, options...)
}

// GetIpv6Nat64Address gets the ipv6.nat64.address value from the UTM
func GetIpv6Nat64Address(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipv6.nat64.address", &val, options...)
	return
}

// UpdateIpv6Nat64Address PUTs the ipv6.nat64.address value to the UTM
func UpdateIpv6Nat64Address(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.nat64.address", val, options...)
}

// GetIpv6Nat64Dns64V6Only gets the ipv6.nat64.dns64_v6only value from the UTM
func GetIpv6Nat64Dns64V6Only(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipv6.nat64.dns64_v6only", &val, options...)
	return
}

// UpdateIpv6Nat64Dns64V6Only PUTs the ipv6.nat64.dns64_v6only value to the UTM
func UpdateIpv6Nat64Dns64V6Only(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.nat64.dns64_v6only", val, options...)
}

// GetIpv6Nat64Prefix gets the ipv6.nat64.prefix value from the UTM
func GetIpv6Nat64Prefix(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipv6.nat64.prefix", &val, options...)
	return
}

// UpdateIpv6Nat64Prefix PUTs the ipv6.nat64.prefix value to the UTM
func UpdateIpv6Nat64Prefix(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.nat64.prefix", val, options...)
}

// GetIpv6Nat64Status gets the ipv6.nat64.status value from the UTM
func GetIpv6Nat64Status(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipv6.nat64.status", &val, options...)
	return
}

// UpdateIpv6Nat64Status PUTs the ipv6.nat64.status value to the UTM
func UpdateIpv6Nat64Status(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.nat64.status", val, options...)
}

// GetIpv6Prefer gets the ipv6.prefer value from the UTM
func GetIpv6Prefer(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipv6.prefer", &val, options...)
	return
}

// UpdateIpv6Prefer PUTs the ipv6.prefer value to the UTM
func UpdateIpv6Prefer(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.prefer", val, options...)
}

// GetIpv6Prefixes gets the ipv6.prefixes value from the UTM
func GetIpv6Prefixes(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ipv6.prefixes", &val, options...)
	return
}

// UpdateIpv6Prefixes PUTs the ipv6.prefixes value to the UTM
func UpdateIpv6Prefixes(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.prefixes", val, options...)
}

// GetIpv6Renumbering gets the ipv6.renumbering value from the UTM
func GetIpv6Renumbering(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipv6.renumbering", &val, options...)
	return
}

// UpdateIpv6Renumbering PUTs the ipv6.renumbering value to the UTM
func UpdateIpv6Renumbering(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.renumbering", val, options...)
}

// GetIpv6Six2FourInterface gets the ipv6.six2four.interface value from the UTM
func GetIpv6Six2FourInterface(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipv6.six2four.interface", &val, options...)
	return
}

// UpdateIpv6Six2FourInterface PUTs the ipv6.six2four.interface value to the UTM
func UpdateIpv6Six2FourInterface(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.six2four.interface", val, options...)
}

// GetIpv6Six2FourServer gets the ipv6.six2four.server value from the UTM
func GetIpv6Six2FourServer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ipv6.six2four.server", &val, options...)
	return
}

// UpdateIpv6Six2FourServer PUTs the ipv6.six2four.server value to the UTM
func UpdateIpv6Six2FourServer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.six2four.server", val, options...)
}

// GetIpv6Six2FourStatus gets the ipv6.six2four.status value from the UTM
func GetIpv6Six2FourStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipv6.six2four.status", &val, options...)
	return
}

// UpdateIpv6Six2FourStatus PUTs the ipv6.six2four.status value to the UTM
func UpdateIpv6Six2FourStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.six2four.status", val, options...)
}

// GetIpv6Status gets the ipv6.status value from the UTM
func GetIpv6Status(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ipv6.status", &val, options...)
	return
}

// UpdateIpv6Status PUTs the ipv6.status value to the UTM
func UpdateIpv6Status(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ipv6.status", val, options...)
}

// GetLicensingActiveIps gets the licensing.active_ips value from the UTM
func GetLicensingActiveIps(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/licensing.active_ips", &val, options...)
	return
}

// UpdateLicensingActiveIps PUTs the licensing.active_ips value to the UTM
func UpdateLicensingActiveIps(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/licensing.active_ips", val, options...)
}

// GetLicensingLicense gets the licensing.license value from the UTM
func GetLicensingLicense(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/licensing.license", &val, options...)
	return
}

// UpdateLicensingLicense PUTs the licensing.license value to the UTM
func UpdateLicensingLicense(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/licensing.license", val, options...)
}

// GetLicensingUserLimitExceeded gets the licensing.user_limit_exceeded value from the UTM
func GetLicensingUserLimitExceeded(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/licensing.user_limit_exceeded", &val, options...)
	return
}

// UpdateLicensingUserLimitExceeded PUTs the licensing.user_limit_exceeded value to the UTM
func UpdateLicensingUserLimitExceeded(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/licensing.user_limit_exceeded", val, options...)
}

// GetLinkAggregationGroups gets the link_aggregation.groups value from the UTM
func GetLinkAggregationGroups(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/link_aggregation.groups", &val, options...)
	return
}

// UpdateLinkAggregationGroups PUTs the link_aggregation.groups value to the UTM
func UpdateLinkAggregationGroups(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/link_aggregation.groups", val, options...)
}

// GetLoadbalanceHttpErrorCode gets the loadbalance.http_error_code value from the UTM
func GetLoadbalanceHttpErrorCode(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/loadbalance.http_error_code", &val, options...)
	return
}

// UpdateLoadbalanceHttpErrorCode PUTs the loadbalance.http_error_code value to the UTM
func UpdateLoadbalanceHttpErrorCode(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/loadbalance.http_error_code", val, options...)
}

// GetLoadbalanceRules gets the loadbalance.rules value from the UTM
func GetLoadbalanceRules(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/loadbalance.rules", &val, options...)
	return
}

// UpdateLoadbalanceRules PUTs the loadbalance.rules value to the UTM
func UpdateLoadbalanceRules(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/loadbalance.rules", val, options...)
}

// GetLogfilesLocalActionOne gets the logfiles.local.action_one value from the UTM
func GetLogfilesLocalActionOne(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/logfiles.local.action_one", &val, options...)
	return
}

// UpdateLogfilesLocalActionOne PUTs the logfiles.local.action_one value to the UTM
func UpdateLogfilesLocalActionOne(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.local.action_one", val, options...)
}

// GetLogfilesLocalActionThree gets the logfiles.local.action_three value from the UTM
func GetLogfilesLocalActionThree(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/logfiles.local.action_three", &val, options...)
	return
}

// UpdateLogfilesLocalActionThree PUTs the logfiles.local.action_three value to the UTM
func UpdateLogfilesLocalActionThree(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.local.action_three", val, options...)
}

// GetLogfilesLocalActionTwo gets the logfiles.local.action_two value from the UTM
func GetLogfilesLocalActionTwo(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/logfiles.local.action_two", &val, options...)
	return
}

// UpdateLogfilesLocalActionTwo PUTs the logfiles.local.action_two value to the UTM
func UpdateLogfilesLocalActionTwo(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.local.action_two", val, options...)
}

// GetLogfilesLocalDeleteAfterDays gets the logfiles.local.delete_after_days value from the UTM
func GetLogfilesLocalDeleteAfterDays(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/logfiles.local.delete_after_days", &val, options...)
	return
}

// UpdateLogfilesLocalDeleteAfterDays PUTs the logfiles.local.delete_after_days value to the UTM
func UpdateLogfilesLocalDeleteAfterDays(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.local.delete_after_days", val, options...)
}

// GetLogfilesLocalPercentageOne gets the logfiles.local.percentage_one value from the UTM
func GetLogfilesLocalPercentageOne(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/logfiles.local.percentage_one", &val, options...)
	return
}

// UpdateLogfilesLocalPercentageOne PUTs the logfiles.local.percentage_one value to the UTM
func UpdateLogfilesLocalPercentageOne(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.local.percentage_one", val, options...)
}

// GetLogfilesLocalPercentageThree gets the logfiles.local.percentage_three value from the UTM
func GetLogfilesLocalPercentageThree(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/logfiles.local.percentage_three", &val, options...)
	return
}

// UpdateLogfilesLocalPercentageThree PUTs the logfiles.local.percentage_three value to the UTM
func UpdateLogfilesLocalPercentageThree(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.local.percentage_three", val, options...)
}

// GetLogfilesLocalPercentageTwo gets the logfiles.local.percentage_two value from the UTM
func GetLogfilesLocalPercentageTwo(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/logfiles.local.percentage_two", &val, options...)
	return
}

// UpdateLogfilesLocalPercentageTwo PUTs the logfiles.local.percentage_two value to the UTM
func UpdateLogfilesLocalPercentageTwo(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.local.percentage_two", val, options...)
}

// GetLogfilesLocalStatus gets the logfiles.local.status value from the UTM
func GetLogfilesLocalStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/logfiles.local.status", &val, options...)
	return
}

// UpdateLogfilesLocalStatus PUTs the logfiles.local.status value to the UTM
func UpdateLogfilesLocalStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.local.status", val, options...)
}

// GetLogfilesLocalSyslogMaxConnections gets the logfiles.local.syslog_max_connections value from the UTM
func GetLogfilesLocalSyslogMaxConnections(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/logfiles.local.syslog_max_connections", &val, options...)
	return
}

// UpdateLogfilesLocalSyslogMaxConnections PUTs the logfiles.local.syslog_max_connections value to the UTM
func UpdateLogfilesLocalSyslogMaxConnections(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.local.syslog_max_connections", val, options...)
}

// GetLogfilesRemoteFtpService gets the logfiles.remote.ftp_service value from the UTM
func GetLogfilesRemoteFtpService(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/logfiles.remote.ftp_service", &val, options...)
	return
}

// UpdateLogfilesRemoteFtpService PUTs the logfiles.remote.ftp_service value to the UTM
func UpdateLogfilesRemoteFtpService(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.remote.ftp_service", val, options...)
}

// GetLogfilesRemoteHost gets the logfiles.remote.host value from the UTM
func GetLogfilesRemoteHost(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/logfiles.remote.host", &val, options...)
	return
}

// UpdateLogfilesRemoteHost PUTs the logfiles.remote.host value to the UTM
func UpdateLogfilesRemoteHost(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.remote.host", val, options...)
}

// GetLogfilesRemotePass gets the logfiles.remote.pass value from the UTM
func GetLogfilesRemotePass(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/logfiles.remote.pass", &val, options...)
	return
}

// UpdateLogfilesRemotePass PUTs the logfiles.remote.pass value to the UTM
func UpdateLogfilesRemotePass(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.remote.pass", val, options...)
}

// GetLogfilesRemotePath gets the logfiles.remote.path value from the UTM
func GetLogfilesRemotePath(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/logfiles.remote.path", &val, options...)
	return
}

// UpdateLogfilesRemotePath PUTs the logfiles.remote.path value to the UTM
func UpdateLogfilesRemotePath(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.remote.path", val, options...)
}

// GetLogfilesRemoteSmbMaxProtocolLevel gets the logfiles.remote.smb_max_protocol_level value from the UTM
func GetLogfilesRemoteSmbMaxProtocolLevel(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/logfiles.remote.smb_max_protocol_level", &val, options...)
	return
}

// UpdateLogfilesRemoteSmbMaxProtocolLevel PUTs the logfiles.remote.smb_max_protocol_level value to the UTM
func UpdateLogfilesRemoteSmbMaxProtocolLevel(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.remote.smb_max_protocol_level", val, options...)
}

// GetLogfilesRemoteSmbWorkgroup gets the logfiles.remote.smb_workgroup value from the UTM
func GetLogfilesRemoteSmbWorkgroup(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/logfiles.remote.smb_workgroup", &val, options...)
	return
}

// UpdateLogfilesRemoteSmbWorkgroup PUTs the logfiles.remote.smb_workgroup value to the UTM
func UpdateLogfilesRemoteSmbWorkgroup(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.remote.smb_workgroup", val, options...)
}

// GetLogfilesRemoteSmtpAddress gets the logfiles.remote.smtp_address value from the UTM
func GetLogfilesRemoteSmtpAddress(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/logfiles.remote.smtp_address", &val, options...)
	return
}

// UpdateLogfilesRemoteSmtpAddress PUTs the logfiles.remote.smtp_address value to the UTM
func UpdateLogfilesRemoteSmtpAddress(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.remote.smtp_address", val, options...)
}

// GetLogfilesRemoteStatus gets the logfiles.remote.status value from the UTM
func GetLogfilesRemoteStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/logfiles.remote.status", &val, options...)
	return
}

// UpdateLogfilesRemoteStatus PUTs the logfiles.remote.status value to the UTM
func UpdateLogfilesRemoteStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.remote.status", val, options...)
}

// GetLogfilesRemoteType gets the logfiles.remote.type value from the UTM
func GetLogfilesRemoteType(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/logfiles.remote.type", &val, options...)
	return
}

// UpdateLogfilesRemoteType PUTs the logfiles.remote.type value to the UTM
func UpdateLogfilesRemoteType(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.remote.type", val, options...)
}

// GetLogfilesRemoteUser gets the logfiles.remote.user value from the UTM
func GetLogfilesRemoteUser(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/logfiles.remote.user", &val, options...)
	return
}

// UpdateLogfilesRemoteUser PUTs the logfiles.remote.user value to the UTM
func UpdateLogfilesRemoteUser(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/logfiles.remote.user", val, options...)
}

// GetMasqRules gets the masq.rules value from the UTM
func GetMasqRules(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/masq.rules", &val, options...)
	return
}

// UpdateMasqRules PUTs the masq.rules value to the UTM
func UpdateMasqRules(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/masq.rules", val, options...)
}

// GetMigrationAccessToken gets the migration.access_token value from the UTM
func GetMigrationAccessToken(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/migration.access_token", &val, options...)
	return
}

// UpdateMigrationAccessToken PUTs the migration.access_token value to the UTM
func UpdateMigrationAccessToken(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/migration.access_token", val, options...)
}

// GetMigrationLocalOverride gets the migration.local_override value from the UTM
func GetMigrationLocalOverride(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/migration.local_override", &val, options...)
	return
}

// UpdateMigrationLocalOverride PUTs the migration.local_override value to the UTM
func UpdateMigrationLocalOverride(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/migration.local_override", val, options...)
}

// GetMigrationRefreshToken gets the migration.refresh_token value from the UTM
func GetMigrationRefreshToken(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/migration.refresh_token", &val, options...)
	return
}

// UpdateMigrationRefreshToken PUTs the migration.refresh_token value to the UTM
func UpdateMigrationRefreshToken(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/migration.refresh_token", val, options...)
}

// GetMigrationTabVisibility gets the migration.tab_visibility value from the UTM
func GetMigrationTabVisibility(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/migration.tab_visibility", &val, options...)
	return
}

// UpdateMigrationTabVisibility PUTs the migration.tab_visibility value to the UTM
func UpdateMigrationTabVisibility(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/migration.tab_visibility", val, options...)
}

// GetMigrationToolsetVersion gets the migration.toolset_version value from the UTM
func GetMigrationToolsetVersion(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/migration.toolset_version", &val, options...)
	return
}

// UpdateMigrationToolsetVersion PUTs the migration.toolset_version value to the UTM
func UpdateMigrationToolsetVersion(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/migration.toolset_version", val, options...)
}

// GetMigrationUtmVersion gets the migration.utm_version value from the UTM
func GetMigrationUtmVersion(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/migration.utm_version", &val, options...)
	return
}

// UpdateMigrationUtmVersion PUTs the migration.utm_version value to the UTM
func UpdateMigrationUtmVersion(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/migration.utm_version", val, options...)
}

// GetMobileControlCa gets the mobile_control.ca value from the UTM
func GetMobileControlCa(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/mobile_control.ca", &val, options...)
	return
}

// UpdateMobileControlCa PUTs the mobile_control.ca value to the UTM
func UpdateMobileControlCa(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.ca", val, options...)
}

// GetMobileControlConfigCisco gets the mobile_control.config.cisco value from the UTM
func GetMobileControlConfigCisco(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/mobile_control.config.cisco", &val, options...)
	return
}

// UpdateMobileControlConfigCisco PUTs the mobile_control.config.cisco value to the UTM
func UpdateMobileControlConfigCisco(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.config.cisco", val, options...)
}

// GetMobileControlConfigEapMethod gets the mobile_control.config.eap_method value from the UTM
func GetMobileControlConfigEapMethod(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/mobile_control.config.eap_method", &val, options...)
	return
}

// UpdateMobileControlConfigEapMethod PUTs the mobile_control.config.eap_method value to the UTM
func UpdateMobileControlConfigEapMethod(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.config.eap_method", val, options...)
}

// GetMobileControlConfigForcePush gets the mobile_control.config.force_push value from the UTM
func GetMobileControlConfigForcePush(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/mobile_control.config.force_push", &val, options...)
	return
}

// UpdateMobileControlConfigForcePush PUTs the mobile_control.config.force_push value to the UTM
func UpdateMobileControlConfigForcePush(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.config.force_push", val, options...)
}

// GetMobileControlConfigL2Tp gets the mobile_control.config.l2tp value from the UTM
func GetMobileControlConfigL2Tp(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/mobile_control.config.l2tp", &val, options...)
	return
}

// UpdateMobileControlConfigL2Tp PUTs the mobile_control.config.l2tp value to the UTM
func UpdateMobileControlConfigL2Tp(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.config.l2tp", val, options...)
}

// GetMobileControlConfigWifiNetworks gets the mobile_control.config.wifi_networks value from the UTM
func GetMobileControlConfigWifiNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/mobile_control.config.wifi_networks", &val, options...)
	return
}

// UpdateMobileControlConfigWifiNetworks PUTs the mobile_control.config.wifi_networks value to the UTM
func UpdateMobileControlConfigWifiNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.config.wifi_networks", val, options...)
}

// GetMobileControlCustomer gets the mobile_control.customer value from the UTM
func GetMobileControlCustomer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/mobile_control.customer", &val, options...)
	return
}

// UpdateMobileControlCustomer PUTs the mobile_control.customer value to the UTM
func UpdateMobileControlCustomer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.customer", val, options...)
}

// GetMobileControlDebug gets the mobile_control.debug value from the UTM
func GetMobileControlDebug(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/mobile_control.debug", &val, options...)
	return
}

// UpdateMobileControlDebug PUTs the mobile_control.debug value to the UTM
func UpdateMobileControlDebug(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.debug", val, options...)
}

// GetMobileControlNacCisco gets the mobile_control.nac.cisco value from the UTM
func GetMobileControlNacCisco(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/mobile_control.nac.cisco", &val, options...)
	return
}

// UpdateMobileControlNacCisco PUTs the mobile_control.nac.cisco value to the UTM
func UpdateMobileControlNacCisco(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.nac.cisco", val, options...)
}

// GetMobileControlNacDenyAllVpn gets the mobile_control.nac.deny_all_vpn value from the UTM
func GetMobileControlNacDenyAllVpn(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/mobile_control.nac.deny_all_vpn", &val, options...)
	return
}

// UpdateMobileControlNacDenyAllVpn PUTs the mobile_control.nac.deny_all_vpn value to the UTM
func UpdateMobileControlNacDenyAllVpn(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.nac.deny_all_vpn", val, options...)
}

// GetMobileControlNacL2Tp gets the mobile_control.nac.l2tp value from the UTM
func GetMobileControlNacL2Tp(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/mobile_control.nac.l2tp", &val, options...)
	return
}

// UpdateMobileControlNacL2Tp PUTs the mobile_control.nac.l2tp value to the UTM
func UpdateMobileControlNacL2Tp(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.nac.l2tp", val, options...)
}

// GetMobileControlNacMacsAllowed gets the mobile_control.nac.macs_allowed value from the UTM
func GetMobileControlNacMacsAllowed(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/mobile_control.nac.macs_allowed", &val, options...)
	return
}

// UpdateMobileControlNacMacsAllowed PUTs the mobile_control.nac.macs_allowed value to the UTM
func UpdateMobileControlNacMacsAllowed(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.nac.macs_allowed", val, options...)
}

// GetMobileControlNacMacsDenied gets the mobile_control.nac.macs_denied value from the UTM
func GetMobileControlNacMacsDenied(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/mobile_control.nac.macs_denied", &val, options...)
	return
}

// UpdateMobileControlNacMacsDenied PUTs the mobile_control.nac.macs_denied value to the UTM
func UpdateMobileControlNacMacsDenied(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.nac.macs_denied", val, options...)
}

// GetMobileControlNacPollInterval gets the mobile_control.nac.poll_interval value from the UTM
func GetMobileControlNacPollInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/mobile_control.nac.poll_interval", &val, options...)
	return
}

// UpdateMobileControlNacPollInterval PUTs the mobile_control.nac.poll_interval value to the UTM
func UpdateMobileControlNacPollInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.nac.poll_interval", val, options...)
}

// GetMobileControlNacUsersDenied gets the mobile_control.nac.users_denied value from the UTM
func GetMobileControlNacUsersDenied(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/mobile_control.nac.users_denied", &val, options...)
	return
}

// UpdateMobileControlNacUsersDenied PUTs the mobile_control.nac.users_denied value to the UTM
func UpdateMobileControlNacUsersDenied(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.nac.users_denied", val, options...)
}

// GetMobileControlNacWifiNetworks gets the mobile_control.nac.wifi_networks value from the UTM
func GetMobileControlNacWifiNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/mobile_control.nac.wifi_networks", &val, options...)
	return
}

// UpdateMobileControlNacWifiNetworks PUTs the mobile_control.nac.wifi_networks value to the UTM
func UpdateMobileControlNacWifiNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.nac.wifi_networks", val, options...)
}

// GetMobileControlPassword gets the mobile_control.password value from the UTM
func GetMobileControlPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/mobile_control.password", &val, options...)
	return
}

// UpdateMobileControlPassword PUTs the mobile_control.password value to the UTM
func UpdateMobileControlPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.password", val, options...)
}

// GetMobileControlServer gets the mobile_control.server value from the UTM
func GetMobileControlServer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/mobile_control.server", &val, options...)
	return
}

// UpdateMobileControlServer PUTs the mobile_control.server value to the UTM
func UpdateMobileControlServer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.server", val, options...)
}

// GetMobileControlStatus gets the mobile_control.status value from the UTM
func GetMobileControlStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/mobile_control.status", &val, options...)
	return
}

// UpdateMobileControlStatus PUTs the mobile_control.status value to the UTM
func UpdateMobileControlStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.status", val, options...)
}

// GetMobileControlUsername gets the mobile_control.username value from the UTM
func GetMobileControlUsername(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/mobile_control.username", &val, options...)
	return
}

// UpdateMobileControlUsername PUTs the mobile_control.username value to the UTM
func UpdateMobileControlUsername(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/mobile_control.username", val, options...)
}

// GetNatEnableCacheAutonat gets the nat.enable_cache_autonat value from the UTM
func GetNatEnableCacheAutonat(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/nat.enable_cache_autonat", &val, options...)
	return
}

// UpdateNatEnableCacheAutonat PUTs the nat.enable_cache_autonat value to the UTM
func UpdateNatEnableCacheAutonat(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/nat.enable_cache_autonat", val, options...)
}

// GetNatRules gets the nat.rules value from the UTM
func GetNatRules(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/nat.rules", &val, options...)
	return
}

// UpdateNatRules PUTs the nat.rules value to the UTM
func UpdateNatRules(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/nat.rules", val, options...)
}

// GetNotificationsDeviceInfo gets the notifications.device_info value from the UTM
func GetNotificationsDeviceInfo(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/notifications.device_info", &val, options...)
	return
}

// UpdateNotificationsDeviceInfo PUTs the notifications.device_info value to the UTM
func UpdateNotificationsDeviceInfo(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.device_info", val, options...)
}

// GetNotificationsLimiting gets the notifications.limiting value from the UTM
func GetNotificationsLimiting(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/notifications.limiting", &val, options...)
	return
}

// UpdateNotificationsLimiting PUTs the notifications.limiting value to the UTM
func UpdateNotificationsLimiting(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.limiting", val, options...)
}

// GetNotificationsOverlay gets the notifications.overlay value from the UTM
func GetNotificationsOverlay(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/notifications.overlay", &val, options...)
	return
}

// UpdateNotificationsOverlay PUTs the notifications.overlay value to the UTM
func UpdateNotificationsOverlay(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.overlay", val, options...)
}

// GetNotificationsRebootReason gets the notifications.reboot_reason value from the UTM
func GetNotificationsRebootReason(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/notifications.reboot_reason", &val, options...)
	return
}

// UpdateNotificationsRebootReason PUTs the notifications.reboot_reason value to the UTM
func UpdateNotificationsRebootReason(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.reboot_reason", val, options...)
}

// GetNotificationsRecipients gets the notifications.recipients value from the UTM
func GetNotificationsRecipients(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/notifications.recipients", &val, options...)
	return
}

// UpdateNotificationsRecipients PUTs the notifications.recipients value to the UTM
func UpdateNotificationsRecipients(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.recipients", val, options...)
}

// GetNotificationsSender gets the notifications.sender value from the UTM
func GetNotificationsSender(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/notifications.sender", &val, options...)
	return
}

// UpdateNotificationsSender PUTs the notifications.sender value to the UTM
func UpdateNotificationsSender(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.sender", val, options...)
}

// GetNotificationsSmtpAuthentication gets the notifications.smtp.authentication value from the UTM
func GetNotificationsSmtpAuthentication(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/notifications.smtp.authentication", &val, options...)
	return
}

// UpdateNotificationsSmtpAuthentication PUTs the notifications.smtp.authentication value to the UTM
func UpdateNotificationsSmtpAuthentication(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.smtp.authentication", val, options...)
}

// GetNotificationsSmtpPassword gets the notifications.smtp.password value from the UTM
func GetNotificationsSmtpPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/notifications.smtp.password", &val, options...)
	return
}

// UpdateNotificationsSmtpPassword PUTs the notifications.smtp.password value to the UTM
func UpdateNotificationsSmtpPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.smtp.password", val, options...)
}

// GetNotificationsSmtpPort gets the notifications.smtp.port value from the UTM
func GetNotificationsSmtpPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/notifications.smtp.port", &val, options...)
	return
}

// UpdateNotificationsSmtpPort PUTs the notifications.smtp.port value to the UTM
func UpdateNotificationsSmtpPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.smtp.port", val, options...)
}

// GetNotificationsSmtpServer gets the notifications.smtp.server value from the UTM
func GetNotificationsSmtpServer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/notifications.smtp.server", &val, options...)
	return
}

// UpdateNotificationsSmtpServer PUTs the notifications.smtp.server value to the UTM
func UpdateNotificationsSmtpServer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.smtp.server", val, options...)
}

// GetNotificationsSmtpStatus gets the notifications.smtp.status value from the UTM
func GetNotificationsSmtpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/notifications.smtp.status", &val, options...)
	return
}

// UpdateNotificationsSmtpStatus PUTs the notifications.smtp.status value to the UTM
func UpdateNotificationsSmtpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.smtp.status", val, options...)
}

// GetNotificationsSmtpTls gets the notifications.smtp.tls value from the UTM
func GetNotificationsSmtpTls(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/notifications.smtp.tls", &val, options...)
	return
}

// UpdateNotificationsSmtpTls PUTs the notifications.smtp.tls value to the UTM
func UpdateNotificationsSmtpTls(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.smtp.tls", val, options...)
}

// GetNotificationsSmtpUsername gets the notifications.smtp.username value from the UTM
func GetNotificationsSmtpUsername(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/notifications.smtp.username", &val, options...)
	return
}

// UpdateNotificationsSmtpUsername PUTs the notifications.smtp.username value to the UTM
func UpdateNotificationsSmtpUsername(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/notifications.smtp.username", val, options...)
}

// GetNtpAllowedNetworks gets the ntp.allowed_networks value from the UTM
func GetNtpAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/ntp.allowed_networks", &val, options...)
	return
}

// UpdateNtpAllowedNetworks PUTs the ntp.allowed_networks value to the UTM
func UpdateNtpAllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ntp.allowed_networks", val, options...)
}

// GetNtpServers gets the ntp.servers value from the UTM
func GetNtpServers(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/ntp.servers", &val, options...)
	return
}

// UpdateNtpServers PUTs the ntp.servers value to the UTM
func UpdateNtpServers(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ntp.servers", val, options...)
}

// GetNtpStatus gets the ntp.status value from the UTM
func GetNtpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ntp.status", &val, options...)
	return
}

// UpdateNtpStatus PUTs the ntp.status value to the UTM
func UpdateNtpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ntp.status", val, options...)
}

// GetPacketfilterAdvancedBlockInvalidCtPackets gets the packetfilter.advanced.block_invalid_ct_packets value from the UTM
func GetPacketfilterAdvancedBlockInvalidCtPackets(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.block_invalid_ct_packets", &val, options...)
	return
}

// UpdatePacketfilterAdvancedBlockInvalidCtPackets PUTs the packetfilter.advanced.block_invalid_ct_packets value to the UTM
func UpdatePacketfilterAdvancedBlockInvalidCtPackets(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.block_invalid_ct_packets", val, options...)
}

// GetPacketfilterAdvancedCheckPacketLength gets the packetfilter.advanced.check_packet_length value from the UTM
func GetPacketfilterAdvancedCheckPacketLength(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.check_packet_length", &val, options...)
	return
}

// UpdatePacketfilterAdvancedCheckPacketLength PUTs the packetfilter.advanced.check_packet_length value to the UTM
func UpdatePacketfilterAdvancedCheckPacketLength(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.check_packet_length", val, options...)
}

// GetPacketfilterAdvancedConntrackHelpers gets the packetfilter.advanced.conntrack_helpers value from the UTM
func GetPacketfilterAdvancedConntrackHelpers(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.conntrack_helpers", &val, options...)
	return
}

// UpdatePacketfilterAdvancedConntrackHelpers PUTs the packetfilter.advanced.conntrack_helpers value to the UTM
func UpdatePacketfilterAdvancedConntrackHelpers(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.conntrack_helpers", val, options...)
}

// GetPacketfilterAdvancedFtpPorts gets the packetfilter.advanced.ftp_ports value from the UTM
func GetPacketfilterAdvancedFtpPorts(client sophos.ClientInterface, options ...sophos.Option) (val []int64, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.ftp_ports", &val, options...)
	return
}

// UpdatePacketfilterAdvancedFtpPorts PUTs the packetfilter.advanced.ftp_ports value to the UTM
func UpdatePacketfilterAdvancedFtpPorts(client sophos.ClientInterface, val []int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.ftp_ports", val, options...)
}

// GetPacketfilterAdvancedIpSetMax gets the packetfilter.advanced.ip_set_max value from the UTM
func GetPacketfilterAdvancedIpSetMax(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.ip_set_max", &val, options...)
	return
}

// UpdatePacketfilterAdvancedIpSetMax PUTs the packetfilter.advanced.ip_set_max value to the UTM
func UpdatePacketfilterAdvancedIpSetMax(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.ip_set_max", val, options...)
}

// GetPacketfilterAdvancedLogBroadcasts gets the packetfilter.advanced.log_broadcasts value from the UTM
func GetPacketfilterAdvancedLogBroadcasts(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.log_broadcasts", &val, options...)
	return
}

// UpdatePacketfilterAdvancedLogBroadcasts PUTs the packetfilter.advanced.log_broadcasts value to the UTM
func UpdatePacketfilterAdvancedLogBroadcasts(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.log_broadcasts", val, options...)
}

// GetPacketfilterAdvancedLogDnsRequests gets the packetfilter.advanced.log_dns_requests value from the UTM
func GetPacketfilterAdvancedLogDnsRequests(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.log_dns_requests", &val, options...)
	return
}

// UpdatePacketfilterAdvancedLogDnsRequests PUTs the packetfilter.advanced.log_dns_requests value to the UTM
func UpdatePacketfilterAdvancedLogDnsRequests(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.log_dns_requests", val, options...)
}

// GetPacketfilterAdvancedLogFtpData gets the packetfilter.advanced.log_ftp_data value from the UTM
func GetPacketfilterAdvancedLogFtpData(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.log_ftp_data", &val, options...)
	return
}

// UpdatePacketfilterAdvancedLogFtpData PUTs the packetfilter.advanced.log_ftp_data value to the UTM
func UpdatePacketfilterAdvancedLogFtpData(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.log_ftp_data", val, options...)
}

// GetPacketfilterAdvancedLogInvalid gets the packetfilter.advanced.log_invalid value from the UTM
func GetPacketfilterAdvancedLogInvalid(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.log_invalid", &val, options...)
	return
}

// UpdatePacketfilterAdvancedLogInvalid PUTs the packetfilter.advanced.log_invalid value to the UTM
func UpdatePacketfilterAdvancedLogInvalid(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.log_invalid", val, options...)
}

// GetPacketfilterAdvancedLogMcast gets the packetfilter.advanced.log_mcast value from the UTM
func GetPacketfilterAdvancedLogMcast(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.log_mcast", &val, options...)
	return
}

// UpdatePacketfilterAdvancedLogMcast PUTs the packetfilter.advanced.log_mcast value to the UTM
func UpdatePacketfilterAdvancedLogMcast(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.log_mcast", val, options...)
}

// GetPacketfilterAdvancedLogStrictTcpState gets the packetfilter.advanced.log_strict_tcp_state value from the UTM
func GetPacketfilterAdvancedLogStrictTcpState(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.log_strict_tcp_state", &val, options...)
	return
}

// UpdatePacketfilterAdvancedLogStrictTcpState PUTs the packetfilter.advanced.log_strict_tcp_state value to the UTM
func UpdatePacketfilterAdvancedLogStrictTcpState(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.log_strict_tcp_state", val, options...)
}

// GetPacketfilterAdvancedNoErrorReplay gets the packetfilter.advanced.no_error_replay value from the UTM
func GetPacketfilterAdvancedNoErrorReplay(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.no_error_replay", &val, options...)
	return
}

// UpdatePacketfilterAdvancedNoErrorReplay PUTs the packetfilter.advanced.no_error_replay value to the UTM
func UpdatePacketfilterAdvancedNoErrorReplay(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.no_error_replay", val, options...)
}

// GetPacketfilterAdvancedOptimizeIpset gets the packetfilter.advanced.optimize.ipset value from the UTM
func GetPacketfilterAdvancedOptimizeIpset(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.optimize.ipset", &val, options...)
	return
}

// UpdatePacketfilterAdvancedOptimizeIpset PUTs the packetfilter.advanced.optimize.ipset value to the UTM
func UpdatePacketfilterAdvancedOptimizeIpset(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.optimize.ipset", val, options...)
}

// GetPacketfilterAdvancedOptimizePorts gets the packetfilter.advanced.optimize.ports value from the UTM
func GetPacketfilterAdvancedOptimizePorts(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.optimize.ports", &val, options...)
	return
}

// UpdatePacketfilterAdvancedOptimizePorts PUTs the packetfilter.advanced.optimize.ports value to the UTM
func UpdatePacketfilterAdvancedOptimizePorts(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.optimize.ports", val, options...)
}

// GetPacketfilterAdvancedSpoofProtection gets the packetfilter.advanced.spoof_protection value from the UTM
func GetPacketfilterAdvancedSpoofProtection(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.spoof_protection", &val, options...)
	return
}

// UpdatePacketfilterAdvancedSpoofProtection PUTs the packetfilter.advanced.spoof_protection value to the UTM
func UpdatePacketfilterAdvancedSpoofProtection(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.spoof_protection", val, options...)
}

// GetPacketfilterAdvancedStrictTcpState gets the packetfilter.advanced.strict_tcp_state value from the UTM
func GetPacketfilterAdvancedStrictTcpState(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.strict_tcp_state", &val, options...)
	return
}

// UpdatePacketfilterAdvancedStrictTcpState PUTs the packetfilter.advanced.strict_tcp_state value to the UTM
func UpdatePacketfilterAdvancedStrictTcpState(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.strict_tcp_state", val, options...)
}

// GetPacketfilterAdvancedTcpMaxRetrans gets the packetfilter.advanced.tcp_max_retrans value from the UTM
func GetPacketfilterAdvancedTcpMaxRetrans(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.tcp_max_retrans", &val, options...)
	return
}

// UpdatePacketfilterAdvancedTcpMaxRetrans PUTs the packetfilter.advanced.tcp_max_retrans value to the UTM
func UpdatePacketfilterAdvancedTcpMaxRetrans(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.tcp_max_retrans", val, options...)
}

// GetPacketfilterAdvancedTcpWindowScaling gets the packetfilter.advanced.tcp_window_scaling value from the UTM
func GetPacketfilterAdvancedTcpWindowScaling(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/packetfilter.advanced.tcp_window_scaling", &val, options...)
	return
}

// UpdatePacketfilterAdvancedTcpWindowScaling PUTs the packetfilter.advanced.tcp_window_scaling value to the UTM
func UpdatePacketfilterAdvancedTcpWindowScaling(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.advanced.tcp_window_scaling", val, options...)
}

// GetPacketfilterRules gets the packetfilter.rules value from the UTM
func GetPacketfilterRules(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/packetfilter.rules", &val, options...)
	return
}

// UpdatePacketfilterRules PUTs the packetfilter.rules value to the UTM
func UpdatePacketfilterRules(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.rules", val, options...)
}

// GetPacketfilterRulesAuto gets the packetfilter.rules_auto value from the UTM
func GetPacketfilterRulesAuto(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/packetfilter.rules_auto", &val, options...)
	return
}

// UpdatePacketfilterRulesAuto PUTs the packetfilter.rules_auto value to the UTM
func UpdatePacketfilterRulesAuto(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.rules_auto", val, options...)
}

// GetPacketfilterRulesBack gets the packetfilter.rules_back value from the UTM
func GetPacketfilterRulesBack(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/packetfilter.rules_back", &val, options...)
	return
}

// UpdatePacketfilterRulesBack PUTs the packetfilter.rules_back value to the UTM
func UpdatePacketfilterRulesBack(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.rules_back", val, options...)
}

// GetPacketfilterRulesFront gets the packetfilter.rules_front value from the UTM
func GetPacketfilterRulesFront(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/packetfilter.rules_front", &val, options...)
	return
}

// UpdatePacketfilterRulesFront PUTs the packetfilter.rules_front value to the UTM
func UpdatePacketfilterRulesFront(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.rules_front", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackGenericTimeout gets the packetfilter.timeouts.ip_conntrack_generic_timeout value from the UTM
func GetPacketfilterTimeoutsIpConntrackGenericTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_generic_timeout", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackGenericTimeout PUTs the packetfilter.timeouts.ip_conntrack_generic_timeout value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackGenericTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_generic_timeout", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackIcmpTimeout gets the packetfilter.timeouts.ip_conntrack_icmp_timeout value from the UTM
func GetPacketfilterTimeoutsIpConntrackIcmpTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_icmp_timeout", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackIcmpTimeout PUTs the packetfilter.timeouts.ip_conntrack_icmp_timeout value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackIcmpTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_icmp_timeout", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutClose gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_close value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutClose(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_close", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutClose PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_close value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutClose(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_close", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutCloseWait gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_close_wait value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutCloseWait(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_close_wait", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutCloseWait PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_close_wait value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutCloseWait(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_close_wait", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutEstablished gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_established value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutEstablished(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_established", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutEstablished PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_established value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutEstablished(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_established", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutFinWait gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_fin_wait value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutFinWait(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_fin_wait", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutFinWait PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_fin_wait value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutFinWait(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_fin_wait", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutLastAck gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_last_ack value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutLastAck(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_last_ack", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutLastAck PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_last_ack value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutLastAck(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_last_ack", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutMaxRetrans gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_max_retrans value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutMaxRetrans(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_max_retrans", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutMaxRetrans PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_max_retrans value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutMaxRetrans(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_max_retrans", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutSynRecv gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_recv value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutSynRecv(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_recv", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutSynRecv PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_recv value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutSynRecv(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_recv", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutSynSent gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_sent value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutSynSent(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_sent", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutSynSent PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_sent value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutSynSent(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_syn_sent", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackTcpTimeoutTimeWait gets the packetfilter.timeouts.ip_conntrack_tcp_timeout_time_wait value from the UTM
func GetPacketfilterTimeoutsIpConntrackTcpTimeoutTimeWait(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_time_wait", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutTimeWait PUTs the packetfilter.timeouts.ip_conntrack_tcp_timeout_time_wait value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackTcpTimeoutTimeWait(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_tcp_timeout_time_wait", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackUdpTimeout gets the packetfilter.timeouts.ip_conntrack_udp_timeout value from the UTM
func GetPacketfilterTimeoutsIpConntrackUdpTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_udp_timeout", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackUdpTimeout PUTs the packetfilter.timeouts.ip_conntrack_udp_timeout value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackUdpTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_udp_timeout", val, options...)
}

// GetPacketfilterTimeoutsIpConntrackUdpTimeoutStream gets the packetfilter.timeouts.ip_conntrack_udp_timeout_stream value from the UTM
func GetPacketfilterTimeoutsIpConntrackUdpTimeoutStream(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_udp_timeout_stream", &val, options...)
	return
}

// UpdatePacketfilterTimeoutsIpConntrackUdpTimeoutStream PUTs the packetfilter.timeouts.ip_conntrack_udp_timeout_stream value to the UTM
func UpdatePacketfilterTimeoutsIpConntrackUdpTimeoutStream(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/packetfilter.timeouts.ip_conntrack_udp_timeout_stream", val, options...)
}

// GetPasswdLoginuserClearpass gets the passwd.loginuser.clearpass value from the UTM
func GetPasswdLoginuserClearpass(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/passwd.loginuser.clearpass", &val, options...)
	return
}

// UpdatePasswdLoginuserClearpass PUTs the passwd.loginuser.clearpass value to the UTM
func UpdatePasswdLoginuserClearpass(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/passwd.loginuser.clearpass", val, options...)
}

// GetPasswdLoginuserCryptpass gets the passwd.loginuser.cryptpass value from the UTM
func GetPasswdLoginuserCryptpass(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/passwd.loginuser.cryptpass", &val, options...)
	return
}

// UpdatePasswdLoginuserCryptpass PUTs the passwd.loginuser.cryptpass value to the UTM
func UpdatePasswdLoginuserCryptpass(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/passwd.loginuser.cryptpass", val, options...)
}

// GetPasswdRootClearpass gets the passwd.root.clearpass value from the UTM
func GetPasswdRootClearpass(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/passwd.root.clearpass", &val, options...)
	return
}

// UpdatePasswdRootClearpass PUTs the passwd.root.clearpass value to the UTM
func UpdatePasswdRootClearpass(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/passwd.root.clearpass", val, options...)
}

// GetPasswdRootCryptpass gets the passwd.root.cryptpass value from the UTM
func GetPasswdRootCryptpass(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/passwd.root.cryptpass", &val, options...)
	return
}

// UpdatePasswdRootCryptpass PUTs the passwd.root.cryptpass value to the UTM
func UpdatePasswdRootCryptpass(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/passwd.root.cryptpass", val, options...)
}

// GetPdfPaper gets the pdf.paper value from the UTM
func GetPdfPaper(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/pdf.paper", &val, options...)
	return
}

// UpdatePdfPaper PUTs the pdf.paper value to the UTM
func UpdatePdfPaper(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pdf.paper", val, options...)
}

// GetPimSmAutoPfOut gets the pim_sm.auto_pf_out value from the UTM
func GetPimSmAutoPfOut(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/pim_sm.auto_pf_out", &val, options...)
	return
}

// UpdatePimSmAutoPfOut PUTs the pim_sm.auto_pf_out value to the UTM
func UpdatePimSmAutoPfOut(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pim_sm.auto_pf_out", val, options...)
}

// GetPimSmAutoPfrule gets the pim_sm.auto_pfrule value from the UTM
func GetPimSmAutoPfrule(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pim_sm.auto_pfrule", &val, options...)
	return
}

// UpdatePimSmAutoPfrule PUTs the pim_sm.auto_pfrule value to the UTM
func UpdatePimSmAutoPfrule(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pim_sm.auto_pfrule", val, options...)
}

// GetPimSmDebug gets the pim_sm.debug value from the UTM
func GetPimSmDebug(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pim_sm.debug", &val, options...)
	return
}

// UpdatePimSmDebug PUTs the pim_sm.debug value to the UTM
func UpdatePimSmDebug(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pim_sm.debug", val, options...)
}

// GetPimSmEnableSubnetMulticasting gets the pim_sm.enable_subnet_multicasting value from the UTM
func GetPimSmEnableSubnetMulticasting(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pim_sm.enable_subnet_multicasting", &val, options...)
	return
}

// UpdatePimSmEnableSubnetMulticasting PUTs the pim_sm.enable_subnet_multicasting value to the UTM
func UpdatePimSmEnableSubnetMulticasting(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pim_sm.enable_subnet_multicasting", val, options...)
}

// GetPimSmInterfaces gets the pim_sm.interfaces value from the UTM
func GetPimSmInterfaces(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/pim_sm.interfaces", &val, options...)
	return
}

// UpdatePimSmInterfaces PUTs the pim_sm.interfaces value to the UTM
func UpdatePimSmInterfaces(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pim_sm.interfaces", val, options...)
}

// GetPimSmRpRouters gets the pim_sm.rp_routers value from the UTM
func GetPimSmRpRouters(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/pim_sm.rp_routers", &val, options...)
	return
}

// UpdatePimSmRpRouters PUTs the pim_sm.rp_routers value to the UTM
func UpdatePimSmRpRouters(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pim_sm.rp_routers", val, options...)
}

// GetPimSmSptSwitchBytes gets the pim_sm.spt_switch_bytes value from the UTM
func GetPimSmSptSwitchBytes(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/pim_sm.spt_switch_bytes", &val, options...)
	return
}

// UpdatePimSmSptSwitchBytes PUTs the pim_sm.spt_switch_bytes value to the UTM
func UpdatePimSmSptSwitchBytes(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pim_sm.spt_switch_bytes", val, options...)
}

// GetPimSmSptSwitchStatus gets the pim_sm.spt_switch_status value from the UTM
func GetPimSmSptSwitchStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pim_sm.spt_switch_status", &val, options...)
	return
}

// UpdatePimSmSptSwitchStatus PUTs the pim_sm.spt_switch_status value to the UTM
func UpdatePimSmSptSwitchStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pim_sm.spt_switch_status", val, options...)
}

// GetPimSmStatus gets the pim_sm.status value from the UTM
func GetPimSmStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pim_sm.status", &val, options...)
	return
}

// UpdatePimSmStatus PUTs the pim_sm.status value to the UTM
func UpdatePimSmStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pim_sm.status", val, options...)
}

// GetPop3AllowedClients gets the pop3.allowed_clients value from the UTM
func GetPop3AllowedClients(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/pop3.allowed_clients", &val, options...)
	return
}

// UpdatePop3AllowedClients PUTs the pop3.allowed_clients value to the UTM
func UpdatePop3AllowedClients(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.allowed_clients", val, options...)
}

// GetPop3AllowedServers gets the pop3.allowed_servers value from the UTM
func GetPop3AllowedServers(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/pop3.allowed_servers", &val, options...)
	return
}

// UpdatePop3AllowedServers PUTs the pop3.allowed_servers value to the UTM
func UpdatePop3AllowedServers(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.allowed_servers", val, options...)
}

// GetPop3CffAsMarker gets the pop3.cff_as_marker value from the UTM
func GetPop3CffAsMarker(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/pop3.cff_as_marker", &val, options...)
	return
}

// UpdatePop3CffAsMarker PUTs the pop3.cff_as_marker value to the UTM
func UpdatePop3CffAsMarker(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.cff_as_marker", val, options...)
}

// GetPop3CffAv gets the pop3.cff_av value from the UTM
func GetPop3CffAv(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pop3.cff_av", &val, options...)
	return
}

// UpdatePop3CffAv PUTs the pop3.cff_av value to the UTM
func UpdatePop3CffAv(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.cff_av", val, options...)
}

// GetPop3CffAvAction gets the pop3.cff_av_action value from the UTM
func GetPop3CffAvAction(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/pop3.cff_av_action", &val, options...)
	return
}

// UpdatePop3CffAvAction PUTs the pop3.cff_av_action value to the UTM
func UpdatePop3CffAvAction(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.cff_av_action", val, options...)
}

// GetPop3CffAvEngines gets the pop3.cff_av_engines value from the UTM
func GetPop3CffAvEngines(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/pop3.cff_av_engines", &val, options...)
	return
}

// UpdatePop3CffAvEngines PUTs the pop3.cff_av_engines value to the UTM
func UpdatePop3CffAvEngines(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.cff_av_engines", val, options...)
}

// GetPop3CffFileExtensions gets the pop3.cff_file_extensions value from the UTM
func GetPop3CffFileExtensions(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/pop3.cff_file_extensions", &val, options...)
	return
}

// UpdatePop3CffFileExtensions PUTs the pop3.cff_file_extensions value to the UTM
func UpdatePop3CffFileExtensions(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.cff_file_extensions", val, options...)
}

// GetPop3DirectlyDeleteQuarantined gets the pop3.directly_delete_quarantined value from the UTM
func GetPop3DirectlyDeleteQuarantined(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pop3.directly_delete_quarantined", &val, options...)
	return
}

// UpdatePop3DirectlyDeleteQuarantined PUTs the pop3.directly_delete_quarantined value to the UTM
func UpdatePop3DirectlyDeleteQuarantined(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.directly_delete_quarantined", val, options...)
}

// GetPop3Exceptions gets the pop3.exceptions value from the UTM
func GetPop3Exceptions(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/pop3.exceptions", &val, options...)
	return
}

// UpdatePop3Exceptions PUTs the pop3.exceptions value to the UTM
func UpdatePop3Exceptions(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.exceptions", val, options...)
}

// GetPop3KnownServers gets the pop3.known_servers value from the UTM
func GetPop3KnownServers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/pop3.known_servers", &val, options...)
	return
}

// UpdatePop3KnownServers PUTs the pop3.known_servers value to the UTM
func UpdatePop3KnownServers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.known_servers", val, options...)
}

// GetPop3MaxMessageSize gets the pop3.max_message_size value from the UTM
func GetPop3MaxMessageSize(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/pop3.max_message_size", &val, options...)
	return
}

// UpdatePop3MaxMessageSize PUTs the pop3.max_message_size value to the UTM
func UpdatePop3MaxMessageSize(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.max_message_size", val, options...)
}

// GetPop3PrefetchingInterval gets the pop3.prefetching.interval value from the UTM
func GetPop3PrefetchingInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/pop3.prefetching.interval", &val, options...)
	return
}

// UpdatePop3PrefetchingInterval PUTs the pop3.prefetching.interval value to the UTM
func UpdatePop3PrefetchingInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.prefetching.interval", val, options...)
}

// GetPop3PrefetchingOptimizeStorage gets the pop3.prefetching.optimize_storage value from the UTM
func GetPop3PrefetchingOptimizeStorage(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pop3.prefetching.optimize_storage", &val, options...)
	return
}

// UpdatePop3PrefetchingOptimizeStorage PUTs the pop3.prefetching.optimize_storage value to the UTM
func UpdatePop3PrefetchingOptimizeStorage(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.prefetching.optimize_storage", val, options...)
}

// GetPop3PrefetchingStatus gets the pop3.prefetching.status value from the UTM
func GetPop3PrefetchingStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pop3.prefetching.status", &val, options...)
	return
}

// UpdatePop3PrefetchingStatus PUTs the pop3.prefetching.status value to the UTM
func UpdatePop3PrefetchingStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.prefetching.status", val, options...)
}

// GetPop3PrefetchingStorageMinHoldDays gets the pop3.prefetching.storage_min_hold_days value from the UTM
func GetPop3PrefetchingStorageMinHoldDays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/pop3.prefetching.storage_min_hold_days", &val, options...)
	return
}

// UpdatePop3PrefetchingStorageMinHoldDays PUTs the pop3.prefetching.storage_min_hold_days value to the UTM
func UpdatePop3PrefetchingStorageMinHoldDays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.prefetching.storage_min_hold_days", val, options...)
}

// GetPop3QuarantineUnscannable gets the pop3.quarantine_unscannable value from the UTM
func GetPop3QuarantineUnscannable(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pop3.quarantine_unscannable", &val, options...)
	return
}

// UpdatePop3QuarantineUnscannable PUTs the pop3.quarantine_unscannable value to the UTM
func UpdatePop3QuarantineUnscannable(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.quarantine_unscannable", val, options...)
}

// GetPop3SandboxMaxFilesizeMb gets the pop3.sandbox_max_filesize_mb value from the UTM
func GetPop3SandboxMaxFilesizeMb(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/pop3.sandbox_max_filesize_mb", &val, options...)
	return
}

// UpdatePop3SandboxMaxFilesizeMb PUTs the pop3.sandbox_max_filesize_mb value to the UTM
func UpdatePop3SandboxMaxFilesizeMb(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.sandbox_max_filesize_mb", val, options...)
}

// GetPop3SandboxScanStatus gets the pop3.sandbox_scan_status value from the UTM
func GetPop3SandboxScanStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pop3.sandbox_scan_status", &val, options...)
	return
}

// UpdatePop3SandboxScanStatus PUTs the pop3.sandbox_scan_status value to the UTM
func UpdatePop3SandboxScanStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.sandbox_scan_status", val, options...)
}

// GetPop3ScanTls gets the pop3.scan_tls value from the UTM
func GetPop3ScanTls(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pop3.scan_tls", &val, options...)
	return
}

// UpdatePop3ScanTls PUTs the pop3.scan_tls value to the UTM
func UpdatePop3ScanTls(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.scan_tls", val, options...)
}

// GetPop3SenderBlacklist gets the pop3.sender_blacklist value from the UTM
func GetPop3SenderBlacklist(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/pop3.sender_blacklist", &val, options...)
	return
}

// UpdatePop3SenderBlacklist PUTs the pop3.sender_blacklist value to the UTM
func UpdatePop3SenderBlacklist(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.sender_blacklist", val, options...)
}

// GetPop3Spam gets the pop3.spam value from the UTM
func GetPop3Spam(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/pop3.spam", &val, options...)
	return
}

// UpdatePop3Spam PUTs the pop3.spam value to the UTM
func UpdatePop3Spam(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.spam", val, options...)
}

// GetPop3SpamExpressions gets the pop3.spam_expressions value from the UTM
func GetPop3SpamExpressions(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/pop3.spam_expressions", &val, options...)
	return
}

// UpdatePop3SpamExpressions PUTs the pop3.spam_expressions value to the UTM
func UpdatePop3SpamExpressions(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.spam_expressions", val, options...)
}

// GetPop3Spamplus gets the pop3.spamplus value from the UTM
func GetPop3Spamplus(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/pop3.spamplus", &val, options...)
	return
}

// UpdatePop3Spamplus PUTs the pop3.spamplus value to the UTM
func UpdatePop3Spamplus(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.spamplus", val, options...)
}

// GetPop3Spamstatus gets the pop3.spamstatus value from the UTM
func GetPop3Spamstatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pop3.spamstatus", &val, options...)
	return
}

// UpdatePop3Spamstatus PUTs the pop3.spamstatus value to the UTM
func UpdatePop3Spamstatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.spamstatus", val, options...)
}

// GetPop3Status gets the pop3.status value from the UTM
func GetPop3Status(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pop3.status", &val, options...)
	return
}

// UpdatePop3Status PUTs the pop3.status value to the UTM
func UpdatePop3Status(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.status", val, options...)
}

// GetPop3TlsCert gets the pop3.tls_cert value from the UTM
func GetPop3TlsCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/pop3.tls_cert", &val, options...)
	return
}

// UpdatePop3TlsCert PUTs the pop3.tls_cert value to the UTM
func UpdatePop3TlsCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.tls_cert", val, options...)
}

// GetPop3TransparentSkip gets the pop3.transparent_skip value from the UTM
func GetPop3TransparentSkip(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/pop3.transparent_skip", &val, options...)
	return
}

// UpdatePop3TransparentSkip PUTs the pop3.transparent_skip value to the UTM
func UpdatePop3TransparentSkip(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.transparent_skip", val, options...)
}

// GetPop3TransparentSkipAutoPf gets the pop3.transparent_skip_auto_pf value from the UTM
func GetPop3TransparentSkipAutoPf(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/pop3.transparent_skip_auto_pf", &val, options...)
	return
}

// UpdatePop3TransparentSkipAutoPf PUTs the pop3.transparent_skip_auto_pf value to the UTM
func UpdatePop3TransparentSkipAutoPf(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.transparent_skip_auto_pf", val, options...)
}

// GetPop3UserCharset gets the pop3.user_charset value from the UTM
func GetPop3UserCharset(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/pop3.user_charset", &val, options...)
	return
}

// UpdatePop3UserCharset PUTs the pop3.user_charset value to the UTM
func UpdatePop3UserCharset(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/pop3.user_charset", val, options...)
}

// GetPortalAllowAnyUser gets the portal.allow_any_user value from the UTM
func GetPortalAllowAnyUser(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/portal.allow_any_user", &val, options...)
	return
}

// UpdatePortalAllowAnyUser PUTs the portal.allow_any_user value to the UTM
func UpdatePortalAllowAnyUser(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/portal.allow_any_user", val, options...)
}

// GetPortalAllowedNetworks gets the portal.allowed_networks value from the UTM
func GetPortalAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/portal.allowed_networks", &val, options...)
	return
}

// UpdatePortalAllowedNetworks PUTs the portal.allowed_networks value to the UTM
func UpdatePortalAllowedNetworks(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/portal.allowed_networks", val, options...)
}

// GetPortalAllowedUsers gets the portal.allowed_users value from the UTM
func GetPortalAllowedUsers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/portal.allowed_users", &val, options...)
	return
}

// UpdatePortalAllowedUsers PUTs the portal.allowed_users value to the UTM
func UpdatePortalAllowedUsers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/portal.allowed_users", val, options...)
}

// GetPortalCert gets the portal.cert value from the UTM
func GetPortalCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/portal.cert", &val, options...)
	return
}

// UpdatePortalCert PUTs the portal.cert value to the UTM
func UpdatePortalCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/portal.cert", val, options...)
}

// GetPortalHideItems gets the portal.hide_items value from the UTM
func GetPortalHideItems(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/portal.hide_items", &val, options...)
	return
}

// UpdatePortalHideItems PUTs the portal.hide_items value to the UTM
func UpdatePortalHideItems(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/portal.hide_items", val, options...)
}

// GetPortalHostname gets the portal.hostname value from the UTM
func GetPortalHostname(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/portal.hostname", &val, options...)
	return
}

// UpdatePortalHostname PUTs the portal.hostname value to the UTM
func UpdatePortalHostname(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/portal.hostname", val, options...)
}

// GetPortalInterfaceAddress gets the portal.interface_address value from the UTM
func GetPortalInterfaceAddress(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/portal.interface_address", &val, options...)
	return
}

// UpdatePortalInterfaceAddress PUTs the portal.interface_address value to the UTM
func UpdatePortalInterfaceAddress(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/portal.interface_address", val, options...)
}

// GetPortalLanguage gets the portal.language value from the UTM
func GetPortalLanguage(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/portal.language", &val, options...)
	return
}

// UpdatePortalLanguage PUTs the portal.language value to the UTM
func UpdatePortalLanguage(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/portal.language", val, options...)
}

// GetPortalPersistentCookies gets the portal.persistent_cookies value from the UTM
func GetPortalPersistentCookies(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/portal.persistent_cookies", &val, options...)
	return
}

// UpdatePortalPersistentCookies PUTs the portal.persistent_cookies value to the UTM
func UpdatePortalPersistentCookies(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/portal.persistent_cookies", val, options...)
}

// GetPortalPort gets the portal.port value from the UTM
func GetPortalPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/portal.port", &val, options...)
	return
}

// UpdatePortalPort PUTs the portal.port value to the UTM
func UpdatePortalPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/portal.port", val, options...)
}

// GetPortalStatus gets the portal.status value from the UTM
func GetPortalStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/portal.status", &val, options...)
	return
}

// UpdatePortalStatus PUTs the portal.status value to the UTM
func UpdatePortalStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/portal.status", val, options...)
}

// GetPortalWelcomeMsg gets the portal.welcome_msg value from the UTM
func GetPortalWelcomeMsg(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/portal.welcome_msg", &val, options...)
	return
}

// UpdatePortalWelcomeMsg PUTs the portal.welcome_msg value to the UTM
func UpdatePortalWelcomeMsg(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/portal.welcome_msg", val, options...)
}

// GetPsdAction gets the psd.action value from the UTM
func GetPsdAction(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/psd.action", &val, options...)
	return
}

// UpdatePsdAction PUTs the psd.action value to the UTM
func UpdatePsdAction(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/psd.action", val, options...)
}

// GetPsdDelayThreshold gets the psd.delay_threshold value from the UTM
func GetPsdDelayThreshold(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/psd.delay_threshold", &val, options...)
	return
}

// UpdatePsdDelayThreshold PUTs the psd.delay_threshold value to the UTM
func UpdatePsdDelayThreshold(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/psd.delay_threshold", val, options...)
}

// GetPsdHiPortsWeight gets the psd.hi_ports_weight value from the UTM
func GetPsdHiPortsWeight(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/psd.hi_ports_weight", &val, options...)
	return
}

// UpdatePsdHiPortsWeight PUTs the psd.hi_ports_weight value to the UTM
func UpdatePsdHiPortsWeight(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/psd.hi_ports_weight", val, options...)
}

// GetPsdLoPortsWeight gets the psd.lo_ports_weight value from the UTM
func GetPsdLoPortsWeight(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/psd.lo_ports_weight", &val, options...)
	return
}

// UpdatePsdLoPortsWeight PUTs the psd.lo_ports_weight value to the UTM
func UpdatePsdLoPortsWeight(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/psd.lo_ports_weight", val, options...)
}

// GetPsdLogLimiterBurst gets the psd.log_limiter.burst value from the UTM
func GetPsdLogLimiterBurst(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/psd.log_limiter.burst", &val, options...)
	return
}

// UpdatePsdLogLimiterBurst PUTs the psd.log_limiter.burst value to the UTM
func UpdatePsdLogLimiterBurst(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/psd.log_limiter.burst", val, options...)
}

// GetPsdLogLimiterRate gets the psd.log_limiter.rate value from the UTM
func GetPsdLogLimiterRate(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/psd.log_limiter.rate", &val, options...)
	return
}

// UpdatePsdLogLimiterRate PUTs the psd.log_limiter.rate value to the UTM
func UpdatePsdLogLimiterRate(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/psd.log_limiter.rate", val, options...)
}

// GetPsdLogLimiterStatus gets the psd.log_limiter.status value from the UTM
func GetPsdLogLimiterStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/psd.log_limiter.status", &val, options...)
	return
}

// UpdatePsdLogLimiterStatus PUTs the psd.log_limiter.status value to the UTM
func UpdatePsdLogLimiterStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/psd.log_limiter.status", val, options...)
}

// GetPsdStatus gets the psd.status value from the UTM
func GetPsdStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/psd.status", &val, options...)
	return
}

// UpdatePsdStatus PUTs the psd.status value to the UTM
func UpdatePsdStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/psd.status", val, options...)
}

// GetPsdWeightThreshold gets the psd.weight_threshold value from the UTM
func GetPsdWeightThreshold(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/psd.weight_threshold", &val, options...)
	return
}

// UpdatePsdWeightThreshold PUTs the psd.weight_threshold value to the UTM
func UpdatePsdWeightThreshold(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/psd.weight_threshold", val, options...)
}

// GetQosAdvancedEcn gets the qos.advanced.ecn value from the UTM
func GetQosAdvancedEcn(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/qos.advanced.ecn", &val, options...)
	return
}

// UpdateQosAdvancedEcn PUTs the qos.advanced.ecn value to the UTM
func UpdateQosAdvancedEcn(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/qos.advanced.ecn", val, options...)
}

// GetQosAdvancedKeepClassAfterEncap gets the qos.advanced.keep_class_after_encap value from the UTM
func GetQosAdvancedKeepClassAfterEncap(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/qos.advanced.keep_class_after_encap", &val, options...)
	return
}

// UpdateQosAdvancedKeepClassAfterEncap PUTs the qos.advanced.keep_class_after_encap value to the UTM
func UpdateQosAdvancedKeepClassAfterEncap(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/qos.advanced.keep_class_after_encap", val, options...)
}

// GetQosInterfaces gets the qos.interfaces value from the UTM
func GetQosInterfaces(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/qos.interfaces", &val, options...)
	return
}

// UpdateQosInterfaces PUTs the qos.interfaces value to the UTM
func UpdateQosInterfaces(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/qos.interfaces", val, options...)
}

// GetQuarantineKeepDbLogDays gets the quarantine.keep_db_log_days value from the UTM
func GetQuarantineKeepDbLogDays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/quarantine.keep_db_log_days", &val, options...)
	return
}

// UpdateQuarantineKeepDbLogDays PUTs the quarantine.keep_db_log_days value to the UTM
func UpdateQuarantineKeepDbLogDays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/quarantine.keep_db_log_days", val, options...)
}

// GetQuarantineKeepQuarantineDays gets the quarantine.keep_quarantine_days value from the UTM
func GetQuarantineKeepQuarantineDays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/quarantine.keep_quarantine_days", &val, options...)
	return
}

// UpdateQuarantineKeepQuarantineDays PUTs the quarantine.keep_quarantine_days value to the UTM
func UpdateQuarantineKeepQuarantineDays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/quarantine.keep_quarantine_days", val, options...)
}

// GetRedActivateProvFw gets the red.activate_prov_fw value from the UTM
func GetRedActivateProvFw(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/red.activate_prov_fw", &val, options...)
	return
}

// UpdateRedActivateProvFw PUTs the red.activate_prov_fw value to the UTM
func UpdateRedActivateProvFw(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.activate_prov_fw", val, options...)
}

// GetRedAuthorization gets the red.authorization value from the UTM
func GetRedAuthorization(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/red.authorization", &val, options...)
	return
}

// UpdateRedAuthorization PUTs the red.authorization value to the UTM
func UpdateRedAuthorization(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.authorization", val, options...)
}

// GetRedCaSettingsCity gets the red.ca_settings.city value from the UTM
func GetRedCaSettingsCity(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/red.ca_settings.city", &val, options...)
	return
}

// UpdateRedCaSettingsCity PUTs the red.ca_settings.city value to the UTM
func UpdateRedCaSettingsCity(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.ca_settings.city", val, options...)
}

// GetRedCaSettingsCountry gets the red.ca_settings.country value from the UTM
func GetRedCaSettingsCountry(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/red.ca_settings.country", &val, options...)
	return
}

// UpdateRedCaSettingsCountry PUTs the red.ca_settings.country value to the UTM
func UpdateRedCaSettingsCountry(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.ca_settings.country", val, options...)
}

// GetRedCaSettingsEmail gets the red.ca_settings.email value from the UTM
func GetRedCaSettingsEmail(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/red.ca_settings.email", &val, options...)
	return
}

// UpdateRedCaSettingsEmail PUTs the red.ca_settings.email value to the UTM
func UpdateRedCaSettingsEmail(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.ca_settings.email", val, options...)
}

// GetRedCaSettingsOrganization gets the red.ca_settings.organization value from the UTM
func GetRedCaSettingsOrganization(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/red.ca_settings.organization", &val, options...)
	return
}

// UpdateRedCaSettingsOrganization PUTs the red.ca_settings.organization value to the UTM
func UpdateRedCaSettingsOrganization(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.ca_settings.organization", val, options...)
}

// GetRedClients gets the red.clients value from the UTM
func GetRedClients(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/red.clients", &val, options...)
	return
}

// UpdateRedClients PUTs the red.clients value to the UTM
func UpdateRedClients(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.clients", val, options...)
}

// GetRedDeauthTimeout gets the red.deauth_timeout value from the UTM
func GetRedDeauthTimeout(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/red.deauth_timeout", &val, options...)
	return
}

// UpdateRedDeauthTimeout PUTs the red.deauth_timeout value to the UTM
func UpdateRedDeauthTimeout(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.deauth_timeout", val, options...)
}

// GetRedOverlayFwEnabled gets the red.overlay_fw_enabled value from the UTM
func GetRedOverlayFwEnabled(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/red.overlay_fw_enabled", &val, options...)
	return
}

// UpdateRedOverlayFwEnabled PUTs the red.overlay_fw_enabled value to the UTM
func UpdateRedOverlayFwEnabled(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.overlay_fw_enabled", val, options...)
}

// GetRedRegistryCert gets the red.registry_cert value from the UTM
func GetRedRegistryCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/red.registry_cert", &val, options...)
	return
}

// UpdateRedRegistryCert PUTs the red.registry_cert value to the UTM
func UpdateRedRegistryCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.registry_cert", val, options...)
}

// GetRedRegistryId gets the red.registry_id value from the UTM
func GetRedRegistryId(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/red.registry_id", &val, options...)
	return
}

// UpdateRedRegistryId PUTs the red.registry_id value to the UTM
func UpdateRedRegistryId(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.registry_id", val, options...)
}

// GetRedRegistryKey gets the red.registry_key value from the UTM
func GetRedRegistryKey(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/red.registry_key", &val, options...)
	return
}

// UpdateRedRegistryKey PUTs the red.registry_key value to the UTM
func UpdateRedRegistryKey(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.registry_key", val, options...)
}

// GetRedReshowEula gets the red.reshow_eula value from the UTM
func GetRedReshowEula(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/red.reshow_eula", &val, options...)
	return
}

// UpdateRedReshowEula PUTs the red.reshow_eula value to the UTM
func UpdateRedReshowEula(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.reshow_eula", val, options...)
}

// GetRedServerCert gets the red.server_cert value from the UTM
func GetRedServerCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/red.server_cert", &val, options...)
	return
}

// UpdateRedServerCert PUTs the red.server_cert value to the UTM
func UpdateRedServerCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.server_cert", val, options...)
}

// GetRedServers gets the red.servers value from the UTM
func GetRedServers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/red.servers", &val, options...)
	return
}

// UpdateRedServers PUTs the red.servers value to the UTM
func UpdateRedServers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.servers", val, options...)
}

// GetRedStatus gets the red.status value from the UTM
func GetRedStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/red.status", &val, options...)
	return
}

// UpdateRedStatus PUTs the red.status value to the UTM
func UpdateRedStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.status", val, options...)
}

// GetRedTls12Only gets the red.tls_1_2_only value from the UTM
func GetRedTls12Only(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/red.tls_1_2_only", &val, options...)
	return
}

// UpdateRedTls12Only PUTs the red.tls_1_2_only value to the UTM
func UpdateRedTls12Only(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.tls_1_2_only", val, options...)
}

// GetRedUseUnifiedFirmware gets the red.use_unified_firmware value from the UTM
func GetRedUseUnifiedFirmware(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/red.use_unified_firmware", &val, options...)
	return
}

// UpdateRedUseUnifiedFirmware PUTs the red.use_unified_firmware value to the UTM
func UpdateRedUseUnifiedFirmware(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/red.use_unified_firmware", val, options...)
}

// GetRemoteAccessAdvancedMsdns1 gets the remote_access.advanced.msdns1 value from the UTM
func GetRemoteAccessAdvancedMsdns1(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.advanced.msdns1", &val, options...)
	return
}

// UpdateRemoteAccessAdvancedMsdns1 PUTs the remote_access.advanced.msdns1 value to the UTM
func UpdateRemoteAccessAdvancedMsdns1(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.advanced.msdns1", val, options...)
}

// GetRemoteAccessAdvancedMsdns2 gets the remote_access.advanced.msdns2 value from the UTM
func GetRemoteAccessAdvancedMsdns2(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.advanced.msdns2", &val, options...)
	return
}

// UpdateRemoteAccessAdvancedMsdns2 PUTs the remote_access.advanced.msdns2 value to the UTM
func UpdateRemoteAccessAdvancedMsdns2(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.advanced.msdns2", val, options...)
}

// GetRemoteAccessAdvancedMsdomain gets the remote_access.advanced.msdomain value from the UTM
func GetRemoteAccessAdvancedMsdomain(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.advanced.msdomain", &val, options...)
	return
}

// UpdateRemoteAccessAdvancedMsdomain PUTs the remote_access.advanced.msdomain value to the UTM
func UpdateRemoteAccessAdvancedMsdomain(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.advanced.msdomain", val, options...)
}

// GetRemoteAccessAdvancedMswins1 gets the remote_access.advanced.mswins1 value from the UTM
func GetRemoteAccessAdvancedMswins1(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.advanced.mswins1", &val, options...)
	return
}

// UpdateRemoteAccessAdvancedMswins1 PUTs the remote_access.advanced.mswins1 value to the UTM
func UpdateRemoteAccessAdvancedMswins1(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.advanced.mswins1", val, options...)
}

// GetRemoteAccessAdvancedMswins2 gets the remote_access.advanced.mswins2 value from the UTM
func GetRemoteAccessAdvancedMswins2(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.advanced.mswins2", &val, options...)
	return
}

// UpdateRemoteAccessAdvancedMswins2 PUTs the remote_access.advanced.mswins2 value to the UTM
func UpdateRemoteAccessAdvancedMswins2(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.advanced.mswins2", val, options...)
}

// GetRemoteAccessCisco gets the remote_access.cisco value from the UTM
func GetRemoteAccessCisco(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.cisco", &val, options...)
	return
}

// UpdateRemoteAccessCisco PUTs the remote_access.cisco value to the UTM
func UpdateRemoteAccessCisco(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.cisco", val, options...)
}

// GetRemoteAccessClientlessVpnDebug gets the remote_access.clientless_vpn.debug value from the UTM
func GetRemoteAccessClientlessVpnDebug(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/remote_access.clientless_vpn.debug", &val, options...)
	return
}

// UpdateRemoteAccessClientlessVpnDebug PUTs the remote_access.clientless_vpn.debug value to the UTM
func UpdateRemoteAccessClientlessVpnDebug(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.clientless_vpn.debug", val, options...)
}

// GetRemoteAccessClientlessVpnStatus gets the remote_access.clientless_vpn.status value from the UTM
func GetRemoteAccessClientlessVpnStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/remote_access.clientless_vpn.status", &val, options...)
	return
}

// UpdateRemoteAccessClientlessVpnStatus PUTs the remote_access.clientless_vpn.status value to the UTM
func UpdateRemoteAccessClientlessVpnStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.clientless_vpn.status", val, options...)
}

// GetRemoteAccessL2Tp gets the remote_access.l2tp value from the UTM
func GetRemoteAccessL2Tp(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.l2tp", &val, options...)
	return
}

// UpdateRemoteAccessL2Tp PUTs the remote_access.l2tp value to the UTM
func UpdateRemoteAccessL2Tp(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.l2tp", val, options...)
}

// GetRemoteAccessPptpAaa gets the remote_access.pptp.aaa value from the UTM
func GetRemoteAccessPptpAaa(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.aaa", &val, options...)
	return
}

// UpdateRemoteAccessPptpAaa PUTs the remote_access.pptp.aaa value to the UTM
func UpdateRemoteAccessPptpAaa(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.aaa", val, options...)
}

// GetRemoteAccessPptpAuthentication gets the remote_access.pptp.authentication value from the UTM
func GetRemoteAccessPptpAuthentication(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.authentication", &val, options...)
	return
}

// UpdateRemoteAccessPptpAuthentication PUTs the remote_access.pptp.authentication value to the UTM
func UpdateRemoteAccessPptpAuthentication(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.authentication", val, options...)
}

// GetRemoteAccessPptpDebug gets the remote_access.pptp.debug value from the UTM
func GetRemoteAccessPptpDebug(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.debug", &val, options...)
	return
}

// UpdateRemoteAccessPptpDebug PUTs the remote_access.pptp.debug value to the UTM
func UpdateRemoteAccessPptpDebug(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.debug", val, options...)
}

// GetRemoteAccessPptpEncryption gets the remote_access.pptp.encryption value from the UTM
func GetRemoteAccessPptpEncryption(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.encryption", &val, options...)
	return
}

// UpdateRemoteAccessPptpEncryption PUTs the remote_access.pptp.encryption value to the UTM
func UpdateRemoteAccessPptpEncryption(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.encryption", val, options...)
}

// GetRemoteAccessPptpIpAssignmentDhcp gets the remote_access.pptp.ip_assignment_dhcp value from the UTM
func GetRemoteAccessPptpIpAssignmentDhcp(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.ip_assignment_dhcp", &val, options...)
	return
}

// UpdateRemoteAccessPptpIpAssignmentDhcp PUTs the remote_access.pptp.ip_assignment_dhcp value to the UTM
func UpdateRemoteAccessPptpIpAssignmentDhcp(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.ip_assignment_dhcp", val, options...)
}

// GetRemoteAccessPptpIpAssignmentDhcpInterface gets the remote_access.pptp.ip_assignment_dhcp_interface value from the UTM
func GetRemoteAccessPptpIpAssignmentDhcpInterface(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.ip_assignment_dhcp_interface", &val, options...)
	return
}

// UpdateRemoteAccessPptpIpAssignmentDhcpInterface PUTs the remote_access.pptp.ip_assignment_dhcp_interface value to the UTM
func UpdateRemoteAccessPptpIpAssignmentDhcpInterface(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.ip_assignment_dhcp_interface", val, options...)
}

// GetRemoteAccessPptpIpAssignmentMode gets the remote_access.pptp.ip_assignment_mode value from the UTM
func GetRemoteAccessPptpIpAssignmentMode(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.ip_assignment_mode", &val, options...)
	return
}

// UpdateRemoteAccessPptpIpAssignmentMode PUTs the remote_access.pptp.ip_assignment_mode value to the UTM
func UpdateRemoteAccessPptpIpAssignmentMode(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.ip_assignment_mode", val, options...)
}

// GetRemoteAccessPptpIpAssignmentPool gets the remote_access.pptp.ip_assignment_pool value from the UTM
func GetRemoteAccessPptpIpAssignmentPool(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.ip_assignment_pool", &val, options...)
	return
}

// UpdateRemoteAccessPptpIpAssignmentPool PUTs the remote_access.pptp.ip_assignment_pool value to the UTM
func UpdateRemoteAccessPptpIpAssignmentPool(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.ip_assignment_pool", val, options...)
}

// GetRemoteAccessPptpIphoneConnectionName gets the remote_access.pptp.iphone_connection_name value from the UTM
func GetRemoteAccessPptpIphoneConnectionName(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.iphone_connection_name", &val, options...)
	return
}

// UpdateRemoteAccessPptpIphoneConnectionName PUTs the remote_access.pptp.iphone_connection_name value to the UTM
func UpdateRemoteAccessPptpIphoneConnectionName(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.iphone_connection_name", val, options...)
}

// GetRemoteAccessPptpIphoneHostname gets the remote_access.pptp.iphone_hostname value from the UTM
func GetRemoteAccessPptpIphoneHostname(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.iphone_hostname", &val, options...)
	return
}

// UpdateRemoteAccessPptpIphoneHostname PUTs the remote_access.pptp.iphone_hostname value to the UTM
func UpdateRemoteAccessPptpIphoneHostname(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.iphone_hostname", val, options...)
}

// GetRemoteAccessPptpIphoneStatus gets the remote_access.pptp.iphone_status value from the UTM
func GetRemoteAccessPptpIphoneStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.iphone_status", &val, options...)
	return
}

// UpdateRemoteAccessPptpIphoneStatus PUTs the remote_access.pptp.iphone_status value to the UTM
func UpdateRemoteAccessPptpIphoneStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.iphone_status", val, options...)
}

// GetRemoteAccessPptpMtu gets the remote_access.pptp.mtu value from the UTM
func GetRemoteAccessPptpMtu(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.mtu", &val, options...)
	return
}

// UpdateRemoteAccessPptpMtu PUTs the remote_access.pptp.mtu value to the UTM
func UpdateRemoteAccessPptpMtu(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.mtu", val, options...)
}

// GetRemoteAccessPptpStatus gets the remote_access.pptp.status value from the UTM
func GetRemoteAccessPptpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/remote_access.pptp.status", &val, options...)
	return
}

// UpdateRemoteAccessPptpStatus PUTs the remote_access.pptp.status value to the UTM
func UpdateRemoteAccessPptpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_access.pptp.status", val, options...)
}

// GetRemoteSyslogBuffer gets the remote_syslog.buffer value from the UTM
func GetRemoteSyslogBuffer(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/remote_syslog.buffer", &val, options...)
	return
}

// UpdateRemoteSyslogBuffer PUTs the remote_syslog.buffer value to the UTM
func UpdateRemoteSyslogBuffer(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_syslog.buffer", val, options...)
}

// GetRemoteSyslogLogs gets the remote_syslog.logs value from the UTM
func GetRemoteSyslogLogs(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/remote_syslog.logs", &val, options...)
	return
}

// UpdateRemoteSyslogLogs PUTs the remote_syslog.logs value to the UTM
func UpdateRemoteSyslogLogs(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_syslog.logs", val, options...)
}

// GetRemoteSyslogStatus gets the remote_syslog.status value from the UTM
func GetRemoteSyslogStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/remote_syslog.status", &val, options...)
	return
}

// UpdateRemoteSyslogStatus PUTs the remote_syslog.status value to the UTM
func UpdateRemoteSyslogStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_syslog.status", val, options...)
}

// GetRemoteSyslogTarget gets the remote_syslog.target value from the UTM
func GetRemoteSyslogTarget(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/remote_syslog.target", &val, options...)
	return
}

// UpdateRemoteSyslogTarget PUTs the remote_syslog.target value to the UTM
func UpdateRemoteSyslogTarget(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_syslog.target", val, options...)
}

// GetRemoteSyslogTimeReopen gets the remote_syslog.time_reopen value from the UTM
func GetRemoteSyslogTimeReopen(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/remote_syslog.time_reopen", &val, options...)
	return
}

// UpdateRemoteSyslogTimeReopen PUTs the remote_syslog.time_reopen value to the UTM
func UpdateRemoteSyslogTimeReopen(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/remote_syslog.time_reopen", val, options...)
}

// GetReportingAccountingKeepdays gets the reporting.accounting_keepdays value from the UTM
func GetReportingAccountingKeepdays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reporting.accounting_keepdays", &val, options...)
	return
}

// UpdateReportingAccountingKeepdays PUTs the reporting.accounting_keepdays value to the UTM
func UpdateReportingAccountingKeepdays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.accounting_keepdays", val, options...)
}

// GetReportingAccountingStatus gets the reporting.accounting_status value from the UTM
func GetReportingAccountingStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.accounting_status", &val, options...)
	return
}

// UpdateReportingAccountingStatus PUTs the reporting.accounting_status value to the UTM
func UpdateReportingAccountingStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.accounting_status", val, options...)
}

// GetReportingAnonymizing gets the reporting.anonymizing value from the UTM
func GetReportingAnonymizing(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.anonymizing", &val, options...)
	return
}

// UpdateReportingAnonymizing PUTs the reporting.anonymizing value to the UTM
func UpdateReportingAnonymizing(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.anonymizing", val, options...)
}

// GetReportingAppctrlKeepdays gets the reporting.appctrl_keepdays value from the UTM
func GetReportingAppctrlKeepdays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reporting.appctrl_keepdays", &val, options...)
	return
}

// UpdateReportingAppctrlKeepdays PUTs the reporting.appctrl_keepdays value to the UTM
func UpdateReportingAppctrlKeepdays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.appctrl_keepdays", val, options...)
}

// GetReportingAppctrlStatus gets the reporting.appctrl_status value from the UTM
func GetReportingAppctrlStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.appctrl_status", &val, options...)
	return
}

// UpdateReportingAppctrlStatus PUTs the reporting.appctrl_status value to the UTM
func UpdateReportingAppctrlStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.appctrl_status", val, options...)
}

// GetReportingAtpKeepdays gets the reporting.atp_keepdays value from the UTM
func GetReportingAtpKeepdays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reporting.atp_keepdays", &val, options...)
	return
}

// UpdateReportingAtpKeepdays PUTs the reporting.atp_keepdays value to the UTM
func UpdateReportingAtpKeepdays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.atp_keepdays", val, options...)
}

// GetReportingAtpReset gets the reporting.atp_reset value from the UTM
func GetReportingAtpReset(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/reporting.atp_reset", &val, options...)
	return
}

// UpdateReportingAtpReset PUTs the reporting.atp_reset value to the UTM
func UpdateReportingAtpReset(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.atp_reset", val, options...)
}

// GetReportingAtpStatus gets the reporting.atp_status value from the UTM
func GetReportingAtpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.atp_status", &val, options...)
	return
}

// UpdateReportingAtpStatus PUTs the reporting.atp_status value to the UTM
func UpdateReportingAtpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.atp_status", val, options...)
}

// GetReportingAuthenticationKeepdays gets the reporting.authentication_keepdays value from the UTM
func GetReportingAuthenticationKeepdays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reporting.authentication_keepdays", &val, options...)
	return
}

// UpdateReportingAuthenticationKeepdays PUTs the reporting.authentication_keepdays value to the UTM
func UpdateReportingAuthenticationKeepdays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.authentication_keepdays", val, options...)
}

// GetReportingAuthenticationStatus gets the reporting.authentication_status value from the UTM
func GetReportingAuthenticationStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.authentication_status", &val, options...)
	return
}

// UpdateReportingAuthenticationStatus PUTs the reporting.authentication_status value to the UTM
func UpdateReportingAuthenticationStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.authentication_status", val, options...)
}

// GetReportingCsvSeparator gets the reporting.csv_separator value from the UTM
func GetReportingCsvSeparator(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/reporting.csv_separator", &val, options...)
	return
}

// UpdateReportingCsvSeparator PUTs the reporting.csv_separator value to the UTM
func UpdateReportingCsvSeparator(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.csv_separator", val, options...)
}

// GetReportingEmailsecurityImport gets the reporting.emailsecurity_import value from the UTM
func GetReportingEmailsecurityImport(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/reporting.emailsecurity_import", &val, options...)
	return
}

// UpdateReportingEmailsecurityImport PUTs the reporting.emailsecurity_import value to the UTM
func UpdateReportingEmailsecurityImport(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.emailsecurity_import", val, options...)
}

// GetReportingEmailsecurityKeepdays gets the reporting.emailsecurity_keepdays value from the UTM
func GetReportingEmailsecurityKeepdays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reporting.emailsecurity_keepdays", &val, options...)
	return
}

// UpdateReportingEmailsecurityKeepdays PUTs the reporting.emailsecurity_keepdays value to the UTM
func UpdateReportingEmailsecurityKeepdays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.emailsecurity_keepdays", val, options...)
}

// GetReportingEmailsecurityStatus gets the reporting.emailsecurity_status value from the UTM
func GetReportingEmailsecurityStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.emailsecurity_status", &val, options...)
	return
}

// UpdateReportingEmailsecurityStatus PUTs the reporting.emailsecurity_status value to the UTM
func UpdateReportingEmailsecurityStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.emailsecurity_status", val, options...)
}

// GetReportingEnableVpnAccounting gets the reporting.enable_vpn_accounting value from the UTM
func GetReportingEnableVpnAccounting(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.enable_vpn_accounting", &val, options...)
	return
}

// UpdateReportingEnableVpnAccounting PUTs the reporting.enable_vpn_accounting value to the UTM
func UpdateReportingEnableVpnAccounting(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.enable_vpn_accounting", val, options...)
}

// GetReportingHideAccountingips gets the reporting.hide_accountingips value from the UTM
func GetReportingHideAccountingips(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/reporting.hide_accountingips", &val, options...)
	return
}

// UpdateReportingHideAccountingips PUTs the reporting.hide_accountingips value to the UTM
func UpdateReportingHideAccountingips(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.hide_accountingips", val, options...)
}

// GetReportingHideMailaddresses gets the reporting.hide_mailaddresses value from the UTM
func GetReportingHideMailaddresses(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/reporting.hide_mailaddresses", &val, options...)
	return
}

// UpdateReportingHideMailaddresses PUTs the reporting.hide_mailaddresses value to the UTM
func UpdateReportingHideMailaddresses(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.hide_mailaddresses", val, options...)
}

// GetReportingHideMaildomains gets the reporting.hide_maildomains value from the UTM
func GetReportingHideMaildomains(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/reporting.hide_maildomains", &val, options...)
	return
}

// UpdateReportingHideMaildomains PUTs the reporting.hide_maildomains value to the UTM
func UpdateReportingHideMaildomains(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.hide_maildomains", val, options...)
}

// GetReportingHideNetsecips gets the reporting.hide_netsecips value from the UTM
func GetReportingHideNetsecips(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/reporting.hide_netsecips", &val, options...)
	return
}

// UpdateReportingHideNetsecips PUTs the reporting.hide_netsecips value to the UTM
func UpdateReportingHideNetsecips(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.hide_netsecips", val, options...)
}

// GetReportingHideWebdomains gets the reporting.hide_webdomains value from the UTM
func GetReportingHideWebdomains(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/reporting.hide_webdomains", &val, options...)
	return
}

// UpdateReportingHideWebdomains PUTs the reporting.hide_webdomains value to the UTM
func UpdateReportingHideWebdomains(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.hide_webdomains", val, options...)
}

// GetReportingIpsImport gets the reporting.ips_import value from the UTM
func GetReportingIpsImport(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/reporting.ips_import", &val, options...)
	return
}

// UpdateReportingIpsImport PUTs the reporting.ips_import value to the UTM
func UpdateReportingIpsImport(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.ips_import", val, options...)
}

// GetReportingIpsKeepdays gets the reporting.ips_keepdays value from the UTM
func GetReportingIpsKeepdays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reporting.ips_keepdays", &val, options...)
	return
}

// UpdateReportingIpsKeepdays PUTs the reporting.ips_keepdays value to the UTM
func UpdateReportingIpsKeepdays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.ips_keepdays", val, options...)
}

// GetReportingIpsStatus gets the reporting.ips_status value from the UTM
func GetReportingIpsStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.ips_status", &val, options...)
	return
}

// UpdateReportingIpsStatus PUTs the reporting.ips_status value to the UTM
func UpdateReportingIpsStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.ips_status", val, options...)
}

// GetReportingPacketfilterImport gets the reporting.packetfilter_import value from the UTM
func GetReportingPacketfilterImport(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/reporting.packetfilter_import", &val, options...)
	return
}

// UpdateReportingPacketfilterImport PUTs the reporting.packetfilter_import value to the UTM
func UpdateReportingPacketfilterImport(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.packetfilter_import", val, options...)
}

// GetReportingPacketfilterKeepdays gets the reporting.packetfilter_keepdays value from the UTM
func GetReportingPacketfilterKeepdays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reporting.packetfilter_keepdays", &val, options...)
	return
}

// UpdateReportingPacketfilterKeepdays PUTs the reporting.packetfilter_keepdays value to the UTM
func UpdateReportingPacketfilterKeepdays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.packetfilter_keepdays", val, options...)
}

// GetReportingPacketfilterStatus gets the reporting.packetfilter_status value from the UTM
func GetReportingPacketfilterStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.packetfilter_status", &val, options...)
	return
}

// UpdateReportingPacketfilterStatus PUTs the reporting.packetfilter_status value to the UTM
func UpdateReportingPacketfilterStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.packetfilter_status", val, options...)
}

// GetReportingPassword1 gets the reporting.password1 value from the UTM
func GetReportingPassword1(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/reporting.password1", &val, options...)
	return
}

// UpdateReportingPassword1 PUTs the reporting.password1 value to the UTM
func UpdateReportingPassword1(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.password1", val, options...)
}

// GetReportingPassword2 gets the reporting.password2 value from the UTM
func GetReportingPassword2(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/reporting.password2", &val, options...)
	return
}

// UpdateReportingPassword2 PUTs the reporting.password2 value to the UTM
func UpdateReportingPassword2(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.password2", val, options...)
}

// GetReportingSandboxKeepdays gets the reporting.sandbox_keepdays value from the UTM
func GetReportingSandboxKeepdays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reporting.sandbox_keepdays", &val, options...)
	return
}

// UpdateReportingSandboxKeepdays PUTs the reporting.sandbox_keepdays value to the UTM
func UpdateReportingSandboxKeepdays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.sandbox_keepdays", val, options...)
}

// GetReportingUserlogFromLogs gets the reporting.userlog_from_logs value from the UTM
func GetReportingUserlogFromLogs(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.userlog_from_logs", &val, options...)
	return
}

// UpdateReportingUserlogFromLogs PUTs the reporting.userlog_from_logs value to the UTM
func UpdateReportingUserlogFromLogs(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.userlog_from_logs", val, options...)
}

// GetReportingVpnKeepdays gets the reporting.vpn_keepdays value from the UTM
func GetReportingVpnKeepdays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reporting.vpn_keepdays", &val, options...)
	return
}

// UpdateReportingVpnKeepdays PUTs the reporting.vpn_keepdays value to the UTM
func UpdateReportingVpnKeepdays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.vpn_keepdays", val, options...)
}

// GetReportingVpnStatus gets the reporting.vpn_status value from the UTM
func GetReportingVpnStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.vpn_status", &val, options...)
	return
}

// UpdateReportingVpnStatus PUTs the reporting.vpn_status value to the UTM
func UpdateReportingVpnStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.vpn_status", val, options...)
}

// GetReportingWafKeepdays gets the reporting.waf_keepdays value from the UTM
func GetReportingWafKeepdays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reporting.waf_keepdays", &val, options...)
	return
}

// UpdateReportingWafKeepdays PUTs the reporting.waf_keepdays value to the UTM
func UpdateReportingWafKeepdays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.waf_keepdays", val, options...)
}

// GetReportingWafStatus gets the reporting.waf_status value from the UTM
func GetReportingWafStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.waf_status", &val, options...)
	return
}

// UpdateReportingWafStatus PUTs the reporting.waf_status value to the UTM
func UpdateReportingWafStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.waf_status", val, options...)
}

// GetReportingWebsecurityDetail gets the reporting.websecurity_detail value from the UTM
func GetReportingWebsecurityDetail(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/reporting.websecurity_detail", &val, options...)
	return
}

// UpdateReportingWebsecurityDetail PUTs the reporting.websecurity_detail value to the UTM
func UpdateReportingWebsecurityDetail(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.websecurity_detail", val, options...)
}

// GetReportingWebsecurityImport gets the reporting.websecurity_import value from the UTM
func GetReportingWebsecurityImport(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/reporting.websecurity_import", &val, options...)
	return
}

// UpdateReportingWebsecurityImport PUTs the reporting.websecurity_import value to the UTM
func UpdateReportingWebsecurityImport(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.websecurity_import", val, options...)
}

// GetReportingWebsecurityKeepdays gets the reporting.websecurity_keepdays value from the UTM
func GetReportingWebsecurityKeepdays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reporting.websecurity_keepdays", &val, options...)
	return
}

// UpdateReportingWebsecurityKeepdays PUTs the reporting.websecurity_keepdays value to the UTM
func UpdateReportingWebsecurityKeepdays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.websecurity_keepdays", val, options...)
}

// GetReportingWebsecurityStatus gets the reporting.websecurity_status value from the UTM
func GetReportingWebsecurityStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reporting.websecurity_status", &val, options...)
	return
}

// UpdateReportingWebsecurityStatus PUTs the reporting.websecurity_status value to the UTM
func UpdateReportingWebsecurityStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reporting.websecurity_status", val, options...)
}

// GetReverseProxyAuaRefreshEnabled gets the reverse_proxy.aua_refresh_enabled value from the UTM
func GetReverseProxyAuaRefreshEnabled(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reverse_proxy.aua_refresh_enabled", &val, options...)
	return
}

// UpdateReverseProxyAuaRefreshEnabled PUTs the reverse_proxy.aua_refresh_enabled value to the UTM
func UpdateReverseProxyAuaRefreshEnabled(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.aua_refresh_enabled", val, options...)
}

// GetReverseProxyAuaRefreshInterval gets the reverse_proxy.aua_refresh_interval value from the UTM
func GetReverseProxyAuaRefreshInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.aua_refresh_interval", &val, options...)
	return
}

// UpdateReverseProxyAuaRefreshInterval PUTs the reverse_proxy.aua_refresh_interval value to the UTM
func UpdateReverseProxyAuaRefreshInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.aua_refresh_interval", val, options...)
}

// GetReverseProxyBlacklistDnsrblZones gets the reverse_proxy.blacklist.dnsrbl_zones value from the UTM
func GetReverseProxyBlacklistDnsrblZones(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/reverse_proxy.blacklist.dnsrbl_zones", &val, options...)
	return
}

// UpdateReverseProxyBlacklistDnsrblZones PUTs the reverse_proxy.blacklist.dnsrbl_zones value to the UTM
func UpdateReverseProxyBlacklistDnsrblZones(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.blacklist.dnsrbl_zones", val, options...)
}

// GetReverseProxyBlacklistGeoipCodes gets the reverse_proxy.blacklist.geoip_codes value from the UTM
func GetReverseProxyBlacklistGeoipCodes(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/reverse_proxy.blacklist.geoip_codes", &val, options...)
	return
}

// UpdateReverseProxyBlacklistGeoipCodes PUTs the reverse_proxy.blacklist.geoip_codes value to the UTM
func UpdateReverseProxyBlacklistGeoipCodes(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.blacklist.geoip_codes", val, options...)
}

// GetReverseProxyBlacklistSxlBlocksets gets the reverse_proxy.blacklist.sxl_blocksets value from the UTM
func GetReverseProxyBlacklistSxlBlocksets(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/reverse_proxy.blacklist.sxl_blocksets", &val, options...)
	return
}

// UpdateReverseProxyBlacklistSxlBlocksets PUTs the reverse_proxy.blacklist.sxl_blocksets value to the UTM
func UpdateReverseProxyBlacklistSxlBlocksets(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.blacklist.sxl_blocksets", val, options...)
}

// GetReverseProxyBlacklistSxlZone gets the reverse_proxy.blacklist.sxl_zone value from the UTM
func GetReverseProxyBlacklistSxlZone(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/reverse_proxy.blacklist.sxl_zone", &val, options...)
	return
}

// UpdateReverseProxyBlacklistSxlZone PUTs the reverse_proxy.blacklist.sxl_zone value to the UTM
func UpdateReverseProxyBlacklistSxlZone(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.blacklist.sxl_zone", val, options...)
}

// GetReverseProxyCookiesignkey gets the reverse_proxy.cookiesignkey value from the UTM
func GetReverseProxyCookiesignkey(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/reverse_proxy.cookiesignkey", &val, options...)
	return
}

// UpdateReverseProxyCookiesignkey PUTs the reverse_proxy.cookiesignkey value to the UTM
func UpdateReverseProxyCookiesignkey(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.cookiesignkey", val, options...)
}

// GetReverseProxyCssdHostname gets the reverse_proxy.cssd_hostname value from the UTM
func GetReverseProxyCssdHostname(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/reverse_proxy.cssd_hostname", &val, options...)
	return
}

// UpdateReverseProxyCssdHostname PUTs the reverse_proxy.cssd_hostname value to the UTM
func UpdateReverseProxyCssdHostname(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.cssd_hostname", val, options...)
}

// GetReverseProxyCssdPort gets the reverse_proxy.cssd_port value from the UTM
func GetReverseProxyCssdPort(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/reverse_proxy.cssd_port", &val, options...)
	return
}

// UpdateReverseProxyCssdPort PUTs the reverse_proxy.cssd_port value to the UTM
func UpdateReverseProxyCssdPort(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.cssd_port", val, options...)
}

// GetReverseProxyCustomThreatFiltersEnabled gets the reverse_proxy.custom_threat_filters_enabled value from the UTM
func GetReverseProxyCustomThreatFiltersEnabled(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reverse_proxy.custom_threat_filters_enabled", &val, options...)
	return
}

// UpdateReverseProxyCustomThreatFiltersEnabled PUTs the reverse_proxy.custom_threat_filters_enabled value to the UTM
func UpdateReverseProxyCustomThreatFiltersEnabled(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.custom_threat_filters_enabled", val, options...)
}

// GetReverseProxyFormhardeningSecret gets the reverse_proxy.formhardening_secret value from the UTM
func GetReverseProxyFormhardeningSecret(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/reverse_proxy.formhardening_secret", &val, options...)
	return
}

// UpdateReverseProxyFormhardeningSecret PUTs the reverse_proxy.formhardening_secret value to the UTM
func UpdateReverseProxyFormhardeningSecret(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.formhardening_secret", val, options...)
}

// GetReverseProxyKeepalive gets the reverse_proxy.keepalive value from the UTM
func GetReverseProxyKeepalive(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reverse_proxy.keepalive", &val, options...)
	return
}

// UpdateReverseProxyKeepalive PUTs the reverse_proxy.keepalive value to the UTM
func UpdateReverseProxyKeepalive(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.keepalive", val, options...)
}

// GetReverseProxyManualmode gets the reverse_proxy.manualmode value from the UTM
func GetReverseProxyManualmode(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reverse_proxy.manualmode", &val, options...)
	return
}

// UpdateReverseProxyManualmode PUTs the reverse_proxy.manualmode value to the UTM
func UpdateReverseProxyManualmode(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.manualmode", val, options...)
}

// GetReverseProxyMaxConnectionsPerChild gets the reverse_proxy.max_connections_per_child value from the UTM
func GetReverseProxyMaxConnectionsPerChild(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/reverse_proxy.max_connections_per_child", &val, options...)
	return
}

// UpdateReverseProxyMaxConnectionsPerChild PUTs the reverse_proxy.max_connections_per_child value to the UTM
func UpdateReverseProxyMaxConnectionsPerChild(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.max_connections_per_child", val, options...)
}

// GetReverseProxyMaxPreforkProcesses gets the reverse_proxy.max_prefork_processes value from the UTM
func GetReverseProxyMaxPreforkProcesses(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.max_prefork_processes", &val, options...)
	return
}

// UpdateReverseProxyMaxPreforkProcesses PUTs the reverse_proxy.max_prefork_processes value to the UTM
func UpdateReverseProxyMaxPreforkProcesses(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.max_prefork_processes", val, options...)
}

// GetReverseProxyMaxProcesses gets the reverse_proxy.max_processes value from the UTM
func GetReverseProxyMaxProcesses(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.max_processes", &val, options...)
	return
}

// UpdateReverseProxyMaxProcesses PUTs the reverse_proxy.max_processes value to the UTM
func UpdateReverseProxyMaxProcesses(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.max_processes", val, options...)
}

// GetReverseProxyMaxSessionFiles gets the reverse_proxy.max_session_files value from the UTM
func GetReverseProxyMaxSessionFiles(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.max_session_files", &val, options...)
	return
}

// UpdateReverseProxyMaxSessionFiles PUTs the reverse_proxy.max_session_files value to the UTM
func UpdateReverseProxyMaxSessionFiles(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.max_session_files", val, options...)
}

// GetReverseProxyMaxSpareProcesses gets the reverse_proxy.max_spare_processes value from the UTM
func GetReverseProxyMaxSpareProcesses(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.max_spare_processes", &val, options...)
	return
}

// UpdateReverseProxyMaxSpareProcesses PUTs the reverse_proxy.max_spare_processes value to the UTM
func UpdateReverseProxyMaxSpareProcesses(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.max_spare_processes", val, options...)
}

// GetReverseProxyMaxSpareThreads gets the reverse_proxy.max_spare_threads value from the UTM
func GetReverseProxyMaxSpareThreads(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.max_spare_threads", &val, options...)
	return
}

// UpdateReverseProxyMaxSpareThreads PUTs the reverse_proxy.max_spare_threads value to the UTM
func UpdateReverseProxyMaxSpareThreads(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.max_spare_threads", val, options...)
}

// GetReverseProxyMaxThreadsPerProcess gets the reverse_proxy.max_threads_per_process value from the UTM
func GetReverseProxyMaxThreadsPerProcess(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.max_threads_per_process", &val, options...)
	return
}

// UpdateReverseProxyMaxThreadsPerProcess PUTs the reverse_proxy.max_threads_per_process value to the UTM
func UpdateReverseProxyMaxThreadsPerProcess(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.max_threads_per_process", val, options...)
}

// GetReverseProxyMinSpareProcesses gets the reverse_proxy.min_spare_processes value from the UTM
func GetReverseProxyMinSpareProcesses(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.min_spare_processes", &val, options...)
	return
}

// UpdateReverseProxyMinSpareProcesses PUTs the reverse_proxy.min_spare_processes value to the UTM
func UpdateReverseProxyMinSpareProcesses(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.min_spare_processes", val, options...)
}

// GetReverseProxyMinSpareThreads gets the reverse_proxy.min_spare_threads value from the UTM
func GetReverseProxyMinSpareThreads(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.min_spare_threads", &val, options...)
	return
}

// UpdateReverseProxyMinSpareThreads PUTs the reverse_proxy.min_spare_threads value to the UTM
func UpdateReverseProxyMinSpareThreads(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.min_spare_threads", val, options...)
}

// GetReverseProxyMinTls gets the reverse_proxy.min_tls value from the UTM
func GetReverseProxyMinTls(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/reverse_proxy.min_tls", &val, options...)
	return
}

// UpdateReverseProxyMinTls PUTs the reverse_proxy.min_tls value to the UTM
func UpdateReverseProxyMinTls(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.min_tls", val, options...)
}

// GetReverseProxyModsecurityBeta gets the reverse_proxy.modsecurity_beta value from the UTM
func GetReverseProxyModsecurityBeta(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reverse_proxy.modsecurity_beta", &val, options...)
	return
}

// UpdateReverseProxyModsecurityBeta PUTs the reverse_proxy.modsecurity_beta value to the UTM
func UpdateReverseProxyModsecurityBeta(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.modsecurity_beta", val, options...)
}

// GetReverseProxyMpmMode gets the reverse_proxy.mpm_mode value from the UTM
func GetReverseProxyMpmMode(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/reverse_proxy.mpm_mode", &val, options...)
	return
}

// UpdateReverseProxyMpmMode PUTs the reverse_proxy.mpm_mode value to the UTM
func UpdateReverseProxyMpmMode(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.mpm_mode", val, options...)
}

// GetReverseProxyPatternversion gets the reverse_proxy.patternversion value from the UTM
func GetReverseProxyPatternversion(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/reverse_proxy.patternversion", &val, options...)
	return
}

// UpdateReverseProxyPatternversion PUTs the reverse_proxy.patternversion value to the UTM
func UpdateReverseProxyPatternversion(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.patternversion", val, options...)
}

// GetReverseProxyPort gets the reverse_proxy.port value from the UTM
func GetReverseProxyPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.port", &val, options...)
	return
}

// UpdateReverseProxyPort PUTs the reverse_proxy.port value to the UTM
func UpdateReverseProxyPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.port", val, options...)
}

// GetReverseProxyProxypassreverseForFrontend gets the reverse_proxy.proxypassreverse_for_frontend value from the UTM
func GetReverseProxyProxypassreverseForFrontend(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reverse_proxy.proxypassreverse_for_frontend", &val, options...)
	return
}

// UpdateReverseProxyProxypassreverseForFrontend PUTs the reverse_proxy.proxypassreverse_for_frontend value to the UTM
func UpdateReverseProxyProxypassreverseForFrontend(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.proxypassreverse_for_frontend", val, options...)
}

// GetReverseProxyProxyprotocol gets the reverse_proxy.proxyprotocol value from the UTM
func GetReverseProxyProxyprotocol(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reverse_proxy.proxyprotocol", &val, options...)
	return
}

// UpdateReverseProxyProxyprotocol PUTs the reverse_proxy.proxyprotocol value to the UTM
func UpdateReverseProxyProxyprotocol(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.proxyprotocol", val, options...)
}

// GetReverseProxyReloadMethod gets the reverse_proxy.reload_method value from the UTM
func GetReverseProxyReloadMethod(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/reverse_proxy.reload_method", &val, options...)
	return
}

// UpdateReverseProxyReloadMethod PUTs the reverse_proxy.reload_method value to the UTM
func UpdateReverseProxyReloadMethod(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.reload_method", val, options...)
}

// GetReverseProxyRequestLineLimit gets the reverse_proxy.request_line_limit value from the UTM
func GetReverseProxyRequestLineLimit(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/reverse_proxy.request_line_limit", &val, options...)
	return
}

// UpdateReverseProxyRequestLineLimit PUTs the reverse_proxy.request_line_limit value to the UTM
func UpdateReverseProxyRequestLineLimit(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.request_line_limit", val, options...)
}

// GetReverseProxySlowhttpExceptions gets the reverse_proxy.slowhttp_exceptions value from the UTM
func GetReverseProxySlowhttpExceptions(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/reverse_proxy.slowhttp_exceptions", &val, options...)
	return
}

// UpdateReverseProxySlowhttpExceptions PUTs the reverse_proxy.slowhttp_exceptions value to the UTM
func UpdateReverseProxySlowhttpExceptions(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.slowhttp_exceptions", val, options...)
}

// GetReverseProxySlowhttpRequestHeaderTimeoutBase gets the reverse_proxy.slowhttp_request_header_timeout_base value from the UTM
func GetReverseProxySlowhttpRequestHeaderTimeoutBase(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.slowhttp_request_header_timeout_base", &val, options...)
	return
}

// UpdateReverseProxySlowhttpRequestHeaderTimeoutBase PUTs the reverse_proxy.slowhttp_request_header_timeout_base value to the UTM
func UpdateReverseProxySlowhttpRequestHeaderTimeoutBase(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.slowhttp_request_header_timeout_base", val, options...)
}

// GetReverseProxySlowhttpRequestHeaderTimeoutEnabled gets the reverse_proxy.slowhttp_request_header_timeout_enabled value from the UTM
func GetReverseProxySlowhttpRequestHeaderTimeoutEnabled(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reverse_proxy.slowhttp_request_header_timeout_enabled", &val, options...)
	return
}

// UpdateReverseProxySlowhttpRequestHeaderTimeoutEnabled PUTs the reverse_proxy.slowhttp_request_header_timeout_enabled value to the UTM
func UpdateReverseProxySlowhttpRequestHeaderTimeoutEnabled(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.slowhttp_request_header_timeout_enabled", val, options...)
}

// GetReverseProxySlowhttpRequestHeaderTimeoutMax gets the reverse_proxy.slowhttp_request_header_timeout_max value from the UTM
func GetReverseProxySlowhttpRequestHeaderTimeoutMax(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.slowhttp_request_header_timeout_max", &val, options...)
	return
}

// UpdateReverseProxySlowhttpRequestHeaderTimeoutMax PUTs the reverse_proxy.slowhttp_request_header_timeout_max value to the UTM
func UpdateReverseProxySlowhttpRequestHeaderTimeoutMax(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.slowhttp_request_header_timeout_max", val, options...)
}

// GetReverseProxySlowhttpRequestHeaderTimeoutRate gets the reverse_proxy.slowhttp_request_header_timeout_rate value from the UTM
func GetReverseProxySlowhttpRequestHeaderTimeoutRate(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/reverse_proxy.slowhttp_request_header_timeout_rate", &val, options...)
	return
}

// UpdateReverseProxySlowhttpRequestHeaderTimeoutRate PUTs the reverse_proxy.slowhttp_request_header_timeout_rate value to the UTM
func UpdateReverseProxySlowhttpRequestHeaderTimeoutRate(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.slowhttp_request_header_timeout_rate", val, options...)
}

// GetReverseProxyStatus gets the reverse_proxy.status value from the UTM
func GetReverseProxyStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reverse_proxy.status", &val, options...)
	return
}

// UpdateReverseProxyStatus PUTs the reverse_proxy.status value to the UTM
func UpdateReverseProxyStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.status", val, options...)
}

// GetReverseProxyTraceEnabled gets the reverse_proxy.trace_enabled value from the UTM
func GetReverseProxyTraceEnabled(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reverse_proxy.trace_enabled", &val, options...)
	return
}

// UpdateReverseProxyTraceEnabled PUTs the reverse_proxy.trace_enabled value to the UTM
func UpdateReverseProxyTraceEnabled(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.trace_enabled", val, options...)
}

// GetReverseProxyUrlhardeningsignkey gets the reverse_proxy.urlhardeningsignkey value from the UTM
func GetReverseProxyUrlhardeningsignkey(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/reverse_proxy.urlhardeningsignkey", &val, options...)
	return
}

// UpdateReverseProxyUrlhardeningsignkey PUTs the reverse_proxy.urlhardeningsignkey value to the UTM
func UpdateReverseProxyUrlhardeningsignkey(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.urlhardeningsignkey", val, options...)
}

// GetReverseProxyWhatkilledus gets the reverse_proxy.whatkilledus value from the UTM
func GetReverseProxyWhatkilledus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/reverse_proxy.whatkilledus", &val, options...)
	return
}

// UpdateReverseProxyWhatkilledus PUTs the reverse_proxy.whatkilledus value to the UTM
func UpdateReverseProxyWhatkilledus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/reverse_proxy.whatkilledus", val, options...)
}

// GetRoutesPolicy gets the routes.policy value from the UTM
func GetRoutesPolicy(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/routes.policy", &val, options...)
	return
}

// UpdateRoutesPolicy PUTs the routes.policy value to the UTM
func UpdateRoutesPolicy(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routes.policy", val, options...)
}

// GetRoutesStatic gets the routes.static value from the UTM
func GetRoutesStatic(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/routes.static", &val, options...)
	return
}

// UpdateRoutesStatic PUTs the routes.static value to the UTM
func UpdateRoutesStatic(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routes.static", val, options...)
}

// GetRoutingBgpMaximumPaths gets the routing.bgp.maximum_paths value from the UTM
func GetRoutingBgpMaximumPaths(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/routing.bgp.maximum_paths", &val, options...)
	return
}

// UpdateRoutingBgpMaximumPaths PUTs the routing.bgp.maximum_paths value to the UTM
func UpdateRoutingBgpMaximumPaths(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.bgp.maximum_paths", val, options...)
}

// GetRoutingBgpMaximumPathsIbgp gets the routing.bgp.maximum_paths_ibgp value from the UTM
func GetRoutingBgpMaximumPathsIbgp(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/routing.bgp.maximum_paths_ibgp", &val, options...)
	return
}

// UpdateRoutingBgpMaximumPathsIbgp PUTs the routing.bgp.maximum_paths_ibgp value to the UTM
func UpdateRoutingBgpMaximumPathsIbgp(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.bgp.maximum_paths_ibgp", val, options...)
}

// GetRoutingBgpMultipleAs gets the routing.bgp.multiple_as value from the UTM
func GetRoutingBgpMultipleAs(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/routing.bgp.multiple_as", &val, options...)
	return
}

// UpdateRoutingBgpMultipleAs PUTs the routing.bgp.multiple_as value to the UTM
func UpdateRoutingBgpMultipleAs(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.bgp.multiple_as", val, options...)
}

// GetRoutingBgpStatus gets the routing.bgp.status value from the UTM
func GetRoutingBgpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/routing.bgp.status", &val, options...)
	return
}

// UpdateRoutingBgpStatus PUTs the routing.bgp.status value to the UTM
func UpdateRoutingBgpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.bgp.status", val, options...)
}

// GetRoutingBgpStrictMatch gets the routing.bgp.strict_match value from the UTM
func GetRoutingBgpStrictMatch(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/routing.bgp.strict_match", &val, options...)
	return
}

// UpdateRoutingBgpStrictMatch PUTs the routing.bgp.strict_match value to the UTM
func UpdateRoutingBgpStrictMatch(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.bgp.strict_match", val, options...)
}

// GetRoutingBgpSystems gets the routing.bgp.systems value from the UTM
func GetRoutingBgpSystems(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/routing.bgp.systems", &val, options...)
	return
}

// UpdateRoutingBgpSystems PUTs the routing.bgp.systems value to the UTM
func UpdateRoutingBgpSystems(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.bgp.systems", val, options...)
}

// GetRoutingOspfAbrType gets the routing.ospf.abr_type value from the UTM
func GetRoutingOspfAbrType(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/routing.ospf.abr_type", &val, options...)
	return
}

// UpdateRoutingOspfAbrType PUTs the routing.ospf.abr_type value to the UTM
func UpdateRoutingOspfAbrType(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.abr_type", val, options...)
}

// GetRoutingOspfAreas gets the routing.ospf.areas value from the UTM
func GetRoutingOspfAreas(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/routing.ospf.areas", &val, options...)
	return
}

// UpdateRoutingOspfAreas PUTs the routing.ospf.areas value to the UTM
func UpdateRoutingOspfAreas(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.areas", val, options...)
}

// GetRoutingOspfDefaultInformation gets the routing.ospf.default_information value from the UTM
func GetRoutingOspfDefaultInformation(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/routing.ospf.default_information", &val, options...)
	return
}

// UpdateRoutingOspfDefaultInformation PUTs the routing.ospf.default_information value to the UTM
func UpdateRoutingOspfDefaultInformation(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.default_information", val, options...)
}

// GetRoutingOspfDefaultRouteMetric gets the routing.ospf.default_route_metric value from the UTM
func GetRoutingOspfDefaultRouteMetric(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/routing.ospf.default_route_metric", &val, options...)
	return
}

// UpdateRoutingOspfDefaultRouteMetric PUTs the routing.ospf.default_route_metric value to the UTM
func UpdateRoutingOspfDefaultRouteMetric(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.default_route_metric", val, options...)
}

// GetRoutingOspfRedistributeBgpMetric gets the routing.ospf.redistribute.bgp.metric value from the UTM
func GetRoutingOspfRedistributeBgpMetric(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/routing.ospf.redistribute.bgp.metric", &val, options...)
	return
}

// UpdateRoutingOspfRedistributeBgpMetric PUTs the routing.ospf.redistribute.bgp.metric value to the UTM
func UpdateRoutingOspfRedistributeBgpMetric(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.redistribute.bgp.metric", val, options...)
}

// GetRoutingOspfRedistributeBgpStatus gets the routing.ospf.redistribute.bgp.status value from the UTM
func GetRoutingOspfRedistributeBgpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/routing.ospf.redistribute.bgp.status", &val, options...)
	return
}

// UpdateRoutingOspfRedistributeBgpStatus PUTs the routing.ospf.redistribute.bgp.status value to the UTM
func UpdateRoutingOspfRedistributeBgpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.redistribute.bgp.status", val, options...)
}

// GetRoutingOspfRedistributeConnectedMetric gets the routing.ospf.redistribute.connected.metric value from the UTM
func GetRoutingOspfRedistributeConnectedMetric(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/routing.ospf.redistribute.connected.metric", &val, options...)
	return
}

// UpdateRoutingOspfRedistributeConnectedMetric PUTs the routing.ospf.redistribute.connected.metric value to the UTM
func UpdateRoutingOspfRedistributeConnectedMetric(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.redistribute.connected.metric", val, options...)
}

// GetRoutingOspfRedistributeConnectedStatus gets the routing.ospf.redistribute.connected.status value from the UTM
func GetRoutingOspfRedistributeConnectedStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/routing.ospf.redistribute.connected.status", &val, options...)
	return
}

// UpdateRoutingOspfRedistributeConnectedStatus PUTs the routing.ospf.redistribute.connected.status value to the UTM
func UpdateRoutingOspfRedistributeConnectedStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.redistribute.connected.status", val, options...)
}

// GetRoutingOspfRedistributeIpsecStatus gets the routing.ospf.redistribute.ipsec.status value from the UTM
func GetRoutingOspfRedistributeIpsecStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/routing.ospf.redistribute.ipsec.status", &val, options...)
	return
}

// UpdateRoutingOspfRedistributeIpsecStatus PUTs the routing.ospf.redistribute.ipsec.status value to the UTM
func UpdateRoutingOspfRedistributeIpsecStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.redistribute.ipsec.status", val, options...)
}

// GetRoutingOspfRedistributeSslVpnStatus gets the routing.ospf.redistribute.ssl_vpn.status value from the UTM
func GetRoutingOspfRedistributeSslVpnStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/routing.ospf.redistribute.ssl_vpn.status", &val, options...)
	return
}

// UpdateRoutingOspfRedistributeSslVpnStatus PUTs the routing.ospf.redistribute.ssl_vpn.status value to the UTM
func UpdateRoutingOspfRedistributeSslVpnStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.redistribute.ssl_vpn.status", val, options...)
}

// GetRoutingOspfRedistributeStaticMetric gets the routing.ospf.redistribute.static.metric value from the UTM
func GetRoutingOspfRedistributeStaticMetric(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/routing.ospf.redistribute.static.metric", &val, options...)
	return
}

// UpdateRoutingOspfRedistributeStaticMetric PUTs the routing.ospf.redistribute.static.metric value to the UTM
func UpdateRoutingOspfRedistributeStaticMetric(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.redistribute.static.metric", val, options...)
}

// GetRoutingOspfRedistributeStaticStatus gets the routing.ospf.redistribute.static.status value from the UTM
func GetRoutingOspfRedistributeStaticStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/routing.ospf.redistribute.static.status", &val, options...)
	return
}

// UpdateRoutingOspfRedistributeStaticStatus PUTs the routing.ospf.redistribute.static.status value to the UTM
func UpdateRoutingOspfRedistributeStaticStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.redistribute.static.status", val, options...)
}

// GetRoutingOspfReferenceBandwidth gets the routing.ospf.reference_bandwidth value from the UTM
func GetRoutingOspfReferenceBandwidth(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/routing.ospf.reference_bandwidth", &val, options...)
	return
}

// UpdateRoutingOspfReferenceBandwidth PUTs the routing.ospf.reference_bandwidth value to the UTM
func UpdateRoutingOspfReferenceBandwidth(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.reference_bandwidth", val, options...)
}

// GetRoutingOspfRouterId gets the routing.ospf.router_id value from the UTM
func GetRoutingOspfRouterId(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/routing.ospf.router_id", &val, options...)
	return
}

// UpdateRoutingOspfRouterId PUTs the routing.ospf.router_id value to the UTM
func UpdateRoutingOspfRouterId(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.router_id", val, options...)
}

// GetRoutingOspfStatus gets the routing.ospf.status value from the UTM
func GetRoutingOspfStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/routing.ospf.status", &val, options...)
	return
}

// UpdateRoutingOspfStatus PUTs the routing.ospf.status value to the UTM
func UpdateRoutingOspfStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.ospf.status", val, options...)
}

// GetRoutingQuaggaAllowedNetworks gets the routing.quagga.allowed_networks value from the UTM
func GetRoutingQuaggaAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/routing.quagga.allowed_networks", &val, options...)
	return
}

// UpdateRoutingQuaggaAllowedNetworks PUTs the routing.quagga.allowed_networks value to the UTM
func UpdateRoutingQuaggaAllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.quagga.allowed_networks", val, options...)
}

// GetRoutingQuaggaEnablePassword gets the routing.quagga.enable_password value from the UTM
func GetRoutingQuaggaEnablePassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/routing.quagga.enable_password", &val, options...)
	return
}

// UpdateRoutingQuaggaEnablePassword PUTs the routing.quagga.enable_password value to the UTM
func UpdateRoutingQuaggaEnablePassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.quagga.enable_password", val, options...)
}

// GetRoutingQuaggaLinkDetect gets the routing.quagga.link_detect value from the UTM
func GetRoutingQuaggaLinkDetect(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/routing.quagga.link_detect", &val, options...)
	return
}

// UpdateRoutingQuaggaLinkDetect PUTs the routing.quagga.link_detect value to the UTM
func UpdateRoutingQuaggaLinkDetect(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.quagga.link_detect", val, options...)
}

// GetRoutingQuaggaPassword gets the routing.quagga.password value from the UTM
func GetRoutingQuaggaPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/routing.quagga.password", &val, options...)
	return
}

// UpdateRoutingQuaggaPassword PUTs the routing.quagga.password value to the UTM
func UpdateRoutingQuaggaPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/routing.quagga.password", val, options...)
}

// GetSandboxReportdDebug gets the sandbox_reportd.debug value from the UTM
func GetSandboxReportdDebug(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sandbox_reportd.debug", &val, options...)
	return
}

// UpdateSandboxReportdDebug PUTs the sandbox_reportd.debug value to the UTM
func UpdateSandboxReportdDebug(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandbox_reportd.debug", val, options...)
}

// GetSandboxReportdPollInterval gets the sandbox_reportd.poll_interval value from the UTM
func GetSandboxReportdPollInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandbox_reportd.poll_interval", &val, options...)
	return
}

// UpdateSandboxReportdPollInterval PUTs the sandbox_reportd.poll_interval value to the UTM
func UpdateSandboxReportdPollInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandbox_reportd.poll_interval", val, options...)
}

// GetSandboxReportdRequestTimeout gets the sandbox_reportd.request_timeout value from the UTM
func GetSandboxReportdRequestTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandbox_reportd.request_timeout", &val, options...)
	return
}

// UpdateSandboxReportdRequestTimeout PUTs the sandbox_reportd.request_timeout value to the UTM
func UpdateSandboxReportdRequestTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandbox_reportd.request_timeout", val, options...)
}

// GetSandboxReportdRetryTimes gets the sandbox_reportd.retry_times value from the UTM
func GetSandboxReportdRetryTimes(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandbox_reportd.retry_times", &val, options...)
	return
}

// UpdateSandboxReportdRetryTimes PUTs the sandbox_reportd.retry_times value to the UTM
func UpdateSandboxReportdRetryTimes(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandbox_reportd.retry_times", val, options...)
}

// GetSandboxdBypassSandboxLimit gets the sandboxd.bypass_sandbox_limit value from the UTM
func GetSandboxdBypassSandboxLimit(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.bypass_sandbox_limit", &val, options...)
	return
}

// UpdateSandboxdBypassSandboxLimit PUTs the sandboxd.bypass_sandbox_limit value to the UTM
func UpdateSandboxdBypassSandboxLimit(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.bypass_sandbox_limit", val, options...)
}

// GetSandboxdCacheExpire gets the sandboxd.cache_expire value from the UTM
func GetSandboxdCacheExpire(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.cache_expire", &val, options...)
	return
}

// UpdateSandboxdCacheExpire PUTs the sandboxd.cache_expire value to the UTM
func UpdateSandboxdCacheExpire(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.cache_expire", val, options...)
}

// GetSandboxdCertstore gets the sandboxd.certstore value from the UTM
func GetSandboxdCertstore(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sandboxd.certstore", &val, options...)
	return
}

// UpdateSandboxdCertstore PUTs the sandboxd.certstore value to the UTM
func UpdateSandboxdCertstore(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.certstore", val, options...)
}

// GetSandboxdClientTimeout gets the sandboxd.client_timeout value from the UTM
func GetSandboxdClientTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.client_timeout", &val, options...)
	return
}

// UpdateSandboxdClientTimeout PUTs the sandboxd.client_timeout value to the UTM
func UpdateSandboxdClientTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.client_timeout", val, options...)
}

// GetSandboxdCloudPollInterval gets the sandboxd.cloud_poll_interval value from the UTM
func GetSandboxdCloudPollInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.cloud_poll_interval", &val, options...)
	return
}

// UpdateSandboxdCloudPollInterval PUTs the sandboxd.cloud_poll_interval value to the UTM
func UpdateSandboxdCloudPollInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.cloud_poll_interval", val, options...)
}

// GetSandboxdCloudPollRequestMaximum gets the sandboxd.cloud_poll_request_maximum value from the UTM
func GetSandboxdCloudPollRequestMaximum(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.cloud_poll_request_maximum", &val, options...)
	return
}

// UpdateSandboxdCloudPollRequestMaximum PUTs the sandboxd.cloud_poll_request_maximum value to the UTM
func UpdateSandboxdCloudPollRequestMaximum(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.cloud_poll_request_maximum", val, options...)
}

// GetSandboxdCloudPollTimeout gets the sandboxd.cloud_poll_timeout value from the UTM
func GetSandboxdCloudPollTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.cloud_poll_timeout", &val, options...)
	return
}

// UpdateSandboxdCloudPollTimeout PUTs the sandboxd.cloud_poll_timeout value to the UTM
func UpdateSandboxdCloudPollTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.cloud_poll_timeout", val, options...)
}

// GetSandboxdConnectTimeout gets the sandboxd.connect_timeout value from the UTM
func GetSandboxdConnectTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.connect_timeout", &val, options...)
	return
}

// UpdateSandboxdConnectTimeout PUTs the sandboxd.connect_timeout value to the UTM
func UpdateSandboxdConnectTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.connect_timeout", val, options...)
}

// GetSandboxdDebug gets the sandboxd.debug value from the UTM
func GetSandboxdDebug(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/sandboxd.debug", &val, options...)
	return
}

// UpdateSandboxdDebug PUTs the sandboxd.debug value to the UTM
func UpdateSandboxdDebug(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.debug", val, options...)
}

// GetSandboxdDhparams2048 gets the sandboxd.dhparams_2048 value from the UTM
func GetSandboxdDhparams2048(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sandboxd.dhparams_2048", &val, options...)
	return
}

// UpdateSandboxdDhparams2048 PUTs the sandboxd.dhparams_2048 value to the UTM
func UpdateSandboxdDhparams2048(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.dhparams_2048", val, options...)
}

// GetSandboxdEdgeServerCertValidation gets the sandboxd.edge_server_cert_validation value from the UTM
func GetSandboxdEdgeServerCertValidation(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/sandboxd.edge_server_cert_validation", &val, options...)
	return
}

// UpdateSandboxdEdgeServerCertValidation PUTs the sandboxd.edge_server_cert_validation value to the UTM
func UpdateSandboxdEdgeServerCertValidation(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.edge_server_cert_validation", val, options...)
}

// GetSandboxdEdgeServerHost gets the sandboxd.edge_server_host value from the UTM
func GetSandboxdEdgeServerHost(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sandboxd.edge_server_host", &val, options...)
	return
}

// UpdateSandboxdEdgeServerHost PUTs the sandboxd.edge_server_host value to the UTM
func UpdateSandboxdEdgeServerHost(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.edge_server_host", val, options...)
}

// GetSandboxdEdgeServerPort gets the sandboxd.edge_server_port value from the UTM
func GetSandboxdEdgeServerPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.edge_server_port", &val, options...)
	return
}

// UpdateSandboxdEdgeServerPort PUTs the sandboxd.edge_server_port value to the UTM
func UpdateSandboxdEdgeServerPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.edge_server_port", val, options...)
}

// GetSandboxdFiletypeSkiplist gets the sandboxd.filetype_skiplist value from the UTM
func GetSandboxdFiletypeSkiplist(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/sandboxd.filetype_skiplist", &val, options...)
	return
}

// UpdateSandboxdFiletypeSkiplist PUTs the sandboxd.filetype_skiplist value to the UTM
func UpdateSandboxdFiletypeSkiplist(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.filetype_skiplist", val, options...)
}

// GetSandboxdNumOfEventThreads gets the sandboxd.num_of_event_threads value from the UTM
func GetSandboxdNumOfEventThreads(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.num_of_event_threads", &val, options...)
	return
}

// UpdateSandboxdNumOfEventThreads PUTs the sandboxd.num_of_event_threads value to the UTM
func UpdateSandboxdNumOfEventThreads(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.num_of_event_threads", val, options...)
}

// GetSandboxdNumOfWorkerThreads gets the sandboxd.num_of_worker_threads value from the UTM
func GetSandboxdNumOfWorkerThreads(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.num_of_worker_threads", &val, options...)
	return
}

// UpdateSandboxdNumOfWorkerThreads PUTs the sandboxd.num_of_worker_threads value to the UTM
func UpdateSandboxdNumOfWorkerThreads(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.num_of_worker_threads", val, options...)
}

// GetSandboxdResponseTimeout gets the sandboxd.response_timeout value from the UTM
func GetSandboxdResponseTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.response_timeout", &val, options...)
	return
}

// UpdateSandboxdResponseTimeout PUTs the sandboxd.response_timeout value to the UTM
func UpdateSandboxdResponseTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.response_timeout", val, options...)
}

// GetSandboxdRetryIntervals gets the sandboxd.retry_intervals value from the UTM
func GetSandboxdRetryIntervals(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sandboxd.retry_intervals", &val, options...)
	return
}

// UpdateSandboxdRetryIntervals PUTs the sandboxd.retry_intervals value to the UTM
func UpdateSandboxdRetryIntervals(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.retry_intervals", val, options...)
}

// GetSandboxdSandboxEnabled gets the sandboxd.sandbox_enabled value from the UTM
func GetSandboxdSandboxEnabled(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/sandboxd.sandbox_enabled", &val, options...)
	return
}

// UpdateSandboxdSandboxEnabled PUTs the sandboxd.sandbox_enabled value to the UTM
func UpdateSandboxdSandboxEnabled(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.sandbox_enabled", val, options...)
}

// GetSandboxdSandboxdScoreThreshold gets the sandboxd.sandboxd_score_threshold value from the UTM
func GetSandboxdSandboxdScoreThreshold(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.sandboxd_score_threshold", &val, options...)
	return
}

// UpdateSandboxdSandboxdScoreThreshold PUTs the sandboxd.sandboxd_score_threshold value to the UTM
func UpdateSandboxdSandboxdScoreThreshold(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.sandboxd_score_threshold", val, options...)
}

// GetSandboxdSbxVersion gets the sandboxd.sbx_version value from the UTM
func GetSandboxdSbxVersion(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.sbx_version", &val, options...)
	return
}

// UpdateSandboxdSbxVersion PUTs the sandboxd.sbx_version value to the UTM
func UpdateSandboxdSbxVersion(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.sbx_version", val, options...)
}

// GetSandboxdSqliteBusyTimeout gets the sandboxd.sqlite_busy_timeout value from the UTM
func GetSandboxdSqliteBusyTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/sandboxd.sqlite_busy_timeout", &val, options...)
	return
}

// UpdateSandboxdSqliteBusyTimeout PUTs the sandboxd.sqlite_busy_timeout value to the UTM
func UpdateSandboxdSqliteBusyTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.sqlite_busy_timeout", val, options...)
}

// GetSandboxdSslCertFile gets the sandboxd.ssl_cert_file value from the UTM
func GetSandboxdSslCertFile(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sandboxd.ssl_cert_file", &val, options...)
	return
}

// UpdateSandboxdSslCertFile PUTs the sandboxd.ssl_cert_file value to the UTM
func UpdateSandboxdSslCertFile(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.ssl_cert_file", val, options...)
}

// GetSandboxdSslKeyFile gets the sandboxd.ssl_key_file value from the UTM
func GetSandboxdSslKeyFile(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sandboxd.ssl_key_file", &val, options...)
	return
}

// UpdateSandboxdSslKeyFile PUTs the sandboxd.ssl_key_file value to the UTM
func UpdateSandboxdSslKeyFile(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.ssl_key_file", val, options...)
}

// GetSandboxdSslciphers gets the sandboxd.sslciphers value from the UTM
func GetSandboxdSslciphers(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sandboxd.sslciphers", &val, options...)
	return
}

// UpdateSandboxdSslciphers PUTs the sandboxd.sslciphers value to the UTM
func UpdateSandboxdSslciphers(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sandboxd.sslciphers", val, options...)
}

// GetSettingsAdminEmail gets the settings.admin_email value from the UTM
func GetSettingsAdminEmail(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/settings.admin_email", &val, options...)
	return
}

// UpdateSettingsAdminEmail PUTs the settings.admin_email value to the UTM
func UpdateSettingsAdminEmail(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.admin_email", val, options...)
}

// GetSettingsBasicSetup gets the settings.basic_setup value from the UTM
func GetSettingsBasicSetup(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/settings.basic_setup", &val, options...)
	return
}

// UpdateSettingsBasicSetup PUTs the settings.basic_setup value to the UTM
func UpdateSettingsBasicSetup(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.basic_setup", val, options...)
}

// GetSettingsCcMode gets the settings.cc_mode value from the UTM
func GetSettingsCcMode(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/settings.cc_mode", &val, options...)
	return
}

// UpdateSettingsCcMode PUTs the settings.cc_mode value to the UTM
func UpdateSettingsCcMode(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.cc_mode", val, options...)
}

// GetSettingsCity gets the settings.city value from the UTM
func GetSettingsCity(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/settings.city", &val, options...)
	return
}

// UpdateSettingsCity PUTs the settings.city value to the UTM
func UpdateSettingsCity(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.city", val, options...)
}

// GetSettingsCountry gets the settings.country value from the UTM
func GetSettingsCountry(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/settings.country", &val, options...)
	return
}

// UpdateSettingsCountry PUTs the settings.country value to the UTM
func UpdateSettingsCountry(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.country", val, options...)
}

// GetSettingsExtraSwap gets the settings.extra_swap value from the UTM
func GetSettingsExtraSwap(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/settings.extra_swap", &val, options...)
	return
}

// UpdateSettingsExtraSwap PUTs the settings.extra_swap value to the UTM
func UpdateSettingsExtraSwap(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.extra_swap", val, options...)
}

// GetSettingsHostname gets the settings.hostname value from the UTM
func GetSettingsHostname(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/settings.hostname", &val, options...)
	return
}

// UpdateSettingsHostname PUTs the settings.hostname value to the UTM
func UpdateSettingsHostname(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.hostname", val, options...)
}

// GetSettingsIcsaMode gets the settings.icsa_mode value from the UTM
func GetSettingsIcsaMode(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/settings.icsa_mode", &val, options...)
	return
}

// UpdateSettingsIcsaMode PUTs the settings.icsa_mode value to the UTM
func UpdateSettingsIcsaMode(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.icsa_mode", val, options...)
}

// GetSettingsOrganization gets the settings.organization value from the UTM
func GetSettingsOrganization(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/settings.organization", &val, options...)
	return
}

// UpdateSettingsOrganization PUTs the settings.organization value to the UTM
func UpdateSettingsOrganization(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.organization", val, options...)
}

// GetSettingsPasswordComplexityMinDigits gets the settings.password_complexity.min_digits value from the UTM
func GetSettingsPasswordComplexityMinDigits(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/settings.password_complexity.min_digits", &val, options...)
	return
}

// UpdateSettingsPasswordComplexityMinDigits PUTs the settings.password_complexity.min_digits value to the UTM
func UpdateSettingsPasswordComplexityMinDigits(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.password_complexity.min_digits", val, options...)
}

// GetSettingsPasswordComplexityMinLength gets the settings.password_complexity.min_length value from the UTM
func GetSettingsPasswordComplexityMinLength(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/settings.password_complexity.min_length", &val, options...)
	return
}

// UpdateSettingsPasswordComplexityMinLength PUTs the settings.password_complexity.min_length value to the UTM
func UpdateSettingsPasswordComplexityMinLength(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.password_complexity.min_length", val, options...)
}

// GetSettingsPasswordComplexityMinLowerChars gets the settings.password_complexity.min_lower_chars value from the UTM
func GetSettingsPasswordComplexityMinLowerChars(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/settings.password_complexity.min_lower_chars", &val, options...)
	return
}

// UpdateSettingsPasswordComplexityMinLowerChars PUTs the settings.password_complexity.min_lower_chars value to the UTM
func UpdateSettingsPasswordComplexityMinLowerChars(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.password_complexity.min_lower_chars", val, options...)
}

// GetSettingsPasswordComplexityMinSpecialChars gets the settings.password_complexity.min_special_chars value from the UTM
func GetSettingsPasswordComplexityMinSpecialChars(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/settings.password_complexity.min_special_chars", &val, options...)
	return
}

// UpdateSettingsPasswordComplexityMinSpecialChars PUTs the settings.password_complexity.min_special_chars value to the UTM
func UpdateSettingsPasswordComplexityMinSpecialChars(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.password_complexity.min_special_chars", val, options...)
}

// GetSettingsPasswordComplexityMinUpperChars gets the settings.password_complexity.min_upper_chars value from the UTM
func GetSettingsPasswordComplexityMinUpperChars(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/settings.password_complexity.min_upper_chars", &val, options...)
	return
}

// UpdateSettingsPasswordComplexityMinUpperChars PUTs the settings.password_complexity.min_upper_chars value to the UTM
func UpdateSettingsPasswordComplexityMinUpperChars(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.password_complexity.min_upper_chars", val, options...)
}

// GetSettingsPasswordComplexityStatus gets the settings.password_complexity.status value from the UTM
func GetSettingsPasswordComplexityStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/settings.password_complexity.status", &val, options...)
	return
}

// UpdateSettingsPasswordComplexityStatus PUTs the settings.password_complexity.status value to the UTM
func UpdateSettingsPasswordComplexityStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.password_complexity.status", val, options...)
}

// GetSettingsPopularity gets the settings.popularity value from the UTM
func GetSettingsPopularity(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/settings.popularity", &val, options...)
	return
}

// UpdateSettingsPopularity PUTs the settings.popularity value to the UTM
func UpdateSettingsPopularity(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.popularity", val, options...)
}

// GetSettingsRasUpdate gets the settings.ras_update value from the UTM
func GetSettingsRasUpdate(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/settings.ras_update", &val, options...)
	return
}

// UpdateSettingsRasUpdate PUTs the settings.ras_update value to the UTM
func UpdateSettingsRasUpdate(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.ras_update", val, options...)
}

// GetSettingsSystemId gets the settings.system_id value from the UTM
func GetSettingsSystemId(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/settings.system_id", &val, options...)
	return
}

// UpdateSettingsSystemId PUTs the settings.system_id value to the UTM
func UpdateSettingsSystemId(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.system_id", val, options...)
}

// GetSettingsTimezone gets the settings.timezone value from the UTM
func GetSettingsTimezone(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/settings.timezone", &val, options...)
	return
}

// UpdateSettingsTimezone PUTs the settings.timezone value to the UTM
func UpdateSettingsTimezone(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/settings.timezone", val, options...)
}

// GetSipAllowedNetworks gets the sip.allowed_networks value from the UTM
func GetSipAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/sip.allowed_networks", &val, options...)
	return
}

// UpdateSipAllowedNetworks PUTs the sip.allowed_networks value to the UTM
func UpdateSipAllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sip.allowed_networks", val, options...)
}

// GetSipExpectMode gets the sip.expect_mode value from the UTM
func GetSipExpectMode(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sip.expect_mode", &val, options...)
	return
}

// UpdateSipExpectMode PUTs the sip.expect_mode value to the UTM
func UpdateSipExpectMode(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sip.expect_mode", val, options...)
}

// GetSipLogRelated gets the sip.log_related value from the UTM
func GetSipLogRelated(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/sip.log_related", &val, options...)
	return
}

// UpdateSipLogRelated PUTs the sip.log_related value to the UTM
func UpdateSipLogRelated(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sip.log_related", val, options...)
}

// GetSipServers gets the sip.servers value from the UTM
func GetSipServers(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/sip.servers", &val, options...)
	return
}

// UpdateSipServers PUTs the sip.servers value to the UTM
func UpdateSipServers(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sip.servers", val, options...)
}

// GetSipStatus gets the sip.status value from the UTM
func GetSipStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/sip.status", &val, options...)
	return
}

// UpdateSipStatus PUTs the sip.status value to the UTM
func UpdateSipStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sip.status", val, options...)
}

// GetSmsClientHostname gets the sms_client.hostname value from the UTM
func GetSmsClientHostname(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sms_client.hostname", &val, options...)
	return
}

// UpdateSmsClientHostname PUTs the sms_client.hostname value to the UTM
func UpdateSmsClientHostname(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sms_client.hostname", val, options...)
}

// GetSmsClientName gets the sms_client.name value from the UTM
func GetSmsClientName(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sms_client.name", &val, options...)
	return
}

// UpdateSmsClientName PUTs the sms_client.name value to the UTM
func UpdateSmsClientName(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sms_client.name", val, options...)
}

// GetSmsClientPassword gets the sms_client.password value from the UTM
func GetSmsClientPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sms_client.password", &val, options...)
	return
}

// UpdateSmsClientPassword PUTs the sms_client.password value to the UTM
func UpdateSmsClientPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sms_client.password", val, options...)
}

// GetSmsClientPort gets the sms_client.port value from the UTM
func GetSmsClientPort(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/sms_client.port", &val, options...)
	return
}

// UpdateSmsClientPort PUTs the sms_client.port value to the UTM
func UpdateSmsClientPort(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sms_client.port", val, options...)
}

// GetSmsClientStatus gets the sms_client.status value from the UTM
func GetSmsClientStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/sms_client.status", &val, options...)
	return
}

// UpdateSmsClientStatus PUTs the sms_client.status value to the UTM
func UpdateSmsClientStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sms_client.status", val, options...)
}

// GetSmsClientUsername gets the sms_client.username value from the UTM
func GetSmsClientUsername(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/sms_client.username", &val, options...)
	return
}

// UpdateSmsClientUsername PUTs the sms_client.username value to the UTM
func UpdateSmsClientUsername(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/sms_client.username", val, options...)
}

// GetSmtpAdLookupContacts gets the smtp.ad_lookup_contacts value from the UTM
func GetSmtpAdLookupContacts(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.ad_lookup_contacts", &val, options...)
	return
}

// UpdateSmtpAdLookupContacts PUTs the smtp.ad_lookup_contacts value to the UTM
func UpdateSmtpAdLookupContacts(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.ad_lookup_contacts", val, options...)
}

// GetSmtpAuthAaa gets the smtp.auth_aaa value from the UTM
func GetSmtpAuthAaa(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/smtp.auth_aaa", &val, options...)
	return
}

// UpdateSmtpAuthAaa PUTs the smtp.auth_aaa value to the UTM
func UpdateSmtpAuthAaa(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.auth_aaa", val, options...)
}

// GetSmtpAuthStatus gets the smtp.auth_status value from the UTM
func GetSmtpAuthStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.auth_status", &val, options...)
	return
}

// UpdateSmtpAuthStatus PUTs the smtp.auth_status value to the UTM
func UpdateSmtpAuthStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.auth_status", val, options...)
}

// GetSmtpAvFooter gets the smtp.av_footer value from the UTM
func GetSmtpAvFooter(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.av_footer", &val, options...)
	return
}

// UpdateSmtpAvFooter PUTs the smtp.av_footer value to the UTM
func UpdateSmtpAvFooter(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.av_footer", val, options...)
}

// GetSmtpAvFooterStatus gets the smtp.av_footer_status value from the UTM
func GetSmtpAvFooterStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.av_footer_status", &val, options...)
	return
}

// UpdateSmtpAvFooterStatus PUTs the smtp.av_footer_status value to the UTM
func UpdateSmtpAvFooterStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.av_footer_status", val, options...)
}

// GetSmtpBatvSecret gets the smtp.batv_secret value from the UTM
func GetSmtpBatvSecret(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.batv_secret", &val, options...)
	return
}

// UpdateSmtpBatvSecret PUTs the smtp.batv_secret value to the UTM
func UpdateSmtpBatvSecret(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.batv_secret", val, options...)
}

// GetSmtpBlockerService gets the smtp.blocker_service value from the UTM
func GetSmtpBlockerService(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.blocker_service", &val, options...)
	return
}

// UpdateSmtpBlockerService PUTs the smtp.blocker_service value to the UTM
func UpdateSmtpBlockerService(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.blocker_service", val, options...)
}

// GetSmtpCffAsMarker gets the smtp.cff_as_marker value from the UTM
func GetSmtpCffAsMarker(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.cff_as_marker", &val, options...)
	return
}

// UpdateSmtpCffAsMarker PUTs the smtp.cff_as_marker value to the UTM
func UpdateSmtpCffAsMarker(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.cff_as_marker", val, options...)
}

// GetSmtpDkimDomains gets the smtp.dkim_domains value from the UTM
func GetSmtpDkimDomains(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/smtp.dkim_domains", &val, options...)
	return
}

// UpdateSmtpDkimDomains PUTs the smtp.dkim_domains value to the UTM
func UpdateSmtpDkimDomains(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.dkim_domains", val, options...)
}

// GetSmtpDkimPrivateKey gets the smtp.dkim_private_key value from the UTM
func GetSmtpDkimPrivateKey(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.dkim_private_key", &val, options...)
	return
}

// UpdateSmtpDkimPrivateKey PUTs the smtp.dkim_private_key value to the UTM
func UpdateSmtpDkimPrivateKey(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.dkim_private_key", val, options...)
}

// GetSmtpDkimSelector gets the smtp.dkim_selector value from the UTM
func GetSmtpDkimSelector(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.dkim_selector", &val, options...)
	return
}

// UpdateSmtpDkimSelector PUTs the smtp.dkim_selector value to the UTM
func UpdateSmtpDkimSelector(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.dkim_selector", val, options...)
}

// GetSmtpEnableOldExpressionFilter gets the smtp.enable_old_expression_filter value from the UTM
func GetSmtpEnableOldExpressionFilter(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.enable_old_expression_filter", &val, options...)
	return
}

// UpdateSmtpEnableOldExpressionFilter PUTs the smtp.enable_old_expression_filter value to the UTM
func UpdateSmtpEnableOldExpressionFilter(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.enable_old_expression_filter", val, options...)
}

// GetSmtpEncryptionUtility gets the smtp.encryption_utility value from the UTM
func GetSmtpEncryptionUtility(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.encryption_utility", &val, options...)
	return
}

// UpdateSmtpEncryptionUtility PUTs the smtp.encryption_utility value to the UTM
func UpdateSmtpEncryptionUtility(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.encryption_utility", val, options...)
}

// GetSmtpExceptions gets the smtp.exceptions value from the UTM
func GetSmtpExceptions(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/smtp.exceptions", &val, options...)
	return
}

// UpdateSmtpExceptions PUTs the smtp.exceptions value to the UTM
func UpdateSmtpExceptions(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.exceptions", val, options...)
}

// GetSmtpFootersMode gets the smtp.footers_mode value from the UTM
func GetSmtpFootersMode(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.footers_mode", &val, options...)
	return
}

// UpdateSmtpFootersMode PUTs the smtp.footers_mode value to the UTM
func UpdateSmtpFootersMode(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.footers_mode", val, options...)
}

// GetSmtpGlobalAsReject gets the smtp.global_as_reject value from the UTM
func GetSmtpGlobalAsReject(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.global_as_reject", &val, options...)
	return
}

// UpdateSmtpGlobalAsReject PUTs the smtp.global_as_reject value to the UTM
func UpdateSmtpGlobalAsReject(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.global_as_reject", val, options...)
}

// GetSmtpGlobalAvReject gets the smtp.global_av_reject value from the UTM
func GetSmtpGlobalAvReject(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.global_av_reject", &val, options...)
	return
}

// UpdateSmtpGlobalAvReject PUTs the smtp.global_av_reject value to the UTM
func UpdateSmtpGlobalAvReject(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.global_av_reject", val, options...)
}

// GetSmtpGlobalProfile gets the smtp.global_profile value from the UTM
func GetSmtpGlobalProfile(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.global_profile", &val, options...)
	return
}

// UpdateSmtpGlobalProfile PUTs the smtp.global_profile value to the UTM
func UpdateSmtpGlobalProfile(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.global_profile", val, options...)
}

// GetSmtpHostBlacklist gets the smtp.host_blacklist value from the UTM
func GetSmtpHostBlacklist(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/smtp.host_blacklist", &val, options...)
	return
}

// UpdateSmtpHostBlacklist PUTs the smtp.host_blacklist value to the UTM
func UpdateSmtpHostBlacklist(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.host_blacklist", val, options...)
}

// GetSmtpHostname gets the smtp.hostname value from the UTM
func GetSmtpHostname(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.hostname", &val, options...)
	return
}

// UpdateSmtpHostname PUTs the smtp.hostname value to the UTM
func UpdateSmtpHostname(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.hostname", val, options...)
}

// GetSmtpMaxMessageSize gets the smtp.max_message_size value from the UTM
func GetSmtpMaxMessageSize(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/smtp.max_message_size", &val, options...)
	return
}

// UpdateSmtpMaxMessageSize PUTs the smtp.max_message_size value to the UTM
func UpdateSmtpMaxMessageSize(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.max_message_size", val, options...)
}

// GetSmtpMode gets the smtp.mode value from the UTM
func GetSmtpMode(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.mode", &val, options...)
	return
}

// UpdateSmtpMode PUTs the smtp.mode value to the UTM
func UpdateSmtpMode(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.mode", val, options...)
}

// GetSmtpParentProxyAuthPass gets the smtp.parent_proxy_auth_pass value from the UTM
func GetSmtpParentProxyAuthPass(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.parent_proxy_auth_pass", &val, options...)
	return
}

// UpdateSmtpParentProxyAuthPass PUTs the smtp.parent_proxy_auth_pass value to the UTM
func UpdateSmtpParentProxyAuthPass(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.parent_proxy_auth_pass", val, options...)
}

// GetSmtpParentProxyAuthStatus gets the smtp.parent_proxy_auth_status value from the UTM
func GetSmtpParentProxyAuthStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.parent_proxy_auth_status", &val, options...)
	return
}

// UpdateSmtpParentProxyAuthStatus PUTs the smtp.parent_proxy_auth_status value to the UTM
func UpdateSmtpParentProxyAuthStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.parent_proxy_auth_status", val, options...)
}

// GetSmtpParentProxyAuthUser gets the smtp.parent_proxy_auth_user value from the UTM
func GetSmtpParentProxyAuthUser(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.parent_proxy_auth_user", &val, options...)
	return
}

// UpdateSmtpParentProxyAuthUser PUTs the smtp.parent_proxy_auth_user value to the UTM
func UpdateSmtpParentProxyAuthUser(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.parent_proxy_auth_user", val, options...)
}

// GetSmtpParentProxyHost gets the smtp.parent_proxy_host value from the UTM
func GetSmtpParentProxyHost(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.parent_proxy_host", &val, options...)
	return
}

// UpdateSmtpParentProxyHost PUTs the smtp.parent_proxy_host value to the UTM
func UpdateSmtpParentProxyHost(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.parent_proxy_host", val, options...)
}

// GetSmtpParentProxyPort gets the smtp.parent_proxy_port value from the UTM
func GetSmtpParentProxyPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/smtp.parent_proxy_port", &val, options...)
	return
}

// UpdateSmtpParentProxyPort PUTs the smtp.parent_proxy_port value to the UTM
func UpdateSmtpParentProxyPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.parent_proxy_port", val, options...)
}

// GetSmtpParentProxyStatus gets the smtp.parent_proxy_status value from the UTM
func GetSmtpParentProxyStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.parent_proxy_status", &val, options...)
	return
}

// UpdateSmtpParentProxyStatus PUTs the smtp.parent_proxy_status value to the UTM
func UpdateSmtpParentProxyStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.parent_proxy_status", val, options...)
}

// GetSmtpPostmaster gets the smtp.postmaster value from the UTM
func GetSmtpPostmaster(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.postmaster", &val, options...)
	return
}

// UpdateSmtpPostmaster PUTs the smtp.postmaster value to the UTM
func UpdateSmtpPostmaster(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.postmaster", val, options...)
}

// GetSmtpProfiles gets the smtp.profiles value from the UTM
func GetSmtpProfiles(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/smtp.profiles", &val, options...)
	return
}

// UpdateSmtpProfiles PUTs the smtp.profiles value to the UTM
func UpdateSmtpProfiles(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.profiles", val, options...)
}

// GetSmtpRecipientsMax gets the smtp.recipients_max value from the UTM
func GetSmtpRecipientsMax(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/smtp.recipients_max", &val, options...)
	return
}

// UpdateSmtpRecipientsMax PUTs the smtp.recipients_max value to the UTM
func UpdateSmtpRecipientsMax(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.recipients_max", val, options...)
}

// GetSmtpRelays gets the smtp.relays value from the UTM
func GetSmtpRelays(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/smtp.relays", &val, options...)
	return
}

// UpdateSmtpRelays PUTs the smtp.relays value to the UTM
func UpdateSmtpRelays(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.relays", val, options...)
}

// GetSmtpSasiToken gets the smtp.sasi_token value from the UTM
func GetSmtpSasiToken(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.sasi_token", &val, options...)
	return
}

// UpdateSmtpSasiToken PUTs the smtp.sasi_token value to the UTM
func UpdateSmtpSasiToken(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.sasi_token", val, options...)
}

// GetSmtpScanOutgoingEmails gets the smtp.scan_outgoing_emails value from the UTM
func GetSmtpScanOutgoingEmails(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.scan_outgoing_emails", &val, options...)
	return
}

// UpdateSmtpScanOutgoingEmails PUTs the smtp.scan_outgoing_emails value to the UTM
func UpdateSmtpScanOutgoingEmails(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.scan_outgoing_emails", val, options...)
}

// GetSmtpScannerPoolMaxPool gets the smtp.scanner_pool.max_pool value from the UTM
func GetSmtpScannerPoolMaxPool(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/smtp.scanner_pool.max_pool", &val, options...)
	return
}

// UpdateSmtpScannerPoolMaxPool PUTs the smtp.scanner_pool.max_pool value to the UTM
func UpdateSmtpScannerPoolMaxPool(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.scanner_pool.max_pool", val, options...)
}

// GetSmtpScannerPoolThresholds gets the smtp.scanner_pool.thresholds value from the UTM
func GetSmtpScannerPoolThresholds(client sophos.ClientInterface, options ...sophos.Option) (val []int64, err error) {
	err = get(client, "/api/nodes/smtp.scanner_pool.thresholds", &val, options...)
	return
}

// UpdateSmtpScannerPoolThresholds PUTs the smtp.scanner_pool.thresholds value to the UTM
func UpdateSmtpScannerPoolThresholds(client sophos.ClientInterface, val []int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.scanner_pool.thresholds", val, options...)
}

// GetSmtpScannerTimeout gets the smtp.scanner_timeout value from the UTM
func GetSmtpScannerTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/smtp.scanner_timeout", &val, options...)
	return
}

// UpdateSmtpScannerTimeout PUTs the smtp.scanner_timeout value to the UTM
func UpdateSmtpScannerTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.scanner_timeout", val, options...)
}

// GetSmtpSmarthostAuth gets the smtp.smarthost_auth value from the UTM
func GetSmtpSmarthostAuth(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.smarthost_auth", &val, options...)
	return
}

// UpdateSmtpSmarthostAuth PUTs the smtp.smarthost_auth value to the UTM
func UpdateSmtpSmarthostAuth(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.smarthost_auth", val, options...)
}

// GetSmtpSmarthostHostname gets the smtp.smarthost_hostname value from the UTM
func GetSmtpSmarthostHostname(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.smarthost_hostname", &val, options...)
	return
}

// UpdateSmtpSmarthostHostname PUTs the smtp.smarthost_hostname value to the UTM
func UpdateSmtpSmarthostHostname(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.smarthost_hostname", val, options...)
}

// GetSmtpSmarthostPass gets the smtp.smarthost_pass value from the UTM
func GetSmtpSmarthostPass(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.smarthost_pass", &val, options...)
	return
}

// UpdateSmtpSmarthostPass PUTs the smtp.smarthost_pass value to the UTM
func UpdateSmtpSmarthostPass(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.smarthost_pass", val, options...)
}

// GetSmtpSmarthostPort gets the smtp.smarthost_port value from the UTM
func GetSmtpSmarthostPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/smtp.smarthost_port", &val, options...)
	return
}

// UpdateSmtpSmarthostPort PUTs the smtp.smarthost_port value to the UTM
func UpdateSmtpSmarthostPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.smarthost_port", val, options...)
}

// GetSmtpSmarthostStatus gets the smtp.smarthost_status value from the UTM
func GetSmtpSmarthostStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.smarthost_status", &val, options...)
	return
}

// UpdateSmtpSmarthostStatus PUTs the smtp.smarthost_status value to the UTM
func UpdateSmtpSmarthostStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.smarthost_status", val, options...)
}

// GetSmtpSmarthostUser gets the smtp.smarthost_user value from the UTM
func GetSmtpSmarthostUser(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.smarthost_user", &val, options...)
	return
}

// UpdateSmtpSmarthostUser PUTs the smtp.smarthost_user value to the UTM
func UpdateSmtpSmarthostUser(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.smarthost_user", val, options...)
}

// GetSmtpSmtpAcceptMax gets the smtp.smtp_accept_max value from the UTM
func GetSmtpSmtpAcceptMax(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/smtp.smtp_accept_max", &val, options...)
	return
}

// UpdateSmtpSmtpAcceptMax PUTs the smtp.smtp_accept_max value to the UTM
func UpdateSmtpSmtpAcceptMax(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.smtp_accept_max", val, options...)
}

// GetSmtpSmtpAcceptPerConnectionMax gets the smtp.smtp_accept_per_connection_max value from the UTM
func GetSmtpSmtpAcceptPerConnectionMax(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/smtp.smtp_accept_per_connection_max", &val, options...)
	return
}

// UpdateSmtpSmtpAcceptPerConnectionMax PUTs the smtp.smtp_accept_per_connection_max value to the UTM
func UpdateSmtpSmtpAcceptPerConnectionMax(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.smtp_accept_per_connection_max", val, options...)
}

// GetSmtpSmtpAcceptPerHostMax gets the smtp.smtp_accept_per_host_max value from the UTM
func GetSmtpSmtpAcceptPerHostMax(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/smtp.smtp_accept_per_host_max", &val, options...)
	return
}

// UpdateSmtpSmtpAcceptPerHostMax PUTs the smtp.smtp_accept_per_host_max value to the UTM
func UpdateSmtpSmtpAcceptPerHostMax(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.smtp_accept_per_host_max", val, options...)
}

// GetSmtpSmtpAllowedInterfaces gets the smtp.smtp_allowed_interfaces value from the UTM
func GetSmtpSmtpAllowedInterfaces(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/smtp.smtp_allowed_interfaces", &val, options...)
	return
}

// UpdateSmtpSmtpAllowedInterfaces PUTs the smtp.smtp_allowed_interfaces value to the UTM
func UpdateSmtpSmtpAllowedInterfaces(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.smtp_allowed_interfaces", val, options...)
}

// GetSmtpSmtpListenOnItf gets the smtp.smtp_listen_on_itf value from the UTM
func GetSmtpSmtpListenOnItf(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.smtp_listen_on_itf", &val, options...)
	return
}

// UpdateSmtpSmtpListenOnItf PUTs the smtp.smtp_listen_on_itf value to the UTM
func UpdateSmtpSmtpListenOnItf(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.smtp_listen_on_itf", val, options...)
}

// GetSmtpStatus gets the smtp.status value from the UTM
func GetSmtpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.status", &val, options...)
	return
}

// UpdateSmtpStatus PUTs the smtp.status value to the UTM
func UpdateSmtpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.status", val, options...)
}

// GetSmtpTlsAvoid gets the smtp.tls_avoid value from the UTM
func GetSmtpTlsAvoid(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/smtp.tls_avoid", &val, options...)
	return
}

// UpdateSmtpTlsAvoid PUTs the smtp.tls_avoid value to the UTM
func UpdateSmtpTlsAvoid(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.tls_avoid", val, options...)
}

// GetSmtpTlsCert gets the smtp.tls_cert value from the UTM
func GetSmtpTlsCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.tls_cert", &val, options...)
	return
}

// UpdateSmtpTlsCert PUTs the smtp.tls_cert value to the UTM
func UpdateSmtpTlsCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.tls_cert", val, options...)
}

// GetSmtpTlsRequire gets the smtp.tls_require value from the UTM
func GetSmtpTlsRequire(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/smtp.tls_require", &val, options...)
	return
}

// UpdateSmtpTlsRequire PUTs the smtp.tls_require value to the UTM
func UpdateSmtpTlsRequire(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.tls_require", val, options...)
}

// GetSmtpTlsRequireSenderDomains gets the smtp.tls_require_sender_domains value from the UTM
func GetSmtpTlsRequireSenderDomains(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/smtp.tls_require_sender_domains", &val, options...)
	return
}

// UpdateSmtpTlsRequireSenderDomains PUTs the smtp.tls_require_sender_domains value to the UTM
func UpdateSmtpTlsRequireSenderDomains(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.tls_require_sender_domains", val, options...)
}

// GetSmtpTlsVersion gets the smtp.tls_version value from the UTM
func GetSmtpTlsVersion(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.tls_version", &val, options...)
	return
}

// UpdateSmtpTlsVersion PUTs the smtp.tls_version value to the UTM
func UpdateSmtpTlsVersion(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.tls_version", val, options...)
}

// GetSmtpTransparent gets the smtp.transparent value from the UTM
func GetSmtpTransparent(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.transparent", &val, options...)
	return
}

// UpdateSmtpTransparent PUTs the smtp.transparent value to the UTM
func UpdateSmtpTransparent(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.transparent", val, options...)
}

// GetSmtpTransparentSkip gets the smtp.transparent_skip value from the UTM
func GetSmtpTransparentSkip(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/smtp.transparent_skip", &val, options...)
	return
}

// UpdateSmtpTransparentSkip PUTs the smtp.transparent_skip value to the UTM
func UpdateSmtpTransparentSkip(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.transparent_skip", val, options...)
}

// GetSmtpTransparentSkipAutoPf gets the smtp.transparent_skip_auto_pf value from the UTM
func GetSmtpTransparentSkipAutoPf(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.transparent_skip_auto_pf", &val, options...)
	return
}

// UpdateSmtpTransparentSkipAutoPf PUTs the smtp.transparent_skip_auto_pf value to the UTM
func UpdateSmtpTransparentSkipAutoPf(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.transparent_skip_auto_pf", val, options...)
}

// GetSmtpTransparentSmtps gets the smtp.transparent_smtps value from the UTM
func GetSmtpTransparentSmtps(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.transparent_smtps", &val, options...)
	return
}

// UpdateSmtpTransparentSmtps PUTs the smtp.transparent_smtps value to the UTM
func UpdateSmtpTransparentSmtps(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.transparent_smtps", val, options...)
}

// GetSmtpTransparentStartTls gets the smtp.transparent_start_tls value from the UTM
func GetSmtpTransparentStartTls(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.transparent_start_tls", &val, options...)
	return
}

// UpdateSmtpTransparentStartTls PUTs the smtp.transparent_start_tls value to the UTM
func UpdateSmtpTransparentStartTls(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.transparent_start_tls", val, options...)
}

// GetSmtpUpstreamHosts gets the smtp.upstream_hosts value from the UTM
func GetSmtpUpstreamHosts(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/smtp.upstream_hosts", &val, options...)
	return
}

// UpdateSmtpUpstreamHosts PUTs the smtp.upstream_hosts value to the UTM
func UpdateSmtpUpstreamHosts(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.upstream_hosts", val, options...)
}

// GetSmtpUpstreamHostsOnly gets the smtp.upstream_hosts_only value from the UTM
func GetSmtpUpstreamHostsOnly(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/smtp.upstream_hosts_only", &val, options...)
	return
}

// UpdateSmtpUpstreamHostsOnly PUTs the smtp.upstream_hosts_only value to the UTM
func UpdateSmtpUpstreamHostsOnly(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.upstream_hosts_only", val, options...)
}

// GetSmtpUseClientMime gets the smtp.use_client_mime value from the UTM
func GetSmtpUseClientMime(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/smtp.use_client_mime", &val, options...)
	return
}

// UpdateSmtpUseClientMime PUTs the smtp.use_client_mime value to the UTM
func UpdateSmtpUseClientMime(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/smtp.use_client_mime", val, options...)
}

// GetSnmpAllowedNetworks gets the snmp.allowed_networks value from the UTM
func GetSnmpAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/snmp.allowed_networks", &val, options...)
	return
}

// UpdateSnmpAllowedNetworks PUTs the snmp.allowed_networks value to the UTM
func UpdateSnmpAllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.allowed_networks", val, options...)
}

// GetSnmpAuthPassword gets the snmp.auth_password value from the UTM
func GetSnmpAuthPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/snmp.auth_password", &val, options...)
	return
}

// UpdateSnmpAuthPassword PUTs the snmp.auth_password value to the UTM
func UpdateSnmpAuthPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.auth_password", val, options...)
}

// GetSnmpAuthType gets the snmp.auth_type value from the UTM
func GetSnmpAuthType(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/snmp.auth_type", &val, options...)
	return
}

// UpdateSnmpAuthType PUTs the snmp.auth_type value to the UTM
func UpdateSnmpAuthType(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.auth_type", val, options...)
}

// GetSnmpCommunity gets the snmp.community value from the UTM
func GetSnmpCommunity(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/snmp.community", &val, options...)
	return
}

// UpdateSnmpCommunity PUTs the snmp.community value to the UTM
func UpdateSnmpCommunity(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.community", val, options...)
}

// GetSnmpDeviceAdmin gets the snmp.device_admin value from the UTM
func GetSnmpDeviceAdmin(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/snmp.device_admin", &val, options...)
	return
}

// UpdateSnmpDeviceAdmin PUTs the snmp.device_admin value to the UTM
func UpdateSnmpDeviceAdmin(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.device_admin", val, options...)
}

// GetSnmpDeviceLocation gets the snmp.device_location value from the UTM
func GetSnmpDeviceLocation(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/snmp.device_location", &val, options...)
	return
}

// UpdateSnmpDeviceLocation PUTs the snmp.device_location value to the UTM
func UpdateSnmpDeviceLocation(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.device_location", val, options...)
}

// GetSnmpDeviceName gets the snmp.device_name value from the UTM
func GetSnmpDeviceName(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/snmp.device_name", &val, options...)
	return
}

// UpdateSnmpDeviceName PUTs the snmp.device_name value to the UTM
func UpdateSnmpDeviceName(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.device_name", val, options...)
}

// GetSnmpEncryptPassword gets the snmp.encrypt_password value from the UTM
func GetSnmpEncryptPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/snmp.encrypt_password", &val, options...)
	return
}

// UpdateSnmpEncryptPassword PUTs the snmp.encrypt_password value to the UTM
func UpdateSnmpEncryptPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.encrypt_password", val, options...)
}

// GetSnmpEncryptType gets the snmp.encrypt_type value from the UTM
func GetSnmpEncryptType(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/snmp.encrypt_type", &val, options...)
	return
}

// UpdateSnmpEncryptType PUTs the snmp.encrypt_type value to the UTM
func UpdateSnmpEncryptType(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.encrypt_type", val, options...)
}

// GetSnmpStatus gets the snmp.status value from the UTM
func GetSnmpStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/snmp.status", &val, options...)
	return
}

// UpdateSnmpStatus PUTs the snmp.status value to the UTM
func UpdateSnmpStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.status", val, options...)
}

// GetSnmpTraps gets the snmp.traps value from the UTM
func GetSnmpTraps(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/snmp.traps", &val, options...)
	return
}

// UpdateSnmpTraps PUTs the snmp.traps value to the UTM
func UpdateSnmpTraps(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.traps", val, options...)
}

// GetSnmpTrapsUseOldOids gets the snmp.traps_use_old_oids value from the UTM
func GetSnmpTrapsUseOldOids(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/snmp.traps_use_old_oids", &val, options...)
	return
}

// UpdateSnmpTrapsUseOldOids PUTs the snmp.traps_use_old_oids value to the UTM
func UpdateSnmpTrapsUseOldOids(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.traps_use_old_oids", val, options...)
}

// GetSnmpUsername gets the snmp.username value from the UTM
func GetSnmpUsername(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/snmp.username", &val, options...)
	return
}

// UpdateSnmpUsername PUTs the snmp.username value to the UTM
func UpdateSnmpUsername(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.username", val, options...)
}

// GetSnmpVersion gets the snmp.version value from the UTM
func GetSnmpVersion(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/snmp.version", &val, options...)
	return
}

// UpdateSnmpVersion PUTs the snmp.version value to the UTM
func UpdateSnmpVersion(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/snmp.version", val, options...)
}

// GetSocksAaa gets the socks.aaa value from the UTM
func GetSocksAaa(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/socks.aaa", &val, options...)
	return
}

// UpdateSocksAaa PUTs the socks.aaa value to the UTM
func UpdateSocksAaa(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/socks.aaa", val, options...)
}

// GetSocksAllowedNetworks gets the socks.allowed_networks value from the UTM
func GetSocksAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/socks.allowed_networks", &val, options...)
	return
}

// UpdateSocksAllowedNetworks PUTs the socks.allowed_networks value to the UTM
func UpdateSocksAllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/socks.allowed_networks", val, options...)
}

// GetSocksAuthentication gets the socks.authentication value from the UTM
func GetSocksAuthentication(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/socks.authentication", &val, options...)
	return
}

// UpdateSocksAuthentication PUTs the socks.authentication value to the UTM
func UpdateSocksAuthentication(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/socks.authentication", val, options...)
}

// GetSocksStatus gets the socks.status value from the UTM
func GetSocksStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/socks.status", &val, options...)
	return
}

// UpdateSocksStatus PUTs the socks.status value to the UTM
func UpdateSocksStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/socks.status", val, options...)
}

// GetSpxGlobalErrorNotificationTarget gets the spx.global.error_notification_target value from the UTM
func GetSpxGlobalErrorNotificationTarget(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/spx.global.error_notification_target", &val, options...)
	return
}

// UpdateSpxGlobalErrorNotificationTarget PUTs the spx.global.error_notification_target value to the UTM
func UpdateSpxGlobalErrorNotificationTarget(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.error_notification_target", val, options...)
}

// GetSpxGlobalExpirySettingsAllowSecureReplyDays gets the spx.global.expiry_settings.allow_secure_reply_days value from the UTM
func GetSpxGlobalExpirySettingsAllowSecureReplyDays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/spx.global.expiry_settings.allow_secure_reply_days", &val, options...)
	return
}

// UpdateSpxGlobalExpirySettingsAllowSecureReplyDays PUTs the spx.global.expiry_settings.allow_secure_reply_days value to the UTM
func UpdateSpxGlobalExpirySettingsAllowSecureReplyDays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.expiry_settings.allow_secure_reply_days", val, options...)
}

// GetSpxGlobalExpirySettingsKeepDelayedMailDays gets the spx.global.expiry_settings.keep_delayed_mail_days value from the UTM
func GetSpxGlobalExpirySettingsKeepDelayedMailDays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/spx.global.expiry_settings.keep_delayed_mail_days", &val, options...)
	return
}

// UpdateSpxGlobalExpirySettingsKeepDelayedMailDays PUTs the spx.global.expiry_settings.keep_delayed_mail_days value to the UTM
func UpdateSpxGlobalExpirySettingsKeepDelayedMailDays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.expiry_settings.keep_delayed_mail_days", val, options...)
}

// GetSpxGlobalExpirySettingsKeepUnusedPwdDays gets the spx.global.expiry_settings.keep_unused_pwd_days value from the UTM
func GetSpxGlobalExpirySettingsKeepUnusedPwdDays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/spx.global.expiry_settings.keep_unused_pwd_days", &val, options...)
	return
}

// UpdateSpxGlobalExpirySettingsKeepUnusedPwdDays PUTs the spx.global.expiry_settings.keep_unused_pwd_days value to the UTM
func UpdateSpxGlobalExpirySettingsKeepUnusedPwdDays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.expiry_settings.keep_unused_pwd_days", val, options...)
}

// GetSpxGlobalExpirySettingsRegistrationReminderDays gets the spx.global.expiry_settings.registration_reminder_days value from the UTM
func GetSpxGlobalExpirySettingsRegistrationReminderDays(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/spx.global.expiry_settings.registration_reminder_days", &val, options...)
	return
}

// UpdateSpxGlobalExpirySettingsRegistrationReminderDays PUTs the spx.global.expiry_settings.registration_reminder_days value to the UTM
func UpdateSpxGlobalExpirySettingsRegistrationReminderDays(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.expiry_settings.registration_reminder_days", val, options...)
}

// GetSpxGlobalPasswordStrengthMinLength gets the spx.global.password_strength.min_length value from the UTM
func GetSpxGlobalPasswordStrengthMinLength(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/spx.global.password_strength.min_length", &val, options...)
	return
}

// UpdateSpxGlobalPasswordStrengthMinLength PUTs the spx.global.password_strength.min_length value to the UTM
func UpdateSpxGlobalPasswordStrengthMinLength(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.password_strength.min_length", val, options...)
}

// GetSpxGlobalPasswordStrengthRequireSpecChars gets the spx.global.password_strength.require_spec_chars value from the UTM
func GetSpxGlobalPasswordStrengthRequireSpecChars(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/spx.global.password_strength.require_spec_chars", &val, options...)
	return
}

// UpdateSpxGlobalPasswordStrengthRequireSpecChars PUTs the spx.global.password_strength.require_spec_chars value to the UTM
func UpdateSpxGlobalPasswordStrengthRequireSpecChars(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.password_strength.require_spec_chars", val, options...)
}

// GetSpxGlobalPortalSettingsAllowedNetworks gets the spx.global.portal_settings.allowed_networks value from the UTM
func GetSpxGlobalPortalSettingsAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/spx.global.portal_settings.allowed_networks", &val, options...)
	return
}

// UpdateSpxGlobalPortalSettingsAllowedNetworks PUTs the spx.global.portal_settings.allowed_networks value to the UTM
func UpdateSpxGlobalPortalSettingsAllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.portal_settings.allowed_networks", val, options...)
}

// GetSpxGlobalPortalSettingsHostname gets the spx.global.portal_settings.hostname value from the UTM
func GetSpxGlobalPortalSettingsHostname(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/spx.global.portal_settings.hostname", &val, options...)
	return
}

// UpdateSpxGlobalPortalSettingsHostname PUTs the spx.global.portal_settings.hostname value to the UTM
func UpdateSpxGlobalPortalSettingsHostname(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.portal_settings.hostname", val, options...)
}

// GetSpxGlobalPortalSettingsInterfaceAddress gets the spx.global.portal_settings.interface_address value from the UTM
func GetSpxGlobalPortalSettingsInterfaceAddress(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/spx.global.portal_settings.interface_address", &val, options...)
	return
}

// UpdateSpxGlobalPortalSettingsInterfaceAddress PUTs the spx.global.portal_settings.interface_address value to the UTM
func UpdateSpxGlobalPortalSettingsInterfaceAddress(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.portal_settings.interface_address", val, options...)
}

// GetSpxGlobalPortalSettingsPort gets the spx.global.portal_settings.port value from the UTM
func GetSpxGlobalPortalSettingsPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/spx.global.portal_settings.port", &val, options...)
	return
}

// UpdateSpxGlobalPortalSettingsPort PUTs the spx.global.portal_settings.port value to the UTM
func UpdateSpxGlobalPortalSettingsPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.portal_settings.port", val, options...)
}

// GetSpxGlobalPoweredByLogo gets the spx.global.powered_by_logo value from the UTM
func GetSpxGlobalPoweredByLogo(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/spx.global.powered_by_logo", &val, options...)
	return
}

// UpdateSpxGlobalPoweredByLogo PUTs the spx.global.powered_by_logo value to the UTM
func UpdateSpxGlobalPoweredByLogo(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.powered_by_logo", val, options...)
}

// GetSpxGlobalSpxPriority gets the spx.global.spx_priority value from the UTM
func GetSpxGlobalSpxPriority(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/spx.global.spx_priority", &val, options...)
	return
}

// UpdateSpxGlobalSpxPriority PUTs the spx.global.spx_priority value to the UTM
func UpdateSpxGlobalSpxPriority(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.spx_priority", val, options...)
}

// GetSpxGlobalStatus gets the spx.global.status value from the UTM
func GetSpxGlobalStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/spx.global.status", &val, options...)
	return
}

// UpdateSpxGlobalStatus PUTs the spx.global.status value to the UTM
func UpdateSpxGlobalStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.global.status", val, options...)
}

// GetSpxSpxAuthPassword gets the spx.spx_auth.password value from the UTM
func GetSpxSpxAuthPassword(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/spx.spx_auth.password", &val, options...)
	return
}

// UpdateSpxSpxAuthPassword PUTs the spx.spx_auth.password value to the UTM
func UpdateSpxSpxAuthPassword(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.spx_auth.password", val, options...)
}

// GetSpxSpxAuthPort gets the spx.spx_auth.port value from the UTM
func GetSpxSpxAuthPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/spx.spx_auth.port", &val, options...)
	return
}

// UpdateSpxSpxAuthPort PUTs the spx.spx_auth.port value to the UTM
func UpdateSpxSpxAuthPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.spx_auth.port", val, options...)
}

// GetSpxSpxAuthServer gets the spx.spx_auth.server value from the UTM
func GetSpxSpxAuthServer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/spx.spx_auth.server", &val, options...)
	return
}

// UpdateSpxSpxAuthServer PUTs the spx.spx_auth.server value to the UTM
func UpdateSpxSpxAuthServer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.spx_auth.server", val, options...)
}

// GetSpxSpxAuthUrl gets the spx.spx_auth.url value from the UTM
func GetSpxSpxAuthUrl(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/spx.spx_auth.url", &val, options...)
	return
}

// UpdateSpxSpxAuthUrl PUTs the spx.spx_auth.url value to the UTM
func UpdateSpxSpxAuthUrl(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.spx_auth.url", val, options...)
}

// GetSpxTemplates gets the spx.templates value from the UTM
func GetSpxTemplates(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/spx.templates", &val, options...)
	return
}

// UpdateSpxTemplates PUTs the spx.templates value to the UTM
func UpdateSpxTemplates(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/spx.templates", val, options...)
}

// GetSshAllowedNetworks gets the ssh.allowed_networks value from the UTM
func GetSshAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/ssh.allowed_networks", &val, options...)
	return
}

// UpdateSshAllowedNetworks PUTs the ssh.allowed_networks value to the UTM
func UpdateSshAllowedNetworks(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssh.allowed_networks", val, options...)
}

// GetSshLoginKeys gets the ssh.login_keys value from the UTM
func GetSshLoginKeys(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/ssh.login_keys", &val, options...)
	return
}

// UpdateSshLoginKeys PUTs the ssh.login_keys value to the UTM
func UpdateSshLoginKeys(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssh.login_keys", val, options...)
}

// GetSshPasswordAuth gets the ssh.password_auth value from the UTM
func GetSshPasswordAuth(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ssh.password_auth", &val, options...)
	return
}

// UpdateSshPasswordAuth PUTs the ssh.password_auth value to the UTM
func UpdateSshPasswordAuth(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssh.password_auth", val, options...)
}

// GetSshPort gets the ssh.port value from the UTM
func GetSshPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ssh.port", &val, options...)
	return
}

// UpdateSshPort PUTs the ssh.port value to the UTM
func UpdateSshPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssh.port", val, options...)
}

// GetSshPubkeyAuth gets the ssh.pubkey_auth value from the UTM
func GetSshPubkeyAuth(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ssh.pubkey_auth", &val, options...)
	return
}

// UpdateSshPubkeyAuth PUTs the ssh.pubkey_auth value to the UTM
func UpdateSshPubkeyAuth(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssh.pubkey_auth", val, options...)
}

// GetSshRootKeys gets the ssh.root_keys value from the UTM
func GetSshRootKeys(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/ssh.root_keys", &val, options...)
	return
}

// UpdateSshRootKeys PUTs the ssh.root_keys value to the UTM
func UpdateSshRootKeys(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssh.root_keys", val, options...)
}

// GetSshRootLogin gets the ssh.root_login value from the UTM
func GetSshRootLogin(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ssh.root_login", &val, options...)
	return
}

// UpdateSshRootLogin PUTs the ssh.root_login value to the UTM
func UpdateSshRootLogin(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssh.root_login", val, options...)
}

// GetSshStatus gets the ssh.status value from the UTM
func GetSshStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ssh.status", &val, options...)
	return
}

// UpdateSshStatus PUTs the ssh.status value to the UTM
func UpdateSshStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssh.status", val, options...)
}

// GetSslVpnAuthenticationAlgorithm gets the ssl_vpn.authentication_algorithm value from the UTM
func GetSslVpnAuthenticationAlgorithm(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ssl_vpn.authentication_algorithm", &val, options...)
	return
}

// UpdateSslVpnAuthenticationAlgorithm PUTs the ssl_vpn.authentication_algorithm value to the UTM
func UpdateSslVpnAuthenticationAlgorithm(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.authentication_algorithm", val, options...)
}

// GetSslVpnCertificate gets the ssl_vpn.certificate value from the UTM
func GetSslVpnCertificate(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ssl_vpn.certificate", &val, options...)
	return
}

// UpdateSslVpnCertificate PUTs the ssl_vpn.certificate value to the UTM
func UpdateSslVpnCertificate(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.certificate", val, options...)
}

// GetSslVpnCompression gets the ssl_vpn.compression value from the UTM
func GetSslVpnCompression(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ssl_vpn.compression", &val, options...)
	return
}

// UpdateSslVpnCompression PUTs the ssl_vpn.compression value to the UTM
func UpdateSslVpnCompression(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.compression", val, options...)
}

// GetSslVpnDatachannelKeyLifetime gets the ssl_vpn.datachannel_key_lifetime value from the UTM
func GetSslVpnDatachannelKeyLifetime(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ssl_vpn.datachannel_key_lifetime", &val, options...)
	return
}

// UpdateSslVpnDatachannelKeyLifetime PUTs the ssl_vpn.datachannel_key_lifetime value to the UTM
func UpdateSslVpnDatachannelKeyLifetime(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.datachannel_key_lifetime", val, options...)
}

// GetSslVpnDebug gets the ssl_vpn.debug value from the UTM
func GetSslVpnDebug(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ssl_vpn.debug", &val, options...)
	return
}

// UpdateSslVpnDebug PUTs the ssl_vpn.debug value to the UTM
func UpdateSslVpnDebug(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.debug", val, options...)
}

// GetSslVpnDhKeySize gets the ssl_vpn.dh_key_size value from the UTM
func GetSslVpnDhKeySize(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ssl_vpn.dh_key_size", &val, options...)
	return
}

// UpdateSslVpnDhKeySize PUTs the ssl_vpn.dh_key_size value to the UTM
func UpdateSslVpnDhKeySize(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.dh_key_size", val, options...)
}

// GetSslVpnDuplicateCn gets the ssl_vpn.duplicate_cn value from the UTM
func GetSslVpnDuplicateCn(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ssl_vpn.duplicate_cn", &val, options...)
	return
}

// UpdateSslVpnDuplicateCn PUTs the ssl_vpn.duplicate_cn value to the UTM
func UpdateSslVpnDuplicateCn(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.duplicate_cn", val, options...)
}

// GetSslVpnEncryptionAlgorithm gets the ssl_vpn.encryption_algorithm value from the UTM
func GetSslVpnEncryptionAlgorithm(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ssl_vpn.encryption_algorithm", &val, options...)
	return
}

// UpdateSslVpnEncryptionAlgorithm PUTs the ssl_vpn.encryption_algorithm value to the UTM
func UpdateSslVpnEncryptionAlgorithm(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.encryption_algorithm", val, options...)
}

// GetSslVpnFallbackDisable gets the ssl_vpn.fallback_disable value from the UTM
func GetSslVpnFallbackDisable(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ssl_vpn.fallback_disable", &val, options...)
	return
}

// UpdateSslVpnFallbackDisable PUTs the ssl_vpn.fallback_disable value to the UTM
func UpdateSslVpnFallbackDisable(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.fallback_disable", val, options...)
}

// GetSslVpnHostname gets the ssl_vpn.hostname value from the UTM
func GetSslVpnHostname(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ssl_vpn.hostname", &val, options...)
	return
}

// UpdateSslVpnHostname PUTs the ssl_vpn.hostname value to the UTM
func UpdateSslVpnHostname(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.hostname", val, options...)
}

// GetSslVpnInterface gets the ssl_vpn.interface value from the UTM
func GetSslVpnInterface(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ssl_vpn.interface", &val, options...)
	return
}

// UpdateSslVpnInterface PUTs the ssl_vpn.interface value to the UTM
func UpdateSslVpnInterface(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.interface", val, options...)
}

// GetSslVpnInterfaceAddress gets the ssl_vpn.interface_address value from the UTM
func GetSslVpnInterfaceAddress(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ssl_vpn.interface_address", &val, options...)
	return
}

// UpdateSslVpnInterfaceAddress PUTs the ssl_vpn.interface_address value to the UTM
func UpdateSslVpnInterfaceAddress(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.interface_address", val, options...)
}

// GetSslVpnIpAssignmentPool gets the ssl_vpn.ip_assignment_pool value from the UTM
func GetSslVpnIpAssignmentPool(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ssl_vpn.ip_assignment_pool", &val, options...)
	return
}

// UpdateSslVpnIpAssignmentPool PUTs the ssl_vpn.ip_assignment_pool value to the UTM
func UpdateSslVpnIpAssignmentPool(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.ip_assignment_pool", val, options...)
}

// GetSslVpnPort gets the ssl_vpn.port value from the UTM
func GetSslVpnPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/ssl_vpn.port", &val, options...)
	return
}

// UpdateSslVpnPort PUTs the ssl_vpn.port value to the UTM
func UpdateSslVpnPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.port", val, options...)
}

// GetSslVpnProtocol gets the ssl_vpn.protocol value from the UTM
func GetSslVpnProtocol(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/ssl_vpn.protocol", &val, options...)
	return
}

// UpdateSslVpnProtocol PUTs the ssl_vpn.protocol value to the UTM
func UpdateSslVpnProtocol(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.protocol", val, options...)
}

// GetSslVpnUserAuthOptional gets the ssl_vpn.user_auth_optional value from the UTM
func GetSslVpnUserAuthOptional(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/ssl_vpn.user_auth_optional", &val, options...)
	return
}

// UpdateSslVpnUserAuthOptional PUTs the ssl_vpn.user_auth_optional value to the UTM
func UpdateSslVpnUserAuthOptional(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/ssl_vpn.user_auth_optional", val, options...)
}

// GetSupportAccessAccessId gets the support_access.access_id value from the UTM
func GetSupportAccessAccessId(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/support_access.access_id", &val, options...)
	return
}

// UpdateSupportAccessAccessId PUTs the support_access.access_id value to the UTM
func UpdateSupportAccessAccessId(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/support_access.access_id", val, options...)
}

// GetSupportAccessApuServer gets the support_access.apu_server value from the UTM
func GetSupportAccessApuServer(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/support_access.apu_server", &val, options...)
	return
}

// UpdateSupportAccessApuServer PUTs the support_access.apu_server value to the UTM
func UpdateSupportAccessApuServer(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/support_access.apu_server", val, options...)
}

// GetSupportAccessLifetimeDuration gets the support_access.lifetime_duration value from the UTM
func GetSupportAccessLifetimeDuration(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/support_access.lifetime_duration", &val, options...)
	return
}

// UpdateSupportAccessLifetimeDuration PUTs the support_access.lifetime_duration value to the UTM
func UpdateSupportAccessLifetimeDuration(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/support_access.lifetime_duration", val, options...)
}

// GetSupportAccessLifetimeEnd gets the support_access.lifetime_end value from the UTM
func GetSupportAccessLifetimeEnd(client sophos.ClientInterface, options ...sophos.Option) (val map[string]interface{}, err error) {
	err = get(client, "/api/nodes/support_access.lifetime_end", &val, options...)
	return
}

// UpdateSupportAccessLifetimeEnd PUTs the support_access.lifetime_end value to the UTM
func UpdateSupportAccessLifetimeEnd(client sophos.ClientInterface, val map[string]interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/support_access.lifetime_end", val, options...)
}

// GetSupportAccessSshKeys gets the support_access.ssh_keys value from the UTM
func GetSupportAccessSshKeys(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/support_access.ssh_keys", &val, options...)
	return
}

// UpdateSupportAccessSshKeys PUTs the support_access.ssh_keys value to the UTM
func UpdateSupportAccessSshKeys(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/support_access.ssh_keys", val, options...)
}

// GetSupportAccessStatus gets the support_access.status value from the UTM
func GetSupportAccessStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/support_access.status", &val, options...)
	return
}

// UpdateSupportAccessStatus PUTs the support_access.status value to the UTM
func UpdateSupportAccessStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/support_access.status", val, options...)
}

// GetU2DcacheAllowedNetworks gets the u2dcache.allowed_networks value from the UTM
func GetU2DcacheAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/u2dcache.allowed_networks", &val, options...)
	return
}

// UpdateU2DcacheAllowedNetworks PUTs the u2dcache.allowed_networks value to the UTM
func UpdateU2DcacheAllowedNetworks(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/u2dcache.allowed_networks", val, options...)
}

// GetU2DcachePort gets the u2dcache.port value from the UTM
func GetU2DcachePort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/u2dcache.port", &val, options...)
	return
}

// UpdateU2DcachePort PUTs the u2dcache.port value to the UTM
func UpdateU2DcachePort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/u2dcache.port", val, options...)
}

// GetU2DcacheStatus gets the u2dcache.status value from the UTM
func GetU2DcacheStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/u2dcache.status", &val, options...)
	return
}

// UpdateU2DcacheStatus PUTs the u2dcache.status value to the UTM
func UpdateU2DcacheStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/u2dcache.status", val, options...)
}

// GetUp2DateCacheHost gets the up2date.cache_host value from the UTM
func GetUp2DateCacheHost(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/up2date.cache_host", &val, options...)
	return
}

// UpdateUp2DateCacheHost PUTs the up2date.cache_host value to the UTM
func UpdateUp2DateCacheHost(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.cache_host", val, options...)
}

// GetUp2DateCachePort gets the up2date.cache_port value from the UTM
func GetUp2DateCachePort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/up2date.cache_port", &val, options...)
	return
}

// UpdateUp2DateCachePort PUTs the up2date.cache_port value to the UTM
func UpdateUp2DateCachePort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.cache_port", val, options...)
}

// GetUp2DateCacheStatus gets the up2date.cache_status value from the UTM
func GetUp2DateCacheStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/up2date.cache_status", &val, options...)
	return
}

// UpdateUp2DateCacheStatus PUTs the up2date.cache_status value to the UTM
func UpdateUp2DateCacheStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.cache_status", val, options...)
}

// GetUp2DateCacheUseAcc gets the up2date.cache_use_acc value from the UTM
func GetUp2DateCacheUseAcc(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/up2date.cache_use_acc", &val, options...)
	return
}

// UpdateUp2DateCacheUseAcc PUTs the up2date.cache_use_acc value to the UTM
func UpdateUp2DateCacheUseAcc(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.cache_use_acc", val, options...)
}

// GetUp2DateForceInsecureUp2Date gets the up2date.force_insecure_up2date value from the UTM
func GetUp2DateForceInsecureUp2Date(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/up2date.force_insecure_up2date", &val, options...)
	return
}

// UpdateUp2DateForceInsecureUp2Date PUTs the up2date.force_insecure_up2date value to the UTM
func UpdateUp2DateForceInsecureUp2Date(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.force_insecure_up2date", val, options...)
}

// GetUp2DateParentProxyAuthPass gets the up2date.parent_proxy_auth_pass value from the UTM
func GetUp2DateParentProxyAuthPass(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/up2date.parent_proxy_auth_pass", &val, options...)
	return
}

// UpdateUp2DateParentProxyAuthPass PUTs the up2date.parent_proxy_auth_pass value to the UTM
func UpdateUp2DateParentProxyAuthPass(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.parent_proxy_auth_pass", val, options...)
}

// GetUp2DateParentProxyAuthStatus gets the up2date.parent_proxy_auth_status value from the UTM
func GetUp2DateParentProxyAuthStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/up2date.parent_proxy_auth_status", &val, options...)
	return
}

// UpdateUp2DateParentProxyAuthStatus PUTs the up2date.parent_proxy_auth_status value to the UTM
func UpdateUp2DateParentProxyAuthStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.parent_proxy_auth_status", val, options...)
}

// GetUp2DateParentProxyAuthUser gets the up2date.parent_proxy_auth_user value from the UTM
func GetUp2DateParentProxyAuthUser(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/up2date.parent_proxy_auth_user", &val, options...)
	return
}

// UpdateUp2DateParentProxyAuthUser PUTs the up2date.parent_proxy_auth_user value to the UTM
func UpdateUp2DateParentProxyAuthUser(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.parent_proxy_auth_user", val, options...)
}

// GetUp2DateParentProxyHost gets the up2date.parent_proxy_host value from the UTM
func GetUp2DateParentProxyHost(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/up2date.parent_proxy_host", &val, options...)
	return
}

// UpdateUp2DateParentProxyHost PUTs the up2date.parent_proxy_host value to the UTM
func UpdateUp2DateParentProxyHost(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.parent_proxy_host", val, options...)
}

// GetUp2DateParentProxyPort gets the up2date.parent_proxy_port value from the UTM
func GetUp2DateParentProxyPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/up2date.parent_proxy_port", &val, options...)
	return
}

// UpdateUp2DateParentProxyPort PUTs the up2date.parent_proxy_port value to the UTM
func UpdateUp2DateParentProxyPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.parent_proxy_port", val, options...)
}

// GetUp2DateParentProxyStatus gets the up2date.parent_proxy_status value from the UTM
func GetUp2DateParentProxyStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/up2date.parent_proxy_status", &val, options...)
	return
}

// UpdateUp2DateParentProxyStatus PUTs the up2date.parent_proxy_status value to the UTM
func UpdateUp2DateParentProxyStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.parent_proxy_status", val, options...)
}

// GetUp2DatePatternDownloadInterval gets the up2date.pattern_download_interval value from the UTM
func GetUp2DatePatternDownloadInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/up2date.pattern_download_interval", &val, options...)
	return
}

// UpdateUp2DatePatternDownloadInterval PUTs the up2date.pattern_download_interval value to the UTM
func UpdateUp2DatePatternDownloadInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.pattern_download_interval", val, options...)
}

// GetUp2DatePatternDownloadStatus gets the up2date.pattern_download_status value from the UTM
func GetUp2DatePatternDownloadStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/up2date.pattern_download_status", &val, options...)
	return
}

// UpdateUp2DatePatternDownloadStatus PUTs the up2date.pattern_download_status value to the UTM
func UpdateUp2DatePatternDownloadStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.pattern_download_status", val, options...)
}

// GetUp2DateScheduledUp2Date gets the up2date.scheduled_up2date value from the UTM
func GetUp2DateScheduledUp2Date(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/up2date.scheduled_up2date", &val, options...)
	return
}

// UpdateUp2DateScheduledUp2Date PUTs the up2date.scheduled_up2date value to the UTM
func UpdateUp2DateScheduledUp2Date(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.scheduled_up2date", val, options...)
}

// GetUp2DateServers gets the up2date.servers value from the UTM
func GetUp2DateServers(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/up2date.servers", &val, options...)
	return
}

// UpdateUp2DateServers PUTs the up2date.servers value to the UTM
func UpdateUp2DateServers(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.servers", val, options...)
}

// GetUp2DateStatus gets the up2date.status value from the UTM
func GetUp2DateStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/up2date.status", &val, options...)
	return
}

// UpdateUp2DateStatus PUTs the up2date.status value to the UTM
func UpdateUp2DateStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.status", val, options...)
}

// GetUp2DateSystemAutoinstallTime gets the up2date.system_autoinstall_time value from the UTM
func GetUp2DateSystemAutoinstallTime(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/up2date.system_autoinstall_time", &val, options...)
	return
}

// UpdateUp2DateSystemAutoinstallTime PUTs the up2date.system_autoinstall_time value to the UTM
func UpdateUp2DateSystemAutoinstallTime(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.system_autoinstall_time", val, options...)
}

// GetUp2DateSystemDownloadInterval gets the up2date.system_download_interval value from the UTM
func GetUp2DateSystemDownloadInterval(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/up2date.system_download_interval", &val, options...)
	return
}

// UpdateUp2DateSystemDownloadInterval PUTs the up2date.system_download_interval value to the UTM
func UpdateUp2DateSystemDownloadInterval(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.system_download_interval", val, options...)
}

// GetUp2DateSystemDownloadStatus gets the up2date.system_download_status value from the UTM
func GetUp2DateSystemDownloadStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/up2date.system_download_status", &val, options...)
	return
}

// UpdateUp2DateSystemDownloadStatus PUTs the up2date.system_download_status value to the UTM
func UpdateUp2DateSystemDownloadStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/up2date.system_download_status", val, options...)
}

// GetUplinkActions gets the uplink.actions value from the UTM
func GetUplinkActions(client sophos.ClientInterface, options ...sophos.Option) (val []interface{}, err error) {
	err = get(client, "/api/nodes/uplink.actions", &val, options...)
	return
}

// UpdateUplinkActions PUTs the uplink.actions value to the UTM
func UpdateUplinkActions(client sophos.ClientInterface, val []interface{}, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/uplink.actions", val, options...)
}

// GetUplinkActive gets the uplink.active value from the UTM
func GetUplinkActive(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/uplink.active", &val, options...)
	return
}

// UpdateUplinkActive PUTs the uplink.active value to the UTM
func UpdateUplinkActive(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/uplink.active", val, options...)
}

// GetUplinkCondition gets the uplink.condition value from the UTM
func GetUplinkCondition(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/uplink.condition", &val, options...)
	return
}

// UpdateUplinkCondition PUTs the uplink.condition value to the UTM
func UpdateUplinkCondition(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/uplink.condition", val, options...)
}

// GetUplinkEnableMultipathStickyRuleFlagWorkaround gets the uplink.enable_multipath_sticky_rule_flag_workaround value from the UTM
func GetUplinkEnableMultipathStickyRuleFlagWorkaround(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/uplink.enable_multipath_sticky_rule_flag_workaround", &val, options...)
	return
}

// UpdateUplinkEnableMultipathStickyRuleFlagWorkaround PUTs the uplink.enable_multipath_sticky_rule_flag_workaround value to the UTM
func UpdateUplinkEnableMultipathStickyRuleFlagWorkaround(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/uplink.enable_multipath_sticky_rule_flag_workaround", val, options...)
}

// GetUplinkInterfaces gets the uplink.interfaces value from the UTM
func GetUplinkInterfaces(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/uplink.interfaces", &val, options...)
	return
}

// UpdateUplinkInterfaces PUTs the uplink.interfaces value to the UTM
func UpdateUplinkInterfaces(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/uplink.interfaces", val, options...)
}

// GetUplinkMonitoring gets the uplink.monitoring value from the UTM
func GetUplinkMonitoring(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/uplink.monitoring", &val, options...)
	return
}

// UpdateUplinkMonitoring PUTs the uplink.monitoring value to the UTM
func UpdateUplinkMonitoring(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/uplink.monitoring", val, options...)
}

// GetUplinkPassive gets the uplink.passive value from the UTM
func GetUplinkPassive(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/uplink.passive", &val, options...)
	return
}

// UpdateUplinkPassive PUTs the uplink.passive value to the UTM
func UpdateUplinkPassive(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/uplink.passive", val, options...)
}

// GetUplinkPrimary gets the uplink.primary value from the UTM
func GetUplinkPrimary(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/uplink.primary", &val, options...)
	return
}

// UpdateUplinkPrimary PUTs the uplink.primary value to the UTM
func UpdateUplinkPrimary(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/uplink.primary", val, options...)
}

// GetUplinkRules gets the uplink.rules value from the UTM
func GetUplinkRules(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/uplink.rules", &val, options...)
	return
}

// UpdateUplinkRules PUTs the uplink.rules value to the UTM
func UpdateUplinkRules(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/uplink.rules", val, options...)
}

// GetUplinkScheduler gets the uplink.scheduler value from the UTM
func GetUplinkScheduler(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/uplink.scheduler", &val, options...)
	return
}

// UpdateUplinkScheduler PUTs the uplink.scheduler value to the UTM
func UpdateUplinkScheduler(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/uplink.scheduler", val, options...)
}

// GetUplinkStatus gets the uplink.status value from the UTM
func GetUplinkStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/uplink.status", &val, options...)
	return
}

// UpdateUplinkStatus PUTs the uplink.status value to the UTM
func UpdateUplinkStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/uplink.status", val, options...)
}

// GetWebadminAllowedNetworks gets the webadmin.allowed_networks value from the UTM
func GetWebadminAllowedNetworks(client sophos.ClientInterface, options ...sophos.Option) (val []string, err error) {
	err = get(client, "/api/nodes/webadmin.allowed_networks", &val, options...)
	return
}

// UpdateWebadminAllowedNetworks PUTs the webadmin.allowed_networks value to the UTM
func UpdateWebadminAllowedNetworks(client sophos.ClientInterface, val []string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.allowed_networks", val, options...)
}

// GetWebadminCa gets the webadmin.ca value from the UTM
func GetWebadminCa(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/webadmin.ca", &val, options...)
	return
}

// UpdateWebadminCa PUTs the webadmin.ca value to the UTM
func UpdateWebadminCa(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.ca", val, options...)
}

// GetWebadminCert gets the webadmin.cert value from the UTM
func GetWebadminCert(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/webadmin.cert", &val, options...)
	return
}

// UpdateWebadminCert PUTs the webadmin.cert value to the UTM
func UpdateWebadminCert(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.cert", val, options...)
}

// GetWebadminDashboardRefresh gets the webadmin.dashboard_refresh value from the UTM
func GetWebadminDashboardRefresh(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/webadmin.dashboard_refresh", &val, options...)
	return
}

// UpdateWebadminDashboardRefresh PUTs the webadmin.dashboard_refresh value to the UTM
func UpdateWebadminDashboardRefresh(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.dashboard_refresh", val, options...)
}

// GetWebadminLanguage gets the webadmin.language value from the UTM
func GetWebadminLanguage(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/webadmin.language", &val, options...)
	return
}

// UpdateWebadminLanguage PUTs the webadmin.language value to the UTM
func UpdateWebadminLanguage(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.language", val, options...)
}

// GetWebadminLogAdminConnect gets the webadmin.log_admin_connect value from the UTM
func GetWebadminLogAdminConnect(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/webadmin.log_admin_connect", &val, options...)
	return
}

// UpdateWebadminLogAdminConnect PUTs the webadmin.log_admin_connect value to the UTM
func UpdateWebadminLogAdminConnect(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.log_admin_connect", val, options...)
}

// GetWebadminOfferWizard gets the webadmin.offer_wizard value from the UTM
func GetWebadminOfferWizard(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/webadmin.offer_wizard", &val, options...)
	return
}

// UpdateWebadminOfferWizard PUTs the webadmin.offer_wizard value to the UTM
func UpdateWebadminOfferWizard(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.offer_wizard", val, options...)
}

// GetWebadminPort gets the webadmin.port value from the UTM
func GetWebadminPort(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/webadmin.port", &val, options...)
	return
}

// UpdateWebadminPort PUTs the webadmin.port value to the UTM
func UpdateWebadminPort(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.port", val, options...)
}

// GetWebadminRestApi gets the webadmin.rest_api value from the UTM
func GetWebadminRestApi(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/webadmin.rest_api", &val, options...)
	return
}

// UpdateWebadminRestApi PUTs the webadmin.rest_api value to the UTM
func UpdateWebadminRestApi(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.rest_api", val, options...)
}

// GetWebadminTermsOfUseStatus gets the webadmin.terms_of_use.status value from the UTM
func GetWebadminTermsOfUseStatus(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/webadmin.terms_of_use.status", &val, options...)
	return
}

// UpdateWebadminTermsOfUseStatus PUTs the webadmin.terms_of_use.status value to the UTM
func UpdateWebadminTermsOfUseStatus(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.terms_of_use.status", val, options...)
}

// GetWebadminTermsOfUseText gets the webadmin.terms_of_use.text value from the UTM
func GetWebadminTermsOfUseText(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/webadmin.terms_of_use.text", &val, options...)
	return
}

// UpdateWebadminTermsOfUseText PUTs the webadmin.terms_of_use.text value to the UTM
func UpdateWebadminTermsOfUseText(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.terms_of_use.text", val, options...)
}

// GetWebadminTimeout gets the webadmin.timeout value from the UTM
func GetWebadminTimeout(client sophos.ClientInterface, options ...sophos.Option) (val int64, err error) {
	err = get(client, "/api/nodes/webadmin.timeout", &val, options...)
	return
}

// UpdateWebadminTimeout PUTs the webadmin.timeout value to the UTM
func UpdateWebadminTimeout(client sophos.ClientInterface, val int64, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.timeout", val, options...)
}

// GetWebadminTimeoutOnDashboard gets the webadmin.timeout_on_dashboard value from the UTM
func GetWebadminTimeoutOnDashboard(client sophos.ClientInterface, options ...sophos.Option) (val bool, err error) {
	err = get(client, "/api/nodes/webadmin.timeout_on_dashboard", &val, options...)
	return
}

// UpdateWebadminTimeoutOnDashboard PUTs the webadmin.timeout_on_dashboard value to the UTM
func UpdateWebadminTimeoutOnDashboard(client sophos.ClientInterface, val bool, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.timeout_on_dashboard", val, options...)
}

// GetWebadminTlsCiphers gets the webadmin.tls_ciphers value from the UTM
func GetWebadminTlsCiphers(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/webadmin.tls_ciphers", &val, options...)
	return
}

// UpdateWebadminTlsCiphers PUTs the webadmin.tls_ciphers value to the UTM
func UpdateWebadminTlsCiphers(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.tls_ciphers", val, options...)
}

// GetWebadminTlsProtocols gets the webadmin.tls_protocols value from the UTM
func GetWebadminTlsProtocols(client sophos.ClientInterface, options ...sophos.Option) (val string, err error) {
	err = get(client, "/api/nodes/webadmin.tls_protocols", &val, options...)
	return
}

// UpdateWebadminTlsProtocols PUTs the webadmin.tls_protocols value to the UTM
func UpdateWebadminTlsProtocols(client sophos.ClientInterface, val string, options ...sophos.Option) (err error) {
	return put(client, "/api/nodes/webadmin.tls_protocols", val, options...)
}
