package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_stopProductSubscriptionCmd = &cobra.Command{
	Use:   "stop-product-subscription",
	Short: "Stops a product subscription for a user with the specified identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_stopProductSubscriptionCmd).Standalone()

	licenseManagerUserSubscriptions_stopProductSubscriptionCmd.Flags().String("domain", "", "The domain name of the Active Directory that contains the user for whom to stop the product subscription.")
	licenseManagerUserSubscriptions_stopProductSubscriptionCmd.Flags().String("identity-provider", "", "An object that specifies details for the identity provider.")
	licenseManagerUserSubscriptions_stopProductSubscriptionCmd.Flags().String("product", "", "The name of the user-based subscription product.")
	licenseManagerUserSubscriptions_stopProductSubscriptionCmd.Flags().String("product-user-arn", "", "The Amazon Resource Name (ARN) of the product user.")
	licenseManagerUserSubscriptions_stopProductSubscriptionCmd.Flags().String("username", "", "The user name from the identity provider for the user.")
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_stopProductSubscriptionCmd)
}
