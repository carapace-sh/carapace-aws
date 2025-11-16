package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_deregisterIdentityProviderCmd = &cobra.Command{
	Use:   "deregister-identity-provider",
	Short: "Deregisters the Active Directory identity provider from License Manager user-based subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_deregisterIdentityProviderCmd).Standalone()

	licenseManagerUserSubscriptions_deregisterIdentityProviderCmd.Flags().String("identity-provider", "", "An object that specifies details for the Active Directory identity provider.")
	licenseManagerUserSubscriptions_deregisterIdentityProviderCmd.Flags().String("identity-provider-arn", "", "The Amazon Resource Name (ARN) that identifies the identity provider to deregister.")
	licenseManagerUserSubscriptions_deregisterIdentityProviderCmd.Flags().String("product", "", "The name of the user-based subscription product.")
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_deregisterIdentityProviderCmd)
}
