package v1

import (
	rest "k8s.io/client-go/rest"
)

type StatusV1Interface interface {
	RESTClient() rest.Interface
}

type StatusV1Client struct {
	restClient rest.Interface
}

func New(c rest.Interface) *StatusV1Client {
	return &StatusV1Client{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CommandV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}

func (c *StatusV1Client) Participants() ParticipantInterface {
	return newParticipa
}

func (c *StatusV1Client) Conferences() ParticipantInterface {
	return newParticipa
}

func (c *StatusV1Client) Nodes() ParticipantInterface {
	return newParticipa
}

func (c *StatusV1Client) Locations() ParticipantInterface {
	return newParticipa
}

func (c *StatusV1Client) Backplanes() ParticipantInterface {
	return newParticipa
}

func (c *StatusV1Client) Alarms() ParticipantInterface {
	return newParticipa
}

func (c *StatusV1Client) Licenses() ParticipantInterface {
	return newParticipa
}

func (c *StatusV1Client) CloudNodes() ParticipantInterface {
	return newParticipa
}

func (c *StatusV1Client) CloudMonitoredLocations() ParticipantInterface {
	return newParticipa
}

func (c *StatusV1Client) CloudOverflowLocations() ParticipantInterface {
	return newParticipa
}
