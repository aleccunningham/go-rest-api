package v1

import (
	rest "k8s.io/client-go/rest"
)

type StatusV1Interface interface {
	RESTClient() rest.Interface
	ParticipantsGetter
	ConferenceGetter
	NodeGetter
	LocationGetter
	BackplaneGetter
	AlarmGetter
	LicenseGetter
	CloudNodeGetter
	CloudMonitoredLocationGetter
	CloudOverflowLocationGetter
}

type StatusV1Client struct {
	restClient rest.Interface
}

func (c *StatusV1Client) Participants() ParticipantInterface {
	return newParticipa
}
