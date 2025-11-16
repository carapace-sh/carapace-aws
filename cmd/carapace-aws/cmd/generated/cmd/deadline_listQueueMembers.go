package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listQueueMembersCmd = &cobra.Command{
	Use:   "list-queue-members",
	Short: "Lists the members in a queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listQueueMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listQueueMembersCmd).Standalone()

		deadline_listQueueMembersCmd.Flags().String("farm-id", "", "The farm ID for the queue.")
		deadline_listQueueMembersCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listQueueMembersCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listQueueMembersCmd.Flags().String("queue-id", "", "The queue ID to include on the list.")
		deadline_listQueueMembersCmd.MarkFlagRequired("farm-id")
		deadline_listQueueMembersCmd.MarkFlagRequired("queue-id")
	})
	deadlineCmd.AddCommand(deadline_listQueueMembersCmd)
}
