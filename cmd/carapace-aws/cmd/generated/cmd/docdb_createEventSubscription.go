package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_createEventSubscriptionCmd = &cobra.Command{
	Use:   "create-event-subscription",
	Short: "Creates an Amazon DocumentDB event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_createEventSubscriptionCmd).Standalone()

	docdb_createEventSubscriptionCmd.Flags().String("enabled", "", "A Boolean value; set to `true` to activate the subscription, set to `false` to create the subscription but not active it.")
	docdb_createEventSubscriptionCmd.Flags().String("event-categories", "", "A list of event categories for a `SourceType` that you want to subscribe to.")
	docdb_createEventSubscriptionCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) of the SNS topic created for event notification.")
	docdb_createEventSubscriptionCmd.Flags().String("source-ids", "", "The list of identifiers of the event sources for which events are returned.")
	docdb_createEventSubscriptionCmd.Flags().String("source-type", "", "The type of source that is generating the events.")
	docdb_createEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the subscription.")
	docdb_createEventSubscriptionCmd.Flags().String("tags", "", "The tags to be assigned to the event subscription.")
	docdb_createEventSubscriptionCmd.MarkFlagRequired("sns-topic-arn")
	docdb_createEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	docdbCmd.AddCommand(docdb_createEventSubscriptionCmd)
}
