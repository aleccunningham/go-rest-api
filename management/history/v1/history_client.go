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

func (c *StatusV1Client) Backplanes() ParticipantInterface {
	return newParticipa
}

func (c *StatusV1Client) Conferences() ParticipantInterface {
	return newParticipa
}

func (c *StatusV1Client) Participants() ParticipantInterface {
	return newParticipa
}

func (c *StatusV1Client) Nodes() ParticipantInterface {
	return newParticipa
}

// New creates a new NetworkingV1Client for the given RESTClient.
func New(c rest.Interface) *CommandV1Client {
	return &CommandV1Client{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CommandV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
