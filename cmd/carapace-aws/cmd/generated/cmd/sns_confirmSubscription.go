package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_confirmSubscriptionCmd = &cobra.Command{
	Use:   "confirm-subscription",
	Short: "Verifies an endpoint owner's intent to receive messages by validating the token sent to the endpoint by an earlier `Subscribe` action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_confirmSubscriptionCmd).Standalone()

	sns_confirmSubscriptionCmd.Flags().String("authenticate-on-unsubscribe", "", "Disallows unauthenticated unsubscribes of the subscription.")
	sns_confirmSubscriptionCmd.Flags().String("token", "", "Short-lived token sent to an endpoint during the `Subscribe` action.")
	sns_confirmSubscriptionCmd.Flags().String("topic-arn", "", "The ARN of the topic for which you wish to confirm a subscription.")
	sns_confirmSubscriptionCmd.MarkFlagRequired("token")
	sns_confirmSubscriptionCmd.MarkFlagRequired("topic-arn")
	snsCmd.AddCommand(sns_confirmSubscriptionCmd)
}
