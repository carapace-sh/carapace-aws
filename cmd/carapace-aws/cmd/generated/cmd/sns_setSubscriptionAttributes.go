package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_setSubscriptionAttributesCmd = &cobra.Command{
	Use:   "set-subscription-attributes",
	Short: "Allows a subscription owner to set an attribute of the subscription to a new value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_setSubscriptionAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_setSubscriptionAttributesCmd).Standalone()

		sns_setSubscriptionAttributesCmd.Flags().String("attribute-name", "", "A map of attributes with their corresponding values.")
		sns_setSubscriptionAttributesCmd.Flags().String("attribute-value", "", "The new value for the attribute in JSON format.")
		sns_setSubscriptionAttributesCmd.Flags().String("subscription-arn", "", "The ARN of the subscription to modify.")
		sns_setSubscriptionAttributesCmd.MarkFlagRequired("attribute-name")
		sns_setSubscriptionAttributesCmd.MarkFlagRequired("subscription-arn")
	})
	snsCmd.AddCommand(sns_setSubscriptionAttributesCmd)
}
