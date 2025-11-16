package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_deleteCacheSubnetGroupCmd = &cobra.Command{
	Use:   "delete-cache-subnet-group",
	Short: "Deletes a cache subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_deleteCacheSubnetGroupCmd).Standalone()

	elasticache_deleteCacheSubnetGroupCmd.Flags().String("cache-subnet-group-name", "", "The name of the cache subnet group to delete.")
	elasticache_deleteCacheSubnetGroupCmd.MarkFlagRequired("cache-subnet-group-name")
	elasticacheCmd.AddCommand(elasticache_deleteCacheSubnetGroupCmd)
}
