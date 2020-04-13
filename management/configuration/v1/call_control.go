package v1

import (
  "context"

  rest "k8s.io/client-go/rest"
)

type CallControlGetter interface {
  CallControls() CallControlInterface
}

type CallControlInterface interface {
  H323Gatekeeper(ctx context.Context) (interface{}, error)
	SIPCredentials(ctx context.Context) (interface{}, error)
	SIPProxy(ctx context.Context) (interface{}, error)
	TURNServer(ctx context.Context) (interface{}, error)
	STUNServer(ctx context.Context) (interface{}, error)
	PolicyServer(ctx context.Context) (interface{}, error)
}

type callControls struct {
  client rest.Interface
}

func newCallControls(c *ConfigurationV1Client) *callControls {
  return &callControls{
    client: c.RESTClient()
  }
}

func (s *systems) H323Gatekeeper(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *systems) SIPCredentials(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *systems) SIPProxy(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *systems) TURNServer(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *systems) STUNServer(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *systems) PolicyServer(ctx context.Context) (interface{}, error) {
  return nil, nil
}
