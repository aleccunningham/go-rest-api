package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	v1 "github.com/aleccunningham/go-rest-api/types/v1"
)

// StatusURI specifies status endpoint
const StatusURI string = "api/admin/status/v1/"

// StatusV1Interface defines the interface for interacting with
// the management status API
type StatusV1Interface interface {
	Participants() ([]v1.Participant, error)
	Conferences() ([]v1.Conference, error)
	Nodes() ([]v1.Node, error)
	Locations() ([]v1.Location, error)
	Backplanes() ([]v1.Backplane, error)
	Alarms() ([]v1.Alarm, error)
	Licenses() ([]v1.License, error)
	CloudNodes() ([]v1.CloudNode, error)
	CloudMonitoredLocations() ([]v1.CloudMonitoredLocations, error)
	CloudOverflowLocations() ([]v1.CloudOverflowLocations, error)
}

// StatusV1Client defines the client for interacting with the
// management status API
type StatusV1Client struct {
	client  *http.Client
	BaseURL string
}

// New returns a new status client
func New(c *http.Client, b string) *StatusV1Client {
	return &StatusV1Client{
		client:  c,
		BaseURL: b,
	}
}

// Participants return the current participants
func (c *StatusV1Client) Participants() ([]v1.Participant, error) {
	url := fmt.Sprintf("https://%s/%s/participant/", c.BaseURL, StatusURI)
	var p []v1.Participant

	resp, err := c.client.Get(url)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&p)

	return p, nil
}

// Conferences return the current conferences
func (c *StatusV1Client) Conferences() ([]v1.Conference, error) {
	url := fmt.Sprintf("https://%s/%s/conference/", c.BaseURL, StatusURI)
	var p []v1.Conference

	resp, err := c.client.Get(url)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&p)

	return p, nil
}

// Nodes return the current nodes
func (c *StatusV1Client) Nodes() ([]v1.Node, error) {
	url := fmt.Sprintf("https://%s/%s/worker_vm/", c.BaseURL, StatusURI)
	var p []v1.WorkerVM

	resp, err := c.client.Get(url)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&p)

	return p, nil
}

// Locations return the current locations
func (c *StatusV1Client) Locations() ([]v1.Location, error) {
	url := fmt.Sprintf("https://%s/%s/system_location/", c.BaseURL, StatusURI)
	var p []v1.Location

	resp, err := c.client.Get(url)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&p)

	return p, nil
}

// Backplanes return the current backplanes
func (c *StatusV1Client) Backplanes() ([]v1.Backplane, error) {
	url := fmt.Sprintf("https://%s/%s/backplane/", c.BaseURL, StatusURI)
	var p []v1.Backplane

	resp, err := c.client.Get(url)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&p)

	return p, nil
}

// Alarms returns the current alarms
func (c *StatusV1Client) Alarms() ([]v1.Alarm, error) {
	url := fmt.Sprintf("https://%s/%s/alarm/", c.BaseURL, StatusURI)
	var p []v1.Alarm

	resp, err := c.client.Get(url)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&p)

	return p, nil
}

// Licenses return the current licenses
func (c *StatusV1Client) Licenses() ([]v1.License, error) {
	url := fmt.Sprintf("https://%s/%s/license/", c.BaseURL, StatusURI)
	var p []v1.License

	resp, err := c.client.Get(url)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&p)

	return p, nil
}

// CloudNodes return the current active cloud nodes
func (c *StatusV1Client) CloudNodes() ([]v1.CloudNode, error) {
	url := fmt.Sprintf("https://%s/%s/cloud_node/", c.BaseURL, StatusURI)
	var p []v1.CloudNode

	resp, err := c.client.Get(url)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&p)

	return p, nil
}

// CloudMonitoredLocations return a summary of the current cloud monitored locations
func (c *StatusV1Client) CloudMonitoredLocations() ([]v1.CloudMonitoredLocations, error) {
	url := fmt.Sprintf("https://%s/%s/cloud_monitored_location/", c.BaseURL, StatusURI)
	var p []v1.CloudMonitoredLocations

	resp, err := c.client.Get(url)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&p)

	return p, nil
}

// CloudOverflowLocations return the current cloud overflow location
func (c *StatusV1Client) CloudOverflowLocations() ([]v1.CloudOverflowLocations, error) {
	url := fmt.Sprintf("https://%s/%s/cloud_overflow_location/", c.BaseURL, StatusURI)
	var p []v1.CloudOverflowLocations

	resp, err := c.client.Get(url)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&p)

	return p, nil
}
