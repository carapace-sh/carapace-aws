package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_batchDeleteAgentsCmd = &cobra.Command{
	Use:   "batch-delete-agents",
	Short: "Deletes one or more agents or collectors as specified by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_batchDeleteAgentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_batchDeleteAgentsCmd).Standalone()

		discovery_batchDeleteAgentsCmd.Flags().String("delete-agents", "", "The list of agents to delete.")
		discovery_batchDeleteAgentsCmd.MarkFlagRequired("delete-agents")
	})
	discoveryCmd.AddCommand(discovery_batchDeleteAgentsCmd)
}
