package command

import (
  "context"

  rest "k8s.io/client-go/rest"
)

type ParticipantGetter interface {
  Participants() ParticipantInterface
}

type ParticipantInterface interface {
	Dial(ctx context.Context) (interface{}, error)
	Disconnect(ctx context.Context) (interface{}, error)
	Mute(ctx context.Context) (interface{}, error)
	Unmute(ctx context.Context) (interface{}, error)
	Unlock(ctx context.Context) (interface{}, error)
	Transfer(ctx context.Context) (interface{}, error)
	Role(ctx context.Context) (interface{}, error)
}

type participants struct {
  client rest.Interface
}

func newParticipants(c *CommandV1Client) *participants {
  return &participants{
    client: c.RESTClient()
  }
}

func (c *conferences) Dial(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (c *conferences) Disconnect(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (c *conferences) Mute(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (c *conferences) Unmute(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (c *conferences) Unlock(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (c *conferences) Transfer(ctx context.Context) (interface{}, error) {
  return nil, nil
}

func (c *conferences) Role(ctx context.Context) (interface{}, error) {
  return nil, nil
}
