package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeCacheParameterGroupsCmd = &cobra.Command{
	Use:   "describe-cache-parameter-groups",
	Short: "Returns a list of cache parameter group descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeCacheParameterGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_describeCacheParameterGroupsCmd).Standalone()

		elasticache_describeCacheParameterGroupsCmd.Flags().String("cache-parameter-group-name", "", "The name of a specific cache parameter group to return details for.")
		elasticache_describeCacheParameterGroupsCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
		elasticache_describeCacheParameterGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	elasticacheCmd.AddCommand(elasticache_describeCacheParameterGroupsCmd)
}
