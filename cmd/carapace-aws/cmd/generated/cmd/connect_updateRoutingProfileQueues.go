package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateRoutingProfileQueuesCmd = &cobra.Command{
	Use:   "update-routing-profile-queues",
	Short: "Updates the properties associated with a set of queues for a routing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateRoutingProfileQueuesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateRoutingProfileQueuesCmd).Standalone()

		connect_updateRoutingProfileQueuesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateRoutingProfileQueuesCmd.Flags().String("queue-configs", "", "The queues to be updated for this routing profile.")
		connect_updateRoutingProfileQueuesCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile.")
		connect_updateRoutingProfileQueuesCmd.MarkFlagRequired("instance-id")
		connect_updateRoutingProfileQueuesCmd.MarkFlagRequired("queue-configs")
		connect_updateRoutingProfileQueuesCmd.MarkFlagRequired("routing-profile-id")
	})
	connectCmd.AddCommand(connect_updateRoutingProfileQueuesCmd)
}
