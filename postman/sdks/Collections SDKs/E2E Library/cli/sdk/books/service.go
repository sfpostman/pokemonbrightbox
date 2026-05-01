package books

import (
	"context"
	"example.com/e2e-library/sdk/e2elibrarysdkconfig"
	restClient "example.com/e2e-library/sdk/internal/clients/rest"
	"example.com/e2e-library/sdk/internal/clients/rest/hooks"
	"example.com/e2e-library/sdk/internal/clients/rest/httptransport"
	"example.com/e2e-library/sdk/internal/configmanager"
	"example.com/e2e-library/sdk/sharedmodels"
	"time"
)

// Service provides methods to interact with Books-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager                   *configmanager.ConfigManager
	hook                      hooks.Hook
	fetchAListOfBooksConfig   []e2elibrarysdkconfig.RequestOption
	createANewBookConfig      []e2elibrarysdkconfig.RequestOption
	verifyTheBookExistsConfig []e2elibrarysdkconfig.RequestOption
	checkoutNewBookConfig     []e2elibrarysdkconfig.RequestOption
}

func NewService() *Service {
	return &Service{
		manager: configmanager.NewConfigManager(e2elibrarysdkconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *Service) WithConfigManager(manager *configmanager.ConfigManager) *Service {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *Service) WithHook(hook hooks.Hook) *Service {
	api.hook = hook
	return api
}

func (api *Service) config() *e2elibrarysdkconfig.Config {
	return api.manager.GetBooks()
}

func (api *Service) getHook() hooks.Hook {
	return api.hook
}

func (api *Service) SetBaseURL(baseURL string) {
	config := api.config()
	config.SetBaseURL(baseURL)
}

func (api *Service) SetTimeout(timeout time.Duration) {
	config := api.config()
	config.SetTimeout(timeout)
}

func (api *Service) SetAPIKey(apiKey string) {
	config := api.config()
	config.SetAPIKey(apiKey)
}

// SetFetchAListOfBooksConfig sets method-level configuration for FetchAListOfBooks.
// Options are applied to every future call to FetchAListOfBooks and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetFetchAListOfBooksConfig(opts ...e2elibrarysdkconfig.RequestOption) *Service {
	api.fetchAListOfBooksConfig = opts
	return api
}

// SetCreateANewBookConfig sets method-level configuration for CreateANewBook.
// Options are applied to every future call to CreateANewBook and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreateANewBookConfig(opts ...e2elibrarysdkconfig.RequestOption) *Service {
	api.createANewBookConfig = opts
	return api
}

// SetVerifyTheBookExistsConfig sets method-level configuration for VerifyTheBookExists.
// Options are applied to every future call to VerifyTheBookExists and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetVerifyTheBookExistsConfig(opts ...e2elibrarysdkconfig.RequestOption) *Service {
	api.verifyTheBookExistsConfig = opts
	return api
}

// SetCheckoutNewBookConfig sets method-level configuration for CheckoutNewBook.
// Options are applied to every future call to CheckoutNewBook and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCheckoutNewBookConfig(opts ...e2elibrarysdkconfig.RequestOption) *Service {
	api.checkoutNewBookConfig = opts
	return api
}

func (api *Service) FetchAListOfBooks(ctx context.Context, params FetchAListOfBooksRequestParams, opts ...e2elibrarysdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.fetchAListOfBooksConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/books").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewE2eLibrarySDKError[[]byte](err)
	}

	return resp.Data, nil
}

func (api *Service) CreateANewBook(ctx context.Context, createANewBookRequest CreateANewBookRequest, params CreateANewBookRequestParams, opts ...e2elibrarysdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.createANewBookConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/books").
		WithConfig(config).
		WithBody(createANewBookRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewE2eLibrarySDKError[[]byte](err)
	}

	return resp.Data, nil
}

func (api *Service) VerifyTheBookExists(ctx context.Context, id string, params VerifyTheBookExistsRequestParams, opts ...e2elibrarysdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.verifyTheBookExistsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/books/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewE2eLibrarySDKError[[]byte](err)
	}

	return resp.Data, nil
}

func (api *Service) CheckoutNewBook(ctx context.Context, id string, checkoutNewBookRequest CheckoutNewBookRequest, params CheckoutNewBookRequestParams, opts ...e2elibrarysdkconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.checkoutNewBookConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PATCH").
		WithPath("/books/{id}").
		WithConfig(config).
		WithBody(checkoutNewBookRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("id", id).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewE2eLibrarySDKError[[]byte](err)
	}

	return resp.Data, nil
}
