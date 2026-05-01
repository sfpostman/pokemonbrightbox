package e2elibrarysdk

import (
	"example.com/e2e-library/sdk/e2elibrarysdkconfig"
	"example.com/e2e-library/sdk/param"
)

// The type aliases below let consumers use a single import path for the entire SDK.
// Internally the concrete types live in e2elibrarysdkconfig and param.

// Config holds all configuration parameters for the SDK client.
type Config = e2elibrarysdkconfig.Config

// RequestOption is a function that configures a single request.
type RequestOption = e2elibrarysdkconfig.RequestOption

// Environment defines the available API base URLs.
type Environment = e2elibrarysdkconfig.Environment

// RetryConfig holds all runtime-configurable retry parameters.
type RetryConfig = e2elibrarysdkconfig.RetryConfig

// NewConfig creates a Config with spec-derived defaults.
var NewConfig = e2elibrarysdkconfig.NewConfig

// NewRetryConfig returns a RetryConfig initialized with spec-derived defaults.
var NewRetryConfig = e2elibrarysdkconfig.NewRetryConfig

// WithBaseURL returns a RequestOption that overrides BaseURL for a single request.
var WithBaseURL = e2elibrarysdkconfig.WithBaseURL

// WithTimeout returns a RequestOption that overrides Timeout for a single request.
var WithTimeout = e2elibrarysdkconfig.WithTimeout

// WithAPIKey returns a RequestOption that overrides APIKey for a single request.
var WithAPIKey = e2elibrarysdkconfig.WithAPIKey

// WithRetryConfig returns a RequestOption that overrides the RetryConfig for a single request.
var WithRetryConfig = e2elibrarysdkconfig.WithRetryConfig

// Nullable returns a *param.Nullable[T] set to v — use for nullable fields with a value.
func Nullable[T any](v T) *param.Nullable[T] { return &param.Nullable[T]{Value: v} }

// Null returns a *param.Nullable[T] with IsNull set to true, signalling an explicit JSON null.
func Null[T any]() *param.Nullable[T] { return param.Null[T]() }

// Ptr returns a pointer to v — use when no type-specific helper exists.
func Ptr[T any](v T) *T { return param.Ptr(v) }

// Environment constants for the available API base URLs.
const (
	DefaultEnvironment    Environment = e2elibrarysdkconfig.DefaultEnvironment
	LibraryApiEnvironment Environment = e2elibrarysdkconfig.LibraryApiEnvironment
)
