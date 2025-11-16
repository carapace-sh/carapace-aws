package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_createServerlessCacheCmd = &cobra.Command{
	Use:   "create-serverless-cache",
	Short: "Creates a serverless cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_createServerlessCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_createServerlessCacheCmd).Standalone()

		elasticache_createServerlessCacheCmd.Flags().String("cache-usage-limits", "", "Sets the cache usage limits for storage and ElastiCache Processing Units for the cache.")
		elasticache_createServerlessCacheCmd.Flags().String("daily-snapshot-time", "", "The daily time that snapshots will be created from the new serverless cache.")
		elasticache_createServerlessCacheCmd.Flags().String("description", "", "User-provided description for the serverless cache.")
		elasticache_createServerlessCacheCmd.Flags().String("engine", "", "The name of the cache engine to be used for creating the serverless cache.")
		elasticache_createServerlessCacheCmd.Flags().String("kms-key-id", "", "ARN of the customer managed key for encrypting the data at rest.")
		elasticache_createServerlessCacheCmd.Flags().String("major-engine-version", "", "The version of the cache engine that will be used to create the serverless cache.")
		elasticache_createServerlessCacheCmd.Flags().String("security-group-ids", "", "A list of the one or more VPC security groups to be associated with the serverless cache.")
		elasticache_createServerlessCacheCmd.Flags().String("serverless-cache-name", "", "User-provided identifier for the serverless cache.")
		elasticache_createServerlessCacheCmd.Flags().String("snapshot-arns-to-restore", "", "The ARN(s) of the snapshot that the new serverless cache will be created from.")
		elasticache_createServerlessCacheCmd.Flags().String("snapshot-retention-limit", "", "The number of snapshots that will be retained for the serverless cache that is being created.")
		elasticache_createServerlessCacheCmd.Flags().String("subnet-ids", "", "A list of the identifiers of the subnets where the VPC endpoint for the serverless cache will be deployed.")
		elasticache_createServerlessCacheCmd.Flags().String("tags", "", "The list of tags (key, value) pairs to be added to the serverless cache resource.")
		elasticache_createServerlessCacheCmd.Flags().String("user-group-id", "", "The identifier of the UserGroup to be associated with the serverless cache.")
		elasticache_createServerlessCacheCmd.MarkFlagRequired("engine")
		elasticache_createServerlessCacheCmd.MarkFlagRequired("serverless-cache-name")
	})
	elasticacheCmd.AddCommand(elasticache_createServerlessCacheCmd)
}
