// For more information please refer to https://github.com/checkout/checkout-sdk-go
import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/configuration"
	webhooks "github.com/checkout/checkout-sdk-go/webhooks/abc"
)

api, err := checkout.
			Builder().
			Previous().
			WithSecretKey("secret_key").
			WithEnvironment(configuration.Sandbox()). // or Environment.PRODUCTION
			Build()
if err != nil {
	return nil, err
}

request := webhooks.WebhookRequest{
	Url:        "https://docs.checkout.com/webhook",
	EventTypes: []string{"payment_captured", "payment_approved", "payment_declined"},
}

response, err := api.Webhooks.RegisterWebhook(request)
if err != nil {
	return nil, err
}

return response, nil
