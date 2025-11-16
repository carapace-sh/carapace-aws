package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_listAllowedNodeTypeUpdatesCmd = &cobra.Command{
	Use:   "list-allowed-node-type-updates",
	Short: "Lists all available node types that you can scale to from your cluster's current node type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_listAllowedNodeTypeUpdatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_listAllowedNodeTypeUpdatesCmd).Standalone()

		memorydb_listAllowedNodeTypeUpdatesCmd.Flags().String("cluster-name", "", "The name of the cluster you want to scale.")
		memorydb_listAllowedNodeTypeUpdatesCmd.MarkFlagRequired("cluster-name")
	})
	memorydbCmd.AddCommand(memorydb_listAllowedNodeTypeUpdatesCmd)
}
