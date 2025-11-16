package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_createCacheParameterGroupCmd = &cobra.Command{
	Use:   "create-cache-parameter-group",
	Short: "Creates a new Amazon ElastiCache cache parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_createCacheParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_createCacheParameterGroupCmd).Standalone()

		elasticache_createCacheParameterGroupCmd.Flags().String("cache-parameter-group-family", "", "The name of the cache parameter group family that the cache parameter group can be used with.")
		elasticache_createCacheParameterGroupCmd.Flags().String("cache-parameter-group-name", "", "A user-specified name for the cache parameter group.")
		elasticache_createCacheParameterGroupCmd.Flags().String("description", "", "A user-specified description for the cache parameter group.")
		elasticache_createCacheParameterGroupCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
		elasticache_createCacheParameterGroupCmd.MarkFlagRequired("cache-parameter-group-family")
		elasticache_createCacheParameterGroupCmd.MarkFlagRequired("cache-parameter-group-name")
		elasticache_createCacheParameterGroupCmd.MarkFlagRequired("description")
	})
	elasticacheCmd.AddCommand(elasticache_createCacheParameterGroupCmd)
}
