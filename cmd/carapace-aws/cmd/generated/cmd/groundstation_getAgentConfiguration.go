package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstation_getAgentConfigurationCmd = &cobra.Command{
	Use:   "get-agent-configuration",
	Short: "For use by AWS Ground Station Agent and shouldn't be called directly.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstation_getAgentConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(groundstation_getAgentConfigurationCmd).Standalone()

		groundstation_getAgentConfigurationCmd.Flags().String("agent-id", "", "UUID of agent to get configuration information for.")
		groundstation_getAgentConfigurationCmd.MarkFlagRequired("agent-id")
	})
	groundstationCmd.AddCommand(groundstation_getAgentConfigurationCmd)
}
