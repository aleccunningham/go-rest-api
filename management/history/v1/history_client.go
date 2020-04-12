package v1

import (
	rest "k8s.io/client-go/rest"
)

type HistoryV1Interface interface {
	RESTClient() rest.Interface
	AlarmGetter
	BackplaneGetter
	ConferenceGetter
	ParticipantGetter
	NodeGetter
}

type HistoryV1Client struct {
	restClient rest.Interface
}

func (c *StatusV1Client) Alarms() ParticipantInterface {
	return newParticipa
}
