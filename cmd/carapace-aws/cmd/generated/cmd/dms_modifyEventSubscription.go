package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_modifyEventSubscriptionCmd = &cobra.Command{
	Use:   "modify-event-subscription",
	Short: "Modifies an existing DMS event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_modifyEventSubscriptionCmd).Standalone()

	dms_modifyEventSubscriptionCmd.Flags().String("enabled", "", "A Boolean value; set to **true** to activate the subscription.")
	dms_modifyEventSubscriptionCmd.Flags().String("event-categories", "", "A list of event categories for a source type that you want to subscribe to.")
	dms_modifyEventSubscriptionCmd.Flags().String("sns-topic-arn", "", "The Amazon Resource Name (ARN) of the Amazon SNS topic created for event notification.")
	dms_modifyEventSubscriptionCmd.Flags().String("source-type", "", "The type of DMS resource that generates the events you want to subscribe to.")
	dms_modifyEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the DMS event notification subscription to be modified.")
	dms_modifyEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	dmsCmd.AddCommand(dms_modifyEventSubscriptionCmd)
}
