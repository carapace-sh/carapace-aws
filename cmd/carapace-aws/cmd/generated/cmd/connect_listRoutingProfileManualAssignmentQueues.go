package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listRoutingProfileManualAssignmentQueuesCmd = &cobra.Command{
	Use:   "list-routing-profile-manual-assignment-queues",
	Short: "Lists the manual assignment queues associated with a routing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listRoutingProfileManualAssignmentQueuesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listRoutingProfileManualAssignmentQueuesCmd).Standalone()

		connect_listRoutingProfileManualAssignmentQueuesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listRoutingProfileManualAssignmentQueuesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listRoutingProfileManualAssignmentQueuesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listRoutingProfileManualAssignmentQueuesCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile.")
		connect_listRoutingProfileManualAssignmentQueuesCmd.MarkFlagRequired("instance-id")
		connect_listRoutingProfileManualAssignmentQueuesCmd.MarkFlagRequired("routing-profile-id")
	})
	connectCmd.AddCommand(connect_listRoutingProfileManualAssignmentQueuesCmd)
}
