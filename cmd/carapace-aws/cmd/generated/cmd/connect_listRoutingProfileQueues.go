package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listRoutingProfileQueuesCmd = &cobra.Command{
	Use:   "list-routing-profile-queues",
	Short: "Lists the queues associated with a routing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listRoutingProfileQueuesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listRoutingProfileQueuesCmd).Standalone()

		connect_listRoutingProfileQueuesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listRoutingProfileQueuesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listRoutingProfileQueuesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listRoutingProfileQueuesCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile.")
		connect_listRoutingProfileQueuesCmd.MarkFlagRequired("instance-id")
		connect_listRoutingProfileQueuesCmd.MarkFlagRequired("routing-profile-id")
	})
	connectCmd.AddCommand(connect_listRoutingProfileQueuesCmd)
}
