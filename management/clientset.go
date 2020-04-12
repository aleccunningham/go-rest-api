package management

import (
	"net/http"

	"github.com/aleccunningham/go-rest-api/typed/command"
	"github.com/aleccunningham/go-rest-api/typed/configuration"
	"github.com/aleccunningham/go-rest-api/typed/history"
	"github.com/aleccunningham/go-rest-api/typed/status"
)

type PexipAPI struct {
	http *http.Client
}

type Interface interface {
	StatusV1() status.StatusV1Interface
	ConfigurationV1() configuration.ConfigurationV1Interface
	HistoryV1() history.HistoryV1Interface
	CommandV1() command.CommandV1Interface
}
