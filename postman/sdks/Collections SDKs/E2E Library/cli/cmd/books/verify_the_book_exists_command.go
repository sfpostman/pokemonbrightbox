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

var verifyTheBookExistsCmd = &cobra.Command{
	Use: "verify-the-book-exists",
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

		params := books.VerifyTheBookExistsRequestParams{}
		xMockResponseCodeParam := sdkutil.Nullable[string]{Value: xMockResponseCode}
		params.XMockResponseCode = &xMockResponseCodeParam
		acceptParam := sdkutil.Nullable[string]{Value: accept}
		params.Accept = &acceptParam
		client := root.CreateSdkClient()
		response, err := client.Books.VerifyTheBookExists(cmd.Context(), id, params)
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
	verifyTheBookExistsCmd.Flags().StringP("id", "", "", "")
	_ = verifyTheBookExistsCmd.MarkFlagRequired("id")
	verifyTheBookExistsCmd.Flags().StringP("x-mock-response-code", "", "", "")
	_ = verifyTheBookExistsCmd.MarkFlagRequired("x-mock-response-code")
	verifyTheBookExistsCmd.Flags().StringP("accept", "", "", "")
	_ = verifyTheBookExistsCmd.MarkFlagRequired("accept")
}
