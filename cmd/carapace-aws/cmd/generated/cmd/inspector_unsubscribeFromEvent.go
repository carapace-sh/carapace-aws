package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_unsubscribeFromEventCmd = &cobra.Command{
	Use:   "unsubscribe-from-event",
	Short: "Disables the process of sending Amazon Simple Notification Service (SNS) notifications about a specified event to a specified SNS topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_unsubscribeFromEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_unsubscribeFromEventCmd).Standalone()

		inspector_unsubscribeFromEventCmd.Flags().String("event", "", "The event for which you want to stop receiving SNS notifications.")
		inspector_unsubscribeFromEventCmd.Flags().String("resource-arn", "", "The ARN of the assessment template that is used during the event for which you want to stop receiving SNS notifications.")
		inspector_unsubscribeFromEventCmd.Flags().String("topic-arn", "", "The ARN of the SNS topic to which SNS notifications are sent.")
		inspector_unsubscribeFromEventCmd.MarkFlagRequired("event")
		inspector_unsubscribeFromEventCmd.MarkFlagRequired("resource-arn")
		inspector_unsubscribeFromEventCmd.MarkFlagRequired("topic-arn")
	})
	inspectorCmd.AddCommand(inspector_unsubscribeFromEventCmd)
}
