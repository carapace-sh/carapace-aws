package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_increaseReplicaCountCmd = &cobra.Command{
	Use:   "increase-replica-count",
	Short: "Dynamically increases the number of replicas in a Valkey or Redis OSS (cluster mode disabled) replication group or the number of replica nodes in one or more node groups (shards) of a Valkey or Redis OSS (cluster mode enabled) replication group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_increaseReplicaCountCmd).Standalone()

	elasticache_increaseReplicaCountCmd.Flags().Bool("apply-immediately", false, "If `True`, the number of replica nodes is increased immediately.")
	elasticache_increaseReplicaCountCmd.Flags().String("new-replica-count", "", "The number of read replica nodes you want at the completion of this operation.")
	elasticache_increaseReplicaCountCmd.Flags().Bool("no-apply-immediately", false, "If `True`, the number of replica nodes is increased immediately.")
	elasticache_increaseReplicaCountCmd.Flags().String("replica-configuration", "", "A list of `ConfigureShard` objects that can be used to configure each shard in a Valkey or Redis OSS (cluster mode enabled) replication group.")
	elasticache_increaseReplicaCountCmd.Flags().String("replication-group-id", "", "The id of the replication group to which you want to add replica nodes.")
	elasticache_increaseReplicaCountCmd.MarkFlagRequired("apply-immediately")
	elasticache_increaseReplicaCountCmd.Flag("no-apply-immediately").Hidden = true
	elasticache_increaseReplicaCountCmd.MarkFlagRequired("no-apply-immediately")
	elasticache_increaseReplicaCountCmd.MarkFlagRequired("replication-group-id")
	elasticacheCmd.AddCommand(elasticache_increaseReplicaCountCmd)
}
