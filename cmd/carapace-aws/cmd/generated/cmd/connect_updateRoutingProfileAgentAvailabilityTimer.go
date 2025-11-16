package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateRoutingProfileAgentAvailabilityTimerCmd = &cobra.Command{
	Use:   "update-routing-profile-agent-availability-timer",
	Short: "Whether agents with this routing profile will have their routing order calculated based on *time since their last inbound contact* or *longest idle time*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateRoutingProfileAgentAvailabilityTimerCmd).Standalone()

	connect_updateRoutingProfileAgentAvailabilityTimerCmd.Flags().String("agent-availability-timer", "", "Whether agents with this routing profile will have their routing order calculated based on *time since their last inbound contact* or *longest idle time*.")
	connect_updateRoutingProfileAgentAvailabilityTimerCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateRoutingProfileAgentAvailabilityTimerCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile.")
	connect_updateRoutingProfileAgentAvailabilityTimerCmd.MarkFlagRequired("agent-availability-timer")
	connect_updateRoutingProfileAgentAvailabilityTimerCmd.MarkFlagRequired("instance-id")
	connect_updateRoutingProfileAgentAvailabilityTimerCmd.MarkFlagRequired("routing-profile-id")
	connectCmd.AddCommand(connect_updateRoutingProfileAgentAvailabilityTimerCmd)
}
