package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listQueueLimitAssociationsCmd = &cobra.Command{
	Use:   "list-queue-limit-associations",
	Short: "Gets a list of the associations between queues and limits defined in a farm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listQueueLimitAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listQueueLimitAssociationsCmd).Standalone()

		deadline_listQueueLimitAssociationsCmd.Flags().String("farm-id", "", "The unique identifier of the farm that contains the limits and associations.")
		deadline_listQueueLimitAssociationsCmd.Flags().String("limit-id", "", "Specifies that the operation should return only the queue limit associations for the specified limit.")
		deadline_listQueueLimitAssociationsCmd.Flags().String("max-results", "", "The maximum number of associations to return in each page of results.")
		deadline_listQueueLimitAssociationsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listQueueLimitAssociationsCmd.Flags().String("queue-id", "", "Specifies that the operation should return only the queue limit associations for the specified queue.")
		deadline_listQueueLimitAssociationsCmd.MarkFlagRequired("farm-id")
	})
	deadlineCmd.AddCommand(deadline_listQueueLimitAssociationsCmd)
}
