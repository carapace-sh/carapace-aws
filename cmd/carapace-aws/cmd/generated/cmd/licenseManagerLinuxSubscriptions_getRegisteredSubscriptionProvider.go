package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerLinuxSubscriptions_getRegisteredSubscriptionProviderCmd = &cobra.Command{
	Use:   "get-registered-subscription-provider",
	Short: "Get details for a Bring Your Own License (BYOL) subscription that's registered to your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerLinuxSubscriptions_getRegisteredSubscriptionProviderCmd).Standalone()

	licenseManagerLinuxSubscriptions_getRegisteredSubscriptionProviderCmd.Flags().String("subscription-provider-arn", "", "The Amazon Resource Name (ARN) of the BYOL registration resource to get details for.")
	licenseManagerLinuxSubscriptions_getRegisteredSubscriptionProviderCmd.MarkFlagRequired("subscription-provider-arn")
	licenseManagerLinuxSubscriptionsCmd.AddCommand(licenseManagerLinuxSubscriptions_getRegisteredSubscriptionProviderCmd)
}
