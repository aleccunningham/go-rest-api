package configuration

// ManualDeployTemplate defines the required fields for deploying
// a new node using the manual deployment type. POSTing a payload
// containing an instance of this struct will return an XML file
// that can then be POSTed to a conference node to provision it
type ManualDeployTemplate struct {
	Name           string `json:"name"`
	Hostname       string `json:"hostname"`
	Domain         string `json:"domain"`
	NetworkAddress string `json:"address"`
	Netmask        string `json:"netmask"`
	NetworkGateway string `json:"gateway"`
	SSHPassword    string `json:"password"`
	NodeType       string `json:"node_type"`       // 'CONFERENCING'
	SystemLocation string `json:"system_location"` // '/api/admin/configuration/v1/system_location/1/'
	DeploymentType string `json:"deployment_type"` // 'MANUAL-PROVISION-ONLY'
}
