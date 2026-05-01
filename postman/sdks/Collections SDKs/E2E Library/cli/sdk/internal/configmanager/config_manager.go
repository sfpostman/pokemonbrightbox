package configmanager

import (
	"example.com/e2e-library/sdk/e2elibrarysdkconfig"
	"time"
)

// ConfigManager manages configuration across all services with synchronized updates.
// Provides centralized configuration management and OAuth token handling for multiple services.
type ConfigManager struct {
	books e2elibrarysdkconfig.Config
}

// NewConfigManager creates a new configuration manager with the provided config and optional OAuth token service.
// Initializes service-specific configs and sets up OAuth token management if enabled.
func NewConfigManager(config e2elibrarysdkconfig.Config) *ConfigManager {
	return &ConfigManager{
		books: config,
	}
}

// SetBaseURL updates the BaseURL configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetBaseURL(baseURL string) {
	c.books.SetBaseURL(baseURL)
}

// SetTimeout updates the Timeout configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.books.SetTimeout(timeout)
}

// SetAPIKey updates the APIKey configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetAPIKey(apiKey string) {
	c.books.SetAPIKey(apiKey)
}

// SetRetryConfig updates the retry configuration across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetRetryConfig(retry e2elibrarysdkconfig.RetryConfig) {
	c.books.SetRetryConfig(retry)
}

// GetBooks returns the configuration for the Books service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetBooks() *e2elibrarysdkconfig.Config {
	return &c.books
}

// GetBaseURL returns the currently configured base URL.
// All services share the same base URL; this reads it from the first service's config.
func (c *ConfigManager) GetBaseURL() string {
	return c.books.BaseURL
}
