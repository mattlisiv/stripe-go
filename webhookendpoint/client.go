// Package webhookendpoint provides the /webhook_endpoints APIs
package webhookendpoint

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /webhook_endpoints APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new webhook_endpoint.
func New(params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	return getC().New(params)
}

// New creates a new webhook_endpoint.
func (c Client) New(params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	webhook_endpoint := &stripe.WebhookEndpoint{}
	err := c.B.Call(http.MethodPost, "/webhook_endpoints", c.Key, params, webhook_endpoint)
	return webhook_endpoint, err
}

// Get returns the details of a webhook endpoint.
func Get(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	return getC().Get(id, params)
}

// Get returns the details of a webhook endpoint.
func (c Client) Get(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	path := stripe.FormatURLPath("/webhook_endpoints/%s", id)
	webhook_endpoint := &stripe.WebhookEndpoint{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, webhook_endpoint)
	return webhook_endpoint, err
}

// Update updates a webhook endpoint's properties.
func Update(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	return getC().Update(id, params)
}

// Update updates a webhook endpoint's properties.
func (c Client) Update(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	path := stripe.FormatURLPath("/webhook_endpoints/%s", id)
	webhook_endpoint := &stripe.WebhookEndpoint{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, webhook_endpoint)
	return webhook_endpoint, err
}

// Del removes a webhook endpoint.
func Del(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	return getC().Del(id, params)
}

// Del removes a webhook endpoint.
func (c Client) Del(id string, params *stripe.WebhookEndpointParams) (*stripe.WebhookEndpoint, error) {
	path := stripe.FormatURLPath("/webhook_endpoints/%s", id)
	webhook_endpoint := &stripe.WebhookEndpoint{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, webhook_endpoint)
	return webhook_endpoint, err
}

// List returns a list of webhook_endpoints.
func List(params *stripe.WebhookEndpointListParams) *Iter {
	return getC().List(params)
}

// List returns a list of webhook_endpoints.
func (c Client) List(listParams *stripe.WebhookEndpointListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.WebhookEndpointList{}
		err := c.B.CallRaw(http.MethodGet, "/webhook_endpoints", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for webhook_endpoints.
type Iter struct {
	*stripe.Iter
}

// WebhookEndpoint returns the webhook_endpoint which the iterator is currently pointing to.
func (i *Iter) WebhookEndpoint() *stripe.WebhookEndpoint {
	return i.Current().(*stripe.WebhookEndpoint)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
