// For more information please refer to https://github.com/checkout/checkout-sdk-go
import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/customers"
)

// API Keys
api, err := checkout.
			Builder().
			StaticKeys().
			WithSecretKey("secret_key").
			WithEnvironment(configuration.Sandbox()). // or Environment.PRODUCTION
			Build()
if err != nil {
	return nil, err
}

// OAuth
api, err := checkout.
			Builder().
			OAuth().
			WithClientCredentials("client_id", "client_secret").
			WithScopes([]string{configuration.Vault}).
			WithEnvironment(configuration.Sandbox()). // or Environment.PRODUCTION
			Build()
if err != nil {
	return nil, err
}

request := customers.CustomerRequest{
	Email: "email@docs.checkout.com",
	Name:  "Name",
	Phone: &common.Phone{
		CountryCode: "44",
		Number:      "4155552671",
	},
}

response, err := api.Customers.Create(request)
if err != nil {
	return nil, err
}

return response, nil
