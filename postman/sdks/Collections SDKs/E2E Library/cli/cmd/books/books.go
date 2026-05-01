package books

import (
	"example.com/e2e-library/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "books",
}

func init() {
	Cmd.AddCommand(fetchAListOfBooksCmd)
	Cmd.AddCommand(createANewBookCmd)
	Cmd.AddCommand(verifyTheBookExistsCmd)
	Cmd.AddCommand(checkoutNewBookCmd)
	root.AddCommand(Cmd)
}
