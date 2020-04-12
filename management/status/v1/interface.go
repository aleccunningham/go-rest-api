package status

import "context"

type StatusGetter interface {
	Status() StatusInterface
}

type StatusInterface interface {
	Participants(ctx context.Context) (interface{}, error)
	Conferences(ctx context.Context) (interface{}, error)
	Nodes(ctx context.Context) (interface{}, error)
	Locations(ctx context.Context) (interface{}, error)
	LocationStatistics(ctx context.Context) (interface{}, error)
	Backplanes(ctx context.Context) (interface{}, error)
	BackplanesStatistics(ctx context.Context) (interface{}, error)
	Alarms(ctx context.Context) (interface{}, error)
	Licenses(ctx context.Context) (interface{}, error)
	CloudNodes(ctx context.Context) (interface{}, error)
	CloudMonitoredLocations(ctx context.Context) (interface{}, error)
	CloudOverflowLocations(ctx context.Context) (interface{}, error)
}
