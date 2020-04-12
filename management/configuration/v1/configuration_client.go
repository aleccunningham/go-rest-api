package v1

import "k8s.io/client-go/rest"

type ConfigurationV1Interface interface {
	RESTClient() rest.Interface
	SystemGetter
	PlatformGetter
	CallControlGetter
	ServiceGetter
	UserGetter
	UtilitiesGetter
}

type ConfigurationV1Client struct {
	restClient rest.Interface
}

func (c *ConfigurationV1Client) System() SystemInterface {
	return newSystems(c)
}

func (c *ConfigurationV1Client) Platforms() SystemInterface {
	return newPlatforms(c)
}

func (c *ConfigurationV1Client) CallControls() SystemInterface {
	return newCallControls(c)
}

func (c *ConfigurationV1Client) Services() SystemInterface {
	return newServices(c)
}

func (c *ConfigurationV1Client) Users() SystemInterface {
	return newUsers(c)
}

func (c *ConfigurationV1Client) Utilities() SystemInterface {
	return newUtilities(c)
}
