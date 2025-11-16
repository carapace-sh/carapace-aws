package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerLinuxSubscriptions_deregisterSubscriptionProviderCmd = &cobra.Command{
	Use:   "deregister-subscription-provider",
	Short: "Remove a third-party subscription provider from the Bring Your Own License (BYOL) subscriptions registered to your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerLinuxSubscriptions_deregisterSubscriptionProviderCmd).Standalone()

	licenseManagerLinuxSubscriptions_deregisterSubscriptionProviderCmd.Flags().String("subscription-provider-arn", "", "The Amazon Resource Name (ARN) of the subscription provider resource to deregister.")
	licenseManagerLinuxSubscriptions_deregisterSubscriptionProviderCmd.MarkFlagRequired("subscription-provider-arn")
	licenseManagerLinuxSubscriptionsCmd.AddCommand(licenseManagerLinuxSubscriptions_deregisterSubscriptionProviderCmd)
}
