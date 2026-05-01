# E2eLibrarySDK Go SDK 1.0.0

Welcome to the E2eLibrarySDK SDK documentation. This guide will help you get started with integrating and using the E2eLibrarySDK SDK in your project.

## Versions

- SDK version: `1.0.0`

## Table of Contents

- [Setup & Configuration](#setup--configuration)
  - [Supported Language Versions](#supported-language-versions)
- [Authentication](#authentication)
  - [API Key Authentication](#api-key-authentication)
- [Setting a Custom Timeout](#setting-a-custom-timeout)
- [Sample Usage](#sample-usage)
- [Services](#services)
  - [Response Wrappers](#response-wrappers)
- [Models](#models)

# Setup & Configuration

## Supported Language Versions

This SDK is compatible with the following versions: `Go >= 1.19.0`

## Authentication

### API Key Authentication

The e2e-library-sdk API uses API keys as a form of authentication. An API key is a unique identifier used to authenticate a user, developer, or a program that is calling the API.

#### Setting the API key

When you initialize the SDK, you can set the API key as follows:

```go
export E2E_LIBRARY_API_KEY="YOUR-API-KEY"
e2e-library <command>
```

If you need to set or update the API key after initializing the SDK, you can use:

```go
e2e-library config set api_key "YOUR-API-KEY"
```

## Setting a Custom Timeout

You can set a custom timeout for the SDK's HTTP requests as follows:

```go

```

# Sample Usage

Below is a comprehensive example demonstrating how to authenticate and call a simple endpoint:

```go
e2e-library books fetch-a-list-of-books --x-mock-response-code "500" --accept "application/json"

```

## Services

The SDK provides various services to interact with the API.

<details>
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                     |
| :--------------------------------------- |
| [Books](documentation/services/books.md) |

</details>

### Response Wrappers

All services use response wrappers to provide a consistent interface to return the responses from the API.

The response wrapper itself is a generic struct that contains the response data and metadata.

<details>
<summary>Below are the response wrappers used in the SDK:</summary>

#### `E2eLibrarySDKResponse[T]`

This response wrapper is used to return the response data from the API. It contains the following fields:

| Name     | Type                            | Description                                 |
| :------- | :------------------------------ | :------------------------------------------ |
| Data     | `T`                             | The body of the API response                |
| Metadata | `E2eLibrarySDKResponseMetadata` | Status code and headers returned by the API |

#### `E2eLibrarySDKError[T]`

This response wrapper is used to return an error. It contains the following fields:

| Name     | Type                         | Description                                                       |
| :------- | :--------------------------- | :---------------------------------------------------------------- |
| Err      | `error`                      | The error that occurred                                           |
| Data     | `*T`                         | The deserialized error response data (nil if unmarshaling failed) |
| Body     | `[]byte`                     | The raw body of the API response                                  |
| Metadata | `E2eLibrarySDKErrorMetadata` | Status code and headers returned by the API                       |

#### `E2eLibrarySDKResponseMetadata`

This struct is shared by both response wrappers and contains the following fields:

| Name       | Type                | Description                                      |
| :--------- | :------------------ | :----------------------------------------------- |
| Headers    | `map[string]string` | A map containing the headers returned by the API |
| StatusCode | `int`               | The status code returned by the API              |

</details>

## Models

The SDK includes several models that represent the data structures used in API requests and responses. These models help in organizing and managing the data efficiently.

<details>
<summary>Below is a list of all available models with links to their detailed documentation:</summary>

| Name                                                                        | Description |
| :-------------------------------------------------------------------------- | :---------- |
| [CreateANewBookRequest](documentation/models/create_a_new_book_request.md)  |             |
| [CheckoutNewBookRequest](documentation/models/checkout_new_book_request.md) |             |

</details>
