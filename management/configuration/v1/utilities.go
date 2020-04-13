package v1

import (
  "context"

  rest "k8s.io/client-go/rest"
)

type UtilitiesGetter interface {
  Utilities() UtilitiesInterface
}

type UtilitiesInterface interface {
  Upgrade(ctx context.Context) (interface{}, error)
	Backup(ctx context.Context) (interface{}, error)
	AutoBackup(ctx context.Context) (interface{}, error)
}

type utilities struct {
  client rest.Interface
}

func newUtilities(c *ConfigurationV1Client) *utilities {
  return &utilities{
    client: c.RESTClient()
  }
}

func (u *utilities) Upgrade(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (u *utilities) Backup(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (u *utilities) AutoBackup(ctx context.Context) (interface{}, error) {
  return nil, nil
}
