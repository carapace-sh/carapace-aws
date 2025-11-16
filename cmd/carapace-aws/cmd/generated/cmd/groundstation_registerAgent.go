package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_registerAgentCmd = &cobra.Command{
	Use:   "register-agent",
	Short: "For use by AWS Ground Station Agent and shouldn't be called directly.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_registerAgentCmd).Standalone()

	groundstation_registerAgentCmd.Flags().String("agent-details", "", "Detailed information about the agent being registered.")
	groundstation_registerAgentCmd.Flags().String("discovery-data", "", "Data for associating an agent with the capabilities it is managing.")
	groundstation_registerAgentCmd.Flags().String("tags", "", "Tags assigned to an `Agent`.")
	groundstation_registerAgentCmd.MarkFlagRequired("agent-details")
	groundstation_registerAgentCmd.MarkFlagRequired("discovery-data")
	groundstationCmd.AddCommand(groundstation_registerAgentCmd)
}
