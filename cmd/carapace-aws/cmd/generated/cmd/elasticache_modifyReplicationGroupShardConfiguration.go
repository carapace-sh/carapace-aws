package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_modifyReplicationGroupShardConfigurationCmd = &cobra.Command{
	Use:   "modify-replication-group-shard-configuration",
	Short: "Modifies a replication group's shards (node groups) by allowing you to add shards, remove shards, or rebalance the keyspaces among existing shards.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_modifyReplicationGroupShardConfigurationCmd).Standalone()

	elasticache_modifyReplicationGroupShardConfigurationCmd.Flags().Bool("apply-immediately", false, "Indicates that the shard reconfiguration process begins immediately.")
	elasticache_modifyReplicationGroupShardConfigurationCmd.Flags().Bool("no-apply-immediately", false, "Indicates that the shard reconfiguration process begins immediately.")
	elasticache_modifyReplicationGroupShardConfigurationCmd.Flags().String("node-group-count", "", "The number of node groups (shards) that results from the modification of the shard configuration.")
	elasticache_modifyReplicationGroupShardConfigurationCmd.Flags().String("node-groups-to-remove", "", "If the value of `NodeGroupCount` is less than the current number of node groups (shards), then either `NodeGroupsToRemove` or `NodeGroupsToRetain` is required.")
	elasticache_modifyReplicationGroupShardConfigurationCmd.Flags().String("node-groups-to-retain", "", "If the value of `NodeGroupCount` is less than the current number of node groups (shards), then either `NodeGroupsToRemove` or `NodeGroupsToRetain` is required.")
	elasticache_modifyReplicationGroupShardConfigurationCmd.Flags().String("replication-group-id", "", "The name of the Valkey or Redis OSS (cluster mode enabled) cluster (replication group) on which the shards are to be configured.")
	elasticache_modifyReplicationGroupShardConfigurationCmd.Flags().String("resharding-configuration", "", "Specifies the preferred availability zones for each node group in the cluster.")
	elasticache_modifyReplicationGroupShardConfigurationCmd.MarkFlagRequired("apply-immediately")
	elasticache_modifyReplicationGroupShardConfigurationCmd.Flag("no-apply-immediately").Hidden = true
	elasticache_modifyReplicationGroupShardConfigurationCmd.MarkFlagRequired("no-apply-immediately")
	elasticache_modifyReplicationGroupShardConfigurationCmd.MarkFlagRequired("node-group-count")
	elasticache_modifyReplicationGroupShardConfigurationCmd.MarkFlagRequired("replication-group-id")
	elasticacheCmd.AddCommand(elasticache_modifyReplicationGroupShardConfigurationCmd)
}
