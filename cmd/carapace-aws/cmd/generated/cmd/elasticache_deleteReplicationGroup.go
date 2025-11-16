package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_deleteReplicationGroupCmd = &cobra.Command{
	Use:   "delete-replication-group",
	Short: "Deletes an existing replication group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_deleteReplicationGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_deleteReplicationGroupCmd).Standalone()

		elasticache_deleteReplicationGroupCmd.Flags().String("final-snapshot-identifier", "", "The name of a final node group (shard) snapshot.")
		elasticache_deleteReplicationGroupCmd.Flags().String("replication-group-id", "", "The identifier for the cluster to be deleted.")
		elasticache_deleteReplicationGroupCmd.Flags().String("retain-primary-cluster", "", "If set to `true`, all of the read replicas are deleted, but the primary node is retained.")
		elasticache_deleteReplicationGroupCmd.MarkFlagRequired("replication-group-id")
	})
	elasticacheCmd.AddCommand(elasticache_deleteReplicationGroupCmd)
}
