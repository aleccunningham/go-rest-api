package v1

import (
  "context"

  rest "k8s.io/client-go/rest"
)

type UserGetter interface {
  Users() UserInterface
}

type UserInterface interface {
  Authentication(ctx context.Context) (interface{}, error)
	AccountRole(ctx context.Context) (interface{}, error)
	LDAPRole(ctx context.Context) (interface{}, error)
	Permission(ctx context.Context) (interface{}, error)
	EndUser(ctx context.Context) (interface{}, error)
}

type users struct {
  client rest.Interface
}

func newUsers(c *ConfigurationV1Client) *users {
  return &users{
    client: c.RESTClient()
  }
}
