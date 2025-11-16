package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_stopDataCollectionByAgentIdsCmd = &cobra.Command{
	Use:   "stop-data-collection-by-agent-ids",
	Short: "Instructs the specified agents to stop collecting data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_stopDataCollectionByAgentIdsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_stopDataCollectionByAgentIdsCmd).Standalone()

		discovery_stopDataCollectionByAgentIdsCmd.Flags().String("agent-ids", "", "The IDs of the agents from which to stop collecting data.")
		discovery_stopDataCollectionByAgentIdsCmd.MarkFlagRequired("agent-ids")
	})
	discoveryCmd.AddCommand(discovery_stopDataCollectionByAgentIdsCmd)
}
