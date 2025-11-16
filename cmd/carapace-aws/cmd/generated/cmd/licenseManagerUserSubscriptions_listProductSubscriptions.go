package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_listProductSubscriptionsCmd = &cobra.Command{
	Use:   "list-product-subscriptions",
	Short: "Lists the user-based subscription products available from an identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_listProductSubscriptionsCmd).Standalone()

	licenseManagerUserSubscriptions_listProductSubscriptionsCmd.Flags().String("filters", "", "You can use the following filters to streamline results:")
	licenseManagerUserSubscriptions_listProductSubscriptionsCmd.Flags().String("identity-provider", "", "An object that specifies details for the identity provider.")
	licenseManagerUserSubscriptions_listProductSubscriptionsCmd.Flags().String("max-results", "", "The maximum number of results to return from a single request.")
	licenseManagerUserSubscriptions_listProductSubscriptionsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	licenseManagerUserSubscriptions_listProductSubscriptionsCmd.Flags().String("product", "", "The name of the user-based subscription product.")
	licenseManagerUserSubscriptions_listProductSubscriptionsCmd.MarkFlagRequired("identity-provider")
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_listProductSubscriptionsCmd)
}
