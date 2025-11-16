package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd = &cobra.Command{
	Use:   "decrease-node-groups-in-global-replication-group",
	Short: "Decreases the number of node groups in a Global datastore",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd).Standalone()

		elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd.Flags().Bool("apply-immediately", false, "Indicates that the shard reconfiguration process begins immediately.")
		elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd.Flags().String("global-node-groups-to-remove", "", "If the value of NodeGroupCount is less than the current number of node groups (shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.")
		elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd.Flags().String("global-node-groups-to-retain", "", "If the value of NodeGroupCount is less than the current number of node groups (shards), then either NodeGroupsToRemove or NodeGroupsToRetain is required.")
		elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd.Flags().String("global-replication-group-id", "", "The name of the Global datastore")
		elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd.Flags().Bool("no-apply-immediately", false, "Indicates that the shard reconfiguration process begins immediately.")
		elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd.Flags().String("node-group-count", "", "The number of node groups (shards) that results from the modification of the shard configuration")
		elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd.MarkFlagRequired("apply-immediately")
		elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd.MarkFlagRequired("global-replication-group-id")
		elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd.Flag("no-apply-immediately").Hidden = true
		elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd.MarkFlagRequired("no-apply-immediately")
		elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd.MarkFlagRequired("node-group-count")
	})
	elasticacheCmd.AddCommand(elasticache_decreaseNodeGroupsInGlobalReplicationGroupCmd)
}
