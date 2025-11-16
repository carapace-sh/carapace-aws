package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_getAccountSettingsCmd = &cobra.Command{
	Use:   "get-account-settings",
	Short: "Returns information about the status of the `DeletionProtection` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_getAccountSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_getAccountSettingsCmd).Standalone()

	})
	appconfigCmd.AddCommand(appconfig_getAccountSettingsCmd)
}
