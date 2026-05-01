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

var fetchAListOfBooksCmd = &cobra.Command{
	Use: "fetch-a-list-of-books",
	RunE: func(cmd *cobra.Command, args []string) error {
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

		params := books.FetchAListOfBooksRequestParams{}
		xMockResponseCodeParam := sdkutil.Nullable[string]{Value: xMockResponseCode}
		params.XMockResponseCode = &xMockResponseCodeParam
		acceptParam := sdkutil.Nullable[string]{Value: accept}
		params.Accept = &acceptParam
		client := root.CreateSdkClient()
		response, err := client.Books.FetchAListOfBooks(cmd.Context(), params)
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
	fetchAListOfBooksCmd.Flags().StringP("x-mock-response-code", "", "", "")
	_ = fetchAListOfBooksCmd.MarkFlagRequired("x-mock-response-code")
	fetchAListOfBooksCmd.Flags().StringP("accept", "", "", "")
	_ = fetchAListOfBooksCmd.MarkFlagRequired("accept")
}
