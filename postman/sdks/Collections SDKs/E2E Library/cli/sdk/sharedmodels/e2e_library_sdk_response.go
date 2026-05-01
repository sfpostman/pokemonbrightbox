package sharedmodels

import (
	"encoding/json"
	"example.com/e2e-library/sdk/internal/clients/rest/httptransport"
	"net/http"
)

// E2eLibrarySDKResponse is the user-facing wrapper for API responses.
// It contains the deserialized data, raw HTTP response, and metadata like headers and status code.
type E2eLibrarySDKResponse[T any] struct {
	Data     T
	Raw      *http.Response
	Metadata E2eLibrarySDKResponseMetadata
}

// E2eLibrarySDKResponseMetadata contains HTTP metadata from the API response.
// Includes status code and headers for inspection and debugging.
type E2eLibrarySDKResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

// NewE2eLibrarySDKResponse creates a new response wrapper from an internal transport response.
// Extracts data and metadata into a user-facing structure.
func NewE2eLibrarySDKResponse[T any](resp *httptransport.Response[T]) *E2eLibrarySDKResponse[T] {
	return &E2eLibrarySDKResponse[T]{
		Data: resp.Data,
		Raw:  resp.Raw,
		Metadata: E2eLibrarySDKResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

// GetData returns the deserialized response data.
func (r *E2eLibrarySDKResponse[T]) GetData() T {
	return r.Data
}

// String returns a JSON representation of the response for debugging.
// Returns an error message if JSON marshaling fails.
func (r E2eLibrarySDKResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: E2eLibrarySDKResponse to string"
	}
	return string(jsonData)
}
