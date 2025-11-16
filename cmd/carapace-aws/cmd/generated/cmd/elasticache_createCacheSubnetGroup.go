package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_createCacheSubnetGroupCmd = &cobra.Command{
	Use:   "create-cache-subnet-group",
	Short: "Creates a new cache subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_createCacheSubnetGroupCmd).Standalone()

	elasticache_createCacheSubnetGroupCmd.Flags().String("cache-subnet-group-description", "", "A description for the cache subnet group.")
	elasticache_createCacheSubnetGroupCmd.Flags().String("cache-subnet-group-name", "", "A name for the cache subnet group.")
	elasticache_createCacheSubnetGroupCmd.Flags().String("subnet-ids", "", "A list of VPC subnet IDs for the cache subnet group.")
	elasticache_createCacheSubnetGroupCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	elasticache_createCacheSubnetGroupCmd.MarkFlagRequired("cache-subnet-group-description")
	elasticache_createCacheSubnetGroupCmd.MarkFlagRequired("cache-subnet-group-name")
	elasticache_createCacheSubnetGroupCmd.MarkFlagRequired("subnet-ids")
	elasticacheCmd.AddCommand(elasticache_createCacheSubnetGroupCmd)
}
