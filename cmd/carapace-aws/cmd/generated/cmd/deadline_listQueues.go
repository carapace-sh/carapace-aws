package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listQueuesCmd = &cobra.Command{
	Use:   "list-queues",
	Short: "Lists queues.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listQueuesCmd).Standalone()

	deadline_listQueuesCmd.Flags().String("farm-id", "", "The farm ID of the queue.")
	deadline_listQueuesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_listQueuesCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_listQueuesCmd.Flags().String("principal-id", "", "The principal IDs to include in the list of queues.")
	deadline_listQueuesCmd.Flags().String("status", "", "The status of the queues listed.")
	deadline_listQueuesCmd.MarkFlagRequired("farm-id")
	deadlineCmd.AddCommand(deadline_listQueuesCmd)
}
