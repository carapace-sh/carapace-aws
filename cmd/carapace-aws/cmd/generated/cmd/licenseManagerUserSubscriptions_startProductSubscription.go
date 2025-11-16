package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_startProductSubscriptionCmd = &cobra.Command{
	Use:   "start-product-subscription",
	Short: "Starts a product subscription for a user with the specified identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_startProductSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerUserSubscriptions_startProductSubscriptionCmd).Standalone()

		licenseManagerUserSubscriptions_startProductSubscriptionCmd.Flags().String("domain", "", "The domain name of the Active Directory that contains the user for whom to start the product subscription.")
		licenseManagerUserSubscriptions_startProductSubscriptionCmd.Flags().String("identity-provider", "", "An object that specifies details for the identity provider.")
		licenseManagerUserSubscriptions_startProductSubscriptionCmd.Flags().String("product", "", "The name of the user-based subscription product.")
		licenseManagerUserSubscriptions_startProductSubscriptionCmd.Flags().String("tags", "", "The tags that apply to the product subscription.")
		licenseManagerUserSubscriptions_startProductSubscriptionCmd.Flags().String("username", "", "The user name from the identity provider of the user.")
		licenseManagerUserSubscriptions_startProductSubscriptionCmd.MarkFlagRequired("identity-provider")
		licenseManagerUserSubscriptions_startProductSubscriptionCmd.MarkFlagRequired("product")
		licenseManagerUserSubscriptions_startProductSubscriptionCmd.MarkFlagRequired("username")
	})
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_startProductSubscriptionCmd)
}
