package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_setQueueAttributesCmd = &cobra.Command{
	Use:   "set-queue-attributes",
	Short: "Sets the value of one or more queue attributes, like a policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_setQueueAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sqs_setQueueAttributesCmd).Standalone()

		sqs_setQueueAttributesCmd.Flags().String("attributes", "", "A map of attributes to set.")
		sqs_setQueueAttributesCmd.Flags().String("queue-url", "", "The URL of the Amazon SQS queue whose attributes are set.")
		sqs_setQueueAttributesCmd.MarkFlagRequired("attributes")
		sqs_setQueueAttributesCmd.MarkFlagRequired("queue-url")
	})
	sqsCmd.AddCommand(sqs_setQueueAttributesCmd)
}
