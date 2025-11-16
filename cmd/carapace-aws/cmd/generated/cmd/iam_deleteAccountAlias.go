package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteAccountAliasCmd = &cobra.Command{
	Use:   "delete-account-alias",
	Short: "Deletes the specified Amazon Web Services account alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteAccountAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_deleteAccountAliasCmd).Standalone()

		iam_deleteAccountAliasCmd.Flags().String("account-alias", "", "The name of the account alias to delete.")
		iam_deleteAccountAliasCmd.MarkFlagRequired("account-alias")
	})
	iamCmd.AddCommand(iam_deleteAccountAliasCmd)
}
