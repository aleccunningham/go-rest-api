package configuration

import "context"

type ConfigurationGetter interface {
	Configuration() ConfigurationInterface
}

type ConfigurationInterface interface {
	System() SystemInterface
	Platform() PlatformInterface
	CallControl() CallControlInterface
	Service() ServiceInterface
	Users() UserInterface
	Utilities() UtilitiesInterface
}

type SystemInterface interface {
	DNS(ctx context.Context) (interface{}, error)
	NTP(ctx context.Context) (interface{}, error)
	WebProxy(ctx context.Context) (interface{}, error)
	SyslogServer(ctx context.Context) (interface{}, error)
	SNMP(ctx context.Context) (interface{}, error)
	EventSync(ctx context.Context) (interface{}, error)
}

type PlatformInterface interface {
	SystemLocation(ctx context.Context) (interface{}, error)
	ManagementNode(ctx context.Context) (interface{}, error)
	ConferenceNode(ctx context.Context) (interface{}, error)
	License(ctx context.Context) (interface{}, error)
	LicenseRequest(ctx context.Context) (interface{}, error)
	CACertificate(ctx context.Context) (interface{}, error)
	TLSCertificate(ctx context.Context) (interface{}, error)
	CertificateSigningRequest(ctx context.Context) (interface{}, error)
	Global(ctx context.Context) (interface{}, error)
	DiagnosticGraphs(ctx context.Context) (interface{}, error)
}

type CallControlInterface interface {
	H323Gatekeeper(ctx context.Context) (interface{}, error)
	SIPCredentials(ctx context.Context) (interface{}, error)
	SIPProxy(ctx context.Context) (interface{}, error)
	TURNServer(ctx context.Context) (interface{}, error)
	STUNServer(ctx context.Context) (interface{}, error)
	PolicyServer(ctx context.Context) (interface{}, error)
}

type ServiceInterface interface {
	Conference(ctx context.Context) (interface{}, error)
	ConferenceAlias(ctx context.Context) (interface{}, error)
	IVRTheme(ctx context.Context) (interface{}, error)
	GatewayRoutingRule(ctx context.Context) (interface{}, error)
	Registration(ctx context.Context) (interface{}, error)
	Device(ctx context.Context) (interface{}, error)
	ConferenceSyncTemplate(ctx context.Context) (interface{}, error)
	LDAPSyncSource(ctx context.Context) (interface{}, error)
	LDAPSyncField(ctx context.Context) (interface{}, error)
}

type UserInterface interface {
	Authentication(ctx context.Context) (interface{}, error)
	AccountRole(ctx context.Context) (interface{}, error)
	LDAPRole(ctx context.Context) (interface{}, error)
	Permission(ctx context.Context) (interface{}, error)
	EndUser(ctx context.Context) (interface{}, error)
}

type UtilitiesInterface interface {
	Upgrade(ctx context.Context) (interface{}, error)
	Backup(ctx context.Context) (interface{}, error)
	AutoBackup(ctx context.Context) (interface{}, error)
}
