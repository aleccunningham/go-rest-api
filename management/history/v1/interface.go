package history

import "context"

type HistoryGetter interface {
	History() HistoryInterface
}

type HistoryInterface interface {
	Alarms(ctx context.Context) (interface{}, error)
	Backplane(ctx context.Context) (interface{}, error)
	BackplaneStatistics(ctx context.Context) (interface{}, error)
	Conferences(ctx context.Context) (interface{}, error)
	Participants(ctx context.Context) (interface{}, error)
	NodeEvents(ctx context.Context) (interface{}, error)
}
