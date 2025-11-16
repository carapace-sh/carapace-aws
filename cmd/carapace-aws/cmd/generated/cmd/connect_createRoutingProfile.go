package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createRoutingProfileCmd = &cobra.Command{
	Use:   "create-routing-profile",
	Short: "Creates a new routing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createRoutingProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_createRoutingProfileCmd).Standalone()

		connect_createRoutingProfileCmd.Flags().String("agent-availability-timer", "", "Whether agents with this routing profile will have their routing order calculated based on *longest idle time* or *time since their last inbound contact*.")
		connect_createRoutingProfileCmd.Flags().String("default-outbound-queue-id", "", "The default outbound queue for the routing profile.")
		connect_createRoutingProfileCmd.Flags().String("description", "", "Description of the routing profile.")
		connect_createRoutingProfileCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_createRoutingProfileCmd.Flags().String("manual-assignment-queue-configs", "", "The manual assignment queues associated with the routing profile.")
		connect_createRoutingProfileCmd.Flags().String("media-concurrencies", "", "The channels that agents can handle in the Contact Control Panel (CCP) for this routing profile.")
		connect_createRoutingProfileCmd.Flags().String("name", "", "The name of the routing profile.")
		connect_createRoutingProfileCmd.Flags().String("queue-configs", "", "The inbound queues associated with the routing profile.")
		connect_createRoutingProfileCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		connect_createRoutingProfileCmd.MarkFlagRequired("default-outbound-queue-id")
		connect_createRoutingProfileCmd.MarkFlagRequired("description")
		connect_createRoutingProfileCmd.MarkFlagRequired("instance-id")
		connect_createRoutingProfileCmd.MarkFlagRequired("media-concurrencies")
		connect_createRoutingProfileCmd.MarkFlagRequired("name")
	})
	connectCmd.AddCommand(connect_createRoutingProfileCmd)
}
