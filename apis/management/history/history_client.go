package v1

import (
	rest "k8s.io/client-go/rest"
)

// HistoryV1Interface defines the interface for interacting with
// the management history API
type HistoryV1Interface interface {
	RESTClient() rest.Interface
	AlarmGetter
	BackplaneGetter
	ConferenceGetter
	ParticipantGetter
	NodeGetter
}

// HistoryV1Client defines the client for interacting with the
// management history API
type HistoryV1Client struct {
	restClient rest.Interface
}

// Alarms returns the current alarms
func (c *StatusV1Client) Alarms() ([]v1.Alarm, error) {
	var p []v1.Alarm

	return p, nil
}

// Backplanes return historical backplanes
func (c *StatusV1Client) Backplanes() ([]v1.Backplane, error) {
	var p []v1.Backplane

	return p, nil
}

// Conferences return the current conferences
func (c *StatusV1Client) Conferences() ([]v1.Conference, error) {
	var p []v1.Conference

	return p, nil
}

// Participants return the current participants
func (c *StatusV1Client) Participants() ([]v1.Participant, error) {
	var p []v1.Participant

	return p, nil
}

// Nodes return the current nodes
func (c *StatusV1Client) Nodes() ([]v1.Node, error) {
	var p []v1.Node

	return p, nil
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
