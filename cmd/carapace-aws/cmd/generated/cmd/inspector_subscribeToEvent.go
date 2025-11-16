package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_subscribeToEventCmd = &cobra.Command{
	Use:   "subscribe-to-event",
	Short: "Enables the process of sending Amazon Simple Notification Service (SNS) notifications about a specified event to a specified SNS topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_subscribeToEventCmd).Standalone()

	inspector_subscribeToEventCmd.Flags().String("event", "", "The event for which you want to receive SNS notifications.")
	inspector_subscribeToEventCmd.Flags().String("resource-arn", "", "The ARN of the assessment template that is used during the event for which you want to receive SNS notifications.")
	inspector_subscribeToEventCmd.Flags().String("topic-arn", "", "The ARN of the SNS topic to which the SNS notifications are sent.")
	inspector_subscribeToEventCmd.MarkFlagRequired("event")
	inspector_subscribeToEventCmd.MarkFlagRequired("resource-arn")
	inspector_subscribeToEventCmd.MarkFlagRequired("topic-arn")
	inspectorCmd.AddCommand(inspector_subscribeToEventCmd)
}
