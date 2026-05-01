# Books

A list of all methods in the `Books` service. Click on the method name to view detailed information about that method.

| Methods                                     | Description |
| :------------------------------------------ | :---------- |
| [FetchAListOfBooks](#fetchalistofbooks)     |             |
| [CreateANewBook](#createanewbook)           |             |
| [VerifyTheBookExists](#verifythebookexists) |             |
| [CheckoutNewBook](#checkoutnewbook)         |             |

## FetchAListOfBooks

- HTTP Method: `GET`
- Endpoint: `/books`

**Parameters**

| Name   | Type                           | Required | Description                   |
| :----- | :----------------------------- | :------- | :---------------------------- |
| ctx    | Context                        | ✅       | Default go language context   |
| params | FetchAListOfBooksRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
e2e-library books fetch-a-list-of-books --x-mock-response-code "500" --accept "application/json"
```

## CreateANewBook

- HTTP Method: `POST`
- Endpoint: `/books`

**Parameters**

| Name                  | Type                        | Required | Description                   |
| :-------------------- | :-------------------------- | :------- | :---------------------------- |
| ctx                   | Context                     | ✅       | Default go language context   |
| createANewBookRequest | CreateANewBookRequest       | ✅       |                               |
| params                | CreateANewBookRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
e2e-library books create-a-new-book --x-mock-response-code "201" --x-mock-response-name "Dynamic Variables Mock Demo" --accept "application/json" --body '{"title":"{{tempBookTitle}}","author":"{{$randomFirstName}} {{$randomLastName}}","genre":"fiction","yearPublished":"1967"}'
```

## VerifyTheBookExists

- HTTP Method: `GET`
- Endpoint: `/books/{id}`

**Parameters**

| Name   | Type                             | Required | Description                   |
| :----- | :------------------------------- | :------- | :---------------------------- |
| ctx    | Context                          | ✅       | Default go language context   |
| id     | string                           | ✅       |                               |
| params | VerifyTheBookExistsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
e2e-library books verify-the-book-exists --id "id" --x-mock-response-code "200" --accept "application/json"
```

## CheckoutNewBook

- HTTP Method: `PATCH`
- Endpoint: `/books/{id}`

**Parameters**

| Name                   | Type                         | Required | Description                   |
| :--------------------- | :--------------------------- | :------- | :---------------------------- |
| ctx                    | Context                      | ✅       | Default go language context   |
| id                     | string                       | ✅       |                               |
| checkoutNewBookRequest | CheckoutNewBookRequest       | ✅       |                               |
| params                 | CheckoutNewBookRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
e2e-library books checkout-new-book --id "id" --x-mock-response-code "200" --accept "application/json" --body '{"checkedOut":true}'
```
