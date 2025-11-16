package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_getAgentTaskResponseUrlCmd = &cobra.Command{
	Use:   "get-agent-task-response-url",
	Short: "For use by AWS Ground Station Agent and shouldn't be called directly.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_getAgentTaskResponseUrlCmd).Standalone()

	groundstation_getAgentTaskResponseUrlCmd.Flags().String("agent-id", "", "UUID of agent requesting the response URL.")
	groundstation_getAgentTaskResponseUrlCmd.Flags().String("task-id", "", "GUID of the agent task for which the response URL is being requested.")
	groundstation_getAgentTaskResponseUrlCmd.MarkFlagRequired("agent-id")
	groundstation_getAgentTaskResponseUrlCmd.MarkFlagRequired("task-id")
	groundstationCmd.AddCommand(groundstation_getAgentTaskResponseUrlCmd)
}
