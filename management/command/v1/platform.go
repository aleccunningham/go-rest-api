package v1

import (
  "context"

  rest "k8s.io/client-go/rest"
)

type PlatformGetter interface {
  Platforms() PlatformInterface
}

type PlatformInterface interface {
	CreateBackup(ctx context.Context) (interface{}, error)
	RestoreBackup(ctx context.Context) (interface{}, error)
	ImportCertificate(ctx context.Context) (interface{}, error)
	StartCloudNode(ctx context.Context) (interface{}, error)
	TakeSnapshot(ctx context.Context) (interface{}, error)
}

type platforms struct {
	client rest.Interface
}

func newPlatforms(c *CommandV1Client) *platforms {
	return &platforms{
		client: c.RESTClient()
	}
}

func (p *platforms) CreateBackup(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (p *platforms) RestoreBackup(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (p *platforms) ImportCertificate(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (p *platforms) StartCloudNode(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (p *platforms) TakeSnapshot(ctx context.Context) (interface{}, error) {
  return nil, nil
}
