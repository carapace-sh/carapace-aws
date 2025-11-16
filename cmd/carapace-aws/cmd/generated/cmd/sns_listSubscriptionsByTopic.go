package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_listSubscriptionsByTopicCmd = &cobra.Command{
	Use:   "list-subscriptions-by-topic",
	Short: "Returns a list of the subscriptions to a specific topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_listSubscriptionsByTopicCmd).Standalone()

	sns_listSubscriptionsByTopicCmd.Flags().String("next-token", "", "Token returned by the previous `ListSubscriptionsByTopic` request.")
	sns_listSubscriptionsByTopicCmd.Flags().String("topic-arn", "", "The ARN of the topic for which you wish to find subscriptions.")
	sns_listSubscriptionsByTopicCmd.MarkFlagRequired("topic-arn")
	snsCmd.AddCommand(sns_listSubscriptionsByTopicCmd)
}
