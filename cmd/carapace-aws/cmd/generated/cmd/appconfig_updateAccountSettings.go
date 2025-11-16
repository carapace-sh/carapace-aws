package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_updateAccountSettingsCmd = &cobra.Command{
	Use:   "update-account-settings",
	Short: "Updates the value of the `DeletionProtection` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_updateAccountSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_updateAccountSettingsCmd).Standalone()

		appconfig_updateAccountSettingsCmd.Flags().String("deletion-protection", "", "A parameter to configure deletion protection.")
	})
	appconfigCmd.AddCommand(appconfig_updateAccountSettingsCmd)
}
