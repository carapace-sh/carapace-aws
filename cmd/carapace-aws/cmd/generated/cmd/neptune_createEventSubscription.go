package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_createEventSubscriptionCmd = &cobra.Command{
	Use:   "create-event-subscription",
	Short: "Creates an event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_createEventSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_createEventSubscriptionCmd).Standalone()

		neptune_createEventSubscriptionCmd.Flags().String("enabled", "", "A Boolean value; set to **true** to activate the subscription, set to **false** to create the subscription but not active it.")
		neptune_createEventSubscriptionCmd.Flags().String("event-categories", "", "A list of event categories for a SourceType that you want to subscribe to.")
		neptune_createEventSubscriptionCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) of the SNS topic created for event notification.")
		neptune_createEventSubscriptionCmd.Flags().String("source-ids", "", "The list of identifiers of the event sources for which events are returned.")
		neptune_createEventSubscriptionCmd.Flags().String("source-type", "", "The type of source that is generating the events.")
		neptune_createEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the subscription.")
		neptune_createEventSubscriptionCmd.Flags().String("tags", "", "The tags to be applied to the new event subscription.")
		neptune_createEventSubscriptionCmd.MarkFlagRequired("sns-topic-arn")
		neptune_createEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	})
	neptuneCmd.AddCommand(neptune_createEventSubscriptionCmd)
}
