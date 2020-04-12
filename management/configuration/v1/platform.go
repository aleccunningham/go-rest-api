package v1

import (
  "context"

  rest "k8s.io/client-go/rest"
)

type PlatformGetter interface {
  Platforms() PlatformInterface
}

type PlatformInterface interface {
  SystemLocation(ctx context.Context) (interface{}, error)
	ManagementNode(ctx context.Context) (interface{}, error)
	ConferenceNode(ctx context.Context) (interface{}, error)
	License(ctx context.Context) (interface{}, error)
	LicenseRequest(ctx context.Context) (interface{}, error)
	CACertificate(ctx context.Context) (interface{}, error)
	TLSCertificate(ctx context.Context) (interface{}, error)
	CertificateSigningRequest(ctx context.Context) (interface{}, error)
	Global(ctx context.Context) (interface{}, error)
	DiagnosticGraphs(ctx context.Context) (interface{}, error)
}

type platforms struct {
  client rest.Interface
}

func newPlatforms(c *ConfigurationV1Client) *platforms {
  return &platforms{
    client: c.RESTClient()
  }
}
