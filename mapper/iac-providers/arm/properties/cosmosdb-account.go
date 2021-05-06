package properties

// CosmosDBAccount exposes the properties for azurerm_cosmosdb_account resource.
var CosmosDBAccount *cosmosDBAccount

type cosmosDBAccount struct {
	IPAddressOrRange string
	Tags             string
	IPRules          string
}

func init() {
	CosmosDBAccount = &cosmosDBAccount{
		IPAddressOrRange: "ipAddressOrRange",
		Tags:             "tags",
		IPRules:          "ipRules",
	}
}
