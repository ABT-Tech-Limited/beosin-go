package beosin

import (
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the default base URL for Beosin API
	DefaultBaseURL = "https://api.beosin.com"

	// DefaultTimeout is the default timeout for HTTP requests
	DefaultTimeout = 30 * time.Second
)

// Options holds the configuration for the Beosin client
type Options struct {
	// BaseURL is the base URL for the Beosin API
	BaseURL string

	// AppID is the application ID for authentication
	AppID string

	// AppSecret is the application secret for authentication
	AppSecret string

	// Timeout is the timeout for HTTP requests
	Timeout time.Duration

	// HTTPClient is the HTTP client to use for requests
	HTTPClient *http.Client

	// Debug enables debug logging
	Debug bool
}

// Option is a function that configures Options
type Option func(*Options)

// WithBaseURL sets the base URL for the client
func WithBaseURL(url string) Option {
	return func(o *Options) {
		o.BaseURL = url
	}
}

// WithTimeout sets the timeout for HTTP requests
func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(client *http.Client) Option {
	return func(o *Options) {
		o.HTTPClient = client
	}
}

// WithDebug enables or disables debug logging
func WithDebug(debug bool) Option {
	return func(o *Options) {
		o.Debug = debug
	}
}

// applyDefaults applies default values to options
func (o *Options) applyDefaults() {
	if o.BaseURL == "" {
		o.BaseURL = DefaultBaseURL
	}
	if o.Timeout == 0 {
		o.Timeout = DefaultTimeout
	}
	if o.HTTPClient == nil {
		o.HTTPClient = &http.Client{
			Timeout: o.Timeout,
		}
	}
}
