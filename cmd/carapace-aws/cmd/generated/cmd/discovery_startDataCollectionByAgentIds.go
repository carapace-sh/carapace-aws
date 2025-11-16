package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_startDataCollectionByAgentIdsCmd = &cobra.Command{
	Use:   "start-data-collection-by-agent-ids",
	Short: "Instructs the specified agents to start collecting data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_startDataCollectionByAgentIdsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_startDataCollectionByAgentIdsCmd).Standalone()

		discovery_startDataCollectionByAgentIdsCmd.Flags().String("agent-ids", "", "The IDs of the agents from which to start collecting data.")
		discovery_startDataCollectionByAgentIdsCmd.MarkFlagRequired("agent-ids")
	})
	discoveryCmd.AddCommand(discovery_startDataCollectionByAgentIdsCmd)
}
