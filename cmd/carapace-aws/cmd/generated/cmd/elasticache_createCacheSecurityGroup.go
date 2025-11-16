package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_createCacheSecurityGroupCmd = &cobra.Command{
	Use:   "create-cache-security-group",
	Short: "Creates a new cache security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_createCacheSecurityGroupCmd).Standalone()

	elasticache_createCacheSecurityGroupCmd.Flags().String("cache-security-group-name", "", "A name for the cache security group.")
	elasticache_createCacheSecurityGroupCmd.Flags().String("description", "", "A description for the cache security group.")
	elasticache_createCacheSecurityGroupCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	elasticache_createCacheSecurityGroupCmd.MarkFlagRequired("cache-security-group-name")
	elasticache_createCacheSecurityGroupCmd.MarkFlagRequired("description")
	elasticacheCmd.AddCommand(elasticache_createCacheSecurityGroupCmd)
}
