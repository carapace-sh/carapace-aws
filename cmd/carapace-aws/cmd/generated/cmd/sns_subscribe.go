package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribes an endpoint to an Amazon SNS topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_subscribeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_subscribeCmd).Standalone()

		sns_subscribeCmd.Flags().String("attributes", "", "A map of attributes with their corresponding values.")
		sns_subscribeCmd.Flags().String("endpoint", "", "The endpoint that you want to receive notifications.")
		sns_subscribeCmd.Flags().String("protocol", "", "The protocol that you want to use.")
		sns_subscribeCmd.Flags().String("return-subscription-arn", "", "Sets whether the response from the `Subscribe` request includes the subscription ARN, even if the subscription is not yet confirmed.")
		sns_subscribeCmd.Flags().String("topic-arn", "", "The ARN of the topic you want to subscribe to.")
		sns_subscribeCmd.MarkFlagRequired("protocol")
		sns_subscribeCmd.MarkFlagRequired("topic-arn")
	})
	snsCmd.AddCommand(sns_subscribeCmd)
}
