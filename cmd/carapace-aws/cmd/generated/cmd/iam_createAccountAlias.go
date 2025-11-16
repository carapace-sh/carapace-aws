package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createAccountAliasCmd = &cobra.Command{
	Use:   "create-account-alias",
	Short: "Creates an alias for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createAccountAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_createAccountAliasCmd).Standalone()

		iam_createAccountAliasCmd.Flags().String("account-alias", "", "The account alias to create.")
		iam_createAccountAliasCmd.MarkFlagRequired("account-alias")
	})
	iamCmd.AddCommand(iam_createAccountAliasCmd)
}
