package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_rebalanceSlotsInGlobalReplicationGroupCmd = &cobra.Command{
	Use:   "rebalance-slots-in-global-replication-group",
	Short: "Redistribute slots to ensure uniform distribution across existing shards in the cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_rebalanceSlotsInGlobalReplicationGroupCmd).Standalone()

	elasticache_rebalanceSlotsInGlobalReplicationGroupCmd.Flags().Bool("apply-immediately", false, "If `True`, redistribution is applied immediately.")
	elasticache_rebalanceSlotsInGlobalReplicationGroupCmd.Flags().String("global-replication-group-id", "", "The name of the Global datastore")
	elasticache_rebalanceSlotsInGlobalReplicationGroupCmd.Flags().Bool("no-apply-immediately", false, "If `True`, redistribution is applied immediately.")
	elasticache_rebalanceSlotsInGlobalReplicationGroupCmd.MarkFlagRequired("apply-immediately")
	elasticache_rebalanceSlotsInGlobalReplicationGroupCmd.MarkFlagRequired("global-replication-group-id")
	elasticache_rebalanceSlotsInGlobalReplicationGroupCmd.Flag("no-apply-immediately").Hidden = true
	elasticache_rebalanceSlotsInGlobalReplicationGroupCmd.MarkFlagRequired("no-apply-immediately")
	elasticacheCmd.AddCommand(elasticache_rebalanceSlotsInGlobalReplicationGroupCmd)
}
