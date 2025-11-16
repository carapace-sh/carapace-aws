package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_modifyEventSubscriptionCmd = &cobra.Command{
	Use:   "modify-event-subscription",
	Short: "Modifies an existing event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_modifyEventSubscriptionCmd).Standalone()

	neptune_modifyEventSubscriptionCmd.Flags().String("enabled", "", "A Boolean value; set to **true** to activate the subscription.")
	neptune_modifyEventSubscriptionCmd.Flags().String("event-categories", "", "A list of event categories for a SourceType that you want to subscribe to.")
	neptune_modifyEventSubscriptionCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) of the SNS topic created for event notification.")
	neptune_modifyEventSubscriptionCmd.Flags().String("source-type", "", "The type of source that is generating the events.")
	neptune_modifyEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the event notification subscription.")
	neptune_modifyEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	neptuneCmd.AddCommand(neptune_modifyEventSubscriptionCmd)
}
