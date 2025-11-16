package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_modifyGlobalReplicationGroupCmd = &cobra.Command{
	Use:   "modify-global-replication-group",
	Short: "Modifies the settings for a Global datastore.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_modifyGlobalReplicationGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_modifyGlobalReplicationGroupCmd).Standalone()

		elasticache_modifyGlobalReplicationGroupCmd.Flags().Bool("apply-immediately", false, "This parameter causes the modifications in this request and any pending modifications to be applied, asynchronously and as soon as possible.")
		elasticache_modifyGlobalReplicationGroupCmd.Flags().String("automatic-failover-enabled", "", "Determines whether a read replica is automatically promoted to read/write primary if the existing primary encounters a failure.")
		elasticache_modifyGlobalReplicationGroupCmd.Flags().String("cache-node-type", "", "A valid cache node type that you want to scale this Global datastore to.")
		elasticache_modifyGlobalReplicationGroupCmd.Flags().String("cache-parameter-group-name", "", "The name of the cache parameter group to use with the Global datastore.")
		elasticache_modifyGlobalReplicationGroupCmd.Flags().String("engine", "", "Modifies the engine listed in a global replication group message.")
		elasticache_modifyGlobalReplicationGroupCmd.Flags().String("engine-version", "", "The upgraded version of the cache engine to be run on the clusters in the Global datastore.")
		elasticache_modifyGlobalReplicationGroupCmd.Flags().String("global-replication-group-description", "", "A description of the Global datastore")
		elasticache_modifyGlobalReplicationGroupCmd.Flags().String("global-replication-group-id", "", "The name of the Global datastore")
		elasticache_modifyGlobalReplicationGroupCmd.Flags().Bool("no-apply-immediately", false, "This parameter causes the modifications in this request and any pending modifications to be applied, asynchronously and as soon as possible.")
		elasticache_modifyGlobalReplicationGroupCmd.MarkFlagRequired("apply-immediately")
		elasticache_modifyGlobalReplicationGroupCmd.MarkFlagRequired("global-replication-group-id")
		elasticache_modifyGlobalReplicationGroupCmd.Flag("no-apply-immediately").Hidden = true
		elasticache_modifyGlobalReplicationGroupCmd.MarkFlagRequired("no-apply-immediately")
	})
	elasticacheCmd.AddCommand(elasticache_modifyGlobalReplicationGroupCmd)
}
