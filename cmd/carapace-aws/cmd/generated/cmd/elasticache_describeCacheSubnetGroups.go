package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeCacheSubnetGroupsCmd = &cobra.Command{
	Use:   "describe-cache-subnet-groups",
	Short: "Returns a list of cache subnet group descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeCacheSubnetGroupsCmd).Standalone()

	elasticache_describeCacheSubnetGroupsCmd.Flags().String("cache-subnet-group-name", "", "The name of the cache subnet group to return details for.")
	elasticache_describeCacheSubnetGroupsCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
	elasticache_describeCacheSubnetGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	elasticacheCmd.AddCommand(elasticache_describeCacheSubnetGroupsCmd)
}
