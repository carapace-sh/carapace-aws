package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_registerIdentityProviderCmd = &cobra.Command{
	Use:   "register-identity-provider",
	Short: "Registers an identity provider for user-based subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_registerIdentityProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerUserSubscriptions_registerIdentityProviderCmd).Standalone()

		licenseManagerUserSubscriptions_registerIdentityProviderCmd.Flags().String("identity-provider", "", "An object that specifies details for the identity provider to register.")
		licenseManagerUserSubscriptions_registerIdentityProviderCmd.Flags().String("product", "", "The name of the user-based subscription product.")
		licenseManagerUserSubscriptions_registerIdentityProviderCmd.Flags().String("settings", "", "The registered identity provider’s product related configuration settings such as the subnets to provision VPC endpoints.")
		licenseManagerUserSubscriptions_registerIdentityProviderCmd.Flags().String("tags", "", "The tags that apply to the identity provider's registration.")
		licenseManagerUserSubscriptions_registerIdentityProviderCmd.MarkFlagRequired("identity-provider")
		licenseManagerUserSubscriptions_registerIdentityProviderCmd.MarkFlagRequired("product")
	})
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_registerIdentityProviderCmd)
}
