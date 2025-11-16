package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_listAllowedNodeTypeModificationsCmd = &cobra.Command{
	Use:   "list-allowed-node-type-modifications",
	Short: "Lists all available node types that you can scale with your cluster's replication group's current node type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_listAllowedNodeTypeModificationsCmd).Standalone()

	elasticache_listAllowedNodeTypeModificationsCmd.Flags().String("cache-cluster-id", "", "The name of the cluster you want to scale up to a larger node instanced type.")
	elasticache_listAllowedNodeTypeModificationsCmd.Flags().String("replication-group-id", "", "The name of the replication group want to scale up to a larger node type.")
	elasticacheCmd.AddCommand(elasticache_listAllowedNodeTypeModificationsCmd)
}
