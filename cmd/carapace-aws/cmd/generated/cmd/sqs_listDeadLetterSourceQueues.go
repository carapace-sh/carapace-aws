package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_listDeadLetterSourceQueuesCmd = &cobra.Command{
	Use:   "list-dead-letter-source-queues",
	Short: "Returns a list of your queues that have the `RedrivePolicy` queue attribute configured with a dead-letter queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_listDeadLetterSourceQueuesCmd).Standalone()

	sqs_listDeadLetterSourceQueuesCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
	sqs_listDeadLetterSourceQueuesCmd.Flags().String("next-token", "", "Pagination token to request the next set of results.")
	sqs_listDeadLetterSourceQueuesCmd.Flags().String("queue-url", "", "The URL of a dead-letter queue.")
	sqs_listDeadLetterSourceQueuesCmd.MarkFlagRequired("queue-url")
	sqsCmd.AddCommand(sqs_listDeadLetterSourceQueuesCmd)
}
