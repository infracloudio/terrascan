package properties

// ApplicationGateway exposes the properties for application_gateway resource.
var ApplicationGateway *applicationGateway

type applicationGateway struct {
	WAFConfiguration string
	Enabled          string
}

func init() {
	ApplicationGateway = &applicationGateway{
		WAFConfiguration: "webApplicationFirewallConfiguration",
		Enabled:          "enabled",
	}
}
