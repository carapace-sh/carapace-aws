package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerLinuxSubscriptions_listRegisteredSubscriptionProvidersCmd = &cobra.Command{
	Use:   "list-registered-subscription-providers",
	Short: "List Bring Your Own License (BYOL) subscription registration resources for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerLinuxSubscriptions_listRegisteredSubscriptionProvidersCmd).Standalone()

	licenseManagerLinuxSubscriptions_listRegisteredSubscriptionProvidersCmd.Flags().String("max-results", "", "The maximum items to return in a request.")
	licenseManagerLinuxSubscriptions_listRegisteredSubscriptionProvidersCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	licenseManagerLinuxSubscriptions_listRegisteredSubscriptionProvidersCmd.Flags().String("subscription-provider-sources", "", "To filter your results, specify which subscription providers to return in the list.")
	licenseManagerLinuxSubscriptionsCmd.AddCommand(licenseManagerLinuxSubscriptions_listRegisteredSubscriptionProvidersCmd)
}
