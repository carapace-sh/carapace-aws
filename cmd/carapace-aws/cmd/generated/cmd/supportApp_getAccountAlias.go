package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportApp_getAccountAliasCmd = &cobra.Command{
	Use:   "get-account-alias",
	Short: "Retrieves the alias from an Amazon Web Services account ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportApp_getAccountAliasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supportApp_getAccountAliasCmd).Standalone()

	})
	supportAppCmd.AddCommand(supportApp_getAccountAliasCmd)
}
