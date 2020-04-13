package v1

import (
	"a-gitlab-01.hst.awhstg.com/oc-upgrade-automation/test/rest"
	v1 "github.com/aleccunningham/go-rest-api/types/v1"
)

// StatusV1Interface defines the interface for interacting with
// the management status API
type StatusV1Interface interface {
	RESTClient() rest.Interface
}

// StatusV1Client defines the client for interacting with the
// management status API
type StatusV1Client struct {
	restClient rest.Interface
}

// New returns a new status client
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

// Participants return the current participants
func (c *StatusV1Client) Participants() ([]v1.Participant, error) {
	var p []v1.Participant

	return p, nil
}

// Conferences return the current conferences
func (c *StatusV1Client) Conferences() ([]v1.Conference, error) {
	var p []v1.Conference

	return p, nil
}

// Nodes return the current nodes
func (c *StatusV1Client) Nodes() ([]v1.Node, error) {
	var p []v1.Node

	return p, nil
}

// Locations return the current locations
func (c *StatusV1Client) Locations() ([]v1.Location, error) {
	var p []v1.Location

	return p, nil
}

// Backplanes return the current backplanes
func (c *StatusV1Client) Backplanes() ([]v1.Backplane, error) {
	var p []v1.Backplane

	return p, nil
}

// Alarms returns the current alarms
func (c *StatusV1Client) Alarms() ([]v1.Alarm, error) {
	var p []v1.Alarm

	return p, nil
}

// Licenses return the current licenses
func (c *StatusV1Client) Licenses() ([]v1.License, error) {
	var p []v1.License

	return p, nil
}

// CloudNodes return the current active cloud nodes
func (c *StatusV1Client) CloudNodes() ([]v1.CloudNode, error) {
	var p []v1.CloudNode

	return p, nil
}

// CloudMonitoredLocations return a summary of the current cloud monitored locations
func (c *StatusV1Client) CloudMonitoredLocations() ([]v1.CloudMonitoredLocations, error) {
	var p []v1.CloudMonitoredLocations

	return p, nil
}

// CloudOverflowLocations return the current cloud overflow location
func (c *StatusV1Client) CloudOverflowLocations() ([]v1.CloudOverflowLocations, error) {
	var p []v1.CloudOverflowLocations

	return p, nil
}
