package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_rebootCacheClusterCmd = &cobra.Command{
	Use:   "reboot-cache-cluster",
	Short: "Reboots some, or all, of the cache nodes within a provisioned cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_rebootCacheClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_rebootCacheClusterCmd).Standalone()

		elasticache_rebootCacheClusterCmd.Flags().String("cache-cluster-id", "", "The cluster identifier.")
		elasticache_rebootCacheClusterCmd.Flags().String("cache-node-ids-to-reboot", "", "A list of cache node IDs to reboot.")
		elasticache_rebootCacheClusterCmd.MarkFlagRequired("cache-cluster-id")
		elasticache_rebootCacheClusterCmd.MarkFlagRequired("cache-node-ids-to-reboot")
	})
	elasticacheCmd.AddCommand(elasticache_rebootCacheClusterCmd)
}
