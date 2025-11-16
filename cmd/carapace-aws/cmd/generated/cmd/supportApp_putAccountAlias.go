package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supportApp_putAccountAliasCmd = &cobra.Command{
	Use:   "put-account-alias",
	Short: "Creates or updates an individual alias for each Amazon Web Services account ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supportApp_putAccountAliasCmd).Standalone()

	supportApp_putAccountAliasCmd.Flags().String("account-alias", "", "An alias or short name for an Amazon Web Services account.")
	supportApp_putAccountAliasCmd.MarkFlagRequired("account-alias")
	supportAppCmd.AddCommand(supportApp_putAccountAliasCmd)
}
