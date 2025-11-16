package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_deleteCacheParameterGroupCmd = &cobra.Command{
	Use:   "delete-cache-parameter-group",
	Short: "Deletes the specified cache parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_deleteCacheParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_deleteCacheParameterGroupCmd).Standalone()

		elasticache_deleteCacheParameterGroupCmd.Flags().String("cache-parameter-group-name", "", "The name of the cache parameter group to delete.")
		elasticache_deleteCacheParameterGroupCmd.MarkFlagRequired("cache-parameter-group-name")
	})
	elasticacheCmd.AddCommand(elasticache_deleteCacheParameterGroupCmd)
}
