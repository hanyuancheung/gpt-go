// Package gpt provides a client for the OpenAI GPT-3 API
package gpt

import (
	"net/http"
	"time"
)

// Option sets gpt-3 Client option values.
type Option interface {
	apply(client) error
}

// ClientOption are options that can be passed when creating a new client
type ClientOption func(*client) *client

func (fn ClientOption) apply(cli *client) *client {
	return fn(cli)
}

// WithOrg is a client option that allows you to override the organization ID
func WithOrg(id string) ClientOption {
	return func(cli *client) *client {
		cli.idOrg = id
		return cli
	}
}

// WithDefaultEngine is a client option that allows you to override the default engine of the client
func WithDefaultEngine(engine string) ClientOption {
	return func(cli *client) *client {
		cli.defaultEngine = engine
		return cli
	}
}

// WithUserAgent is a client option that allows you to override the default user agent of the client
func WithUserAgent(userAgent string) ClientOption {
	return func(cli *client) *client {
		cli.userAgent = userAgent
		return cli
	}
}

// WithBaseURL is a client option that allows you to override the default base url of the client.
// The default base url is "https://api.openai.com/v1"
func WithBaseURL(baseURL string) ClientOption {
	return func(cli *client) *client {
		cli.baseURL = baseURL
		return cli
	}
}

// WithHTTPClient allows you to override the internal http.Client used
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(cli *client) *client {
		cli.httpClient = httpClient
		return cli
	}
}

// WithTimeout is a client option that allows you to override the default timeout duration of requests
// for the client. The default is 30 seconds. If you are overriding the http client as well, just include
// the timeout there.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(cli *client) *client {
		cli.httpClient.Timeout = timeout
		return cli
	}
}
