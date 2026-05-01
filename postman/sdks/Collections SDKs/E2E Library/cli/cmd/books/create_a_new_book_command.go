package books

import (
	"encoding/json"
	"example.com/e2e-library/root"
	"example.com/e2e-library/sdk/books"
	sdkutil "example.com/e2e-library/sdk/param"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var createANewBookCmd = &cobra.Command{
	Use:  "create-a-new-book",
	Long: "Body schema:\n  {\n    \"title\": \"{{tempBookTitle}}\",\n    \"author\": \"{{$randomFirstName}} {{$randomLastName}}\",\n    \"genre\": \"fiction\",\n    \"yearPublished\": \"1967\"\n  }\n\nExamples:\n  --body '{\"title\":\"{{tempBookTitle}}\",\"author\":\"{{$randomFirstName}} {{$randomLastName}}\",\"genre\":\"fiction\",\"yearPublished\":\"1967\"}'\n  --body-file ./body.json",
	RunE: func(cmd *cobra.Command, args []string) error {
		xMockResponseCode, err := cmd.Flags().GetString("x-mock-response-code")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}
		xMockResponseName, err := cmd.Flags().GetString("x-mock-response-name")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}
		accept, err := cmd.Flags().GetString("accept")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}
		bodyStr, err := cmd.Flags().GetString("body")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to get body param: %v\n", err)
			return err
		}
		bodyFile, err := cmd.Flags().GetString("body-file")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to get body-file param: %v\n", err)
			return err
		}

		var bodyContent []byte
		if bodyFile != "" {
			bodyContent, err = os.ReadFile(bodyFile)
			if err != nil {
				return err
			}
		} else {
			bodyContent = []byte(bodyStr)
		}

		var requestBody books.CreateANewBookRequest
		if len(bodyContent) > 0 {
			if err := json.Unmarshal(bodyContent, &requestBody); err != nil {
				return err
			}
		}

		params := books.CreateANewBookRequestParams{}
		xMockResponseCodeParam := sdkutil.Nullable[string]{Value: xMockResponseCode}
		params.XMockResponseCode = &xMockResponseCodeParam
		xMockResponseNameParam := sdkutil.Nullable[string]{Value: xMockResponseName}
		params.XMockResponseName = &xMockResponseNameParam
		acceptParam := sdkutil.Nullable[string]{Value: accept}
		params.Accept = &acceptParam
		client := root.CreateSdkClient()
		response, err := client.Books.CreateANewBook(cmd.Context(), requestBody, params)
		if err != nil {
			return err
		}

		if len(response) == 0 {
			fmt.Println("[empty response]")
		} else if json.Valid(response) {
			jsonData, err := json.MarshalIndent(json.RawMessage(response), "", "  ")
			if err != nil {
				fmt.Println(string(response))
			} else {
				fmt.Println(string(jsonData))
			}
		} else {
			fmt.Println(string(response))
		}

		return nil
	},
}

func init() {
	createANewBookCmd.Flags().StringP("x-mock-response-code", "", "", "")
	_ = createANewBookCmd.MarkFlagRequired("x-mock-response-code")
	createANewBookCmd.Flags().StringP("x-mock-response-name", "", "", "")
	_ = createANewBookCmd.MarkFlagRequired("x-mock-response-name")
	createANewBookCmd.Flags().StringP("accept", "", "", "")
	_ = createANewBookCmd.MarkFlagRequired("accept")
	createANewBookCmd.Flags().String("body", "", "Request body as inline JSON, e.g. '{\"title\":\"{{tempBookTitle}}\",\"author\":\"{{$randomFirstName}} {{$randomLastName}}\",\"genre\":\"fiction\",\"yearPublished\":\"1967\"}'")
	createANewBookCmd.Flags().String("body-file", "", "Path to a file containing the request body")
	createANewBookCmd.MarkFlagsMutuallyExclusive("body", "body-file")
}
