package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerLinuxSubscriptions_getServiceSettingsCmd = &cobra.Command{
	Use:   "get-service-settings",
	Short: "Lists the Linux subscriptions service settings for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerLinuxSubscriptions_getServiceSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerLinuxSubscriptions_getServiceSettingsCmd).Standalone()

	})
	licenseManagerLinuxSubscriptionsCmd.AddCommand(licenseManagerLinuxSubscriptions_getServiceSettingsCmd)
}
