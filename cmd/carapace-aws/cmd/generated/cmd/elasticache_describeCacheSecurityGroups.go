package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeCacheSecurityGroupsCmd = &cobra.Command{
	Use:   "describe-cache-security-groups",
	Short: "Returns a list of cache security group descriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeCacheSecurityGroupsCmd).Standalone()

	elasticache_describeCacheSecurityGroupsCmd.Flags().String("cache-security-group-name", "", "The name of the cache security group to return details for.")
	elasticache_describeCacheSecurityGroupsCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
	elasticache_describeCacheSecurityGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	elasticacheCmd.AddCommand(elasticache_describeCacheSecurityGroupsCmd)
}
