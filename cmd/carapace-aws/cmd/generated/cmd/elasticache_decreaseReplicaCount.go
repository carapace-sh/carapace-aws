package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_decreaseReplicaCountCmd = &cobra.Command{
	Use:   "decrease-replica-count",
	Short: "Dynamically decreases the number of replicas in a Valkey or Redis OSS (cluster mode disabled) replication group or the number of replica nodes in one or more node groups (shards) of a Valkey or Redis OSS (cluster mode enabled) replication group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_decreaseReplicaCountCmd).Standalone()

	elasticache_decreaseReplicaCountCmd.Flags().Bool("apply-immediately", false, "If `True`, the number of replica nodes is decreased immediately.")
	elasticache_decreaseReplicaCountCmd.Flags().String("new-replica-count", "", "The number of read replica nodes you want at the completion of this operation.")
	elasticache_decreaseReplicaCountCmd.Flags().Bool("no-apply-immediately", false, "If `True`, the number of replica nodes is decreased immediately.")
	elasticache_decreaseReplicaCountCmd.Flags().String("replica-configuration", "", "A list of `ConfigureShard` objects that can be used to configure each shard in a Valkey or Redis OSS replication group.")
	elasticache_decreaseReplicaCountCmd.Flags().String("replicas-to-remove", "", "A list of the node ids to remove from the replication group or node group (shard).")
	elasticache_decreaseReplicaCountCmd.Flags().String("replication-group-id", "", "The id of the replication group from which you want to remove replica nodes.")
	elasticache_decreaseReplicaCountCmd.MarkFlagRequired("apply-immediately")
	elasticache_decreaseReplicaCountCmd.Flag("no-apply-immediately").Hidden = true
	elasticache_decreaseReplicaCountCmd.MarkFlagRequired("no-apply-immediately")
	elasticache_decreaseReplicaCountCmd.MarkFlagRequired("replication-group-id")
	elasticacheCmd.AddCommand(elasticache_decreaseReplicaCountCmd)
}
