package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportApp_deleteAccountAliasCmd = &cobra.Command{
	Use:   "delete-account-alias",
	Short: "Deletes an alias for an Amazon Web Services account ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportApp_deleteAccountAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supportApp_deleteAccountAliasCmd).Standalone()

	})
	supportAppCmd.AddCommand(supportApp_deleteAccountAliasCmd)
}
