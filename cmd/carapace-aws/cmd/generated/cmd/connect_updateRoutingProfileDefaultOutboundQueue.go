package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateRoutingProfileDefaultOutboundQueueCmd = &cobra.Command{
	Use:   "update-routing-profile-default-outbound-queue",
	Short: "Updates the default outbound queue of a routing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateRoutingProfileDefaultOutboundQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateRoutingProfileDefaultOutboundQueueCmd).Standalone()

		connect_updateRoutingProfileDefaultOutboundQueueCmd.Flags().String("default-outbound-queue-id", "", "The identifier for the default outbound queue.")
		connect_updateRoutingProfileDefaultOutboundQueueCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateRoutingProfileDefaultOutboundQueueCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile.")
		connect_updateRoutingProfileDefaultOutboundQueueCmd.MarkFlagRequired("default-outbound-queue-id")
		connect_updateRoutingProfileDefaultOutboundQueueCmd.MarkFlagRequired("instance-id")
		connect_updateRoutingProfileDefaultOutboundQueueCmd.MarkFlagRequired("routing-profile-id")
	})
	connectCmd.AddCommand(connect_updateRoutingProfileDefaultOutboundQueueCmd)
}
