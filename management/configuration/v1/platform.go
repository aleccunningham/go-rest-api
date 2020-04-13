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

func (s *platforms) SystemLocation(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *platforms) ManagementNode(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *platforms) ConferenceNode(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *platforms) License(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *platforms) LicenseRequest(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *platforms) CACertificate(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *platforms) TLSCertificate(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *platforms) CertificateSigningRequest(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *platforms) Global(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (s *platforms) DiagnosticGraphs(ctx context.Context) (interface{}, error) {
  return nil, nil
}
