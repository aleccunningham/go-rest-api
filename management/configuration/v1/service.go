package v1

import (
  "context"

  rest "k8s.io/client-go/rest"
)

type ServiceGetter interface {
  Services() ServiceInterface
}

type ServiceInterface interface {
  Conference(ctx context.Context) (interface{}, error)
	ConferenceAlias(ctx context.Context) (interface{}, error)
	IVRTheme(ctx context.Context) (interface{}, error)
	GatewayRoutingRule(ctx context.Context) (interface{}, error)
	Registration(ctx context.Context) (interface{}, error)
	Device(ctx context.Context) (interface{}, error)
	ConferenceSyncTemplate(ctx context.Context) (interface{}, error)
	LDAPSyncSource(ctx context.Context) (interface{}, error)
	LDAPSyncField(ctx context.Context) (interface{}, error)
}

type services struct {
  client rest.Interface
}

func newServices(c *ConfigurationV1Client) *services {
  return &services{
    client: c.RESTClient()
  }
}
