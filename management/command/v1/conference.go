package v1

import (
  "context"

  rest "k8s.io/client-go/rest"
)

type ConferencesGetter interface {
  Conferences() ConferenceInterface
}

type ConferenceInterface interface {
	Dial(ctx context.Context) (interface{}, error)
	Disconnect(ctx context.Context) (interface{}, error)
	Mute(ctx context.Context) (interface{}, error)
	Lock(ctx context.Context) (interface{}, error)
	Unlock(ctx context.Context) (interface{}, error)
	SendEmail(ctx context.Context) (interface{}, error)
}

type conferences struct {
  client rest.Interface
}

func newConferences(c *CommandV1Client) *conferences {
  return &conferences{
    client: c.RESTClient()
  }
}
