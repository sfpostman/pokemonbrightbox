package e2elibrarysdk

import (
	"example.com/e2e-library/sdk/books"
	"example.com/e2e-library/sdk/internal/clients/rest/hooks"
	"example.com/e2e-library/sdk/internal/configmanager"
	"time"
)

// E2eLibrarySDK is the main SDK client that provides access to all service endpoints.
// It manages configuration, authentication, and service instances with centralized settings.
type E2eLibrarySDK struct {
	Books   *books.Service
	manager *configmanager.ConfigManager
}

func NewE2eLibrarySDK(config Config) *E2eLibrarySDK {
	books := books.NewService()

	manager := configmanager.NewConfigManager(config)
	hook := hooks.NewDefaultHook()
	books.WithConfigManager(manager)
	books.WithHook(hook)

	return &E2eLibrarySDK{
		Books:   books,
		manager: manager,
	}
}

func (e *E2eLibrarySDK) SetBaseURL(baseURL string) {
	e.manager.SetBaseURL(baseURL)
}

func (e *E2eLibrarySDK) SetTimeout(timeout time.Duration) {
	e.manager.SetTimeout(timeout)
}

func (e *E2eLibrarySDK) SetAPIKey(apiKey string) {
	e.manager.SetAPIKey(apiKey)
}

// SetEnvironment configures the SDK to use the specified environment's base URL.
func (e *E2eLibrarySDK) SetEnvironment(environment Environment) {
	e.manager.SetBaseURL(string(environment))
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
