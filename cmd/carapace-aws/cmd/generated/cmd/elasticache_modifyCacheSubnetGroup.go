package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_modifyCacheSubnetGroupCmd = &cobra.Command{
	Use:   "modify-cache-subnet-group",
	Short: "Modifies an existing cache subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_modifyCacheSubnetGroupCmd).Standalone()

	elasticache_modifyCacheSubnetGroupCmd.Flags().String("cache-subnet-group-description", "", "A description of the cache subnet group.")
	elasticache_modifyCacheSubnetGroupCmd.Flags().String("cache-subnet-group-name", "", "The name for the cache subnet group.")
	elasticache_modifyCacheSubnetGroupCmd.Flags().String("subnet-ids", "", "The EC2 subnet IDs for the cache subnet group.")
	elasticache_modifyCacheSubnetGroupCmd.MarkFlagRequired("cache-subnet-group-name")
	elasticacheCmd.AddCommand(elasticache_modifyCacheSubnetGroupCmd)
}
