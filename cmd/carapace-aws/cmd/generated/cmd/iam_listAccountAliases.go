package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listAccountAliasesCmd = &cobra.Command{
	Use:   "list-account-aliases",
	Short: "Lists the account alias associated with the Amazon Web Services account (Note: you can have only one).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listAccountAliasesCmd).Standalone()

	iam_listAccountAliasesCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_listAccountAliasesCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iamCmd.AddCommand(iam_listAccountAliasesCmd)
}
