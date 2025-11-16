package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateRoutingProfileQueuesCmd = &cobra.Command{
	Use:   "associate-routing-profile-queues",
	Short: "Associates a set of queues with a routing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateRoutingProfileQueuesCmd).Standalone()

	connect_associateRoutingProfileQueuesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_associateRoutingProfileQueuesCmd.Flags().String("manual-assignment-queue-configs", "", "The manual assignment queues to associate with this routing profile.")
	connect_associateRoutingProfileQueuesCmd.Flags().String("queue-configs", "", "The queues to associate with this routing profile.")
	connect_associateRoutingProfileQueuesCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile.")
	connect_associateRoutingProfileQueuesCmd.MarkFlagRequired("instance-id")
	connect_associateRoutingProfileQueuesCmd.MarkFlagRequired("routing-profile-id")
	connectCmd.AddCommand(connect_associateRoutingProfileQueuesCmd)
}
