package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listQueueEnvironmentsCmd = &cobra.Command{
	Use:   "list-queue-environments",
	Short: "Lists queue environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listQueueEnvironmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listQueueEnvironmentsCmd).Standalone()

		deadline_listQueueEnvironmentsCmd.Flags().String("farm-id", "", "The farm ID for the queue environment list.")
		deadline_listQueueEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listQueueEnvironmentsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listQueueEnvironmentsCmd.Flags().String("queue-id", "", "The queue ID for the queue environment list.")
		deadline_listQueueEnvironmentsCmd.MarkFlagRequired("farm-id")
		deadline_listQueueEnvironmentsCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_listQueueEnvironmentsCmd)
}
