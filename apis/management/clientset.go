package management

import (
	"github.com/aleccunningham/go-rest-api/apis/management/command"
	"github.com/aleccunningham/go-rest-api/apis/management/configuration"
	"github.com/aleccunningham/go-rest-api/apis/management/history"
	"github.com/aleccunningham/go-rest-api/apis/management/status"
)

// Interface defines the interface for the Management API clientset,
// which is a subset of the APIs provided by Pexip
type Interface interface {
	StatusV1() status.StatusV1Interface
	ConfigurationV1() configuration.ConfigurationV1Interface
	HistoryV1() history.HistoryV1Interface
	CommandV1() command.CommandV1Interface
}
