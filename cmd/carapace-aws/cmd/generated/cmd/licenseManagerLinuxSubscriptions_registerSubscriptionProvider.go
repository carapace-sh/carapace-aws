package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerLinuxSubscriptions_registerSubscriptionProviderCmd = &cobra.Command{
	Use:   "register-subscription-provider",
	Short: "Register the supported third-party subscription provider for your Bring Your Own License (BYOL) subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerLinuxSubscriptions_registerSubscriptionProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(licenseManagerLinuxSubscriptions_registerSubscriptionProviderCmd).Standalone()

		licenseManagerLinuxSubscriptions_registerSubscriptionProviderCmd.Flags().String("secret-arn", "", "The Amazon Resource Name (ARN) of the secret where you've stored your subscription provider's access token.")
		licenseManagerLinuxSubscriptions_registerSubscriptionProviderCmd.Flags().String("subscription-provider-source", "", "The supported Linux subscription provider to register.")
		licenseManagerLinuxSubscriptions_registerSubscriptionProviderCmd.Flags().String("tags", "", "The metadata tags to assign to your registered Linux subscription provider resource.")
		licenseManagerLinuxSubscriptions_registerSubscriptionProviderCmd.MarkFlagRequired("secret-arn")
		licenseManagerLinuxSubscriptions_registerSubscriptionProviderCmd.MarkFlagRequired("subscription-provider-source")
	})
	licenseManagerLinuxSubscriptionsCmd.AddCommand(licenseManagerLinuxSubscriptions_registerSubscriptionProviderCmd)
}
