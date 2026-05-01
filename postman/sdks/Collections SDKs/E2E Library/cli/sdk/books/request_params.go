package books

import (
	"example.com/e2e-library/sdk/param"
)

// FetchAListOfBooksRequestParams holds the optional parameters for the API request.
type FetchAListOfBooksRequestParams struct {
	XMockResponseCode *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"x-mock-response-code" required:"true"`
	Accept            *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

// CreateANewBookRequestParams holds the optional parameters for the API request.
type CreateANewBookRequestParams struct {
	XMockResponseCode *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"x-mock-response-code" required:"true"`
	XMockResponseName *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"x-mock-response-name" required:"true"`
	Accept            *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

// VerifyTheBookExistsRequestParams holds the optional parameters for the API request.
type VerifyTheBookExistsRequestParams struct {
	XMockResponseCode *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"x-mock-response-code" required:"true"`
	Accept            *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}

// CheckoutNewBookRequestParams holds the optional parameters for the API request.
type CheckoutNewBookRequestParams struct {
	XMockResponseCode *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"x-mock-response-code" required:"true"`
	Accept            *param.Nullable[string] `explode:"false" serializationStyle:"simple" headerParam:"Accept" required:"true"`
}
