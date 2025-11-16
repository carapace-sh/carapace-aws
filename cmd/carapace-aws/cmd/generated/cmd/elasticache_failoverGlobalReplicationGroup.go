package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_failoverGlobalReplicationGroupCmd = &cobra.Command{
	Use:   "failover-global-replication-group",
	Short: "Used to failover the primary region to a secondary region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_failoverGlobalReplicationGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_failoverGlobalReplicationGroupCmd).Standalone()

		elasticache_failoverGlobalReplicationGroupCmd.Flags().String("global-replication-group-id", "", "The name of the Global datastore")
		elasticache_failoverGlobalReplicationGroupCmd.Flags().String("primary-region", "", "The Amazon region of the primary cluster of the Global datastore")
		elasticache_failoverGlobalReplicationGroupCmd.Flags().String("primary-replication-group-id", "", "The name of the primary replication group")
		elasticache_failoverGlobalReplicationGroupCmd.MarkFlagRequired("global-replication-group-id")
		elasticache_failoverGlobalReplicationGroupCmd.MarkFlagRequired("primary-region")
		elasticache_failoverGlobalReplicationGroupCmd.MarkFlagRequired("primary-replication-group-id")
	})
	elasticacheCmd.AddCommand(elasticache_failoverGlobalReplicationGroupCmd)
}
