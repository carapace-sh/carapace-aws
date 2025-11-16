package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_deleteCacheClusterCmd = &cobra.Command{
	Use:   "delete-cache-cluster",
	Short: "Deletes a previously provisioned cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_deleteCacheClusterCmd).Standalone()

	elasticache_deleteCacheClusterCmd.Flags().String("cache-cluster-id", "", "The cluster identifier for the cluster to be deleted.")
	elasticache_deleteCacheClusterCmd.Flags().String("final-snapshot-identifier", "", "The user-supplied name of a final cluster snapshot.")
	elasticache_deleteCacheClusterCmd.MarkFlagRequired("cache-cluster-id")
	elasticacheCmd.AddCommand(elasticache_deleteCacheClusterCmd)
}
