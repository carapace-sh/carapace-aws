package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_createGlobalReplicationGroupCmd = &cobra.Command{
	Use:   "create-global-replication-group",
	Short: "Global Datastore offers fully managed, fast, reliable and secure cross-region replication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_createGlobalReplicationGroupCmd).Standalone()

	elasticache_createGlobalReplicationGroupCmd.Flags().String("global-replication-group-description", "", "Provides details of the Global datastore")
	elasticache_createGlobalReplicationGroupCmd.Flags().String("global-replication-group-id-suffix", "", "The suffix name of a Global datastore.")
	elasticache_createGlobalReplicationGroupCmd.Flags().String("primary-replication-group-id", "", "The name of the primary cluster that accepts writes and will replicate updates to the secondary cluster.")
	elasticache_createGlobalReplicationGroupCmd.MarkFlagRequired("global-replication-group-id-suffix")
	elasticache_createGlobalReplicationGroupCmd.MarkFlagRequired("primary-replication-group-id")
	elasticacheCmd.AddCommand(elasticache_createGlobalReplicationGroupCmd)
}
