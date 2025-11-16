package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_updateIdentityProviderSettingsCmd = &cobra.Command{
	Use:   "update-identity-provider-settings",
	Short: "Updates additional product configuration settings for the registered identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_updateIdentityProviderSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerUserSubscriptions_updateIdentityProviderSettingsCmd).Standalone()

		licenseManagerUserSubscriptions_updateIdentityProviderSettingsCmd.Flags().String("identity-provider", "", "")
		licenseManagerUserSubscriptions_updateIdentityProviderSettingsCmd.Flags().String("identity-provider-arn", "", "The Amazon Resource Name (ARN) of the identity provider to update.")
		licenseManagerUserSubscriptions_updateIdentityProviderSettingsCmd.Flags().String("product", "", "The name of the user-based subscription product.")
		licenseManagerUserSubscriptions_updateIdentityProviderSettingsCmd.Flags().String("update-settings", "", "Updates the registered identity provider’s product related configuration settings.")
		licenseManagerUserSubscriptions_updateIdentityProviderSettingsCmd.MarkFlagRequired("update-settings")
	})
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_updateIdentityProviderSettingsCmd)
}
