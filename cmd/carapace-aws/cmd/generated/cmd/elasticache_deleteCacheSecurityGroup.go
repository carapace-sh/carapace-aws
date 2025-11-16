package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_deleteCacheSecurityGroupCmd = &cobra.Command{
	Use:   "delete-cache-security-group",
	Short: "Deletes a cache security group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_deleteCacheSecurityGroupCmd).Standalone()

	elasticache_deleteCacheSecurityGroupCmd.Flags().String("cache-security-group-name", "", "The name of the cache security group to delete.")
	elasticache_deleteCacheSecurityGroupCmd.MarkFlagRequired("cache-security-group-name")
	elasticacheCmd.AddCommand(elasticache_deleteCacheSecurityGroupCmd)
}
