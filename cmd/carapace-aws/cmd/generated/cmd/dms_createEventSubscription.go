package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_createEventSubscriptionCmd = &cobra.Command{
	Use:   "create-event-subscription",
	Short: "Creates an DMS event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_createEventSubscriptionCmd).Standalone()

	dms_createEventSubscriptionCmd.Flags().String("enabled", "", "A Boolean value; set to `true` to activate the subscription, or set to `false` to create the subscription but not activate it.")
	dms_createEventSubscriptionCmd.Flags().String("event-categories", "", "A list of event categories for a source type that you want to subscribe to.")
	dms_createEventSubscriptionCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) of the Amazon SNS topic created for event notification.")
	dms_createEventSubscriptionCmd.Flags().String("source-ids", "", "A list of identifiers for which DMS provides notification events.")
	dms_createEventSubscriptionCmd.Flags().String("source-type", "", "The type of DMS resource that generates the events.")
	dms_createEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the DMS event notification subscription.")
	dms_createEventSubscriptionCmd.Flags().String("tags", "", "One or more tags to be assigned to the event subscription.")
	dms_createEventSubscriptionCmd.MarkFlagRequired("sns-topic-arn")
	dms_createEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	dmsCmd.AddCommand(dms_createEventSubscriptionCmd)
}
