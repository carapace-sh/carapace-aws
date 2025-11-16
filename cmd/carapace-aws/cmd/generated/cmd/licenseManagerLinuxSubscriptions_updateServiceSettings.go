package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerLinuxSubscriptions_updateServiceSettingsCmd = &cobra.Command{
	Use:   "update-service-settings",
	Short: "Updates the service settings for Linux subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerLinuxSubscriptions_updateServiceSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerLinuxSubscriptions_updateServiceSettingsCmd).Standalone()

		licenseManagerLinuxSubscriptions_updateServiceSettingsCmd.Flags().Bool("allow-update", false, "Describes if updates are allowed to the service settings for Linux subscriptions.")
		licenseManagerLinuxSubscriptions_updateServiceSettingsCmd.Flags().String("linux-subscriptions-discovery", "", "Describes if the discovery of Linux subscriptions is enabled.")
		licenseManagerLinuxSubscriptions_updateServiceSettingsCmd.Flags().String("linux-subscriptions-discovery-settings", "", "The settings defined for Linux subscriptions discovery.")
		licenseManagerLinuxSubscriptions_updateServiceSettingsCmd.Flags().Bool("no-allow-update", false, "Describes if updates are allowed to the service settings for Linux subscriptions.")
		licenseManagerLinuxSubscriptions_updateServiceSettingsCmd.MarkFlagRequired("linux-subscriptions-discovery")
		licenseManagerLinuxSubscriptions_updateServiceSettingsCmd.MarkFlagRequired("linux-subscriptions-discovery-settings")
		licenseManagerLinuxSubscriptions_updateServiceSettingsCmd.Flag("no-allow-update").Hidden = true
	})
	licenseManagerLinuxSubscriptionsCmd.AddCommand(licenseManagerLinuxSubscriptions_updateServiceSettingsCmd)
}
