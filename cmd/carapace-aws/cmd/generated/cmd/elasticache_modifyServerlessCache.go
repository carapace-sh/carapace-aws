package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_modifyServerlessCacheCmd = &cobra.Command{
	Use:   "modify-serverless-cache",
	Short: "This API modifies the attributes of a serverless cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_modifyServerlessCacheCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_modifyServerlessCacheCmd).Standalone()

		elasticache_modifyServerlessCacheCmd.Flags().String("cache-usage-limits", "", "Modify the cache usage limit for the serverless cache.")
		elasticache_modifyServerlessCacheCmd.Flags().String("daily-snapshot-time", "", "The daily time during which Elasticache begins taking a daily snapshot of the serverless cache.")
		elasticache_modifyServerlessCacheCmd.Flags().String("description", "", "User provided description for the serverless cache.")
		elasticache_modifyServerlessCacheCmd.Flags().String("engine", "", "Modifies the engine listed in a serverless cache request.")
		elasticache_modifyServerlessCacheCmd.Flags().String("major-engine-version", "", "Modifies the engine vesion listed in a serverless cache request.")
		elasticache_modifyServerlessCacheCmd.Flags().String("remove-user-group", "", "The identifier of the UserGroup to be removed from association with the Valkey and Redis OSS serverless cache.")
		elasticache_modifyServerlessCacheCmd.Flags().String("security-group-ids", "", "The new list of VPC security groups to be associated with the serverless cache.")
		elasticache_modifyServerlessCacheCmd.Flags().String("serverless-cache-name", "", "User-provided identifier for the serverless cache to be modified.")
		elasticache_modifyServerlessCacheCmd.Flags().String("snapshot-retention-limit", "", "The number of days for which Elasticache retains automatic snapshots before deleting them.")
		elasticache_modifyServerlessCacheCmd.Flags().String("user-group-id", "", "The identifier of the UserGroup to be associated with the serverless cache.")
		elasticache_modifyServerlessCacheCmd.MarkFlagRequired("serverless-cache-name")
	})
	elasticacheCmd.AddCommand(elasticache_modifyServerlessCacheCmd)
}
