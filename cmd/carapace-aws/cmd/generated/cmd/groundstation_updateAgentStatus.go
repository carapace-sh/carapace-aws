package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_updateAgentStatusCmd = &cobra.Command{
	Use:   "update-agent-status",
	Short: "For use by AWS Ground Station Agent and shouldn't be called directly.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_updateAgentStatusCmd).Standalone()

	groundstation_updateAgentStatusCmd.Flags().String("agent-id", "", "UUID of agent to update.")
	groundstation_updateAgentStatusCmd.Flags().String("aggregate-status", "", "Aggregate status for agent.")
	groundstation_updateAgentStatusCmd.Flags().String("component-statuses", "", "List of component statuses for agent.")
	groundstation_updateAgentStatusCmd.Flags().String("task-id", "", "GUID of agent task.")
	groundstation_updateAgentStatusCmd.MarkFlagRequired("agent-id")
	groundstation_updateAgentStatusCmd.MarkFlagRequired("aggregate-status")
	groundstation_updateAgentStatusCmd.MarkFlagRequired("component-statuses")
	groundstation_updateAgentStatusCmd.MarkFlagRequired("task-id")
	groundstationCmd.AddCommand(groundstation_updateAgentStatusCmd)
}
