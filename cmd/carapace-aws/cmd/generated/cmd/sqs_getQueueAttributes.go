package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_getQueueAttributesCmd = &cobra.Command{
	Use:   "get-queue-attributes",
	Short: "Gets attributes for the specified queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_getQueueAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sqs_getQueueAttributesCmd).Standalone()

		sqs_getQueueAttributesCmd.Flags().String("attribute-names", "", "A list of attributes for which to retrieve information.")
		sqs_getQueueAttributesCmd.Flags().String("queue-url", "", "The URL of the Amazon SQS queue whose attribute information is retrieved.")
		sqs_getQueueAttributesCmd.MarkFlagRequired("queue-url")
	})
	sqsCmd.AddCommand(sqs_getQueueAttributesCmd)
}
