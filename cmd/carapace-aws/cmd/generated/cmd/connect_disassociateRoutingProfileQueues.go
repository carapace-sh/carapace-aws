package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateRoutingProfileQueuesCmd = &cobra.Command{
	Use:   "disassociate-routing-profile-queues",
	Short: "Disassociates a set of queues from a routing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateRoutingProfileQueuesCmd).Standalone()

	connect_disassociateRoutingProfileQueuesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_disassociateRoutingProfileQueuesCmd.Flags().String("manual-assignment-queue-references", "", "The manual assignment queues to disassociate with this routing profile.")
	connect_disassociateRoutingProfileQueuesCmd.Flags().String("queue-references", "", "The queues to disassociate from this routing profile.")
	connect_disassociateRoutingProfileQueuesCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile.")
	connect_disassociateRoutingProfileQueuesCmd.MarkFlagRequired("instance-id")
	connect_disassociateRoutingProfileQueuesCmd.MarkFlagRequired("routing-profile-id")
	connectCmd.AddCommand(connect_disassociateRoutingProfileQueuesCmd)
}
