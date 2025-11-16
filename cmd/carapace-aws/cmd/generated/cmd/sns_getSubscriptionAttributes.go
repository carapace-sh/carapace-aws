package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_getSubscriptionAttributesCmd = &cobra.Command{
	Use:   "get-subscription-attributes",
	Short: "Returns all of the properties of a subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_getSubscriptionAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_getSubscriptionAttributesCmd).Standalone()

		sns_getSubscriptionAttributesCmd.Flags().String("subscription-arn", "", "The ARN of the subscription whose properties you want to get.")
		sns_getSubscriptionAttributesCmd.MarkFlagRequired("subscription-arn")
	})
	snsCmd.AddCommand(sns_getSubscriptionAttributesCmd)
}
