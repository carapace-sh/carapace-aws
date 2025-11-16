package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd = &cobra.Command{
	Use:   "increase-node-groups-in-global-replication-group",
	Short: "Increase the number of node groups in the Global datastore",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd).Standalone()

		elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd.Flags().Bool("apply-immediately", false, "Indicates that the process begins immediately.")
		elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd.Flags().String("global-replication-group-id", "", "The name of the Global datastore")
		elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd.Flags().Bool("no-apply-immediately", false, "Indicates that the process begins immediately.")
		elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd.Flags().String("node-group-count", "", "Total number of node groups you want")
		elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd.Flags().String("regional-configurations", "", "Describes the replication group IDs, the Amazon regions where they are stored and the shard configuration for each that comprise the Global datastore")
		elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd.MarkFlagRequired("apply-immediately")
		elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd.MarkFlagRequired("global-replication-group-id")
		elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd.Flag("no-apply-immediately").Hidden = true
		elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd.MarkFlagRequired("no-apply-immediately")
		elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd.MarkFlagRequired("node-group-count")
	})
	elasticacheCmd.AddCommand(elasticache_increaseNodeGroupsInGlobalReplicationGroupCmd)
}
