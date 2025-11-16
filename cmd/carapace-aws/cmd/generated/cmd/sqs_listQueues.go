package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqs_listQueuesCmd = &cobra.Command{
	Use:   "list-queues",
	Short: "Returns a list of your queues in the current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqs_listQueuesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sqs_listQueuesCmd).Standalone()

		sqs_listQueuesCmd.Flags().String("max-results", "", "Maximum number of results to include in the response.")
		sqs_listQueuesCmd.Flags().String("next-token", "", "Pagination token to request the next set of results.")
		sqs_listQueuesCmd.Flags().String("queue-name-prefix", "", "A string to use for filtering the list results.")
	})
	sqsCmd.AddCommand(sqs_listQueuesCmd)
}
