package v1

import (
  "context"

  rest "k8s.io/client-go/rest"
)

type UtilitiesGetter interface {
  Utilities() UtilitiesInterface
}

type UtilitiesInterface interface {

}

type utilities struct {
  client rest.Interface
}

func newUtilities(c *ConfigurationV1Client) *utilities {
  return &utilities{
    client: c.RESTClient()
  }
}
