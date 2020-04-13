package configuration

import (
  "context"

  rest "k8s.io/client-go/rest"
)

type SystemGetter interface {
  Systems() SystemInterface
}

type SystemInterface interface {
  DNS(ctx context.Context) (interface{}, error)
	NTP(ctx context.Context) (interface{}, error)
  SNMP(ctx context.Context) (interface{}, error)
	WebProxy(ctx context.Context) (interface{}, error)
	SyslogServer(ctx context.Context) (interface{}, error)
	EventSync(ctx context.Context) (interface{}, error)
}

type systems struct {
  client rest.Interface
}

func newSystsems(c *ConfigurationV1Client) *systems {
  return &systems{
    client: c.RESTClient()
  }
}

func (s *systems) DNS(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *systems) NTP(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *systems) SNMP(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *systems) SyslogServer(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *systems) WebProxy(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *systems) EventSync(ctx context.Context) (interface{}, error) {
  return nil, nil
}
