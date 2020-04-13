package v1

// ResponseMeta encompasses the meta object returned
// with all API response calls
type ResponseMeta struct {
	Limit      string `json:"limit,omitempty"`
	Next       string `json:"next,omitempty"`
	Offset     int    `json:"offset,omitempty"`
	Previous   int    `json:"previous,omitempty"`
	TotalCount int    `json:"total_count,omitempty"`
}

// ConferencingNodes defines a conference nodes properties
// https://<manageraddress>/api/admin/status/v1/worker_vm/
type ConferencingNodes struct {
	Meta  ResponseMeta `json:"meta"`
	Nodes []WorkerVM   `json:"objects"`
}

// SystemLocations encompasses all system locations
// https://<manageraddress>/api/admin/status/v1/system_location/
type SystemLocations struct {
	Meta      ResponseMeta `json:"meta"`
	Locations []Location   `json:"objects"`
}

// LocationStatistics encompasses a system locations statistics
// https://<manageraddress>/api/admin/status/v1/system_location/<system_location_id>/statistics/
type LocationStatistics struct {
	MediaLoad float32 `json:"media_load,omitempty"`
}

// CloudMonitoredLocations encompasses all locations that
// are being monitored for dynamic bursting
// https://<manageraddress>/api/admin/status/v1/cloud_monitored_location/
type CloudMonitoredLocations struct {
	Meta      ResponseMeta `json:"meta"`
	Locations []Location   `json:"objects"`
}

// CloudOverflowLocations encompasses all locations that
// may be used for dynamic bursting
// https://<manageraddress>/api/admin/status/v1/cloud_overflow_location/
type CloudOverflowLocations struct {
	Meta      ResponseMeta `json:"meta"`
	Locations []Location   `json:"objects"`
}

// WorkerVM defines a nodes attributes
type WorkerVM struct {
	Meta            ResponseMeta `json:"meta"`
	ID              int          `json:"id,omitempty"`
	Name            string       `json:"name,omitempty"`
	ResourceURI     string       `json:"resource_uri,omitempty"`
	Version         string       `json:"version,omitempty"`
	MaintenanceMode bool         `json:"maintenance_mode,omitempty"`
	ConfigurationID int          `json:"configuration_id,omitempty"`
	SystemLocation  string       `json:"system_location,omitempty"`
	LastUpdated     string       `json:"last_updated,omitempty"`
	LastReported    string       `json:"last_reported,omitempty"`
	MediaLoad       int          `json:"media_load,omitempty"`
	MediaTokensUsed int          `json:"media_tokens_used,omitempty"`
	SignalingCount  int          `json:"signaling_count,omitempty"`
	MaxAudioCalls   int          `json:"max_audio_calls,omitempty"`
	MaxSDCalls      int          `json:"max_sd_calls,omitempty"`
	MaxMediaTokens  int          `json:"max_media_tokens,omitempty"`
	MaxHDCalls      int          `json:"max_hd_calls,omitempty"`
	MaxFullHDCalls  int          `json:"max_full_hd_calls,omitempty"`
	Hypervisor      string       `json:"hypervisor,omitempty"`
	NodeType        string       `json:"node_type,omitempty"`
	TotalRAM        int          `json:"total_ram,omitempty"`
	CPUCount        int          `json:"cpu_count,omitempty"`
	CPUModel        string       `json:"cpu_model,omitempty"`
	CPUCapabilities string       `json:"cpu_capabilities,omitempty"`
	DeployProgress  int          `json:"deploy_progress,omitempty"`
	DeployStatus    string       `json:"deploy_status,omitempty"`
	DeployError     string       `json:"deploy_error,omitempty"`
	UpgradeStatus   string       `json:"upgrade_status,omitempty"`
}

type Participant struct {
	UUID                  string  `json:"uuid"`
	URI                   string  `json:"uri"`
	DisplayName           string  `json:"display_name"`
	HasMedia              bool    `json:"has_media"`
	StartTime             uint    `json:"start_time"`
	Vad                   uint    `json:"vad"`
	IsMuted               string  `json:"is_muted"`
	IsPresenting          string  `json:"is_presenting"`
	RxPresentationPolicy  string  `json:"rx_presentation_policy"`
	Role                  string  `json:"role"`
	ServiceType           string  `json:"service_type"`
	Type                  string  `json:"type"`
	PresentationSupported string  `json:"presentation_supported"`
	StageIndex            int     `json:"stage_index"`
	OverlayText           string  `json:"overlay_text"`
	Spotlight             float64 `json:"spotlight"`
}

type ParticipantsResult struct {
	Status  string        `json:"statusx"`
	Objects []Participant `json:"result"`
}

type ConferencesResult struct {
	Objects []struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		StartTime string `json:"start_time"`
		IsLocked  bool   `json:"is_locked"`
	} `json:"objects"`
}

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

// Location encompasses a Pexip location
type Location struct {
	ID               int    `json:"id,omitempty"`
	SystemLocationID int    `json:"systemlocation_id,omitempty"`
	ResourceURI      string `json:"resource_uri,omitempty"`
	Name             string `json:"name,omitempty"`
	MaxHDCalls       int    `json:"max_hd_calls,omitempty"`
	FreeHDCalls      int    `json:"free_hd_calls,omitempty"`
	MediaLoad        int    `json:"media_load,omitempty"`
}

// CloudNode encompases a cloud overflow Conference Node
// https://<manageraddress>/api/admin/status/v1/cloud_node/
type CloudNode struct {
	CloudInstanceID                   string `json:"cloud_instance_id,omitempty"`
	CloudInstanceIP                   string `json:"cloud_instance_ip,omitempty"`
	CloudInstanceLaunchTime           string `json:"cloud_instance_launch_time,omitempty"`
	MaxHDCalls                        int    `json:"max_hd_calls,omitempty"`
	MediaLoad                         int    `json:"media_load,omitempty"`
	ResourceURI                       string `json:"resource_uri,omitempty"`
	WorkerVMConfigurationID           int    `json:"workervm_configuration_id,omitempty"`
	WorkerVMConfigurationLocationName string `json:"workervm_configuration_location_name,omitempty"`
	WorkerVMConfigurationName         string `json:"workervm_configuration_name,omitempty"`
}
