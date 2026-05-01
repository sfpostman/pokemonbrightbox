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

var checkoutNewBookCmd = &cobra.Command{
	Use:  "checkout-new-book",
	Long: "Body schema:\n  {\n    \"checkedOut\": true\n  }\n\nExamples:\n  --body '{\"checkedOut\":true}'\n  --body-file ./body.json",
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return err
		}
		xMockResponseCode, err := cmd.Flags().GetString("x-mock-response-code")
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

		var requestBody books.CheckoutNewBookRequest
		if len(bodyContent) > 0 {
			if err := json.Unmarshal(bodyContent, &requestBody); err != nil {
				return err
			}
		}

		params := books.CheckoutNewBookRequestParams{}
		xMockResponseCodeParam := sdkutil.Nullable[string]{Value: xMockResponseCode}
		params.XMockResponseCode = &xMockResponseCodeParam
		acceptParam := sdkutil.Nullable[string]{Value: accept}
		params.Accept = &acceptParam
		client := root.CreateSdkClient()
		response, err := client.Books.CheckoutNewBook(cmd.Context(), id, requestBody, params)
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
	checkoutNewBookCmd.Flags().StringP("id", "", "", "")
	_ = checkoutNewBookCmd.MarkFlagRequired("id")
	checkoutNewBookCmd.Flags().StringP("x-mock-response-code", "", "", "")
	_ = checkoutNewBookCmd.MarkFlagRequired("x-mock-response-code")
	checkoutNewBookCmd.Flags().StringP("accept", "", "", "")
	_ = checkoutNewBookCmd.MarkFlagRequired("accept")
	checkoutNewBookCmd.Flags().String("body", "", "Request body as inline JSON, e.g. '{\"checkedOut\":true}'")
	checkoutNewBookCmd.Flags().String("body-file", "", "Path to a file containing the request body")
	checkoutNewBookCmd.MarkFlagsMutuallyExclusive("body", "body-file")
}
