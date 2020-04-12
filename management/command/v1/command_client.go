package v1

import (
	rest "k8s.io/client-go/rest"
)

type CommandV1Interface interface {
	RESTClient() rest.Interface
	ParticipantsGetter
	ConferencesGetter
	PlatformsGetter
}

type CommandV1Client struct {
	restClient rest.Interface
}

func (c *CommandV1Client) Participants() ParticipantInterface {
	return newParticipants(c)
}

func (c *CommandV1Client) Conferences() ParticipantInterface {
	return newConferences(c)
}

func (c *CommandV1Client) Platforms() ParticipantInterface {
	return newPlatforms(c)
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
