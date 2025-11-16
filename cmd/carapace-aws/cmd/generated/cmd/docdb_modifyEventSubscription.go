package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_modifyEventSubscriptionCmd = &cobra.Command{
	Use:   "modify-event-subscription",
	Short: "Modifies an existing Amazon DocumentDB event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_modifyEventSubscriptionCmd).Standalone()

	docdb_modifyEventSubscriptionCmd.Flags().String("enabled", "", "A Boolean value; set to `true` to activate the subscription.")
	docdb_modifyEventSubscriptionCmd.Flags().String("event-categories", "", "A list of event categories for a `SourceType` that you want to subscribe to.")
	docdb_modifyEventSubscriptionCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) of the SNS topic created for event notification.")
	docdb_modifyEventSubscriptionCmd.Flags().String("source-type", "", "The type of source that is generating the events.")
	docdb_modifyEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the Amazon DocumentDB event notification subscription.")
	docdb_modifyEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	docdbCmd.AddCommand(docdb_modifyEventSubscriptionCmd)
}
