package v1

import (
  "context"

  rest "k8s.io/client-go/rest"
)

type SystemGetter interface {
  Systems() SystemInterface
}

type SystemInterface interface {
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

type systems struct {
  client rest.Interface
}

func newSystsems(c *ConfigurationV1Client) *systems {
  return &systems{
    client: c.RESTClient()
  }
}
